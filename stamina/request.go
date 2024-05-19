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

package stamina

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
	SourceRequestId       *string     `json:"sourceRequestId"`
	RequestId             *string     `json:"requestId"`
	ContextStack          *string     `json:"contextStack"`
	Name                  *string     `json:"name"`
	Description           *string     `json:"description"`
	OverflowTriggerScript *string     `json:"overflowTriggerScript"`
	LogSetting            *LogSetting `json:"logSetting"`
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
		if v, ok := d["overflowTriggerScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.OverflowTriggerScript = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.OverflowTriggerScript = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.OverflowTriggerScript = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.OverflowTriggerScript = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.OverflowTriggerScript = &strValue
			default:
				if err := json.Unmarshal(*v, &p.OverflowTriggerScript); err != nil {
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
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		OverflowTriggerScript: core.CastString(data["overflowTriggerScript"]),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                  p.Name,
		"description":           p.Description,
		"overflowTriggerScript": p.OverflowTriggerScript,
		"logSetting":            p.LogSetting.ToDict(),
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
	SourceRequestId       *string     `json:"sourceRequestId"`
	RequestId             *string     `json:"requestId"`
	ContextStack          *string     `json:"contextStack"`
	NamespaceName         *string     `json:"namespaceName"`
	Description           *string     `json:"description"`
	OverflowTriggerScript *string     `json:"overflowTriggerScript"`
	LogSetting            *LogSetting `json:"logSetting"`
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
		if v, ok := d["overflowTriggerScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.OverflowTriggerScript = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.OverflowTriggerScript = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.OverflowTriggerScript = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.OverflowTriggerScript = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.OverflowTriggerScript = &strValue
			default:
				if err := json.Unmarshal(*v, &p.OverflowTriggerScript); err != nil {
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
		NamespaceName:         core.CastString(data["namespaceName"]),
		Description:           core.CastString(data["description"]),
		OverflowTriggerScript: core.CastString(data["overflowTriggerScript"]),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"description":           p.Description,
		"overflowTriggerScript": p.OverflowTriggerScript,
		"logSetting":            p.LogSetting.ToDict(),
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

type DumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UploadToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
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
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["uploadToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UploadToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
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

type DescribeStaminaModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeStaminaModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminaModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminaModelMastersRequest{}
	} else {
		*p = DescribeStaminaModelMastersRequest{}
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

func NewDescribeStaminaModelMastersRequestFromJson(data string) (DescribeStaminaModelMastersRequest, error) {
	req := DescribeStaminaModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminaModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminaModelMastersRequestFromDict(data map[string]interface{}) DescribeStaminaModelMastersRequest {
	return DescribeStaminaModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStaminaModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStaminaModelMastersRequest) Pointer() *DescribeStaminaModelMastersRequest {
	return &p
}

type CreateStaminaModelMasterRequest struct {
	SourceRequestId          *string `json:"sourceRequestId"`
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	Name                     *string `json:"name"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
}

func (p *CreateStaminaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateStaminaModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateStaminaModelMasterRequest{}
	} else {
		*p = CreateStaminaModelMasterRequest{}
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverIntervalMinutes); err != nil {
				return err
			}
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverValue); err != nil {
				return err
			}
		}
		if v, ok := d["initialCapacity"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.InitialCapacity); err != nil {
				return err
			}
		}
		if v, ok := d["isOverflow"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.IsOverflow); err != nil {
				return err
			}
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.MaxCapacity); err != nil {
				return err
			}
		}
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.MaxStaminaTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.MaxStaminaTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.MaxStaminaTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.MaxStaminaTableName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverIntervalTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverIntervalTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverIntervalTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverIntervalTableName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverValueTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverValueTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverValueTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverValueTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCreateStaminaModelMasterRequestFromJson(data string) (CreateStaminaModelMasterRequest, error) {
	req := CreateStaminaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateStaminaModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateStaminaModelMasterRequestFromDict(data map[string]interface{}) CreateStaminaModelMasterRequest {
	return CreateStaminaModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		RecoverIntervalMinutes:   core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:             core.CastInt32(data["recoverValue"]),
		InitialCapacity:          core.CastInt32(data["initialCapacity"]),
		IsOverflow:               core.CastBool(data["isOverflow"]),
		MaxCapacity:              core.CastInt32(data["maxCapacity"]),
		MaxStaminaTableName:      core.CastString(data["maxStaminaTableName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		RecoverValueTableName:    core.CastString(data["recoverValueTableName"]),
	}
}

func (p CreateStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"recoverIntervalMinutes":   p.RecoverIntervalMinutes,
		"recoverValue":             p.RecoverValue,
		"initialCapacity":          p.InitialCapacity,
		"isOverflow":               p.IsOverflow,
		"maxCapacity":              p.MaxCapacity,
		"maxStaminaTableName":      p.MaxStaminaTableName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"recoverValueTableName":    p.RecoverValueTableName,
	}
}

func (p CreateStaminaModelMasterRequest) Pointer() *CreateStaminaModelMasterRequest {
	return &p
}

type GetStaminaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
}

func (p *GetStaminaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaModelMasterRequest{}
	} else {
		*p = GetStaminaModelMasterRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetStaminaModelMasterRequestFromJson(data string) (GetStaminaModelMasterRequest, error) {
	req := GetStaminaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetStaminaModelMasterRequestFromDict(data map[string]interface{}) GetStaminaModelMasterRequest {
	return GetStaminaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p GetStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p GetStaminaModelMasterRequest) Pointer() *GetStaminaModelMasterRequest {
	return &p
}

type UpdateStaminaModelMasterRequest struct {
	SourceRequestId          *string `json:"sourceRequestId"`
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	StaminaName              *string `json:"staminaName"`
	Description              *string `json:"description"`
	Metadata                 *string `json:"metadata"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
}

func (p *UpdateStaminaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStaminaModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateStaminaModelMasterRequest{}
	} else {
		*p = UpdateStaminaModelMasterRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverIntervalMinutes); err != nil {
				return err
			}
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverValue); err != nil {
				return err
			}
		}
		if v, ok := d["initialCapacity"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.InitialCapacity); err != nil {
				return err
			}
		}
		if v, ok := d["isOverflow"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.IsOverflow); err != nil {
				return err
			}
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.MaxCapacity); err != nil {
				return err
			}
		}
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.MaxStaminaTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.MaxStaminaTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.MaxStaminaTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.MaxStaminaTableName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverIntervalTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverIntervalTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverIntervalTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverIntervalTableName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverValueTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverValueTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverValueTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverValueTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateStaminaModelMasterRequestFromJson(data string) (UpdateStaminaModelMasterRequest, error) {
	req := UpdateStaminaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStaminaModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateStaminaModelMasterRequestFromDict(data map[string]interface{}) UpdateStaminaModelMasterRequest {
	return UpdateStaminaModelMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		StaminaName:              core.CastString(data["staminaName"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		RecoverIntervalMinutes:   core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:             core.CastInt32(data["recoverValue"]),
		InitialCapacity:          core.CastInt32(data["initialCapacity"]),
		IsOverflow:               core.CastBool(data["isOverflow"]),
		MaxCapacity:              core.CastInt32(data["maxCapacity"]),
		MaxStaminaTableName:      core.CastString(data["maxStaminaTableName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		RecoverValueTableName:    core.CastString(data["recoverValueTableName"]),
	}
}

func (p UpdateStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"staminaName":              p.StaminaName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"recoverIntervalMinutes":   p.RecoverIntervalMinutes,
		"recoverValue":             p.RecoverValue,
		"initialCapacity":          p.InitialCapacity,
		"isOverflow":               p.IsOverflow,
		"maxCapacity":              p.MaxCapacity,
		"maxStaminaTableName":      p.MaxStaminaTableName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"recoverValueTableName":    p.RecoverValueTableName,
	}
}

func (p UpdateStaminaModelMasterRequest) Pointer() *UpdateStaminaModelMasterRequest {
	return &p
}

type DeleteStaminaModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
}

func (p *DeleteStaminaModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStaminaModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteStaminaModelMasterRequest{}
	} else {
		*p = DeleteStaminaModelMasterRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteStaminaModelMasterRequestFromJson(data string) (DeleteStaminaModelMasterRequest, error) {
	req := DeleteStaminaModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStaminaModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteStaminaModelMasterRequestFromDict(data map[string]interface{}) DeleteStaminaModelMasterRequest {
	return DeleteStaminaModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p DeleteStaminaModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p DeleteStaminaModelMasterRequest) Pointer() *DeleteStaminaModelMasterRequest {
	return &p
}

type DescribeMaxStaminaTableMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeMaxStaminaTableMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeMaxStaminaTableMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeMaxStaminaTableMastersRequest{}
	} else {
		*p = DescribeMaxStaminaTableMastersRequest{}
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

func NewDescribeMaxStaminaTableMastersRequestFromJson(data string) (DescribeMaxStaminaTableMastersRequest, error) {
	req := DescribeMaxStaminaTableMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeMaxStaminaTableMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeMaxStaminaTableMastersRequestFromDict(data map[string]interface{}) DescribeMaxStaminaTableMastersRequest {
	return DescribeMaxStaminaTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMaxStaminaTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMaxStaminaTableMastersRequest) Pointer() *DescribeMaxStaminaTableMastersRequest {
	return &p
}

type CreateMaxStaminaTableMasterRequest struct {
	SourceRequestId   *string  `json:"sourceRequestId"`
	RequestId         *string  `json:"requestId"`
	ContextStack      *string  `json:"contextStack"`
	NamespaceName     *string  `json:"namespaceName"`
	Name              *string  `json:"name"`
	Description       *string  `json:"description"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *CreateMaxStaminaTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateMaxStaminaTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateMaxStaminaTableMasterRequest{}
	} else {
		*p = CreateMaxStaminaTableMasterRequest{}
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCreateMaxStaminaTableMasterRequestFromJson(data string) (CreateMaxStaminaTableMasterRequest, error) {
	req := CreateMaxStaminaTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateMaxStaminaTableMasterRequest{}, err
	}
	return req, nil
}

func NewCreateMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) CreateMaxStaminaTableMasterRequest {
	return CreateMaxStaminaTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateMaxStaminaTableMasterRequest) Pointer() *CreateMaxStaminaTableMasterRequest {
	return &p
}

type GetMaxStaminaTableMasterRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
}

func (p *GetMaxStaminaTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetMaxStaminaTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetMaxStaminaTableMasterRequest{}
	} else {
		*p = GetMaxStaminaTableMasterRequest{}
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
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.MaxStaminaTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.MaxStaminaTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.MaxStaminaTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.MaxStaminaTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetMaxStaminaTableMasterRequestFromJson(data string) (GetMaxStaminaTableMasterRequest, error) {
	req := GetMaxStaminaTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetMaxStaminaTableMasterRequest{}, err
	}
	return req, nil
}

func NewGetMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) GetMaxStaminaTableMasterRequest {
	return GetMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
	}
}

func (p GetMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
	}
}

func (p GetMaxStaminaTableMasterRequest) Pointer() *GetMaxStaminaTableMasterRequest {
	return &p
}

type UpdateMaxStaminaTableMasterRequest struct {
	SourceRequestId     *string  `json:"sourceRequestId"`
	RequestId           *string  `json:"requestId"`
	ContextStack        *string  `json:"contextStack"`
	NamespaceName       *string  `json:"namespaceName"`
	MaxStaminaTableName *string  `json:"maxStaminaTableName"`
	Description         *string  `json:"description"`
	Metadata            *string  `json:"metadata"`
	ExperienceModelId   *string  `json:"experienceModelId"`
	Values              []*int32 `json:"values"`
}

func (p *UpdateMaxStaminaTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateMaxStaminaTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateMaxStaminaTableMasterRequest{}
	} else {
		*p = UpdateMaxStaminaTableMasterRequest{}
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
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.MaxStaminaTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.MaxStaminaTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.MaxStaminaTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.MaxStaminaTableName); err != nil {
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateMaxStaminaTableMasterRequestFromJson(data string) (UpdateMaxStaminaTableMasterRequest, error) {
	req := UpdateMaxStaminaTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateMaxStaminaTableMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) UpdateMaxStaminaTableMasterRequest {
	return UpdateMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		ExperienceModelId:   core.CastString(data["experienceModelId"]),
		Values:              core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"experienceModelId":   p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateMaxStaminaTableMasterRequest) Pointer() *UpdateMaxStaminaTableMasterRequest {
	return &p
}

type DeleteMaxStaminaTableMasterRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
}

func (p *DeleteMaxStaminaTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteMaxStaminaTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteMaxStaminaTableMasterRequest{}
	} else {
		*p = DeleteMaxStaminaTableMasterRequest{}
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
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.MaxStaminaTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.MaxStaminaTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.MaxStaminaTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.MaxStaminaTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.MaxStaminaTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteMaxStaminaTableMasterRequestFromJson(data string) (DeleteMaxStaminaTableMasterRequest, error) {
	req := DeleteMaxStaminaTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteMaxStaminaTableMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteMaxStaminaTableMasterRequestFromDict(data map[string]interface{}) DeleteMaxStaminaTableMasterRequest {
	return DeleteMaxStaminaTableMasterRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
	}
}

func (p DeleteMaxStaminaTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"maxStaminaTableName": p.MaxStaminaTableName,
	}
}

func (p DeleteMaxStaminaTableMasterRequest) Pointer() *DeleteMaxStaminaTableMasterRequest {
	return &p
}

type DescribeRecoverIntervalTableMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeRecoverIntervalTableMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRecoverIntervalTableMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRecoverIntervalTableMastersRequest{}
	} else {
		*p = DescribeRecoverIntervalTableMastersRequest{}
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

func NewDescribeRecoverIntervalTableMastersRequestFromJson(data string) (DescribeRecoverIntervalTableMastersRequest, error) {
	req := DescribeRecoverIntervalTableMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRecoverIntervalTableMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeRecoverIntervalTableMastersRequestFromDict(data map[string]interface{}) DescribeRecoverIntervalTableMastersRequest {
	return DescribeRecoverIntervalTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRecoverIntervalTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRecoverIntervalTableMastersRequest) Pointer() *DescribeRecoverIntervalTableMastersRequest {
	return &p
}

type CreateRecoverIntervalTableMasterRequest struct {
	SourceRequestId   *string  `json:"sourceRequestId"`
	RequestId         *string  `json:"requestId"`
	ContextStack      *string  `json:"contextStack"`
	NamespaceName     *string  `json:"namespaceName"`
	Name              *string  `json:"name"`
	Description       *string  `json:"description"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *CreateRecoverIntervalTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateRecoverIntervalTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateRecoverIntervalTableMasterRequest{}
	} else {
		*p = CreateRecoverIntervalTableMasterRequest{}
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCreateRecoverIntervalTableMasterRequestFromJson(data string) (CreateRecoverIntervalTableMasterRequest, error) {
	req := CreateRecoverIntervalTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateRecoverIntervalTableMasterRequest{}, err
	}
	return req, nil
}

func NewCreateRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) CreateRecoverIntervalTableMasterRequest {
	return CreateRecoverIntervalTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateRecoverIntervalTableMasterRequest) Pointer() *CreateRecoverIntervalTableMasterRequest {
	return &p
}

type GetRecoverIntervalTableMasterRequest struct {
	SourceRequestId          *string `json:"sourceRequestId"`
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
}

func (p *GetRecoverIntervalTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRecoverIntervalTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRecoverIntervalTableMasterRequest{}
	} else {
		*p = GetRecoverIntervalTableMasterRequest{}
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
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverIntervalTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverIntervalTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverIntervalTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverIntervalTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetRecoverIntervalTableMasterRequestFromJson(data string) (GetRecoverIntervalTableMasterRequest, error) {
	req := GetRecoverIntervalTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRecoverIntervalTableMasterRequest{}, err
	}
	return req, nil
}

func NewGetRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) GetRecoverIntervalTableMasterRequest {
	return GetRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
	}
}

func (p GetRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
	}
}

func (p GetRecoverIntervalTableMasterRequest) Pointer() *GetRecoverIntervalTableMasterRequest {
	return &p
}

type UpdateRecoverIntervalTableMasterRequest struct {
	SourceRequestId          *string  `json:"sourceRequestId"`
	RequestId                *string  `json:"requestId"`
	ContextStack             *string  `json:"contextStack"`
	NamespaceName            *string  `json:"namespaceName"`
	RecoverIntervalTableName *string  `json:"recoverIntervalTableName"`
	Description              *string  `json:"description"`
	Metadata                 *string  `json:"metadata"`
	ExperienceModelId        *string  `json:"experienceModelId"`
	Values                   []*int32 `json:"values"`
}

func (p *UpdateRecoverIntervalTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateRecoverIntervalTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateRecoverIntervalTableMasterRequest{}
	} else {
		*p = UpdateRecoverIntervalTableMasterRequest{}
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
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverIntervalTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverIntervalTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverIntervalTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverIntervalTableName); err != nil {
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateRecoverIntervalTableMasterRequestFromJson(data string) (UpdateRecoverIntervalTableMasterRequest, error) {
	req := UpdateRecoverIntervalTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateRecoverIntervalTableMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) UpdateRecoverIntervalTableMasterRequest {
	return UpdateRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		Description:              core.CastString(data["description"]),
		Metadata:                 core.CastString(data["metadata"]),
		ExperienceModelId:        core.CastString(data["experienceModelId"]),
		Values:                   core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"experienceModelId":        p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateRecoverIntervalTableMasterRequest) Pointer() *UpdateRecoverIntervalTableMasterRequest {
	return &p
}

type DeleteRecoverIntervalTableMasterRequest struct {
	SourceRequestId          *string `json:"sourceRequestId"`
	RequestId                *string `json:"requestId"`
	ContextStack             *string `json:"contextStack"`
	NamespaceName            *string `json:"namespaceName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
}

func (p *DeleteRecoverIntervalTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRecoverIntervalTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRecoverIntervalTableMasterRequest{}
	} else {
		*p = DeleteRecoverIntervalTableMasterRequest{}
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
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverIntervalTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverIntervalTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverIntervalTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverIntervalTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverIntervalTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteRecoverIntervalTableMasterRequestFromJson(data string) (DeleteRecoverIntervalTableMasterRequest, error) {
	req := DeleteRecoverIntervalTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRecoverIntervalTableMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteRecoverIntervalTableMasterRequestFromDict(data map[string]interface{}) DeleteRecoverIntervalTableMasterRequest {
	return DeleteRecoverIntervalTableMasterRequest{
		NamespaceName:            core.CastString(data["namespaceName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
	}
}

func (p DeleteRecoverIntervalTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":            p.NamespaceName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
	}
}

func (p DeleteRecoverIntervalTableMasterRequest) Pointer() *DeleteRecoverIntervalTableMasterRequest {
	return &p
}

type DescribeRecoverValueTableMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeRecoverValueTableMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeRecoverValueTableMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeRecoverValueTableMastersRequest{}
	} else {
		*p = DescribeRecoverValueTableMastersRequest{}
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

func NewDescribeRecoverValueTableMastersRequestFromJson(data string) (DescribeRecoverValueTableMastersRequest, error) {
	req := DescribeRecoverValueTableMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeRecoverValueTableMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeRecoverValueTableMastersRequestFromDict(data map[string]interface{}) DescribeRecoverValueTableMastersRequest {
	return DescribeRecoverValueTableMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRecoverValueTableMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRecoverValueTableMastersRequest) Pointer() *DescribeRecoverValueTableMastersRequest {
	return &p
}

type CreateRecoverValueTableMasterRequest struct {
	SourceRequestId   *string  `json:"sourceRequestId"`
	RequestId         *string  `json:"requestId"`
	ContextStack      *string  `json:"contextStack"`
	NamespaceName     *string  `json:"namespaceName"`
	Name              *string  `json:"name"`
	Description       *string  `json:"description"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *CreateRecoverValueTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateRecoverValueTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateRecoverValueTableMasterRequest{}
	} else {
		*p = CreateRecoverValueTableMasterRequest{}
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCreateRecoverValueTableMasterRequestFromJson(data string) (CreateRecoverValueTableMasterRequest, error) {
	req := CreateRecoverValueTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateRecoverValueTableMasterRequest{}, err
	}
	return req, nil
}

func NewCreateRecoverValueTableMasterRequestFromDict(data map[string]interface{}) CreateRecoverValueTableMasterRequest {
	return CreateRecoverValueTableMasterRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p CreateRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p CreateRecoverValueTableMasterRequest) Pointer() *CreateRecoverValueTableMasterRequest {
	return &p
}

type GetRecoverValueTableMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
}

func (p *GetRecoverValueTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetRecoverValueTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetRecoverValueTableMasterRequest{}
	} else {
		*p = GetRecoverValueTableMasterRequest{}
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
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverValueTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverValueTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverValueTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverValueTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetRecoverValueTableMasterRequestFromJson(data string) (GetRecoverValueTableMasterRequest, error) {
	req := GetRecoverValueTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetRecoverValueTableMasterRequest{}, err
	}
	return req, nil
}

func NewGetRecoverValueTableMasterRequestFromDict(data map[string]interface{}) GetRecoverValueTableMasterRequest {
	return GetRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
	}
}

func (p GetRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
	}
}

