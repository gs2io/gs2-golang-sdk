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

package ranking

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId       *string        `json:"namespaceId"`
	Name              *string        `json:"name"`
	Description       *string        `json:"description"`
	LastCalculatedAts []CalculatedAt `json:"lastCalculatedAts"`
	LogSetting        *LogSetting    `json:"logSetting"`
	CreatedAt         *int64         `json:"createdAt"`
	UpdatedAt         *int64         `json:"updatedAt"`
	Revision          *int64         `json:"revision"`
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
		if v, ok := d["lastCalculatedAts"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LastCalculatedAts)
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
		NamespaceId:       core.CastString(data["namespaceId"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		LastCalculatedAts: CastCalculatedAts(core.CastArray(data["lastCalculatedAts"])),
		LogSetting:        NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
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
	var lastCalculatedAts []interface{}
	if p.LastCalculatedAts != nil {
		lastCalculatedAts = CastCalculatedAtsFromDict(
			p.LastCalculatedAts,
		)
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
		"namespaceId":       namespaceId,
		"name":              name,
		"description":       description,
		"lastCalculatedAts": lastCalculatedAts,
		"logSetting":        logSetting,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
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

type CategoryModel struct {
	CategoryModelId      *string               `json:"categoryModelId"`
	Name                 *string               `json:"name"`
	Metadata             *string               `json:"metadata"`
	MinimumValue         *int64                `json:"minimumValue"`
	MaximumValue         *int64                `json:"maximumValue"`
	Sum                  *bool                 `json:"sum"`
	OrderDirection       *string               `json:"orderDirection"`
	Scope                *string               `json:"scope"`
	GlobalRankingSetting *GlobalRankingSetting `json:"globalRankingSetting"`
	EntryPeriodEventId   *string               `json:"entryPeriodEventId"`
	AccessPeriodEventId  *string               `json:"accessPeriodEventId"`
	// Deprecated: should not be used
	UniqueByUserId *bool `json:"uniqueByUserId"`
	// Deprecated: should not be used
	CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
	// Deprecated: should not be used
	CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
	// Deprecated: should not be used
	CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
	// Deprecated: should not be used
	AdditionalScopes []Scope `json:"additionalScopes"`
	// Deprecated: should not be used
	IgnoreUserIds []*string `json:"ignoreUserIds"`
	// Deprecated: should not be used
	Generation *string `json:"generation"`
}

func (p *CategoryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CategoryModel{}
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
		*p = CategoryModel{}
	} else {
		*p = CategoryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["categoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryModelId)
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["globalRankingSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GlobalRankingSetting)
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
		if v, ok := d["uniqueByUserId"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UniqueByUserId)
		}
		if v, ok := d["calculateFixedTimingHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateFixedTimingHour)
		}
		if v, ok := d["calculateFixedTimingMinute"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateFixedTimingMinute)
		}
		if v, ok := d["calculateIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateIntervalMinutes)
		}
		if v, ok := d["additionalScopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AdditionalScopes)
		}
		if v, ok := d["ignoreUserIds"]; ok && v != nil {
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
				p.IgnoreUserIds = l
			}
		}
		if v, ok := d["generation"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Generation = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Generation = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Generation = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Generation)
				}
			}
		}
	}
	return nil
}

