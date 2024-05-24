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

package megaField

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
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
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
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
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
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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

type DescribeAreaModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeAreaModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAreaModelsRequest{}
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
		*p = DescribeAreaModelsRequest{}
	} else {
		*p = DescribeAreaModelsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewDescribeAreaModelsRequestFromJson(data string) (DescribeAreaModelsRequest, error) {
	req := DescribeAreaModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAreaModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeAreaModelsRequestFromDict(data map[string]interface{}) DescribeAreaModelsRequest {
	return DescribeAreaModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeAreaModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeAreaModelsRequest) Pointer() *DescribeAreaModelsRequest {
	return &p
}

type GetAreaModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
}

func (p *GetAreaModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetAreaModelRequest{}
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
		*p = GetAreaModelRequest{}
	} else {
		*p = GetAreaModelRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
	}
	return nil
}

func NewGetAreaModelRequestFromJson(data string) (GetAreaModelRequest, error) {
	req := GetAreaModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetAreaModelRequest{}, err
	}
	return req, nil
}

func NewGetAreaModelRequestFromDict(data map[string]interface{}) GetAreaModelRequest {
	return GetAreaModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p GetAreaModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p GetAreaModelRequest) Pointer() *GetAreaModelRequest {
	return &p
}

type DescribeAreaModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeAreaModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAreaModelMastersRequest{}
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
		*p = DescribeAreaModelMastersRequest{}
	} else {
		*p = DescribeAreaModelMastersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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

func NewDescribeAreaModelMastersRequestFromJson(data string) (DescribeAreaModelMastersRequest, error) {
	req := DescribeAreaModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAreaModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeAreaModelMastersRequestFromDict(data map[string]interface{}) DescribeAreaModelMastersRequest {
	return DescribeAreaModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeAreaModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAreaModelMastersRequest) Pointer() *DescribeAreaModelMastersRequest {
	return &p
}

type CreateAreaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateAreaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateAreaModelMasterRequest{}
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
		*p = CreateAreaModelMasterRequest{}
	} else {
		*p = CreateAreaModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
	}
	return nil
}

func NewCreateAreaModelMasterRequestFromJson(data string) (CreateAreaModelMasterRequest, error) {
	req := CreateAreaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateAreaModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateAreaModelMasterRequestFromDict(data map[string]interface{}) CreateAreaModelMasterRequest {
	return CreateAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateAreaModelMasterRequest) Pointer() *CreateAreaModelMasterRequest {
	return &p
}

type GetAreaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
}

func (p *GetAreaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetAreaModelMasterRequest{}
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
		*p = GetAreaModelMasterRequest{}
	} else {
		*p = GetAreaModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
	}
	return nil
}

func NewGetAreaModelMasterRequestFromJson(data string) (GetAreaModelMasterRequest, error) {
	req := GetAreaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetAreaModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetAreaModelMasterRequestFromDict(data map[string]interface{}) GetAreaModelMasterRequest {
	return GetAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p GetAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p GetAreaModelMasterRequest) Pointer() *GetAreaModelMasterRequest {
	return &p
}

type UpdateAreaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateAreaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateAreaModelMasterRequest{}
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
		*p = UpdateAreaModelMasterRequest{}
	} else {
		*p = UpdateAreaModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
	}
	return nil
}

func NewUpdateAreaModelMasterRequestFromJson(data string) (UpdateAreaModelMasterRequest, error) {
	req := UpdateAreaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateAreaModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateAreaModelMasterRequestFromDict(data map[string]interface{}) UpdateAreaModelMasterRequest {
	return UpdateAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateAreaModelMasterRequest) Pointer() *UpdateAreaModelMasterRequest {
	return &p
}

type DeleteAreaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
}

func (p *DeleteAreaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteAreaModelMasterRequest{}
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
		*p = DeleteAreaModelMasterRequest{}
	} else {
		*p = DeleteAreaModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
	}
	return nil
}

func NewDeleteAreaModelMasterRequestFromJson(data string) (DeleteAreaModelMasterRequest, error) {
	req := DeleteAreaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteAreaModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteAreaModelMasterRequestFromDict(data map[string]interface{}) DeleteAreaModelMasterRequest {
	return DeleteAreaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p DeleteAreaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p DeleteAreaModelMasterRequest) Pointer() *DeleteAreaModelMasterRequest {
	return &p
}

type DescribeLayerModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
}

