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

package project

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type CreateAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	FullName        *string `json:"fullName"`
	CompanyName     *string `json:"companyName"`
	Password        *string `json:"password"`
	Lang            *string `json:"lang"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreateAccountRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateAccountRequest{}
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
		*p = CreateAccountRequest{}
	} else {
		*p = CreateAccountRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["email"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Email = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Email = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Email = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Email)
				}
			}
		}
		if v, ok := d["fullName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FullName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FullName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FullName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FullName)
				}
			}
		}
		if v, ok := d["companyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompanyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompanyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompanyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompanyName)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["lang"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Lang = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Lang = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Lang = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Lang = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Lang = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Lang)
				}
			}
		}
	}
	return nil
}

func NewCreateAccountRequestFromJson(data string) (CreateAccountRequest, error) {
	req := CreateAccountRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateAccountRequest{}, err
	}
	return req, nil
}

func NewCreateAccountRequestFromDict(data map[string]interface{}) CreateAccountRequest {
	return CreateAccountRequest{
		Email: func() *string {
			v, ok := data["email"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["email"])
		}(),
		FullName: func() *string {
			v, ok := data["fullName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fullName"])
		}(),
		CompanyName: func() *string {
			v, ok := data["companyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["companyName"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		Lang: func() *string {
			v, ok := data["lang"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lang"])
		}(),
	}
}

func (p CreateAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":       p.Email,
		"fullName":    p.FullName,
		"companyName": p.CompanyName,
		"password":    p.Password,
		"lang":        p.Lang,
	}
}

func (p CreateAccountRequest) Pointer() *CreateAccountRequest {
	return &p
}

type VerifyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	VerifyToken     *string `json:"verifyToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *VerifyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRequest{}
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
		*p = VerifyRequest{}
	} else {
		*p = VerifyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["verifyToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyToken)
				}
			}
		}
	}
	return nil
}

func NewVerifyRequestFromJson(data string) (VerifyRequest, error) {
	req := VerifyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRequest{}, err
	}
	return req, nil
}

func NewVerifyRequestFromDict(data map[string]interface{}) VerifyRequest {
	return VerifyRequest{
		VerifyToken: func() *string {
			v, ok := data["verifyToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyToken"])
		}(),
	}
}

func (p VerifyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"verifyToken": p.VerifyToken,
	}
}

func (p VerifyRequest) Pointer() *VerifyRequest {
	return &p
}

type SignInRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	Password        *string `json:"password"`
	Otp             *string `json:"otp"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *SignInRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SignInRequest{}
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
		*p = SignInRequest{}
	} else {
		*p = SignInRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["email"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Email = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Email = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Email = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Email)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["otp"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Otp = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Otp = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Otp = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Otp = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Otp = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Otp)
				}
			}
		}
	}
	return nil
}

func NewSignInRequestFromJson(data string) (SignInRequest, error) {
	req := SignInRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SignInRequest{}, err
	}
	return req, nil
}

func NewSignInRequestFromDict(data map[string]interface{}) SignInRequest {
	return SignInRequest{
		Email: func() *string {
			v, ok := data["email"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["email"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		Otp: func() *string {
			v, ok := data["otp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["otp"])
		}(),
	}
}

func (p SignInRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":    p.Email,
		"password": p.Password,
		"otp":      p.Otp,
	}
}

func (p SignInRequest) Pointer() *SignInRequest {
	return &p
}

type ForgetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	Lang            *string `json:"lang"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *ForgetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ForgetRequest{}
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
		*p = ForgetRequest{}
	} else {
		*p = ForgetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["email"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Email = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Email = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Email = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Email)
				}
			}
		}
		if v, ok := d["lang"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Lang = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Lang = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Lang = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Lang = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Lang = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Lang)
				}
			}
		}
	}
	return nil
}

func NewForgetRequestFromJson(data string) (ForgetRequest, error) {
	req := ForgetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ForgetRequest{}, err
	}
	return req, nil
}

func NewForgetRequestFromDict(data map[string]interface{}) ForgetRequest {
	return ForgetRequest{
		Email: func() *string {
			v, ok := data["email"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["email"])
		}(),
		Lang: func() *string {
			v, ok := data["lang"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lang"])
		}(),
	}
}

func (p ForgetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email": p.Email,
		"lang":  p.Lang,
	}
}

func (p ForgetRequest) Pointer() *ForgetRequest {
	return &p
}

type IssuePasswordRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	IssuePasswordToken *string `json:"issuePasswordToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *IssuePasswordRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IssuePasswordRequest{}
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
		*p = IssuePasswordRequest{}
	} else {
		*p = IssuePasswordRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["issuePasswordToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.IssuePasswordToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.IssuePasswordToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.IssuePasswordToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.IssuePasswordToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.IssuePasswordToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.IssuePasswordToken)
				}
			}
		}
	}
	return nil
}

func NewIssuePasswordRequestFromJson(data string) (IssuePasswordRequest, error) {
	req := IssuePasswordRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return IssuePasswordRequest{}, err
	}
	return req, nil
}

func NewIssuePasswordRequestFromDict(data map[string]interface{}) IssuePasswordRequest {
	return IssuePasswordRequest{
		IssuePasswordToken: func() *string {
			v, ok := data["issuePasswordToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["issuePasswordToken"])
		}(),
	}
}

func (p IssuePasswordRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"issuePasswordToken": p.IssuePasswordToken,
	}
}

func (p IssuePasswordRequest) Pointer() *IssuePasswordRequest {
	return &p
}

type UpdateAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Email           *string `json:"email"`
	FullName        *string `json:"fullName"`
	CompanyName     *string `json:"companyName"`
	Password        *string `json:"password"`
	AccountToken    *string `json:"accountToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *UpdateAccountRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateAccountRequest{}
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
		*p = UpdateAccountRequest{}
	} else {
		*p = UpdateAccountRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["email"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Email = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Email = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Email = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Email = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Email)
				}
			}
		}
		if v, ok := d["fullName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FullName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FullName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FullName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FullName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FullName)
				}
			}
		}
		if v, ok := d["companyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompanyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompanyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompanyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompanyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompanyName)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
	}
	return nil
}

