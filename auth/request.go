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
	"strconv"

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

func (p *LoginRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LoginRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LoginRequest{}
	} else {
		*p = LoginRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.UserId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.UserId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.UserId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffset"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.TimeOffset); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewLoginRequestFromJson(data string) (LoginRequest, error) {
	req := LoginRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return LoginRequest{}, err
	}
	return req, nil
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

func (p *LoginBySignatureRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LoginBySignatureRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LoginBySignatureRequest{}
	} else {
		*p = LoginBySignatureRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["body"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.Body = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.Body = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.Body = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.Body = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.Body = &strValue
			default:
				if err := json.Unmarshal(*v, &p.Body); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.Signature = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.Signature = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.Signature = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.Signature = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.Signature = &strValue
			default:
				if err := json.Unmarshal(*v, &p.Signature); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewLoginBySignatureRequestFromJson(data string) (LoginBySignatureRequest, error) {
	req := LoginBySignatureRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return LoginBySignatureRequest{}, err
	}
	return req, nil
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

func (p *FederationRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FederationRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = FederationRequest{}
	} else {
		*p = FederationRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["originalUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.OriginalUserId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.OriginalUserId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.OriginalUserId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.OriginalUserId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.OriginalUserId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.OriginalUserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.UserId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.UserId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.UserId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["policyDocument"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PolicyDocument = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PolicyDocument = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PolicyDocument = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PolicyDocument = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PolicyDocument = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PolicyDocument); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffset"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.TimeOffset); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewFederationRequestFromJson(data string) (FederationRequest, error) {
	req := FederationRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return FederationRequest{}, err
	}
	return req, nil
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

func (p *IssueTimeOffsetTokenByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IssueTimeOffsetTokenByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = IssueTimeOffsetTokenByUserIdRequest{}
	} else {
		*p = IssueTimeOffsetTokenByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.UserId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.UserId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.UserId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.UserId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffset"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.TimeOffset); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewIssueTimeOffsetTokenByUserIdRequestFromJson(data string) (IssueTimeOffsetTokenByUserIdRequest, error) {
	req := IssueTimeOffsetTokenByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return IssueTimeOffsetTokenByUserIdRequest{}, err
	}
	return req, nil
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
