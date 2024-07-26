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

package mission

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Complete struct {
	CompleteId                *string   `json:"completeId"`
	UserId                    *string   `json:"userId"`
	MissionGroupName          *string   `json:"missionGroupName"`
	CompletedMissionTaskNames []*string `json:"completedMissionTaskNames"`
	ReceivedMissionTaskNames  []*string `json:"receivedMissionTaskNames"`
	NextResetAt               *int64    `json:"nextResetAt"`
	CreatedAt                 *int64    `json:"createdAt"`
	UpdatedAt                 *int64    `json:"updatedAt"`
	Revision                  *int64    `json:"revision"`
}

func (p *Complete) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Complete{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Complete{}
	} else {
		*p = Complete{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["completeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteId)
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
		if v, ok := d["missionGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionGroupName)
				}
			}
		}
		if v, ok := d["completedMissionTaskNames"]; ok && v != nil {
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
				p.CompletedMissionTaskNames = l
			}
		}
		if v, ok := d["receivedMissionTaskNames"]; ok && v != nil {
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
				p.ReceivedMissionTaskNames = l
			}
		}
		if v, ok := d["nextResetAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NextResetAt)
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

func NewCompleteFromJson(data string) Complete {
	req := Complete{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCompleteFromDict(data map[string]interface{}) Complete {
	return Complete{
		CompleteId:                core.CastString(data["completeId"]),
		UserId:                    core.CastString(data["userId"]),
		MissionGroupName:          core.CastString(data["missionGroupName"]),
		CompletedMissionTaskNames: core.CastStrings(core.CastArray(data["completedMissionTaskNames"])),
		ReceivedMissionTaskNames:  core.CastStrings(core.CastArray(data["receivedMissionTaskNames"])),
		NextResetAt:               core.CastInt64(data["nextResetAt"]),
		CreatedAt:                 core.CastInt64(data["createdAt"]),
		UpdatedAt:                 core.CastInt64(data["updatedAt"]),
		Revision:                  core.CastInt64(data["revision"]),
	}
}

func (p Complete) ToDict() map[string]interface{} {

	var completeId *string
	if p.CompleteId != nil {
		completeId = p.CompleteId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var missionGroupName *string
	if p.MissionGroupName != nil {
		missionGroupName = p.MissionGroupName
	}
	var completedMissionTaskNames []interface{}
	if p.CompletedMissionTaskNames != nil {
		completedMissionTaskNames = core.CastStringsFromDict(
			p.CompletedMissionTaskNames,
		)
	}
	var receivedMissionTaskNames []interface{}
	if p.ReceivedMissionTaskNames != nil {
		receivedMissionTaskNames = core.CastStringsFromDict(
			p.ReceivedMissionTaskNames,
		)
	}
	var nextResetAt *int64
	if p.NextResetAt != nil {
		nextResetAt = p.NextResetAt
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
		"completeId":                completeId,
		"userId":                    userId,
		"missionGroupName":          missionGroupName,
		"completedMissionTaskNames": completedMissionTaskNames,
		"receivedMissionTaskNames":  receivedMissionTaskNames,
		"nextResetAt":               nextResetAt,
		"createdAt":                 createdAt,
		"updatedAt":                 updatedAt,
		"revision":                  revision,
	}
}

func (p Complete) Pointer() *Complete {
	return &p
}

func CastCompletes(data []interface{}) []Complete {
	v := make([]Complete, 0)
	for _, d := range data {
		v = append(v, NewCompleteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCompletesFromDict(data []Complete) []interface{} {
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

type CounterModelMaster struct {
	CounterId              *string             `json:"counterId"`
	Name                   *string             `json:"name"`
	Metadata               *string             `json:"metadata"`
	Description            *string             `json:"description"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
	CreatedAt              *int64              `json:"createdAt"`
	UpdatedAt              *int64              `json:"updatedAt"`
	Revision               *int64              `json:"revision"`
}

func (p *CounterModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CounterModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CounterModelMaster{}
	} else {
		*p = CounterModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["counterId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterId)
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
		if v, ok := d["scopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Scopes)
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

func NewCounterModelMasterFromJson(data string) CounterModelMaster {
	req := CounterModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCounterModelMasterFromDict(data map[string]interface{}) CounterModelMaster {
	return CounterModelMaster{
		CounterId:              core.CastString(data["counterId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		Scopes:                 CastCounterScopeModels(core.CastArray(data["scopes"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p CounterModelMaster) ToDict() map[string]interface{} {

	var counterId *string
	if p.CounterId != nil {
		counterId = p.CounterId
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
	var scopes []interface{}
	if p.Scopes != nil {
		scopes = CastCounterScopeModelsFromDict(
			p.Scopes,
		)
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
		"counterId":              counterId,
		"name":                   name,
		"metadata":               metadata,
		"description":            description,
		"scopes":                 scopes,
		"challengePeriodEventId": challengePeriodEventId,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
}

func (p CounterModelMaster) Pointer() *CounterModelMaster {
	return &p
}

func CastCounterModelMasters(data []interface{}) []CounterModelMaster {
	v := make([]CounterModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCounterModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterModelMastersFromDict(data []CounterModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CounterScopeModel struct {
	ResetType       *string `json:"resetType"`
	ResetDayOfMonth *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek  *string `json:"resetDayOfWeek"`
	ResetHour       *int32  `json:"resetHour"`
}

func (p *CounterScopeModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CounterScopeModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CounterScopeModel{}
	} else {
		*p = CounterScopeModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["resetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetType)
				}
			}
		}
		if v, ok := d["resetDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetDayOfMonth)
		}
		if v, ok := d["resetDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetDayOfWeek)
				}
			}
		}
		if v, ok := d["resetHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetHour)
		}
	}
	return nil
}

func NewCounterScopeModelFromJson(data string) CounterScopeModel {
	req := CounterScopeModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCounterScopeModelFromDict(data map[string]interface{}) CounterScopeModel {
	return CounterScopeModel{
		ResetType:       core.CastString(data["resetType"]),
		ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:  core.CastString(data["resetDayOfWeek"]),
		ResetHour:       core.CastInt32(data["resetHour"]),
	}
}

func (p CounterScopeModel) ToDict() map[string]interface{} {

	var resetType *string
	if p.ResetType != nil {
		resetType = p.ResetType
	}
	var resetDayOfMonth *int32
	if p.ResetDayOfMonth != nil {
		resetDayOfMonth = p.ResetDayOfMonth
	}
	var resetDayOfWeek *string
	if p.ResetDayOfWeek != nil {
		resetDayOfWeek = p.ResetDayOfWeek
	}
	var resetHour *int32
	if p.ResetHour != nil {
		resetHour = p.ResetHour
	}
	return map[string]interface{}{
		"resetType":       resetType,
		"resetDayOfMonth": resetDayOfMonth,
		"resetDayOfWeek":  resetDayOfWeek,
		"resetHour":       resetHour,
	}
}

func (p CounterScopeModel) Pointer() *CounterScopeModel {
	return &p
}

func CastCounterScopeModels(data []interface{}) []CounterScopeModel {
	v := make([]CounterScopeModel, 0)
	for _, d := range data {
		v = append(v, NewCounterScopeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterScopeModelsFromDict(data []CounterScopeModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionGroupModelMaster struct {
	MissionGroupId                  *string `json:"missionGroupId"`
	Name                            *string `json:"name"`
	Metadata                        *string `json:"metadata"`
	Description                     *string `json:"description"`
	ResetType                       *string `json:"resetType"`
	ResetDayOfMonth                 *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string `json:"resetDayOfWeek"`
	ResetHour                       *int32  `json:"resetHour"`
	CompleteNotificationNamespaceId *string `json:"completeNotificationNamespaceId"`
	CreatedAt                       *int64  `json:"createdAt"`
	UpdatedAt                       *int64  `json:"updatedAt"`
	Revision                        *int64  `json:"revision"`
}

func (p *MissionGroupModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MissionGroupModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MissionGroupModelMaster{}
	} else {
		*p = MissionGroupModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["missionGroupId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionGroupId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionGroupId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionGroupId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionGroupId)
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
		if v, ok := d["resetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetType)
				}
			}
		}
		if v, ok := d["resetDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetDayOfMonth)
		}
		if v, ok := d["resetDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetDayOfWeek)
				}
			}
		}
		if v, ok := d["resetHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetHour)
		}
		if v, ok := d["completeNotificationNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteNotificationNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteNotificationNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteNotificationNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteNotificationNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteNotificationNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteNotificationNamespaceId)
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

func NewMissionGroupModelMasterFromJson(data string) MissionGroupModelMaster {
	req := MissionGroupModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMissionGroupModelMasterFromDict(data map[string]interface{}) MissionGroupModelMaster {
	return MissionGroupModelMaster{
		MissionGroupId:                  core.CastString(data["missionGroupId"]),
		Name:                            core.CastString(data["name"]),
		Metadata:                        core.CastString(data["metadata"]),
		Description:                     core.CastString(data["description"]),
		ResetType:                       core.CastString(data["resetType"]),
		ResetDayOfMonth:                 core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:                  core.CastString(data["resetDayOfWeek"]),
		ResetHour:                       core.CastInt32(data["resetHour"]),
		CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
		CreatedAt:                       core.CastInt64(data["createdAt"]),
		UpdatedAt:                       core.CastInt64(data["updatedAt"]),
		Revision:                        core.CastInt64(data["revision"]),
	}
}

func (p MissionGroupModelMaster) ToDict() map[string]interface{} {

	var missionGroupId *string
	if p.MissionGroupId != nil {
		missionGroupId = p.MissionGroupId
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
	var resetType *string
	if p.ResetType != nil {
		resetType = p.ResetType
	}
	var resetDayOfMonth *int32
	if p.ResetDayOfMonth != nil {
		resetDayOfMonth = p.ResetDayOfMonth
	}
	var resetDayOfWeek *string
	if p.ResetDayOfWeek != nil {
		resetDayOfWeek = p.ResetDayOfWeek
	}
	var resetHour *int32
	if p.ResetHour != nil {
		resetHour = p.ResetHour
	}
	var completeNotificationNamespaceId *string
	if p.CompleteNotificationNamespaceId != nil {
		completeNotificationNamespaceId = p.CompleteNotificationNamespaceId
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
		"missionGroupId":                  missionGroupId,
		"name":                            name,
		"metadata":                        metadata,
		"description":                     description,
		"resetType":                       resetType,
		"resetDayOfMonth":                 resetDayOfMonth,
		"resetDayOfWeek":                  resetDayOfWeek,
		"resetHour":                       resetHour,
		"completeNotificationNamespaceId": completeNotificationNamespaceId,
		"createdAt":                       createdAt,
		"updatedAt":                       updatedAt,
		"revision":                        revision,
	}
}

func (p MissionGroupModelMaster) Pointer() *MissionGroupModelMaster {
	return &p
}

func CastMissionGroupModelMasters(data []interface{}) []MissionGroupModelMaster {
	v := make([]MissionGroupModelMaster, 0)
	for _, d := range data {
		v = append(v, NewMissionGroupModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionGroupModelMastersFromDict(data []MissionGroupModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Namespace struct {
	NamespaceId            *string              `json:"namespaceId"`
	Name                   *string              `json:"name"`
	Description            *string              `json:"description"`
	TransactionSetting     *TransactionSetting  `json:"transactionSetting"`
	MissionCompleteScript  *ScriptSetting       `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting       `json:"counterIncrementScript"`
	ReceiveRewardsScript   *ScriptSetting       `json:"receiveRewardsScript"`
	CompleteNotification   *NotificationSetting `json:"completeNotification"`
	LogSetting             *LogSetting          `json:"logSetting"`
	CreatedAt              *int64               `json:"createdAt"`
	UpdatedAt              *int64               `json:"updatedAt"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId    *string `json:"keyId"`
	Revision *int64  `json:"revision"`
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
		if v, ok := d["missionCompleteScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MissionCompleteScript)
		}
		if v, ok := d["counterIncrementScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CounterIncrementScript)
		}
		if v, ok := d["receiveRewardsScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceiveRewardsScript)
		}
		if v, ok := d["completeNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteNotification)
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
		NamespaceId:            core.CastString(data["namespaceId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		TransactionSetting:     NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		MissionCompleteScript:  NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
		CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
		ReceiveRewardsScript:   NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
		CompleteNotification:   NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:             NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		Revision:               core.CastInt64(data["revision"]),
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
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var missionCompleteScript map[string]interface{}
	if p.MissionCompleteScript != nil {
		missionCompleteScript = p.MissionCompleteScript.ToDict()
	}
	var counterIncrementScript map[string]interface{}
	if p.CounterIncrementScript != nil {
		counterIncrementScript = p.CounterIncrementScript.ToDict()
	}
	var receiveRewardsScript map[string]interface{}
	if p.ReceiveRewardsScript != nil {
		receiveRewardsScript = p.ReceiveRewardsScript.ToDict()
	}
	var completeNotification map[string]interface{}
	if p.CompleteNotification != nil {
		completeNotification = p.CompleteNotification.ToDict()
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
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":            namespaceId,
		"name":                   name,
		"description":            description,
		"transactionSetting":     transactionSetting,
		"missionCompleteScript":  missionCompleteScript,
		"counterIncrementScript": counterIncrementScript,
		"receiveRewardsScript":   receiveRewardsScript,
		"completeNotification":   completeNotification,
		"logSetting":             logSetting,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"queueNamespaceId":       queueNamespaceId,
		"keyId":                  keyId,
		"revision":               revision,
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

type Counter struct {
	CounterId *string       `json:"counterId"`
	UserId    *string       `json:"userId"`
	Name      *string       `json:"name"`
	Values    []ScopedValue `json:"values"`
	CreatedAt *int64        `json:"createdAt"`
	UpdatedAt *int64        `json:"updatedAt"`
	Revision  *int64        `json:"revision"`
}

func (p *Counter) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Counter{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Counter{}
	} else {
		*p = Counter{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["counterId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterId)
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
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

func NewCounterFromJson(data string) Counter {
	req := Counter{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCounterFromDict(data map[string]interface{}) Counter {
	return Counter{
		CounterId: core.CastString(data["counterId"]),
		UserId:    core.CastString(data["userId"]),
		Name:      core.CastString(data["name"]),
		Values:    CastScopedValues(core.CastArray(data["values"])),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Counter) ToDict() map[string]interface{} {

	var counterId *string
	if p.CounterId != nil {
		counterId = p.CounterId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var values []interface{}
	if p.Values != nil {
		values = CastScopedValuesFromDict(
			p.Values,
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"counterId": counterId,
		"userId":    userId,
		"name":      name,
		"values":    values,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
}

func (p Counter) Pointer() *Counter {
	return &p
}

func CastCounters(data []interface{}) []Counter {
	v := make([]Counter, 0)
	for _, d := range data {
		v = append(v, NewCounterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCountersFromDict(data []Counter) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentMissionMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentMissionMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentMissionMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentMissionMaster{}
	} else {
		*p = CurrentMissionMaster{}
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

func NewCurrentMissionMasterFromJson(data string) CurrentMissionMaster {
	req := CurrentMissionMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentMissionMasterFromDict(data map[string]interface{}) CurrentMissionMaster {
	return CurrentMissionMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentMissionMaster) ToDict() map[string]interface{} {

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

func (p CurrentMissionMaster) Pointer() *CurrentMissionMaster {
	return &p
}

func CastCurrentMissionMasters(data []interface{}) []CurrentMissionMaster {
	v := make([]CurrentMissionMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentMissionMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentMissionMastersFromDict(data []CurrentMissionMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CounterModel struct {
	CounterId              *string             `json:"counterId"`
	Name                   *string             `json:"name"`
	Metadata               *string             `json:"metadata"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
}

func (p *CounterModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CounterModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CounterModel{}
	} else {
		*p = CounterModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["counterId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterId)
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
		if v, ok := d["scopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Scopes)
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

func NewCounterModelFromJson(data string) CounterModel {
	req := CounterModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCounterModelFromDict(data map[string]interface{}) CounterModel {
	return CounterModel{
		CounterId:              core.CastString(data["counterId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Scopes:                 CastCounterScopeModels(core.CastArray(data["scopes"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p CounterModel) ToDict() map[string]interface{} {

	var counterId *string
	if p.CounterId != nil {
		counterId = p.CounterId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var scopes []interface{}
	if p.Scopes != nil {
		scopes = CastCounterScopeModelsFromDict(
			p.Scopes,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	return map[string]interface{}{
		"counterId":              counterId,
		"name":                   name,
		"metadata":               metadata,
		"scopes":                 scopes,
		"challengePeriodEventId": challengePeriodEventId,
	}
}

func (p CounterModel) Pointer() *CounterModel {
	return &p
}

func CastCounterModels(data []interface{}) []CounterModel {
	v := make([]CounterModel, 0)
	for _, d := range data {
		v = append(v, NewCounterModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCounterModelsFromDict(data []CounterModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionGroupModel struct {
	MissionGroupId                  *string            `json:"missionGroupId"`
	Name                            *string            `json:"name"`
	Metadata                        *string            `json:"metadata"`
	Tasks                           []MissionTaskModel `json:"tasks"`
	ResetType                       *string            `json:"resetType"`
	ResetDayOfMonth                 *int32             `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string            `json:"resetDayOfWeek"`
	ResetHour                       *int32             `json:"resetHour"`
	CompleteNotificationNamespaceId *string            `json:"completeNotificationNamespaceId"`
}

func (p *MissionGroupModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MissionGroupModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MissionGroupModel{}
	} else {
		*p = MissionGroupModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["missionGroupId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionGroupId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionGroupId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionGroupId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionGroupId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionGroupId)
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
		if v, ok := d["tasks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tasks)
		}
		if v, ok := d["resetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetType)
				}
			}
		}
		if v, ok := d["resetDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetDayOfMonth)
		}
		if v, ok := d["resetDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetDayOfWeek)
				}
			}
		}
		if v, ok := d["resetHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetHour)
		}
		if v, ok := d["completeNotificationNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CompleteNotificationNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CompleteNotificationNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CompleteNotificationNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CompleteNotificationNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CompleteNotificationNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CompleteNotificationNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewMissionGroupModelFromJson(data string) MissionGroupModel {
	req := MissionGroupModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMissionGroupModelFromDict(data map[string]interface{}) MissionGroupModel {
	return MissionGroupModel{
		MissionGroupId:                  core.CastString(data["missionGroupId"]),
		Name:                            core.CastString(data["name"]),
		Metadata:                        core.CastString(data["metadata"]),
		Tasks:                           CastMissionTaskModels(core.CastArray(data["tasks"])),
		ResetType:                       core.CastString(data["resetType"]),
		ResetDayOfMonth:                 core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:                  core.CastString(data["resetDayOfWeek"]),
		ResetHour:                       core.CastInt32(data["resetHour"]),
		CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
	}
}

func (p MissionGroupModel) ToDict() map[string]interface{} {

	var missionGroupId *string
	if p.MissionGroupId != nil {
		missionGroupId = p.MissionGroupId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var tasks []interface{}
	if p.Tasks != nil {
		tasks = CastMissionTaskModelsFromDict(
			p.Tasks,
		)
	}
	var resetType *string
	if p.ResetType != nil {
		resetType = p.ResetType
	}
	var resetDayOfMonth *int32
	if p.ResetDayOfMonth != nil {
		resetDayOfMonth = p.ResetDayOfMonth
	}
	var resetDayOfWeek *string
	if p.ResetDayOfWeek != nil {
		resetDayOfWeek = p.ResetDayOfWeek
	}
	var resetHour *int32
	if p.ResetHour != nil {
		resetHour = p.ResetHour
	}
	var completeNotificationNamespaceId *string
	if p.CompleteNotificationNamespaceId != nil {
		completeNotificationNamespaceId = p.CompleteNotificationNamespaceId
	}
	return map[string]interface{}{
		"missionGroupId":                  missionGroupId,
		"name":                            name,
		"metadata":                        metadata,
		"tasks":                           tasks,
		"resetType":                       resetType,
		"resetDayOfMonth":                 resetDayOfMonth,
		"resetDayOfWeek":                  resetDayOfWeek,
		"resetHour":                       resetHour,
		"completeNotificationNamespaceId": completeNotificationNamespaceId,
	}
}

func (p MissionGroupModel) Pointer() *MissionGroupModel {
	return &p
}

func CastMissionGroupModels(data []interface{}) []MissionGroupModel {
	v := make([]MissionGroupModel, 0)
	for _, d := range data {
		v = append(v, NewMissionGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionGroupModelsFromDict(data []MissionGroupModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionTaskModel struct {
	MissionTaskId                *string             `json:"missionTaskId"`
	Name                         *string             `json:"name"`
	Metadata                     *string             `json:"metadata"`
	VerifyCompleteType           *string             `json:"verifyCompleteType"`
	TargetCounter                *TargetCounterModel `json:"targetCounter"`
	VerifyCompleteConsumeActions []VerifyAction      `json:"verifyCompleteConsumeActions"`
	CompleteAcquireActions       []AcquireAction     `json:"completeAcquireActions"`
	ChallengePeriodEventId       *string             `json:"challengePeriodEventId"`
	PremiseMissionTaskName       *string             `json:"premiseMissionTaskName"`
	// Deprecated: should not be used
	CounterName *string `json:"counterName"`
	// Deprecated: should not be used
	TargetResetType *string `json:"targetResetType"`
	// Deprecated: should not be used
	TargetValue *int64 `json:"targetValue"`
}

func (p *MissionTaskModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MissionTaskModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MissionTaskModel{}
	} else {
		*p = MissionTaskModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["missionTaskId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionTaskId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionTaskId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionTaskId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionTaskId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionTaskId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionTaskId)
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
		if v, ok := d["verifyCompleteType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyCompleteType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyCompleteType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyCompleteType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyCompleteType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyCompleteType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyCompleteType)
				}
			}
		}
		if v, ok := d["targetCounter"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetCounter)
		}
		if v, ok := d["verifyCompleteConsumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyCompleteConsumeActions)
		}
		if v, ok := d["completeAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteAcquireActions)
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
		if v, ok := d["premiseMissionTaskName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PremiseMissionTaskName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PremiseMissionTaskName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PremiseMissionTaskName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PremiseMissionTaskName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PremiseMissionTaskName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PremiseMissionTaskName)
				}
			}
		}
		if v, ok := d["counterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterName)
				}
			}
		}
		if v, ok := d["targetResetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetResetType)
				}
			}
		}
		if v, ok := d["targetValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetValue)
		}
	}
	return nil
}

func NewMissionTaskModelFromJson(data string) MissionTaskModel {
	req := MissionTaskModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMissionTaskModelFromDict(data map[string]interface{}) MissionTaskModel {
	return MissionTaskModel{
		MissionTaskId:                core.CastString(data["missionTaskId"]),
		Name:                         core.CastString(data["name"]),
		Metadata:                     core.CastString(data["metadata"]),
		VerifyCompleteType:           core.CastString(data["verifyCompleteType"]),
		TargetCounter:                NewTargetCounterModelFromDict(core.CastMap(data["targetCounter"])).Pointer(),
		VerifyCompleteConsumeActions: CastVerifyActions(core.CastArray(data["verifyCompleteConsumeActions"])),
		CompleteAcquireActions:       CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		ChallengePeriodEventId:       core.CastString(data["challengePeriodEventId"]),
		PremiseMissionTaskName:       core.CastString(data["premiseMissionTaskName"]),
		CounterName:                  core.CastString(data["counterName"]),
		TargetResetType:              core.CastString(data["targetResetType"]),
		TargetValue:                  core.CastInt64(data["targetValue"]),
	}
}

func (p MissionTaskModel) ToDict() map[string]interface{} {

	var missionTaskId *string
	if p.MissionTaskId != nil {
		missionTaskId = p.MissionTaskId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var verifyCompleteType *string
	if p.VerifyCompleteType != nil {
		verifyCompleteType = p.VerifyCompleteType
	}
	var targetCounter map[string]interface{}
	if p.TargetCounter != nil {
		targetCounter = p.TargetCounter.ToDict()
	}
	var verifyCompleteConsumeActions []interface{}
	if p.VerifyCompleteConsumeActions != nil {
		verifyCompleteConsumeActions = CastVerifyActionsFromDict(
			p.VerifyCompleteConsumeActions,
		)
	}
	var completeAcquireActions []interface{}
	if p.CompleteAcquireActions != nil {
		completeAcquireActions = CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var premiseMissionTaskName *string
	if p.PremiseMissionTaskName != nil {
		premiseMissionTaskName = p.PremiseMissionTaskName
	}
	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var targetResetType *string
	if p.TargetResetType != nil {
		targetResetType = p.TargetResetType
	}
	var targetValue *int64
	if p.TargetValue != nil {
		targetValue = p.TargetValue
	}
	return map[string]interface{}{
		"missionTaskId":                missionTaskId,
		"name":                         name,
		"metadata":                     metadata,
		"verifyCompleteType":           verifyCompleteType,
		"targetCounter":                targetCounter,
		"verifyCompleteConsumeActions": verifyCompleteConsumeActions,
		"completeAcquireActions":       completeAcquireActions,
		"challengePeriodEventId":       challengePeriodEventId,
		"premiseMissionTaskName":       premiseMissionTaskName,
		"counterName":                  counterName,
		"targetResetType":              targetResetType,
		"targetValue":                  targetValue,
	}
}

func (p MissionTaskModel) Pointer() *MissionTaskModel {
	return &p
}

func CastMissionTaskModels(data []interface{}) []MissionTaskModel {
	v := make([]MissionTaskModel, 0)
	for _, d := range data {
		v = append(v, NewMissionTaskModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionTaskModelsFromDict(data []MissionTaskModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MissionTaskModelMaster struct {
	MissionTaskId                *string             `json:"missionTaskId"`
	Name                         *string             `json:"name"`
	Metadata                     *string             `json:"metadata"`
	Description                  *string             `json:"description"`
	VerifyCompleteType           *string             `json:"verifyCompleteType"`
	TargetCounter                *TargetCounterModel `json:"targetCounter"`
	VerifyCompleteConsumeActions []VerifyAction      `json:"verifyCompleteConsumeActions"`
	CompleteAcquireActions       []AcquireAction     `json:"completeAcquireActions"`
	ChallengePeriodEventId       *string             `json:"challengePeriodEventId"`
	PremiseMissionTaskName       *string             `json:"premiseMissionTaskName"`
	CreatedAt                    *int64              `json:"createdAt"`
	UpdatedAt                    *int64              `json:"updatedAt"`
	Revision                     *int64              `json:"revision"`
	// Deprecated: should not be used
	CounterName *string `json:"counterName"`
	// Deprecated: should not be used
	TargetResetType *string `json:"targetResetType"`
	// Deprecated: should not be used
	TargetValue *int64 `json:"targetValue"`
}

func (p *MissionTaskModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MissionTaskModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MissionTaskModelMaster{}
	} else {
		*p = MissionTaskModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["missionTaskId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MissionTaskId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MissionTaskId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MissionTaskId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MissionTaskId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MissionTaskId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MissionTaskId)
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
		if v, ok := d["verifyCompleteType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyCompleteType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyCompleteType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyCompleteType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyCompleteType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyCompleteType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyCompleteType)
				}
			}
		}
		if v, ok := d["targetCounter"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetCounter)
		}
		if v, ok := d["verifyCompleteConsumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyCompleteConsumeActions)
		}
		if v, ok := d["completeAcquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CompleteAcquireActions)
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
		if v, ok := d["premiseMissionTaskName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PremiseMissionTaskName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PremiseMissionTaskName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PremiseMissionTaskName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PremiseMissionTaskName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PremiseMissionTaskName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PremiseMissionTaskName)
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
		if v, ok := d["counterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterName)
				}
			}
		}
		if v, ok := d["targetResetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetResetType)
				}
			}
		}
		if v, ok := d["targetValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetValue)
		}
	}
	return nil
}

func NewMissionTaskModelMasterFromJson(data string) MissionTaskModelMaster {
	req := MissionTaskModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMissionTaskModelMasterFromDict(data map[string]interface{}) MissionTaskModelMaster {
	return MissionTaskModelMaster{
		MissionTaskId:                core.CastString(data["missionTaskId"]),
		Name:                         core.CastString(data["name"]),
		Metadata:                     core.CastString(data["metadata"]),
		Description:                  core.CastString(data["description"]),
		VerifyCompleteType:           core.CastString(data["verifyCompleteType"]),
		TargetCounter:                NewTargetCounterModelFromDict(core.CastMap(data["targetCounter"])).Pointer(),
		VerifyCompleteConsumeActions: CastVerifyActions(core.CastArray(data["verifyCompleteConsumeActions"])),
		CompleteAcquireActions:       CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		ChallengePeriodEventId:       core.CastString(data["challengePeriodEventId"]),
		PremiseMissionTaskName:       core.CastString(data["premiseMissionTaskName"]),
		CreatedAt:                    core.CastInt64(data["createdAt"]),
		UpdatedAt:                    core.CastInt64(data["updatedAt"]),
		Revision:                     core.CastInt64(data["revision"]),
		CounterName:                  core.CastString(data["counterName"]),
		TargetResetType:              core.CastString(data["targetResetType"]),
		TargetValue:                  core.CastInt64(data["targetValue"]),
	}
}

func (p MissionTaskModelMaster) ToDict() map[string]interface{} {

	var missionTaskId *string
	if p.MissionTaskId != nil {
		missionTaskId = p.MissionTaskId
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
	var verifyCompleteType *string
	if p.VerifyCompleteType != nil {
		verifyCompleteType = p.VerifyCompleteType
	}
	var targetCounter map[string]interface{}
	if p.TargetCounter != nil {
		targetCounter = p.TargetCounter.ToDict()
	}
	var verifyCompleteConsumeActions []interface{}
	if p.VerifyCompleteConsumeActions != nil {
		verifyCompleteConsumeActions = CastVerifyActionsFromDict(
			p.VerifyCompleteConsumeActions,
		)
	}
	var completeAcquireActions []interface{}
	if p.CompleteAcquireActions != nil {
		completeAcquireActions = CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		)
	}
	var challengePeriodEventId *string
	if p.ChallengePeriodEventId != nil {
		challengePeriodEventId = p.ChallengePeriodEventId
	}
	var premiseMissionTaskName *string
	if p.PremiseMissionTaskName != nil {
		premiseMissionTaskName = p.PremiseMissionTaskName
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
	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var targetResetType *string
	if p.TargetResetType != nil {
		targetResetType = p.TargetResetType
	}
	var targetValue *int64
	if p.TargetValue != nil {
		targetValue = p.TargetValue
	}
	return map[string]interface{}{
		"missionTaskId":                missionTaskId,
		"name":                         name,
		"metadata":                     metadata,
		"description":                  description,
		"verifyCompleteType":           verifyCompleteType,
		"targetCounter":                targetCounter,
		"verifyCompleteConsumeActions": verifyCompleteConsumeActions,
		"completeAcquireActions":       completeAcquireActions,
		"challengePeriodEventId":       challengePeriodEventId,
		"premiseMissionTaskName":       premiseMissionTaskName,
		"createdAt":                    createdAt,
		"updatedAt":                    updatedAt,
		"revision":                     revision,
		"counterName":                  counterName,
		"targetResetType":              targetResetType,
		"targetValue":                  targetValue,
	}
}

func (p MissionTaskModelMaster) Pointer() *MissionTaskModelMaster {
	return &p
}

func CastMissionTaskModelMasters(data []interface{}) []MissionTaskModelMaster {
	v := make([]MissionTaskModelMaster, 0)
	for _, d := range data {
		v = append(v, NewMissionTaskModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMissionTaskModelMastersFromDict(data []MissionTaskModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScopedValue struct {
	ResetType   *string `json:"resetType"`
	Value       *int64  `json:"value"`
	NextResetAt *int64  `json:"nextResetAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func (p *ScopedValue) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScopedValue{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ScopedValue{}
	} else {
		*p = ScopedValue{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["resetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetType)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
		if v, ok := d["nextResetAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NextResetAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewScopedValueFromJson(data string) ScopedValue {
	req := ScopedValue{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScopedValueFromDict(data map[string]interface{}) ScopedValue {
	return ScopedValue{
		ResetType:   core.CastString(data["resetType"]),
		Value:       core.CastInt64(data["value"]),
		NextResetAt: core.CastInt64(data["nextResetAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
	}
}

func (p ScopedValue) ToDict() map[string]interface{} {

	var resetType *string
	if p.ResetType != nil {
		resetType = p.ResetType
	}
	var value *int64
	if p.Value != nil {
		value = p.Value
	}
	var nextResetAt *int64
	if p.NextResetAt != nil {
		nextResetAt = p.NextResetAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	return map[string]interface{}{
		"resetType":   resetType,
		"value":       value,
		"nextResetAt": nextResetAt,
		"updatedAt":   updatedAt,
	}
}

func (p ScopedValue) Pointer() *ScopedValue {
	return &p
}

func CastScopedValues(data []interface{}) []ScopedValue {
	v := make([]ScopedValue, 0)
	for _, d := range data {
		v = append(v, NewScopedValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScopedValuesFromDict(data []ScopedValue) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TargetCounterModel struct {
	CounterName *string `json:"counterName"`
	ResetType   *string `json:"resetType"`
	Value       *int64  `json:"value"`
}

func (p *TargetCounterModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TargetCounterModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = TargetCounterModel{}
	} else {
		*p = TargetCounterModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["counterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CounterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CounterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CounterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CounterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CounterName)
				}
			}
		}
		if v, ok := d["resetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResetType)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewTargetCounterModelFromJson(data string) TargetCounterModel {
	req := TargetCounterModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTargetCounterModelFromDict(data map[string]interface{}) TargetCounterModel {
	return TargetCounterModel{
		CounterName: core.CastString(data["counterName"]),
		ResetType:   core.CastString(data["resetType"]),
		Value:       core.CastInt64(data["value"]),
	}
}

func (p TargetCounterModel) ToDict() map[string]interface{} {

	var counterName *string
	if p.CounterName != nil {
		counterName = p.CounterName
	}
	var resetType *string
	if p.ResetType != nil {
		resetType = p.ResetType
	}
	var value *int64
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"counterName": counterName,
		"resetType":   resetType,
		"value":       value,
	}
}

func (p TargetCounterModel) Pointer() *TargetCounterModel {
	return &p
}

func CastTargetCounterModels(data []interface{}) []TargetCounterModel {
	v := make([]TargetCounterModel, 0)
	for _, d := range data {
		v = append(v, NewTargetCounterModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTargetCounterModelsFromDict(data []TargetCounterModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func (p *AcquireAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireAction{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcquireAction{}
	} else {
		*p = AcquireAction{}
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
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
		}
	}
	return nil
}

func NewAcquireActionFromJson(data string) AcquireAction {
	req := AcquireAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p AcquireAction) Pointer() *AcquireAction {
	return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func (p *ConsumeAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeAction{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeAction{}
	} else {
		*p = ConsumeAction{}
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
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
		}
	}
	return nil
}

func NewConsumeActionFromJson(data string) ConsumeAction {
	req := ConsumeAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
	return ConsumeAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p ConsumeAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p ConsumeAction) Pointer() *ConsumeAction {
	return &p
}

func CastConsumeActions(data []interface{}) []ConsumeAction {
	v := make([]ConsumeAction, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionsFromDict(data []ConsumeAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VerifyAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func (p *VerifyAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyAction{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = VerifyAction{}
	} else {
		*p = VerifyAction{}
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
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
		}
	}
	return nil
}

func NewVerifyActionFromJson(data string) VerifyAction {
	req := VerifyAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVerifyActionFromDict(data map[string]interface{}) VerifyAction {
	return VerifyAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p VerifyAction) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p VerifyAction) Pointer() *VerifyAction {
	return &p
}

func CastVerifyActions(data []interface{}) []VerifyAction {
	v := make([]VerifyAction, 0)
	for _, d := range data {
		v = append(v, NewVerifyActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVerifyActionsFromDict(data []VerifyAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Config struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *Config) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Config{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Config{}
	} else {
		*p = Config{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
	}
	return nil
}

func NewConfigFromJson(data string) Config {
	req := Config{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Config) ToDict() map[string]interface{} {

	var key *string
	if p.Key != nil {
		key = p.Key
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"key":   key,
		"value": value,
	}
}

func (p Config) Pointer() *Config {
	return &p
}

func CastConfigs(data []interface{}) []Config {
	v := make([]Config, 0)
	for _, d := range data {
		v = append(v, NewConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConfigsFromDict(data []Config) []interface{} {
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

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
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
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
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
