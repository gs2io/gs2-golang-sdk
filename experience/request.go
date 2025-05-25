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

package experience

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
	ContextStack             *string             `json:"contextStack"`
	Name                     *string             `json:"name"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	RankCapScriptId          *string             `json:"rankCapScriptId"`
	ChangeExperienceScript   *ScriptSetting      `json:"changeExperienceScript"`
	ChangeRankScript         *ScriptSetting      `json:"changeRankScript"`
	ChangeRankCapScript      *ScriptSetting      `json:"changeRankCapScript"`
	OverflowExperienceScript *string             `json:"overflowExperienceScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
	DryRun                   *bool               `json:"dryRun"`
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["rankCapScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankCapScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankCapScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankCapScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankCapScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankCapScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankCapScriptId)
				}
			}
		}
		if v, ok := d["changeExperienceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeExperienceScript)
		}
		if v, ok := d["changeRankScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRankScript)
		}
		if v, ok := d["changeRankCapScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRankCapScript)
		}
		if v, ok := d["overflowExperienceScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OverflowExperienceScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OverflowExperienceScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OverflowExperienceScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OverflowExperienceScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OverflowExperienceScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OverflowExperienceScript)
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
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		RankCapScriptId: func() *string {
			v, ok := data["rankCapScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankCapScriptId"])
		}(),
		ChangeExperienceScript: func() *ScriptSetting {
			v, ok := data["changeExperienceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer()
		}(),
		ChangeRankScript: func() *ScriptSetting {
			v, ok := data["changeRankScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer()
		}(),
		ChangeRankCapScript: func() *ScriptSetting {
			v, ok := data["changeRankCapScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer()
		}(),
		OverflowExperienceScript: func() *string {
			v, ok := data["overflowExperienceScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["overflowExperienceScript"])
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
		"name":        p.Name,
		"description": p.Description,
		"transactionSetting": func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}(),
		"rankCapScriptId": p.RankCapScriptId,
		"changeExperienceScript": func() map[string]interface{} {
			if p.ChangeExperienceScript == nil {
				return nil
			}
			return p.ChangeExperienceScript.ToDict()
		}(),
		"changeRankScript": func() map[string]interface{} {
			if p.ChangeRankScript == nil {
				return nil
			}
			return p.ChangeRankScript.ToDict()
		}(),
		"changeRankCapScript": func() map[string]interface{} {
			if p.ChangeRankCapScript == nil {
				return nil
			}
			return p.ChangeRankCapScript.ToDict()
		}(),
		"overflowExperienceScript": p.OverflowExperienceScript,
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
	ContextStack             *string             `json:"contextStack"`
	NamespaceName            *string             `json:"namespaceName"`
	Description              *string             `json:"description"`
	TransactionSetting       *TransactionSetting `json:"transactionSetting"`
	RankCapScriptId          *string             `json:"rankCapScriptId"`
	ChangeExperienceScript   *ScriptSetting      `json:"changeExperienceScript"`
	ChangeRankScript         *ScriptSetting      `json:"changeRankScript"`
	ChangeRankCapScript      *ScriptSetting      `json:"changeRankCapScript"`
	OverflowExperienceScript *string             `json:"overflowExperienceScript"`
	LogSetting               *LogSetting         `json:"logSetting"`
	DryRun                   *bool               `json:"dryRun"`
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["rankCapScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankCapScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankCapScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankCapScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankCapScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankCapScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankCapScriptId)
				}
			}
		}
		if v, ok := d["changeExperienceScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeExperienceScript)
		}
		if v, ok := d["changeRankScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRankScript)
		}
		if v, ok := d["changeRankCapScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRankCapScript)
		}
		if v, ok := d["overflowExperienceScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OverflowExperienceScript = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OverflowExperienceScript = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OverflowExperienceScript = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OverflowExperienceScript = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OverflowExperienceScript = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OverflowExperienceScript)
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
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		RankCapScriptId: func() *string {
			v, ok := data["rankCapScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankCapScriptId"])
		}(),
		ChangeExperienceScript: func() *ScriptSetting {
			v, ok := data["changeExperienceScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer()
		}(),
		ChangeRankScript: func() *ScriptSetting {
			v, ok := data["changeRankScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer()
		}(),
		ChangeRankCapScript: func() *ScriptSetting {
			v, ok := data["changeRankCapScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer()
		}(),
		OverflowExperienceScript: func() *string {
			v, ok := data["overflowExperienceScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["overflowExperienceScript"])
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
		"transactionSetting": func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}(),
		"rankCapScriptId": p.RankCapScriptId,
		"changeExperienceScript": func() map[string]interface{} {
			if p.ChangeExperienceScript == nil {
				return nil
			}
			return p.ChangeExperienceScript.ToDict()
		}(),
		"changeRankScript": func() map[string]interface{} {
			if p.ChangeRankScript == nil {
				return nil
			}
			return p.ChangeRankScript.ToDict()
		}(),
		"changeRankCapScript": func() map[string]interface{} {
			if p.ChangeRankCapScript == nil {
				return nil
			}
			return p.ChangeRankCapScript.ToDict()
		}(),
		"overflowExperienceScript": p.OverflowExperienceScript,
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

type GetServiceVersionRequest struct {
	ContextStack *string `json:"contextStack"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *GetServiceVersionRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetServiceVersionRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetServiceVersionRequest{}
	} else {
		*p = GetServiceVersionRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewGetServiceVersionRequestFromJson(data string) (GetServiceVersionRequest, error) {
	req := GetServiceVersionRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetServiceVersionRequest{}, err
	}
	return req, nil
}

func NewGetServiceVersionRequestFromDict(data map[string]interface{}) GetServiceVersionRequest {
	return GetServiceVersionRequest{}
}

func (p GetServiceVersionRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p GetServiceVersionRequest) Pointer() *GetServiceVersionRequest {
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

type DescribeExperienceModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeExperienceModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExperienceModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExperienceModelMastersRequest{}
	} else {
		*p = DescribeExperienceModelMastersRequest{}
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

func NewDescribeExperienceModelMastersRequestFromJson(data string) (DescribeExperienceModelMastersRequest, error) {
	req := DescribeExperienceModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExperienceModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeExperienceModelMastersRequestFromDict(data map[string]interface{}) DescribeExperienceModelMastersRequest {
	return DescribeExperienceModelMastersRequest{
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

func (p DescribeExperienceModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeExperienceModelMastersRequest) Pointer() *DescribeExperienceModelMastersRequest {
	return &p
}

type CreateExperienceModelMasterRequest struct {
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultExperience  *int64              `json:"defaultExperience"`
	DefaultRankCap     *int64              `json:"defaultRankCap"`
	MaxRankCap         *int64              `json:"maxRankCap"`
	RankThresholdName  *string             `json:"rankThresholdName"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
	DryRun             *bool               `json:"dryRun"`
}

func (p *CreateExperienceModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateExperienceModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateExperienceModelMasterRequest{}
	} else {
		*p = CreateExperienceModelMasterRequest{}
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
		if v, ok := d["defaultExperience"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultExperience)
		}
		if v, ok := d["defaultRankCap"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultRankCap)
		}
		if v, ok := d["maxRankCap"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxRankCap)
		}
		if v, ok := d["rankThresholdName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankThresholdName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankThresholdName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankThresholdName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankThresholdName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankThresholdName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankThresholdName)
				}
			}
		}
		if v, ok := d["acquireActionRates"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionRates)
		}
	}
	return nil
}