func (p *DescribeLayerModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLayerModelsRequest{}
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
		*p = DescribeLayerModelsRequest{}
	} else {
		*p = DescribeLayerModelsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
	}
	return nil
}

func NewDescribeLayerModelsRequestFromJson(data string) (DescribeLayerModelsRequest, error) {
	req := DescribeLayerModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLayerModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeLayerModelsRequestFromDict(data map[string]interface{}) DescribeLayerModelsRequest {
	return DescribeLayerModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
	}
}

func (p DescribeLayerModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
	}
}

func (p DescribeLayerModelsRequest) Pointer() *DescribeLayerModelsRequest {
	return &p
}

type GetLayerModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	LayerModelName  *string `json:"layerModelName"`
}

func (p *GetLayerModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLayerModelRequest{}
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
		*p = GetLayerModelRequest{}
	} else {
		*p = GetLayerModelRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
	}
	return nil
}

func NewGetLayerModelRequestFromJson(data string) (GetLayerModelRequest, error) {
	req := GetLayerModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLayerModelRequest{}, err
	}
	return req, nil
}

func NewGetLayerModelRequestFromDict(data map[string]interface{}) GetLayerModelRequest {
	return GetLayerModelRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p GetLayerModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p GetLayerModelRequest) Pointer() *GetLayerModelRequest {
	return &p
}

type DescribeLayerModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeLayerModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeLayerModelMastersRequest{}
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
		*p = DescribeLayerModelMastersRequest{}
	} else {
		*p = DescribeLayerModelMastersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
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

func NewDescribeLayerModelMastersRequestFromJson(data string) (DescribeLayerModelMastersRequest, error) {
	req := DescribeLayerModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeLayerModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeLayerModelMastersRequestFromDict(data map[string]interface{}) DescribeLayerModelMastersRequest {
	return DescribeLayerModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeLayerModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeLayerModelMastersRequest) Pointer() *DescribeLayerModelMastersRequest {
	return &p
}

type CreateLayerModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateLayerModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateLayerModelMasterRequest{}
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
		*p = CreateLayerModelMasterRequest{}
	} else {
		*p = CreateLayerModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
	}
	return nil
}

