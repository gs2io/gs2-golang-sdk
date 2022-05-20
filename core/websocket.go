/**
 * Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
 * Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License").
 * You may not use this file except in compliance with the License.
 * A copy of the License is located at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * or in the "license" file accompanying this file. This file is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
 * express or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package core

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type WebSocketRequestId string
type WebSocketBodies map[string]interface{}

type WebSocketNetworkJob struct {
	RequestId WebSocketRequestId
	Bodies    WebSocketBodies
	Callback  chan<- AsyncResult
}

type WebSocketMessageContainerBody string
type WebSocketMessageContainerMessage string
type WebSocketMessageContainerType string
type WebSocketMessageContainerStatus int

type WebSocketMessageContainer struct {
	RequestId WebSocketRequestId               `json:"requestId"`
	Message   WebSocketMessageContainerMessage `json:"message"`
	Body      map[string]interface{}           `json:"body"`
	Type      WebSocketMessageContainerType    `json:"type"`
	Status    WebSocketMessageContainerStatus  `json:"status"`
}

type Gs2WebSocketSession struct {
	Credential   IGs2Credential
	Region       Region
	projectToken ProjectToken
	connection   *WebsocketConnection
	Jobs         []*WebSocketNetworkJob

	notificationHandler []func(message Notification)
}

func (p Gs2WebSocketSession) EndpointHost(service string) Url {
	return Url(strings.ReplaceAll(strings.ReplaceAll(EndpointHost, "{service}", service), "{region}", string(p.Region)))
}

func (p Gs2WebSocketSession) CreateAuthorizationHeader() map[string]string {
	return map[string]string{
		"xGs2ClientId":     string(*p.Credential.GetClientId()),
		"xGs2ProjectToken": fmt.Sprintf("%s", p.projectToken),
	}
}

func (p *Gs2WebSocketSession) AddNotificationHandler(handler func(message Notification)) {
	p.notificationHandler = append(p.notificationHandler, handler)
}

func readWebSocketErrors(response map[string]interface{}) ([]RequestError, error) {
	byteArray, err := json.Marshal(response)
	if err != nil {
		return []RequestError{}, err
	}
	var result ErrorResult
	err = json.Unmarshal(byteArray, &result)
	if err != nil {
		return []RequestError{}, err
	}

	var errors []RequestError
	err = json.Unmarshal([]byte(result.Message), &errors)
	if err != nil {
		return []RequestError{}, err
	}
	return errors, nil
}

func parseWebSocketResponse(response WebSocketMessageContainer) (string, error) {
	switch response.Status {
	case 200:
		byteArray, err := json.Marshal(response.Body)
		if err != nil {
			return "", err
		}
		return string(byteArray), nil
	case 400:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", BadRequestException{}
		}
		return "", BadRequestException{
			Errors: errors,
		}
	case 401:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", UnauthorizedException{}
		}
		return "", UnauthorizedException{
			Errors: errors,
		}
	case 402:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", QuotaExceedException{}
		}
		return "", QuotaExceedException{
			Errors: errors,
		}
	case 404:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", NotFoundException{}
		}
		return "", NotFoundException{
			Errors: errors,
		}
	case 405:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", ConflictException{}
		}
		return "", ConflictException{
			Errors: errors,
		}
	case 500:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", InternalServerErrorException{}
		}
		return "", InternalServerErrorException{
			Errors: errors,
		}
	case 502:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", BadGatewayException{}
		}
		return "", BadGatewayException{
			Errors: errors,
		}
	case 503:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", ServiceUnavailableException{}
		}
		return "", ServiceUnavailableException{
			Errors: errors,
		}
	case 504:
		errors, err := readWebSocketErrors(response.Body)
		if err != nil {
			return "", RequestTimeoutException{}
		}
		return "", RequestTimeoutException{
			Errors: errors,
		}
	}
	errors, err := readWebSocketErrors(response.Body)
	if err != nil {
		return "", RequestTimeoutException{}
	}
	return "", RequestTimeoutException{
		Errors: errors,
	}
}

func (p *Gs2WebSocketSession) receive() {

	exit := make(chan string)

	defer func() {
		err := recover()
		if err == "repeated read on failed websocket connection" {
			exit <- "disconnect"
		}
	}()

	for {
		if p.connection == nil || p.connection.client == nil {
			continue
		}
		_, payload, err := p.connection.client.ReadMessage()
		if err != nil {
			log.Error(errors.New("read error"))
		}

		container := WebSocketMessageContainer{}
		err = json.Unmarshal(payload, &container)
		if err != nil {
			log.Error(errors.New("read error"))
		}

		if container.RequestId == "" {
			for _, notificationHandler := range p.notificationHandler {
				body, _ := json.Marshal(container.Body)
				notification := Notification{}
				_ = json.Unmarshal(body, &notification)
				notificationHandler(notification)
			}
			continue
		}

		for _, job := range p.Jobs {
			if job.RequestId == container.RequestId {
				if container.Body == nil {
					job.Callback <- AsyncResult{
						Payload: "",
						Err:     errors.New(string(container.Message)),
					}
				}

				payload_, err := parseWebSocketResponse(
					container,
				)
				job.Callback <- AsyncResult{
					Payload: payload_,
					Err:     err,
				}
				break
			}
		}
		select {
		case <-exit:
			os.Exit(0)
		case <-time.After(1 * time.Microsecond):
			break
		}
	}
}

func (p *Gs2WebSocketSession) send(
	job *WebSocketNetworkJob,
) error {

	if p.connection == nil {
		err := ConnectionBroken{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	jsonText, err := json.Marshal(job.Bodies)
	if err != nil {
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	p.Jobs = append(p.Jobs, job)

	err = p.connection.client.WriteMessage(websocket.TextMessage, jsonText)
	if err != nil {
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	return err
}

func (p *Gs2WebSocketSession) Send(
	job *WebSocketNetworkJob,
	isBlocking bool,
) error {

	if p.connection == nil {
		err := ConnectionBroken{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	if isBlocking {
		if err := p.send(job); err != nil {
			return err
		}
	} else {
		go p.send(job)
	}

	return nil
}

type WebsocketConnection struct {
	client *websocket.Conn
}

func (p *Gs2WebSocketSession) connectAsync(
	callback chan<- AsyncResult,
	isBlocking bool,
) {
	if p.connection == nil {
		connection, _, err := websocket.DefaultDialer.Dial(strings.ReplaceAll(WsEndpointHost, "{region}", string(p.Region)), nil)
		if err != nil {
			callback <- AsyncResult{
				Err: err,
			}
			return
		}
		p.connection = &WebsocketConnection{
			client: connection,
		}
	}

	go p.receive()

	projectTokenCredential, isProjectTokenCredential := p.Credential.(ProjectTokenGs2Credential)
	if isProjectTokenCredential {
		p.projectToken = projectTokenCredential.ProjectToken
		callback <- AsyncResult{}
		return
	} else {
		requestId := WebSocketRequestId(uuid.New().String())
		job := WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: map[string]interface{}{
				"client_id":     p.Credential.GetClientId(),
				"client_secret": p.Credential.GetClientSecret(),
				"x_gs2": map[string]interface{}{
					"service":     "identifier",
					"component":   "projectToken",
					"function":    "login",
					"contentType": "application/json",
					"requestId":   requestId,
				},
			},
			Callback: callback,
		}
		err := p.Send(&job, isBlocking)
		if err != nil {
			return
		}
	}
}

func (p *Gs2WebSocketSession) connectAsyncHandler(
	callback chan<- ConnectAsyncResult,
) {

	internalCallback := make(chan AsyncResult, 1)

	go p.connectAsync(
		internalCallback,
		true,
	)

	result := <-internalCallback
	if result.Err != nil {
		p.Disconnect()
		callback <- ConnectAsyncResult{
			err: result.Err,
		}
		return
	}

	if p.projectToken == "" {
		var loginResult = LoginResult{}
		if err := json.Unmarshal([]byte(result.Payload), &loginResult); err != nil {
			callback <- ConnectAsyncResult{
				err: err,
			}
			return
		}
		p.projectToken = loginResult.AccessToken
	}

	callback <- ConnectAsyncResult{
		err: nil,
	}
}

func (p *Gs2WebSocketSession) ConnectAsync(
	callback chan<- ConnectAsyncResult,
) {

	if p.connection != nil {
		callback <- ConnectAsyncResult{
			err: ConnectionBroken{},
		}
		return
	}

	go p.connectAsyncHandler(
		callback,
	)
}

func (p *Gs2WebSocketSession) Connect() error {

	callback := make(chan ConnectAsyncResult, 1)
	go p.ConnectAsync(
		callback,
	)

	result := <-callback
	if result.err != nil {
		return result.err
	}

	return nil
}

func (p *Gs2WebSocketSession) Disconnect() {
	p.connection = nil
	p.projectToken = ""
}
