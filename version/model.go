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

package version

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId                 *string        `json:"namespaceId"`
	Name                        *string        `json:"name"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
	CreatedAt                   *int64         `json:"createdAt"`
	UpdatedAt                   *int64         `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:                 core.CastString(data["namespaceId"]),
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		AssumeUserId:                core.CastString(data["assumeUserId"]),
		AcceptVersionScript:         NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer(),
		CheckVersionTriggerScriptId: core.CastString(data["checkVersionTriggerScriptId"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                   core.CastInt64(data["createdAt"]),
		UpdatedAt:                   core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":                 p.NamespaceId,
		"name":                        p.Name,
		"description":                 p.Description,
		"assumeUserId":                p.AssumeUserId,
		"acceptVersionScript":         p.AcceptVersionScript.ToDict(),
		"checkVersionTriggerScriptId": p.CheckVersionTriggerScriptId,
		"logSetting":                  p.LogSetting.ToDict(),
		"createdAt":                   p.CreatedAt,
		"updatedAt":                   p.UpdatedAt,
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

type VersionModelMaster struct {
	VersionModelId *string  `json:"versionModelId"`
	Name           *string  `json:"name"`
	Description    *string  `json:"description"`
	Metadata       *string  `json:"metadata"`
	WarningVersion *Version `json:"warningVersion"`
	ErrorVersion   *Version `json:"errorVersion"`
	Scope          *string  `json:"scope"`
	CurrentVersion *Version `json:"currentVersion"`
	NeedSignature  *bool    `json:"needSignature"`
	SignatureKeyId *string  `json:"signatureKeyId"`
	CreatedAt      *int64   `json:"createdAt"`
	UpdatedAt      *int64   `json:"updatedAt"`
}

func NewVersionModelMasterFromDict(data map[string]interface{}) VersionModelMaster {
	return VersionModelMaster{
		VersionModelId: core.CastString(data["versionModelId"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		WarningVersion: NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer(),
		ErrorVersion:   NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer(),
		Scope:          core.CastString(data["scope"]),
		CurrentVersion: NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer(),
		NeedSignature:  core.CastBool(data["needSignature"]),
		SignatureKeyId: core.CastString(data["signatureKeyId"]),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
	}
}

func (p VersionModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"versionModelId": p.VersionModelId,
		"name":           p.Name,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"warningVersion": p.WarningVersion.ToDict(),
		"errorVersion":   p.ErrorVersion.ToDict(),
		"scope":          p.Scope,
		"currentVersion": p.CurrentVersion.ToDict(),
		"needSignature":  p.NeedSignature,
		"signatureKeyId": p.SignatureKeyId,
		"createdAt":      p.CreatedAt,
		"updatedAt":      p.UpdatedAt,
	}
}

func (p VersionModelMaster) Pointer() *VersionModelMaster {
	return &p
}

func CastVersionModelMasters(data []interface{}) []VersionModelMaster {
	v := make([]VersionModelMaster, 0)
	for _, d := range data {
		v = append(v, NewVersionModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionModelMastersFromDict(data []VersionModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Version struct {
	Major *int32 `json:"major"`
	Minor *int32 `json:"minor"`
	Micro *int32 `json:"micro"`
}

func NewVersionFromDict(data map[string]interface{}) Version {
	return Version{
		Major: core.CastInt32(data["major"]),
		Minor: core.CastInt32(data["minor"]),
		Micro: core.CastInt32(data["micro"]),
	}
}

func (p Version) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"major": p.Major,
		"minor": p.Minor,
		"micro": p.Micro,
	}
}

func (p Version) Pointer() *Version {
	return &p
}

func CastVersions(data []interface{}) []Version {
	v := make([]Version, 0)
	for _, d := range data {
		v = append(v, NewVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionsFromDict(data []Version) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VersionModel struct {
	VersionModelId *string  `json:"versionModelId"`
	Name           *string  `json:"name"`
	Metadata       *string  `json:"metadata"`
	WarningVersion *Version `json:"warningVersion"`
	ErrorVersion   *Version `json:"errorVersion"`
	Scope          *string  `json:"scope"`
	CurrentVersion *Version `json:"currentVersion"`
	NeedSignature  *bool    `json:"needSignature"`
	SignatureKeyId *string  `json:"signatureKeyId"`
}

func NewVersionModelFromDict(data map[string]interface{}) VersionModel {
	return VersionModel{
		VersionModelId: core.CastString(data["versionModelId"]),
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		WarningVersion: NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer(),
		ErrorVersion:   NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer(),
		Scope:          core.CastString(data["scope"]),
		CurrentVersion: NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer(),
		NeedSignature:  core.CastBool(data["needSignature"]),
		SignatureKeyId: core.CastString(data["signatureKeyId"]),
	}
}

func (p VersionModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"versionModelId": p.VersionModelId,
		"name":           p.Name,
		"metadata":       p.Metadata,
		"warningVersion": p.WarningVersion.ToDict(),
		"errorVersion":   p.ErrorVersion.ToDict(),
		"scope":          p.Scope,
		"currentVersion": p.CurrentVersion.ToDict(),
		"needSignature":  p.NeedSignature,
		"signatureKeyId": p.SignatureKeyId,
	}
}

func (p VersionModel) Pointer() *VersionModel {
	return &p
}

func CastVersionModels(data []interface{}) []VersionModel {
	v := make([]VersionModel, 0)
	for _, d := range data {
		v = append(v, NewVersionModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionModelsFromDict(data []VersionModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcceptVersion struct {
	AcceptVersionId *string  `json:"acceptVersionId"`
	VersionName     *string  `json:"versionName"`
	UserId          *string  `json:"userId"`
	Version         *Version `json:"version"`
	CreatedAt       *int64   `json:"createdAt"`
	UpdatedAt       *int64   `json:"updatedAt"`
}

func NewAcceptVersionFromDict(data map[string]interface{}) AcceptVersion {
	return AcceptVersion{
		AcceptVersionId: core.CastString(data["acceptVersionId"]),
		VersionName:     core.CastString(data["versionName"]),
		UserId:          core.CastString(data["userId"]),
		Version:         NewVersionFromDict(core.CastMap(data["version"])).Pointer(),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
	}
}

func (p AcceptVersion) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"acceptVersionId": p.AcceptVersionId,
		"versionName":     p.VersionName,
		"userId":          p.UserId,
		"version":         p.Version.ToDict(),
		"createdAt":       p.CreatedAt,
		"updatedAt":       p.UpdatedAt,
	}
}

func (p AcceptVersion) Pointer() *AcceptVersion {
	return &p
}

func CastAcceptVersions(data []interface{}) []AcceptVersion {
	v := make([]AcceptVersion, 0)
	for _, d := range data {
		v = append(v, NewAcceptVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcceptVersionsFromDict(data []AcceptVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	VersionModel   *VersionModel `json:"versionModel"`
	CurrentVersion *Version      `json:"currentVersion"`
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		VersionModel:   NewVersionModelFromDict(core.CastMap(data["versionModel"])).Pointer(),
		CurrentVersion: NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer(),
	}
}

func (p Status) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"versionModel":   p.VersionModel.ToDict(),
		"currentVersion": p.CurrentVersion.ToDict(),
	}
}

func (p Status) Pointer() *Status {
	return &p
}

func CastStatuses(data []interface{}) []Status {
	v := make([]Status, 0)
	for _, d := range data {
		v = append(v, NewStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStatusesFromDict(data []Status) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TargetVersion struct {
	VersionName *string  `json:"versionName"`
	Version     *Version `json:"version"`
	Body        *string  `json:"body"`
	Signature   *string  `json:"signature"`
}

func NewTargetVersionFromDict(data map[string]interface{}) TargetVersion {
	return TargetVersion{
		VersionName: core.CastString(data["versionName"]),
		Version:     NewVersionFromDict(core.CastMap(data["version"])).Pointer(),
		Body:        core.CastString(data["body"]),
		Signature:   core.CastString(data["signature"]),
	}
}

func (p TargetVersion) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"versionName": p.VersionName,
		"version":     p.Version.ToDict(),
		"body":        p.Body,
		"signature":   p.Signature,
	}
}

func (p TargetVersion) Pointer() *TargetVersion {
	return &p
}

func CastTargetVersions(data []interface{}) []TargetVersion {
	v := make([]TargetVersion, 0)
	for _, d := range data {
		v = append(v, NewTargetVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTargetVersionsFromDict(data []TargetVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SignTargetVersion struct {
	Region        *string  `json:"region"`
	NamespaceName *string  `json:"namespaceName"`
	VersionName   *string  `json:"versionName"`
	Version       *Version `json:"version"`
}

func NewSignTargetVersionFromDict(data map[string]interface{}) SignTargetVersion {
	return SignTargetVersion{
		Region:        core.CastString(data["region"]),
		NamespaceName: core.CastString(data["namespaceName"]),
		VersionName:   core.CastString(data["versionName"]),
		Version:       NewVersionFromDict(core.CastMap(data["version"])).Pointer(),
	}
}

func (p SignTargetVersion) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"region":        p.Region,
		"namespaceName": p.NamespaceName,
		"versionName":   p.VersionName,
		"version":       p.Version.ToDict(),
	}
}

func (p SignTargetVersion) Pointer() *SignTargetVersion {
	return &p
}

func CastSignTargetVersions(data []interface{}) []SignTargetVersion {
	v := make([]SignTargetVersion, 0)
	for _, d := range data {
		v = append(v, NewSignTargetVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignTargetVersionsFromDict(data []SignTargetVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentVersionMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentVersionMasterFromDict(data map[string]interface{}) CurrentVersionMaster {
	return CurrentVersionMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentVersionMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentVersionMaster) Pointer() *CurrentVersionMaster {
	return &p
}

func CastCurrentVersionMasters(data []interface{}) []CurrentVersionMaster {
	v := make([]CurrentVersionMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentVersionMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentVersionMastersFromDict(data []CurrentVersionMaster) []interface{} {
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