func NewCreateExperienceModelMasterRequestFromJson(data string) (CreateExperienceModelMasterRequest, error) {
	req := CreateExperienceModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateExperienceModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateExperienceModelMasterRequestFromDict(data map[string]interface{}) CreateExperienceModelMasterRequest {
	return CreateExperienceModelMasterRequest{
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
		DefaultExperience: func() *int64 {
			v, ok := data["defaultExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["defaultExperience"])
		}(),
		DefaultRankCap: func() *int64 {
			v, ok := data["defaultRankCap"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["defaultRankCap"])
		}(),
		MaxRankCap: func() *int64 {
			v, ok := data["maxRankCap"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maxRankCap"])
		}(),
		RankThresholdName: func() *string {
			v, ok := data["rankThresholdName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankThresholdName"])
		}(),
		AcquireActionRates: func() []AcquireActionRate {
			if data["acquireActionRates"] == nil {
				return nil
			}
			return CastAcquireActionRates(core.CastArray(data["acquireActionRates"]))
		}(),
	}
}

func (p CreateExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThresholdName": p.RankThresholdName,
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p CreateExperienceModelMasterRequest) Pointer() *CreateExperienceModelMasterRequest {
	return &p
}

type GetExperienceModelMasterRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *GetExperienceModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExperienceModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExperienceModelMasterRequest{}
	} else {
		*p = GetExperienceModelMasterRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
	}
	return nil
}

func NewGetExperienceModelMasterRequestFromJson(data string) (GetExperienceModelMasterRequest, error) {
	req := GetExperienceModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExperienceModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetExperienceModelMasterRequestFromDict(data map[string]interface{}) GetExperienceModelMasterRequest {
	return GetExperienceModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
	}
}

func (p GetExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p GetExperienceModelMasterRequest) Pointer() *GetExperienceModelMasterRequest {
	return &p
}

type UpdateExperienceModelMasterRequest struct {
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	ExperienceName     *string             `json:"experienceName"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultExperience  *int64              `json:"defaultExperience"`
	DefaultRankCap     *int64              `json:"defaultRankCap"`
	MaxRankCap         *int64              `json:"maxRankCap"`
	RankThresholdName  *string             `json:"rankThresholdName"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
	DryRun             *bool               `json:"dryRun"`
}

func (p *UpdateExperienceModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateExperienceModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateExperienceModelMasterRequest{}
	} else {
		*p = UpdateExperienceModelMasterRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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
		if v, ok := d["defaultExperience"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultExperience)
		}
		if v, ok := d["defaultRankCap"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultRankCap)
		}
		if v, ok := d["maxRankCap"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxRankCap)
		}
		if v, ok := d["rankThresholdName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankThresholdName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankThresholdName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankThresholdName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankThresholdName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankThresholdName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankThresholdName)
				}
			}
		}
		if v, ok := d["acquireActionRates"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionRates)
		}
	}
	return nil
}

func NewUpdateExperienceModelMasterRequestFromJson(data string) (UpdateExperienceModelMasterRequest, error) {
	req := UpdateExperienceModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateExperienceModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateExperienceModelMasterRequestFromDict(data map[string]interface{}) UpdateExperienceModelMasterRequest {
	return UpdateExperienceModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
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
		DefaultExperience: func() *int64 {
			v, ok := data["defaultExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["defaultExperience"])
		}(),
		DefaultRankCap: func() *int64 {
			v, ok := data["defaultRankCap"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["defaultRankCap"])
		}(),
		MaxRankCap: func() *int64 {
			v, ok := data["maxRankCap"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maxRankCap"])
		}(),
		RankThresholdName: func() *string {
			v, ok := data["rankThresholdName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankThresholdName"])
		}(),
		AcquireActionRates: func() []AcquireActionRate {
			if data["acquireActionRates"] == nil {
				return nil
			}
			return CastAcquireActionRates(core.CastArray(data["acquireActionRates"]))
		}(),
	}
}

func (p UpdateExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"experienceName":    p.ExperienceName,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThresholdName": p.RankThresholdName,
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p UpdateExperienceModelMasterRequest) Pointer() *UpdateExperienceModelMasterRequest {
	return &p
}

type DeleteExperienceModelMasterRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *DeleteExperienceModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteExperienceModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteExperienceModelMasterRequest{}
	} else {
		*p = DeleteExperienceModelMasterRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
	}
	return nil
}

