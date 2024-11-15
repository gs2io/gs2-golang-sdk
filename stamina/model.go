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

type Namespace struct {
	NamespaceId           *string     `json:"namespaceId"`
	Name                  *string     `json:"name"`
	Description           *string     `json:"description"`
	OverflowTriggerScript *string     `json:"overflowTriggerScript"`
	LogSetting            *LogSetting `json:"logSetting"`
	CreatedAt             *int64      `json:"createdAt"`
	UpdatedAt             *int64      `json:"updatedAt"`
	Revision              *int64      `json:"revision"`
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
		if v, ok := d["overflowTriggerScript"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.OverflowTriggerScript)
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
		OverflowTriggerScript: func() *string {
			v, ok := data["overflowTriggerScript"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["overflowTriggerScript"])
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
	if p.OverflowTriggerScript != nil {
		m["overflowTriggerScript"] = p.OverflowTriggerScript
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

type StaminaModelMaster struct {
	StaminaModelId           *string `json:"staminaModelId"`
	Name                     *string `json:"name"`
	Metadata                 *string `json:"metadata"`
	Description              *string `json:"description"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
	CreatedAt                *int64  `json:"createdAt"`
	UpdatedAt                *int64  `json:"updatedAt"`
	Revision                 *int64  `json:"revision"`
}

func (p *StaminaModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StaminaModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = StaminaModelMaster{}
	} else {
		*p = StaminaModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["staminaModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StaminaModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StaminaModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StaminaModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StaminaModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StaminaModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StaminaModelId)
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
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverIntervalMinutes)
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverValue)
		}
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["isOverflow"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IsOverflow)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["maxStaminaTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.MaxStaminaTableName)
				}
			}
		}
		if v, ok := d["recoverIntervalTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.RecoverIntervalTableName)
				}
			}
		}
		if v, ok := d["recoverValueTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.RecoverValueTableName)
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

func NewStaminaModelMasterFromJson(data string) StaminaModelMaster {
	req := StaminaModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStaminaModelMasterFromDict(data map[string]interface{}) StaminaModelMaster {
	return StaminaModelMaster{
		StaminaModelId: func() *string {
			v, ok := data["staminaModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["staminaModelId"])
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
		RecoverIntervalMinutes: func() *int32 {
			v, ok := data["recoverIntervalMinutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverIntervalMinutes"])
		}(),
		RecoverValue: func() *int32 {
			v, ok := data["recoverValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverValue"])
		}(),
		InitialCapacity: func() *int32 {
			v, ok := data["initialCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialCapacity"])
		}(),
		IsOverflow: func() *bool {
			v, ok := data["isOverflow"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isOverflow"])
		}(),
		MaxCapacity: func() *int32 {
			v, ok := data["maxCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxCapacity"])
		}(),
		MaxStaminaTableName: func() *string {
			v, ok := data["maxStaminaTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["maxStaminaTableName"])
		}(),
		RecoverIntervalTableName: func() *string {
			v, ok := data["recoverIntervalTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["recoverIntervalTableName"])
		}(),
		RecoverValueTableName: func() *string {
			v, ok := data["recoverValueTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["recoverValueTableName"])
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

func (p StaminaModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StaminaModelId != nil {
		m["staminaModelId"] = p.StaminaModelId
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
	if p.RecoverIntervalMinutes != nil {
		m["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
	}
	if p.RecoverValue != nil {
		m["recoverValue"] = p.RecoverValue
	}
	if p.InitialCapacity != nil {
		m["initialCapacity"] = p.InitialCapacity
	}
	if p.IsOverflow != nil {
		m["isOverflow"] = p.IsOverflow
	}
	if p.MaxCapacity != nil {
		m["maxCapacity"] = p.MaxCapacity
	}
	if p.MaxStaminaTableName != nil {
		m["maxStaminaTableName"] = p.MaxStaminaTableName
	}
	if p.RecoverIntervalTableName != nil {
		m["recoverIntervalTableName"] = p.RecoverIntervalTableName
	}
	if p.RecoverValueTableName != nil {
		m["recoverValueTableName"] = p.RecoverValueTableName
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

func (p StaminaModelMaster) Pointer() *StaminaModelMaster {
	return &p
}

func CastStaminaModelMasters(data []interface{}) []StaminaModelMaster {
	v := make([]StaminaModelMaster, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelMastersFromDict(data []StaminaModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MaxStaminaTableMaster struct {
	MaxStaminaTableId *string  `json:"maxStaminaTableId"`
	Name              *string  `json:"name"`
	Metadata          *string  `json:"metadata"`
	Description       *string  `json:"description"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
	CreatedAt         *int64   `json:"createdAt"`
	UpdatedAt         *int64   `json:"updatedAt"`
	Revision          *int64   `json:"revision"`
}

func (p *MaxStaminaTableMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MaxStaminaTableMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MaxStaminaTableMaster{}
	} else {
		*p = MaxStaminaTableMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["maxStaminaTableId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MaxStaminaTableId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MaxStaminaTableId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MaxStaminaTableId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MaxStaminaTableId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MaxStaminaTableId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MaxStaminaTableId)
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

func NewMaxStaminaTableMasterFromJson(data string) MaxStaminaTableMaster {
	req := MaxStaminaTableMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMaxStaminaTableMasterFromDict(data map[string]interface{}) MaxStaminaTableMaster {
	return MaxStaminaTableMaster{
		MaxStaminaTableId: func() *string {
			v, ok := data["maxStaminaTableId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["maxStaminaTableId"])
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
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

func (p MaxStaminaTableMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.MaxStaminaTableId != nil {
		m["maxStaminaTableId"] = p.MaxStaminaTableId
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
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
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

func (p MaxStaminaTableMaster) Pointer() *MaxStaminaTableMaster {
	return &p
}

func CastMaxStaminaTableMasters(data []interface{}) []MaxStaminaTableMaster {
	v := make([]MaxStaminaTableMaster, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTableMastersFromDict(data []MaxStaminaTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverIntervalTableMaster struct {
	RecoverIntervalTableId *string  `json:"recoverIntervalTableId"`
	Name                   *string  `json:"name"`
	Metadata               *string  `json:"metadata"`
	Description            *string  `json:"description"`
	ExperienceModelId      *string  `json:"experienceModelId"`
	Values                 []*int32 `json:"values"`
	CreatedAt              *int64   `json:"createdAt"`
	UpdatedAt              *int64   `json:"updatedAt"`
	Revision               *int64   `json:"revision"`
}

func (p *RecoverIntervalTableMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverIntervalTableMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverIntervalTableMaster{}
	} else {
		*p = RecoverIntervalTableMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["recoverIntervalTableId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RecoverIntervalTableId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RecoverIntervalTableId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RecoverIntervalTableId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RecoverIntervalTableId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RecoverIntervalTableId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RecoverIntervalTableId)
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

func NewRecoverIntervalTableMasterFromJson(data string) RecoverIntervalTableMaster {
	req := RecoverIntervalTableMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRecoverIntervalTableMasterFromDict(data map[string]interface{}) RecoverIntervalTableMaster {
	return RecoverIntervalTableMaster{
		RecoverIntervalTableId: func() *string {
			v, ok := data["recoverIntervalTableId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["recoverIntervalTableId"])
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
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

func (p RecoverIntervalTableMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RecoverIntervalTableId != nil {
		m["recoverIntervalTableId"] = p.RecoverIntervalTableId
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
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
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

func (p RecoverIntervalTableMaster) Pointer() *RecoverIntervalTableMaster {
	return &p
}

func CastRecoverIntervalTableMasters(data []interface{}) []RecoverIntervalTableMaster {
	v := make([]RecoverIntervalTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTableMastersFromDict(data []RecoverIntervalTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverValueTableMaster struct {
	RecoverValueTableId *string  `json:"recoverValueTableId"`
	Name                *string  `json:"name"`
	Metadata            *string  `json:"metadata"`
	Description         *string  `json:"description"`
	ExperienceModelId   *string  `json:"experienceModelId"`
	Values              []*int32 `json:"values"`
	CreatedAt           *int64   `json:"createdAt"`
	UpdatedAt           *int64   `json:"updatedAt"`
	Revision            *int64   `json:"revision"`
}

func (p *RecoverValueTableMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverValueTableMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverValueTableMaster{}
	} else {
		*p = RecoverValueTableMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["recoverValueTableId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RecoverValueTableId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RecoverValueTableId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RecoverValueTableId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RecoverValueTableId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RecoverValueTableId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RecoverValueTableId)
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

func NewRecoverValueTableMasterFromJson(data string) RecoverValueTableMaster {
	req := RecoverValueTableMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRecoverValueTableMasterFromDict(data map[string]interface{}) RecoverValueTableMaster {
	return RecoverValueTableMaster{
		RecoverValueTableId: func() *string {
			v, ok := data["recoverValueTableId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["recoverValueTableId"])
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
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

func (p RecoverValueTableMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RecoverValueTableId != nil {
		m["recoverValueTableId"] = p.RecoverValueTableId
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
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
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

func (p RecoverValueTableMaster) Pointer() *RecoverValueTableMaster {
	return &p
}

func CastRecoverValueTableMasters(data []interface{}) []RecoverValueTableMaster {
	v := make([]RecoverValueTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTableMastersFromDict(data []RecoverValueTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentStaminaMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentStaminaMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentStaminaMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentStaminaMaster{}
	} else {
		*p = CurrentStaminaMaster{}
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

func NewCurrentStaminaMasterFromJson(data string) CurrentStaminaMaster {
	req := CurrentStaminaMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentStaminaMasterFromDict(data map[string]interface{}) CurrentStaminaMaster {
	return CurrentStaminaMaster{
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

func (p CurrentStaminaMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentStaminaMaster) Pointer() *CurrentStaminaMaster {
	return &p
}

func CastCurrentStaminaMasters(data []interface{}) []CurrentStaminaMaster {
	v := make([]CurrentStaminaMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentStaminaMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentStaminaMastersFromDict(data []CurrentStaminaMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaModel struct {
	StaminaModelId         *string               `json:"staminaModelId"`
	Name                   *string               `json:"name"`
	Metadata               *string               `json:"metadata"`
	RecoverIntervalMinutes *int32                `json:"recoverIntervalMinutes"`
	RecoverValue           *int32                `json:"recoverValue"`
	InitialCapacity        *int32                `json:"initialCapacity"`
	IsOverflow             *bool                 `json:"isOverflow"`
	MaxCapacity            *int32                `json:"maxCapacity"`
	MaxStaminaTable        *MaxStaminaTable      `json:"maxStaminaTable"`
	RecoverIntervalTable   *RecoverIntervalTable `json:"recoverIntervalTable"`
	RecoverValueTable      *RecoverValueTable    `json:"recoverValueTable"`
}

func (p *StaminaModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StaminaModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = StaminaModel{}
	} else {
		*p = StaminaModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["staminaModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StaminaModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StaminaModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StaminaModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StaminaModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StaminaModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StaminaModelId)
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
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverIntervalMinutes)
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverValue)
		}
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["isOverflow"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.IsOverflow)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["maxStaminaTable"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxStaminaTable)
		}
		if v, ok := d["recoverIntervalTable"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverIntervalTable)
		}
		if v, ok := d["recoverValueTable"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverValueTable)
		}
	}
	return nil
}

func NewStaminaModelFromJson(data string) StaminaModel {
	req := StaminaModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStaminaModelFromDict(data map[string]interface{}) StaminaModel {
	return StaminaModel{
		StaminaModelId: func() *string {
			v, ok := data["staminaModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["staminaModelId"])
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
		RecoverIntervalMinutes: func() *int32 {
			v, ok := data["recoverIntervalMinutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverIntervalMinutes"])
		}(),
		RecoverValue: func() *int32 {
			v, ok := data["recoverValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverValue"])
		}(),
		InitialCapacity: func() *int32 {
			v, ok := data["initialCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialCapacity"])
		}(),
		IsOverflow: func() *bool {
			v, ok := data["isOverflow"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["isOverflow"])
		}(),
		MaxCapacity: func() *int32 {
			v, ok := data["maxCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxCapacity"])
		}(),
		MaxStaminaTable: func() *MaxStaminaTable {
			v, ok := data["maxStaminaTable"]
			if !ok || v == nil {
				return nil
			}
			return NewMaxStaminaTableFromDict(core.CastMap(data["maxStaminaTable"])).Pointer()
		}(),
		RecoverIntervalTable: func() *RecoverIntervalTable {
			v, ok := data["recoverIntervalTable"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverIntervalTableFromDict(core.CastMap(data["recoverIntervalTable"])).Pointer()
		}(),
		RecoverValueTable: func() *RecoverValueTable {
			v, ok := data["recoverValueTable"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverValueTableFromDict(core.CastMap(data["recoverValueTable"])).Pointer()
		}(),
	}
}

func (p StaminaModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StaminaModelId != nil {
		m["staminaModelId"] = p.StaminaModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.RecoverIntervalMinutes != nil {
		m["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
	}
	if p.RecoverValue != nil {
		m["recoverValue"] = p.RecoverValue
	}
	if p.InitialCapacity != nil {
		m["initialCapacity"] = p.InitialCapacity
	}
	if p.IsOverflow != nil {
		m["isOverflow"] = p.IsOverflow
	}
	if p.MaxCapacity != nil {
		m["maxCapacity"] = p.MaxCapacity
	}
	if p.MaxStaminaTable != nil {
		m["maxStaminaTable"] = func() map[string]interface{} {
			if p.MaxStaminaTable == nil {
				return nil
			}
			return p.MaxStaminaTable.ToDict()
		}()
	}
	if p.RecoverIntervalTable != nil {
		m["recoverIntervalTable"] = func() map[string]interface{} {
			if p.RecoverIntervalTable == nil {
				return nil
			}
			return p.RecoverIntervalTable.ToDict()
		}()
	}
	if p.RecoverValueTable != nil {
		m["recoverValueTable"] = func() map[string]interface{} {
			if p.RecoverValueTable == nil {
				return nil
			}
			return p.RecoverValueTable.ToDict()
		}()
	}
	return m
}

func (p StaminaModel) Pointer() *StaminaModel {
	return &p
}

func CastStaminaModels(data []interface{}) []StaminaModel {
	v := make([]StaminaModel, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelsFromDict(data []StaminaModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MaxStaminaTable struct {
	Name              *string  `json:"name"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *MaxStaminaTable) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MaxStaminaTable{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = MaxStaminaTable{}
	} else {
		*p = MaxStaminaTable{}
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewMaxStaminaTableFromJson(data string) MaxStaminaTable {
	req := MaxStaminaTable{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMaxStaminaTableFromDict(data map[string]interface{}) MaxStaminaTable {
	return MaxStaminaTable{
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
	}
}

func (p MaxStaminaTable) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
		)
	}
	return m
}

func (p MaxStaminaTable) Pointer() *MaxStaminaTable {
	return &p
}

func CastMaxStaminaTables(data []interface{}) []MaxStaminaTable {
	v := make([]MaxStaminaTable, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTablesFromDict(data []MaxStaminaTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverIntervalTable struct {
	Name              *string  `json:"name"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *RecoverIntervalTable) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverIntervalTable{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverIntervalTable{}
	} else {
		*p = RecoverIntervalTable{}
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewRecoverIntervalTableFromJson(data string) RecoverIntervalTable {
	req := RecoverIntervalTable{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRecoverIntervalTableFromDict(data map[string]interface{}) RecoverIntervalTable {
	return RecoverIntervalTable{
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
	}
}

func (p RecoverIntervalTable) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
		)
	}
	return m
}

func (p RecoverIntervalTable) Pointer() *RecoverIntervalTable {
	return &p
}

func CastRecoverIntervalTables(data []interface{}) []RecoverIntervalTable {
	v := make([]RecoverIntervalTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTablesFromDict(data []RecoverIntervalTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverValueTable struct {
	Name              *string  `json:"name"`
	Metadata          *string  `json:"metadata"`
	ExperienceModelId *string  `json:"experienceModelId"`
	Values            []*int32 `json:"values"`
}

func (p *RecoverValueTable) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RecoverValueTable{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RecoverValueTable{}
	} else {
		*p = RecoverValueTable{}
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
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewRecoverValueTableFromJson(data string) RecoverValueTable {
	req := RecoverValueTable{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRecoverValueTableFromDict(data map[string]interface{}) RecoverValueTable {
	return RecoverValueTable{
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
		ExperienceModelId: func() *string {
			v, ok := data["experienceModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceModelId"])
		}(),
		Values: func() []*int32 {
			v, ok := data["values"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32s(core.CastArray(v))
		}(),
	}
}

func (p RecoverValueTable) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ExperienceModelId != nil {
		m["experienceModelId"] = p.ExperienceModelId
	}
	if p.Values != nil {
		m["values"] = core.CastInt32sFromDict(
			p.Values,
		)
	}
	return m
}

func (p RecoverValueTable) Pointer() *RecoverValueTable {
	return &p
}

func CastRecoverValueTables(data []interface{}) []RecoverValueTable {
	v := make([]RecoverValueTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTablesFromDict(data []RecoverValueTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Stamina struct {
	StaminaId              *string `json:"staminaId"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	Value                  *int32  `json:"value"`
	MaxValue               *int32  `json:"maxValue"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
	RecoverValue           *int32  `json:"recoverValue"`
	OverflowValue          *int32  `json:"overflowValue"`
	NextRecoverAt          *int64  `json:"nextRecoverAt"`
	LastRecoveredAt        *int64  `json:"lastRecoveredAt"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
	Revision               *int64  `json:"revision"`
}

func (p *Stamina) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Stamina{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Stamina{}
	} else {
		*p = Stamina{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["staminaId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StaminaId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StaminaId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StaminaId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StaminaId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StaminaId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StaminaId)
				}
			}
		}
		if v, ok := d["staminaName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.StaminaName)
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
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
		if v, ok := d["maxValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxValue)
		}
		if v, ok := d["recoverIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverIntervalMinutes)
		}
		if v, ok := d["recoverValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RecoverValue)
		}
		if v, ok := d["overflowValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OverflowValue)
		}
		if v, ok := d["nextRecoverAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NextRecoverAt)
		}
		if v, ok := d["lastRecoveredAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LastRecoveredAt)
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

func NewStaminaFromJson(data string) Stamina {
	req := Stamina{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStaminaFromDict(data map[string]interface{}) Stamina {
	return Stamina{
		StaminaId: func() *string {
			v, ok := data["staminaId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["staminaId"])
		}(),
		StaminaName: func() *string {
			v, ok := data["staminaName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["staminaName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Value: func() *int32 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["value"])
		}(),
		MaxValue: func() *int32 {
			v, ok := data["maxValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxValue"])
		}(),
		RecoverIntervalMinutes: func() *int32 {
			v, ok := data["recoverIntervalMinutes"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverIntervalMinutes"])
		}(),
		RecoverValue: func() *int32 {
			v, ok := data["recoverValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["recoverValue"])
		}(),
		OverflowValue: func() *int32 {
			v, ok := data["overflowValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["overflowValue"])
		}(),
		NextRecoverAt: func() *int64 {
			v, ok := data["nextRecoverAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["nextRecoverAt"])
		}(),
		LastRecoveredAt: func() *int64 {
			v, ok := data["lastRecoveredAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["lastRecoveredAt"])
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

func (p Stamina) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StaminaId != nil {
		m["staminaId"] = p.StaminaId
	}
	if p.StaminaName != nil {
		m["staminaName"] = p.StaminaName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	if p.MaxValue != nil {
		m["maxValue"] = p.MaxValue
	}
	if p.RecoverIntervalMinutes != nil {
		m["recoverIntervalMinutes"] = p.RecoverIntervalMinutes
	}
	if p.RecoverValue != nil {
		m["recoverValue"] = p.RecoverValue
	}
	if p.OverflowValue != nil {
		m["overflowValue"] = p.OverflowValue
	}
	if p.NextRecoverAt != nil {
		m["nextRecoverAt"] = p.NextRecoverAt
	}
	if p.LastRecoveredAt != nil {
		m["lastRecoveredAt"] = p.LastRecoveredAt
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

func (p Stamina) Pointer() *Stamina {
	return &p
}

func CastStaminas(data []interface{}) []Stamina {
	v := make([]Stamina, 0)
	for _, d := range data {
		v = append(v, NewStaminaFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminasFromDict(data []Stamina) []interface{} {
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
