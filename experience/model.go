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

package experience

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId              *string        `json:"namespaceId"`
	Name                     *string        `json:"name"`
	Description              *string        `json:"description"`
	ExperienceCapScriptId    *string        `json:"experienceCapScriptId"`
	ChangeExperienceScript   *ScriptSetting `json:"changeExperienceScript"`
	ChangeRankScript         *ScriptSetting `json:"changeRankScript"`
	ChangeRankCapScript      *ScriptSetting `json:"changeRankCapScript"`
	OverflowExperienceScript *ScriptSetting `json:"overflowExperienceScript"`
	LogSetting               *LogSetting    `json:"logSetting"`
	CreatedAt                *int64         `json:"createdAt"`
	UpdatedAt                *int64         `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:              core.CastString(data["namespaceId"]),
		Name:                     core.CastString(data["name"]),
		Description:              core.CastString(data["description"]),
		ExperienceCapScriptId:    core.CastString(data["experienceCapScriptId"]),
		ChangeExperienceScript:   NewScriptSettingFromDict(core.CastMap(data["changeExperienceScript"])).Pointer(),
		ChangeRankScript:         NewScriptSettingFromDict(core.CastMap(data["changeRankScript"])).Pointer(),
		ChangeRankCapScript:      NewScriptSettingFromDict(core.CastMap(data["changeRankCapScript"])).Pointer(),
		OverflowExperienceScript: NewScriptSettingFromDict(core.CastMap(data["overflowExperienceScript"])).Pointer(),
		LogSetting:               NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                core.CastInt64(data["createdAt"]),
		UpdatedAt:                core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":              p.NamespaceId,
		"name":                     p.Name,
		"description":              p.Description,
		"experienceCapScriptId":    p.ExperienceCapScriptId,
		"changeExperienceScript":   p.ChangeExperienceScript.ToDict(),
		"changeRankScript":         p.ChangeRankScript.ToDict(),
		"changeRankCapScript":      p.ChangeRankCapScript.ToDict(),
		"overflowExperienceScript": p.OverflowExperienceScript.ToDict(),
		"logSetting":               p.LogSetting.ToDict(),
		"createdAt":                p.CreatedAt,
		"updatedAt":                p.UpdatedAt,
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

type ExperienceModelMaster struct {
	ExperienceModelId *string `json:"experienceModelId"`
	Name              *string `json:"name"`
	Description       *string `json:"description"`
	Metadata          *string `json:"metadata"`
	DefaultExperience *int64  `json:"defaultExperience"`
	DefaultRankCap    *int64  `json:"defaultRankCap"`
	MaxRankCap        *int64  `json:"maxRankCap"`
	RankThresholdName *string `json:"rankThresholdName"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
}

func NewExperienceModelMasterFromJson(data string) ExperienceModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceModelMasterFromDict(dict)
}

