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

package ranking2

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

type GlobalRankingModel struct {
	GlobalRankingModelId *string         `json:"globalRankingModelId"`
	Name                 *string         `json:"name"`
	Metadata             *string         `json:"metadata"`
	MinimumValue         *int64          `json:"minimumValue"`
	MaximumValue         *int64          `json:"maximumValue"`
	Sum                  *bool           `json:"sum"`
	OrderDirection       *string         `json:"orderDirection"`
	EntryPeriodEventId   *string         `json:"entryPeriodEventId"`
	RankingRewards       []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId  *string         `json:"accessPeriodEventId"`
}

func (p *GlobalRankingModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingModel{}
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
		*p = GlobalRankingModel{}
	} else {
		*p = GlobalRankingModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalRankingModelId)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewGlobalRankingModelFromJson(data string) GlobalRankingModel {
	req := GlobalRankingModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingModelFromDict(data map[string]interface{}) GlobalRankingModel {
	return GlobalRankingModel{
		GlobalRankingModelId: core.CastString(data["globalRankingModelId"]),
		Name:                 core.CastString(data["name"]),
		Metadata:             core.CastString(data["metadata"]),
		MinimumValue:         core.CastInt64(data["minimumValue"]),
		MaximumValue:         core.CastInt64(data["maximumValue"]),
		Sum:                  core.CastBool(data["sum"]),
		OrderDirection:       core.CastString(data["orderDirection"]),
		EntryPeriodEventId:   core.CastString(data["entryPeriodEventId"]),
		RankingRewards:       CastRankingRewards(core.CastArray(data["rankingRewards"])),
		AccessPeriodEventId:  core.CastString(data["accessPeriodEventId"]),
	}
}

func (p GlobalRankingModel) ToDict() map[string]interface{} {

	var globalRankingModelId *string
	if p.GlobalRankingModelId != nil {
		globalRankingModelId = p.GlobalRankingModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var rankingRewards []interface{}
	if p.RankingRewards != nil {
		rankingRewards = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
	}
	return map[string]interface{}{
		"globalRankingModelId": globalRankingModelId,
		"name":                 name,
		"metadata":             metadata,
		"minimumValue":         minimumValue,
		"maximumValue":         maximumValue,
		"sum":                  sum,
		"orderDirection":       orderDirection,
		"entryPeriodEventId":   entryPeriodEventId,
		"rankingRewards":       rankingRewards,
		"accessPeriodEventId":  accessPeriodEventId,
	}
}

func (p GlobalRankingModel) Pointer() *GlobalRankingModel {
	return &p
}

func CastGlobalRankingModels(data []interface{}) []GlobalRankingModel {
	v := make([]GlobalRankingModel, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingModelsFromDict(data []GlobalRankingModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalRankingModelMaster struct {
	GlobalRankingModelId *string         `json:"globalRankingModelId"`
	Name                 *string         `json:"name"`
	Description          *string         `json:"description"`
	Metadata             *string         `json:"metadata"`
	MinimumValue         *int64          `json:"minimumValue"`
	MaximumValue         *int64          `json:"maximumValue"`
	Sum                  *bool           `json:"sum"`
	OrderDirection       *string         `json:"orderDirection"`
	EntryPeriodEventId   *string         `json:"entryPeriodEventId"`
	RankingRewards       []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId  *string         `json:"accessPeriodEventId"`
	CreatedAt            *int64          `json:"createdAt"`
	UpdatedAt            *int64          `json:"updatedAt"`
	Revision             *int64          `json:"revision"`
}

func (p *GlobalRankingModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingModelMaster{}
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
		*p = GlobalRankingModelMaster{}
	} else {
		*p = GlobalRankingModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalRankingModelId)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
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

func NewGlobalRankingModelMasterFromJson(data string) GlobalRankingModelMaster {
	req := GlobalRankingModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingModelMasterFromDict(data map[string]interface{}) GlobalRankingModelMaster {
	return GlobalRankingModelMaster{
		GlobalRankingModelId: core.CastString(data["globalRankingModelId"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		Metadata:             core.CastString(data["metadata"]),
		MinimumValue:         core.CastInt64(data["minimumValue"]),
		MaximumValue:         core.CastInt64(data["maximumValue"]),
		Sum:                  core.CastBool(data["sum"]),
		OrderDirection:       core.CastString(data["orderDirection"]),
		EntryPeriodEventId:   core.CastString(data["entryPeriodEventId"]),
		RankingRewards:       CastRankingRewards(core.CastArray(data["rankingRewards"])),
		AccessPeriodEventId:  core.CastString(data["accessPeriodEventId"]),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p GlobalRankingModelMaster) ToDict() map[string]interface{} {

	var globalRankingModelId *string
	if p.GlobalRankingModelId != nil {
		globalRankingModelId = p.GlobalRankingModelId
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
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var rankingRewards []interface{}
	if p.RankingRewards != nil {
		rankingRewards = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
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
		"globalRankingModelId": globalRankingModelId,
		"name":                 name,
		"description":          description,
		"metadata":             metadata,
		"minimumValue":         minimumValue,
		"maximumValue":         maximumValue,
		"sum":                  sum,
		"orderDirection":       orderDirection,
		"entryPeriodEventId":   entryPeriodEventId,
		"rankingRewards":       rankingRewards,
		"accessPeriodEventId":  accessPeriodEventId,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
}

func (p GlobalRankingModelMaster) Pointer() *GlobalRankingModelMaster {
	return &p
}

func CastGlobalRankingModelMasters(data []interface{}) []GlobalRankingModelMaster {
	v := make([]GlobalRankingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingModelMastersFromDict(data []GlobalRankingModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalRankingScore struct {
	GlobalRankingScoreId *string `json:"globalRankingScoreId"`
	RankingName          *string `json:"rankingName"`
	UserId               *string `json:"userId"`
	Season               *int64  `json:"season"`
	Score                *int64  `json:"score"`
	Metadata             *string `json:"metadata"`
	CreatedAt            *int64  `json:"createdAt"`
	UpdatedAt            *int64  `json:"updatedAt"`
	Revision             *int64  `json:"revision"`
}

func (p *GlobalRankingScore) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingScore{}
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
		*p = GlobalRankingScore{}
	} else {
		*p = GlobalRankingScore{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalRankingScoreId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalRankingScoreId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalRankingScoreId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalRankingScoreId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingScoreId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingScoreId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalRankingScoreId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewGlobalRankingScoreFromJson(data string) GlobalRankingScore {
	req := GlobalRankingScore{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingScoreFromDict(data map[string]interface{}) GlobalRankingScore {
	return GlobalRankingScore{
		GlobalRankingScoreId: core.CastString(data["globalRankingScoreId"]),
		RankingName:          core.CastString(data["rankingName"]),
		UserId:               core.CastString(data["userId"]),
		Season:               core.CastInt64(data["season"]),
		Score:                core.CastInt64(data["score"]),
		Metadata:             core.CastString(data["metadata"]),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p GlobalRankingScore) ToDict() map[string]interface{} {

	var globalRankingScoreId *string
	if p.GlobalRankingScoreId != nil {
		globalRankingScoreId = p.GlobalRankingScoreId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"globalRankingScoreId": globalRankingScoreId,
		"rankingName":          rankingName,
		"userId":               userId,
		"season":               season,
		"score":                score,
		"metadata":             metadata,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
}

func (p GlobalRankingScore) Pointer() *GlobalRankingScore {
	return &p
}

func CastGlobalRankingScores(data []interface{}) []GlobalRankingScore {
	v := make([]GlobalRankingScore, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingScoreFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingScoresFromDict(data []GlobalRankingScore) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalRankingReceivedReward struct {
	GlobalRankingReceivedRewardId *string `json:"globalRankingReceivedRewardId"`
	RankingName                   *string `json:"rankingName"`
	UserId                        *string `json:"userId"`
	Season                        *int64  `json:"season"`
	ReceivedAt                    *int64  `json:"receivedAt"`
	Revision                      *int64  `json:"revision"`
}

func (p *GlobalRankingReceivedReward) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingReceivedReward{}
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
		*p = GlobalRankingReceivedReward{}
	} else {
		*p = GlobalRankingReceivedReward{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalRankingReceivedRewardId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalRankingReceivedRewardId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalRankingReceivedRewardId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalRankingReceivedRewardId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingReceivedRewardId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingReceivedRewardId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalRankingReceivedRewardId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
		}
		if v, ok := d["receivedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceivedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewGlobalRankingReceivedRewardFromJson(data string) GlobalRankingReceivedReward {
	req := GlobalRankingReceivedReward{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingReceivedRewardFromDict(data map[string]interface{}) GlobalRankingReceivedReward {
	return GlobalRankingReceivedReward{
		GlobalRankingReceivedRewardId: core.CastString(data["globalRankingReceivedRewardId"]),
		RankingName:                   core.CastString(data["rankingName"]),
		UserId:                        core.CastString(data["userId"]),
		Season:                        core.CastInt64(data["season"]),
		ReceivedAt:                    core.CastInt64(data["receivedAt"]),
		Revision:                      core.CastInt64(data["revision"]),
	}
}

func (p GlobalRankingReceivedReward) ToDict() map[string]interface{} {

	var globalRankingReceivedRewardId *string
	if p.GlobalRankingReceivedRewardId != nil {
		globalRankingReceivedRewardId = p.GlobalRankingReceivedRewardId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var receivedAt *int64
	if p.ReceivedAt != nil {
		receivedAt = p.ReceivedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"globalRankingReceivedRewardId": globalRankingReceivedRewardId,
		"rankingName":                   rankingName,
		"userId":                        userId,
		"season":                        season,
		"receivedAt":                    receivedAt,
		"revision":                      revision,
	}
}

func (p GlobalRankingReceivedReward) Pointer() *GlobalRankingReceivedReward {
	return &p
}

func CastGlobalRankingReceivedRewards(data []interface{}) []GlobalRankingReceivedReward {
	v := make([]GlobalRankingReceivedReward, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingReceivedRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingReceivedRewardsFromDict(data []GlobalRankingReceivedReward) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalRankingData struct {
	GlobalRankingDataId *string `json:"globalRankingDataId"`
	RankingName         *string `json:"rankingName"`
	Season              *int64  `json:"season"`
	UserId              *string `json:"userId"`
	Index               *int32  `json:"index"`
	Rank                *int32  `json:"rank"`
	Score               *int64  `json:"score"`
	Metadata            *string `json:"metadata"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
	Revision            *int64  `json:"revision"`
}

func (p *GlobalRankingData) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingData{}
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
		*p = GlobalRankingData{}
	} else {
		*p = GlobalRankingData{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["globalRankingDataId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GlobalRankingDataId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GlobalRankingDataId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GlobalRankingDataId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingDataId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GlobalRankingDataId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GlobalRankingDataId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["index"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Index)
		}
		if v, ok := d["rank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rank)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewGlobalRankingDataFromJson(data string) GlobalRankingData {
	req := GlobalRankingData{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingDataFromDict(data map[string]interface{}) GlobalRankingData {
	return GlobalRankingData{
		GlobalRankingDataId: core.CastString(data["globalRankingDataId"]),
		RankingName:         core.CastString(data["rankingName"]),
		Season:              core.CastInt64(data["season"]),
		UserId:              core.CastString(data["userId"]),
		Index:               core.CastInt32(data["index"]),
		Rank:                core.CastInt32(data["rank"]),
		Score:               core.CastInt64(data["score"]),
		Metadata:            core.CastString(data["metadata"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		Revision:            core.CastInt64(data["revision"]),
	}
}

func (p GlobalRankingData) ToDict() map[string]interface{} {

	var globalRankingDataId *string
	if p.GlobalRankingDataId != nil {
		globalRankingDataId = p.GlobalRankingDataId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var index *int32
	if p.Index != nil {
		index = p.Index
	}
	var rank *int32
	if p.Rank != nil {
		rank = p.Rank
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"globalRankingDataId": globalRankingDataId,
		"rankingName":         rankingName,
		"season":              season,
		"userId":              userId,
		"index":               index,
		"rank":                rank,
		"score":               score,
		"metadata":            metadata,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"revision":            revision,
	}
}

func (p GlobalRankingData) Pointer() *GlobalRankingData {
	return &p
}

func CastGlobalRankingDatas(data []interface{}) []GlobalRankingData {
	v := make([]GlobalRankingData, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingDataFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingDatasFromDict(data []GlobalRankingData) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ClusterRankingModel struct {
	ClusterRankingModelId *string         `json:"clusterRankingModelId"`
	Name                  *string         `json:"name"`
	Metadata              *string         `json:"metadata"`
	ClusterType           *string         `json:"clusterType"`
	MinimumValue          *int64          `json:"minimumValue"`
	MaximumValue          *int64          `json:"maximumValue"`
	Sum                   *bool           `json:"sum"`
	OrderDirection        *string         `json:"orderDirection"`
	EntryPeriodEventId    *string         `json:"entryPeriodEventId"`
	RankingRewards        []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId   *string         `json:"accessPeriodEventId"`
}

func (p *ClusterRankingModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ClusterRankingModel{}
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
		*p = ClusterRankingModel{}
	} else {
		*p = ClusterRankingModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clusterRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterRankingModelId)
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
		if v, ok := d["clusterType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterType)
				}
			}
		}
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewClusterRankingModelFromJson(data string) ClusterRankingModel {
	req := ClusterRankingModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewClusterRankingModelFromDict(data map[string]interface{}) ClusterRankingModel {
	return ClusterRankingModel{
		ClusterRankingModelId: core.CastString(data["clusterRankingModelId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		ClusterType:           core.CastString(data["clusterType"]),
		MinimumValue:          core.CastInt64(data["minimumValue"]),
		MaximumValue:          core.CastInt64(data["maximumValue"]),
		Sum:                   core.CastBool(data["sum"]),
		OrderDirection:        core.CastString(data["orderDirection"]),
		EntryPeriodEventId:    core.CastString(data["entryPeriodEventId"]),
		RankingRewards:        CastRankingRewards(core.CastArray(data["rankingRewards"])),
		AccessPeriodEventId:   core.CastString(data["accessPeriodEventId"]),
	}
}

func (p ClusterRankingModel) ToDict() map[string]interface{} {

	var clusterRankingModelId *string
	if p.ClusterRankingModelId != nil {
		clusterRankingModelId = p.ClusterRankingModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var clusterType *string
	if p.ClusterType != nil {
		clusterType = p.ClusterType
	}
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var rankingRewards []interface{}
	if p.RankingRewards != nil {
		rankingRewards = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
	}
	return map[string]interface{}{
		"clusterRankingModelId": clusterRankingModelId,
		"name":                  name,
		"metadata":              metadata,
		"clusterType":           clusterType,
		"minimumValue":          minimumValue,
		"maximumValue":          maximumValue,
		"sum":                   sum,
		"orderDirection":        orderDirection,
		"entryPeriodEventId":    entryPeriodEventId,
		"rankingRewards":        rankingRewards,
		"accessPeriodEventId":   accessPeriodEventId,
	}
}

func (p ClusterRankingModel) Pointer() *ClusterRankingModel {
	return &p
}

func CastClusterRankingModels(data []interface{}) []ClusterRankingModel {
	v := make([]ClusterRankingModel, 0)
	for _, d := range data {
		v = append(v, NewClusterRankingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastClusterRankingModelsFromDict(data []ClusterRankingModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ClusterRankingModelMaster struct {
	ClusterRankingModelId *string         `json:"clusterRankingModelId"`
	Name                  *string         `json:"name"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ClusterType           *string         `json:"clusterType"`
	MinimumValue          *int64          `json:"minimumValue"`
	MaximumValue          *int64          `json:"maximumValue"`
	Sum                   *bool           `json:"sum"`
	ScoreTtlDays          *int32          `json:"scoreTtlDays"`
	OrderDirection        *string         `json:"orderDirection"`
	EntryPeriodEventId    *string         `json:"entryPeriodEventId"`
	RankingRewards        []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId   *string         `json:"accessPeriodEventId"`
	CreatedAt             *int64          `json:"createdAt"`
	UpdatedAt             *int64          `json:"updatedAt"`
	Revision              *int64          `json:"revision"`
}

func (p *ClusterRankingModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ClusterRankingModelMaster{}
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
		*p = ClusterRankingModelMaster{}
	} else {
		*p = ClusterRankingModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clusterRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterRankingModelId)
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
		if v, ok := d["clusterType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterType)
				}
			}
		}
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["rankingRewards"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RankingRewards)
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
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

func NewClusterRankingModelMasterFromJson(data string) ClusterRankingModelMaster {
	req := ClusterRankingModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewClusterRankingModelMasterFromDict(data map[string]interface{}) ClusterRankingModelMaster {
	return ClusterRankingModelMaster{
		ClusterRankingModelId: core.CastString(data["clusterRankingModelId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ClusterType:           core.CastString(data["clusterType"]),
		MinimumValue:          core.CastInt64(data["minimumValue"]),
		MaximumValue:          core.CastInt64(data["maximumValue"]),
		Sum:                   core.CastBool(data["sum"]),
		ScoreTtlDays:          core.CastInt32(data["scoreTtlDays"]),
		OrderDirection:        core.CastString(data["orderDirection"]),
		EntryPeriodEventId:    core.CastString(data["entryPeriodEventId"]),
		RankingRewards:        CastRankingRewards(core.CastArray(data["rankingRewards"])),
		AccessPeriodEventId:   core.CastString(data["accessPeriodEventId"]),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p ClusterRankingModelMaster) ToDict() map[string]interface{} {

	var clusterRankingModelId *string
	if p.ClusterRankingModelId != nil {
		clusterRankingModelId = p.ClusterRankingModelId
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
	var clusterType *string
	if p.ClusterType != nil {
		clusterType = p.ClusterType
	}
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var scoreTtlDays *int32
	if p.ScoreTtlDays != nil {
		scoreTtlDays = p.ScoreTtlDays
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var rankingRewards []interface{}
	if p.RankingRewards != nil {
		rankingRewards = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
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
		"clusterRankingModelId": clusterRankingModelId,
		"name":                  name,
		"description":           description,
		"metadata":              metadata,
		"clusterType":           clusterType,
		"minimumValue":          minimumValue,
		"maximumValue":          maximumValue,
		"sum":                   sum,
		"scoreTtlDays":          scoreTtlDays,
		"orderDirection":        orderDirection,
		"entryPeriodEventId":    entryPeriodEventId,
		"rankingRewards":        rankingRewards,
		"accessPeriodEventId":   accessPeriodEventId,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
	}
}

func (p ClusterRankingModelMaster) Pointer() *ClusterRankingModelMaster {
	return &p
}

func CastClusterRankingModelMasters(data []interface{}) []ClusterRankingModelMaster {
	v := make([]ClusterRankingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewClusterRankingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastClusterRankingModelMastersFromDict(data []ClusterRankingModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ClusterRankingScore struct {
	ClusterRankingScoreId *string `json:"clusterRankingScoreId"`
	RankingName           *string `json:"rankingName"`
	ClusterName           *string `json:"clusterName"`
	Season                *int64  `json:"season"`
	UserId                *string `json:"userId"`
	Score                 *int64  `json:"score"`
	Metadata              *string `json:"metadata"`
	CreatedAt             *int64  `json:"createdAt"`
	UpdatedAt             *int64  `json:"updatedAt"`
	Revision              *int64  `json:"revision"`
}

func (p *ClusterRankingScore) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ClusterRankingScore{}
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
		*p = ClusterRankingScore{}
	} else {
		*p = ClusterRankingScore{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clusterRankingScoreId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterRankingScoreId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterRankingScoreId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterRankingScoreId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingScoreId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingScoreId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterRankingScoreId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewClusterRankingScoreFromJson(data string) ClusterRankingScore {
	req := ClusterRankingScore{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewClusterRankingScoreFromDict(data map[string]interface{}) ClusterRankingScore {
	return ClusterRankingScore{
		ClusterRankingScoreId: core.CastString(data["clusterRankingScoreId"]),
		RankingName:           core.CastString(data["rankingName"]),
		ClusterName:           core.CastString(data["clusterName"]),
		Season:                core.CastInt64(data["season"]),
		UserId:                core.CastString(data["userId"]),
		Score:                 core.CastInt64(data["score"]),
		Metadata:              core.CastString(data["metadata"]),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p ClusterRankingScore) ToDict() map[string]interface{} {

	var clusterRankingScoreId *string
	if p.ClusterRankingScoreId != nil {
		clusterRankingScoreId = p.ClusterRankingScoreId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var clusterName *string
	if p.ClusterName != nil {
		clusterName = p.ClusterName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"clusterRankingScoreId": clusterRankingScoreId,
		"rankingName":           rankingName,
		"clusterName":           clusterName,
		"season":                season,
		"userId":                userId,
		"score":                 score,
		"metadata":              metadata,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
	}
}

func (p ClusterRankingScore) Pointer() *ClusterRankingScore {
	return &p
}

func CastClusterRankingScores(data []interface{}) []ClusterRankingScore {
	v := make([]ClusterRankingScore, 0)
	for _, d := range data {
		v = append(v, NewClusterRankingScoreFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastClusterRankingScoresFromDict(data []ClusterRankingScore) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ClusterRankingReceivedReward struct {
	ClusterRankingReceivedRewardId *string `json:"clusterRankingReceivedRewardId"`
	RankingName                    *string `json:"rankingName"`
	ClusterName                    *string `json:"clusterName"`
	Season                         *int64  `json:"season"`
	UserId                         *string `json:"userId"`
	ReceivedAt                     *int64  `json:"receivedAt"`
	Revision                       *int64  `json:"revision"`
}

func (p *ClusterRankingReceivedReward) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ClusterRankingReceivedReward{}
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
		*p = ClusterRankingReceivedReward{}
	} else {
		*p = ClusterRankingReceivedReward{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clusterRankingReceivedRewardId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterRankingReceivedRewardId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterRankingReceivedRewardId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterRankingReceivedRewardId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingReceivedRewardId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingReceivedRewardId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterRankingReceivedRewardId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["receivedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ReceivedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewClusterRankingReceivedRewardFromJson(data string) ClusterRankingReceivedReward {
	req := ClusterRankingReceivedReward{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewClusterRankingReceivedRewardFromDict(data map[string]interface{}) ClusterRankingReceivedReward {
	return ClusterRankingReceivedReward{
		ClusterRankingReceivedRewardId: core.CastString(data["clusterRankingReceivedRewardId"]),
		RankingName:                    core.CastString(data["rankingName"]),
		ClusterName:                    core.CastString(data["clusterName"]),
		Season:                         core.CastInt64(data["season"]),
		UserId:                         core.CastString(data["userId"]),
		ReceivedAt:                     core.CastInt64(data["receivedAt"]),
		Revision:                       core.CastInt64(data["revision"]),
	}
}

func (p ClusterRankingReceivedReward) ToDict() map[string]interface{} {

	var clusterRankingReceivedRewardId *string
	if p.ClusterRankingReceivedRewardId != nil {
		clusterRankingReceivedRewardId = p.ClusterRankingReceivedRewardId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var clusterName *string
	if p.ClusterName != nil {
		clusterName = p.ClusterName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var receivedAt *int64
	if p.ReceivedAt != nil {
		receivedAt = p.ReceivedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"clusterRankingReceivedRewardId": clusterRankingReceivedRewardId,
		"rankingName":                    rankingName,
		"clusterName":                    clusterName,
		"season":                         season,
		"userId":                         userId,
		"receivedAt":                     receivedAt,
		"revision":                       revision,
	}
}

func (p ClusterRankingReceivedReward) Pointer() *ClusterRankingReceivedReward {
	return &p
}

func CastClusterRankingReceivedRewards(data []interface{}) []ClusterRankingReceivedReward {
	v := make([]ClusterRankingReceivedReward, 0)
	for _, d := range data {
		v = append(v, NewClusterRankingReceivedRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastClusterRankingReceivedRewardsFromDict(data []ClusterRankingReceivedReward) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ClusterRankingData struct {
	ClusterRankingDataId *string `json:"clusterRankingDataId"`
	RankingName          *string `json:"rankingName"`
	ClusterName          *string `json:"clusterName"`
	Season               *int64  `json:"season"`
	UserId               *string `json:"userId"`
	Index                *int32  `json:"index"`
	Rank                 *int32  `json:"rank"`
	Score                *int64  `json:"score"`
	Metadata             *string `json:"metadata"`
	CreatedAt            *int64  `json:"createdAt"`
	UpdatedAt            *int64  `json:"updatedAt"`
	Revision             *int64  `json:"revision"`
}

func (p *ClusterRankingData) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ClusterRankingData{}
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
		*p = ClusterRankingData{}
	} else {
		*p = ClusterRankingData{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["clusterRankingDataId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterRankingDataId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterRankingDataId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterRankingDataId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingDataId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterRankingDataId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterRankingDataId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["clusterName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ClusterName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ClusterName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ClusterName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ClusterName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ClusterName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["index"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Index)
		}
		if v, ok := d["rank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rank)
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewClusterRankingDataFromJson(data string) ClusterRankingData {
	req := ClusterRankingData{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewClusterRankingDataFromDict(data map[string]interface{}) ClusterRankingData {
	return ClusterRankingData{
		ClusterRankingDataId: core.CastString(data["clusterRankingDataId"]),
		RankingName:          core.CastString(data["rankingName"]),
		ClusterName:          core.CastString(data["clusterName"]),
		Season:               core.CastInt64(data["season"]),
		UserId:               core.CastString(data["userId"]),
		Index:                core.CastInt32(data["index"]),
		Rank:                 core.CastInt32(data["rank"]),
		Score:                core.CastInt64(data["score"]),
		Metadata:             core.CastString(data["metadata"]),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p ClusterRankingData) ToDict() map[string]interface{} {

	var clusterRankingDataId *string
	if p.ClusterRankingDataId != nil {
		clusterRankingDataId = p.ClusterRankingDataId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var clusterName *string
	if p.ClusterName != nil {
		clusterName = p.ClusterName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var index *int32
	if p.Index != nil {
		index = p.Index
	}
	var rank *int32
	if p.Rank != nil {
		rank = p.Rank
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"clusterRankingDataId": clusterRankingDataId,
		"rankingName":          rankingName,
		"clusterName":          clusterName,
		"season":               season,
		"userId":               userId,
		"index":                index,
		"rank":                 rank,
		"score":                score,
		"metadata":             metadata,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
}

func (p ClusterRankingData) Pointer() *ClusterRankingData {
	return &p
}

func CastClusterRankingDatas(data []interface{}) []ClusterRankingData {
	v := make([]ClusterRankingData, 0)
	for _, d := range data {
		v = append(v, NewClusterRankingDataFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastClusterRankingDatasFromDict(data []ClusterRankingData) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeRankingModel struct {
	SubscribeRankingModelId *string `json:"subscribeRankingModelId"`
	Name                    *string `json:"name"`
	Metadata                *string `json:"metadata"`
	MinimumValue            *int64  `json:"minimumValue"`
	MaximumValue            *int64  `json:"maximumValue"`
	Sum                     *bool   `json:"sum"`
	OrderDirection          *string `json:"orderDirection"`
	EntryPeriodEventId      *string `json:"entryPeriodEventId"`
	AccessPeriodEventId     *string `json:"accessPeriodEventId"`
}

func (p *SubscribeRankingModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeRankingModel{}
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
		*p = SubscribeRankingModel{}
	} else {
		*p = SubscribeRankingModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeRankingModelId)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
				}
			}
		}
	}
	return nil
}

func NewSubscribeRankingModelFromJson(data string) SubscribeRankingModel {
	req := SubscribeRankingModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeRankingModelFromDict(data map[string]interface{}) SubscribeRankingModel {
	return SubscribeRankingModel{
		SubscribeRankingModelId: core.CastString(data["subscribeRankingModelId"]),
		Name:                    core.CastString(data["name"]),
		Metadata:                core.CastString(data["metadata"]),
		MinimumValue:            core.CastInt64(data["minimumValue"]),
		MaximumValue:            core.CastInt64(data["maximumValue"]),
		Sum:                     core.CastBool(data["sum"]),
		OrderDirection:          core.CastString(data["orderDirection"]),
		EntryPeriodEventId:      core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:     core.CastString(data["accessPeriodEventId"]),
	}
}

func (p SubscribeRankingModel) ToDict() map[string]interface{} {

	var subscribeRankingModelId *string
	if p.SubscribeRankingModelId != nil {
		subscribeRankingModelId = p.SubscribeRankingModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
	}
	return map[string]interface{}{
		"subscribeRankingModelId": subscribeRankingModelId,
		"name":                    name,
		"metadata":                metadata,
		"minimumValue":            minimumValue,
		"maximumValue":            maximumValue,
		"sum":                     sum,
		"orderDirection":          orderDirection,
		"entryPeriodEventId":      entryPeriodEventId,
		"accessPeriodEventId":     accessPeriodEventId,
	}
}

func (p SubscribeRankingModel) Pointer() *SubscribeRankingModel {
	return &p
}

func CastSubscribeRankingModels(data []interface{}) []SubscribeRankingModel {
	v := make([]SubscribeRankingModel, 0)
	for _, d := range data {
		v = append(v, NewSubscribeRankingModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeRankingModelsFromDict(data []SubscribeRankingModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeRankingModelMaster struct {
	SubscribeRankingModelId *string `json:"subscribeRankingModelId"`
	Name                    *string `json:"name"`
	Description             *string `json:"description"`
	Metadata                *string `json:"metadata"`
	MinimumValue            *int64  `json:"minimumValue"`
	MaximumValue            *int64  `json:"maximumValue"`
	Sum                     *bool   `json:"sum"`
	ScoreTtlDays            *int32  `json:"scoreTtlDays"`
	OrderDirection          *string `json:"orderDirection"`
	EntryPeriodEventId      *string `json:"entryPeriodEventId"`
	AccessPeriodEventId     *string `json:"accessPeriodEventId"`
	CreatedAt               *int64  `json:"createdAt"`
	UpdatedAt               *int64  `json:"updatedAt"`
	Revision                *int64  `json:"revision"`
}

func (p *SubscribeRankingModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeRankingModelMaster{}
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
		*p = SubscribeRankingModelMaster{}
	} else {
		*p = SubscribeRankingModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeRankingModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeRankingModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeRankingModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeRankingModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeRankingModelId)
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
		if v, ok := d["minimumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MinimumValue)
		}
		if v, ok := d["maximumValue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaximumValue)
		}
		if v, ok := d["sum"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Sum)
		}
		if v, ok := d["scoreTtlDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScoreTtlDays)
		}
		if v, ok := d["orderDirection"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OrderDirection = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OrderDirection = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OrderDirection = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OrderDirection = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OrderDirection)
				}
			}
		}
		if v, ok := d["entryPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EntryPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EntryPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EntryPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EntryPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EntryPeriodEventId)
				}
			}
		}
		if v, ok := d["accessPeriodEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AccessPeriodEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AccessPeriodEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AccessPeriodEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AccessPeriodEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AccessPeriodEventId)
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

func NewSubscribeRankingModelMasterFromJson(data string) SubscribeRankingModelMaster {
	req := SubscribeRankingModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeRankingModelMasterFromDict(data map[string]interface{}) SubscribeRankingModelMaster {
	return SubscribeRankingModelMaster{
		SubscribeRankingModelId: core.CastString(data["subscribeRankingModelId"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Metadata:                core.CastString(data["metadata"]),
		MinimumValue:            core.CastInt64(data["minimumValue"]),
		MaximumValue:            core.CastInt64(data["maximumValue"]),
		Sum:                     core.CastBool(data["sum"]),
		ScoreTtlDays:            core.CastInt32(data["scoreTtlDays"]),
		OrderDirection:          core.CastString(data["orderDirection"]),
		EntryPeriodEventId:      core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:     core.CastString(data["accessPeriodEventId"]),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
	}
}

func (p SubscribeRankingModelMaster) ToDict() map[string]interface{} {

	var subscribeRankingModelId *string
	if p.SubscribeRankingModelId != nil {
		subscribeRankingModelId = p.SubscribeRankingModelId
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
	var minimumValue *int64
	if p.MinimumValue != nil {
		minimumValue = p.MinimumValue
	}
	var maximumValue *int64
	if p.MaximumValue != nil {
		maximumValue = p.MaximumValue
	}
	var sum *bool
	if p.Sum != nil {
		sum = p.Sum
	}
	var scoreTtlDays *int32
	if p.ScoreTtlDays != nil {
		scoreTtlDays = p.ScoreTtlDays
	}
	var orderDirection *string
	if p.OrderDirection != nil {
		orderDirection = p.OrderDirection
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
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
		"subscribeRankingModelId": subscribeRankingModelId,
		"name":                    name,
		"description":             description,
		"metadata":                metadata,
		"minimumValue":            minimumValue,
		"maximumValue":            maximumValue,
		"sum":                     sum,
		"scoreTtlDays":            scoreTtlDays,
		"orderDirection":          orderDirection,
		"entryPeriodEventId":      entryPeriodEventId,
		"accessPeriodEventId":     accessPeriodEventId,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
	}
}

func (p SubscribeRankingModelMaster) Pointer() *SubscribeRankingModelMaster {
	return &p
}

func CastSubscribeRankingModelMasters(data []interface{}) []SubscribeRankingModelMaster {
	v := make([]SubscribeRankingModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSubscribeRankingModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeRankingModelMastersFromDict(data []SubscribeRankingModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Subscribe struct {
	SubscribeId   *string   `json:"subscribeId"`
	RankingName   *string   `json:"rankingName"`
	UserId        *string   `json:"userId"`
	TargetUserIds []*string `json:"targetUserIds"`
	FromUserIds   []*string `json:"fromUserIds"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
	Revision      *int64    `json:"revision"`
}

func (p *Subscribe) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Subscribe{}
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
		*p = Subscribe{}
	} else {
		*p = Subscribe{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserIds"]; ok && v != nil {
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
				p.TargetUserIds = l
			}
		}
		if v, ok := d["fromUserIds"]; ok && v != nil {
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
				p.FromUserIds = l
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

func NewSubscribeFromJson(data string) Subscribe {
	req := Subscribe{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
	return Subscribe{
		SubscribeId:   core.CastString(data["subscribeId"]),
		RankingName:   core.CastString(data["rankingName"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		FromUserIds:   core.CastStrings(core.CastArray(data["fromUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p Subscribe) ToDict() map[string]interface{} {

	var subscribeId *string
	if p.SubscribeId != nil {
		subscribeId = p.SubscribeId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserIds []interface{}
	if p.TargetUserIds != nil {
		targetUserIds = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	var fromUserIds []interface{}
	if p.FromUserIds != nil {
		fromUserIds = core.CastStringsFromDict(
			p.FromUserIds,
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
		"subscribeId":   subscribeId,
		"rankingName":   rankingName,
		"userId":        userId,
		"targetUserIds": targetUserIds,
		"fromUserIds":   fromUserIds,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p Subscribe) Pointer() *Subscribe {
	return &p
}

func CastSubscribes(data []interface{}) []Subscribe {
	v := make([]Subscribe, 0)
	for _, d := range data {
		v = append(v, NewSubscribeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribesFromDict(data []Subscribe) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeRankingScore struct {
	SubscribeRankingScoreId *string `json:"subscribeRankingScoreId"`
	RankingName             *string `json:"rankingName"`
	Season                  *int64  `json:"season"`
	UserId                  *string `json:"userId"`
	Score                   *int64  `json:"score"`
	Metadata                *string `json:"metadata"`
	CreatedAt               *int64  `json:"createdAt"`
	UpdatedAt               *int64  `json:"updatedAt"`
	Revision                *int64  `json:"revision"`
}

func (p *SubscribeRankingScore) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeRankingScore{}
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
		*p = SubscribeRankingScore{}
	} else {
		*p = SubscribeRankingScore{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeRankingScoreId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeRankingScoreId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeRankingScoreId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeRankingScoreId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingScoreId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingScoreId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeRankingScoreId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewSubscribeRankingScoreFromJson(data string) SubscribeRankingScore {
	req := SubscribeRankingScore{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeRankingScoreFromDict(data map[string]interface{}) SubscribeRankingScore {
	return SubscribeRankingScore{
		SubscribeRankingScoreId: core.CastString(data["subscribeRankingScoreId"]),
		RankingName:             core.CastString(data["rankingName"]),
		Season:                  core.CastInt64(data["season"]),
		UserId:                  core.CastString(data["userId"]),
		Score:                   core.CastInt64(data["score"]),
		Metadata:                core.CastString(data["metadata"]),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
	}
}

func (p SubscribeRankingScore) ToDict() map[string]interface{} {

	var subscribeRankingScoreId *string
	if p.SubscribeRankingScoreId != nil {
		subscribeRankingScoreId = p.SubscribeRankingScoreId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"subscribeRankingScoreId": subscribeRankingScoreId,
		"rankingName":             rankingName,
		"season":                  season,
		"userId":                  userId,
		"score":                   score,
		"metadata":                metadata,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
	}
}

func (p SubscribeRankingScore) Pointer() *SubscribeRankingScore {
	return &p
}

func CastSubscribeRankingScores(data []interface{}) []SubscribeRankingScore {
	v := make([]SubscribeRankingScore, 0)
	for _, d := range data {
		v = append(v, NewSubscribeRankingScoreFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeRankingScoresFromDict(data []SubscribeRankingScore) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeRankingData struct {
	SubscribeRankingDataId *string `json:"subscribeRankingDataId"`
	RankingName            *string `json:"rankingName"`
	Season                 *int64  `json:"season"`
	UserId                 *string `json:"userId"`
	Index                  *int32  `json:"index"`
	Rank                   *int32  `json:"rank"`
	ScorerUserId           *string `json:"scorerUserId"`
	Score                  *int64  `json:"score"`
	Metadata               *string `json:"metadata"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
	Revision               *int64  `json:"revision"`
}

func (p *SubscribeRankingData) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeRankingData{}
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
		*p = SubscribeRankingData{}
	} else {
		*p = SubscribeRankingData{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["subscribeRankingDataId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeRankingDataId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeRankingDataId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeRankingDataId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingDataId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeRankingDataId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeRankingDataId)
				}
			}
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
				}
			}
		}
		if v, ok := d["season"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Season)
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
		if v, ok := d["index"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Index)
		}
		if v, ok := d["rank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rank)
		}
		if v, ok := d["scorerUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScorerUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScorerUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScorerUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScorerUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScorerUserId)
				}
			}
		}
		if v, ok := d["score"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Score)
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

func NewSubscribeRankingDataFromJson(data string) SubscribeRankingData {
	req := SubscribeRankingData{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeRankingDataFromDict(data map[string]interface{}) SubscribeRankingData {
	return SubscribeRankingData{
		SubscribeRankingDataId: core.CastString(data["subscribeRankingDataId"]),
		RankingName:            core.CastString(data["rankingName"]),
		Season:                 core.CastInt64(data["season"]),
		UserId:                 core.CastString(data["userId"]),
		Index:                  core.CastInt32(data["index"]),
		Rank:                   core.CastInt32(data["rank"]),
		ScorerUserId:           core.CastString(data["scorerUserId"]),
		Score:                  core.CastInt64(data["score"]),
		Metadata:               core.CastString(data["metadata"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p SubscribeRankingData) ToDict() map[string]interface{} {

	var subscribeRankingDataId *string
	if p.SubscribeRankingDataId != nil {
		subscribeRankingDataId = p.SubscribeRankingDataId
	}
	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var season *int64
	if p.Season != nil {
		season = p.Season
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var index *int32
	if p.Index != nil {
		index = p.Index
	}
	var rank *int32
	if p.Rank != nil {
		rank = p.Rank
	}
	var scorerUserId *string
	if p.ScorerUserId != nil {
		scorerUserId = p.ScorerUserId
	}
	var score *int64
	if p.Score != nil {
		score = p.Score
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
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
		"subscribeRankingDataId": subscribeRankingDataId,
		"rankingName":            rankingName,
		"season":                 season,
		"userId":                 userId,
		"index":                  index,
		"rank":                   rank,
		"scorerUserId":           scorerUserId,
		"score":                  score,
		"metadata":               metadata,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
}

func (p SubscribeRankingData) Pointer() *SubscribeRankingData {
	return &p
}

func CastSubscribeRankingDatas(data []interface{}) []SubscribeRankingData {
	v := make([]SubscribeRankingData, 0)
	for _, d := range data {
		v = append(v, NewSubscribeRankingDataFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeRankingDatasFromDict(data []SubscribeRankingData) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentRankingMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentRankingMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentRankingMaster{}
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
		*p = CurrentRankingMaster{}
	} else {
		*p = CurrentRankingMaster{}
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

func NewCurrentRankingMasterFromJson(data string) CurrentRankingMaster {
	req := CurrentRankingMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentRankingMasterFromDict(data map[string]interface{}) CurrentRankingMaster {
	return CurrentRankingMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRankingMaster) ToDict() map[string]interface{} {

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

func (p CurrentRankingMaster) Pointer() *CurrentRankingMaster {
	return &p
}

func CastCurrentRankingMasters(data []interface{}) []CurrentRankingMaster {
	v := make([]CurrentRankingMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentRankingMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentRankingMastersFromDict(data []CurrentRankingMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeUser struct {
	RankingName  *string `json:"rankingName"`
	UserId       *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func (p *SubscribeUser) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SubscribeUser{}
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
		*p = SubscribeUser{}
	} else {
		*p = SubscribeUser{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rankingName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RankingName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RankingName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RankingName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RankingName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RankingName)
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
		if v, ok := d["targetUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetUserId)
				}
			}
		}
	}
	return nil
}

func NewSubscribeUserFromJson(data string) SubscribeUser {
	req := SubscribeUser{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeUserFromDict(data map[string]interface{}) SubscribeUser {
	return SubscribeUser{
		RankingName:  core.CastString(data["rankingName"]),
		UserId:       core.CastString(data["userId"]),
		TargetUserId: core.CastString(data["targetUserId"]),
	}
}

func (p SubscribeUser) ToDict() map[string]interface{} {

	var rankingName *string
	if p.RankingName != nil {
		rankingName = p.RankingName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var targetUserId *string
	if p.TargetUserId != nil {
		targetUserId = p.TargetUserId
	}
	return map[string]interface{}{
		"rankingName":  rankingName,
		"userId":       userId,
		"targetUserId": targetUserId,
	}
}

func (p SubscribeUser) Pointer() *SubscribeUser {
	return &p
}

func CastSubscribeUsers(data []interface{}) []SubscribeUser {
	v := make([]SubscribeUser, 0)
	for _, d := range data {
		v = append(v, NewSubscribeUserFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribeUsersFromDict(data []SubscribeUser) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RankingReward struct {
	ThresholdRank  *int32          `json:"thresholdRank"`
	Metadata       *string         `json:"metadata"`
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func (p *RankingReward) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RankingReward{}
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
		*p = RankingReward{}
	} else {
		*p = RankingReward{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["thresholdRank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ThresholdRank)
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
		if v, ok := d["acquireActions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActions)
		}
	}
	return nil
}

func NewRankingRewardFromJson(data string) RankingReward {
	req := RankingReward{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRankingRewardFromDict(data map[string]interface{}) RankingReward {
	return RankingReward{
		ThresholdRank:  core.CastInt32(data["thresholdRank"]),
		Metadata:       core.CastString(data["metadata"]),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p RankingReward) ToDict() map[string]interface{} {

	var thresholdRank *int32
	if p.ThresholdRank != nil {
		thresholdRank = p.ThresholdRank
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"thresholdRank":  thresholdRank,
		"metadata":       metadata,
		"acquireActions": acquireActions,
	}
}

func (p RankingReward) Pointer() *RankingReward {
	return &p
}

func CastRankingRewards(data []interface{}) []RankingReward {
	v := make([]RankingReward, 0)
	for _, d := range data {
		v = append(v, NewRankingRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingRewardsFromDict(data []RankingReward) []interface{} {
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