func (p GetRecoverValueTableMasterRequest) Pointer() *GetRecoverValueTableMasterRequest {
	return &p
}

type UpdateRecoverValueTableMasterRequest struct {
	SourceRequestId       *string  `json:"sourceRequestId"`
	RequestId             *string  `json:"requestId"`
	ContextStack          *string  `json:"contextStack"`
	NamespaceName         *string  `json:"namespaceName"`
	RecoverValueTableName *string  `json:"recoverValueTableName"`
	Description           *string  `json:"description"`
	Metadata              *string  `json:"metadata"`
	ExperienceModelId     *string  `json:"experienceModelId"`
	Values                []*int32 `json:"values"`
}

func (p *UpdateRecoverValueTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateRecoverValueTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateRecoverValueTableMasterRequest{}
	} else {
		*p = UpdateRecoverValueTableMasterRequest{}
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
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverValueTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverValueTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverValueTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverValueTableName); err != nil {
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Metadata); err != nil {
					return err
				}
			}
		}
		if v, ok := d["experienceModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.ExperienceModelId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Values); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateRecoverValueTableMasterRequestFromJson(data string) (UpdateRecoverValueTableMasterRequest, error) {
	req := UpdateRecoverValueTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateRecoverValueTableMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateRecoverValueTableMasterRequestFromDict(data map[string]interface{}) UpdateRecoverValueTableMasterRequest {
	return UpdateRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ExperienceModelId:     core.CastString(data["experienceModelId"]),
		Values:                core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p UpdateRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
		"description":           p.Description,
		"metadata":              p.Metadata,
		"experienceModelId":     p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p UpdateRecoverValueTableMasterRequest) Pointer() *UpdateRecoverValueTableMasterRequest {
	return &p
}

type DeleteRecoverValueTableMasterRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	NamespaceName         *string `json:"namespaceName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
}

func (p *DeleteRecoverValueTableMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRecoverValueTableMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRecoverValueTableMasterRequest{}
	} else {
		*p = DeleteRecoverValueTableMasterRequest{}
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
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.RecoverValueTableName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.RecoverValueTableName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.RecoverValueTableName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.RecoverValueTableName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.RecoverValueTableName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteRecoverValueTableMasterRequestFromJson(data string) (DeleteRecoverValueTableMasterRequest, error) {
	req := DeleteRecoverValueTableMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRecoverValueTableMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteRecoverValueTableMasterRequestFromDict(data map[string]interface{}) DeleteRecoverValueTableMasterRequest {
	return DeleteRecoverValueTableMasterRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
	}
}

func (p DeleteRecoverValueTableMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"recoverValueTableName": p.RecoverValueTableName,
	}
}

func (p DeleteRecoverValueTableMasterRequest) Pointer() *DeleteRecoverValueTableMasterRequest {
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

type GetCurrentStaminaMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentStaminaMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentStaminaMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentStaminaMasterRequest{}
	} else {
		*p = GetCurrentStaminaMasterRequest{}
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

func NewGetCurrentStaminaMasterRequestFromJson(data string) (GetCurrentStaminaMasterRequest, error) {
	req := GetCurrentStaminaMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentStaminaMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentStaminaMasterRequestFromDict(data map[string]interface{}) GetCurrentStaminaMasterRequest {
	return GetCurrentStaminaMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentStaminaMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentStaminaMasterRequest) Pointer() *GetCurrentStaminaMasterRequest {
	return &p
}

type UpdateCurrentStaminaMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentStaminaMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentStaminaMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentStaminaMasterRequest{}
	} else {
		*p = UpdateCurrentStaminaMasterRequest{}
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
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Settings); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateCurrentStaminaMasterRequestFromJson(data string) (UpdateCurrentStaminaMasterRequest, error) {
	req := UpdateCurrentStaminaMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentStaminaMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentStaminaMasterRequestFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterRequest {
	return UpdateCurrentStaminaMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentStaminaMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentStaminaMasterRequest) Pointer() *UpdateCurrentStaminaMasterRequest {
	return &p
}

type UpdateCurrentStaminaMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentStaminaMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentStaminaMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentStaminaMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentStaminaMasterFromGitHubRequest{}
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
		if v, ok := d["checkoutSetting"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.CheckoutSetting); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateCurrentStaminaMasterFromGitHubRequestFromJson(data string) (UpdateCurrentStaminaMasterFromGitHubRequest, error) {
	req := UpdateCurrentStaminaMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentStaminaMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentStaminaMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterFromGitHubRequest {
	return UpdateCurrentStaminaMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubRequest) Pointer() *UpdateCurrentStaminaMasterFromGitHubRequest {
	return &p
}

type DescribeStaminaModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeStaminaModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminaModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminaModelsRequest{}
	} else {
		*p = DescribeStaminaModelsRequest{}
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

func NewDescribeStaminaModelsRequestFromJson(data string) (DescribeStaminaModelsRequest, error) {
	req := DescribeStaminaModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminaModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminaModelsRequestFromDict(data map[string]interface{}) DescribeStaminaModelsRequest {
	return DescribeStaminaModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeStaminaModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeStaminaModelsRequest) Pointer() *DescribeStaminaModelsRequest {
	return &p
}

type GetStaminaModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
}

func (p *GetStaminaModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaModelRequest{}
	} else {
		*p = GetStaminaModelRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetStaminaModelRequestFromJson(data string) (GetStaminaModelRequest, error) {
	req := GetStaminaModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaModelRequest{}, err
	}
	return req, nil
}

func NewGetStaminaModelRequestFromDict(data map[string]interface{}) GetStaminaModelRequest {
	return GetStaminaModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
	}
}

func (p GetStaminaModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
	}
}

func (p GetStaminaModelRequest) Pointer() *GetStaminaModelRequest {
	return &p
}

type DescribeStaminasRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeStaminasRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminasRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminasRequest{}
	} else {
		*p = DescribeStaminasRequest{}
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
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
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

func NewDescribeStaminasRequestFromJson(data string) (DescribeStaminasRequest, error) {
	req := DescribeStaminasRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminasRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminasRequestFromDict(data map[string]interface{}) DescribeStaminasRequest {
	return DescribeStaminasRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeStaminasRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeStaminasRequest) Pointer() *DescribeStaminasRequest {
	return &p
}

type DescribeStaminasByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeStaminasByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStaminasByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeStaminasByUserIdRequest{}
	} else {
		*p = DescribeStaminasByUserIdRequest{}
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
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
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
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDescribeStaminasByUserIdRequestFromJson(data string) (DescribeStaminasByUserIdRequest, error) {
	req := DescribeStaminasByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStaminasByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeStaminasByUserIdRequestFromDict(data map[string]interface{}) DescribeStaminasByUserIdRequest {
	return DescribeStaminasByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeStaminasByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeStaminasByUserIdRequest) Pointer() *DescribeStaminasByUserIdRequest {
	return &p
}

type GetStaminaRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
	AccessToken     *string `json:"accessToken"`
}

func (p *GetStaminaRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaRequest{}
	} else {
		*p = GetStaminaRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetStaminaRequestFromJson(data string) (GetStaminaRequest, error) {
	req := GetStaminaRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaRequest{}, err
	}
	return req, nil
}

func NewGetStaminaRequestFromDict(data map[string]interface{}) GetStaminaRequest {
	return GetStaminaRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetStaminaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetStaminaRequest) Pointer() *GetStaminaRequest {
	return &p
}

type GetStaminaByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StaminaName     *string `json:"staminaName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetStaminaByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStaminaByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStaminaByUserIdRequest{}
	} else {
		*p = GetStaminaByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetStaminaByUserIdRequestFromJson(data string) (GetStaminaByUserIdRequest, error) {
	req := GetStaminaByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStaminaByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetStaminaByUserIdRequestFromDict(data map[string]interface{}) GetStaminaByUserIdRequest {
	return GetStaminaByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetStaminaByUserIdRequest) Pointer() *GetStaminaByUserIdRequest {
	return &p
}

type UpdateStaminaByUserIdRequest struct {
	SourceRequestId        *string `json:"sourceRequestId"`
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	DuplicationAvoider     *string `json:"duplicationAvoider"`
	NamespaceName          *string `json:"namespaceName"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	Value                  *int32  `json:"value"`
	MaxValue               *int32  `json:"maxValue"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
	RecoverValue           *int32  `json:"recoverValue"`
	TimeOffsetToken        *string `json:"timeOffsetToken"`
}

func (p *UpdateStaminaByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateStaminaByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateStaminaByUserIdRequest{}
	} else {
		*p = UpdateStaminaByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.Value); err != nil {
				return err
			}
		}
		if v, ok := d["maxValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.MaxValue); err != nil {
				return err
			}
		}
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverIntervalMinutes); err != nil {
				return err
			}
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateStaminaByUserIdRequestFromJson(data string) (UpdateStaminaByUserIdRequest, error) {
	req := UpdateStaminaByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateStaminaByUserIdRequest{}, err
	}
	return req, nil
}