func NewDeleteExperienceModelMasterRequestFromJson(data string) (DeleteExperienceModelMasterRequest, error) {
	req := DeleteExperienceModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteExperienceModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteExperienceModelMasterRequestFromDict(data map[string]interface{}) DeleteExperienceModelMasterRequest {
	return DeleteExperienceModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
	}
}

func (p DeleteExperienceModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p DeleteExperienceModelMasterRequest) Pointer() *DeleteExperienceModelMasterRequest {
	return &p
}

type DescribeExperienceModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeExperienceModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeExperienceModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeExperienceModelsRequest{}
	} else {
		*p = DescribeExperienceModelsRequest{}
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

func NewDescribeExperienceModelsRequestFromJson(data string) (DescribeExperienceModelsRequest, error) {
	req := DescribeExperienceModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeExperienceModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeExperienceModelsRequestFromDict(data map[string]interface{}) DescribeExperienceModelsRequest {
	return DescribeExperienceModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeExperienceModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeExperienceModelsRequest) Pointer() *DescribeExperienceModelsRequest {
	return &p
}

type GetExperienceModelRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *GetExperienceModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetExperienceModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetExperienceModelRequest{}
	} else {
		*p = GetExperienceModelRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
	}
	return nil
}

func NewGetExperienceModelRequestFromJson(data string) (GetExperienceModelRequest, error) {
	req := GetExperienceModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetExperienceModelRequest{}, err
	}
	return req, nil
}

func NewGetExperienceModelRequestFromDict(data map[string]interface{}) GetExperienceModelRequest {
	return GetExperienceModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
	}
}

func (p GetExperienceModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
	}
}

func (p GetExperienceModelRequest) Pointer() *GetExperienceModelRequest {
	return &p
}

type DescribeThresholdMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeThresholdMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeThresholdMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeThresholdMastersRequest{}
	} else {
		*p = DescribeThresholdMastersRequest{}
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

func NewDescribeThresholdMastersRequestFromJson(data string) (DescribeThresholdMastersRequest, error) {
	req := DescribeThresholdMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeThresholdMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeThresholdMastersRequestFromDict(data map[string]interface{}) DescribeThresholdMastersRequest {
	return DescribeThresholdMastersRequest{
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

func (p DescribeThresholdMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeThresholdMastersRequest) Pointer() *DescribeThresholdMastersRequest {
	return &p
}

type CreateThresholdMasterRequest struct {
	ContextStack  *string  `json:"contextStack"`
	NamespaceName *string  `json:"namespaceName"`
	Name          *string  `json:"name"`
	Description   *string  `json:"description"`
	Metadata      *string  `json:"metadata"`
	Values        []*int64 `json:"values"`
	DryRun        *bool    `json:"dryRun"`
}

func (p *CreateThresholdMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateThresholdMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateThresholdMasterRequest{}
	} else {
		*p = CreateThresholdMasterRequest{}
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewCreateThresholdMasterRequestFromJson(data string) (CreateThresholdMasterRequest, error) {
	req := CreateThresholdMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateThresholdMasterRequest{}, err
	}
	return req, nil
}

func NewCreateThresholdMasterRequestFromDict(data map[string]interface{}) CreateThresholdMasterRequest {
	return CreateThresholdMasterRequest{
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
		Values: func() []*int64 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64s(core.CastArray(v))
		}(),
	}
}

func (p CreateThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
	}
}

func (p CreateThresholdMasterRequest) Pointer() *CreateThresholdMasterRequest {
	return &p
}

type GetThresholdMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ThresholdName *string `json:"thresholdName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetThresholdMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetThresholdMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetThresholdMasterRequest{}
	} else {
		*p = GetThresholdMasterRequest{}
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
		if v, ok := d["thresholdName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ThresholdName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ThresholdName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ThresholdName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ThresholdName)
				}
			}
		}
	}
	return nil
}

func NewGetThresholdMasterRequestFromJson(data string) (GetThresholdMasterRequest, error) {
	req := GetThresholdMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetThresholdMasterRequest{}, err
	}
	return req, nil
}

func NewGetThresholdMasterRequestFromDict(data map[string]interface{}) GetThresholdMasterRequest {
	return GetThresholdMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ThresholdName: func() *string {
			v, ok := data["thresholdName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["thresholdName"])
		}(),
	}
}

func (p GetThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
	}
}

func (p GetThresholdMasterRequest) Pointer() *GetThresholdMasterRequest {
	return &p
}

type UpdateThresholdMasterRequest struct {
	ContextStack  *string  `json:"contextStack"`
	NamespaceName *string  `json:"namespaceName"`
	ThresholdName *string  `json:"thresholdName"`
	Description   *string  `json:"description"`
	Metadata      *string  `json:"metadata"`
	Values        []*int64 `json:"values"`
	DryRun        *bool    `json:"dryRun"`
}

func (p *UpdateThresholdMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateThresholdMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateThresholdMasterRequest{}
	} else {
		*p = UpdateThresholdMasterRequest{}
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
		if v, ok := d["thresholdName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ThresholdName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ThresholdName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ThresholdName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ThresholdName)
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewUpdateThresholdMasterRequestFromJson(data string) (UpdateThresholdMasterRequest, error) {
	req := UpdateThresholdMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateThresholdMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateThresholdMasterRequestFromDict(data map[string]interface{}) UpdateThresholdMasterRequest {
	return UpdateThresholdMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ThresholdName: func() *string {
			v, ok := data["thresholdName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["thresholdName"])
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
		Values: func() []*int64 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64s(core.CastArray(v))
		}(),
	}
}

func (p UpdateThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
	}
}

func (p UpdateThresholdMasterRequest) Pointer() *UpdateThresholdMasterRequest {
	return &p
}

type DeleteThresholdMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	ThresholdName *string `json:"thresholdName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteThresholdMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteThresholdMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteThresholdMasterRequest{}
	} else {
		*p = DeleteThresholdMasterRequest{}
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
		if v, ok := d["thresholdName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ThresholdName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ThresholdName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ThresholdName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ThresholdName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ThresholdName)
				}
			}
		}
	}
	return nil
}

