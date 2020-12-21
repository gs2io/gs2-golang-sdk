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

import "encoding/json"

type Region string
type OwnerId string

const (
	ApNortheast1 Region = "ap-northeast-1"
	UsEast1      Region = "us-east-1"
	EuWest1      Region = "eu-west-1"
	ApSoutheast1 Region = "ap-southeast-1"
)

type ClientId string
type ClientSecret string
type ProjectToken string

type RequestId = string
type ContextStack = string
type AccessToken = string

type AccessTokenType string
type ExpiresIn int64

var (
	EndpointHost   = "https://{service}.{region}.gen2.gs2io.com"
	WsEndpointHost = "wss://gateway-ws.{region}.gen2.gs2io.com"
)

type IGs2Credential interface {
	GetClientId() *ClientId
	GetClientSecret() *ClientSecret
}

type BasicGs2Credential struct {
	ClientId     ClientId
	ClientSecret ClientSecret
}

func (p BasicGs2Credential) GetClientId() *ClientId {
	return &p.ClientId
}

func (p BasicGs2Credential) GetClientSecret() *ClientSecret {
	return &p.ClientSecret
}

type ProjectTokenGs2Credential struct {
	IGs2Credential
	ownerId      OwnerId
	projectToken ProjectToken
}

func (p ProjectTokenGs2Credential) GetClientId() *ClientId {
	clientId := ClientId(p.ownerId)
	return &clientId
}

func (p ProjectTokenGs2Credential) GetClientSecret() *ClientSecret {
	return nil
}

func (p ProjectTokenGs2Credential) getProjectToken() *ProjectToken {
	return &p.projectToken
}

type Gs2Result interface {
	toDict() map[string]*interface{}
}

type Gs2Model interface {
	toDict() map[string]*interface{}
}

type LoginResult struct {
	Gs2Result
	AccessToken ProjectToken    `json:"access_token"`
	TokenType   AccessTokenType `json:"token_type"`
	ExpiresIn   ExpiresIn       `json:"expires_in"`
}

func (p LoginResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accessToken": p.AccessToken,
		"tokenType":   p.TokenType,
		"expiresIn":   p.ExpiresIn,
	}
}

type ISession interface {
	getCredential() *IGs2Credential
	getRegion() *Region
	getProjectToken() *ProjectToken

	connect()
	connectAsync()
	disconnect()
}

type Notification struct {
	Issuer  string `json:"issuer"`
	Subject string `json:"subject"`
	Payload string `json:"payload"`
}

type String string

func (p *String) UnmarshalJSON(data []byte) error {
	println(string(data))
	if data[0] == '"' {
		var d string
		err := json.Unmarshal(data, &d)
		if err != nil {
			return err
		}
		data = []byte(d)
	}
	*p = String(data)
	return nil
}