func NewExperienceModelMasterFromDict(data map[string]interface{}) ExperienceModelMaster {
	return ExperienceModelMaster{
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Name:              core.CastString(data["name"]),
		Description:       core.CastString(data["description"]),
		Metadata:          core.CastString(data["metadata"]),
		DefaultExperience: core.CastInt64(data["defaultExperience"]),
		DefaultRankCap:    core.CastInt64(data["defaultRankCap"]),
		MaxRankCap:        core.CastInt64(data["maxRankCap"]),
		RankThresholdName: core.CastString(data["rankThresholdName"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
	}
}

func (p ExperienceModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"experienceModelId": p.ExperienceModelId,
		"name":              p.Name,
		"description":       p.Description,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThresholdName": p.RankThresholdName,
		"createdAt":         p.CreatedAt,
		"updatedAt":         p.UpdatedAt,
	}
}

func (p ExperienceModelMaster) Pointer() *ExperienceModelMaster {
	return &p
}

func CastExperienceModelMasters(data []interface{}) []ExperienceModelMaster {
	v := make([]ExperienceModelMaster, 0)
	for _, d := range data {
		v = append(v, NewExperienceModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceModelMastersFromDict(data []ExperienceModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ExperienceModel struct {
	ExperienceModelId *string    `json:"experienceModelId"`
	Name              *string    `json:"name"`
	Metadata          *string    `json:"metadata"`
	DefaultExperience *int64     `json:"defaultExperience"`
	DefaultRankCap    *int64     `json:"defaultRankCap"`
	MaxRankCap        *int64     `json:"maxRankCap"`
	RankThreshold     *Threshold `json:"rankThreshold"`
}

func NewExperienceModelFromJson(data string) ExperienceModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExperienceModelFromDict(dict)
}

func NewExperienceModelFromDict(data map[string]interface{}) ExperienceModel {
	return ExperienceModel{
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		DefaultExperience: core.CastInt64(data["defaultExperience"]),
		DefaultRankCap:    core.CastInt64(data["defaultRankCap"]),
		MaxRankCap:        core.CastInt64(data["maxRankCap"]),
		RankThreshold:     NewThresholdFromDict(core.CastMap(data["rankThreshold"])).Pointer(),
	}
}

func (p ExperienceModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"experienceModelId": p.ExperienceModelId,
		"name":              p.Name,
		"metadata":          p.Metadata,
		"defaultExperience": p.DefaultExperience,
		"defaultRankCap":    p.DefaultRankCap,
		"maxRankCap":        p.MaxRankCap,
		"rankThreshold":     p.RankThreshold.ToDict(),
	}
}

func (p ExperienceModel) Pointer() *ExperienceModel {
	return &p
}

func CastExperienceModels(data []interface{}) []ExperienceModel {
	v := make([]ExperienceModel, 0)
	for _, d := range data {
		v = append(v, NewExperienceModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExperienceModelsFromDict(data []ExperienceModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ThresholdMaster struct {
	ThresholdId *string `json:"thresholdId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Metadata    *string `json:"metadata"`
	Values      []int64 `json:"values"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func NewThresholdMasterFromJson(data string) ThresholdMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewThresholdMasterFromDict(dict)
}

func NewThresholdMasterFromDict(data map[string]interface{}) ThresholdMaster {
	return ThresholdMaster{
		ThresholdId: core.CastString(data["thresholdId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Metadata:    core.CastString(data["metadata"]),
		Values:      core.CastInt64s(core.CastArray(data["values"])),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
	}
}

func (p ThresholdMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"thresholdId": p.ThresholdId,
		"name":        p.Name,
		"description": p.Description,
		"metadata":    p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p ThresholdMaster) Pointer() *ThresholdMaster {
	return &p
}

func CastThresholdMasters(data []interface{}) []ThresholdMaster {
	v := make([]ThresholdMaster, 0)
	for _, d := range data {
		v = append(v, NewThresholdMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastThresholdMastersFromDict(data []ThresholdMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Threshold struct {
	Metadata *string `json:"metadata"`
	Values   []int64 `json:"values"`
}

func NewThresholdFromJson(data string) Threshold {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewThresholdFromDict(dict)
}

func NewThresholdFromDict(data map[string]interface{}) Threshold {
	return Threshold{
		Metadata: core.CastString(data["metadata"]),
		Values:   core.CastInt64s(core.CastArray(data["values"])),
	}
}

func (p Threshold) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": p.Metadata,
		"values": core.CastInt64sFromDict(
			p.Values,
		),
	}
}

func (p Threshold) Pointer() *Threshold {
	return &p
}

func CastThresholds(data []interface{}) []Threshold {
	v := make([]Threshold, 0)
	for _, d := range data {
		v = append(v, NewThresholdFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastThresholdsFromDict(data []Threshold) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentExperienceMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentExperienceMasterFromJson(data string) CurrentExperienceMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentExperienceMasterFromDict(dict)
}

func NewCurrentExperienceMasterFromDict(data map[string]interface{}) CurrentExperienceMaster {
	return CurrentExperienceMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentExperienceMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentExperienceMaster) Pointer() *CurrentExperienceMaster {
	return &p
}

func CastCurrentExperienceMasters(data []interface{}) []CurrentExperienceMaster {
	v := make([]CurrentExperienceMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentExperienceMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentExperienceMastersFromDict(data []CurrentExperienceMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	StatusId        *string `json:"statusId"`
	ExperienceName  *string `json:"experienceName"`
	UserId          *string `json:"userId"`
	PropertyId      *string `json:"propertyId"`
	ExperienceValue *int64  `json:"experienceValue"`
	RankValue       *int64  `json:"rankValue"`
	RankCapValue    *int64  `json:"rankCapValue"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
}

func NewStatusFromJson(data string) Status {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		StatusId:        core.CastString(data["statusId"]),
		ExperienceName:  core.CastString(data["experienceName"]),
		UserId:          core.CastString(data["userId"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt64(data["experienceValue"]),
		RankValue:       core.CastInt64(data["rankValue"]),
		RankCapValue:    core.CastInt64(data["rankCapValue"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
	}
}

func (p Status) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"statusId":        p.StatusId,
		"experienceName":  p.ExperienceName,
		"userId":          p.UserId,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
		"rankValue":       p.RankValue,
		"rankCapValue":    p.RankCapValue,
		"createdAt":       p.CreatedAt,
		"updatedAt":       p.UpdatedAt,
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
