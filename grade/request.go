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

package grade

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
	ContextStack       *string             `json:"contextStack"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ChangeGradeScript  *ScriptSetting      `json:"changeGradeScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	DryRun             *bool               `json:"dryRun"`
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
		if v, ok := d["changeGradeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeGradeScript)
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
		ChangeGradeScript: func() *ScriptSetting {
			v, ok := data["changeGradeScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeGradeScript"])).Pointer()
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
		"changeGradeScript": func() map[string]interface{} {
			if p.ChangeGradeScript == nil {
				return nil
			}
			return p.ChangeGradeScript.ToDict()
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
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ChangeGradeScript  *ScriptSetting      `json:"changeGradeScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	DryRun             *bool               `json:"dryRun"`
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
		if v, ok := d["changeGradeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeGradeScript)
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
		ChangeGradeScript: func() *ScriptSetting {
			v, ok := data["changeGradeScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeGradeScript"])).Pointer()
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
		"changeGradeScript": func() map[string]interface{} {
			if p.ChangeGradeScript == nil {
				return nil
			}
			return p.ChangeGradeScript.ToDict()
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

type DescribeGradeModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGradeModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGradeModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGradeModelMastersRequest{}
	} else {
		*p = DescribeGradeModelMastersRequest{}
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

func NewDescribeGradeModelMastersRequestFromJson(data string) (DescribeGradeModelMastersRequest, error) {
	req := DescribeGradeModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGradeModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeGradeModelMastersRequestFromDict(data map[string]interface{}) DescribeGradeModelMastersRequest {
	return DescribeGradeModelMastersRequest{
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

func (p DescribeGradeModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGradeModelMastersRequest) Pointer() *DescribeGradeModelMastersRequest {
	return &p
}

type CreateGradeModelMasterRequest struct {
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
	DryRun             *bool               `json:"dryRun"`
}

func (p *CreateGradeModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGradeModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGradeModelMasterRequest{}
	} else {
		*p = CreateGradeModelMasterRequest{}
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
		if v, ok := d["defaultGrades"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultGrades)
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceModelId)
				}
			}
		}
		if v, ok := d["gradeEntries"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeEntries)
		}
		if v, ok := d["acquireActionRates"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionRates)
		}
	}
	return nil
}

func NewCreateGradeModelMasterRequestFromJson(data string) (CreateGradeModelMasterRequest, error) {
	req := CreateGradeModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGradeModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateGradeModelMasterRequestFromDict(data map[string]interface{}) CreateGradeModelMasterRequest {
	return CreateGradeModelMasterRequest{
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
		DefaultGrades: func() []DefaultGradeModel {
			if data["defaultGrades"] == nil {
				return nil
			}
			return CastDefaultGradeModels(core.CastArray(data["defaultGrades"]))
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		GradeEntries: func() []GradeEntryModel {
			if data["gradeEntries"] == nil {
				return nil
			}
			return CastGradeEntryModels(core.CastArray(data["gradeEntries"]))
		}(),
		AcquireActionRates: func() []AcquireActionRate {
			if data["acquireActionRates"] == nil {
				return nil
			}
			return CastAcquireActionRates(core.CastArray(data["acquireActionRates"]))
		}(),
	}
}

func (p CreateGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"defaultGrades": CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		),
		"experienceModelId": p.ExperienceModelId,
		"gradeEntries": CastGradeEntryModelsFromDict(
			p.GradeEntries,
		),
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p CreateGradeModelMasterRequest) Pointer() *CreateGradeModelMasterRequest {
	return &p
}

type GetGradeModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGradeModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGradeModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGradeModelMasterRequest{}
	} else {
		*p = GetGradeModelMasterRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
				}
			}
		}
	}
	return nil
}

func NewGetGradeModelMasterRequestFromJson(data string) (GetGradeModelMasterRequest, error) {
	req := GetGradeModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGradeModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetGradeModelMasterRequestFromDict(data map[string]interface{}) GetGradeModelMasterRequest {
	return GetGradeModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
	}
}

func (p GetGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p GetGradeModelMasterRequest) Pointer() *GetGradeModelMasterRequest {
	return &p
}

type UpdateGradeModelMasterRequest struct {
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	GradeName          *string             `json:"gradeName"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
	DryRun             *bool               `json:"dryRun"`
}

