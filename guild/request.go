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

package guild

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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	ChangeNotification         *NotificationSetting `json:"changeNotification"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
	CreateGuildScript          *ScriptSetting       `json:"createGuildScript"`
	UpdateGuildScript          *ScriptSetting       `json:"updateGuildScript"`
	JoinGuildScript            *ScriptSetting       `json:"joinGuildScript"`
	LeaveGuildScript           *ScriptSetting       `json:"leaveGuildScript"`
	ChangeRoleScript           *ScriptSetting       `json:"changeRoleScript"`
	LogSetting                 *LogSetting          `json:"logSetting"`
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
		if v, ok := d["changeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeNotification)
		}
		if v, ok := d["joinNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinNotification)
		}
		if v, ok := d["leaveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveNotification)
		}
		if v, ok := d["changeMemberNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeMemberNotification)
		}
		if v, ok := d["receiveRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveRequestNotification)
		}
		if v, ok := d["removeRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RemoveRequestNotification)
		}
		if v, ok := d["createGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateGuildScript)
		}
		if v, ok := d["updateGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdateGuildScript)
		}
		if v, ok := d["joinGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinGuildScript)
		}
		if v, ok := d["leaveGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveGuildScript)
		}
		if v, ok := d["changeRoleScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRoleScript)
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
		ChangeNotification: func() *NotificationSetting {
			v, ok := data["changeNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeNotification"])).Pointer()
		}(),
		JoinNotification: func() *NotificationSetting {
			v, ok := data["joinNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer()
		}(),
		LeaveNotification: func() *NotificationSetting {
			v, ok := data["leaveNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer()
		}(),
		ChangeMemberNotification: func() *NotificationSetting {
			v, ok := data["changeMemberNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeMemberNotification"])).Pointer()
		}(),
		ReceiveRequestNotification: func() *NotificationSetting {
			v, ok := data["receiveRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer()
		}(),
		RemoveRequestNotification: func() *NotificationSetting {
			v, ok := data["removeRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["removeRequestNotification"])).Pointer()
		}(),
		CreateGuildScript: func() *ScriptSetting {
			v, ok := data["createGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["createGuildScript"])).Pointer()
		}(),
		UpdateGuildScript: func() *ScriptSetting {
			v, ok := data["updateGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["updateGuildScript"])).Pointer()
		}(),
		JoinGuildScript: func() *ScriptSetting {
			v, ok := data["joinGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["joinGuildScript"])).Pointer()
		}(),
		LeaveGuildScript: func() *ScriptSetting {
			v, ok := data["leaveGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["leaveGuildScript"])).Pointer()
		}(),
		ChangeRoleScript: func() *ScriptSetting {
			v, ok := data["changeRoleScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRoleScript"])).Pointer()
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
		"changeNotification": func() map[string]interface{} {
			if p.ChangeNotification == nil {
				return nil
			}
			return p.ChangeNotification.ToDict()
		}(),
		"joinNotification": func() map[string]interface{} {
			if p.JoinNotification == nil {
				return nil
			}
			return p.JoinNotification.ToDict()
		}(),
		"leaveNotification": func() map[string]interface{} {
			if p.LeaveNotification == nil {
				return nil
			}
			return p.LeaveNotification.ToDict()
		}(),
		"changeMemberNotification": func() map[string]interface{} {
			if p.ChangeMemberNotification == nil {
				return nil
			}
			return p.ChangeMemberNotification.ToDict()
		}(),
		"receiveRequestNotification": func() map[string]interface{} {
			if p.ReceiveRequestNotification == nil {
				return nil
			}
			return p.ReceiveRequestNotification.ToDict()
		}(),
		"removeRequestNotification": func() map[string]interface{} {
			if p.RemoveRequestNotification == nil {
				return nil
			}
			return p.RemoveRequestNotification.ToDict()
		}(),
		"createGuildScript": func() map[string]interface{} {
			if p.CreateGuildScript == nil {
				return nil
			}
			return p.CreateGuildScript.ToDict()
		}(),
		"updateGuildScript": func() map[string]interface{} {
			if p.UpdateGuildScript == nil {
				return nil
			}
			return p.UpdateGuildScript.ToDict()
		}(),
		"joinGuildScript": func() map[string]interface{} {
			if p.JoinGuildScript == nil {
				return nil
			}
			return p.JoinGuildScript.ToDict()
		}(),
		"leaveGuildScript": func() map[string]interface{} {
			if p.LeaveGuildScript == nil {
				return nil
			}
			return p.LeaveGuildScript.ToDict()
		}(),
		"changeRoleScript": func() map[string]interface{} {
			if p.ChangeRoleScript == nil {
				return nil
			}
			return p.ChangeRoleScript.ToDict()
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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	NamespaceName              *string              `json:"namespaceName"`
	Description                *string              `json:"description"`
	ChangeNotification         *NotificationSetting `json:"changeNotification"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
	CreateGuildScript          *ScriptSetting       `json:"createGuildScript"`
	UpdateGuildScript          *ScriptSetting       `json:"updateGuildScript"`
	JoinGuildScript            *ScriptSetting       `json:"joinGuildScript"`
	LeaveGuildScript           *ScriptSetting       `json:"leaveGuildScript"`
	ChangeRoleScript           *ScriptSetting       `json:"changeRoleScript"`
	LogSetting                 *LogSetting          `json:"logSetting"`
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
		if v, ok := d["changeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeNotification)
		}
		if v, ok := d["joinNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinNotification)
		}
		if v, ok := d["leaveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveNotification)
		}
		if v, ok := d["changeMemberNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeMemberNotification)
		}
		if v, ok := d["receiveRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveRequestNotification)
		}
		if v, ok := d["removeRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RemoveRequestNotification)
		}
		if v, ok := d["createGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreateGuildScript)
		}
		if v, ok := d["updateGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdateGuildScript)
		}
		if v, ok := d["joinGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinGuildScript)
		}
		if v, ok := d["leaveGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveGuildScript)
		}
		if v, ok := d["changeRoleScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRoleScript)
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
		ChangeNotification: func() *NotificationSetting {
			v, ok := data["changeNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeNotification"])).Pointer()
		}(),
		JoinNotification: func() *NotificationSetting {
			v, ok := data["joinNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer()
		}(),
		LeaveNotification: func() *NotificationSetting {
			v, ok := data["leaveNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer()
		}(),
		ChangeMemberNotification: func() *NotificationSetting {
			v, ok := data["changeMemberNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeMemberNotification"])).Pointer()
		}(),
		ReceiveRequestNotification: func() *NotificationSetting {
			v, ok := data["receiveRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer()
		}(),
		RemoveRequestNotification: func() *NotificationSetting {
			v, ok := data["removeRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["removeRequestNotification"])).Pointer()
		}(),
		CreateGuildScript: func() *ScriptSetting {
			v, ok := data["createGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["createGuildScript"])).Pointer()
		}(),
		UpdateGuildScript: func() *ScriptSetting {
			v, ok := data["updateGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["updateGuildScript"])).Pointer()
		}(),
		JoinGuildScript: func() *ScriptSetting {
			v, ok := data["joinGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["joinGuildScript"])).Pointer()
		}(),
		LeaveGuildScript: func() *ScriptSetting {
			v, ok := data["leaveGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["leaveGuildScript"])).Pointer()
		}(),
		ChangeRoleScript: func() *ScriptSetting {
			v, ok := data["changeRoleScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRoleScript"])).Pointer()
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
		"changeNotification": func() map[string]interface{} {
			if p.ChangeNotification == nil {
				return nil
			}
			return p.ChangeNotification.ToDict()
		}(),
		"joinNotification": func() map[string]interface{} {
			if p.JoinNotification == nil {
				return nil
			}
			return p.JoinNotification.ToDict()
		}(),
		"leaveNotification": func() map[string]interface{} {
			if p.LeaveNotification == nil {
				return nil
			}
			return p.LeaveNotification.ToDict()
		}(),
		"changeMemberNotification": func() map[string]interface{} {
			if p.ChangeMemberNotification == nil {
				return nil
			}
			return p.ChangeMemberNotification.ToDict()
		}(),
		"receiveRequestNotification": func() map[string]interface{} {
			if p.ReceiveRequestNotification == nil {
				return nil
			}
			return p.ReceiveRequestNotification.ToDict()
		}(),
		"removeRequestNotification": func() map[string]interface{} {
			if p.RemoveRequestNotification == nil {
				return nil
			}
			return p.RemoveRequestNotification.ToDict()
		}(),
		"createGuildScript": func() map[string]interface{} {
			if p.CreateGuildScript == nil {
				return nil
			}
			return p.CreateGuildScript.ToDict()
		}(),
		"updateGuildScript": func() map[string]interface{} {
			if p.UpdateGuildScript == nil {
				return nil
			}
			return p.UpdateGuildScript.ToDict()
		}(),
		"joinGuildScript": func() map[string]interface{} {
			if p.JoinGuildScript == nil {
				return nil
			}
			return p.JoinGuildScript.ToDict()
		}(),
		"leaveGuildScript": func() map[string]interface{} {
			if p.LeaveGuildScript == nil {
				return nil
			}
			return p.LeaveGuildScript.ToDict()
		}(),
		"changeRoleScript": func() map[string]interface{} {
			if p.ChangeRoleScript == nil {
				return nil
			}
			return p.ChangeRoleScript.ToDict()
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

type DescribeGuildModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeGuildModelMastersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGuildModelMastersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGuildModelMastersRequest{}
	} else {
		*p = DescribeGuildModelMastersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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

func NewDescribeGuildModelMastersRequestFromJson(data string) (DescribeGuildModelMastersRequest, error) {
	req := DescribeGuildModelMastersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGuildModelMastersRequest{}, err
	}
	return req, nil
}

func NewDescribeGuildModelMastersRequestFromDict(data map[string]interface{}) DescribeGuildModelMastersRequest {
	return DescribeGuildModelMastersRequest{
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

func (p DescribeGuildModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGuildModelMastersRequest) Pointer() *DescribeGuildModelMastersRequest {
	return &p
}

type CreateGuildModelMasterRequest struct {
	SourceRequestId               *string     `json:"sourceRequestId"`
	RequestId                     *string     `json:"requestId"`
	ContextStack                  *string     `json:"contextStack"`
	NamespaceName                 *string     `json:"namespaceName"`
	Name                          *string     `json:"name"`
	Description                   *string     `json:"description"`
	Metadata                      *string     `json:"metadata"`
	DefaultMaximumMemberCount     *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount            *int32      `json:"maximumMemberCount"`
	InactivityPeriodDays          *int32      `json:"inactivityPeriodDays"`
	Roles                         []RoleModel `json:"roles"`
	GuildMasterRole               *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole        *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes         *int32      `json:"rejoinCoolTimeMinutes"`
	MaxConcurrentJoinGuilds       *int32      `json:"maxConcurrentJoinGuilds"`
	MaxConcurrentGuildMasterCount *int32      `json:"maxConcurrentGuildMasterCount"`
}

func (p *CreateGuildModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGuildModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGuildModelMasterRequest{}
	} else {
		*p = CreateGuildModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["defaultMaximumMemberCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultMaximumMemberCount)
		}
		if v, ok := d["maximumMemberCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumMemberCount)
		}
		if v, ok := d["inactivityPeriodDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InactivityPeriodDays)
		}
		if v, ok := d["roles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Roles)
		}
		if v, ok := d["guildMasterRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMasterRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMasterRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMasterRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMasterRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMasterRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMasterRole)
				}
			}
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
				}
			}
		}
		if v, ok := d["rejoinCoolTimeMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RejoinCoolTimeMinutes)
		}
		if v, ok := d["maxConcurrentJoinGuilds"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxConcurrentJoinGuilds)
		}
		if v, ok := d["maxConcurrentGuildMasterCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxConcurrentGuildMasterCount)
		}
	}
	return nil
}

func NewCreateGuildModelMasterRequestFromJson(data string) (CreateGuildModelMasterRequest, error) {
	req := CreateGuildModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGuildModelMasterRequest{}, err
	}
	return req, nil
}

