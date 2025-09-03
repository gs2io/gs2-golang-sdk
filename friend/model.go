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

package friend

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
	FollowScript               *ScriptSetting       `json:"followScript"`
	UnfollowScript             *ScriptSetting       `json:"unfollowScript"`
	SendRequestScript          *ScriptSetting       `json:"sendRequestScript"`
	CancelRequestScript        *ScriptSetting       `json:"cancelRequestScript"`
	AcceptRequestScript        *ScriptSetting       `json:"acceptRequestScript"`
	RejectRequestScript        *ScriptSetting       `json:"rejectRequestScript"`
	DeleteFriendScript         *ScriptSetting       `json:"deleteFriendScript"`
	UpdateProfileScript        *ScriptSetting       `json:"updateProfileScript"`
	FollowNotification         *NotificationSetting `json:"followNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	CancelRequestNotification  *NotificationSetting `json:"cancelRequestNotification"`
	AcceptRequestNotification  *NotificationSetting `json:"acceptRequestNotification"`
	RejectRequestNotification  *NotificationSetting `json:"rejectRequestNotification"`
	DeleteFriendNotification   *NotificationSetting `json:"deleteFriendNotification"`
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
		if v, ok := d["followScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FollowScript)
		}
		if v, ok := d["unfollowScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UnfollowScript)
		}
		if v, ok := d["sendRequestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SendRequestScript)
		}
		if v, ok := d["cancelRequestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CancelRequestScript)
		}
		if v, ok := d["acceptRequestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcceptRequestScript)
		}
		if v, ok := d["rejectRequestScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RejectRequestScript)
		}
		if v, ok := d["deleteFriendScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DeleteFriendScript)
		}
		if v, ok := d["updateProfileScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdateProfileScript)
		}
		if v, ok := d["followNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.FollowNotification)
		}
		if v, ok := d["receiveRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveRequestNotification)
		}
		if v, ok := d["cancelRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CancelRequestNotification)
		}
		if v, ok := d["acceptRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcceptRequestNotification)
		}
		if v, ok := d["rejectRequestNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RejectRequestNotification)
		}
		if v, ok := d["deleteFriendNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DeleteFriendNotification)
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
		FollowScript: func() *ScriptSetting {
			v, ok := data["followScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["followScript"])).Pointer()
		}(),
		UnfollowScript: func() *ScriptSetting {
			v, ok := data["unfollowScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["unfollowScript"])).Pointer()
		}(),
		SendRequestScript: func() *ScriptSetting {
			v, ok := data["sendRequestScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["sendRequestScript"])).Pointer()
		}(),
		CancelRequestScript: func() *ScriptSetting {
			v, ok := data["cancelRequestScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["cancelRequestScript"])).Pointer()
		}(),
		AcceptRequestScript: func() *ScriptSetting {
			v, ok := data["acceptRequestScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["acceptRequestScript"])).Pointer()
		}(),
		RejectRequestScript: func() *ScriptSetting {
			v, ok := data["rejectRequestScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["rejectRequestScript"])).Pointer()
		}(),
		DeleteFriendScript: func() *ScriptSetting {
			v, ok := data["deleteFriendScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["deleteFriendScript"])).Pointer()
		}(),
		UpdateProfileScript: func() *ScriptSetting {
			v, ok := data["updateProfileScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["updateProfileScript"])).Pointer()
		}(),
		FollowNotification: func() *NotificationSetting {
			v, ok := data["followNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["followNotification"])).Pointer()
		}(),
		ReceiveRequestNotification: func() *NotificationSetting {
			v, ok := data["receiveRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer()
		}(),
		CancelRequestNotification: func() *NotificationSetting {
			v, ok := data["cancelRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["cancelRequestNotification"])).Pointer()
		}(),
		AcceptRequestNotification: func() *NotificationSetting {
			v, ok := data["acceptRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["acceptRequestNotification"])).Pointer()
		}(),
		RejectRequestNotification: func() *NotificationSetting {
			v, ok := data["rejectRequestNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["rejectRequestNotification"])).Pointer()
		}(),
		DeleteFriendNotification: func() *NotificationSetting {
			v, ok := data["deleteFriendNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["deleteFriendNotification"])).Pointer()
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
	if p.FollowScript != nil {
		m["followScript"] = func() map[string]interface{} {
			if p.FollowScript == nil {
				return nil
			}
			return p.FollowScript.ToDict()
		}()
	}
	if p.UnfollowScript != nil {
		m["unfollowScript"] = func() map[string]interface{} {
			if p.UnfollowScript == nil {
				return nil
			}
			return p.UnfollowScript.ToDict()
		}()
	}
	if p.SendRequestScript != nil {
		m["sendRequestScript"] = func() map[string]interface{} {
			if p.SendRequestScript == nil {
				return nil
			}
			return p.SendRequestScript.ToDict()
		}()
	}
	if p.CancelRequestScript != nil {
		m["cancelRequestScript"] = func() map[string]interface{} {
			if p.CancelRequestScript == nil {
				return nil
			}
			return p.CancelRequestScript.ToDict()
		}()
	}
	if p.AcceptRequestScript != nil {
		m["acceptRequestScript"] = func() map[string]interface{} {
			if p.AcceptRequestScript == nil {
				return nil
			}
			return p.AcceptRequestScript.ToDict()
		}()
	}
	if p.RejectRequestScript != nil {
		m["rejectRequestScript"] = func() map[string]interface{} {
			if p.RejectRequestScript == nil {
				return nil
			}
			return p.RejectRequestScript.ToDict()
		}()
	}
	if p.DeleteFriendScript != nil {
		m["deleteFriendScript"] = func() map[string]interface{} {
			if p.DeleteFriendScript == nil {
				return nil
			}
			return p.DeleteFriendScript.ToDict()
		}()
	}
	if p.UpdateProfileScript != nil {
		m["updateProfileScript"] = func() map[string]interface{} {
			if p.UpdateProfileScript == nil {
				return nil
			}
			return p.UpdateProfileScript.ToDict()
		}()
	}
	if p.FollowNotification != nil {
		m["followNotification"] = func() map[string]interface{} {
			if p.FollowNotification == nil {
				return nil
			}
			return p.FollowNotification.ToDict()
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
	if p.CancelRequestNotification != nil {
		m["cancelRequestNotification"] = func() map[string]interface{} {
			if p.CancelRequestNotification == nil {
				return nil
			}
			return p.CancelRequestNotification.ToDict()
		}()
	}
	if p.AcceptRequestNotification != nil {
		m["acceptRequestNotification"] = func() map[string]interface{} {
			if p.AcceptRequestNotification == nil {
				return nil
			}
			return p.AcceptRequestNotification.ToDict()
		}()
	}
	if p.RejectRequestNotification != nil {
		m["rejectRequestNotification"] = func() map[string]interface{} {
			if p.RejectRequestNotification == nil {
				return nil
			}
			return p.RejectRequestNotification.ToDict()
		}()
	}
	if p.DeleteFriendNotification != nil {
		m["deleteFriendNotification"] = func() map[string]interface{} {
			if p.DeleteFriendNotification == nil {
				return nil
			}
			return p.DeleteFriendNotification.ToDict()
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

type Profile struct {
	ProfileId       *string `json:"profileId"`
	UserId          *string `json:"userId"`
	PublicProfile   *string `json:"publicProfile"`
	FollowerProfile *string `json:"followerProfile"`
	FriendProfile   *string `json:"friendProfile"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
	Revision        *int64  `json:"revision"`
}

