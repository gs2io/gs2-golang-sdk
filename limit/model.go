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

package limit

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
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

type Counter struct {
	CounterId *string `json:"counterId"`
	LimitName *string `json:"limitName"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	Count *int32 `json:"count"`
	NextResetAt *int64 `json:"nextResetAt"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewCounterFromJson(data string) Counter {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCounterFromDict(dict)
}

func NewCounterFromDict(data map[string]interface{}) Counter {
    return Counter {
        CounterId: core.CastString(data["counterId"]),
        LimitName: core.CastString(data["limitName"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        Count: core.CastInt32(data["count"]),
        NextResetAt: core.CastInt64(data["nextResetAt"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Counter) ToDict() map[string]interface{} {
    
    var counterId *string
    if p.CounterId != nil {
        counterId = p.CounterId
    }
    var limitName *string
    if p.LimitName != nil {
        limitName = p.LimitName
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var count *int32
    if p.Count != nil {
        count = p.Count
    }
    var nextResetAt *int64
    if p.NextResetAt != nil {
        nextResetAt = p.NextResetAt
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
        "counterId": counterId,
        "limitName": limitName,
        "name": name,
        "userId": userId,
        "count": count,
        "nextResetAt": nextResetAt,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Counter) Pointer() *Counter {
    return &p
}

func CastCounters(data []interface{}) []Counter {
	v := make([]Counter, 0)
	for _, d := range data {
		v = append(v, NewCounterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCountersFromDict(data []Counter) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LimitModelMaster struct {
	LimitModelId *string `json:"limitModelId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	ResetType *string `json:"resetType"`
	ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
	ResetDayOfWeek *string `json:"resetDayOfWeek"`
	ResetHour *int32 `json:"resetHour"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewLimitModelMasterFromJson(data string) LimitModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLimitModelMasterFromDict(dict)
}

func NewLimitModelMasterFromDict(data map[string]interface{}) LimitModelMaster {
    return LimitModelMaster {
        LimitModelId: core.CastString(data["limitModelId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p LimitModelMaster) ToDict() map[string]interface{} {
    
    var limitModelId *string
    if p.LimitModelId != nil {
        limitModelId = p.LimitModelId
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
    var resetType *string
    if p.ResetType != nil {
        resetType = p.ResetType
    }
    var resetDayOfMonth *int32
    if p.ResetDayOfMonth != nil {
        resetDayOfMonth = p.ResetDayOfMonth
    }
    var resetDayOfWeek *string
    if p.ResetDayOfWeek != nil {
        resetDayOfWeek = p.ResetDayOfWeek
    }
    var resetHour *int32
    if p.ResetHour != nil {
        resetHour = p.ResetHour
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
        "limitModelId": limitModelId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "resetType": resetType,
        "resetDayOfMonth": resetDayOfMonth,
        "resetDayOfWeek": resetDayOfWeek,
        "resetHour": resetHour,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p LimitModelMaster) Pointer() *LimitModelMaster {
    return &p
}

func CastLimitModelMasters(data []interface{}) []LimitModelMaster {
	v := make([]LimitModelMaster, 0)
	for _, d := range data {
		v = append(v, NewLimitModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitModelMastersFromDict(data []LimitModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentLimitMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentLimitMasterFromJson(data string) CurrentLimitMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentLimitMasterFromDict(dict)
}

func NewCurrentLimitMasterFromDict(data map[string]interface{}) CurrentLimitMaster {
    return CurrentLimitMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentLimitMaster) ToDict() map[string]interface{} {
    
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

func (p CurrentLimitMaster) Pointer() *CurrentLimitMaster {
    return &p
}

func CastCurrentLimitMasters(data []interface{}) []CurrentLimitMaster {
	v := make([]CurrentLimitMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentLimitMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentLimitMastersFromDict(data []CurrentLimitMaster) []interface{} {
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

type LimitModel struct {
	LimitModelId *string `json:"limitModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	ResetType *string `json:"resetType"`
	ResetDayOfMonth *int32 `json:"resetDayOfMonth"`
	ResetDayOfWeek *string `json:"resetDayOfWeek"`
	ResetHour *int32 `json:"resetHour"`
}

func NewLimitModelFromJson(data string) LimitModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLimitModelFromDict(dict)
}

func NewLimitModelFromDict(data map[string]interface{}) LimitModel {
    return LimitModel {
        LimitModelId: core.CastString(data["limitModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        ResetType: core.CastString(data["resetType"]),
        ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
        ResetDayOfWeek: core.CastString(data["resetDayOfWeek"]),
        ResetHour: core.CastInt32(data["resetHour"]),
    }
}

func (p LimitModel) ToDict() map[string]interface{} {
    
    var limitModelId *string
    if p.LimitModelId != nil {
        limitModelId = p.LimitModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var resetType *string
    if p.ResetType != nil {
        resetType = p.ResetType
    }
    var resetDayOfMonth *int32
    if p.ResetDayOfMonth != nil {
        resetDayOfMonth = p.ResetDayOfMonth
    }
    var resetDayOfWeek *string
    if p.ResetDayOfWeek != nil {
        resetDayOfWeek = p.ResetDayOfWeek
    }
    var resetHour *int32
    if p.ResetHour != nil {
        resetHour = p.ResetHour
    }
    return map[string]interface{} {
        "limitModelId": limitModelId,
        "name": name,
        "metadata": metadata,
        "resetType": resetType,
        "resetDayOfMonth": resetDayOfMonth,
        "resetDayOfWeek": resetDayOfWeek,
        "resetHour": resetHour,
    }
}

func (p LimitModel) Pointer() *LimitModel {
    return &p
}

func CastLimitModels(data []interface{}) []LimitModel {
	v := make([]LimitModel, 0)
	for _, d := range data {
		v = append(v, NewLimitModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLimitModelsFromDict(data []LimitModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}