func NewCreateGuildModelMasterRequestFromDict(data map[string]interface{}) CreateGuildModelMasterRequest {
	return CreateGuildModelMasterRequest{
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
		DefaultMaximumMemberCount: func() *int32 {
			v, ok := data["defaultMaximumMemberCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["defaultMaximumMemberCount"])
		}(),
		MaximumMemberCount: func() *int32 {
			v, ok := data["maximumMemberCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumMemberCount"])
		}(),
		InactivityPeriodDays: func() *int32 {
			v, ok := data["inactivityPeriodDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["inactivityPeriodDays"])
		}(),
		Roles: func() []RoleModel {
			if data["roles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["roles"]))
		}(),
		GuildMasterRole: func() *string {
			v, ok := data["guildMasterRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMasterRole"])
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
		}(),
		RejoinCoolTimeMinutes: func() *int32 {
			v, ok := data["rejoinCoolTimeMinutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rejoinCoolTimeMinutes"])
		}(),
		MaxConcurrentJoinGuilds: func() *int32 {
			v, ok := data["maxConcurrentJoinGuilds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxConcurrentJoinGuilds"])
		}(),
		MaxConcurrentGuildMasterCount: func() *int32 {
			v, ok := data["maxConcurrentGuildMasterCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxConcurrentGuildMasterCount"])
		}(),
	}
}

func (p CreateGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"name":                      p.Name,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"defaultMaximumMemberCount": p.DefaultMaximumMemberCount,
		"maximumMemberCount":        p.MaximumMemberCount,
		"inactivityPeriodDays":      p.InactivityPeriodDays,
		"roles": CastRoleModelsFromDict(
			p.Roles,
		),
		"guildMasterRole":               p.GuildMasterRole,
		"guildMemberDefaultRole":        p.GuildMemberDefaultRole,
		"rejoinCoolTimeMinutes":         p.RejoinCoolTimeMinutes,
		"maxConcurrentJoinGuilds":       p.MaxConcurrentJoinGuilds,
		"maxConcurrentGuildMasterCount": p.MaxConcurrentGuildMasterCount,
	}
}

func (p CreateGuildModelMasterRequest) Pointer() *CreateGuildModelMasterRequest {
	return &p
}

type GetGuildModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func (p *GetGuildModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGuildModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGuildModelMasterRequest{}
	} else {
		*p = GetGuildModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
	}
	return nil
}

func NewGetGuildModelMasterRequestFromJson(data string) (GetGuildModelMasterRequest, error) {
	req := GetGuildModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGuildModelMasterRequest{}, err
	}
	return req, nil
}

func NewGetGuildModelMasterRequestFromDict(data map[string]interface{}) GetGuildModelMasterRequest {
	return GetGuildModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
	}
}

func (p GetGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p GetGuildModelMasterRequest) Pointer() *GetGuildModelMasterRequest {
	return &p
}

type UpdateGuildModelMasterRequest struct {
	SourceRequestId               *string     `json:"sourceRequestId"`
	RequestId                     *string     `json:"requestId"`
	ContextStack                  *string     `json:"contextStack"`
	NamespaceName                 *string     `json:"namespaceName"`
	GuildModelName                *string     `json:"guildModelName"`
	Description                   *string     `json:"description"`
	Metadata                      *string     `json:"metadata"`
	DefaultMaximumMemberCount     *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount            *int32      `json:"maximumMemberCount"`
	InactivityPeriodDays          *int32      `json:"inactivityPeriodDays"`
	Roles                         []RoleModel `json:"roles"`
	GuildMasterRole               *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole        *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes         *int32      `json:"rejoinCoolTimeMinutes"`
	MaxConcurrentJoinGuilds       *int32      `json:"maxConcurrentJoinGuilds"`
	MaxConcurrentGuildMasterCount *int32      `json:"maxConcurrentGuildMasterCount"`
}

func (p *UpdateGuildModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGuildModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGuildModelMasterRequest{}
	} else {
		*p = UpdateGuildModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["defaultMaximumMemberCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DefaultMaximumMemberCount)
		}
		if v, ok := d["maximumMemberCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumMemberCount)
		}
		if v, ok := d["inactivityPeriodDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InactivityPeriodDays)
		}
		if v, ok := d["roles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Roles)
		}
		if v, ok := d["guildMasterRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMasterRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMasterRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMasterRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMasterRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMasterRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMasterRole)
				}
			}
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
				}
			}
		}
		if v, ok := d["rejoinCoolTimeMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RejoinCoolTimeMinutes)
		}
		if v, ok := d["maxConcurrentJoinGuilds"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxConcurrentJoinGuilds)
		}
		if v, ok := d["maxConcurrentGuildMasterCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxConcurrentGuildMasterCount)
		}
	}
	return nil
}

func NewUpdateGuildModelMasterRequestFromJson(data string) (UpdateGuildModelMasterRequest, error) {
	req := UpdateGuildModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGuildModelMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateGuildModelMasterRequestFromDict(data map[string]interface{}) UpdateGuildModelMasterRequest {
	return UpdateGuildModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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
		DefaultMaximumMemberCount: func() *int32 {
			v, ok := data["defaultMaximumMemberCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["defaultMaximumMemberCount"])
		}(),
		MaximumMemberCount: func() *int32 {
			v, ok := data["maximumMemberCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumMemberCount"])
		}(),
		InactivityPeriodDays: func() *int32 {
			v, ok := data["inactivityPeriodDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["inactivityPeriodDays"])
		}(),
		Roles: func() []RoleModel {
			if data["roles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["roles"]))
		}(),
		GuildMasterRole: func() *string {
			v, ok := data["guildMasterRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMasterRole"])
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
		}(),
		RejoinCoolTimeMinutes: func() *int32 {
			v, ok := data["rejoinCoolTimeMinutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rejoinCoolTimeMinutes"])
		}(),
		MaxConcurrentJoinGuilds: func() *int32 {
			v, ok := data["maxConcurrentJoinGuilds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxConcurrentJoinGuilds"])
		}(),
		MaxConcurrentGuildMasterCount: func() *int32 {
			v, ok := data["maxConcurrentGuildMasterCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxConcurrentGuildMasterCount"])
		}(),
	}
}

func (p UpdateGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"guildModelName":            p.GuildModelName,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"defaultMaximumMemberCount": p.DefaultMaximumMemberCount,
		"maximumMemberCount":        p.MaximumMemberCount,
		"inactivityPeriodDays":      p.InactivityPeriodDays,
		"roles": CastRoleModelsFromDict(
			p.Roles,
		),
		"guildMasterRole":               p.GuildMasterRole,
		"guildMemberDefaultRole":        p.GuildMemberDefaultRole,
		"rejoinCoolTimeMinutes":         p.RejoinCoolTimeMinutes,
		"maxConcurrentJoinGuilds":       p.MaxConcurrentJoinGuilds,
		"maxConcurrentGuildMasterCount": p.MaxConcurrentGuildMasterCount,
	}
}

func (p UpdateGuildModelMasterRequest) Pointer() *UpdateGuildModelMasterRequest {
	return &p
}

type DeleteGuildModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func (p *DeleteGuildModelMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGuildModelMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGuildModelMasterRequest{}
	} else {
		*p = DeleteGuildModelMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
	}
	return nil
}

func NewDeleteGuildModelMasterRequestFromJson(data string) (DeleteGuildModelMasterRequest, error) {
	req := DeleteGuildModelMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGuildModelMasterRequest{}, err
	}
	return req, nil
}

func NewDeleteGuildModelMasterRequestFromDict(data map[string]interface{}) DeleteGuildModelMasterRequest {
	return DeleteGuildModelMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
	}
}

func (p DeleteGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p DeleteGuildModelMasterRequest) Pointer() *DeleteGuildModelMasterRequest {
	return &p
}

type DescribeGuildModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DescribeGuildModelsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeGuildModelsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeGuildModelsRequest{}
	} else {
		*p = DescribeGuildModelsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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

func NewDescribeGuildModelsRequestFromJson(data string) (DescribeGuildModelsRequest, error) {
	req := DescribeGuildModelsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeGuildModelsRequest{}, err
	}
	return req, nil
}

func NewDescribeGuildModelsRequestFromDict(data map[string]interface{}) DescribeGuildModelsRequest {
	return DescribeGuildModelsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p DescribeGuildModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGuildModelsRequest) Pointer() *DescribeGuildModelsRequest {
	return &p
}

type GetGuildModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func (p *GetGuildModelRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGuildModelRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGuildModelRequest{}
	} else {
		*p = GetGuildModelRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
	}
	return nil
}

func NewGetGuildModelRequestFromJson(data string) (GetGuildModelRequest, error) {
	req := GetGuildModelRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGuildModelRequest{}, err
	}
	return req, nil
}

func NewGetGuildModelRequestFromDict(data map[string]interface{}) GetGuildModelRequest {
	return GetGuildModelRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
	}
}

func (p GetGuildModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p GetGuildModelRequest) Pointer() *GetGuildModelRequest {
	return &p
}

