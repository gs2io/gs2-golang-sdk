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

package key

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
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"name":        p.Name,
		"description": p.Description,
		"logSetting":  p.LogSetting.ToDict(),
		"createdAt":   p.CreatedAt,
		"updatedAt":   p.UpdatedAt,
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

type Key struct {
	KeyId       *string `json:"keyId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func NewKeyFromJson(data string) Key {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewKeyFromDict(dict)
}

func NewKeyFromDict(data map[string]interface{}) Key {
	return Key{
		KeyId:       core.CastString(data["keyId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
	}
}

func (p Key) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"keyId":       p.KeyId,
		"name":        p.Name,
		"description": p.Description,
		"createdAt":   p.CreatedAt,
		"updatedAt":   p.UpdatedAt,
	}
}

func (p Key) Pointer() *Key {
	return &p
}

func CastKeys(data []interface{}) []Key {
	v := make([]Key, 0)
	for _, d := range data {
		v = append(v, NewKeyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastKeysFromDict(data []Key) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GitHubApiKey struct {
	ApiKeyId          *string `json:"apiKeyId"`
	Name              *string `json:"name"`
	Description       *string `json:"description"`
	EncryptionKeyName *string `json:"encryptionKeyName"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
}

func NewGitHubApiKeyFromJson(data string) GitHubApiKey {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGitHubApiKeyFromDict(dict)
}

func NewGitHubApiKeyFromDict(data map[string]interface{}) GitHubApiKey {
	return GitHubApiKey{
		ApiKeyId:          core.CastString(data["apiKeyId"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		EncryptionKeyName: core.CastString(data["encryptionKeyName"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
	}
}

func (p GitHubApiKey) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"apiKeyId":          p.ApiKeyId,
		"name":              p.Name,
		"description":       p.Description,
		"encryptionKeyName": p.EncryptionKeyName,
		"createdAt":         p.CreatedAt,
		"updatedAt":         p.UpdatedAt,
	}
}

func (p GitHubApiKey) Pointer() *GitHubApiKey {
	return &p
}

func CastGitHubApiKeys(data []interface{}) []GitHubApiKey {
	v := make([]GitHubApiKey, 0)
	for _, d := range data {
		v = append(v, NewGitHubApiKeyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubApiKeysFromDict(data []GitHubApiKey) []interface{} {
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
