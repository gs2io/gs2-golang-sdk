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

type Counter struct {
	CounterId *string `json:"counterId"`
	LimitName *string `json:"limitName"`
	Name      *string `json:"name"`
	UserId    *string `json:"userId"`
	Count     *int32  `json:"count"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
}

func NewCounterFromJson(data string) Counter {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCounterFromDict(dict)
}

func NewCounterFromDict(data map[string]interface{}) Counter {
	return Counter{
		CounterId: core.CastString(data["counterId"]),
		LimitName: core.CastString(data["limitName"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		Count:     core.CastInt32(data["count"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
	}
}

func (p Counter) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"counterId": p.CounterId,
		"limitName": p.LimitName,
		"name":      p.Name,
		"userId":    p.UserId,
		"count":     p.Count,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
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
	LimitModelId    *string `json:"limitModelId"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	ResetType       *string `json:"resetType"`
	ResetDayOfMonth *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek  *string `json:"resetDayOfWeek"`
	ResetHour       *int32  `json:"resetHour"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
}

func NewLimitModelMasterFromJson(data string) LimitModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitModelMasterFromDict(dict)
}

func NewLimitModelMasterFromDict(data map[string]interface{}) LimitModelMaster {
	return LimitModelMaster{
		LimitModelId:    core.CastString(data["limitModelId"]),
		Name:            core.CastString(data["name"]),
		Description:     core.CastString(data["description"]),
		Metadata:        core.CastString(data["metadata"]),
		ResetType:       core.CastString(data["resetType"]),
		ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:  core.CastString(data["resetDayOfWeek"]),
		ResetHour:       core.CastInt32(data["resetHour"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
	}
}

func (p LimitModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"limitModelId":    p.LimitModelId,
		"name":            p.Name,
		"description":     p.Description,
		"metadata":        p.Metadata,
		"resetType":       p.ResetType,
		"resetDayOfMonth": p.ResetDayOfMonth,
		"resetDayOfWeek":  p.ResetDayOfWeek,
		"resetHour":       p.ResetHour,
		"createdAt":       p.CreatedAt,
		"updatedAt":       p.UpdatedAt,
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
	Settings    *string `json:"settings"`
}

func NewCurrentLimitMasterFromJson(data string) CurrentLimitMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentLimitMasterFromDict(dict)
}

func NewCurrentLimitMasterFromDict(data map[string]interface{}) CurrentLimitMaster {
	return CurrentLimitMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentLimitMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
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

type LimitModel struct {
	LimitModelId    *string `json:"limitModelId"`
	Name            *string `json:"name"`
	Metadata        *string `json:"metadata"`
	ResetType       *string `json:"resetType"`
	ResetDayOfMonth *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek  *string `json:"resetDayOfWeek"`
	ResetHour       *int32  `json:"resetHour"`
}

func NewLimitModelFromJson(data string) LimitModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewLimitModelFromDict(dict)
}

func NewLimitModelFromDict(data map[string]interface{}) LimitModel {
	return LimitModel{
		LimitModelId:    core.CastString(data["limitModelId"]),
		Name:            core.CastString(data["name"]),
		Metadata:        core.CastString(data["metadata"]),
		ResetType:       core.CastString(data["resetType"]),
		ResetDayOfMonth: core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:  core.CastString(data["resetDayOfWeek"]),
		ResetHour:       core.CastInt32(data["resetHour"]),
	}
}

func (p LimitModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"limitModelId":    p.LimitModelId,
		"name":            p.Name,
		"metadata":        p.Metadata,
		"resetType":       p.ResetType,
		"resetDayOfMonth": p.ResetDayOfMonth,
		"resetDayOfWeek":  p.ResetDayOfWeek,
		"resetHour":       p.ResetHour,
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
