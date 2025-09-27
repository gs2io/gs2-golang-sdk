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

package lottery

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId            *string             `json:"namespaceId"`
	Name                   *string             `json:"name"`
	Description            *string             `json:"description"`
	TransactionSetting     *TransactionSetting `json:"transactionSetting"`
	LotteryTriggerScriptId *string             `json:"lotteryTriggerScriptId"`
	LogSetting             *LogSetting         `json:"logSetting"`
	CreatedAt              *int64              `json:"createdAt"`
	UpdatedAt              *int64              `json:"updatedAt"`
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
		if v, ok := d["lotteryTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LotteryTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LotteryTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LotteryTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LotteryTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LotteryTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LotteryTriggerScriptId)
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
		LotteryTriggerScriptId: func() *string {
			v, ok := data["lotteryTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lotteryTriggerScriptId"])
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
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
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
	if p.LotteryTriggerScriptId != nil {
		m["lotteryTriggerScriptId"] = p.LotteryTriggerScriptId
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
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
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

type LotteryModelMaster struct {
	LotteryModelId           *string `json:"lotteryModelId"`
	Name                     *string `json:"name"`
	Metadata                 *string `json:"metadata"`
	Description              *string `json:"description"`
	Mode                     *string `json:"mode"`
	Method                   *string `json:"method"`
	PrizeTableName           *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
	CreatedAt                *int64  `json:"createdAt"`
	UpdatedAt                *int64  `json:"updatedAt"`
	Revision                 *int64  `json:"revision"`
}

func (p *LotteryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LotteryModelMaster{}
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
		*p = LotteryModelMaster{}
	} else {
		*p = LotteryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["lotteryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LotteryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LotteryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LotteryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LotteryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LotteryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LotteryModelId)
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
		if v, ok := d["mode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Mode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Mode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Mode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Mode)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
				}
			}
		}
		if v, ok := d["prizeTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableName)
				}
			}
		}
		if v, ok := d["choicePrizeTableScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChoicePrizeTableScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChoicePrizeTableScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChoicePrizeTableScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChoicePrizeTableScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChoicePrizeTableScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChoicePrizeTableScriptId)
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

func NewLotteryModelMasterFromJson(data string) LotteryModelMaster {
	req := LotteryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLotteryModelMasterFromDict(data map[string]interface{}) LotteryModelMaster {
	return LotteryModelMaster{
		LotteryModelId: func() *string {
			v, ok := data["lotteryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lotteryModelId"])
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
		Mode: func() *string {
			v, ok := data["mode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["mode"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		PrizeTableName: func() *string {
			v, ok := data["prizeTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableName"])
		}(),
		ChoicePrizeTableScriptId: func() *string {
			v, ok := data["choicePrizeTableScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["choicePrizeTableScriptId"])
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

func (p LotteryModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LotteryModelId != nil {
		m["lotteryModelId"] = p.LotteryModelId
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
	if p.Mode != nil {
		m["mode"] = p.Mode
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.PrizeTableName != nil {
		m["prizeTableName"] = p.PrizeTableName
	}
	if p.ChoicePrizeTableScriptId != nil {
		m["choicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
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

func (p LotteryModelMaster) Pointer() *LotteryModelMaster {
	return &p
}

func CastLotteryModelMasters(data []interface{}) []LotteryModelMaster {
	v := make([]LotteryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewLotteryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryModelMastersFromDict(data []LotteryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PrizeTableMaster struct {
	PrizeTableId *string `json:"prizeTableId"`
	Name         *string `json:"name"`
	Metadata     *string `json:"metadata"`
	Description  *string `json:"description"`
	Prizes       []Prize `json:"prizes"`
	CreatedAt    *int64  `json:"createdAt"`
	UpdatedAt    *int64  `json:"updatedAt"`
	Revision     *int64  `json:"revision"`
}

func (p *PrizeTableMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrizeTableMaster{}
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
		*p = PrizeTableMaster{}
	} else {
		*p = PrizeTableMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeTableId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableId)
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
		if v, ok := d["prizes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Prizes)
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

func NewPrizeTableMasterFromJson(data string) PrizeTableMaster {
	req := PrizeTableMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPrizeTableMasterFromDict(data map[string]interface{}) PrizeTableMaster {
	return PrizeTableMaster{
		PrizeTableId: func() *string {
			v, ok := data["prizeTableId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableId"])
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
		Prizes: func() []Prize {
			if data["prizes"] == nil {
				return nil
			}
			return CastPrizes(core.CastArray(data["prizes"]))
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

func (p PrizeTableMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeTableId != nil {
		m["prizeTableId"] = p.PrizeTableId
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
	if p.Prizes != nil {
		m["prizes"] = CastPrizesFromDict(
			p.Prizes,
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

func (p PrizeTableMaster) Pointer() *PrizeTableMaster {
	return &p
}

func CastPrizeTableMasters(data []interface{}) []PrizeTableMaster {
	v := make([]PrizeTableMaster, 0)
	for _, d := range data {
		v = append(v, NewPrizeTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizeTableMastersFromDict(data []PrizeTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LotteryModel struct {
	LotteryModelId           *string `json:"lotteryModelId"`
	Name                     *string `json:"name"`
	Metadata                 *string `json:"metadata"`
	Mode                     *string `json:"mode"`
	Method                   *string `json:"method"`
	PrizeTableName           *string `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string `json:"choicePrizeTableScriptId"`
}

func (p *LotteryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LotteryModel{}
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
		*p = LotteryModel{}
	} else {
		*p = LotteryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["lotteryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LotteryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LotteryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LotteryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LotteryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LotteryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LotteryModelId)
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
		if v, ok := d["mode"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Mode = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Mode = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Mode = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Mode = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Mode)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
				}
			}
		}
		if v, ok := d["prizeTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableName)
				}
			}
		}
		if v, ok := d["choicePrizeTableScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ChoicePrizeTableScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ChoicePrizeTableScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ChoicePrizeTableScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ChoicePrizeTableScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ChoicePrizeTableScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ChoicePrizeTableScriptId)
				}
			}
		}
	}
	return nil
}

func NewLotteryModelFromJson(data string) LotteryModel {
	req := LotteryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLotteryModelFromDict(data map[string]interface{}) LotteryModel {
	return LotteryModel{
		LotteryModelId: func() *string {
			v, ok := data["lotteryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lotteryModelId"])
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
		Mode: func() *string {
			v, ok := data["mode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["mode"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		PrizeTableName: func() *string {
			v, ok := data["prizeTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableName"])
		}(),
		ChoicePrizeTableScriptId: func() *string {
			v, ok := data["choicePrizeTableScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["choicePrizeTableScriptId"])
		}(),
	}
}

func (p LotteryModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LotteryModelId != nil {
		m["lotteryModelId"] = p.LotteryModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.Mode != nil {
		m["mode"] = p.Mode
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.PrizeTableName != nil {
		m["prizeTableName"] = p.PrizeTableName
	}
	if p.ChoicePrizeTableScriptId != nil {
		m["choicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
	}
	return m
}

func (p LotteryModel) Pointer() *LotteryModel {
	return &p
}

func CastLotteryModels(data []interface{}) []LotteryModel {
	v := make([]LotteryModel, 0)
	for _, d := range data {
		v = append(v, NewLotteryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLotteryModelsFromDict(data []LotteryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PrizeTable struct {
	PrizeTableId *string `json:"prizeTableId"`
	Name         *string `json:"name"`
	Metadata     *string `json:"metadata"`
	Prizes       []Prize `json:"prizes"`
}

func (p *PrizeTable) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrizeTable{}
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
		*p = PrizeTable{}
	} else {
		*p = PrizeTable{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeTableId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableId)
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
		if v, ok := d["prizes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Prizes)
		}
	}
	return nil
}

func NewPrizeTableFromJson(data string) PrizeTable {
	req := PrizeTable{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPrizeTableFromDict(data map[string]interface{}) PrizeTable {
	return PrizeTable{
		PrizeTableId: func() *string {
			v, ok := data["prizeTableId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableId"])
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
		Prizes: func() []Prize {
			if data["prizes"] == nil {
				return nil
			}
			return CastPrizes(core.CastArray(data["prizes"]))
		}(),
	}
}

func (p PrizeTable) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeTableId != nil {
		m["prizeTableId"] = p.PrizeTableId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.Prizes != nil {
		m["prizes"] = CastPrizesFromDict(
			p.Prizes,
		)
	}
	return m
}

func (p PrizeTable) Pointer() *PrizeTable {
	return &p
}

func CastPrizeTables(data []interface{}) []PrizeTable {
	v := make([]PrizeTable, 0)
	for _, d := range data {
		v = append(v, NewPrizeTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizeTablesFromDict(data []PrizeTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Probability struct {
	Prize *DrawnPrize `json:"prize"`
	Rate  *float32    `json:"rate"`
}

func (p *Probability) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Probability{}
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
		*p = Probability{}
	} else {
		*p = Probability{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prize"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Prize)
		}
		if v, ok := d["rate"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rate)
		}
	}
	return nil
}

func NewProbabilityFromJson(data string) Probability {
	req := Probability{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewProbabilityFromDict(data map[string]interface{}) Probability {
	return Probability{
		Prize: func() *DrawnPrize {
			v, ok := data["prize"]
			if !ok || v == nil {
				return nil
			}
			return NewDrawnPrizeFromDict(core.CastMap(data["prize"])).Pointer()
		}(),
		Rate: func() *float32 {
			v, ok := data["rate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["rate"])
		}(),
	}
}

func (p Probability) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Prize != nil {
		m["prize"] = func() map[string]interface{} {
			if p.Prize == nil {
				return nil
			}
			return p.Prize.ToDict()
		}()
	}
	if p.Rate != nil {
		m["rate"] = p.Rate
	}
	return m
}

func (p Probability) Pointer() *Probability {
	return &p
}

func CastProbabilities(data []interface{}) []Probability {
	v := make([]Probability, 0)
	for _, d := range data {
		v = append(v, NewProbabilityFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProbabilitiesFromDict(data []Probability) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentLotteryMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentLotteryMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentLotteryMaster{}
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
		*p = CurrentLotteryMaster{}
	} else {
		*p = CurrentLotteryMaster{}
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

func NewCurrentLotteryMasterFromJson(data string) CurrentLotteryMaster {
	req := CurrentLotteryMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentLotteryMasterFromDict(data map[string]interface{}) CurrentLotteryMaster {
	return CurrentLotteryMaster{
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

func (p CurrentLotteryMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentLotteryMaster) Pointer() *CurrentLotteryMaster {
	return &p
}

func CastCurrentLotteryMasters(data []interface{}) []CurrentLotteryMaster {
	v := make([]CurrentLotteryMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentLotteryMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentLotteryMastersFromDict(data []CurrentLotteryMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Prize struct {
	PrizeId              *string         `json:"prizeId"`
	Type                 *string         `json:"type"`
	AcquireActions       []AcquireAction `json:"acquireActions"`
	DrawnLimit           *int32          `json:"drawnLimit"`
	LimitFailOverPrizeId *string         `json:"limitFailOverPrizeId"`
	PrizeTableName       *string         `json:"prizeTableName"`
	Weight               *int32          `json:"weight"`
}

func (p *Prize) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Prize{}
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
		*p = Prize{}
	} else {
		*p = Prize{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeId)
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
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
		if v, ok := d["drawnLimit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DrawnLimit)
		}
		if v, ok := d["limitFailOverPrizeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LimitFailOverPrizeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LimitFailOverPrizeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LimitFailOverPrizeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LimitFailOverPrizeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LimitFailOverPrizeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LimitFailOverPrizeId)
				}
			}
		}
		if v, ok := d["prizeTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableName)
				}
			}
		}
		if v, ok := d["weight"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Weight)
		}
	}
	return nil
}

func NewPrizeFromJson(data string) Prize {
	req := Prize{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPrizeFromDict(data map[string]interface{}) Prize {
	return Prize{
		PrizeId: func() *string {
			v, ok := data["prizeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeId"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
		}(),
		DrawnLimit: func() *int32 {
			v, ok := data["drawnLimit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["drawnLimit"])
		}(),
		LimitFailOverPrizeId: func() *string {
			v, ok := data["limitFailOverPrizeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["limitFailOverPrizeId"])
		}(),
		PrizeTableName: func() *string {
			v, ok := data["prizeTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableName"])
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

func (p Prize) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeId != nil {
		m["prizeId"] = p.PrizeId
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.AcquireActions != nil {
		m["acquireActions"] = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	if p.DrawnLimit != nil {
		m["drawnLimit"] = p.DrawnLimit
	}
	if p.LimitFailOverPrizeId != nil {
		m["limitFailOverPrizeId"] = p.LimitFailOverPrizeId
	}
	if p.PrizeTableName != nil {
		m["prizeTableName"] = p.PrizeTableName
	}
	if p.Weight != nil {
		m["weight"] = p.Weight
	}
	return m
}

func (p Prize) Pointer() *Prize {
	return &p
}

func CastPrizes(data []interface{}) []Prize {
	v := make([]Prize, 0)
	for _, d := range data {
		v = append(v, NewPrizeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizesFromDict(data []Prize) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PrizeLimit struct {
	PrizeLimitId *string `json:"prizeLimitId"`
	PrizeId      *string `json:"prizeId"`
	DrawnCount   *int32  `json:"drawnCount"`
	CreatedAt    *int64  `json:"createdAt"`
	UpdatedAt    *int64  `json:"updatedAt"`
	Revision     *int64  `json:"revision"`
}

func (p *PrizeLimit) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PrizeLimit{}
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
		*p = PrizeLimit{}
	} else {
		*p = PrizeLimit{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeLimitId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeLimitId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeLimitId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeLimitId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeLimitId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeLimitId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeLimitId)
				}
			}
		}
		if v, ok := d["prizeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeId)
				}
			}
		}
		if v, ok := d["drawnCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.DrawnCount)
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

func NewPrizeLimitFromJson(data string) PrizeLimit {
	req := PrizeLimit{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPrizeLimitFromDict(data map[string]interface{}) PrizeLimit {
	return PrizeLimit{
		PrizeLimitId: func() *string {
			v, ok := data["prizeLimitId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeLimitId"])
		}(),
		PrizeId: func() *string {
			v, ok := data["prizeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeId"])
		}(),
		DrawnCount: func() *int32 {
			v, ok := data["drawnCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["drawnCount"])
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

func (p PrizeLimit) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeLimitId != nil {
		m["prizeLimitId"] = p.PrizeLimitId
	}
	if p.PrizeId != nil {
		m["prizeId"] = p.PrizeId
	}
	if p.DrawnCount != nil {
		m["drawnCount"] = p.DrawnCount
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

func (p PrizeLimit) Pointer() *PrizeLimit {
	return &p
}

func CastPrizeLimits(data []interface{}) []PrizeLimit {
	v := make([]PrizeLimit, 0)
	for _, d := range data {
		v = append(v, NewPrizeLimitFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPrizeLimitsFromDict(data []PrizeLimit) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DrawnPrize struct {
	PrizeId        *string         `json:"prizeId"`
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func (p *DrawnPrize) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DrawnPrize{}
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
		*p = DrawnPrize{}
	} else {
		*p = DrawnPrize{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeId)
				}
			}
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
	}
	return nil
}

func NewDrawnPrizeFromJson(data string) DrawnPrize {
	req := DrawnPrize{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDrawnPrizeFromDict(data map[string]interface{}) DrawnPrize {
	return DrawnPrize{
		PrizeId: func() *string {
			v, ok := data["prizeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeId"])
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
		}(),
	}
}

func (p DrawnPrize) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeId != nil {
		m["prizeId"] = p.PrizeId
	}
	if p.AcquireActions != nil {
		m["acquireActions"] = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return m
}

func (p DrawnPrize) Pointer() *DrawnPrize {
	return &p
}

func CastDrawnPrizes(data []interface{}) []DrawnPrize {
	v := make([]DrawnPrize, 0)
	for _, d := range data {
		v = append(v, NewDrawnPrizeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDrawnPrizesFromDict(data []DrawnPrize) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BoxItem struct {
	PrizeId        *string         `json:"prizeId"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	Remaining      *int32          `json:"remaining"`
	Initial        *int32          `json:"initial"`
}

func (p *BoxItem) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BoxItem{}
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
		*p = BoxItem{}
	} else {
		*p = BoxItem{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["prizeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeId)
				}
			}
		}
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
		if v, ok := d["remaining"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Remaining)
		}
		if v, ok := d["initial"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Initial)
		}
	}
	return nil
}

func NewBoxItemFromJson(data string) BoxItem {
	req := BoxItem{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBoxItemFromDict(data map[string]interface{}) BoxItem {
	return BoxItem{
		PrizeId: func() *string {
			v, ok := data["prizeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeId"])
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
		}(),
		Remaining: func() *int32 {
			v, ok := data["remaining"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["remaining"])
		}(),
		Initial: func() *int32 {
			v, ok := data["initial"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["initial"])
		}(),
	}
}

func (p BoxItem) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.PrizeId != nil {
		m["prizeId"] = p.PrizeId
	}
	if p.AcquireActions != nil {
		m["acquireActions"] = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	if p.Remaining != nil {
		m["remaining"] = p.Remaining
	}
	if p.Initial != nil {
		m["initial"] = p.Initial
	}
	return m
}

func (p BoxItem) Pointer() *BoxItem {
	return &p
}

func CastBoxItems(data []interface{}) []BoxItem {
	v := make([]BoxItem, 0)
	for _, d := range data {
		v = append(v, NewBoxItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBoxItemsFromDict(data []BoxItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BoxItems struct {
	BoxId          *string   `json:"boxId"`
	PrizeTableName *string   `json:"prizeTableName"`
	UserId         *string   `json:"userId"`
	Items          []BoxItem `json:"items"`
}

func (p *BoxItems) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BoxItems{}
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
		*p = BoxItems{}
	} else {
		*p = BoxItems{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["boxId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BoxId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BoxId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BoxId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BoxId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BoxId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BoxId)
				}
			}
		}
		if v, ok := d["prizeTableName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PrizeTableName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PrizeTableName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PrizeTableName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PrizeTableName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PrizeTableName)
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
		if v, ok := d["items"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Items)
		}
	}
	return nil
}

func NewBoxItemsFromJson(data string) BoxItems {
	req := BoxItems{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBoxItemsFromDict(data map[string]interface{}) BoxItems {
	return BoxItems{
		BoxId: func() *string {
			v, ok := data["boxId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["boxId"])
		}(),
		PrizeTableName: func() *string {
			v, ok := data["prizeTableName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["prizeTableName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Items: func() []BoxItem {
			if data["items"] == nil {
				return nil
			}
			return CastBoxItems(core.CastArray(data["items"]))
		}(),
	}
}

func (p BoxItems) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.BoxId != nil {
		m["boxId"] = p.BoxId
	}
	if p.PrizeTableName != nil {
		m["prizeTableName"] = p.PrizeTableName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Items != nil {
		m["items"] = CastBoxItemsFromDict(
			p.Items,
		)
	}
	return m
}

func (p BoxItems) Pointer() *BoxItems {
	return &p
}

func CastBoxItemses(data []interface{}) []BoxItems {
	v := make([]BoxItems, 0)
	for _, d := range data {
		v = append(v, NewBoxItemsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBoxItemsesFromDict(data []BoxItems) []interface{} {
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
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Request: func() *string {
			v, ok := data["request"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["request"])
		}(),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Request != nil {
		m["request"] = p.Request
	}
	return m
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
