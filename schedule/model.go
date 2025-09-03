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
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	LogSetting         *LogSetting         `json:"logSetting"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
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
		EventId: func() *string {
			v, ok := data["eventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventId"])
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
		ScheduleType: func() *string {
			v, ok := data["scheduleType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleType"])
		}(),
		AbsoluteBegin: func() *int64 {
			v, ok := data["absoluteBegin"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["absoluteBegin"])
		}(),
		AbsoluteEnd: func() *int64 {
			v, ok := data["absoluteEnd"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["absoluteEnd"])
		}(),
		RelativeTriggerName: func() *string {
			v, ok := data["relativeTriggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["relativeTriggerName"])
		}(),
		RepeatSetting: func() *RepeatSetting {
			v, ok := data["repeatSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewRepeatSettingFromDict(core.CastMap(data["repeatSetting"])).Pointer()
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
		RepeatType: func() *string {
			v, ok := data["repeatType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatType"])
		}(),
		RepeatBeginDayOfMonth: func() *int32 {
			v, ok := data["repeatBeginDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatBeginDayOfMonth"])
		}(),
		RepeatEndDayOfMonth: func() *int32 {
			v, ok := data["repeatEndDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatEndDayOfMonth"])
		}(),
		RepeatBeginDayOfWeek: func() *string {
			v, ok := data["repeatBeginDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatBeginDayOfWeek"])
		}(),
		RepeatEndDayOfWeek: func() *string {
			v, ok := data["repeatEndDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatEndDayOfWeek"])
		}(),
		RepeatBeginHour: func() *int32 {
			v, ok := data["repeatBeginHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatBeginHour"])
		}(),
		RepeatEndHour: func() *int32 {
			v, ok := data["repeatEndHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatEndHour"])
		}(),
	}
}

func (p EventMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EventId != nil {
		m["eventId"] = p.EventId
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
	if p.ScheduleType != nil {
		m["scheduleType"] = p.ScheduleType
	}
	if p.AbsoluteBegin != nil {
		m["absoluteBegin"] = p.AbsoluteBegin
	}
	if p.AbsoluteEnd != nil {
		m["absoluteEnd"] = p.AbsoluteEnd
	}
	if p.RelativeTriggerName != nil {
		m["relativeTriggerName"] = p.RelativeTriggerName
	}
	if p.RepeatSetting != nil {
		m["repeatSetting"] = func() map[string]interface{} {
			if p.RepeatSetting == nil {
				return nil
			}
			return p.RepeatSetting.ToDict()
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
	if p.RepeatType != nil {
		m["repeatType"] = p.RepeatType
	}
	if p.RepeatBeginDayOfMonth != nil {
		m["repeatBeginDayOfMonth"] = p.RepeatBeginDayOfMonth
	}
	if p.RepeatEndDayOfMonth != nil {
		m["repeatEndDayOfMonth"] = p.RepeatEndDayOfMonth
	}
	if p.RepeatBeginDayOfWeek != nil {
		m["repeatBeginDayOfWeek"] = p.RepeatBeginDayOfWeek
	}
	if p.RepeatEndDayOfWeek != nil {
		m["repeatEndDayOfWeek"] = p.RepeatEndDayOfWeek
	}
	if p.RepeatBeginHour != nil {
		m["repeatBeginHour"] = p.RepeatBeginHour
	}
	if p.RepeatEndHour != nil {
		m["repeatEndHour"] = p.RepeatEndHour
	}
	return m
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
	TriggerId   *string `json:"triggerId"`
	Name        *string `json:"name"`
	UserId      *string `json:"userId"`
	TriggeredAt *int64  `json:"triggeredAt"`
	ExpiresAt   *int64  `json:"expiresAt"`
	CreatedAt   *int64  `json:"createdAt"`
	Revision    *int64  `json:"revision"`
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
		if v, ok := d["triggeredAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TriggeredAt)
		}
		if v, ok := d["expiresAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ExpiresAt)
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

func NewTriggerFromJson(data string) Trigger {
	req := Trigger{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTriggerFromDict(data map[string]interface{}) Trigger {
	return Trigger{
		TriggerId: func() *string {
			v, ok := data["triggerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerId"])
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
		TriggeredAt: func() *int64 {
			v, ok := data["triggeredAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["triggeredAt"])
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
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Trigger) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TriggerId != nil {
		m["triggerId"] = p.TriggerId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TriggeredAt != nil {
		m["triggeredAt"] = p.TriggeredAt
	}
	if p.ExpiresAt != nil {
		m["expiresAt"] = p.ExpiresAt
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		EventId: func() *string {
			v, ok := data["eventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventId"])
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
		ScheduleType: func() *string {
			v, ok := data["scheduleType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleType"])
		}(),
		AbsoluteBegin: func() *int64 {
			v, ok := data["absoluteBegin"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["absoluteBegin"])
		}(),
		AbsoluteEnd: func() *int64 {
			v, ok := data["absoluteEnd"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["absoluteEnd"])
		}(),
		RelativeTriggerName: func() *string {
			v, ok := data["relativeTriggerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["relativeTriggerName"])
		}(),
		RepeatSetting: func() *RepeatSetting {
			v, ok := data["repeatSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewRepeatSettingFromDict(core.CastMap(data["repeatSetting"])).Pointer()
		}(),
		RepeatType: func() *string {
			v, ok := data["repeatType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatType"])
		}(),
		RepeatBeginDayOfMonth: func() *int32 {
			v, ok := data["repeatBeginDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatBeginDayOfMonth"])
		}(),
		RepeatEndDayOfMonth: func() *int32 {
			v, ok := data["repeatEndDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatEndDayOfMonth"])
		}(),
		RepeatBeginDayOfWeek: func() *string {
			v, ok := data["repeatBeginDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatBeginDayOfWeek"])
		}(),
		RepeatEndDayOfWeek: func() *string {
			v, ok := data["repeatEndDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatEndDayOfWeek"])
		}(),
		RepeatBeginHour: func() *int32 {
			v, ok := data["repeatBeginHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatBeginHour"])
		}(),
		RepeatEndHour: func() *int32 {
			v, ok := data["repeatEndHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatEndHour"])
		}(),
	}
}

func (p Event) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EventId != nil {
		m["eventId"] = p.EventId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ScheduleType != nil {
		m["scheduleType"] = p.ScheduleType
	}
	if p.AbsoluteBegin != nil {
		m["absoluteBegin"] = p.AbsoluteBegin
	}
	if p.AbsoluteEnd != nil {
		m["absoluteEnd"] = p.AbsoluteEnd
	}
	if p.RelativeTriggerName != nil {
		m["relativeTriggerName"] = p.RelativeTriggerName
	}
	if p.RepeatSetting != nil {
		m["repeatSetting"] = func() map[string]interface{} {
			if p.RepeatSetting == nil {
				return nil
			}
			return p.RepeatSetting.ToDict()
		}()
	}
	if p.RepeatType != nil {
		m["repeatType"] = p.RepeatType
	}
	if p.RepeatBeginDayOfMonth != nil {
		m["repeatBeginDayOfMonth"] = p.RepeatBeginDayOfMonth
	}
	if p.RepeatEndDayOfMonth != nil {
		m["repeatEndDayOfMonth"] = p.RepeatEndDayOfMonth
	}
	if p.RepeatBeginDayOfWeek != nil {
		m["repeatBeginDayOfWeek"] = p.RepeatBeginDayOfWeek
	}
	if p.RepeatEndDayOfWeek != nil {
		m["repeatEndDayOfWeek"] = p.RepeatEndDayOfWeek
	}
	if p.RepeatBeginHour != nil {
		m["repeatBeginHour"] = p.RepeatBeginHour
	}
	if p.RepeatEndHour != nil {
		m["repeatEndHour"] = p.RepeatEndHour
	}
	return m
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
	AnchorTimestamp *int64  `json:"anchorTimestamp"`
	ActiveDays      *int32  `json:"activeDays"`
	InactiveDays    *int32  `json:"inactiveDays"`
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
		if v, ok := d["anchorTimestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AnchorTimestamp)
		}
		if v, ok := d["activeDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ActiveDays)
		}
		if v, ok := d["inactiveDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InactiveDays)
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
		RepeatType: func() *string {
			v, ok := data["repeatType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repeatType"])
		}(),
		BeginDayOfMonth: func() *int32 {
			v, ok := data["beginDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["beginDayOfMonth"])
		}(),
		EndDayOfMonth: func() *int32 {
			v, ok := data["endDayOfMonth"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["endDayOfMonth"])
		}(),
		BeginDayOfWeek: func() *string {
			v, ok := data["beginDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["beginDayOfWeek"])
		}(),
		EndDayOfWeek: func() *string {
			v, ok := data["endDayOfWeek"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["endDayOfWeek"])
		}(),
		BeginHour: func() *int32 {
			v, ok := data["beginHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["beginHour"])
		}(),
		EndHour: func() *int32 {
			v, ok := data["endHour"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["endHour"])
		}(),
		AnchorTimestamp: func() *int64 {
			v, ok := data["anchorTimestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["anchorTimestamp"])
		}(),
		ActiveDays: func() *int32 {
			v, ok := data["activeDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["activeDays"])
		}(),
		InactiveDays: func() *int32 {
			v, ok := data["inactiveDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["inactiveDays"])
		}(),
	}
}

func (p RepeatSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RepeatType != nil {
		m["repeatType"] = p.RepeatType
	}
	if p.BeginDayOfMonth != nil {
		m["beginDayOfMonth"] = p.BeginDayOfMonth
	}
	if p.EndDayOfMonth != nil {
		m["endDayOfMonth"] = p.EndDayOfMonth
	}
	if p.BeginDayOfWeek != nil {
		m["beginDayOfWeek"] = p.BeginDayOfWeek
	}
	if p.EndDayOfWeek != nil {
		m["endDayOfWeek"] = p.EndDayOfWeek
	}
	if p.BeginHour != nil {
		m["beginHour"] = p.BeginHour
	}
	if p.EndHour != nil {
		m["endHour"] = p.EndHour
	}
	if p.AnchorTimestamp != nil {
		m["anchorTimestamp"] = p.AnchorTimestamp
	}
	if p.ActiveDays != nil {
		m["activeDays"] = p.ActiveDays
	}
	if p.InactiveDays != nil {
		m["inactiveDays"] = p.InactiveDays
	}
	return m
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
		RepeatCount: func() *int32 {
			v, ok := data["repeatCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["repeatCount"])
		}(),
		CurrentRepeatStartAt: func() *int64 {
			v, ok := data["currentRepeatStartAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["currentRepeatStartAt"])
		}(),
		CurrentRepeatEndAt: func() *int64 {
			v, ok := data["currentRepeatEndAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["currentRepeatEndAt"])
		}(),
		LastRepeatEndAt: func() *int64 {
			v, ok := data["lastRepeatEndAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["lastRepeatEndAt"])
		}(),
		NextRepeatStartAt: func() *int64 {
			v, ok := data["nextRepeatStartAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["nextRepeatStartAt"])
		}(),
	}
}

func (p RepeatSchedule) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RepeatCount != nil {
		m["repeatCount"] = p.RepeatCount
	}
	if p.CurrentRepeatStartAt != nil {
		m["currentRepeatStartAt"] = p.CurrentRepeatStartAt
	}
	if p.CurrentRepeatEndAt != nil {
		m["currentRepeatEndAt"] = p.CurrentRepeatEndAt
	}
	if p.LastRepeatEndAt != nil {
		m["lastRepeatEndAt"] = p.LastRepeatEndAt
	}
	if p.NextRepeatStartAt != nil {
		m["nextRepeatStartAt"] = p.NextRepeatStartAt
	}
	return m
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

func (p CurrentEventMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
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
