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

package news

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Version *string `json:"version"`
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
        Version: core.CastString(data["version"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Namespace) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceId": p.NamespaceId,
        "name": p.Name,
        "description": p.Description,
        "version": p.Version,
        "logSetting": p.LogSetting.ToDict(),
        "createdAt": p.CreatedAt,
        "updatedAt": p.UpdatedAt,
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

type News struct {
	Section *string `json:"section"`
	Content *string `json:"content"`
	Title *string `json:"title"`
	ScheduleEventId *string `json:"scheduleEventId"`
	Timestamp *int64 `json:"timestamp"`
	FrontMatter *string `json:"frontMatter"`
}

func NewNewsFromJson(data string) News {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNewsFromDict(dict)
}

func NewNewsFromDict(data map[string]interface{}) News {
    return News {
        Section: core.CastString(data["section"]),
        Content: core.CastString(data["content"]),
        Title: core.CastString(data["title"]),
        ScheduleEventId: core.CastString(data["scheduleEventId"]),
        Timestamp: core.CastInt64(data["timestamp"]),
        FrontMatter: core.CastString(data["frontMatter"]),
    }
}

func (p News) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "section": p.Section,
        "content": p.Content,
        "title": p.Title,
        "scheduleEventId": p.ScheduleEventId,
        "timestamp": p.Timestamp,
        "frontMatter": p.FrontMatter,
    }
}

func (p News) Pointer() *News {
    return &p
}

func CastNewses(data []interface{}) []News {
	v := make([]News, 0)
	for _, d := range data {
		v = append(v, NewNewsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNewsesFromDict(data []News) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type SetCookieRequestEntry struct {
	Key *string `json:"key"`
	Value *string `json:"value"`
}

func NewSetCookieRequestEntryFromJson(data string) SetCookieRequestEntry {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetCookieRequestEntryFromDict(dict)
}

func NewSetCookieRequestEntryFromDict(data map[string]interface{}) SetCookieRequestEntry {
    return SetCookieRequestEntry {
        Key: core.CastString(data["key"]),
        Value: core.CastString(data["value"]),
    }
}

func (p SetCookieRequestEntry) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "key": p.Key,
        "value": p.Value,
    }
}

func (p SetCookieRequestEntry) Pointer() *SetCookieRequestEntry {
    return &p
}

func CastSetCookieRequestEntries(data []interface{}) []SetCookieRequestEntry {
	v := make([]SetCookieRequestEntry, 0)
	for _, d := range data {
		v = append(v, NewSetCookieRequestEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSetCookieRequestEntriesFromDict(data []SetCookieRequestEntry) []interface{} {
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
    return map[string]interface{} {
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
    return map[string]interface{} {
        "apiKeyId": p.ApiKeyId,
        "repositoryName": p.RepositoryName,
        "sourcePath": p.SourcePath,
        "referenceType": p.ReferenceType,
        "commitHash": p.CommitHash,
        "branchName": p.BranchName,
        "tagName": p.TagName,
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