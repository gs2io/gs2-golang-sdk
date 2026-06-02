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
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
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

type ErrorResult struct {
	Message  string          `json:"message"`
	Metadata *ResultMetadata `json:"metadata,omitempty"`
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
	Credential                IGs2Credential
	Region                    Region
	projectToken              ProjectToken
	connection                IConnection
	DisableCompressRequest    bool
	DisableDecompressResponse bool
}

func NewGs2RestSession(credential IGs2Credential, region Region) *Gs2RestSession {
	return &Gs2RestSession{
		Credential:                credential,
		Region:                    region,
		DisableCompressRequest:    false,
		DisableDecompressResponse: false,
	}
}

func compressGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p Gs2RestSession) EndpointHost(service string, endpointHost *string) Url {
	if endpointHost == nil {
		endpointHost = &EndpointHost
	}
	return Url(strings.ReplaceAll(strings.ReplaceAll(*endpointHost, "{service}", service), "{region}", string(p.Region)))
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

	var bodyReader io.Reader = nil
	compressedRequest := false
	if job.Method == Post || job.Method == Put {
		if !p.DisableCompressRequest && len(bodies) > 0 {
			compressedBody, compressErr := compressGzip(bodies)
			if compressErr == nil {
				bodyReader = bytes.NewReader(compressedBody)
				compressedRequest = true
			} else {
				bodyReader = bytes.NewReader(bodies)
			}
		} else {
			bodyReader = bytes.NewReader(bodies)
		}
	}

	request, err := http.NewRequest(
		string(job.Method),
		httpUrl,
		bodyReader,
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

	if compressedRequest {
		request.Header.Set("Content-Encoding", "gzip")
	}

	if !p.DisableDecompressResponse {
		request.Header.Set("Accept-Encoding", "gzip")
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
	payload, err := p.parseResponseWithDecompression(response)
	job.Callback <- AsyncResult{
		Payload: payload,
		Err:     err,
	}
	return err
}

func (p Gs2RestSession) parseResponseWithDecompression(response *http.Response) (string, error) {
	defer response.Body.Close()

	var bodyReader io.Reader = response.Body
	if !p.DisableDecompressResponse && response.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(response.Body)
		if err != nil {
			return "", err
		}
		defer gzipReader.Close()
		bodyReader = gzipReader
	}

	byteArray, err := io.ReadAll(bodyReader)
	if err != nil {
		return "", err
	}

	if response.StatusCode == 200 {
		return string(byteArray), nil
	}

	var result ErrorResult
	err = json.Unmarshal(byteArray, &result)
	if err != nil {
		return "", createExceptionByStatusCode(response.StatusCode, nil, nil)
	}

	var errors []RequestError
	err = json.Unmarshal([]byte(result.Message), &errors)
	if err != nil {
		return "", createExceptionByStatusCode(response.StatusCode, nil, result.Metadata)
	}

	return "", createExceptionByStatusCode(response.StatusCode, errors, result.Metadata)
}

func createExceptionByStatusCode(statusCode int, errors []RequestError, metadata *ResultMetadata) error {
	switch statusCode {
	case 400:
		return BadRequestException{Errors: errors, Metadata: metadata}
	case 401:
		return UnauthorizedException{Errors: errors, Metadata: metadata}
	case 402:
		return QuotaExceedException{Errors: errors, Metadata: metadata}
	case 404:
		return NotFoundException{Errors: errors, Metadata: metadata}
	case 409:
		return ConflictException{Errors: errors, Metadata: metadata}
	case 500:
		return InternalServerErrorException{Errors: errors, Metadata: metadata}
	case 502:
		return BadGatewayException{Errors: errors, Metadata: metadata}
	case 503:
		return ServiceUnavailableException{Errors: errors, Metadata: metadata}
	case 504:
		return RequestTimeoutException{Errors: errors, Metadata: metadata}
	default:
		return RequestTimeoutException{Errors: errors, Metadata: metadata}
	}
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
