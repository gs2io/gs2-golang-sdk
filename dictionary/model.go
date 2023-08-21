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

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	EntryScript *ScriptSetting `json:"entryScript"`
	DuplicateEntryScript *ScriptSetting `json:"duplicateEntryScript"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
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
        EntryScript: NewScriptSettingFromDict(core.CastMap(data["entryScript"])).Pointer(),
        DuplicateEntryScript: NewScriptSettingFromDict(core.CastMap(data["duplicateEntryScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
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
    var entryScript map[string]interface{}
    if p.EntryScript != nil {
        entryScript = p.EntryScript.ToDict()
    }
    var duplicateEntryScript map[string]interface{}
    if p.DuplicateEntryScript != nil {
        duplicateEntryScript = p.DuplicateEntryScript.ToDict()
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
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "entryScript": entryScript,
        "duplicateEntryScript": duplicateEntryScript,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
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
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
}

func NewEntryModelFromJson(data string) EntryModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEntryModelFromDict(dict)
}

func NewEntryModelFromDict(data map[string]interface{}) EntryModel {
    return EntryModel {
        EntryModelId: core.CastString(data["entryModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
    }
}

func (p EntryModel) ToDict() map[string]interface{} {
    
    var entryModelId *string
    if p.EntryModelId != nil {
        entryModelId = p.EntryModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    return map[string]interface{} {
        "entryModelId": entryModelId,
        "name": name,
        "metadata": metadata,
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
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewEntryModelMasterFromJson(data string) EntryModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEntryModelMasterFromDict(dict)
}

func NewEntryModelMasterFromDict(data map[string]interface{}) EntryModelMaster {
    return EntryModelMaster {
        EntryModelId: core.CastString(data["entryModelId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p EntryModelMaster) ToDict() map[string]interface{} {
    
    var entryModelId *string
    if p.EntryModelId != nil {
        entryModelId = p.EntryModelId
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
    return map[string]interface{} {
        "entryModelId": entryModelId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
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
	EntryId *string `json:"entryId"`
	UserId *string `json:"userId"`
	Name *string `json:"name"`
	AcquiredAt *int64 `json:"acquiredAt"`
	Revision *int64 `json:"revision"`
}

func NewEntryFromJson(data string) Entry {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewEntryFromDict(dict)
}

func NewEntryFromDict(data map[string]interface{}) Entry {
    return Entry {
        EntryId: core.CastString(data["entryId"]),
        UserId: core.CastString(data["userId"]),
        Name: core.CastString(data["name"]),
        AcquiredAt: core.CastInt64(data["acquiredAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Entry) ToDict() map[string]interface{} {
    
    var entryId *string
    if p.EntryId != nil {
        entryId = p.EntryId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var acquiredAt *int64
    if p.AcquiredAt != nil {
        acquiredAt = p.AcquiredAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "entryId": entryId,
        "userId": userId,
        "name": name,
        "acquiredAt": acquiredAt,
        "revision": revision,
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
	Settings *string `json:"settings"`
}

func NewCurrentEntryMasterFromJson(data string) CurrentEntryMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentEntryMasterFromDict(dict)
}

func NewCurrentEntryMasterFromDict(data map[string]interface{}) CurrentEntryMaster {
    return CurrentEntryMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentEntryMaster) ToDict() map[string]interface{} {
    
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
	Key *string `json:"key"`
	Value *string `json:"value"`
}

func NewConfigFromJson(data string) Config {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewConfigFromDict(dict)
}

func NewConfigFromDict(data map[string]interface{}) Config {
    return Config {
        Key: core.CastString(data["key"]),
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
    return map[string]interface{} {
        "key": key,
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

type ScriptSetting struct {
	TriggerScriptId *string `json:"triggerScriptId"`
	DoneTriggerTargetType *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScriptSettingFromDict(dict)
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
    return ScriptSetting {
        TriggerScriptId: core.CastString(data["triggerScriptId"]),
        DoneTriggerTargetType: core.CastString(data["doneTriggerTargetType"]),
        DoneTriggerScriptId: core.CastString(data["doneTriggerScriptId"]),
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
    return map[string]interface{} {
        "triggerScriptId": triggerScriptId,
        "doneTriggerTargetType": doneTriggerTargetType,
        "doneTriggerScriptId": doneTriggerScriptId,
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