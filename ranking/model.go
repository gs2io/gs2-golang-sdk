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
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	LastCalculatedAts []CalculatedAt `json:"lastCalculatedAts"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
    return Namespace {
        NamespaceId: core.CastString(data["namespaceId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        LastCalculatedAts: CastCalculatedAts(core.CastArray(data["lastCalculatedAts"])),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "lastCalculatedAts": lastCalculatedAts,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	CategoryModelId *string `json:"categoryModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	MinimumValue *int64 `json:"minimumValue"`
	MaximumValue *int64 `json:"maximumValue"`
	OrderDirection *string `json:"orderDirection"`
	Scope *string `json:"scope"`
	UniqueByUserId *bool `json:"uniqueByUserId"`
	CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
	EntryPeriodEventId *string `json:"entryPeriodEventId"`
	AccessPeriodEventId *string `json:"accessPeriodEventId"`
	Generation *string `json:"generation"`
}

func NewCategoryModelFromJson(data string) CategoryModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCategoryModelFromDict(dict)
}

func NewCategoryModelFromDict(data map[string]interface{}) CategoryModel {
    return CategoryModel {
        CategoryModelId: core.CastString(data["categoryModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        MinimumValue: core.CastInt64(data["minimumValue"]),
        MaximumValue: core.CastInt64(data["maximumValue"]),
        OrderDirection: core.CastString(data["orderDirection"]),
        Scope: core.CastString(data["scope"]),
        UniqueByUserId: core.CastBool(data["uniqueByUserId"]),
        CalculateFixedTimingHour: core.CastInt32(data["calculateFixedTimingHour"]),
        CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
        CalculateIntervalMinutes: core.CastInt32(data["calculateIntervalMinutes"]),
        EntryPeriodEventId: core.CastString(data["entryPeriodEventId"]),
        AccessPeriodEventId: core.CastString(data["accessPeriodEventId"]),
        Generation: core.CastString(data["generation"]),
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
    var orderDirection *string
    if p.OrderDirection != nil {
        orderDirection = p.OrderDirection
    }
    var scope *string
    if p.Scope != nil {
        scope = p.Scope
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
    var entryPeriodEventId *string
    if p.EntryPeriodEventId != nil {
        entryPeriodEventId = p.EntryPeriodEventId
    }
    var accessPeriodEventId *string
    if p.AccessPeriodEventId != nil {
        accessPeriodEventId = p.AccessPeriodEventId
    }
    var generation *string
    if p.Generation != nil {
        generation = p.Generation
    }
    return map[string]interface{} {
        "categoryModelId": categoryModelId,
        "name": name,
        "metadata": metadata,
        "minimumValue": minimumValue,
        "maximumValue": maximumValue,
        "orderDirection": orderDirection,
        "scope": scope,
        "uniqueByUserId": uniqueByUserId,
        "calculateFixedTimingHour": calculateFixedTimingHour,
        "calculateFixedTimingMinute": calculateFixedTimingMinute,
        "calculateIntervalMinutes": calculateIntervalMinutes,
        "entryPeriodEventId": entryPeriodEventId,
        "accessPeriodEventId": accessPeriodEventId,
        "generation": generation,
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
	CategoryModelId *string `json:"categoryModelId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	MinimumValue *int64 `json:"minimumValue"`
	MaximumValue *int64 `json:"maximumValue"`
	OrderDirection *string `json:"orderDirection"`
	Scope *string `json:"scope"`
	UniqueByUserId *bool `json:"uniqueByUserId"`
	CalculateFixedTimingHour *int32 `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32 `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes *int32 `json:"calculateIntervalMinutes"`
	EntryPeriodEventId *string `json:"entryPeriodEventId"`
	AccessPeriodEventId *string `json:"accessPeriodEventId"`
	Generation *string `json:"generation"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCategoryModelMasterFromJson(data string) CategoryModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCategoryModelMasterFromDict(dict)
}

func NewCategoryModelMasterFromDict(data map[string]interface{}) CategoryModelMaster {
    return CategoryModelMaster {
        CategoryModelId: core.CastString(data["categoryModelId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        MinimumValue: core.CastInt64(data["minimumValue"]),
        MaximumValue: core.CastInt64(data["maximumValue"]),
        OrderDirection: core.CastString(data["orderDirection"]),
        Scope: core.CastString(data["scope"]),
        UniqueByUserId: core.CastBool(data["uniqueByUserId"]),
        CalculateFixedTimingHour: core.CastInt32(data["calculateFixedTimingHour"]),
        CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
        CalculateIntervalMinutes: core.CastInt32(data["calculateIntervalMinutes"]),
        EntryPeriodEventId: core.CastString(data["entryPeriodEventId"]),
        AccessPeriodEventId: core.CastString(data["accessPeriodEventId"]),
        Generation: core.CastString(data["generation"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    var orderDirection *string
    if p.OrderDirection != nil {
        orderDirection = p.OrderDirection
    }
    var scope *string
    if p.Scope != nil {
        scope = p.Scope
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
    var entryPeriodEventId *string
    if p.EntryPeriodEventId != nil {
        entryPeriodEventId = p.EntryPeriodEventId
    }
    var accessPeriodEventId *string
    if p.AccessPeriodEventId != nil {
        accessPeriodEventId = p.AccessPeriodEventId
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
    return map[string]interface{} {
        "categoryModelId": categoryModelId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "minimumValue": minimumValue,
        "maximumValue": maximumValue,
        "orderDirection": orderDirection,
        "scope": scope,
        "uniqueByUserId": uniqueByUserId,
        "calculateFixedTimingHour": calculateFixedTimingHour,
        "calculateFixedTimingMinute": calculateFixedTimingMinute,
        "calculateIntervalMinutes": calculateIntervalMinutes,
        "entryPeriodEventId": entryPeriodEventId,
        "accessPeriodEventId": accessPeriodEventId,
        "generation": generation,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	SubscribeId *string `json:"subscribeId"`
	CategoryName *string `json:"categoryName"`
	UserId *string `json:"userId"`
	TargetUserIds []string `json:"targetUserIds"`
	SubscribedUserIds []string `json:"subscribedUserIds"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewSubscribeFromJson(data string) Subscribe {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSubscribeFromDict(dict)
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
    return Subscribe {
        SubscribeId: core.CastString(data["subscribeId"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
        SubscribedUserIds: core.CastStrings(core.CastArray(data["subscribedUserIds"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
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
    return map[string]interface{} {
        "subscribeId": subscribeId,
        "categoryName": categoryName,
        "userId": userId,
        "targetUserIds": targetUserIds,
        "subscribedUserIds": subscribedUserIds,
        "createdAt": createdAt,
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
	ScoreId *string `json:"scoreId"`
	CategoryName *string `json:"categoryName"`
	UserId *string `json:"userId"`
	UniqueId *string `json:"uniqueId"`
	ScorerUserId *string `json:"scorerUserId"`
	Score *int64 `json:"score"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewScoreFromJson(data string) Score {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScoreFromDict(dict)
}

func NewScoreFromDict(data map[string]interface{}) Score {
    return Score {
        ScoreId: core.CastString(data["scoreId"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        UniqueId: core.CastString(data["uniqueId"]),
        ScorerUserId: core.CastString(data["scorerUserId"]),
        Score: core.CastInt64(data["score"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
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
    return map[string]interface{} {
        "scoreId": scoreId,
        "categoryName": categoryName,
        "userId": userId,
        "uniqueId": uniqueId,
        "scorerUserId": scorerUserId,
        "score": score,
        "metadata": metadata,
        "createdAt": createdAt,
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
	Rank *int64 `json:"rank"`
	Index *int64 `json:"index"`
	CategoryName *string `json:"categoryName"`
	UserId *string `json:"userId"`
	Score *int64 `json:"score"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewRankingFromJson(data string) Ranking {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRankingFromDict(dict)
}

func NewRankingFromDict(data map[string]interface{}) Ranking {
    return Ranking {
        Rank: core.CastInt64(data["rank"]),
        Index: core.CastInt64(data["index"]),
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        Score: core.CastInt64(data["score"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
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
    return map[string]interface{} {
        "rank": rank,
        "index": index,
        "categoryName": categoryName,
        "userId": userId,
        "score": score,
        "metadata": metadata,
        "createdAt": createdAt,
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
	Settings *string `json:"settings"`
}

func NewCurrentRankingMasterFromJson(data string) CurrentRankingMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentRankingMasterFromDict(dict)
}

func NewCurrentRankingMasterFromDict(data map[string]interface{}) CurrentRankingMaster {
    return CurrentRankingMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
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
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "settings": settings,
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

type CalculatedAt struct {
	CategoryName *string `json:"categoryName"`
	CalculatedAt *int64 `json:"calculatedAt"`
}

func NewCalculatedAtFromJson(data string) CalculatedAt {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCalculatedAtFromDict(dict)
}

func NewCalculatedAtFromDict(data map[string]interface{}) CalculatedAt {
    return CalculatedAt {
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
    return map[string]interface{} {
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
	CategoryName *string `json:"categoryName"`
	UserId *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func NewSubscribeUserFromJson(data string) SubscribeUser {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSubscribeUserFromDict(dict)
}

func NewSubscribeUserFromDict(data map[string]interface{}) SubscribeUser {
    return SubscribeUser {
        CategoryName: core.CastString(data["categoryName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p SubscribeUser) ToDict() map[string]interface{} {
    
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
    return map[string]interface{} {
        "categoryName": categoryName,
        "userId": userId,
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

type GitHubCheckoutSetting struct {
	ApiKeyId *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath *string `json:"sourcePath"`
	ReferenceType *string `json:"referenceType"`
	CommitHash *string `json:"commitHash"`
	BranchName *string `json:"branchName"`
	TagName *string `json:"tagName"`
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGitHubCheckoutSettingFromDict(dict)
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
    return GitHubCheckoutSetting {
        ApiKeyId: core.CastString(data["apiKeyId"]),
        RepositoryName: core.CastString(data["repositoryName"]),
        SourcePath: core.CastString(data["sourcePath"]),
        ReferenceType: core.CastString(data["referenceType"]),
        CommitHash: core.CastString(data["commitHash"]),
        BranchName: core.CastString(data["branchName"]),
        TagName: core.CastString(data["tagName"]),
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
    return map[string]interface{} {
        "apiKeyId": apiKeyId,
        "repositoryName": repositoryName,
        "sourcePath": sourcePath,
        "referenceType": referenceType,
        "commitHash": commitHash,
        "branchName": branchName,
        "tagName": tagName,
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

func NewLogSettingFromJson(data string) LogSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    
    var loggingNamespaceId *string
    if p.LoggingNamespaceId != nil {
        loggingNamespaceId = p.LoggingNamespaceId
    }
    return map[string]interface{} {
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