func (p *UpdateGradeModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGradeModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGradeModelMasterRequest{}
	} else {
		*p = UpdateGradeModelMasterRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["defaultGrades"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultGrades)
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ExperienceModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ExperienceModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ExperienceModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ExperienceModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ExperienceModelId)
				}
			}
		}
		if v, ok := d["gradeEntries"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeEntries)
		}
		if v, ok := d["acquireActionRates"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionRates)
		}
	}
	return nil
}

func NewUpdateGradeModelMasterRequestFromJson(data string) (UpdateGradeModelMasterRequest, error) {
	req := UpdateGradeModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGradeModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateGradeModelMasterRequestFromDict(data map[string]interface{}) UpdateGradeModelMasterRequest {
	return UpdateGradeModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		DefaultGrades: func() []DefaultGradeModel {
			if data["defaultGrades"] == nil {
				return nil
			}
			return CastDefaultGradeModels(core.CastArray(data["defaultGrades"]))
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		GradeEntries: func() []GradeEntryModel {
			if data["gradeEntries"] == nil {
				return nil
			}
			return CastGradeEntryModels(core.CastArray(data["gradeEntries"]))
		}(),
		AcquireActionRates: func() []AcquireActionRate {
			if data["acquireActionRates"] == nil {
				return nil
			}
			return CastAcquireActionRates(core.CastArray(data["acquireActionRates"]))
		}(),
	}
}

func (p UpdateGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"defaultGrades": CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		),
		"experienceModelId": p.ExperienceModelId,
		"gradeEntries": CastGradeEntryModelsFromDict(
			p.GradeEntries,
		),
		"acquireActionRates": CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		),
	}
}

func (p UpdateGradeModelMasterRequest) Pointer() *UpdateGradeModelMasterRequest {
	return &p
}

type DeleteGradeModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteGradeModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGradeModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGradeModelMasterRequest{}
	} else {
		*p = DeleteGradeModelMasterRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
				}
			}
		}
	}
	return nil
}

func NewDeleteGradeModelMasterRequestFromJson(data string) (DeleteGradeModelMasterRequest, error) {
	req := DeleteGradeModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGradeModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteGradeModelMasterRequestFromDict(data map[string]interface{}) DeleteGradeModelMasterRequest {
	return DeleteGradeModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
	}
}

func (p DeleteGradeModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p DeleteGradeModelMasterRequest) Pointer() *DeleteGradeModelMasterRequest {
	return &p
}

type DescribeGradeModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGradeModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGradeModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGradeModelsRequest{}
	} else {
		*p = DescribeGradeModelsRequest{}
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

func NewDescribeGradeModelsRequestFromJson(data string) (DescribeGradeModelsRequest, error) {
	req := DescribeGradeModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGradeModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeGradeModelsRequestFromDict(data map[string]interface{}) DescribeGradeModelsRequest {
	return DescribeGradeModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeGradeModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGradeModelsRequest) Pointer() *DescribeGradeModelsRequest {
	return &p
}

type GetGradeModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGradeModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGradeModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGradeModelRequest{}
	} else {
		*p = GetGradeModelRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
				}
			}
		}
	}
	return nil
}

func NewGetGradeModelRequestFromJson(data string) (GetGradeModelRequest, error) {
	req := GetGradeModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGradeModelRequest{}, err
	}
	return req, nil
}

func NewGetGradeModelRequestFromDict(data map[string]interface{}) GetGradeModelRequest {
	return GetGradeModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
	}
}

func (p GetGradeModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
	}
}

func (p GetGradeModelRequest) Pointer() *GetGradeModelRequest {
	return &p
}

type DescribeStatusesRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	GradeName     *string `json:"gradeName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"namespaceName": p.NamespaceName,
		"gradeName":     p.GradeName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStatusesRequest) Pointer() *DescribeStatusesRequest {
	return &p
}

type DescribeStatusesByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GradeName       *string `json:"gradeName"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"gradeName":       p.GradeName,
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
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	GradeName     *string `json:"gradeName"`
	PropertyId    *string `json:"propertyId"`
	DryRun        *bool   `json:"dryRun"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p GetStatusRequest) Pointer() *GetStatusRequest {
	return &p
}

type GetStatusByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GradeName       *string `json:"gradeName"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetStatusByUserIdRequest) Pointer() *GetStatusByUserIdRequest {
	return &p
}

type AddGradeByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AddGradeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddGradeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddGradeByUserIdRequest{}
	} else {
		*p = AddGradeByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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

func NewAddGradeByUserIdRequestFromJson(data string) (AddGradeByUserIdRequest, error) {
	req := AddGradeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddGradeByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddGradeByUserIdRequestFromDict(data map[string]interface{}) AddGradeByUserIdRequest {
	return AddGradeByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
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

func (p AddGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"gradeValue":      p.GradeValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddGradeByUserIdRequest) Pointer() *AddGradeByUserIdRequest {
	return &p
}

type SubGradeRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubGradeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubGradeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubGradeRequest{}
	} else {
		*p = SubGradeRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["gradeValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeValue)
		}
	}
	return nil
}

func NewSubGradeRequestFromJson(data string) (SubGradeRequest, error) {
	req := SubGradeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubGradeRequest{}, err
	}
	return req, nil
}

func NewSubGradeRequestFromDict(data map[string]interface{}) SubGradeRequest {
	return SubGradeRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
		}(),
	}
}

func (p SubGradeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"gradeValue":    p.GradeValue,
	}
}

func (p SubGradeRequest) Pointer() *SubGradeRequest {
	return &p
}

type SubGradeByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SubGradeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubGradeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubGradeByUserIdRequest{}
	} else {
		*p = SubGradeByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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

func NewSubGradeByUserIdRequestFromJson(data string) (SubGradeByUserIdRequest, error) {
	req := SubGradeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubGradeByUserIdRequest{}, err
	}
	return req, nil
}

func NewSubGradeByUserIdRequestFromDict(data map[string]interface{}) SubGradeByUserIdRequest {
	return SubGradeByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
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

func (p SubGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"gradeValue":      p.GradeValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SubGradeByUserIdRequest) Pointer() *SubGradeByUserIdRequest {
	return &p
}

type SetGradeByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	GradeValue         *int64  `json:"gradeValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *SetGradeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetGradeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetGradeByUserIdRequest{}
	} else {
		*p = SetGradeByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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

func NewSetGradeByUserIdRequestFromJson(data string) (SetGradeByUserIdRequest, error) {
	req := SetGradeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetGradeByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetGradeByUserIdRequestFromDict(data map[string]interface{}) SetGradeByUserIdRequest {
	return SetGradeByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
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

func (p SetGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"gradeValue":      p.GradeValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetGradeByUserIdRequest) Pointer() *SetGradeByUserIdRequest {
	return &p
}

type ApplyRankCapRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *ApplyRankCapRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ApplyRankCapRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ApplyRankCapRequest{}
	} else {
		*p = ApplyRankCapRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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

func NewApplyRankCapRequestFromJson(data string) (ApplyRankCapRequest, error) {
	req := ApplyRankCapRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ApplyRankCapRequest{}, err
	}
	return req, nil
}

func NewApplyRankCapRequestFromDict(data map[string]interface{}) ApplyRankCapRequest {
	return ApplyRankCapRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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

func (p ApplyRankCapRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
	}
}

func (p ApplyRankCapRequest) Pointer() *ApplyRankCapRequest {
	return &p
}

type ApplyRankCapByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	PropertyId         *string `json:"propertyId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *ApplyRankCapByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ApplyRankCapByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ApplyRankCapByUserIdRequest{}
	} else {
		*p = ApplyRankCapByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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

func NewApplyRankCapByUserIdRequestFromJson(data string) (ApplyRankCapByUserIdRequest, error) {
	req := ApplyRankCapByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ApplyRankCapByUserIdRequest{}, err
	}
	return req, nil
}

func NewApplyRankCapByUserIdRequestFromDict(data map[string]interface{}) ApplyRankCapByUserIdRequest {
	return ApplyRankCapByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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

func (p ApplyRankCapByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ApplyRankCapByUserIdRequest) Pointer() *ApplyRankCapByUserIdRequest {
	return &p
}

type DeleteStatusByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"gradeName":       p.GradeName,
		"propertyId":      p.PropertyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteStatusByUserIdRequest) Pointer() *DeleteStatusByUserIdRequest {
	return &p
}

type VerifyGradeRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	GradeName                       *string `json:"gradeName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	GradeValue                      *int64  `json:"gradeValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyGradeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeRequest{}
	} else {
		*p = VerifyGradeRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["gradeValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeValue)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyGradeRequestFromJson(data string) (VerifyGradeRequest, error) {
	req := VerifyGradeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeRequestFromDict(data map[string]interface{}) VerifyGradeRequest {
	return VerifyGradeRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
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

func (p VerifyGradeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"gradeName":                       p.GradeName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"gradeValue":                      p.GradeValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyGradeRequest) Pointer() *VerifyGradeRequest {
	return &p
}

type VerifyGradeByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	GradeName                       *string `json:"gradeName"`
	VerifyType                      *string `json:"verifyType"`
	PropertyId                      *string `json:"propertyId"`
	GradeValue                      *int64  `json:"gradeValue"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyGradeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeByUserIdRequest{}
	} else {
		*p = VerifyGradeByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["gradeValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GradeValue)
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

func NewVerifyGradeByUserIdRequestFromJson(data string) (VerifyGradeByUserIdRequest, error) {
	req := VerifyGradeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeByUserIdRequestFromDict(data map[string]interface{}) VerifyGradeByUserIdRequest {
	return VerifyGradeByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		GradeValue: func() *int64 {
			v, ok := data["gradeValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["gradeValue"])
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

func (p VerifyGradeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"gradeName":                       p.GradeName,
		"verifyType":                      p.VerifyType,
		"propertyId":                      p.PropertyId,
		"gradeValue":                      p.GradeValue,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyGradeByUserIdRequest) Pointer() *VerifyGradeByUserIdRequest {
	return &p
}

type VerifyGradeUpMaterialRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	MaterialPropertyId *string `json:"materialPropertyId"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *VerifyGradeUpMaterialRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeUpMaterialRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeUpMaterialRequest{}
	} else {
		*p = VerifyGradeUpMaterialRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["materialPropertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MaterialPropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MaterialPropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MaterialPropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MaterialPropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MaterialPropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MaterialPropertyId)
				}
			}
		}
	}
	return nil
}

func NewVerifyGradeUpMaterialRequestFromJson(data string) (VerifyGradeUpMaterialRequest, error) {
	req := VerifyGradeUpMaterialRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeUpMaterialRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeUpMaterialRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialRequest {
	return VerifyGradeUpMaterialRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		MaterialPropertyId: func() *string {
			v, ok := data["materialPropertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["materialPropertyId"])
		}(),
	}
}

func (p VerifyGradeUpMaterialRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"accessToken":        p.AccessToken,
		"gradeName":          p.GradeName,
		"verifyType":         p.VerifyType,
		"propertyId":         p.PropertyId,
		"materialPropertyId": p.MaterialPropertyId,
	}
}

func (p VerifyGradeUpMaterialRequest) Pointer() *VerifyGradeUpMaterialRequest {
	return &p
}

type VerifyGradeUpMaterialByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GradeName          *string `json:"gradeName"`
	VerifyType         *string `json:"verifyType"`
	PropertyId         *string `json:"propertyId"`
	MaterialPropertyId *string `json:"materialPropertyId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *VerifyGradeUpMaterialByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeUpMaterialByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeUpMaterialByUserIdRequest{}
	} else {
		*p = VerifyGradeUpMaterialByUserIdRequest{}
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		if v, ok := d["materialPropertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MaterialPropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MaterialPropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MaterialPropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MaterialPropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MaterialPropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MaterialPropertyId)
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

func NewVerifyGradeUpMaterialByUserIdRequestFromJson(data string) (VerifyGradeUpMaterialByUserIdRequest, error) {
	req := VerifyGradeUpMaterialByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeUpMaterialByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeUpMaterialByUserIdRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialByUserIdRequest {
	return VerifyGradeUpMaterialByUserIdRequest{
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		MaterialPropertyId: func() *string {
			v, ok := data["materialPropertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["materialPropertyId"])
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

func (p VerifyGradeUpMaterialByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"userId":             p.UserId,
		"gradeName":          p.GradeName,
		"verifyType":         p.VerifyType,
		"propertyId":         p.PropertyId,
		"materialPropertyId": p.MaterialPropertyId,
		"timeOffsetToken":    p.TimeOffsetToken,
	}
}

func (p VerifyGradeUpMaterialByUserIdRequest) Pointer() *VerifyGradeUpMaterialByUserIdRequest {
	return &p
}

type AddGradeByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *AddGradeByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddGradeByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddGradeByStampSheetRequest{}
	} else {
		*p = AddGradeByStampSheetRequest{}
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

func NewAddGradeByStampSheetRequestFromJson(data string) (AddGradeByStampSheetRequest, error) {
	req := AddGradeByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddGradeByStampSheetRequest{}, err
	}
	return req, nil
}

func NewAddGradeByStampSheetRequestFromDict(data map[string]interface{}) AddGradeByStampSheetRequest {
	return AddGradeByStampSheetRequest{
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

func (p AddGradeByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p AddGradeByStampSheetRequest) Pointer() *AddGradeByStampSheetRequest {
	return &p
}

type ApplyRankCapByStampSheetRequest struct {
	ContextStack *string `json:"contextStack"`
	StampSheet   *string `json:"stampSheet"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *ApplyRankCapByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ApplyRankCapByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ApplyRankCapByStampSheetRequest{}
	} else {
		*p = ApplyRankCapByStampSheetRequest{}
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

func NewApplyRankCapByStampSheetRequestFromJson(data string) (ApplyRankCapByStampSheetRequest, error) {
	req := ApplyRankCapByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ApplyRankCapByStampSheetRequest{}, err
	}
	return req, nil
}

