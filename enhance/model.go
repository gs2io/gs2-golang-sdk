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

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId         *string             `json:"namespaceId"`
	Name                *string             `json:"name"`
	Description         *string             `json:"description"`
	EnableDirectEnhance *bool               `json:"enableDirectEnhance"`
	TransactionSetting  *TransactionSetting `json:"transactionSetting"`
	EnhanceScript       *ScriptSetting      `json:"enhanceScript"`
	LogSetting          *LogSetting         `json:"logSetting"`
	CreatedAt           *int64              `json:"createdAt"`
	UpdatedAt           *int64              `json:"updatedAt"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId    *string `json:"keyId"`
	Revision *int64  `json:"revision"`
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
		EnableDirectEnhance: core.CastBool(data["enableDirectEnhance"]),
		TransactionSetting:  NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		EnhanceScript:       NewScriptSettingFromDict(core.CastMap(data["enhanceScript"])).Pointer(),
		LogSetting:          NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:    core.CastString(data["queueNamespaceId"]),
		KeyId:               core.CastString(data["keyId"]),
		Revision:            core.CastInt64(data["revision"]),
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
	var enableDirectEnhance *bool
	if p.EnableDirectEnhance != nil {
		enableDirectEnhance = p.EnableDirectEnhance
	}
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = p.TransactionSetting.ToDict()
	}
	var enhanceScript map[string]interface{}
	if p.EnhanceScript != nil {
		enhanceScript = p.EnhanceScript.ToDict()
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
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":         namespaceId,
		"name":                name,
		"description":         description,
		"enableDirectEnhance": enableDirectEnhance,
		"transactionSetting":  transactionSetting,
		"enhanceScript":       enhanceScript,
		"logSetting":          logSetting,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"queueNamespaceId":    queueNamespaceId,
		"keyId":               keyId,
		"revision":            revision,
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
	AcquireExperienceHierarchy []*string   `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
}

func NewRateModelFromJson(data string) RateModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRateModelFromDict(dict)
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

	var rateModelId *string
	if p.RateModelId != nil {
		rateModelId = p.RateModelId
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
	var targetInventoryModelId *string
	if p.TargetInventoryModelId != nil {
		targetInventoryModelId = p.TargetInventoryModelId
	}
	var acquireExperienceSuffix *string
	if p.AcquireExperienceSuffix != nil {
		acquireExperienceSuffix = p.AcquireExperienceSuffix
	}
	var materialInventoryModelId *string
	if p.MaterialInventoryModelId != nil {
		materialInventoryModelId = p.MaterialInventoryModelId
	}
	var acquireExperienceHierarchy []interface{}
	if p.AcquireExperienceHierarchy != nil {
		acquireExperienceHierarchy = core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var bonusRates []interface{}
	if p.BonusRates != nil {
		bonusRates = CastBonusRatesFromDict(
			p.BonusRates,
		)
	}
	return map[string]interface{}{
		"rateModelId":                rateModelId,
		"name":                       name,
		"description":                description,
		"metadata":                   metadata,
		"targetInventoryModelId":     targetInventoryModelId,
		"acquireExperienceSuffix":    acquireExperienceSuffix,
		"materialInventoryModelId":   materialInventoryModelId,
		"acquireExperienceHierarchy": acquireExperienceHierarchy,
		"experienceModelId":          experienceModelId,
		"bonusRates":                 bonusRates,
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
	AcquireExperienceHierarchy []*string   `json:"acquireExperienceHierarchy"`
	ExperienceModelId          *string     `json:"experienceModelId"`
	BonusRates                 []BonusRate `json:"bonusRates"`
	CreatedAt                  *int64      `json:"createdAt"`
	UpdatedAt                  *int64      `json:"updatedAt"`
	Revision                   *int64      `json:"revision"`
}

func NewRateModelMasterFromJson(data string) RateModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRateModelMasterFromDict(dict)
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
		Revision:                   core.CastInt64(data["revision"]),
	}
}

