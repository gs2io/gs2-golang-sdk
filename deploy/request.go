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

package deploy

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeStacksRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeStacksRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStacksRequest{}
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
		*p = DescribeStacksRequest{}
	} else {
		*p = DescribeStacksRequest{}
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

func NewDescribeStacksRequestFromJson(data string) (DescribeStacksRequest, error) {
	req := DescribeStacksRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStacksRequest{}, err
	}
	return req, nil
}

func NewDescribeStacksRequestFromDict(data map[string]interface{}) DescribeStacksRequest {
	return DescribeStacksRequest{
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

func (p DescribeStacksRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeStacksRequest) Pointer() *DescribeStacksRequest {
	return &p
}

type CreateStackRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Template        *string `json:"template"`
}

func (p *CreateStackRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateStackRequest{}
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
		*p = CreateStackRequest{}
	} else {
		*p = CreateStackRequest{}
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
		if v, ok := d["template"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Template = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Template = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Template = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Template)
				}
			}
		}
	}
	return nil
}

func NewCreateStackRequestFromJson(data string) (CreateStackRequest, error) {
	req := CreateStackRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateStackRequest{}, err
	}
	return req, nil
}

func NewCreateStackRequestFromDict(data map[string]interface{}) CreateStackRequest {
	return CreateStackRequest{
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
		Template: func() *string {
			v, ok := data["template"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["template"])
		}(),
	}
}

func (p CreateStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"template":    p.Template,
	}
}

func (p CreateStackRequest) Pointer() *CreateStackRequest {
	return &p
}

type CreateStackFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	Name            *string                `json:"name"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *CreateStackFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateStackFromGitHubRequest{}
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
		*p = CreateStackFromGitHubRequest{}
	} else {
		*p = CreateStackFromGitHubRequest{}
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
		if v, ok := d["checkoutSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CheckoutSetting)
		}
	}
	return nil
}

func NewCreateStackFromGitHubRequestFromJson(data string) (CreateStackFromGitHubRequest, error) {
	req := CreateStackFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateStackFromGitHubRequest{}, err
	}
	return req, nil
}

func NewCreateStackFromGitHubRequestFromDict(data map[string]interface{}) CreateStackFromGitHubRequest {
	return CreateStackFromGitHubRequest{
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
		CheckoutSetting: func() *GitHubCheckoutSetting {
			v, ok := data["checkoutSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer()
		}(),
	}
}

func (p CreateStackFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"checkoutSetting": func() map[string]interface{} {
			if p.CheckoutSetting == nil {
				return nil
			}
			return p.CheckoutSetting.ToDict()
		}(),
	}
}

func (p CreateStackFromGitHubRequest) Pointer() *CreateStackFromGitHubRequest {
	return &p
}

type ValidateRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	Template        *string `json:"template"`
}

func (p *ValidateRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ValidateRequest{}
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
		*p = ValidateRequest{}
	} else {
		*p = ValidateRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["template"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Template = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Template = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Template = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Template)
				}
			}
		}
	}
	return nil
}

func NewValidateRequestFromJson(data string) (ValidateRequest, error) {
	req := ValidateRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ValidateRequest{}, err
	}
	return req, nil
}

func NewValidateRequestFromDict(data map[string]interface{}) ValidateRequest {
	return ValidateRequest{
		Template: func() *string {
			v, ok := data["template"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["template"])
		}(),
	}
}

func (p ValidateRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"template": p.Template,
	}
}

func (p ValidateRequest) Pointer() *ValidateRequest {
	return &p
}

type GetStackStatusRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *GetStackStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStackStatusRequest{}
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
		*p = GetStackStatusRequest{}
	} else {
		*p = GetStackStatusRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewGetStackStatusRequestFromJson(data string) (GetStackStatusRequest, error) {
	req := GetStackStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStackStatusRequest{}, err
	}
	return req, nil
}

func NewGetStackStatusRequestFromDict(data map[string]interface{}) GetStackStatusRequest {
	return GetStackStatusRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p GetStackStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p GetStackStatusRequest) Pointer() *GetStackStatusRequest {
	return &p
}

type GetStackRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *GetStackRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStackRequest{}
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
		*p = GetStackRequest{}
	} else {
		*p = GetStackRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewGetStackRequestFromJson(data string) (GetStackRequest, error) {
	req := GetStackRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStackRequest{}, err
	}
	return req, nil
}

func NewGetStackRequestFromDict(data map[string]interface{}) GetStackRequest {
	return GetStackRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p GetStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p GetStackRequest) Pointer() *GetStackRequest {
	return &p
}

type UpdateStackRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	Description     *string `json:"description"`
	Template        *string `json:"template"`
}

func (p *UpdateStackRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStackRequest{}
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
		*p = UpdateStackRequest{}
	} else {
		*p = UpdateStackRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
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
		if v, ok := d["template"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Template = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Template = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Template = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Template)
				}
			}
		}
	}
	return nil
}

func NewUpdateStackRequestFromJson(data string) (UpdateStackRequest, error) {
	req := UpdateStackRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStackRequest{}, err
	}
	return req, nil
}

func NewUpdateStackRequestFromDict(data map[string]interface{}) UpdateStackRequest {
	return UpdateStackRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Template: func() *string {
			v, ok := data["template"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["template"])
		}(),
	}
}

func (p UpdateStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":   p.StackName,
		"description": p.Description,
		"template":    p.Template,
	}
}

func (p UpdateStackRequest) Pointer() *UpdateStackRequest {
	return &p
}

type ChangeSetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	Template        *string `json:"template"`
}

func (p *ChangeSetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ChangeSetRequest{}
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
		*p = ChangeSetRequest{}
	} else {
		*p = ChangeSetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
		if v, ok := d["template"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Template = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Template = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Template = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Template)
				}
			}
		}
	}
	return nil
}

func NewChangeSetRequestFromJson(data string) (ChangeSetRequest, error) {
	req := ChangeSetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ChangeSetRequest{}, err
	}
	return req, nil
}

func NewChangeSetRequestFromDict(data map[string]interface{}) ChangeSetRequest {
	return ChangeSetRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		Template: func() *string {
			v, ok := data["template"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["template"])
		}(),
	}
}

func (p ChangeSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"template":  p.Template,
	}
}

func (p ChangeSetRequest) Pointer() *ChangeSetRequest {
	return &p
}

type UpdateStackFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	StackName       *string                `json:"stackName"`
	Description     *string                `json:"description"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateStackFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStackFromGitHubRequest{}
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
		*p = UpdateStackFromGitHubRequest{}
	} else {
		*p = UpdateStackFromGitHubRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
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
		if v, ok := d["checkoutSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CheckoutSetting)
		}
	}
	return nil
}

