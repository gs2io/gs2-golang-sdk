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
		NamespaceId:        core.CastString(data["namespaceId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
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
		transactionSetting = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	var logSetting map[string]interface{}
	if p.LogSetting != nil {
		logSetting = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
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
		"namespaceId":        namespaceId,
		"name":               name,
		"description":        description,
		"transactionSetting": transactionSetting,
		"logSetting":         logSetting,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
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
		BalanceParameterModelId: core.CastString(data["balanceParameterModelId"]),
		Name:                    core.CastString(data["name"]),
		Metadata:                core.CastString(data["metadata"]),
		TotalValue:              core.CastInt64(data["totalValue"]),
		InitialValueStrategy:    core.CastString(data["initialValueStrategy"]),
		Parameters:              CastBalanceParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p BalanceParameterModel) ToDict() map[string]interface{} {

	var balanceParameterModelId *string
	if p.BalanceParameterModelId != nil {
		balanceParameterModelId = p.BalanceParameterModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var totalValue *int64
	if p.TotalValue != nil {
		totalValue = p.TotalValue
	}
	var initialValueStrategy *string
	if p.InitialValueStrategy != nil {
		initialValueStrategy = p.InitialValueStrategy
	}
	var parameters []interface{}
	if p.Parameters != nil {
		parameters = CastBalanceParameterValueModelsFromDict(
			p.Parameters,
		)
	}
	return map[string]interface{}{
		"balanceParameterModelId": balanceParameterModelId,
		"name":                    name,
		"metadata":                metadata,
		"totalValue":              totalValue,
		"initialValueStrategy":    initialValueStrategy,
		"parameters":              parameters,
	}
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
		BalanceParameterModelId: core.CastString(data["balanceParameterModelId"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Metadata:                core.CastString(data["metadata"]),
		TotalValue:              core.CastInt64(data["totalValue"]),
		InitialValueStrategy:    core.CastString(data["initialValueStrategy"]),
		Parameters:              CastBalanceParameterValueModels(core.CastArray(data["parameters"])),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
	}
}

func (p BalanceParameterModelMaster) ToDict() map[string]interface{} {

	var balanceParameterModelId *string
	if p.BalanceParameterModelId != nil {
		balanceParameterModelId = p.BalanceParameterModelId
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
	var totalValue *int64
	if p.TotalValue != nil {
		totalValue = p.TotalValue
	}
	var initialValueStrategy *string
	if p.InitialValueStrategy != nil {
		initialValueStrategy = p.InitialValueStrategy
	}
	var parameters []interface{}
	if p.Parameters != nil {
		parameters = CastBalanceParameterValueModelsFromDict(
			p.Parameters,
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
		"balanceParameterModelId": balanceParameterModelId,
		"name":                    name,
		"description":             description,
		"metadata":                metadata,
		"totalValue":              totalValue,
		"initialValueStrategy":    initialValueStrategy,
		"parameters":              parameters,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
	}
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
		RarityParameterModelId: core.CastString(data["rarityParameterModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		MaximumParameterCount:  core.CastInt32(data["maximumParameterCount"]),
		ParameterCounts:        CastRarityParameterCountModels(core.CastArray(data["parameterCounts"])),
		Parameters:             CastRarityParameterValueModels(core.CastArray(data["parameters"])),
	}
}

func (p RarityParameterModel) ToDict() map[string]interface{} {

	var rarityParameterModelId *string
	if p.RarityParameterModelId != nil {
		rarityParameterModelId = p.RarityParameterModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var maximumParameterCount *int32
	if p.MaximumParameterCount != nil {
		maximumParameterCount = p.MaximumParameterCount
	}
	var parameterCounts []interface{}
	if p.ParameterCounts != nil {
		parameterCounts = CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		)
	}
	var parameters []interface{}
	if p.Parameters != nil {
		parameters = CastRarityParameterValueModelsFromDict(
			p.Parameters,
		)
	}
	return map[string]interface{}{
		"rarityParameterModelId": rarityParameterModelId,
		"name":                   name,
		"metadata":               metadata,
		"maximumParameterCount":  maximumParameterCount,
		"parameterCounts":        parameterCounts,
		"parameters":             parameters,
	}
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
		RarityParameterModelId: core.CastString(data["rarityParameterModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		MaximumParameterCount:  core.CastInt32(data["maximumParameterCount"]),
		ParameterCounts:        CastRarityParameterCountModels(core.CastArray(data["parameterCounts"])),
		Parameters:             CastRarityParameterValueModels(core.CastArray(data["parameters"])),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p RarityParameterModelMaster) ToDict() map[string]interface{} {

	var rarityParameterModelId *string
	if p.RarityParameterModelId != nil {
		rarityParameterModelId = p.RarityParameterModelId
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
	var maximumParameterCount *int32
	if p.MaximumParameterCount != nil {
		maximumParameterCount = p.MaximumParameterCount
	}
	var parameterCounts []interface{}
	if p.ParameterCounts != nil {
		parameterCounts = CastRarityParameterCountModelsFromDict(
			p.ParameterCounts,
		)
	}
	var parameters []interface{}
	if p.Parameters != nil {
		parameters = CastRarityParameterValueModelsFromDict(
			p.Parameters,
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
		"rarityParameterModelId": rarityParameterModelId,
		"name":                   name,
		"description":            description,
		"metadata":               metadata,
		"maximumParameterCount":  maximumParameterCount,
		"parameterCounts":        parameterCounts,
		"parameters":             parameters,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
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
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentParameterMaster) ToDict() map[string]interface{} {

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
		BalanceParameterStatusId: core.CastString(data["balanceParameterStatusId"]),
		UserId:                   core.CastString(data["userId"]),
		ParameterName:            core.CastString(data["parameterName"]),
		PropertyId:               core.CastString(data["propertyId"]),
		ParameterValues:          CastBalanceParameterValues(core.CastArray(data["parameterValues"])),
		CreatedAt:                core.CastInt64(data["createdAt"]),
		UpdatedAt:                core.CastInt64(data["updatedAt"]),
		Revision:                 core.CastInt64(data["revision"]),
	}
}

func (p BalanceParameterStatus) ToDict() map[string]interface{} {

	var balanceParameterStatusId *string
	if p.BalanceParameterStatusId != nil {
		balanceParameterStatusId = p.BalanceParameterStatusId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var parameterName *string
	if p.ParameterName != nil {
		parameterName = p.ParameterName
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var parameterValues []interface{}
	if p.ParameterValues != nil {
		parameterValues = CastBalanceParameterValuesFromDict(
			p.ParameterValues,
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
		"balanceParameterStatusId": balanceParameterStatusId,
		"userId":                   userId,
		"parameterName":            parameterName,
		"propertyId":               propertyId,
		"parameterValues":          parameterValues,
		"createdAt":                createdAt,
		"updatedAt":                updatedAt,
		"revision":                 revision,
	}
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
		RarityParameterStatusId: core.CastString(data["rarityParameterStatusId"]),
		UserId:                  core.CastString(data["userId"]),
		ParameterName:           core.CastString(data["parameterName"]),
		PropertyId:              core.CastString(data["propertyId"]),
		ParameterValues:         CastRarityParameterValues(core.CastArray(data["parameterValues"])),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
	}
}

func (p RarityParameterStatus) ToDict() map[string]interface{} {

	var rarityParameterStatusId *string
	if p.RarityParameterStatusId != nil {
		rarityParameterStatusId = p.RarityParameterStatusId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var parameterName *string
	if p.ParameterName != nil {
		parameterName = p.ParameterName
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var parameterValues []interface{}
	if p.ParameterValues != nil {
		parameterValues = CastRarityParameterValuesFromDict(
			p.ParameterValues,
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
		"rarityParameterStatusId": rarityParameterStatusId,
		"userId":                  userId,
		"parameterName":           parameterName,
		"propertyId":              propertyId,
		"parameterValues":         parameterValues,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
	}
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
		Name:     core.CastString(data["name"]),
		Metadata: core.CastString(data["metadata"]),
	}
}

func (p BalanceParameterValueModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"name":     name,
		"metadata": metadata,
	}
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
		Count:  core.CastInt32(data["count"]),
		Weight: core.CastInt32(data["weight"]),
	}
}

func (p RarityParameterCountModel) ToDict() map[string]interface{} {

	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"count":  count,
		"weight": weight,
	}
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
		Name:          core.CastString(data["name"]),
		Metadata:      core.CastString(data["metadata"]),
		ResourceName:  core.CastString(data["resourceName"]),
		ResourceValue: core.CastInt64(data["resourceValue"]),
		Weight:        core.CastInt32(data["weight"]),
	}
}

func (p RarityParameterValueModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var resourceName *string
	if p.ResourceName != nil {
		resourceName = p.ResourceName
	}
	var resourceValue *int64
	if p.ResourceValue != nil {
		resourceValue = p.ResourceValue
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"name":          name,
		"metadata":      metadata,
		"resourceName":  resourceName,
		"resourceValue": resourceValue,
		"weight":        weight,
	}
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
		Name:  core.CastString(data["name"]),
		Value: core.CastInt64(data["value"]),
	}
}

func (p BalanceParameterValue) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var value *int64
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"name":  name,
		"value": value,
	}
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
		Name:          core.CastString(data["name"]),
		ResourceName:  core.CastString(data["resourceName"]),
		ResourceValue: core.CastInt64(data["resourceValue"]),
	}
}

func (p RarityParameterValue) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var resourceName *string
	if p.ResourceName != nil {
		resourceName = p.ResourceName
	}
	var resourceValue *int64
	if p.ResourceValue != nil {
		resourceValue = p.ResourceValue
	}
	return map[string]interface{}{
		"name":          name,
		"resourceName":  resourceName,
		"resourceValue": resourceValue,
	}
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