func (p RateModelMaster) ToDict() map[string]interface{} {

	var rateModelId *string
	if p.RateModelId != nil {
		rateModelId = p.RateModelId
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
	var targetInventoryModelId *string
	if p.TargetInventoryModelId != nil {
		targetInventoryModelId = p.TargetInventoryModelId
	}
	var acquireExperienceSuffix *string
	if p.AcquireExperienceSuffix != nil {
		acquireExperienceSuffix = p.AcquireExperienceSuffix
	}
	var materialInventoryModelId *string
	if p.MaterialInventoryModelId != nil {
		materialInventoryModelId = p.MaterialInventoryModelId
	}
	var acquireExperienceHierarchy []interface{}
	if p.AcquireExperienceHierarchy != nil {
		acquireExperienceHierarchy = core.CastStringsFromDict(
			p.AcquireExperienceHierarchy,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var bonusRates []interface{}
	if p.BonusRates != nil {
		bonusRates = CastBonusRatesFromDict(
			p.BonusRates,
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
		"rateModelId":                rateModelId,
		"name":                       name,
		"description":                description,
		"metadata":                   metadata,
		"targetInventoryModelId":     targetInventoryModelId,
		"acquireExperienceSuffix":    acquireExperienceSuffix,
		"materialInventoryModelId":   materialInventoryModelId,
		"acquireExperienceHierarchy": acquireExperienceHierarchy,
		"experienceModelId":          experienceModelId,
		"bonusRates":                 bonusRates,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
		"revision":                   revision,
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

type UnleashRateModel struct {
	UnleashRateModelId     *string                 `json:"unleashRateModelId"`
	Name                   *string                 `json:"name"`
	Description            *string                 `json:"description"`
	Metadata               *string                 `json:"metadata"`
	TargetInventoryModelId *string                 `json:"targetInventoryModelId"`
	GradeModelId           *string                 `json:"gradeModelId"`
	GradeEntries           []UnleashRateEntryModel `json:"gradeEntries"`
}

func NewUnleashRateModelFromJson(data string) UnleashRateModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashRateModelFromDict(dict)
}

func NewUnleashRateModelFromDict(data map[string]interface{}) UnleashRateModel {
	return UnleashRateModel{
		UnleashRateModelId:     core.CastString(data["unleashRateModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		TargetInventoryModelId: core.CastString(data["targetInventoryModelId"]),
		GradeModelId:           core.CastString(data["gradeModelId"]),
		GradeEntries:           CastUnleashRateEntryModels(core.CastArray(data["gradeEntries"])),
	}
}

func (p UnleashRateModel) ToDict() map[string]interface{} {

	var unleashRateModelId *string
	if p.UnleashRateModelId != nil {
		unleashRateModelId = p.UnleashRateModelId
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
	var targetInventoryModelId *string
	if p.TargetInventoryModelId != nil {
		targetInventoryModelId = p.TargetInventoryModelId
	}
	var gradeModelId *string
	if p.GradeModelId != nil {
		gradeModelId = p.GradeModelId
	}
	var gradeEntries []interface{}
	if p.GradeEntries != nil {
		gradeEntries = CastUnleashRateEntryModelsFromDict(
			p.GradeEntries,
		)
	}
	return map[string]interface{}{
		"unleashRateModelId":     unleashRateModelId,
		"name":                   name,
		"description":            description,
		"metadata":               metadata,
		"targetInventoryModelId": targetInventoryModelId,
		"gradeModelId":           gradeModelId,
		"gradeEntries":           gradeEntries,
	}
}

func (p UnleashRateModel) Pointer() *UnleashRateModel {
	return &p
}

func CastUnleashRateModels(data []interface{}) []UnleashRateModel {
	v := make([]UnleashRateModel, 0)
	for _, d := range data {
		v = append(v, NewUnleashRateModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUnleashRateModelsFromDict(data []UnleashRateModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type UnleashRateModelMaster struct {
	UnleashRateModelId     *string                 `json:"unleashRateModelId"`
	Name                   *string                 `json:"name"`
	Description            *string                 `json:"description"`
	Metadata               *string                 `json:"metadata"`
	TargetInventoryModelId *string                 `json:"targetInventoryModelId"`
	GradeModelId           *string                 `json:"gradeModelId"`
	GradeEntries           []UnleashRateEntryModel `json:"gradeEntries"`
	CreatedAt              *int64                  `json:"createdAt"`
	UpdatedAt              *int64                  `json:"updatedAt"`
	Revision               *int64                  `json:"revision"`
}

func NewUnleashRateModelMasterFromJson(data string) UnleashRateModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashRateModelMasterFromDict(dict)
}

func NewUnleashRateModelMasterFromDict(data map[string]interface{}) UnleashRateModelMaster {
	return UnleashRateModelMaster{
		UnleashRateModelId:     core.CastString(data["unleashRateModelId"]),
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		Metadata:               core.CastString(data["metadata"]),
		TargetInventoryModelId: core.CastString(data["targetInventoryModelId"]),
		GradeModelId:           core.CastString(data["gradeModelId"]),
		GradeEntries:           CastUnleashRateEntryModels(core.CastArray(data["gradeEntries"])),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
		Revision:               core.CastInt64(data["revision"]),
	}
}

func (p UnleashRateModelMaster) ToDict() map[string]interface{} {

	var unleashRateModelId *string
	if p.UnleashRateModelId != nil {
		unleashRateModelId = p.UnleashRateModelId
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
	var targetInventoryModelId *string
	if p.TargetInventoryModelId != nil {
		targetInventoryModelId = p.TargetInventoryModelId
	}
	var gradeModelId *string
	if p.GradeModelId != nil {
		gradeModelId = p.GradeModelId
	}
	var gradeEntries []interface{}
	if p.GradeEntries != nil {
		gradeEntries = CastUnleashRateEntryModelsFromDict(
			p.GradeEntries,
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
		"unleashRateModelId":     unleashRateModelId,
		"name":                   name,
		"description":            description,
		"metadata":               metadata,
		"targetInventoryModelId": targetInventoryModelId,
		"gradeModelId":           gradeModelId,
		"gradeEntries":           gradeEntries,
		"createdAt":              createdAt,
		"updatedAt":              updatedAt,
		"revision":               revision,
	}
}

func (p UnleashRateModelMaster) Pointer() *UnleashRateModelMaster {
	return &p
}

func CastUnleashRateModelMasters(data []interface{}) []UnleashRateModelMaster {
	v := make([]UnleashRateModelMaster, 0)
	for _, d := range data {
		v = append(v, NewUnleashRateModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUnleashRateModelMastersFromDict(data []UnleashRateModelMaster) []interface{} {
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
	Name            *string  `json:"name"`
	PropertyId      *string  `json:"propertyId"`
	ExperienceValue *int64   `json:"experienceValue"`
	Rate            *float32 `json:"rate"`
	CreatedAt       *int64   `json:"createdAt"`
	UpdatedAt       *int64   `json:"updatedAt"`
	Revision        *int64   `json:"revision"`
}

func NewProgressFromJson(data string) Progress {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewProgressFromDict(dict)
}

func NewProgressFromDict(data map[string]interface{}) Progress {
	return Progress{
		ProgressId:      core.CastString(data["progressId"]),
		UserId:          core.CastString(data["userId"]),
		RateName:        core.CastString(data["rateName"]),
		Name:            core.CastString(data["name"]),
		PropertyId:      core.CastString(data["propertyId"]),
		ExperienceValue: core.CastInt64(data["experienceValue"]),
		Rate:            core.CastFloat32(data["rate"]),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
		Revision:        core.CastInt64(data["revision"]),
	}
}

func (p Progress) ToDict() map[string]interface{} {

	var progressId *string
	if p.ProgressId != nil {
		progressId = p.ProgressId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var rateName *string
	if p.RateName != nil {
		rateName = p.RateName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var experienceValue *int64
	if p.ExperienceValue != nil {
		experienceValue = p.ExperienceValue
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
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
		"progressId":      progressId,
		"userId":          userId,
		"rateName":        rateName,
		"name":            name,
		"propertyId":      propertyId,
		"experienceValue": experienceValue,
		"rate":            rate,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
		"revision":        revision,
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

func NewCurrentRateMasterFromJson(data string) CurrentRateMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentRateMasterFromDict(dict)
}

func NewCurrentRateMasterFromDict(data map[string]interface{}) CurrentRateMaster {
	return CurrentRateMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentRateMaster) ToDict() map[string]interface{} {

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

type BonusRate struct {
	Rate   *float32 `json:"rate"`
	Weight *int32   `json:"weight"`
}

func NewBonusRateFromJson(data string) BonusRate {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBonusRateFromDict(dict)
}

func NewBonusRateFromDict(data map[string]interface{}) BonusRate {
	return BonusRate{
		Rate:   core.CastFloat32(data["rate"]),
		Weight: core.CastInt32(data["weight"]),
	}
}

func (p BonusRate) ToDict() map[string]interface{} {

	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"rate":   rate,
		"weight": weight,
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

func NewMaterialFromJson(data string) Material {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMaterialFromDict(dict)
}

func NewMaterialFromDict(data map[string]interface{}) Material {
	return Material{
		MaterialItemSetId: core.CastString(data["materialItemSetId"]),
		Count:             core.CastInt32(data["count"]),
	}
}

func (p Material) ToDict() map[string]interface{} {

	var materialItemSetId *string
	if p.MaterialItemSetId != nil {
		materialItemSetId = p.MaterialItemSetId
	}
	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"materialItemSetId": materialItemSetId,
		"count":             count,
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

type UnleashRateEntryModel struct {
	GradeValue *int64 `json:"gradeValue"`
	NeedCount  *int32 `json:"needCount"`
}

func NewUnleashRateEntryModelFromJson(data string) UnleashRateEntryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashRateEntryModelFromDict(dict)
}

func NewUnleashRateEntryModelFromDict(data map[string]interface{}) UnleashRateEntryModel {
	return UnleashRateEntryModel{
		GradeValue: core.CastInt64(data["gradeValue"]),
		NeedCount:  core.CastInt32(data["needCount"]),
	}
}

func (p UnleashRateEntryModel) ToDict() map[string]interface{} {

	var gradeValue *int64
	if p.GradeValue != nil {
		gradeValue = p.GradeValue
	}
	var needCount *int32
	if p.NeedCount != nil {
		needCount = p.NeedCount
	}
	return map[string]interface{}{
		"gradeValue": gradeValue,
		"needCount":  needCount,
	}
}

func (p UnleashRateEntryModel) Pointer() *UnleashRateEntryModel {
	return &p
}

func CastUnleashRateEntryModels(data []interface{}) []UnleashRateEntryModel {
	v := make([]UnleashRateEntryModel, 0)
	for _, d := range data {
		v = append(v, NewUnleashRateEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUnleashRateEntryModelsFromDict(data []UnleashRateEntryModel) []interface{} {
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
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
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
