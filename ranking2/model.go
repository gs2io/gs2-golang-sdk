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

type GlobalRankingModel struct {
	GlobalRankingModelId   *string         `json:"globalRankingModelId"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	MinimumValue           *int64          `json:"minimumValue"`
	MaximumValue           *int64          `json:"maximumValue"`
	Sum                    *bool           `json:"sum"`
	OrderDirection         *string         `json:"orderDirection"`
	EntryPeriodEventId     *string         `json:"entryPeriodEventId"`
	RankingRewards         []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId    *string         `json:"accessPeriodEventId"`
	RewardCalculationIndex *string         `json:"rewardCalculationIndex"`
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
		if v, ok := d["rewardCalculationIndex"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RewardCalculationIndex = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RewardCalculationIndex = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RewardCalculationIndex = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RewardCalculationIndex)
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
		GlobalRankingModelId: func() *string {
			v, ok := data["globalRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalRankingModelId"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
		RewardCalculationIndex: func() *string {
			v, ok := data["rewardCalculationIndex"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rewardCalculationIndex"])
		}(),
	}
}

func (p GlobalRankingModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalRankingModelId != nil {
		m["globalRankingModelId"] = p.GlobalRankingModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.RankingRewards != nil {
		m["rankingRewards"] = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
	}
	if p.RewardCalculationIndex != nil {
		m["rewardCalculationIndex"] = p.RewardCalculationIndex
	}
	return m
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
	GlobalRankingModelId   *string         `json:"globalRankingModelId"`
	Name                   *string         `json:"name"`
	Description            *string         `json:"description"`
	Metadata               *string         `json:"metadata"`
	MinimumValue           *int64          `json:"minimumValue"`
	MaximumValue           *int64          `json:"maximumValue"`
	Sum                    *bool           `json:"sum"`
	OrderDirection         *string         `json:"orderDirection"`
	EntryPeriodEventId     *string         `json:"entryPeriodEventId"`
	RankingRewards         []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId    *string         `json:"accessPeriodEventId"`
	RewardCalculationIndex *string         `json:"rewardCalculationIndex"`
	CreatedAt              *int64          `json:"createdAt"`
	UpdatedAt              *int64          `json:"updatedAt"`
	Revision               *int64          `json:"revision"`
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
		if v, ok := d["rewardCalculationIndex"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RewardCalculationIndex = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RewardCalculationIndex = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RewardCalculationIndex = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RewardCalculationIndex)
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
		GlobalRankingModelId: func() *string {
			v, ok := data["globalRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalRankingModelId"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
		RewardCalculationIndex: func() *string {
			v, ok := data["rewardCalculationIndex"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rewardCalculationIndex"])
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

func (p GlobalRankingModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalRankingModelId != nil {
		m["globalRankingModelId"] = p.GlobalRankingModelId
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
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.RankingRewards != nil {
		m["rankingRewards"] = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
	}
	if p.RewardCalculationIndex != nil {
		m["rewardCalculationIndex"] = p.RewardCalculationIndex
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
		GlobalRankingScoreId: func() *string {
			v, ok := data["globalRankingScoreId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalRankingScoreId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p GlobalRankingScore) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalRankingScoreId != nil {
		m["globalRankingScoreId"] = p.GlobalRankingScoreId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.Score != nil {
		m["score"] = p.Score
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
		GlobalRankingReceivedRewardId: func() *string {
			v, ok := data["globalRankingReceivedRewardId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalRankingReceivedRewardId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		ReceivedAt: func() *int64 {
			v, ok := data["receivedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["receivedAt"])
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

func (p GlobalRankingReceivedReward) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalRankingReceivedRewardId != nil {
		m["globalRankingReceivedRewardId"] = p.GlobalRankingReceivedRewardId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.ReceivedAt != nil {
		m["receivedAt"] = p.ReceivedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
	InvertUpdatedAt     *int64  `json:"invertUpdatedAt"`
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
		if v, ok := d["invertUpdatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InvertUpdatedAt)
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
		GlobalRankingDataId: func() *string {
			v, ok := data["globalRankingDataId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["globalRankingDataId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Index: func() *int32 {
			v, ok := data["index"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["index"])
		}(),
		Rank: func() *int32 {
			v, ok := data["rank"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rank"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		InvertUpdatedAt: func() *int64 {
			v, ok := data["invertUpdatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["invertUpdatedAt"])
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

func (p GlobalRankingData) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GlobalRankingDataId != nil {
		m["globalRankingDataId"] = p.GlobalRankingDataId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Index != nil {
		m["index"] = p.Index
	}
	if p.Rank != nil {
		m["rank"] = p.Rank
	}
	if p.Score != nil {
		m["score"] = p.Score
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.InvertUpdatedAt != nil {
		m["invertUpdatedAt"] = p.InvertUpdatedAt
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
	ClusterRankingModelId  *string         `json:"clusterRankingModelId"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	ClusterType            *string         `json:"clusterType"`
	MinimumValue           *int64          `json:"minimumValue"`
	MaximumValue           *int64          `json:"maximumValue"`
	Sum                    *bool           `json:"sum"`
	OrderDirection         *string         `json:"orderDirection"`
	EntryPeriodEventId     *string         `json:"entryPeriodEventId"`
	RankingRewards         []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId    *string         `json:"accessPeriodEventId"`
	RewardCalculationIndex *string         `json:"rewardCalculationIndex"`
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
		if v, ok := d["rewardCalculationIndex"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RewardCalculationIndex = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RewardCalculationIndex = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RewardCalculationIndex = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RewardCalculationIndex)
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
		ClusterRankingModelId: func() *string {
			v, ok := data["clusterRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterRankingModelId"])
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
		ClusterType: func() *string {
			v, ok := data["clusterType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterType"])
		}(),
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
		RewardCalculationIndex: func() *string {
			v, ok := data["rewardCalculationIndex"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rewardCalculationIndex"])
		}(),
	}
}

