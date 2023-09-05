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

package loginReward

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ReceiveScript      *ScriptSetting      `json:"receiveScript"`
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
		ReceiveScript:      NewScriptSettingFromDict(core.CastMap(data["receiveScript"])).Pointer(),
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
	var receiveScript map[string]interface{}
	if p.ReceiveScript != nil {
		receiveScript = p.ReceiveScript.ToDict()
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
		"receiveScript":      receiveScript,
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

type BonusModelMaster struct {
	BonusModelId                      *string         `json:"bonusModelId"`
	Name                              *string         `json:"name"`
	Description                       *string         `json:"description"`
	Metadata                          *string         `json:"metadata"`
	Mode                              *string         `json:"mode"`
	PeriodEventId                     *string         `json:"periodEventId"`
	ResetHour                         *int32          `json:"resetHour"`
	Repeat                            *string         `json:"repeat"`
	Rewards                           []Reward        `json:"rewards"`
	MissedReceiveRelief               *string         `json:"missedReceiveRelief"`
	MissedReceiveReliefConsumeActions []ConsumeAction `json:"missedReceiveReliefConsumeActions"`
	CreatedAt                         *int64          `json:"createdAt"`
	UpdatedAt                         *int64          `json:"updatedAt"`
	Revision                          *int64          `json:"revision"`
}

func NewBonusModelMasterFromJson(data string) BonusModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBonusModelMasterFromDict(dict)
}

func NewBonusModelMasterFromDict(data map[string]interface{}) BonusModelMaster {
	return BonusModelMaster{
		BonusModelId:                      core.CastString(data["bonusModelId"]),
		Name:                              core.CastString(data["name"]),
		Description:                       core.CastString(data["description"]),
		Metadata:                          core.CastString(data["metadata"]),
		Mode:                              core.CastString(data["mode"]),
		PeriodEventId:                     core.CastString(data["periodEventId"]),
		ResetHour:                         core.CastInt32(data["resetHour"]),
		Repeat:                            core.CastString(data["repeat"]),
		Rewards:                           CastRewards(core.CastArray(data["rewards"])),
		MissedReceiveRelief:               core.CastString(data["missedReceiveRelief"]),
		MissedReceiveReliefConsumeActions: CastConsumeActions(core.CastArray(data["missedReceiveReliefConsumeActions"])),
		CreatedAt:                         core.CastInt64(data["createdAt"]),
		UpdatedAt:                         core.CastInt64(data["updatedAt"]),
		Revision:                          core.CastInt64(data["revision"]),
	}
}

