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

package ranking2

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

type DescribeGlobalRankingModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingModelsRequest{}
	} else {
		*p = DescribeGlobalRankingModelsRequest{}
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

func NewDescribeGlobalRankingModelsRequestFromJson(data string) (DescribeGlobalRankingModelsRequest, error) {
	req := DescribeGlobalRankingModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingModelsRequestFromDict(data map[string]interface{}) DescribeGlobalRankingModelsRequest {
	return DescribeGlobalRankingModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeGlobalRankingModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGlobalRankingModelsRequest) Pointer() *DescribeGlobalRankingModelsRequest {
	return &p
}

type GetGlobalRankingModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingModelRequest{}
	} else {
		*p = GetGlobalRankingModelRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetGlobalRankingModelRequestFromJson(data string) (GetGlobalRankingModelRequest, error) {
	req := GetGlobalRankingModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingModelRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingModelRequestFromDict(data map[string]interface{}) GetGlobalRankingModelRequest {
	return GetGlobalRankingModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetGlobalRankingModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetGlobalRankingModelRequest) Pointer() *GetGlobalRankingModelRequest {
	return &p
}

type DescribeGlobalRankingModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingModelMastersRequest{}
	} else {
		*p = DescribeGlobalRankingModelMastersRequest{}
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

func NewDescribeGlobalRankingModelMastersRequestFromJson(data string) (DescribeGlobalRankingModelMastersRequest, error) {
	req := DescribeGlobalRankingModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingModelMastersRequestFromDict(data map[string]interface{}) DescribeGlobalRankingModelMastersRequest {
	return DescribeGlobalRankingModelMastersRequest{
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

func (p DescribeGlobalRankingModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGlobalRankingModelMastersRequest) Pointer() *DescribeGlobalRankingModelMastersRequest {
	return &p
}

type CreateGlobalRankingModelMasterRequest struct {
	ContextStack        *string         `json:"contextStack"`
	NamespaceName       *string         `json:"namespaceName"`
	Name                *string         `json:"name"`
	Description         *string         `json:"description"`
	Metadata            *string         `json:"metadata"`
	MinimumValue        *int64          `json:"minimumValue"`
	MaximumValue        *int64          `json:"maximumValue"`
	Sum                 *bool           `json:"sum"`
	OrderDirection      *string         `json:"orderDirection"`
	RankingRewards      []RankingReward `json:"rankingRewards"`
	EntryPeriodEventId  *string         `json:"entryPeriodEventId"`
	AccessPeriodEventId *string         `json:"accessPeriodEventId"`
	DryRun              *bool           `json:"dryRun"`
}

func (p *CreateGlobalRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGlobalRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGlobalRankingModelMasterRequest{}
	} else {
		*p = CreateGlobalRankingModelMasterRequest{}
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewCreateGlobalRankingModelMasterRequestFromJson(data string) (CreateGlobalRankingModelMasterRequest, error) {
	req := CreateGlobalRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGlobalRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateGlobalRankingModelMasterRequestFromDict(data map[string]interface{}) CreateGlobalRankingModelMasterRequest {
	return CreateGlobalRankingModelMasterRequest{
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p CreateGlobalRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"name":           p.Name,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"minimumValue":   p.MinimumValue,
		"maximumValue":   p.MaximumValue,
		"sum":            p.Sum,
		"orderDirection": p.OrderDirection,
		"rankingRewards": CastRankingRewardsFromDict(
			p.RankingRewards,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p CreateGlobalRankingModelMasterRequest) Pointer() *CreateGlobalRankingModelMasterRequest {
	return &p
}

type GetGlobalRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingModelMasterRequest{}
	} else {
		*p = GetGlobalRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetGlobalRankingModelMasterRequestFromJson(data string) (GetGlobalRankingModelMasterRequest, error) {
	req := GetGlobalRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingModelMasterRequestFromDict(data map[string]interface{}) GetGlobalRankingModelMasterRequest {
	return GetGlobalRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetGlobalRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetGlobalRankingModelMasterRequest) Pointer() *GetGlobalRankingModelMasterRequest {
	return &p
}

type UpdateGlobalRankingModelMasterRequest struct {
	ContextStack        *string         `json:"contextStack"`
	NamespaceName       *string         `json:"namespaceName"`
	RankingName         *string         `json:"rankingName"`
	Description         *string         `json:"description"`
	Metadata            *string         `json:"metadata"`
	MinimumValue        *int64          `json:"minimumValue"`
	MaximumValue        *int64          `json:"maximumValue"`
	Sum                 *bool           `json:"sum"`
	OrderDirection      *string         `json:"orderDirection"`
	RankingRewards      []RankingReward `json:"rankingRewards"`
	EntryPeriodEventId  *string         `json:"entryPeriodEventId"`
	AccessPeriodEventId *string         `json:"accessPeriodEventId"`
	DryRun              *bool           `json:"dryRun"`
}

func (p *UpdateGlobalRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGlobalRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGlobalRankingModelMasterRequest{}
	} else {
		*p = UpdateGlobalRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewUpdateGlobalRankingModelMasterRequestFromJson(data string) (UpdateGlobalRankingModelMasterRequest, error) {
	req := UpdateGlobalRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGlobalRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateGlobalRankingModelMasterRequestFromDict(data map[string]interface{}) UpdateGlobalRankingModelMasterRequest {
	return UpdateGlobalRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p UpdateGlobalRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"rankingName":    p.RankingName,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"minimumValue":   p.MinimumValue,
		"maximumValue":   p.MaximumValue,
		"sum":            p.Sum,
		"orderDirection": p.OrderDirection,
		"rankingRewards": CastRankingRewardsFromDict(
			p.RankingRewards,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p UpdateGlobalRankingModelMasterRequest) Pointer() *UpdateGlobalRankingModelMasterRequest {
	return &p
}

type DeleteGlobalRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteGlobalRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGlobalRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGlobalRankingModelMasterRequest{}
	} else {
		*p = DeleteGlobalRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewDeleteGlobalRankingModelMasterRequestFromJson(data string) (DeleteGlobalRankingModelMasterRequest, error) {
	req := DeleteGlobalRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGlobalRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteGlobalRankingModelMasterRequestFromDict(data map[string]interface{}) DeleteGlobalRankingModelMasterRequest {
	return DeleteGlobalRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p DeleteGlobalRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p DeleteGlobalRankingModelMasterRequest) Pointer() *DeleteGlobalRankingModelMasterRequest {
	return &p
}

type DescribeGlobalRankingScoresRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingScoresRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingScoresRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingScoresRequest{}
	} else {
		*p = DescribeGlobalRankingScoresRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeGlobalRankingScoresRequestFromJson(data string) (DescribeGlobalRankingScoresRequest, error) {
	req := DescribeGlobalRankingScoresRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingScoresRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingScoresRequestFromDict(data map[string]interface{}) DescribeGlobalRankingScoresRequest {
	return DescribeGlobalRankingScoresRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeGlobalRankingScoresRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGlobalRankingScoresRequest) Pointer() *DescribeGlobalRankingScoresRequest {
	return &p
}

type DescribeGlobalRankingScoresByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingScoresByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingScoresByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingScoresByUserIdRequest{}
	} else {
		*p = DescribeGlobalRankingScoresByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeGlobalRankingScoresByUserIdRequestFromJson(data string) (DescribeGlobalRankingScoresByUserIdRequest, error) {
	req := DescribeGlobalRankingScoresByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingScoresByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingScoresByUserIdRequestFromDict(data map[string]interface{}) DescribeGlobalRankingScoresByUserIdRequest {
	return DescribeGlobalRankingScoresByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeGlobalRankingScoresByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeGlobalRankingScoresByUserIdRequest) Pointer() *DescribeGlobalRankingScoresByUserIdRequest {
	return &p
}

type PutGlobalRankingScoreRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	AccessToken        *string `json:"accessToken"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutGlobalRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutGlobalRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutGlobalRankingScoreRequest{}
	} else {
		*p = PutGlobalRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutGlobalRankingScoreRequestFromJson(data string) (PutGlobalRankingScoreRequest, error) {
	req := PutGlobalRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutGlobalRankingScoreRequest{}, err
	}
	return req, nil
}

func NewPutGlobalRankingScoreRequestFromDict(data map[string]interface{}) PutGlobalRankingScoreRequest {
	return PutGlobalRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
	}
}

func (p PutGlobalRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"score":         p.Score,
		"metadata":      p.Metadata,
	}
}

func (p PutGlobalRankingScoreRequest) Pointer() *PutGlobalRankingScoreRequest {
	return &p
}

type PutGlobalRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutGlobalRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutGlobalRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutGlobalRankingScoreByUserIdRequest{}
	} else {
		*p = PutGlobalRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutGlobalRankingScoreByUserIdRequestFromJson(data string) (PutGlobalRankingScoreByUserIdRequest, error) {
	req := PutGlobalRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutGlobalRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewPutGlobalRankingScoreByUserIdRequestFromDict(data map[string]interface{}) PutGlobalRankingScoreByUserIdRequest {
	return PutGlobalRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
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

func (p PutGlobalRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"score":           p.Score,
		"metadata":        p.Metadata,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PutGlobalRankingScoreByUserIdRequest) Pointer() *PutGlobalRankingScoreByUserIdRequest {
	return &p
}

type GetGlobalRankingScoreRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingScoreRequest{}
	} else {
		*p = GetGlobalRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetGlobalRankingScoreRequestFromJson(data string) (GetGlobalRankingScoreRequest, error) {
	req := GetGlobalRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingScoreRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingScoreRequestFromDict(data map[string]interface{}) GetGlobalRankingScoreRequest {
	return GetGlobalRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetGlobalRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetGlobalRankingScoreRequest) Pointer() *GetGlobalRankingScoreRequest {
	return &p
}

type GetGlobalRankingScoreByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingScoreByUserIdRequest{}
	} else {
		*p = GetGlobalRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetGlobalRankingScoreByUserIdRequestFromJson(data string) (GetGlobalRankingScoreByUserIdRequest, error) {
	req := GetGlobalRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingScoreByUserIdRequestFromDict(data map[string]interface{}) GetGlobalRankingScoreByUserIdRequest {
	return GetGlobalRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetGlobalRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetGlobalRankingScoreByUserIdRequest) Pointer() *GetGlobalRankingScoreByUserIdRequest {
	return &p
}

type DeleteGlobalRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteGlobalRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGlobalRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGlobalRankingScoreByUserIdRequest{}
	} else {
		*p = DeleteGlobalRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDeleteGlobalRankingScoreByUserIdRequestFromJson(data string) (DeleteGlobalRankingScoreByUserIdRequest, error) {
	req := DeleteGlobalRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGlobalRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteGlobalRankingScoreByUserIdRequestFromDict(data map[string]interface{}) DeleteGlobalRankingScoreByUserIdRequest {
	return DeleteGlobalRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DeleteGlobalRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteGlobalRankingScoreByUserIdRequest) Pointer() *DeleteGlobalRankingScoreByUserIdRequest {
	return &p
}

type VerifyGlobalRankingScoreRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	RankingName                     *string `json:"rankingName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyGlobalRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGlobalRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGlobalRankingScoreRequest{}
	} else {
		*p = VerifyGlobalRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyGlobalRankingScoreRequestFromJson(data string) (VerifyGlobalRankingScoreRequest, error) {
	req := VerifyGlobalRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGlobalRankingScoreRequest{}, err
	}
	return req, nil
}

func NewVerifyGlobalRankingScoreRequestFromDict(data map[string]interface{}) VerifyGlobalRankingScoreRequest {
	return VerifyGlobalRankingScoreRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifyGlobalRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"rankingName":                     p.RankingName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyGlobalRankingScoreRequest) Pointer() *VerifyGlobalRankingScoreRequest {
	return &p
}

type VerifyGlobalRankingScoreByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	RankingName                     *string `json:"rankingName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyGlobalRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGlobalRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGlobalRankingScoreByUserIdRequest{}
	} else {
		*p = VerifyGlobalRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewVerifyGlobalRankingScoreByUserIdRequestFromJson(data string) (VerifyGlobalRankingScoreByUserIdRequest, error) {
	req := VerifyGlobalRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGlobalRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyGlobalRankingScoreByUserIdRequestFromDict(data map[string]interface{}) VerifyGlobalRankingScoreByUserIdRequest {
	return VerifyGlobalRankingScoreByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifyGlobalRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"rankingName":                     p.RankingName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyGlobalRankingScoreByUserIdRequest) Pointer() *VerifyGlobalRankingScoreByUserIdRequest {
	return &p
}

type VerifyGlobalRankingScoreByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyGlobalRankingScoreByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyGlobalRankingScoreByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyGlobalRankingScoreByStampTaskRequest{}
	} else {
		*p = VerifyGlobalRankingScoreByStampTaskRequest{}
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

func NewVerifyGlobalRankingScoreByStampTaskRequestFromJson(data string) (VerifyGlobalRankingScoreByStampTaskRequest, error) {
	req := VerifyGlobalRankingScoreByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyGlobalRankingScoreByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyGlobalRankingScoreByStampTaskRequestFromDict(data map[string]interface{}) VerifyGlobalRankingScoreByStampTaskRequest {
	return VerifyGlobalRankingScoreByStampTaskRequest{
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

func (p VerifyGlobalRankingScoreByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyGlobalRankingScoreByStampTaskRequest) Pointer() *VerifyGlobalRankingScoreByStampTaskRequest {
	return &p
}

type DescribeGlobalRankingReceivedRewardsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingReceivedRewardsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingReceivedRewardsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingReceivedRewardsRequest{}
	} else {
		*p = DescribeGlobalRankingReceivedRewardsRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeGlobalRankingReceivedRewardsRequestFromJson(data string) (DescribeGlobalRankingReceivedRewardsRequest, error) {
	req := DescribeGlobalRankingReceivedRewardsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingReceivedRewardsRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingReceivedRewardsRequestFromDict(data map[string]interface{}) DescribeGlobalRankingReceivedRewardsRequest {
	return DescribeGlobalRankingReceivedRewardsRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeGlobalRankingReceivedRewardsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGlobalRankingReceivedRewardsRequest) Pointer() *DescribeGlobalRankingReceivedRewardsRequest {
	return &p
}

type DescribeGlobalRankingReceivedRewardsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingReceivedRewardsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingReceivedRewardsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingReceivedRewardsByUserIdRequest{}
	} else {
		*p = DescribeGlobalRankingReceivedRewardsByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeGlobalRankingReceivedRewardsByUserIdRequestFromJson(data string) (DescribeGlobalRankingReceivedRewardsByUserIdRequest, error) {
	req := DescribeGlobalRankingReceivedRewardsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingReceivedRewardsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingReceivedRewardsByUserIdRequestFromDict(data map[string]interface{}) DescribeGlobalRankingReceivedRewardsByUserIdRequest {
	return DescribeGlobalRankingReceivedRewardsByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeGlobalRankingReceivedRewardsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeGlobalRankingReceivedRewardsByUserIdRequest) Pointer() *DescribeGlobalRankingReceivedRewardsByUserIdRequest {
	return &p
}

type CreateGlobalRankingReceivedRewardRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	AccessToken        *string `json:"accessToken"`
	Season             *int64  `json:"season"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *CreateGlobalRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGlobalRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGlobalRankingReceivedRewardRequest{}
	} else {
		*p = CreateGlobalRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewCreateGlobalRankingReceivedRewardRequestFromJson(data string) (CreateGlobalRankingReceivedRewardRequest, error) {
	req := CreateGlobalRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGlobalRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewCreateGlobalRankingReceivedRewardRequestFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardRequest {
	return CreateGlobalRankingReceivedRewardRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p CreateGlobalRankingReceivedRewardRequest) Pointer() *CreateGlobalRankingReceivedRewardRequest {
	return &p
}

type CreateGlobalRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *CreateGlobalRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGlobalRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGlobalRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = CreateGlobalRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewCreateGlobalRankingReceivedRewardByUserIdRequestFromJson(data string) (CreateGlobalRankingReceivedRewardByUserIdRequest, error) {
	req := CreateGlobalRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGlobalRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewCreateGlobalRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardByUserIdRequest {
	return CreateGlobalRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p CreateGlobalRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CreateGlobalRankingReceivedRewardByUserIdRequest) Pointer() *CreateGlobalRankingReceivedRewardByUserIdRequest {
	return &p
}

type ReceiveGlobalRankingReceivedRewardRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	RankingName        *string  `json:"rankingName"`
	Season             *int64   `json:"season"`
	Config             []Config `json:"config"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *ReceiveGlobalRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveGlobalRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveGlobalRankingReceivedRewardRequest{}
	} else {
		*p = ReceiveGlobalRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["config"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Config)
		}
	}
	return nil
}

func NewReceiveGlobalRankingReceivedRewardRequestFromJson(data string) (ReceiveGlobalRankingReceivedRewardRequest, error) {
	req := ReceiveGlobalRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ReceiveGlobalRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewReceiveGlobalRankingReceivedRewardRequestFromDict(data map[string]interface{}) ReceiveGlobalRankingReceivedRewardRequest {
	return ReceiveGlobalRankingReceivedRewardRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Config: func() []Config {
			if data["config"] == nil {
				return nil
			}
			return CastConfigs(core.CastArray(data["config"]))
		}(),
	}
}

func (p ReceiveGlobalRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"season":        p.Season,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReceiveGlobalRankingReceivedRewardRequest) Pointer() *ReceiveGlobalRankingReceivedRewardRequest {
	return &p
}

type ReceiveGlobalRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	RankingName        *string  `json:"rankingName"`
	Season             *int64   `json:"season"`
	Config             []Config `json:"config"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *ReceiveGlobalRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveGlobalRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveGlobalRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = ReceiveGlobalRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["config"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Config)
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

func NewReceiveGlobalRankingReceivedRewardByUserIdRequestFromJson(data string) (ReceiveGlobalRankingReceivedRewardByUserIdRequest, error) {
	req := ReceiveGlobalRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ReceiveGlobalRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewReceiveGlobalRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) ReceiveGlobalRankingReceivedRewardByUserIdRequest {
	return ReceiveGlobalRankingReceivedRewardByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Config: func() []Config {
			if data["config"] == nil {
				return nil
			}
			return CastConfigs(core.CastArray(data["config"]))
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

func (p ReceiveGlobalRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"rankingName":   p.RankingName,
		"season":        p.Season,
		"config": CastConfigsFromDict(
			p.Config,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ReceiveGlobalRankingReceivedRewardByUserIdRequest) Pointer() *ReceiveGlobalRankingReceivedRewardByUserIdRequest {
	return &p
}

type GetGlobalRankingReceivedRewardRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingReceivedRewardRequest{}
	} else {
		*p = GetGlobalRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetGlobalRankingReceivedRewardRequestFromJson(data string) (GetGlobalRankingReceivedRewardRequest, error) {
	req := GetGlobalRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingReceivedRewardRequestFromDict(data map[string]interface{}) GetGlobalRankingReceivedRewardRequest {
	return GetGlobalRankingReceivedRewardRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetGlobalRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetGlobalRankingReceivedRewardRequest) Pointer() *GetGlobalRankingReceivedRewardRequest {
	return &p
}

type GetGlobalRankingReceivedRewardByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = GetGlobalRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetGlobalRankingReceivedRewardByUserIdRequestFromJson(data string) (GetGlobalRankingReceivedRewardByUserIdRequest, error) {
	req := GetGlobalRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) GetGlobalRankingReceivedRewardByUserIdRequest {
	return GetGlobalRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetGlobalRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetGlobalRankingReceivedRewardByUserIdRequest) Pointer() *GetGlobalRankingReceivedRewardByUserIdRequest {
	return &p
}

type DeleteGlobalRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteGlobalRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGlobalRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGlobalRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = DeleteGlobalRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDeleteGlobalRankingReceivedRewardByUserIdRequestFromJson(data string) (DeleteGlobalRankingReceivedRewardByUserIdRequest, error) {
	req := DeleteGlobalRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGlobalRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteGlobalRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) DeleteGlobalRankingReceivedRewardByUserIdRequest {
	return DeleteGlobalRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DeleteGlobalRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteGlobalRankingReceivedRewardByUserIdRequest) Pointer() *DeleteGlobalRankingReceivedRewardByUserIdRequest {
	return &p
}

type CreateGlobalRankingReceivedRewardByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *CreateGlobalRankingReceivedRewardByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGlobalRankingReceivedRewardByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGlobalRankingReceivedRewardByStampTaskRequest{}
	} else {
		*p = CreateGlobalRankingReceivedRewardByStampTaskRequest{}
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

func NewCreateGlobalRankingReceivedRewardByStampTaskRequestFromJson(data string) (CreateGlobalRankingReceivedRewardByStampTaskRequest, error) {
	req := CreateGlobalRankingReceivedRewardByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGlobalRankingReceivedRewardByStampTaskRequest{}, err
	}
	return req, nil
}

func NewCreateGlobalRankingReceivedRewardByStampTaskRequestFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardByStampTaskRequest {
	return CreateGlobalRankingReceivedRewardByStampTaskRequest{
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

func (p CreateGlobalRankingReceivedRewardByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p CreateGlobalRankingReceivedRewardByStampTaskRequest) Pointer() *CreateGlobalRankingReceivedRewardByStampTaskRequest {
	return &p
}

type DescribeGlobalRankingsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingsRequest{}
	} else {
		*p = DescribeGlobalRankingsRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeGlobalRankingsRequestFromJson(data string) (DescribeGlobalRankingsRequest, error) {
	req := DescribeGlobalRankingsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingsRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingsRequestFromDict(data map[string]interface{}) DescribeGlobalRankingsRequest {
	return DescribeGlobalRankingsRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeGlobalRankingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGlobalRankingsRequest) Pointer() *DescribeGlobalRankingsRequest {
	return &p
}

type DescribeGlobalRankingsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeGlobalRankingsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGlobalRankingsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGlobalRankingsByUserIdRequest{}
	} else {
		*p = DescribeGlobalRankingsByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeGlobalRankingsByUserIdRequestFromJson(data string) (DescribeGlobalRankingsByUserIdRequest, error) {
	req := DescribeGlobalRankingsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGlobalRankingsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeGlobalRankingsByUserIdRequestFromDict(data map[string]interface{}) DescribeGlobalRankingsByUserIdRequest {
	return DescribeGlobalRankingsByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeGlobalRankingsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeGlobalRankingsByUserIdRequest) Pointer() *DescribeGlobalRankingsByUserIdRequest {
	return &p
}

type GetGlobalRankingRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingRequest{}
	} else {
		*p = GetGlobalRankingRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetGlobalRankingRequestFromJson(data string) (GetGlobalRankingRequest, error) {
	req := GetGlobalRankingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingRequestFromDict(data map[string]interface{}) GetGlobalRankingRequest {
	return GetGlobalRankingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetGlobalRankingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetGlobalRankingRequest) Pointer() *GetGlobalRankingRequest {
	return &p
}

type GetGlobalRankingByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetGlobalRankingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGlobalRankingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGlobalRankingByUserIdRequest{}
	} else {
		*p = GetGlobalRankingByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetGlobalRankingByUserIdRequestFromJson(data string) (GetGlobalRankingByUserIdRequest, error) {
	req := GetGlobalRankingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGlobalRankingByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetGlobalRankingByUserIdRequestFromDict(data map[string]interface{}) GetGlobalRankingByUserIdRequest {
	return GetGlobalRankingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetGlobalRankingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetGlobalRankingByUserIdRequest) Pointer() *GetGlobalRankingByUserIdRequest {
	return &p
}

type DescribeClusterRankingModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingModelsRequest{}
	} else {
		*p = DescribeClusterRankingModelsRequest{}
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

func NewDescribeClusterRankingModelsRequestFromJson(data string) (DescribeClusterRankingModelsRequest, error) {
	req := DescribeClusterRankingModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingModelsRequestFromDict(data map[string]interface{}) DescribeClusterRankingModelsRequest {
	return DescribeClusterRankingModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeClusterRankingModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeClusterRankingModelsRequest) Pointer() *DescribeClusterRankingModelsRequest {
	return &p
}

type GetClusterRankingModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetClusterRankingModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingModelRequest{}
	} else {
		*p = GetClusterRankingModelRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetClusterRankingModelRequestFromJson(data string) (GetClusterRankingModelRequest, error) {
	req := GetClusterRankingModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingModelRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingModelRequestFromDict(data map[string]interface{}) GetClusterRankingModelRequest {
	return GetClusterRankingModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetClusterRankingModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetClusterRankingModelRequest) Pointer() *GetClusterRankingModelRequest {
	return &p
}

type DescribeClusterRankingModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingModelMastersRequest{}
	} else {
		*p = DescribeClusterRankingModelMastersRequest{}
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

func NewDescribeClusterRankingModelMastersRequestFromJson(data string) (DescribeClusterRankingModelMastersRequest, error) {
	req := DescribeClusterRankingModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingModelMastersRequestFromDict(data map[string]interface{}) DescribeClusterRankingModelMastersRequest {
	return DescribeClusterRankingModelMastersRequest{
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

func (p DescribeClusterRankingModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeClusterRankingModelMastersRequest) Pointer() *DescribeClusterRankingModelMastersRequest {
	return &p
}

type CreateClusterRankingModelMasterRequest struct {
	ContextStack        *string         `json:"contextStack"`
	NamespaceName       *string         `json:"namespaceName"`
	Name                *string         `json:"name"`
	Description         *string         `json:"description"`
	Metadata            *string         `json:"metadata"`
	ClusterType         *string         `json:"clusterType"`
	MinimumValue        *int64          `json:"minimumValue"`
	MaximumValue        *int64          `json:"maximumValue"`
	Sum                 *bool           `json:"sum"`
	ScoreTtlDays        *int32          `json:"scoreTtlDays"`
	OrderDirection      *string         `json:"orderDirection"`
	RankingRewards      []RankingReward `json:"rankingRewards"`
	EntryPeriodEventId  *string         `json:"entryPeriodEventId"`
	AccessPeriodEventId *string         `json:"accessPeriodEventId"`
	DryRun              *bool           `json:"dryRun"`
}

func (p *CreateClusterRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateClusterRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateClusterRankingModelMasterRequest{}
	} else {
		*p = CreateClusterRankingModelMasterRequest{}
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
		if v, ok := d["clusterType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterType)
				}
			}
		}
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewCreateClusterRankingModelMasterRequestFromJson(data string) (CreateClusterRankingModelMasterRequest, error) {
	req := CreateClusterRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateClusterRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateClusterRankingModelMasterRequestFromDict(data map[string]interface{}) CreateClusterRankingModelMasterRequest {
	return CreateClusterRankingModelMasterRequest{
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
		ClusterType: func() *string {
			v, ok := data["clusterType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterType"])
		}(),
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p CreateClusterRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"name":           p.Name,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"clusterType":    p.ClusterType,
		"minimumValue":   p.MinimumValue,
		"maximumValue":   p.MaximumValue,
		"sum":            p.Sum,
		"scoreTtlDays":   p.ScoreTtlDays,
		"orderDirection": p.OrderDirection,
		"rankingRewards": CastRankingRewardsFromDict(
			p.RankingRewards,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p CreateClusterRankingModelMasterRequest) Pointer() *CreateClusterRankingModelMasterRequest {
	return &p
}

type GetClusterRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetClusterRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingModelMasterRequest{}
	} else {
		*p = GetClusterRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetClusterRankingModelMasterRequestFromJson(data string) (GetClusterRankingModelMasterRequest, error) {
	req := GetClusterRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingModelMasterRequestFromDict(data map[string]interface{}) GetClusterRankingModelMasterRequest {
	return GetClusterRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetClusterRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetClusterRankingModelMasterRequest) Pointer() *GetClusterRankingModelMasterRequest {
	return &p
}

type UpdateClusterRankingModelMasterRequest struct {
	ContextStack        *string         `json:"contextStack"`
	NamespaceName       *string         `json:"namespaceName"`
	RankingName         *string         `json:"rankingName"`
	Description         *string         `json:"description"`
	Metadata            *string         `json:"metadata"`
	ClusterType         *string         `json:"clusterType"`
	MinimumValue        *int64          `json:"minimumValue"`
	MaximumValue        *int64          `json:"maximumValue"`
	Sum                 *bool           `json:"sum"`
	ScoreTtlDays        *int32          `json:"scoreTtlDays"`
	OrderDirection      *string         `json:"orderDirection"`
	RankingRewards      []RankingReward `json:"rankingRewards"`
	EntryPeriodEventId  *string         `json:"entryPeriodEventId"`
	AccessPeriodEventId *string         `json:"accessPeriodEventId"`
	DryRun              *bool           `json:"dryRun"`
}

func (p *UpdateClusterRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateClusterRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateClusterRankingModelMasterRequest{}
	} else {
		*p = UpdateClusterRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["clusterType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterType)
				}
			}
		}
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewUpdateClusterRankingModelMasterRequestFromJson(data string) (UpdateClusterRankingModelMasterRequest, error) {
	req := UpdateClusterRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateClusterRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateClusterRankingModelMasterRequestFromDict(data map[string]interface{}) UpdateClusterRankingModelMasterRequest {
	return UpdateClusterRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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
		ClusterType: func() *string {
			v, ok := data["clusterType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterType"])
		}(),
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p UpdateClusterRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"rankingName":    p.RankingName,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"clusterType":    p.ClusterType,
		"minimumValue":   p.MinimumValue,
		"maximumValue":   p.MaximumValue,
		"sum":            p.Sum,
		"scoreTtlDays":   p.ScoreTtlDays,
		"orderDirection": p.OrderDirection,
		"rankingRewards": CastRankingRewardsFromDict(
			p.RankingRewards,
		),
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p UpdateClusterRankingModelMasterRequest) Pointer() *UpdateClusterRankingModelMasterRequest {
	return &p
}

type DeleteClusterRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteClusterRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteClusterRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteClusterRankingModelMasterRequest{}
	} else {
		*p = DeleteClusterRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewDeleteClusterRankingModelMasterRequestFromJson(data string) (DeleteClusterRankingModelMasterRequest, error) {
	req := DeleteClusterRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteClusterRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteClusterRankingModelMasterRequestFromDict(data map[string]interface{}) DeleteClusterRankingModelMasterRequest {
	return DeleteClusterRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p DeleteClusterRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p DeleteClusterRankingModelMasterRequest) Pointer() *DeleteClusterRankingModelMasterRequest {
	return &p
}

type DescribeClusterRankingScoresRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingScoresRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingScoresRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingScoresRequest{}
	} else {
		*p = DescribeClusterRankingScoresRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingScoresRequestFromJson(data string) (DescribeClusterRankingScoresRequest, error) {
	req := DescribeClusterRankingScoresRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingScoresRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingScoresRequestFromDict(data map[string]interface{}) DescribeClusterRankingScoresRequest {
	return DescribeClusterRankingScoresRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingScoresRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeClusterRankingScoresRequest) Pointer() *DescribeClusterRankingScoresRequest {
	return &p
}

type DescribeClusterRankingScoresByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingScoresByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingScoresByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingScoresByUserIdRequest{}
	} else {
		*p = DescribeClusterRankingScoresByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingScoresByUserIdRequestFromJson(data string) (DescribeClusterRankingScoresByUserIdRequest, error) {
	req := DescribeClusterRankingScoresByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingScoresByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingScoresByUserIdRequestFromDict(data map[string]interface{}) DescribeClusterRankingScoresByUserIdRequest {
	return DescribeClusterRankingScoresByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingScoresByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeClusterRankingScoresByUserIdRequest) Pointer() *DescribeClusterRankingScoresByUserIdRequest {
	return &p
}

type PutClusterRankingScoreRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	AccessToken        *string `json:"accessToken"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutClusterRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutClusterRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutClusterRankingScoreRequest{}
	} else {
		*p = PutClusterRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutClusterRankingScoreRequestFromJson(data string) (PutClusterRankingScoreRequest, error) {
	req := PutClusterRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutClusterRankingScoreRequest{}, err
	}
	return req, nil
}

func NewPutClusterRankingScoreRequestFromDict(data map[string]interface{}) PutClusterRankingScoreRequest {
	return PutClusterRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
	}
}

func (p PutClusterRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"accessToken":   p.AccessToken,
		"score":         p.Score,
		"metadata":      p.Metadata,
	}
}

func (p PutClusterRankingScoreRequest) Pointer() *PutClusterRankingScoreRequest {
	return &p
}

type PutClusterRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	UserId             *string `json:"userId"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutClusterRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutClusterRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutClusterRankingScoreByUserIdRequest{}
	} else {
		*p = PutClusterRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutClusterRankingScoreByUserIdRequestFromJson(data string) (PutClusterRankingScoreByUserIdRequest, error) {
	req := PutClusterRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutClusterRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewPutClusterRankingScoreByUserIdRequestFromDict(data map[string]interface{}) PutClusterRankingScoreByUserIdRequest {
	return PutClusterRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
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

func (p PutClusterRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"score":           p.Score,
		"metadata":        p.Metadata,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PutClusterRankingScoreByUserIdRequest) Pointer() *PutClusterRankingScoreByUserIdRequest {
	return &p
}

type GetClusterRankingScoreRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetClusterRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingScoreRequest{}
	} else {
		*p = GetClusterRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetClusterRankingScoreRequestFromJson(data string) (GetClusterRankingScoreRequest, error) {
	req := GetClusterRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingScoreRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingScoreRequestFromDict(data map[string]interface{}) GetClusterRankingScoreRequest {
	return GetClusterRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetClusterRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetClusterRankingScoreRequest) Pointer() *GetClusterRankingScoreRequest {
	return &p
}

type GetClusterRankingScoreByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetClusterRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingScoreByUserIdRequest{}
	} else {
		*p = GetClusterRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetClusterRankingScoreByUserIdRequestFromJson(data string) (GetClusterRankingScoreByUserIdRequest, error) {
	req := GetClusterRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingScoreByUserIdRequestFromDict(data map[string]interface{}) GetClusterRankingScoreByUserIdRequest {
	return GetClusterRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetClusterRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetClusterRankingScoreByUserIdRequest) Pointer() *GetClusterRankingScoreByUserIdRequest {
	return &p
}

type DeleteClusterRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteClusterRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteClusterRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteClusterRankingScoreByUserIdRequest{}
	} else {
		*p = DeleteClusterRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDeleteClusterRankingScoreByUserIdRequestFromJson(data string) (DeleteClusterRankingScoreByUserIdRequest, error) {
	req := DeleteClusterRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteClusterRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteClusterRankingScoreByUserIdRequestFromDict(data map[string]interface{}) DeleteClusterRankingScoreByUserIdRequest {
	return DeleteClusterRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DeleteClusterRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteClusterRankingScoreByUserIdRequest) Pointer() *DeleteClusterRankingScoreByUserIdRequest {
	return &p
}

type VerifyClusterRankingScoreRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	RankingName                     *string `json:"rankingName"`
	ClusterName                     *string `json:"clusterName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyClusterRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyClusterRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyClusterRankingScoreRequest{}
	} else {
		*p = VerifyClusterRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyClusterRankingScoreRequestFromJson(data string) (VerifyClusterRankingScoreRequest, error) {
	req := VerifyClusterRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyClusterRankingScoreRequest{}, err
	}
	return req, nil
}

func NewVerifyClusterRankingScoreRequestFromDict(data map[string]interface{}) VerifyClusterRankingScoreRequest {
	return VerifyClusterRankingScoreRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifyClusterRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"rankingName":                     p.RankingName,
		"clusterName":                     p.ClusterName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyClusterRankingScoreRequest) Pointer() *VerifyClusterRankingScoreRequest {
	return &p
}

type VerifyClusterRankingScoreByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	RankingName                     *string `json:"rankingName"`
	ClusterName                     *string `json:"clusterName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifyClusterRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyClusterRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyClusterRankingScoreByUserIdRequest{}
	} else {
		*p = VerifyClusterRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewVerifyClusterRankingScoreByUserIdRequestFromJson(data string) (VerifyClusterRankingScoreByUserIdRequest, error) {
	req := VerifyClusterRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyClusterRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyClusterRankingScoreByUserIdRequestFromDict(data map[string]interface{}) VerifyClusterRankingScoreByUserIdRequest {
	return VerifyClusterRankingScoreByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifyClusterRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"rankingName":                     p.RankingName,
		"clusterName":                     p.ClusterName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifyClusterRankingScoreByUserIdRequest) Pointer() *VerifyClusterRankingScoreByUserIdRequest {
	return &p
}

type VerifyClusterRankingScoreByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifyClusterRankingScoreByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyClusterRankingScoreByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyClusterRankingScoreByStampTaskRequest{}
	} else {
		*p = VerifyClusterRankingScoreByStampTaskRequest{}
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

func NewVerifyClusterRankingScoreByStampTaskRequestFromJson(data string) (VerifyClusterRankingScoreByStampTaskRequest, error) {
	req := VerifyClusterRankingScoreByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyClusterRankingScoreByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyClusterRankingScoreByStampTaskRequestFromDict(data map[string]interface{}) VerifyClusterRankingScoreByStampTaskRequest {
	return VerifyClusterRankingScoreByStampTaskRequest{
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

func (p VerifyClusterRankingScoreByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyClusterRankingScoreByStampTaskRequest) Pointer() *VerifyClusterRankingScoreByStampTaskRequest {
	return &p
}

type DescribeClusterRankingReceivedRewardsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingReceivedRewardsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingReceivedRewardsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingReceivedRewardsRequest{}
	} else {
		*p = DescribeClusterRankingReceivedRewardsRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingReceivedRewardsRequestFromJson(data string) (DescribeClusterRankingReceivedRewardsRequest, error) {
	req := DescribeClusterRankingReceivedRewardsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingReceivedRewardsRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingReceivedRewardsRequestFromDict(data map[string]interface{}) DescribeClusterRankingReceivedRewardsRequest {
	return DescribeClusterRankingReceivedRewardsRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingReceivedRewardsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeClusterRankingReceivedRewardsRequest) Pointer() *DescribeClusterRankingReceivedRewardsRequest {
	return &p
}

type DescribeClusterRankingReceivedRewardsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingReceivedRewardsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingReceivedRewardsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingReceivedRewardsByUserIdRequest{}
	} else {
		*p = DescribeClusterRankingReceivedRewardsByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingReceivedRewardsByUserIdRequestFromJson(data string) (DescribeClusterRankingReceivedRewardsByUserIdRequest, error) {
	req := DescribeClusterRankingReceivedRewardsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingReceivedRewardsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingReceivedRewardsByUserIdRequestFromDict(data map[string]interface{}) DescribeClusterRankingReceivedRewardsByUserIdRequest {
	return DescribeClusterRankingReceivedRewardsByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingReceivedRewardsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeClusterRankingReceivedRewardsByUserIdRequest) Pointer() *DescribeClusterRankingReceivedRewardsByUserIdRequest {
	return &p
}

type CreateClusterRankingReceivedRewardRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	AccessToken        *string `json:"accessToken"`
	Season             *int64  `json:"season"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *CreateClusterRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateClusterRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateClusterRankingReceivedRewardRequest{}
	} else {
		*p = CreateClusterRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewCreateClusterRankingReceivedRewardRequestFromJson(data string) (CreateClusterRankingReceivedRewardRequest, error) {
	req := CreateClusterRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateClusterRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewCreateClusterRankingReceivedRewardRequestFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardRequest {
	return CreateClusterRankingReceivedRewardRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p CreateClusterRankingReceivedRewardRequest) Pointer() *CreateClusterRankingReceivedRewardRequest {
	return &p
}

type CreateClusterRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *CreateClusterRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateClusterRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateClusterRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = CreateClusterRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewCreateClusterRankingReceivedRewardByUserIdRequestFromJson(data string) (CreateClusterRankingReceivedRewardByUserIdRequest, error) {
	req := CreateClusterRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateClusterRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewCreateClusterRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardByUserIdRequest {
	return CreateClusterRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p CreateClusterRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CreateClusterRankingReceivedRewardByUserIdRequest) Pointer() *CreateClusterRankingReceivedRewardByUserIdRequest {
	return &p
}

type ReceiveClusterRankingReceivedRewardRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	AccessToken        *string  `json:"accessToken"`
	RankingName        *string  `json:"rankingName"`
	ClusterName        *string  `json:"clusterName"`
	Season             *int64   `json:"season"`
	Config             []Config `json:"config"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *ReceiveClusterRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveClusterRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveClusterRankingReceivedRewardRequest{}
	} else {
		*p = ReceiveClusterRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["config"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Config)
		}
	}
	return nil
}

func NewReceiveClusterRankingReceivedRewardRequestFromJson(data string) (ReceiveClusterRankingReceivedRewardRequest, error) {
	req := ReceiveClusterRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ReceiveClusterRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewReceiveClusterRankingReceivedRewardRequestFromDict(data map[string]interface{}) ReceiveClusterRankingReceivedRewardRequest {
	return ReceiveClusterRankingReceivedRewardRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Config: func() []Config {
			if data["config"] == nil {
				return nil
			}
			return CastConfigs(core.CastArray(data["config"]))
		}(),
	}
}

func (p ReceiveClusterRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"season":        p.Season,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p ReceiveClusterRankingReceivedRewardRequest) Pointer() *ReceiveClusterRankingReceivedRewardRequest {
	return &p
}

type ReceiveClusterRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	UserId             *string  `json:"userId"`
	RankingName        *string  `json:"rankingName"`
	ClusterName        *string  `json:"clusterName"`
	Season             *int64   `json:"season"`
	Config             []Config `json:"config"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
	DryRun             *bool    `json:"dryRun"`
}

func (p *ReceiveClusterRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveClusterRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveClusterRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = ReceiveClusterRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["config"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Config)
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

func NewReceiveClusterRankingReceivedRewardByUserIdRequestFromJson(data string) (ReceiveClusterRankingReceivedRewardByUserIdRequest, error) {
	req := ReceiveClusterRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ReceiveClusterRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewReceiveClusterRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) ReceiveClusterRankingReceivedRewardByUserIdRequest {
	return ReceiveClusterRankingReceivedRewardByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Config: func() []Config {
			if data["config"] == nil {
				return nil
			}
			return CastConfigs(core.CastArray(data["config"]))
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

func (p ReceiveClusterRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"season":        p.Season,
		"config": CastConfigsFromDict(
			p.Config,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ReceiveClusterRankingReceivedRewardByUserIdRequest) Pointer() *ReceiveClusterRankingReceivedRewardByUserIdRequest {
	return &p
}

type GetClusterRankingReceivedRewardRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetClusterRankingReceivedRewardRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingReceivedRewardRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingReceivedRewardRequest{}
	} else {
		*p = GetClusterRankingReceivedRewardRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetClusterRankingReceivedRewardRequestFromJson(data string) (GetClusterRankingReceivedRewardRequest, error) {
	req := GetClusterRankingReceivedRewardRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingReceivedRewardRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingReceivedRewardRequestFromDict(data map[string]interface{}) GetClusterRankingReceivedRewardRequest {
	return GetClusterRankingReceivedRewardRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetClusterRankingReceivedRewardRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetClusterRankingReceivedRewardRequest) Pointer() *GetClusterRankingReceivedRewardRequest {
	return &p
}

type GetClusterRankingReceivedRewardByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetClusterRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = GetClusterRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetClusterRankingReceivedRewardByUserIdRequestFromJson(data string) (GetClusterRankingReceivedRewardByUserIdRequest, error) {
	req := GetClusterRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) GetClusterRankingReceivedRewardByUserIdRequest {
	return GetClusterRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetClusterRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetClusterRankingReceivedRewardByUserIdRequest) Pointer() *GetClusterRankingReceivedRewardByUserIdRequest {
	return &p
}

type DeleteClusterRankingReceivedRewardByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	ClusterName        *string `json:"clusterName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteClusterRankingReceivedRewardByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteClusterRankingReceivedRewardByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteClusterRankingReceivedRewardByUserIdRequest{}
	} else {
		*p = DeleteClusterRankingReceivedRewardByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDeleteClusterRankingReceivedRewardByUserIdRequestFromJson(data string) (DeleteClusterRankingReceivedRewardByUserIdRequest, error) {
	req := DeleteClusterRankingReceivedRewardByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteClusterRankingReceivedRewardByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteClusterRankingReceivedRewardByUserIdRequestFromDict(data map[string]interface{}) DeleteClusterRankingReceivedRewardByUserIdRequest {
	return DeleteClusterRankingReceivedRewardByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DeleteClusterRankingReceivedRewardByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteClusterRankingReceivedRewardByUserIdRequest) Pointer() *DeleteClusterRankingReceivedRewardByUserIdRequest {
	return &p
}

type CreateClusterRankingReceivedRewardByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *CreateClusterRankingReceivedRewardByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateClusterRankingReceivedRewardByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateClusterRankingReceivedRewardByStampTaskRequest{}
	} else {
		*p = CreateClusterRankingReceivedRewardByStampTaskRequest{}
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

func NewCreateClusterRankingReceivedRewardByStampTaskRequestFromJson(data string) (CreateClusterRankingReceivedRewardByStampTaskRequest, error) {
	req := CreateClusterRankingReceivedRewardByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateClusterRankingReceivedRewardByStampTaskRequest{}, err
	}
	return req, nil
}

func NewCreateClusterRankingReceivedRewardByStampTaskRequestFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardByStampTaskRequest {
	return CreateClusterRankingReceivedRewardByStampTaskRequest{
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

func (p CreateClusterRankingReceivedRewardByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p CreateClusterRankingReceivedRewardByStampTaskRequest) Pointer() *CreateClusterRankingReceivedRewardByStampTaskRequest {
	return &p
}

type DescribeClusterRankingsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingsRequest{}
	} else {
		*p = DescribeClusterRankingsRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingsRequestFromJson(data string) (DescribeClusterRankingsRequest, error) {
	req := DescribeClusterRankingsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingsRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingsRequestFromDict(data map[string]interface{}) DescribeClusterRankingsRequest {
	return DescribeClusterRankingsRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeClusterRankingsRequest) Pointer() *DescribeClusterRankingsRequest {
	return &p
}

type DescribeClusterRankingsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeClusterRankingsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeClusterRankingsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeClusterRankingsByUserIdRequest{}
	} else {
		*p = DescribeClusterRankingsByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeClusterRankingsByUserIdRequestFromJson(data string) (DescribeClusterRankingsByUserIdRequest, error) {
	req := DescribeClusterRankingsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeClusterRankingsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeClusterRankingsByUserIdRequestFromDict(data map[string]interface{}) DescribeClusterRankingsByUserIdRequest {
	return DescribeClusterRankingsByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeClusterRankingsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeClusterRankingsByUserIdRequest) Pointer() *DescribeClusterRankingsByUserIdRequest {
	return &p
}

type GetClusterRankingRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	ClusterName   *string `json:"clusterName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetClusterRankingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingRequest{}
	} else {
		*p = GetClusterRankingRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetClusterRankingRequestFromJson(data string) (GetClusterRankingRequest, error) {
	req := GetClusterRankingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingRequestFromDict(data map[string]interface{}) GetClusterRankingRequest {
	return GetClusterRankingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetClusterRankingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"clusterName":   p.ClusterName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetClusterRankingRequest) Pointer() *GetClusterRankingRequest {
	return &p
}

type GetClusterRankingByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	ClusterName     *string `json:"clusterName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetClusterRankingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetClusterRankingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetClusterRankingByUserIdRequest{}
	} else {
		*p = GetClusterRankingByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetClusterRankingByUserIdRequestFromJson(data string) (GetClusterRankingByUserIdRequest, error) {
	req := GetClusterRankingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetClusterRankingByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetClusterRankingByUserIdRequestFromDict(data map[string]interface{}) GetClusterRankingByUserIdRequest {
	return GetClusterRankingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetClusterRankingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"clusterName":     p.ClusterName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetClusterRankingByUserIdRequest) Pointer() *GetClusterRankingByUserIdRequest {
	return &p
}

type DescribeSubscribeRankingModelsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingModelsRequest{}
	} else {
		*p = DescribeSubscribeRankingModelsRequest{}
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

func NewDescribeSubscribeRankingModelsRequestFromJson(data string) (DescribeSubscribeRankingModelsRequest, error) {
	req := DescribeSubscribeRankingModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingModelsRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingModelsRequest {
	return DescribeSubscribeRankingModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeSubscribeRankingModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeSubscribeRankingModelsRequest) Pointer() *DescribeSubscribeRankingModelsRequest {
	return &p
}

type GetSubscribeRankingModelRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingModelRequest{}
	} else {
		*p = GetSubscribeRankingModelRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetSubscribeRankingModelRequestFromJson(data string) (GetSubscribeRankingModelRequest, error) {
	req := GetSubscribeRankingModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingModelRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingModelRequestFromDict(data map[string]interface{}) GetSubscribeRankingModelRequest {
	return GetSubscribeRankingModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetSubscribeRankingModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetSubscribeRankingModelRequest) Pointer() *GetSubscribeRankingModelRequest {
	return &p
}

type DescribeSubscribeRankingModelMastersRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingModelMastersRequest{}
	} else {
		*p = DescribeSubscribeRankingModelMastersRequest{}
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

func NewDescribeSubscribeRankingModelMastersRequestFromJson(data string) (DescribeSubscribeRankingModelMastersRequest, error) {
	req := DescribeSubscribeRankingModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingModelMastersRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingModelMastersRequest {
	return DescribeSubscribeRankingModelMastersRequest{
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

func (p DescribeSubscribeRankingModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribeRankingModelMastersRequest) Pointer() *DescribeSubscribeRankingModelMastersRequest {
	return &p
}

type CreateSubscribeRankingModelMasterRequest struct {
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	MinimumValue        *int64  `json:"minimumValue"`
	MaximumValue        *int64  `json:"maximumValue"`
	Sum                 *bool   `json:"sum"`
	ScoreTtlDays        *int32  `json:"scoreTtlDays"`
	OrderDirection      *string `json:"orderDirection"`
	EntryPeriodEventId  *string `json:"entryPeriodEventId"`
	AccessPeriodEventId *string `json:"accessPeriodEventId"`
	DryRun              *bool   `json:"dryRun"`
}

func (p *CreateSubscribeRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateSubscribeRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateSubscribeRankingModelMasterRequest{}
	} else {
		*p = CreateSubscribeRankingModelMasterRequest{}
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewCreateSubscribeRankingModelMasterRequestFromJson(data string) (CreateSubscribeRankingModelMasterRequest, error) {
	req := CreateSubscribeRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateSubscribeRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateSubscribeRankingModelMasterRequestFromDict(data map[string]interface{}) CreateSubscribeRankingModelMasterRequest {
	return CreateSubscribeRankingModelMasterRequest{
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p CreateSubscribeRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"name":                p.Name,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"minimumValue":        p.MinimumValue,
		"maximumValue":        p.MaximumValue,
		"sum":                 p.Sum,
		"scoreTtlDays":        p.ScoreTtlDays,
		"orderDirection":      p.OrderDirection,
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p CreateSubscribeRankingModelMasterRequest) Pointer() *CreateSubscribeRankingModelMasterRequest {
	return &p
}

type GetSubscribeRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingModelMasterRequest{}
	} else {
		*p = GetSubscribeRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewGetSubscribeRankingModelMasterRequestFromJson(data string) (GetSubscribeRankingModelMasterRequest, error) {
	req := GetSubscribeRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingModelMasterRequestFromDict(data map[string]interface{}) GetSubscribeRankingModelMasterRequest {
	return GetSubscribeRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p GetSubscribeRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p GetSubscribeRankingModelMasterRequest) Pointer() *GetSubscribeRankingModelMasterRequest {
	return &p
}

type UpdateSubscribeRankingModelMasterRequest struct {
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	RankingName         *string `json:"rankingName"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	MinimumValue        *int64  `json:"minimumValue"`
	MaximumValue        *int64  `json:"maximumValue"`
	Sum                 *bool   `json:"sum"`
	ScoreTtlDays        *int32  `json:"scoreTtlDays"`
	OrderDirection      *string `json:"orderDirection"`
	EntryPeriodEventId  *string `json:"entryPeriodEventId"`
	AccessPeriodEventId *string `json:"accessPeriodEventId"`
	DryRun              *bool   `json:"dryRun"`
}

func (p *UpdateSubscribeRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateSubscribeRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateSubscribeRankingModelMasterRequest{}
	} else {
		*p = UpdateSubscribeRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewUpdateSubscribeRankingModelMasterRequestFromJson(data string) (UpdateSubscribeRankingModelMasterRequest, error) {
	req := UpdateSubscribeRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateSubscribeRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateSubscribeRankingModelMasterRequestFromDict(data map[string]interface{}) UpdateSubscribeRankingModelMasterRequest {
	return UpdateSubscribeRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p UpdateSubscribeRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"rankingName":         p.RankingName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"minimumValue":        p.MinimumValue,
		"maximumValue":        p.MaximumValue,
		"sum":                 p.Sum,
		"scoreTtlDays":        p.ScoreTtlDays,
		"orderDirection":      p.OrderDirection,
		"entryPeriodEventId":  p.EntryPeriodEventId,
		"accessPeriodEventId": p.AccessPeriodEventId,
	}
}

func (p UpdateSubscribeRankingModelMasterRequest) Pointer() *UpdateSubscribeRankingModelMasterRequest {
	return &p
}

type DeleteSubscribeRankingModelMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DeleteSubscribeRankingModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSubscribeRankingModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSubscribeRankingModelMasterRequest{}
	} else {
		*p = DeleteSubscribeRankingModelMasterRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
	}
	return nil
}

func NewDeleteSubscribeRankingModelMasterRequestFromJson(data string) (DeleteSubscribeRankingModelMasterRequest, error) {
	req := DeleteSubscribeRankingModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSubscribeRankingModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteSubscribeRankingModelMasterRequestFromDict(data map[string]interface{}) DeleteSubscribeRankingModelMasterRequest {
	return DeleteSubscribeRankingModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
	}
}

func (p DeleteSubscribeRankingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
	}
}

func (p DeleteSubscribeRankingModelMasterRequest) Pointer() *DeleteSubscribeRankingModelMasterRequest {
	return &p
}

type DescribeSubscribesRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscribesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribesRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribesRequest{}
	} else {
		*p = DescribeSubscribesRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeSubscribesRequestFromJson(data string) (DescribeSubscribesRequest, error) {
	req := DescribeSubscribesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribesRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribesRequestFromDict(data map[string]interface{}) DescribeSubscribesRequest {
	return DescribeSubscribesRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeSubscribesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribesRequest) Pointer() *DescribeSubscribesRequest {
	return &p
}

type DescribeSubscribesByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeSubscribesByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribesByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribesByUserIdRequest{}
	} else {
		*p = DescribeSubscribesByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeSubscribesByUserIdRequestFromJson(data string) (DescribeSubscribesByUserIdRequest, error) {
	req := DescribeSubscribesByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribesByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribesByUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribesByUserIdRequest {
	return DescribeSubscribesByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeSubscribesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSubscribesByUserIdRequest) Pointer() *DescribeSubscribesByUserIdRequest {
	return &p
}

type AddSubscribeRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AddSubscribeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddSubscribeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddSubscribeRequest{}
	} else {
		*p = AddSubscribeRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
				}
			}
		}
	}
	return nil
}

func NewAddSubscribeRequestFromJson(data string) (AddSubscribeRequest, error) {
	req := AddSubscribeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddSubscribeRequest{}, err
	}
	return req, nil
}

func NewAddSubscribeRequestFromDict(data map[string]interface{}) AddSubscribeRequest {
	return AddSubscribeRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
		}(),
	}
}

func (p AddSubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p AddSubscribeRequest) Pointer() *AddSubscribeRequest {
	return &p
}

type AddSubscribeByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	TargetUserId       *string `json:"targetUserId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *AddSubscribeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddSubscribeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddSubscribeByUserIdRequest{}
	} else {
		*p = AddSubscribeByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
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

func NewAddSubscribeByUserIdRequestFromJson(data string) (AddSubscribeByUserIdRequest, error) {
	req := AddSubscribeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddSubscribeByUserIdRequest{}, err
	}
	return req, nil
}

func NewAddSubscribeByUserIdRequestFromDict(data map[string]interface{}) AddSubscribeByUserIdRequest {
	return AddSubscribeByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
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

func (p AddSubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddSubscribeByUserIdRequest) Pointer() *AddSubscribeByUserIdRequest {
	return &p
}

type DescribeSubscribeRankingScoresRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingScoresRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingScoresRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingScoresRequest{}
	} else {
		*p = DescribeSubscribeRankingScoresRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeSubscribeRankingScoresRequestFromJson(data string) (DescribeSubscribeRankingScoresRequest, error) {
	req := DescribeSubscribeRankingScoresRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingScoresRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingScoresRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingScoresRequest {
	return DescribeSubscribeRankingScoresRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeSubscribeRankingScoresRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribeRankingScoresRequest) Pointer() *DescribeSubscribeRankingScoresRequest {
	return &p
}

type DescribeSubscribeRankingScoresByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingScoresByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingScoresByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingScoresByUserIdRequest{}
	} else {
		*p = DescribeSubscribeRankingScoresByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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

func NewDescribeSubscribeRankingScoresByUserIdRequestFromJson(data string) (DescribeSubscribeRankingScoresByUserIdRequest, error) {
	req := DescribeSubscribeRankingScoresByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingScoresByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingScoresByUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingScoresByUserIdRequest {
	return DescribeSubscribeRankingScoresByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
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

func (p DescribeSubscribeRankingScoresByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSubscribeRankingScoresByUserIdRequest) Pointer() *DescribeSubscribeRankingScoresByUserIdRequest {
	return &p
}

type PutSubscribeRankingScoreRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	AccessToken        *string `json:"accessToken"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutSubscribeRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutSubscribeRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutSubscribeRankingScoreRequest{}
	} else {
		*p = PutSubscribeRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutSubscribeRankingScoreRequestFromJson(data string) (PutSubscribeRankingScoreRequest, error) {
	req := PutSubscribeRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutSubscribeRankingScoreRequest{}, err
	}
	return req, nil
}

func NewPutSubscribeRankingScoreRequestFromDict(data map[string]interface{}) PutSubscribeRankingScoreRequest {
	return PutSubscribeRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
	}
}

func (p PutSubscribeRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"score":         p.Score,
		"metadata":      p.Metadata,
	}
}

func (p PutSubscribeRankingScoreRequest) Pointer() *PutSubscribeRankingScoreRequest {
	return &p
}

type PutSubscribeRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Score              *int64  `json:"score"`
	Metadata           *string `json:"metadata"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *PutSubscribeRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutSubscribeRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PutSubscribeRankingScoreByUserIdRequest{}
	} else {
		*p = PutSubscribeRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewPutSubscribeRankingScoreByUserIdRequestFromJson(data string) (PutSubscribeRankingScoreByUserIdRequest, error) {
	req := PutSubscribeRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutSubscribeRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewPutSubscribeRankingScoreByUserIdRequestFromDict(data map[string]interface{}) PutSubscribeRankingScoreByUserIdRequest {
	return PutSubscribeRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
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

func (p PutSubscribeRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"score":           p.Score,
		"metadata":        p.Metadata,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PutSubscribeRankingScoreByUserIdRequest) Pointer() *PutSubscribeRankingScoreByUserIdRequest {
	return &p
}

type GetSubscribeRankingScoreRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingScoreRequest{}
	} else {
		*p = GetSubscribeRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
	}
	return nil
}

func NewGetSubscribeRankingScoreRequestFromJson(data string) (GetSubscribeRankingScoreRequest, error) {
	req := GetSubscribeRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingScoreRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingScoreRequestFromDict(data map[string]interface{}) GetSubscribeRankingScoreRequest {
	return GetSubscribeRankingScoreRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
	}
}

func (p GetSubscribeRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
	}
}

func (p GetSubscribeRankingScoreRequest) Pointer() *GetSubscribeRankingScoreRequest {
	return &p
}

type GetSubscribeRankingScoreByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingScoreByUserIdRequest{}
	} else {
		*p = GetSubscribeRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewGetSubscribeRankingScoreByUserIdRequestFromJson(data string) (GetSubscribeRankingScoreByUserIdRequest, error) {
	req := GetSubscribeRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingScoreByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeRankingScoreByUserIdRequest {
	return GetSubscribeRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p GetSubscribeRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSubscribeRankingScoreByUserIdRequest) Pointer() *GetSubscribeRankingScoreByUserIdRequest {
	return &p
}

type DeleteSubscribeRankingScoreByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	Season             *int64  `json:"season"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteSubscribeRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSubscribeRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSubscribeRankingScoreByUserIdRequest{}
	} else {
		*p = DeleteSubscribeRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDeleteSubscribeRankingScoreByUserIdRequestFromJson(data string) (DeleteSubscribeRankingScoreByUserIdRequest, error) {
	req := DeleteSubscribeRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSubscribeRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteSubscribeRankingScoreByUserIdRequestFromDict(data map[string]interface{}) DeleteSubscribeRankingScoreByUserIdRequest {
	return DeleteSubscribeRankingScoreByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DeleteSubscribeRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteSubscribeRankingScoreByUserIdRequest) Pointer() *DeleteSubscribeRankingScoreByUserIdRequest {
	return &p
}

type VerifySubscribeRankingScoreRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	AccessToken                     *string `json:"accessToken"`
	RankingName                     *string `json:"rankingName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifySubscribeRankingScoreRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySubscribeRankingScoreRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifySubscribeRankingScoreRequest{}
	} else {
		*p = VerifySubscribeRankingScoreRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifySubscribeRankingScoreRequestFromJson(data string) (VerifySubscribeRankingScoreRequest, error) {
	req := VerifySubscribeRankingScoreRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySubscribeRankingScoreRequest{}, err
	}
	return req, nil
}

func NewVerifySubscribeRankingScoreRequestFromDict(data map[string]interface{}) VerifySubscribeRankingScoreRequest {
	return VerifySubscribeRankingScoreRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifySubscribeRankingScoreRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"accessToken":                     p.AccessToken,
		"rankingName":                     p.RankingName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifySubscribeRankingScoreRequest) Pointer() *VerifySubscribeRankingScoreRequest {
	return &p
}

type VerifySubscribeRankingScoreByUserIdRequest struct {
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	UserId                          *string `json:"userId"`
	RankingName                     *string `json:"rankingName"`
	VerifyType                      *string `json:"verifyType"`
	Season                          *int64  `json:"season"`
	Score                           *int64  `json:"score"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
	TimeOffsetToken                 *string `json:"timeOffsetToken"`
	DryRun                          *bool   `json:"dryRun"`
}

func (p *VerifySubscribeRankingScoreByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySubscribeRankingScoreByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifySubscribeRankingScoreByUserIdRequest{}
	} else {
		*p = VerifySubscribeRankingScoreByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewVerifySubscribeRankingScoreByUserIdRequestFromJson(data string) (VerifySubscribeRankingScoreByUserIdRequest, error) {
	req := VerifySubscribeRankingScoreByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySubscribeRankingScoreByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifySubscribeRankingScoreByUserIdRequestFromDict(data map[string]interface{}) VerifySubscribeRankingScoreByUserIdRequest {
	return VerifySubscribeRankingScoreByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p VerifySubscribeRankingScoreByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"userId":                          p.UserId,
		"rankingName":                     p.RankingName,
		"verifyType":                      p.VerifyType,
		"season":                          p.Season,
		"score":                           p.Score,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
		"timeOffsetToken":                 p.TimeOffsetToken,
	}
}

func (p VerifySubscribeRankingScoreByUserIdRequest) Pointer() *VerifySubscribeRankingScoreByUserIdRequest {
	return &p
}

type VerifySubscribeRankingScoreByStampTaskRequest struct {
	ContextStack *string `json:"contextStack"`
	StampTask    *string `json:"stampTask"`
	KeyId        *string `json:"keyId"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *VerifySubscribeRankingScoreByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifySubscribeRankingScoreByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifySubscribeRankingScoreByStampTaskRequest{}
	} else {
		*p = VerifySubscribeRankingScoreByStampTaskRequest{}
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

func NewVerifySubscribeRankingScoreByStampTaskRequestFromJson(data string) (VerifySubscribeRankingScoreByStampTaskRequest, error) {
	req := VerifySubscribeRankingScoreByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifySubscribeRankingScoreByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifySubscribeRankingScoreByStampTaskRequestFromDict(data map[string]interface{}) VerifySubscribeRankingScoreByStampTaskRequest {
	return VerifySubscribeRankingScoreByStampTaskRequest{
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

func (p VerifySubscribeRankingScoreByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifySubscribeRankingScoreByStampTaskRequest) Pointer() *VerifySubscribeRankingScoreByStampTaskRequest {
	return &p
}

type DescribeSubscribeRankingsRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	RankingName   *string `json:"rankingName"`
	Season        *int64  `json:"season"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingsRequest{}
	} else {
		*p = DescribeSubscribeRankingsRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeSubscribeRankingsRequestFromJson(data string) (DescribeSubscribeRankingsRequest, error) {
	req := DescribeSubscribeRankingsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingsRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingsRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingsRequest {
	return DescribeSubscribeRankingsRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeSubscribeRankingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"rankingName":   p.RankingName,
		"season":        p.Season,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribeRankingsRequest) Pointer() *DescribeSubscribeRankingsRequest {
	return &p
}

type DescribeSubscribeRankingsByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RankingName     *string `json:"rankingName"`
	Season          *int64  `json:"season"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *DescribeSubscribeRankingsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSubscribeRankingsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSubscribeRankingsByUserIdRequest{}
	} else {
		*p = DescribeSubscribeRankingsByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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

func NewDescribeSubscribeRankingsByUserIdRequestFromJson(data string) (DescribeSubscribeRankingsByUserIdRequest, error) {
	req := DescribeSubscribeRankingsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSubscribeRankingsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSubscribeRankingsByUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribeRankingsByUserIdRequest {
	return DescribeSubscribeRankingsByUserIdRequest{
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
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

func (p DescribeSubscribeRankingsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"rankingName":     p.RankingName,
		"season":          p.Season,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSubscribeRankingsByUserIdRequest) Pointer() *DescribeSubscribeRankingsByUserIdRequest {
	return &p
}

type GetSubscribeRankingRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	Season        *int64  `json:"season"`
	ScorerUserId  *string `json:"scorerUserId"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingRequest{}
	} else {
		*p = GetSubscribeRankingRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["scorerUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScorerUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScorerUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScorerUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScorerUserId)
				}
			}
		}
	}
	return nil
}

func NewGetSubscribeRankingRequestFromJson(data string) (GetSubscribeRankingRequest, error) {
	req := GetSubscribeRankingRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingRequestFromDict(data map[string]interface{}) GetSubscribeRankingRequest {
	return GetSubscribeRankingRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		ScorerUserId: func() *string {
			v, ok := data["scorerUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scorerUserId"])
		}(),
	}
}

func (p GetSubscribeRankingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"season":        p.Season,
		"scorerUserId":  p.ScorerUserId,
	}
}

func (p GetSubscribeRankingRequest) Pointer() *GetSubscribeRankingRequest {
	return &p
}

type GetSubscribeRankingByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	Season          *int64  `json:"season"`
	ScorerUserId    *string `json:"scorerUserId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetSubscribeRankingByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRankingByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRankingByUserIdRequest{}
	} else {
		*p = GetSubscribeRankingByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["scorerUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScorerUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScorerUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScorerUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScorerUserId)
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

func NewGetSubscribeRankingByUserIdRequestFromJson(data string) (GetSubscribeRankingByUserIdRequest, error) {
	req := GetSubscribeRankingByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRankingByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRankingByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeRankingByUserIdRequest {
	return GetSubscribeRankingByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		ScorerUserId: func() *string {
			v, ok := data["scorerUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scorerUserId"])
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

func (p GetSubscribeRankingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"season":          p.Season,
		"scorerUserId":    p.ScorerUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSubscribeRankingByUserIdRequest) Pointer() *GetSubscribeRankingByUserIdRequest {
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

type GetCurrentRankingMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetCurrentRankingMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentRankingMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentRankingMasterRequest{}
	} else {
		*p = GetCurrentRankingMasterRequest{}
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

func NewGetCurrentRankingMasterRequestFromJson(data string) (GetCurrentRankingMasterRequest, error) {
	req := GetCurrentRankingMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentRankingMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentRankingMasterRequestFromDict(data map[string]interface{}) GetCurrentRankingMasterRequest {
	return GetCurrentRankingMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentRankingMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentRankingMasterRequest) Pointer() *GetCurrentRankingMasterRequest {
	return &p
}

type UpdateCurrentRankingMasterRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	Settings      *string `json:"settings"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *UpdateCurrentRankingMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentRankingMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentRankingMasterRequest{}
	} else {
		*p = UpdateCurrentRankingMasterRequest{}
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

func NewUpdateCurrentRankingMasterRequestFromJson(data string) (UpdateCurrentRankingMasterRequest, error) {
	req := UpdateCurrentRankingMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentRankingMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentRankingMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterRequest {
	return UpdateCurrentRankingMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
	}
}

func (p UpdateCurrentRankingMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentRankingMasterRequest) Pointer() *UpdateCurrentRankingMasterRequest {
	return &p
}

type UpdateCurrentRankingMasterFromGitHubRequest struct {
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
	DryRun          *bool                  `json:"dryRun"`
}

func (p *UpdateCurrentRankingMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentRankingMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentRankingMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentRankingMasterFromGitHubRequest{}
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

func NewUpdateCurrentRankingMasterFromGitHubRequestFromJson(data string) (UpdateCurrentRankingMasterFromGitHubRequest, error) {
	req := UpdateCurrentRankingMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentRankingMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentRankingMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubRequest {
	return UpdateCurrentRankingMasterFromGitHubRequest{
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

func (p UpdateCurrentRankingMasterFromGitHubRequest) ToDict() map[string]interface{} {
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

func (p UpdateCurrentRankingMasterFromGitHubRequest) Pointer() *UpdateCurrentRankingMasterFromGitHubRequest {
	return &p
}

type GetSubscribeRequest struct {
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RankingName   *string `json:"rankingName"`
	AccessToken   *string `json:"accessToken"`
	TargetUserId  *string `json:"targetUserId"`
	DryRun        *bool   `json:"dryRun"`
}

func (p *GetSubscribeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeRequest{}
	} else {
		*p = GetSubscribeRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
				}
			}
		}
	}
	return nil
}

func NewGetSubscribeRequestFromJson(data string) (GetSubscribeRequest, error) {
	req := GetSubscribeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeRequestFromDict(data map[string]interface{}) GetSubscribeRequest {
	return GetSubscribeRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
		}(),
	}
}

func (p GetSubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p GetSubscribeRequest) Pointer() *GetSubscribeRequest {
	return &p
}

type GetSubscribeByUserIdRequest struct {
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RankingName     *string `json:"rankingName"`
	UserId          *string `json:"userId"`
	TargetUserId    *string `json:"targetUserId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
	DryRun          *bool   `json:"dryRun"`
}

func (p *GetSubscribeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSubscribeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSubscribeByUserIdRequest{}
	} else {
		*p = GetSubscribeByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
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

func NewGetSubscribeByUserIdRequestFromJson(data string) (GetSubscribeByUserIdRequest, error) {
	req := GetSubscribeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSubscribeByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSubscribeByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeByUserIdRequest {
	return GetSubscribeByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
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

func (p GetSubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSubscribeByUserIdRequest) Pointer() *GetSubscribeByUserIdRequest {
	return &p
}

type DeleteSubscribeRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteSubscribeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSubscribeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSubscribeRequest{}
	} else {
		*p = DeleteSubscribeRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
				}
			}
		}
	}
	return nil
}

func NewDeleteSubscribeRequestFromJson(data string) (DeleteSubscribeRequest, error) {
	req := DeleteSubscribeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSubscribeRequest{}, err
	}
	return req, nil
}

func NewDeleteSubscribeRequestFromDict(data map[string]interface{}) DeleteSubscribeRequest {
	return DeleteSubscribeRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
		}(),
	}
}

func (p DeleteSubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"rankingName":   p.RankingName,
		"accessToken":   p.AccessToken,
		"targetUserId":  p.TargetUserId,
	}
}

func (p DeleteSubscribeRequest) Pointer() *DeleteSubscribeRequest {
	return &p
}

type DeleteSubscribeByUserIdRequest struct {
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RankingName        *string `json:"rankingName"`
	UserId             *string `json:"userId"`
	TargetUserId       *string `json:"targetUserId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
	DryRun             *bool   `json:"dryRun"`
}

func (p *DeleteSubscribeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteSubscribeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteSubscribeByUserIdRequest{}
	} else {
		*p = DeleteSubscribeByUserIdRequest{}
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
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
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

func NewDeleteSubscribeByUserIdRequestFromJson(data string) (DeleteSubscribeByUserIdRequest, error) {
	req := DeleteSubscribeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteSubscribeByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteSubscribeByUserIdRequestFromDict(data map[string]interface{}) DeleteSubscribeByUserIdRequest {
	return DeleteSubscribeByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
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

func (p DeleteSubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"rankingName":     p.RankingName,
		"userId":          p.UserId,
		"targetUserId":    p.TargetUserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteSubscribeByUserIdRequest) Pointer() *DeleteSubscribeByUserIdRequest {
	return &p
}
