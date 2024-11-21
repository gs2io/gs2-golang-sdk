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

package identifier

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeUsersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeUsersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeUsersRequest{}
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
		*p = DescribeUsersRequest{}
	} else {
		*p = DescribeUsersRequest{}
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

func NewDescribeUsersRequestFromJson(data string) (DescribeUsersRequest, error) {
	req := DescribeUsersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeUsersRequest{}, err
	}
	return req, nil
}

func NewDescribeUsersRequestFromDict(data map[string]interface{}) DescribeUsersRequest {
	return DescribeUsersRequest{
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

func (p DescribeUsersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeUsersRequest) Pointer() *DescribeUsersRequest {
	return &p
}

type CreateUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreateUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateUserRequest{}
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
		*p = CreateUserRequest{}
	} else {
		*p = CreateUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewCreateUserRequestFromJson(data string) (CreateUserRequest, error) {
	req := CreateUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateUserRequest{}, err
	}
	return req, nil
}

func NewCreateUserRequestFromDict(data map[string]interface{}) CreateUserRequest {
	return CreateUserRequest{
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
	}
}

func (p CreateUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
	}
}

func (p CreateUserRequest) Pointer() *CreateUserRequest {
	return &p
}

type UpdateUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	Description     *string `json:"description"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *UpdateUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateUserRequest{}
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
		*p = UpdateUserRequest{}
	} else {
		*p = UpdateUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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

func NewUpdateUserRequestFromJson(data string) (UpdateUserRequest, error) {
	req := UpdateUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateUserRequest{}, err
	}
	return req, nil
}

func NewUpdateUserRequestFromDict(data map[string]interface{}) UpdateUserRequest {
	return UpdateUserRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
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

func (p UpdateUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":    p.UserName,
		"description": p.Description,
	}
}

func (p UpdateUserRequest) Pointer() *UpdateUserRequest {
	return &p
}

type GetUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetUserRequest{}
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
		*p = GetUserRequest{}
	} else {
		*p = GetUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewGetUserRequestFromJson(data string) (GetUserRequest, error) {
	req := GetUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetUserRequest{}, err
	}
	return req, nil
}

func NewGetUserRequestFromDict(data map[string]interface{}) GetUserRequest {
	return GetUserRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p GetUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p GetUserRequest) Pointer() *GetUserRequest {
	return &p
}

type DeleteUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DeleteUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteUserRequest{}
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
		*p = DeleteUserRequest{}
	} else {
		*p = DeleteUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewDeleteUserRequestFromJson(data string) (DeleteUserRequest, error) {
	req := DeleteUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteUserRequest{}, err
	}
	return req, nil
}

func NewDeleteUserRequestFromDict(data map[string]interface{}) DeleteUserRequest {
	return DeleteUserRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p DeleteUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p DeleteUserRequest) Pointer() *DeleteUserRequest {
	return &p
}

type DescribeSecurityPoliciesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeSecurityPoliciesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSecurityPoliciesRequest{}
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
		*p = DescribeSecurityPoliciesRequest{}
	} else {
		*p = DescribeSecurityPoliciesRequest{}
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

func NewDescribeSecurityPoliciesRequestFromJson(data string) (DescribeSecurityPoliciesRequest, error) {
	req := DescribeSecurityPoliciesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSecurityPoliciesRequest{}, err
	}
	return req, nil
}

func NewDescribeSecurityPoliciesRequestFromDict(data map[string]interface{}) DescribeSecurityPoliciesRequest {
	return DescribeSecurityPoliciesRequest{
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

func (p DescribeSecurityPoliciesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeSecurityPoliciesRequest) Pointer() *DescribeSecurityPoliciesRequest {
	return &p
}

type DescribeCommonSecurityPoliciesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeCommonSecurityPoliciesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeCommonSecurityPoliciesRequest{}
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
		*p = DescribeCommonSecurityPoliciesRequest{}
	} else {
		*p = DescribeCommonSecurityPoliciesRequest{}
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

func NewDescribeCommonSecurityPoliciesRequestFromJson(data string) (DescribeCommonSecurityPoliciesRequest, error) {
	req := DescribeCommonSecurityPoliciesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeCommonSecurityPoliciesRequest{}, err
	}
	return req, nil
}

func NewDescribeCommonSecurityPoliciesRequestFromDict(data map[string]interface{}) DescribeCommonSecurityPoliciesRequest {
	return DescribeCommonSecurityPoliciesRequest{
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

func (p DescribeCommonSecurityPoliciesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeCommonSecurityPoliciesRequest) Pointer() *DescribeCommonSecurityPoliciesRequest {
	return &p
}

type CreateSecurityPolicyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Policy          *string `json:"policy"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreateSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateSecurityPolicyRequest{}
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
		*p = CreateSecurityPolicyRequest{}
	} else {
		*p = CreateSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["policy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Policy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Policy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Policy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Policy)
				}
			}
		}
	}
	return nil
}

