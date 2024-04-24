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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffset      *int32  `json:"timeOffset"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewLoginRequestFromJson(data string) LoginRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLoginRequestFromDict(dict)
}

func NewLoginRequestFromDict(data map[string]interface{}) LoginRequest {
	return LoginRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffset:      core.CastInt32(data["timeOffset"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p LoginRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffset":      p.TimeOffset,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p LoginRequest) Pointer() *LoginRequest {
	return &p
}

type LoginBySignatureRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	KeyId           *string `json:"keyId"`
	Body            *string `json:"body"`
	Signature       *string `json:"signature"`
}

func NewLoginBySignatureRequestFromJson(data string) LoginBySignatureRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLoginBySignatureRequestFromDict(dict)
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

type FederationRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	OriginalUserId  *string `json:"originalUserId"`
	UserId          *string `json:"userId"`
	PolicyDocument  *string `json:"policyDocument"`
	TimeOffset      *int32  `json:"timeOffset"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewFederationRequestFromJson(data string) FederationRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFederationRequestFromDict(dict)
}

func NewFederationRequestFromDict(data map[string]interface{}) FederationRequest {
	return FederationRequest{
		OriginalUserId:  core.CastString(data["originalUserId"]),
		UserId:          core.CastString(data["userId"]),
		PolicyDocument:  core.CastString(data["policyDocument"]),
		TimeOffset:      core.CastInt32(data["timeOffset"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p FederationRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"originalUserId":  p.OriginalUserId,
		"userId":          p.UserId,
		"policyDocument":  p.PolicyDocument,
		"timeOffset":      p.TimeOffset,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p FederationRequest) Pointer() *FederationRequest {
	return &p
}

type IssueTimeOffsetTokenByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffset      *int32  `json:"timeOffset"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewIssueTimeOffsetTokenByUserIdRequestFromJson(data string) IssueTimeOffsetTokenByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueTimeOffsetTokenByUserIdRequestFromDict(dict)
}

func NewIssueTimeOffsetTokenByUserIdRequestFromDict(data map[string]interface{}) IssueTimeOffsetTokenByUserIdRequest {
	return IssueTimeOffsetTokenByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffset:      core.CastInt32(data["timeOffset"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p IssueTimeOffsetTokenByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffset":      p.TimeOffset,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p IssueTimeOffsetTokenByUserIdRequest) Pointer() *IssueTimeOffsetTokenByUserIdRequest {
	return &p
}