type SearchGuildsRequest struct {
	SourceRequestId         *string   `json:"sourceRequestId"`
	RequestId               *string   `json:"requestId"`
	ContextStack            *string   `json:"contextStack"`
	DuplicationAvoider      *string   `json:"duplicationAvoider"`
	NamespaceName           *string   `json:"namespaceName"`
	GuildModelName          *string   `json:"guildModelName"`
	AccessToken             *string   `json:"accessToken"`
	DisplayName             *string   `json:"displayName"`
	Attributes1             []*int32  `json:"attributes1"`
	Attributes2             []*int32  `json:"attributes2"`
	Attributes3             []*int32  `json:"attributes3"`
	Attributes4             []*int32  `json:"attributes4"`
	Attributes5             []*int32  `json:"attributes5"`
	JoinPolicies            []*string `json:"joinPolicies"`
	IncludeFullMembersGuild *bool     `json:"includeFullMembersGuild"`
	OrderBy                 *string   `json:"orderBy"`
	PageToken               *string   `json:"pageToken"`
	Limit                   *int32    `json:"limit"`
}

func (p *SearchGuildsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SearchGuildsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SearchGuildsRequest{}
	} else {
		*p = SearchGuildsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attributes1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes1)
		}
		if v, ok := d["attributes2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes2)
		}
		if v, ok := d["attributes3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes3)
		}
		if v, ok := d["attributes4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes4)
		}
		if v, ok := d["attributes5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes5)
		}
		if v, ok := d["joinPolicies"]; ok && v != nil {
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
				p.JoinPolicies = l
			}
		}
		if v, ok := d["includeFullMembersGuild"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IncludeFullMembersGuild)
		}
		if v, ok := d["orderBy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderBy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderBy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderBy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderBy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderBy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderBy)
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

func NewSearchGuildsRequestFromJson(data string) (SearchGuildsRequest, error) {
	req := SearchGuildsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SearchGuildsRequest{}, err
	}
	return req, nil
}

func NewSearchGuildsRequestFromDict(data map[string]interface{}) SearchGuildsRequest {
	return SearchGuildsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attributes1: func() []*int32 {
			v, ok := data["attributes1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes2: func() []*int32 {
			v, ok := data["attributes2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes3: func() []*int32 {
			v, ok := data["attributes3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes4: func() []*int32 {
			v, ok := data["attributes4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes5: func() []*int32 {
			v, ok := data["attributes5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		JoinPolicies: func() []*string {
			v, ok := data["joinPolicies"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		IncludeFullMembersGuild: func() *bool {
			v, ok := data["includeFullMembersGuild"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["includeFullMembersGuild"])
		}(),
		OrderBy: func() *string {
			v, ok := data["orderBy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderBy"])
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

func (p SearchGuildsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"displayName":    p.DisplayName,
		"attributes1": core.CastInt32sFromDict(
			p.Attributes1,
		),
		"attributes2": core.CastInt32sFromDict(
			p.Attributes2,
		),
		"attributes3": core.CastInt32sFromDict(
			p.Attributes3,
		),
		"attributes4": core.CastInt32sFromDict(
			p.Attributes4,
		),
		"attributes5": core.CastInt32sFromDict(
			p.Attributes5,
		),
		"joinPolicies": core.CastStringsFromDict(
			p.JoinPolicies,
		),
		"includeFullMembersGuild": p.IncludeFullMembersGuild,
		"orderBy":                 p.OrderBy,
		"pageToken":               p.PageToken,
		"limit":                   p.Limit,
	}
}

func (p SearchGuildsRequest) Pointer() *SearchGuildsRequest {
	return &p
}

type SearchGuildsByUserIdRequest struct {
	SourceRequestId         *string   `json:"sourceRequestId"`
	RequestId               *string   `json:"requestId"`
	ContextStack            *string   `json:"contextStack"`
	DuplicationAvoider      *string   `json:"duplicationAvoider"`
	NamespaceName           *string   `json:"namespaceName"`
	GuildModelName          *string   `json:"guildModelName"`
	UserId                  *string   `json:"userId"`
	DisplayName             *string   `json:"displayName"`
	Attributes1             []*int32  `json:"attributes1"`
	Attributes2             []*int32  `json:"attributes2"`
	Attributes3             []*int32  `json:"attributes3"`
	Attributes4             []*int32  `json:"attributes4"`
	Attributes5             []*int32  `json:"attributes5"`
	JoinPolicies            []*string `json:"joinPolicies"`
	IncludeFullMembersGuild *bool     `json:"includeFullMembersGuild"`
	OrderBy                 *string   `json:"orderBy"`
	PageToken               *string   `json:"pageToken"`
	Limit                   *int32    `json:"limit"`
	TimeOffsetToken         *string   `json:"timeOffsetToken"`
}

func (p *SearchGuildsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SearchGuildsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SearchGuildsByUserIdRequest{}
	} else {
		*p = SearchGuildsByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attributes1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes1)
		}
		if v, ok := d["attributes2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes2)
		}
		if v, ok := d["attributes3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes3)
		}
		if v, ok := d["attributes4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes4)
		}
		if v, ok := d["attributes5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes5)
		}
		if v, ok := d["joinPolicies"]; ok && v != nil {
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
				p.JoinPolicies = l
			}
		}
		if v, ok := d["includeFullMembersGuild"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IncludeFullMembersGuild)
		}
		if v, ok := d["orderBy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderBy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderBy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderBy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderBy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderBy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderBy)
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

func NewSearchGuildsByUserIdRequestFromJson(data string) (SearchGuildsByUserIdRequest, error) {
	req := SearchGuildsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SearchGuildsByUserIdRequest{}, err
	}
	return req, nil
}

func NewSearchGuildsByUserIdRequestFromDict(data map[string]interface{}) SearchGuildsByUserIdRequest {
	return SearchGuildsByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attributes1: func() []*int32 {
			v, ok := data["attributes1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes2: func() []*int32 {
			v, ok := data["attributes2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes3: func() []*int32 {
			v, ok := data["attributes3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes4: func() []*int32 {
			v, ok := data["attributes4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		Attributes5: func() []*int32 {
			v, ok := data["attributes5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
		JoinPolicies: func() []*string {
			v, ok := data["joinPolicies"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		IncludeFullMembersGuild: func() *bool {
			v, ok := data["includeFullMembersGuild"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["includeFullMembersGuild"])
		}(),
		OrderBy: func() *string {
			v, ok := data["orderBy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderBy"])
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

func (p SearchGuildsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"userId":         p.UserId,
		"displayName":    p.DisplayName,
		"attributes1": core.CastInt32sFromDict(
			p.Attributes1,
		),
		"attributes2": core.CastInt32sFromDict(
			p.Attributes2,
		),
		"attributes3": core.CastInt32sFromDict(
			p.Attributes3,
		),
		"attributes4": core.CastInt32sFromDict(
			p.Attributes4,
		),
		"attributes5": core.CastInt32sFromDict(
			p.Attributes5,
		),
		"joinPolicies": core.CastStringsFromDict(
			p.JoinPolicies,
		),
		"includeFullMembersGuild": p.IncludeFullMembersGuild,
		"orderBy":                 p.OrderBy,
		"pageToken":               p.PageToken,
		"limit":                   p.Limit,
		"timeOffsetToken":         p.TimeOffsetToken,
	}
}

func (p SearchGuildsByUserIdRequest) Pointer() *SearchGuildsByUserIdRequest {
	return &p
}

type CreateGuildRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	AccessToken            *string     `json:"accessToken"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	Attribute1             *int32      `json:"attribute1"`
	Attribute2             *int32      `json:"attribute2"`
	Attribute3             *int32      `json:"attribute3"`
	Attribute4             *int32      `json:"attribute4"`
	Attribute5             *int32      `json:"attribute5"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func (p *CreateGuildRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGuildRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGuildRequest{}
	} else {
		*p = CreateGuildRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attribute1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute1)
		}
		if v, ok := d["attribute2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute2)
		}
		if v, ok := d["attribute3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute3)
		}
		if v, ok := d["attribute4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute4)
		}
		if v, ok := d["attribute5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute5)
		}
		if v, ok := d["joinPolicy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinPolicy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinPolicy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinPolicy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinPolicy)
				}
			}
		}
		if v, ok := d["customRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CustomRoles)
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
				}
			}
		}
	}
	return nil
}

func NewCreateGuildRequestFromJson(data string) (CreateGuildRequest, error) {
	req := CreateGuildRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGuildRequest{}, err
	}
	return req, nil
}

func NewCreateGuildRequestFromDict(data map[string]interface{}) CreateGuildRequest {
	return CreateGuildRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attribute1: func() *int32 {
			v, ok := data["attribute1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute1"])
		}(),
		Attribute2: func() *int32 {
			v, ok := data["attribute2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute2"])
		}(),
		Attribute3: func() *int32 {
			v, ok := data["attribute3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute3"])
		}(),
		Attribute4: func() *int32 {
			v, ok := data["attribute4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute4"])
		}(),
		Attribute5: func() *int32 {
			v, ok := data["attribute5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute5"])
		}(),
		JoinPolicy: func() *string {
			v, ok := data["joinPolicy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinPolicy"])
		}(),
		CustomRoles: func() []RoleModel {
			if data["customRoles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["customRoles"]))
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
		}(),
	}
}

func (p CreateGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"displayName":    p.DisplayName,
		"attribute1":     p.Attribute1,
		"attribute2":     p.Attribute2,
		"attribute3":     p.Attribute3,
		"attribute4":     p.Attribute4,
		"attribute5":     p.Attribute5,
		"joinPolicy":     p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p CreateGuildRequest) Pointer() *CreateGuildRequest {
	return &p
}

type CreateGuildByUserIdRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	UserId                 *string     `json:"userId"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	Attribute1             *int32      `json:"attribute1"`
	Attribute2             *int32      `json:"attribute2"`
	Attribute3             *int32      `json:"attribute3"`
	Attribute4             *int32      `json:"attribute4"`
	Attribute5             *int32      `json:"attribute5"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
	TimeOffsetToken        *string     `json:"timeOffsetToken"`
}

func (p *CreateGuildByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateGuildByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CreateGuildByUserIdRequest{}
	} else {
		*p = CreateGuildByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attribute1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute1)
		}
		if v, ok := d["attribute2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute2)
		}
		if v, ok := d["attribute3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute3)
		}
		if v, ok := d["attribute4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute4)
		}
		if v, ok := d["attribute5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute5)
		}
		if v, ok := d["joinPolicy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinPolicy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinPolicy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinPolicy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinPolicy)
				}
			}
		}
		if v, ok := d["customRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CustomRoles)
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
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

func NewCreateGuildByUserIdRequestFromJson(data string) (CreateGuildByUserIdRequest, error) {
	req := CreateGuildByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateGuildByUserIdRequest{}, err
	}
	return req, nil
}

