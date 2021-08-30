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

package enhance

import "core"

type Namespace struct {
	NamespaceId         *string     `json:"namespaceId"`
	Name                *string     `json:"name"`
	Description         *string     `json:"description"`
	EnableDirectEnhance *bool       `json:"enableDirectEnhance"`
	QueueNamespaceId    *string     `json:"queueNamespaceId"`
	KeyId               *string     `json:"keyId"`
	LogSetting          *LogSetting `json:"logSetting"`
	CreatedAt           *int64      `json:"createdAt"`
	UpdatedAt           *int64      `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:         core.CastString(data["namespaceId"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		EnableDirectEnhance: core.CastBool(data["enableDirectEnhance"]),
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
		"enableDirectEnhance": p.EnableDirectEnhance,
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

type RateModel struct {
	RateModelId                *string     `json:"rateModelId"`
	Name                       *string     `json:"name"`
	Description                *string     `json:"description"`
	Metadata                   *string     `json:"metadata"`
	TargetInventoryModelId     *string     `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string     `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string     `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy []string    `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
}

func NewRateModelFromDict(data map[string]interface{}) RateModel {
	return RateModel{
		RateModelId:                core.CastString(data["rateModelId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetInventoryModelId:     core.CastString(data["targetInventoryModelId"]),
		AcquireExperienceSuffix:    core.CastString(data["acquireExperienceSuffix"]),
		MaterialInventoryModelId:   core.CastString(data["materialInventoryModelId"]),
		AcquireExperienceHierarchy: core.CastStrings(core.CastArray(data["acquireExperienceHierarchy"])),
		ExperienceModelId:          core.CastString(data["experienceModelId"]),
		BonusRates:                 CastBonusRates(core.CastArray(data["bonusRates"])),
	}
}

func (p RateModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rateModelId":              p.RateModelId,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"targetInventoryModelId":   p.TargetInventoryModelId,
		"acquireExperienceSuffix":  p.AcquireExperienceSuffix,
		"materialInventoryModelId": p.MaterialInventoryModelId,
		"acquireExperienceHierarchy": core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		),
		"experienceModelId": p.ExperienceModelId,
		"bonusRates": CastBonusRatesFromDict(
			p.BonusRates,
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
	RateModelId                *string     `json:"rateModelId"`
	Name                       *string     `json:"name"`
	Description                *string     `json:"description"`
	Metadata                   *string     `json:"metadata"`
	TargetInventoryModelId     *string     `json:"targetInventoryModelId"`
	AcquireExperienceSuffix    *string     `json:"acquireExperienceSuffix"`
	MaterialInventoryModelId   *string     `json:"materialInventoryModelId"`
	AcquireExperienceHierarchy []string    `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
	CreatedAt                  *int64      `json:"createdAt"`
	UpdatedAt                  *int64      `json:"updatedAt"`
}

func NewRateModelMasterFromDict(data map[string]interface{}) RateModelMaster {
	return RateModelMaster{
		RateModelId:                core.CastString(data["rateModelId"]),
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		Metadata:                   core.CastString(data["metadata"]),
		TargetInventoryModelId:     core.CastString(data["targetInventoryModelId"]),
		AcquireExperienceSuffix:    core.CastString(data["acquireExperienceSuffix"]),
		MaterialInventoryModelId:   core.CastString(data["materialInventoryModelId"]),
		AcquireExperienceHierarchy: core.CastStrings(core.CastArray(data["acquireExperienceHierarchy"])),
		ExperienceModelId:          core.CastString(data["experienceModelId"]),
		BonusRates:                 CastBonusRates(core.CastArray(data["bonusRates"])),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
	}
}

func (p RateModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rateModelId":              p.RateModelId,
		"name":                     p.Name,
		"description":              p.Description,
		"metadata":                 p.Metadata,
		"targetInventoryModelId":   p.TargetInventoryModelId,
		"acquireExperienceSuffix":  p.AcquireExperienceSuffix,
		"materialInventoryModelId": p.MaterialInventoryModelId,
		"acquireExperienceHierarchy": core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		),
		"experienceModelId": p.ExperienceModelId,
		"bonusRates": CastBonusRatesFromDict(
			p.BonusRates,
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

type Progress struct {
	ProgressId      *string  `json:"progressId"`
	UserId          *string  `json:"userId"`
	RateName        *string  `json:"rateName"`
	PropertyId      *string  `json:"propertyId"`
	ExperienceValue *int32   `json:"experienceValue"`
	Rate            *float32 `json:"rate"`
	CreatedAt       *int64   `json:"createdAt"`
	UpdatedAt       *int64   `json:"updatedAt"`
}

func NewProgressFromDict(data map[string]interface{}) Progress {
	return Progress{
		ProgressId:      core.CastString(data["progressId"]),
		UserId:          core.CastString(data["userId"]),
		RateName:        core.CastString(data["rateName"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt32(data["experienceValue"]),
		Rate:            core.CastFloat32(data["rate"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
	}
}

func (p Progress) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"progressId":      p.ProgressId,
		"userId":          p.UserId,
		"rateName":        p.RateName,
		"propertyId":      p.PropertyId,
		"experienceValue": p.ExperienceValue,
		"rate":            p.Rate,
		"createdAt":       p.CreatedAt,
		"updatedAt":       p.UpdatedAt,
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

type BonusRate struct {
	Rate   *float32 `json:"rate"`
	Weight *int32   `json:"weight"`
}

func NewBonusRateFromDict(data map[string]interface{}) BonusRate {
	return BonusRate{
		Rate:   core.CastFloat32(data["rate"]),
		Weight: core.CastInt32(data["weight"]),
	}
}

func (p BonusRate) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"rate":   p.Rate,
		"weight": p.Weight,
	}
}

func (p BonusRate) Pointer() *BonusRate {
	return &p
}

func CastBonusRates(data []interface{}) []BonusRate {
	v := make([]BonusRate, 0)
	for _, d := range data {
		v = append(v, NewBonusRateFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBonusRatesFromDict(data []BonusRate) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Material struct {
	MaterialItemSetId *string `json:"materialItemSetId"`
	Count             *int32  `json:"count"`
}

func NewMaterialFromDict(data map[string]interface{}) Material {
	return Material{
		MaterialItemSetId: core.CastString(data["materialItemSetId"]),
		Count:             core.CastInt32(data["count"]),
	}
}

func (p Material) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"materialItemSetId": p.MaterialItemSetId,
		"count":             p.Count,
	}
}

func (p Material) Pointer() *Material {
	return &p
}

func CastMaterials(data []interface{}) []Material {
	v := make([]Material, 0)
	for _, d := range data {
		v = append(v, NewMaterialFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaterialsFromDict(data []Material) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