func NewCreateSecurityPolicyRequestFromJson(data string) (CreateSecurityPolicyRequest, error) {
	req := CreateSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewCreateSecurityPolicyRequestFromDict(data map[string]interface{}) CreateSecurityPolicyRequest {
	return CreateSecurityPolicyRequest{
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
		Policy: func() *string {
			v, ok := data["policy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["policy"])
		}(),
	}
}

func (p CreateSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"policy":      p.Policy,
	}
}

func (p CreateSecurityPolicyRequest) Pointer() *CreateSecurityPolicyRequest {
	return &p
}

type UpdateSecurityPolicyRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	SecurityPolicyName *string `json:"securityPolicyName"`
	Description        *string `json:"description"`
	Policy             *string `json:"policy"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *UpdateSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateSecurityPolicyRequest{}
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
		*p = UpdateSecurityPolicyRequest{}
	} else {
		*p = UpdateSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["securityPolicyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyName)
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
		if v, ok := d["policy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Policy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Policy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Policy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Policy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Policy)
				}
			}
		}
	}
	return nil
}

func NewUpdateSecurityPolicyRequestFromJson(data string) (UpdateSecurityPolicyRequest, error) {
	req := UpdateSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewUpdateSecurityPolicyRequestFromDict(data map[string]interface{}) UpdateSecurityPolicyRequest {
	return UpdateSecurityPolicyRequest{
		SecurityPolicyName: func() *string {
			v, ok := data["securityPolicyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["securityPolicyName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Policy: func() *string {
			v, ok := data["policy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["policy"])
		}(),
	}
}

func (p UpdateSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"securityPolicyName": p.SecurityPolicyName,
		"description":        p.Description,
		"policy":             p.Policy,
	}
}

func (p UpdateSecurityPolicyRequest) Pointer() *UpdateSecurityPolicyRequest {
	return &p
}

type GetSecurityPolicyRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	SecurityPolicyName *string `json:"securityPolicyName"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *GetSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSecurityPolicyRequest{}
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
		*p = GetSecurityPolicyRequest{}
	} else {
		*p = GetSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["securityPolicyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyName)
				}
			}
		}
	}
	return nil
}

func NewGetSecurityPolicyRequestFromJson(data string) (GetSecurityPolicyRequest, error) {
	req := GetSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewGetSecurityPolicyRequestFromDict(data map[string]interface{}) GetSecurityPolicyRequest {
	return GetSecurityPolicyRequest{
		SecurityPolicyName: func() *string {
			v, ok := data["securityPolicyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["securityPolicyName"])
		}(),
	}
}

func (p GetSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"securityPolicyName": p.SecurityPolicyName,
	}
}

func (p GetSecurityPolicyRequest) Pointer() *GetSecurityPolicyRequest {
	return &p
}

type DeleteSecurityPolicyRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	SecurityPolicyName *string `json:"securityPolicyName"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSecurityPolicyRequest{}
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
		*p = DeleteSecurityPolicyRequest{}
	} else {
		*p = DeleteSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["securityPolicyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSecurityPolicyRequestFromJson(data string) (DeleteSecurityPolicyRequest, error) {
	req := DeleteSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewDeleteSecurityPolicyRequestFromDict(data map[string]interface{}) DeleteSecurityPolicyRequest {
	return DeleteSecurityPolicyRequest{
		SecurityPolicyName: func() *string {
			v, ok := data["securityPolicyName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["securityPolicyName"])
		}(),
	}
}

func (p DeleteSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"securityPolicyName": p.SecurityPolicyName,
	}
}

func (p DeleteSecurityPolicyRequest) Pointer() *DeleteSecurityPolicyRequest {
	return &p
}

type DescribeIdentifiersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeIdentifiersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeIdentifiersRequest{}
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
		*p = DescribeIdentifiersRequest{}
	} else {
		*p = DescribeIdentifiersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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

func NewDescribeIdentifiersRequestFromJson(data string) (DescribeIdentifiersRequest, error) {
	req := DescribeIdentifiersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeIdentifiersRequest{}, err
	}
	return req, nil
}

func NewDescribeIdentifiersRequestFromDict(data map[string]interface{}) DescribeIdentifiersRequest {
	return DescribeIdentifiersRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
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

func (p DescribeIdentifiersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":  p.UserName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeIdentifiersRequest) Pointer() *DescribeIdentifiersRequest {
	return &p
}

type CreateIdentifierRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreateIdentifierRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateIdentifierRequest{}
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
		*p = CreateIdentifierRequest{}
	} else {
		*p = CreateIdentifierRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewCreateIdentifierRequestFromJson(data string) (CreateIdentifierRequest, error) {
	req := CreateIdentifierRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateIdentifierRequest{}, err
	}
	return req, nil
}

func NewCreateIdentifierRequestFromDict(data map[string]interface{}) CreateIdentifierRequest {
	return CreateIdentifierRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p CreateIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p CreateIdentifierRequest) Pointer() *CreateIdentifierRequest {
	return &p
}

type GetIdentifierRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	ClientId        *string `json:"clientId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetIdentifierRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetIdentifierRequest{}
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
		*p = GetIdentifierRequest{}
	} else {
		*p = GetIdentifierRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
	}
	return nil
}

