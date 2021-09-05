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

package quest

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId         *string        `json:"namespaceId"`
	Name                *string        `json:"name"`
	Description         *string        `json:"description"`
	StartQuestScript    *ScriptSetting `json:"startQuestScript"`
	CompleteQuestScript *ScriptSetting `json:"completeQuestScript"`
	FailedQuestScript   *ScriptSetting `json:"failedQuestScript"`
	QueueNamespaceId    *string        `json:"queueNamespaceId"`
	KeyId               *string        `json:"keyId"`
	LogSetting          *LogSetting    `json:"logSetting"`
	CreatedAt           *int64         `json:"createdAt"`
	UpdatedAt           *int64         `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:         core.CastString(data["namespaceId"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		StartQuestScript:    NewScriptSettingFromDict(core.CastMap(data["startQuestScript"])).Pointer(),
		CompleteQuestScript: NewScriptSettingFromDict(core.CastMap(data["completeQuestScript"])).Pointer(),
		FailedQuestScript:   NewScriptSettingFromDict(core.CastMap(data["failedQuestScript"])).Pointer(),
		QueueNamespaceId:    core.CastString(data["queueNamespaceId"]),
		KeyId:               core.CastString(data["keyId"]),
		LogSetting:          NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":         p.NamespaceId,
		"name":                p.Name,
		"description":         p.Description,
		"startQuestScript":    p.StartQuestScript.ToDict(),
		"completeQuestScript": p.CompleteQuestScript.ToDict(),
		"failedQuestScript":   p.FailedQuestScript.ToDict(),
		"queueNamespaceId":    p.QueueNamespaceId,
		"keyId":               p.KeyId,
		"logSetting":          p.LogSetting.ToDict(),
		"createdAt":           p.CreatedAt,
		"updatedAt":           p.UpdatedAt,
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

type QuestGroupModelMaster struct {
	QuestGroupModelId      *string `json:"questGroupModelId"`
	Name                   *string `json:"name"`
	Description            *string `json:"description"`
	Metadata               *string `json:"metadata"`
	ChallengePeriodEventId *string `json:"challengePeriodEventId"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
}

func NewQuestGroupModelMasterFromJson(data string) QuestGroupModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestGroupModelMasterFromDict(dict)
}

