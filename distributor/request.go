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

package distributor

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
	SourceRequestId               *string              `json:"sourceRequestId"`
	RequestId                     *string              `json:"requestId"`
	ContextStack                  *string              `json:"contextStack"`
	Name                          *string              `json:"name"`
	Description                   *string              `json:"description"`
	AssumeUserId                  *string              `json:"assumeUserId"`
	AutoRunStampSheetNotification *NotificationSetting `json:"autoRunStampSheetNotification"`
	LogSetting                    *LogSetting          `json:"logSetting"`
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
		if v, ok := d["assumeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AssumeUserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["autoRunStampSheetNotification"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.AutoRunStampSheetNotification); err != nil {
				return err
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
		Name:                          core.CastString(data["name"]),
		Description:                   core.CastString(data["description"]),
		AssumeUserId:                  core.CastString(data["assumeUserId"]),
		AutoRunStampSheetNotification: NewNotificationSettingFromDict(core.CastMap(data["autoRunStampSheetNotification"])).Pointer(),
		LogSetting:                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                          p.Name,
		"description":                   p.Description,
		"assumeUserId":                  p.AssumeUserId,
		"autoRunStampSheetNotification": p.AutoRunStampSheetNotification.ToDict(),
		"logSetting":                    p.LogSetting.ToDict(),
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
	SourceRequestId               *string              `json:"sourceRequestId"`
	RequestId                     *string              `json:"requestId"`
	ContextStack                  *string              `json:"contextStack"`
	NamespaceName                 *string              `json:"namespaceName"`
	Description                   *string              `json:"description"`
	AssumeUserId                  *string              `json:"assumeUserId"`
	AutoRunStampSheetNotification *NotificationSetting `json:"autoRunStampSheetNotification"`
	LogSetting                    *LogSetting          `json:"logSetting"`
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
		if v, ok := d["assumeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AssumeUserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["autoRunStampSheetNotification"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.AutoRunStampSheetNotification); err != nil {
				return err
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
		NamespaceName:                 core.CastString(data["namespaceName"]),
		Description:                   core.CastString(data["description"]),
		AssumeUserId:                  core.CastString(data["assumeUserId"]),
		AutoRunStampSheetNotification: NewNotificationSettingFromDict(core.CastMap(data["autoRunStampSheetNotification"])).Pointer(),
		LogSetting:                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                 p.NamespaceName,
		"description":                   p.Description,
		"assumeUserId":                  p.AssumeUserId,
		"autoRunStampSheetNotification": p.AutoRunStampSheetNotification.ToDict(),
		"logSetting":                    p.LogSetting.ToDict(),
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

type DescribeDistributorModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeDistributorModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDistributorModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeDistributorModelMastersRequest{}
	} else {
		*p = DescribeDistributorModelMastersRequest{}
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

func NewDescribeDistributorModelMastersRequestFromJson(data string) (DescribeDistributorModelMastersRequest, error) {
	req := DescribeDistributorModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDistributorModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeDistributorModelMastersRequestFromDict(data map[string]interface{}) DescribeDistributorModelMastersRequest {
	return DescribeDistributorModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeDistributorModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeDistributorModelMastersRequest) Pointer() *DescribeDistributorModelMastersRequest {
	return &p
}

type CreateDistributorModelMasterRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	NamespaceName      *string   `json:"namespaceName"`
	Name               *string   `json:"name"`
	Description        *string   `json:"description"`
	Metadata           *string   `json:"metadata"`
	InboxNamespaceId   *string   `json:"inboxNamespaceId"`
	WhiteListTargetIds []*string `json:"whiteListTargetIds"`
}

func (p *CreateDistributorModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateDistributorModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateDistributorModelMasterRequest{}
	} else {
		*p = CreateDistributorModelMasterRequest{}
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
		if v, ok := d["inboxNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.InboxNamespaceId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.InboxNamespaceId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.InboxNamespaceId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.InboxNamespaceId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.InboxNamespaceId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.InboxNamespaceId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["whiteListTargetIds"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.WhiteListTargetIds); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewCreateDistributorModelMasterRequestFromJson(data string) (CreateDistributorModelMasterRequest, error) {
	req := CreateDistributorModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateDistributorModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateDistributorModelMasterRequestFromDict(data map[string]interface{}) CreateDistributorModelMasterRequest {
	return CreateDistributorModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		InboxNamespaceId:   core.CastString(data["inboxNamespaceId"]),
		WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
	}
}

func (p CreateDistributorModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"name":             p.Name,
		"description":      p.Description,
		"metadata":         p.Metadata,
		"inboxNamespaceId": p.InboxNamespaceId,
		"whiteListTargetIds": core.CastStringsFromDict(
			p.WhiteListTargetIds,
		),
	}
}

func (p CreateDistributorModelMasterRequest) Pointer() *CreateDistributorModelMasterRequest {
	return &p
}

type GetDistributorModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	DistributorName *string `json:"distributorName"`
}

func (p *GetDistributorModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDistributorModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetDistributorModelMasterRequest{}
	} else {
		*p = GetDistributorModelMasterRequest{}
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
		if v, ok := d["distributorName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.DistributorName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.DistributorName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.DistributorName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.DistributorName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetDistributorModelMasterRequestFromJson(data string) (GetDistributorModelMasterRequest, error) {
	req := GetDistributorModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDistributorModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetDistributorModelMasterRequestFromDict(data map[string]interface{}) GetDistributorModelMasterRequest {
	return GetDistributorModelMasterRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		DistributorName: core.CastString(data["distributorName"]),
	}
}

func (p GetDistributorModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"distributorName": p.DistributorName,
	}
}

func (p GetDistributorModelMasterRequest) Pointer() *GetDistributorModelMasterRequest {
	return &p
}

type UpdateDistributorModelMasterRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	NamespaceName      *string   `json:"namespaceName"`
	DistributorName    *string   `json:"distributorName"`
	Description        *string   `json:"description"`
	Metadata           *string   `json:"metadata"`
	InboxNamespaceId   *string   `json:"inboxNamespaceId"`
	WhiteListTargetIds []*string `json:"whiteListTargetIds"`
}

