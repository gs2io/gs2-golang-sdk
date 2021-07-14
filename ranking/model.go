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

import "core"

type Namespace struct {
	NamespaceId       *string        `json:"namespaceId"`
	Name              *string        `json:"name"`
	Description       *string        `json:"description"`
	LastCalculatedAts []CalculatedAt `json:"lastCalculatedAts"`
	LogSetting        *LogSetting    `json:"logSetting"`
	CreatedAt         *int64         `json:"createdAt"`
	UpdatedAt         *int64         `json:"updatedAt"`
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
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"name":        p.Name,
		"description": p.Description,
		"lastCalculatedAts": CastCalculatedAtsFromDict(
			p.LastCalculatedAts,
		),
		"logSetting": p.LogSetting.ToDict(),
		"createdAt":  p.CreatedAt,
		"updatedAt":  p.UpdatedAt,
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
	CategoryModelId            *string `json:"categoryModelId"`
	Name                       *string `json:"name"`
	Metadata                   *string `json:"metadata"`
	MinimumValue               *int64  `json:"minimumValue"`
	MaximumValue               *int64  `json:"maximumValue"`
	OrderDirection             *string `json:"orderDirection"`
	Scope                      *string `json:"scope"`
	UniqueByUserId             *bool   `json:"uniqueByUserId"`
	CalculateFixedTimingHour   *int32  `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32  `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes   *int32  `json:"calculateIntervalMinutes"`
	EntryPeriodEventId         *string `json:"entryPeriodEventId"`
	AccessPeriodEventId        *string `json:"accessPeriodEventId"`
	Generation                 *string `json:"generation"`
}