func NewApplyRankCapByStampSheetRequestFromDict(data map[string]interface{}) ApplyRankCapByStampSheetRequest {
	return ApplyRankCapByStampSheetRequest{
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

func (p ApplyRankCapByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p ApplyRankCapByStampSheetRequest) Pointer() *ApplyRankCapByStampSheetRequest {
	return &p
}

type SubGradeByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *SubGradeByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubGradeByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SubGradeByStampTaskRequest{}
	} else {
		*p = SubGradeByStampTaskRequest{}
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

func NewSubGradeByStampTaskRequestFromJson(data string) (SubGradeByStampTaskRequest, error) {
	req := SubGradeByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SubGradeByStampTaskRequest{}, err
	}
	return req, nil
}

func NewSubGradeByStampTaskRequestFromDict(data map[string]interface{}) SubGradeByStampTaskRequest {
	return SubGradeByStampTaskRequest{
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

func (p SubGradeByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p SubGradeByStampTaskRequest) Pointer() *SubGradeByStampTaskRequest {
	return &p
}

type MultiplyAcquireActionsByUserIdRequest struct {
	ContextStack       *string         `json:"contextStack"`
	DuplicationAvoider *string         `json:"duplicationAvoider"`
	NamespaceName      *string         `json:"namespaceName"`
	UserId             *string         `json:"userId"`
	GradeName          *string         `json:"gradeName"`
	PropertyId         *string         `json:"propertyId"`
	RateName           *string         `json:"rateName"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
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
		if v, ok := d["gradeName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GradeName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GradeName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GradeName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GradeName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GradeName)
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
		GradeName: func() *string {
			v, ok := data["gradeName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gradeName"])
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
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"gradeName":     p.GradeName,
		"propertyId":    p.PropertyId,
		"rateName":      p.RateName,
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
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

type VerifyGradeByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyGradeByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeByStampTaskRequest{}
	} else {
		*p = VerifyGradeByStampTaskRequest{}
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

func NewVerifyGradeByStampTaskRequestFromJson(data string) (VerifyGradeByStampTaskRequest, error) {
	req := VerifyGradeByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeByStampTaskRequestFromDict(data map[string]interface{}) VerifyGradeByStampTaskRequest {
	return VerifyGradeByStampTaskRequest{
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

func (p VerifyGradeByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyGradeByStampTaskRequest) Pointer() *VerifyGradeByStampTaskRequest {
	return &p
}

type VerifyGradeUpMaterialByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyGradeUpMaterialByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGradeUpMaterialByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGradeUpMaterialByStampTaskRequest{}
	} else {
		*p = VerifyGradeUpMaterialByStampTaskRequest{}
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

func NewVerifyGradeUpMaterialByStampTaskRequestFromJson(data string) (VerifyGradeUpMaterialByStampTaskRequest, error) {
	req := VerifyGradeUpMaterialByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGradeUpMaterialByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyGradeUpMaterialByStampTaskRequestFromDict(data map[string]interface{}) VerifyGradeUpMaterialByStampTaskRequest {
	return VerifyGradeUpMaterialByStampTaskRequest{
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

func (p VerifyGradeUpMaterialByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyGradeUpMaterialByStampTaskRequest) Pointer() *VerifyGradeUpMaterialByStampTaskRequest {
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

type GetCurrentGradeMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetCurrentGradeMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentGradeMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentGradeMasterRequest{}
	} else {
		*p = GetCurrentGradeMasterRequest{}
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

func NewGetCurrentGradeMasterRequestFromJson(data string) (GetCurrentGradeMasterRequest, error) {
	req := GetCurrentGradeMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentGradeMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentGradeMasterRequestFromDict(data map[string]interface{}) GetCurrentGradeMasterRequest {
	return GetCurrentGradeMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentGradeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentGradeMasterRequest) Pointer() *GetCurrentGradeMasterRequest {
	return &p
}

type PreUpdateCurrentGradeMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *PreUpdateCurrentGradeMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PreUpdateCurrentGradeMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PreUpdateCurrentGradeMasterRequest{}
	} else {
		*p = PreUpdateCurrentGradeMasterRequest{}
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

func NewPreUpdateCurrentGradeMasterRequestFromJson(data string) (PreUpdateCurrentGradeMasterRequest, error) {
	req := PreUpdateCurrentGradeMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PreUpdateCurrentGradeMasterRequest{}, err
	}
	return req, nil
}

func NewPreUpdateCurrentGradeMasterRequestFromDict(data map[string]interface{}) PreUpdateCurrentGradeMasterRequest {
	return PreUpdateCurrentGradeMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p PreUpdateCurrentGradeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p PreUpdateCurrentGradeMasterRequest) Pointer() *PreUpdateCurrentGradeMasterRequest {
	return &p
}

type UpdateCurrentGradeMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Mode          *string `json:"mode"`
	Settings      *string `json:"settings"`
	UploadToken   *string `json:"uploadToken"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *UpdateCurrentGradeMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentGradeMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentGradeMasterRequest{}
	} else {
		*p = UpdateCurrentGradeMasterRequest{}
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

func NewUpdateCurrentGradeMasterRequestFromJson(data string) (UpdateCurrentGradeMasterRequest, error) {
	req := UpdateCurrentGradeMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentGradeMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentGradeMasterRequestFromDict(data map[string]interface{}) UpdateCurrentGradeMasterRequest {
	return UpdateCurrentGradeMasterRequest{
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

func (p UpdateCurrentGradeMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"mode":          p.Mode,
		"settings":      p.Settings,
		"uploadToken":   p.UploadToken,
	}
}

func (p UpdateCurrentGradeMasterRequest) Pointer() *UpdateCurrentGradeMasterRequest {
	return &p
}

type UpdateCurrentGradeMasterFromGitHubRequest struct {
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
	DryRun          *bool                  `json:"dryRun"`
}

func (p *UpdateCurrentGradeMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentGradeMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentGradeMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentGradeMasterFromGitHubRequest{}
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

func NewUpdateCurrentGradeMasterFromGitHubRequestFromJson(data string) (UpdateCurrentGradeMasterFromGitHubRequest, error) {
	req := UpdateCurrentGradeMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentGradeMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentGradeMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentGradeMasterFromGitHubRequest {
	return UpdateCurrentGradeMasterFromGitHubRequest{
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

func (p UpdateCurrentGradeMasterFromGitHubRequest) ToDict() map[string]interface{} {
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

func (p UpdateCurrentGradeMasterFromGitHubRequest) Pointer() *UpdateCurrentGradeMasterFromGitHubRequest {
	return &p
}
