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

package distributor

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId  *string     `json:"namespaceId"`
	Name         *string     `json:"name"`
	Description  *string     `json:"description"`
	AssumeUserId *string     `json:"assumeUserId"`
	LogSetting   *LogSetting `json:"logSetting"`
	CreatedAt    *int64      `json:"createdAt"`
	UpdatedAt    *int64      `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:  core.CastString(data["namespaceId"]),
		Name:         core.CastString(data["name"]),
		Description:  core.CastString(data["description"]),
		AssumeUserId: core.CastString(data["assumeUserId"]),
		LogSetting:   NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:    core.CastInt64(data["createdAt"]),
		UpdatedAt:    core.CastInt64(data["updatedAt"]),
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
	var assumeUserId *string
	if p.AssumeUserId != nil {
		assumeUserId = p.AssumeUserId
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
	return map[string]interface{}{
		"namespaceId":  namespaceId,
		"name":         name,
		"description":  description,
		"assumeUserId": assumeUserId,
		"logSetting":   logSetting,
		"createdAt":    createdAt,
		"updatedAt":    updatedAt,
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

type DistributorModelMaster struct {
	DistributorModelId *string  `json:"distributorModelId"`
	Name               *string  `json:"name"`
	Description        *string  `json:"description"`
	Metadata           *string  `json:"metadata"`
	InboxNamespaceId   *string  `json:"inboxNamespaceId"`
	WhiteListTargetIds []string `json:"whiteListTargetIds"`
	CreatedAt          *int64   `json:"createdAt"`
	UpdatedAt          *int64   `json:"updatedAt"`
}

func NewDistributorModelMasterFromJson(data string) DistributorModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributorModelMasterFromDict(dict)
}

func NewDistributorModelMasterFromDict(data map[string]interface{}) DistributorModelMaster {
	return DistributorModelMaster{
		DistributorModelId: core.CastString(data["distributorModelId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		InboxNamespaceId:   core.CastString(data["inboxNamespaceId"]),
		WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
	}
}

func (p DistributorModelMaster) ToDict() map[string]interface{} {

	var distributorModelId *string
	if p.DistributorModelId != nil {
		distributorModelId = p.DistributorModelId
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
	var inboxNamespaceId *string
	if p.InboxNamespaceId != nil {
		inboxNamespaceId = p.InboxNamespaceId
	}
	var whiteListTargetIds []interface{}
	if p.WhiteListTargetIds != nil {
		whiteListTargetIds = core.CastStringsFromDict(
			p.WhiteListTargetIds,
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
	return map[string]interface{}{
		"distributorModelId": distributorModelId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"inboxNamespaceId":   inboxNamespaceId,
		"whiteListTargetIds": whiteListTargetIds,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
	}
}

func (p DistributorModelMaster) Pointer() *DistributorModelMaster {
	return &p
}

func CastDistributorModelMasters(data []interface{}) []DistributorModelMaster {
	v := make([]DistributorModelMaster, 0)
	for _, d := range data {
		v = append(v, NewDistributorModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDistributorModelMastersFromDict(data []DistributorModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DistributorModel struct {
	DistributorModelId *string  `json:"distributorModelId"`
	Name               *string  `json:"name"`
	Metadata           *string  `json:"metadata"`
	InboxNamespaceId   *string  `json:"inboxNamespaceId"`
	WhiteListTargetIds []string `json:"whiteListTargetIds"`
}

func NewDistributorModelFromJson(data string) DistributorModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributorModelFromDict(dict)
}

func NewDistributorModelFromDict(data map[string]interface{}) DistributorModel {
	return DistributorModel{
		DistributorModelId: core.CastString(data["distributorModelId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		InboxNamespaceId:   core.CastString(data["inboxNamespaceId"]),
		WhiteListTargetIds: core.CastStrings(core.CastArray(data["whiteListTargetIds"])),
	}
}

func (p DistributorModel) ToDict() map[string]interface{} {

	var distributorModelId *string
	if p.DistributorModelId != nil {
		distributorModelId = p.DistributorModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var inboxNamespaceId *string
	if p.InboxNamespaceId != nil {
		inboxNamespaceId = p.InboxNamespaceId
	}
	var whiteListTargetIds []interface{}
	if p.WhiteListTargetIds != nil {
		whiteListTargetIds = core.CastStringsFromDict(
			p.WhiteListTargetIds,
		)
	}
	return map[string]interface{}{
		"distributorModelId": distributorModelId,
		"name":               name,
		"metadata":           metadata,
		"inboxNamespaceId":   inboxNamespaceId,
		"whiteListTargetIds": whiteListTargetIds,
	}
}

func (p DistributorModel) Pointer() *DistributorModel {
	return &p
}

func CastDistributorModels(data []interface{}) []DistributorModel {
	v := make([]DistributorModel, 0)
	for _, d := range data {
		v = append(v, NewDistributorModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDistributorModelsFromDict(data []DistributorModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentDistributorMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentDistributorMasterFromJson(data string) CurrentDistributorMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentDistributorMasterFromDict(dict)
}

func NewCurrentDistributorMasterFromDict(data map[string]interface{}) CurrentDistributorMaster {
	return CurrentDistributorMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentDistributorMaster) ToDict() map[string]interface{} {

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

func (p CurrentDistributorMaster) Pointer() *CurrentDistributorMaster {
	return &p
}

func CastCurrentDistributorMasters(data []interface{}) []CurrentDistributorMaster {
	v := make([]CurrentDistributorMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentDistributorMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentDistributorMastersFromDict(data []CurrentDistributorMaster) []interface{} {
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

type DistributeResource struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewDistributeResourceFromJson(data string) DistributeResource {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDistributeResourceFromDict(dict)
}

func NewDistributeResourceFromDict(data map[string]interface{}) DistributeResource {
	return DistributeResource{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p DistributeResource) ToDict() map[string]interface{} {

	var action *string
	if p.Action != nil {
		action = p.Action
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	return map[string]interface{}{
		"action":  action,
		"request": request,
	}
}

func (p DistributeResource) Pointer() *DistributeResource {
	return &p
}

func CastDistributeResources(data []interface{}) []DistributeResource {
	v := make([]DistributeResource, 0)
	for _, d := range data {
		v = append(v, NewDistributeResourceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDistributeResourcesFromDict(data []DistributeResource) []interface{} {
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
