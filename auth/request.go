/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package auth

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type LoginRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	UserId       *string `json:"userId"`
	TimeOffset   *int32  `json:"timeOffset"`
}

func NewLoginRequestFromJson(data string) LoginRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewLoginRequestFromDict(dict2)
}

func NewLoginRequestFromDict(data map[string]interface{}) LoginRequest {
	return LoginRequest{
		UserId:     core.CastString(data["userId"]),
		TimeOffset: core.CastInt32(data["timeOffset"]),
	}
}

func (p LoginRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":     p.UserId,
		"timeOffset": p.TimeOffset,
	}
}

func (p LoginRequest) Pointer() *LoginRequest {
	return &p
}

type LoginBySignatureRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	KeyId        *string `json:"keyId"`
	Body         *string `json:"body"`
	Signature    *string `json:"signature"`
}

func NewLoginBySignatureRequestFromJson(data string) LoginBySignatureRequest {
	dict := map[string]string{}
	_ = json.Unmarshal([]byte(data), &dict)
	dict2 := map[string]interface{}{}
	data2, _ := json.Marshal(&dict2)
	_ = json.Unmarshal(data2, &dict2)
	return NewLoginBySignatureRequestFromDict(dict2)
}

func NewLoginBySignatureRequestFromDict(data map[string]interface{}) LoginBySignatureRequest {
	return LoginBySignatureRequest{
		KeyId:     core.CastString(data["keyId"]),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p LoginBySignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"keyId":     p.KeyId,
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p LoginBySignatureRequest) Pointer() *LoginBySignatureRequest {
	return &p
}