func NewUpdateStaminaByUserIdRequestFromDict(data map[string]interface{}) UpdateStaminaByUserIdRequest {
	return UpdateStaminaByUserIdRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		StaminaName:            core.CastString(data["staminaName"]),
		UserId:                 core.CastString(data["userId"]),
		Value:                  core.CastInt32(data["value"]),
		MaxValue:               core.CastInt32(data["maxValue"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:           core.CastInt32(data["recoverValue"]),
		TimeOffsetToken:        core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"staminaName":            p.StaminaName,
		"userId":                 p.UserId,
		"value":                  p.Value,
		"maxValue":               p.MaxValue,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
		"recoverValue":           p.RecoverValue,
		"timeOffsetToken":        p.TimeOffsetToken,
	}
}

func (p UpdateStaminaByUserIdRequest) Pointer() *UpdateStaminaByUserIdRequest {
	return &p
}

type ConsumeStaminaRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	AccessToken        *string `json:"accessToken"`
	ConsumeValue       *int32  `json:"consumeValue"`
}

func (p *ConsumeStaminaRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeStaminaRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeStaminaRequest{}
	} else {
		*p = ConsumeStaminaRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["consumeValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.ConsumeValue); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewConsumeStaminaRequestFromJson(data string) (ConsumeStaminaRequest, error) {
	req := ConsumeStaminaRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeStaminaRequest{}, err
	}
	return req, nil
}