func NewGetIdentifierRequestFromJson(data string) (GetIdentifierRequest, error) {
	req := GetIdentifierRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetIdentifierRequest{}, err
	}
	return req, nil
}

func NewGetIdentifierRequestFromDict(data map[string]interface{}) GetIdentifierRequest {
	return GetIdentifierRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
	}
}

func (p GetIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
		"clientId": p.ClientId,
	}
}

func (p GetIdentifierRequest) Pointer() *GetIdentifierRequest {
	return &p
}

type DeleteIdentifierRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	ClientId        *string `json:"clientId"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DeleteIdentifierRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteIdentifierRequest{}
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
		*p = DeleteIdentifierRequest{}
	} else {
		*p = DeleteIdentifierRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
	}
	return nil
}

func NewDeleteIdentifierRequestFromJson(data string) (DeleteIdentifierRequest, error) {
	req := DeleteIdentifierRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteIdentifierRequest{}, err
	}
	return req, nil
}

func NewDeleteIdentifierRequestFromDict(data map[string]interface{}) DeleteIdentifierRequest {
	return DeleteIdentifierRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
	}
}

func (p DeleteIdentifierRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
		"clientId": p.ClientId,
	}
}

func (p DeleteIdentifierRequest) Pointer() *DeleteIdentifierRequest {
	return &p
}

type DescribeAttachedGuardsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ClientId        *string `json:"clientId"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeAttachedGuardsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAttachedGuardsRequest{}
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
		*p = DescribeAttachedGuardsRequest{}
	} else {
		*p = DescribeAttachedGuardsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
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
	}
	return nil
}

func NewDescribeAttachedGuardsRequestFromJson(data string) (DescribeAttachedGuardsRequest, error) {
	req := DescribeAttachedGuardsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAttachedGuardsRequest{}, err
	}
	return req, nil
}

func NewDescribeAttachedGuardsRequestFromDict(data map[string]interface{}) DescribeAttachedGuardsRequest {
	return DescribeAttachedGuardsRequest{
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p DescribeAttachedGuardsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"clientId": p.ClientId,
		"userName": p.UserName,
	}
}

func (p DescribeAttachedGuardsRequest) Pointer() *DescribeAttachedGuardsRequest {
	return &p
}

type AttachGuardRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	UserName         *string `json:"userName"`
	ClientId         *string `json:"clientId"`
	GuardNamespaceId *string `json:"guardNamespaceId"`
	DryRun           *bool   `json:"dryRun"`
}

func (p *AttachGuardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AttachGuardRequest{}
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
		*p = AttachGuardRequest{}
	} else {
		*p = AttachGuardRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
		if v, ok := d["guardNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuardNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuardNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuardNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuardNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuardNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuardNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewAttachGuardRequestFromJson(data string) (AttachGuardRequest, error) {
	req := AttachGuardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AttachGuardRequest{}, err
	}
	return req, nil
}

