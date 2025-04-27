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

package version

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
	DryRun       *bool   `json:"dryRun"`
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
	ContextStack                *string        `json:"contextStack"`
	Name                        *string        `json:"name"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
	DryRun                      *bool          `json:"dryRun"`
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
		if v, ok := d["assumeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AssumeUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AssumeUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AssumeUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AssumeUserId)
				}
			}
		}
		if v, ok := d["acceptVersionScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcceptVersionScript)
		}
		if v, ok := d["checkVersionTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CheckVersionTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CheckVersionTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CheckVersionTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CheckVersionTriggerScriptId)
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
		AssumeUserId: func() *string {
			v, ok := data["assumeUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["assumeUserId"])
		}(),
		AcceptVersionScript: func() *ScriptSetting {
			v, ok := data["acceptVersionScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer()
		}(),
		CheckVersionTriggerScriptId: func() *string {
			v, ok := data["checkVersionTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["checkVersionTriggerScriptId"])
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":         p.Name,
		"description":  p.Description,
		"assumeUserId": p.AssumeUserId,
		"acceptVersionScript": func() map[string]interface{} {
			if p.AcceptVersionScript == nil {
				return nil
			}
			return p.AcceptVersionScript.ToDict()
		}(),
		"checkVersionTriggerScriptId": p.CheckVersionTriggerScriptId,
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
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
	ContextStack                *string        `json:"contextStack"`
	NamespaceName               *string        `json:"namespaceName"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
	DryRun                      *bool          `json:"dryRun"`
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
		if v, ok := d["assumeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AssumeUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AssumeUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AssumeUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AssumeUserId)
				}
			}
		}
		if v, ok := d["acceptVersionScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcceptVersionScript)
		}
		if v, ok := d["checkVersionTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CheckVersionTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CheckVersionTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CheckVersionTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CheckVersionTriggerScriptId)
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
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		AssumeUserId: func() *string {
			v, ok := data["assumeUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["assumeUserId"])
		}(),
		AcceptVersionScript: func() *ScriptSetting {
			v, ok := data["acceptVersionScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer()
		}(),
		CheckVersionTriggerScriptId: func() *string {
			v, ok := data["checkVersionTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["checkVersionTriggerScriptId"])
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"description":   p.Description,
		"assumeUserId":  p.AssumeUserId,
		"acceptVersionScript": func() map[string]interface{} {
			if p.AcceptVersionScript == nil {
				return nil
			}
			return p.AcceptVersionScript.ToDict()
		}(),
		"checkVersionTriggerScriptId": p.CheckVersionTriggerScriptId,
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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
	ContextStack    *string `json:"contextStack"`
	UserId          *string `json:"userId"`
	UploadToken     *string `json:"uploadToken"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
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

type DescribeVersionModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeVersionModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeVersionModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeVersionModelMastersRequest{}
	} else {
		*p = DescribeVersionModelMastersRequest{}
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

func NewDescribeVersionModelMastersRequestFromJson(data string) (DescribeVersionModelMastersRequest, error) {
	req := DescribeVersionModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeVersionModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeVersionModelMastersRequestFromDict(data map[string]interface{}) DescribeVersionModelMastersRequest {
	return DescribeVersionModelMastersRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
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

func (p DescribeVersionModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeVersionModelMastersRequest) Pointer() *DescribeVersionModelMastersRequest {
	return &p
}

type CreateVersionModelMasterRequest struct {
	ContextStack       *string           `json:"contextStack"`
	NamespaceName      *string           `json:"namespaceName"`
	Name               *string           `json:"name"`
	Description        *string           `json:"description"`
	Metadata           *string           `json:"metadata"`
	Scope              *string           `json:"scope"`
	Type               *string           `json:"type"`
	CurrentVersion     *Version          `json:"currentVersion"`
	WarningVersion     *Version          `json:"warningVersion"`
	ErrorVersion       *Version          `json:"errorVersion"`
	ScheduleVersions   []ScheduleVersion `json:"scheduleVersions"`
	NeedSignature      *bool             `json:"needSignature"`
	SignatureKeyId     *string           `json:"signatureKeyId"`
	ApproveRequirement *string           `json:"approveRequirement"`
	DryRun             *bool             `json:"dryRun"`
}

func (p *CreateVersionModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateVersionModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateVersionModelMasterRequest{}
	} else {
		*p = CreateVersionModelMasterRequest{}
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
		}
		if v, ok := d["warningVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WarningVersion)
		}
		if v, ok := d["errorVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorVersion)
		}
		if v, ok := d["scheduleVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScheduleVersions)
		}
		if v, ok := d["needSignature"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NeedSignature)
		}
		if v, ok := d["signatureKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SignatureKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SignatureKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SignatureKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SignatureKeyId)
				}
			}
		}
		if v, ok := d["approveRequirement"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApproveRequirement = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApproveRequirement = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApproveRequirement = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApproveRequirement = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApproveRequirement = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApproveRequirement)
				}
			}
		}
	}
	return nil
}

func NewCreateVersionModelMasterRequestFromJson(data string) (CreateVersionModelMasterRequest, error) {
	req := CreateVersionModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateVersionModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateVersionModelMasterRequestFromDict(data map[string]interface{}) CreateVersionModelMasterRequest {
	return CreateVersionModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
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
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		Scope: func() *string {
			v, ok := data["scope"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scope"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
		WarningVersion: func() *Version {
			v, ok := data["warningVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer()
		}(),
		ErrorVersion: func() *Version {
			v, ok := data["errorVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer()
		}(),
		ScheduleVersions: func() []ScheduleVersion {
			if data["scheduleVersions"] == nil {
				return nil
			}
			return CastScheduleVersions(core.CastArray(data["scheduleVersions"]))
		}(),
		NeedSignature: func() *bool {
			v, ok := data["needSignature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["needSignature"])
		}(),
		SignatureKeyId: func() *string {
			v, ok := data["signatureKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signatureKeyId"])
		}(),
		ApproveRequirement: func() *string {
			v, ok := data["approveRequirement"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["approveRequirement"])
		}(),
	}
}

func (p CreateVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"scope":         p.Scope,
		"type":          p.Type,
		"currentVersion": func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}(),
		"warningVersion": func() map[string]interface{} {
			if p.WarningVersion == nil {
				return nil
			}
			return p.WarningVersion.ToDict()
		}(),
		"errorVersion": func() map[string]interface{} {
			if p.ErrorVersion == nil {
				return nil
			}
			return p.ErrorVersion.ToDict()
		}(),
		"scheduleVersions": CastScheduleVersionsFromDict(
			p.ScheduleVersions,
		),
		"needSignature":      p.NeedSignature,
		"signatureKeyId":     p.SignatureKeyId,
		"approveRequirement": p.ApproveRequirement,
	}
}

func (p CreateVersionModelMasterRequest) Pointer() *CreateVersionModelMasterRequest {
	return &p
}

type GetVersionModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	VersionName   *string `json:"versionName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetVersionModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetVersionModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetVersionModelMasterRequest{}
	} else {
		*p = GetVersionModelMasterRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
	}
	return nil
}

func NewGetVersionModelMasterRequestFromJson(data string) (GetVersionModelMasterRequest, error) {
	req := GetVersionModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetVersionModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetVersionModelMasterRequestFromDict(data map[string]interface{}) GetVersionModelMasterRequest {
	return GetVersionModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
	}
}

func (p GetVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p GetVersionModelMasterRequest) Pointer() *GetVersionModelMasterRequest {
	return &p
}

type UpdateVersionModelMasterRequest struct {
	ContextStack       *string           `json:"contextStack"`
	NamespaceName      *string           `json:"namespaceName"`
	VersionName        *string           `json:"versionName"`
	Description        *string           `json:"description"`
	Metadata           *string           `json:"metadata"`
	Scope              *string           `json:"scope"`
	Type               *string           `json:"type"`
	CurrentVersion     *Version          `json:"currentVersion"`
	WarningVersion     *Version          `json:"warningVersion"`
	ErrorVersion       *Version          `json:"errorVersion"`
	ScheduleVersions   []ScheduleVersion `json:"scheduleVersions"`
	NeedSignature      *bool             `json:"needSignature"`
	SignatureKeyId     *string           `json:"signatureKeyId"`
	ApproveRequirement *string           `json:"approveRequirement"`
	DryRun             *bool             `json:"dryRun"`
}

func (p *UpdateVersionModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateVersionModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateVersionModelMasterRequest{}
	} else {
		*p = UpdateVersionModelMasterRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
		}
		if v, ok := d["warningVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WarningVersion)
		}
		if v, ok := d["errorVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorVersion)
		}
		if v, ok := d["scheduleVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScheduleVersions)
		}
		if v, ok := d["needSignature"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NeedSignature)
		}
		if v, ok := d["signatureKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SignatureKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SignatureKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SignatureKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SignatureKeyId)
				}
			}
		}
		if v, ok := d["approveRequirement"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApproveRequirement = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApproveRequirement = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApproveRequirement = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApproveRequirement = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApproveRequirement = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApproveRequirement)
				}
			}
		}
	}
	return nil
}

