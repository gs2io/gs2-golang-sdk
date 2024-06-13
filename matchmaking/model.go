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
		NamespaceId:                       core.CastString(data["namespaceId"]),
		Name:                              core.CastString(data["name"]),
		Description:                       core.CastString(data["description"]),
		EnableRating:                      core.CastBool(data["enableRating"]),
		EnableDisconnectDetection:         core.CastString(data["enableDisconnectDetection"]),
		DisconnectDetectionTimeoutSeconds: core.CastInt32(data["disconnectDetectionTimeoutSeconds"]),
		CreateGatheringTriggerType:        core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		EnableCollaborateSeasonRating:                 core.CastString(data["enableCollaborateSeasonRating"]),
		CollaborateSeasonRatingNamespaceId:            core.CastString(data["collaborateSeasonRatingNamespaceId"]),
		CollaborateSeasonRatingTtl:                    core.CastInt32(data["collaborateSeasonRatingTtl"]),
		ChangeRatingScript:                            NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		ChangeRatingNotification:                      NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                                     core.CastInt64(data["createdAt"]),
		UpdatedAt:                                     core.CastInt64(data["updatedAt"]),
		Revision:                                      core.CastInt64(data["revision"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var enableRating *bool
	if p.EnableRating != nil {
		enableRating = p.EnableRating
	}
	var enableDisconnectDetection *string
	if p.EnableDisconnectDetection != nil {
		enableDisconnectDetection = p.EnableDisconnectDetection
	}
	var disconnectDetectionTimeoutSeconds *int32
	if p.DisconnectDetectionTimeoutSeconds != nil {
		disconnectDetectionTimeoutSeconds = p.DisconnectDetectionTimeoutSeconds
	}
	var createGatheringTriggerType *string
	if p.CreateGatheringTriggerType != nil {
		createGatheringTriggerType = p.CreateGatheringTriggerType
	}
	var createGatheringTriggerRealtimeNamespaceId *string
	if p.CreateGatheringTriggerRealtimeNamespaceId != nil {
		createGatheringTriggerRealtimeNamespaceId = p.CreateGatheringTriggerRealtimeNamespaceId
	}
	var createGatheringTriggerScriptId *string
	if p.CreateGatheringTriggerScriptId != nil {
		createGatheringTriggerScriptId = p.CreateGatheringTriggerScriptId
	}
	var completeMatchmakingTriggerType *string
	if p.CompleteMatchmakingTriggerType != nil {
		completeMatchmakingTriggerType = p.CompleteMatchmakingTriggerType
	}
	var completeMatchmakingTriggerRealtimeNamespaceId *string
	if p.CompleteMatchmakingTriggerRealtimeNamespaceId != nil {
		completeMatchmakingTriggerRealtimeNamespaceId = p.CompleteMatchmakingTriggerRealtimeNamespaceId
	}
	var completeMatchmakingTriggerScriptId *string
	if p.CompleteMatchmakingTriggerScriptId != nil {
		completeMatchmakingTriggerScriptId = p.CompleteMatchmakingTriggerScriptId
	}
	var enableCollaborateSeasonRating *string
	if p.EnableCollaborateSeasonRating != nil {
		enableCollaborateSeasonRating = p.EnableCollaborateSeasonRating
	}
	var collaborateSeasonRatingNamespaceId *string
	if p.CollaborateSeasonRatingNamespaceId != nil {
		collaborateSeasonRatingNamespaceId = p.CollaborateSeasonRatingNamespaceId
	}
	var collaborateSeasonRatingTtl *int32
	if p.CollaborateSeasonRatingTtl != nil {
		collaborateSeasonRatingTtl = p.CollaborateSeasonRatingTtl
	}
	var changeRatingScript map[string]interface{}
	if p.ChangeRatingScript != nil {
		changeRatingScript = p.ChangeRatingScript.ToDict()
	}
	var joinNotification map[string]interface{}
	if p.JoinNotification != nil {
		joinNotification = p.JoinNotification.ToDict()
	}
	var leaveNotification map[string]interface{}
	if p.LeaveNotification != nil {
		leaveNotification = p.LeaveNotification.ToDict()
	}
	var completeNotification map[string]interface{}
	if p.CompleteNotification != nil {
		completeNotification = p.CompleteNotification.ToDict()
	}
	var changeRatingNotification map[string]interface{}
	if p.ChangeRatingNotification != nil {
		changeRatingNotification = p.ChangeRatingNotification.ToDict()
	}
	var logSetting map[string]interface{}
	if p.LogSetting != nil {
		logSetting = p.LogSetting.ToDict()
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":                       namespaceId,
		"name":                              name,
		"description":                       description,
		"enableRating":                      enableRating,
		"enableDisconnectDetection":         enableDisconnectDetection,
		"disconnectDetectionTimeoutSeconds": disconnectDetectionTimeoutSeconds,
		"createGatheringTriggerType":        createGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     createGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                createGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                completeMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": completeMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            completeMatchmakingTriggerScriptId,
		"enableCollaborateSeasonRating":                 enableCollaborateSeasonRating,
		"collaborateSeasonRatingNamespaceId":            collaborateSeasonRatingNamespaceId,
		"collaborateSeasonRatingTtl":                    collaborateSeasonRatingTtl,
		"changeRatingScript":                            changeRatingScript,
		"joinNotification":                              joinNotification,
		"leaveNotification":                             leaveNotification,
		"completeNotification":                          completeNotification,
		"changeRatingNotification":                      changeRatingNotification,
		"logSetting":                                    logSetting,
		"createdAt":                                     createdAt,
		"updatedAt":                                     updatedAt,
		"revision":                                      revision,
	}
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
		GatheringId:     core.CastString(data["gatheringId"]),
		Name:            core.CastString(data["name"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles: CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:    core.CastStrings(core.CastArray(data["allowUserIds"])),
		Metadata:        core.CastString(data["metadata"]),
		ExpiresAt:       core.CastInt64(data["expiresAt"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
		Revision:        core.CastInt64(data["revision"]),
	}
}

func (p Gathering) ToDict() map[string]interface{} {

	var gatheringId *string
	if p.GatheringId != nil {
		gatheringId = p.GatheringId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var attributeRanges []interface{}
	if p.AttributeRanges != nil {
		attributeRanges = CastAttributeRangesFromDict(
			p.AttributeRanges,
		)
	}
	var capacityOfRoles []interface{}
	if p.CapacityOfRoles != nil {
		capacityOfRoles = CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		)
	}
	var allowUserIds []interface{}
	if p.AllowUserIds != nil {
		allowUserIds = core.CastStringsFromDict(
			p.AllowUserIds,
		)
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"gatheringId":     gatheringId,
		"name":            name,
		"attributeRanges": attributeRanges,
		"capacityOfRoles": capacityOfRoles,
		"allowUserIds":    allowUserIds,
		"metadata":        metadata,
		"expiresAt":       expiresAt,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
		"revision":        revision,
	}
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
		RatingModelId: core.CastString(data["ratingModelId"]),
		Name:          core.CastString(data["name"]),
		Metadata:      core.CastString(data["metadata"]),
		Description:   core.CastString(data["description"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p RatingModelMaster) ToDict() map[string]interface{} {

	var ratingModelId *string
	if p.RatingModelId != nil {
		ratingModelId = p.RatingModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var initialValue *int32
	if p.InitialValue != nil {
		initialValue = p.InitialValue
	}
	var volatility *int32
	if p.Volatility != nil {
		volatility = p.Volatility
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"ratingModelId": ratingModelId,
		"name":          name,
		"metadata":      metadata,
		"description":   description,
		"initialValue":  initialValue,
		"volatility":    volatility,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
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
		RatingModelId: core.CastString(data["ratingModelId"]),
		Name:          core.CastString(data["name"]),
		Metadata:      core.CastString(data["metadata"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
	}
}

func (p RatingModel) ToDict() map[string]interface{} {

	var ratingModelId *string
	if p.RatingModelId != nil {
		ratingModelId = p.RatingModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var initialValue *int32
	if p.InitialValue != nil {
		initialValue = p.InitialValue
	}
	var volatility *int32
	if p.Volatility != nil {
		volatility = p.Volatility
	}
	return map[string]interface{}{
		"ratingModelId": ratingModelId,
		"name":          name,
		"metadata":      metadata,
		"initialValue":  initialValue,
		"volatility":    volatility,
	}
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
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentModelMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
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
		SeasonModelId:          core.CastString(data["seasonModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		MaximumParticipants:    core.CastInt32(data["maximumParticipants"]),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p SeasonModel) ToDict() map[string]interface{} {

	var seasonModelId *string
	if p.SeasonModelId != nil {
		seasonModelId = p.SeasonModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var maximumParticipants *int32
	if p.MaximumParticipants != nil {
		maximumParticipants = p.MaximumParticipants
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	return map[string]interface{}{
		"seasonModelId":          seasonModelId,
		"name":                   name,
		"metadata":               metadata,
		"maximumParticipants":    maximumParticipants,
		"experienceModelId":      experienceModelId,
		"challengePeriodEventId": challengePeriodEventId,
	}
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
		SeasonModelId:          core.CastString(data["seasonModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		MaximumParticipants:    core.CastInt32(data["maximumParticipants"]),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p SeasonModelMaster) ToDict() map[string]interface{} {

	var seasonModelId *string
	if p.SeasonModelId != nil {
		seasonModelId = p.SeasonModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var maximumParticipants *int32
	if p.MaximumParticipants != nil {
		maximumParticipants = p.MaximumParticipants
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"seasonModelId":          seasonModelId,
		"name":                   name,
		"metadata":               metadata,
		"description":            description,
		"maximumParticipants":    maximumParticipants,
		"experienceModelId":      experienceModelId,
		"challengePeriodEventId": challengePeriodEventId,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
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
		SeasonGatheringId: core.CastString(data["seasonGatheringId"]),
		SeasonName:        core.CastString(data["seasonName"]),
		Season:            core.CastInt64(data["season"]),
		Tier:              core.CastInt64(data["tier"]),
		Name:              core.CastString(data["name"]),
		Participants:      core.CastStrings(core.CastArray(data["participants"])),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p SeasonGathering) ToDict() map[string]interface{} {

	var seasonGatheringId *string
	if p.SeasonGatheringId != nil {
		seasonGatheringId = p.SeasonGatheringId
	}
	var seasonName *string
	if p.SeasonName != nil {
		seasonName = p.SeasonName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var tier *int64
	if p.Tier != nil {
		tier = p.Tier
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var participants []interface{}
	if p.Participants != nil {
		participants = core.CastStringsFromDict(
			p.Participants,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"seasonGatheringId": seasonGatheringId,
		"seasonName":        seasonName,
		"season":            season,
		"tier":              tier,
		"name":              name,
		"participants":      participants,
		"createdAt":         createdAt,
		"revision":          revision,
	}
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
		JoinedSeasonGatheringId: core.CastString(data["joinedSeasonGatheringId"]),
		UserId:                  core.CastString(data["userId"]),
		SeasonName:              core.CastString(data["seasonName"]),
		Season:                  core.CastInt64(data["season"]),
		Tier:                    core.CastInt64(data["tier"]),
		SeasonGatheringName:     core.CastString(data["seasonGatheringName"]),
		CreatedAt:               core.CastInt64(data["createdAt"]),
	}
}

func (p JoinedSeasonGathering) ToDict() map[string]interface{} {

	var joinedSeasonGatheringId *string
	if p.JoinedSeasonGatheringId != nil {
		joinedSeasonGatheringId = p.JoinedSeasonGatheringId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var seasonName *string
	if p.SeasonName != nil {
		seasonName = p.SeasonName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var tier *int64
	if p.Tier != nil {
		tier = p.Tier
	}
	var seasonGatheringName *string
	if p.SeasonGatheringName != nil {
		seasonGatheringName = p.SeasonGatheringName
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"joinedSeasonGatheringId": joinedSeasonGatheringId,
		"userId":                  userId,
		"seasonName":              seasonName,
		"season":                  season,
		"tier":                    tier,
		"seasonGatheringName":     seasonGatheringName,
		"createdAt":               createdAt,
	}
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
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {

	var triggerScriptId *string
	if p.TriggerScriptId != nil {
		triggerScriptId = p.TriggerScriptId
	}
	var doneTriggerTargetType *string
	if p.DoneTriggerTargetType != nil {
		doneTriggerTargetType = p.DoneTriggerTargetType
	}
	var doneTriggerScriptId *string
	if p.DoneTriggerScriptId != nil {
		doneTriggerScriptId = p.DoneTriggerScriptId
	}
	var doneTriggerQueueNamespaceId *string
	if p.DoneTriggerQueueNamespaceId != nil {
		doneTriggerQueueNamespaceId = p.DoneTriggerQueueNamespaceId
	}
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
		"doneTriggerQueueNamespaceId": doneTriggerQueueNamespaceId,
	}
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
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {

	var gatewayNamespaceId *string
	if p.GatewayNamespaceId != nil {
		gatewayNamespaceId = p.GatewayNamespaceId
	}
	var enableTransferMobileNotification *bool
	if p.EnableTransferMobileNotification != nil {
		enableTransferMobileNotification = p.EnableTransferMobileNotification
	}
	var sound *string
	if p.Sound != nil {
		sound = p.Sound
	}
	return map[string]interface{}{
		"gatewayNamespaceId":               gatewayNamespaceId,
		"enableTransferMobileNotification": enableTransferMobileNotification,
		"sound":                            sound,
	}
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
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {

	var loggingNamespaceId *string
	if p.LoggingNamespaceId != nil {
		loggingNamespaceId = p.LoggingNamespaceId
	}
	return map[string]interface{}{
		"loggingNamespaceId": loggingNamespaceId,
	}
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
		ApiKeyId:       core.CastString(data["apiKeyId"]),
		RepositoryName: core.CastString(data["repositoryName"]),
		SourcePath:     core.CastString(data["sourcePath"]),
		ReferenceType:  core.CastString(data["referenceType"]),
		CommitHash:     core.CastString(data["commitHash"]),
		BranchName:     core.CastString(data["branchName"]),
		TagName:        core.CastString(data["tagName"]),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
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
		Name: core.CastString(data["name"]),
		Min:  core.CastInt32(data["min"]),
		Max:  core.CastInt32(data["max"]),
	}
}

func (p AttributeRange) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var min *int32
	if p.Min != nil {
		min = p.Min
	}
	var max *int32
	if p.Max != nil {
		max = p.Max
	}
	return map[string]interface{}{
		"name": name,
		"min":  min,
		"max":  max,
	}
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
		RoleName:     core.CastString(data["roleName"]),
		RoleAliases:  core.CastStrings(core.CastArray(data["roleAliases"])),
		Capacity:     core.CastInt32(data["capacity"]),
		Participants: CastPlayers(core.CastArray(data["participants"])),
	}
}

func (p CapacityOfRole) ToDict() map[string]interface{} {

	var roleName *string
	if p.RoleName != nil {
		roleName = p.RoleName
	}
	var roleAliases []interface{}
	if p.RoleAliases != nil {
		roleAliases = core.CastStringsFromDict(
			p.RoleAliases,
		)
	}
	var capacity *int32
	if p.Capacity != nil {
		capacity = p.Capacity
	}
	var participants []interface{}
	if p.Participants != nil {
		participants = CastPlayersFromDict(
			p.Participants,
		)
	}
	return map[string]interface{}{
		"roleName":     roleName,
		"roleAliases":  roleAliases,
		"capacity":     capacity,
		"participants": participants,
	}
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
		Name:  core.CastString(data["name"]),
		Value: core.CastInt32(data["value"]),
	}
}

func (p Attribute) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var value *int32
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"name":  name,
		"value": value,
	}
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
		UserId:      core.CastString(data["userId"]),
		Attributes:  CastAttributes(core.CastArray(data["attributes"])),
		RoleName:    core.CastString(data["roleName"]),
		DenyUserIds: core.CastStrings(core.CastArray(data["denyUserIds"])),
		CreatedAt:   core.CastInt64(data["createdAt"]),
	}
}

func (p Player) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var attributes []interface{}
	if p.Attributes != nil {
		attributes = CastAttributesFromDict(
			p.Attributes,
		)
	}
	var roleName *string
	if p.RoleName != nil {
		roleName = p.RoleName
	}
	var denyUserIds []interface{}
	if p.DenyUserIds != nil {
		denyUserIds = core.CastStringsFromDict(
			p.DenyUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"userId":      userId,
		"attributes":  attributes,
		"roleName":    roleName,
		"denyUserIds": denyUserIds,
		"createdAt":   createdAt,
	}
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
		RatingId:  core.CastString(data["ratingId"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		RateValue: core.CastFloat32(data["rateValue"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Rating) ToDict() map[string]interface{} {

	var ratingId *string
	if p.RatingId != nil {
		ratingId = p.RatingId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var rateValue *float32
	if p.RateValue != nil {
		rateValue = p.RateValue
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"ratingId":  ratingId,
		"name":      name,
		"userId":    userId,
		"rateValue": rateValue,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
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
		Rank:   core.CastInt32(data["rank"]),
		UserId: core.CastString(data["userId"]),
	}
}

func (p GameResult) ToDict() map[string]interface{} {

	var rank *int32
	if p.Rank != nil {
		rank = p.Rank
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	return map[string]interface{}{
		"rank":   rank,
		"userId": userId,
	}
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
		UserId:         core.CastString(data["userId"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
	}
}

func (p Ballot) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var ratingName *string
	if p.RatingName != nil {
		ratingName = p.RatingName
	}
	var gatheringName *string
	if p.GatheringName != nil {
		gatheringName = p.GatheringName
	}
	var numberOfPlayer *int32
	if p.NumberOfPlayer != nil {
		numberOfPlayer = p.NumberOfPlayer
	}
	return map[string]interface{}{
		"userId":         userId,
		"ratingName":     ratingName,
		"gatheringName":  gatheringName,
		"numberOfPlayer": numberOfPlayer,
	}
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
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p SignedBallot) ToDict() map[string]interface{} {

	var body *string
	if p.Body != nil {
		body = p.Body
	}
	var signature *string
	if p.Signature != nil {
		signature = p.Signature
	}
	return map[string]interface{}{
		"body":      body,
		"signature": signature,
	}
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
		Ballot:      NewBallotFromDict(core.CastMap(data["ballot"])).Pointer(),
		GameResults: CastGameResults(core.CastArray(data["gameResults"])),
	}
}

func (p WrittenBallot) ToDict() map[string]interface{} {

	var ballot map[string]interface{}
	if p.Ballot != nil {
		ballot = p.Ballot.ToDict()
	}
	var gameResults []interface{}
	if p.GameResults != nil {
		gameResults = CastGameResultsFromDict(
			p.GameResults,
		)
	}
	return map[string]interface{}{
		"ballot":      ballot,
		"gameResults": gameResults,
	}
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
		VoteId:         core.CastString(data["voteId"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		WrittenBallots: CastWrittenBallots(core.CastArray(data["writtenBallots"])),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
	}
}

func (p Vote) ToDict() map[string]interface{} {

	var voteId *string
	if p.VoteId != nil {
		voteId = p.VoteId
	}
	var ratingName *string
	if p.RatingName != nil {
		ratingName = p.RatingName
	}
	var gatheringName *string
	if p.GatheringName != nil {
		gatheringName = p.GatheringName
	}
	var writtenBallots []interface{}
	if p.WrittenBallots != nil {
		writtenBallots = CastWrittenBallotsFromDict(
			p.WrittenBallots,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	return map[string]interface{}{
		"voteId":         voteId,
		"ratingName":     ratingName,
		"gatheringName":  gatheringName,
		"writtenBallots": writtenBallots,
		"createdAt":      createdAt,
		"updatedAt":      updatedAt,
	}
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
		Days:    core.CastInt32(data["days"]),
		Hours:   core.CastInt32(data["hours"]),
		Minutes: core.CastInt32(data["minutes"]),
	}
}

func (p TimeSpan) ToDict() map[string]interface{} {

	var days *int32
	if p.Days != nil {
		days = p.Days
	}
	var hours *int32
	if p.Hours != nil {
		hours = p.Hours
	}
	var minutes *int32
	if p.Minutes != nil {
		minutes = p.Minutes
	}
	return map[string]interface{}{
		"days":    days,
		"hours":   hours,
		"minutes": minutes,
	}
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
