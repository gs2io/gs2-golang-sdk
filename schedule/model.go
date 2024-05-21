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

package schedule

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string     `json:"namespaceId"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	LogSetting  *LogSetting `json:"logSetting"`
	CreatedAt   *int64      `json:"createdAt"`
	UpdatedAt   *int64      `json:"updatedAt"`
	Revision    *int64      `json:"revision"`
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
		NamespaceId: core.CastString(data["namespaceId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
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
		"namespaceId": namespaceId,
		"name":        name,
		"description": description,
		"logSetting":  logSetting,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
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

type EventMaster struct {
	EventId             *string        `json:"eventId"`
	Name                *string        `json:"name"`
	Description         *string        `json:"description"`
	Metadata            *string        `json:"metadata"`
	ScheduleType        *string        `json:"scheduleType"`
	AbsoluteBegin       *int64         `json:"absoluteBegin"`
	AbsoluteEnd         *int64         `json:"absoluteEnd"`
	RelativeTriggerName *string        `json:"relativeTriggerName"`
	RepeatSetting       *RepeatSetting `json:"repeatSetting"`
	CreatedAt           *int64         `json:"createdAt"`
	UpdatedAt           *int64         `json:"updatedAt"`
	Revision            *int64         `json:"revision"`
	// Deprecated: should not be used
	RepeatType *string `json:"repeatType"`
	// Deprecated: should not be used
	RepeatBeginDayOfMonth *int32 `json:"repeatBeginDayOfMonth"`
	// Deprecated: should not be used
	RepeatEndDayOfMonth *int32 `json:"repeatEndDayOfMonth"`
	// Deprecated: should not be used
	RepeatBeginDayOfWeek *string `json:"repeatBeginDayOfWeek"`
	// Deprecated: should not be used
	RepeatEndDayOfWeek *string `json:"repeatEndDayOfWeek"`
	// Deprecated: should not be used
	RepeatBeginHour *int32 `json:"repeatBeginHour"`
	// Deprecated: should not be used
	RepeatEndHour *int32 `json:"repeatEndHour"`
}

func (p *EventMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EventMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = EventMaster{}
	} else {
		*p = EventMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["eventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventId)
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
		if v, ok := d["scheduleType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleType)
				}
			}
		}
		if v, ok := d["absoluteBegin"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AbsoluteBegin)
		}
		if v, ok := d["absoluteEnd"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AbsoluteEnd)
		}
		if v, ok := d["relativeTriggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RelativeTriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RelativeTriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RelativeTriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RelativeTriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RelativeTriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RelativeTriggerName)
				}
			}
		}
		if v, ok := d["repeatSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatSetting)
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
		if v, ok := d["repeatType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatType)
				}
			}
		}
		if v, ok := d["repeatBeginDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatBeginDayOfMonth)
		}
		if v, ok := d["repeatEndDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatEndDayOfMonth)
		}
		if v, ok := d["repeatBeginDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatBeginDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatBeginDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatBeginDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatBeginDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatBeginDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatBeginDayOfWeek)
				}
			}
		}
		if v, ok := d["repeatEndDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatEndDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatEndDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatEndDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatEndDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatEndDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatEndDayOfWeek)
				}
			}
		}
		if v, ok := d["repeatBeginHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatBeginHour)
		}
		if v, ok := d["repeatEndHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatEndHour)
		}
	}
	return nil
}

func NewEventMasterFromJson(data string) EventMaster {
	req := EventMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewEventMasterFromDict(data map[string]interface{}) EventMaster {
	return EventMaster{
		EventId:               core.CastString(data["eventId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ScheduleType:          core.CastString(data["scheduleType"]),
		AbsoluteBegin:         core.CastInt64(data["absoluteBegin"]),
		AbsoluteEnd:           core.CastInt64(data["absoluteEnd"]),
		RelativeTriggerName:   core.CastString(data["relativeTriggerName"]),
		RepeatSetting:         NewRepeatSettingFromDict(core.CastMap(data["repeatSetting"])).Pointer(),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
		RepeatType:            core.CastString(data["repeatType"]),
		RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
		RepeatEndDayOfMonth:   core.CastInt32(data["repeatEndDayOfMonth"]),
		RepeatBeginDayOfWeek:  core.CastString(data["repeatBeginDayOfWeek"]),
		RepeatEndDayOfWeek:    core.CastString(data["repeatEndDayOfWeek"]),
		RepeatBeginHour:       core.CastInt32(data["repeatBeginHour"]),
		RepeatEndHour:         core.CastInt32(data["repeatEndHour"]),
	}
}

func (p EventMaster) ToDict() map[string]interface{} {

	var eventId *string
	if p.EventId != nil {
		eventId = p.EventId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var scheduleType *string
	if p.ScheduleType != nil {
		scheduleType = p.ScheduleType
	}
	var absoluteBegin *int64
	if p.AbsoluteBegin != nil {
		absoluteBegin = p.AbsoluteBegin
	}
	var absoluteEnd *int64
	if p.AbsoluteEnd != nil {
		absoluteEnd = p.AbsoluteEnd
	}
	var relativeTriggerName *string
	if p.RelativeTriggerName != nil {
		relativeTriggerName = p.RelativeTriggerName
	}
	var repeatSetting map[string]interface{}
	if p.RepeatSetting != nil {
		repeatSetting = p.RepeatSetting.ToDict()
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
	var repeatType *string
	if p.RepeatType != nil {
		repeatType = p.RepeatType
	}
	var repeatBeginDayOfMonth *int32
	if p.RepeatBeginDayOfMonth != nil {
		repeatBeginDayOfMonth = p.RepeatBeginDayOfMonth
	}
	var repeatEndDayOfMonth *int32
	if p.RepeatEndDayOfMonth != nil {
		repeatEndDayOfMonth = p.RepeatEndDayOfMonth
	}
	var repeatBeginDayOfWeek *string
	if p.RepeatBeginDayOfWeek != nil {
		repeatBeginDayOfWeek = p.RepeatBeginDayOfWeek
	}
	var repeatEndDayOfWeek *string
	if p.RepeatEndDayOfWeek != nil {
		repeatEndDayOfWeek = p.RepeatEndDayOfWeek
	}
	var repeatBeginHour *int32
	if p.RepeatBeginHour != nil {
		repeatBeginHour = p.RepeatBeginHour
	}
	var repeatEndHour *int32
	if p.RepeatEndHour != nil {
		repeatEndHour = p.RepeatEndHour
	}
	return map[string]interface{}{
		"eventId":               eventId,
		"name":                  name,
		"description":           description,
		"metadata":              metadata,
		"scheduleType":          scheduleType,
		"absoluteBegin":         absoluteBegin,
		"absoluteEnd":           absoluteEnd,
		"relativeTriggerName":   relativeTriggerName,
		"repeatSetting":         repeatSetting,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
		"repeatType":            repeatType,
		"repeatBeginDayOfMonth": repeatBeginDayOfMonth,
		"repeatEndDayOfMonth":   repeatEndDayOfMonth,
		"repeatBeginDayOfWeek":  repeatBeginDayOfWeek,
		"repeatEndDayOfWeek":    repeatEndDayOfWeek,
		"repeatBeginHour":       repeatBeginHour,
		"repeatEndHour":         repeatEndHour,
	}
}

func (p EventMaster) Pointer() *EventMaster {
	return &p
}

func CastEventMasters(data []interface{}) []EventMaster {
	v := make([]EventMaster, 0)
	for _, d := range data {
		v = append(v, NewEventMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventMastersFromDict(data []EventMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Trigger struct {
	TriggerId *string `json:"triggerId"`
	Name      *string `json:"name"`
	UserId    *string `json:"userId"`
	CreatedAt *int64  `json:"createdAt"`
	ExpiresAt *int64  `json:"expiresAt"`
	Revision  *int64  `json:"revision"`
}

func (p *Trigger) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Trigger{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Trigger{}
	} else {
		*p = Trigger{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["triggerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerId)
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewTriggerFromJson(data string) Trigger {
	req := Trigger{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTriggerFromDict(data map[string]interface{}) Trigger {
	return Trigger{
		TriggerId: core.CastString(data["triggerId"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		ExpiresAt: core.CastInt64(data["expiresAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Trigger) ToDict() map[string]interface{} {

	var triggerId *string
	if p.TriggerId != nil {
		triggerId = p.TriggerId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"triggerId": triggerId,
		"name":      name,
		"userId":    userId,
		"createdAt": createdAt,
		"expiresAt": expiresAt,
		"revision":  revision,
	}
}

func (p Trigger) Pointer() *Trigger {
	return &p
}

func CastTriggers(data []interface{}) []Trigger {
	v := make([]Trigger, 0)
	for _, d := range data {
		v = append(v, NewTriggerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTriggersFromDict(data []Trigger) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Event struct {
	EventId             *string        `json:"eventId"`
	Name                *string        `json:"name"`
	Metadata            *string        `json:"metadata"`
	ScheduleType        *string        `json:"scheduleType"`
	AbsoluteBegin       *int64         `json:"absoluteBegin"`
	AbsoluteEnd         *int64         `json:"absoluteEnd"`
	RelativeTriggerName *string        `json:"relativeTriggerName"`
	RepeatSetting       *RepeatSetting `json:"repeatSetting"`
	// Deprecated: should not be used
	RepeatType *string `json:"repeatType"`
	// Deprecated: should not be used
	RepeatBeginDayOfMonth *int32 `json:"repeatBeginDayOfMonth"`
	// Deprecated: should not be used
	RepeatEndDayOfMonth *int32 `json:"repeatEndDayOfMonth"`
	// Deprecated: should not be used
	RepeatBeginDayOfWeek *string `json:"repeatBeginDayOfWeek"`
	// Deprecated: should not be used
	RepeatEndDayOfWeek *string `json:"repeatEndDayOfWeek"`
	// Deprecated: should not be used
	RepeatBeginHour *int32 `json:"repeatBeginHour"`
	// Deprecated: should not be used
	RepeatEndHour *int32 `json:"repeatEndHour"`
}

func (p *Event) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Event{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Event{}
	} else {
		*p = Event{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["eventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventId)
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
		if v, ok := d["scheduleType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleType)
				}
			}
		}
		if v, ok := d["absoluteBegin"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AbsoluteBegin)
		}
		if v, ok := d["absoluteEnd"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AbsoluteEnd)
		}
		if v, ok := d["relativeTriggerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RelativeTriggerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RelativeTriggerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RelativeTriggerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RelativeTriggerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RelativeTriggerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RelativeTriggerName)
				}
			}
		}
		if v, ok := d["repeatSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatSetting)
		}
		if v, ok := d["repeatType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatType)
				}
			}
		}
		if v, ok := d["repeatBeginDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatBeginDayOfMonth)
		}
		if v, ok := d["repeatEndDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatEndDayOfMonth)
		}
		if v, ok := d["repeatBeginDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatBeginDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatBeginDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatBeginDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatBeginDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatBeginDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatBeginDayOfWeek)
				}
			}
		}
		if v, ok := d["repeatEndDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatEndDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatEndDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatEndDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatEndDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatEndDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatEndDayOfWeek)
				}
			}
		}
		if v, ok := d["repeatBeginHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatBeginHour)
		}
		if v, ok := d["repeatEndHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatEndHour)
		}
	}
	return nil
}

func NewEventFromJson(data string) Event {
	req := Event{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewEventFromDict(data map[string]interface{}) Event {
	return Event{
		EventId:               core.CastString(data["eventId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		ScheduleType:          core.CastString(data["scheduleType"]),
		AbsoluteBegin:         core.CastInt64(data["absoluteBegin"]),
		AbsoluteEnd:           core.CastInt64(data["absoluteEnd"]),
		RelativeTriggerName:   core.CastString(data["relativeTriggerName"]),
		RepeatSetting:         NewRepeatSettingFromDict(core.CastMap(data["repeatSetting"])).Pointer(),
		RepeatType:            core.CastString(data["repeatType"]),
		RepeatBeginDayOfMonth: core.CastInt32(data["repeatBeginDayOfMonth"]),
		RepeatEndDayOfMonth:   core.CastInt32(data["repeatEndDayOfMonth"]),
		RepeatBeginDayOfWeek:  core.CastString(data["repeatBeginDayOfWeek"]),
		RepeatEndDayOfWeek:    core.CastString(data["repeatEndDayOfWeek"]),
		RepeatBeginHour:       core.CastInt32(data["repeatBeginHour"]),
		RepeatEndHour:         core.CastInt32(data["repeatEndHour"]),
	}
}

func (p Event) ToDict() map[string]interface{} {

	var eventId *string
	if p.EventId != nil {
		eventId = p.EventId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var scheduleType *string
	if p.ScheduleType != nil {
		scheduleType = p.ScheduleType
	}
	var absoluteBegin *int64
	if p.AbsoluteBegin != nil {
		absoluteBegin = p.AbsoluteBegin
	}
	var absoluteEnd *int64
	if p.AbsoluteEnd != nil {
		absoluteEnd = p.AbsoluteEnd
	}
	var relativeTriggerName *string
	if p.RelativeTriggerName != nil {
		relativeTriggerName = p.RelativeTriggerName
	}
	var repeatSetting map[string]interface{}
	if p.RepeatSetting != nil {
		repeatSetting = p.RepeatSetting.ToDict()
	}
	var repeatType *string
	if p.RepeatType != nil {
		repeatType = p.RepeatType
	}
	var repeatBeginDayOfMonth *int32
	if p.RepeatBeginDayOfMonth != nil {
		repeatBeginDayOfMonth = p.RepeatBeginDayOfMonth
	}
	var repeatEndDayOfMonth *int32
	if p.RepeatEndDayOfMonth != nil {
		repeatEndDayOfMonth = p.RepeatEndDayOfMonth
	}
	var repeatBeginDayOfWeek *string
	if p.RepeatBeginDayOfWeek != nil {
		repeatBeginDayOfWeek = p.RepeatBeginDayOfWeek
	}
	var repeatEndDayOfWeek *string
	if p.RepeatEndDayOfWeek != nil {
		repeatEndDayOfWeek = p.RepeatEndDayOfWeek
	}
	var repeatBeginHour *int32
	if p.RepeatBeginHour != nil {
		repeatBeginHour = p.RepeatBeginHour
	}
	var repeatEndHour *int32
	if p.RepeatEndHour != nil {
		repeatEndHour = p.RepeatEndHour
	}
	return map[string]interface{}{
		"eventId":               eventId,
		"name":                  name,
		"metadata":              metadata,
		"scheduleType":          scheduleType,
		"absoluteBegin":         absoluteBegin,
		"absoluteEnd":           absoluteEnd,
		"relativeTriggerName":   relativeTriggerName,
		"repeatSetting":         repeatSetting,
		"repeatType":            repeatType,
		"repeatBeginDayOfMonth": repeatBeginDayOfMonth,
		"repeatEndDayOfMonth":   repeatEndDayOfMonth,
		"repeatBeginDayOfWeek":  repeatBeginDayOfWeek,
		"repeatEndDayOfWeek":    repeatEndDayOfWeek,
		"repeatBeginHour":       repeatBeginHour,
		"repeatEndHour":         repeatEndHour,
	}
}

func (p Event) Pointer() *Event {
	return &p
}

func CastEvents(data []interface{}) []Event {
	v := make([]Event, 0)
	for _, d := range data {
		v = append(v, NewEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventsFromDict(data []Event) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RepeatSetting struct {
	RepeatType      *string `json:"repeatType"`
	BeginDayOfMonth *int32  `json:"beginDayOfMonth"`
	EndDayOfMonth   *int32  `json:"endDayOfMonth"`
	BeginDayOfWeek  *string `json:"beginDayOfWeek"`
	EndDayOfWeek    *string `json:"endDayOfWeek"`
	BeginHour       *int32  `json:"beginHour"`
	EndHour         *int32  `json:"endHour"`
}

func (p *RepeatSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RepeatSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RepeatSetting{}
	} else {
		*p = RepeatSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["repeatType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepeatType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepeatType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepeatType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepeatType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepeatType)
				}
			}
		}
		if v, ok := d["beginDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BeginDayOfMonth)
		}
		if v, ok := d["endDayOfMonth"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EndDayOfMonth)
		}
		if v, ok := d["beginDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BeginDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BeginDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BeginDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BeginDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BeginDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BeginDayOfWeek)
				}
			}
		}
		if v, ok := d["endDayOfWeek"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EndDayOfWeek = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EndDayOfWeek = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EndDayOfWeek = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EndDayOfWeek = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EndDayOfWeek = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EndDayOfWeek)
				}
			}
		}
		if v, ok := d["beginHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BeginHour)
		}
		if v, ok := d["endHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EndHour)
		}
	}
	return nil
}

func NewRepeatSettingFromJson(data string) RepeatSetting {
	req := RepeatSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRepeatSettingFromDict(data map[string]interface{}) RepeatSetting {
	return RepeatSetting{
		RepeatType:      core.CastString(data["repeatType"]),
		BeginDayOfMonth: core.CastInt32(data["beginDayOfMonth"]),
		EndDayOfMonth:   core.CastInt32(data["endDayOfMonth"]),
		BeginDayOfWeek:  core.CastString(data["beginDayOfWeek"]),
		EndDayOfWeek:    core.CastString(data["endDayOfWeek"]),
		BeginHour:       core.CastInt32(data["beginHour"]),
		EndHour:         core.CastInt32(data["endHour"]),
	}
}

func (p RepeatSetting) ToDict() map[string]interface{} {

	var repeatType *string
	if p.RepeatType != nil {
		repeatType = p.RepeatType
	}
	var beginDayOfMonth *int32
	if p.BeginDayOfMonth != nil {
		beginDayOfMonth = p.BeginDayOfMonth
	}
	var endDayOfMonth *int32
	if p.EndDayOfMonth != nil {
		endDayOfMonth = p.EndDayOfMonth
	}
	var beginDayOfWeek *string
	if p.BeginDayOfWeek != nil {
		beginDayOfWeek = p.BeginDayOfWeek
	}
	var endDayOfWeek *string
	if p.EndDayOfWeek != nil {
		endDayOfWeek = p.EndDayOfWeek
	}
	var beginHour *int32
	if p.BeginHour != nil {
		beginHour = p.BeginHour
	}
	var endHour *int32
	if p.EndHour != nil {
		endHour = p.EndHour
	}
	return map[string]interface{}{
		"repeatType":      repeatType,
		"beginDayOfMonth": beginDayOfMonth,
		"endDayOfMonth":   endDayOfMonth,
		"beginDayOfWeek":  beginDayOfWeek,
		"endDayOfWeek":    endDayOfWeek,
		"beginHour":       beginHour,
		"endHour":         endHour,
	}
}

func (p RepeatSetting) Pointer() *RepeatSetting {
	return &p
}

func CastRepeatSettings(data []interface{}) []RepeatSetting {
	v := make([]RepeatSetting, 0)
	for _, d := range data {
		v = append(v, NewRepeatSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRepeatSettingsFromDict(data []RepeatSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RepeatSchedule struct {
	RepeatCount          *int32 `json:"repeatCount"`
	CurrentRepeatStartAt *int64 `json:"currentRepeatStartAt"`
	CurrentRepeatEndAt   *int64 `json:"currentRepeatEndAt"`
	LastRepeatEndAt      *int64 `json:"lastRepeatEndAt"`
	NextRepeatStartAt    *int64 `json:"nextRepeatStartAt"`
}

func (p *RepeatSchedule) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RepeatSchedule{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RepeatSchedule{}
	} else {
		*p = RepeatSchedule{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["repeatCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RepeatCount)
		}
		if v, ok := d["currentRepeatStartAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentRepeatStartAt)
		}
		if v, ok := d["currentRepeatEndAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentRepeatEndAt)
		}
		if v, ok := d["lastRepeatEndAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LastRepeatEndAt)
		}
		if v, ok := d["nextRepeatStartAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NextRepeatStartAt)
		}
	}
	return nil
}

func NewRepeatScheduleFromJson(data string) RepeatSchedule {
	req := RepeatSchedule{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRepeatScheduleFromDict(data map[string]interface{}) RepeatSchedule {
	return RepeatSchedule{
		RepeatCount:          core.CastInt32(data["repeatCount"]),
		CurrentRepeatStartAt: core.CastInt64(data["currentRepeatStartAt"]),
		CurrentRepeatEndAt:   core.CastInt64(data["currentRepeatEndAt"]),
		LastRepeatEndAt:      core.CastInt64(data["lastRepeatEndAt"]),
		NextRepeatStartAt:    core.CastInt64(data["nextRepeatStartAt"]),
	}
}

func (p RepeatSchedule) ToDict() map[string]interface{} {

	var repeatCount *int32
	if p.RepeatCount != nil {
		repeatCount = p.RepeatCount
	}
	var currentRepeatStartAt *int64
	if p.CurrentRepeatStartAt != nil {
		currentRepeatStartAt = p.CurrentRepeatStartAt
	}
	var currentRepeatEndAt *int64
	if p.CurrentRepeatEndAt != nil {
		currentRepeatEndAt = p.CurrentRepeatEndAt
	}
	var lastRepeatEndAt *int64
	if p.LastRepeatEndAt != nil {
		lastRepeatEndAt = p.LastRepeatEndAt
	}
	var nextRepeatStartAt *int64
	if p.NextRepeatStartAt != nil {
		nextRepeatStartAt = p.NextRepeatStartAt
	}
	return map[string]interface{}{
		"repeatCount":          repeatCount,
		"currentRepeatStartAt": currentRepeatStartAt,
		"currentRepeatEndAt":   currentRepeatEndAt,
		"lastRepeatEndAt":      lastRepeatEndAt,
		"nextRepeatStartAt":    nextRepeatStartAt,
	}
}

func (p RepeatSchedule) Pointer() *RepeatSchedule {
	return &p
}

func CastRepeatSchedules(data []interface{}) []RepeatSchedule {
	v := make([]RepeatSchedule, 0)
	for _, d := range data {
		v = append(v, NewRepeatScheduleFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRepeatSchedulesFromDict(data []RepeatSchedule) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentEventMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentEventMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentEventMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentEventMaster{}
	} else {
		*p = CurrentEventMaster{}
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

func NewCurrentEventMasterFromJson(data string) CurrentEventMaster {
	req := CurrentEventMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentEventMasterFromDict(data map[string]interface{}) CurrentEventMaster {
	return CurrentEventMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentEventMaster) ToDict() map[string]interface{} {

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

func (p CurrentEventMaster) Pointer() *CurrentEventMaster {
	return &p
}

func CastCurrentEventMasters(data []interface{}) []CurrentEventMaster {
	v := make([]CurrentEventMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentEventMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentEventMastersFromDict(data []CurrentEventMaster) []interface{} {
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