func NewUpdateStackFromGitHubRequestFromJson(data string) (UpdateStackFromGitHubRequest, error) {
	req := UpdateStackFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStackFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateStackFromGitHubRequestFromDict(data map[string]interface{}) UpdateStackFromGitHubRequest {
	return UpdateStackFromGitHubRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		CheckoutSetting: func() *GitHubCheckoutSetting {
			v, ok := data["checkoutSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer()
		}(),
	}
}

func (p UpdateStackFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":   p.StackName,
		"description": p.Description,
		"checkoutSetting": func() map[string]interface{} {
			if p.CheckoutSetting == nil {
				return nil
			}
			return p.CheckoutSetting.ToDict()
		}(),
	}
}

func (p UpdateStackFromGitHubRequest) Pointer() *UpdateStackFromGitHubRequest {
	return &p
}

type DeleteStackRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *DeleteStackRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStackRequest{}
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
		*p = DeleteStackRequest{}
	} else {
		*p = DeleteStackRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewDeleteStackRequestFromJson(data string) (DeleteStackRequest, error) {
	req := DeleteStackRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStackRequest{}, err
	}
	return req, nil
}

func NewDeleteStackRequestFromDict(data map[string]interface{}) DeleteStackRequest {
	return DeleteStackRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p DeleteStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackRequest) Pointer() *DeleteStackRequest {
	return &p
}

type ForceDeleteStackRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *ForceDeleteStackRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ForceDeleteStackRequest{}
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
		*p = ForceDeleteStackRequest{}
	} else {
		*p = ForceDeleteStackRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewForceDeleteStackRequestFromJson(data string) (ForceDeleteStackRequest, error) {
	req := ForceDeleteStackRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ForceDeleteStackRequest{}, err
	}
	return req, nil
}

func NewForceDeleteStackRequestFromDict(data map[string]interface{}) ForceDeleteStackRequest {
	return ForceDeleteStackRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p ForceDeleteStackRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p ForceDeleteStackRequest) Pointer() *ForceDeleteStackRequest {
	return &p
}

type DeleteStackResourcesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *DeleteStackResourcesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStackResourcesRequest{}
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
		*p = DeleteStackResourcesRequest{}
	} else {
		*p = DeleteStackResourcesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewDeleteStackResourcesRequestFromJson(data string) (DeleteStackResourcesRequest, error) {
	req := DeleteStackResourcesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStackResourcesRequest{}, err
	}
	return req, nil
}

func NewDeleteStackResourcesRequestFromDict(data map[string]interface{}) DeleteStackResourcesRequest {
	return DeleteStackResourcesRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p DeleteStackResourcesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackResourcesRequest) Pointer() *DeleteStackResourcesRequest {
	return &p
}

type DeleteStackEntityRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
}

func (p *DeleteStackEntityRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStackEntityRequest{}
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
		*p = DeleteStackEntityRequest{}
	} else {
		*p = DeleteStackEntityRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
	}
	return nil
}

func NewDeleteStackEntityRequestFromJson(data string) (DeleteStackEntityRequest, error) {
	req := DeleteStackEntityRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStackEntityRequest{}, err
	}
	return req, nil
}