func NewCreateGuildByUserIdRequestFromDict(data map[string]interface{}) CreateGuildByUserIdRequest {
	return CreateGuildByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attribute1: func() *int32 {
			v, ok := data["attribute1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute1"])
		}(),
		Attribute2: func() *int32 {
			v, ok := data["attribute2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute2"])
		}(),
		Attribute3: func() *int32 {
			v, ok := data["attribute3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute3"])
		}(),
		Attribute4: func() *int32 {
			v, ok := data["attribute4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute4"])
		}(),
		Attribute5: func() *int32 {
			v, ok := data["attribute5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute5"])
		}(),
		JoinPolicy: func() *string {
			v, ok := data["joinPolicy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinPolicy"])
		}(),
		CustomRoles: func() []RoleModel {
			if data["customRoles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["customRoles"]))
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
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

func (p CreateGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"userId":         p.UserId,
		"guildModelName": p.GuildModelName,
		"displayName":    p.DisplayName,
		"attribute1":     p.Attribute1,
		"attribute2":     p.Attribute2,
		"attribute3":     p.Attribute3,
		"attribute4":     p.Attribute4,
		"attribute5":     p.Attribute5,
		"joinPolicy":     p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
		"timeOffsetToken":        p.TimeOffsetToken,
	}
}

func (p CreateGuildByUserIdRequest) Pointer() *CreateGuildByUserIdRequest {
	return &p
}

type GetGuildRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
}

func (p *GetGuildRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGuildRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGuildRequest{}
	} else {
		*p = GetGuildRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewGetGuildRequestFromJson(data string) (GetGuildRequest, error) {
	req := GetGuildRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGuildRequest{}, err
	}
	return req, nil
}

func NewGetGuildRequestFromDict(data map[string]interface{}) GetGuildRequest {
	return GetGuildRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p GetGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p GetGuildRequest) Pointer() *GetGuildRequest {
	return &p
}

type GetGuildByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetGuildByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetGuildByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetGuildByUserIdRequest{}
	} else {
		*p = GetGuildByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewGetGuildByUserIdRequestFromJson(data string) (GetGuildByUserIdRequest, error) {
	req := GetGuildByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetGuildByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetGuildByUserIdRequestFromDict(data map[string]interface{}) GetGuildByUserIdRequest {
	return GetGuildByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p GetGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetGuildByUserIdRequest) Pointer() *GetGuildByUserIdRequest {
	return &p
}

type UpdateGuildRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	AccessToken            *string     `json:"accessToken"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	Attribute1             *int32      `json:"attribute1"`
	Attribute2             *int32      `json:"attribute2"`
	Attribute3             *int32      `json:"attribute3"`
	Attribute4             *int32      `json:"attribute4"`
	Attribute5             *int32      `json:"attribute5"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func (p *UpdateGuildRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGuildRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGuildRequest{}
	} else {
		*p = UpdateGuildRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attribute1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute1)
		}
		if v, ok := d["attribute2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute2)
		}
		if v, ok := d["attribute3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute3)
		}
		if v, ok := d["attribute4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute4)
		}
		if v, ok := d["attribute5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute5)
		}
		if v, ok := d["joinPolicy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinPolicy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinPolicy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinPolicy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinPolicy)
				}
			}
		}
		if v, ok := d["customRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CustomRoles)
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
				}
			}
		}
	}
	return nil
}

func NewUpdateGuildRequestFromJson(data string) (UpdateGuildRequest, error) {
	req := UpdateGuildRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGuildRequest{}, err
	}
	return req, nil
}

func NewUpdateGuildRequestFromDict(data map[string]interface{}) UpdateGuildRequest {
	return UpdateGuildRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attribute1: func() *int32 {
			v, ok := data["attribute1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute1"])
		}(),
		Attribute2: func() *int32 {
			v, ok := data["attribute2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute2"])
		}(),
		Attribute3: func() *int32 {
			v, ok := data["attribute3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute3"])
		}(),
		Attribute4: func() *int32 {
			v, ok := data["attribute4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute4"])
		}(),
		Attribute5: func() *int32 {
			v, ok := data["attribute5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute5"])
		}(),
		JoinPolicy: func() *string {
			v, ok := data["joinPolicy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinPolicy"])
		}(),
		CustomRoles: func() []RoleModel {
			if data["customRoles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["customRoles"]))
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
		}(),
	}
}

func (p UpdateGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"displayName":    p.DisplayName,
		"attribute1":     p.Attribute1,
		"attribute2":     p.Attribute2,
		"attribute3":     p.Attribute3,
		"attribute4":     p.Attribute4,
		"attribute5":     p.Attribute5,
		"joinPolicy":     p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p UpdateGuildRequest) Pointer() *UpdateGuildRequest {
	return &p
}

type UpdateGuildByGuildNameRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	GuildName              *string     `json:"guildName"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	Attribute1             *int32      `json:"attribute1"`
	Attribute2             *int32      `json:"attribute2"`
	Attribute3             *int32      `json:"attribute3"`
	Attribute4             *int32      `json:"attribute4"`
	Attribute5             *int32      `json:"attribute5"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func (p *UpdateGuildByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateGuildByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateGuildByGuildNameRequest{}
	} else {
		*p = UpdateGuildByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["attribute1"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute1)
		}
		if v, ok := d["attribute2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute2)
		}
		if v, ok := d["attribute3"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute3)
		}
		if v, ok := d["attribute4"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute4)
		}
		if v, ok := d["attribute5"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attribute5)
		}
		if v, ok := d["joinPolicy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinPolicy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinPolicy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinPolicy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinPolicy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinPolicy)
				}
			}
		}
		if v, ok := d["customRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CustomRoles)
		}
		if v, ok := d["guildMemberDefaultRole"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildMemberDefaultRole = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildMemberDefaultRole = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildMemberDefaultRole = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildMemberDefaultRole = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildMemberDefaultRole)
				}
			}
		}
	}
	return nil
}

func NewUpdateGuildByGuildNameRequestFromJson(data string) (UpdateGuildByGuildNameRequest, error) {
	req := UpdateGuildByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateGuildByGuildNameRequest{}, err
	}
	return req, nil
}

func NewUpdateGuildByGuildNameRequestFromDict(data map[string]interface{}) UpdateGuildByGuildNameRequest {
	return UpdateGuildByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Attribute1: func() *int32 {
			v, ok := data["attribute1"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute1"])
		}(),
		Attribute2: func() *int32 {
			v, ok := data["attribute2"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute2"])
		}(),
		Attribute3: func() *int32 {
			v, ok := data["attribute3"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute3"])
		}(),
		Attribute4: func() *int32 {
			v, ok := data["attribute4"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute4"])
		}(),
		Attribute5: func() *int32 {
			v, ok := data["attribute5"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["attribute5"])
		}(),
		JoinPolicy: func() *string {
			v, ok := data["joinPolicy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinPolicy"])
		}(),
		CustomRoles: func() []RoleModel {
			if data["customRoles"] == nil {
				return nil
			}
			return CastRoleModels(core.CastArray(data["customRoles"]))
		}(),
		GuildMemberDefaultRole: func() *string {
			v, ok := data["guildMemberDefaultRole"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildMemberDefaultRole"])
		}(),
	}
}

func (p UpdateGuildByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildName":      p.GuildName,
		"guildModelName": p.GuildModelName,
		"displayName":    p.DisplayName,
		"attribute1":     p.Attribute1,
		"attribute2":     p.Attribute2,
		"attribute3":     p.Attribute3,
		"attribute4":     p.Attribute4,
		"attribute5":     p.Attribute5,
		"joinPolicy":     p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p UpdateGuildByGuildNameRequest) Pointer() *UpdateGuildByGuildNameRequest {
	return &p
}

type DeleteMemberRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
}

func (p *DeleteMemberRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteMemberRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteMemberRequest{}
	} else {
		*p = DeleteMemberRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDeleteMemberRequestFromJson(data string) (DeleteMemberRequest, error) {
	req := DeleteMemberRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteMemberRequest{}, err
	}
	return req, nil
}

func NewDeleteMemberRequestFromDict(data map[string]interface{}) DeleteMemberRequest {
	return DeleteMemberRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DeleteMemberRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"targetUserId":   p.TargetUserId,
	}
}

func (p DeleteMemberRequest) Pointer() *DeleteMemberRequest {
	return &p
}

type DeleteMemberByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TargetUserId       *string `json:"targetUserId"`
}

func (p *DeleteMemberByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteMemberByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteMemberByGuildNameRequest{}
	} else {
		*p = DeleteMemberByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewDeleteMemberByGuildNameRequestFromJson(data string) (DeleteMemberByGuildNameRequest, error) {
	req := DeleteMemberByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteMemberByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDeleteMemberByGuildNameRequestFromDict(data map[string]interface{}) DeleteMemberByGuildNameRequest {
	return DeleteMemberByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p DeleteMemberByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"targetUserId":   p.TargetUserId,
	}
}

func (p DeleteMemberByGuildNameRequest) Pointer() *DeleteMemberByGuildNameRequest {
	return &p
}

type UpdateMemberRoleRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
	RoleName           *string `json:"roleName"`
}

func (p *UpdateMemberRoleRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateMemberRoleRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateMemberRoleRequest{}
	} else {
		*p = UpdateMemberRoleRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["roleName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RoleName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RoleName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RoleName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RoleName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RoleName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RoleName)
				}
			}
		}
	}
	return nil
}

func NewUpdateMemberRoleRequestFromJson(data string) (UpdateMemberRoleRequest, error) {
	req := UpdateMemberRoleRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateMemberRoleRequest{}, err
	}
	return req, nil
}

func NewUpdateMemberRoleRequestFromDict(data map[string]interface{}) UpdateMemberRoleRequest {
	return UpdateMemberRoleRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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
		RoleName: func() *string {
			v, ok := data["roleName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["roleName"])
		}(),
	}
}

func (p UpdateMemberRoleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"targetUserId":   p.TargetUserId,
		"roleName":       p.RoleName,
	}
}

func (p UpdateMemberRoleRequest) Pointer() *UpdateMemberRoleRequest {
	return &p
}

type UpdateMemberRoleByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TargetUserId       *string `json:"targetUserId"`
	RoleName           *string `json:"roleName"`
}

func (p *UpdateMemberRoleByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateMemberRoleByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateMemberRoleByGuildNameRequest{}
	} else {
		*p = UpdateMemberRoleByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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
		if v, ok := d["roleName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RoleName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RoleName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RoleName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RoleName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RoleName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RoleName)
				}
			}
		}
	}
	return nil
}

func NewUpdateMemberRoleByGuildNameRequestFromJson(data string) (UpdateMemberRoleByGuildNameRequest, error) {
	req := UpdateMemberRoleByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateMemberRoleByGuildNameRequest{}, err
	}
	return req, nil
}

