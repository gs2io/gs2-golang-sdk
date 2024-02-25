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

package seasonRating

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string     `json:"namespaceId"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	LogSetting  *LogSetting `json:"logSetting"`
	CreatedAt   *int64      `json:"createdAt"`
	UpdatedAt   *int64      `json:"updatedAt"`
	Revision    *int64      `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId: core.CastString(data["namespaceId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
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
		"namespaceId": namespaceId,
		"name":        name,
		"description": description,
		"logSetting":  logSetting,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
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

type MatchSession struct {
	SessionId *string `json:"sessionId"`
	Name      *string `json:"name"`
	CreatedAt *int64  `json:"createdAt"`
	Revision  *int64  `json:"revision"`
}

func NewMatchSessionFromJson(data string) MatchSession {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMatchSessionFromDict(dict)
}

func NewMatchSessionFromDict(data map[string]interface{}) MatchSession {
	return MatchSession{
		SessionId: core.CastString(data["sessionId"]),
		Name:      core.CastString(data["name"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p MatchSession) ToDict() map[string]interface{} {

	var sessionId *string
	if p.SessionId != nil {
		sessionId = p.SessionId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
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
		"sessionId": sessionId,
		"name":      name,
		"createdAt": createdAt,
		"revision":  revision,
	}
}

func (p MatchSession) Pointer() *MatchSession {
	return &p
}

func CastMatchSessions(data []interface{}) []MatchSession {
	v := make([]MatchSession, 0)
	for _, d := range data {
		v = append(v, NewMatchSessionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMatchSessionsFromDict(data []MatchSession) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SeasonModelMaster struct {
	SeasonModelId     *string     `json:"seasonModelId"`
	Name              *string     `json:"name"`
	Metadata          *string     `json:"metadata"`
	Description       *string     `json:"description"`
	Tiers             []TierModel `json:"tiers"`
	ExperienceModelId *string     `json:"experienceModelId"`
	CreatedAt         *int64      `json:"createdAt"`
	UpdatedAt         *int64      `json:"updatedAt"`
	Revision          *int64      `json:"revision"`
}

func NewSeasonModelMasterFromJson(data string) SeasonModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSeasonModelMasterFromDict(dict)
}

func NewSeasonModelMasterFromDict(data map[string]interface{}) SeasonModelMaster {
	return SeasonModelMaster{
		SeasonModelId:     core.CastString(data["seasonModelId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		Description:       core.CastString(data["description"]),
		Tiers:             CastTierModels(core.CastArray(data["tiers"])),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p SeasonModelMaster) ToDict() map[string]interface{} {

	var seasonModelId *string
	if p.SeasonModelId != nil {
		seasonModelId = p.SeasonModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var tiers []interface{}
	if p.Tiers != nil {
		tiers = CastTierModelsFromDict(
			p.Tiers,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
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
		"seasonModelId":     seasonModelId,
		"name":              name,
		"metadata":          metadata,
		"description":       description,
		"tiers":             tiers,
		"experienceModelId": experienceModelId,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
	}
}

func (p SeasonModelMaster) Pointer() *SeasonModelMaster {
	return &p
}

func CastSeasonModelMasters(data []interface{}) []SeasonModelMaster {
	v := make([]SeasonModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSeasonModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSeasonModelMastersFromDict(data []SeasonModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SeasonModel struct {
	SeasonModelId     *string     `json:"seasonModelId"`
	Name              *string     `json:"name"`
	Metadata          *string     `json:"metadata"`
	Tiers             []TierModel `json:"tiers"`
	ExperienceModelId *string     `json:"experienceModelId"`
}

func NewSeasonModelFromJson(data string) SeasonModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSeasonModelFromDict(dict)
}

func NewSeasonModelFromDict(data map[string]interface{}) SeasonModel {
	return SeasonModel{
		SeasonModelId:     core.CastString(data["seasonModelId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		Tiers:             CastTierModels(core.CastArray(data["tiers"])),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
	}
}

func (p SeasonModel) ToDict() map[string]interface{} {

	var seasonModelId *string
	if p.SeasonModelId != nil {
		seasonModelId = p.SeasonModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var tiers []interface{}
	if p.Tiers != nil {
		tiers = CastTierModelsFromDict(
			p.Tiers,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	return map[string]interface{}{
		"seasonModelId":     seasonModelId,
		"name":              name,
		"metadata":          metadata,
		"tiers":             tiers,
		"experienceModelId": experienceModelId,
	}
}

func (p SeasonModel) Pointer() *SeasonModel {
	return &p
}

func CastSeasonModels(data []interface{}) []SeasonModel {
	v := make([]SeasonModel, 0)
	for _, d := range data {
		v = append(v, NewSeasonModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSeasonModelsFromDict(data []SeasonModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TierModel struct {
	Metadata           *string `json:"metadata"`
	RaiseRankBonus     *int32  `json:"raiseRankBonus"`
	EntryFee           *int32  `json:"entryFee"`
	MinimumChangePoint *int32  `json:"minimumChangePoint"`
	MaximumChangePoint *int32  `json:"maximumChangePoint"`
}

func NewTierModelFromJson(data string) TierModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTierModelFromDict(dict)
}

func NewTierModelFromDict(data map[string]interface{}) TierModel {
	return TierModel{
		Metadata:           core.CastString(data["metadata"]),
		RaiseRankBonus:     core.CastInt32(data["raiseRankBonus"]),
		EntryFee:           core.CastInt32(data["entryFee"]),
		MinimumChangePoint: core.CastInt32(data["minimumChangePoint"]),
		MaximumChangePoint: core.CastInt32(data["maximumChangePoint"]),
	}
}

func (p TierModel) ToDict() map[string]interface{} {

	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var raiseRankBonus *int32
	if p.RaiseRankBonus != nil {
		raiseRankBonus = p.RaiseRankBonus
	}
	var entryFee *int32
	if p.EntryFee != nil {
		entryFee = p.EntryFee
	}
	var minimumChangePoint *int32
	if p.MinimumChangePoint != nil {
		minimumChangePoint = p.MinimumChangePoint
	}
	var maximumChangePoint *int32
	if p.MaximumChangePoint != nil {
		maximumChangePoint = p.MaximumChangePoint
	}
	return map[string]interface{}{
		"metadata":           metadata,
		"raiseRankBonus":     raiseRankBonus,
		"entryFee":           entryFee,
		"minimumChangePoint": minimumChangePoint,
		"maximumChangePoint": maximumChangePoint,
	}
}

func (p TierModel) Pointer() *TierModel {
	return &p
}

func CastTierModels(data []interface{}) []TierModel {
	v := make([]TierModel, 0)
	for _, d := range data {
		v = append(v, NewTierModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTierModelsFromDict(data []TierModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentSeasonModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentSeasonModelMasterFromJson(data string) CurrentSeasonModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentSeasonModelMasterFromDict(dict)
}

func NewCurrentSeasonModelMasterFromDict(data map[string]interface{}) CurrentSeasonModelMaster {
	return CurrentSeasonModelMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentSeasonModelMaster) ToDict() map[string]interface{} {

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

func (p CurrentSeasonModelMaster) Pointer() *CurrentSeasonModelMaster {
	return &p
}

func CastCurrentSeasonModelMasters(data []interface{}) []CurrentSeasonModelMaster {
	v := make([]CurrentSeasonModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentSeasonModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentSeasonModelMastersFromDict(data []CurrentSeasonModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	KeyId                  *string `json:"keyId"`
	QueueNamespaceId       *string `json:"queueNamespaceId"`
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTransactionSettingFromDict(dict)
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

func NewScriptSettingFromJson(data string) ScriptSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewScriptSettingFromDict(dict)
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

func NewLogSettingFromJson(data string) LogSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLogSettingFromDict(dict)
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

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGitHubCheckoutSettingFromDict(dict)
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

type GameResult struct {
	Rank   *int32  `json:"rank"`
	UserId *string `json:"userId"`
}

func NewGameResultFromJson(data string) GameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGameResultFromDict(dict)
}

func NewGameResultFromDict(data map[string]interface{}) GameResult {
	return GameResult{
		Rank:   core.CastInt32(data["rank"]),
		UserId: core.CastString(data["userId"]),
	}
}

func (p GameResult) ToDict() map[string]interface{} {

	var rank *int32
	if p.Rank != nil {
		rank = p.Rank
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	return map[string]interface{}{
		"rank":   rank,
		"userId": userId,
	}
}

func (p GameResult) Pointer() *GameResult {
	return &p
}

func CastGameResults(data []interface{}) []GameResult {
	v := make([]GameResult, 0)
	for _, d := range data {
		v = append(v, NewGameResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGameResultsFromDict(data []GameResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Ballot struct {
	UserId         *string `json:"userId"`
	SeasonName     *string `json:"seasonName"`
	SessionName    *string `json:"sessionName"`
	NumberOfPlayer *int32  `json:"numberOfPlayer"`
}

func NewBallotFromJson(data string) Ballot {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBallotFromDict(dict)
}

func NewBallotFromDict(data map[string]interface{}) Ballot {
	return Ballot{
		UserId:         core.CastString(data["userId"]),
		SeasonName:     core.CastString(data["seasonName"]),
		SessionName:    core.CastString(data["sessionName"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
	}
}

func (p Ballot) ToDict() map[string]interface{} {

	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var seasonName *string
	if p.SeasonName != nil {
		seasonName = p.SeasonName
	}
	var sessionName *string
	if p.SessionName != nil {
		sessionName = p.SessionName
	}
	var numberOfPlayer *int32
	if p.NumberOfPlayer != nil {
		numberOfPlayer = p.NumberOfPlayer
	}
	return map[string]interface{}{
		"userId":         userId,
		"seasonName":     seasonName,
		"sessionName":    sessionName,
		"numberOfPlayer": numberOfPlayer,
	}
}

func (p Ballot) Pointer() *Ballot {
	return &p
}

func CastBallots(data []interface{}) []Ballot {
	v := make([]Ballot, 0)
	for _, d := range data {
		v = append(v, NewBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBallotsFromDict(data []Ballot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SignedBallot struct {
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

func NewSignedBallotFromJson(data string) SignedBallot {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSignedBallotFromDict(dict)
}

func NewSignedBallotFromDict(data map[string]interface{}) SignedBallot {
	return SignedBallot{
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p SignedBallot) ToDict() map[string]interface{} {

	var body *string
	if p.Body != nil {
		body = p.Body
	}
	var signature *string
	if p.Signature != nil {
		signature = p.Signature
	}
	return map[string]interface{}{
		"body":      body,
		"signature": signature,
	}
}

func (p SignedBallot) Pointer() *SignedBallot {
	return &p
}

func CastSignedBallots(data []interface{}) []SignedBallot {
	v := make([]SignedBallot, 0)
	for _, d := range data {
		v = append(v, NewSignedBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignedBallotsFromDict(data []SignedBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type WrittenBallot struct {
	Ballot      *Ballot      `json:"ballot"`
	GameResults []GameResult `json:"gameResults"`
}

func NewWrittenBallotFromJson(data string) WrittenBallot {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWrittenBallotFromDict(dict)
}

func NewWrittenBallotFromDict(data map[string]interface{}) WrittenBallot {
	return WrittenBallot{
		Ballot:      NewBallotFromDict(core.CastMap(data["ballot"])).Pointer(),
		GameResults: CastGameResults(core.CastArray(data["gameResults"])),
	}
}

func (p WrittenBallot) ToDict() map[string]interface{} {

	var ballot map[string]interface{}
	if p.Ballot != nil {
		ballot = p.Ballot.ToDict()
	}
	var gameResults []interface{}
	if p.GameResults != nil {
		gameResults = CastGameResultsFromDict(
			p.GameResults,
		)
	}
	return map[string]interface{}{
		"ballot":      ballot,
		"gameResults": gameResults,
	}
}

func (p WrittenBallot) Pointer() *WrittenBallot {
	return &p
}

func CastWrittenBallots(data []interface{}) []WrittenBallot {
	v := make([]WrittenBallot, 0)
	for _, d := range data {
		v = append(v, NewWrittenBallotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastWrittenBallotsFromDict(data []WrittenBallot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Vote struct {
	VoteId         *string         `json:"voteId"`
	SeasonName     *string         `json:"seasonName"`
	SessionName    *string         `json:"sessionName"`
	WrittenBallots []WrittenBallot `json:"writtenBallots"`
	CreatedAt      *int64          `json:"createdAt"`
	UpdatedAt      *int64          `json:"updatedAt"`
	Revision       *int64          `json:"revision"`
}

func NewVoteFromJson(data string) Vote {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteFromDict(dict)
}

func NewVoteFromDict(data map[string]interface{}) Vote {
	return Vote{
		VoteId:         core.CastString(data["voteId"]),
		SeasonName:     core.CastString(data["seasonName"]),
		SessionName:    core.CastString(data["sessionName"]),
		WrittenBallots: CastWrittenBallots(core.CastArray(data["writtenBallots"])),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
		Revision:       core.CastInt64(data["revision"]),
	}
}

func (p Vote) ToDict() map[string]interface{} {

	var voteId *string
	if p.VoteId != nil {
		voteId = p.VoteId
	}
	var seasonName *string
	if p.SeasonName != nil {
		seasonName = p.SeasonName
	}
	var sessionName *string
	if p.SessionName != nil {
		sessionName = p.SessionName
	}
	var writtenBallots []interface{}
	if p.WrittenBallots != nil {
		writtenBallots = CastWrittenBallotsFromDict(
			p.WrittenBallots,
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
		"voteId":         voteId,
		"seasonName":     seasonName,
		"sessionName":    sessionName,
		"writtenBallots": writtenBallots,
		"createdAt":      createdAt,
		"updatedAt":      updatedAt,
		"revision":       revision,
	}
}

func (p Vote) Pointer() *Vote {
	return &p
}

func CastVotes(data []interface{}) []Vote {
	v := make([]Vote, 0)
	for _, d := range data {
		v = append(v, NewVoteFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVotesFromDict(data []Vote) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