func NewDeleteStackEntityRequestFromDict(data map[string]interface{}) DeleteStackEntityRequest {
	return DeleteStackEntityRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
	}
}

func (p DeleteStackEntityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
	}
}

func (p DeleteStackEntityRequest) Pointer() *DeleteStackEntityRequest {
	return &p
}

type DescribeResourcesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeResourcesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeResourcesRequest{}
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
		*p = DescribeResourcesRequest{}
	} else {
		*p = DescribeResourcesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
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

func NewDescribeResourcesRequestFromJson(data string) (DescribeResourcesRequest, error) {
	req := DescribeResourcesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeResourcesRequest{}, err
	}
	return req, nil
}

func NewDescribeResourcesRequestFromDict(data map[string]interface{}) DescribeResourcesRequest {
	return DescribeResourcesRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
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

func (p DescribeResourcesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeResourcesRequest) Pointer() *DescribeResourcesRequest {
	return &p
}

type GetResourceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	ResourceName    *string `json:"resourceName"`
}

func (p *GetResourceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetResourceRequest{}
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
		*p = GetResourceRequest{}
	} else {
		*p = GetResourceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
		if v, ok := d["resourceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceName)
				}
			}
		}
	}
	return nil
}

func NewGetResourceRequestFromJson(data string) (GetResourceRequest, error) {
	req := GetResourceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetResourceRequest{}, err
	}
	return req, nil
}

func NewGetResourceRequestFromDict(data map[string]interface{}) GetResourceRequest {
	return GetResourceRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		ResourceName: func() *string {
			v, ok := data["resourceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceName"])
		}(),
	}
}

func (p GetResourceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":    p.StackName,
		"resourceName": p.ResourceName,
	}
}

func (p GetResourceRequest) Pointer() *GetResourceRequest {
	return &p
}

type DescribeEventsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeEventsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeEventsRequest{}
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
		*p = DescribeEventsRequest{}
	} else {
		*p = DescribeEventsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
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

func NewDescribeEventsRequestFromJson(data string) (DescribeEventsRequest, error) {
	req := DescribeEventsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeEventsRequest{}, err
	}
	return req, nil
}

func NewDescribeEventsRequestFromDict(data map[string]interface{}) DescribeEventsRequest {
	return DescribeEventsRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
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

func (p DescribeEventsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeEventsRequest) Pointer() *DescribeEventsRequest {
	return &p
}

type GetEventRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	EventName       *string `json:"eventName"`
}

func (p *GetEventRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetEventRequest{}
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
		*p = GetEventRequest{}
	} else {
		*p = GetEventRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
		if v, ok := d["eventName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventName)
				}
			}
		}
	}
	return nil
}

func NewGetEventRequestFromJson(data string) (GetEventRequest, error) {
	req := GetEventRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetEventRequest{}, err
	}
	return req, nil
}

func NewGetEventRequestFromDict(data map[string]interface{}) GetEventRequest {
	return GetEventRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		EventName: func() *string {
			v, ok := data["eventName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventName"])
		}(),
	}
}

func (p GetEventRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"eventName": p.EventName,
	}
}

func (p GetEventRequest) Pointer() *GetEventRequest {
	return &p
}

type DescribeOutputsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeOutputsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeOutputsRequest{}
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
		*p = DescribeOutputsRequest{}
	} else {
		*p = DescribeOutputsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
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

func NewDescribeOutputsRequestFromJson(data string) (DescribeOutputsRequest, error) {
	req := DescribeOutputsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeOutputsRequest{}, err
	}
	return req, nil
}

func NewDescribeOutputsRequestFromDict(data map[string]interface{}) DescribeOutputsRequest {
	return DescribeOutputsRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
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

func (p DescribeOutputsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName": p.StackName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeOutputsRequest) Pointer() *DescribeOutputsRequest {
	return &p
}

type GetOutputRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StackName       *string `json:"stackName"`
	OutputName      *string `json:"outputName"`
}

func (p *GetOutputRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetOutputRequest{}
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
		*p = GetOutputRequest{}
	} else {
		*p = GetOutputRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackName)
				}
			}
		}
		if v, ok := d["outputName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OutputName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OutputName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OutputName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OutputName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OutputName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OutputName)
				}
			}
		}
	}
	return nil
}

func NewGetOutputRequestFromJson(data string) (GetOutputRequest, error) {
	req := GetOutputRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetOutputRequest{}, err
	}
	return req, nil
}

func NewGetOutputRequestFromDict(data map[string]interface{}) GetOutputRequest {
	return GetOutputRequest{
		StackName: func() *string {
			v, ok := data["stackName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackName"])
		}(),
		OutputName: func() *string {
			v, ok := data["outputName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["outputName"])
		}(),
	}
}

func (p GetOutputRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stackName":  p.StackName,
		"outputName": p.OutputName,
	}
}

func (p GetOutputRequest) Pointer() *GetOutputRequest {
	return &p
}