func NewUpdateVersionModelMasterRequestFromJson(data string) (UpdateVersionModelMasterRequest, error) {
	req := UpdateVersionModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateVersionModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateVersionModelMasterRequestFromDict(data map[string]interface{}) UpdateVersionModelMasterRequest {
	return UpdateVersionModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		Scope: func() *string {
			v, ok := data["scope"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scope"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
		WarningVersion: func() *Version {
			v, ok := data["warningVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer()
		}(),
		ErrorVersion: func() *Version {
			v, ok := data["errorVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer()
		}(),
		ScheduleVersions: func() []ScheduleVersion {
			if data["scheduleVersions"] == nil {
				return nil
			}
			return CastScheduleVersions(core.CastArray(data["scheduleVersions"]))
		}(),
		NeedSignature: func() *bool {
			v, ok := data["needSignature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["needSignature"])
		}(),
		SignatureKeyId: func() *string {
			v, ok := data["signatureKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signatureKeyId"])
		}(),
		ApproveRequirement: func() *string {
			v, ok := data["approveRequirement"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["approveRequirement"])
		}(),
	}
}

func (p UpdateVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"scope":         p.Scope,
		"type":          p.Type,
		"currentVersion": func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}(),
		"warningVersion": func() map[string]interface{} {
			if p.WarningVersion == nil {
				return nil
			}
			return p.WarningVersion.ToDict()
		}(),
		"errorVersion": func() map[string]interface{} {
			if p.ErrorVersion == nil {
				return nil
			}
			return p.ErrorVersion.ToDict()
		}(),
		"scheduleVersions": CastScheduleVersionsFromDict(
			p.ScheduleVersions,
		),
		"needSignature":      p.NeedSignature,
		"signatureKeyId":     p.SignatureKeyId,
		"approveRequirement": p.ApproveRequirement,
	}
}

func (p UpdateVersionModelMasterRequest) Pointer() *UpdateVersionModelMasterRequest {
	return &p
}

type DeleteVersionModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	VersionName   *string `json:"versionName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteVersionModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteVersionModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteVersionModelMasterRequest{}
	} else {
		*p = DeleteVersionModelMasterRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
	}
	return nil
}

func NewDeleteVersionModelMasterRequestFromJson(data string) (DeleteVersionModelMasterRequest, error) {
	req := DeleteVersionModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteVersionModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteVersionModelMasterRequestFromDict(data map[string]interface{}) DeleteVersionModelMasterRequest {
	return DeleteVersionModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
	}
}

func (p DeleteVersionModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p DeleteVersionModelMasterRequest) Pointer() *DeleteVersionModelMasterRequest {
	return &p
}

type DescribeVersionModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeVersionModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeVersionModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeVersionModelsRequest{}
	} else {
		*p = DescribeVersionModelsRequest{}
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

func NewDescribeVersionModelsRequestFromJson(data string) (DescribeVersionModelsRequest, error) {
	req := DescribeVersionModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeVersionModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeVersionModelsRequestFromDict(data map[string]interface{}) DescribeVersionModelsRequest {
	return DescribeVersionModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeVersionModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeVersionModelsRequest) Pointer() *DescribeVersionModelsRequest {
	return &p
}

type GetVersionModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	VersionName   *string `json:"versionName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetVersionModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetVersionModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetVersionModelRequest{}
	} else {
		*p = GetVersionModelRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
	}
	return nil
}