func (p ClusterRankingModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ClusterRankingModelId != nil {
		m["clusterRankingModelId"] = p.ClusterRankingModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.ClusterType != nil {
		m["clusterType"] = p.ClusterType
	}
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.RankingRewards != nil {
		m["rankingRewards"] = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
	}
	if p.RewardCalculationIndex != nil {
		m["rewardCalculationIndex"] = p.RewardCalculationIndex
	}
	return m
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
	ClusterRankingModelId  *string         `json:"clusterRankingModelId"`
	Name                   *string         `json:"name"`
	Description            *string         `json:"description"`
	Metadata               *string         `json:"metadata"`
	ClusterType            *string         `json:"clusterType"`
	MinimumValue           *int64          `json:"minimumValue"`
	MaximumValue           *int64          `json:"maximumValue"`
	Sum                    *bool           `json:"sum"`
	ScoreTtlDays           *int32          `json:"scoreTtlDays"`
	OrderDirection         *string         `json:"orderDirection"`
	EntryPeriodEventId     *string         `json:"entryPeriodEventId"`
	RankingRewards         []RankingReward `json:"rankingRewards"`
	AccessPeriodEventId    *string         `json:"accessPeriodEventId"`
	RewardCalculationIndex *string         `json:"rewardCalculationIndex"`
	CreatedAt              *int64          `json:"createdAt"`
	UpdatedAt              *int64          `json:"updatedAt"`
	Revision               *int64          `json:"revision"`
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
		if v, ok := d["rewardCalculationIndex"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RewardCalculationIndex = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RewardCalculationIndex = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RewardCalculationIndex = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RewardCalculationIndex = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RewardCalculationIndex)
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
		ClusterRankingModelId: func() *string {
			v, ok := data["clusterRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterRankingModelId"])
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
		ClusterType: func() *string {
			v, ok := data["clusterType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterType"])
		}(),
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		RankingRewards: func() []RankingReward {
			if data["rankingRewards"] == nil {
				return nil
			}
			return CastRankingRewards(core.CastArray(data["rankingRewards"]))
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
		RewardCalculationIndex: func() *string {
			v, ok := data["rewardCalculationIndex"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rewardCalculationIndex"])
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

func (p ClusterRankingModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ClusterRankingModelId != nil {
		m["clusterRankingModelId"] = p.ClusterRankingModelId
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
	if p.ClusterType != nil {
		m["clusterType"] = p.ClusterType
	}
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.ScoreTtlDays != nil {
		m["scoreTtlDays"] = p.ScoreTtlDays
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.RankingRewards != nil {
		m["rankingRewards"] = CastRankingRewardsFromDict(
			p.RankingRewards,
		)
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
	}
	if p.RewardCalculationIndex != nil {
		m["rewardCalculationIndex"] = p.RewardCalculationIndex
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
		ClusterRankingScoreId: func() *string {
			v, ok := data["clusterRankingScoreId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterRankingScoreId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p ClusterRankingScore) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ClusterRankingScoreId != nil {
		m["clusterRankingScoreId"] = p.ClusterRankingScoreId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.ClusterName != nil {
		m["clusterName"] = p.ClusterName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Score != nil {
		m["score"] = p.Score
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
		ClusterRankingReceivedRewardId: func() *string {
			v, ok := data["clusterRankingReceivedRewardId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterRankingReceivedRewardId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ReceivedAt: func() *int64 {
			v, ok := data["receivedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["receivedAt"])
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

func (p ClusterRankingReceivedReward) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ClusterRankingReceivedRewardId != nil {
		m["clusterRankingReceivedRewardId"] = p.ClusterRankingReceivedRewardId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.ClusterName != nil {
		m["clusterName"] = p.ClusterName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ReceivedAt != nil {
		m["receivedAt"] = p.ReceivedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
	InvertUpdatedAt      *int64  `json:"invertUpdatedAt"`
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
		if v, ok := d["invertUpdatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InvertUpdatedAt)
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
		ClusterRankingDataId: func() *string {
			v, ok := data["clusterRankingDataId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterRankingDataId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		ClusterName: func() *string {
			v, ok := data["clusterName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["clusterName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Index: func() *int32 {
			v, ok := data["index"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["index"])
		}(),
		Rank: func() *int32 {
			v, ok := data["rank"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rank"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		InvertUpdatedAt: func() *int64 {
			v, ok := data["invertUpdatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["invertUpdatedAt"])
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

func (p ClusterRankingData) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ClusterRankingDataId != nil {
		m["clusterRankingDataId"] = p.ClusterRankingDataId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.ClusterName != nil {
		m["clusterName"] = p.ClusterName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Index != nil {
		m["index"] = p.Index
	}
	if p.Rank != nil {
		m["rank"] = p.Rank
	}
	if p.Score != nil {
		m["score"] = p.Score
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.InvertUpdatedAt != nil {
		m["invertUpdatedAt"] = p.InvertUpdatedAt
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
		SubscribeRankingModelId: func() *string {
			v, ok := data["subscribeRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeRankingModelId"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
		}(),
	}
}

func (p SubscribeRankingModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeRankingModelId != nil {
		m["subscribeRankingModelId"] = p.SubscribeRankingModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
	}
	return m
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
		SubscribeRankingModelId: func() *string {
			v, ok := data["subscribeRankingModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeRankingModelId"])
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
		MinimumValue: func() *int64 {
			v, ok := data["minimumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["minimumValue"])
		}(),
		MaximumValue: func() *int64 {
			v, ok := data["maximumValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["maximumValue"])
		}(),
		Sum: func() *bool {
			v, ok := data["sum"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["sum"])
		}(),
		ScoreTtlDays: func() *int32 {
			v, ok := data["scoreTtlDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["scoreTtlDays"])
		}(),
		OrderDirection: func() *string {
			v, ok := data["orderDirection"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["orderDirection"])
		}(),
		EntryPeriodEventId: func() *string {
			v, ok := data["entryPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["entryPeriodEventId"])
		}(),
		AccessPeriodEventId: func() *string {
			v, ok := data["accessPeriodEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["accessPeriodEventId"])
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

func (p SubscribeRankingModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeRankingModelId != nil {
		m["subscribeRankingModelId"] = p.SubscribeRankingModelId
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
	if p.MinimumValue != nil {
		m["minimumValue"] = p.MinimumValue
	}
	if p.MaximumValue != nil {
		m["maximumValue"] = p.MaximumValue
	}
	if p.Sum != nil {
		m["sum"] = p.Sum
	}
	if p.ScoreTtlDays != nil {
		m["scoreTtlDays"] = p.ScoreTtlDays
	}
	if p.OrderDirection != nil {
		m["orderDirection"] = p.OrderDirection
	}
	if p.EntryPeriodEventId != nil {
		m["entryPeriodEventId"] = p.EntryPeriodEventId
	}
	if p.AccessPeriodEventId != nil {
		m["accessPeriodEventId"] = p.AccessPeriodEventId
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
		SubscribeId: func() *string {
			v, ok := data["subscribeId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserIds: func() []*string {
			v, ok := data["targetUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		FromUserIds: func() []*string {
			v, ok := data["fromUserIds"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p Subscribe) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeId != nil {
		m["subscribeId"] = p.SubscribeId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserIds != nil {
		m["targetUserIds"] = core.CastStringsFromDict(
			p.TargetUserIds,
		)
	}
	if p.FromUserIds != nil {
		m["fromUserIds"] = core.CastStringsFromDict(
			p.FromUserIds,
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
		SubscribeRankingScoreId: func() *string {
			v, ok := data["subscribeRankingScoreId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeRankingScoreId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
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

func (p SubscribeRankingScore) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeRankingScoreId != nil {
		m["subscribeRankingScoreId"] = p.SubscribeRankingScoreId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Score != nil {
		m["score"] = p.Score
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
	InvertUpdatedAt        *int64  `json:"invertUpdatedAt"`
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
		if v, ok := d["invertUpdatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.InvertUpdatedAt)
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
		SubscribeRankingDataId: func() *string {
			v, ok := data["subscribeRankingDataId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["subscribeRankingDataId"])
		}(),
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		Season: func() *int64 {
			v, ok := data["season"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["season"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Index: func() *int32 {
			v, ok := data["index"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["index"])
		}(),
		Rank: func() *int32 {
			v, ok := data["rank"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["rank"])
		}(),
		ScorerUserId: func() *string {
			v, ok := data["scorerUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scorerUserId"])
		}(),
		Score: func() *int64 {
			v, ok := data["score"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["score"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		InvertUpdatedAt: func() *int64 {
			v, ok := data["invertUpdatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["invertUpdatedAt"])
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

func (p SubscribeRankingData) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SubscribeRankingDataId != nil {
		m["subscribeRankingDataId"] = p.SubscribeRankingDataId
	}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.Season != nil {
		m["season"] = p.Season
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Index != nil {
		m["index"] = p.Index
	}
	if p.Rank != nil {
		m["rank"] = p.Rank
	}
	if p.ScorerUserId != nil {
		m["scorerUserId"] = p.ScorerUserId
	}
	if p.Score != nil {
		m["score"] = p.Score
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.InvertUpdatedAt != nil {
		m["invertUpdatedAt"] = p.InvertUpdatedAt
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

func (p CurrentRankingMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
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
		RankingName: func() *string {
			v, ok := data["rankingName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rankingName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		TargetUserId: func() *string {
			v, ok := data["targetUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetUserId"])
		}(),
	}
}

func (p SubscribeUser) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.RankingName != nil {
		m["rankingName"] = p.RankingName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.TargetUserId != nil {
		m["targetUserId"] = p.TargetUserId
	}
	return m
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
		ThresholdRank: func() *int32 {
			v, ok := data["thresholdRank"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["thresholdRank"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
		}(),
	}
}

func (p RankingReward) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ThresholdRank != nil {
		m["thresholdRank"] = p.ThresholdRank
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.AcquireActions != nil {
		m["acquireActions"] = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return m
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