func NewAttachGuardRequestFromDict(data map[string]interface{}) AttachGuardRequest {
	return AttachGuardRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
		GuardNamespaceId: func() *string {
			v, ok := data["guardNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guardNamespaceId"])
		}(),
	}
}

func (p AttachGuardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":         p.UserName,
		"clientId":         p.ClientId,
		"guardNamespaceId": p.GuardNamespaceId,
	}
}

func (p AttachGuardRequest) Pointer() *AttachGuardRequest {
	return &p
}

type DetachGuardRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	UserName         *string `json:"userName"`
	ClientId         *string `json:"clientId"`
	GuardNamespaceId *string `json:"guardNamespaceId"`
	DryRun           *bool   `json:"dryRun"`
}

func (p *DetachGuardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DetachGuardRequest{}
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
		*p = DetachGuardRequest{}
	} else {
		*p = DetachGuardRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
		if v, ok := d["guardNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuardNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuardNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuardNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuardNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuardNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuardNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewDetachGuardRequestFromJson(data string) (DetachGuardRequest, error) {
	req := DetachGuardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DetachGuardRequest{}, err
	}
	return req, nil
}

func NewDetachGuardRequestFromDict(data map[string]interface{}) DetachGuardRequest {
	return DetachGuardRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
		GuardNamespaceId: func() *string {
			v, ok := data["guardNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guardNamespaceId"])
		}(),
	}
}

func (p DetachGuardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":         p.UserName,
		"clientId":         p.ClientId,
		"guardNamespaceId": p.GuardNamespaceId,
	}
}

func (p DetachGuardRequest) Pointer() *DetachGuardRequest {
	return &p
}

type CreatePasswordRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	Password        *string `json:"password"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *CreatePasswordRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreatePasswordRequest{}
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
		*p = CreatePasswordRequest{}
	} else {
		*p = CreatePasswordRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewCreatePasswordRequestFromJson(data string) (CreatePasswordRequest, error) {
	req := CreatePasswordRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreatePasswordRequest{}, err
	}
	return req, nil
}

func NewCreatePasswordRequestFromDict(data map[string]interface{}) CreatePasswordRequest {
	return CreatePasswordRequest{
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
	}
}

func (p CreatePasswordRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
		"password": p.Password,
	}
}

func (p CreatePasswordRequest) Pointer() *CreatePasswordRequest {
	return &p
}

type GetPasswordRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetPasswordRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetPasswordRequest{}
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
		*p = GetPasswordRequest{}
	} else {
		*p = GetPasswordRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewGetPasswordRequestFromJson(data string) (GetPasswordRequest, error) {
	req := GetPasswordRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetPasswordRequest{}, err
	}
	return req, nil
}

func NewGetPasswordRequestFromDict(data map[string]interface{}) GetPasswordRequest {
	return GetPasswordRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p GetPasswordRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p GetPasswordRequest) Pointer() *GetPasswordRequest {
	return &p
}

type EnableMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
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
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p EnableMfaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p EnableMfaRequest) Pointer() *EnableMfaRequest {
	return &p
}

type ChallengeMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
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
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
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
		"userName": p.UserName,
		"passcode": p.Passcode,
	}
}

func (p ChallengeMfaRequest) Pointer() *ChallengeMfaRequest {
	return &p
}

type DisableMfaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
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
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p DisableMfaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p DisableMfaRequest) Pointer() *DisableMfaRequest {
	return &p
}

type DeletePasswordRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DeletePasswordRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeletePasswordRequest{}
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
		*p = DeletePasswordRequest{}
	} else {
		*p = DeletePasswordRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewDeletePasswordRequestFromJson(data string) (DeletePasswordRequest, error) {
	req := DeletePasswordRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeletePasswordRequest{}, err
	}
	return req, nil
}

func NewDeletePasswordRequestFromDict(data map[string]interface{}) DeletePasswordRequest {
	return DeletePasswordRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p DeletePasswordRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p DeletePasswordRequest) Pointer() *DeletePasswordRequest {
	return &p
}

type GetHasSecurityPolicyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetHasSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetHasSecurityPolicyRequest{}
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
		*p = GetHasSecurityPolicyRequest{}
	} else {
		*p = GetHasSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
	}
	return nil
}

func NewGetHasSecurityPolicyRequestFromJson(data string) (GetHasSecurityPolicyRequest, error) {
	req := GetHasSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetHasSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewGetHasSecurityPolicyRequestFromDict(data map[string]interface{}) GetHasSecurityPolicyRequest {
	return GetHasSecurityPolicyRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
	}
}

func (p GetHasSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
	}
}

func (p GetHasSecurityPolicyRequest) Pointer() *GetHasSecurityPolicyRequest {
	return &p
}

type AttachSecurityPolicyRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	UserName         *string `json:"userName"`
	SecurityPolicyId *string `json:"securityPolicyId"`
	DryRun           *bool   `json:"dryRun"`
}

func (p *AttachSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AttachSecurityPolicyRequest{}
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
		*p = AttachSecurityPolicyRequest{}
	} else {
		*p = AttachSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["securityPolicyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyId)
				}
			}
		}
	}
	return nil
}

func NewAttachSecurityPolicyRequestFromJson(data string) (AttachSecurityPolicyRequest, error) {
	req := AttachSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AttachSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewAttachSecurityPolicyRequestFromDict(data map[string]interface{}) AttachSecurityPolicyRequest {
	return AttachSecurityPolicyRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		SecurityPolicyId: func() *string {
			v, ok := data["securityPolicyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["securityPolicyId"])
		}(),
	}
}

func (p AttachSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":         p.UserName,
		"securityPolicyId": p.SecurityPolicyId,
	}
}

func (p AttachSecurityPolicyRequest) Pointer() *AttachSecurityPolicyRequest {
	return &p
}

type DetachSecurityPolicyRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	UserName         *string `json:"userName"`
	SecurityPolicyId *string `json:"securityPolicyId"`
	DryRun           *bool   `json:"dryRun"`
}

func (p *DetachSecurityPolicyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DetachSecurityPolicyRequest{}
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
		*p = DetachSecurityPolicyRequest{}
	} else {
		*p = DetachSecurityPolicyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["securityPolicyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SecurityPolicyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SecurityPolicyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SecurityPolicyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SecurityPolicyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SecurityPolicyId)
				}
			}
		}
	}
	return nil
}

func NewDetachSecurityPolicyRequestFromJson(data string) (DetachSecurityPolicyRequest, error) {
	req := DetachSecurityPolicyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DetachSecurityPolicyRequest{}, err
	}
	return req, nil
}

func NewDetachSecurityPolicyRequestFromDict(data map[string]interface{}) DetachSecurityPolicyRequest {
	return DetachSecurityPolicyRequest{
		UserName: func() *string {
			v, ok := data["userName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userName"])
		}(),
		SecurityPolicyId: func() *string {
			v, ok := data["securityPolicyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["securityPolicyId"])
		}(),
	}
}

func (p DetachSecurityPolicyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName":         p.UserName,
		"securityPolicyId": p.SecurityPolicyId,
	}
}

func (p DetachSecurityPolicyRequest) Pointer() *DetachSecurityPolicyRequest {
	return &p
}

type LoginRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	ClientId        *string `json:"clientId"`
	ClientSecret    *string `json:"clientSecret"`
	DryRun          *bool   `json:"dryRun"`
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
		if v, ok := d["clientId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientId)
				}
			}
		}
		if v, ok := d["clientSecret"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClientSecret = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClientSecret = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClientSecret = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClientSecret = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClientSecret = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClientSecret)
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
		ClientId: func() *string {
			v, ok := data["clientId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientId"])
		}(),
		ClientSecret: func() *string {
			v, ok := data["clientSecret"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clientSecret"])
		}(),
	}
}

func (p LoginRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"clientId":     p.ClientId,
		"clientSecret": p.ClientSecret,
	}
}

func (p LoginRequest) Pointer() *LoginRequest {
	return &p
}

type LoginByUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserName        *string `json:"userName"`
	Password        *string `json:"password"`
	Otp             *string `json:"otp"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *LoginByUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LoginByUserRequest{}
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
		*p = LoginByUserRequest{}
	} else {
		*p = LoginByUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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

func NewLoginByUserRequestFromJson(data string) (LoginByUserRequest, error) {
	req := LoginByUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return LoginByUserRequest{}, err
	}
	return req, nil
}

func NewLoginByUserRequestFromDict(data map[string]interface{}) LoginByUserRequest {
	return LoginByUserRequest{
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

func (p LoginByUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userName": p.UserName,
		"password": p.Password,
		"otp":      p.Otp,
	}
}

func (p LoginByUserRequest) Pointer() *LoginByUserRequest {
	return &p
}