func NewCategoryModelFromJson(data string) CategoryModel {
	req := CategoryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCategoryModelFromDict(data map[string]interface{}) CategoryModel {
	return CategoryModel{
		CategoryModelId:            core.CastString(data["categoryModelId"]),
		Name:                       core.CastString(data["name"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		Sum:                        core.CastBool(data["sum"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		GlobalRankingSetting:       NewGlobalRankingSettingFromDict(core.CastMap(data["globalRankingSetting"])).Pointer(),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		AdditionalScopes:           CastScopes(core.CastArray(data["additionalScopes"])),
		IgnoreUserIds:              core.CastStrings(core.CastArray(data["ignoreUserIds"])),
		Generation:                 core.CastString(data["generation"]),
	}
}

func (p CategoryModel) ToDict() map[string]interface{} {

	var categoryModelId *string
	if p.CategoryModelId != nil {
		categoryModelId = p.CategoryModelId
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
	var scope *string
	if p.Scope != nil {
		scope = p.Scope
	}
	var globalRankingSetting map[string]interface{}
	if p.GlobalRankingSetting != nil {
		globalRankingSetting = p.GlobalRankingSetting.ToDict()
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
	}
	var uniqueByUserId *bool
	if p.UniqueByUserId != nil {
		uniqueByUserId = p.UniqueByUserId
	}
	var calculateFixedTimingHour *int32
	if p.CalculateFixedTimingHour != nil {
		calculateFixedTimingHour = p.CalculateFixedTimingHour
	}
	var calculateFixedTimingMinute *int32
	if p.CalculateFixedTimingMinute != nil {
		calculateFixedTimingMinute = p.CalculateFixedTimingMinute
	}
	var calculateIntervalMinutes *int32
	if p.CalculateIntervalMinutes != nil {
		calculateIntervalMinutes = p.CalculateIntervalMinutes
	}
	var additionalScopes []interface{}
	if p.AdditionalScopes != nil {
		additionalScopes = CastScopesFromDict(
			p.AdditionalScopes,
		)
	}
	var ignoreUserIds []interface{}
	if p.IgnoreUserIds != nil {
		ignoreUserIds = core.CastStringsFromDict(
			p.IgnoreUserIds,
		)
	}
	var generation *string
	if p.Generation != nil {
		generation = p.Generation
	}
	return map[string]interface{}{
		"categoryModelId":            categoryModelId,
		"name":                       name,
		"metadata":                   metadata,
		"minimumValue":               minimumValue,
		"maximumValue":               maximumValue,
		"sum":                        sum,
		"orderDirection":             orderDirection,
		"scope":                      scope,
		"globalRankingSetting":       globalRankingSetting,
		"entryPeriodEventId":         entryPeriodEventId,
		"accessPeriodEventId":        accessPeriodEventId,
		"uniqueByUserId":             uniqueByUserId,
		"calculateFixedTimingHour":   calculateFixedTimingHour,
		"calculateFixedTimingMinute": calculateFixedTimingMinute,
		"calculateIntervalMinutes":   calculateIntervalMinutes,
		"additionalScopes":           additionalScopes,
		"ignoreUserIds":              ignoreUserIds,
		"generation":                 generation,
	}
}

func (p CategoryModel) Pointer() *CategoryModel {
	return &p
}

func CastCategoryModels(data []interface{}) []CategoryModel {
	v := make([]CategoryModel, 0)
	for _, d := range data {
		v = append(v, NewCategoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCategoryModelsFromDict(data []CategoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CategoryModelMaster struct {
	CategoryModelId      *string               `json:"categoryModelId"`
	Name                 *string               `json:"name"`
	Description          *string               `json:"description"`
	Metadata             *string               `json:"metadata"`
	MinimumValue         *int64                `json:"minimumValue"`
	MaximumValue         *int64                `json:"maximumValue"`
	Sum                  *bool                 `json:"sum"`
	OrderDirection       *string               `json:"orderDirection"`
	Scope                *string               `json:"scope"`
	GlobalRankingSetting *GlobalRankingSetting `json:"globalRankingSetting"`
	EntryPeriodEventId   *string               `json:"entryPeriodEventId"`
	AccessPeriodEventId  *string               `json:"accessPeriodEventId"`
	// Deprecated: should not be used
	UniqueByUserId *bool `json:"uniqueByUserId"`
	// Deprecated: should not be used
	CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
	// Deprecated: should not be used
	CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
	// Deprecated: should not be used
	CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
	// Deprecated: should not be used
	AdditionalScopes []Scope `json:"additionalScopes"`
	// Deprecated: should not be used
	IgnoreUserIds []*string `json:"ignoreUserIds"`
	// Deprecated: should not be used
	Generation *string `json:"generation"`
	CreatedAt  *int64  `json:"createdAt"`
	UpdatedAt  *int64  `json:"updatedAt"`
	Revision   *int64  `json:"revision"`
}

func (p *CategoryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CategoryModelMaster{}
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
		*p = CategoryModelMaster{}
	} else {
		*p = CategoryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["categoryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryModelId)
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["globalRankingSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GlobalRankingSetting)
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
		if v, ok := d["uniqueByUserId"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UniqueByUserId)
		}
		if v, ok := d["calculateFixedTimingHour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateFixedTimingHour)
		}
		if v, ok := d["calculateFixedTimingMinute"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateFixedTimingMinute)
		}
		if v, ok := d["calculateIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateIntervalMinutes)
		}
		if v, ok := d["additionalScopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AdditionalScopes)
		}
		if v, ok := d["ignoreUserIds"]; ok && v != nil {
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
				p.IgnoreUserIds = l
			}
		}
		if v, ok := d["generation"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Generation = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Generation = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Generation = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Generation)
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

func NewCategoryModelMasterFromJson(data string) CategoryModelMaster {
	req := CategoryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCategoryModelMasterFromDict(data map[string]interface{}) CategoryModelMaster {
	return CategoryModelMaster{
		CategoryModelId:            core.CastString(data["categoryModelId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		Sum:                        core.CastBool(data["sum"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		GlobalRankingSetting:       NewGlobalRankingSettingFromDict(core.CastMap(data["globalRankingSetting"])).Pointer(),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		AdditionalScopes:           CastScopes(core.CastArray(data["additionalScopes"])),
		IgnoreUserIds:              core.CastStrings(core.CastArray(data["ignoreUserIds"])),
		Generation:                 core.CastString(data["generation"]),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
		Revision:                   core.CastInt64(data["revision"]),
	}
}

func (p CategoryModelMaster) ToDict() map[string]interface{} {

	var categoryModelId *string
	if p.CategoryModelId != nil {
		categoryModelId = p.CategoryModelId
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
	var scope *string
	if p.Scope != nil {
		scope = p.Scope
	}
	var globalRankingSetting map[string]interface{}
	if p.GlobalRankingSetting != nil {
		globalRankingSetting = p.GlobalRankingSetting.ToDict()
	}
	var entryPeriodEventId *string
	if p.EntryPeriodEventId != nil {
		entryPeriodEventId = p.EntryPeriodEventId
	}
	var accessPeriodEventId *string
	if p.AccessPeriodEventId != nil {
		accessPeriodEventId = p.AccessPeriodEventId
	}
	var uniqueByUserId *bool
	if p.UniqueByUserId != nil {
		uniqueByUserId = p.UniqueByUserId
	}
	var calculateFixedTimingHour *int32
	if p.CalculateFixedTimingHour != nil {
		calculateFixedTimingHour = p.CalculateFixedTimingHour
	}
	var calculateFixedTimingMinute *int32
	if p.CalculateFixedTimingMinute != nil {
		calculateFixedTimingMinute = p.CalculateFixedTimingMinute
	}
	var calculateIntervalMinutes *int32
	if p.CalculateIntervalMinutes != nil {
		calculateIntervalMinutes = p.CalculateIntervalMinutes
	}
	var additionalScopes []interface{}
	if p.AdditionalScopes != nil {
		additionalScopes = CastScopesFromDict(
			p.AdditionalScopes,
		)
	}
	var ignoreUserIds []interface{}
	if p.IgnoreUserIds != nil {
		ignoreUserIds = core.CastStringsFromDict(
			p.IgnoreUserIds,
		)
	}
	var generation *string
	if p.Generation != nil {
		generation = p.Generation
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
		"categoryModelId":            categoryModelId,
		"name":                       name,
		"description":                description,
		"metadata":                   metadata,
		"minimumValue":               minimumValue,
		"maximumValue":               maximumValue,
		"sum":                        sum,
		"orderDirection":             orderDirection,
		"scope":                      scope,
		"globalRankingSetting":       globalRankingSetting,
		"entryPeriodEventId":         entryPeriodEventId,
		"accessPeriodEventId":        accessPeriodEventId,
		"uniqueByUserId":             uniqueByUserId,
		"calculateFixedTimingHour":   calculateFixedTimingHour,
		"calculateFixedTimingMinute": calculateFixedTimingMinute,
		"calculateIntervalMinutes":   calculateIntervalMinutes,
		"additionalScopes":           additionalScopes,
		"ignoreUserIds":              ignoreUserIds,
		"generation":                 generation,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
		"revision":                   revision,
	}
}

func (p CategoryModelMaster) Pointer() *CategoryModelMaster {
	return &p
}

func CastCategoryModelMasters(data []interface{}) []CategoryModelMaster {
	v := make([]CategoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCategoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCategoryModelMastersFromDict(data []CategoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Subscribe struct {
	SubscribeId       *string   `json:"subscribeId"`
	CategoryName      *string   `json:"categoryName"`
	UserId            *string   `json:"userId"`
	TargetUserIds     []*string `json:"targetUserIds"`
	SubscribedUserIds []*string `json:"subscribedUserIds"`
	CreatedAt         *int64    `json:"createdAt"`
	Revision          *int64    `json:"revision"`
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
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
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
		if v, ok := d["subscribedUserIds"]; ok && v != nil {
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
				p.SubscribedUserIds = l
			}
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

func NewSubscribeFromJson(data string) Subscribe {
	req := Subscribe{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
	return Subscribe{
		SubscribeId:       core.CastString(data["subscribeId"]),
		CategoryName:      core.CastString(data["categoryName"]),
		UserId:            core.CastString(data["userId"]),
		TargetUserIds:     core.CastStrings(core.CastArray(data["targetUserIds"])),
		SubscribedUserIds: core.CastStrings(core.CastArray(data["subscribedUserIds"])),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p Subscribe) ToDict() map[string]interface{} {

	var subscribeId *string
	if p.SubscribeId != nil {
		subscribeId = p.SubscribeId
	}
	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
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
	var subscribedUserIds []interface{}
	if p.SubscribedUserIds != nil {
		subscribedUserIds = core.CastStringsFromDict(
			p.SubscribedUserIds,
		)
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"subscribeId":       subscribeId,
		"categoryName":      categoryName,
		"userId":            userId,
		"targetUserIds":     targetUserIds,
		"subscribedUserIds": subscribedUserIds,
		"createdAt":         createdAt,
		"revision":          revision,
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

type Score struct {
	ScoreId      *string `json:"scoreId"`
	CategoryName *string `json:"categoryName"`
	UserId       *string `json:"userId"`
	UniqueId     *string `json:"uniqueId"`
	ScorerUserId *string `json:"scorerUserId"`
	Score        *int64  `json:"score"`
	Metadata     *string `json:"metadata"`
	CreatedAt    *int64  `json:"createdAt"`
	Revision     *int64  `json:"revision"`
}

func (p *Score) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Score{}
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
		*p = Score{}
	} else {
		*p = Score{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["scoreId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScoreId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScoreId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScoreId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScoreId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScoreId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScoreId)
				}
			}
		}
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
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
		if v, ok := d["uniqueId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UniqueId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UniqueId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UniqueId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UniqueId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UniqueId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UniqueId)
				}
			}
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
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewScoreFromJson(data string) Score {
	req := Score{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScoreFromDict(data map[string]interface{}) Score {
	return Score{
		ScoreId:      core.CastString(data["scoreId"]),
		CategoryName: core.CastString(data["categoryName"]),
		UserId:       core.CastString(data["userId"]),
		UniqueId:     core.CastString(data["uniqueId"]),
		ScorerUserId: core.CastString(data["scorerUserId"]),
		Score:        core.CastInt64(data["score"]),
		Metadata:     core.CastString(data["metadata"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		Revision:     core.CastInt64(data["revision"]),
	}
}

func (p Score) ToDict() map[string]interface{} {

	var scoreId *string
	if p.ScoreId != nil {
		scoreId = p.ScoreId
	}
	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var uniqueId *string
	if p.UniqueId != nil {
		uniqueId = p.UniqueId
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"scoreId":      scoreId,
		"categoryName": categoryName,
		"userId":       userId,
		"uniqueId":     uniqueId,
		"scorerUserId": scorerUserId,
		"score":        score,
		"metadata":     metadata,
		"createdAt":    createdAt,
		"revision":     revision,
	}
}

func (p Score) Pointer() *Score {
	return &p
}

func CastScores(data []interface{}) []Score {
	v := make([]Score, 0)
	for _, d := range data {
		v = append(v, NewScoreFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScoresFromDict(data []Score) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Ranking struct {
	Rank         *int64  `json:"rank"`
	Index        *int64  `json:"index"`
	CategoryName *string `json:"categoryName"`
	UserId       *string `json:"userId"`
	Score        *int64  `json:"score"`
	Metadata     *string `json:"metadata"`
	CreatedAt    *int64  `json:"createdAt"`
}

func (p *Ranking) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Ranking{}
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
		*p = Ranking{}
	} else {
		*p = Ranking{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["rank"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rank)
		}
		if v, ok := d["index"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Index)
		}
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
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
	}
	return nil
}

func NewRankingFromJson(data string) Ranking {
	req := Ranking{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRankingFromDict(data map[string]interface{}) Ranking {
	return Ranking{
		Rank:         core.CastInt64(data["rank"]),
		Index:        core.CastInt64(data["index"]),
		CategoryName: core.CastString(data["categoryName"]),
		UserId:       core.CastString(data["userId"]),
		Score:        core.CastInt64(data["score"]),
		Metadata:     core.CastString(data["metadata"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
	}
}

func (p Ranking) ToDict() map[string]interface{} {

	var rank *int64
	if p.Rank != nil {
		rank = p.Rank
	}
	var index *int64
	if p.Index != nil {
		index = p.Index
	}
	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
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
	return map[string]interface{}{
		"rank":         rank,
		"index":        index,
		"categoryName": categoryName,
		"userId":       userId,
		"score":        score,
		"metadata":     metadata,
		"createdAt":    createdAt,
	}
}

func (p Ranking) Pointer() *Ranking {
	return &p
}

func CastRankings(data []interface{}) []Ranking {
	v := make([]Ranking, 0)
	for _, d := range data {
		v = append(v, NewRankingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRankingsFromDict(data []Ranking) []interface{} {
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

type Scope struct {
	Name       *string `json:"name"`
	TargetDays *int64  `json:"targetDays"`
}

func (p *Scope) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Scope{}
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
		*p = Scope{}
	} else {
		*p = Scope{}
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
		if v, ok := d["targetDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetDays)
		}
	}
	return nil
}

func NewScopeFromJson(data string) Scope {
	req := Scope{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScopeFromDict(data map[string]interface{}) Scope {
	return Scope{
		Name:       core.CastString(data["name"]),
		TargetDays: core.CastInt64(data["targetDays"]),
	}
}

func (p Scope) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var targetDays *int64
	if p.TargetDays != nil {
		targetDays = p.TargetDays
	}
	return map[string]interface{}{
		"name":       name,
		"targetDays": targetDays,
	}
}

func (p Scope) Pointer() *Scope {
	return &p
}

func CastScopes(data []interface{}) []Scope {
	v := make([]Scope, 0)
	for _, d := range data {
		v = append(v, NewScopeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScopesFromDict(data []Scope) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GlobalRankingSetting struct {
	UniqueByUserId           *bool        `json:"uniqueByUserId"`
	CalculateIntervalMinutes *int32       `json:"calculateIntervalMinutes"`
	CalculateFixedTiming     *FixedTiming `json:"calculateFixedTiming"`
	AdditionalScopes         []Scope      `json:"additionalScopes"`
	IgnoreUserIds            []*string    `json:"ignoreUserIds"`
	Generation               *string      `json:"generation"`
}

func (p *GlobalRankingSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GlobalRankingSetting{}
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
		*p = GlobalRankingSetting{}
	} else {
		*p = GlobalRankingSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["uniqueByUserId"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UniqueByUserId)
		}
		if v, ok := d["calculateIntervalMinutes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateIntervalMinutes)
		}
		if v, ok := d["calculateFixedTiming"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculateFixedTiming)
		}
		if v, ok := d["additionalScopes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AdditionalScopes)
		}
		if v, ok := d["ignoreUserIds"]; ok && v != nil {
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
				p.IgnoreUserIds = l
			}
		}
		if v, ok := d["generation"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Generation = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Generation = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Generation = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Generation = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Generation)
				}
			}
		}
	}
	return nil
}

func NewGlobalRankingSettingFromJson(data string) GlobalRankingSetting {
	req := GlobalRankingSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGlobalRankingSettingFromDict(data map[string]interface{}) GlobalRankingSetting {
	return GlobalRankingSetting{
		UniqueByUserId:           core.CastBool(data["uniqueByUserId"]),
		CalculateIntervalMinutes: core.CastInt32(data["calculateIntervalMinutes"]),
		CalculateFixedTiming:     NewFixedTimingFromDict(core.CastMap(data["calculateFixedTiming"])).Pointer(),
		AdditionalScopes:         CastScopes(core.CastArray(data["additionalScopes"])),
		IgnoreUserIds:            core.CastStrings(core.CastArray(data["ignoreUserIds"])),
		Generation:               core.CastString(data["generation"]),
	}
}

func (p GlobalRankingSetting) ToDict() map[string]interface{} {

	var uniqueByUserId *bool
	if p.UniqueByUserId != nil {
		uniqueByUserId = p.UniqueByUserId
	}
	var calculateIntervalMinutes *int32
	if p.CalculateIntervalMinutes != nil {
		calculateIntervalMinutes = p.CalculateIntervalMinutes
	}
	var calculateFixedTiming map[string]interface{}
	if p.CalculateFixedTiming != nil {
		calculateFixedTiming = p.CalculateFixedTiming.ToDict()
	}
	var additionalScopes []interface{}
	if p.AdditionalScopes != nil {
		additionalScopes = CastScopesFromDict(
			p.AdditionalScopes,
		)
	}
	var ignoreUserIds []interface{}
	if p.IgnoreUserIds != nil {
		ignoreUserIds = core.CastStringsFromDict(
			p.IgnoreUserIds,
		)
	}
	var generation *string
	if p.Generation != nil {
		generation = p.Generation
	}
	return map[string]interface{}{
		"uniqueByUserId":           uniqueByUserId,
		"calculateIntervalMinutes": calculateIntervalMinutes,
		"calculateFixedTiming":     calculateFixedTiming,
		"additionalScopes":         additionalScopes,
		"ignoreUserIds":            ignoreUserIds,
		"generation":               generation,
	}
}

func (p GlobalRankingSetting) Pointer() *GlobalRankingSetting {
	return &p
}

func CastGlobalRankingSettings(data []interface{}) []GlobalRankingSetting {
	v := make([]GlobalRankingSetting, 0)
	for _, d := range data {
		v = append(v, NewGlobalRankingSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGlobalRankingSettingsFromDict(data []GlobalRankingSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FixedTiming struct {
	Hour   *int32 `json:"hour"`
	Minute *int32 `json:"minute"`
}

func (p *FixedTiming) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FixedTiming{}
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
		*p = FixedTiming{}
	} else {
		*p = FixedTiming{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["hour"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Hour)
		}
		if v, ok := d["minute"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Minute)
		}
	}
	return nil
}

func NewFixedTimingFromJson(data string) FixedTiming {
	req := FixedTiming{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFixedTimingFromDict(data map[string]interface{}) FixedTiming {
	return FixedTiming{
		Hour:   core.CastInt32(data["hour"]),
		Minute: core.CastInt32(data["minute"]),
	}
}

func (p FixedTiming) ToDict() map[string]interface{} {

	var hour *int32
	if p.Hour != nil {
		hour = p.Hour
	}
	var minute *int32
	if p.Minute != nil {
		minute = p.Minute
	}
	return map[string]interface{}{
		"hour":   hour,
		"minute": minute,
	}
}

func (p FixedTiming) Pointer() *FixedTiming {
	return &p
}

func CastFixedTimings(data []interface{}) []FixedTiming {
	v := make([]FixedTiming, 0)
	for _, d := range data {
		v = append(v, NewFixedTimingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFixedTimingsFromDict(data []FixedTiming) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CalculatedAt struct {
	CategoryName *string `json:"categoryName"`
	CalculatedAt *int64  `json:"calculatedAt"`
}

func (p *CalculatedAt) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CalculatedAt{}
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
		*p = CalculatedAt{}
	} else {
		*p = CalculatedAt{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
				}
			}
		}
		if v, ok := d["calculatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CalculatedAt)
		}
	}
	return nil
}

func NewCalculatedAtFromJson(data string) CalculatedAt {
	req := CalculatedAt{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCalculatedAtFromDict(data map[string]interface{}) CalculatedAt {
	return CalculatedAt{
		CategoryName: core.CastString(data["categoryName"]),
		CalculatedAt: core.CastInt64(data["calculatedAt"]),
	}
}

func (p CalculatedAt) ToDict() map[string]interface{} {

	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
	}
	var calculatedAt *int64
	if p.CalculatedAt != nil {
		calculatedAt = p.CalculatedAt
	}
	return map[string]interface{}{
		"categoryName": categoryName,
		"calculatedAt": calculatedAt,
	}
}

func (p CalculatedAt) Pointer() *CalculatedAt {
	return &p
}

func CastCalculatedAts(data []interface{}) []CalculatedAt {
	v := make([]CalculatedAt, 0)
	for _, d := range data {
		v = append(v, NewCalculatedAtFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCalculatedAtsFromDict(data []CalculatedAt) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SubscribeUser struct {
	SubscribeUserId *string `json:"subscribeUserId"`
	CategoryName    *string `json:"categoryName"`
	UserId          *string `json:"userId"`
	TargetUserId    *string `json:"targetUserId"`
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
		if v, ok := d["subscribeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SubscribeUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SubscribeUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SubscribeUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SubscribeUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SubscribeUserId)
				}
			}
		}
		if v, ok := d["categoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CategoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CategoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CategoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CategoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CategoryName)
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
		SubscribeUserId: core.CastString(data["subscribeUserId"]),
		CategoryName:    core.CastString(data["categoryName"]),
		UserId:          core.CastString(data["userId"]),
		TargetUserId:    core.CastString(data["targetUserId"]),
	}
}

func (p SubscribeUser) ToDict() map[string]interface{} {

	var subscribeUserId *string
	if p.SubscribeUserId != nil {
		subscribeUserId = p.SubscribeUserId
	}
	var categoryName *string
	if p.CategoryName != nil {
		categoryName = p.CategoryName
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
		"subscribeUserId": subscribeUserId,
		"categoryName":    categoryName,
		"userId":          userId,
		"targetUserId":    targetUserId,
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