func NewUpdateMemberRoleByGuildNameRequestFromDict(data map[string]interface{}) UpdateMemberRoleByGuildNameRequest {
	return UpdateMemberRoleByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
		}(),
		RoleName: func() *string {
			v, ok := data["roleName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["roleName"])
		}(),
	}
}

func (p UpdateMemberRoleByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"targetUserId":   p.TargetUserId,
		"roleName":       p.RoleName,
	}
}

func (p UpdateMemberRoleByGuildNameRequest) Pointer() *UpdateMemberRoleByGuildNameRequest {
	return &p
}

type BatchUpdateMemberRoleRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	GuildModelName     *string  `json:"guildModelName"`
	AccessToken        *string  `json:"accessToken"`
	Members            []Member `json:"members"`
}

func (p *BatchUpdateMemberRoleRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BatchUpdateMemberRoleRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BatchUpdateMemberRoleRequest{}
	} else {
		*p = BatchUpdateMemberRoleRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["members"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Members)
		}
	}
	return nil
}

func NewBatchUpdateMemberRoleRequestFromJson(data string) (BatchUpdateMemberRoleRequest, error) {
	req := BatchUpdateMemberRoleRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return BatchUpdateMemberRoleRequest{}, err
	}
	return req, nil
}

func NewBatchUpdateMemberRoleRequestFromDict(data map[string]interface{}) BatchUpdateMemberRoleRequest {
	return BatchUpdateMemberRoleRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Members: func() []Member {
			if data["members"] == nil {
				return nil
			}
			return CastMembers(core.CastArray(data["members"]))
		}(),
	}
}

func (p BatchUpdateMemberRoleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"members": CastMembersFromDict(
			p.Members,
		),
	}
}

func (p BatchUpdateMemberRoleRequest) Pointer() *BatchUpdateMemberRoleRequest {
	return &p
}

type BatchUpdateMemberRoleByGuildNameRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	GuildModelName     *string  `json:"guildModelName"`
	GuildName          *string  `json:"guildName"`
	Members            []Member `json:"members"`
}

func (p *BatchUpdateMemberRoleByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BatchUpdateMemberRoleByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BatchUpdateMemberRoleByGuildNameRequest{}
	} else {
		*p = BatchUpdateMemberRoleByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["members"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Members)
		}
	}
	return nil
}

func NewBatchUpdateMemberRoleByGuildNameRequestFromJson(data string) (BatchUpdateMemberRoleByGuildNameRequest, error) {
	req := BatchUpdateMemberRoleByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return BatchUpdateMemberRoleByGuildNameRequest{}, err
	}
	return req, nil
}

func NewBatchUpdateMemberRoleByGuildNameRequestFromDict(data map[string]interface{}) BatchUpdateMemberRoleByGuildNameRequest {
	return BatchUpdateMemberRoleByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		Members: func() []Member {
			if data["members"] == nil {
				return nil
			}
			return CastMembers(core.CastArray(data["members"]))
		}(),
	}
}

func (p BatchUpdateMemberRoleByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"members": CastMembersFromDict(
			p.Members,
		),
	}
}

func (p BatchUpdateMemberRoleByGuildNameRequest) Pointer() *BatchUpdateMemberRoleByGuildNameRequest {
	return &p
}

type DeleteGuildRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
}

func (p *DeleteGuildRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGuildRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGuildRequest{}
	} else {
		*p = DeleteGuildRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDeleteGuildRequestFromJson(data string) (DeleteGuildRequest, error) {
	req := DeleteGuildRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGuildRequest{}, err
	}
	return req, nil
}

func NewDeleteGuildRequestFromDict(data map[string]interface{}) DeleteGuildRequest {
	return DeleteGuildRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p DeleteGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
	}
}

func (p DeleteGuildRequest) Pointer() *DeleteGuildRequest {
	return &p
}

type DeleteGuildByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func (p *DeleteGuildByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteGuildByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteGuildByGuildNameRequest{}
	} else {
		*p = DeleteGuildByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewDeleteGuildByGuildNameRequestFromJson(data string) (DeleteGuildByGuildNameRequest, error) {
	req := DeleteGuildByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteGuildByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDeleteGuildByGuildNameRequestFromDict(data map[string]interface{}) DeleteGuildByGuildNameRequest {
	return DeleteGuildByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p DeleteGuildByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p DeleteGuildByGuildNameRequest) Pointer() *DeleteGuildByGuildNameRequest {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	Value              *int32  `json:"value"`
}

func (p *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	} else {
		*p = IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) (IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest, error) {
	req := IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}, err
	}
	return req, nil
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
		}(),
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"value":          p.Value,
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	Value              *int32  `json:"value"`
}

func (p *DecreaseMaximumCurrentMaximumMemberCountRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecreaseMaximumCurrentMaximumMemberCountRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DecreaseMaximumCurrentMaximumMemberCountRequest{}
	} else {
		*p = DecreaseMaximumCurrentMaximumMemberCountRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewDecreaseMaximumCurrentMaximumMemberCountRequestFromJson(data string) (DecreaseMaximumCurrentMaximumMemberCountRequest, error) {
	req := DecreaseMaximumCurrentMaximumMemberCountRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecreaseMaximumCurrentMaximumMemberCountRequest{}, err
	}
	return req, nil
}

func NewDecreaseMaximumCurrentMaximumMemberCountRequestFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountRequest {
	return DecreaseMaximumCurrentMaximumMemberCountRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
		}(),
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"value":          p.Value,
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountRequest) Pointer() *DecreaseMaximumCurrentMaximumMemberCountRequest {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	Value              *int32  `json:"value"`
}

func (p *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	} else {
		*p = DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) (DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest, error) {
	req := DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
		}(),
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"value":          p.Value,
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type VerifyCurrentMaximumMemberCountRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	GuildModelName                  *string `json:"guildModelName"`
	AccessToken                     *string `json:"accessToken"`
	VerifyType                      *string `json:"verifyType"`
	Value                           *int32  `json:"value"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifyCurrentMaximumMemberCountRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyCurrentMaximumMemberCountRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyCurrentMaximumMemberCountRequest{}
	} else {
		*p = VerifyCurrentMaximumMemberCountRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyCurrentMaximumMemberCountRequestFromJson(data string) (VerifyCurrentMaximumMemberCountRequest, error) {
	req := VerifyCurrentMaximumMemberCountRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyCurrentMaximumMemberCountRequest{}, err
	}
	return req, nil
}

func NewVerifyCurrentMaximumMemberCountRequestFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountRequest {
	return VerifyCurrentMaximumMemberCountRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
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

func (p VerifyCurrentMaximumMemberCountRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"guildModelName":                  p.GuildModelName,
		"accessToken":                     p.AccessToken,
		"verifyType":                      p.VerifyType,
		"value":                           p.Value,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyCurrentMaximumMemberCountRequest) Pointer() *VerifyCurrentMaximumMemberCountRequest {
	return &p
}

type VerifyCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	GuildModelName                  *string `json:"guildModelName"`
	GuildName                       *string `json:"guildName"`
	VerifyType                      *string `json:"verifyType"`
	Value                           *int32  `json:"value"`
	MultiplyValueSpecifyingQuantity *bool   `json:"multiplyValueSpecifyingQuantity"`
}

func (p *VerifyCurrentMaximumMemberCountByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyCurrentMaximumMemberCountByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyCurrentMaximumMemberCountByGuildNameRequest{}
	} else {
		*p = VerifyCurrentMaximumMemberCountByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
		if v, ok := d["multiplyValueSpecifyingQuantity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MultiplyValueSpecifyingQuantity)
		}
	}
	return nil
}

func NewVerifyCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) (VerifyCurrentMaximumMemberCountByGuildNameRequest, error) {
	req := VerifyCurrentMaximumMemberCountByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyCurrentMaximumMemberCountByGuildNameRequest{}, err
	}
	return req, nil
}

func NewVerifyCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountByGuildNameRequest {
	return VerifyCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
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

func (p VerifyCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"guildModelName":                  p.GuildModelName,
		"guildName":                       p.GuildName,
		"verifyType":                      p.VerifyType,
		"value":                           p.Value,
		"multiplyValueSpecifyingQuantity": p.MultiplyValueSpecifyingQuantity,
	}
}

func (p VerifyCurrentMaximumMemberCountByGuildNameRequest) Pointer() *VerifyCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type VerifyIncludeMemberRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	AccessToken        *string `json:"accessToken"`
	VerifyType         *string `json:"verifyType"`
}

func (p *VerifyIncludeMemberRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeMemberRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeMemberRequest{}
	} else {
		*p = VerifyIncludeMemberRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewVerifyIncludeMemberRequestFromJson(data string) (VerifyIncludeMemberRequest, error) {
	req := VerifyIncludeMemberRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeMemberRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeMemberRequestFromDict(data map[string]interface{}) VerifyIncludeMemberRequest {
	return VerifyIncludeMemberRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
		}(),
	}
}

func (p VerifyIncludeMemberRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"accessToken":    p.AccessToken,
		"verifyType":     p.VerifyType,
	}
}

func (p VerifyIncludeMemberRequest) Pointer() *VerifyIncludeMemberRequest {
	return &p
}

type VerifyIncludeMemberByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	UserId             *string `json:"userId"`
	VerifyType         *string `json:"verifyType"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *VerifyIncludeMemberByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeMemberByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeMemberByUserIdRequest{}
	} else {
		*p = VerifyIncludeMemberByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewVerifyIncludeMemberByUserIdRequestFromJson(data string) (VerifyIncludeMemberByUserIdRequest, error) {
	req := VerifyIncludeMemberByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeMemberByUserIdRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeMemberByUserIdRequestFromDict(data map[string]interface{}) VerifyIncludeMemberByUserIdRequest {
	return VerifyIncludeMemberByUserIdRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		VerifyType: func() *string {
			v, ok := data["verifyType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyType"])
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

func (p VerifyIncludeMemberByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"userId":          p.UserId,
		"verifyType":      p.VerifyType,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p VerifyIncludeMemberByUserIdRequest) Pointer() *VerifyIncludeMemberByUserIdRequest {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildName          *string `json:"guildName"`
	GuildModelName     *string `json:"guildModelName"`
	Value              *int32  `json:"value"`
}

func (p *SetMaximumCurrentMaximumMemberCountByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	} else {
		*p = SetMaximumCurrentMaximumMemberCountByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) (SetMaximumCurrentMaximumMemberCountByGuildNameRequest, error) {
	req := SetMaximumCurrentMaximumMemberCountByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetMaximumCurrentMaximumMemberCountByGuildNameRequest{}, err
	}
	return req, nil
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return SetMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
		}(),
	}
}