func (p *UpdateDistributorModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateDistributorModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateDistributorModelMasterRequest{}
	} else {
		*p = UpdateDistributorModelMasterRequest{}
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
		if v, ok := d["distributorName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.DistributorName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.DistributorName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.DistributorName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.DistributorName); err != nil {
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
		if v, ok := d["inboxNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.InboxNamespaceId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.InboxNamespaceId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.InboxNamespaceId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.InboxNamespaceId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.InboxNamespaceId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.InboxNamespaceId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["whiteListTargetIds"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.WhiteListTargetIds); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUpdateDistributorModelMasterRequestFromJson(data string) (UpdateDistributorModelMasterRequest, error) {
	req := UpdateDistributorModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateDistributorModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateDistributorModelMasterRequestFromDict(data map[string]interface{}) UpdateDistributorModelMasterRequest {
	return UpdateDistributorModelMasterRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		DistributorName:    core.CastString(data["distributorName"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		InboxNamespaceId:   core.CastString(data["inboxNamespaceId"]),
		WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
	}
}

func (p UpdateDistributorModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"distributorName":  p.DistributorName,
		"description":      p.Description,
		"metadata":         p.Metadata,
		"inboxNamespaceId": p.InboxNamespaceId,
		"whiteListTargetIds": core.CastStringsFromDict(
			p.WhiteListTargetIds,
		),
	}
}

func (p UpdateDistributorModelMasterRequest) Pointer() *UpdateDistributorModelMasterRequest {
	return &p
}

type DeleteDistributorModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	DistributorName *string `json:"distributorName"`
}

func (p *DeleteDistributorModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteDistributorModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteDistributorModelMasterRequest{}
	} else {
		*p = DeleteDistributorModelMasterRequest{}
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
		if v, ok := d["distributorName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.DistributorName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.DistributorName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.DistributorName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.DistributorName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteDistributorModelMasterRequestFromJson(data string) (DeleteDistributorModelMasterRequest, error) {
	req := DeleteDistributorModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteDistributorModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteDistributorModelMasterRequestFromDict(data map[string]interface{}) DeleteDistributorModelMasterRequest {
	return DeleteDistributorModelMasterRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		DistributorName: core.CastString(data["distributorName"]),
	}
}

func (p DeleteDistributorModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"distributorName": p.DistributorName,
	}
}

func (p DeleteDistributorModelMasterRequest) Pointer() *DeleteDistributorModelMasterRequest {
	return &p
}

type DescribeDistributorModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeDistributorModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeDistributorModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeDistributorModelsRequest{}
	} else {
		*p = DescribeDistributorModelsRequest{}
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

func NewDescribeDistributorModelsRequestFromJson(data string) (DescribeDistributorModelsRequest, error) {
	req := DescribeDistributorModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeDistributorModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeDistributorModelsRequestFromDict(data map[string]interface{}) DescribeDistributorModelsRequest {
	return DescribeDistributorModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeDistributorModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeDistributorModelsRequest) Pointer() *DescribeDistributorModelsRequest {
	return &p
}

type GetDistributorModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	DistributorName *string `json:"distributorName"`
}

func (p *GetDistributorModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetDistributorModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetDistributorModelRequest{}
	} else {
		*p = GetDistributorModelRequest{}
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
		if v, ok := d["distributorName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.DistributorName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.DistributorName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.DistributorName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.DistributorName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetDistributorModelRequestFromJson(data string) (GetDistributorModelRequest, error) {
	req := GetDistributorModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetDistributorModelRequest{}, err
	}
	return req, nil
}

func NewGetDistributorModelRequestFromDict(data map[string]interface{}) GetDistributorModelRequest {
	return GetDistributorModelRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		DistributorName: core.CastString(data["distributorName"]),
	}
}

func (p GetDistributorModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"distributorName": p.DistributorName,
	}
}

