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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Url string

func (p Url) AppendPath(
	path string,
	stringReplacer *strings.Replacer,
) Url {
	if stringReplacer == nil {
		return Url(string(p) + path)
	}
	return Url(stringReplacer.Replace(string(p) + path))
}

type HttpMethod string
type HttpHeaders map[string]string
type QueryStrings map[string]string
type Bodies map[string]interface{}

const (
	Get    = HttpMethod("GET")
	Post   = HttpMethod("POST")
	Put    = HttpMethod("PUT")
	Delete = HttpMethod("DELETE")
)

type ConnectionBroken struct {
}

func (p ConnectionBroken) Error() string {
	return "connection broken"
}

func readBody(response *http.Response) (string, error) {
	byteArray, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}

type ErrorResult struct {
	Message string `json:"message"`
}

func readErrors(response *http.Response) ([]RequestError, error) {
	byteArray, err := ioutil.ReadAll(response.Body)
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

func parseResponse(response *http.Response) (string, error) {
	if response.StatusCode == 200 {
		body, err := readBody(response)
		if err != nil {
			return "", err
		}
		return body, nil
	}
	if response.StatusCode == 400 {
		errors, err := readErrors(response)
		if err != nil {
			return "", BadRequestException{}
		}
		return "", BadRequestException{
			Errors: errors,
		}
	}
	if response.StatusCode == 401 {
		errors, err := readErrors(response)
		if err != nil {
			return "", UnauthorizedException{}
		}
		return "", UnauthorizedException{
			Errors: errors,
		}
	}
	if response.StatusCode == 402 {
		errors, err := readErrors(response)
		if err != nil {
			return "", QuotaExceedException{}
		}
		return "", QuotaExceedException{
			Errors: errors,
		}
	}
	if response.StatusCode == 404 {
		errors, err := readErrors(response)
		if err != nil {
			return "", NotFoundException{}
		}
		return "", NotFoundException{
			Errors: errors,
		}
	}
	if response.StatusCode == 409 {
		errors, err := readErrors(response)
		if err != nil {
			return "", ConflictException{}
		}
		return "", ConflictException{
			Errors: errors,
		}
	}
	if response.StatusCode == 500 {
		errors, err := readErrors(response)
		if err != nil {
			return "", InternalServerErrorException{}
		}
		return "", InternalServerErrorException{
			Errors: errors,
		}
	}
	if response.StatusCode == 502 {
		errors, err := readErrors(response)
		if err != nil {
			return "", BadGatewayException{}
		}
		return "", BadGatewayException{
			Errors: errors,
		}
	}
	if response.StatusCode == 503 {
		errors, err := readErrors(response)
		if err != nil {
			return "", ServiceUnavailableException{}
		}
		return "", ServiceUnavailableException{
			Errors: errors,
		}
	}
	if response.StatusCode == 504 {
		errors, err := readErrors(response)
		if err != nil {
			return "", RequestTimeoutException{}
		}
		return "", RequestTimeoutException{
			Errors: errors,
		}
	}
	errors, err := readErrors(response)
	if err != nil {
		return "", RequestTimeoutException{}
	}
	return "", RequestTimeoutException{
		Errors: errors,
	}
}

type AsyncResult struct {
	Payload string
	Err     error
}

type NetworkJob struct {
	Url          Url
	Method       HttpMethod
	Headers      HttpHeaders
	QueryStrings QueryStrings
	Bodies       Bodies
	Callback     chan<- AsyncResult
}

type IConnection interface {
	Client() *http.Client
}

type Connection struct {
	client *http.Client
}

func (p Connection) Client() *http.Client {
	return p.client
}

type Gs2RestSession struct {
	Credential   IGs2Credential
	Region       Region
	projectToken ProjectToken
	connection   IConnection
}

func (p Gs2RestSession) EndpointHost(service string) Url {
	return Url(strings.ReplaceAll(strings.ReplaceAll(EndpointHost, "{service}", service), "{region}", string(p.Region)))
}

func (p Gs2RestSession) CreateAuthorizationHeader() map[string]string {
	return map[string]string{
		"Content-Type":    "application/json",
		"X-GS2-CLIENT-ID": string(*p.Credential.GetClientId()),
		"Authorization":   fmt.Sprintf("Bearer %s", p.projectToken),
	}
}

func MarshalJson(v interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (p Gs2RestSession) send(job *NetworkJob) error {

	if p.connection == nil {
		err := ConnectionBroken{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	bodies, err := MarshalJson(job.Bodies)
	if err != nil {
		err := BadRequestException{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	httpUrl := string(job.Url)
	if len(job.QueryStrings) > 0 {
		query := url.Values{}

		for key, value := range job.QueryStrings {
			query.Add(key, value)
		}

		httpUrl += "?" + query.Encode()
	}

	var reader *strings.Reader = nil
	if bodies != nil {
		reader = strings.NewReader(string(bodies))
	}

	request, err := http.NewRequest(
		string(job.Method),
		httpUrl,
		func() io.Reader {
			if job.Method == Post || job.Method == Put {
				return reader
			} else {
				return nil
			}
		}(),
	)
	if err != nil {
		err := BadRequestException{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}
	for key, value := range job.Headers {
		request.Header.Add(key, value)
	}

	response, err := p.connection.Client().Do(request)
	if err != nil {
		err := UnknownException{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}
	payload, err := parseResponse(
		response,
	)
	job.Callback <- AsyncResult{
		Payload: payload,
		Err:     err,
	}
	return err
}

func (p *Gs2RestSession) Send(
	job *NetworkJob,
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

func (p *Gs2RestSession) connectWithCustomConnectionAsync(
	createConnection func() IConnection,
	callback chan<- AsyncResult,
	isBlocking bool,
) {
	if p.connection == nil {
		p.connection = createConnection()
	}

	projectTokenCredential, isProjectTokenCredential := p.Credential.(ProjectTokenGs2Credential)
	if isProjectTokenCredential {
		p.projectToken = projectTokenCredential.ProjectToken
		callback <- AsyncResult{}
		return
	} else {
		job := NetworkJob{
			Url:    Url(strings.ReplaceAll(strings.ReplaceAll(EndpointHost, "{service}", "identifier"), "{region}", string(p.Region))).AppendPath("/projectToken/login", nil),
			Method: Post,
			Bodies: map[string]interface{}{
				"client_id":     p.Credential.GetClientId(),
				"client_secret": p.Credential.GetClientSecret(),
			},
			Headers: HttpHeaders{
				"Content-Type": "application/json",
			},
			Callback: callback,
		}
		err := p.Send(&job, isBlocking)
		if err != nil {
			return
		}
	}
}

type ConnectAsyncResult struct {
	err error
}

func connectAsyncHandler(
	callback chan<- ConnectAsyncResult,
	session *Gs2RestSession,
) {
	connectWithCustomConnectionAsyncHandler(
		func() IConnection {
			return &Connection{
				client: new(http.Client),
			}
		},
		callback,
		session,
	)
}

func connectWithCustomConnectionAsyncHandler(
	createConnection func() IConnection,
	callback chan<- ConnectAsyncResult,
	session *Gs2RestSession,
) {

	internalCallback := make(chan AsyncResult, 1)

	go session.connectWithCustomConnectionAsync(
		createConnection,
		internalCallback,
		true,
	)

	result := <-internalCallback
	if result.Err != nil {
		session.Disconnect()
		callback <- ConnectAsyncResult{
			err: result.Err,
		}
		return
	}

	if session.projectToken == "" {
		var loginResult = LoginResult{}
		if err := json.Unmarshal([]byte(result.Payload), &loginResult); err != nil {
			callback <- ConnectAsyncResult{
				err: err,
			}
			return
		}
		session.projectToken = loginResult.AccessToken
	}

	callback <- ConnectAsyncResult{
		err: nil,
	}
}

func (p *Gs2RestSession) ConnectWithCustomConnectionAsync(
	createConnection func() IConnection,
	callback chan<- ConnectAsyncResult,
) {

	if p.connection != nil {
		callback <- ConnectAsyncResult{
			err: ConnectionBroken{},
		}
		return
	}

	go connectWithCustomConnectionAsyncHandler(
		createConnection,
		callback,
		p,
	)
}

func (p *Gs2RestSession) ConnectAsync(
	callback chan<- ConnectAsyncResult,
) {

	if p.connection != nil {
		callback <- ConnectAsyncResult{
			err: ConnectionBroken{},
		}
		return
	}

	go connectAsyncHandler(
		callback,
		p,
	)
}

func (p *Gs2RestSession) ConnectWithCustomConnection(
	createConnection func() IConnection,
) error {

	callback := make(chan ConnectAsyncResult, 1)
	go p.ConnectWithCustomConnectionAsync(
		createConnection,
		callback,
	)

	result := <-callback
	if result.err != nil {
		return result.err
	}

	return nil
}

func (p *Gs2RestSession) Connect() error {

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

func (p *Gs2RestSession) Disconnect() {
	p.connection = nil
	p.projectToken = ""
}
