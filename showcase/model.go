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

package showcase

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
	BuyScript          *ScriptSetting      `json:"buyScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
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
		if v, ok := d["buyScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BuyScript)
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
		NamespaceId:        core.CastString(data["namespaceId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		BuyScript:          NewScriptSettingFromDict(core.CastMap(data["buyScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:   core.CastString(data["queueNamespaceId"]),
		KeyId:              core.CastString(data["keyId"]),
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
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var buyScript map[string]interface{}
	if p.BuyScript != nil {
		buyScript = p.BuyScript.ToDict()
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
		"namespaceId":        namespaceId,
		"name":               name,
		"description":        description,
		"transactionSetting": transactionSetting,
		"buyScript":          buyScript,
		"logSetting":         logSetting,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"queueNamespaceId":   queueNamespaceId,
		"keyId":              keyId,
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

type SalesItemMaster struct {
	SalesItemId    *string         `json:"salesItemId"`
	Name           *string         `json:"name"`
	Description    *string         `json:"description"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	CreatedAt      *int64          `json:"createdAt"`
	UpdatedAt      *int64          `json:"updatedAt"`
	Revision       *int64          `json:"revision"`
}

func (p *SalesItemMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SalesItemMaster{}
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
		*p = SalesItemMaster{}
	} else {
		*p = SalesItemMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["salesItemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesItemId)
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
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
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

func NewSalesItemMasterFromJson(data string) SalesItemMaster {
	req := SalesItemMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSalesItemMasterFromDict(data map[string]interface{}) SalesItemMaster {
	return SalesItemMaster{
		SalesItemId:    core.CastString(data["salesItemId"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
		Revision:       core.CastInt64(data["revision"]),
	}
}

func (p SalesItemMaster) ToDict() map[string]interface{} {

	var salesItemId *string
	if p.SalesItemId != nil {
		salesItemId = p.SalesItemId
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
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
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
		"salesItemId":    salesItemId,
		"name":           name,
		"description":    description,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
		"createdAt":      createdAt,
		"updatedAt":      updatedAt,
		"revision":       revision,
	}
}

func (p SalesItemMaster) Pointer() *SalesItemMaster {
	return &p
}

func CastSalesItemMasters(data []interface{}) []SalesItemMaster {
	v := make([]SalesItemMaster, 0)
	for _, d := range data {
		v = append(v, NewSalesItemMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemMastersFromDict(data []SalesItemMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItemGroupMaster struct {
	SalesItemGroupId *string   `json:"salesItemGroupId"`
	Name             *string   `json:"name"`
	Description      *string   `json:"description"`
	Metadata         *string   `json:"metadata"`
	SalesItemNames   []*string `json:"salesItemNames"`
	CreatedAt        *int64    `json:"createdAt"`
	UpdatedAt        *int64    `json:"updatedAt"`
	Revision         *int64    `json:"revision"`
}

func (p *SalesItemGroupMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SalesItemGroupMaster{}
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
		*p = SalesItemGroupMaster{}
	} else {
		*p = SalesItemGroupMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["salesItemGroupId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesItemGroupId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesItemGroupId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesItemGroupId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemGroupId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemGroupId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesItemGroupId)
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
		if v, ok := d["salesItemNames"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SalesItemNames)
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

func NewSalesItemGroupMasterFromJson(data string) SalesItemGroupMaster {
	req := SalesItemGroupMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSalesItemGroupMasterFromDict(data map[string]interface{}) SalesItemGroupMaster {
	return SalesItemGroupMaster{
		SalesItemGroupId: core.CastString(data["salesItemGroupId"]),
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		Metadata:         core.CastString(data["metadata"]),
		SalesItemNames:   core.CastStrings(core.CastArray(data["salesItemNames"])),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
	}
}

func (p SalesItemGroupMaster) ToDict() map[string]interface{} {

	var salesItemGroupId *string
	if p.SalesItemGroupId != nil {
		salesItemGroupId = p.SalesItemGroupId
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
	var salesItemNames []interface{}
	if p.SalesItemNames != nil {
		salesItemNames = core.CastStringsFromDict(
			p.SalesItemNames,
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
		"salesItemGroupId": salesItemGroupId,
		"name":             name,
		"description":      description,
		"metadata":         metadata,
		"salesItemNames":   salesItemNames,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
	}
}

func (p SalesItemGroupMaster) Pointer() *SalesItemGroupMaster {
	return &p
}

func CastSalesItemGroupMasters(data []interface{}) []SalesItemGroupMaster {
	v := make([]SalesItemGroupMaster, 0)
	for _, d := range data {
		v = append(v, NewSalesItemGroupMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemGroupMastersFromDict(data []SalesItemGroupMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseMaster struct {
	ShowcaseId         *string             `json:"showcaseId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	SalesPeriodEventId *string             `json:"salesPeriodEventId"`
	DisplayItems       []DisplayItemMaster `json:"displayItems"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
}

func (p *ShowcaseMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ShowcaseMaster{}
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
		*p = ShowcaseMaster{}
	} else {
		*p = ShowcaseMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["showcaseId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseId)
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
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
				}
			}
		}
		if v, ok := d["displayItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisplayItems)
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

func NewShowcaseMasterFromJson(data string) ShowcaseMaster {
	req := ShowcaseMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewShowcaseMasterFromDict(data map[string]interface{}) ShowcaseMaster {
	return ShowcaseMaster{
		ShowcaseId:         core.CastString(data["showcaseId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		DisplayItems:       CastDisplayItemMasters(core.CastArray(data["displayItems"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p ShowcaseMaster) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
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
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastDisplayItemMastersFromDict(
			p.DisplayItems,
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
		"showcaseId":         showcaseId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"salesPeriodEventId": salesPeriodEventId,
		"displayItems":       displayItems,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p ShowcaseMaster) Pointer() *ShowcaseMaster {
	return &p
}

func CastShowcaseMasters(data []interface{}) []ShowcaseMaster {
	v := make([]ShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseMastersFromDict(data []ShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentShowcaseMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentShowcaseMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentShowcaseMaster{}
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
		*p = CurrentShowcaseMaster{}
	} else {
		*p = CurrentShowcaseMaster{}
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

func NewCurrentShowcaseMasterFromJson(data string) CurrentShowcaseMaster {
	req := CurrentShowcaseMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentShowcaseMasterFromDict(data map[string]interface{}) CurrentShowcaseMaster {
	return CurrentShowcaseMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentShowcaseMaster) ToDict() map[string]interface{} {

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

func (p CurrentShowcaseMaster) Pointer() *CurrentShowcaseMaster {
	return &p
}

func CastCurrentShowcaseMasters(data []interface{}) []CurrentShowcaseMaster {
	v := make([]CurrentShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentShowcaseMastersFromDict(data []CurrentShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItem struct {
	Name           *string         `json:"name"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func (p *SalesItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SalesItem{}
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
		*p = SalesItem{}
	} else {
		*p = SalesItem{}
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
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
	}
	return nil
}

func NewSalesItemFromJson(data string) SalesItem {
	req := SalesItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSalesItemFromDict(data map[string]interface{}) SalesItem {
	return SalesItem{
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p SalesItem) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"name":           name,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
	}
}

func (p SalesItem) Pointer() *SalesItem {
	return &p
}

func CastSalesItems(data []interface{}) []SalesItem {
	v := make([]SalesItem, 0)
	for _, d := range data {
		v = append(v, NewSalesItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemsFromDict(data []SalesItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItemGroup struct {
	Name       *string     `json:"name"`
	Metadata   *string     `json:"metadata"`
	SalesItems []SalesItem `json:"salesItems"`
}

func (p *SalesItemGroup) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SalesItemGroup{}
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
		*p = SalesItemGroup{}
	} else {
		*p = SalesItemGroup{}
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
		if v, ok := d["salesItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SalesItems)
		}
	}
	return nil
}

func NewSalesItemGroupFromJson(data string) SalesItemGroup {
	req := SalesItemGroup{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSalesItemGroupFromDict(data map[string]interface{}) SalesItemGroup {
	return SalesItemGroup{
		Name:       core.CastString(data["name"]),
		Metadata:   core.CastString(data["metadata"]),
		SalesItems: CastSalesItems(core.CastArray(data["salesItems"])),
	}
}

func (p SalesItemGroup) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var salesItems []interface{}
	if p.SalesItems != nil {
		salesItems = CastSalesItemsFromDict(
			p.SalesItems,
		)
	}
	return map[string]interface{}{
		"name":       name,
		"metadata":   metadata,
		"salesItems": salesItems,
	}
}

func (p SalesItemGroup) Pointer() *SalesItemGroup {
	return &p
}

func CastSalesItemGroups(data []interface{}) []SalesItemGroup {
	v := make([]SalesItemGroup, 0)
	for _, d := range data {
		v = append(v, NewSalesItemGroupFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemGroupsFromDict(data []SalesItemGroup) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Showcase struct {
	ShowcaseId         *string       `json:"showcaseId"`
	Name               *string       `json:"name"`
	Metadata           *string       `json:"metadata"`
	SalesPeriodEventId *string       `json:"salesPeriodEventId"`
	DisplayItems       []DisplayItem `json:"displayItems"`
}

func (p *Showcase) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Showcase{}
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
		*p = Showcase{}
	} else {
		*p = Showcase{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["showcaseId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseId)
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
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
				}
			}
		}
		if v, ok := d["displayItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisplayItems)
		}
	}
	return nil
}

func NewShowcaseFromJson(data string) Showcase {
	req := Showcase{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewShowcaseFromDict(data map[string]interface{}) Showcase {
	return Showcase{
		ShowcaseId:         core.CastString(data["showcaseId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		DisplayItems:       CastDisplayItems(core.CastArray(data["displayItems"])),
	}
}

func (p Showcase) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastDisplayItemsFromDict(
			p.DisplayItems,
		)
	}
	return map[string]interface{}{
		"showcaseId":         showcaseId,
		"name":               name,
		"metadata":           metadata,
		"salesPeriodEventId": salesPeriodEventId,
		"displayItems":       displayItems,
	}
}

func (p Showcase) Pointer() *Showcase {
	return &p
}

func CastShowcases(data []interface{}) []Showcase {
	v := make([]Showcase, 0)
	for _, d := range data {
		v = append(v, NewShowcaseFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcasesFromDict(data []Showcase) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DisplayItem struct {
	DisplayItemId      *string         `json:"displayItemId"`
	Type               *string         `json:"type"`
	SalesItem          *SalesItem      `json:"salesItem"`
	SalesItemGroup     *SalesItemGroup `json:"salesItemGroup"`
	SalesPeriodEventId *string         `json:"salesPeriodEventId"`
}

func (p *DisplayItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DisplayItem{}
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
		*p = DisplayItem{}
	} else {
		*p = DisplayItem{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["displayItemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayItemId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["salesItem"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SalesItem)
		}
		if v, ok := d["salesItemGroup"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SalesItemGroup)
		}
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewDisplayItemFromJson(data string) DisplayItem {
	req := DisplayItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDisplayItemFromDict(data map[string]interface{}) DisplayItem {
	return DisplayItem{
		DisplayItemId:      core.CastString(data["displayItemId"]),
		Type:               core.CastString(data["type"]),
		SalesItem:          NewSalesItemFromDict(core.CastMap(data["salesItem"])).Pointer(),
		SalesItemGroup:     NewSalesItemGroupFromDict(core.CastMap(data["salesItemGroup"])).Pointer(),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
	}
}

func (p DisplayItem) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var salesItem map[string]interface{}
	if p.SalesItem != nil {
		salesItem = p.SalesItem.ToDict()
	}
	var salesItemGroup map[string]interface{}
	if p.SalesItemGroup != nil {
		salesItemGroup = p.SalesItemGroup.ToDict()
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	return map[string]interface{}{
		"displayItemId":      displayItemId,
		"type":               _type,
		"salesItem":          salesItem,
		"salesItemGroup":     salesItemGroup,
		"salesPeriodEventId": salesPeriodEventId,
	}
}

func (p DisplayItem) Pointer() *DisplayItem {
	return &p
}

func CastDisplayItems(data []interface{}) []DisplayItem {
	v := make([]DisplayItem, 0)
	for _, d := range data {
		v = append(v, NewDisplayItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDisplayItemsFromDict(data []DisplayItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DisplayItemMaster struct {
	DisplayItemId      *string `json:"displayItemId"`
	Type               *string `json:"type"`
	SalesItemName      *string `json:"salesItemName"`
	SalesItemGroupName *string `json:"salesItemGroupName"`
	SalesPeriodEventId *string `json:"salesPeriodEventId"`
	Revision           *int64  `json:"revision"`
}

func (p *DisplayItemMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DisplayItemMaster{}
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
		*p = DisplayItemMaster{}
	} else {
		*p = DisplayItemMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["displayItemId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayItemId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayItemId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayItemId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayItemId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayItemId)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["salesItemName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesItemName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesItemName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesItemName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesItemName)
				}
			}
		}
		if v, ok := d["salesItemGroupName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesItemGroupName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesItemGroupName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesItemGroupName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemGroupName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesItemGroupName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesItemGroupName)
				}
			}
		}
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
				}
			}
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewDisplayItemMasterFromJson(data string) DisplayItemMaster {
	req := DisplayItemMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDisplayItemMasterFromDict(data map[string]interface{}) DisplayItemMaster {
	return DisplayItemMaster{
		DisplayItemId:      core.CastString(data["displayItemId"]),
		Type:               core.CastString(data["type"]),
		SalesItemName:      core.CastString(data["salesItemName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p DisplayItemMaster) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var salesItemName *string
	if p.SalesItemName != nil {
		salesItemName = p.SalesItemName
	}
	var salesItemGroupName *string
	if p.SalesItemGroupName != nil {
		salesItemGroupName = p.SalesItemGroupName
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"displayItemId":      displayItemId,
		"type":               _type,
		"salesItemName":      salesItemName,
		"salesItemGroupName": salesItemGroupName,
		"salesPeriodEventId": salesPeriodEventId,
		"revision":           revision,
	}
}

func (p DisplayItemMaster) Pointer() *DisplayItemMaster {
	return &p
}

func CastDisplayItemMasters(data []interface{}) []DisplayItemMaster {
	v := make([]DisplayItemMaster, 0)
	for _, d := range data {
		v = append(v, NewDisplayItemMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDisplayItemMastersFromDict(data []DisplayItemMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomShowcaseMaster struct {
	ShowcaseId            *string                  `json:"showcaseId"`
	Name                  *string                  `json:"name"`
	Description           *string                  `json:"description"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
	CreatedAt             *int64                   `json:"createdAt"`
	UpdatedAt             *int64                   `json:"updatedAt"`
	Revision              *int64                   `json:"revision"`
}

func (p *RandomShowcaseMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomShowcaseMaster{}
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
		*p = RandomShowcaseMaster{}
	} else {
		*p = RandomShowcaseMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["showcaseId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseId)
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
		if v, ok := d["maximumNumberOfChoice"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumNumberOfChoice)
		}
		if v, ok := d["displayItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisplayItems)
		}
		if v, ok := d["baseTimestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BaseTimestamp)
		}
		if v, ok := d["resetIntervalHours"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetIntervalHours)
		}
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
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

func NewRandomShowcaseMasterFromJson(data string) RandomShowcaseMaster {
	req := RandomShowcaseMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomShowcaseMasterFromDict(data map[string]interface{}) RandomShowcaseMaster {
	return RandomShowcaseMaster{
		ShowcaseId:            core.CastString(data["showcaseId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p RandomShowcaseMaster) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
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
	var maximumNumberOfChoice *int32
	if p.MaximumNumberOfChoice != nil {
		maximumNumberOfChoice = p.MaximumNumberOfChoice
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		)
	}
	var baseTimestamp *int64
	if p.BaseTimestamp != nil {
		baseTimestamp = p.BaseTimestamp
	}
	var resetIntervalHours *int32
	if p.ResetIntervalHours != nil {
		resetIntervalHours = p.ResetIntervalHours
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
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
		"showcaseId":            showcaseId,
		"name":                  name,
		"description":           description,
		"metadata":              metadata,
		"maximumNumberOfChoice": maximumNumberOfChoice,
		"displayItems":          displayItems,
		"baseTimestamp":         baseTimestamp,
		"resetIntervalHours":    resetIntervalHours,
		"salesPeriodEventId":    salesPeriodEventId,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
	}
}

func (p RandomShowcaseMaster) Pointer() *RandomShowcaseMaster {
	return &p
}

func CastRandomShowcaseMasters(data []interface{}) []RandomShowcaseMaster {
	v := make([]RandomShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewRandomShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomShowcaseMastersFromDict(data []RandomShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomShowcase struct {
	RandomShowcaseId      *string                  `json:"randomShowcaseId"`
	Name                  *string                  `json:"name"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
}

func (p *RandomShowcase) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomShowcase{}
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
		*p = RandomShowcase{}
	} else {
		*p = RandomShowcase{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["randomShowcaseId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RandomShowcaseId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RandomShowcaseId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RandomShowcaseId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RandomShowcaseId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RandomShowcaseId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RandomShowcaseId)
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
		if v, ok := d["maximumNumberOfChoice"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumNumberOfChoice)
		}
		if v, ok := d["displayItems"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DisplayItems)
		}
		if v, ok := d["baseTimestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.BaseTimestamp)
		}
		if v, ok := d["resetIntervalHours"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ResetIntervalHours)
		}
		if v, ok := d["salesPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SalesPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SalesPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SalesPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SalesPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SalesPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewRandomShowcaseFromJson(data string) RandomShowcase {
	req := RandomShowcase{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomShowcaseFromDict(data map[string]interface{}) RandomShowcase {
	return RandomShowcase{
		RandomShowcaseId:      core.CastString(data["randomShowcaseId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
	}
}

func (p RandomShowcase) ToDict() map[string]interface{} {

	var randomShowcaseId *string
	if p.RandomShowcaseId != nil {
		randomShowcaseId = p.RandomShowcaseId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var maximumNumberOfChoice *int32
	if p.MaximumNumberOfChoice != nil {
		maximumNumberOfChoice = p.MaximumNumberOfChoice
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		)
	}
	var baseTimestamp *int64
	if p.BaseTimestamp != nil {
		baseTimestamp = p.BaseTimestamp
	}
	var resetIntervalHours *int32
	if p.ResetIntervalHours != nil {
		resetIntervalHours = p.ResetIntervalHours
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	return map[string]interface{}{
		"randomShowcaseId":      randomShowcaseId,
		"name":                  name,
		"metadata":              metadata,
		"maximumNumberOfChoice": maximumNumberOfChoice,
		"displayItems":          displayItems,
		"baseTimestamp":         baseTimestamp,
		"resetIntervalHours":    resetIntervalHours,
		"salesPeriodEventId":    salesPeriodEventId,
	}
}

func (p RandomShowcase) Pointer() *RandomShowcase {
	return &p
}

func CastRandomShowcases(data []interface{}) []RandomShowcase {
	v := make([]RandomShowcase, 0)
	for _, d := range data {
		v = append(v, NewRandomShowcaseFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomShowcasesFromDict(data []RandomShowcase) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PurchaseCount struct {
	Name  *string `json:"name"`
	Count *int32  `json:"count"`
}

func (p *PurchaseCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PurchaseCount{}
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
		*p = PurchaseCount{}
	} else {
		*p = PurchaseCount{}
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
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewPurchaseCountFromJson(data string) PurchaseCount {
	req := PurchaseCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPurchaseCountFromDict(data map[string]interface{}) PurchaseCount {
	return PurchaseCount{
		Name:  core.CastString(data["name"]),
		Count: core.CastInt32(data["count"]),
	}
}

func (p PurchaseCount) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"name":  name,
		"count": count,
	}
}

func (p PurchaseCount) Pointer() *PurchaseCount {
	return &p
}

func CastPurchaseCounts(data []interface{}) []PurchaseCount {
	v := make([]PurchaseCount, 0)
	for _, d := range data {
		v = append(v, NewPurchaseCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPurchaseCountsFromDict(data []PurchaseCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomDisplayItem struct {
	ShowcaseName         *string         `json:"showcaseName"`
	Name                 *string         `json:"name"`
	Metadata             *string         `json:"metadata"`
	ConsumeActions       []ConsumeAction `json:"consumeActions"`
	AcquireActions       []AcquireAction `json:"acquireActions"`
	CurrentPurchaseCount *int32          `json:"currentPurchaseCount"`
	MaximumPurchaseCount *int32          `json:"maximumPurchaseCount"`
}

func (p *RandomDisplayItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomDisplayItem{}
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
		*p = RandomDisplayItem{}
	} else {
		*p = RandomDisplayItem{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["showcaseName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ShowcaseName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ShowcaseName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ShowcaseName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ShowcaseName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ShowcaseName)
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
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
		if v, ok := d["currentPurchaseCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentPurchaseCount)
		}
		if v, ok := d["maximumPurchaseCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumPurchaseCount)
		}
	}
	return nil
}

func NewRandomDisplayItemFromJson(data string) RandomDisplayItem {
	req := RandomDisplayItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomDisplayItemFromDict(data map[string]interface{}) RandomDisplayItem {
	return RandomDisplayItem{
		ShowcaseName:         core.CastString(data["showcaseName"]),
		Name:                 core.CastString(data["name"]),
		Metadata:             core.CastString(data["metadata"]),
		ConsumeActions:       CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions:       CastAcquireActions(core.CastArray(data["acquireActions"])),
		CurrentPurchaseCount: core.CastInt32(data["currentPurchaseCount"]),
		MaximumPurchaseCount: core.CastInt32(data["maximumPurchaseCount"]),
	}
}

func (p RandomDisplayItem) ToDict() map[string]interface{} {

	var showcaseName *string
	if p.ShowcaseName != nil {
		showcaseName = p.ShowcaseName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var currentPurchaseCount *int32
	if p.CurrentPurchaseCount != nil {
		currentPurchaseCount = p.CurrentPurchaseCount
	}
	var maximumPurchaseCount *int32
	if p.MaximumPurchaseCount != nil {
		maximumPurchaseCount = p.MaximumPurchaseCount
	}
	return map[string]interface{}{
		"showcaseName":         showcaseName,
		"name":                 name,
		"metadata":             metadata,
		"consumeActions":       consumeActions,
		"acquireActions":       acquireActions,
		"currentPurchaseCount": currentPurchaseCount,
		"maximumPurchaseCount": maximumPurchaseCount,
	}
}

func (p RandomDisplayItem) Pointer() *RandomDisplayItem {
	return &p
}

func CastRandomDisplayItems(data []interface{}) []RandomDisplayItem {
	v := make([]RandomDisplayItem, 0)
	for _, d := range data {
		v = append(v, NewRandomDisplayItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomDisplayItemsFromDict(data []RandomDisplayItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomDisplayItemModel struct {
	Name           *string         `json:"name"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	Stock          *int32          `json:"stock"`
	Weight         *int32          `json:"weight"`
}

func (p *RandomDisplayItemModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomDisplayItemModel{}
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
		*p = RandomDisplayItemModel{}
	} else {
		*p = RandomDisplayItemModel{}
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
		if v, ok := d["consumeActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeActions)
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
		if v, ok := d["stock"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Stock)
		}
		if v, ok := d["weight"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Weight)
		}
	}
	return nil
}

func NewRandomDisplayItemModelFromJson(data string) RandomDisplayItemModel {
	req := RandomDisplayItemModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomDisplayItemModelFromDict(data map[string]interface{}) RandomDisplayItemModel {
	return RandomDisplayItemModel{
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
		Stock:          core.CastInt32(data["stock"]),
		Weight:         core.CastInt32(data["weight"]),
	}
}

func (p RandomDisplayItemModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var stock *int32
	if p.Stock != nil {
		stock = p.Stock
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"name":           name,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
		"stock":          stock,
		"weight":         weight,
	}
}

func (p RandomDisplayItemModel) Pointer() *RandomDisplayItemModel {
	return &p
}

func CastRandomDisplayItemModels(data []interface{}) []RandomDisplayItemModel {
	v := make([]RandomDisplayItemModel, 0)
	for _, d := range data {
		v = append(v, NewRandomDisplayItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomDisplayItemModelsFromDict(data []RandomDisplayItemModel) []interface{} {
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
