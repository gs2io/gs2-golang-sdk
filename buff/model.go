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

package buff

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

type BuffTargetModel struct {
	TargetModelName *string         `json:"targetModelName"`
	TargetFieldName *string         `json:"targetFieldName"`
	ConditionGrns   []BuffTargetGrn `json:"conditionGrns"`
	Rate            *float32        `json:"rate"`
}

func NewBuffTargetModelFromJson(data string) BuffTargetModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuffTargetModelFromDict(dict)
}

func NewBuffTargetModelFromDict(data map[string]interface{}) BuffTargetModel {
	return BuffTargetModel{
		TargetModelName: core.CastString(data["targetModelName"]),
		TargetFieldName: core.CastString(data["targetFieldName"]),
		ConditionGrns:   CastBuffTargetGrns(core.CastArray(data["conditionGrns"])),
		Rate:            core.CastFloat32(data["rate"]),
	}
}

func (p BuffTargetModel) ToDict() map[string]interface{} {

	var targetModelName *string
	if p.TargetModelName != nil {
		targetModelName = p.TargetModelName
	}
	var targetFieldName *string
	if p.TargetFieldName != nil {
		targetFieldName = p.TargetFieldName
	}
	var conditionGrns []interface{}
	if p.ConditionGrns != nil {
		conditionGrns = CastBuffTargetGrnsFromDict(
			p.ConditionGrns,
		)
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	return map[string]interface{}{
		"targetModelName": targetModelName,
		"targetFieldName": targetFieldName,
		"conditionGrns":   conditionGrns,
		"rate":            rate,
	}
}

func (p BuffTargetModel) Pointer() *BuffTargetModel {
	return &p
}

func CastBuffTargetModels(data []interface{}) []BuffTargetModel {
	v := make([]BuffTargetModel, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetModelsFromDict(data []BuffTargetModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffTargetAction struct {
	TargetActionName *string         `json:"targetActionName"`
	TargetFieldName  *string         `json:"targetFieldName"`
	ConditionGrns    []BuffTargetGrn `json:"conditionGrns"`
	Rate             *float32        `json:"rate"`
}

func NewBuffTargetActionFromJson(data string) BuffTargetAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuffTargetActionFromDict(dict)
}

func NewBuffTargetActionFromDict(data map[string]interface{}) BuffTargetAction {
	return BuffTargetAction{
		TargetActionName: core.CastString(data["targetActionName"]),
		TargetFieldName:  core.CastString(data["targetFieldName"]),
		ConditionGrns:    CastBuffTargetGrns(core.CastArray(data["conditionGrns"])),
		Rate:             core.CastFloat32(data["rate"]),
	}
}

func (p BuffTargetAction) ToDict() map[string]interface{} {

	var targetActionName *string
	if p.TargetActionName != nil {
		targetActionName = p.TargetActionName
	}
	var targetFieldName *string
	if p.TargetFieldName != nil {
		targetFieldName = p.TargetFieldName
	}
	var conditionGrns []interface{}
	if p.ConditionGrns != nil {
		conditionGrns = CastBuffTargetGrnsFromDict(
			p.ConditionGrns,
		)
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	return map[string]interface{}{
		"targetActionName": targetActionName,
		"targetFieldName":  targetFieldName,
		"conditionGrns":    conditionGrns,
		"rate":             rate,
	}
}

func (p BuffTargetAction) Pointer() *BuffTargetAction {
	return &p
}

func CastBuffTargetActions(data []interface{}) []BuffTargetAction {
	v := make([]BuffTargetAction, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetActionsFromDict(data []BuffTargetAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffTargetGrn struct {
	TargetModelName *string `json:"targetModelName"`
	TargetGrn       *string `json:"targetGrn"`
}

func NewBuffTargetGrnFromJson(data string) BuffTargetGrn {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuffTargetGrnFromDict(dict)
}

func NewBuffTargetGrnFromDict(data map[string]interface{}) BuffTargetGrn {
	return BuffTargetGrn{
		TargetModelName: core.CastString(data["targetModelName"]),
		TargetGrn:       core.CastString(data["targetGrn"]),
	}
}

func (p BuffTargetGrn) ToDict() map[string]interface{} {

	var targetModelName *string
	if p.TargetModelName != nil {
		targetModelName = p.TargetModelName
	}
	var targetGrn *string
	if p.TargetGrn != nil {
		targetGrn = p.TargetGrn
	}
	return map[string]interface{}{
		"targetModelName": targetModelName,
		"targetGrn":       targetGrn,
	}
}

func (p BuffTargetGrn) Pointer() *BuffTargetGrn {
	return &p
}

func CastBuffTargetGrns(data []interface{}) []BuffTargetGrn {
	v := make([]BuffTargetGrn, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetGrnFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetGrnsFromDict(data []BuffTargetGrn) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffEntryModel struct {
	BuffEntryModelId           *string           `json:"buffEntryModelId"`
	Name                       *string           `json:"name"`
	Metadata                   *string           `json:"metadata"`
	TargetType                 *string           `json:"targetType"`
	TargetModel                *BuffTargetModel  `json:"targetModel"`
	TargetAction               *BuffTargetAction `json:"targetAction"`
	Expression                 *string           `json:"expression"`
	Priority                   *int32            `json:"priority"`
	ApplyPeriodScheduleEventId *string           `json:"applyPeriodScheduleEventId"`
}

func NewBuffEntryModelFromJson(data string) BuffEntryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuffEntryModelFromDict(dict)
}

func NewBuffEntryModelFromDict(data map[string]interface{}) BuffEntryModel {
	return BuffEntryModel{
		BuffEntryModelId:           core.CastString(data["buffEntryModelId"]),
		Name:                       core.CastString(data["name"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetType:                 core.CastString(data["targetType"]),
		TargetModel:                NewBuffTargetModelFromDict(core.CastMap(data["targetModel"])).Pointer(),
		TargetAction:               NewBuffTargetActionFromDict(core.CastMap(data["targetAction"])).Pointer(),
		Expression:                 core.CastString(data["expression"]),
		Priority:                   core.CastInt32(data["priority"]),
		ApplyPeriodScheduleEventId: core.CastString(data["applyPeriodScheduleEventId"]),
	}
}

func (p BuffEntryModel) ToDict() map[string]interface{} {

	var buffEntryModelId *string
	if p.BuffEntryModelId != nil {
		buffEntryModelId = p.BuffEntryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var targetType *string
	if p.TargetType != nil {
		targetType = p.TargetType
	}
	var targetModel map[string]interface{}
	if p.TargetModel != nil {
		targetModel = p.TargetModel.ToDict()
	}
	var targetAction map[string]interface{}
	if p.TargetAction != nil {
		targetAction = p.TargetAction.ToDict()
	}
	var expression *string
	if p.Expression != nil {
		expression = p.Expression
	}
	var priority *int32
	if p.Priority != nil {
		priority = p.Priority
	}
	var applyPeriodScheduleEventId *string
	if p.ApplyPeriodScheduleEventId != nil {
		applyPeriodScheduleEventId = p.ApplyPeriodScheduleEventId
	}
	return map[string]interface{}{
		"buffEntryModelId":           buffEntryModelId,
		"name":                       name,
		"metadata":                   metadata,
		"targetType":                 targetType,
		"targetModel":                targetModel,
		"targetAction":               targetAction,
		"expression":                 expression,
		"priority":                   priority,
		"applyPeriodScheduleEventId": applyPeriodScheduleEventId,
	}
}

func (p BuffEntryModel) Pointer() *BuffEntryModel {
	return &p
}

func CastBuffEntryModels(data []interface{}) []BuffEntryModel {
	v := make([]BuffEntryModel, 0)
	for _, d := range data {
		v = append(v, NewBuffEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffEntryModelsFromDict(data []BuffEntryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffEntryModelMaster struct {
	BuffEntryModelId           *string           `json:"buffEntryModelId"`
	Name                       *string           `json:"name"`
	Description                *string           `json:"description"`
	Metadata                   *string           `json:"metadata"`
	TargetType                 *string           `json:"targetType"`
	TargetModel                *BuffTargetModel  `json:"targetModel"`
	TargetAction               *BuffTargetAction `json:"targetAction"`
	Expression                 *string           `json:"expression"`
	Priority                   *int32            `json:"priority"`
	ApplyPeriodScheduleEventId *string           `json:"applyPeriodScheduleEventId"`
	CreatedAt                  *int64            `json:"createdAt"`
	UpdatedAt                  *int64            `json:"updatedAt"`
	Revision                   *int64            `json:"revision"`
}

func NewBuffEntryModelMasterFromJson(data string) BuffEntryModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuffEntryModelMasterFromDict(dict)
}

func NewBuffEntryModelMasterFromDict(data map[string]interface{}) BuffEntryModelMaster {
	return BuffEntryModelMaster{
		BuffEntryModelId:           core.CastString(data["buffEntryModelId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetType:                 core.CastString(data["targetType"]),
		TargetModel:                NewBuffTargetModelFromDict(core.CastMap(data["targetModel"])).Pointer(),
		TargetAction:               NewBuffTargetActionFromDict(core.CastMap(data["targetAction"])).Pointer(),
		Expression:                 core.CastString(data["expression"]),
		Priority:                   core.CastInt32(data["priority"]),
		ApplyPeriodScheduleEventId: core.CastString(data["applyPeriodScheduleEventId"]),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
		Revision:                   core.CastInt64(data["revision"]),
	}
}

func (p BuffEntryModelMaster) ToDict() map[string]interface{} {

	var buffEntryModelId *string
	if p.BuffEntryModelId != nil {
		buffEntryModelId = p.BuffEntryModelId
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
	var targetType *string
	if p.TargetType != nil {
		targetType = p.TargetType
	}
	var targetModel map[string]interface{}
	if p.TargetModel != nil {
		targetModel = p.TargetModel.ToDict()
	}
	var targetAction map[string]interface{}
	if p.TargetAction != nil {
		targetAction = p.TargetAction.ToDict()
	}
	var expression *string
	if p.Expression != nil {
		expression = p.Expression
	}
	var priority *int32
	if p.Priority != nil {
		priority = p.Priority
	}
	var applyPeriodScheduleEventId *string
	if p.ApplyPeriodScheduleEventId != nil {
		applyPeriodScheduleEventId = p.ApplyPeriodScheduleEventId
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
		"buffEntryModelId":           buffEntryModelId,
		"name":                       name,
		"description":                description,
		"metadata":                   metadata,
		"targetType":                 targetType,
		"targetModel":                targetModel,
		"targetAction":               targetAction,
		"expression":                 expression,
		"priority":                   priority,
		"applyPeriodScheduleEventId": applyPeriodScheduleEventId,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
		"revision":                   revision,
	}
}

func (p BuffEntryModelMaster) Pointer() *BuffEntryModelMaster {
	return &p
}

func CastBuffEntryModelMasters(data []interface{}) []BuffEntryModelMaster {
	v := make([]BuffEntryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBuffEntryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffEntryModelMastersFromDict(data []BuffEntryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentBuffMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentBuffMasterFromJson(data string) CurrentBuffMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentBuffMasterFromDict(dict)
}

func NewCurrentBuffMasterFromDict(data map[string]interface{}) CurrentBuffMaster {
	return CurrentBuffMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentBuffMaster) ToDict() map[string]interface{} {

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

func (p CurrentBuffMaster) Pointer() *CurrentBuffMaster {
	return &p
}

func CastCurrentBuffMasters(data []interface{}) []CurrentBuffMaster {
	v := make([]CurrentBuffMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentBuffMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentBuffMastersFromDict(data []CurrentBuffMaster) []interface{} {
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