func NewQuestGroupModelMasterFromDict(data map[string]interface{}) QuestGroupModelMaster {
	return QuestGroupModelMaster{
		QuestGroupModelId:      core.CastString(data["questGroupModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
	}
}

func (p QuestGroupModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"questGroupModelId":      p.QuestGroupModelId,
		"name":                   p.Name,
		"description":            p.Description,
		"metadata":               p.Metadata,
		"challengePeriodEventId": p.ChallengePeriodEventId,
		"createdAt":              p.CreatedAt,
		"updatedAt":              p.UpdatedAt,
	}
}

func (p QuestGroupModelMaster) Pointer() *QuestGroupModelMaster {
	return &p
}

func CastQuestGroupModelMasters(data []interface{}) []QuestGroupModelMaster {
	v := make([]QuestGroupModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelMastersFromDict(data []QuestGroupModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestModelMaster struct {
	QuestModelId           *string         `json:"questModelId"`
	QuestGroupName         *string         `json:"questGroupName"`
	Name                   *string         `json:"name"`
	Description            *string         `json:"description"`
	Metadata               *string         `json:"metadata"`
	Contents               []Contents      `json:"contents"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	ConsumeActions         []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions   []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames      []string        `json:"premiseQuestNames"`
	CreatedAt              *int64          `json:"createdAt"`
	UpdatedAt              *int64          `json:"updatedAt"`
}

func NewQuestModelMasterFromJson(data string) QuestModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestModelMasterFromDict(dict)
}

func NewQuestModelMasterFromDict(data map[string]interface{}) QuestModelMaster {
	return QuestModelMaster{
		QuestModelId:           core.CastString(data["questModelId"]),
		QuestGroupName:         core.CastString(data["questGroupName"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		Contents:               CastContentses(core.CastArray(data["contents"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		ConsumeActions:         CastConsumeActions(core.CastArray(data["consumeActions"])),
		FailedAcquireActions:   CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
		PremiseQuestNames:      core.CastStrings(core.CastArray(data["premiseQuestNames"])),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
	}
}

func (p QuestModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"questModelId":   p.QuestModelId,
		"questGroupName": p.QuestGroupName,
		"name":           p.Name,
		"description":    p.Description,
		"metadata":       p.Metadata,
		"contents": CastContentsesFromDict(
			p.Contents,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"failedAcquireActions": CastAcquireActionsFromDict(
			p.FailedAcquireActions,
		),
		"premiseQuestNames": core.CastStringsFromDict(
			p.PremiseQuestNames,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p QuestModelMaster) Pointer() *QuestModelMaster {
	return &p
}

func CastQuestModelMasters(data []interface{}) []QuestModelMaster {
	v := make([]QuestModelMaster, 0)
	for _, d := range data {
		v = append(v, NewQuestModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelMastersFromDict(data []QuestModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentQuestMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentQuestMasterFromJson(data string) CurrentQuestMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentQuestMasterFromDict(dict)
}

func NewCurrentQuestMasterFromDict(data map[string]interface{}) CurrentQuestMaster {
	return CurrentQuestMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentQuestMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentQuestMaster) Pointer() *CurrentQuestMaster {
	return &p
}

func CastCurrentQuestMasters(data []interface{}) []CurrentQuestMaster {
	v := make([]CurrentQuestMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentQuestMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentQuestMastersFromDict(data []CurrentQuestMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Contents struct {
	Metadata               *string         `json:"metadata"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	Weight                 *int32          `json:"weight"`
}

func NewContentsFromJson(data string) Contents {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewContentsFromDict(dict)
}

func NewContentsFromDict(data map[string]interface{}) Contents {
	return Contents{
		Metadata:               core.CastString(data["metadata"]),
		CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		Weight:                 core.CastInt32(data["weight"]),
	}
}

func (p Contents) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": p.Metadata,
		"completeAcquireActions": CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		),
		"weight": p.Weight,
	}
}

func (p Contents) Pointer() *Contents {
	return &p
}

func CastContentses(data []interface{}) []Contents {
	v := make([]Contents, 0)
	for _, d := range data {
		v = append(v, NewContentsFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastContentsesFromDict(data []Contents) []interface{} {
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

type Reward struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
	ItemId  *string `json:"itemId"`
	Value   *int32  `json:"value"`
}

func NewRewardFromJson(data string) Reward {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRewardFromDict(dict)
}

func NewRewardFromDict(data map[string]interface{}) Reward {
	return Reward{
		Action:  core.CastString(data["action"]),
		Request: core.CastString(data["request"]),
		ItemId:  core.CastString(data["itemId"]),
		Value:   core.CastInt32(data["value"]),
	}
}

func (p Reward) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"action":  p.Action,
		"request": p.Request,
		"itemId":  p.ItemId,
		"value":   p.Value,
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

type Progress struct {
	ProgressId    *string  `json:"progressId"`
	UserId        *string  `json:"userId"`
	TransactionId *string  `json:"transactionId"`
	QuestModelId  *string  `json:"questModelId"`
	RandomSeed    *int64   `json:"randomSeed"`
	Rewards       []Reward `json:"rewards"`
	Metadata      *string  `json:"metadata"`
	CreatedAt     *int64   `json:"createdAt"`
	UpdatedAt     *int64   `json:"updatedAt"`
}

func NewProgressFromJson(data string) Progress {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewProgressFromDict(dict)
}

func NewProgressFromDict(data map[string]interface{}) Progress {
	return Progress{
		ProgressId:    core.CastString(data["progressId"]),
		UserId:        core.CastString(data["userId"]),
		TransactionId: core.CastString(data["transactionId"]),
		QuestModelId:  core.CastString(data["questModelId"]),
		RandomSeed:    core.CastInt64(data["randomSeed"]),
		Rewards:       CastRewards(core.CastArray(data["rewards"])),
		Metadata:      core.CastString(data["metadata"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
	}
}

func (p Progress) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"progressId":    p.ProgressId,
		"userId":        p.UserId,
		"transactionId": p.TransactionId,
		"questModelId":  p.QuestModelId,
		"randomSeed":    p.RandomSeed,
		"rewards": CastRewardsFromDict(
			p.Rewards,
		),
		"metadata":  p.Metadata,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p Progress) Pointer() *Progress {
	return &p
}

func CastProgresses(data []interface{}) []Progress {
	v := make([]Progress, 0)
	for _, d := range data {
		v = append(v, NewProgressFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastProgressesFromDict(data []Progress) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CompletedQuestList struct {
	CompletedQuestListId *string  `json:"completedQuestListId"`
	UserId               *string  `json:"userId"`
	QuestGroupName       *string  `json:"questGroupName"`
	CompleteQuestNames   []string `json:"completeQuestNames"`
	CreatedAt            *int64   `json:"createdAt"`
	UpdatedAt            *int64   `json:"updatedAt"`
}

func NewCompletedQuestListFromJson(data string) CompletedQuestList {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCompletedQuestListFromDict(dict)
}

func NewCompletedQuestListFromDict(data map[string]interface{}) CompletedQuestList {
	return CompletedQuestList{
		CompletedQuestListId: core.CastString(data["completedQuestListId"]),
		UserId:               core.CastString(data["userId"]),
		QuestGroupName:       core.CastString(data["questGroupName"]),
		CompleteQuestNames:   core.CastStrings(core.CastArray(data["completeQuestNames"])),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
	}
}

func (p CompletedQuestList) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"completedQuestListId": p.CompletedQuestListId,
		"userId":               p.UserId,
		"questGroupName":       p.QuestGroupName,
		"completeQuestNames": core.CastStringsFromDict(
			p.CompleteQuestNames,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p CompletedQuestList) Pointer() *CompletedQuestList {
	return &p
}

func CastCompletedQuestLists(data []interface{}) []CompletedQuestList {
	v := make([]CompletedQuestList, 0)
	for _, d := range data {
		v = append(v, NewCompletedQuestListFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCompletedQuestListsFromDict(data []CompletedQuestList) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestGroupModel struct {
	QuestGroupModelId      *string      `json:"questGroupModelId"`
	Name                   *string      `json:"name"`
	Metadata               *string      `json:"metadata"`
	Quests                 []QuestModel `json:"quests"`
	ChallengePeriodEventId *string      `json:"challengePeriodEventId"`
}

func NewQuestGroupModelFromJson(data string) QuestGroupModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestGroupModelFromDict(dict)
}

func NewQuestGroupModelFromDict(data map[string]interface{}) QuestGroupModel {
	return QuestGroupModel{
		QuestGroupModelId:      core.CastString(data["questGroupModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Quests:                 CastQuestModels(core.CastArray(data["quests"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p QuestGroupModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"questGroupModelId": p.QuestGroupModelId,
		"name":              p.Name,
		"metadata":          p.Metadata,
		"quests": CastQuestModelsFromDict(
			p.Quests,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p QuestGroupModel) Pointer() *QuestGroupModel {
	return &p
}

func CastQuestGroupModels(data []interface{}) []QuestGroupModel {
	v := make([]QuestGroupModel, 0)
	for _, d := range data {
		v = append(v, NewQuestGroupModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestGroupModelsFromDict(data []QuestGroupModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type QuestModel struct {
	QuestModelId           *string         `json:"questModelId"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	Contents               []Contents      `json:"contents"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	ConsumeActions         []ConsumeAction `json:"consumeActions"`
	FailedAcquireActions   []AcquireAction `json:"failedAcquireActions"`
	PremiseQuestNames      []string        `json:"premiseQuestNames"`
}

func NewQuestModelFromJson(data string) QuestModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewQuestModelFromDict(dict)
}

func NewQuestModelFromDict(data map[string]interface{}) QuestModel {
	return QuestModel{
		QuestModelId:           core.CastString(data["questModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Contents:               CastContentses(core.CastArray(data["contents"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		ConsumeActions:         CastConsumeActions(core.CastArray(data["consumeActions"])),
		FailedAcquireActions:   CastAcquireActions(core.CastArray(data["failedAcquireActions"])),
		PremiseQuestNames:      core.CastStrings(core.CastArray(data["premiseQuestNames"])),
	}
}

func (p QuestModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"questModelId": p.QuestModelId,
		"name":         p.Name,
		"metadata":     p.Metadata,
		"contents": CastContentsesFromDict(
			p.Contents,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"failedAcquireActions": CastAcquireActionsFromDict(
			p.FailedAcquireActions,
		),
		"premiseQuestNames": core.CastStringsFromDict(
			p.PremiseQuestNames,
		),
	}
}

func (p QuestModel) Pointer() *QuestModel {
	return &p
}

func CastQuestModels(data []interface{}) []QuestModel {
	v := make([]QuestModel, 0)
	for _, d := range data {
		v = append(v, NewQuestModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastQuestModelsFromDict(data []QuestModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