func NewDeleteThresholdMasterRequestFromJson(data string) (DeleteThresholdMasterRequest, error) {
	req := DeleteThresholdMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteThresholdMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteThresholdMasterRequestFromDict(data map[string]interface{}) DeleteThresholdMasterRequest {
	return DeleteThresholdMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ThresholdName: func() *string {
			v, ok := data["thresholdName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["thresholdName"])
		}(),
	}
}

func (p DeleteThresholdMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"thresholdName": p.ThresholdName,
	}
}

func (p DeleteThresholdMasterRequest) Pointer() *DeleteThresholdMasterRequest {
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

type GetCurrentExperienceMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetCurrentExperienceMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentExperienceMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentExperienceMasterRequest{}
	} else {
		*p = GetCurrentExperienceMasterRequest{}
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

func NewGetCurrentExperienceMasterRequestFromJson(data string) (GetCurrentExperienceMasterRequest, error) {
	req := GetCurrentExperienceMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentExperienceMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentExperienceMasterRequestFromDict(data map[string]interface{}) GetCurrentExperienceMasterRequest {
	return GetCurrentExperienceMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentExperienceMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentExperienceMasterRequest) Pointer() *GetCurrentExperienceMasterRequest {
	return &p
}

type PreUpdateCurrentExperienceMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *PreUpdateCurrentExperienceMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PreUpdateCurrentExperienceMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PreUpdateCurrentExperienceMasterRequest{}
	} else {
		*p = PreUpdateCurrentExperienceMasterRequest{}
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

func NewPreUpdateCurrentExperienceMasterRequestFromJson(data string) (PreUpdateCurrentExperienceMasterRequest, error) {
	req := PreUpdateCurrentExperienceMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PreUpdateCurrentExperienceMasterRequest{}, err
	}
	return req, nil
}

func NewPreUpdateCurrentExperienceMasterRequestFromDict(data map[string]interface{}) PreUpdateCurrentExperienceMasterRequest {
	return PreUpdateCurrentExperienceMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p PreUpdateCurrentExperienceMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p PreUpdateCurrentExperienceMasterRequest) Pointer() *PreUpdateCurrentExperienceMasterRequest {
	return &p
}

type UpdateCurrentExperienceMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Mode          *string `json:"mode"`
	Settings      *string `json:"settings"`
	UploadToken   *string `json:"uploadToken"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *UpdateCurrentExperienceMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentExperienceMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentExperienceMasterRequest{}
	} else {
		*p = UpdateCurrentExperienceMasterRequest{}
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

func NewUpdateCurrentExperienceMasterRequestFromJson(data string) (UpdateCurrentExperienceMasterRequest, error) {
	req := UpdateCurrentExperienceMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentExperienceMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentExperienceMasterRequestFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterRequest {
	return UpdateCurrentExperienceMasterRequest{
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

func (p UpdateCurrentExperienceMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"mode":          p.Mode,
		"settings":      p.Settings,
		"uploadToken":   p.UploadToken,
	}
}

func (p UpdateCurrentExperienceMasterRequest) Pointer() *UpdateCurrentExperienceMasterRequest {
	return &p
}

type UpdateCurrentExperienceMasterFromGitHubRequest struct {
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
	DryRun          *bool                  `json:"dryRun"`
}

func (p *UpdateCurrentExperienceMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentExperienceMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentExperienceMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentExperienceMasterFromGitHubRequest{}
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

func NewUpdateCurrentExperienceMasterFromGitHubRequestFromJson(data string) (UpdateCurrentExperienceMasterFromGitHubRequest, error) {
	req := UpdateCurrentExperienceMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentExperienceMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentExperienceMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentExperienceMasterFromGitHubRequest {
	return UpdateCurrentExperienceMasterFromGitHubRequest{
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

func (p UpdateCurrentExperienceMasterFromGitHubRequest) ToDict() map[string]interface{} {
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

func (p UpdateCurrentExperienceMasterFromGitHubRequest) Pointer() *UpdateCurrentExperienceMasterFromGitHubRequest {
	return &p
}

type DescribeStatusesRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	ExperienceName *string `json:"experienceName"`
	AccessToken    *string `json:"accessToken"`
	PageToken      *string `json:"pageToken"`
	Limit          *int32  `json:"limit"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *DescribeStatusesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStatusesRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStatusesRequest{}
	} else {
		*p = DescribeStatusesRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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

func NewDescribeStatusesRequestFromJson(data string) (DescribeStatusesRequest, error) {
	req := DescribeStatusesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStatusesRequest{}, err
	}
	return req, nil
}

func NewDescribeStatusesRequestFromDict(data map[string]interface{}) DescribeStatusesRequest {
	return DescribeStatusesRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
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

func (p DescribeStatusesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"experienceName": p.ExperienceName,
		"accessToken":    p.AccessToken,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
	return &p
}

type DescribeStatusesByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	ExperienceName  *string `json:"experienceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeStatusesByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStatusesByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStatusesByUserIdRequest{}
	} else {
		*p = DescribeStatusesByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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

func NewDescribeStatusesByUserIdRequestFromJson(data string) (DescribeStatusesByUserIdRequest, error) {
	req := DescribeStatusesByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStatusesByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeStatusesByUserIdRequestFromDict(data map[string]interface{}) DescribeStatusesByUserIdRequest {
	return DescribeStatusesByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
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

func (p DescribeStatusesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"experienceName":  p.ExperienceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeStatusesByUserIdRequest) Pointer() *DescribeStatusesByUserIdRequest {
	return &p
}

type GetStatusRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AccessToken    *string `json:"accessToken"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *GetStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStatusRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStatusRequest{}
	} else {
		*p = GetStatusRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
	}
	return nil
}