func (p SetMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildName":      p.GuildName,
		"guildModelName": p.GuildModelName,
		"value":          p.Value,
	}
}

func (p SetMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *SetMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type AssumeRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func (p *AssumeRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AssumeRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AssumeRequest{}
	} else {
		*p = AssumeRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewAssumeRequestFromJson(data string) (AssumeRequest, error) {
	req := AssumeRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AssumeRequest{}, err
	}
	return req, nil
}

func NewAssumeRequestFromDict(data map[string]interface{}) AssumeRequest {
	return AssumeRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p AssumeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p AssumeRequest) Pointer() *AssumeRequest {
	return &p
}

type AssumeByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AssumeByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AssumeByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AssumeByUserIdRequest{}
	} else {
		*p = AssumeByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewAssumeByUserIdRequestFromJson(data string) (AssumeByUserIdRequest, error) {
	req := AssumeByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AssumeByUserIdRequest{}, err
	}
	return req, nil
}

func NewAssumeByUserIdRequestFromDict(data map[string]interface{}) AssumeByUserIdRequest {
	return AssumeByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p AssumeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AssumeByUserIdRequest) Pointer() *AssumeByUserIdRequest {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{}
	} else {
		*p = IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{}
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

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetRequestFromJson(data string) (IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest, error) {
	req := IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{}, err
	}
	return req, nil
}

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{
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

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{}
	} else {
		*p = DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{}
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

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskRequestFromJson(data string) (DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest, error) {
	req := DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{}, err
	}
	return req, nil
}

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskRequestFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest {
	return DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{
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

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func (p *SetMaximumCurrentMaximumMemberCountByStampSheetRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SetMaximumCurrentMaximumMemberCountByStampSheetRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SetMaximumCurrentMaximumMemberCountByStampSheetRequest{}
	} else {
		*p = SetMaximumCurrentMaximumMemberCountByStampSheetRequest{}
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

func NewSetMaximumCurrentMaximumMemberCountByStampSheetRequestFromJson(data string) (SetMaximumCurrentMaximumMemberCountByStampSheetRequest, error) {
	req := SetMaximumCurrentMaximumMemberCountByStampSheetRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SetMaximumCurrentMaximumMemberCountByStampSheetRequest{}, err
	}
	return req, nil
}

func NewSetMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return SetMaximumCurrentMaximumMemberCountByStampSheetRequest{
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

func (p SetMaximumCurrentMaximumMemberCountByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetMaximumCurrentMaximumMemberCountByStampSheetRequest) Pointer() *SetMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return &p
}

type VerifyCurrentMaximumMemberCountByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyCurrentMaximumMemberCountByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyCurrentMaximumMemberCountByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyCurrentMaximumMemberCountByStampTaskRequest{}
	} else {
		*p = VerifyCurrentMaximumMemberCountByStampTaskRequest{}
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

func NewVerifyCurrentMaximumMemberCountByStampTaskRequestFromJson(data string) (VerifyCurrentMaximumMemberCountByStampTaskRequest, error) {
	req := VerifyCurrentMaximumMemberCountByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyCurrentMaximumMemberCountByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyCurrentMaximumMemberCountByStampTaskRequestFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountByStampTaskRequest {
	return VerifyCurrentMaximumMemberCountByStampTaskRequest{
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

func (p VerifyCurrentMaximumMemberCountByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyCurrentMaximumMemberCountByStampTaskRequest) Pointer() *VerifyCurrentMaximumMemberCountByStampTaskRequest {
	return &p
}

type VerifyIncludeMemberByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func (p *VerifyIncludeMemberByStampTaskRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyIncludeMemberByStampTaskRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyIncludeMemberByStampTaskRequest{}
	} else {
		*p = VerifyIncludeMemberByStampTaskRequest{}
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

func NewVerifyIncludeMemberByStampTaskRequestFromJson(data string) (VerifyIncludeMemberByStampTaskRequest, error) {
	req := VerifyIncludeMemberByStampTaskRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return VerifyIncludeMemberByStampTaskRequest{}, err
	}
	return req, nil
}

func NewVerifyIncludeMemberByStampTaskRequestFromDict(data map[string]interface{}) VerifyIncludeMemberByStampTaskRequest {
	return VerifyIncludeMemberByStampTaskRequest{
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

func (p VerifyIncludeMemberByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p VerifyIncludeMemberByStampTaskRequest) Pointer() *VerifyIncludeMemberByStampTaskRequest {
	return &p
}

type DescribeJoinedGuildsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeJoinedGuildsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeJoinedGuildsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeJoinedGuildsRequest{}
	} else {
		*p = DescribeJoinedGuildsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeJoinedGuildsRequestFromJson(data string) (DescribeJoinedGuildsRequest, error) {
	req := DescribeJoinedGuildsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeJoinedGuildsRequest{}, err
	}
	return req, nil
}

func NewDescribeJoinedGuildsRequestFromDict(data map[string]interface{}) DescribeJoinedGuildsRequest {
	return DescribeJoinedGuildsRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeJoinedGuildsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeJoinedGuildsRequest) Pointer() *DescribeJoinedGuildsRequest {
	return &p
}

type DescribeJoinedGuildsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeJoinedGuildsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeJoinedGuildsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeJoinedGuildsByUserIdRequest{}
	} else {
		*p = DescribeJoinedGuildsByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeJoinedGuildsByUserIdRequestFromJson(data string) (DescribeJoinedGuildsByUserIdRequest, error) {
	req := DescribeJoinedGuildsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeJoinedGuildsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeJoinedGuildsByUserIdRequestFromDict(data map[string]interface{}) DescribeJoinedGuildsByUserIdRequest {
	return DescribeJoinedGuildsByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeJoinedGuildsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeJoinedGuildsByUserIdRequest) Pointer() *DescribeJoinedGuildsByUserIdRequest {
	return &p
}

type GetJoinedGuildRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
}

func (p *GetJoinedGuildRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetJoinedGuildRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetJoinedGuildRequest{}
	} else {
		*p = GetJoinedGuildRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewGetJoinedGuildRequestFromJson(data string) (GetJoinedGuildRequest, error) {
	req := GetJoinedGuildRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetJoinedGuildRequest{}, err
	}
	return req, nil
}

func NewGetJoinedGuildRequestFromDict(data map[string]interface{}) GetJoinedGuildRequest {
	return GetJoinedGuildRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p GetJoinedGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p GetJoinedGuildRequest) Pointer() *GetJoinedGuildRequest {
	return &p
}

type GetJoinedGuildByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetJoinedGuildByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetJoinedGuildByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetJoinedGuildByUserIdRequest{}
	} else {
		*p = GetJoinedGuildByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewGetJoinedGuildByUserIdRequestFromJson(data string) (GetJoinedGuildByUserIdRequest, error) {
	req := GetJoinedGuildByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetJoinedGuildByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetJoinedGuildByUserIdRequestFromDict(data map[string]interface{}) GetJoinedGuildByUserIdRequest {
	return GetJoinedGuildByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p GetJoinedGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetJoinedGuildByUserIdRequest) Pointer() *GetJoinedGuildByUserIdRequest {
	return &p
}

type WithdrawalRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func (p *WithdrawalRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawalRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = WithdrawalRequest{}
	} else {
		*p = WithdrawalRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewWithdrawalRequestFromJson(data string) (WithdrawalRequest, error) {
	req := WithdrawalRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WithdrawalRequest{}, err
	}
	return req, nil
}

func NewWithdrawalRequestFromDict(data map[string]interface{}) WithdrawalRequest {
	return WithdrawalRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p WithdrawalRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p WithdrawalRequest) Pointer() *WithdrawalRequest {
	return &p
}

type WithdrawalByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *WithdrawalByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WithdrawalByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = WithdrawalByUserIdRequest{}
	} else {
		*p = WithdrawalByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewWithdrawalByUserIdRequestFromJson(data string) (WithdrawalByUserIdRequest, error) {
	req := WithdrawalByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return WithdrawalByUserIdRequest{}, err
	}
	return req, nil
}

func NewWithdrawalByUserIdRequestFromDict(data map[string]interface{}) WithdrawalByUserIdRequest {
	return WithdrawalByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p WithdrawalByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p WithdrawalByUserIdRequest) Pointer() *WithdrawalByUserIdRequest {
	return &p
}

type GetLastGuildMasterActivityRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
}

func (p *GetLastGuildMasterActivityRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLastGuildMasterActivityRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLastGuildMasterActivityRequest{}
	} else {
		*p = GetLastGuildMasterActivityRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewGetLastGuildMasterActivityRequestFromJson(data string) (GetLastGuildMasterActivityRequest, error) {
	req := GetLastGuildMasterActivityRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLastGuildMasterActivityRequest{}, err
	}
	return req, nil
}

func NewGetLastGuildMasterActivityRequestFromDict(data map[string]interface{}) GetLastGuildMasterActivityRequest {
	return GetLastGuildMasterActivityRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p GetLastGuildMasterActivityRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
	}
}

func (p GetLastGuildMasterActivityRequest) Pointer() *GetLastGuildMasterActivityRequest {
	return &p
}

type GetLastGuildMasterActivityByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
}

func (p *GetLastGuildMasterActivityByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetLastGuildMasterActivityByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetLastGuildMasterActivityByGuildNameRequest{}
	} else {
		*p = GetLastGuildMasterActivityByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewGetLastGuildMasterActivityByGuildNameRequestFromJson(data string) (GetLastGuildMasterActivityByGuildNameRequest, error) {
	req := GetLastGuildMasterActivityByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetLastGuildMasterActivityByGuildNameRequest{}, err
	}
	return req, nil
}

func NewGetLastGuildMasterActivityByGuildNameRequestFromDict(data map[string]interface{}) GetLastGuildMasterActivityByGuildNameRequest {
	return GetLastGuildMasterActivityByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p GetLastGuildMasterActivityByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p GetLastGuildMasterActivityByGuildNameRequest) Pointer() *GetLastGuildMasterActivityByGuildNameRequest {
	return &p
}

type PromoteSeniorMemberRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
}

func (p *PromoteSeniorMemberRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PromoteSeniorMemberRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PromoteSeniorMemberRequest{}
	} else {
		*p = PromoteSeniorMemberRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewPromoteSeniorMemberRequestFromJson(data string) (PromoteSeniorMemberRequest, error) {
	req := PromoteSeniorMemberRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PromoteSeniorMemberRequest{}, err
	}
	return req, nil
}

func NewPromoteSeniorMemberRequestFromDict(data map[string]interface{}) PromoteSeniorMemberRequest {
	return PromoteSeniorMemberRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
	}
}

func (p PromoteSeniorMemberRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
	}
}

