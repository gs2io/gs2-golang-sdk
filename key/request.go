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

package key

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeNamespacesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeNamespacesRequest{}
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
		*p = DescribeNamespacesRequest{}
	} else {
		*p = DescribeNamespacesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewDescribeNamespacesRequestFromJson(data string) (DescribeNamespacesRequest, error) {
	req := DescribeNamespacesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeNamespacesRequest{}, err
	}
	return req, nil
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	Name            *string     `json:"name"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func (p *CreateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateNamespaceRequest{}
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
		*p = CreateNamespaceRequest{}
	} else {
		*p = CreateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Name); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.LogSetting); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCreateNamespaceRequestFromJson(data string) (CreateNamespaceRequest, error) {
	req := CreateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateNamespaceRequest{}, err
	}
	return req, nil
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceStatusRequest{}
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
		*p = GetNamespaceStatusRequest{}
	} else {
		*p = GetNamespaceStatusRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceStatusRequestFromJson(data string) (GetNamespaceStatusRequest, error) {
	req := GetNamespaceStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceStatusRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceRequest{}
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
		*p = GetNamespaceRequest{}
	} else {
		*p = GetNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceRequestFromJson(data string) (GetNamespaceRequest, error) {
	req := GetNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	SourceRequestId *string     `json:"sourceRequestId"`
	RequestId       *string     `json:"requestId"`
	ContextStack    *string     `json:"contextStack"`
	NamespaceName   *string     `json:"namespaceName"`
	Description     *string     `json:"description"`
	LogSetting      *LogSetting `json:"logSetting"`
}

func (p *UpdateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateNamespaceRequest{}
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
		*p = UpdateNamespaceRequest{}
	} else {
		*p = UpdateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.LogSetting); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateNamespaceRequestFromJson(data string) (UpdateNamespaceRequest, error) {
	req := UpdateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateNamespaceRequest{}, err
	}
	return req, nil
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Description:   core.CastString(data["description"]),
		LogSetting:    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"logSetting":    p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DeleteNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteNamespaceRequest{}
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
		*p = DeleteNamespaceRequest{}
	} else {
		*p = DeleteNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteNamespaceRequestFromJson(data string) (DeleteNamespaceRequest, error) {
	req := DeleteNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteNamespaceRequest{}, err
	}
	return req, nil
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type DescribeKeysRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeKeysRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeKeysRequest{}
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
		*p = DescribeKeysRequest{}
	} else {
		*p = DescribeKeysRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewDescribeKeysRequestFromJson(data string) (DescribeKeysRequest, error) {
	req := DescribeKeysRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeKeysRequest{}, err
	}
	return req, nil
}

func NewDescribeKeysRequestFromDict(data map[string]interface{}) DescribeKeysRequest {
	return DescribeKeysRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeKeysRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeKeysRequest) Pointer() *DescribeKeysRequest {
	return &p
}

type CreateKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
}

func (p *CreateKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateKeyRequest{}
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
		*p = CreateKeyRequest{}
	} else {
		*p = CreateKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Name); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCreateKeyRequestFromJson(data string) (CreateKeyRequest, error) {
	req := CreateKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateKeyRequest{}, err
	}
	return req, nil
}

func NewCreateKeyRequestFromDict(data map[string]interface{}) CreateKeyRequest {
	return CreateKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
	}
}

func (p CreateKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
	}
}

func (p CreateKeyRequest) Pointer() *CreateKeyRequest {
	return &p
}

type UpdateKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	KeyName         *string `json:"keyName"`
	Description     *string `json:"description"`
}

func (p *UpdateKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateKeyRequest{}
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
		*p = UpdateKeyRequest{}
	} else {
		*p = UpdateKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateKeyRequestFromJson(data string) (UpdateKeyRequest, error) {
	req := UpdateKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateKeyRequest{}, err
	}
	return req, nil
}

func NewUpdateKeyRequestFromDict(data map[string]interface{}) UpdateKeyRequest {
	return UpdateKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		KeyName:       core.CastString(data["keyName"]),
		Description:   core.CastString(data["description"]),
	}
}

func (p UpdateKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"keyName":       p.KeyName,
		"description":   p.Description,
	}
}

func (p UpdateKeyRequest) Pointer() *UpdateKeyRequest {
	return &p
}

type GetKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	KeyName         *string `json:"keyName"`
}

func (p *GetKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetKeyRequest{}
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
		*p = GetKeyRequest{}
	} else {
		*p = GetKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetKeyRequestFromJson(data string) (GetKeyRequest, error) {
	req := GetKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetKeyRequest{}, err
	}
	return req, nil
}

func NewGetKeyRequestFromDict(data map[string]interface{}) GetKeyRequest {
	return GetKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		KeyName:       core.CastString(data["keyName"]),
	}
}

func (p GetKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"keyName":       p.KeyName,
	}
}

func (p GetKeyRequest) Pointer() *GetKeyRequest {
	return &p
}

type DeleteKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	KeyName         *string `json:"keyName"`
}

func (p *DeleteKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteKeyRequest{}
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
		*p = DeleteKeyRequest{}
	} else {
		*p = DeleteKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteKeyRequestFromJson(data string) (DeleteKeyRequest, error) {
	req := DeleteKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteKeyRequest{}, err
	}
	return req, nil
}

func NewDeleteKeyRequestFromDict(data map[string]interface{}) DeleteKeyRequest {
	return DeleteKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		KeyName:       core.CastString(data["keyName"]),
	}
}

func (p DeleteKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"keyName":       p.KeyName,
	}
}

func (p DeleteKeyRequest) Pointer() *DeleteKeyRequest {
	return &p
}

type EncryptRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	KeyName         *string `json:"keyName"`
	Data            *string `json:"data"`
}

func (p *EncryptRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EncryptRequest{}
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
		*p = EncryptRequest{}
	} else {
		*p = EncryptRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["data"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.Data = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.Data = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.Data = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.Data = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.Data = &strValue
			default:
				if err := json.Unmarshal(*v, &p.Data); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewEncryptRequestFromJson(data string) (EncryptRequest, error) {
	req := EncryptRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return EncryptRequest{}, err
	}
	return req, nil
}

func NewEncryptRequestFromDict(data map[string]interface{}) EncryptRequest {
	return EncryptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		KeyName:       core.CastString(data["keyName"]),
		Data:          core.CastString(data["data"]),
	}
}

func (p EncryptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"keyName":       p.KeyName,
		"data":          p.Data,
	}
}

func (p EncryptRequest) Pointer() *EncryptRequest {
	return &p
}

type DecryptRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	KeyName         *string `json:"keyName"`
	Data            *string `json:"data"`
}

func (p *DecryptRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecryptRequest{}
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
		*p = DecryptRequest{}
	} else {
		*p = DecryptRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.KeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.KeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.KeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.KeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.KeyName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["data"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.Data = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.Data = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.Data = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.Data = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.Data = &strValue
			default:
				if err := json.Unmarshal(*v, &p.Data); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDecryptRequestFromJson(data string) (DecryptRequest, error) {
	req := DecryptRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecryptRequest{}, err
	}
	return req, nil
}

func NewDecryptRequestFromDict(data map[string]interface{}) DecryptRequest {
	return DecryptRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		KeyName:       core.CastString(data["keyName"]),
		Data:          core.CastString(data["data"]),
	}
}

func (p DecryptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"keyName":       p.KeyName,
		"data":          p.Data,
	}
}

func (p DecryptRequest) Pointer() *DecryptRequest {
	return &p
}

type DescribeGitHubApiKeysRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeGitHubApiKeysRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGitHubApiKeysRequest{}
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
		*p = DescribeGitHubApiKeysRequest{}
	} else {
		*p = DescribeGitHubApiKeysRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewDescribeGitHubApiKeysRequestFromJson(data string) (DescribeGitHubApiKeysRequest, error) {
	req := DescribeGitHubApiKeysRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGitHubApiKeysRequest{}, err
	}
	return req, nil
}

func NewDescribeGitHubApiKeysRequestFromDict(data map[string]interface{}) DescribeGitHubApiKeysRequest {
	return DescribeGitHubApiKeysRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeGitHubApiKeysRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGitHubApiKeysRequest) Pointer() *DescribeGitHubApiKeysRequest {
	return &p
}

type CreateGitHubApiKeyRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	Name              *string `json:"name"`
	Description       *string `json:"description"`
	ApiKey            *string `json:"apiKey"`
	EncryptionKeyName *string `json:"encryptionKeyName"`
}

func (p *CreateGitHubApiKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGitHubApiKeyRequest{}
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
		*p = CreateGitHubApiKeyRequest{}
	} else {
		*p = CreateGitHubApiKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Name); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["apiKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.ApiKey = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.ApiKey = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.ApiKey = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.ApiKey = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.ApiKey = &strValue
			default:
				if err := json.Unmarshal(*v, &p.ApiKey); err != nil {
					return err
				}
			}
		}
		if v, ok := d["encryptionKeyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.EncryptionKeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.EncryptionKeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.EncryptionKeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.EncryptionKeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.EncryptionKeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.EncryptionKeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCreateGitHubApiKeyRequestFromJson(data string) (CreateGitHubApiKeyRequest, error) {
	req := CreateGitHubApiKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGitHubApiKeyRequest{}, err
	}
	return req, nil
}

func NewCreateGitHubApiKeyRequestFromDict(data map[string]interface{}) CreateGitHubApiKeyRequest {
	return CreateGitHubApiKeyRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		ApiKey:            core.CastString(data["apiKey"]),
		EncryptionKeyName: core.CastString(data["encryptionKeyName"]),
	}
}

func (p CreateGitHubApiKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"apiKey":            p.ApiKey,
		"encryptionKeyName": p.EncryptionKeyName,
	}
}

func (p CreateGitHubApiKeyRequest) Pointer() *CreateGitHubApiKeyRequest {
	return &p
}

type UpdateGitHubApiKeyRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	ApiKeyName        *string `json:"apiKeyName"`
	Description       *string `json:"description"`
	ApiKey            *string `json:"apiKey"`
	EncryptionKeyName *string `json:"encryptionKeyName"`
}

func (p *UpdateGitHubApiKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGitHubApiKeyRequest{}
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
		*p = UpdateGitHubApiKeyRequest{}
	} else {
		*p = UpdateGitHubApiKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["apiKeyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.ApiKeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.ApiKeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.ApiKeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.ApiKeyName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["apiKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.ApiKey = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.ApiKey = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.ApiKey = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.ApiKey = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.ApiKey = &strValue
			default:
				if err := json.Unmarshal(*v, &p.ApiKey); err != nil {
					return err
				}
			}
		}
		if v, ok := d["encryptionKeyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.EncryptionKeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.EncryptionKeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.EncryptionKeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.EncryptionKeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.EncryptionKeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.EncryptionKeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateGitHubApiKeyRequestFromJson(data string) (UpdateGitHubApiKeyRequest, error) {
	req := UpdateGitHubApiKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGitHubApiKeyRequest{}, err
	}
	return req, nil
}

func NewUpdateGitHubApiKeyRequestFromDict(data map[string]interface{}) UpdateGitHubApiKeyRequest {
	return UpdateGitHubApiKeyRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		ApiKeyName:        core.CastString(data["apiKeyName"]),
		Description:       core.CastString(data["description"]),
		ApiKey:            core.CastString(data["apiKey"]),
		EncryptionKeyName: core.CastString(data["encryptionKeyName"]),
	}
}

func (p UpdateGitHubApiKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"apiKeyName":        p.ApiKeyName,
		"description":       p.Description,
		"apiKey":            p.ApiKey,
		"encryptionKeyName": p.EncryptionKeyName,
	}
}

func (p UpdateGitHubApiKeyRequest) Pointer() *UpdateGitHubApiKeyRequest {
	return &p
}

type GetGitHubApiKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ApiKeyName      *string `json:"apiKeyName"`
}

func (p *GetGitHubApiKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGitHubApiKeyRequest{}
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
		*p = GetGitHubApiKeyRequest{}
	} else {
		*p = GetGitHubApiKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["apiKeyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.ApiKeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.ApiKeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.ApiKeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.ApiKeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetGitHubApiKeyRequestFromJson(data string) (GetGitHubApiKeyRequest, error) {
	req := GetGitHubApiKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGitHubApiKeyRequest{}, err
	}
	return req, nil
}

func NewGetGitHubApiKeyRequestFromDict(data map[string]interface{}) GetGitHubApiKeyRequest {
	return GetGitHubApiKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ApiKeyName:    core.CastString(data["apiKeyName"]),
	}
}

func (p GetGitHubApiKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"apiKeyName":    p.ApiKeyName,
	}
}

func (p GetGitHubApiKeyRequest) Pointer() *GetGitHubApiKeyRequest {
	return &p
}

type DeleteGitHubApiKeyRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ApiKeyName      *string `json:"apiKeyName"`
}

func (p *DeleteGitHubApiKeyRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGitHubApiKeyRequest{}
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
		*p = DeleteGitHubApiKeyRequest{}
	} else {
		*p = DeleteGitHubApiKeyRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["apiKeyName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.ApiKeyName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.ApiKeyName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.ApiKeyName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.ApiKeyName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.ApiKeyName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteGitHubApiKeyRequestFromJson(data string) (DeleteGitHubApiKeyRequest, error) {
	req := DeleteGitHubApiKeyRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGitHubApiKeyRequest{}, err
	}
	return req, nil
}

func NewDeleteGitHubApiKeyRequestFromDict(data map[string]interface{}) DeleteGitHubApiKeyRequest {
	return DeleteGitHubApiKeyRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		ApiKeyName:    core.CastString(data["apiKeyName"]),
	}
}

func (p DeleteGitHubApiKeyRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"apiKeyName":    p.ApiKeyName,
	}
}

func (p DeleteGitHubApiKeyRequest) Pointer() *DeleteGitHubApiKeyRequest {
	return &p
}
