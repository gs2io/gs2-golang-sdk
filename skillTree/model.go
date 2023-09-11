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

package skillTree

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ReleaseScript      *ScriptSetting      `json:"releaseScript"`
	RestrainScript     *ScriptSetting      `json:"restrainScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:        core.CastString(data["namespaceId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		ReleaseScript:      NewScriptSettingFromDict(core.CastMap(data["releaseScript"])).Pointer(),
		RestrainScript:     NewScriptSettingFromDict(core.CastMap(data["restrainScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
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
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var releaseScript map[string]interface{}
	if p.ReleaseScript != nil {
		releaseScript = p.ReleaseScript.ToDict()
	}
	var restrainScript map[string]interface{}
	if p.RestrainScript != nil {
		restrainScript = p.RestrainScript.ToDict()
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
		"namespaceId":        namespaceId,
		"name":               name,
		"description":        description,
		"transactionSetting": transactionSetting,
		"releaseScript":      releaseScript,
		"restrainScript":     restrainScript,
		"logSetting":         logSetting,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
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

type NodeModel struct {
	NodeModelId           *string         `json:"nodeModelId"`
	Name                  *string         `json:"name"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	ReturnAcquireActions  []AcquireAction `json:"returnAcquireActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
}

func NewNodeModelFromJson(data string) NodeModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNodeModelFromDict(dict)
}

func NewNodeModelFromDict(data map[string]interface{}) NodeModel {
	return NodeModel{
		NodeModelId:           core.CastString(data["nodeModelId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		ReturnAcquireActions:  CastAcquireActions(core.CastArray(data["returnAcquireActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
	}
}

func (p NodeModel) ToDict() map[string]interface{} {

	var nodeModelId *string
	if p.NodeModelId != nil {
		nodeModelId = p.NodeModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var releaseConsumeActions []interface{}
	if p.ReleaseConsumeActions != nil {
		releaseConsumeActions = CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		)
	}
	var returnAcquireActions []interface{}
	if p.ReturnAcquireActions != nil {
		returnAcquireActions = CastAcquireActionsFromDict(
			p.ReturnAcquireActions,
		)
	}
	var restrainReturnRate *float32
	if p.RestrainReturnRate != nil {
		restrainReturnRate = p.RestrainReturnRate
	}
	var premiseNodeNames []interface{}
	if p.PremiseNodeNames != nil {
		premiseNodeNames = core.CastStringsFromDict(
			p.PremiseNodeNames,
		)
	}
	return map[string]interface{}{
		"nodeModelId":           nodeModelId,
		"name":                  name,
		"metadata":              metadata,
		"releaseConsumeActions": releaseConsumeActions,
		"returnAcquireActions":  returnAcquireActions,
		"restrainReturnRate":    restrainReturnRate,
		"premiseNodeNames":      premiseNodeNames,
	}
}

func (p NodeModel) Pointer() *NodeModel {
	return &p
}

func CastNodeModels(data []interface{}) []NodeModel {
	v := make([]NodeModel, 0)
	for _, d := range data {
		v = append(v, NewNodeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNodeModelsFromDict(data []NodeModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NodeModelMaster struct {
	NodeModelId           *string         `json:"nodeModelId"`
	Name                  *string         `json:"name"`
	Description           *string         `json:"description"`
	Metadata              *string         `json:"metadata"`
	ReleaseConsumeActions []ConsumeAction `json:"releaseConsumeActions"`
	RestrainReturnRate    *float32        `json:"restrainReturnRate"`
	PremiseNodeNames      []*string       `json:"premiseNodeNames"`
	CreatedAt             *int64          `json:"createdAt"`
	UpdatedAt             *int64          `json:"updatedAt"`
	Revision              *int64          `json:"revision"`
}

func NewNodeModelMasterFromJson(data string) NodeModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNodeModelMasterFromDict(dict)
}

func NewNodeModelMasterFromDict(data map[string]interface{}) NodeModelMaster {
	return NodeModelMaster{
		NodeModelId:           core.CastString(data["nodeModelId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		ReleaseConsumeActions: CastConsumeActions(core.CastArray(data["releaseConsumeActions"])),
		RestrainReturnRate:    core.CastFloat32(data["restrainReturnRate"]),
		PremiseNodeNames:      core.CastStrings(core.CastArray(data["premiseNodeNames"])),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p NodeModelMaster) ToDict() map[string]interface{} {

	var nodeModelId *string
	if p.NodeModelId != nil {
		nodeModelId = p.NodeModelId
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
	var releaseConsumeActions []interface{}
	if p.ReleaseConsumeActions != nil {
		releaseConsumeActions = CastConsumeActionsFromDict(
			p.ReleaseConsumeActions,
		)
	}
	var restrainReturnRate *float32
	if p.RestrainReturnRate != nil {
		restrainReturnRate = p.RestrainReturnRate
	}
	var premiseNodeNames []interface{}
	if p.PremiseNodeNames != nil {
		premiseNodeNames = core.CastStringsFromDict(
			p.PremiseNodeNames,
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"nodeModelId":           nodeModelId,
		"name":                  name,
		"description":           description,
		"metadata":              metadata,
		"releaseConsumeActions": releaseConsumeActions,
		"restrainReturnRate":    restrainReturnRate,
		"premiseNodeNames":      premiseNodeNames,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
	}
}

func (p NodeModelMaster) Pointer() *NodeModelMaster {
	return &p
}

func CastNodeModelMasters(data []interface{}) []NodeModelMaster {
	v := make([]NodeModelMaster, 0)
	for _, d := range data {
		v = append(v, NewNodeModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNodeModelMastersFromDict(data []NodeModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	StatusId          *string   `json:"statusId"`
	UserId            *string   `json:"userId"`
	ReleasedNodeNames []*string `json:"releasedNodeNames"`
	CreatedAt         *int64    `json:"createdAt"`
	UpdatedAt         *int64    `json:"updatedAt"`
	Revision          *int64    `json:"revision"`
}

func NewStatusFromJson(data string) Status {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		StatusId:          core.CastString(data["statusId"]),
		UserId:            core.CastString(data["userId"]),
		ReleasedNodeNames: core.CastStrings(core.CastArray(data["releasedNodeNames"])),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
		Revision:          core.CastInt64(data["revision"]),
	}
}

func (p Status) ToDict() map[string]interface{} {

	var statusId *string
	if p.StatusId != nil {
		statusId = p.StatusId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var releasedNodeNames []interface{}
	if p.ReleasedNodeNames != nil {
		releasedNodeNames = core.CastStringsFromDict(
			p.ReleasedNodeNames,
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"statusId":          statusId,
		"userId":            userId,
		"releasedNodeNames": releasedNodeNames,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
		"revision":          revision,
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

type CurrentTreeMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentTreeMasterFromJson(data string) CurrentTreeMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentTreeMasterFromDict(dict)
}

func NewCurrentTreeMasterFromDict(data map[string]interface{}) CurrentTreeMaster {
	return CurrentTreeMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentTreeMaster) ToDict() map[string]interface{} {

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

func (p CurrentTreeMaster) Pointer() *CurrentTreeMaster {
	return &p
}

func CastCurrentTreeMasters(data []interface{}) []CurrentTreeMaster {
	v := make([]CurrentTreeMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentTreeMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentTreeMastersFromDict(data []CurrentTreeMaster) []interface{} {
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

func NewConfigFromJson(data string) Config {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConfigFromDict(dict)
}

func NewConfigFromDict(data map[string]interface{}) Config {
	return Config{
		Key:   core.CastString(data["key"]),
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
	return map[string]interface{}{
		"key":   key,
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

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewConsumeActionFromJson(data string) ConsumeAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeActionFromDict(dict)
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
	return ConsumeAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p ConsumeAction) ToDict() map[string]interface{} {

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

func (p ConsumeAction) Pointer() *ConsumeAction {
	return &p
}

func CastConsumeActions(data []interface{}) []ConsumeAction {
	v := make([]ConsumeAction, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionsFromDict(data []ConsumeAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewAcquireActionFromJson(data string) AcquireAction {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionFromDict(dict)
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {

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

func (p AcquireAction) Pointer() *AcquireAction {
	return &p
}

func CastAcquireActions(data []interface{}) []AcquireAction {
	v := make([]AcquireAction, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionsFromDict(data []AcquireAction) []interface{} {
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

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewScriptSettingFromDict(dict)
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
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
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

type TransactionSetting struct {
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
	KeyId                  *string `json:"keyId"`
	QueueNamespaceId       *string `json:"queueNamespaceId"`
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTransactionSettingFromDict(dict)
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
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