func (p PromoteSeniorMemberRequest) Pointer() *PromoteSeniorMemberRequest {
	return &p
}

type PromoteSeniorMemberByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func (p *PromoteSeniorMemberByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PromoteSeniorMemberByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PromoteSeniorMemberByGuildNameRequest{}
	} else {
		*p = PromoteSeniorMemberByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
	}
	return nil
}

func NewPromoteSeniorMemberByGuildNameRequestFromJson(data string) (PromoteSeniorMemberByGuildNameRequest, error) {
	req := PromoteSeniorMemberByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PromoteSeniorMemberByGuildNameRequest{}, err
	}
	return req, nil
}

func NewPromoteSeniorMemberByGuildNameRequestFromDict(data map[string]interface{}) PromoteSeniorMemberByGuildNameRequest {
	return PromoteSeniorMemberByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
	}
}

func (p PromoteSeniorMemberByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p PromoteSeniorMemberByGuildNameRequest) Pointer() *PromoteSeniorMemberByGuildNameRequest {
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

type GetCurrentGuildMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetCurrentGuildMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetCurrentGuildMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetCurrentGuildMasterRequest{}
	} else {
		*p = GetCurrentGuildMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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

func NewGetCurrentGuildMasterRequestFromJson(data string) (GetCurrentGuildMasterRequest, error) {
	req := GetCurrentGuildMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetCurrentGuildMasterRequest{}, err
	}
	return req, nil
}

func NewGetCurrentGuildMasterRequestFromDict(data map[string]interface{}) GetCurrentGuildMasterRequest {
	return GetCurrentGuildMasterRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
	}
}

func (p GetCurrentGuildMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentGuildMasterRequest) Pointer() *GetCurrentGuildMasterRequest {
	return &p
}

type UpdateCurrentGuildMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func (p *UpdateCurrentGuildMasterRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentGuildMasterRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentGuildMasterRequest{}
	} else {
		*p = UpdateCurrentGuildMasterRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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

func NewUpdateCurrentGuildMasterRequestFromJson(data string) (UpdateCurrentGuildMasterRequest, error) {
	req := UpdateCurrentGuildMasterRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentGuildMasterRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentGuildMasterRequestFromDict(data map[string]interface{}) UpdateCurrentGuildMasterRequest {
	return UpdateCurrentGuildMasterRequest{
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

func (p UpdateCurrentGuildMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentGuildMasterRequest) Pointer() *UpdateCurrentGuildMasterRequest {
	return &p
}

type UpdateCurrentGuildMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func (p *UpdateCurrentGuildMasterFromGitHubRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateCurrentGuildMasterFromGitHubRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = UpdateCurrentGuildMasterFromGitHubRequest{}
	} else {
		*p = UpdateCurrentGuildMasterFromGitHubRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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

func NewUpdateCurrentGuildMasterFromGitHubRequestFromJson(data string) (UpdateCurrentGuildMasterFromGitHubRequest, error) {
	req := UpdateCurrentGuildMasterFromGitHubRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateCurrentGuildMasterFromGitHubRequest{}, err
	}
	return req, nil
}

func NewUpdateCurrentGuildMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentGuildMasterFromGitHubRequest {
	return UpdateCurrentGuildMasterFromGitHubRequest{
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

func (p UpdateCurrentGuildMasterFromGitHubRequest) ToDict() map[string]interface{} {
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

func (p UpdateCurrentGuildMasterFromGitHubRequest) Pointer() *UpdateCurrentGuildMasterFromGitHubRequest {
	return &p
}

type DescribeReceiveRequestsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeReceiveRequestsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeReceiveRequestsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeReceiveRequestsRequest{}
	} else {
		*p = DescribeReceiveRequestsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeReceiveRequestsRequestFromJson(data string) (DescribeReceiveRequestsRequest, error) {
	req := DescribeReceiveRequestsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeReceiveRequestsRequest{}, err
	}
	return req, nil
}

func NewDescribeReceiveRequestsRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsRequest {
	return DescribeReceiveRequestsRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeReceiveRequestsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeReceiveRequestsRequest) Pointer() *DescribeReceiveRequestsRequest {
	return &p
}

type DescribeReceiveRequestsByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeReceiveRequestsByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeReceiveRequestsByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeReceiveRequestsByGuildNameRequest{}
	} else {
		*p = DescribeReceiveRequestsByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewDescribeReceiveRequestsByGuildNameRequestFromJson(data string) (DescribeReceiveRequestsByGuildNameRequest, error) {
	req := DescribeReceiveRequestsByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeReceiveRequestsByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDescribeReceiveRequestsByGuildNameRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsByGuildNameRequest {
	return DescribeReceiveRequestsByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p DescribeReceiveRequestsByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeReceiveRequestsByGuildNameRequest) Pointer() *DescribeReceiveRequestsByGuildNameRequest {
	return &p
}

type GetReceiveRequestRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	FromUserId      *string `json:"fromUserId"`
}

func (p *GetReceiveRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetReceiveRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetReceiveRequestRequest{}
	} else {
		*p = GetReceiveRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewGetReceiveRequestRequestFromJson(data string) (GetReceiveRequestRequest, error) {
	req := GetReceiveRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetReceiveRequestRequest{}, err
	}
	return req, nil
}

func NewGetReceiveRequestRequestFromDict(data map[string]interface{}) GetReceiveRequestRequest {
	return GetReceiveRequestRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p GetReceiveRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p GetReceiveRequestRequest) Pointer() *GetReceiveRequestRequest {
	return &p
}

type GetReceiveRequestByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	FromUserId      *string `json:"fromUserId"`
}

func (p *GetReceiveRequestByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetReceiveRequestByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetReceiveRequestByGuildNameRequest{}
	} else {
		*p = GetReceiveRequestByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewGetReceiveRequestByGuildNameRequestFromJson(data string) (GetReceiveRequestByGuildNameRequest, error) {
	req := GetReceiveRequestByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetReceiveRequestByGuildNameRequest{}, err
	}
	return req, nil
}

func NewGetReceiveRequestByGuildNameRequestFromDict(data map[string]interface{}) GetReceiveRequestByGuildNameRequest {
	return GetReceiveRequestByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p GetReceiveRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p GetReceiveRequestByGuildNameRequest) Pointer() *GetReceiveRequestByGuildNameRequest {
	return &p
}

type AcceptRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	FromUserId         *string `json:"fromUserId"`
}

func (p *AcceptRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcceptRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcceptRequestRequest{}
	} else {
		*p = AcceptRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewAcceptRequestRequestFromJson(data string) (AcceptRequestRequest, error) {
	req := AcceptRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcceptRequestRequest{}, err
	}
	return req, nil
}

func NewAcceptRequestRequestFromDict(data map[string]interface{}) AcceptRequestRequest {
	return AcceptRequestRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p AcceptRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p AcceptRequestRequest) Pointer() *AcceptRequestRequest {
	return &p
}

type AcceptRequestByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	FromUserId         *string `json:"fromUserId"`
}

func (p *AcceptRequestByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcceptRequestByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcceptRequestByGuildNameRequest{}
	} else {
		*p = AcceptRequestByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewAcceptRequestByGuildNameRequestFromJson(data string) (AcceptRequestByGuildNameRequest, error) {
	req := AcceptRequestByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AcceptRequestByGuildNameRequest{}, err
	}
	return req, nil
}

func NewAcceptRequestByGuildNameRequestFromDict(data map[string]interface{}) AcceptRequestByGuildNameRequest {
	return AcceptRequestByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p AcceptRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p AcceptRequestByGuildNameRequest) Pointer() *AcceptRequestByGuildNameRequest {
	return &p
}

type RejectRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	FromUserId         *string `json:"fromUserId"`
}

func (p *RejectRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RejectRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RejectRequestRequest{}
	} else {
		*p = RejectRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewRejectRequestRequestFromJson(data string) (RejectRequestRequest, error) {
	req := RejectRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RejectRequestRequest{}, err
	}
	return req, nil
}

func NewRejectRequestRequestFromDict(data map[string]interface{}) RejectRequestRequest {
	return RejectRequestRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p RejectRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p RejectRequestRequest) Pointer() *RejectRequestRequest {
	return &p
}

type RejectRequestByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	FromUserId         *string `json:"fromUserId"`
}

func (p *RejectRequestByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RejectRequestByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RejectRequestByGuildNameRequest{}
	} else {
		*p = RejectRequestByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
				}
			}
		}
		if v, ok := d["fromUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FromUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FromUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FromUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FromUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FromUserId)
				}
			}
		}
	}
	return nil
}

func NewRejectRequestByGuildNameRequestFromJson(data string) (RejectRequestByGuildNameRequest, error) {
	req := RejectRequestByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RejectRequestByGuildNameRequest{}, err
	}
	return req, nil
}

func NewRejectRequestByGuildNameRequestFromDict(data map[string]interface{}) RejectRequestByGuildNameRequest {
	return RejectRequestByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
		FromUserId: func() *string {
			v, ok := data["fromUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fromUserId"])
		}(),
	}
}

func (p RejectRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p RejectRequestByGuildNameRequest) Pointer() *RejectRequestByGuildNameRequest {
	return &p
}

type DescribeSendRequestsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeSendRequestsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSendRequestsRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSendRequestsRequest{}
	} else {
		*p = DescribeSendRequestsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeSendRequestsRequestFromJson(data string) (DescribeSendRequestsRequest, error) {
	req := DescribeSendRequestsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSendRequestsRequest{}, err
	}
	return req, nil
}

func NewDescribeSendRequestsRequestFromDict(data map[string]interface{}) DescribeSendRequestsRequest {
	return DescribeSendRequestsRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeSendRequestsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeSendRequestsRequest) Pointer() *DescribeSendRequestsRequest {
	return &p
}

type DescribeSendRequestsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *DescribeSendRequestsByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeSendRequestsByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeSendRequestsByUserIdRequest{}
	} else {
		*p = DescribeSendRequestsByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeSendRequestsByUserIdRequestFromJson(data string) (DescribeSendRequestsByUserIdRequest, error) {
	req := DescribeSendRequestsByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeSendRequestsByUserIdRequest{}, err
	}
	return req, nil
}

func NewDescribeSendRequestsByUserIdRequestFromDict(data map[string]interface{}) DescribeSendRequestsByUserIdRequest {
	return DescribeSendRequestsByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeSendRequestsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSendRequestsByUserIdRequest) Pointer() *DescribeSendRequestsByUserIdRequest {
	return &p
}

type GetSendRequestRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	TargetGuildName *string `json:"targetGuildName"`
}

func (p *GetSendRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSendRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSendRequestRequest{}
	} else {
		*p = GetSendRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
				}
			}
		}
	}
	return nil
}

func NewGetSendRequestRequestFromJson(data string) (GetSendRequestRequest, error) {
	req := GetSendRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSendRequestRequest{}, err
	}
	return req, nil
}

func NewGetSendRequestRequestFromDict(data map[string]interface{}) GetSendRequestRequest {
	return GetSendRequestRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
		}(),
	}
}

func (p GetSendRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p GetSendRequestRequest) Pointer() *GetSendRequestRequest {
	return &p
}

type GetSendRequestByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	TargetGuildName *string `json:"targetGuildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetSendRequestByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetSendRequestByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetSendRequestByUserIdRequest{}
	} else {
		*p = GetSendRequestByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
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

func NewGetSendRequestByUserIdRequestFromJson(data string) (GetSendRequestByUserIdRequest, error) {
	req := GetSendRequestByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetSendRequestByUserIdRequest{}, err
	}
	return req, nil
}

func NewGetSendRequestByUserIdRequestFromDict(data map[string]interface{}) GetSendRequestByUserIdRequest {
	return GetSendRequestByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
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

func (p GetSendRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSendRequestByUserIdRequest) Pointer() *GetSendRequestByUserIdRequest {
	return &p
}

type SendRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
}

func (p *SendRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SendRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SendRequestRequest{}
	} else {
		*p = SendRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
				}
			}
		}
	}
	return nil
}

func NewSendRequestRequestFromJson(data string) (SendRequestRequest, error) {
	req := SendRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SendRequestRequest{}, err
	}
	return req, nil
}

func NewSendRequestRequestFromDict(data map[string]interface{}) SendRequestRequest {
	return SendRequestRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
		}(),
	}
}

func (p SendRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p SendRequestRequest) Pointer() *SendRequestRequest {
	return &p
}

type SendRequestByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *SendRequestByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SendRequestByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SendRequestByUserIdRequest{}
	} else {
		*p = SendRequestByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
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

func NewSendRequestByUserIdRequestFromJson(data string) (SendRequestByUserIdRequest, error) {
	req := SendRequestByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return SendRequestByUserIdRequest{}, err
	}
	return req, nil
}

func NewSendRequestByUserIdRequestFromDict(data map[string]interface{}) SendRequestByUserIdRequest {
	return SendRequestByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
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

func (p SendRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SendRequestByUserIdRequest) Pointer() *SendRequestByUserIdRequest {
	return &p
}

type DeleteRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
}

func (p *DeleteRequestRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRequestRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRequestRequest{}
	} else {
		*p = DeleteRequestRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
				}
			}
		}
	}
	return nil
}

func NewDeleteRequestRequestFromJson(data string) (DeleteRequestRequest, error) {
	req := DeleteRequestRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRequestRequest{}, err
	}
	return req, nil
}

func NewDeleteRequestRequestFromDict(data map[string]interface{}) DeleteRequestRequest {
	return DeleteRequestRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
		}(),
	}
}

func (p DeleteRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p DeleteRequestRequest) Pointer() *DeleteRequestRequest {
	return &p
}

type DeleteRequestByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteRequestByUserIdRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteRequestByUserIdRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteRequestByUserIdRequest{}
	} else {
		*p = DeleteRequestByUserIdRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
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
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["targetGuildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGuildName)
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

func NewDeleteRequestByUserIdRequestFromJson(data string) (DeleteRequestByUserIdRequest, error) {
	req := DeleteRequestByUserIdRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteRequestByUserIdRequest{}, err
	}
	return req, nil
}

func NewDeleteRequestByUserIdRequestFromDict(data map[string]interface{}) DeleteRequestByUserIdRequest {
	return DeleteRequestByUserIdRequest{
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
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
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

func (p DeleteRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteRequestByUserIdRequest) Pointer() *DeleteRequestByUserIdRequest {
	return &p
}

type DescribeIgnoreUsersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeIgnoreUsersRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeIgnoreUsersRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeIgnoreUsersRequest{}
	} else {
		*p = DescribeIgnoreUsersRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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

func NewDescribeIgnoreUsersRequestFromJson(data string) (DescribeIgnoreUsersRequest, error) {
	req := DescribeIgnoreUsersRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeIgnoreUsersRequest{}, err
	}
	return req, nil
}

func NewDescribeIgnoreUsersRequestFromDict(data map[string]interface{}) DescribeIgnoreUsersRequest {
	return DescribeIgnoreUsersRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
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

func (p DescribeIgnoreUsersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeIgnoreUsersRequest) Pointer() *DescribeIgnoreUsersRequest {
	return &p
}

type DescribeIgnoreUsersByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeIgnoreUsersByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeIgnoreUsersByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DescribeIgnoreUsersByGuildNameRequest{}
	} else {
		*p = DescribeIgnoreUsersByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewDescribeIgnoreUsersByGuildNameRequestFromJson(data string) (DescribeIgnoreUsersByGuildNameRequest, error) {
	req := DescribeIgnoreUsersByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeIgnoreUsersByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDescribeIgnoreUsersByGuildNameRequestFromDict(data map[string]interface{}) DescribeIgnoreUsersByGuildNameRequest {
	return DescribeIgnoreUsersByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
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

func (p DescribeIgnoreUsersByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeIgnoreUsersByGuildNameRequest) Pointer() *DescribeIgnoreUsersByGuildNameRequest {
	return &p
}

type GetIgnoreUserRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	UserId          *string `json:"userId"`
}

func (p *GetIgnoreUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetIgnoreUserRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetIgnoreUserRequest{}
	} else {
		*p = GetIgnoreUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
	}
	return nil
}

func NewGetIgnoreUserRequestFromJson(data string) (GetIgnoreUserRequest, error) {
	req := GetIgnoreUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetIgnoreUserRequest{}, err
	}
	return req, nil
}

func NewGetIgnoreUserRequestFromDict(data map[string]interface{}) GetIgnoreUserRequest {
	return GetIgnoreUserRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
	}
}

func (p GetIgnoreUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"userId":         p.UserId,
	}
}

func (p GetIgnoreUserRequest) Pointer() *GetIgnoreUserRequest {
	return &p
}

type GetIgnoreUserByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *GetIgnoreUserByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetIgnoreUserByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GetIgnoreUserByGuildNameRequest{}
	} else {
		*p = GetIgnoreUserByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewGetIgnoreUserByGuildNameRequestFromJson(data string) (GetIgnoreUserByGuildNameRequest, error) {
	req := GetIgnoreUserByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetIgnoreUserByGuildNameRequest{}, err
	}
	return req, nil
}

func NewGetIgnoreUserByGuildNameRequestFromDict(data map[string]interface{}) GetIgnoreUserByGuildNameRequest {
	return GetIgnoreUserByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
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

func (p GetIgnoreUserByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetIgnoreUserByGuildNameRequest) Pointer() *GetIgnoreUserByGuildNameRequest {
	return &p
}

type AddIgnoreUserRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	UserId             *string `json:"userId"`
}

func (p *AddIgnoreUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddIgnoreUserRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddIgnoreUserRequest{}
	} else {
		*p = AddIgnoreUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
	}
	return nil
}

func NewAddIgnoreUserRequestFromJson(data string) (AddIgnoreUserRequest, error) {
	req := AddIgnoreUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddIgnoreUserRequest{}, err
	}
	return req, nil
}

func NewAddIgnoreUserRequestFromDict(data map[string]interface{}) AddIgnoreUserRequest {
	return AddIgnoreUserRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
	}
}

func (p AddIgnoreUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"userId":         p.UserId,
	}
}

func (p AddIgnoreUserRequest) Pointer() *AddIgnoreUserRequest {
	return &p
}

type AddIgnoreUserByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *AddIgnoreUserByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AddIgnoreUserByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AddIgnoreUserByGuildNameRequest{}
	} else {
		*p = AddIgnoreUserByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewAddIgnoreUserByGuildNameRequestFromJson(data string) (AddIgnoreUserByGuildNameRequest, error) {
	req := AddIgnoreUserByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return AddIgnoreUserByGuildNameRequest{}, err
	}
	return req, nil
}

func NewAddIgnoreUserByGuildNameRequestFromDict(data map[string]interface{}) AddIgnoreUserByGuildNameRequest {
	return AddIgnoreUserByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
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

func (p AddIgnoreUserByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AddIgnoreUserByGuildNameRequest) Pointer() *AddIgnoreUserByGuildNameRequest {
	return &p
}

type DeleteIgnoreUserRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	UserId             *string `json:"userId"`
}

func (p *DeleteIgnoreUserRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteIgnoreUserRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteIgnoreUserRequest{}
	} else {
		*p = DeleteIgnoreUserRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
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
	}
	return nil
}

func NewDeleteIgnoreUserRequestFromJson(data string) (DeleteIgnoreUserRequest, error) {
	req := DeleteIgnoreUserRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteIgnoreUserRequest{}, err
	}
	return req, nil
}

func NewDeleteIgnoreUserRequestFromDict(data map[string]interface{}) DeleteIgnoreUserRequest {
	return DeleteIgnoreUserRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		AccessToken: func() *string {
			v, ok := data["accessToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessToken"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
	}
}

func (p DeleteIgnoreUserRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"userId":         p.UserId,
	}
}

func (p DeleteIgnoreUserRequest) Pointer() *DeleteIgnoreUserRequest {
	return &p
}

type DeleteIgnoreUserByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *DeleteIgnoreUserByGuildNameRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteIgnoreUserByGuildNameRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeleteIgnoreUserByGuildNameRequest{}
	} else {
		*p = DeleteIgnoreUserByGuildNameRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["guildModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelName)
				}
			}
		}
		if v, ok := d["guildName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildName)
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

func NewDeleteIgnoreUserByGuildNameRequestFromJson(data string) (DeleteIgnoreUserByGuildNameRequest, error) {
	req := DeleteIgnoreUserByGuildNameRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteIgnoreUserByGuildNameRequest{}, err
	}
	return req, nil
}

func NewDeleteIgnoreUserByGuildNameRequestFromDict(data map[string]interface{}) DeleteIgnoreUserByGuildNameRequest {
	return DeleteIgnoreUserByGuildNameRequest{
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		GuildName: func() *string {
			v, ok := data["guildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildName"])
		}(),
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

func (p DeleteIgnoreUserByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteIgnoreUserByGuildNameRequest) Pointer() *DeleteIgnoreUserByGuildNameRequest {
	return &p
}
