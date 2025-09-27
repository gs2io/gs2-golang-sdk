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

package enchant

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

type BalanceParameterModel struct {
	BalanceParameterModelId *string                      `json:"balanceParameterModelId"`
	Name                    *string                      `json:"name"`
	Metadata                *string                      `json:"metadata"`
	TotalValue              *int64                       `json:"totalValue"`
	InitialValueStrategy    *string                      `json:"initialValueStrategy"`
	Parameters              []BalanceParameterValueModel `json:"parameters"`
}

func (p *BalanceParameterModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BalanceParameterModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BalanceParameterModel{}
	} else {
		*p = BalanceParameterModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["balanceParameterModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BalanceParameterModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BalanceParameterModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BalanceParameterModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BalanceParameterModelId)
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
		if v, ok := d["totalValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TotalValue)
		}
		if v, ok := d["initialValueStrategy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InitialValueStrategy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InitialValueStrategy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InitialValueStrategy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InitialValueStrategy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InitialValueStrategy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InitialValueStrategy)
				}
			}
		}
		if v, ok := d["parameters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Parameters)
		}
	}
	return nil
}

func NewBalanceParameterModelFromJson(data string) BalanceParameterModel {
	req := BalanceParameterModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBalanceParameterModelFromDict(data map[string]interface{}) BalanceParameterModel {
	return BalanceParameterModel{
		BalanceParameterModelId: func() *string {
			v, ok := data["balanceParameterModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["balanceParameterModelId"])
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
		TotalValue: func() *int64 {
			v, ok := data["totalValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalValue"])
		}(),
		InitialValueStrategy: func() *string {
			v, ok := data["initialValueStrategy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["initialValueStrategy"])
		}(),
		Parameters: func() []BalanceParameterValueModel {
			if data["parameters"] == nil {
				return nil
			}
			return CastBalanceParameterValueModels(core.CastArray(data["parameters"]))
		}(),
	}
}

func (p BalanceParameterModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BalanceParameterModelId != nil {
		m["balanceParameterModelId"] = p.BalanceParameterModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.TotalValue != nil {
		m["totalValue"] = p.TotalValue
	}
	if p.InitialValueStrategy != nil {
		m["initialValueStrategy"] = p.InitialValueStrategy
	}
	if p.Parameters != nil {
		m["parameters"] = CastBalanceParameterValueModelsFromDict(
			p.Parameters,
		)
	}
	return m
}

func (p BalanceParameterModel) Pointer() *BalanceParameterModel {
	return &p
}

func CastBalanceParameterModels(data []interface{}) []BalanceParameterModel {
	v := make([]BalanceParameterModel, 0)
	for _, d := range data {
		v = append(v, NewBalanceParameterModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBalanceParameterModelsFromDict(data []BalanceParameterModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BalanceParameterModelMaster struct {
	BalanceParameterModelId *string                      `json:"balanceParameterModelId"`
	Name                    *string                      `json:"name"`
	Description             *string                      `json:"description"`
	Metadata                *string                      `json:"metadata"`
	TotalValue              *int64                       `json:"totalValue"`
	InitialValueStrategy    *string                      `json:"initialValueStrategy"`
	Parameters              []BalanceParameterValueModel `json:"parameters"`
	CreatedAt               *int64                       `json:"createdAt"`
	UpdatedAt               *int64                       `json:"updatedAt"`
	Revision                *int64                       `json:"revision"`
}

func (p *BalanceParameterModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BalanceParameterModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BalanceParameterModelMaster{}
	} else {
		*p = BalanceParameterModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["balanceParameterModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BalanceParameterModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BalanceParameterModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BalanceParameterModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BalanceParameterModelId)
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
		if v, ok := d["totalValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TotalValue)
		}
		if v, ok := d["initialValueStrategy"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InitialValueStrategy = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InitialValueStrategy = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InitialValueStrategy = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InitialValueStrategy = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InitialValueStrategy = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InitialValueStrategy)
				}
			}
		}
		if v, ok := d["parameters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Parameters)
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

func NewBalanceParameterModelMasterFromJson(data string) BalanceParameterModelMaster {
	req := BalanceParameterModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBalanceParameterModelMasterFromDict(data map[string]interface{}) BalanceParameterModelMaster {
	return BalanceParameterModelMaster{
		BalanceParameterModelId: func() *string {
			v, ok := data["balanceParameterModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["balanceParameterModelId"])
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
		TotalValue: func() *int64 {
			v, ok := data["totalValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["totalValue"])
		}(),
		InitialValueStrategy: func() *string {
			v, ok := data["initialValueStrategy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["initialValueStrategy"])
		}(),
		Parameters: func() []BalanceParameterValueModel {
			if data["parameters"] == nil {
				return nil
			}
			return CastBalanceParameterValueModels(core.CastArray(data["parameters"]))
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

func (p BalanceParameterModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BalanceParameterModelId != nil {
		m["balanceParameterModelId"] = p.BalanceParameterModelId
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
	if p.TotalValue != nil {
		m["totalValue"] = p.TotalValue
	}
	if p.InitialValueStrategy != nil {
		m["initialValueStrategy"] = p.InitialValueStrategy
	}
	if p.Parameters != nil {
		m["parameters"] = CastBalanceParameterValueModelsFromDict(
			p.Parameters,
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

func (p BalanceParameterModelMaster) Pointer() *BalanceParameterModelMaster {
	return &p
}

func CastBalanceParameterModelMasters(data []interface{}) []BalanceParameterModelMaster {
	v := make([]BalanceParameterModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBalanceParameterModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBalanceParameterModelMastersFromDict(data []BalanceParameterModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterModel struct {
	RarityParameterModelId *string                     `json:"rarityParameterModelId"`
	Name                   *string                     `json:"name"`
	Metadata               *string                     `json:"metadata"`
	MaximumParameterCount  *int32                      `json:"maximumParameterCount"`
	ParameterCounts        []RarityParameterCountModel `json:"parameterCounts"`
	Parameters             []RarityParameterValueModel `json:"parameters"`
}

func (p *RarityParameterModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterModel{}
	} else {
		*p = RarityParameterModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rarityParameterModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RarityParameterModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RarityParameterModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RarityParameterModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RarityParameterModelId)
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
		if v, ok := d["maximumParameterCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParameterCount)
		}
		if v, ok := d["parameterCounts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ParameterCounts)
		}
		if v, ok := d["parameters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Parameters)
		}
	}
	return nil
}

func NewRarityParameterModelFromJson(data string) RarityParameterModel {
	req := RarityParameterModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterModelFromDict(data map[string]interface{}) RarityParameterModel {
	return RarityParameterModel{
		RarityParameterModelId: func() *string {
			v, ok := data["rarityParameterModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rarityParameterModelId"])
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
		MaximumParameterCount: func() *int32 {
			v, ok := data["maximumParameterCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParameterCount"])
		}(),
		ParameterCounts: func() []RarityParameterCountModel {
			if data["parameterCounts"] == nil {
				return nil
			}
			return CastRarityParameterCountModels(core.CastArray(data["parameterCounts"]))
		}(),
		Parameters: func() []RarityParameterValueModel {
			if data["parameters"] == nil {
				return nil
			}
			return CastRarityParameterValueModels(core.CastArray(data["parameters"]))
		}(),
	}
}

func (p RarityParameterModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RarityParameterModelId != nil {
		m["rarityParameterModelId"] = p.RarityParameterModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.MaximumParameterCount != nil {
		m["maximumParameterCount"] = p.MaximumParameterCount
	}
	if p.ParameterCounts != nil {
		m["parameterCounts"] = CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		)
	}
	if p.Parameters != nil {
		m["parameters"] = CastRarityParameterValueModelsFromDict(
			p.Parameters,
		)
	}
	return m
}

func (p RarityParameterModel) Pointer() *RarityParameterModel {
	return &p
}

func CastRarityParameterModels(data []interface{}) []RarityParameterModel {
	v := make([]RarityParameterModel, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterModelsFromDict(data []RarityParameterModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterModelMaster struct {
	RarityParameterModelId *string                     `json:"rarityParameterModelId"`
	Name                   *string                     `json:"name"`
	Description            *string                     `json:"description"`
	Metadata               *string                     `json:"metadata"`
	MaximumParameterCount  *int32                      `json:"maximumParameterCount"`
	ParameterCounts        []RarityParameterCountModel `json:"parameterCounts"`
	Parameters             []RarityParameterValueModel `json:"parameters"`
	CreatedAt              *int64                      `json:"createdAt"`
	UpdatedAt              *int64                      `json:"updatedAt"`
	Revision               *int64                      `json:"revision"`
}

func (p *RarityParameterModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterModelMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterModelMaster{}
	} else {
		*p = RarityParameterModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rarityParameterModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RarityParameterModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RarityParameterModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RarityParameterModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RarityParameterModelId)
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
		if v, ok := d["maximumParameterCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumParameterCount)
		}
		if v, ok := d["parameterCounts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ParameterCounts)
		}
		if v, ok := d["parameters"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Parameters)
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

func NewRarityParameterModelMasterFromJson(data string) RarityParameterModelMaster {
	req := RarityParameterModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterModelMasterFromDict(data map[string]interface{}) RarityParameterModelMaster {
	return RarityParameterModelMaster{
		RarityParameterModelId: func() *string {
			v, ok := data["rarityParameterModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rarityParameterModelId"])
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
		MaximumParameterCount: func() *int32 {
			v, ok := data["maximumParameterCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maximumParameterCount"])
		}(),
		ParameterCounts: func() []RarityParameterCountModel {
			if data["parameterCounts"] == nil {
				return nil
			}
			return CastRarityParameterCountModels(core.CastArray(data["parameterCounts"]))
		}(),
		Parameters: func() []RarityParameterValueModel {
			if data["parameters"] == nil {
				return nil
			}
			return CastRarityParameterValueModels(core.CastArray(data["parameters"]))
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

func (p RarityParameterModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RarityParameterModelId != nil {
		m["rarityParameterModelId"] = p.RarityParameterModelId
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
	if p.MaximumParameterCount != nil {
		m["maximumParameterCount"] = p.MaximumParameterCount
	}
	if p.ParameterCounts != nil {
		m["parameterCounts"] = CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		)
	}
	if p.Parameters != nil {
		m["parameters"] = CastRarityParameterValueModelsFromDict(
			p.Parameters,
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

func (p RarityParameterModelMaster) Pointer() *RarityParameterModelMaster {
	return &p
}

func CastRarityParameterModelMasters(data []interface{}) []RarityParameterModelMaster {
	v := make([]RarityParameterModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterModelMastersFromDict(data []RarityParameterModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentParameterMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentParameterMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentParameterMaster{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = CurrentParameterMaster{}
	} else {
		*p = CurrentParameterMaster{}
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

func NewCurrentParameterMasterFromJson(data string) CurrentParameterMaster {
	req := CurrentParameterMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentParameterMasterFromDict(data map[string]interface{}) CurrentParameterMaster {
	return CurrentParameterMaster{
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

func (p CurrentParameterMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentParameterMaster) Pointer() *CurrentParameterMaster {
	return &p
}

func CastCurrentParameterMasters(data []interface{}) []CurrentParameterMaster {
	v := make([]CurrentParameterMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentParameterMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentParameterMastersFromDict(data []CurrentParameterMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BalanceParameterStatus struct {
	BalanceParameterStatusId *string                 `json:"balanceParameterStatusId"`
	UserId                   *string                 `json:"userId"`
	ParameterName            *string                 `json:"parameterName"`
	PropertyId               *string                 `json:"propertyId"`
	ParameterValues          []BalanceParameterValue `json:"parameterValues"`
	CreatedAt                *int64                  `json:"createdAt"`
	UpdatedAt                *int64                  `json:"updatedAt"`
	Revision                 *int64                  `json:"revision"`
}

func (p *BalanceParameterStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BalanceParameterStatus{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BalanceParameterStatus{}
	} else {
		*p = BalanceParameterStatus{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["balanceParameterStatusId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BalanceParameterStatusId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BalanceParameterStatusId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BalanceParameterStatusId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterStatusId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BalanceParameterStatusId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BalanceParameterStatusId)
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
		if v, ok := d["parameterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ParameterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ParameterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ParameterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ParameterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ParameterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ParameterName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["parameterValues"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ParameterValues)
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

func NewBalanceParameterStatusFromJson(data string) BalanceParameterStatus {
	req := BalanceParameterStatus{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBalanceParameterStatusFromDict(data map[string]interface{}) BalanceParameterStatus {
	return BalanceParameterStatus{
		BalanceParameterStatusId: func() *string {
			v, ok := data["balanceParameterStatusId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["balanceParameterStatusId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ParameterName: func() *string {
			v, ok := data["parameterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["parameterName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ParameterValues: func() []BalanceParameterValue {
			if data["parameterValues"] == nil {
				return nil
			}
			return CastBalanceParameterValues(core.CastArray(data["parameterValues"]))
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

func (p BalanceParameterStatus) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BalanceParameterStatusId != nil {
		m["balanceParameterStatusId"] = p.BalanceParameterStatusId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ParameterName != nil {
		m["parameterName"] = p.ParameterName
	}
	if p.PropertyId != nil {
		m["propertyId"] = p.PropertyId
	}
	if p.ParameterValues != nil {
		m["parameterValues"] = CastBalanceParameterValuesFromDict(
			p.ParameterValues,
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

func (p BalanceParameterStatus) Pointer() *BalanceParameterStatus {
	return &p
}

func CastBalanceParameterStatuses(data []interface{}) []BalanceParameterStatus {
	v := make([]BalanceParameterStatus, 0)
	for _, d := range data {
		v = append(v, NewBalanceParameterStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBalanceParameterStatusesFromDict(data []BalanceParameterStatus) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterStatus struct {
	RarityParameterStatusId *string                `json:"rarityParameterStatusId"`
	UserId                  *string                `json:"userId"`
	ParameterName           *string                `json:"parameterName"`
	PropertyId              *string                `json:"propertyId"`
	ParameterValues         []RarityParameterValue `json:"parameterValues"`
	CreatedAt               *int64                 `json:"createdAt"`
	UpdatedAt               *int64                 `json:"updatedAt"`
	Revision                *int64                 `json:"revision"`
}

func (p *RarityParameterStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterStatus{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterStatus{}
	} else {
		*p = RarityParameterStatus{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rarityParameterStatusId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RarityParameterStatusId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RarityParameterStatusId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RarityParameterStatusId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterStatusId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RarityParameterStatusId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RarityParameterStatusId)
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
		if v, ok := d["parameterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ParameterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ParameterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ParameterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ParameterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ParameterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ParameterName)
				}
			}
		}
		if v, ok := d["propertyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PropertyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PropertyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PropertyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PropertyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PropertyId)
				}
			}
		}
		if v, ok := d["parameterValues"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ParameterValues)
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

func NewRarityParameterStatusFromJson(data string) RarityParameterStatus {
	req := RarityParameterStatus{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterStatusFromDict(data map[string]interface{}) RarityParameterStatus {
	return RarityParameterStatus{
		RarityParameterStatusId: func() *string {
			v, ok := data["rarityParameterStatusId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rarityParameterStatusId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ParameterName: func() *string {
			v, ok := data["parameterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["parameterName"])
		}(),
		PropertyId: func() *string {
			v, ok := data["propertyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["propertyId"])
		}(),
		ParameterValues: func() []RarityParameterValue {
			if data["parameterValues"] == nil {
				return nil
			}
			return CastRarityParameterValues(core.CastArray(data["parameterValues"]))
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

func (p RarityParameterStatus) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RarityParameterStatusId != nil {
		m["rarityParameterStatusId"] = p.RarityParameterStatusId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ParameterName != nil {
		m["parameterName"] = p.ParameterName
	}
	if p.PropertyId != nil {
		m["propertyId"] = p.PropertyId
	}
	if p.ParameterValues != nil {
		m["parameterValues"] = CastRarityParameterValuesFromDict(
			p.ParameterValues,
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

func (p RarityParameterStatus) Pointer() *RarityParameterStatus {
	return &p
}

func CastRarityParameterStatuses(data []interface{}) []RarityParameterStatus {
	v := make([]RarityParameterStatus, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterStatusesFromDict(data []RarityParameterStatus) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BalanceParameterValueModel struct {
	Name     *string `json:"name"`
	Metadata *string `json:"metadata"`
}

func (p *BalanceParameterValueModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BalanceParameterValueModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BalanceParameterValueModel{}
	} else {
		*p = BalanceParameterValueModel{}
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
	}
	return nil
}

func NewBalanceParameterValueModelFromJson(data string) BalanceParameterValueModel {
	req := BalanceParameterValueModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBalanceParameterValueModelFromDict(data map[string]interface{}) BalanceParameterValueModel {
	return BalanceParameterValueModel{
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

func (p BalanceParameterValueModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	return m
}

func (p BalanceParameterValueModel) Pointer() *BalanceParameterValueModel {
	return &p
}

func CastBalanceParameterValueModels(data []interface{}) []BalanceParameterValueModel {
	v := make([]BalanceParameterValueModel, 0)
	for _, d := range data {
		v = append(v, NewBalanceParameterValueModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBalanceParameterValueModelsFromDict(data []BalanceParameterValueModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterCountModel struct {
	Count  *int32 `json:"count"`
	Weight *int32 `json:"weight"`
}

func (p *RarityParameterCountModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterCountModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterCountModel{}
	} else {
		*p = RarityParameterCountModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
		if v, ok := d["weight"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Weight)
		}
	}
	return nil
}

func NewRarityParameterCountModelFromJson(data string) RarityParameterCountModel {
	req := RarityParameterCountModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterCountModelFromDict(data map[string]interface{}) RarityParameterCountModel {
	return RarityParameterCountModel{
		Count: func() *int32 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["count"])
		}(),
		Weight: func() *int32 {
			v, ok := data["weight"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["weight"])
		}(),
	}
}

func (p RarityParameterCountModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Count != nil {
		m["count"] = p.Count
	}
	if p.Weight != nil {
		m["weight"] = p.Weight
	}
	return m
}

func (p RarityParameterCountModel) Pointer() *RarityParameterCountModel {
	return &p
}

func CastRarityParameterCountModels(data []interface{}) []RarityParameterCountModel {
	v := make([]RarityParameterCountModel, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterCountModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterCountModelsFromDict(data []RarityParameterCountModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterValueModel struct {
	Name          *string `json:"name"`
	Metadata      *string `json:"metadata"`
	ResourceName  *string `json:"resourceName"`
	ResourceValue *int64  `json:"resourceValue"`
	Weight        *int32  `json:"weight"`
}

func (p *RarityParameterValueModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterValueModel{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterValueModel{}
	} else {
		*p = RarityParameterValueModel{}
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
		if v, ok := d["resourceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceName)
				}
			}
		}
		if v, ok := d["resourceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResourceValue)
		}
		if v, ok := d["weight"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Weight)
		}
	}
	return nil
}

func NewRarityParameterValueModelFromJson(data string) RarityParameterValueModel {
	req := RarityParameterValueModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterValueModelFromDict(data map[string]interface{}) RarityParameterValueModel {
	return RarityParameterValueModel{
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
		ResourceName: func() *string {
			v, ok := data["resourceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceName"])
		}(),
		ResourceValue: func() *int64 {
			v, ok := data["resourceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["resourceValue"])
		}(),
		Weight: func() *int32 {
			v, ok := data["weight"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["weight"])
		}(),
	}
}

func (p RarityParameterValueModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ResourceName != nil {
		m["resourceName"] = p.ResourceName
	}
	if p.ResourceValue != nil {
		m["resourceValue"] = p.ResourceValue
	}
	if p.Weight != nil {
		m["weight"] = p.Weight
	}
	return m
}

func (p RarityParameterValueModel) Pointer() *RarityParameterValueModel {
	return &p
}

func CastRarityParameterValueModels(data []interface{}) []RarityParameterValueModel {
	v := make([]RarityParameterValueModel, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterValueModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterValueModelsFromDict(data []RarityParameterValueModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BalanceParameterValue struct {
	Name  *string `json:"name"`
	Value *int64  `json:"value"`
}

func (p *BalanceParameterValue) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BalanceParameterValue{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = BalanceParameterValue{}
	} else {
		*p = BalanceParameterValue{}
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

func NewBalanceParameterValueFromJson(data string) BalanceParameterValue {
	req := BalanceParameterValue{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBalanceParameterValueFromDict(data map[string]interface{}) BalanceParameterValue {
	return BalanceParameterValue{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Value: func() *int64 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["value"])
		}(),
	}
}

func (p BalanceParameterValue) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p BalanceParameterValue) Pointer() *BalanceParameterValue {
	return &p
}

func CastBalanceParameterValues(data []interface{}) []BalanceParameterValue {
	v := make([]BalanceParameterValue, 0)
	for _, d := range data {
		v = append(v, NewBalanceParameterValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBalanceParameterValuesFromDict(data []BalanceParameterValue) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RarityParameterValue struct {
	Name          *string `json:"name"`
	ResourceName  *string `json:"resourceName"`
	ResourceValue *int64  `json:"resourceValue"`
}

func (p *RarityParameterValue) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RarityParameterValue{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = RarityParameterValue{}
	} else {
		*p = RarityParameterValue{}
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
		if v, ok := d["resourceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceName)
				}
			}
		}
		if v, ok := d["resourceValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResourceValue)
		}
	}
	return nil
}

func NewRarityParameterValueFromJson(data string) RarityParameterValue {
	req := RarityParameterValue{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRarityParameterValueFromDict(data map[string]interface{}) RarityParameterValue {
	return RarityParameterValue{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		ResourceName: func() *string {
			v, ok := data["resourceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceName"])
		}(),
		ResourceValue: func() *int64 {
			v, ok := data["resourceValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["resourceValue"])
		}(),
	}
}

func (p RarityParameterValue) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.ResourceName != nil {
		m["resourceName"] = p.ResourceName
	}
	if p.ResourceValue != nil {
		m["resourceValue"] = p.ResourceValue
	}
	return m
}

func (p RarityParameterValue) Pointer() *RarityParameterValue {
	return &p
}

func CastRarityParameterValues(data []interface{}) []RarityParameterValue {
	v := make([]RarityParameterValue, 0)
	for _, d := range data {
		v = append(v, NewRarityParameterValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRarityParameterValuesFromDict(data []RarityParameterValue) []interface{} {
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
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
	}
}

func (p Config) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
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
	EnableAutoRun                      *bool   `json:"enableAutoRun"`
	EnableAtomicCommit                 *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor          *bool   `json:"transactionUseDistributor"`
	CommitScriptResultInUseDistributor *bool   `json:"commitScriptResultInUseDistributor"`
	AcquireActionUseJobQueue           *bool   `json:"acquireActionUseJobQueue"`
	DistributorNamespaceId             *string `json:"distributorNamespaceId"`
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
		if v, ok := d["commitScriptResultInUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CommitScriptResultInUseDistributor)
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
		CommitScriptResultInUseDistributor: func() *bool {
			v, ok := data["commitScriptResultInUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["commitScriptResultInUseDistributor"])
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
	if p.CommitScriptResultInUseDistributor != nil {
		m["commitScriptResultInUseDistributor"] = p.CommitScriptResultInUseDistributor
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
