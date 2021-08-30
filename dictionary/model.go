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

package dictionary

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId          *string        `json:"namespaceId"`
	Name                 *string        `json:"name"`
	Description          *string        `json:"description"`
	EntryScript          *ScriptSetting `json:"entryScript"`
	DuplicateEntryScript *ScriptSetting `json:"duplicateEntryScript"`
	LogSetting           *LogSetting    `json:"logSetting"`
	CreatedAt            *int64         `json:"createdAt"`
	UpdatedAt            *int64         `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:          core.CastString(data["namespaceId"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		EntryScript:          NewScriptSettingFromDict(core.CastMap(data["entryScript"])).Pointer(),
		DuplicateEntryScript: NewScriptSettingFromDict(core.CastMap(data["duplicateEntryScript"])).Pointer(),
		LogSetting:           NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":          p.NamespaceId,
		"name":                 p.Name,
		"description":          p.Description,
		"entryScript":          p.EntryScript.ToDict(),
		"duplicateEntryScript": p.DuplicateEntryScript.ToDict(),
		"logSetting":           p.LogSetting.ToDict(),
		"createdAt":            p.CreatedAt,
		"updatedAt":            p.UpdatedAt,
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

type EntryModel struct {
	EntryModelId *string `json:"entryModelId"`
	Name         *string `json:"name"`
	Metadata     *string `json:"metadata"`
}

func NewEntryModelFromDict(data map[string]interface{}) EntryModel {
	return EntryModel{
		EntryModelId: core.CastString(data["entryModelId"]),
		Name:         core.CastString(data["name"]),
		Metadata:     core.CastString(data["metadata"]),
	}
}

func (p EntryModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"entryModelId": p.EntryModelId,
		"name":         p.Name,
		"metadata":     p.Metadata,
	}
}

func (p EntryModel) Pointer() *EntryModel {
	return &p
}

func CastEntryModels(data []interface{}) []EntryModel {
	v := make([]EntryModel, 0)
	for _, d := range data {
		v = append(v, NewEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEntryModelsFromDict(data []EntryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type EntryModelMaster struct {
	EntryModelId *string `json:"entryModelId"`
	Name         *string `json:"name"`
	Description  *string `json:"description"`
	Metadata     *string `json:"metadata"`
	CreatedAt    *int64  `json:"createdAt"`
	UpdatedAt    *int64  `json:"updatedAt"`
}

func NewEntryModelMasterFromDict(data map[string]interface{}) EntryModelMaster {
	return EntryModelMaster{
		EntryModelId: core.CastString(data["entryModelId"]),
		Name:         core.CastString(data["name"]),
		Description:  core.CastString(data["description"]),
		Metadata:     core.CastString(data["metadata"]),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		UpdatedAt:    core.CastInt64(data["updatedAt"]),
	}
}

func (p EntryModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"entryModelId": p.EntryModelId,
		"name":         p.Name,
		"description":  p.Description,
		"metadata":     p.Metadata,
		"createdAt":    p.CreatedAt,
		"updatedAt":    p.UpdatedAt,
	}
}

func (p EntryModelMaster) Pointer() *EntryModelMaster {
	return &p
}

func CastEntryModelMasters(data []interface{}) []EntryModelMaster {
	v := make([]EntryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewEntryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEntryModelMastersFromDict(data []EntryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Entry struct {
	EntryId    *string `json:"entryId"`
	UserId     *string `json:"userId"`
	Name       *string `json:"name"`
	AcquiredAt *int64  `json:"acquiredAt"`
}

func NewEntryFromDict(data map[string]interface{}) Entry {
	return Entry{
		EntryId:    core.CastString(data["entryId"]),
		UserId:     core.CastString(data["userId"]),
		Name:       core.CastString(data["name"]),
		AcquiredAt: core.CastInt64(data["acquiredAt"]),
	}
}

func (p Entry) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"entryId":    p.EntryId,
		"userId":     p.UserId,
		"name":       p.Name,
		"acquiredAt": p.AcquiredAt,
	}
}

func (p Entry) Pointer() *Entry {
	return &p
}

func CastEntries(data []interface{}) []Entry {
	v := make([]Entry, 0)
	for _, d := range data {
		v = append(v, NewEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEntriesFromDict(data []Entry) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentEntryMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentEntryMasterFromDict(data map[string]interface{}) CurrentEntryMaster {
	return CurrentEntryMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentEntryMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentEntryMaster) Pointer() *CurrentEntryMaster {
	return &p
}

func CastCurrentEntryMasters(data []interface{}) []CurrentEntryMaster {
	v := make([]CurrentEntryMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentEntryMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentEntryMastersFromDict(data []CurrentEntryMaster) []interface{} {
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

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key:   core.CastString(data["key"]),
		Value: core.CastString(data["value"]),
	}
}

func (p Config) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"key":   p.Key,
		"value": p.Value,
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

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
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
	return map[string]interface{}{
		"triggerScriptId":             p.TriggerScriptId,
		"doneTriggerTargetType":       p.DoneTriggerTargetType,
		"doneTriggerScriptId":         p.DoneTriggerScriptId,
		"doneTriggerQueueNamespaceId": p.DoneTriggerQueueNamespaceId,
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