func NewUpdateAccountRequestFromJson(data string) (UpdateAccountRequest, error) {
	req := UpdateAccountRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateAccountRequest{}, err
	}
	return req, nil
}

func NewUpdateAccountRequestFromDict(data map[string]interface{}) UpdateAccountRequest {
	return UpdateAccountRequest{
		Email: func() *string {
			v, ok := data["email"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["email"])
		}(),
		FullName: func() *string {
			v, ok := data["fullName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fullName"])
		}(),
		CompanyName: func() *string {
			v, ok := data["companyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["companyName"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p UpdateAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"email":        p.Email,
		"fullName":     p.FullName,
		"companyName":  p.CompanyName,
		"password":     p.Password,
		"accountToken": p.AccountToken,
	}
}

func (p UpdateAccountRequest) Pointer() *UpdateAccountRequest {
	return &p
}

type EnableMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *EnableMfaRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EnableMfaRequest{}
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
		*p = EnableMfaRequest{}
	} else {
		*p = EnableMfaRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
	}
	return nil
}

func NewEnableMfaRequestFromJson(data string) (EnableMfaRequest, error) {
	req := EnableMfaRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return EnableMfaRequest{}, err
	}
	return req, nil
}

func NewEnableMfaRequestFromDict(data map[string]interface{}) EnableMfaRequest {
	return EnableMfaRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p EnableMfaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
	}
}

func (p EnableMfaRequest) Pointer() *EnableMfaRequest {
	return &p
}

type ChallengeMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	Passcode        *string `json:"passcode"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *ChallengeMfaRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ChallengeMfaRequest{}
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
		*p = ChallengeMfaRequest{}
	} else {
		*p = ChallengeMfaRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["passcode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Passcode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Passcode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Passcode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Passcode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Passcode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Passcode)
				}
			}
		}
	}
	return nil
}

func NewChallengeMfaRequestFromJson(data string) (ChallengeMfaRequest, error) {
	req := ChallengeMfaRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ChallengeMfaRequest{}, err
	}
	return req, nil
}

func NewChallengeMfaRequestFromDict(data map[string]interface{}) ChallengeMfaRequest {
	return ChallengeMfaRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		Passcode: func() *string {
			v, ok := data["passcode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["passcode"])
		}(),
	}
}

func (p ChallengeMfaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"passcode":     p.Passcode,
	}
}

func (p ChallengeMfaRequest) Pointer() *ChallengeMfaRequest {
	return &p
}

type DisableMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DisableMfaRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DisableMfaRequest{}
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
		*p = DisableMfaRequest{}
	} else {
		*p = DisableMfaRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
	}
	return nil
}

func NewDisableMfaRequestFromJson(data string) (DisableMfaRequest, error) {
	req := DisableMfaRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DisableMfaRequest{}, err
	}
	return req, nil
}

func NewDisableMfaRequestFromDict(data map[string]interface{}) DisableMfaRequest {
	return DisableMfaRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p DisableMfaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
	}
}

func (p DisableMfaRequest) Pointer() *DisableMfaRequest {
	return &p
}

type DeleteAccountRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DeleteAccountRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteAccountRequest{}
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
		*p = DeleteAccountRequest{}
	} else {
		*p = DeleteAccountRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
	}
	return nil
}

func NewDeleteAccountRequestFromJson(data string) (DeleteAccountRequest, error) {
	req := DeleteAccountRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteAccountRequest{}, err
	}
	return req, nil
}

func NewDeleteAccountRequestFromDict(data map[string]interface{}) DeleteAccountRequest {
	return DeleteAccountRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p DeleteAccountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
	}
}

func (p DeleteAccountRequest) Pointer() *DeleteAccountRequest {
	return &p
}

type DescribeProjectsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeProjectsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeProjectsRequest{}
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
		*p = DescribeProjectsRequest{}
	} else {
		*p = DescribeProjectsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeProjectsRequestFromJson(data string) (DescribeProjectsRequest, error) {
	req := DescribeProjectsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeProjectsRequest{}, err
	}
	return req, nil
}

func NewDescribeProjectsRequestFromDict(data map[string]interface{}) DescribeProjectsRequest {
	return DescribeProjectsRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeProjectsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeProjectsRequest) Pointer() *DescribeProjectsRequest {
	return &p
}

type CreateProjectRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	AccountToken            *string `json:"accountToken"`
	Name                    *string `json:"name"`
	Description             *string `json:"description"`
	Plan                    *string `json:"plan"`
	Currency                *string `json:"currency"`
	ActivateRegionName      *string `json:"activateRegionName"`
	BillingMethodName       *string `json:"billingMethodName"`
	EnableEventBridge       *string `json:"enableEventBridge"`
	EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion    *string `json:"eventBridgeAwsRegion"`
	DryRun                  *bool   `json:"dryRun"`
}

func (p *CreateProjectRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateProjectRequest{}
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
		*p = CreateProjectRequest{}
	} else {
		*p = CreateProjectRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["plan"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Plan = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Plan = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Plan = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Plan)
				}
			}
		}
		if v, ok := d["currency"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Currency = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Currency = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Currency = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Currency = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Currency)
				}
			}
		}
		if v, ok := d["activateRegionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ActivateRegionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ActivateRegionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ActivateRegionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ActivateRegionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ActivateRegionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ActivateRegionName)
				}
			}
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
		if v, ok := d["enableEventBridge"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableEventBridge = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableEventBridge = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableEventBridge = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableEventBridge)
				}
			}
		}
		if v, ok := d["eventBridgeAwsAccountId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsAccountId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsAccountId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsAccountId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsAccountId)
				}
			}
		}
		if v, ok := d["eventBridgeAwsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsRegion = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsRegion = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsRegion = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsRegion)
				}
			}
		}
	}
	return nil
}

func NewCreateProjectRequestFromJson(data string) (CreateProjectRequest, error) {
	req := CreateProjectRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateProjectRequest{}, err
	}
	return req, nil
}

func NewCreateProjectRequestFromDict(data map[string]interface{}) CreateProjectRequest {
	return CreateProjectRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Plan: func() *string {
			v, ok := data["plan"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["plan"])
		}(),
		Currency: func() *string {
			v, ok := data["currency"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["currency"])
		}(),
		ActivateRegionName: func() *string {
			v, ok := data["activateRegionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["activateRegionName"])
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
		EnableEventBridge: func() *string {
			v, ok := data["enableEventBridge"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableEventBridge"])
		}(),
		EventBridgeAwsAccountId: func() *string {
			v, ok := data["eventBridgeAwsAccountId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsAccountId"])
		}(),
		EventBridgeAwsRegion: func() *string {
			v, ok := data["eventBridgeAwsRegion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsRegion"])
		}(),
	}
}

func (p CreateProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":            p.AccountToken,
		"name":                    p.Name,
		"description":             p.Description,
		"plan":                    p.Plan,
		"currency":                p.Currency,
		"activateRegionName":      p.ActivateRegionName,
		"billingMethodName":       p.BillingMethodName,
		"enableEventBridge":       p.EnableEventBridge,
		"eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
		"eventBridgeAwsRegion":    p.EventBridgeAwsRegion,
	}
}

func (p CreateProjectRequest) Pointer() *CreateProjectRequest {
	return &p
}

type GetProjectRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetProjectRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetProjectRequest{}
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
		*p = GetProjectRequest{}
	} else {
		*p = GetProjectRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
	}
	return nil
}

func NewGetProjectRequestFromJson(data string) (GetProjectRequest, error) {
	req := GetProjectRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetProjectRequest{}, err
	}
	return req, nil
}

func NewGetProjectRequestFromDict(data map[string]interface{}) GetProjectRequest {
	return GetProjectRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
	}
}

func (p GetProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
	}
}

func (p GetProjectRequest) Pointer() *GetProjectRequest {
	return &p
}

type GetProjectTokenRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ProjectName     *string `json:"projectName"`
	AccountToken    *string `json:"accountToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetProjectTokenRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetProjectTokenRequest{}
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
		*p = GetProjectTokenRequest{}
	} else {
		*p = GetProjectTokenRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
	}
	return nil
}

func NewGetProjectTokenRequestFromJson(data string) (GetProjectTokenRequest, error) {
	req := GetProjectTokenRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetProjectTokenRequest{}, err
	}
	return req, nil
}

func NewGetProjectTokenRequestFromDict(data map[string]interface{}) GetProjectTokenRequest {
	return GetProjectTokenRequest{
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
	}
}

func (p GetProjectTokenRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectName":  p.ProjectName,
		"accountToken": p.AccountToken,
	}
}

func (p GetProjectTokenRequest) Pointer() *GetProjectTokenRequest {
	return &p
}

type GetProjectTokenByIdentifierRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountName     *string `json:"accountName"`
	ProjectName     *string `json:"projectName"`
	UserName        *string `json:"userName"`
	Password        *string `json:"password"`
	Otp             *string `json:"otp"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetProjectTokenByIdentifierRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetProjectTokenByIdentifierRequest{}
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
		*p = GetProjectTokenByIdentifierRequest{}
	} else {
		*p = GetProjectTokenByIdentifierRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountName)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["userName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserName)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["otp"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Otp = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Otp = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Otp = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Otp = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Otp = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Otp)
				}
			}
		}
	}
	return nil
}

func NewGetProjectTokenByIdentifierRequestFromJson(data string) (GetProjectTokenByIdentifierRequest, error) {
	req := GetProjectTokenByIdentifierRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetProjectTokenByIdentifierRequest{}, err
	}
	return req, nil
}

func NewGetProjectTokenByIdentifierRequestFromDict(data map[string]interface{}) GetProjectTokenByIdentifierRequest {
	return GetProjectTokenByIdentifierRequest{
		AccountName: func() *string {
			v, ok := data["accountName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountName"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		Otp: func() *string {
			v, ok := data["otp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["otp"])
		}(),
	}
}

func (p GetProjectTokenByIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountName": p.AccountName,
		"projectName": p.ProjectName,
		"userName":    p.UserName,
		"password":    p.Password,
		"otp":         p.Otp,
	}
}

func (p GetProjectTokenByIdentifierRequest) Pointer() *GetProjectTokenByIdentifierRequest {
	return &p
}

type UpdateProjectRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	AccountToken            *string `json:"accountToken"`
	ProjectName             *string `json:"projectName"`
	Description             *string `json:"description"`
	Plan                    *string `json:"plan"`
	BillingMethodName       *string `json:"billingMethodName"`
	EnableEventBridge       *string `json:"enableEventBridge"`
	EventBridgeAwsAccountId *string `json:"eventBridgeAwsAccountId"`
	EventBridgeAwsRegion    *string `json:"eventBridgeAwsRegion"`
	DryRun                  *bool   `json:"dryRun"`
}

func (p *UpdateProjectRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateProjectRequest{}
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
		*p = UpdateProjectRequest{}
	} else {
		*p = UpdateProjectRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["plan"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Plan = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Plan = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Plan = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Plan = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Plan)
				}
			}
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
		if v, ok := d["enableEventBridge"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableEventBridge = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableEventBridge = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableEventBridge = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableEventBridge = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableEventBridge)
				}
			}
		}
		if v, ok := d["eventBridgeAwsAccountId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsAccountId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsAccountId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsAccountId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsAccountId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsAccountId)
				}
			}
		}
		if v, ok := d["eventBridgeAwsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventBridgeAwsRegion = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventBridgeAwsRegion = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventBridgeAwsRegion = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventBridgeAwsRegion = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventBridgeAwsRegion)
				}
			}
		}
	}
	return nil
}

func NewUpdateProjectRequestFromJson(data string) (UpdateProjectRequest, error) {
	req := UpdateProjectRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateProjectRequest{}, err
	}
	return req, nil
}

func NewUpdateProjectRequestFromDict(data map[string]interface{}) UpdateProjectRequest {
	return UpdateProjectRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Plan: func() *string {
			v, ok := data["plan"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["plan"])
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
		EnableEventBridge: func() *string {
			v, ok := data["enableEventBridge"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableEventBridge"])
		}(),
		EventBridgeAwsAccountId: func() *string {
			v, ok := data["eventBridgeAwsAccountId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsAccountId"])
		}(),
		EventBridgeAwsRegion: func() *string {
			v, ok := data["eventBridgeAwsRegion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventBridgeAwsRegion"])
		}(),
	}
}

func (p UpdateProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":            p.AccountToken,
		"projectName":             p.ProjectName,
		"description":             p.Description,
		"plan":                    p.Plan,
		"billingMethodName":       p.BillingMethodName,
		"enableEventBridge":       p.EnableEventBridge,
		"eventBridgeAwsAccountId": p.EventBridgeAwsAccountId,
		"eventBridgeAwsRegion":    p.EventBridgeAwsRegion,
	}
}

func (p UpdateProjectRequest) Pointer() *UpdateProjectRequest {
	return &p
}

type ActivateRegionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	RegionName      *string `json:"regionName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *ActivateRegionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ActivateRegionRequest{}
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
		*p = ActivateRegionRequest{}
	} else {
		*p = ActivateRegionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["regionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RegionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RegionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RegionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RegionName)
				}
			}
		}
	}
	return nil
}

func NewActivateRegionRequestFromJson(data string) (ActivateRegionRequest, error) {
	req := ActivateRegionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ActivateRegionRequest{}, err
	}
	return req, nil
}

func NewActivateRegionRequestFromDict(data map[string]interface{}) ActivateRegionRequest {
	return ActivateRegionRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		RegionName: func() *string {
			v, ok := data["regionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["regionName"])
		}(),
	}
}

func (p ActivateRegionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
		"regionName":   p.RegionName,
	}
}

func (p ActivateRegionRequest) Pointer() *ActivateRegionRequest {
	return &p
}

type WaitActivateRegionRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ProjectName     *string `json:"projectName"`
	RegionName      *string `json:"regionName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *WaitActivateRegionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WaitActivateRegionRequest{}
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
		*p = WaitActivateRegionRequest{}
	} else {
		*p = WaitActivateRegionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["regionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RegionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RegionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RegionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RegionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RegionName)
				}
			}
		}
	}
	return nil
}

func NewWaitActivateRegionRequestFromJson(data string) (WaitActivateRegionRequest, error) {
	req := WaitActivateRegionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WaitActivateRegionRequest{}, err
	}
	return req, nil
}

func NewWaitActivateRegionRequestFromDict(data map[string]interface{}) WaitActivateRegionRequest {
	return WaitActivateRegionRequest{
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		RegionName: func() *string {
			v, ok := data["regionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["regionName"])
		}(),
	}
}

func (p WaitActivateRegionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectName": p.ProjectName,
		"regionName":  p.RegionName,
	}
}

func (p WaitActivateRegionRequest) Pointer() *WaitActivateRegionRequest {
	return &p
}

type DeleteProjectRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DeleteProjectRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteProjectRequest{}
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
		*p = DeleteProjectRequest{}
	} else {
		*p = DeleteProjectRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
	}
	return nil
}

func NewDeleteProjectRequestFromJson(data string) (DeleteProjectRequest, error) {
	req := DeleteProjectRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteProjectRequest{}, err
	}
	return req, nil
}

func NewDeleteProjectRequestFromDict(data map[string]interface{}) DeleteProjectRequest {
	return DeleteProjectRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
	}
}

func (p DeleteProjectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
	}
}

func (p DeleteProjectRequest) Pointer() *DeleteProjectRequest {
	return &p
}

type DescribeBillingMethodsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeBillingMethodsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBillingMethodsRequest{}
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
		*p = DescribeBillingMethodsRequest{}
	} else {
		*p = DescribeBillingMethodsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeBillingMethodsRequestFromJson(data string) (DescribeBillingMethodsRequest, error) {
	req := DescribeBillingMethodsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBillingMethodsRequest{}, err
	}
	return req, nil
}

func NewDescribeBillingMethodsRequestFromDict(data map[string]interface{}) DescribeBillingMethodsRequest {
	return DescribeBillingMethodsRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeBillingMethodsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeBillingMethodsRequest) Pointer() *DescribeBillingMethodsRequest {
	return &p
}

type CreateBillingMethodRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	Description     *string `json:"description"`
	MethodType      *string `json:"methodType"`
	CardCustomerId  *string `json:"cardCustomerId"`
	PartnerId       *string `json:"partnerId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreateBillingMethodRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateBillingMethodRequest{}
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
		*p = CreateBillingMethodRequest{}
	} else {
		*p = CreateBillingMethodRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["methodType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MethodType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MethodType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MethodType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MethodType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MethodType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MethodType)
				}
			}
		}
		if v, ok := d["cardCustomerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CardCustomerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CardCustomerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CardCustomerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CardCustomerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CardCustomerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CardCustomerId)
				}
			}
		}
		if v, ok := d["partnerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PartnerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PartnerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PartnerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PartnerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PartnerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PartnerId)
				}
			}
		}
	}
	return nil
}

func NewCreateBillingMethodRequestFromJson(data string) (CreateBillingMethodRequest, error) {
	req := CreateBillingMethodRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateBillingMethodRequest{}, err
	}
	return req, nil
}

func NewCreateBillingMethodRequestFromDict(data map[string]interface{}) CreateBillingMethodRequest {
	return CreateBillingMethodRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		MethodType: func() *string {
			v, ok := data["methodType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["methodType"])
		}(),
		CardCustomerId: func() *string {
			v, ok := data["cardCustomerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["cardCustomerId"])
		}(),
		PartnerId: func() *string {
			v, ok := data["partnerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["partnerId"])
		}(),
	}
}

func (p CreateBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":   p.AccountToken,
		"description":    p.Description,
		"methodType":     p.MethodType,
		"cardCustomerId": p.CardCustomerId,
		"partnerId":      p.PartnerId,
	}
}

func (p CreateBillingMethodRequest) Pointer() *CreateBillingMethodRequest {
	return &p
}

type GetBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
	DryRun            *bool   `json:"dryRun"`
}

func (p *GetBillingMethodRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBillingMethodRequest{}
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
		*p = GetBillingMethodRequest{}
	} else {
		*p = GetBillingMethodRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
	}
	return nil
}

func NewGetBillingMethodRequestFromJson(data string) (GetBillingMethodRequest, error) {
	req := GetBillingMethodRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBillingMethodRequest{}, err
	}
	return req, nil
}

func NewGetBillingMethodRequestFromDict(data map[string]interface{}) GetBillingMethodRequest {
	return GetBillingMethodRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
	}
}

func (p GetBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
	}
}

func (p GetBillingMethodRequest) Pointer() *GetBillingMethodRequest {
	return &p
}

type UpdateBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
	Description       *string `json:"description"`
	DryRun            *bool   `json:"dryRun"`
}

func (p *UpdateBillingMethodRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateBillingMethodRequest{}
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
		*p = UpdateBillingMethodRequest{}
	} else {
		*p = UpdateBillingMethodRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
	}
	return nil
}

func NewUpdateBillingMethodRequestFromJson(data string) (UpdateBillingMethodRequest, error) {
	req := UpdateBillingMethodRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateBillingMethodRequest{}, err
	}
	return req, nil
}

func NewUpdateBillingMethodRequestFromDict(data map[string]interface{}) UpdateBillingMethodRequest {
	return UpdateBillingMethodRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
	}
}

func (p UpdateBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
		"description":       p.Description,
	}
}

func (p UpdateBillingMethodRequest) Pointer() *UpdateBillingMethodRequest {
	return &p
}

type DeleteBillingMethodRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	AccountToken      *string `json:"accountToken"`
	BillingMethodName *string `json:"billingMethodName"`
	DryRun            *bool   `json:"dryRun"`
}

func (p *DeleteBillingMethodRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteBillingMethodRequest{}
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
		*p = DeleteBillingMethodRequest{}
	} else {
		*p = DeleteBillingMethodRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["billingMethodName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BillingMethodName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BillingMethodName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BillingMethodName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BillingMethodName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BillingMethodName)
				}
			}
		}
	}
	return nil
}

