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

package serialKey

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

type IssueJob struct {
	IssueJobId        *string `json:"issueJobId"`
	Name              *string `json:"name"`
	Metadata          *string `json:"metadata"`
	IssuedCount       *int32  `json:"issuedCount"`
	IssueRequestCount *int32  `json:"issueRequestCount"`
	Status            *string `json:"status"`
	CreatedAt         *int64  `json:"createdAt"`
	Revision          *int64  `json:"revision"`
}

func NewIssueJobFromJson(data string) IssueJob {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIssueJobFromDict(dict)
}

func NewIssueJobFromDict(data map[string]interface{}) IssueJob {
	return IssueJob{
		IssueJobId:        core.CastString(data["issueJobId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		IssuedCount:       core.CastInt32(data["issuedCount"]),
		IssueRequestCount: core.CastInt32(data["issueRequestCount"]),
		Status:            core.CastString(data["status"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p IssueJob) ToDict() map[string]interface{} {

	var issueJobId *string
	if p.IssueJobId != nil {
		issueJobId = p.IssueJobId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var issuedCount *int32
	if p.IssuedCount != nil {
		issuedCount = p.IssuedCount
	}
	var issueRequestCount *int32
	if p.IssueRequestCount != nil {
		issueRequestCount = p.IssueRequestCount
	}
	var status *string
	if p.Status != nil {
		status = p.Status
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
		"issueJobId":        issueJobId,
		"name":              name,
		"metadata":          metadata,
		"issuedCount":       issuedCount,
		"issueRequestCount": issueRequestCount,
		"status":            status,
		"createdAt":         createdAt,
		"revision":          revision,
	}
}

func (p IssueJob) Pointer() *IssueJob {
	return &p
}

func CastIssueJobs(data []interface{}) []IssueJob {
	v := make([]IssueJob, 0)
	for _, d := range data {
		v = append(v, NewIssueJobFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIssueJobsFromDict(data []IssueJob) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SerialKey struct {
	SerialKeyId       *string `json:"serialKeyId"`
	CampaignModelName *string `json:"campaignModelName"`
	Code              *string `json:"code"`
	Metadata          *string `json:"metadata"`
	Status            *string `json:"status"`
	UsedUserId        *string `json:"usedUserId"`
	CreatedAt         *int64  `json:"createdAt"`
	UsedAt            *int64  `json:"usedAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
	Revision          *int64  `json:"revision"`
}

func NewSerialKeyFromJson(data string) SerialKey {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSerialKeyFromDict(dict)
}

func NewSerialKeyFromDict(data map[string]interface{}) SerialKey {
	return SerialKey{
		SerialKeyId:       core.CastString(data["serialKeyId"]),
		CampaignModelName: core.CastString(data["campaignModelName"]),
		Code:              core.CastString(data["code"]),
		Metadata:          core.CastString(data["metadata"]),
		Status:            core.CastString(data["status"]),
		UsedUserId:        core.CastString(data["usedUserId"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UsedAt:            core.CastInt64(data["usedAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p SerialKey) ToDict() map[string]interface{} {

	var serialKeyId *string
	if p.SerialKeyId != nil {
		serialKeyId = p.SerialKeyId
	}
	var campaignModelName *string
	if p.CampaignModelName != nil {
		campaignModelName = p.CampaignModelName
	}
	var code *string
	if p.Code != nil {
		code = p.Code
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var status *string
	if p.Status != nil {
		status = p.Status
	}
	var usedUserId *string
	if p.UsedUserId != nil {
		usedUserId = p.UsedUserId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var usedAt *int64
	if p.UsedAt != nil {
		usedAt = p.UsedAt
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
		"serialKeyId":       serialKeyId,
		"campaignModelName": campaignModelName,
		"code":              code,
		"metadata":          metadata,
		"status":            status,
		"usedUserId":        usedUserId,
		"createdAt":         createdAt,
		"usedAt":            usedAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
	}
}

func (p SerialKey) Pointer() *SerialKey {
	return &p
}

func CastSerialKeys(data []interface{}) []SerialKey {
	v := make([]SerialKey, 0)
	for _, d := range data {
		v = append(v, NewSerialKeyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSerialKeysFromDict(data []SerialKey) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CampaignModel struct {
	CampaignId         *string `json:"campaignId"`
	Name               *string `json:"name"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
}

func NewCampaignModelFromJson(data string) CampaignModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCampaignModelFromDict(dict)
}

func NewCampaignModelFromDict(data map[string]interface{}) CampaignModel {
	return CampaignModel{
		CampaignId:         core.CastString(data["campaignId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		EnableCampaignCode: core.CastBool(data["enableCampaignCode"]),
	}
}

func (p CampaignModel) ToDict() map[string]interface{} {

	var campaignId *string
	if p.CampaignId != nil {
		campaignId = p.CampaignId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var enableCampaignCode *bool
	if p.EnableCampaignCode != nil {
		enableCampaignCode = p.EnableCampaignCode
	}
	return map[string]interface{}{
		"campaignId":         campaignId,
		"name":               name,
		"metadata":           metadata,
		"enableCampaignCode": enableCampaignCode,
	}
}

func (p CampaignModel) Pointer() *CampaignModel {
	return &p
}

func CastCampaignModels(data []interface{}) []CampaignModel {
	v := make([]CampaignModel, 0)
	for _, d := range data {
		v = append(v, NewCampaignModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCampaignModelsFromDict(data []CampaignModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CampaignModelMaster struct {
	CampaignId         *string `json:"campaignId"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	EnableCampaignCode *bool   `json:"enableCampaignCode"`
	CreatedAt          *int64  `json:"createdAt"`
	UpdatedAt          *int64  `json:"updatedAt"`
	Revision           *int64  `json:"revision"`
}

func NewCampaignModelMasterFromJson(data string) CampaignModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCampaignModelMasterFromDict(dict)
}

func NewCampaignModelMasterFromDict(data map[string]interface{}) CampaignModelMaster {
	return CampaignModelMaster{
		CampaignId:         core.CastString(data["campaignId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		EnableCampaignCode: core.CastBool(data["enableCampaignCode"]),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p CampaignModelMaster) ToDict() map[string]interface{} {

	var campaignId *string
	if p.CampaignId != nil {
		campaignId = p.CampaignId
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
	var enableCampaignCode *bool
	if p.EnableCampaignCode != nil {
		enableCampaignCode = p.EnableCampaignCode
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
		"campaignId":         campaignId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"enableCampaignCode": enableCampaignCode,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p CampaignModelMaster) Pointer() *CampaignModelMaster {
	return &p
}

func CastCampaignModelMasters(data []interface{}) []CampaignModelMaster {
	v := make([]CampaignModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCampaignModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCampaignModelMastersFromDict(data []CampaignModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentCampaignMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentCampaignMasterFromJson(data string) CurrentCampaignMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentCampaignMasterFromDict(dict)
}

func NewCurrentCampaignMasterFromDict(data map[string]interface{}) CurrentCampaignMaster {
	return CurrentCampaignMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentCampaignMaster) ToDict() map[string]interface{} {

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

func (p CurrentCampaignMaster) Pointer() *CurrentCampaignMaster {
	return &p
}

func CastCurrentCampaignMasters(data []interface{}) []CurrentCampaignMaster {
	v := make([]CurrentCampaignMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentCampaignMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentCampaignMastersFromDict(data []CurrentCampaignMaster) []interface{} {
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
