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

type Namespace struct {
	NamespaceId                *string              `json:"namespaceId"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	TransactionSetting         *TransactionSetting  `json:"transactionSetting"`
	ChangeNotification         *NotificationSetting `json:"changeNotification"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
	CreateGuildScript          *ScriptSetting       `json:"createGuildScript"`
	UpdateGuildScript          *ScriptSetting       `json:"updateGuildScript"`
	JoinGuildScript            *ScriptSetting       `json:"joinGuildScript"`
	ReceiveJoinRequestScript   *ScriptSetting       `json:"receiveJoinRequestScript"`
	LeaveGuildScript           *ScriptSetting       `json:"leaveGuildScript"`
	ChangeRoleScript           *ScriptSetting       `json:"changeRoleScript"`
	DeleteGuildScript          *ScriptSetting       `json:"deleteGuildScript"`
	LogSetting                 *LogSetting          `json:"logSetting"`
	CreatedAt                  *int64               `json:"createdAt"`
	UpdatedAt                  *int64               `json:"updatedAt"`
	Revision                   *int64               `json:"revision"`
}

func (p *Namespace) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Namespace{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Namespace{}
	} else {
		*p = Namespace{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
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
		if v, ok := d["receiveJoinRequestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveJoinRequestScript)
		}
		if v, ok := d["leaveGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveGuildScript)
		}
		if v, ok := d["changeRoleScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRoleScript)
		}
		if v, ok := d["deleteGuildScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DeleteGuildScript)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewNamespaceFromJson(data string) Namespace {
	req := Namespace{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
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
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
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
		ReceiveJoinRequestScript: func() *ScriptSetting {
			v, ok := data["receiveJoinRequestScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["receiveJoinRequestScript"])).Pointer()
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
		DeleteGuildScript: func() *ScriptSetting {
			v, ok := data["deleteGuildScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["deleteGuildScript"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.TransactionSetting != nil {
		m["transactionSetting"] = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	if p.ChangeNotification != nil {
		m["changeNotification"] = func() map[string]interface{} {
			if p.ChangeNotification == nil {
				return nil
			}
			return p.ChangeNotification.ToDict()
		}()
	}
	if p.JoinNotification != nil {
		m["joinNotification"] = func() map[string]interface{} {
			if p.JoinNotification == nil {
				return nil
			}
			return p.JoinNotification.ToDict()
		}()
	}
	if p.LeaveNotification != nil {
		m["leaveNotification"] = func() map[string]interface{} {
			if p.LeaveNotification == nil {
				return nil
			}
			return p.LeaveNotification.ToDict()
		}()
	}
	if p.ChangeMemberNotification != nil {
		m["changeMemberNotification"] = func() map[string]interface{} {
			if p.ChangeMemberNotification == nil {
				return nil
			}
			return p.ChangeMemberNotification.ToDict()
		}()
	}
	if p.ReceiveRequestNotification != nil {
		m["receiveRequestNotification"] = func() map[string]interface{} {
			if p.ReceiveRequestNotification == nil {
				return nil
			}
			return p.ReceiveRequestNotification.ToDict()
		}()
	}
	if p.RemoveRequestNotification != nil {
		m["removeRequestNotification"] = func() map[string]interface{} {
			if p.RemoveRequestNotification == nil {
				return nil
			}
			return p.RemoveRequestNotification.ToDict()
		}()
	}
	if p.CreateGuildScript != nil {
		m["createGuildScript"] = func() map[string]interface{} {
			if p.CreateGuildScript == nil {
				return nil
			}
			return p.CreateGuildScript.ToDict()
		}()
	}
	if p.UpdateGuildScript != nil {
		m["updateGuildScript"] = func() map[string]interface{} {
			if p.UpdateGuildScript == nil {
				return nil
			}
			return p.UpdateGuildScript.ToDict()
		}()
	}
	if p.JoinGuildScript != nil {
		m["joinGuildScript"] = func() map[string]interface{} {
			if p.JoinGuildScript == nil {
				return nil
			}
			return p.JoinGuildScript.ToDict()
		}()
	}
	if p.ReceiveJoinRequestScript != nil {
		m["receiveJoinRequestScript"] = func() map[string]interface{} {
			if p.ReceiveJoinRequestScript == nil {
				return nil
			}
			return p.ReceiveJoinRequestScript.ToDict()
		}()
	}
	if p.LeaveGuildScript != nil {
		m["leaveGuildScript"] = func() map[string]interface{} {
			if p.LeaveGuildScript == nil {
				return nil
			}
			return p.LeaveGuildScript.ToDict()
		}()
	}
	if p.ChangeRoleScript != nil {
		m["changeRoleScript"] = func() map[string]interface{} {
			if p.ChangeRoleScript == nil {
				return nil
			}
			return p.ChangeRoleScript.ToDict()
		}()
	}
	if p.DeleteGuildScript != nil {
		m["deleteGuildScript"] = func() map[string]interface{} {
			if p.DeleteGuildScript == nil {
				return nil
			}
			return p.DeleteGuildScript.ToDict()
		}()
	}
	if p.LogSetting != nil {
		m["logSetting"] = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p Namespace) Pointer() *Namespace {
	return &p
}

func CastNamespaces(data []interface{}) []Namespace {
	v := make([]Namespace, 0)
	for _, d := range data {
		v = append(v, NewNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNamespacesFromDict(data []Namespace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GuildModelMaster struct {
	GuildModelId                  *string     `json:"guildModelId"`
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
	CreatedAt                     *int64      `json:"createdAt"`
	UpdatedAt                     *int64      `json:"updatedAt"`
	Revision                      *int64      `json:"revision"`
}

func (p *GuildModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GuildModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GuildModelMaster{}
	} else {
		*p = GuildModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["guildModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelId)
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewGuildModelMasterFromJson(data string) GuildModelMaster {
	req := GuildModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGuildModelMasterFromDict(data map[string]interface{}) GuildModelMaster {
	return GuildModelMaster{
		GuildModelId: func() *string {
			v, ok := data["guildModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelId"])
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
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p GuildModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GuildModelId != nil {
		m["guildModelId"] = p.GuildModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.DefaultMaximumMemberCount != nil {
		m["defaultMaximumMemberCount"] = p.DefaultMaximumMemberCount
	}
	if p.MaximumMemberCount != nil {
		m["maximumMemberCount"] = p.MaximumMemberCount
	}
	if p.InactivityPeriodDays != nil {
		m["inactivityPeriodDays"] = p.InactivityPeriodDays
	}
	if p.Roles != nil {
		m["roles"] = CastRoleModelsFromDict(
			p.Roles,
		)
	}
	if p.GuildMasterRole != nil {
		m["guildMasterRole"] = p.GuildMasterRole
	}
	if p.GuildMemberDefaultRole != nil {
		m["guildMemberDefaultRole"] = p.GuildMemberDefaultRole
	}
	if p.RejoinCoolTimeMinutes != nil {
		m["rejoinCoolTimeMinutes"] = p.RejoinCoolTimeMinutes
	}
	if p.MaxConcurrentJoinGuilds != nil {
		m["maxConcurrentJoinGuilds"] = p.MaxConcurrentJoinGuilds
	}
	if p.MaxConcurrentGuildMasterCount != nil {
		m["maxConcurrentGuildMasterCount"] = p.MaxConcurrentGuildMasterCount
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p GuildModelMaster) Pointer() *GuildModelMaster {
	return &p
}

func CastGuildModelMasters(data []interface{}) []GuildModelMaster {
	v := make([]GuildModelMaster, 0)
	for _, d := range data {
		v = append(v, NewGuildModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildModelMastersFromDict(data []GuildModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GuildModel struct {
	GuildModelId                  *string     `json:"guildModelId"`
	Name                          *string     `json:"name"`
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

func (p *GuildModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GuildModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GuildModel{}
	} else {
		*p = GuildModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["guildModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildModelId)
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

func NewGuildModelFromJson(data string) GuildModel {
	req := GuildModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGuildModelFromDict(data map[string]interface{}) GuildModel {
	return GuildModel{
		GuildModelId: func() *string {
			v, ok := data["guildModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
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

func (p GuildModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GuildModelId != nil {
		m["guildModelId"] = p.GuildModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.DefaultMaximumMemberCount != nil {
		m["defaultMaximumMemberCount"] = p.DefaultMaximumMemberCount
	}
	if p.MaximumMemberCount != nil {
		m["maximumMemberCount"] = p.MaximumMemberCount
	}
	if p.InactivityPeriodDays != nil {
		m["inactivityPeriodDays"] = p.InactivityPeriodDays
	}
	if p.Roles != nil {
		m["roles"] = CastRoleModelsFromDict(
			p.Roles,
		)
	}
	if p.GuildMasterRole != nil {
		m["guildMasterRole"] = p.GuildMasterRole
	}
	if p.GuildMemberDefaultRole != nil {
		m["guildMemberDefaultRole"] = p.GuildMemberDefaultRole
	}
	if p.RejoinCoolTimeMinutes != nil {
		m["rejoinCoolTimeMinutes"] = p.RejoinCoolTimeMinutes
	}
	if p.MaxConcurrentJoinGuilds != nil {
		m["maxConcurrentJoinGuilds"] = p.MaxConcurrentJoinGuilds
	}
	if p.MaxConcurrentGuildMasterCount != nil {
		m["maxConcurrentGuildMasterCount"] = p.MaxConcurrentGuildMasterCount
	}
	return m
}

func (p GuildModel) Pointer() *GuildModel {
	return &p
}

func CastGuildModels(data []interface{}) []GuildModel {
	v := make([]GuildModel, 0)
	for _, d := range data {
		v = append(v, NewGuildModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildModelsFromDict(data []GuildModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Guild struct {
	GuildId                   *string     `json:"guildId"`
	GuildModelName            *string     `json:"guildModelName"`
	Name                      *string     `json:"name"`
	DisplayName               *string     `json:"displayName"`
	Attribute1                *int32      `json:"attribute1"`
	Attribute2                *int32      `json:"attribute2"`
	Attribute3                *int32      `json:"attribute3"`
	Attribute4                *int32      `json:"attribute4"`
	Attribute5                *int32      `json:"attribute5"`
	Metadata                  *string     `json:"metadata"`
	JoinPolicy                *string     `json:"joinPolicy"`
	CustomRoles               []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	CurrentMaximumMemberCount *int32      `json:"currentMaximumMemberCount"`
	Members                   []Member    `json:"members"`
	CreatedAt                 *int64      `json:"createdAt"`
	UpdatedAt                 *int64      `json:"updatedAt"`
	Revision                  *int64      `json:"revision"`
}

func (p *Guild) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Guild{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Guild{}
	} else {
		*p = Guild{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["guildId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GuildId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GuildId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GuildId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GuildId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GuildId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GuildId)
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
		if v, ok := d["currentMaximumMemberCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentMaximumMemberCount)
		}
		if v, ok := d["members"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Members)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewGuildFromJson(data string) Guild {
	req := Guild{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGuildFromDict(data map[string]interface{}) Guild {
	return Guild{
		GuildId: func() *string {
			v, ok := data["guildId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildId"])
		}(),
		GuildModelName: func() *string {
			v, ok := data["guildModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["guildModelName"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
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
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
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
		CurrentMaximumMemberCount: func() *int32 {
			v, ok := data["currentMaximumMemberCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["currentMaximumMemberCount"])
		}(),
		Members: func() []Member {
			if data["members"] == nil {
				return nil
			}
			return CastMembers(core.CastArray(data["members"]))
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Guild) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GuildId != nil {
		m["guildId"] = p.GuildId
	}
	if p.GuildModelName != nil {
		m["guildModelName"] = p.GuildModelName
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.DisplayName != nil {
		m["displayName"] = p.DisplayName
	}
	if p.Attribute1 != nil {
		m["attribute1"] = p.Attribute1
	}
	if p.Attribute2 != nil {
		m["attribute2"] = p.Attribute2
	}
	if p.Attribute3 != nil {
		m["attribute3"] = p.Attribute3
	}
	if p.Attribute4 != nil {
		m["attribute4"] = p.Attribute4
	}
	if p.Attribute5 != nil {
		m["attribute5"] = p.Attribute5
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.JoinPolicy != nil {
		m["joinPolicy"] = p.JoinPolicy
	}
	if p.CustomRoles != nil {
		m["customRoles"] = CastRoleModelsFromDict(
			p.CustomRoles,
		)
	}
	if p.GuildMemberDefaultRole != nil {
		m["guildMemberDefaultRole"] = p.GuildMemberDefaultRole
	}
	if p.CurrentMaximumMemberCount != nil {
		m["currentMaximumMemberCount"] = p.CurrentMaximumMemberCount
	}
	if p.Members != nil {
		m["members"] = CastMembersFromDict(
			p.Members,
		)
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p Guild) Pointer() *Guild {
	return &p
}

func CastGuilds(data []interface{}) []Guild {
	v := make([]Guild, 0)
	for _, d := range data {
		v = append(v, NewGuildFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGuildsFromDict(data []Guild) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type JoinedGuild struct {
	JoinedGuildId  *string `json:"joinedGuildId"`
	GuildModelName *string `json:"guildModelName"`
	GuildName      *string `json:"guildName"`
	UserId         *string `json:"userId"`
	CreatedAt      *int64  `json:"createdAt"`
}

func (p *JoinedGuild) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = JoinedGuild{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = JoinedGuild{}
	} else {
		*p = JoinedGuild{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["joinedGuildId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinedGuildId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinedGuildId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinedGuildId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinedGuildId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinedGuildId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinedGuildId)
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewJoinedGuildFromJson(data string) JoinedGuild {
	req := JoinedGuild{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJoinedGuildFromDict(data map[string]interface{}) JoinedGuild {
	return JoinedGuild{
		JoinedGuildId: func() *string {
			v, ok := data["joinedGuildId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinedGuildId"])
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
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p JoinedGuild) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.JoinedGuildId != nil {
		m["joinedGuildId"] = p.JoinedGuildId
	}
	if p.GuildModelName != nil {
		m["guildModelName"] = p.GuildModelName
	}
	if p.GuildName != nil {
		m["guildName"] = p.GuildName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p JoinedGuild) Pointer() *JoinedGuild {
	return &p
}

func CastJoinedGuilds(data []interface{}) []JoinedGuild {
	v := make([]JoinedGuild, 0)
	for _, d := range data {
		v = append(v, NewJoinedGuildFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJoinedGuildsFromDict(data []JoinedGuild) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LastGuildMasterActivity struct {
	UserId    *string `json:"userId"`
	UpdatedAt *int64  `json:"updatedAt"`
}

func (p *LastGuildMasterActivity) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LastGuildMasterActivity{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LastGuildMasterActivity{}
	} else {
		*p = LastGuildMasterActivity{}
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
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewLastGuildMasterActivityFromJson(data string) LastGuildMasterActivity {
	req := LastGuildMasterActivity{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLastGuildMasterActivityFromDict(data map[string]interface{}) LastGuildMasterActivity {
	return LastGuildMasterActivity{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
	}
}

func (p LastGuildMasterActivity) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p LastGuildMasterActivity) Pointer() *LastGuildMasterActivity {
	return &p
}

func CastLastGuildMasterActivities(data []interface{}) []LastGuildMasterActivity {
	v := make([]LastGuildMasterActivity, 0)
	for _, d := range data {
		v = append(v, NewLastGuildMasterActivityFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLastGuildMasterActivitiesFromDict(data []LastGuildMasterActivity) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentGuildMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentGuildMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentGuildMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentGuildMaster{}
	} else {
		*p = CurrentGuildMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
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

func NewCurrentGuildMasterFromJson(data string) CurrentGuildMaster {
	req := CurrentGuildMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentGuildMasterFromDict(data map[string]interface{}) CurrentGuildMaster {
	return CurrentGuildMaster{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
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

func (p CurrentGuildMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentGuildMaster) Pointer() *CurrentGuildMaster {
	return &p
}

func CastCurrentGuildMasters(data []interface{}) []CurrentGuildMaster {
	v := make([]CurrentGuildMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentGuildMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentGuildMastersFromDict(data []CurrentGuildMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RoleModel struct {
	Name           *string `json:"name"`
	Metadata       *string `json:"metadata"`
	PolicyDocument *string `json:"policyDocument"`
}

func (p *RoleModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RoleModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RoleModel{}
	} else {
		*p = RoleModel{}
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
		if v, ok := d["policyDocument"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PolicyDocument = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PolicyDocument = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PolicyDocument = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PolicyDocument = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PolicyDocument = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PolicyDocument)
				}
			}
		}
	}
	return nil
}

func NewRoleModelFromJson(data string) RoleModel {
	req := RoleModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRoleModelFromDict(data map[string]interface{}) RoleModel {
	return RoleModel{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		PolicyDocument: func() *string {
			v, ok := data["policyDocument"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["policyDocument"])
		}(),
	}
}

func (p RoleModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.PolicyDocument != nil {
		m["policyDocument"] = p.PolicyDocument
	}
	return m
}

func (p RoleModel) Pointer() *RoleModel {
	return &p
}

func CastRoleModels(data []interface{}) []RoleModel {
	v := make([]RoleModel, 0)
	for _, d := range data {
		v = append(v, NewRoleModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRoleModelsFromDict(data []RoleModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Member struct {
	UserId   *string `json:"userId"`
	RoleName *string `json:"roleName"`
	Metadata *string `json:"metadata"`
	JoinedAt *int64  `json:"joinedAt"`
}

func (p *Member) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Member{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Member{}
	} else {
		*p = Member{}
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
		if v, ok := d["joinedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinedAt)
		}
	}
	return nil
}

func NewMemberFromJson(data string) Member {
	req := Member{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMemberFromDict(data map[string]interface{}) Member {
	return Member{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		RoleName: func() *string {
			v, ok := data["roleName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["roleName"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		JoinedAt: func() *int64 {
			v, ok := data["joinedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["joinedAt"])
		}(),
	}
}

func (p Member) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.RoleName != nil {
		m["roleName"] = p.RoleName
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.JoinedAt != nil {
		m["joinedAt"] = p.JoinedAt
	}
	return m
}

func (p Member) Pointer() *Member {
	return &p
}

func CastMembers(data []interface{}) []Member {
	v := make([]Member, 0)
	for _, d := range data {
		v = append(v, NewMemberFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMembersFromDict(data []Member) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReceiveMemberRequest struct {
	UserId          *string `json:"userId"`
	TargetGuildName *string `json:"targetGuildName"`
	Metadata        *string `json:"metadata"`
	CreatedAt       *int64  `json:"createdAt"`
}

func (p *ReceiveMemberRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveMemberRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveMemberRequest{}
	} else {
		*p = ReceiveMemberRequest{}
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewReceiveMemberRequestFromJson(data string) ReceiveMemberRequest {
	req := ReceiveMemberRequest{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewReceiveMemberRequestFromDict(data map[string]interface{}) ReceiveMemberRequest {
	return ReceiveMemberRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p ReceiveMemberRequest) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetGuildName != nil {
		m["targetGuildName"] = p.TargetGuildName
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p ReceiveMemberRequest) Pointer() *ReceiveMemberRequest {
	return &p
}

func CastReceiveMemberRequests(data []interface{}) []ReceiveMemberRequest {
	v := make([]ReceiveMemberRequest, 0)
	for _, d := range data {
		v = append(v, NewReceiveMemberRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiveMemberRequestsFromDict(data []ReceiveMemberRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendMemberRequest struct {
	UserId          *string `json:"userId"`
	TargetGuildName *string `json:"targetGuildName"`
	Metadata        *string `json:"metadata"`
	CreatedAt       *int64  `json:"createdAt"`
}

func (p *SendMemberRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SendMemberRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SendMemberRequest{}
	} else {
		*p = SendMemberRequest{}
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewSendMemberRequestFromJson(data string) SendMemberRequest {
	req := SendMemberRequest{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSendMemberRequestFromDict(data map[string]interface{}) SendMemberRequest {
	return SendMemberRequest{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetGuildName: func() *string {
			v, ok := data["targetGuildName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGuildName"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p SendMemberRequest) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetGuildName != nil {
		m["targetGuildName"] = p.TargetGuildName
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p SendMemberRequest) Pointer() *SendMemberRequest {
	return &p
}

func CastSendMemberRequests(data []interface{}) []SendMemberRequest {
	v := make([]SendMemberRequest, 0)
	for _, d := range data {
		v = append(v, NewSendMemberRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendMemberRequestsFromDict(data []SendMemberRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type IgnoreUser struct {
	UserId    *string `json:"userId"`
	CreatedAt *int64  `json:"createdAt"`
}

func (p *IgnoreUser) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IgnoreUser{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = IgnoreUser{}
	} else {
		*p = IgnoreUser{}
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewIgnoreUserFromJson(data string) IgnoreUser {
	req := IgnoreUser{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewIgnoreUserFromDict(data map[string]interface{}) IgnoreUser {
	return IgnoreUser{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p IgnoreUser) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p IgnoreUser) Pointer() *IgnoreUser {
	return &p
}

func CastIgnoreUsers(data []interface{}) []IgnoreUser {
	v := make([]IgnoreUser, 0)
	for _, d := range data {
		v = append(v, NewIgnoreUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIgnoreUsersFromDict(data []IgnoreUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VerifyActionResult struct {
	Action        *string `json:"action"`
	VerifyRequest *string `json:"verifyRequest"`
	StatusCode    *int32  `json:"statusCode"`
	VerifyResult  *string `json:"verifyResult"`
}

func (p *VerifyActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyActionResult{}
	} else {
		*p = VerifyActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["verifyRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["verifyResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyResult)
				}
			}
		}
	}
	return nil
}

func NewVerifyActionResultFromJson(data string) VerifyActionResult {
	req := VerifyActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVerifyActionResultFromDict(data map[string]interface{}) VerifyActionResult {
	return VerifyActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		VerifyRequest: func() *string {
			v, ok := data["verifyRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		VerifyResult: func() *string {
			v, ok := data["verifyResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyResult"])
		}(),
	}
}

func (p VerifyActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.VerifyRequest != nil {
		m["verifyRequest"] = p.VerifyRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.VerifyResult != nil {
		m["verifyResult"] = p.VerifyResult
	}
	return m
}

func (p VerifyActionResult) Pointer() *VerifyActionResult {
	return &p
}

func CastVerifyActionResults(data []interface{}) []VerifyActionResult {
	v := make([]VerifyActionResult, 0)
	for _, d := range data {
		v = append(v, NewVerifyActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVerifyActionResultsFromDict(data []VerifyActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeActionResult struct {
	Action         *string `json:"action"`
	ConsumeRequest *string `json:"consumeRequest"`
	StatusCode     *int32  `json:"statusCode"`
	ConsumeResult  *string `json:"consumeResult"`
}

func (p *ConsumeActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeActionResult{}
	} else {
		*p = ConsumeActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["consumeRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["consumeResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeResult)
				}
			}
		}
	}
	return nil
}

func NewConsumeActionResultFromJson(data string) ConsumeActionResult {
	req := ConsumeActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeActionResultFromDict(data map[string]interface{}) ConsumeActionResult {
	return ConsumeActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		ConsumeRequest: func() *string {
			v, ok := data["consumeRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		ConsumeResult: func() *string {
			v, ok := data["consumeResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeResult"])
		}(),
	}
}

func (p ConsumeActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.ConsumeRequest != nil {
		m["consumeRequest"] = p.ConsumeRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.ConsumeResult != nil {
		m["consumeResult"] = p.ConsumeResult
	}
	return m
}

func (p ConsumeActionResult) Pointer() *ConsumeActionResult {
	return &p
}

func CastConsumeActionResults(data []interface{}) []ConsumeActionResult {
	v := make([]ConsumeActionResult, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionResultsFromDict(data []ConsumeActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireActionResult struct {
	Action         *string `json:"action"`
	AcquireRequest *string `json:"acquireRequest"`
	StatusCode     *int32  `json:"statusCode"`
	AcquireResult  *string `json:"acquireResult"`
}

func (p *AcquireActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireActionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcquireActionResult{}
	} else {
		*p = AcquireActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["acquireRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["acquireResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireResult)
				}
			}
		}
	}
	return nil
}

func NewAcquireActionResultFromJson(data string) AcquireActionResult {
	req := AcquireActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireActionResultFromDict(data map[string]interface{}) AcquireActionResult {
	return AcquireActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		AcquireRequest: func() *string {
			v, ok := data["acquireRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		AcquireResult: func() *string {
			v, ok := data["acquireResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireResult"])
		}(),
	}
}

func (p AcquireActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.AcquireRequest != nil {
		m["acquireRequest"] = p.AcquireRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.AcquireResult != nil {
		m["acquireResult"] = p.AcquireResult
	}
	return m
}

func (p AcquireActionResult) Pointer() *AcquireActionResult {
	return &p
}

func CastAcquireActionResults(data []interface{}) []AcquireActionResult {
	v := make([]AcquireActionResult, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionResultsFromDict(data []AcquireActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionResult struct {
	TransactionId  *string               `json:"transactionId"`
	VerifyResults  []VerifyActionResult  `json:"verifyResults"`
	ConsumeResults []ConsumeActionResult `json:"consumeResults"`
	AcquireResults []AcquireActionResult `json:"acquireResults"`
	HasError       *bool                 `json:"hasError"`
}

func (p *TransactionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TransactionResult{}
	} else {
		*p = TransactionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["verifyResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyResults)
		}
		if v, ok := d["consumeResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeResults)
		}
		if v, ok := d["acquireResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireResults)
		}
		if v, ok := d["hasError"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.HasError)
		}
	}
	return nil
}

func NewTransactionResultFromJson(data string) TransactionResult {
	req := TransactionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionResultFromDict(data map[string]interface{}) TransactionResult {
	return TransactionResult{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		VerifyResults: func() []VerifyActionResult {
			if data["verifyResults"] == nil {
				return nil
			}
			return CastVerifyActionResults(core.CastArray(data["verifyResults"]))
		}(),
		ConsumeResults: func() []ConsumeActionResult {
			if data["consumeResults"] == nil {
				return nil
			}
			return CastConsumeActionResults(core.CastArray(data["consumeResults"]))
		}(),
		AcquireResults: func() []AcquireActionResult {
			if data["acquireResults"] == nil {
				return nil
			}
			return CastAcquireActionResults(core.CastArray(data["acquireResults"]))
		}(),
		HasError: func() *bool {
			v, ok := data["hasError"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["hasError"])
		}(),
	}
}

func (p TransactionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.VerifyResults != nil {
		m["verifyResults"] = CastVerifyActionResultsFromDict(
			p.VerifyResults,
		)
	}
	if p.ConsumeResults != nil {
		m["consumeResults"] = CastConsumeActionResultsFromDict(
			p.ConsumeResults,
		)
	}
	if p.AcquireResults != nil {
		m["acquireResults"] = CastAcquireActionResultsFromDict(
			p.AcquireResults,
		)
	}
	if p.HasError != nil {
		m["hasError"] = p.HasError
	}
	return m
}

func (p TransactionResult) Pointer() *TransactionResult {
	return &p
}

func CastTransactionResults(data []interface{}) []TransactionResult {
	v := make([]TransactionResult, 0)
	for _, d := range data {
		v = append(v, NewTransactionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionResultsFromDict(data []TransactionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun             *bool   `json:"enableAutoRun"`
	EnableAtomicCommit        *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor *bool   `json:"transactionUseDistributor"`
	AcquireActionUseJobQueue  *bool   `json:"acquireActionUseJobQueue"`
	DistributorNamespaceId    *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func (p *TransactionSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TransactionSetting{}
	} else {
		*p = TransactionSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["enableAutoRun"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAutoRun)
		}
		if v, ok := d["enableAtomicCommit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAtomicCommit)
		}
		if v, ok := d["transactionUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionUseDistributor)
		}
		if v, ok := d["acquireActionUseJobQueue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionUseJobQueue)
		}
		if v, ok := d["distributorNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DistributorNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DistributorNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DistributorNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DistributorNamespaceId)
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
		if v, ok := d["queueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	req := TransactionSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun: func() *bool {
			v, ok := data["enableAutoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAutoRun"])
		}(),
		EnableAtomicCommit: func() *bool {
			v, ok := data["enableAtomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAtomicCommit"])
		}(),
		TransactionUseDistributor: func() *bool {
			v, ok := data["transactionUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["transactionUseDistributor"])
		}(),
		AcquireActionUseJobQueue: func() *bool {
			v, ok := data["acquireActionUseJobQueue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["acquireActionUseJobQueue"])
		}(),
		DistributorNamespaceId: func() *string {
			v, ok := data["distributorNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["distributorNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EnableAutoRun != nil {
		m["enableAutoRun"] = p.EnableAutoRun
	}
	if p.EnableAtomicCommit != nil {
		m["enableAtomicCommit"] = p.EnableAtomicCommit
	}
	if p.TransactionUseDistributor != nil {
		m["transactionUseDistributor"] = p.TransactionUseDistributor
	}
	if p.AcquireActionUseJobQueue != nil {
		m["acquireActionUseJobQueue"] = p.AcquireActionUseJobQueue
	}
	if p.DistributorNamespaceId != nil {
		m["distributorNamespaceId"] = p.DistributorNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	return m
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NotificationSetting struct {
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func (p *NotificationSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NotificationSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = NotificationSetting{}
	} else {
		*p = NotificationSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["gatewayNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatewayNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatewayNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatewayNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatewayNamespaceId)
				}
			}
		}
		if v, ok := d["enableTransferMobileNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableTransferMobileNotification)
		}
		if v, ok := d["sound"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Sound = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Sound = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Sound = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Sound)
				}
			}
		}
	}
	return nil
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	req := NotificationSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId: func() *string {
			v, ok := data["gatewayNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatewayNamespaceId"])
		}(),
		EnableTransferMobileNotification: func() *bool {
			v, ok := data["enableTransferMobileNotification"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableTransferMobileNotification"])
		}(),
		Sound: func() *string {
			v, ok := data["sound"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sound"])
		}(),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GatewayNamespaceId != nil {
		m["gatewayNamespaceId"] = p.GatewayNamespaceId
	}
	if p.EnableTransferMobileNotification != nil {
		m["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
	}
	if p.Sound != nil {
		m["sound"] = p.Sound
	}
	return m
}

func (p NotificationSetting) Pointer() *NotificationSetting {
	return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScriptSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ScriptSetting{}
	} else {
		*p = ScriptSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["triggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerTargetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerTargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerTargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerTargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerTargetType)
				}
			}
		}
		if v, ok := d["doneTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerQueueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerQueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerQueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	req := ScriptSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId: func() *string {
			v, ok := data["triggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerScriptId"])
		}(),
		DoneTriggerTargetType: func() *string {
			v, ok := data["doneTriggerTargetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerTargetType"])
		}(),
		DoneTriggerScriptId: func() *string {
			v, ok := data["doneTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerScriptId"])
		}(),
		DoneTriggerQueueNamespaceId: func() *string {
			v, ok := data["doneTriggerQueueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerQueueNamespaceId"])
		}(),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		m["triggerScriptId"] = p.TriggerScriptId
	}
	if p.DoneTriggerTargetType != nil {
		m["doneTriggerTargetType"] = p.DoneTriggerTargetType
	}
	if p.DoneTriggerScriptId != nil {
		m["doneTriggerScriptId"] = p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		m["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
	}
	return m
}

func (p ScriptSetting) Pointer() *ScriptSetting {
	return &p
}

func CastScriptSettings(data []interface{}) []ScriptSetting {
	v := make([]ScriptSetting, 0)
	for _, d := range data {
		v = append(v, NewScriptSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScriptSettingsFromDict(data []ScriptSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LogSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LogSetting{}
	} else {
		*p = LogSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["loggingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LoggingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LoggingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LoggingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LoggingNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewLogSettingFromJson(data string) LogSetting {
	req := LogSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: func() *string {
			v, ok := data["loggingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["loggingNamespaceId"])
		}(),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LoggingNamespaceId != nil {
		m["loggingNamespaceId"] = p.LoggingNamespaceId
	}
	return m
}

func (p LogSetting) Pointer() *LogSetting {
	return &p
}

func CastLogSettings(data []interface{}) []LogSetting {
	v := make([]LogSetting, 0)
	for _, d := range data {
		v = append(v, NewLogSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogSettingsFromDict(data []LogSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func (p *GitHubCheckoutSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GitHubCheckoutSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GitHubCheckoutSetting{}
	} else {
		*p = GitHubCheckoutSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["apiKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApiKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApiKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApiKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApiKeyId)
				}
			}
		}
		if v, ok := d["repositoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepositoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepositoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepositoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepositoryName)
				}
			}
		}
		if v, ok := d["sourcePath"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourcePath = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourcePath = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourcePath = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourcePath)
				}
			}
		}
		if v, ok := d["referenceType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceType)
				}
			}
		}
		if v, ok := d["commitHash"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CommitHash = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CommitHash = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CommitHash = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CommitHash)
				}
			}
		}
		if v, ok := d["branchName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BranchName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BranchName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BranchName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BranchName)
				}
			}
		}
		if v, ok := d["tagName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TagName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TagName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TagName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TagName)
				}
			}
		}
	}
	return nil
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	req := GitHubCheckoutSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId: func() *string {
			v, ok := data["apiKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["apiKeyId"])
		}(),
		RepositoryName: func() *string {
			v, ok := data["repositoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repositoryName"])
		}(),
		SourcePath: func() *string {
			v, ok := data["sourcePath"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourcePath"])
		}(),
		ReferenceType: func() *string {
			v, ok := data["referenceType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["referenceType"])
		}(),
		CommitHash: func() *string {
			v, ok := data["commitHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["commitHash"])
		}(),
		BranchName: func() *string {
			v, ok := data["branchName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["branchName"])
		}(),
		TagName: func() *string {
			v, ok := data["tagName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["tagName"])
		}(),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ApiKeyId != nil {
		m["apiKeyId"] = p.ApiKeyId
	}
	if p.RepositoryName != nil {
		m["repositoryName"] = p.RepositoryName
	}
	if p.SourcePath != nil {
		m["sourcePath"] = p.SourcePath
	}
	if p.ReferenceType != nil {
		m["referenceType"] = p.ReferenceType
	}
	if p.CommitHash != nil {
		m["commitHash"] = p.CommitHash
	}
	if p.BranchName != nil {
		m["branchName"] = p.BranchName
	}
	if p.TagName != nil {
		m["tagName"] = p.TagName
	}
	return m
}

func (p GitHubCheckoutSetting) Pointer() *GitHubCheckoutSetting {
	return &p
}

func CastGitHubCheckoutSettings(data []interface{}) []GitHubCheckoutSetting {
	v := make([]GitHubCheckoutSetting, 0)
	for _, d := range data {
		v = append(v, NewGitHubCheckoutSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubCheckoutSettingsFromDict(data []GitHubCheckoutSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