func NewDeleteBillingMethodRequestFromJson(data string) (DeleteBillingMethodRequest, error) {
	req := DeleteBillingMethodRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteBillingMethodRequest{}, err
	}
	return req, nil
}

func NewDeleteBillingMethodRequestFromDict(data map[string]interface{}) DeleteBillingMethodRequest {
	return DeleteBillingMethodRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		BillingMethodName: func() *string {
			v, ok := data["billingMethodName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["billingMethodName"])
		}(),
	}
}

func (p DeleteBillingMethodRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken":      p.AccountToken,
		"billingMethodName": p.BillingMethodName,
	}
}

func (p DeleteBillingMethodRequest) Pointer() *DeleteBillingMethodRequest {
	return &p
}

type DescribeReceiptsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeReceiptsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeReceiptsRequest{}
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
		*p = DescribeReceiptsRequest{}
	} else {
		*p = DescribeReceiptsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeReceiptsRequestFromJson(data string) (DescribeReceiptsRequest, error) {
	req := DescribeReceiptsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeReceiptsRequest{}, err
	}
	return req, nil
}

func NewDescribeReceiptsRequestFromDict(data map[string]interface{}) DescribeReceiptsRequest {
	return DescribeReceiptsRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeReceiptsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"pageToken":    p.PageToken,
		"limit":        p.Limit,
	}
}

func (p DescribeReceiptsRequest) Pointer() *DescribeReceiptsRequest {
	return &p
}

type DescribeBillingsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	AccountToken    *string `json:"accountToken"`
	ProjectName     *string `json:"projectName"`
	Year            *int32  `json:"year"`
	Month           *int32  `json:"month"`
	Region          *string `json:"region"`
	Service         *string `json:"service"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeBillingsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBillingsRequest{}
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
		*p = DescribeBillingsRequest{}
	} else {
		*p = DescribeBillingsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["accountToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccountToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccountToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccountToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccountToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccountToken)
				}
			}
		}
		if v, ok := d["projectName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProjectName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProjectName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProjectName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProjectName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProjectName)
				}
			}
		}
		if v, ok := d["year"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Year)
		}
		if v, ok := d["month"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Month)
		}
		if v, ok := d["region"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Region = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Region = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Region = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Region)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
	}
	return nil
}

func NewDescribeBillingsRequestFromJson(data string) (DescribeBillingsRequest, error) {
	req := DescribeBillingsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBillingsRequest{}, err
	}
	return req, nil
}

func NewDescribeBillingsRequestFromDict(data map[string]interface{}) DescribeBillingsRequest {
	return DescribeBillingsRequest{
		AccountToken: func() *string {
			v, ok := data["accountToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accountToken"])
		}(),
		ProjectName: func() *string {
			v, ok := data["projectName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectName"])
		}(),
		Year: func() *int32 {
			v, ok := data["year"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["year"])
		}(),
		Month: func() *int32 {
			v, ok := data["month"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["month"])
		}(),
		Region: func() *string {
			v, ok := data["region"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["region"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
	}
}

func (p DescribeBillingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"accountToken": p.AccountToken,
		"projectName":  p.ProjectName,
		"year":         p.Year,
		"month":        p.Month,
		"region":       p.Region,
		"service":      p.Service,
	}
}

func (p DescribeBillingsRequest) Pointer() *DescribeBillingsRequest {
	return &p
}

type DescribeDumpProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeDumpProgressesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDumpProgressesRequest{}
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
		*p = DescribeDumpProgressesRequest{}
	} else {
		*p = DescribeDumpProgressesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeDumpProgressesRequestFromJson(data string) (DescribeDumpProgressesRequest, error) {
	req := DescribeDumpProgressesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDumpProgressesRequest{}, err
	}
	return req, nil
}

func NewDescribeDumpProgressesRequestFromDict(data map[string]interface{}) DescribeDumpProgressesRequest {
	return DescribeDumpProgressesRequest{
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeDumpProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeDumpProgressesRequest) Pointer() *DescribeDumpProgressesRequest {
	return &p
}

type GetDumpProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetDumpProgressRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDumpProgressRequest{}
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
		*p = GetDumpProgressRequest{}
	} else {
		*p = GetDumpProgressRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetDumpProgressRequestFromJson(data string) (GetDumpProgressRequest, error) {
	req := GetDumpProgressRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDumpProgressRequest{}, err
	}
	return req, nil
}

func NewGetDumpProgressRequestFromDict(data map[string]interface{}) GetDumpProgressRequest {
	return GetDumpProgressRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetDumpProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetDumpProgressRequest) Pointer() *GetDumpProgressRequest {
	return &p
}

type WaitDumpUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *WaitDumpUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WaitDumpUserDataRequest{}
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
		*p = WaitDumpUserDataRequest{}
	} else {
		*p = WaitDumpUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["microserviceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MicroserviceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MicroserviceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MicroserviceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MicroserviceName)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewWaitDumpUserDataRequestFromJson(data string) (WaitDumpUserDataRequest, error) {
	req := WaitDumpUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WaitDumpUserDataRequest{}, err
	}
	return req, nil
}

func NewWaitDumpUserDataRequestFromDict(data map[string]interface{}) WaitDumpUserDataRequest {
	return WaitDumpUserDataRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		MicroserviceName: func() *string {
			v, ok := data["microserviceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["microserviceName"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p WaitDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitDumpUserDataRequest) Pointer() *WaitDumpUserDataRequest {
	return &p
}

type ArchiveDumpUserDataRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *ArchiveDumpUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ArchiveDumpUserDataRequest{}
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
		*p = ArchiveDumpUserDataRequest{}
	} else {
		*p = ArchiveDumpUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewArchiveDumpUserDataRequestFromJson(data string) (ArchiveDumpUserDataRequest, error) {
	req := ArchiveDumpUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ArchiveDumpUserDataRequest{}, err
	}
	return req, nil
}

func NewArchiveDumpUserDataRequestFromDict(data map[string]interface{}) ArchiveDumpUserDataRequest {
	return ArchiveDumpUserDataRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p ArchiveDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p ArchiveDumpUserDataRequest) Pointer() *ArchiveDumpUserDataRequest {
	return &p
}

type DumpUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DumpUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DumpUserDataRequest{}
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
		*p = DumpUserDataRequest{}
	} else {
		*p = DumpUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewDumpUserDataRequestFromJson(data string) (DumpUserDataRequest, error) {
	req := DumpUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DumpUserDataRequest{}, err
	}
	return req, nil
}

func NewDumpUserDataRequestFromDict(data map[string]interface{}) DumpUserDataRequest {
	return DumpUserDataRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataRequest) Pointer() *DumpUserDataRequest {
	return &p
}

type GetDumpUserDataRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetDumpUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDumpUserDataRequest{}
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
		*p = GetDumpUserDataRequest{}
	} else {
		*p = GetDumpUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetDumpUserDataRequestFromJson(data string) (GetDumpUserDataRequest, error) {
	req := GetDumpUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDumpUserDataRequest{}, err
	}
	return req, nil
}

func NewGetDumpUserDataRequestFromDict(data map[string]interface{}) GetDumpUserDataRequest {
	return GetDumpUserDataRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetDumpUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetDumpUserDataRequest) Pointer() *GetDumpUserDataRequest {
	return &p
}

type DescribeCleanProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeCleanProgressesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeCleanProgressesRequest{}
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
		*p = DescribeCleanProgressesRequest{}
	} else {
		*p = DescribeCleanProgressesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeCleanProgressesRequestFromJson(data string) (DescribeCleanProgressesRequest, error) {
	req := DescribeCleanProgressesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeCleanProgressesRequest{}, err
	}
	return req, nil
}

func NewDescribeCleanProgressesRequestFromDict(data map[string]interface{}) DescribeCleanProgressesRequest {
	return DescribeCleanProgressesRequest{
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeCleanProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeCleanProgressesRequest) Pointer() *DescribeCleanProgressesRequest {
	return &p
}

type GetCleanProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetCleanProgressRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCleanProgressRequest{}
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
		*p = GetCleanProgressRequest{}
	} else {
		*p = GetCleanProgressRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetCleanProgressRequestFromJson(data string) (GetCleanProgressRequest, error) {
	req := GetCleanProgressRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCleanProgressRequest{}, err
	}
	return req, nil
}

func NewGetCleanProgressRequestFromDict(data map[string]interface{}) GetCleanProgressRequest {
	return GetCleanProgressRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetCleanProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetCleanProgressRequest) Pointer() *GetCleanProgressRequest {
	return &p
}

type WaitCleanUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *WaitCleanUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WaitCleanUserDataRequest{}
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
		*p = WaitCleanUserDataRequest{}
	} else {
		*p = WaitCleanUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["microserviceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MicroserviceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MicroserviceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MicroserviceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MicroserviceName)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewWaitCleanUserDataRequestFromJson(data string) (WaitCleanUserDataRequest, error) {
	req := WaitCleanUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WaitCleanUserDataRequest{}, err
	}
	return req, nil
}

func NewWaitCleanUserDataRequestFromDict(data map[string]interface{}) WaitCleanUserDataRequest {
	return WaitCleanUserDataRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		MicroserviceName: func() *string {
			v, ok := data["microserviceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["microserviceName"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p WaitCleanUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitCleanUserDataRequest) Pointer() *WaitCleanUserDataRequest {
	return &p
}

type CleanUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *CleanUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CleanUserDataRequest{}
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
		*p = CleanUserDataRequest{}
	} else {
		*p = CleanUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewCleanUserDataRequestFromJson(data string) (CleanUserDataRequest, error) {
	req := CleanUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CleanUserDataRequest{}, err
	}
	return req, nil
}

func NewCleanUserDataRequestFromDict(data map[string]interface{}) CleanUserDataRequest {
	return CleanUserDataRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p CleanUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataRequest) Pointer() *CleanUserDataRequest {
	return &p
}

type DescribeImportProgressesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeImportProgressesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeImportProgressesRequest{}
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
		*p = DescribeImportProgressesRequest{}
	} else {
		*p = DescribeImportProgressesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeImportProgressesRequestFromJson(data string) (DescribeImportProgressesRequest, error) {
	req := DescribeImportProgressesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeImportProgressesRequest{}, err
	}
	return req, nil
}

func NewDescribeImportProgressesRequestFromDict(data map[string]interface{}) DescribeImportProgressesRequest {
	return DescribeImportProgressesRequest{
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeImportProgressesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeImportProgressesRequest) Pointer() *DescribeImportProgressesRequest {
	return &p
}

type GetImportProgressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetImportProgressRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetImportProgressRequest{}
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
		*p = GetImportProgressRequest{}
	} else {
		*p = GetImportProgressRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
	}
	return nil
}

func NewGetImportProgressRequestFromJson(data string) (GetImportProgressRequest, error) {
	req := GetImportProgressRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetImportProgressRequest{}, err
	}
	return req, nil
}

func NewGetImportProgressRequestFromDict(data map[string]interface{}) GetImportProgressRequest {
	return GetImportProgressRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
	}
}

func (p GetImportProgressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
	}
}

func (p GetImportProgressRequest) Pointer() *GetImportProgressRequest {
	return &p
}

type WaitImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	TransactionId      *string `json:"transactionId"`
	UserId             *string `json:"userId"`
	MicroserviceName   *string `json:"microserviceName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *WaitImportUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WaitImportUserDataRequest{}
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
		*p = WaitImportUserDataRequest{}
	} else {
		*p = WaitImportUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["microserviceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MicroserviceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MicroserviceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MicroserviceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MicroserviceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MicroserviceName)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewWaitImportUserDataRequestFromJson(data string) (WaitImportUserDataRequest, error) {
	req := WaitImportUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WaitImportUserDataRequest{}, err
	}
	return req, nil
}

