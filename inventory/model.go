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

package inventory

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId             *string             `json:"namespaceId"`
	Name                    *string             `json:"name"`
	Description             *string             `json:"description"`
	TransactionSetting      *TransactionSetting `json:"transactionSetting"`
	AcquireScript           *ScriptSetting      `json:"acquireScript"`
	OverflowScript          *ScriptSetting      `json:"overflowScript"`
	ConsumeScript           *ScriptSetting      `json:"consumeScript"`
	SimpleItemAcquireScript *ScriptSetting      `json:"simpleItemAcquireScript"`
	SimpleItemConsumeScript *ScriptSetting      `json:"simpleItemConsumeScript"`
	BigItemAcquireScript    *ScriptSetting      `json:"bigItemAcquireScript"`
	BigItemConsumeScript    *ScriptSetting      `json:"bigItemConsumeScript"`
	LogSetting              *LogSetting         `json:"logSetting"`
	CreatedAt               *int64              `json:"createdAt"`
	UpdatedAt               *int64              `json:"updatedAt"`
	Revision                *int64              `json:"revision"`
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
		if v, ok := d["acquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireScript)
		}
		if v, ok := d["overflowScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OverflowScript)
		}
		if v, ok := d["consumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeScript)
		}
		if v, ok := d["simpleItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemAcquireScript)
		}
		if v, ok := d["simpleItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemConsumeScript)
		}
		if v, ok := d["bigItemAcquireScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemAcquireScript)
		}
		if v, ok := d["bigItemConsumeScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemConsumeScript)
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
		AcquireScript: func() *ScriptSetting {
			v, ok := data["acquireScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer()
		}(),
		OverflowScript: func() *ScriptSetting {
			v, ok := data["overflowScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer()
		}(),
		ConsumeScript: func() *ScriptSetting {
			v, ok := data["consumeScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer()
		}(),
		SimpleItemAcquireScript: func() *ScriptSetting {
			v, ok := data["simpleItemAcquireScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["simpleItemAcquireScript"])).Pointer()
		}(),
		SimpleItemConsumeScript: func() *ScriptSetting {
			v, ok := data["simpleItemConsumeScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["simpleItemConsumeScript"])).Pointer()
		}(),
		BigItemAcquireScript: func() *ScriptSetting {
			v, ok := data["bigItemAcquireScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["bigItemAcquireScript"])).Pointer()
		}(),
		BigItemConsumeScript: func() *ScriptSetting {
			v, ok := data["bigItemConsumeScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["bigItemConsumeScript"])).Pointer()
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
	if p.AcquireScript != nil {
		m["acquireScript"] = func() map[string]interface{} {
			if p.AcquireScript == nil {
				return nil
			}
			return p.AcquireScript.ToDict()
		}()
	}
	if p.OverflowScript != nil {
		m["overflowScript"] = func() map[string]interface{} {
			if p.OverflowScript == nil {
				return nil
			}
			return p.OverflowScript.ToDict()
		}()
	}
	if p.ConsumeScript != nil {
		m["consumeScript"] = func() map[string]interface{} {
			if p.ConsumeScript == nil {
				return nil
			}
			return p.ConsumeScript.ToDict()
		}()
	}
	if p.SimpleItemAcquireScript != nil {
		m["simpleItemAcquireScript"] = func() map[string]interface{} {
			if p.SimpleItemAcquireScript == nil {
				return nil
			}
			return p.SimpleItemAcquireScript.ToDict()
		}()
	}
	if p.SimpleItemConsumeScript != nil {
		m["simpleItemConsumeScript"] = func() map[string]interface{} {
			if p.SimpleItemConsumeScript == nil {
				return nil
			}
			return p.SimpleItemConsumeScript.ToDict()
		}()
	}
	if p.BigItemAcquireScript != nil {
		m["bigItemAcquireScript"] = func() map[string]interface{} {
			if p.BigItemAcquireScript == nil {
				return nil
			}
			return p.BigItemAcquireScript.ToDict()
		}()
	}
	if p.BigItemConsumeScript != nil {
		m["bigItemConsumeScript"] = func() map[string]interface{} {
			if p.BigItemConsumeScript == nil {
				return nil
			}
			return p.BigItemConsumeScript.ToDict()
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

type InventoryModelMaster struct {
	InventoryModelId      *string `json:"inventoryModelId"`
	Name                  *string `json:"name"`
	Metadata              *string `json:"metadata"`
	Description           *string `json:"description"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
	CreatedAt             *int64  `json:"createdAt"`
	UpdatedAt             *int64  `json:"updatedAt"`
	Revision              *int64  `json:"revision"`
}

func (p *InventoryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = InventoryModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = InventoryModelMaster{}
	} else {
		*p = InventoryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["protectReferencedItem"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ProtectReferencedItem)
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

func NewInventoryModelMasterFromJson(data string) InventoryModelMaster {
	req := InventoryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInventoryModelMasterFromDict(data map[string]interface{}) InventoryModelMaster {
	return InventoryModelMaster{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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
		InitialCapacity: func() *int32 {
			v, ok := data["initialCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialCapacity"])
		}(),
		MaxCapacity: func() *int32 {
			v, ok := data["maxCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxCapacity"])
		}(),
		ProtectReferencedItem: func() *bool {
			v, ok := data["protectReferencedItem"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["protectReferencedItem"])
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

func (p InventoryModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
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
	if p.InitialCapacity != nil {
		m["initialCapacity"] = p.InitialCapacity
	}
	if p.MaxCapacity != nil {
		m["maxCapacity"] = p.MaxCapacity
	}
	if p.ProtectReferencedItem != nil {
		m["protectReferencedItem"] = p.ProtectReferencedItem
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

func (p InventoryModelMaster) Pointer() *InventoryModelMaster {
	return &p
}

func CastInventoryModelMasters(data []interface{}) []InventoryModelMaster {
	v := make([]InventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryModelMastersFromDict(data []InventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryModel struct {
	InventoryModelId      *string     `json:"inventoryModelId"`
	Name                  *string     `json:"name"`
	Metadata              *string     `json:"metadata"`
	InitialCapacity       *int32      `json:"initialCapacity"`
	MaxCapacity           *int32      `json:"maxCapacity"`
	ProtectReferencedItem *bool       `json:"protectReferencedItem"`
	ItemModels            []ItemModel `json:"itemModels"`
}

func (p *InventoryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = InventoryModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = InventoryModel{}
	} else {
		*p = InventoryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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
		if v, ok := d["initialCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InitialCapacity)
		}
		if v, ok := d["maxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxCapacity)
		}
		if v, ok := d["protectReferencedItem"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ProtectReferencedItem)
		}
		if v, ok := d["itemModels"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ItemModels)
		}
	}
	return nil
}

func NewInventoryModelFromJson(data string) InventoryModel {
	req := InventoryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInventoryModelFromDict(data map[string]interface{}) InventoryModel {
	return InventoryModel{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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
		InitialCapacity: func() *int32 {
			v, ok := data["initialCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initialCapacity"])
		}(),
		MaxCapacity: func() *int32 {
			v, ok := data["maxCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxCapacity"])
		}(),
		ProtectReferencedItem: func() *bool {
			v, ok := data["protectReferencedItem"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["protectReferencedItem"])
		}(),
		ItemModels: func() []ItemModel {
			if data["itemModels"] == nil {
				return nil
			}
			return CastItemModels(core.CastArray(data["itemModels"]))
		}(),
	}
}

func (p InventoryModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.InitialCapacity != nil {
		m["initialCapacity"] = p.InitialCapacity
	}
	if p.MaxCapacity != nil {
		m["maxCapacity"] = p.MaxCapacity
	}
	if p.ProtectReferencedItem != nil {
		m["protectReferencedItem"] = p.ProtectReferencedItem
	}
	if p.ItemModels != nil {
		m["itemModels"] = CastItemModelsFromDict(
			p.ItemModels,
		)
	}
	return m
}

func (p InventoryModel) Pointer() *InventoryModel {
	return &p
}

func CastInventoryModels(data []interface{}) []InventoryModel {
	v := make([]InventoryModel, 0)
	for _, d := range data {
		v = append(v, NewInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryModelsFromDict(data []InventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemModelMaster struct {
	ItemModelId         *string `json:"itemModelId"`
	InventoryName       *string `json:"inventoryName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
	Revision            *int64  `json:"revision"`
}

func (p *ItemModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ItemModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ItemModelMaster{}
	} else {
		*p = ItemModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
				}
			}
		}
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["stackingLimit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StackingLimit)
		}
		if v, ok := d["allowMultipleStacks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AllowMultipleStacks)
		}
		if v, ok := d["sortValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortValue)
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

func NewItemModelMasterFromJson(data string) ItemModelMaster {
	req := ItemModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewItemModelMasterFromDict(data map[string]interface{}) ItemModelMaster {
	return ItemModelMaster{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
		}(),
		InventoryName: func() *string {
			v, ok := data["inventoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryName"])
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
		StackingLimit: func() *int64 {
			v, ok := data["stackingLimit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["stackingLimit"])
		}(),
		AllowMultipleStacks: func() *bool {
			v, ok := data["allowMultipleStacks"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["allowMultipleStacks"])
		}(),
		SortValue: func() *int32 {
			v, ok := data["sortValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sortValue"])
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

func (p ItemModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
	}
	if p.InventoryName != nil {
		m["inventoryName"] = p.InventoryName
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
	if p.StackingLimit != nil {
		m["stackingLimit"] = p.StackingLimit
	}
	if p.AllowMultipleStacks != nil {
		m["allowMultipleStacks"] = p.AllowMultipleStacks
	}
	if p.SortValue != nil {
		m["sortValue"] = p.SortValue
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

func (p ItemModelMaster) Pointer() *ItemModelMaster {
	return &p
}

func CastItemModelMasters(data []interface{}) []ItemModelMaster {
	v := make([]ItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemModelMastersFromDict(data []ItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemModel struct {
	ItemModelId         *string `json:"itemModelId"`
	Name                *string `json:"name"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func (p *ItemModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ItemModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ItemModel{}
	} else {
		*p = ItemModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
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
		if v, ok := d["stackingLimit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StackingLimit)
		}
		if v, ok := d["allowMultipleStacks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AllowMultipleStacks)
		}
		if v, ok := d["sortValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortValue)
		}
	}
	return nil
}

func NewItemModelFromJson(data string) ItemModel {
	req := ItemModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewItemModelFromDict(data map[string]interface{}) ItemModel {
	return ItemModel{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
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
		StackingLimit: func() *int64 {
			v, ok := data["stackingLimit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["stackingLimit"])
		}(),
		AllowMultipleStacks: func() *bool {
			v, ok := data["allowMultipleStacks"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["allowMultipleStacks"])
		}(),
		SortValue: func() *int32 {
			v, ok := data["sortValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sortValue"])
		}(),
	}
}

func (p ItemModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.StackingLimit != nil {
		m["stackingLimit"] = p.StackingLimit
	}
	if p.AllowMultipleStacks != nil {
		m["allowMultipleStacks"] = p.AllowMultipleStacks
	}
	if p.SortValue != nil {
		m["sortValue"] = p.SortValue
	}
	return m
}

func (p ItemModel) Pointer() *ItemModel {
	return &p
}

func CastItemModels(data []interface{}) []ItemModel {
	v := make([]ItemModel, 0)
	for _, d := range data {
		v = append(v, NewItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemModelsFromDict(data []ItemModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleInventoryModelMaster struct {
	InventoryModelId *string `json:"inventoryModelId"`
	Name             *string `json:"name"`
	Metadata         *string `json:"metadata"`
	Description      *string `json:"description"`
	CreatedAt        *int64  `json:"createdAt"`
	UpdatedAt        *int64  `json:"updatedAt"`
	Revision         *int64  `json:"revision"`
}

func (p *SimpleInventoryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleInventoryModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleInventoryModelMaster{}
	} else {
		*p = SimpleInventoryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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

func NewSimpleInventoryModelMasterFromJson(data string) SimpleInventoryModelMaster {
	req := SimpleInventoryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleInventoryModelMasterFromDict(data map[string]interface{}) SimpleInventoryModelMaster {
	return SimpleInventoryModelMaster{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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

func (p SimpleInventoryModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
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

func (p SimpleInventoryModelMaster) Pointer() *SimpleInventoryModelMaster {
	return &p
}

func CastSimpleInventoryModelMasters(data []interface{}) []SimpleInventoryModelMaster {
	v := make([]SimpleInventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoryModelMastersFromDict(data []SimpleInventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleInventoryModel struct {
	InventoryModelId *string           `json:"inventoryModelId"`
	Name             *string           `json:"name"`
	Metadata         *string           `json:"metadata"`
	SimpleItemModels []SimpleItemModel `json:"simpleItemModels"`
}

func (p *SimpleInventoryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleInventoryModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleInventoryModel{}
	} else {
		*p = SimpleInventoryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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
		if v, ok := d["simpleItemModels"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItemModels)
		}
	}
	return nil
}

func NewSimpleInventoryModelFromJson(data string) SimpleInventoryModel {
	req := SimpleInventoryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleInventoryModelFromDict(data map[string]interface{}) SimpleInventoryModel {
	return SimpleInventoryModel{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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
		SimpleItemModels: func() []SimpleItemModel {
			if data["simpleItemModels"] == nil {
				return nil
			}
			return CastSimpleItemModels(core.CastArray(data["simpleItemModels"]))
		}(),
	}
}

func (p SimpleInventoryModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.SimpleItemModels != nil {
		m["simpleItemModels"] = CastSimpleItemModelsFromDict(
			p.SimpleItemModels,
		)
	}
	return m
}

func (p SimpleInventoryModel) Pointer() *SimpleInventoryModel {
	return &p
}

func CastSimpleInventoryModels(data []interface{}) []SimpleInventoryModel {
	v := make([]SimpleInventoryModel, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoryModelsFromDict(data []SimpleInventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItemModelMaster struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Metadata    *string `json:"metadata"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func (p *SimpleItemModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleItemModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleItemModelMaster{}
	} else {
		*p = SimpleItemModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
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

func NewSimpleItemModelMasterFromJson(data string) SimpleItemModelMaster {
	req := SimpleItemModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleItemModelMasterFromDict(data map[string]interface{}) SimpleItemModelMaster {
	return SimpleItemModelMaster{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
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

func (p SimpleItemModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
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

func (p SimpleItemModelMaster) Pointer() *SimpleItemModelMaster {
	return &p
}

func CastSimpleItemModelMasters(data []interface{}) []SimpleItemModelMaster {
	v := make([]SimpleItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemModelMastersFromDict(data []SimpleItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItemModel struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Metadata    *string `json:"metadata"`
}

func (p *SimpleItemModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleItemModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleItemModel{}
	} else {
		*p = SimpleItemModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
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
	}
	return nil
}

func NewSimpleItemModelFromJson(data string) SimpleItemModel {
	req := SimpleItemModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleItemModelFromDict(data map[string]interface{}) SimpleItemModel {
	return SimpleItemModel{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
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
	}
}

func (p SimpleItemModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	return m
}

func (p SimpleItemModel) Pointer() *SimpleItemModel {
	return &p
}

func CastSimpleItemModels(data []interface{}) []SimpleItemModel {
	v := make([]SimpleItemModel, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemModelsFromDict(data []SimpleItemModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventoryModelMaster struct {
	InventoryModelId *string `json:"inventoryModelId"`
	Name             *string `json:"name"`
	Metadata         *string `json:"metadata"`
	Description      *string `json:"description"`
	CreatedAt        *int64  `json:"createdAt"`
	UpdatedAt        *int64  `json:"updatedAt"`
	Revision         *int64  `json:"revision"`
}

func (p *BigInventoryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigInventoryModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigInventoryModelMaster{}
	} else {
		*p = BigInventoryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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

func NewBigInventoryModelMasterFromJson(data string) BigInventoryModelMaster {
	req := BigInventoryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigInventoryModelMasterFromDict(data map[string]interface{}) BigInventoryModelMaster {
	return BigInventoryModelMaster{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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

func (p BigInventoryModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
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

func (p BigInventoryModelMaster) Pointer() *BigInventoryModelMaster {
	return &p
}

func CastBigInventoryModelMasters(data []interface{}) []BigInventoryModelMaster {
	v := make([]BigInventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoryModelMastersFromDict(data []BigInventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventoryModel struct {
	InventoryModelId *string        `json:"inventoryModelId"`
	Name             *string        `json:"name"`
	Metadata         *string        `json:"metadata"`
	BigItemModels    []BigItemModel `json:"bigItemModels"`
}

func (p *BigInventoryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigInventoryModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigInventoryModel{}
	} else {
		*p = BigInventoryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryModelId)
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
		if v, ok := d["bigItemModels"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItemModels)
		}
	}
	return nil
}

func NewBigInventoryModelFromJson(data string) BigInventoryModel {
	req := BigInventoryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigInventoryModelFromDict(data map[string]interface{}) BigInventoryModel {
	return BigInventoryModel{
		InventoryModelId: func() *string {
			v, ok := data["inventoryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryModelId"])
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
		BigItemModels: func() []BigItemModel {
			if data["bigItemModels"] == nil {
				return nil
			}
			return CastBigItemModels(core.CastArray(data["bigItemModels"]))
		}(),
	}
}

func (p BigInventoryModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryModelId != nil {
		m["inventoryModelId"] = p.InventoryModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.BigItemModels != nil {
		m["bigItemModels"] = CastBigItemModelsFromDict(
			p.BigItemModels,
		)
	}
	return m
}

func (p BigInventoryModel) Pointer() *BigInventoryModel {
	return &p
}

func CastBigInventoryModels(data []interface{}) []BigInventoryModel {
	v := make([]BigInventoryModel, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoryModelsFromDict(data []BigInventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItemModelMaster struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Metadata    *string `json:"metadata"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func (p *BigItemModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigItemModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigItemModelMaster{}
	} else {
		*p = BigItemModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
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

func NewBigItemModelMasterFromJson(data string) BigItemModelMaster {
	req := BigItemModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigItemModelMasterFromDict(data map[string]interface{}) BigItemModelMaster {
	return BigItemModelMaster{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
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

func (p BigItemModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
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

func (p BigItemModelMaster) Pointer() *BigItemModelMaster {
	return &p
}

func CastBigItemModelMasters(data []interface{}) []BigItemModelMaster {
	v := make([]BigItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBigItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemModelMastersFromDict(data []BigItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItemModel struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Metadata    *string `json:"metadata"`
}

func (p *BigItemModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigItemModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigItemModel{}
	} else {
		*p = BigItemModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemModelId)
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
	}
	return nil
}

func NewBigItemModelFromJson(data string) BigItemModel {
	req := BigItemModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigItemModelFromDict(data map[string]interface{}) BigItemModel {
	return BigItemModel{
		ItemModelId: func() *string {
			v, ok := data["itemModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemModelId"])
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
	}
}

func (p BigItemModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemModelId != nil {
		m["itemModelId"] = p.ItemModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	return m
}

func (p BigItemModel) Pointer() *BigItemModel {
	return &p
}

func CastBigItemModels(data []interface{}) []BigItemModel {
	v := make([]BigItemModel, 0)
	for _, d := range data {
		v = append(v, NewBigItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemModelsFromDict(data []BigItemModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentItemModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentItemModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentItemModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentItemModelMaster{}
	} else {
		*p = CurrentItemModelMaster{}
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

func NewCurrentItemModelMasterFromJson(data string) CurrentItemModelMaster {
	req := CurrentItemModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentItemModelMasterFromDict(data map[string]interface{}) CurrentItemModelMaster {
	return CurrentItemModelMaster{
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

func (p CurrentItemModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentItemModelMaster) Pointer() *CurrentItemModelMaster {
	return &p
}

func CastCurrentItemModelMasters(data []interface{}) []CurrentItemModelMaster {
	v := make([]CurrentItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentItemModelMastersFromDict(data []CurrentItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Inventory struct {
	InventoryId                   *string `json:"inventoryId"`
	InventoryName                 *string `json:"inventoryName"`
	UserId                        *string `json:"userId"`
	CurrentInventoryCapacityUsage *int32  `json:"currentInventoryCapacityUsage"`
	CurrentInventoryMaxCapacity   *int32  `json:"currentInventoryMaxCapacity"`
	CreatedAt                     *int64  `json:"createdAt"`
	UpdatedAt                     *int64  `json:"updatedAt"`
	Revision                      *int64  `json:"revision"`
}

func (p *Inventory) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Inventory{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Inventory{}
	} else {
		*p = Inventory{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryId)
				}
			}
		}
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["currentInventoryCapacityUsage"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentInventoryCapacityUsage)
		}
		if v, ok := d["currentInventoryMaxCapacity"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentInventoryMaxCapacity)
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

func NewInventoryFromJson(data string) Inventory {
	req := Inventory{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInventoryFromDict(data map[string]interface{}) Inventory {
	return Inventory{
		InventoryId: func() *string {
			v, ok := data["inventoryId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryId"])
		}(),
		InventoryName: func() *string {
			v, ok := data["inventoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		CurrentInventoryCapacityUsage: func() *int32 {
			v, ok := data["currentInventoryCapacityUsage"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["currentInventoryCapacityUsage"])
		}(),
		CurrentInventoryMaxCapacity: func() *int32 {
			v, ok := data["currentInventoryMaxCapacity"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["currentInventoryMaxCapacity"])
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

func (p Inventory) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryId != nil {
		m["inventoryId"] = p.InventoryId
	}
	if p.InventoryName != nil {
		m["inventoryName"] = p.InventoryName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.CurrentInventoryCapacityUsage != nil {
		m["currentInventoryCapacityUsage"] = p.CurrentInventoryCapacityUsage
	}
	if p.CurrentInventoryMaxCapacity != nil {
		m["currentInventoryMaxCapacity"] = p.CurrentInventoryMaxCapacity
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

func (p Inventory) Pointer() *Inventory {
	return &p
}

func CastInventories(data []interface{}) []Inventory {
	v := make([]Inventory, 0)
	for _, d := range data {
		v = append(v, NewInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoriesFromDict(data []Inventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemSet struct {
	ItemSetId     *string   `json:"itemSetId"`
	Name          *string   `json:"name"`
	InventoryName *string   `json:"inventoryName"`
	UserId        *string   `json:"userId"`
	ItemName      *string   `json:"itemName"`
	Count         *int64    `json:"count"`
	ReferenceOf   []*string `json:"referenceOf"`
	SortValue     *int32    `json:"sortValue"`
	ExpiresAt     *int64    `json:"expiresAt"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
}

func (p *ItemSet) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ItemSet{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ItemSet{}
	} else {
		*p = ItemSet{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemSetId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemSetId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemSetId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemSetId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemSetId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemSetId)
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
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["referenceOf"]; ok && v != nil {
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
				p.ReferenceOf = l
			}
		}
		if v, ok := d["sortValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortValue)
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
	}
	return nil
}

func NewItemSetFromJson(data string) ItemSet {
	req := ItemSet{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewItemSetFromDict(data map[string]interface{}) ItemSet {
	return ItemSet{
		ItemSetId: func() *string {
			v, ok := data["itemSetId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemSetId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		InventoryName: func() *string {
			v, ok := data["inventoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
		ReferenceOf: func() []*string {
			v, ok := data["referenceOf"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		SortValue: func() *int32 {
			v, ok := data["sortValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sortValue"])
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
	}
}

func (p ItemSet) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemSetId != nil {
		m["itemSetId"] = p.ItemSetId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.InventoryName != nil {
		m["inventoryName"] = p.InventoryName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	if p.ReferenceOf != nil {
		m["referenceOf"] = core.CastStringsFromDict(
			p.ReferenceOf,
		)
	}
	if p.SortValue != nil {
		m["sortValue"] = p.SortValue
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
	return m
}

func (p ItemSet) Pointer() *ItemSet {
	return &p
}

func CastItemSets(data []interface{}) []ItemSet {
	v := make([]ItemSet, 0)
	for _, d := range data {
		v = append(v, NewItemSetFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemSetsFromDict(data []ItemSet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReferenceOf struct {
	ReferenceOfId *string `json:"referenceOfId"`
	Name          *string `json:"name"`
}

func (p *ReferenceOf) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ReferenceOf{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ReferenceOf{}
	} else {
		*p = ReferenceOf{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["referenceOfId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceOfId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceOfId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceOfId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOfId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceOfId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceOfId)
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
	}
	return nil
}

func NewReferenceOfFromJson(data string) ReferenceOf {
	req := ReferenceOf{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewReferenceOfFromDict(data map[string]interface{}) ReferenceOf {
	return ReferenceOf{
		ReferenceOfId: func() *string {
			v, ok := data["referenceOfId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["referenceOfId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
	}
}

func (p ReferenceOf) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ReferenceOfId != nil {
		m["referenceOfId"] = p.ReferenceOfId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	return m
}

func (p ReferenceOf) Pointer() *ReferenceOf {
	return &p
}

func CastReferenceOves(data []interface{}) []ReferenceOf {
	v := make([]ReferenceOf, 0)
	for _, d := range data {
		v = append(v, NewReferenceOfFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReferenceOvesFromDict(data []ReferenceOf) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleInventory struct {
	InventoryId   *string      `json:"inventoryId"`
	InventoryName *string      `json:"inventoryName"`
	UserId        *string      `json:"userId"`
	SimpleItems   []SimpleItem `json:"simpleItems"`
	CreatedAt     *int64       `json:"createdAt"`
	UpdatedAt     *int64       `json:"updatedAt"`
	Revision      *int64       `json:"revision"`
}

func (p *SimpleInventory) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleInventory{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleInventory{}
	} else {
		*p = SimpleInventory{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryId)
				}
			}
		}
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["simpleItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SimpleItems)
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

func NewSimpleInventoryFromJson(data string) SimpleInventory {
	req := SimpleInventory{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleInventoryFromDict(data map[string]interface{}) SimpleInventory {
	return SimpleInventory{
		InventoryId: func() *string {
			v, ok := data["inventoryId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryId"])
		}(),
		InventoryName: func() *string {
			v, ok := data["inventoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		SimpleItems: func() []SimpleItem {
			if data["simpleItems"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["simpleItems"]))
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

func (p SimpleInventory) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryId != nil {
		m["inventoryId"] = p.InventoryId
	}
	if p.InventoryName != nil {
		m["inventoryName"] = p.InventoryName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.SimpleItems != nil {
		m["simpleItems"] = CastSimpleItemsFromDict(
			p.SimpleItems,
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

func (p SimpleInventory) Pointer() *SimpleInventory {
	return &p
}

func CastSimpleInventories(data []interface{}) []SimpleInventory {
	v := make([]SimpleInventory, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoriesFromDict(data []SimpleInventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItem struct {
	ItemId   *string `json:"itemId"`
	UserId   *string `json:"userId"`
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
	Revision *int64  `json:"revision"`
}

func (p *SimpleItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SimpleItem{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = SimpleItem{}
	} else {
		*p = SimpleItem{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemId)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewSimpleItemFromJson(data string) SimpleItem {
	req := SimpleItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSimpleItemFromDict(data map[string]interface{}) SimpleItem {
	return SimpleItem{
		ItemId: func() *string {
			v, ok := data["itemId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
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

func (p SimpleItem) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemId != nil {
		m["itemId"] = p.ItemId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p SimpleItem) Pointer() *SimpleItem {
	return &p
}

func CastSimpleItems(data []interface{}) []SimpleItem {
	v := make([]SimpleItem, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemsFromDict(data []SimpleItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventory struct {
	InventoryId   *string   `json:"inventoryId"`
	InventoryName *string   `json:"inventoryName"`
	UserId        *string   `json:"userId"`
	BigItems      []BigItem `json:"bigItems"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
}

func (p *BigInventory) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigInventory{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigInventory{}
	} else {
		*p = BigInventory{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["inventoryId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryId)
				}
			}
		}
		if v, ok := d["inventoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InventoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InventoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InventoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InventoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InventoryName)
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
		if v, ok := d["bigItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BigItems)
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

func NewBigInventoryFromJson(data string) BigInventory {
	req := BigInventory{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigInventoryFromDict(data map[string]interface{}) BigInventory {
	return BigInventory{
		InventoryId: func() *string {
			v, ok := data["inventoryId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryId"])
		}(),
		InventoryName: func() *string {
			v, ok := data["inventoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["inventoryName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		BigItems: func() []BigItem {
			if data["bigItems"] == nil {
				return nil
			}
			return CastBigItems(core.CastArray(data["bigItems"]))
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

func (p BigInventory) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InventoryId != nil {
		m["inventoryId"] = p.InventoryId
	}
	if p.InventoryName != nil {
		m["inventoryName"] = p.InventoryName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.BigItems != nil {
		m["bigItems"] = CastBigItemsFromDict(
			p.BigItems,
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

func (p BigInventory) Pointer() *BigInventory {
	return &p
}

func CastBigInventories(data []interface{}) []BigInventory {
	v := make([]BigInventory, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoriesFromDict(data []BigInventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItem struct {
	ItemId    *string `json:"itemId"`
	UserId    *string `json:"userId"`
	ItemName  *string `json:"itemName"`
	Count     *string `json:"count"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
	Revision  *int64  `json:"revision"`
}

func (p *BigItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BigItem{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BigItem{}
	} else {
		*p = BigItem{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemId)
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
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Count = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Count = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Count = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Count = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Count)
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

func NewBigItemFromJson(data string) BigItem {
	req := BigItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBigItemFromDict(data map[string]interface{}) BigItem {
	return BigItem{
		ItemId: func() *string {
			v, ok := data["itemId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *string {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["count"])
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

func (p BigItem) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemId != nil {
		m["itemId"] = p.ItemId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
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

func (p BigItem) Pointer() *BigItem {
	return &p
}

func CastBigItems(data []interface{}) []BigItem {
	v := make([]BigItem, 0)
	for _, d := range data {
		v = append(v, NewBigItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemsFromDict(data []BigItem) []interface{} {
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

type AcquireCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func (p *AcquireCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireCount{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = AcquireCount{}
	} else {
		*p = AcquireCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewAcquireCountFromJson(data string) AcquireCount {
	req := AcquireCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireCountFromDict(data map[string]interface{}) AcquireCount {
	return AcquireCount{
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p AcquireCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
}

func (p AcquireCount) Pointer() *AcquireCount {
	return &p
}

func CastAcquireCounts(data []interface{}) []AcquireCount {
	v := make([]AcquireCount, 0)
	for _, d := range data {
		v = append(v, NewAcquireCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireCountsFromDict(data []AcquireCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func (p *ConsumeCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeCount{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = ConsumeCount{}
	} else {
		*p = ConsumeCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewConsumeCountFromJson(data string) ConsumeCount {
	req := ConsumeCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeCountFromDict(data map[string]interface{}) ConsumeCount {
	return ConsumeCount{
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p ConsumeCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
}

func (p ConsumeCount) Pointer() *ConsumeCount {
	return &p
}

func CastConsumeCounts(data []interface{}) []ConsumeCount {
	v := make([]ConsumeCount, 0)
	for _, d := range data {
		v = append(v, NewConsumeCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeCountsFromDict(data []ConsumeCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type HeldCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func (p *HeldCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = HeldCount{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = HeldCount{}
	} else {
		*p = HeldCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["itemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ItemName)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewHeldCountFromJson(data string) HeldCount {
	req := HeldCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewHeldCountFromDict(data map[string]interface{}) HeldCount {
	return HeldCount{
		ItemName: func() *string {
			v, ok := data["itemName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["itemName"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p HeldCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ItemName != nil {
		m["itemName"] = p.ItemName
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
}

func (p HeldCount) Pointer() *HeldCount {
	return &p
}

func CastHeldCounts(data []interface{}) []HeldCount {
	v := make([]HeldCount, 0)
	for _, d := range data {
		v = append(v, NewHeldCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastHeldCountsFromDict(data []HeldCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