func (p GetDistributorModelRequest) Pointer() *GetDistributorModelRequest {
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

type GetCurrentDistributorMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentDistributorMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentDistributorMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentDistributorMasterRequest{}
	} else {
		*p = GetCurrentDistributorMasterRequest{}
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

func NewGetCurrentDistributorMasterRequestFromJson(data string) (GetCurrentDistributorMasterRequest, error) {
	req := GetCurrentDistributorMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentDistributorMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentDistributorMasterRequestFromDict(data map[string]interface{}) GetCurrentDistributorMasterRequest {
	return GetCurrentDistributorMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentDistributorMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentDistributorMasterRequest) Pointer() *GetCurrentDistributorMasterRequest {
	return &p
}

type UpdateCurrentDistributorMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentDistributorMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentDistributorMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentDistributorMasterRequest{}
	} else {
		*p = UpdateCurrentDistributorMasterRequest{}
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

func NewUpdateCurrentDistributorMasterRequestFromJson(data string) (UpdateCurrentDistributorMasterRequest, error) {
	req := UpdateCurrentDistributorMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentDistributorMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentDistributorMasterRequestFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterRequest {
	return UpdateCurrentDistributorMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentDistributorMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentDistributorMasterRequest) Pointer() *UpdateCurrentDistributorMasterRequest {
	return &p
}

type UpdateCurrentDistributorMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentDistributorMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentDistributorMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentDistributorMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentDistributorMasterFromGitHubRequest{}
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

func NewUpdateCurrentDistributorMasterFromGitHubRequestFromJson(data string) (UpdateCurrentDistributorMasterFromGitHubRequest, error) {
	req := UpdateCurrentDistributorMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentDistributorMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentDistributorMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentDistributorMasterFromGitHubRequest {
	return UpdateCurrentDistributorMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentDistributorMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentDistributorMasterFromGitHubRequest) Pointer() *UpdateCurrentDistributorMasterFromGitHubRequest {
	return &p
}

type DistributeRequest struct {
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	NamespaceName      *string             `json:"namespaceName"`
	DistributorName    *string             `json:"distributorName"`
	UserId             *string             `json:"userId"`
	DistributeResource *DistributeResource `json:"distributeResource"`
	TimeOffsetToken    *string             `json:"timeOffsetToken"`
}

func (p *DistributeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DistributeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DistributeRequest{}
	} else {
		*p = DistributeRequest{}
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
		if v, ok := d["distributorName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.DistributorName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.DistributorName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.DistributorName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.DistributorName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.DistributorName); err != nil {
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
		if v, ok := d["distributeResource"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.DistributeResource); err != nil {
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

func NewDistributeRequestFromJson(data string) (DistributeRequest, error) {
	req := DistributeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DistributeRequest{}, err
	}
	return req, nil
}

func NewDistributeRequestFromDict(data map[string]interface{}) DistributeRequest {
	return DistributeRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		DistributorName:    core.CastString(data["distributorName"]),
		UserId:             core.CastString(data["userId"]),
		DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
		TimeOffsetToken:    core.CastString(data["timeOffsetToken"]),
	}
}

func (p DistributeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"distributorName":    p.DistributorName,
		"userId":             p.UserId,
		"distributeResource": p.DistributeResource.ToDict(),
		"timeOffsetToken":    p.TimeOffsetToken,
	}
}

func (p DistributeRequest) Pointer() *DistributeRequest {
	return &p
}

type DistributeWithoutOverflowProcessRequest struct {
	SourceRequestId    *string             `json:"sourceRequestId"`
	RequestId          *string             `json:"requestId"`
	ContextStack       *string             `json:"contextStack"`
	UserId             *string             `json:"userId"`
	DistributeResource *DistributeResource `json:"distributeResource"`
	TimeOffsetToken    *string             `json:"timeOffsetToken"`
}

func (p *DistributeWithoutOverflowProcessRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DistributeWithoutOverflowProcessRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DistributeWithoutOverflowProcessRequest{}
	} else {
		*p = DistributeWithoutOverflowProcessRequest{}
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
		if v, ok := d["distributeResource"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			if err := json.Unmarshal(*v, &p.DistributeResource); err != nil {
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

func NewDistributeWithoutOverflowProcessRequestFromJson(data string) (DistributeWithoutOverflowProcessRequest, error) {
	req := DistributeWithoutOverflowProcessRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DistributeWithoutOverflowProcessRequest{}, err
	}
	return req, nil
}

func NewDistributeWithoutOverflowProcessRequestFromDict(data map[string]interface{}) DistributeWithoutOverflowProcessRequest {
	return DistributeWithoutOverflowProcessRequest{
		UserId:             core.CastString(data["userId"]),
		DistributeResource: NewDistributeResourceFromDict(core.CastMap(data["distributeResource"])).Pointer(),
		TimeOffsetToken:    core.CastString(data["timeOffsetToken"]),
	}
}

func (p DistributeWithoutOverflowProcessRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":             p.UserId,
		"distributeResource": p.DistributeResource.ToDict(),
		"timeOffsetToken":    p.TimeOffsetToken,
	}
}

func (p DistributeWithoutOverflowProcessRequest) Pointer() *DistributeWithoutOverflowProcessRequest {
	return &p
}

type RunStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampTaskRequest{}
	} else {
		*p = RunStampTaskRequest{}
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

func NewRunStampTaskRequestFromJson(data string) (RunStampTaskRequest, error) {
	req := RunStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampTaskRequest{}, err
	}
	return req, nil
}

func NewRunStampTaskRequestFromDict(data map[string]interface{}) RunStampTaskRequest {
	return RunStampTaskRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StampTask:     core.CastString(data["stampTask"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p RunStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"stampTask":     p.StampTask,
		"keyId":         p.KeyId,
	}
}

func (p RunStampTaskRequest) Pointer() *RunStampTaskRequest {
	return &p
}

type RunStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampSheetRequest{}
	} else {
		*p = RunStampSheetRequest{}
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

func NewRunStampSheetRequestFromJson(data string) (RunStampSheetRequest, error) {
	req := RunStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampSheetRequest{}, err
	}
	return req, nil
}

func NewRunStampSheetRequestFromDict(data map[string]interface{}) RunStampSheetRequest {
	return RunStampSheetRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StampSheet:    core.CastString(data["stampSheet"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p RunStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"stampSheet":    p.StampSheet,
		"keyId":         p.KeyId,
	}
}

func (p RunStampSheetRequest) Pointer() *RunStampSheetRequest {
	return &p
}

type RunStampSheetExpressRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampSheetExpressRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampSheetExpressRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampSheetExpressRequest{}
	} else {
		*p = RunStampSheetExpressRequest{}
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

func NewRunStampSheetExpressRequestFromJson(data string) (RunStampSheetExpressRequest, error) {
	req := RunStampSheetExpressRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampSheetExpressRequest{}, err
	}
	return req, nil
}

func NewRunStampSheetExpressRequestFromDict(data map[string]interface{}) RunStampSheetExpressRequest {
	return RunStampSheetExpressRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		StampSheet:    core.CastString(data["stampSheet"]),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p RunStampSheetExpressRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"stampSheet":    p.StampSheet,
		"keyId":         p.KeyId,
	}
}

func (p RunStampSheetExpressRequest) Pointer() *RunStampSheetExpressRequest {
	return &p
}

type RunStampTaskWithoutNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampTaskWithoutNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampTaskWithoutNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampTaskWithoutNamespaceRequest{}
	} else {
		*p = RunStampTaskWithoutNamespaceRequest{}
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

func NewRunStampTaskWithoutNamespaceRequestFromJson(data string) (RunStampTaskWithoutNamespaceRequest, error) {
	req := RunStampTaskWithoutNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampTaskWithoutNamespaceRequest{}, err
	}
	return req, nil
}

func NewRunStampTaskWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampTaskWithoutNamespaceRequest {
	return RunStampTaskWithoutNamespaceRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p RunStampTaskWithoutNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p RunStampTaskWithoutNamespaceRequest) Pointer() *RunStampTaskWithoutNamespaceRequest {
	return &p
}

type RunStampSheetWithoutNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampSheetWithoutNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampSheetWithoutNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampSheetWithoutNamespaceRequest{}
	} else {
		*p = RunStampSheetWithoutNamespaceRequest{}
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

func NewRunStampSheetWithoutNamespaceRequestFromJson(data string) (RunStampSheetWithoutNamespaceRequest, error) {
	req := RunStampSheetWithoutNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampSheetWithoutNamespaceRequest{}, err
	}
	return req, nil
}

func NewRunStampSheetWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampSheetWithoutNamespaceRequest {
	return RunStampSheetWithoutNamespaceRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RunStampSheetWithoutNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RunStampSheetWithoutNamespaceRequest) Pointer() *RunStampSheetWithoutNamespaceRequest {
	return &p
}

type RunStampSheetExpressWithoutNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *RunStampSheetExpressWithoutNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RunStampSheetExpressWithoutNamespaceRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RunStampSheetExpressWithoutNamespaceRequest{}
	} else {
		*p = RunStampSheetExpressWithoutNamespaceRequest{}
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

func NewRunStampSheetExpressWithoutNamespaceRequestFromJson(data string) (RunStampSheetExpressWithoutNamespaceRequest, error) {
	req := RunStampSheetExpressWithoutNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RunStampSheetExpressWithoutNamespaceRequest{}, err
	}
	return req, nil
}