func NewCategoryModelFromDict(data map[string]interface{}) CategoryModel {
	return CategoryModel{
		CategoryModelId:            core.CastString(data["categoryModelId"]),
		Name:                       core.CastString(data["name"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		Generation:                 core.CastString(data["generation"]),
	}
}

func (p CategoryModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"categoryModelId":            p.CategoryModelId,
		"name":                       p.Name,
		"metadata":                   p.Metadata,
		"minimumValue":               p.MinimumValue,
		"maximumValue":               p.MaximumValue,
		"orderDirection":             p.OrderDirection,
		"scope":                      p.Scope,
		"uniqueByUserId":             p.UniqueByUserId,
		"calculateFixedTimingHour":   p.CalculateFixedTimingHour,
		"calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
		"calculateIntervalMinutes":   p.CalculateIntervalMinutes,
		"entryPeriodEventId":         p.EntryPeriodEventId,
		"accessPeriodEventId":        p.AccessPeriodEventId,
		"generation":                 p.Generation,
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
	CategoryModelId            *string `json:"categoryModelId"`
	Name                       *string `json:"name"`
	Description                *string `json:"description"`
	Metadata                   *string `json:"metadata"`
	MinimumValue               *int64  `json:"minimumValue"`
	MaximumValue               *int64  `json:"maximumValue"`
	OrderDirection             *string `json:"orderDirection"`
	Scope                      *string `json:"scope"`
	UniqueByUserId             *bool   `json:"uniqueByUserId"`
	CalculateFixedTimingHour   *int32  `json:"calculateFixedTimingHour"`
	CalculateFixedTimingMinute *int32  `json:"calculateFixedTimingMinute"`
	CalculateIntervalMinutes   *int32  `json:"calculateIntervalMinutes"`
	EntryPeriodEventId         *string `json:"entryPeriodEventId"`
	AccessPeriodEventId        *string `json:"accessPeriodEventId"`
	Generation                 *string `json:"generation"`
	CreatedAt                  *int64  `json:"createdAt"`
	UpdatedAt                  *int64  `json:"updatedAt"`
}

func NewCategoryModelMasterFromDict(data map[string]interface{}) CategoryModelMaster {
	return CategoryModelMaster{
		CategoryModelId:            core.CastString(data["categoryModelId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		MinimumValue:               core.CastInt64(data["minimumValue"]),
		MaximumValue:               core.CastInt64(data["maximumValue"]),
		OrderDirection:             core.CastString(data["orderDirection"]),
		Scope:                      core.CastString(data["scope"]),
		UniqueByUserId:             core.CastBool(data["uniqueByUserId"]),
		CalculateFixedTimingHour:   core.CastInt32(data["calculateFixedTimingHour"]),
		CalculateFixedTimingMinute: core.CastInt32(data["calculateFixedTimingMinute"]),
		CalculateIntervalMinutes:   core.CastInt32(data["calculateIntervalMinutes"]),
		EntryPeriodEventId:         core.CastString(data["entryPeriodEventId"]),
		AccessPeriodEventId:        core.CastString(data["accessPeriodEventId"]),
		Generation:                 core.CastString(data["generation"]),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
	}
}

func (p CategoryModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"categoryModelId":            p.CategoryModelId,
		"name":                       p.Name,
		"description":                p.Description,
		"metadata":                   p.Metadata,
		"minimumValue":               p.MinimumValue,
		"maximumValue":               p.MaximumValue,
		"orderDirection":             p.OrderDirection,
		"scope":                      p.Scope,
		"uniqueByUserId":             p.UniqueByUserId,
		"calculateFixedTimingHour":   p.CalculateFixedTimingHour,
		"calculateFixedTimingMinute": p.CalculateFixedTimingMinute,
		"calculateIntervalMinutes":   p.CalculateIntervalMinutes,
		"entryPeriodEventId":         p.EntryPeriodEventId,
		"accessPeriodEventId":        p.AccessPeriodEventId,
		"generation":                 p.Generation,
		"createdAt":                  p.CreatedAt,
		"updatedAt":                  p.UpdatedAt,
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
	SubscribeId   *string  `json:"subscribeId"`
	CategoryName  *string  `json:"categoryName"`
	UserId        *string  `json:"userId"`
	TargetUserIds []string `json:"targetUserIds"`
	CreatedAt     *int64   `json:"createdAt"`
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
	return Subscribe{
		SubscribeId:   core.CastString(data["subscribeId"]),
		CategoryName:  core.CastString(data["categoryName"]),
		UserId:        core.CastString(data["userId"]),
		TargetUserIds: core.CastStrings(core.CastArray(data["targetUserIds"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
	}
}

func (p Subscribe) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"subscribeId":  p.SubscribeId,
		"categoryName": p.CategoryName,
		"userId":       p.UserId,
		"targetUserIds": core.CastStringsFromDict(
			p.TargetUserIds,
		),
		"createdAt": p.CreatedAt,
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
	}
}

func (p Score) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"scoreId":      p.ScoreId,
		"categoryName": p.CategoryName,
		"userId":       p.UserId,
		"uniqueId":     p.UniqueId,
		"scorerUserId": p.ScorerUserId,
		"score":        p.Score,
		"metadata":     p.Metadata,
		"createdAt":    p.CreatedAt,
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
	Rank      *int64  `json:"rank"`
	Index     *int64  `json:"index"`
	UserId    *string `json:"userId"`
	Score     *int64  `json:"score"`
	Metadata  *string `json:"metadata"`
	CreatedAt *int64  `json:"createdAt"`
}

func NewRankingFromDict(data map[string]interface{}) Ranking {
	return Ranking{
		Rank:      core.CastInt64(data["rank"]),
		Index:     core.CastInt64(data["index"]),
		UserId:    core.CastString(data["userId"]),
		Score:     core.CastInt64(data["score"]),
		Metadata:  core.CastString(data["metadata"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
	}
}

func (p Ranking) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rank":      p.Rank,
		"index":     p.Index,
		"userId":    p.UserId,
		"score":     p.Score,
		"metadata":  p.Metadata,
		"createdAt": p.CreatedAt,
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

func NewCurrentRankingMasterFromDict(data map[string]interface{}) CurrentRankingMaster {
	return CurrentRankingMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRankingMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
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
	CalculatedAt *int64  `json:"calculatedAt"`
}

func NewCalculatedAtFromDict(data map[string]interface{}) CalculatedAt {
	return CalculatedAt{
		CategoryName: core.CastString(data["categoryName"]),
		CalculatedAt: core.CastInt64(data["calculatedAt"]),
	}
}

func (p CalculatedAt) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"categoryName": p.CategoryName,
		"calculatedAt": p.CalculatedAt,
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
	UserId       *string `json:"userId"`
	TargetUserId *string `json:"targetUserId"`
}

func NewSubscribeUserFromDict(data map[string]interface{}) SubscribeUser {
	return SubscribeUser{
		CategoryName: core.CastString(data["categoryName"]),
		UserId:       core.CastString(data["userId"]),
		TargetUserId: core.CastString(data["targetUserId"]),
	}
}

func (p SubscribeUser) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"categoryName": p.CategoryName,
		"userId":       p.UserId,
		"targetUserId": p.TargetUserId,
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
	return map[string]interface{}{
		"apiKeyId":       p.ApiKeyId,
		"repositoryName": p.RepositoryName,
		"sourcePath":     p.SourcePath,
		"referenceType":  p.ReferenceType,
		"commitHash":     p.CommitHash,
		"branchName":     p.BranchName,
		"tagName":        p.TagName,
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

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
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
