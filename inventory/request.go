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

package inventory

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
	SourceRequestId         *string        `json:"sourceRequestId"`
	RequestId               *string        `json:"requestId"`
	ContextStack            *string        `json:"contextStack"`
	Name                    *string        `json:"name"`
	Description             *string        `json:"description"`
	AcquireScript           *ScriptSetting `json:"acquireScript"`
	OverflowScript          *ScriptSetting `json:"overflowScript"`
	ConsumeScript           *ScriptSetting `json:"consumeScript"`
	SimpleItemAcquireScript *ScriptSetting `json:"simpleItemAcquireScript"`
	SimpleItemConsumeScript *ScriptSetting `json:"simpleItemConsumeScript"`
	BigItemAcquireScript    *ScriptSetting `json:"bigItemAcquireScript"`
	BigItemConsumeScript    *ScriptSetting `json:"bigItemConsumeScript"`
	LogSetting              *LogSetting    `json:"logSetting"`
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
		if v, ok := d["acquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireScript)
		}
		if v, ok := d["overflowScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OverflowScript)
		}
		if v, ok := d["consumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeScript)
		}
		if v, ok := d["simpleItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemAcquireScript)
		}
		if v, ok := d["simpleItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemConsumeScript)
		}
		if v, ok := d["bigItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemAcquireScript)
		}
		if v, ok := d["bigItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemConsumeScript)
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
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		AcquireScript:           NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript:          NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:           NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		SimpleItemAcquireScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemAcquireScript"])).Pointer(),
		SimpleItemConsumeScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemConsumeScript"])).Pointer(),
		BigItemAcquireScript:    NewScriptSettingFromDict(core.CastMap(data["bigItemAcquireScript"])).Pointer(),
		BigItemConsumeScript:    NewScriptSettingFromDict(core.CastMap(data["bigItemConsumeScript"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":        p.Name,
		"description": p.Description,
		"acquireScript": func() map[string]interface{} {
			if p.AcquireScript == nil {
				return nil
			}
			return p.AcquireScript.ToDict()
		}(),
		"overflowScript": func() map[string]interface{} {
			if p.OverflowScript == nil {
				return nil
			}
			return p.OverflowScript.ToDict()
		}(),
		"consumeScript": func() map[string]interface{} {
			if p.ConsumeScript == nil {
				return nil
			}
			return p.ConsumeScript.ToDict()
		}(),
		"simpleItemAcquireScript": func() map[string]interface{} {
			if p.SimpleItemAcquireScript == nil {
				return nil
			}
			return p.SimpleItemAcquireScript.ToDict()
		}(),
		"simpleItemConsumeScript": func() map[string]interface{} {
			if p.SimpleItemConsumeScript == nil {
				return nil
			}
			return p.SimpleItemConsumeScript.ToDict()
		}(),
		"bigItemAcquireScript": func() map[string]interface{} {
			if p.BigItemAcquireScript == nil {
				return nil
			}
			return p.BigItemAcquireScript.ToDict()
		}(),
		"bigItemConsumeScript": func() map[string]interface{} {
			if p.BigItemConsumeScript == nil {
				return nil
			}
			return p.BigItemConsumeScript.ToDict()
		}(),
		"logSetting": func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}(),
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
	SourceRequestId         *string        `json:"sourceRequestId"`
	RequestId               *string        `json:"requestId"`
	ContextStack            *string        `json:"contextStack"`
	NamespaceName           *string        `json:"namespaceName"`
	Description             *string        `json:"description"`
	AcquireScript           *ScriptSetting `json:"acquireScript"`
	OverflowScript          *ScriptSetting `json:"overflowScript"`
	ConsumeScript           *ScriptSetting `json:"consumeScript"`
	SimpleItemAcquireScript *ScriptSetting `json:"simpleItemAcquireScript"`
	SimpleItemConsumeScript *ScriptSetting `json:"simpleItemConsumeScript"`
	BigItemAcquireScript    *ScriptSetting `json:"bigItemAcquireScript"`
	BigItemConsumeScript    *ScriptSetting `json:"bigItemConsumeScript"`
	LogSetting              *LogSetting    `json:"logSetting"`
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
		if v, ok := d["acquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireScript)
		}
		if v, ok := d["overflowScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OverflowScript)
		}
		if v, ok := d["consumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeScript)
		}
		if v, ok := d["simpleItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemAcquireScript)
		}
		if v, ok := d["simpleItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemConsumeScript)
		}
		if v, ok := d["bigItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemAcquireScript)
		}
		if v, ok := d["bigItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemConsumeScript)
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
		NamespaceName:           core.CastString(data["namespaceName"]),
		Description:             core.CastString(data["description"]),
		AcquireScript:           NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript:          NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:           NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		SimpleItemAcquireScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemAcquireScript"])).Pointer(),
		SimpleItemConsumeScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemConsumeScript"])).Pointer(),
		BigItemAcquireScript:    NewScriptSettingFromDict(core.CastMap(data["bigItemAcquireScript"])).Pointer(),
		BigItemConsumeScript:    NewScriptSettingFromDict(core.CastMap(data["bigItemConsumeScript"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"acquireScript": func() map[string]interface{} {
			if p.AcquireScript == nil {
				return nil
			}
			return p.AcquireScript.ToDict()
		}(),
		"overflowScript": func() map[string]interface{} {
			if p.OverflowScript == nil {
				return nil
			}
			return p.OverflowScript.ToDict()
		}(),
		"consumeScript": func() map[string]interface{} {
			if p.ConsumeScript == nil {
				return nil
			}
			return p.ConsumeScript.ToDict()
		}(),
		"simpleItemAcquireScript": func() map[string]interface{} {
			if p.SimpleItemAcquireScript == nil {
				return nil
			}
			return p.SimpleItemAcquireScript.ToDict()
		}(),
		"simpleItemConsumeScript": func() map[string]interface{} {
			if p.SimpleItemConsumeScript == nil {
				return nil
			}
			return p.SimpleItemConsumeScript.ToDict()
		}(),
		"bigItemAcquireScript": func() map[string]interface{} {
			if p.BigItemAcquireScript == nil {
				return nil
			}
			return p.BigItemAcquireScript.ToDict()
		}(),
		"bigItemConsumeScript": func() map[string]interface{} {
			if p.BigItemConsumeScript == nil {
				return nil
			}
			return p.BigItemConsumeScript.ToDict()
		}(),
		"logSetting": func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}(),
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

type DumpUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DumpUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DumpUserDataByUserIdRequest{}
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
		*p = DumpUserDataByUserIdRequest{}
	} else {
		*p = DumpUserDataByUserIdRequest{}
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

func NewDumpUserDataByUserIdRequestFromJson(data string) (DumpUserDataByUserIdRequest, error) {
	req := DumpUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DumpUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckDumpUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckDumpUserDataByUserIdRequest{}
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
		*p = CheckDumpUserDataByUserIdRequest{}
	} else {
		*p = CheckDumpUserDataByUserIdRequest{}
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

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) (CheckDumpUserDataByUserIdRequest, error) {
	req := CheckDumpUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckDumpUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CleanUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CleanUserDataByUserIdRequest{}
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
		*p = CleanUserDataByUserIdRequest{}
	} else {
		*p = CleanUserDataByUserIdRequest{}
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

func NewCleanUserDataByUserIdRequestFromJson(data string) (CleanUserDataByUserIdRequest, error) {
	req := CleanUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CleanUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckCleanUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckCleanUserDataByUserIdRequest{}
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
		*p = CheckCleanUserDataByUserIdRequest{}
	} else {
		*p = CheckCleanUserDataByUserIdRequest{}
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

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) (CheckCleanUserDataByUserIdRequest, error) {
	req := CheckCleanUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckCleanUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *PrepareImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrepareImportUserDataByUserIdRequest{}
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
		*p = PrepareImportUserDataByUserIdRequest{}
	} else {
		*p = PrepareImportUserDataByUserIdRequest{}
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

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) (PrepareImportUserDataByUserIdRequest, error) {
	req := PrepareImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PrepareImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *ImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ImportUserDataByUserIdRequest{}
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
		*p = ImportUserDataByUserIdRequest{}
	} else {
		*p = ImportUserDataByUserIdRequest{}
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

func NewImportUserDataByUserIdRequestFromJson(data string) (ImportUserDataByUserIdRequest, error) {
	req := ImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CheckImportUserDataByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckImportUserDataByUserIdRequest{}
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
		*p = CheckImportUserDataByUserIdRequest{}
	} else {
		*p = CheckImportUserDataByUserIdRequest{}
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

func NewCheckImportUserDataByUserIdRequestFromJson(data string) (CheckImportUserDataByUserIdRequest, error) {
	req := CheckImportUserDataByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckImportUserDataByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeInventoryModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeInventoryModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoryModelMastersRequest{}
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
		*p = DescribeInventoryModelMastersRequest{}
	} else {
		*p = DescribeInventoryModelMastersRequest{}
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

func NewDescribeInventoryModelMastersRequestFromJson(data string) (DescribeInventoryModelMastersRequest, error) {
	req := DescribeInventoryModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoryModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeInventoryModelMastersRequest {
	return DescribeInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInventoryModelMastersRequest) Pointer() *DescribeInventoryModelMastersRequest {
	return &p
}

type CreateInventoryModelMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	Name                  *string `json:"name"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
}

func (p *CreateInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateInventoryModelMasterRequest{}
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
		*p = CreateInventoryModelMasterRequest{}
	} else {
		*p = CreateInventoryModelMasterRequest{}
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
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["protectReferencedItem"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ProtectReferencedItem)
		}
	}
	return nil
}

func NewCreateInventoryModelMasterRequestFromJson(data string) (CreateInventoryModelMasterRequest, error) {
	req := CreateInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateInventoryModelMasterRequest {
	return CreateInventoryModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
	}
}

func (p CreateInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"name":                  p.Name,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
	}
}

func (p CreateInventoryModelMasterRequest) Pointer() *CreateInventoryModelMasterRequest {
	return &p
}

type GetInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInventoryModelMasterRequest{}
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
		*p = GetInventoryModelMasterRequest{}
	} else {
		*p = GetInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetInventoryModelMasterRequestFromJson(data string) (GetInventoryModelMasterRequest, error) {
	req := GetInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetInventoryModelMasterRequestFromDict(data map[string]interface{}) GetInventoryModelMasterRequest {
	return GetInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetInventoryModelMasterRequest) Pointer() *GetInventoryModelMasterRequest {
	return &p
}

type UpdateInventoryModelMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	InventoryName         *string `json:"inventoryName"`
	Description           *string `json:"description"`
	Metadata              *string `json:"metadata"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
}

func (p *UpdateInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateInventoryModelMasterRequest{}
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
		*p = UpdateInventoryModelMasterRequest{}
	} else {
		*p = UpdateInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["protectReferencedItem"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ProtectReferencedItem)
		}
	}
	return nil
}

func NewUpdateInventoryModelMasterRequestFromJson(data string) (UpdateInventoryModelMasterRequest, error) {
	req := UpdateInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateInventoryModelMasterRequest {
	return UpdateInventoryModelMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		InventoryName:         core.CastString(data["inventoryName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
	}
}

func (p UpdateInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"inventoryName":         p.InventoryName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
	}
}

func (p UpdateInventoryModelMasterRequest) Pointer() *UpdateInventoryModelMasterRequest {
	return &p
}

type DeleteInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DeleteInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteInventoryModelMasterRequest{}
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
		*p = DeleteInventoryModelMasterRequest{}
	} else {
		*p = DeleteInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDeleteInventoryModelMasterRequestFromJson(data string) (DeleteInventoryModelMasterRequest, error) {
	req := DeleteInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteInventoryModelMasterRequest {
	return DeleteInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteInventoryModelMasterRequest) Pointer() *DeleteInventoryModelMasterRequest {
	return &p
}

type DescribeInventoryModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeInventoryModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoryModelsRequest{}
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
		*p = DescribeInventoryModelsRequest{}
	} else {
		*p = DescribeInventoryModelsRequest{}
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

func NewDescribeInventoryModelsRequestFromJson(data string) (DescribeInventoryModelsRequest, error) {
	req := DescribeInventoryModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoryModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoryModelsRequestFromDict(data map[string]interface{}) DescribeInventoryModelsRequest {
	return DescribeInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeInventoryModelsRequest) Pointer() *DescribeInventoryModelsRequest {
	return &p
}

type GetInventoryModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetInventoryModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInventoryModelRequest{}
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
		*p = GetInventoryModelRequest{}
	} else {
		*p = GetInventoryModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetInventoryModelRequestFromJson(data string) (GetInventoryModelRequest, error) {
	req := GetInventoryModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInventoryModelRequest{}, err
	}
	return req, nil
}

func NewGetInventoryModelRequestFromDict(data map[string]interface{}) GetInventoryModelRequest {
	return GetInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetInventoryModelRequest) Pointer() *GetInventoryModelRequest {
	return &p
}

type DescribeItemModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeItemModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeItemModelMastersRequest{}
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
		*p = DescribeItemModelMastersRequest{}
	} else {
		*p = DescribeItemModelMastersRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeItemModelMastersRequestFromJson(data string) (DescribeItemModelMastersRequest, error) {
	req := DescribeItemModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeItemModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeItemModelMastersRequestFromDict(data map[string]interface{}) DescribeItemModelMastersRequest {
	return DescribeItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeItemModelMastersRequest) Pointer() *DescribeItemModelMastersRequest {
	return &p
}

type CreateItemModelMasterRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func (p *CreateItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateItemModelMasterRequest{}
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
		*p = CreateItemModelMasterRequest{}
	} else {
		*p = CreateItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["stackingLimit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StackingLimit)
		}
		if v, ok := d["allowMultipleStacks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AllowMultipleStacks)
		}
		if v, ok := d["sortValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortValue)
		}
	}
	return nil
}

func NewCreateItemModelMasterRequestFromJson(data string) (CreateItemModelMasterRequest, error) {
	req := CreateItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateItemModelMasterRequestFromDict(data map[string]interface{}) CreateItemModelMasterRequest {
	return CreateItemModelMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		InventoryName:       core.CastString(data["inventoryName"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
	}
}

func (p CreateItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"inventoryName":       p.InventoryName,
		"name":                p.Name,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
	}
}

func (p CreateItemModelMasterRequest) Pointer() *CreateItemModelMasterRequest {
	return &p
}

type GetItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemModelMasterRequest{}
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
		*p = GetItemModelMasterRequest{}
	} else {
		*p = GetItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetItemModelMasterRequestFromJson(data string) (GetItemModelMasterRequest, error) {
	req := GetItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetItemModelMasterRequestFromDict(data map[string]interface{}) GetItemModelMasterRequest {
	return GetItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetItemModelMasterRequest) Pointer() *GetItemModelMasterRequest {
	return &p
}

type UpdateItemModelMasterRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	InventoryName       *string `json:"inventoryName"`
	ItemName            *string `json:"itemName"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func (p *UpdateItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateItemModelMasterRequest{}
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
		*p = UpdateItemModelMasterRequest{}
	} else {
		*p = UpdateItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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
		if v, ok := d["stackingLimit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StackingLimit)
		}
		if v, ok := d["allowMultipleStacks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AllowMultipleStacks)
		}
		if v, ok := d["sortValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortValue)
		}
	}
	return nil
}

func NewUpdateItemModelMasterRequestFromJson(data string) (UpdateItemModelMasterRequest, error) {
	req := UpdateItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateItemModelMasterRequestFromDict(data map[string]interface{}) UpdateItemModelMasterRequest {
	return UpdateItemModelMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		InventoryName:       core.CastString(data["inventoryName"]),
		ItemName:            core.CastString(data["itemName"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
	}
}

func (p UpdateItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"inventoryName":       p.InventoryName,
		"itemName":            p.ItemName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
	}
}

func (p UpdateItemModelMasterRequest) Pointer() *UpdateItemModelMasterRequest {
	return &p
}

type DeleteItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *DeleteItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteItemModelMasterRequest{}
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
		*p = DeleteItemModelMasterRequest{}
	} else {
		*p = DeleteItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewDeleteItemModelMasterRequestFromJson(data string) (DeleteItemModelMasterRequest, error) {
	req := DeleteItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteItemModelMasterRequestFromDict(data map[string]interface{}) DeleteItemModelMasterRequest {
	return DeleteItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteItemModelMasterRequest) Pointer() *DeleteItemModelMasterRequest {
	return &p
}

type DescribeItemModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DescribeItemModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeItemModelsRequest{}
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
		*p = DescribeItemModelsRequest{}
	} else {
		*p = DescribeItemModelsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDescribeItemModelsRequestFromJson(data string) (DescribeItemModelsRequest, error) {
	req := DescribeItemModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeItemModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeItemModelsRequestFromDict(data map[string]interface{}) DescribeItemModelsRequest {
	return DescribeItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeItemModelsRequest) Pointer() *DescribeItemModelsRequest {
	return &p
}

type GetItemModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetItemModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemModelRequest{}
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
		*p = GetItemModelRequest{}
	} else {
		*p = GetItemModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetItemModelRequestFromJson(data string) (GetItemModelRequest, error) {
	req := GetItemModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemModelRequest{}, err
	}
	return req, nil
}

func NewGetItemModelRequestFromDict(data map[string]interface{}) GetItemModelRequest {
	return GetItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetItemModelRequest) Pointer() *GetItemModelRequest {
	return &p
}

type DescribeSimpleInventoryModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSimpleInventoryModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleInventoryModelMastersRequest{}
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
		*p = DescribeSimpleInventoryModelMastersRequest{}
	} else {
		*p = DescribeSimpleInventoryModelMastersRequest{}
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

func NewDescribeSimpleInventoryModelMastersRequestFromJson(data string) (DescribeSimpleInventoryModelMastersRequest, error) {
	req := DescribeSimpleInventoryModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleInventoryModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeSimpleInventoryModelMastersRequest {
	return DescribeSimpleInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleInventoryModelMastersRequest) Pointer() *DescribeSimpleInventoryModelMastersRequest {
	return &p
}

type CreateSimpleInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateSimpleInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateSimpleInventoryModelMasterRequest{}
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
		*p = CreateSimpleInventoryModelMasterRequest{}
	} else {
		*p = CreateSimpleInventoryModelMasterRequest{}
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

func NewCreateSimpleInventoryModelMasterRequestFromJson(data string) (CreateSimpleInventoryModelMasterRequest, error) {
	req := CreateSimpleInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateSimpleInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateSimpleInventoryModelMasterRequest {
	return CreateSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateSimpleInventoryModelMasterRequest) Pointer() *CreateSimpleInventoryModelMasterRequest {
	return &p
}

type GetSimpleInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetSimpleInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleInventoryModelMasterRequest{}
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
		*p = GetSimpleInventoryModelMasterRequest{}
	} else {
		*p = GetSimpleInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleInventoryModelMasterRequestFromJson(data string) (GetSimpleInventoryModelMasterRequest, error) {
	req := GetSimpleInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) GetSimpleInventoryModelMasterRequest {
	return GetSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetSimpleInventoryModelMasterRequest) Pointer() *GetSimpleInventoryModelMasterRequest {
	return &p
}

type UpdateSimpleInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateSimpleInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateSimpleInventoryModelMasterRequest{}
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
		*p = UpdateSimpleInventoryModelMasterRequest{}
	} else {
		*p = UpdateSimpleInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewUpdateSimpleInventoryModelMasterRequestFromJson(data string) (UpdateSimpleInventoryModelMasterRequest, error) {
	req := UpdateSimpleInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateSimpleInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateSimpleInventoryModelMasterRequest {
	return UpdateSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateSimpleInventoryModelMasterRequest) Pointer() *UpdateSimpleInventoryModelMasterRequest {
	return &p
}

type DeleteSimpleInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DeleteSimpleInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSimpleInventoryModelMasterRequest{}
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
		*p = DeleteSimpleInventoryModelMasterRequest{}
	} else {
		*p = DeleteSimpleInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSimpleInventoryModelMasterRequestFromJson(data string) (DeleteSimpleInventoryModelMasterRequest, error) {
	req := DeleteSimpleInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSimpleInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteSimpleInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteSimpleInventoryModelMasterRequest {
	return DeleteSimpleInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteSimpleInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteSimpleInventoryModelMasterRequest) Pointer() *DeleteSimpleInventoryModelMasterRequest {
	return &p
}

type DescribeSimpleInventoryModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeSimpleInventoryModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleInventoryModelsRequest{}
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
		*p = DescribeSimpleInventoryModelsRequest{}
	} else {
		*p = DescribeSimpleInventoryModelsRequest{}
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

func NewDescribeSimpleInventoryModelsRequestFromJson(data string) (DescribeSimpleInventoryModelsRequest, error) {
	req := DescribeSimpleInventoryModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleInventoryModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleInventoryModelsRequestFromDict(data map[string]interface{}) DescribeSimpleInventoryModelsRequest {
	return DescribeSimpleInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeSimpleInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeSimpleInventoryModelsRequest) Pointer() *DescribeSimpleInventoryModelsRequest {
	return &p
}

type GetSimpleInventoryModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetSimpleInventoryModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleInventoryModelRequest{}
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
		*p = GetSimpleInventoryModelRequest{}
	} else {
		*p = GetSimpleInventoryModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleInventoryModelRequestFromJson(data string) (GetSimpleInventoryModelRequest, error) {
	req := GetSimpleInventoryModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleInventoryModelRequest{}, err
	}
	return req, nil
}

func NewGetSimpleInventoryModelRequestFromDict(data map[string]interface{}) GetSimpleInventoryModelRequest {
	return GetSimpleInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetSimpleInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetSimpleInventoryModelRequest) Pointer() *GetSimpleInventoryModelRequest {
	return &p
}

type DescribeSimpleItemModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSimpleItemModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleItemModelMastersRequest{}
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
		*p = DescribeSimpleItemModelMastersRequest{}
	} else {
		*p = DescribeSimpleItemModelMastersRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeSimpleItemModelMastersRequestFromJson(data string) (DescribeSimpleItemModelMastersRequest, error) {
	req := DescribeSimpleItemModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleItemModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleItemModelMastersRequestFromDict(data map[string]interface{}) DescribeSimpleItemModelMastersRequest {
	return DescribeSimpleItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleItemModelMastersRequest) Pointer() *DescribeSimpleItemModelMastersRequest {
	return &p
}

type CreateSimpleItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateSimpleItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateSimpleItemModelMasterRequest{}
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
		*p = CreateSimpleItemModelMasterRequest{}
	} else {
		*p = CreateSimpleItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewCreateSimpleItemModelMasterRequestFromJson(data string) (CreateSimpleItemModelMasterRequest, error) {
	req := CreateSimpleItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateSimpleItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateSimpleItemModelMasterRequestFromDict(data map[string]interface{}) CreateSimpleItemModelMasterRequest {
	return CreateSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateSimpleItemModelMasterRequest) Pointer() *CreateSimpleItemModelMasterRequest {
	return &p
}

type GetSimpleItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetSimpleItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemModelMasterRequest{}
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
		*p = GetSimpleItemModelMasterRequest{}
	} else {
		*p = GetSimpleItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleItemModelMasterRequestFromJson(data string) (GetSimpleItemModelMasterRequest, error) {
	req := GetSimpleItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemModelMasterRequestFromDict(data map[string]interface{}) GetSimpleItemModelMasterRequest {
	return GetSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemModelMasterRequest) Pointer() *GetSimpleItemModelMasterRequest {
	return &p
}

type UpdateSimpleItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateSimpleItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateSimpleItemModelMasterRequest{}
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
		*p = UpdateSimpleItemModelMasterRequest{}
	} else {
		*p = UpdateSimpleItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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

func NewUpdateSimpleItemModelMasterRequestFromJson(data string) (UpdateSimpleItemModelMasterRequest, error) {
	req := UpdateSimpleItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateSimpleItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateSimpleItemModelMasterRequestFromDict(data map[string]interface{}) UpdateSimpleItemModelMasterRequest {
	return UpdateSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateSimpleItemModelMasterRequest) Pointer() *UpdateSimpleItemModelMasterRequest {
	return &p
}

type DeleteSimpleItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *DeleteSimpleItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSimpleItemModelMasterRequest{}
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
		*p = DeleteSimpleItemModelMasterRequest{}
	} else {
		*p = DeleteSimpleItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSimpleItemModelMasterRequestFromJson(data string) (DeleteSimpleItemModelMasterRequest, error) {
	req := DeleteSimpleItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSimpleItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteSimpleItemModelMasterRequestFromDict(data map[string]interface{}) DeleteSimpleItemModelMasterRequest {
	return DeleteSimpleItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteSimpleItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteSimpleItemModelMasterRequest) Pointer() *DeleteSimpleItemModelMasterRequest {
	return &p
}

type DescribeSimpleItemModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DescribeSimpleItemModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleItemModelsRequest{}
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
		*p = DescribeSimpleItemModelsRequest{}
	} else {
		*p = DescribeSimpleItemModelsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDescribeSimpleItemModelsRequestFromJson(data string) (DescribeSimpleItemModelsRequest, error) {
	req := DescribeSimpleItemModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleItemModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleItemModelsRequestFromDict(data map[string]interface{}) DescribeSimpleItemModelsRequest {
	return DescribeSimpleItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeSimpleItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeSimpleItemModelsRequest) Pointer() *DescribeSimpleItemModelsRequest {
	return &p
}

type GetSimpleItemModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetSimpleItemModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemModelRequest{}
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
		*p = GetSimpleItemModelRequest{}
	} else {
		*p = GetSimpleItemModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleItemModelRequestFromJson(data string) (GetSimpleItemModelRequest, error) {
	req := GetSimpleItemModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemModelRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemModelRequestFromDict(data map[string]interface{}) GetSimpleItemModelRequest {
	return GetSimpleItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemModelRequest) Pointer() *GetSimpleItemModelRequest {
	return &p
}

type DescribeBigInventoryModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeBigInventoryModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigInventoryModelMastersRequest{}
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
		*p = DescribeBigInventoryModelMastersRequest{}
	} else {
		*p = DescribeBigInventoryModelMastersRequest{}
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

func NewDescribeBigInventoryModelMastersRequestFromJson(data string) (DescribeBigInventoryModelMastersRequest, error) {
	req := DescribeBigInventoryModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigInventoryModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeBigInventoryModelMastersRequestFromDict(data map[string]interface{}) DescribeBigInventoryModelMastersRequest {
	return DescribeBigInventoryModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigInventoryModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigInventoryModelMastersRequest) Pointer() *DescribeBigInventoryModelMastersRequest {
	return &p
}

type CreateBigInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateBigInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateBigInventoryModelMasterRequest{}
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
		*p = CreateBigInventoryModelMasterRequest{}
	} else {
		*p = CreateBigInventoryModelMasterRequest{}
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

func NewCreateBigInventoryModelMasterRequestFromJson(data string) (CreateBigInventoryModelMasterRequest, error) {
	req := CreateBigInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateBigInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateBigInventoryModelMasterRequestFromDict(data map[string]interface{}) CreateBigInventoryModelMasterRequest {
	return CreateBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateBigInventoryModelMasterRequest) Pointer() *CreateBigInventoryModelMasterRequest {
	return &p
}

type GetBigInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetBigInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigInventoryModelMasterRequest{}
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
		*p = GetBigInventoryModelMasterRequest{}
	} else {
		*p = GetBigInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetBigInventoryModelMasterRequestFromJson(data string) (GetBigInventoryModelMasterRequest, error) {
	req := GetBigInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetBigInventoryModelMasterRequestFromDict(data map[string]interface{}) GetBigInventoryModelMasterRequest {
	return GetBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetBigInventoryModelMasterRequest) Pointer() *GetBigInventoryModelMasterRequest {
	return &p
}

type UpdateBigInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateBigInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateBigInventoryModelMasterRequest{}
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
		*p = UpdateBigInventoryModelMasterRequest{}
	} else {
		*p = UpdateBigInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewUpdateBigInventoryModelMasterRequestFromJson(data string) (UpdateBigInventoryModelMasterRequest, error) {
	req := UpdateBigInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateBigInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateBigInventoryModelMasterRequestFromDict(data map[string]interface{}) UpdateBigInventoryModelMasterRequest {
	return UpdateBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateBigInventoryModelMasterRequest) Pointer() *UpdateBigInventoryModelMasterRequest {
	return &p
}

type DeleteBigInventoryModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DeleteBigInventoryModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteBigInventoryModelMasterRequest{}
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
		*p = DeleteBigInventoryModelMasterRequest{}
	} else {
		*p = DeleteBigInventoryModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDeleteBigInventoryModelMasterRequestFromJson(data string) (DeleteBigInventoryModelMasterRequest, error) {
	req := DeleteBigInventoryModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteBigInventoryModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteBigInventoryModelMasterRequestFromDict(data map[string]interface{}) DeleteBigInventoryModelMasterRequest {
	return DeleteBigInventoryModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DeleteBigInventoryModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DeleteBigInventoryModelMasterRequest) Pointer() *DeleteBigInventoryModelMasterRequest {
	return &p
}

type DescribeBigInventoryModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeBigInventoryModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigInventoryModelsRequest{}
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
		*p = DescribeBigInventoryModelsRequest{}
	} else {
		*p = DescribeBigInventoryModelsRequest{}
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

func NewDescribeBigInventoryModelsRequestFromJson(data string) (DescribeBigInventoryModelsRequest, error) {
	req := DescribeBigInventoryModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigInventoryModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeBigInventoryModelsRequestFromDict(data map[string]interface{}) DescribeBigInventoryModelsRequest {
	return DescribeBigInventoryModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeBigInventoryModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeBigInventoryModelsRequest) Pointer() *DescribeBigInventoryModelsRequest {
	return &p
}

type GetBigInventoryModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *GetBigInventoryModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigInventoryModelRequest{}
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
		*p = GetBigInventoryModelRequest{}
	} else {
		*p = GetBigInventoryModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewGetBigInventoryModelRequestFromJson(data string) (GetBigInventoryModelRequest, error) {
	req := GetBigInventoryModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigInventoryModelRequest{}, err
	}
	return req, nil
}

func NewGetBigInventoryModelRequestFromDict(data map[string]interface{}) GetBigInventoryModelRequest {
	return GetBigInventoryModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p GetBigInventoryModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p GetBigInventoryModelRequest) Pointer() *GetBigInventoryModelRequest {
	return &p
}

type DescribeBigItemModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeBigItemModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigItemModelMastersRequest{}
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
		*p = DescribeBigItemModelMastersRequest{}
	} else {
		*p = DescribeBigItemModelMastersRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeBigItemModelMastersRequestFromJson(data string) (DescribeBigItemModelMastersRequest, error) {
	req := DescribeBigItemModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigItemModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeBigItemModelMastersRequestFromDict(data map[string]interface{}) DescribeBigItemModelMastersRequest {
	return DescribeBigItemModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigItemModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigItemModelMastersRequest) Pointer() *DescribeBigItemModelMastersRequest {
	return &p
}

type CreateBigItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *CreateBigItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateBigItemModelMasterRequest{}
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
		*p = CreateBigItemModelMasterRequest{}
	} else {
		*p = CreateBigItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewCreateBigItemModelMasterRequestFromJson(data string) (CreateBigItemModelMasterRequest, error) {
	req := CreateBigItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateBigItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateBigItemModelMasterRequestFromDict(data map[string]interface{}) CreateBigItemModelMasterRequest {
	return CreateBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p CreateBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p CreateBigItemModelMasterRequest) Pointer() *CreateBigItemModelMasterRequest {
	return &p
}

type GetBigItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetBigItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigItemModelMasterRequest{}
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
		*p = GetBigItemModelMasterRequest{}
	} else {
		*p = GetBigItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetBigItemModelMasterRequestFromJson(data string) (GetBigItemModelMasterRequest, error) {
	req := GetBigItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetBigItemModelMasterRequestFromDict(data map[string]interface{}) GetBigItemModelMasterRequest {
	return GetBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemModelMasterRequest) Pointer() *GetBigItemModelMasterRequest {
	return &p
}

type UpdateBigItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
}

func (p *UpdateBigItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateBigItemModelMasterRequest{}
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
		*p = UpdateBigItemModelMasterRequest{}
	} else {
		*p = UpdateBigItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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

func NewUpdateBigItemModelMasterRequestFromJson(data string) (UpdateBigItemModelMasterRequest, error) {
	req := UpdateBigItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateBigItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateBigItemModelMasterRequestFromDict(data map[string]interface{}) UpdateBigItemModelMasterRequest {
	return UpdateBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p UpdateBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
		"description":   p.Description,
		"metadata":      p.Metadata,
	}
}

func (p UpdateBigItemModelMasterRequest) Pointer() *UpdateBigItemModelMasterRequest {
	return &p
}

type DeleteBigItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *DeleteBigItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteBigItemModelMasterRequest{}
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
		*p = DeleteBigItemModelMasterRequest{}
	} else {
		*p = DeleteBigItemModelMasterRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewDeleteBigItemModelMasterRequestFromJson(data string) (DeleteBigItemModelMasterRequest, error) {
	req := DeleteBigItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteBigItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteBigItemModelMasterRequestFromDict(data map[string]interface{}) DeleteBigItemModelMasterRequest {
	return DeleteBigItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p DeleteBigItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p DeleteBigItemModelMasterRequest) Pointer() *DeleteBigItemModelMasterRequest {
	return &p
}

type DescribeBigItemModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
}

func (p *DescribeBigItemModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigItemModelsRequest{}
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
		*p = DescribeBigItemModelsRequest{}
	} else {
		*p = DescribeBigItemModelsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
	}
	return nil
}

func NewDescribeBigItemModelsRequestFromJson(data string) (DescribeBigItemModelsRequest, error) {
	req := DescribeBigItemModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigItemModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeBigItemModelsRequestFromDict(data map[string]interface{}) DescribeBigItemModelsRequest {
	return DescribeBigItemModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
	}
}

func (p DescribeBigItemModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
	}
}

func (p DescribeBigItemModelsRequest) Pointer() *DescribeBigItemModelsRequest {
	return &p
}

type GetBigItemModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	ItemName        *string `json:"itemName"`
}

func (p *GetBigItemModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigItemModelRequest{}
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
		*p = GetBigItemModelRequest{}
	} else {
		*p = GetBigItemModelRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetBigItemModelRequestFromJson(data string) (GetBigItemModelRequest, error) {
	req := GetBigItemModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigItemModelRequest{}, err
	}
	return req, nil
}

func NewGetBigItemModelRequestFromDict(data map[string]interface{}) GetBigItemModelRequest {
	return GetBigItemModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemModelRequest) Pointer() *GetBigItemModelRequest {
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

type GetCurrentItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentItemModelMasterRequest{}
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
		*p = GetCurrentItemModelMasterRequest{}
	} else {
		*p = GetCurrentItemModelMasterRequest{}
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

func NewGetCurrentItemModelMasterRequestFromJson(data string) (GetCurrentItemModelMasterRequest, error) {
	req := GetCurrentItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentItemModelMasterRequestFromDict(data map[string]interface{}) GetCurrentItemModelMasterRequest {
	return GetCurrentItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentItemModelMasterRequest) Pointer() *GetCurrentItemModelMasterRequest {
	return &p
}

type UpdateCurrentItemModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentItemModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentItemModelMasterRequest{}
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
		*p = UpdateCurrentItemModelMasterRequest{}
	} else {
		*p = UpdateCurrentItemModelMasterRequest{}
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

func NewUpdateCurrentItemModelMasterRequestFromJson(data string) (UpdateCurrentItemModelMasterRequest, error) {
	req := UpdateCurrentItemModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentItemModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentItemModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterRequest {
	return UpdateCurrentItemModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentItemModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentItemModelMasterRequest) Pointer() *UpdateCurrentItemModelMasterRequest {
	return &p
}

type UpdateCurrentItemModelMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentItemModelMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentItemModelMasterFromGitHubRequest{}
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
		*p = UpdateCurrentItemModelMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentItemModelMasterFromGitHubRequest{}
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

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromJson(data string) (UpdateCurrentItemModelMasterFromGitHubRequest, error) {
	req := UpdateCurrentItemModelMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentItemModelMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentItemModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterFromGitHubRequest {
	return UpdateCurrentItemModelMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"checkoutSetting": func() map[string]interface{} {
			if p.CheckoutSetting == nil {
				return nil
			}
			return p.CheckoutSetting.ToDict()
		}(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubRequest) Pointer() *UpdateCurrentItemModelMasterFromGitHubRequest {
	return &p
}

type DescribeInventoriesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeInventoriesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoriesRequest{}
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
		*p = DescribeInventoriesRequest{}
	} else {
		*p = DescribeInventoriesRequest{}
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

func NewDescribeInventoriesRequestFromJson(data string) (DescribeInventoriesRequest, error) {
	req := DescribeInventoriesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoriesRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoriesRequestFromDict(data map[string]interface{}) DescribeInventoriesRequest {
	return DescribeInventoriesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInventoriesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInventoriesRequest) Pointer() *DescribeInventoriesRequest {
	return &p
}

type DescribeInventoriesByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeInventoriesByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInventoriesByUserIdRequest{}
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
		*p = DescribeInventoriesByUserIdRequest{}
	} else {
		*p = DescribeInventoriesByUserIdRequest{}
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

func NewDescribeInventoriesByUserIdRequestFromJson(data string) (DescribeInventoriesByUserIdRequest, error) {
	req := DescribeInventoriesByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInventoriesByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeInventoriesByUserIdRequestFromDict(data map[string]interface{}) DescribeInventoriesByUserIdRequest {
	return DescribeInventoriesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeInventoriesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeInventoriesByUserIdRequest) Pointer() *DescribeInventoriesByUserIdRequest {
	return &p
}

type GetInventoryRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
}

func (p *GetInventoryRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInventoryRequest{}
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
		*p = GetInventoryRequest{}
	} else {
		*p = GetInventoryRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
	}
	return nil
}

func NewGetInventoryRequestFromJson(data string) (GetInventoryRequest, error) {
	req := GetInventoryRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInventoryRequest{}, err
	}
	return req, nil
}

func NewGetInventoryRequestFromDict(data map[string]interface{}) GetInventoryRequest {
	return GetInventoryRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetInventoryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetInventoryRequest) Pointer() *GetInventoryRequest {
	return &p
}

type GetInventoryByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetInventoryByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInventoryByUserIdRequest{}
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
		*p = GetInventoryByUserIdRequest{}
	} else {
		*p = GetInventoryByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewGetInventoryByUserIdRequestFromJson(data string) (GetInventoryByUserIdRequest, error) {
	req := GetInventoryByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInventoryByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetInventoryByUserIdRequestFromDict(data map[string]interface{}) GetInventoryByUserIdRequest {
	return GetInventoryByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetInventoryByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetInventoryByUserIdRequest) Pointer() *GetInventoryByUserIdRequest {
	return &p
}

type AddCapacityByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	AddCapacityValue   *int32  `json:"addCapacityValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AddCapacityByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddCapacityByUserIdRequest{}
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
		*p = AddCapacityByUserIdRequest{}
	} else {
		*p = AddCapacityByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["addCapacityValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AddCapacityValue)
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

func NewAddCapacityByUserIdRequestFromJson(data string) (AddCapacityByUserIdRequest, error) {
	req := AddCapacityByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddCapacityByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddCapacityByUserIdRequestFromDict(data map[string]interface{}) AddCapacityByUserIdRequest {
	return AddCapacityByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		UserId:           core.CastString(data["userId"]),
		AddCapacityValue: core.CastInt32(data["addCapacityValue"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p AddCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"userId":           p.UserId,
		"addCapacityValue": p.AddCapacityValue,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p AddCapacityByUserIdRequest) Pointer() *AddCapacityByUserIdRequest {
	return &p
}

type SetCapacityByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	NewCapacityValue   *int32  `json:"newCapacityValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *SetCapacityByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetCapacityByUserIdRequest{}
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
		*p = SetCapacityByUserIdRequest{}
	} else {
		*p = SetCapacityByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["newCapacityValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NewCapacityValue)
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

func NewSetCapacityByUserIdRequestFromJson(data string) (SetCapacityByUserIdRequest, error) {
	req := SetCapacityByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetCapacityByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetCapacityByUserIdRequestFromDict(data map[string]interface{}) SetCapacityByUserIdRequest {
	return SetCapacityByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		UserId:           core.CastString(data["userId"]),
		NewCapacityValue: core.CastInt32(data["newCapacityValue"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"userId":           p.UserId,
		"newCapacityValue": p.NewCapacityValue,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p SetCapacityByUserIdRequest) Pointer() *SetCapacityByUserIdRequest {
	return &p
}

type DeleteInventoryByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteInventoryByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteInventoryByUserIdRequest{}
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
		*p = DeleteInventoryByUserIdRequest{}
	} else {
		*p = DeleteInventoryByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDeleteInventoryByUserIdRequestFromJson(data string) (DeleteInventoryByUserIdRequest, error) {
	req := DeleteInventoryByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteInventoryByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteInventoryByUserIdRequestFromDict(data map[string]interface{}) DeleteInventoryByUserIdRequest {
	return DeleteInventoryByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteInventoryByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteInventoryByUserIdRequest) Pointer() *DeleteInventoryByUserIdRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	InventoryName                   *string `json:"inventoryName"`
	VerifyType                      *string `json:"verifyType"`
	CurrentInventoryMaxCapacity     *int32  `json:"currentInventoryMaxCapacity"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifyInventoryCurrentMaxCapacityRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyInventoryCurrentMaxCapacityRequest{}
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
		*p = VerifyInventoryCurrentMaxCapacityRequest{}
	} else {
		*p = VerifyInventoryCurrentMaxCapacityRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["currentInventoryMaxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentInventoryMaxCapacity)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyInventoryCurrentMaxCapacityRequestFromJson(data string) (VerifyInventoryCurrentMaxCapacityRequest, error) {
	req := VerifyInventoryCurrentMaxCapacityRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyInventoryCurrentMaxCapacityRequest{}, err
	}
	return req, nil
}

func NewVerifyInventoryCurrentMaxCapacityRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityRequest {
	return VerifyInventoryCurrentMaxCapacityRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		AccessToken:                     core.CastString(data["accessToken"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		CurrentInventoryMaxCapacity:     core.CastInt32(data["currentInventoryMaxCapacity"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"inventoryName":                   p.InventoryName,
		"verifyType":                      p.VerifyType,
		"currentInventoryMaxCapacity":     p.CurrentInventoryMaxCapacity,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyInventoryCurrentMaxCapacityRequest) Pointer() *VerifyInventoryCurrentMaxCapacityRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByUserIdRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	InventoryName                   *string `json:"inventoryName"`
	VerifyType                      *string `json:"verifyType"`
	CurrentInventoryMaxCapacity     *int32  `json:"currentInventoryMaxCapacity"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
}

func (p *VerifyInventoryCurrentMaxCapacityByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyInventoryCurrentMaxCapacityByUserIdRequest{}
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
		*p = VerifyInventoryCurrentMaxCapacityByUserIdRequest{}
	} else {
		*p = VerifyInventoryCurrentMaxCapacityByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["currentInventoryMaxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentInventoryMaxCapacity)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
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

func NewVerifyInventoryCurrentMaxCapacityByUserIdRequestFromJson(data string) (VerifyInventoryCurrentMaxCapacityByUserIdRequest, error) {
	req := VerifyInventoryCurrentMaxCapacityByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyInventoryCurrentMaxCapacityByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyInventoryCurrentMaxCapacityByUserIdRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByUserIdRequest {
	return VerifyInventoryCurrentMaxCapacityByUserIdRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		UserId:                          core.CastString(data["userId"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		CurrentInventoryMaxCapacity:     core.CastInt32(data["currentInventoryMaxCapacity"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
		TimeOffsetToken:                 core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"inventoryName":                   p.InventoryName,
		"verifyType":                      p.VerifyType,
		"currentInventoryMaxCapacity":     p.CurrentInventoryMaxCapacity,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdRequest) Pointer() *VerifyInventoryCurrentMaxCapacityByUserIdRequest {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyInventoryCurrentMaxCapacityByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyInventoryCurrentMaxCapacityByStampTaskRequest{}
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
		*p = VerifyInventoryCurrentMaxCapacityByStampTaskRequest{}
	} else {
		*p = VerifyInventoryCurrentMaxCapacityByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskRequestFromJson(data string) (VerifyInventoryCurrentMaxCapacityByStampTaskRequest, error) {
	req := VerifyInventoryCurrentMaxCapacityByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyInventoryCurrentMaxCapacityByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskRequestFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByStampTaskRequest {
	return VerifyInventoryCurrentMaxCapacityByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyInventoryCurrentMaxCapacityByStampTaskRequest) Pointer() *VerifyInventoryCurrentMaxCapacityByStampTaskRequest {
	return &p
}

type AddCapacityByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AddCapacityByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddCapacityByStampSheetRequest{}
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
		*p = AddCapacityByStampSheetRequest{}
	} else {
		*p = AddCapacityByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAddCapacityByStampSheetRequestFromJson(data string) (AddCapacityByStampSheetRequest, error) {
	req := AddCapacityByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddCapacityByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAddCapacityByStampSheetRequestFromDict(data map[string]interface{}) AddCapacityByStampSheetRequest {
	return AddCapacityByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddCapacityByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddCapacityByStampSheetRequest) Pointer() *AddCapacityByStampSheetRequest {
	return &p
}

type SetCapacityByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetCapacityByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetCapacityByStampSheetRequest{}
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
		*p = SetCapacityByStampSheetRequest{}
	} else {
		*p = SetCapacityByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewSetCapacityByStampSheetRequestFromJson(data string) (SetCapacityByStampSheetRequest, error) {
	req := SetCapacityByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetCapacityByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetCapacityByStampSheetRequestFromDict(data map[string]interface{}) SetCapacityByStampSheetRequest {
	return SetCapacityByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetCapacityByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetCapacityByStampSheetRequest) Pointer() *SetCapacityByStampSheetRequest {
	return &p
}

type DescribeItemSetsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeItemSetsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeItemSetsRequest{}
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
		*p = DescribeItemSetsRequest{}
	} else {
		*p = DescribeItemSetsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeItemSetsRequestFromJson(data string) (DescribeItemSetsRequest, error) {
	req := DescribeItemSetsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeItemSetsRequest{}, err
	}
	return req, nil
}

func NewDescribeItemSetsRequestFromDict(data map[string]interface{}) DescribeItemSetsRequest {
	return DescribeItemSetsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeItemSetsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeItemSetsRequest) Pointer() *DescribeItemSetsRequest {
	return &p
}

type DescribeItemSetsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeItemSetsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeItemSetsByUserIdRequest{}
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
		*p = DescribeItemSetsByUserIdRequest{}
	} else {
		*p = DescribeItemSetsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeItemSetsByUserIdRequestFromJson(data string) (DescribeItemSetsByUserIdRequest, error) {
	req := DescribeItemSetsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeItemSetsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeItemSetsByUserIdRequestFromDict(data map[string]interface{}) DescribeItemSetsByUserIdRequest {
	return DescribeItemSetsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeItemSetsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeItemSetsByUserIdRequest) Pointer() *DescribeItemSetsByUserIdRequest {
	return &p
}

type GetItemSetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
}

func (p *GetItemSetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemSetRequest{}
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
		*p = GetItemSetRequest{}
	} else {
		*p = GetItemSetRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
	}
	return nil
}

func NewGetItemSetRequestFromJson(data string) (GetItemSetRequest, error) {
	req := GetItemSetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemSetRequest{}, err
	}
	return req, nil
}

func NewGetItemSetRequestFromDict(data map[string]interface{}) GetItemSetRequest {
	return GetItemSetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p GetItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p GetItemSetRequest) Pointer() *GetItemSetRequest {
	return &p
}

type GetItemSetByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetItemSetByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemSetByUserIdRequest{}
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
		*p = GetItemSetByUserIdRequest{}
	} else {
		*p = GetItemSetByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
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

func NewGetItemSetByUserIdRequestFromJson(data string) (GetItemSetByUserIdRequest, error) {
	req := GetItemSetByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemSetByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetItemSetByUserIdRequestFromDict(data map[string]interface{}) GetItemSetByUserIdRequest {
	return GetItemSetByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetItemSetByUserIdRequest) Pointer() *GetItemSetByUserIdRequest {
	return &p
}

type GetItemWithSignatureRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	KeyId           *string `json:"keyId"`
}

func (p *GetItemWithSignatureRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemWithSignatureRequest{}
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
		*p = GetItemWithSignatureRequest{}
	} else {
		*p = GetItemWithSignatureRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewGetItemWithSignatureRequestFromJson(data string) (GetItemWithSignatureRequest, error) {
	req := GetItemWithSignatureRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemWithSignatureRequest{}, err
	}
	return req, nil
}

func NewGetItemWithSignatureRequestFromDict(data map[string]interface{}) GetItemWithSignatureRequest {
	return GetItemWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetItemWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"keyId":         p.KeyId,
	}
}

func (p GetItemWithSignatureRequest) Pointer() *GetItemWithSignatureRequest {
	return &p
}

type GetItemWithSignatureByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetItemWithSignatureByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetItemWithSignatureByUserIdRequest{}
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
		*p = GetItemWithSignatureByUserIdRequest{}
	} else {
		*p = GetItemWithSignatureByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
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

func NewGetItemWithSignatureByUserIdRequestFromJson(data string) (GetItemWithSignatureByUserIdRequest, error) {
	req := GetItemWithSignatureByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetItemWithSignatureByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetItemWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetItemWithSignatureByUserIdRequest {
	return GetItemWithSignatureByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetItemWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetItemWithSignatureByUserIdRequest) Pointer() *GetItemWithSignatureByUserIdRequest {
	return &p
}

type AcquireItemSetByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	UserId             *string `json:"userId"`
	AcquireCount       *int64  `json:"acquireCount"`
	ExpiresAt          *int64  `json:"expiresAt"`
	CreateNewItemSet   *bool   `json:"createNewItemSet"`
	ItemSetName        *string `json:"itemSetName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AcquireItemSetByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireItemSetByUserIdRequest{}
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
		*p = AcquireItemSetByUserIdRequest{}
	} else {
		*p = AcquireItemSetByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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
		if v, ok := d["acquireCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireCount)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["createNewItemSet"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateNewItemSet)
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
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

func NewAcquireItemSetByUserIdRequestFromJson(data string) (AcquireItemSetByUserIdRequest, error) {
	req := AcquireItemSetByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireItemSetByUserIdRequest{}, err
	}
	return req, nil
}

func NewAcquireItemSetByUserIdRequestFromDict(data map[string]interface{}) AcquireItemSetByUserIdRequest {
	return AcquireItemSetByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		InventoryName:    core.CastString(data["inventoryName"]),
		ItemName:         core.CastString(data["itemName"]),
		UserId:           core.CastString(data["userId"]),
		AcquireCount:     core.CastInt64(data["acquireCount"]),
		ExpiresAt:        core.CastInt64(data["expiresAt"]),
		CreateNewItemSet: core.CastBool(data["createNewItemSet"]),
		ItemSetName:      core.CastString(data["itemSetName"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"inventoryName":    p.InventoryName,
		"itemName":         p.ItemName,
		"userId":           p.UserId,
		"acquireCount":     p.AcquireCount,
		"expiresAt":        p.ExpiresAt,
		"createNewItemSet": p.CreateNewItemSet,
		"itemSetName":      p.ItemSetName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p AcquireItemSetByUserIdRequest) Pointer() *AcquireItemSetByUserIdRequest {
	return &p
}

type AcquireItemSetWithGradeByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	ItemName           *string `json:"itemName"`
	UserId             *string `json:"userId"`
	GradeModelId       *string `json:"gradeModelId"`
	GradeValue         *int64  `json:"gradeValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AcquireItemSetWithGradeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireItemSetWithGradeByUserIdRequest{}
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
		*p = AcquireItemSetWithGradeByUserIdRequest{}
	} else {
		*p = AcquireItemSetWithGradeByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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
		if v, ok := d["gradeModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeModelId)
				}
			}
		}
		if v, ok := d["gradeValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeValue)
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

func NewAcquireItemSetWithGradeByUserIdRequestFromJson(data string) (AcquireItemSetWithGradeByUserIdRequest, error) {
	req := AcquireItemSetWithGradeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireItemSetWithGradeByUserIdRequest{}, err
	}
	return req, nil
}

func NewAcquireItemSetWithGradeByUserIdRequestFromDict(data map[string]interface{}) AcquireItemSetWithGradeByUserIdRequest {
	return AcquireItemSetWithGradeByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		ItemName:        core.CastString(data["itemName"]),
		UserId:          core.CastString(data["userId"]),
		GradeModelId:    core.CastString(data["gradeModelId"]),
		GradeValue:      core.CastInt64(data["gradeValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireItemSetWithGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"itemName":        p.ItemName,
		"userId":          p.UserId,
		"gradeModelId":    p.GradeModelId,
		"gradeValue":      p.GradeValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcquireItemSetWithGradeByUserIdRequest) Pointer() *AcquireItemSetWithGradeByUserIdRequest {
	return &p
}

type ConsumeItemSetRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *int64  `json:"consumeCount"`
	ItemSetName        *string `json:"itemSetName"`
}

func (p *ConsumeItemSetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeItemSetRequest{}
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
		*p = ConsumeItemSetRequest{}
	} else {
		*p = ConsumeItemSetRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["consumeCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeCount)
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
	}
	return nil
}

func NewConsumeItemSetRequestFromJson(data string) (ConsumeItemSetRequest, error) {
	req := ConsumeItemSetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeItemSetRequest{}, err
	}
	return req, nil
}

func NewConsumeItemSetRequestFromDict(data map[string]interface{}) ConsumeItemSetRequest {
	return ConsumeItemSetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastInt64(data["consumeCount"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p ConsumeItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
		"itemSetName":   p.ItemSetName,
	}
}

func (p ConsumeItemSetRequest) Pointer() *ConsumeItemSetRequest {
	return &p
}

type ConsumeItemSetByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *int64  `json:"consumeCount"`
	ItemSetName        *string `json:"itemSetName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *ConsumeItemSetByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeItemSetByUserIdRequest{}
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
		*p = ConsumeItemSetByUserIdRequest{}
	} else {
		*p = ConsumeItemSetByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["consumeCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeCount)
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
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

func NewConsumeItemSetByUserIdRequestFromJson(data string) (ConsumeItemSetByUserIdRequest, error) {
	req := ConsumeItemSetByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeItemSetByUserIdRequest{}, err
	}
	return req, nil
}

func NewConsumeItemSetByUserIdRequestFromDict(data map[string]interface{}) ConsumeItemSetByUserIdRequest {
	return ConsumeItemSetByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ConsumeCount:    core.CastInt64(data["consumeCount"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ConsumeItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"consumeCount":    p.ConsumeCount,
		"itemSetName":     p.ItemSetName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ConsumeItemSetByUserIdRequest) Pointer() *ConsumeItemSetByUserIdRequest {
	return &p
}

type DeleteItemSetByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteItemSetByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteItemSetByUserIdRequest{}
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
		*p = DeleteItemSetByUserIdRequest{}
	} else {
		*p = DeleteItemSetByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
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

func NewDeleteItemSetByUserIdRequestFromJson(data string) (DeleteItemSetByUserIdRequest, error) {
	req := DeleteItemSetByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteItemSetByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteItemSetByUserIdRequestFromDict(data map[string]interface{}) DeleteItemSetByUserIdRequest {
	return DeleteItemSetByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteItemSetByUserIdRequest) Pointer() *DeleteItemSetByUserIdRequest {
	return &p
}

type VerifyItemSetRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	ItemSetName                     *string `json:"itemSetName"`
	Count                           *int64  `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifyItemSetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyItemSetRequest{}
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
		*p = VerifyItemSetRequest{}
	} else {
		*p = VerifyItemSetRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyItemSetRequestFromJson(data string) (VerifyItemSetRequest, error) {
	req := VerifyItemSetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyItemSetRequest{}, err
	}
	return req, nil
}

func NewVerifyItemSetRequestFromDict(data map[string]interface{}) VerifyItemSetRequest {
	return VerifyItemSetRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		AccessToken:                     core.CastString(data["accessToken"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		ItemSetName:                     core.CastString(data["itemSetName"]),
		Count:                           core.CastInt64(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
	}
}

func (p VerifyItemSetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"itemSetName":                     p.ItemSetName,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyItemSetRequest) Pointer() *VerifyItemSetRequest {
	return &p
}

type VerifyItemSetByUserIdRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	ItemSetName                     *string `json:"itemSetName"`
	Count                           *int64  `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
}

func (p *VerifyItemSetByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyItemSetByUserIdRequest{}
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
		*p = VerifyItemSetByUserIdRequest{}
	} else {
		*p = VerifyItemSetByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
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

func NewVerifyItemSetByUserIdRequestFromJson(data string) (VerifyItemSetByUserIdRequest, error) {
	req := VerifyItemSetByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyItemSetByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyItemSetByUserIdRequestFromDict(data map[string]interface{}) VerifyItemSetByUserIdRequest {
	return VerifyItemSetByUserIdRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		UserId:                          core.CastString(data["userId"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		ItemSetName:                     core.CastString(data["itemSetName"]),
		Count:                           core.CastInt64(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
		TimeOffsetToken:                 core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifyItemSetByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"itemSetName":                     p.ItemSetName,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyItemSetByUserIdRequest) Pointer() *VerifyItemSetByUserIdRequest {
	return &p
}

type AcquireItemSetByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AcquireItemSetByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireItemSetByStampSheetRequest{}
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
		*p = AcquireItemSetByStampSheetRequest{}
	} else {
		*p = AcquireItemSetByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAcquireItemSetByStampSheetRequestFromJson(data string) (AcquireItemSetByStampSheetRequest, error) {
	req := AcquireItemSetByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireItemSetByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAcquireItemSetByStampSheetRequestFromDict(data map[string]interface{}) AcquireItemSetByStampSheetRequest {
	return AcquireItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireItemSetByStampSheetRequest) Pointer() *AcquireItemSetByStampSheetRequest {
	return &p
}

type AcquireItemSetWithGradeByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AcquireItemSetWithGradeByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireItemSetWithGradeByStampSheetRequest{}
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
		*p = AcquireItemSetWithGradeByStampSheetRequest{}
	} else {
		*p = AcquireItemSetWithGradeByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAcquireItemSetWithGradeByStampSheetRequestFromJson(data string) (AcquireItemSetWithGradeByStampSheetRequest, error) {
	req := AcquireItemSetWithGradeByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireItemSetWithGradeByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAcquireItemSetWithGradeByStampSheetRequestFromDict(data map[string]interface{}) AcquireItemSetWithGradeByStampSheetRequest {
	return AcquireItemSetWithGradeByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireItemSetWithGradeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireItemSetWithGradeByStampSheetRequest) Pointer() *AcquireItemSetWithGradeByStampSheetRequest {
	return &p
}

type ConsumeItemSetByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *ConsumeItemSetByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeItemSetByStampTaskRequest{}
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
		*p = ConsumeItemSetByStampTaskRequest{}
	} else {
		*p = ConsumeItemSetByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewConsumeItemSetByStampTaskRequestFromJson(data string) (ConsumeItemSetByStampTaskRequest, error) {
	req := ConsumeItemSetByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeItemSetByStampTaskRequest{}, err
	}
	return req, nil
}

func NewConsumeItemSetByStampTaskRequestFromDict(data map[string]interface{}) ConsumeItemSetByStampTaskRequest {
	return ConsumeItemSetByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeItemSetByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeItemSetByStampTaskRequest) Pointer() *ConsumeItemSetByStampTaskRequest {
	return &p
}

type VerifyItemSetByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyItemSetByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyItemSetByStampTaskRequest{}
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
		*p = VerifyItemSetByStampTaskRequest{}
	} else {
		*p = VerifyItemSetByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyItemSetByStampTaskRequestFromJson(data string) (VerifyItemSetByStampTaskRequest, error) {
	req := VerifyItemSetByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyItemSetByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyItemSetByStampTaskRequestFromDict(data map[string]interface{}) VerifyItemSetByStampTaskRequest {
	return VerifyItemSetByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyItemSetByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyItemSetByStampTaskRequest) Pointer() *VerifyItemSetByStampTaskRequest {
	return &p
}

type DescribeReferenceOfRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
}

func (p *DescribeReferenceOfRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeReferenceOfRequest{}
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
		*p = DescribeReferenceOfRequest{}
	} else {
		*p = DescribeReferenceOfRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
	}
	return nil
}

func NewDescribeReferenceOfRequestFromJson(data string) (DescribeReferenceOfRequest, error) {
	req := DescribeReferenceOfRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeReferenceOfRequest{}, err
	}
	return req, nil
}

func NewDescribeReferenceOfRequestFromDict(data map[string]interface{}) DescribeReferenceOfRequest {
	return DescribeReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
	}
}

func (p DescribeReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
	}
}

func (p DescribeReferenceOfRequest) Pointer() *DescribeReferenceOfRequest {
	return &p
}

type DescribeReferenceOfByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeReferenceOfByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeReferenceOfByUserIdRequest{}
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
		*p = DescribeReferenceOfByUserIdRequest{}
	} else {
		*p = DescribeReferenceOfByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
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

func NewDescribeReferenceOfByUserIdRequestFromJson(data string) (DescribeReferenceOfByUserIdRequest, error) {
	req := DescribeReferenceOfByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeReferenceOfByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DescribeReferenceOfByUserIdRequest {
	return DescribeReferenceOfByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeReferenceOfByUserIdRequest) Pointer() *DescribeReferenceOfByUserIdRequest {
	return &p
}

type GetReferenceOfRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	ReferenceOf     *string `json:"referenceOf"`
}

func (p *GetReferenceOfRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetReferenceOfRequest{}
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
		*p = GetReferenceOfRequest{}
	} else {
		*p = GetReferenceOfRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
				}
			}
		}
	}
	return nil
}

func NewGetReferenceOfRequestFromJson(data string) (GetReferenceOfRequest, error) {
	req := GetReferenceOfRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetReferenceOfRequest{}, err
	}
	return req, nil
}

func NewGetReferenceOfRequestFromDict(data map[string]interface{}) GetReferenceOfRequest {
	return GetReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p GetReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p GetReferenceOfRequest) Pointer() *GetReferenceOfRequest {
	return &p
}

type GetReferenceOfByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	ItemSetName     *string `json:"itemSetName"`
	ReferenceOf     *string `json:"referenceOf"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetReferenceOfByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetReferenceOfByUserIdRequest{}
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
		*p = GetReferenceOfByUserIdRequest{}
	} else {
		*p = GetReferenceOfByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
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

func NewGetReferenceOfByUserIdRequestFromJson(data string) (GetReferenceOfByUserIdRequest, error) {
	req := GetReferenceOfByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetReferenceOfByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetReferenceOfByUserIdRequestFromDict(data map[string]interface{}) GetReferenceOfByUserIdRequest {
	return GetReferenceOfByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		ReferenceOf:     core.CastString(data["referenceOf"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"referenceOf":     p.ReferenceOf,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetReferenceOfByUserIdRequest) Pointer() *GetReferenceOfByUserIdRequest {
	return &p
}

type VerifyReferenceOfRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	VerifyType         *string `json:"verifyType"`
}

func (p *VerifyReferenceOfRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReferenceOfRequest{}
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
		*p = VerifyReferenceOfRequest{}
	} else {
		*p = VerifyReferenceOfRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
	}
	return nil
}

func NewVerifyReferenceOfRequestFromJson(data string) (VerifyReferenceOfRequest, error) {
	req := VerifyReferenceOfRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReferenceOfRequest{}, err
	}
	return req, nil
}

func NewVerifyReferenceOfRequestFromDict(data map[string]interface{}) VerifyReferenceOfRequest {
	return VerifyReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
		VerifyType:    core.CastString(data["verifyType"]),
	}
}

func (p VerifyReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
		"verifyType":    p.VerifyType,
	}
}

func (p VerifyReferenceOfRequest) Pointer() *VerifyReferenceOfRequest {
	return &p
}

type VerifyReferenceOfByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	VerifyType         *string `json:"verifyType"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *VerifyReferenceOfByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReferenceOfByUserIdRequest{}
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
		*p = VerifyReferenceOfByUserIdRequest{}
	} else {
		*p = VerifyReferenceOfByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
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

func NewVerifyReferenceOfByUserIdRequestFromJson(data string) (VerifyReferenceOfByUserIdRequest, error) {
	req := VerifyReferenceOfByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReferenceOfByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyReferenceOfByUserIdRequestFromDict(data map[string]interface{}) VerifyReferenceOfByUserIdRequest {
	return VerifyReferenceOfByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		ReferenceOf:     core.CastString(data["referenceOf"]),
		VerifyType:      core.CastString(data["verifyType"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifyReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"referenceOf":     p.ReferenceOf,
		"verifyType":      p.VerifyType,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p VerifyReferenceOfByUserIdRequest) Pointer() *VerifyReferenceOfByUserIdRequest {
	return &p
}

type AddReferenceOfRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func (p *AddReferenceOfRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddReferenceOfRequest{}
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
		*p = AddReferenceOfRequest{}
	} else {
		*p = AddReferenceOfRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
				}
			}
		}
	}
	return nil
}

func NewAddReferenceOfRequestFromJson(data string) (AddReferenceOfRequest, error) {
	req := AddReferenceOfRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddReferenceOfRequest{}, err
	}
	return req, nil
}

func NewAddReferenceOfRequestFromDict(data map[string]interface{}) AddReferenceOfRequest {
	return AddReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p AddReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p AddReferenceOfRequest) Pointer() *AddReferenceOfRequest {
	return &p
}

type AddReferenceOfByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AddReferenceOfByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddReferenceOfByUserIdRequest{}
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
		*p = AddReferenceOfByUserIdRequest{}
	} else {
		*p = AddReferenceOfByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
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

func NewAddReferenceOfByUserIdRequestFromJson(data string) (AddReferenceOfByUserIdRequest, error) {
	req := AddReferenceOfByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddReferenceOfByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddReferenceOfByUserIdRequestFromDict(data map[string]interface{}) AddReferenceOfByUserIdRequest {
	return AddReferenceOfByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		ReferenceOf:     core.CastString(data["referenceOf"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AddReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"referenceOf":     p.ReferenceOf,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddReferenceOfByUserIdRequest) Pointer() *AddReferenceOfByUserIdRequest {
	return &p
}

type DeleteReferenceOfRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
}

func (p *DeleteReferenceOfRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteReferenceOfRequest{}
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
		*p = DeleteReferenceOfRequest{}
	} else {
		*p = DeleteReferenceOfRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
				}
			}
		}
	}
	return nil
}

func NewDeleteReferenceOfRequestFromJson(data string) (DeleteReferenceOfRequest, error) {
	req := DeleteReferenceOfRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteReferenceOfRequest{}, err
	}
	return req, nil
}

func NewDeleteReferenceOfRequestFromDict(data map[string]interface{}) DeleteReferenceOfRequest {
	return DeleteReferenceOfRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ItemSetName:   core.CastString(data["itemSetName"]),
		ReferenceOf:   core.CastString(data["referenceOf"]),
	}
}

func (p DeleteReferenceOfRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"itemSetName":   p.ItemSetName,
		"referenceOf":   p.ReferenceOf,
	}
}

func (p DeleteReferenceOfRequest) Pointer() *DeleteReferenceOfRequest {
	return &p
}

type DeleteReferenceOfByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ItemSetName        *string `json:"itemSetName"`
	ReferenceOf        *string `json:"referenceOf"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteReferenceOfByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteReferenceOfByUserIdRequest{}
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
		*p = DeleteReferenceOfByUserIdRequest{}
	} else {
		*p = DeleteReferenceOfByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["itemSetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetName)
				}
			}
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOf = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOf = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOf = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOf = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOf)
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

func NewDeleteReferenceOfByUserIdRequestFromJson(data string) (DeleteReferenceOfByUserIdRequest, error) {
	req := DeleteReferenceOfByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteReferenceOfByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteReferenceOfByUserIdRequestFromDict(data map[string]interface{}) DeleteReferenceOfByUserIdRequest {
	return DeleteReferenceOfByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ItemSetName:     core.CastString(data["itemSetName"]),
		ReferenceOf:     core.CastString(data["referenceOf"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteReferenceOfByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"itemSetName":     p.ItemSetName,
		"referenceOf":     p.ReferenceOf,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteReferenceOfByUserIdRequest) Pointer() *DeleteReferenceOfByUserIdRequest {
	return &p
}

type AddReferenceOfItemSetByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AddReferenceOfItemSetByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddReferenceOfItemSetByStampSheetRequest{}
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
		*p = AddReferenceOfItemSetByStampSheetRequest{}
	} else {
		*p = AddReferenceOfItemSetByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAddReferenceOfItemSetByStampSheetRequestFromJson(data string) (AddReferenceOfItemSetByStampSheetRequest, error) {
	req := AddReferenceOfItemSetByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddReferenceOfItemSetByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAddReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) AddReferenceOfItemSetByStampSheetRequest {
	return AddReferenceOfItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AddReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddReferenceOfItemSetByStampSheetRequest) Pointer() *AddReferenceOfItemSetByStampSheetRequest {
	return &p
}

type DeleteReferenceOfItemSetByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *DeleteReferenceOfItemSetByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteReferenceOfItemSetByStampSheetRequest{}
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
		*p = DeleteReferenceOfItemSetByStampSheetRequest{}
	} else {
		*p = DeleteReferenceOfItemSetByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromJson(data string) (DeleteReferenceOfItemSetByStampSheetRequest, error) {
	req := DeleteReferenceOfItemSetByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteReferenceOfItemSetByStampSheetRequest{}, err
	}
	return req, nil
}

func NewDeleteReferenceOfItemSetByStampSheetRequestFromDict(data map[string]interface{}) DeleteReferenceOfItemSetByStampSheetRequest {
	return DeleteReferenceOfItemSetByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p DeleteReferenceOfItemSetByStampSheetRequest) Pointer() *DeleteReferenceOfItemSetByStampSheetRequest {
	return &p
}

type VerifyReferenceOfByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyReferenceOfByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyReferenceOfByStampTaskRequest{}
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
		*p = VerifyReferenceOfByStampTaskRequest{}
	} else {
		*p = VerifyReferenceOfByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyReferenceOfByStampTaskRequestFromJson(data string) (VerifyReferenceOfByStampTaskRequest, error) {
	req := VerifyReferenceOfByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyReferenceOfByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyReferenceOfByStampTaskRequestFromDict(data map[string]interface{}) VerifyReferenceOfByStampTaskRequest {
	return VerifyReferenceOfByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyReferenceOfByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyReferenceOfByStampTaskRequest) Pointer() *VerifyReferenceOfByStampTaskRequest {
	return &p
}

type DescribeSimpleItemsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSimpleItemsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleItemsRequest{}
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
		*p = DescribeSimpleItemsRequest{}
	} else {
		*p = DescribeSimpleItemsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeSimpleItemsRequestFromJson(data string) (DescribeSimpleItemsRequest, error) {
	req := DescribeSimpleItemsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleItemsRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleItemsRequestFromDict(data map[string]interface{}) DescribeSimpleItemsRequest {
	return DescribeSimpleItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSimpleItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSimpleItemsRequest) Pointer() *DescribeSimpleItemsRequest {
	return &p
}

type DescribeSimpleItemsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeSimpleItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSimpleItemsByUserIdRequest{}
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
		*p = DescribeSimpleItemsByUserIdRequest{}
	} else {
		*p = DescribeSimpleItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeSimpleItemsByUserIdRequestFromJson(data string) (DescribeSimpleItemsByUserIdRequest, error) {
	req := DescribeSimpleItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSimpleItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) DescribeSimpleItemsByUserIdRequest {
	return DescribeSimpleItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSimpleItemsByUserIdRequest) Pointer() *DescribeSimpleItemsByUserIdRequest {
	return &p
}

type GetSimpleItemRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
}

func (p *GetSimpleItemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemRequest{}
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
		*p = GetSimpleItemRequest{}
	} else {
		*p = GetSimpleItemRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleItemRequestFromJson(data string) (GetSimpleItemRequest, error) {
	req := GetSimpleItemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemRequestFromDict(data map[string]interface{}) GetSimpleItemRequest {
	return GetSimpleItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetSimpleItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
	}
}

func (p GetSimpleItemRequest) Pointer() *GetSimpleItemRequest {
	return &p
}

type GetSimpleItemByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetSimpleItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemByUserIdRequest{}
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
		*p = GetSimpleItemByUserIdRequest{}
	} else {
		*p = GetSimpleItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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

func NewGetSimpleItemByUserIdRequestFromJson(data string) (GetSimpleItemByUserIdRequest, error) {
	req := GetSimpleItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemByUserIdRequestFromDict(data map[string]interface{}) GetSimpleItemByUserIdRequest {
	return GetSimpleItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetSimpleItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSimpleItemByUserIdRequest) Pointer() *GetSimpleItemByUserIdRequest {
	return &p
}

type GetSimpleItemWithSignatureRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
	KeyId           *string `json:"keyId"`
}

func (p *GetSimpleItemWithSignatureRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemWithSignatureRequest{}
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
		*p = GetSimpleItemWithSignatureRequest{}
	} else {
		*p = GetSimpleItemWithSignatureRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewGetSimpleItemWithSignatureRequestFromJson(data string) (GetSimpleItemWithSignatureRequest, error) {
	req := GetSimpleItemWithSignatureRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemWithSignatureRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemWithSignatureRequestFromDict(data map[string]interface{}) GetSimpleItemWithSignatureRequest {
	return GetSimpleItemWithSignatureRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p GetSimpleItemWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"keyId":         p.KeyId,
	}
}

func (p GetSimpleItemWithSignatureRequest) Pointer() *GetSimpleItemWithSignatureRequest {
	return &p
}

type GetSimpleItemWithSignatureByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetSimpleItemWithSignatureByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSimpleItemWithSignatureByUserIdRequest{}
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
		*p = GetSimpleItemWithSignatureByUserIdRequest{}
	} else {
		*p = GetSimpleItemWithSignatureByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
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

func NewGetSimpleItemWithSignatureByUserIdRequestFromJson(data string) (GetSimpleItemWithSignatureByUserIdRequest, error) {
	req := GetSimpleItemWithSignatureByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSimpleItemWithSignatureByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSimpleItemWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetSimpleItemWithSignatureByUserIdRequest {
	return GetSimpleItemWithSignatureByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetSimpleItemWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSimpleItemWithSignatureByUserIdRequest) Pointer() *GetSimpleItemWithSignatureByUserIdRequest {
	return &p
}

type AcquireSimpleItemsByUserIdRequest struct {
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	UserId             *string        `json:"userId"`
	AcquireCounts      []AcquireCount `json:"acquireCounts"`
	TimeOffsetToken    *string        `json:"timeOffsetToken"`
}

func (p *AcquireSimpleItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireSimpleItemsByUserIdRequest{}
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
		*p = AcquireSimpleItemsByUserIdRequest{}
	} else {
		*p = AcquireSimpleItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["acquireCounts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireCounts)
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

func NewAcquireSimpleItemsByUserIdRequestFromJson(data string) (AcquireSimpleItemsByUserIdRequest, error) {
	req := AcquireSimpleItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireSimpleItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewAcquireSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) AcquireSimpleItemsByUserIdRequest {
	return AcquireSimpleItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		AcquireCounts:   CastAcquireCounts(core.CastArray(data["acquireCounts"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"acquireCounts": CastAcquireCountsFromDict(
			p.AcquireCounts,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcquireSimpleItemsByUserIdRequest) Pointer() *AcquireSimpleItemsByUserIdRequest {
	return &p
}

type ConsumeSimpleItemsRequest struct {
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	AccessToken        *string        `json:"accessToken"`
	ConsumeCounts      []ConsumeCount `json:"consumeCounts"`
}

func (p *ConsumeSimpleItemsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeSimpleItemsRequest{}
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
		*p = ConsumeSimpleItemsRequest{}
	} else {
		*p = ConsumeSimpleItemsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["consumeCounts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeCounts)
		}
	}
	return nil
}

func NewConsumeSimpleItemsRequestFromJson(data string) (ConsumeSimpleItemsRequest, error) {
	req := ConsumeSimpleItemsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeSimpleItemsRequest{}, err
	}
	return req, nil
}

func NewConsumeSimpleItemsRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsRequest {
	return ConsumeSimpleItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ConsumeCounts: CastConsumeCounts(core.CastArray(data["consumeCounts"])),
	}
}

func (p ConsumeSimpleItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"consumeCounts": CastConsumeCountsFromDict(
			p.ConsumeCounts,
		),
	}
}

func (p ConsumeSimpleItemsRequest) Pointer() *ConsumeSimpleItemsRequest {
	return &p
}

type ConsumeSimpleItemsByUserIdRequest struct {
	SourceRequestId    *string        `json:"sourceRequestId"`
	RequestId          *string        `json:"requestId"`
	ContextStack       *string        `json:"contextStack"`
	DuplicationAvoider *string        `json:"duplicationAvoider"`
	NamespaceName      *string        `json:"namespaceName"`
	InventoryName      *string        `json:"inventoryName"`
	UserId             *string        `json:"userId"`
	ConsumeCounts      []ConsumeCount `json:"consumeCounts"`
	TimeOffsetToken    *string        `json:"timeOffsetToken"`
}

func (p *ConsumeSimpleItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeSimpleItemsByUserIdRequest{}
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
		*p = ConsumeSimpleItemsByUserIdRequest{}
	} else {
		*p = ConsumeSimpleItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["consumeCounts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeCounts)
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

func NewConsumeSimpleItemsByUserIdRequestFromJson(data string) (ConsumeSimpleItemsByUserIdRequest, error) {
	req := ConsumeSimpleItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeSimpleItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewConsumeSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsByUserIdRequest {
	return ConsumeSimpleItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ConsumeCounts:   CastConsumeCounts(core.CastArray(data["consumeCounts"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ConsumeSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"consumeCounts": CastConsumeCountsFromDict(
			p.ConsumeCounts,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ConsumeSimpleItemsByUserIdRequest) Pointer() *ConsumeSimpleItemsByUserIdRequest {
	return &p
}

type SetSimpleItemsByUserIdRequest struct {
	SourceRequestId    *string     `json:"sourceRequestId"`
	RequestId          *string     `json:"requestId"`
	ContextStack       *string     `json:"contextStack"`
	DuplicationAvoider *string     `json:"duplicationAvoider"`
	NamespaceName      *string     `json:"namespaceName"`
	InventoryName      *string     `json:"inventoryName"`
	UserId             *string     `json:"userId"`
	Counts             []HeldCount `json:"counts"`
	TimeOffsetToken    *string     `json:"timeOffsetToken"`
}

func (p *SetSimpleItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetSimpleItemsByUserIdRequest{}
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
		*p = SetSimpleItemsByUserIdRequest{}
	} else {
		*p = SetSimpleItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["counts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Counts)
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

func NewSetSimpleItemsByUserIdRequestFromJson(data string) (SetSimpleItemsByUserIdRequest, error) {
	req := SetSimpleItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetSimpleItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) SetSimpleItemsByUserIdRequest {
	return SetSimpleItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		Counts:          CastHeldCounts(core.CastArray(data["counts"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"counts": CastHeldCountsFromDict(
			p.Counts,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetSimpleItemsByUserIdRequest) Pointer() *SetSimpleItemsByUserIdRequest {
	return &p
}

type DeleteSimpleItemsByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteSimpleItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSimpleItemsByUserIdRequest{}
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
		*p = DeleteSimpleItemsByUserIdRequest{}
	} else {
		*p = DeleteSimpleItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDeleteSimpleItemsByUserIdRequestFromJson(data string) (DeleteSimpleItemsByUserIdRequest, error) {
	req := DeleteSimpleItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSimpleItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteSimpleItemsByUserIdRequestFromDict(data map[string]interface{}) DeleteSimpleItemsByUserIdRequest {
	return DeleteSimpleItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteSimpleItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteSimpleItemsByUserIdRequest) Pointer() *DeleteSimpleItemsByUserIdRequest {
	return &p
}

type VerifySimpleItemRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	Count                           *int64  `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifySimpleItemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySimpleItemRequest{}
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
		*p = VerifySimpleItemRequest{}
	} else {
		*p = VerifySimpleItemRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifySimpleItemRequestFromJson(data string) (VerifySimpleItemRequest, error) {
	req := VerifySimpleItemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySimpleItemRequest{}, err
	}
	return req, nil
}

func NewVerifySimpleItemRequestFromDict(data map[string]interface{}) VerifySimpleItemRequest {
	return VerifySimpleItemRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		AccessToken:                     core.CastString(data["accessToken"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		Count:                           core.CastInt64(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
	}
}

func (p VerifySimpleItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifySimpleItemRequest) Pointer() *VerifySimpleItemRequest {
	return &p
}

type VerifySimpleItemByUserIdRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	Count                           *int64  `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
}

func (p *VerifySimpleItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySimpleItemByUserIdRequest{}
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
		*p = VerifySimpleItemByUserIdRequest{}
	} else {
		*p = VerifySimpleItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
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

func NewVerifySimpleItemByUserIdRequestFromJson(data string) (VerifySimpleItemByUserIdRequest, error) {
	req := VerifySimpleItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySimpleItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifySimpleItemByUserIdRequestFromDict(data map[string]interface{}) VerifySimpleItemByUserIdRequest {
	return VerifySimpleItemByUserIdRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		UserId:                          core.CastString(data["userId"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		Count:                           core.CastInt64(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
		TimeOffsetToken:                 core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifySimpleItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifySimpleItemByUserIdRequest) Pointer() *VerifySimpleItemByUserIdRequest {
	return &p
}

type AcquireSimpleItemsByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AcquireSimpleItemsByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireSimpleItemsByStampSheetRequest{}
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
		*p = AcquireSimpleItemsByStampSheetRequest{}
	} else {
		*p = AcquireSimpleItemsByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAcquireSimpleItemsByStampSheetRequestFromJson(data string) (AcquireSimpleItemsByStampSheetRequest, error) {
	req := AcquireSimpleItemsByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireSimpleItemsByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAcquireSimpleItemsByStampSheetRequestFromDict(data map[string]interface{}) AcquireSimpleItemsByStampSheetRequest {
	return AcquireSimpleItemsByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireSimpleItemsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireSimpleItemsByStampSheetRequest) Pointer() *AcquireSimpleItemsByStampSheetRequest {
	return &p
}

type ConsumeSimpleItemsByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *ConsumeSimpleItemsByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeSimpleItemsByStampTaskRequest{}
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
		*p = ConsumeSimpleItemsByStampTaskRequest{}
	} else {
		*p = ConsumeSimpleItemsByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewConsumeSimpleItemsByStampTaskRequestFromJson(data string) (ConsumeSimpleItemsByStampTaskRequest, error) {
	req := ConsumeSimpleItemsByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeSimpleItemsByStampTaskRequest{}, err
	}
	return req, nil
}

func NewConsumeSimpleItemsByStampTaskRequestFromDict(data map[string]interface{}) ConsumeSimpleItemsByStampTaskRequest {
	return ConsumeSimpleItemsByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeSimpleItemsByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeSimpleItemsByStampTaskRequest) Pointer() *ConsumeSimpleItemsByStampTaskRequest {
	return &p
}

type SetSimpleItemsByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetSimpleItemsByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetSimpleItemsByStampSheetRequest{}
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
		*p = SetSimpleItemsByStampSheetRequest{}
	} else {
		*p = SetSimpleItemsByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewSetSimpleItemsByStampSheetRequestFromJson(data string) (SetSimpleItemsByStampSheetRequest, error) {
	req := SetSimpleItemsByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetSimpleItemsByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetSimpleItemsByStampSheetRequestFromDict(data map[string]interface{}) SetSimpleItemsByStampSheetRequest {
	return SetSimpleItemsByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetSimpleItemsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetSimpleItemsByStampSheetRequest) Pointer() *SetSimpleItemsByStampSheetRequest {
	return &p
}

type VerifySimpleItemByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifySimpleItemByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySimpleItemByStampTaskRequest{}
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
		*p = VerifySimpleItemByStampTaskRequest{}
	} else {
		*p = VerifySimpleItemByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifySimpleItemByStampTaskRequestFromJson(data string) (VerifySimpleItemByStampTaskRequest, error) {
	req := VerifySimpleItemByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySimpleItemByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifySimpleItemByStampTaskRequestFromDict(data map[string]interface{}) VerifySimpleItemByStampTaskRequest {
	return VerifySimpleItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifySimpleItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifySimpleItemByStampTaskRequest) Pointer() *VerifySimpleItemByStampTaskRequest {
	return &p
}

type DescribeBigItemsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeBigItemsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigItemsRequest{}
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
		*p = DescribeBigItemsRequest{}
	} else {
		*p = DescribeBigItemsRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeBigItemsRequestFromJson(data string) (DescribeBigItemsRequest, error) {
	req := DescribeBigItemsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigItemsRequest{}, err
	}
	return req, nil
}

func NewDescribeBigItemsRequestFromDict(data map[string]interface{}) DescribeBigItemsRequest {
	return DescribeBigItemsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeBigItemsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeBigItemsRequest) Pointer() *DescribeBigItemsRequest {
	return &p
}

type DescribeBigItemsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeBigItemsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeBigItemsByUserIdRequest{}
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
		*p = DescribeBigItemsByUserIdRequest{}
	} else {
		*p = DescribeBigItemsByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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

func NewDescribeBigItemsByUserIdRequestFromJson(data string) (DescribeBigItemsByUserIdRequest, error) {
	req := DescribeBigItemsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeBigItemsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeBigItemsByUserIdRequestFromDict(data map[string]interface{}) DescribeBigItemsByUserIdRequest {
	return DescribeBigItemsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeBigItemsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeBigItemsByUserIdRequest) Pointer() *DescribeBigItemsByUserIdRequest {
	return &p
}

type GetBigItemRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	AccessToken     *string `json:"accessToken"`
	ItemName        *string `json:"itemName"`
}

func (p *GetBigItemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigItemRequest{}
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
		*p = GetBigItemRequest{}
	} else {
		*p = GetBigItemRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
	}
	return nil
}

func NewGetBigItemRequestFromJson(data string) (GetBigItemRequest, error) {
	req := GetBigItemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigItemRequest{}, err
	}
	return req, nil
}

func NewGetBigItemRequestFromDict(data map[string]interface{}) GetBigItemRequest {
	return GetBigItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
	}
}

func (p GetBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
	}
}

func (p GetBigItemRequest) Pointer() *GetBigItemRequest {
	return &p
}

type GetBigItemByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InventoryName   *string `json:"inventoryName"`
	UserId          *string `json:"userId"`
	ItemName        *string `json:"itemName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetBigItemByUserIdRequest{}
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
		*p = GetBigItemByUserIdRequest{}
	} else {
		*p = GetBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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

func NewGetBigItemByUserIdRequestFromJson(data string) (GetBigItemByUserIdRequest, error) {
	req := GetBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetBigItemByUserIdRequestFromDict(data map[string]interface{}) GetBigItemByUserIdRequest {
	return GetBigItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetBigItemByUserIdRequest) Pointer() *GetBigItemByUserIdRequest {
	return &p
}

type AcquireBigItemByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	AcquireCount       *string `json:"acquireCount"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AcquireBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireBigItemByUserIdRequest{}
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
		*p = AcquireBigItemByUserIdRequest{}
	} else {
		*p = AcquireBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["acquireCount"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireCount = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireCount = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireCount = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireCount = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireCount = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireCount)
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

func NewAcquireBigItemByUserIdRequestFromJson(data string) (AcquireBigItemByUserIdRequest, error) {
	req := AcquireBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewAcquireBigItemByUserIdRequestFromDict(data map[string]interface{}) AcquireBigItemByUserIdRequest {
	return AcquireBigItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		AcquireCount:    core.CastString(data["acquireCount"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AcquireBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"acquireCount":    p.AcquireCount,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcquireBigItemByUserIdRequest) Pointer() *AcquireBigItemByUserIdRequest {
	return &p
}

type ConsumeBigItemRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	AccessToken        *string `json:"accessToken"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *string `json:"consumeCount"`
}

func (p *ConsumeBigItemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeBigItemRequest{}
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
		*p = ConsumeBigItemRequest{}
	} else {
		*p = ConsumeBigItemRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["consumeCount"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeCount = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeCount = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeCount = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeCount = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeCount = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeCount)
				}
			}
		}
	}
	return nil
}

func NewConsumeBigItemRequestFromJson(data string) (ConsumeBigItemRequest, error) {
	req := ConsumeBigItemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeBigItemRequest{}, err
	}
	return req, nil
}

func NewConsumeBigItemRequestFromDict(data map[string]interface{}) ConsumeBigItemRequest {
	return ConsumeBigItemRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InventoryName: core.CastString(data["inventoryName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ItemName:      core.CastString(data["itemName"]),
		ConsumeCount:  core.CastString(data["consumeCount"]),
	}
}

func (p ConsumeBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"inventoryName": p.InventoryName,
		"accessToken":   p.AccessToken,
		"itemName":      p.ItemName,
		"consumeCount":  p.ConsumeCount,
	}
}

func (p ConsumeBigItemRequest) Pointer() *ConsumeBigItemRequest {
	return &p
}

type ConsumeBigItemByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	ConsumeCount       *string `json:"consumeCount"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *ConsumeBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeBigItemByUserIdRequest{}
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
		*p = ConsumeBigItemByUserIdRequest{}
	} else {
		*p = ConsumeBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["consumeCount"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeCount = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeCount = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeCount = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeCount = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeCount = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeCount)
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

func NewConsumeBigItemByUserIdRequestFromJson(data string) (ConsumeBigItemByUserIdRequest, error) {
	req := ConsumeBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewConsumeBigItemByUserIdRequestFromDict(data map[string]interface{}) ConsumeBigItemByUserIdRequest {
	return ConsumeBigItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		ConsumeCount:    core.CastString(data["consumeCount"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ConsumeBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"consumeCount":    p.ConsumeCount,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ConsumeBigItemByUserIdRequest) Pointer() *ConsumeBigItemByUserIdRequest {
	return &p
}

type SetBigItemByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	Count              *string `json:"count"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *SetBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetBigItemByUserIdRequest{}
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
		*p = SetBigItemByUserIdRequest{}
	} else {
		*p = SetBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Count = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Count = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Count = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Count)
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

func NewSetBigItemByUserIdRequestFromJson(data string) (SetBigItemByUserIdRequest, error) {
	req := SetBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetBigItemByUserIdRequestFromDict(data map[string]interface{}) SetBigItemByUserIdRequest {
	return SetBigItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		Count:           core.CastString(data["count"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"count":           p.Count,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetBigItemByUserIdRequest) Pointer() *SetBigItemByUserIdRequest {
	return &p
}

type DeleteBigItemByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	InventoryName      *string `json:"inventoryName"`
	UserId             *string `json:"userId"`
	ItemName           *string `json:"itemName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteBigItemByUserIdRequest{}
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
		*p = DeleteBigItemByUserIdRequest{}
	} else {
		*p = DeleteBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
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

func NewDeleteBigItemByUserIdRequestFromJson(data string) (DeleteBigItemByUserIdRequest, error) {
	req := DeleteBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteBigItemByUserIdRequestFromDict(data map[string]interface{}) DeleteBigItemByUserIdRequest {
	return DeleteBigItemByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		InventoryName:   core.CastString(data["inventoryName"]),
		UserId:          core.CastString(data["userId"]),
		ItemName:        core.CastString(data["itemName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"inventoryName":   p.InventoryName,
		"userId":          p.UserId,
		"itemName":        p.ItemName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteBigItemByUserIdRequest) Pointer() *DeleteBigItemByUserIdRequest {
	return &p
}

type VerifyBigItemRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	Count                           *string `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifyBigItemRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyBigItemRequest{}
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
		*p = VerifyBigItemRequest{}
	} else {
		*p = VerifyBigItemRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Count = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Count = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Count = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Count)
				}
			}
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyBigItemRequestFromJson(data string) (VerifyBigItemRequest, error) {
	req := VerifyBigItemRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyBigItemRequest{}, err
	}
	return req, nil
}

func NewVerifyBigItemRequestFromDict(data map[string]interface{}) VerifyBigItemRequest {
	return VerifyBigItemRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		AccessToken:                     core.CastString(data["accessToken"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		Count:                           core.CastString(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
	}
}

func (p VerifyBigItemRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyBigItemRequest) Pointer() *VerifyBigItemRequest {
	return &p
}

type VerifyBigItemByUserIdRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	InventoryName                   *string `json:"inventoryName"`
	ItemName                        *string `json:"itemName"`
	VerifyType                      *string `json:"verifyType"`
	Count                           *string `json:"count"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
}

func (p *VerifyBigItemByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyBigItemByUserIdRequest{}
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
		*p = VerifyBigItemByUserIdRequest{}
	} else {
		*p = VerifyBigItemByUserIdRequest{}
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
				}
			}
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["verifyType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyType)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Count = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Count = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Count = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Count)
				}
			}
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
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

func NewVerifyBigItemByUserIdRequestFromJson(data string) (VerifyBigItemByUserIdRequest, error) {
	req := VerifyBigItemByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyBigItemByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyBigItemByUserIdRequestFromDict(data map[string]interface{}) VerifyBigItemByUserIdRequest {
	return VerifyBigItemByUserIdRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		UserId:                          core.CastString(data["userId"]),
		InventoryName:                   core.CastString(data["inventoryName"]),
		ItemName:                        core.CastString(data["itemName"]),
		VerifyType:                      core.CastString(data["verifyType"]),
		Count:                           core.CastString(data["count"]),
		MultiplyValueSpecifyingQuantity: core.CastBool(data["multiplyValueSpecifyingQuantity"]),
		TimeOffsetToken:                 core.CastString(data["timeOffsetToken"]),
	}
}

func (p VerifyBigItemByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"inventoryName":                   p.InventoryName,
		"itemName":                        p.ItemName,
		"verifyType":                      p.VerifyType,
		"count":                           p.Count,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyBigItemByUserIdRequest) Pointer() *VerifyBigItemByUserIdRequest {
	return &p
}

type AcquireBigItemByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *AcquireBigItemByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireBigItemByStampSheetRequest{}
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
		*p = AcquireBigItemByStampSheetRequest{}
	} else {
		*p = AcquireBigItemByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewAcquireBigItemByStampSheetRequestFromJson(data string) (AcquireBigItemByStampSheetRequest, error) {
	req := AcquireBigItemByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcquireBigItemByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAcquireBigItemByStampSheetRequestFromDict(data map[string]interface{}) AcquireBigItemByStampSheetRequest {
	return AcquireBigItemByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p AcquireBigItemByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AcquireBigItemByStampSheetRequest) Pointer() *AcquireBigItemByStampSheetRequest {
	return &p
}

type ConsumeBigItemByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *ConsumeBigItemByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeBigItemByStampTaskRequest{}
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
		*p = ConsumeBigItemByStampTaskRequest{}
	} else {
		*p = ConsumeBigItemByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewConsumeBigItemByStampTaskRequestFromJson(data string) (ConsumeBigItemByStampTaskRequest, error) {
	req := ConsumeBigItemByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeBigItemByStampTaskRequest{}, err
	}
	return req, nil
}

func NewConsumeBigItemByStampTaskRequestFromDict(data map[string]interface{}) ConsumeBigItemByStampTaskRequest {
	return ConsumeBigItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeBigItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeBigItemByStampTaskRequest) Pointer() *ConsumeBigItemByStampTaskRequest {
	return &p
}

type SetBigItemByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetBigItemByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetBigItemByStampSheetRequest{}
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
		*p = SetBigItemByStampSheetRequest{}
	} else {
		*p = SetBigItemByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampSheet = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampSheet = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampSheet = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampSheet = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampSheet)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewSetBigItemByStampSheetRequestFromJson(data string) (SetBigItemByStampSheetRequest, error) {
	req := SetBigItemByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetBigItemByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetBigItemByStampSheetRequestFromDict(data map[string]interface{}) SetBigItemByStampSheetRequest {
	return SetBigItemByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetBigItemByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetBigItemByStampSheetRequest) Pointer() *SetBigItemByStampSheetRequest {
	return &p
}

type VerifyBigItemByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyBigItemByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyBigItemByStampTaskRequest{}
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
		*p = VerifyBigItemByStampTaskRequest{}
	} else {
		*p = VerifyBigItemByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StampTask = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StampTask = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StampTask = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StampTask = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StampTask)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyBigItemByStampTaskRequestFromJson(data string) (VerifyBigItemByStampTaskRequest, error) {
	req := VerifyBigItemByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyBigItemByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyBigItemByStampTaskRequestFromDict(data map[string]interface{}) VerifyBigItemByStampTaskRequest {
	return VerifyBigItemByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p VerifyBigItemByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyBigItemByStampTaskRequest) Pointer() *VerifyBigItemByStampTaskRequest {
	return &p
}
