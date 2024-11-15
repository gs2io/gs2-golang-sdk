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

package matchmaking

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                                   *string              `json:"namespaceId"`
	Name                                          *string              `json:"name"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	EnableDisconnectDetection                     *string              `json:"enableDisconnectDetection"`
	DisconnectDetectionTimeoutSeconds             *int32               `json:"disconnectDetectionTimeoutSeconds"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	EnableCollaborateSeasonRating                 *string              `json:"enableCollaborateSeasonRating"`
	CollaborateSeasonRatingNamespaceId            *string              `json:"collaborateSeasonRatingNamespaceId"`
	CollaborateSeasonRatingTtl                    *int32               `json:"collaborateSeasonRatingTtl"`
	ChangeRatingScript                            *ScriptSetting       `json:"changeRatingScript"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification                      *NotificationSetting `json:"changeRatingNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
	CreatedAt                                     *int64               `json:"createdAt"`
	UpdatedAt                                     *int64               `json:"updatedAt"`
	Revision                                      *int64               `json:"revision"`
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
		if v, ok := d["enableRating"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableRating)
		}
		if v, ok := d["enableDisconnectDetection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableDisconnectDetection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableDisconnectDetection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableDisconnectDetection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableDisconnectDetection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableDisconnectDetection)
				}
			}
		}
		if v, ok := d["disconnectDetectionTimeoutSeconds"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisconnectDetectionTimeoutSeconds)
		}
		if v, ok := d["createGatheringTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerType)
				}
			}
		}
		if v, ok := d["createGatheringTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["createGatheringTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CreateGatheringTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CreateGatheringTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CreateGatheringTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CreateGatheringTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CreateGatheringTriggerScriptId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerType)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerRealtimeNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerRealtimeNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerRealtimeNamespaceId)
				}
			}
		}
		if v, ok := d["completeMatchmakingTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteMatchmakingTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteMatchmakingTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteMatchmakingTriggerScriptId)
				}
			}
		}
		if v, ok := d["enableCollaborateSeasonRating"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableCollaborateSeasonRating = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableCollaborateSeasonRating = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableCollaborateSeasonRating = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableCollaborateSeasonRating = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableCollaborateSeasonRating)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CollaborateSeasonRatingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CollaborateSeasonRatingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingNamespaceId)
				}
			}
		}
		if v, ok := d["collaborateSeasonRatingTtl"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CollaborateSeasonRatingTtl)
		}
		if v, ok := d["changeRatingScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingScript)
		}
		if v, ok := d["joinNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.JoinNotification)
		}
		if v, ok := d["leaveNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LeaveNotification)
		}
		if v, ok := d["completeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteNotification)
		}
		if v, ok := d["changeRatingNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeRatingNotification)
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
		EnableRating: func() *bool {
			v, ok := data["enableRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableRating"])
		}(),
		EnableDisconnectDetection: func() *string {
			v, ok := data["enableDisconnectDetection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableDisconnectDetection"])
		}(),
		DisconnectDetectionTimeoutSeconds: func() *int32 {
			v, ok := data["disconnectDetectionTimeoutSeconds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["disconnectDetectionTimeoutSeconds"])
		}(),
		CreateGatheringTriggerType: func() *string {
			v, ok := data["createGatheringTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerType"])
		}(),
		CreateGatheringTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["createGatheringTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerRealtimeNamespaceId"])
		}(),
		CreateGatheringTriggerScriptId: func() *string {
			v, ok := data["createGatheringTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["createGatheringTriggerScriptId"])
		}(),
		CompleteMatchmakingTriggerType: func() *string {
			v, ok := data["completeMatchmakingTriggerType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerType"])
		}(),
		CompleteMatchmakingTriggerRealtimeNamespaceId: func() *string {
			v, ok := data["completeMatchmakingTriggerRealtimeNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"])
		}(),
		CompleteMatchmakingTriggerScriptId: func() *string {
			v, ok := data["completeMatchmakingTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["completeMatchmakingTriggerScriptId"])
		}(),
		EnableCollaborateSeasonRating: func() *string {
			v, ok := data["enableCollaborateSeasonRating"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableCollaborateSeasonRating"])
		}(),
		CollaborateSeasonRatingNamespaceId: func() *string {
			v, ok := data["collaborateSeasonRatingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["collaborateSeasonRatingNamespaceId"])
		}(),
		CollaborateSeasonRatingTtl: func() *int32 {
			v, ok := data["collaborateSeasonRatingTtl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["collaborateSeasonRatingTtl"])
		}(),
		ChangeRatingScript: func() *ScriptSetting {
			v, ok := data["changeRatingScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer()
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
		CompleteNotification: func() *NotificationSetting {
			v, ok := data["completeNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer()
		}(),
		ChangeRatingNotification: func() *NotificationSetting {
			v, ok := data["changeRatingNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer()
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
	if p.EnableRating != nil {
		m["enableRating"] = p.EnableRating
	}
	if p.EnableDisconnectDetection != nil {
		m["enableDisconnectDetection"] = p.EnableDisconnectDetection
	}
	if p.DisconnectDetectionTimeoutSeconds != nil {
		m["disconnectDetectionTimeoutSeconds"] = p.DisconnectDetectionTimeoutSeconds
	}
	if p.CreateGatheringTriggerType != nil {
		m["createGatheringTriggerType"] = p.CreateGatheringTriggerType
	}
	if p.CreateGatheringTriggerRealtimeNamespaceId != nil {
		m["createGatheringTriggerRealtimeNamespaceId"] = p.CreateGatheringTriggerRealtimeNamespaceId
	}
	if p.CreateGatheringTriggerScriptId != nil {
		m["createGatheringTriggerScriptId"] = p.CreateGatheringTriggerScriptId
	}
	if p.CompleteMatchmakingTriggerType != nil {
		m["completeMatchmakingTriggerType"] = p.CompleteMatchmakingTriggerType
	}
	if p.CompleteMatchmakingTriggerRealtimeNamespaceId != nil {
		m["completeMatchmakingTriggerRealtimeNamespaceId"] = p.CompleteMatchmakingTriggerRealtimeNamespaceId
	}
	if p.CompleteMatchmakingTriggerScriptId != nil {
		m["completeMatchmakingTriggerScriptId"] = p.CompleteMatchmakingTriggerScriptId
	}
	if p.EnableCollaborateSeasonRating != nil {
		m["enableCollaborateSeasonRating"] = p.EnableCollaborateSeasonRating
	}
	if p.CollaborateSeasonRatingNamespaceId != nil {
		m["collaborateSeasonRatingNamespaceId"] = p.CollaborateSeasonRatingNamespaceId
	}
	if p.CollaborateSeasonRatingTtl != nil {
		m["collaborateSeasonRatingTtl"] = p.CollaborateSeasonRatingTtl
	}
	if p.ChangeRatingScript != nil {
		m["changeRatingScript"] = func() map[string]interface{} {
			if p.ChangeRatingScript == nil {
				return nil
			}
			return p.ChangeRatingScript.ToDict()
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
	if p.CompleteNotification != nil {
		m["completeNotification"] = func() map[string]interface{} {
			if p.CompleteNotification == nil {
				return nil
			}
			return p.CompleteNotification.ToDict()
		}()
	}
	if p.ChangeRatingNotification != nil {
		m["changeRatingNotification"] = func() map[string]interface{} {
			if p.ChangeRatingNotification == nil {
				return nil
			}
			return p.ChangeRatingNotification.ToDict()
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

type Gathering struct {
	GatheringId     *string          `json:"gatheringId"`
	Name            *string          `json:"name"`
	AttributeRanges []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds    []*string        `json:"allowUserIds"`
	Metadata        *string          `json:"metadata"`
	ExpiresAt       *int64           `json:"expiresAt"`
	CreatedAt       *int64           `json:"createdAt"`
	UpdatedAt       *int64           `json:"updatedAt"`
	Revision        *int64           `json:"revision"`
}

func (p *Gathering) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Gathering{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Gathering{}
	} else {
		*p = Gathering{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["gatheringId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringId)
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
		if v, ok := d["attributeRanges"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AttributeRanges)
		}
		if v, ok := d["capacityOfRoles"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CapacityOfRoles)
		}
		if v, ok := d["allowUserIds"]; ok && v != nil {
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
				p.AllowUserIds = l
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
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
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

func NewGatheringFromJson(data string) Gathering {
	req := Gathering{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGatheringFromDict(data map[string]interface{}) Gathering {
	return Gathering{
		GatheringId: func() *string {
			v, ok := data["gatheringId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		AttributeRanges: func() []AttributeRange {
			if data["attributeRanges"] == nil {
				return nil
			}
			return CastAttributeRanges(core.CastArray(data["attributeRanges"]))
		}(),
		CapacityOfRoles: func() []CapacityOfRole {
			if data["capacityOfRoles"] == nil {
				return nil
			}
			return CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"]))
		}(),
		AllowUserIds: func() []*string {
			v, ok := data["allowUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		ExpiresAt: func() *int64 {
			v, ok := data["expiresAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expiresAt"])
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

func (p Gathering) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GatheringId != nil {
		m["gatheringId"] = p.GatheringId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.AttributeRanges != nil {
		m["attributeRanges"] = CastAttributeRangesFromDict(
			p.AttributeRanges,
		)
	}
	if p.CapacityOfRoles != nil {
		m["capacityOfRoles"] = CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		)
	}
	if p.AllowUserIds != nil {
		m["allowUserIds"] = core.CastStringsFromDict(
			p.AllowUserIds,
		)
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
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

func (p Gathering) Pointer() *Gathering {
	return &p
}

func CastGatherings(data []interface{}) []Gathering {
	v := make([]Gathering, 0)
	for _, d := range data {
		v = append(v, NewGatheringFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGatheringsFromDict(data []Gathering) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RatingModelMaster struct {
	RatingModelId *string `json:"ratingModelId"`
	Name          *string `json:"name"`
	Metadata      *string `json:"metadata"`
	Description   *string `json:"description"`
	InitialValue  *int32  `json:"initialValue"`
	Volatility    *int32  `json:"volatility"`
	CreatedAt     *int64  `json:"createdAt"`
	UpdatedAt     *int64  `json:"updatedAt"`
	Revision      *int64  `json:"revision"`
}

func (p *RatingModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RatingModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RatingModelMaster{}
	} else {
		*p = RatingModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["ratingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingModelId)
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
		if v, ok := d["initialValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialValue)
		}
		if v, ok := d["volatility"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Volatility)
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

func NewRatingModelMasterFromJson(data string) RatingModelMaster {
	req := RatingModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRatingModelMasterFromDict(data map[string]interface{}) RatingModelMaster {
	return RatingModelMaster{
		RatingModelId: func() *string {
			v, ok := data["ratingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingModelId"])
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
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		InitialValue: func() *int32 {
			v, ok := data["initialValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialValue"])
		}(),
		Volatility: func() *int32 {
			v, ok := data["volatility"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["volatility"])
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

func (p RatingModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RatingModelId != nil {
		m["ratingModelId"] = p.RatingModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.InitialValue != nil {
		m["initialValue"] = p.InitialValue
	}
	if p.Volatility != nil {
		m["volatility"] = p.Volatility
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

func (p RatingModelMaster) Pointer() *RatingModelMaster {
	return &p
}

func CastRatingModelMasters(data []interface{}) []RatingModelMaster {
	v := make([]RatingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRatingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelMastersFromDict(data []RatingModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RatingModel struct {
	RatingModelId *string `json:"ratingModelId"`
	Name          *string `json:"name"`
	Metadata      *string `json:"metadata"`
	InitialValue  *int32  `json:"initialValue"`
	Volatility    *int32  `json:"volatility"`
}

func (p *RatingModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RatingModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RatingModel{}
	} else {
		*p = RatingModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["ratingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingModelId)
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
		if v, ok := d["initialValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialValue)
		}
		if v, ok := d["volatility"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Volatility)
		}
	}
	return nil
}

func NewRatingModelFromJson(data string) RatingModel {
	req := RatingModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRatingModelFromDict(data map[string]interface{}) RatingModel {
	return RatingModel{
		RatingModelId: func() *string {
			v, ok := data["ratingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingModelId"])
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
		InitialValue: func() *int32 {
			v, ok := data["initialValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialValue"])
		}(),
		Volatility: func() *int32 {
			v, ok := data["volatility"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["volatility"])
		}(),
	}
}

func (p RatingModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RatingModelId != nil {
		m["ratingModelId"] = p.RatingModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.InitialValue != nil {
		m["initialValue"] = p.InitialValue
	}
	if p.Volatility != nil {
		m["volatility"] = p.Volatility
	}
	return m
}

func (p RatingModel) Pointer() *RatingModel {
	return &p
}

func CastRatingModels(data []interface{}) []RatingModel {
	v := make([]RatingModel, 0)
	for _, d := range data {
		v = append(v, NewRatingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingModelsFromDict(data []RatingModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentModelMaster{}
	} else {
		*p = CurrentModelMaster{}
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

func NewCurrentModelMasterFromJson(data string) CurrentModelMaster {
	req := CurrentModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentModelMasterFromDict(data map[string]interface{}) CurrentModelMaster {
	return CurrentModelMaster{
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

func (p CurrentModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentModelMaster) Pointer() *CurrentModelMaster {
	return &p
}

func CastCurrentModelMasters(data []interface{}) []CurrentModelMaster {
	v := make([]CurrentModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentModelMastersFromDict(data []CurrentModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SeasonModel struct {
	SeasonModelId          *string `json:"seasonModelId"`
	Name                   *string `json:"name"`
	Metadata               *string `json:"metadata"`
	MaximumParticipants    *int32  `json:"maximumParticipants"`
	ExperienceModelId      *string `json:"experienceModelId"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
}

func (p *SeasonModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SeasonModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SeasonModel{}
	} else {
		*p = SeasonModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["seasonModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonModelId)
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
		if v, ok := d["maximumParticipants"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParticipants)
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
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewSeasonModelFromJson(data string) SeasonModel {
	req := SeasonModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSeasonModelFromDict(data map[string]interface{}) SeasonModel {
	return SeasonModel{
		SeasonModelId: func() *string {
			v, ok := data["seasonModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonModelId"])
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
		MaximumParticipants: func() *int32 {
			v, ok := data["maximumParticipants"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParticipants"])
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		ChallengePeriodEventId: func() *string {
			v, ok := data["challengePeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["challengePeriodEventId"])
		}(),
	}
}

func (p SeasonModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SeasonModelId != nil {
		m["seasonModelId"] = p.SeasonModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.MaximumParticipants != nil {
		m["maximumParticipants"] = p.MaximumParticipants
	}
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.ChallengePeriodEventId != nil {
		m["challengePeriodEventId"] = p.ChallengePeriodEventId
	}
	return m
}

func (p SeasonModel) Pointer() *SeasonModel {
	return &p
}

func CastSeasonModels(data []interface{}) []SeasonModel {
	v := make([]SeasonModel, 0)
	for _, d := range data {
		v = append(v, NewSeasonModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSeasonModelsFromDict(data []SeasonModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SeasonModelMaster struct {
	SeasonModelId          *string `json:"seasonModelId"`
	Name                   *string `json:"name"`
	Metadata               *string `json:"metadata"`
	Description            *string `json:"description"`
	MaximumParticipants    *int32  `json:"maximumParticipants"`
	ExperienceModelId      *string `json:"experienceModelId"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
	Revision               *int64  `json:"revision"`
}

func (p *SeasonModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SeasonModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SeasonModelMaster{}
	} else {
		*p = SeasonModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["seasonModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonModelId)
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
		if v, ok := d["maximumParticipants"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParticipants)
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
		if v, ok := d["challengePeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChallengePeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChallengePeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChallengePeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChallengePeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChallengePeriodEventId)
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

func NewSeasonModelMasterFromJson(data string) SeasonModelMaster {
	req := SeasonModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSeasonModelMasterFromDict(data map[string]interface{}) SeasonModelMaster {
	return SeasonModelMaster{
		SeasonModelId: func() *string {
			v, ok := data["seasonModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonModelId"])
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
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		MaximumParticipants: func() *int32 {
			v, ok := data["maximumParticipants"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParticipants"])
		}(),
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		ChallengePeriodEventId: func() *string {
			v, ok := data["challengePeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["challengePeriodEventId"])
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

func (p SeasonModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SeasonModelId != nil {
		m["seasonModelId"] = p.SeasonModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.MaximumParticipants != nil {
		m["maximumParticipants"] = p.MaximumParticipants
	}
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.ChallengePeriodEventId != nil {
		m["challengePeriodEventId"] = p.ChallengePeriodEventId
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

func (p SeasonModelMaster) Pointer() *SeasonModelMaster {
	return &p
}

func CastSeasonModelMasters(data []interface{}) []SeasonModelMaster {
	v := make([]SeasonModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSeasonModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSeasonModelMastersFromDict(data []SeasonModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SeasonGathering struct {
	SeasonGatheringId *string   `json:"seasonGatheringId"`
	SeasonName        *string   `json:"seasonName"`
	Season            *int64    `json:"season"`
	Tier              *int64    `json:"tier"`
	Name              *string   `json:"name"`
	Participants      []*string `json:"participants"`
	CreatedAt         *int64    `json:"createdAt"`
	Revision          *int64    `json:"revision"`
}

func (p *SeasonGathering) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SeasonGathering{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SeasonGathering{}
	} else {
		*p = SeasonGathering{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["seasonGatheringId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringId)
				}
			}
		}
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
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
		if v, ok := d["participants"]; ok && v != nil {
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
				p.Participants = l
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewSeasonGatheringFromJson(data string) SeasonGathering {
	req := SeasonGathering{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSeasonGatheringFromDict(data map[string]interface{}) SeasonGathering {
	return SeasonGathering{
		SeasonGatheringId: func() *string {
			v, ok := data["seasonGatheringId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringId"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Participants: func() []*string {
			v, ok := data["participants"]
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
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p SeasonGathering) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SeasonGatheringId != nil {
		m["seasonGatheringId"] = p.SeasonGatheringId
	}
	if p.SeasonName != nil {
		m["seasonName"] = p.SeasonName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.Tier != nil {
		m["tier"] = p.Tier
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Participants != nil {
		m["participants"] = core.CastStringsFromDict(
			p.Participants,
		)
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p SeasonGathering) Pointer() *SeasonGathering {
	return &p
}

func CastSeasonGatherings(data []interface{}) []SeasonGathering {
	v := make([]SeasonGathering, 0)
	for _, d := range data {
		v = append(v, NewSeasonGatheringFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSeasonGatheringsFromDict(data []SeasonGathering) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type JoinedSeasonGathering struct {
	JoinedSeasonGatheringId *string `json:"joinedSeasonGatheringId"`
	UserId                  *string `json:"userId"`
	SeasonName              *string `json:"seasonName"`
	Season                  *int64  `json:"season"`
	Tier                    *int64  `json:"tier"`
	SeasonGatheringName     *string `json:"seasonGatheringName"`
	CreatedAt               *int64  `json:"createdAt"`
}

func (p *JoinedSeasonGathering) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = JoinedSeasonGathering{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = JoinedSeasonGathering{}
	} else {
		*p = JoinedSeasonGathering{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["joinedSeasonGatheringId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JoinedSeasonGatheringId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JoinedSeasonGatheringId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JoinedSeasonGatheringId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JoinedSeasonGatheringId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JoinedSeasonGatheringId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JoinedSeasonGatheringId)
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
		if v, ok := d["seasonName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["tier"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tier)
		}
		if v, ok := d["seasonGatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SeasonGatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SeasonGatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SeasonGatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SeasonGatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SeasonGatheringName)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewJoinedSeasonGatheringFromJson(data string) JoinedSeasonGathering {
	req := JoinedSeasonGathering{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJoinedSeasonGatheringFromDict(data map[string]interface{}) JoinedSeasonGathering {
	return JoinedSeasonGathering{
		JoinedSeasonGatheringId: func() *string {
			v, ok := data["joinedSeasonGatheringId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["joinedSeasonGatheringId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		SeasonName: func() *string {
			v, ok := data["seasonName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Tier: func() *int64 {
			v, ok := data["tier"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tier"])
		}(),
		SeasonGatheringName: func() *string {
			v, ok := data["seasonGatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["seasonGatheringName"])
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

func (p JoinedSeasonGathering) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.JoinedSeasonGatheringId != nil {
		m["joinedSeasonGatheringId"] = p.JoinedSeasonGatheringId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.SeasonName != nil {
		m["seasonName"] = p.SeasonName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.Tier != nil {
		m["tier"] = p.Tier
	}
	if p.SeasonGatheringName != nil {
		m["seasonGatheringName"] = p.SeasonGatheringName
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p JoinedSeasonGathering) Pointer() *JoinedSeasonGathering {
	return &p
}

func CastJoinedSeasonGatherings(data []interface{}) []JoinedSeasonGathering {
	v := make([]JoinedSeasonGathering, 0)
	for _, d := range data {
		v = append(v, NewJoinedSeasonGatheringFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJoinedSeasonGatheringsFromDict(data []JoinedSeasonGathering) []interface{} {
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

type AttributeRange struct {
	Name *string `json:"name"`
	Min  *int32  `json:"min"`
	Max  *int32  `json:"max"`
}

func (p *AttributeRange) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AttributeRange{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AttributeRange{}
	} else {
		*p = AttributeRange{}
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
		if v, ok := d["min"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Min)
		}
		if v, ok := d["max"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Max)
		}
	}
	return nil
}

func NewAttributeRangeFromJson(data string) AttributeRange {
	req := AttributeRange{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAttributeRangeFromDict(data map[string]interface{}) AttributeRange {
	return AttributeRange{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Min: func() *int32 {
			v, ok := data["min"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["min"])
		}(),
		Max: func() *int32 {
			v, ok := data["max"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["max"])
		}(),
	}
}

func (p AttributeRange) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Min != nil {
		m["min"] = p.Min
	}
	if p.Max != nil {
		m["max"] = p.Max
	}
	return m
}

func (p AttributeRange) Pointer() *AttributeRange {
	return &p
}

func CastAttributeRanges(data []interface{}) []AttributeRange {
	v := make([]AttributeRange, 0)
	for _, d := range data {
		v = append(v, NewAttributeRangeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributeRangesFromDict(data []AttributeRange) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CapacityOfRole struct {
	RoleName     *string   `json:"roleName"`
	RoleAliases  []*string `json:"roleAliases"`
	Capacity     *int32    `json:"capacity"`
	Participants []Player  `json:"participants"`
}

func (p *CapacityOfRole) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CapacityOfRole{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CapacityOfRole{}
	} else {
		*p = CapacityOfRole{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["roleAliases"]; ok && v != nil {
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
				p.RoleAliases = l
			}
		}
		if v, ok := d["capacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Capacity)
		}
		if v, ok := d["participants"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Participants)
		}
	}
	return nil
}

func NewCapacityOfRoleFromJson(data string) CapacityOfRole {
	req := CapacityOfRole{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCapacityOfRoleFromDict(data map[string]interface{}) CapacityOfRole {
	return CapacityOfRole{
		RoleName: func() *string {
			v, ok := data["roleName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["roleName"])
		}(),
		RoleAliases: func() []*string {
			v, ok := data["roleAliases"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		Capacity: func() *int32 {
			v, ok := data["capacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["capacity"])
		}(),
		Participants: func() []Player {
			if data["participants"] == nil {
				return nil
			}
			return CastPlayers(core.CastArray(data["participants"]))
		}(),
	}
}

func (p CapacityOfRole) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RoleName != nil {
		m["roleName"] = p.RoleName
	}
	if p.RoleAliases != nil {
		m["roleAliases"] = core.CastStringsFromDict(
			p.RoleAliases,
		)
	}
	if p.Capacity != nil {
		m["capacity"] = p.Capacity
	}
	if p.Participants != nil {
		m["participants"] = CastPlayersFromDict(
			p.Participants,
		)
	}
	return m
}

func (p CapacityOfRole) Pointer() *CapacityOfRole {
	return &p
}

func CastCapacityOfRoles(data []interface{}) []CapacityOfRole {
	v := make([]CapacityOfRole, 0)
	for _, d := range data {
		v = append(v, NewCapacityOfRoleFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCapacityOfRolesFromDict(data []CapacityOfRole) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Attribute struct {
	Name  *string `json:"name"`
	Value *int32  `json:"value"`
}

func (p *Attribute) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Attribute{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Attribute{}
	} else {
		*p = Attribute{}
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
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewAttributeFromJson(data string) Attribute {
	req := Attribute{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAttributeFromDict(data map[string]interface{}) Attribute {
	return Attribute{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
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

func (p Attribute) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p Attribute) Pointer() *Attribute {
	return &p
}

func CastAttributes(data []interface{}) []Attribute {
	v := make([]Attribute, 0)
	for _, d := range data {
		v = append(v, NewAttributeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAttributesFromDict(data []Attribute) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Player struct {
	UserId      *string     `json:"userId"`
	Attributes  []Attribute `json:"attributes"`
	RoleName    *string     `json:"roleName"`
	DenyUserIds []*string   `json:"denyUserIds"`
	CreatedAt   *int64      `json:"createdAt"`
}

func (p *Player) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Player{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Player{}
	} else {
		*p = Player{}
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
		if v, ok := d["attributes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Attributes)
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
		if v, ok := d["denyUserIds"]; ok && v != nil {
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
				p.DenyUserIds = l
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewPlayerFromJson(data string) Player {
	req := Player{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPlayerFromDict(data map[string]interface{}) Player {
	return Player{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Attributes: func() []Attribute {
			if data["attributes"] == nil {
				return nil
			}
			return CastAttributes(core.CastArray(data["attributes"]))
		}(),
		RoleName: func() *string {
			v, ok := data["roleName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["roleName"])
		}(),
		DenyUserIds: func() []*string {
			v, ok := data["denyUserIds"]
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
	}
}

func (p Player) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Attributes != nil {
		m["attributes"] = CastAttributesFromDict(
			p.Attributes,
		)
	}
	if p.RoleName != nil {
		m["roleName"] = p.RoleName
	}
	if p.DenyUserIds != nil {
		m["denyUserIds"] = core.CastStringsFromDict(
			p.DenyUserIds,
		)
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
}

func (p Player) Pointer() *Player {
	return &p
}

func CastPlayers(data []interface{}) []Player {
	v := make([]Player, 0)
	for _, d := range data {
		v = append(v, NewPlayerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPlayersFromDict(data []Player) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Rating struct {
	RatingId  *string  `json:"ratingId"`
	Name      *string  `json:"name"`
	UserId    *string  `json:"userId"`
	RateValue *float32 `json:"rateValue"`
	CreatedAt *int64   `json:"createdAt"`
	UpdatedAt *int64   `json:"updatedAt"`
	Revision  *int64   `json:"revision"`
}

func (p *Rating) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Rating{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Rating{}
	} else {
		*p = Rating{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["ratingId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingId)
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
		if v, ok := d["rateValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RateValue)
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

func NewRatingFromJson(data string) Rating {
	req := Rating{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRatingFromDict(data map[string]interface{}) Rating {
	return Rating{
		RatingId: func() *string {
			v, ok := data["ratingId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		RateValue: func() *float32 {
			v, ok := data["rateValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["rateValue"])
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

func (p Rating) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RatingId != nil {
		m["ratingId"] = p.RatingId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.RateValue != nil {
		m["rateValue"] = p.RateValue
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

func (p Rating) Pointer() *Rating {
	return &p
}

func CastRatings(data []interface{}) []Rating {
	v := make([]Rating, 0)
	for _, d := range data {
		v = append(v, NewRatingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRatingsFromDict(data []Rating) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GameResult struct {
	Rank   *int32  `json:"rank"`
	UserId *string `json:"userId"`
}

func (p *GameResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GameResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = GameResult{}
	} else {
		*p = GameResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rank)
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

func NewGameResultFromJson(data string) GameResult {
	req := GameResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGameResultFromDict(data map[string]interface{}) GameResult {
	return GameResult{
		Rank: func() *int32 {
			v, ok := data["rank"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rank"])
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

func (p GameResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Rank != nil {
		m["rank"] = p.Rank
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	return m
}

func (p GameResult) Pointer() *GameResult {
	return &p
}

func CastGameResults(data []interface{}) []GameResult {
	v := make([]GameResult, 0)
	for _, d := range data {
		v = append(v, NewGameResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGameResultsFromDict(data []GameResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Ballot struct {
	UserId         *string `json:"userId"`
	RatingName     *string `json:"ratingName"`
	GatheringName  *string `json:"gatheringName"`
	NumberOfPlayer *int32  `json:"numberOfPlayer"`
}

func (p *Ballot) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Ballot{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Ballot{}
	} else {
		*p = Ballot{}
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
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["numberOfPlayer"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NumberOfPlayer)
		}
	}
	return nil
}

func NewBallotFromJson(data string) Ballot {
	req := Ballot{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBallotFromDict(data map[string]interface{}) Ballot {
	return Ballot{
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		NumberOfPlayer: func() *int32 {
			v, ok := data["numberOfPlayer"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["numberOfPlayer"])
		}(),
	}
}

func (p Ballot) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.RatingName != nil {
		m["ratingName"] = p.RatingName
	}
	if p.GatheringName != nil {
		m["gatheringName"] = p.GatheringName
	}
	if p.NumberOfPlayer != nil {
		m["numberOfPlayer"] = p.NumberOfPlayer
	}
	return m
}

func (p Ballot) Pointer() *Ballot {
	return &p
}

func CastBallots(data []interface{}) []Ballot {
	v := make([]Ballot, 0)
	for _, d := range data {
		v = append(v, NewBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBallotsFromDict(data []Ballot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SignedBallot struct {
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

func (p *SignedBallot) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SignedBallot{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SignedBallot{}
	} else {
		*p = SignedBallot{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["body"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Body = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Body = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Body = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Body = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Body = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Body)
				}
			}
		}
		if v, ok := d["signature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Signature = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Signature = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Signature = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Signature = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Signature = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Signature)
				}
			}
		}
	}
	return nil
}

func NewSignedBallotFromJson(data string) SignedBallot {
	req := SignedBallot{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSignedBallotFromDict(data map[string]interface{}) SignedBallot {
	return SignedBallot{
		Body: func() *string {
			v, ok := data["body"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["body"])
		}(),
		Signature: func() *string {
			v, ok := data["signature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signature"])
		}(),
	}
}

func (p SignedBallot) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Body != nil {
		m["body"] = p.Body
	}
	if p.Signature != nil {
		m["signature"] = p.Signature
	}
	return m
}

func (p SignedBallot) Pointer() *SignedBallot {
	return &p
}

func CastSignedBallots(data []interface{}) []SignedBallot {
	v := make([]SignedBallot, 0)
	for _, d := range data {
		v = append(v, NewSignedBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignedBallotsFromDict(data []SignedBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type WrittenBallot struct {
	Ballot      *Ballot      `json:"ballot"`
	GameResults []GameResult `json:"gameResults"`
}

func (p *WrittenBallot) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = WrittenBallot{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = WrittenBallot{}
	} else {
		*p = WrittenBallot{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["ballot"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Ballot)
		}
		if v, ok := d["gameResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GameResults)
		}
	}
	return nil
}

func NewWrittenBallotFromJson(data string) WrittenBallot {
	req := WrittenBallot{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewWrittenBallotFromDict(data map[string]interface{}) WrittenBallot {
	return WrittenBallot{
		Ballot: func() *Ballot {
			v, ok := data["ballot"]
			if !ok || v == nil {
				return nil
			}
			return NewBallotFromDict(core.CastMap(data["ballot"])).Pointer()
		}(),
		GameResults: func() []GameResult {
			if data["gameResults"] == nil {
				return nil
			}
			return CastGameResults(core.CastArray(data["gameResults"]))
		}(),
	}
}

func (p WrittenBallot) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Ballot != nil {
		m["ballot"] = func() map[string]interface{} {
			if p.Ballot == nil {
				return nil
			}
			return p.Ballot.ToDict()
		}()
	}
	if p.GameResults != nil {
		m["gameResults"] = CastGameResultsFromDict(
			p.GameResults,
		)
	}
	return m
}

func (p WrittenBallot) Pointer() *WrittenBallot {
	return &p
}

func CastWrittenBallots(data []interface{}) []WrittenBallot {
	v := make([]WrittenBallot, 0)
	for _, d := range data {
		v = append(v, NewWrittenBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWrittenBallotsFromDict(data []WrittenBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Vote struct {
	VoteId         *string         `json:"voteId"`
	RatingName     *string         `json:"ratingName"`
	GatheringName  *string         `json:"gatheringName"`
	WrittenBallots []WrittenBallot `json:"writtenBallots"`
	CreatedAt      *int64          `json:"createdAt"`
	UpdatedAt      *int64          `json:"updatedAt"`
}

func (p *Vote) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Vote{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Vote{}
	} else {
		*p = Vote{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["voteId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VoteId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VoteId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VoteId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VoteId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VoteId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VoteId)
				}
			}
		}
		if v, ok := d["ratingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RatingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RatingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RatingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RatingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RatingName)
				}
			}
		}
		if v, ok := d["gatheringName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatheringName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatheringName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatheringName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatheringName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatheringName)
				}
			}
		}
		if v, ok := d["writtenBallots"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WrittenBallots)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewVoteFromJson(data string) Vote {
	req := Vote{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVoteFromDict(data map[string]interface{}) Vote {
	return Vote{
		VoteId: func() *string {
			v, ok := data["voteId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["voteId"])
		}(),
		RatingName: func() *string {
			v, ok := data["ratingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["ratingName"])
		}(),
		GatheringName: func() *string {
			v, ok := data["gatheringName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatheringName"])
		}(),
		WrittenBallots: func() []WrittenBallot {
			if data["writtenBallots"] == nil {
				return nil
			}
			return CastWrittenBallots(core.CastArray(data["writtenBallots"]))
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
	}
}

func (p Vote) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.VoteId != nil {
		m["voteId"] = p.VoteId
	}
	if p.RatingName != nil {
		m["ratingName"] = p.RatingName
	}
	if p.GatheringName != nil {
		m["gatheringName"] = p.GatheringName
	}
	if p.WrittenBallots != nil {
		m["writtenBallots"] = CastWrittenBallotsFromDict(
			p.WrittenBallots,
		)
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Vote) Pointer() *Vote {
	return &p
}

func CastVotes(data []interface{}) []Vote {
	v := make([]Vote, 0)
	for _, d := range data {
		v = append(v, NewVoteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVotesFromDict(data []Vote) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TimeSpan struct {
	Days    *int32 `json:"days"`
	Hours   *int32 `json:"hours"`
	Minutes *int32 `json:"minutes"`
}

func (p *TimeSpan) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TimeSpan{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TimeSpan{}
	} else {
		*p = TimeSpan{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["days"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Days)
		}
		if v, ok := d["hours"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Hours)
		}
		if v, ok := d["minutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Minutes)
		}
	}
	return nil
}

func NewTimeSpanFromJson(data string) TimeSpan {
	req := TimeSpan{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTimeSpanFromDict(data map[string]interface{}) TimeSpan {
	return TimeSpan{
		Days: func() *int32 {
			v, ok := data["days"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["days"])
		}(),
		Hours: func() *int32 {
			v, ok := data["hours"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["hours"])
		}(),
		Minutes: func() *int32 {
			v, ok := data["minutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["minutes"])
		}(),
	}
}

func (p TimeSpan) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Days != nil {
		m["days"] = p.Days
	}
	if p.Hours != nil {
		m["hours"] = p.Hours
	}
	if p.Minutes != nil {
		m["minutes"] = p.Minutes
	}
	return m
}

func (p TimeSpan) Pointer() *TimeSpan {
	return &p
}

func CastTimeSpans(data []interface{}) []TimeSpan {
	v := make([]TimeSpan, 0)
	for _, d := range data {
		v = append(v, NewTimeSpanFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeSpansFromDict(data []TimeSpan) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