func NewConsumeStaminaRequestFromDict(data map[string]interface{}) ConsumeStaminaRequest {
	return ConsumeStaminaRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StaminaName:   core.CastString(data["staminaName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		ConsumeValue:  core.CastInt32(data["consumeValue"]),
	}
}

func (p ConsumeStaminaRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"staminaName":   p.StaminaName,
		"accessToken":   p.AccessToken,
		"consumeValue":  p.ConsumeValue,
	}
}

func (p ConsumeStaminaRequest) Pointer() *ConsumeStaminaRequest {
	return &p
}

type ConsumeStaminaByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	ConsumeValue       *int32  `json:"consumeValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *ConsumeStaminaByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeStaminaByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeStaminaByUserIdRequest{}
	} else {
		*p = ConsumeStaminaByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["consumeValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.ConsumeValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewConsumeStaminaByUserIdRequestFromJson(data string) (ConsumeStaminaByUserIdRequest, error) {
	req := ConsumeStaminaByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeStaminaByUserIdRequest{}, err
	}
	return req, nil
}

func NewConsumeStaminaByUserIdRequestFromDict(data map[string]interface{}) ConsumeStaminaByUserIdRequest {
	return ConsumeStaminaByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		ConsumeValue:    core.CastInt32(data["consumeValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ConsumeStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"consumeValue":    p.ConsumeValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p ConsumeStaminaByUserIdRequest) Pointer() *ConsumeStaminaByUserIdRequest {
	return &p
}

type RecoverStaminaByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RecoverValue       *int32  `json:"recoverValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *RecoverStaminaByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverStaminaByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverStaminaByUserIdRequest{}
	} else {
		*p = RecoverStaminaByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewRecoverStaminaByUserIdRequestFromJson(data string) (RecoverStaminaByUserIdRequest, error) {
	req := RecoverStaminaByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RecoverStaminaByUserIdRequest{}, err
	}
	return req, nil
}