func NewCreateLayerModelMasterRequestFromJson(data string) (CreateLayerModelMasterRequest, error) {
	req := CreateLayerModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateLayerModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateLayerModelMasterRequestFromDict(data map[string]interface{}) CreateLayerModelMasterRequest {
	return CreateLayerModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AreaModelName: core.CastString(data["areaModelName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"areaModelName": p.AreaModelName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateLayerModelMasterRequest) Pointer() *CreateLayerModelMasterRequest {
	return &p
}

type GetLayerModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	LayerModelName  *string `json:"layerModelName"`
}

func (p *GetLayerModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLayerModelMasterRequest{}
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
		*p = GetLayerModelMasterRequest{}
	} else {
		*p = GetLayerModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
	}
	return nil
}

func NewGetLayerModelMasterRequestFromJson(data string) (GetLayerModelMasterRequest, error) {
	req := GetLayerModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLayerModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetLayerModelMasterRequestFromDict(data map[string]interface{}) GetLayerModelMasterRequest {
	return GetLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p GetLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p GetLayerModelMasterRequest) Pointer() *GetLayerModelMasterRequest {
	return &p
}

type UpdateLayerModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	LayerModelName  *string `json:"layerModelName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateLayerModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateLayerModelMasterRequest{}
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
		*p = UpdateLayerModelMasterRequest{}
	} else {
		*p = UpdateLayerModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
	}
	return nil
}

func NewUpdateLayerModelMasterRequestFromJson(data string) (UpdateLayerModelMasterRequest, error) {
	req := UpdateLayerModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateLayerModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateLayerModelMasterRequestFromDict(data map[string]interface{}) UpdateLayerModelMasterRequest {
	return UpdateLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
	}
}

func (p UpdateLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"description":    p.Description,
		"metadata":       p.Metadata,
	}
}

func (p UpdateLayerModelMasterRequest) Pointer() *UpdateLayerModelMasterRequest {
	return &p
}

type DeleteLayerModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AreaModelName   *string `json:"areaModelName"`
	LayerModelName  *string `json:"layerModelName"`
}

func (p *DeleteLayerModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteLayerModelMasterRequest{}
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
		*p = DeleteLayerModelMasterRequest{}
	} else {
		*p = DeleteLayerModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
	}
	return nil
}

func NewDeleteLayerModelMasterRequestFromJson(data string) (DeleteLayerModelMasterRequest, error) {
	req := DeleteLayerModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteLayerModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteLayerModelMasterRequestFromDict(data map[string]interface{}) DeleteLayerModelMasterRequest {
	return DeleteLayerModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
	}
}

func (p DeleteLayerModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
	}
}

func (p DeleteLayerModelMasterRequest) Pointer() *DeleteLayerModelMasterRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *ExportMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExportMasterRequest{}
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
		*p = ExportMasterRequest{}
	} else {
		*p = ExportMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewExportMasterRequestFromJson(data string) (ExportMasterRequest, error) {
	req := ExportMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ExportMasterRequest{}, err
	}
	return req, nil
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
	return ExportMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
	return &p
}

type GetCurrentFieldMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentFieldMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentFieldMasterRequest{}
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
		*p = GetCurrentFieldMasterRequest{}
	} else {
		*p = GetCurrentFieldMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
	}
	return nil
}

func NewGetCurrentFieldMasterRequestFromJson(data string) (GetCurrentFieldMasterRequest, error) {
	req := GetCurrentFieldMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentFieldMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentFieldMasterRequestFromDict(data map[string]interface{}) GetCurrentFieldMasterRequest {
	return GetCurrentFieldMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentFieldMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentFieldMasterRequest) Pointer() *GetCurrentFieldMasterRequest {
	return &p
}

type UpdateCurrentFieldMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentFieldMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentFieldMasterRequest{}
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
		*p = UpdateCurrentFieldMasterRequest{}
	} else {
		*p = UpdateCurrentFieldMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Settings = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Settings = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Settings = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Settings)
				}
			}
		}
	}
	return nil
}

func NewUpdateCurrentFieldMasterRequestFromJson(data string) (UpdateCurrentFieldMasterRequest, error) {
	req := UpdateCurrentFieldMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentFieldMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentFieldMasterRequestFromDict(data map[string]interface{}) UpdateCurrentFieldMasterRequest {
	return UpdateCurrentFieldMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentFieldMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentFieldMasterRequest) Pointer() *UpdateCurrentFieldMasterRequest {
	return &p
}

type UpdateCurrentFieldMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentFieldMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentFieldMasterFromGitHubRequest{}
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
		*p = UpdateCurrentFieldMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentFieldMasterFromGitHubRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["checkoutSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CheckoutSetting)
		}
	}
	return nil
}

func NewUpdateCurrentFieldMasterFromGitHubRequestFromJson(data string) (UpdateCurrentFieldMasterFromGitHubRequest, error) {
	req := UpdateCurrentFieldMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentFieldMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentFieldMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentFieldMasterFromGitHubRequest {
	return UpdateCurrentFieldMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentFieldMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentFieldMasterFromGitHubRequest) Pointer() *UpdateCurrentFieldMasterFromGitHubRequest {
	return &p
}

type PutPositionRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Position           *Position `json:"position"`
	Vector             *Vector   `json:"vector"`
	R                  *float32  `json:"r"`
}

func (p *PutPositionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutPositionRequest{}
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
		*p = PutPositionRequest{}
	} else {
		*p = PutPositionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["vector"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Vector)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
	}
	return nil
}

func NewPutPositionRequestFromJson(data string) (PutPositionRequest, error) {
	req := PutPositionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutPositionRequest{}, err
	}
	return req, nil
}

func NewPutPositionRequestFromDict(data map[string]interface{}) PutPositionRequest {
	return PutPositionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Vector:         NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
	}
}

func (p PutPositionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"vector":         p.Vector.ToDict(),
		"r":              p.R,
	}
}

func (p PutPositionRequest) Pointer() *PutPositionRequest {
	return &p
}

type PutPositionByUserIdRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	UserId             *string   `json:"userId"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Position           *Position `json:"position"`
	Vector             *Vector   `json:"vector"`
	R                  *float32  `json:"r"`
	TimeOffsetToken    *string   `json:"timeOffsetToken"`
}

func (p *PutPositionByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutPositionByUserIdRequest{}
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
		*p = PutPositionByUserIdRequest{}
	} else {
		*p = PutPositionByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["vector"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Vector)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
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

func NewPutPositionByUserIdRequestFromJson(data string) (PutPositionByUserIdRequest, error) {
	req := PutPositionByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutPositionByUserIdRequest{}, err
	}
	return req, nil
}

func NewPutPositionByUserIdRequestFromDict(data map[string]interface{}) PutPositionByUserIdRequest {
	return PutPositionByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		AreaModelName:   core.CastString(data["areaModelName"]),
		LayerModelName:  core.CastString(data["layerModelName"]),
		Position:        NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Vector:          NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
		R:               core.CastFloat32(data["r"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PutPositionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"areaModelName":   p.AreaModelName,
		"layerModelName":  p.LayerModelName,
		"position":        p.Position.ToDict(),
		"vector":          p.Vector.ToDict(),
		"r":               p.R,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PutPositionByUserIdRequest) Pointer() *PutPositionByUserIdRequest {
	return &p
}

type FetchPositionRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	UserIds            []*string `json:"userIds"`
}

func (p *FetchPositionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FetchPositionRequest{}
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
		*p = FetchPositionRequest{}
	} else {
		*p = FetchPositionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["userIds"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.UserIds = l
			}
		}
	}
	return nil
}