func NewRunStampSheetExpressWithoutNamespaceRequestFromDict(data map[string]interface{}) RunStampSheetExpressWithoutNamespaceRequest {
	return RunStampSheetExpressWithoutNamespaceRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RunStampSheetExpressWithoutNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RunStampSheetExpressWithoutNamespaceRequest) Pointer() *RunStampSheetExpressWithoutNamespaceRequest {
	return &p
}

type GetStampSheetResultRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	TransactionId   *string `json:"transactionId"`
}

func (p *GetStampSheetResultRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStampSheetResultRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStampSheetResultRequest{}
	} else {
		*p = GetStampSheetResultRequest{}
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
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TransactionId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TransactionId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TransactionId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TransactionId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TransactionId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TransactionId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetStampSheetResultRequestFromJson(data string) (GetStampSheetResultRequest, error) {
	req := GetStampSheetResultRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStampSheetResultRequest{}, err
	}
	return req, nil
}

func NewGetStampSheetResultRequestFromDict(data map[string]interface{}) GetStampSheetResultRequest {
	return GetStampSheetResultRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		TransactionId: core.CastString(data["transactionId"]),
	}
}

func (p GetStampSheetResultRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"transactionId": p.TransactionId,
	}
}

func (p GetStampSheetResultRequest) Pointer() *GetStampSheetResultRequest {
	return &p
}

type GetStampSheetResultByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	TransactionId   *string `json:"transactionId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetStampSheetResultByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStampSheetResultByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetStampSheetResultByUserIdRequest{}
	} else {
		*p = GetStampSheetResultByUserIdRequest{}
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
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TransactionId = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TransactionId = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TransactionId = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TransactionId = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TransactionId = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TransactionId); err != nil {
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

func NewGetStampSheetResultByUserIdRequestFromJson(data string) (GetStampSheetResultByUserIdRequest, error) {
	req := GetStampSheetResultByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStampSheetResultByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetStampSheetResultByUserIdRequestFromDict(data map[string]interface{}) GetStampSheetResultByUserIdRequest {
	return GetStampSheetResultByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TransactionId:   core.CastString(data["transactionId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetStampSheetResultByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"transactionId":   p.TransactionId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetStampSheetResultByUserIdRequest) Pointer() *GetStampSheetResultByUserIdRequest {
	return &p
}