func (p BonusModelMaster) ToDict() map[string]interface{} {

	var bonusModelId *string
	if p.BonusModelId != nil {
		bonusModelId = p.BonusModelId
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
	var mode *string
	if p.Mode != nil {
		mode = p.Mode
	}
	var periodEventId *string
	if p.PeriodEventId != nil {
		periodEventId = p.PeriodEventId
	}
	var resetHour *int32
	if p.ResetHour != nil {
		resetHour = p.ResetHour
	}
	var repeat *string
	if p.Repeat != nil {
		repeat = p.Repeat
	}
	var rewards []interface{}
	if p.Rewards != nil {
		rewards = CastRewardsFromDict(
			p.Rewards,
		)
	}
	var missedReceiveRelief *string
	if p.MissedReceiveRelief != nil {
		missedReceiveRelief = p.MissedReceiveRelief
	}
	var missedReceiveReliefConsumeActions []interface{}
	if p.MissedReceiveReliefConsumeActions != nil {
		missedReceiveReliefConsumeActions = CastConsumeActionsFromDict(
			p.MissedReceiveReliefConsumeActions,
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
		"bonusModelId":                      bonusModelId,
		"name":                              name,
		"description":                       description,
		"metadata":                          metadata,
		"mode":                              mode,
		"periodEventId":                     periodEventId,
		"resetHour":                         resetHour,
		"repeat":                            repeat,
		"rewards":                           rewards,
		"missedReceiveRelief":               missedReceiveRelief,
		"missedReceiveReliefConsumeActions": missedReceiveReliefConsumeActions,
		"createdAt":                         createdAt,
		"updatedAt":                         updatedAt,
		"revision":                          revision,
	}
}

func (p BonusModelMaster) Pointer() *BonusModelMaster {
	return &p
}

func CastBonusModelMasters(data []interface{}) []BonusModelMaster {
	v := make([]BonusModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBonusModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBonusModelMastersFromDict(data []BonusModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentBonusMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentBonusMasterFromJson(data string) CurrentBonusMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentBonusMasterFromDict(dict)
}

func NewCurrentBonusMasterFromDict(data map[string]interface{}) CurrentBonusMaster {
	return CurrentBonusMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentBonusMaster) ToDict() map[string]interface{} {

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

func (p CurrentBonusMaster) Pointer() *CurrentBonusMaster {
	return &p
}

func CastCurrentBonusMasters(data []interface{}) []CurrentBonusMaster {
	v := make([]CurrentBonusMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentBonusMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentBonusMastersFromDict(data []CurrentBonusMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BonusModel struct {
	BonusModelId                      *string         `json:"bonusModelId"`
	Name                              *string         `json:"name"`
	Metadata                          *string         `json:"metadata"`
	Mode                              *string         `json:"mode"`
	PeriodEventId                     *string         `json:"periodEventId"`
	ResetHour                         *int32          `json:"resetHour"`
	Repeat                            *string         `json:"repeat"`
	Rewards                           []Reward        `json:"rewards"`
	MissedReceiveRelief               *string         `json:"missedReceiveRelief"`
	MissedReceiveReliefConsumeActions []ConsumeAction `json:"missedReceiveReliefConsumeActions"`
}

func NewBonusModelFromJson(data string) BonusModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBonusModelFromDict(dict)
}

func NewBonusModelFromDict(data map[string]interface{}) BonusModel {
	return BonusModel{
		BonusModelId:                      core.CastString(data["bonusModelId"]),
		Name:                              core.CastString(data["name"]),
		Metadata:                          core.CastString(data["metadata"]),
		Mode:                              core.CastString(data["mode"]),
		PeriodEventId:                     core.CastString(data["periodEventId"]),
		ResetHour:                         core.CastInt32(data["resetHour"]),
		Repeat:                            core.CastString(data["repeat"]),
		Rewards:                           CastRewards(core.CastArray(data["rewards"])),
		MissedReceiveRelief:               core.CastString(data["missedReceiveRelief"]),
		MissedReceiveReliefConsumeActions: CastConsumeActions(core.CastArray(data["missedReceiveReliefConsumeActions"])),
	}
}

func (p BonusModel) ToDict() map[string]interface{} {

	var bonusModelId *string
	if p.BonusModelId != nil {
		bonusModelId = p.BonusModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var mode *string
	if p.Mode != nil {
		mode = p.Mode
	}
	var periodEventId *string
	if p.PeriodEventId != nil {
		periodEventId = p.PeriodEventId
	}
	var resetHour *int32
	if p.ResetHour != nil {
		resetHour = p.ResetHour
	}
	var repeat *string
	if p.Repeat != nil {
		repeat = p.Repeat
	}
	var rewards []interface{}
	if p.Rewards != nil {
		rewards = CastRewardsFromDict(
			p.Rewards,
		)
	}
	var missedReceiveRelief *string
	if p.MissedReceiveRelief != nil {
		missedReceiveRelief = p.MissedReceiveRelief
	}
	var missedReceiveReliefConsumeActions []interface{}
	if p.MissedReceiveReliefConsumeActions != nil {
		missedReceiveReliefConsumeActions = CastConsumeActionsFromDict(
			p.MissedReceiveReliefConsumeActions,
		)
	}
	return map[string]interface{}{
		"bonusModelId":                      bonusModelId,
		"name":                              name,
		"metadata":                          metadata,
		"mode":                              mode,
		"periodEventId":                     periodEventId,
		"resetHour":                         resetHour,
		"repeat":                            repeat,
		"rewards":                           rewards,
		"missedReceiveRelief":               missedReceiveRelief,
		"missedReceiveReliefConsumeActions": missedReceiveReliefConsumeActions,
	}
}

func (p BonusModel) Pointer() *BonusModel {
	return &p
}

func CastBonusModels(data []interface{}) []BonusModel {
	v := make([]BonusModel, 0)
	for _, d := range data {
		v = append(v, NewBonusModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBonusModelsFromDict(data []BonusModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Reward struct {
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func NewRewardFromJson(data string) Reward {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRewardFromDict(dict)
}

func NewRewardFromDict(data map[string]interface{}) Reward {
	return Reward{
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p Reward) ToDict() map[string]interface{} {

	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"acquireActions": acquireActions,
	}
}

func (p Reward) Pointer() *Reward {
	return &p
}

func CastRewards(data []interface{}) []Reward {
	v := make([]Reward, 0)
	for _, d := range data {
		v = append(v, NewRewardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRewardsFromDict(data []Reward) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ReceiveStatus struct {
	ReceiveStatusId *string `json:"receiveStatusId"`
	BonusModelName  *string `json:"bonusModelName"`
	UserId          *string `json:"userId"`
	ReceivedSteps   []*bool `json:"receivedSteps"`
	LastReceivedAt  *int64  `json:"lastReceivedAt"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
	Revision        *int64  `json:"revision"`
}

func NewReceiveStatusFromJson(data string) ReceiveStatus {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveStatusFromDict(dict)
}

func NewReceiveStatusFromDict(data map[string]interface{}) ReceiveStatus {
	return ReceiveStatus{
		ReceiveStatusId: core.CastString(data["receiveStatusId"]),
		BonusModelName:  core.CastString(data["bonusModelName"]),
		UserId:          core.CastString(data["userId"]),
		ReceivedSteps:   core.CastBools(core.CastArray(data["receivedSteps"])),
		LastReceivedAt:  core.CastInt64(data["lastReceivedAt"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
		Revision:        core.CastInt64(data["revision"]),
	}
}

func (p ReceiveStatus) ToDict() map[string]interface{} {

	var receiveStatusId *string
	if p.ReceiveStatusId != nil {
		receiveStatusId = p.ReceiveStatusId
	}
	var bonusModelName *string
	if p.BonusModelName != nil {
		bonusModelName = p.BonusModelName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var receivedSteps []interface{}
	if p.ReceivedSteps != nil {
		receivedSteps = core.CastBoolsFromDict(
			p.ReceivedSteps,
		)
	}
	var lastReceivedAt *int64
	if p.LastReceivedAt != nil {
		lastReceivedAt = p.LastReceivedAt
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
		"receiveStatusId": receiveStatusId,
		"bonusModelName":  bonusModelName,
		"userId":          userId,
		"receivedSteps":   receivedSteps,
		"lastReceivedAt":  lastReceivedAt,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
		"revision":        revision,
	}
}

func (p ReceiveStatus) Pointer() *ReceiveStatus {
	return &p
}

func CastReceiveStatuses(data []interface{}) []ReceiveStatus {
	v := make([]ReceiveStatus, 0)
	for _, d := range data {
		v = append(v, NewReceiveStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReceiveStatusesFromDict(data []ReceiveStatus) []interface{} {
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