func (p *Profile) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Profile{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Profile{}
	} else {
		*p = Profile{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["profileId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ProfileId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ProfileId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ProfileId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ProfileId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ProfileId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ProfileId)
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
		if v, ok := d["followerProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FollowerProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FollowerProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FollowerProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FollowerProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FollowerProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FollowerProfile)
				}
			}
		}
		if v, ok := d["friendProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FriendProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FriendProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FriendProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FriendProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FriendProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FriendProfile)
				}
			}
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

func NewProfileFromJson(data string) Profile {
	req := Profile{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewProfileFromDict(data map[string]interface{}) Profile {
	return Profile{
		ProfileId: func() *string {
			v, ok := data["profileId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["profileId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
		FollowerProfile: func() *string {
			v, ok := data["followerProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["followerProfile"])
		}(),
		FriendProfile: func() *string {
			v, ok := data["friendProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["friendProfile"])
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

func (p Profile) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ProfileId != nil {
		m["profileId"] = p.ProfileId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	if p.FollowerProfile != nil {
		m["followerProfile"] = p.FollowerProfile
	}
	if p.FriendProfile != nil {
		m["friendProfile"] = p.FriendProfile
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

func (p Profile) Pointer() *Profile {
	return &p
}

func CastProfiles(data []interface{}) []Profile {
	v := make([]Profile, 0)
	for _, d := range data {
		v = append(v, NewProfileFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProfilesFromDict(data []Profile) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Follow struct {
	FollowId      *string   `json:"followId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func (p *Follow) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Follow{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Follow{}
	} else {
		*p = Follow{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["followId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FollowId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FollowId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FollowId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FollowId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FollowId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FollowId)
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
		if v, ok := d["targetUserIds"]; ok && v != nil {
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
				p.TargetUserIds = l
			}
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

func NewFollowFromJson(data string) Follow {
	req := Follow{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFollowFromDict(data map[string]interface{}) Follow {
	return Follow{
		FollowId: func() *string {
			v, ok := data["followId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["followId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserIds: func() []*string {
			v, ok := data["targetUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p Follow) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.FollowId != nil {
		m["followId"] = p.FollowId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserIds != nil {
		m["targetUserIds"] = core.CastStringsFromDict(
			p.TargetUserIds,
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

func (p Follow) Pointer() *Follow {
	return &p
}

func CastFollows(data []interface{}) []Follow {
	v := make([]Follow, 0)
	for _, d := range data {
		v = append(v, NewFollowFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFollowsFromDict(data []Follow) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Friend struct {
	FriendId      *string   `json:"friendId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func (p *Friend) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Friend{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Friend{}
	} else {
		*p = Friend{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["friendId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FriendId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FriendId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FriendId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FriendId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FriendId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FriendId)
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
		if v, ok := d["targetUserIds"]; ok && v != nil {
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
				p.TargetUserIds = l
			}
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

func NewFriendFromJson(data string) Friend {
	req := Friend{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFriendFromDict(data map[string]interface{}) Friend {
	return Friend{
		FriendId: func() *string {
			v, ok := data["friendId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["friendId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserIds: func() []*string {
			v, ok := data["targetUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p Friend) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.FriendId != nil {
		m["friendId"] = p.FriendId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserIds != nil {
		m["targetUserIds"] = core.CastStringsFromDict(
			p.TargetUserIds,
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

func (p Friend) Pointer() *Friend {
	return &p
}

func CastFriends(data []interface{}) []Friend {
	v := make([]Friend, 0)
	for _, d := range data {
		v = append(v, NewFriendFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendsFromDict(data []Friend) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendBox struct {
	SendBoxId     *string   `json:"sendBoxId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func (p *SendBox) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SendBox{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SendBox{}
	} else {
		*p = SendBox{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["sendBoxId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SendBoxId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SendBoxId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SendBoxId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SendBoxId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SendBoxId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SendBoxId)
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
		if v, ok := d["targetUserIds"]; ok && v != nil {
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
				p.TargetUserIds = l
			}
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

func NewSendBoxFromJson(data string) SendBox {
	req := SendBox{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSendBoxFromDict(data map[string]interface{}) SendBox {
	return SendBox{
		SendBoxId: func() *string {
			v, ok := data["sendBoxId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sendBoxId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserIds: func() []*string {
			v, ok := data["targetUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p SendBox) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SendBoxId != nil {
		m["sendBoxId"] = p.SendBoxId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserIds != nil {
		m["targetUserIds"] = core.CastStringsFromDict(
			p.TargetUserIds,
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

func (p SendBox) Pointer() *SendBox {
	return &p
}

func CastSendBoxes(data []interface{}) []SendBox {
	v := make([]SendBox, 0)
	for _, d := range data {
		v = append(v, NewSendBoxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendBoxesFromDict(data []SendBox) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Inbox struct {
	InboxId     *string   `json:"inboxId"`
	UserId      *string   `json:"userId"`
	FromUserIds []*string `json:"fromUserIds"`
	CreatedAt   *int64    `json:"createdAt"`
	UpdatedAt   *int64    `json:"updatedAt"`
	Revision    *int64    `json:"revision"`
}

func (p *Inbox) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Inbox{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Inbox{}
	} else {
		*p = Inbox{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inboxId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InboxId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InboxId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InboxId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InboxId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InboxId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InboxId)
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
		if v, ok := d["fromUserIds"]; ok && v != nil {
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
				p.FromUserIds = l
			}
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

func NewInboxFromJson(data string) Inbox {
	req := Inbox{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInboxFromDict(data map[string]interface{}) Inbox {
	return Inbox{
		InboxId: func() *string {
			v, ok := data["inboxId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inboxId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		FromUserIds: func() []*string {
			v, ok := data["fromUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p Inbox) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InboxId != nil {
		m["inboxId"] = p.InboxId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.FromUserIds != nil {
		m["fromUserIds"] = core.CastStringsFromDict(
			p.FromUserIds,
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

func (p Inbox) Pointer() *Inbox {
	return &p
}

func CastInboxes(data []interface{}) []Inbox {
	v := make([]Inbox, 0)
	for _, d := range data {
		v = append(v, NewInboxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInboxesFromDict(data []Inbox) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BlackList struct {
	BlackListId   *string   `json:"blackListId"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func (p *BlackList) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BlackList{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BlackList{}
	} else {
		*p = BlackList{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["blackListId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BlackListId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BlackListId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BlackListId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BlackListId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BlackListId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BlackListId)
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
		if v, ok := d["targetUserIds"]; ok && v != nil {
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
				p.TargetUserIds = l
			}
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

func NewBlackListFromJson(data string) BlackList {
	req := BlackList{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBlackListFromDict(data map[string]interface{}) BlackList {
	return BlackList{
		BlackListId: func() *string {
			v, ok := data["blackListId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["blackListId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserIds: func() []*string {
			v, ok := data["targetUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p BlackList) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BlackListId != nil {
		m["blackListId"] = p.BlackListId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserIds != nil {
		m["targetUserIds"] = core.CastStringsFromDict(
			p.TargetUserIds,
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

func (p BlackList) Pointer() *BlackList {
	return &p
}

func CastBlackList(data []interface{}) []BlackList {
	v := make([]BlackList, 0)
	for _, d := range data {
		v = append(v, NewBlackListFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBlackListFromDict(data []BlackList) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FollowUser struct {
	UserId          *string `json:"userId"`
	PublicProfile   *string `json:"publicProfile"`
	FollowerProfile *string `json:"followerProfile"`
}

func (p *FollowUser) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FollowUser{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = FollowUser{}
	} else {
		*p = FollowUser{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
		if v, ok := d["followerProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FollowerProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FollowerProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FollowerProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FollowerProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FollowerProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FollowerProfile)
				}
			}
		}
	}
	return nil
}

func NewFollowUserFromJson(data string) FollowUser {
	req := FollowUser{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFollowUserFromDict(data map[string]interface{}) FollowUser {
	return FollowUser{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
		FollowerProfile: func() *string {
			v, ok := data["followerProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["followerProfile"])
		}(),
	}
}

func (p FollowUser) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	if p.FollowerProfile != nil {
		m["followerProfile"] = p.FollowerProfile
	}
	return m
}

func (p FollowUser) Pointer() *FollowUser {
	return &p
}

func CastFollowUsers(data []interface{}) []FollowUser {
	v := make([]FollowUser, 0)
	for _, d := range data {
		v = append(v, NewFollowUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFollowUsersFromDict(data []FollowUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendUser struct {
	UserId        *string `json:"userId"`
	PublicProfile *string `json:"publicProfile"`
	FriendProfile *string `json:"friendProfile"`
}

func (p *FriendUser) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FriendUser{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = FriendUser{}
	} else {
		*p = FriendUser{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
		if v, ok := d["friendProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FriendProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FriendProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FriendProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FriendProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FriendProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FriendProfile)
				}
			}
		}
	}
	return nil
}

func NewFriendUserFromJson(data string) FriendUser {
	req := FriendUser{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFriendUserFromDict(data map[string]interface{}) FriendUser {
	return FriendUser{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
		FriendProfile: func() *string {
			v, ok := data["friendProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["friendProfile"])
		}(),
	}
}

func (p FriendUser) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	if p.FriendProfile != nil {
		m["friendProfile"] = p.FriendProfile
	}
	return m
}

func (p FriendUser) Pointer() *FriendUser {
	return &p
}

func CastFriendUsers(data []interface{}) []FriendUser {
	v := make([]FriendUser, 0)
	for _, d := range data {
		v = append(v, NewFriendUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendUsersFromDict(data []FriendUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FriendRequest struct {
	UserId        *string `json:"userId"`
	TargetUserId  *string `json:"targetUserId"`
	PublicProfile *string `json:"publicProfile"`
}

func (p *FriendRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FriendRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = FriendRequest{}
	} else {
		*p = FriendRequest{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
	}
	return nil
}

func NewFriendRequestFromJson(data string) FriendRequest {
	req := FriendRequest{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFriendRequestFromDict(data map[string]interface{}) FriendRequest {
	return FriendRequest{
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
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
	}
}

func (p FriendRequest) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserId != nil {
		m["targetUserId"] = p.TargetUserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	return m
}

func (p FriendRequest) Pointer() *FriendRequest {
	return &p
}

func CastFriendRequests(data []interface{}) []FriendRequest {
	v := make([]FriendRequest, 0)
	for _, d := range data {
		v = append(v, NewFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFriendRequestsFromDict(data []FriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SendFriendRequest struct {
	UserId        *string `json:"userId"`
	TargetUserId  *string `json:"targetUserId"`
	PublicProfile *string `json:"publicProfile"`
}

func (p *SendFriendRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SendFriendRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SendFriendRequest{}
	} else {
		*p = SendFriendRequest{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
	}
	return nil
}

func NewSendFriendRequestFromJson(data string) SendFriendRequest {
	req := SendFriendRequest{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSendFriendRequestFromDict(data map[string]interface{}) SendFriendRequest {
	return SendFriendRequest{
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
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
	}
}

func (p SendFriendRequest) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserId != nil {
		m["targetUserId"] = p.TargetUserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	return m
}

func (p SendFriendRequest) Pointer() *SendFriendRequest {
	return &p
}

func CastSendFriendRequests(data []interface{}) []SendFriendRequest {
	v := make([]SendFriendRequest, 0)
	for _, d := range data {
		v = append(v, NewSendFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSendFriendRequestsFromDict(data []SendFriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReceiveFriendRequest struct {
	UserId        *string `json:"userId"`
	TargetUserId  *string `json:"targetUserId"`
	PublicProfile *string `json:"publicProfile"`
}

func (p *ReceiveFriendRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReceiveFriendRequest{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReceiveFriendRequest{}
	} else {
		*p = ReceiveFriendRequest{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
	}
	return nil
}

func NewReceiveFriendRequestFromJson(data string) ReceiveFriendRequest {
	req := ReceiveFriendRequest{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewReceiveFriendRequestFromDict(data map[string]interface{}) ReceiveFriendRequest {
	return ReceiveFriendRequest{
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
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
	}
}

func (p ReceiveFriendRequest) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserId != nil {
		m["targetUserId"] = p.TargetUserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	return m
}

func (p ReceiveFriendRequest) Pointer() *ReceiveFriendRequest {
	return &p
}

func CastReceiveFriendRequests(data []interface{}) []ReceiveFriendRequest {
	v := make([]ReceiveFriendRequest, 0)
	for _, d := range data {
		v = append(v, NewReceiveFriendRequestFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiveFriendRequestsFromDict(data []ReceiveFriendRequest) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PublicProfile struct {
	UserId        *string `json:"userId"`
	PublicProfile *string `json:"publicProfile"`
}

func (p *PublicProfile) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PublicProfile{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = PublicProfile{}
	} else {
		*p = PublicProfile{}
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
		if v, ok := d["publicProfile"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PublicProfile = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PublicProfile = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PublicProfile = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PublicProfile = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PublicProfile)
				}
			}
		}
	}
	return nil
}

func NewPublicProfileFromJson(data string) PublicProfile {
	req := PublicProfile{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPublicProfileFromDict(data map[string]interface{}) PublicProfile {
	return PublicProfile{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		PublicProfile: func() *string {
			v, ok := data["publicProfile"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["publicProfile"])
		}(),
	}
}

func (p PublicProfile) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.PublicProfile != nil {
		m["publicProfile"] = p.PublicProfile
	}
	return m
}

func (p PublicProfile) Pointer() *PublicProfile {
	return &p
}

func CastPublicProfiles(data []interface{}) []PublicProfile {
	v := make([]PublicProfile, 0)
	for _, d := range data {
		v = append(v, NewPublicProfileFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPublicProfilesFromDict(data []PublicProfile) []interface{} {
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
