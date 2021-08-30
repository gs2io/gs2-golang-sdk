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

package exchange

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId          *string     `json:"namespaceId"`
	Name                 *string     `json:"name"`
	Description          *string     `json:"description"`
	EnableDirectExchange *bool       `json:"enableDirectExchange"`
	EnableAwaitExchange  *bool       `json:"enableAwaitExchange"`
	QueueNamespaceId     *string     `json:"queueNamespaceId"`
	KeyId                *string     `json:"keyId"`
	LogSetting           *LogSetting `json:"logSetting"`
	CreatedAt            *int64      `json:"createdAt"`
	UpdatedAt            *int64      `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:          core.CastString(data["namespaceId"]),
		Name:                 core.CastString(data["name"]),
		Description:          core.CastString(data["description"]),
		EnableDirectExchange: core.CastBool(data["enableDirectExchange"]),
		EnableAwaitExchange:  core.CastBool(data["enableAwaitExchange"]),
		QueueNamespaceId:     core.CastString(data["queueNamespaceId"]),
		KeyId:                core.CastString(data["keyId"]),
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
		"enableDirectExchange": p.EnableDirectExchange,
		"enableAwaitExchange":  p.EnableAwaitExchange,
		"queueNamespaceId":     p.QueueNamespaceId,
		"keyId":                p.KeyId,
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

type RateModel struct {
	RateModelId        *string         `json:"rateModelId"`
	Name               *string         `json:"name"`
	Metadata           *string         `json:"metadata"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
}

func NewRateModelFromDict(data map[string]interface{}) RateModel {
	return RateModel{
		RateModelId:        core.CastString(data["rateModelId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p RateModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rateModelId": p.RateModelId,
		"name":        p.Name,
		"metadata":    p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"timingType": p.TimingType,
		"lockTime":   p.LockTime,
		"enableSkip": p.EnableSkip,
		"skipConsumeActions": CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
	}
}

func (p RateModel) Pointer() *RateModel {
	return &p
}

func CastRateModels(data []interface{}) []RateModel {
	v := make([]RateModel, 0)
	for _, d := range data {
		v = append(v, NewRateModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRateModelsFromDict(data []RateModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RateModelMaster struct {
	RateModelId        *string         `json:"rateModelId"`
	Name               *string         `json:"name"`
	Description        *string         `json:"description"`
	Metadata           *string         `json:"metadata"`
	ConsumeActions     []ConsumeAction `json:"consumeActions"`
	TimingType         *string         `json:"timingType"`
	LockTime           *int32          `json:"lockTime"`
	EnableSkip         *bool           `json:"enableSkip"`
	SkipConsumeActions []ConsumeAction `json:"skipConsumeActions"`
	AcquireActions     []AcquireAction `json:"acquireActions"`
	CreatedAt          *int64          `json:"createdAt"`
	UpdatedAt          *int64          `json:"updatedAt"`
}

func NewRateModelMasterFromDict(data map[string]interface{}) RateModelMaster {
	return RateModelMaster{
		RateModelId:        core.CastString(data["rateModelId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		ConsumeActions:     CastConsumeActions(core.CastArray(data["consumeActions"])),
		TimingType:         core.CastString(data["timingType"]),
		LockTime:           core.CastInt32(data["lockTime"]),
		EnableSkip:         core.CastBool(data["enableSkip"]),
		SkipConsumeActions: CastConsumeActions(core.CastArray(data["skipConsumeActions"])),
		AcquireActions:     CastAcquireActions(core.CastArray(data["acquireActions"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
	}
}

func (p RateModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rateModelId": p.RateModelId,
		"name":        p.Name,
		"description": p.Description,
		"metadata":    p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"timingType": p.TimingType,
		"lockTime":   p.LockTime,
		"enableSkip": p.EnableSkip,
		"skipConsumeActions": CastConsumeActionsFromDict(
			p.SkipConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p RateModelMaster) Pointer() *RateModelMaster {
	return &p
}

func CastRateModelMasters(data []interface{}) []RateModelMaster {
	v := make([]RateModelMaster, 0)
	for _, d := range data {
		v = append(v, NewRateModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRateModelMastersFromDict(data []RateModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentRateMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentRateMasterFromDict(data map[string]interface{}) CurrentRateMaster {
	return CurrentRateMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRateMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentRateMaster) Pointer() *CurrentRateMaster {
	return &p
}

func CastCurrentRateMasters(data []interface{}) []CurrentRateMaster {
	v := make([]CurrentRateMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentRateMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentRateMastersFromDict(data []CurrentRateMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Await struct {
	AwaitId     *string `json:"awaitId"`
	UserId      *string `json:"userId"`
	RateName    *string `json:"rateName"`
	Name        *string `json:"name"`
	Count       *int32  `json:"count"`
	ExchangedAt *int64  `json:"exchangedAt"`
}

func NewAwaitFromDict(data map[string]interface{}) Await {
	return Await{
		AwaitId:     core.CastString(data["awaitId"]),
		UserId:      core.CastString(data["userId"]),
		RateName:    core.CastString(data["rateName"]),
		Name:        core.CastString(data["name"]),
		Count:       core.CastInt32(data["count"]),
		ExchangedAt: core.CastInt64(data["exchangedAt"]),
	}
}

func (p Await) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"awaitId":     p.AwaitId,
		"userId":      p.UserId,
		"rateName":    p.RateName,
		"name":        p.Name,
		"count":       p.Count,
		"exchangedAt": p.ExchangedAt,
	}
}

func (p Await) Pointer() *Await {
	return &p
}

func CastAwaits(data []interface{}) []Await {
	v := make([]Await, 0)
	for _, d := range data {
		v = append(v, NewAwaitFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAwaitsFromDict(data []Await) []interface{} {
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

type AcquireAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewAcquireActionFromDict(data map[string]interface{}) AcquireAction {
	return AcquireAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p AcquireAction) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"action":  p.Action,
		"request": p.Request,
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

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
}

func NewConsumeActionFromDict(data map[string]interface{}) ConsumeAction {
	return ConsumeAction{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
	}
}

func (p ConsumeAction) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"action":  p.Action,
		"request": p.Request,
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