func NewFetchPositionRequestFromJson(data string) (FetchPositionRequest, error) {
	req := FetchPositionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return FetchPositionRequest{}, err
	}
	return req, nil
}

func NewFetchPositionRequestFromDict(data map[string]interface{}) FetchPositionRequest {
	return FetchPositionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		UserIds:        core.CastStrings(core.CastArray(data["userIds"])),
	}
}

func (p FetchPositionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"userIds": core.CastStringsFromDict(
			p.UserIds,
		),
	}
}

func (p FetchPositionRequest) Pointer() *FetchPositionRequest {
	return &p
}

type FetchPositionFromSystemRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	UserIds            []*string `json:"userIds"`
}

func (p *FetchPositionFromSystemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FetchPositionFromSystemRequest{}
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
		*p = FetchPositionFromSystemRequest{}
	} else {
		*p = FetchPositionFromSystemRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["userIds"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.UserIds = l
			}
		}
	}
	return nil
}

func NewFetchPositionFromSystemRequestFromJson(data string) (FetchPositionFromSystemRequest, error) {
	req := FetchPositionFromSystemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return FetchPositionFromSystemRequest{}, err
	}
	return req, nil
}

func NewFetchPositionFromSystemRequestFromDict(data map[string]interface{}) FetchPositionFromSystemRequest {
	return FetchPositionFromSystemRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		UserIds:        core.CastStrings(core.CastArray(data["userIds"])),
	}
}

func (p FetchPositionFromSystemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"userIds": core.CastStringsFromDict(
			p.UserIds,
		),
	}
}

func (p FetchPositionFromSystemRequest) Pointer() *FetchPositionFromSystemRequest {
	return &p
}

type NearUserIdsRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Point              *Position `json:"point"`
	R                  *float32  `json:"r"`
	Limit              *int32    `json:"limit"`
}

func (p *NearUserIdsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NearUserIdsRequest{}
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
		*p = NearUserIdsRequest{}
	} else {
		*p = NearUserIdsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["point"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Point)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewNearUserIdsRequestFromJson(data string) (NearUserIdsRequest, error) {
	req := NearUserIdsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return NearUserIdsRequest{}, err
	}
	return req, nil
}