func NewGetVersionModelRequestFromJson(data string) (GetVersionModelRequest, error) {
	req := GetVersionModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetVersionModelRequest{}, err
	}
	return req, nil
}

func NewGetVersionModelRequestFromDict(data map[string]interface{}) GetVersionModelRequest {
	return GetVersionModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
	}
}

func (p GetVersionModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
	}
}

func (p GetVersionModelRequest) Pointer() *GetVersionModelRequest {
	return &p
}

type DescribeAcceptVersionsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeAcceptVersionsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAcceptVersionsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeAcceptVersionsRequest{}
	} else {
		*p = DescribeAcceptVersionsRequest{}
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

func NewDescribeAcceptVersionsRequestFromJson(data string) (DescribeAcceptVersionsRequest, error) {
	req := DescribeAcceptVersionsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAcceptVersionsRequest{}, err
	}
	return req, nil
}

func NewDescribeAcceptVersionsRequestFromDict(data map[string]interface{}) DescribeAcceptVersionsRequest {
	return DescribeAcceptVersionsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
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

func (p DescribeAcceptVersionsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeAcceptVersionsRequest) Pointer() *DescribeAcceptVersionsRequest {
	return &p
}

type DescribeAcceptVersionsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeAcceptVersionsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeAcceptVersionsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeAcceptVersionsByUserIdRequest{}
	} else {
		*p = DescribeAcceptVersionsByUserIdRequest{}
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

func NewDescribeAcceptVersionsByUserIdRequestFromJson(data string) (DescribeAcceptVersionsByUserIdRequest, error) {
	req := DescribeAcceptVersionsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeAcceptVersionsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeAcceptVersionsByUserIdRequestFromDict(data map[string]interface{}) DescribeAcceptVersionsByUserIdRequest {
	return DescribeAcceptVersionsByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
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
		TimeOffsetToken: func() *string {
			v, ok := data["timeOffsetToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["timeOffsetToken"])
		}(),
	}
}

func (p DescribeAcceptVersionsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeAcceptVersionsByUserIdRequest) Pointer() *DescribeAcceptVersionsByUserIdRequest {
	return &p
}

type AcceptRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	AccessToken        *string  `json:"accessToken"`
	Version            *Version `json:"version"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *AcceptRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcceptRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcceptRequest{}
	} else {
		*p = AcceptRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
		}
	}
	return nil
}

func NewAcceptRequestFromJson(data string) (AcceptRequest, error) {
	req := AcceptRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcceptRequest{}, err
	}
	return req, nil
}

func NewAcceptRequestFromDict(data map[string]interface{}) AcceptRequest {
	return AcceptRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
		}(),
	}
}

func (p AcceptRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"accessToken":   p.AccessToken,
		"version": func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}(),
	}
}

func (p AcceptRequest) Pointer() *AcceptRequest {
	return &p
}

type AcceptByUserIdRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	UserId             *string  `json:"userId"`
	Version            *Version `json:"version"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *AcceptByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcceptByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcceptByUserIdRequest{}
	} else {
		*p = AcceptByUserIdRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
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

func NewAcceptByUserIdRequestFromJson(data string) (AcceptByUserIdRequest, error) {
	req := AcceptByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcceptByUserIdRequest{}, err
	}
	return req, nil
}

func NewAcceptByUserIdRequestFromDict(data map[string]interface{}) AcceptByUserIdRequest {
	return AcceptByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
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

func (p AcceptByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"userId":        p.UserId,
		"version": func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AcceptByUserIdRequest) Pointer() *AcceptByUserIdRequest {
	return &p
}

type RejectRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	AccessToken        *string  `json:"accessToken"`
	Version            *Version `json:"version"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *RejectRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RejectRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RejectRequest{}
	} else {
		*p = RejectRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
		}
	}
	return nil
}

func NewRejectRequestFromJson(data string) (RejectRequest, error) {
	req := RejectRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RejectRequest{}, err
	}
	return req, nil
}

func NewRejectRequestFromDict(data map[string]interface{}) RejectRequest {
	return RejectRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
		}(),
	}
}

func (p RejectRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"accessToken":   p.AccessToken,
		"version": func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}(),
	}
}

func (p RejectRequest) Pointer() *RejectRequest {
	return &p
}

type RejectByUserIdRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	VersionName        *string  `json:"versionName"`
	UserId             *string  `json:"userId"`
	Version            *Version `json:"version"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *RejectByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RejectByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RejectByUserIdRequest{}
	} else {
		*p = RejectByUserIdRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
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

func NewRejectByUserIdRequestFromJson(data string) (RejectByUserIdRequest, error) {
	req := RejectByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RejectByUserIdRequest{}, err
	}
	return req, nil
}

func NewRejectByUserIdRequestFromDict(data map[string]interface{}) RejectByUserIdRequest {
	return RejectByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
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

func (p RejectByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"userId":        p.UserId,
		"version": func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}(),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RejectByUserIdRequest) Pointer() *RejectByUserIdRequest {
	return &p
}

type GetAcceptVersionRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	VersionName   *string `json:"versionName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetAcceptVersionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetAcceptVersionRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetAcceptVersionRequest{}
	} else {
		*p = GetAcceptVersionRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
	}
	return nil
}

func NewGetAcceptVersionRequestFromJson(data string) (GetAcceptVersionRequest, error) {
	req := GetAcceptVersionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetAcceptVersionRequest{}, err
	}
	return req, nil
}

func NewGetAcceptVersionRequestFromDict(data map[string]interface{}) GetAcceptVersionRequest {
	return GetAcceptVersionRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
	}
}

func (p GetAcceptVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"versionName":   p.VersionName,
	}
}

func (p GetAcceptVersionRequest) Pointer() *GetAcceptVersionRequest {
	return &p
}

type GetAcceptVersionByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	VersionName     *string `json:"versionName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetAcceptVersionByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetAcceptVersionByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetAcceptVersionByUserIdRequest{}
	} else {
		*p = GetAcceptVersionByUserIdRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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

func NewGetAcceptVersionByUserIdRequestFromJson(data string) (GetAcceptVersionByUserIdRequest, error) {
	req := GetAcceptVersionByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetAcceptVersionByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetAcceptVersionByUserIdRequestFromDict(data map[string]interface{}) GetAcceptVersionByUserIdRequest {
	return GetAcceptVersionByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
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

func (p GetAcceptVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"versionName":     p.VersionName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetAcceptVersionByUserIdRequest) Pointer() *GetAcceptVersionByUserIdRequest {
	return &p
}

type DeleteAcceptVersionRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	VersionName        *string `json:"versionName"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteAcceptVersionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteAcceptVersionRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteAcceptVersionRequest{}
	} else {
		*p = DeleteAcceptVersionRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
	}
	return nil
}

func NewDeleteAcceptVersionRequestFromJson(data string) (DeleteAcceptVersionRequest, error) {
	req := DeleteAcceptVersionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteAcceptVersionRequest{}, err
	}
	return req, nil
}

func NewDeleteAcceptVersionRequestFromDict(data map[string]interface{}) DeleteAcceptVersionRequest {
	return DeleteAcceptVersionRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
	}
}

func (p DeleteAcceptVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"versionName":   p.VersionName,
	}
}

func (p DeleteAcceptVersionRequest) Pointer() *DeleteAcceptVersionRequest {
	return &p
}

type DeleteAcceptVersionByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	VersionName        *string `json:"versionName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteAcceptVersionByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteAcceptVersionByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteAcceptVersionByUserIdRequest{}
	} else {
		*p = DeleteAcceptVersionByUserIdRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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

func NewDeleteAcceptVersionByUserIdRequestFromJson(data string) (DeleteAcceptVersionByUserIdRequest, error) {
	req := DeleteAcceptVersionByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteAcceptVersionByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteAcceptVersionByUserIdRequestFromDict(data map[string]interface{}) DeleteAcceptVersionByUserIdRequest {
	return DeleteAcceptVersionByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
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

func (p DeleteAcceptVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"versionName":     p.VersionName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteAcceptVersionByUserIdRequest) Pointer() *DeleteAcceptVersionByUserIdRequest {
	return &p
}

type CheckVersionRequest struct {
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	AccessToken        *string         `json:"accessToken"`
	TargetVersions     []TargetVersion `json:"targetVersions"`
	DryRun             *bool           `json:"dryRun"`
}

func (p *CheckVersionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckVersionRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CheckVersionRequest{}
	} else {
		*p = CheckVersionRequest{}
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
		if v, ok := d["targetVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetVersions)
		}
	}
	return nil
}

func NewCheckVersionRequestFromJson(data string) (CheckVersionRequest, error) {
	req := CheckVersionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckVersionRequest{}, err
	}
	return req, nil
}

func NewCheckVersionRequestFromDict(data map[string]interface{}) CheckVersionRequest {
	return CheckVersionRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		TargetVersions: func() []TargetVersion {
			if data["targetVersions"] == nil {
				return nil
			}
			return CastTargetVersions(core.CastArray(data["targetVersions"]))
		}(),
	}
}

func (p CheckVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"targetVersions": CastTargetVersionsFromDict(
			p.TargetVersions,
		),
	}
}

func (p CheckVersionRequest) Pointer() *CheckVersionRequest {
	return &p
}

type CheckVersionByUserIdRequest struct {
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	TargetVersions     []TargetVersion `json:"targetVersions"`
	TimeOffsetToken    *string         `json:"timeOffsetToken"`
	DryRun             *bool           `json:"dryRun"`
}

func (p *CheckVersionByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CheckVersionByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CheckVersionByUserIdRequest{}
	} else {
		*p = CheckVersionByUserIdRequest{}
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
		if v, ok := d["targetVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetVersions)
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

func NewCheckVersionByUserIdRequestFromJson(data string) (CheckVersionByUserIdRequest, error) {
	req := CheckVersionByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CheckVersionByUserIdRequest{}, err
	}
	return req, nil
}

func NewCheckVersionByUserIdRequestFromDict(data map[string]interface{}) CheckVersionByUserIdRequest {
	return CheckVersionByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetVersions: func() []TargetVersion {
			if data["targetVersions"] == nil {
				return nil
			}
			return CastTargetVersions(core.CastArray(data["targetVersions"]))
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

func (p CheckVersionByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"targetVersions": CastTargetVersionsFromDict(
			p.TargetVersions,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckVersionByUserIdRequest) Pointer() *CheckVersionByUserIdRequest {
	return &p
}

type CalculateSignatureRequest struct {
	ContextStack  *string  `json:"contextStack"`
	NamespaceName *string  `json:"namespaceName"`
	VersionName   *string  `json:"versionName"`
	Version       *Version `json:"version"`
	DryRun        *bool    `json:"dryRun"`
}

func (p *CalculateSignatureRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CalculateSignatureRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CalculateSignatureRequest{}
	} else {
		*p = CalculateSignatureRequest{}
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
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
		}
	}
	return nil
}

func NewCalculateSignatureRequestFromJson(data string) (CalculateSignatureRequest, error) {
	req := CalculateSignatureRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CalculateSignatureRequest{}, err
	}
	return req, nil
}

func NewCalculateSignatureRequestFromDict(data map[string]interface{}) CalculateSignatureRequest {
	return CalculateSignatureRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
		}(),
	}
}