func NewWaitImportUserDataRequestFromDict(data map[string]interface{}) WaitImportUserDataRequest {
	return WaitImportUserDataRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		MicroserviceName: func() *string {
			v, ok := data["microserviceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["microserviceName"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p WaitImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId":    p.TransactionId,
		"userId":           p.UserId,
		"microserviceName": p.MicroserviceName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p WaitImportUserDataRequest) Pointer() *WaitImportUserDataRequest {
	return &p
}

type PrepareImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PrepareImportUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrepareImportUserDataRequest{}
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
		*p = PrepareImportUserDataRequest{}
	} else {
		*p = PrepareImportUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewPrepareImportUserDataRequestFromJson(data string) (PrepareImportUserDataRequest, error) {
	req := PrepareImportUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PrepareImportUserDataRequest{}, err
	}
	return req, nil
}

func NewPrepareImportUserDataRequestFromDict(data map[string]interface{}) PrepareImportUserDataRequest {
	return PrepareImportUserDataRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p PrepareImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataRequest) Pointer() *PrepareImportUserDataRequest {
	return &p
}

type ImportUserDataRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *ImportUserDataRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ImportUserDataRequest{}
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
		*p = ImportUserDataRequest{}
	} else {
		*p = ImportUserDataRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UploadToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UploadToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UploadToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UploadToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UploadToken)
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TimeOffsetToken)
				}
			}
		}
	}
	return nil
}

func NewImportUserDataRequestFromJson(data string) (ImportUserDataRequest, error) {
	req := ImportUserDataRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ImportUserDataRequest{}, err
	}
	return req, nil
}

func NewImportUserDataRequestFromDict(data map[string]interface{}) ImportUserDataRequest {
	return ImportUserDataRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p ImportUserDataRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataRequest) Pointer() *ImportUserDataRequest {
	return &p
}

type DescribeImportErrorLogsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeImportErrorLogsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeImportErrorLogsRequest{}
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
		*p = DescribeImportErrorLogsRequest{}
	} else {
		*p = DescribeImportErrorLogsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PageToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PageToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PageToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PageToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeImportErrorLogsRequestFromJson(data string) (DescribeImportErrorLogsRequest, error) {
	req := DescribeImportErrorLogsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeImportErrorLogsRequest{}, err
	}
	return req, nil
}

func NewDescribeImportErrorLogsRequestFromDict(data map[string]interface{}) DescribeImportErrorLogsRequest {
	return DescribeImportErrorLogsRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeImportErrorLogsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeImportErrorLogsRequest) Pointer() *DescribeImportErrorLogsRequest {
	return &p
}

type GetImportErrorLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	TransactionId   *string `json:"transactionId"`
	ErrorLogName    *string `json:"errorLogName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetImportErrorLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetImportErrorLogRequest{}
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
		*p = GetImportErrorLogRequest{}
	} else {
		*p = GetImportErrorLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["errorLogName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ErrorLogName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ErrorLogName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ErrorLogName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ErrorLogName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ErrorLogName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ErrorLogName)
				}
			}
		}
	}
	return nil
}

func NewGetImportErrorLogRequestFromJson(data string) (GetImportErrorLogRequest, error) {
	req := GetImportErrorLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetImportErrorLogRequest{}, err
	}
	return req, nil
}

func NewGetImportErrorLogRequestFromDict(data map[string]interface{}) GetImportErrorLogRequest {
	return GetImportErrorLogRequest{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		ErrorLogName: func() *string {
			v, ok := data["errorLogName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["errorLogName"])
		}(),
	}
}

func (p GetImportErrorLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"transactionId": p.TransactionId,
		"errorLogName":  p.ErrorLogName,
	}
}

func (p GetImportErrorLogRequest) Pointer() *GetImportErrorLogRequest {
	return &p
}