func NewNearUserIdsRequestFromDict(data map[string]interface{}) NearUserIdsRequest {
	return NearUserIdsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Point:          NewPositionFromDict(core.CastMap(data["point"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p NearUserIdsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"point":          p.Point.ToDict(),
		"r":              p.R,
		"limit":          p.Limit,
	}
}

func (p NearUserIdsRequest) Pointer() *NearUserIdsRequest {
	return &p
}

type NearUserIdsFromSystemRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AreaModelName      *string   `json:"areaModelName"`
	LayerModelName     *string   `json:"layerModelName"`
	Point              *Position `json:"point"`
	R                  *float32  `json:"r"`
	Limit              *int32    `json:"limit"`
}

func (p *NearUserIdsFromSystemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NearUserIdsFromSystemRequest{}
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
		*p = NearUserIdsFromSystemRequest{}
	} else {
		*p = NearUserIdsFromSystemRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["point"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Point)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewNearUserIdsFromSystemRequestFromJson(data string) (NearUserIdsFromSystemRequest, error) {
	req := NearUserIdsFromSystemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return NearUserIdsFromSystemRequest{}, err
	}
	return req, nil
}

func NewNearUserIdsFromSystemRequestFromDict(data map[string]interface{}) NearUserIdsFromSystemRequest {
	return NearUserIdsFromSystemRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Point:          NewPositionFromDict(core.CastMap(data["point"])).Pointer(),
		R:              core.CastFloat32(data["r"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p NearUserIdsFromSystemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"point":          p.Point.ToDict(),
		"r":              p.R,
		"limit":          p.Limit,
	}
}

func (p NearUserIdsFromSystemRequest) Pointer() *NearUserIdsFromSystemRequest {
	return &p
}

type ActionRequest struct {
	SourceRequestId    *string     `json:"sourceRequestId"`
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	AccessToken        *string     `json:"accessToken"`
	AreaModelName      *string     `json:"areaModelName"`
	LayerModelName     *string     `json:"layerModelName"`
	Position           *MyPosition `json:"position"`
	Scopes             []Scope     `json:"scopes"`
}

func (p *ActionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ActionRequest{}
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
		*p = ActionRequest{}
	} else {
		*p = ActionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessToken = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessToken = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessToken = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessToken = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessToken)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["scopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Scopes)
		}
	}
	return nil
}

func NewActionRequestFromJson(data string) (ActionRequest, error) {
	req := ActionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ActionRequest{}, err
	}
	return req, nil
}

func NewActionRequestFromDict(data map[string]interface{}) ActionRequest {
	return ActionRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		AreaModelName:  core.CastString(data["areaModelName"]),
		LayerModelName: core.CastString(data["layerModelName"]),
		Position:       NewMyPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Scopes:         CastScopes(core.CastArray(data["scopes"])),
	}
}

func (p ActionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"scopes": CastScopesFromDict(
			p.Scopes,
		),
	}
}

func (p ActionRequest) Pointer() *ActionRequest {
	return &p
}

type ActionByUserIdRequest struct {
	SourceRequestId    *string     `json:"sourceRequestId"`
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	UserId             *string     `json:"userId"`
	AreaModelName      *string     `json:"areaModelName"`
	LayerModelName     *string     `json:"layerModelName"`
	Position           *MyPosition `json:"position"`
	Scopes             []Scope     `json:"scopes"`
	TimeOffsetToken    *string     `json:"timeOffsetToken"`
}

func (p *ActionByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ActionByUserIdRequest{}
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
		*p = ActionByUserIdRequest{}
	} else {
		*p = ActionByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.NamespaceName)
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
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["scopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Scopes)
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

func NewActionByUserIdRequestFromJson(data string) (ActionByUserIdRequest, error) {
	req := ActionByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ActionByUserIdRequest{}, err
	}
	return req, nil
}

func NewActionByUserIdRequestFromDict(data map[string]interface{}) ActionByUserIdRequest {
	return ActionByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		AreaModelName:   core.CastString(data["areaModelName"]),
		LayerModelName:  core.CastString(data["layerModelName"]),
		Position:        NewMyPositionFromDict(core.CastMap(data["position"])).Pointer(),
		Scopes:          CastScopes(core.CastArray(data["scopes"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ActionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"areaModelName":  p.AreaModelName,
		"layerModelName": p.LayerModelName,
		"position":       p.Position.ToDict(),
		"scopes": CastScopesFromDict(
			p.Scopes,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ActionByUserIdRequest) Pointer() *ActionByUserIdRequest {
	return &p
}
