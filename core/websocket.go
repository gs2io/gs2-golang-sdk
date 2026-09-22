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
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
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
	// Jobs は応答待ちの要求。★send と receive が別の goroutine から触るので jobsMu で守り、
	// 応答が来たもの・接続が切れたものは取り除く（以前は増え続けていた）。
	Jobs []*WebSocketNetworkJob
	// jobsMu は Jobs と connection を守る。★ポインタで持つ（構造体リテラルで作られ、値レシーバの
	// メソッドもあるので、値として複製されても同じ錠を指すように）。初期化は mu() で遅延して行う。
	jobsMu *sync.Mutex

	notificationHandler []func(message Notification)

	// SteadyEndpoint は Steady（専用フリート）の基点（https://<host>）。空なら共有クラウド。
	// 接続先は wss://<host>/ になり、handshake に上限（SteadyConnectTimeout）が付く。
	SteadyEndpoint string
	// DialContext は省略可。指定すると WebSocket の TCP 接続にこれを使う（プロキシや
	// 名前解決の差し替え、試験でのアドレス選択など）。TLS の SNI と検証は URL のホストのまま。
	DialContext func(ctx context.Context, network, addr string) (net.Conn, error)
}

// webSocketUrl は接続先。SteadyEndpoint ＞ WsEndpointHost。
func (p *Gs2WebSocketSession) webSocketUrl() string {
	if u := steadyWebSocketUrl(p.SteadyEndpoint); u != "" {
		return u
	}
	return strings.ReplaceAll(WsEndpointHost, "{region}", string(p.Region))
}

// dialer は Steady のときだけ handshake に上限を置く。共有クラウドは従来どおり DefaultDialer。
func (p *Gs2WebSocketSession) dialer() *websocket.Dialer {
	if normalizeSteadyEndpoint(p.SteadyEndpoint) == "" && p.DialContext == nil {
		return websocket.DefaultDialer
	}
	d := &websocket.Dialer{Proxy: http.ProxyFromEnvironment, NetDialContext: p.DialContext}
	if normalizeSteadyEndpoint(p.SteadyEndpoint) != "" {
		d.HandshakeTimeout = SteadyConnectTimeout
	}
	return d
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

// receive は接続から読み続け、応答を待ち中の要求へ配る。
//
// ★接続が切れたら（読み取りの誤り）、**待ち中の要求すべてに ConnectionBroken を返して**
// 抜ける。以前は誤りを記録して読み続け、次の読み取りで gorilla が panic する経路に入り、
// 待ち中の呼び出し（SetUserId 等の同期呼び出し）が永久に返らなかった。
// サーバーが応答の前に接続を閉じた場合（gateway の setUserId が自分自身の接続を切る形、
// ノードの停止、ネットワーク断）に当たる。
func (p *Gs2WebSocketSession) receive(connection *WebsocketConnection) {
	for {
		_, payload, err := connection.client.ReadMessage()
		if err != nil {
			p.dropConnection(connection, err)
			return
		}

		container := WebSocketMessageContainer{}
		if err := json.Unmarshal(payload, &container); err != nil {
			log.Error(errors.New("read error"))
			continue
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

		if job := p.takeJob(container.RequestId); job != nil {
			if container.Body == nil {
				job.Callback <- AsyncResult{
					Payload: "",
					Err:     errors.New(string(container.Message)),
				}
				continue
			}
			payload_, err := parseWebSocketResponse(container)
			job.Callback <- AsyncResult{
				Payload: payload_,
				Err:     err,
			}
		}
	}
}

// sessionLocks は jobsMu の遅延初期化を守る（全セッション共通。持つのは一瞬）。
var sessionLocks sync.Mutex

// mu は Jobs / connection を守る錠（無ければ作る）。
func (p *Gs2WebSocketSession) mu() *sync.Mutex {
	sessionLocks.Lock()
	defer sessionLocks.Unlock()
	if p.jobsMu == nil {
		p.jobsMu = &sync.Mutex{}
	}
	return p.jobsMu
}

// takeJob は requestId の要求を待ち行列から外して返す（無ければ nil）。
func (p *Gs2WebSocketSession) takeJob(requestId WebSocketRequestId) *WebSocketNetworkJob {
	p.mu().Lock()
	defer p.mu().Unlock()
	for i, job := range p.Jobs {
		if job.RequestId == requestId {
			p.Jobs = append(p.Jobs[:i], p.Jobs[i+1:]...)
			return job
		}
	}
	return nil
}

// dropConnection は切れた接続を捨て、待ち中の要求すべてに ConnectionBroken を返す。
// ★既に別の接続に差し替わっていたら（Disconnect → Connect の後）、何もしない。
func (p *Gs2WebSocketSession) dropConnection(connection *WebsocketConnection, cause error) {
	p.mu().Lock()
	if p.connection != connection {
		p.mu().Unlock()
		return
	}
	p.connection = nil
	p.projectToken = ""
	jobs := p.Jobs
	p.Jobs = nil
	p.mu().Unlock()
	_ = connection.client.Close()
	for _, job := range jobs {
		job.Callback <- AsyncResult{Payload: "", Err: ConnectionBroken{}}
	}
	if len(jobs) > 0 {
		log.Error(fmt.Errorf("websocket closed with %d pending request(s): %v", len(jobs), cause))
	}
}

func (p *Gs2WebSocketSession) send(
	job *WebSocketNetworkJob,
) error {

	p.mu().Lock()
	connection := p.connection
	if connection == nil {
		p.mu().Unlock()
		err := ConnectionBroken{}
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}
	p.Jobs = append(p.Jobs, job)
	p.mu().Unlock()

	jsonText, err := json.Marshal(job.Bodies)
	if err != nil {
		p.takeJob(job.RequestId)
		job.Callback <- AsyncResult{
			Payload: "",
			Err:     err,
		}
		return err
	}

	// ★gorilla の WriteMessage は並行に呼べない（1 本の接続に書くのは 1 goroutine まで）。
	connection.writeMu.Lock()
	err = connection.client.WriteMessage(websocket.TextMessage, jsonText)
	connection.writeMu.Unlock()
	if err != nil {
		// 書けなかった要求は届いていないので、待ち行列から外してその場で返す。
		if p.takeJob(job.RequestId) != nil {
			job.Callback <- AsyncResult{
				Payload: "",
				Err:     err,
			}
		}
		return err
	}

	return nil
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
	client  *websocket.Conn
	writeMu sync.Mutex
}

func (p *Gs2WebSocketSession) connectAsync(
	callback chan<- AsyncResult,
	isBlocking bool,
) {
	if p.connection == nil {
		connection, _, err := p.dialer().Dial(p.webSocketUrl(), nil)
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

	go p.receive(p.connection)

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

// Disconnect は接続を閉じる。待ち中の要求には ConnectionBroken が返る。
func (p *Gs2WebSocketSession) Disconnect() {
	p.mu().Lock()
	connection := p.connection
	p.mu().Unlock()
	if connection != nil {
		p.dropConnection(connection, errors.New("disconnect"))
	}
	p.projectToken = ""
}