func NewRecoverStaminaByUserIdRequestFromDict(data map[string]interface{}) RecoverStaminaByUserIdRequest {
	return RecoverStaminaByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		RecoverValue:    core.CastInt32(data["recoverValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RecoverStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"recoverValue":    p.RecoverValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RecoverStaminaByUserIdRequest) Pointer() *RecoverStaminaByUserIdRequest {
	return &p
}

type RaiseMaxValueByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RaiseValue         *int32  `json:"raiseValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *RaiseMaxValueByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RaiseMaxValueByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RaiseMaxValueByUserIdRequest{}
	} else {
		*p = RaiseMaxValueByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["raiseValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RaiseValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewRaiseMaxValueByUserIdRequestFromJson(data string) (RaiseMaxValueByUserIdRequest, error) {
	req := RaiseMaxValueByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RaiseMaxValueByUserIdRequest{}, err
	}
	return req, nil
}

func NewRaiseMaxValueByUserIdRequestFromDict(data map[string]interface{}) RaiseMaxValueByUserIdRequest {
	return RaiseMaxValueByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		RaiseValue:      core.CastInt32(data["raiseValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RaiseMaxValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"raiseValue":      p.RaiseValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RaiseMaxValueByUserIdRequest) Pointer() *RaiseMaxValueByUserIdRequest {
	return &p
}

type DecreaseMaxValueByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	DecreaseValue      *int32  `json:"decreaseValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DecreaseMaxValueByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecreaseMaxValueByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DecreaseMaxValueByUserIdRequest{}
	} else {
		*p = DecreaseMaxValueByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["decreaseValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.DecreaseValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDecreaseMaxValueByUserIdRequestFromJson(data string) (DecreaseMaxValueByUserIdRequest, error) {
	req := DecreaseMaxValueByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecreaseMaxValueByUserIdRequest{}, err
	}
	return req, nil
}

func NewDecreaseMaxValueByUserIdRequestFromDict(data map[string]interface{}) DecreaseMaxValueByUserIdRequest {
	return DecreaseMaxValueByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		DecreaseValue:   core.CastInt32(data["decreaseValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DecreaseMaxValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"decreaseValue":   p.DecreaseValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DecreaseMaxValueByUserIdRequest) Pointer() *DecreaseMaxValueByUserIdRequest {
	return &p
}

type SetMaxValueByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	MaxValue           *int32  `json:"maxValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *SetMaxValueByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetMaxValueByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetMaxValueByUserIdRequest{}
	} else {
		*p = SetMaxValueByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["maxValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.MaxValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetMaxValueByUserIdRequestFromJson(data string) (SetMaxValueByUserIdRequest, error) {
	req := SetMaxValueByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetMaxValueByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetMaxValueByUserIdRequestFromDict(data map[string]interface{}) SetMaxValueByUserIdRequest {
	return SetMaxValueByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		MaxValue:        core.CastInt32(data["maxValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetMaxValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"maxValue":        p.MaxValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetMaxValueByUserIdRequest) Pointer() *SetMaxValueByUserIdRequest {
	return &p
}

type SetRecoverIntervalByUserIdRequest struct {
	SourceRequestId        *string `json:"sourceRequestId"`
	RequestId              *string `json:"requestId"`
	ContextStack           *string `json:"contextStack"`
	DuplicationAvoider     *string `json:"duplicationAvoider"`
	NamespaceName          *string `json:"namespaceName"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
	TimeOffsetToken        *string `json:"timeOffsetToken"`
}

func (p *SetRecoverIntervalByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverIntervalByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverIntervalByUserIdRequest{}
	} else {
		*p = SetRecoverIntervalByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverIntervalMinutes); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverIntervalByUserIdRequestFromJson(data string) (SetRecoverIntervalByUserIdRequest, error) {
	req := SetRecoverIntervalByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverIntervalByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetRecoverIntervalByUserIdRequestFromDict(data map[string]interface{}) SetRecoverIntervalByUserIdRequest {
	return SetRecoverIntervalByUserIdRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		StaminaName:            core.CastString(data["staminaName"]),
		UserId:                 core.CastString(data["userId"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
		TimeOffsetToken:        core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetRecoverIntervalByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"staminaName":            p.StaminaName,
		"userId":                 p.UserId,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
		"timeOffsetToken":        p.TimeOffsetToken,
	}
}

func (p SetRecoverIntervalByUserIdRequest) Pointer() *SetRecoverIntervalByUserIdRequest {
	return &p
}

type SetRecoverValueByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	RecoverValue       *int32  `json:"recoverValue"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *SetRecoverValueByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverValueByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverValueByUserIdRequest{}
	} else {
		*p = SetRecoverValueByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.RecoverValue); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverValueByUserIdRequestFromJson(data string) (SetRecoverValueByUserIdRequest, error) {
	req := SetRecoverValueByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverValueByUserIdRequest{}, err
	}
	return req, nil
}

func NewSetRecoverValueByUserIdRequestFromDict(data map[string]interface{}) SetRecoverValueByUserIdRequest {
	return SetRecoverValueByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		RecoverValue:    core.CastInt32(data["recoverValue"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SetRecoverValueByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"recoverValue":    p.RecoverValue,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SetRecoverValueByUserIdRequest) Pointer() *SetRecoverValueByUserIdRequest {
	return &p
}

type SetMaxValueByStatusRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func (p *SetMaxValueByStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetMaxValueByStatusRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetMaxValueByStatusRequest{}
	} else {
		*p = SetMaxValueByStatusRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusBody"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusBody = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusBody = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusBody = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusBody); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusSignature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusSignature = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusSignature = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusSignature = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusSignature); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetMaxValueByStatusRequestFromJson(data string) (SetMaxValueByStatusRequest, error) {
	req := SetMaxValueByStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetMaxValueByStatusRequest{}, err
	}
	return req, nil
}

func NewSetMaxValueByStatusRequestFromDict(data map[string]interface{}) SetMaxValueByStatusRequest {
	return SetMaxValueByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetMaxValueByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetMaxValueByStatusRequest) Pointer() *SetMaxValueByStatusRequest {
	return &p
}

type SetRecoverIntervalByStatusRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func (p *SetRecoverIntervalByStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverIntervalByStatusRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverIntervalByStatusRequest{}
	} else {
		*p = SetRecoverIntervalByStatusRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusBody"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusBody = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusBody = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusBody = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusBody); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusSignature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusSignature = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusSignature = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusSignature = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusSignature); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverIntervalByStatusRequestFromJson(data string) (SetRecoverIntervalByStatusRequest, error) {
	req := SetRecoverIntervalByStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverIntervalByStatusRequest{}, err
	}
	return req, nil
}

func NewSetRecoverIntervalByStatusRequestFromDict(data map[string]interface{}) SetRecoverIntervalByStatusRequest {
	return SetRecoverIntervalByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetRecoverIntervalByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetRecoverIntervalByStatusRequest) Pointer() *SetRecoverIntervalByStatusRequest {
	return &p
}

type SetRecoverValueByStatusRequest struct {
	SourceRequestId       *string `json:"sourceRequestId"`
	RequestId             *string `json:"requestId"`
	ContextStack          *string `json:"contextStack"`
	DuplicationAvoider    *string `json:"duplicationAvoider"`
	NamespaceName         *string `json:"namespaceName"`
	StaminaName           *string `json:"staminaName"`
	AccessToken           *string `json:"accessToken"`
	KeyId                 *string `json:"keyId"`
	SignedStatusBody      *string `json:"signedStatusBody"`
	SignedStatusSignature *string `json:"signedStatusSignature"`
}

func (p *SetRecoverValueByStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverValueByStatusRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverValueByStatusRequest{}
	} else {
		*p = SetRecoverValueByStatusRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["accessToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AccessToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusBody"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusBody = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusBody = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusBody = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusBody = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusBody); err != nil {
					return err
				}
			}
		}
		if v, ok := d["signedStatusSignature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.SignedStatusSignature = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.SignedStatusSignature = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.SignedStatusSignature = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.SignedStatusSignature = &strValue
			default:
				if err := json.Unmarshal(*v, &p.SignedStatusSignature); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverValueByStatusRequestFromJson(data string) (SetRecoverValueByStatusRequest, error) {
	req := SetRecoverValueByStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverValueByStatusRequest{}, err
	}
	return req, nil
}

func NewSetRecoverValueByStatusRequestFromDict(data map[string]interface{}) SetRecoverValueByStatusRequest {
	return SetRecoverValueByStatusRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		StaminaName:           core.CastString(data["staminaName"]),
		AccessToken:           core.CastString(data["accessToken"]),
		KeyId:                 core.CastString(data["keyId"]),
		SignedStatusBody:      core.CastString(data["signedStatusBody"]),
		SignedStatusSignature: core.CastString(data["signedStatusSignature"]),
	}
}

func (p SetRecoverValueByStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"staminaName":           p.StaminaName,
		"accessToken":           p.AccessToken,
		"keyId":                 p.KeyId,
		"signedStatusBody":      p.SignedStatusBody,
		"signedStatusSignature": p.SignedStatusSignature,
	}
}

func (p SetRecoverValueByStatusRequest) Pointer() *SetRecoverValueByStatusRequest {
	return &p
}

type DeleteStaminaByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	StaminaName        *string `json:"staminaName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteStaminaByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteStaminaByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteStaminaByUserIdRequest{}
	} else {
		*p = DeleteStaminaByUserIdRequest{}
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
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.StaminaName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.StaminaName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.StaminaName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.StaminaName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.StaminaName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteStaminaByUserIdRequestFromJson(data string) (DeleteStaminaByUserIdRequest, error) {
	req := DeleteStaminaByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteStaminaByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteStaminaByUserIdRequestFromDict(data map[string]interface{}) DeleteStaminaByUserIdRequest {
	return DeleteStaminaByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		StaminaName:     core.CastString(data["staminaName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteStaminaByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"staminaName":     p.StaminaName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteStaminaByUserIdRequest) Pointer() *DeleteStaminaByUserIdRequest {
	return &p
}

type RecoverStaminaByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RecoverStaminaByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverStaminaByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverStaminaByStampSheetRequest{}
	} else {
		*p = RecoverStaminaByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampSheet); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewRecoverStaminaByStampSheetRequestFromJson(data string) (RecoverStaminaByStampSheetRequest, error) {
	req := RecoverStaminaByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RecoverStaminaByStampSheetRequest{}, err
	}
	return req, nil
}

func NewRecoverStaminaByStampSheetRequestFromDict(data map[string]interface{}) RecoverStaminaByStampSheetRequest {
	return RecoverStaminaByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RecoverStaminaByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RecoverStaminaByStampSheetRequest) Pointer() *RecoverStaminaByStampSheetRequest {
	return &p
}

type RaiseMaxValueByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RaiseMaxValueByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RaiseMaxValueByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RaiseMaxValueByStampSheetRequest{}
	} else {
		*p = RaiseMaxValueByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampSheet); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewRaiseMaxValueByStampSheetRequestFromJson(data string) (RaiseMaxValueByStampSheetRequest, error) {
	req := RaiseMaxValueByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RaiseMaxValueByStampSheetRequest{}, err
	}
	return req, nil
}

func NewRaiseMaxValueByStampSheetRequestFromDict(data map[string]interface{}) RaiseMaxValueByStampSheetRequest {
	return RaiseMaxValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RaiseMaxValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RaiseMaxValueByStampSheetRequest) Pointer() *RaiseMaxValueByStampSheetRequest {
	return &p
}

type DecreaseMaxValueByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *DecreaseMaxValueByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecreaseMaxValueByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DecreaseMaxValueByStampTaskRequest{}
	} else {
		*p = DecreaseMaxValueByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampTask); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDecreaseMaxValueByStampTaskRequestFromJson(data string) (DecreaseMaxValueByStampTaskRequest, error) {
	req := DecreaseMaxValueByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecreaseMaxValueByStampTaskRequest{}, err
	}
	return req, nil
}

func NewDecreaseMaxValueByStampTaskRequestFromDict(data map[string]interface{}) DecreaseMaxValueByStampTaskRequest {
	return DecreaseMaxValueByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DecreaseMaxValueByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DecreaseMaxValueByStampTaskRequest) Pointer() *DecreaseMaxValueByStampTaskRequest {
	return &p
}

type SetMaxValueByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetMaxValueByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetMaxValueByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetMaxValueByStampSheetRequest{}
	} else {
		*p = SetMaxValueByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampSheet); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetMaxValueByStampSheetRequestFromJson(data string) (SetMaxValueByStampSheetRequest, error) {
	req := SetMaxValueByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetMaxValueByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetMaxValueByStampSheetRequestFromDict(data map[string]interface{}) SetMaxValueByStampSheetRequest {
	return SetMaxValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetMaxValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetMaxValueByStampSheetRequest) Pointer() *SetMaxValueByStampSheetRequest {
	return &p
}

type SetRecoverIntervalByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetRecoverIntervalByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverIntervalByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverIntervalByStampSheetRequest{}
	} else {
		*p = SetRecoverIntervalByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampSheet); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverIntervalByStampSheetRequestFromJson(data string) (SetRecoverIntervalByStampSheetRequest, error) {
	req := SetRecoverIntervalByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverIntervalByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetRecoverIntervalByStampSheetRequestFromDict(data map[string]interface{}) SetRecoverIntervalByStampSheetRequest {
	return SetRecoverIntervalByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRecoverIntervalByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRecoverIntervalByStampSheetRequest) Pointer() *SetRecoverIntervalByStampSheetRequest {
	return &p
}

type SetRecoverValueByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetRecoverValueByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetRecoverValueByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetRecoverValueByStampSheetRequest{}
	} else {
		*p = SetRecoverValueByStampSheetRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampSheet"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampSheet); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewSetRecoverValueByStampSheetRequestFromJson(data string) (SetRecoverValueByStampSheetRequest, error) {
	req := SetRecoverValueByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetRecoverValueByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetRecoverValueByStampSheetRequestFromDict(data map[string]interface{}) SetRecoverValueByStampSheetRequest {
	return SetRecoverValueByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetRecoverValueByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetRecoverValueByStampSheetRequest) Pointer() *SetRecoverValueByStampSheetRequest {
	return &p
}

type ConsumeStaminaByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *ConsumeStaminaByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeStaminaByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeStaminaByStampTaskRequest{}
	} else {
		*p = ConsumeStaminaByStampTaskRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stampTask"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.StampTask); err != nil {
					return err
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.KeyId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewConsumeStaminaByStampTaskRequestFromJson(data string) (ConsumeStaminaByStampTaskRequest, error) {
	req := ConsumeStaminaByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return ConsumeStaminaByStampTaskRequest{}, err
	}
	return req, nil
}

func NewConsumeStaminaByStampTaskRequestFromDict(data map[string]interface{}) ConsumeStaminaByStampTaskRequest {
	return ConsumeStaminaByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ConsumeStaminaByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ConsumeStaminaByStampTaskRequest) Pointer() *ConsumeStaminaByStampTaskRequest {
	return &p
}