func (p CalculateSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"version": func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}(),
	}
}

func (p CalculateSignatureRequest) Pointer() *CalculateSignatureRequest {
	return &p
}

type ExportMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
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
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
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

type GetCurrentVersionMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetCurrentVersionMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentVersionMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentVersionMasterRequest{}
	} else {
		*p = GetCurrentVersionMasterRequest{}
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

func NewGetCurrentVersionMasterRequestFromJson(data string) (GetCurrentVersionMasterRequest, error) {
	req := GetCurrentVersionMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentVersionMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentVersionMasterRequestFromDict(data map[string]interface{}) GetCurrentVersionMasterRequest {
	return GetCurrentVersionMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentVersionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentVersionMasterRequest) Pointer() *GetCurrentVersionMasterRequest {
	return &p
}

type PreUpdateCurrentVersionMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *PreUpdateCurrentVersionMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PreUpdateCurrentVersionMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PreUpdateCurrentVersionMasterRequest{}
	} else {
		*p = PreUpdateCurrentVersionMasterRequest{}
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

func NewPreUpdateCurrentVersionMasterRequestFromJson(data string) (PreUpdateCurrentVersionMasterRequest, error) {
	req := PreUpdateCurrentVersionMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PreUpdateCurrentVersionMasterRequest{}, err
	}
	return req, nil
}

func NewPreUpdateCurrentVersionMasterRequestFromDict(data map[string]interface{}) PreUpdateCurrentVersionMasterRequest {
	return PreUpdateCurrentVersionMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p PreUpdateCurrentVersionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p PreUpdateCurrentVersionMasterRequest) Pointer() *PreUpdateCurrentVersionMasterRequest {
	return &p
}

type UpdateCurrentVersionMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Mode          *string `json:"mode"`
	Settings      *string `json:"settings"`
	UploadToken   *string `json:"uploadToken"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *UpdateCurrentVersionMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentVersionMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentVersionMasterRequest{}
	} else {
		*p = UpdateCurrentVersionMasterRequest{}
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
		if v, ok := d["mode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Mode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Mode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Mode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Mode)
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
	}
	return nil
}

func NewUpdateCurrentVersionMasterRequestFromJson(data string) (UpdateCurrentVersionMasterRequest, error) {
	req := UpdateCurrentVersionMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentVersionMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentVersionMasterRequestFromDict(data map[string]interface{}) UpdateCurrentVersionMasterRequest {
	return UpdateCurrentVersionMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Mode: func() *string {
			v, ok := data["mode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["mode"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
	}
}

func (p UpdateCurrentVersionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"mode":          p.Mode,
		"settings":      p.Settings,
		"uploadToken":   p.UploadToken,
	}
}

func (p UpdateCurrentVersionMasterRequest) Pointer() *UpdateCurrentVersionMasterRequest {
	return &p
}

type UpdateCurrentVersionMasterFromGitHubRequest struct {
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
	DryRun          *bool                  `json:"dryRun"`
}

func (p *UpdateCurrentVersionMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentVersionMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentVersionMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentVersionMasterFromGitHubRequest{}
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

func NewUpdateCurrentVersionMasterFromGitHubRequestFromJson(data string) (UpdateCurrentVersionMasterFromGitHubRequest, error) {
	req := UpdateCurrentVersionMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentVersionMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentVersionMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentVersionMasterFromGitHubRequest {
	return UpdateCurrentVersionMasterFromGitHubRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
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

func (p UpdateCurrentVersionMasterFromGitHubRequest) ToDict() map[string]interface{} {
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

func (p UpdateCurrentVersionMasterFromGitHubRequest) Pointer() *UpdateCurrentVersionMasterFromGitHubRequest {
	return &p
}