func NewGetStatusRequestFromJson(data string) (GetStatusRequest, error) {
	req := GetStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStatusRequest{}, err
	}
	return req, nil
}

func NewGetStatusRequestFromDict(data map[string]interface{}) GetStatusRequest {
	return GetStatusRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
	}
}

func (p GetStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
	}
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
	return &p
}

type GetStatusByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	ExperienceName  *string `json:"experienceName"`
	PropertyId      *string `json:"propertyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetStatusByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStatusByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStatusByUserIdRequest{}
	} else {
		*p = GetStatusByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
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

func NewGetStatusByUserIdRequestFromJson(data string) (GetStatusByUserIdRequest, error) {
	req := GetStatusByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStatusByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetStatusByUserIdRequestFromDict(data map[string]interface{}) GetStatusByUserIdRequest {
	return GetStatusByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
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

func (p GetStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type GetStatusWithSignatureRequest struct {
	ContextStack   *string `json:"contextStack"`
	NamespaceName  *string `json:"namespaceName"`
	AccessToken    *string `json:"accessToken"`
	ExperienceName *string `json:"experienceName"`
	PropertyId     *string `json:"propertyId"`
	KeyId          *string `json:"keyId"`
	DryRun         *bool   `json:"dryRun"`
}

func (p *GetStatusWithSignatureRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStatusWithSignatureRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStatusWithSignatureRequest{}
	} else {
		*p = GetStatusWithSignatureRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
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

func NewGetStatusWithSignatureRequestFromJson(data string) (GetStatusWithSignatureRequest, error) {
	req := GetStatusWithSignatureRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStatusWithSignatureRequest{}, err
	}
	return req, nil
}

func NewGetStatusWithSignatureRequestFromDict(data map[string]interface{}) GetStatusWithSignatureRequest {
	return GetStatusWithSignatureRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p GetStatusWithSignatureRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"keyId":          p.KeyId,
	}
}

func (p GetStatusWithSignatureRequest) Pointer() *GetStatusWithSignatureRequest {
	return &p
}

type GetStatusWithSignatureByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	ExperienceName  *string `json:"experienceName"`
	PropertyId      *string `json:"propertyId"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetStatusWithSignatureByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStatusWithSignatureByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStatusWithSignatureByUserIdRequest{}
	} else {
		*p = GetStatusWithSignatureByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
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

func NewGetStatusWithSignatureByUserIdRequestFromJson(data string) (GetStatusWithSignatureByUserIdRequest, error) {
	req := GetStatusWithSignatureByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStatusWithSignatureByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetStatusWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetStatusWithSignatureByUserIdRequest {
	return GetStatusWithSignatureByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
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

func (p GetStatusWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetStatusWithSignatureByUserIdRequest) Pointer() *GetStatusWithSignatureByUserIdRequest {
	return &p
}

type AddExperienceByUserIdRequest struct {
	ContextStack                 *string `json:"contextStack"`
	DuplicationAvoider           *string `json:"duplicationAvoider"`
	NamespaceName                *string `json:"namespaceName"`
	UserId                       *string `json:"userId"`
	ExperienceName               *string `json:"experienceName"`
	PropertyId                   *string `json:"propertyId"`
	ExperienceValue              *int64  `json:"experienceValue"`
	TruncateExperienceWhenRankUp *bool   `json:"truncateExperienceWhenRankUp"`
	TimeOffsetToken              *string `json:"timeOffsetToken"`
	DryRun                       *bool   `json:"dryRun"`
}

func (p *AddExperienceByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddExperienceByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddExperienceByUserIdRequest{}
	} else {
		*p = AddExperienceByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["experienceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExperienceValue)
		}
		if v, ok := d["truncateExperienceWhenRankUp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TruncateExperienceWhenRankUp)
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

func NewAddExperienceByUserIdRequestFromJson(data string) (AddExperienceByUserIdRequest, error) {
	req := AddExperienceByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddExperienceByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddExperienceByUserIdRequestFromDict(data map[string]interface{}) AddExperienceByUserIdRequest {
	return AddExperienceByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ExperienceValue: func() *int64 {
			v, ok := data["experienceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["experienceValue"])
		}(),
		TruncateExperienceWhenRankUp: func() *bool {
			v, ok := data["truncateExperienceWhenRankUp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["truncateExperienceWhenRankUp"])
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

func (p AddExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                p.NamespaceName,
		"userId":                       p.UserId,
		"experienceName":               p.ExperienceName,
		"propertyId":                   p.PropertyId,
		"experienceValue":              p.ExperienceValue,
		"truncateExperienceWhenRankUp": p.TruncateExperienceWhenRankUp,
		"timeOffsetToken":              p.TimeOffsetToken,
	}
}

func (p AddExperienceByUserIdRequest) Pointer() *AddExperienceByUserIdRequest {
	return &p
}

type SubExperienceRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubExperienceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubExperienceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubExperienceRequest{}
	} else {
		*p = SubExperienceRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["experienceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExperienceValue)
		}
	}
	return nil
}

func NewSubExperienceRequestFromJson(data string) (SubExperienceRequest, error) {
	req := SubExperienceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubExperienceRequest{}, err
	}
	return req, nil
}

func NewSubExperienceRequestFromDict(data map[string]interface{}) SubExperienceRequest {
	return SubExperienceRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ExperienceValue: func() *int64 {
			v, ok := data["experienceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["experienceValue"])
		}(),
	}
}

func (p SubExperienceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
	}
}

func (p SubExperienceRequest) Pointer() *SubExperienceRequest {
	return &p
}

type SubExperienceByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubExperienceByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubExperienceByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubExperienceByUserIdRequest{}
	} else {
		*p = SubExperienceByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["experienceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExperienceValue)
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

func NewSubExperienceByUserIdRequestFromJson(data string) (SubExperienceByUserIdRequest, error) {
	req := SubExperienceByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubExperienceByUserIdRequest{}, err
	}
	return req, nil
}

func NewSubExperienceByUserIdRequestFromDict(data map[string]interface{}) SubExperienceByUserIdRequest {
	return SubExperienceByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ExperienceValue: func() *int64 {
			v, ok := data["experienceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["experienceValue"])
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

func (p SubExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SubExperienceByUserIdRequest) Pointer() *SubExperienceByUserIdRequest {
	return &p
}

type SetExperienceByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	ExperienceValue    *int64  `json:"experienceValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SetExperienceByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetExperienceByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetExperienceByUserIdRequest{}
	} else {
		*p = SetExperienceByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["experienceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExperienceValue)
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

func NewSetExperienceByUserIdRequestFromJson(data string) (SetExperienceByUserIdRequest, error) {
	req := SetExperienceByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetExperienceByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetExperienceByUserIdRequestFromDict(data map[string]interface{}) SetExperienceByUserIdRequest {
	return SetExperienceByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ExperienceValue: func() *int64 {
			v, ok := data["experienceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["experienceValue"])
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

func (p SetExperienceByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetExperienceByUserIdRequest) Pointer() *SetExperienceByUserIdRequest {
	return &p
}

type AddRankCapByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AddRankCapByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddRankCapByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddRankCapByUserIdRequest{}
	} else {
		*p = AddRankCapByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
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

func NewAddRankCapByUserIdRequestFromJson(data string) (AddRankCapByUserIdRequest, error) {
	req := AddRankCapByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddRankCapByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddRankCapByUserIdRequestFromDict(data map[string]interface{}) AddRankCapByUserIdRequest {
	return AddRankCapByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
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

func (p AddRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"rankCapValue":    p.RankCapValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddRankCapByUserIdRequest) Pointer() *AddRankCapByUserIdRequest {
	return &p
}

type SubRankCapRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubRankCapRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubRankCapRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubRankCapRequest{}
	} else {
		*p = SubRankCapRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
		}
	}
	return nil
}

func NewSubRankCapRequestFromJson(data string) (SubRankCapRequest, error) {
	req := SubRankCapRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubRankCapRequest{}, err
	}
	return req, nil
}

func NewSubRankCapRequestFromDict(data map[string]interface{}) SubRankCapRequest {
	return SubRankCapRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
		}(),
	}
}

func (p SubRankCapRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rankCapValue":   p.RankCapValue,
	}
}

func (p SubRankCapRequest) Pointer() *SubRankCapRequest {
	return &p
}

type SubRankCapByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubRankCapByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubRankCapByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubRankCapByUserIdRequest{}
	} else {
		*p = SubRankCapByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
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

func NewSubRankCapByUserIdRequestFromJson(data string) (SubRankCapByUserIdRequest, error) {
	req := SubRankCapByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubRankCapByUserIdRequest{}, err
	}
	return req, nil
}

func NewSubRankCapByUserIdRequestFromDict(data map[string]interface{}) SubRankCapByUserIdRequest {
	return SubRankCapByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
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

func (p SubRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"rankCapValue":    p.RankCapValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SubRankCapByUserIdRequest) Pointer() *SubRankCapByUserIdRequest {
	return &p
}

type SetRankCapByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	RankCapValue       *int64  `json:"rankCapValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SetRankCapByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRankCapByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRankCapByUserIdRequest{}
	} else {
		*p = SetRankCapByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
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

func NewSetRankCapByUserIdRequestFromJson(data string) (SetRankCapByUserIdRequest, error) {
	req := SetRankCapByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRankCapByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetRankCapByUserIdRequestFromDict(data map[string]interface{}) SetRankCapByUserIdRequest {
	return SetRankCapByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
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

func (p SetRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"rankCapValue":    p.RankCapValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetRankCapByUserIdRequest) Pointer() *SetRankCapByUserIdRequest {
	return &p
}

type DeleteStatusByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	ExperienceName     *string `json:"experienceName"`
	PropertyId         *string `json:"propertyId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteStatusByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStatusByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteStatusByUserIdRequest{}
	} else {
		*p = DeleteStatusByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
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

func NewDeleteStatusByUserIdRequestFromJson(data string) (DeleteStatusByUserIdRequest, error) {
	req := DeleteStatusByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStatusByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteStatusByUserIdRequestFromDict(data map[string]interface{}) DeleteStatusByUserIdRequest {
	return DeleteStatusByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
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

func (p DeleteStatusByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"experienceName":  p.ExperienceName,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
	return &p
}

type VerifyRankRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	ExperienceName                  *string `json:"experienceName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	RankValue                       *int64  `json:"rankValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyRankRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankRequest{}
	} else {
		*p = VerifyRankRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankValue)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyRankRequestFromJson(data string) (VerifyRankRequest, error) {
	req := VerifyRankRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankRequest{}, err
	}
	return req, nil
}

func NewVerifyRankRequestFromDict(data map[string]interface{}) VerifyRankRequest {
	return VerifyRankRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankValue: func() *int64 {
			v, ok := data["rankValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankValue"])
		}(),
		MultiplyValueSpecifyingQuantity: func() *bool {
			v, ok := data["multiplyValueSpecifyingQuantity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["multiplyValueSpecifyingQuantity"])
		}(),
	}
}

func (p VerifyRankRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"experienceName":                  p.ExperienceName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"rankValue":                       p.RankValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyRankRequest) Pointer() *VerifyRankRequest {
	return &p
}

type VerifyRankByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	ExperienceName                  *string `json:"experienceName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	RankValue                       *int64  `json:"rankValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyRankByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankByUserIdRequest{}
	} else {
		*p = VerifyRankByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankValue)
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

func NewVerifyRankByUserIdRequestFromJson(data string) (VerifyRankByUserIdRequest, error) {
	req := VerifyRankByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyRankByUserIdRequestFromDict(data map[string]interface{}) VerifyRankByUserIdRequest {
	return VerifyRankByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankValue: func() *int64 {
			v, ok := data["rankValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankValue"])
		}(),
		MultiplyValueSpecifyingQuantity: func() *bool {
			v, ok := data["multiplyValueSpecifyingQuantity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["multiplyValueSpecifyingQuantity"])
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

func (p VerifyRankByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"experienceName":                  p.ExperienceName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"rankValue":                       p.RankValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyRankByUserIdRequest) Pointer() *VerifyRankByUserIdRequest {
	return &p
}

type VerifyRankCapRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	ExperienceName                  *string `json:"experienceName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	RankCapValue                    *int64  `json:"rankCapValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyRankCapRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankCapRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankCapRequest{}
	} else {
		*p = VerifyRankCapRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyRankCapRequestFromJson(data string) (VerifyRankCapRequest, error) {
	req := VerifyRankCapRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankCapRequest{}, err
	}
	return req, nil
}

func NewVerifyRankCapRequestFromDict(data map[string]interface{}) VerifyRankCapRequest {
	return VerifyRankCapRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
		}(),
		MultiplyValueSpecifyingQuantity: func() *bool {
			v, ok := data["multiplyValueSpecifyingQuantity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["multiplyValueSpecifyingQuantity"])
		}(),
	}
}

func (p VerifyRankCapRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"experienceName":                  p.ExperienceName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"rankCapValue":                    p.RankCapValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyRankCapRequest) Pointer() *VerifyRankCapRequest {
	return &p
}

type VerifyRankCapByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	ExperienceName                  *string `json:"experienceName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	RankCapValue                    *int64  `json:"rankCapValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyRankCapByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankCapByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankCapByUserIdRequest{}
	} else {
		*p = VerifyRankCapByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
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
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rankCapValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankCapValue)
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

func NewVerifyRankCapByUserIdRequestFromJson(data string) (VerifyRankCapByUserIdRequest, error) {
	req := VerifyRankCapByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankCapByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyRankCapByUserIdRequestFromDict(data map[string]interface{}) VerifyRankCapByUserIdRequest {
	return VerifyRankCapByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RankCapValue: func() *int64 {
			v, ok := data["rankCapValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["rankCapValue"])
		}(),
		MultiplyValueSpecifyingQuantity: func() *bool {
			v, ok := data["multiplyValueSpecifyingQuantity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["multiplyValueSpecifyingQuantity"])
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

func (p VerifyRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"experienceName":                  p.ExperienceName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"rankCapValue":                    p.RankCapValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyRankCapByUserIdRequest) Pointer() *VerifyRankCapByUserIdRequest {
	return &p
}

type AddExperienceByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *AddExperienceByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddExperienceByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddExperienceByStampSheetRequest{}
	} else {
		*p = AddExperienceByStampSheetRequest{}
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

func NewAddExperienceByStampSheetRequestFromJson(data string) (AddExperienceByStampSheetRequest, error) {
	req := AddExperienceByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddExperienceByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAddExperienceByStampSheetRequestFromDict(data map[string]interface{}) AddExperienceByStampSheetRequest {
	return AddExperienceByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p AddExperienceByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddExperienceByStampSheetRequest) Pointer() *AddExperienceByStampSheetRequest {
	return &p
}

type SetExperienceByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *SetExperienceByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetExperienceByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetExperienceByStampSheetRequest{}
	} else {
		*p = SetExperienceByStampSheetRequest{}
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

func NewSetExperienceByStampSheetRequestFromJson(data string) (SetExperienceByStampSheetRequest, error) {
	req := SetExperienceByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetExperienceByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetExperienceByStampSheetRequestFromDict(data map[string]interface{}) SetExperienceByStampSheetRequest {
	return SetExperienceByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p SetExperienceByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetExperienceByStampSheetRequest) Pointer() *SetExperienceByStampSheetRequest {
	return &p
}

type SubExperienceByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *SubExperienceByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubExperienceByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubExperienceByStampTaskRequest{}
	} else {
		*p = SubExperienceByStampTaskRequest{}
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

func NewSubExperienceByStampTaskRequestFromJson(data string) (SubExperienceByStampTaskRequest, error) {
	req := SubExperienceByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubExperienceByStampTaskRequest{}, err
	}
	return req, nil
}

func NewSubExperienceByStampTaskRequestFromDict(data map[string]interface{}) SubExperienceByStampTaskRequest {
	return SubExperienceByStampTaskRequest{
		StampTask: func() *string {
			v, ok := data["stampTask"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampTask"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p SubExperienceByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubExperienceByStampTaskRequest) Pointer() *SubExperienceByStampTaskRequest {
	return &p
}

type AddRankCapByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *AddRankCapByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddRankCapByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddRankCapByStampSheetRequest{}
	} else {
		*p = AddRankCapByStampSheetRequest{}
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

func NewAddRankCapByStampSheetRequestFromJson(data string) (AddRankCapByStampSheetRequest, error) {
	req := AddRankCapByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddRankCapByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAddRankCapByStampSheetRequestFromDict(data map[string]interface{}) AddRankCapByStampSheetRequest {
	return AddRankCapByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p AddRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddRankCapByStampSheetRequest) Pointer() *AddRankCapByStampSheetRequest {
	return &p
}

type SubRankCapByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *SubRankCapByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubRankCapByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubRankCapByStampTaskRequest{}
	} else {
		*p = SubRankCapByStampTaskRequest{}
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

func NewSubRankCapByStampTaskRequestFromJson(data string) (SubRankCapByStampTaskRequest, error) {
	req := SubRankCapByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubRankCapByStampTaskRequest{}, err
	}
	return req, nil
}

func NewSubRankCapByStampTaskRequestFromDict(data map[string]interface{}) SubRankCapByStampTaskRequest {
	return SubRankCapByStampTaskRequest{
		StampTask: func() *string {
			v, ok := data["stampTask"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampTask"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p SubRankCapByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubRankCapByStampTaskRequest) Pointer() *SubRankCapByStampTaskRequest {
	return &p
}

type SetRankCapByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *SetRankCapByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRankCapByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRankCapByStampSheetRequest{}
	} else {
		*p = SetRankCapByStampSheetRequest{}
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

func NewSetRankCapByStampSheetRequestFromJson(data string) (SetRankCapByStampSheetRequest, error) {
	req := SetRankCapByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRankCapByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetRankCapByStampSheetRequestFromDict(data map[string]interface{}) SetRankCapByStampSheetRequest {
	return SetRankCapByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p SetRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRankCapByStampSheetRequest) Pointer() *SetRankCapByStampSheetRequest {
	return &p
}

type MultiplyAcquireActionsByUserIdRequest struct {
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	ExperienceName     *string         `json:"experienceName"`
	PropertyId         *string         `json:"propertyId"`
	RateName           *string         `json:"rateName"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
	BaseRate           *float32        `json:"baseRate"`
	TimeOffsetToken    *string         `json:"timeOffsetToken"`
	DryRun             *bool           `json:"dryRun"`
}

func (p *MultiplyAcquireActionsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MultiplyAcquireActionsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MultiplyAcquireActionsByUserIdRequest{}
	} else {
		*p = MultiplyAcquireActionsByUserIdRequest{}
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
		if v, ok := d["experienceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["rateName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RateName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RateName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RateName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RateName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RateName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RateName)
				}
			}
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
		if v, ok := d["baseRate"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BaseRate)
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

func NewMultiplyAcquireActionsByUserIdRequestFromJson(data string) (MultiplyAcquireActionsByUserIdRequest, error) {
	req := MultiplyAcquireActionsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return MultiplyAcquireActionsByUserIdRequest{}, err
	}
	return req, nil
}

func NewMultiplyAcquireActionsByUserIdRequestFromDict(data map[string]interface{}) MultiplyAcquireActionsByUserIdRequest {
	return MultiplyAcquireActionsByUserIdRequest{
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
		ExperienceName: func() *string {
			v, ok := data["experienceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		RateName: func() *string {
			v, ok := data["rateName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rateName"])
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
		}(),
		BaseRate: func() *float32 {
			v, ok := data["baseRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["baseRate"])
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

func (p MultiplyAcquireActionsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"experienceName": p.ExperienceName,
		"propertyId":     p.PropertyId,
		"rateName":       p.RateName,
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
		"baseRate":        p.BaseRate,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p MultiplyAcquireActionsByUserIdRequest) Pointer() *MultiplyAcquireActionsByUserIdRequest {
	return &p
}

type MultiplyAcquireActionsByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *MultiplyAcquireActionsByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MultiplyAcquireActionsByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MultiplyAcquireActionsByStampSheetRequest{}
	} else {
		*p = MultiplyAcquireActionsByStampSheetRequest{}
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

func NewMultiplyAcquireActionsByStampSheetRequestFromJson(data string) (MultiplyAcquireActionsByStampSheetRequest, error) {
	req := MultiplyAcquireActionsByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return MultiplyAcquireActionsByStampSheetRequest{}, err
	}
	return req, nil
}

func NewMultiplyAcquireActionsByStampSheetRequestFromDict(data map[string]interface{}) MultiplyAcquireActionsByStampSheetRequest {
	return MultiplyAcquireActionsByStampSheetRequest{
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p MultiplyAcquireActionsByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p MultiplyAcquireActionsByStampSheetRequest) Pointer() *MultiplyAcquireActionsByStampSheetRequest {
	return &p
}

type VerifyRankByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyRankByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankByStampTaskRequest{}
	} else {
		*p = VerifyRankByStampTaskRequest{}
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

func NewVerifyRankByStampTaskRequestFromJson(data string) (VerifyRankByStampTaskRequest, error) {
	req := VerifyRankByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyRankByStampTaskRequestFromDict(data map[string]interface{}) VerifyRankByStampTaskRequest {
	return VerifyRankByStampTaskRequest{
		StampTask: func() *string {
			v, ok := data["stampTask"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampTask"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p VerifyRankByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyRankByStampTaskRequest) Pointer() *VerifyRankByStampTaskRequest {
	return &p
}

type VerifyRankCapByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyRankCapByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyRankCapByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyRankCapByStampTaskRequest{}
	} else {
		*p = VerifyRankCapByStampTaskRequest{}
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

func NewVerifyRankCapByStampTaskRequestFromJson(data string) (VerifyRankCapByStampTaskRequest, error) {
	req := VerifyRankCapByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyRankCapByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyRankCapByStampTaskRequestFromDict(data map[string]interface{}) VerifyRankCapByStampTaskRequest {
	return VerifyRankCapByStampTaskRequest{
		StampTask: func() *string {
			v, ok := data["stampTask"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampTask"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
	}
}

func (p VerifyRankCapByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyRankCapByStampTaskRequest) Pointer() *VerifyRankCapByStampTaskRequest {
	return &p
}
