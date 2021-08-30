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

package stamina

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId           *string        `json:"namespaceId"`
	Name                  *string        `json:"name"`
	Description           *string        `json:"description"`
	OverflowTriggerScript *ScriptSetting `json:"overflowTriggerScript"`
	LogSetting            *LogSetting    `json:"logSetting"`
	CreatedAt             *int64         `json:"createdAt"`
	UpdatedAt             *int64         `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:           core.CastString(data["namespaceId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		OverflowTriggerScript: NewScriptSettingFromDict(core.CastMap(data["overflowTriggerScript"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":           p.NamespaceId,
		"name":                  p.Name,
		"description":           p.Description,
		"overflowTriggerScript": p.OverflowTriggerScript.ToDict(),
		"logSetting":            p.LogSetting.ToDict(),
		"createdAt":             p.CreatedAt,
		"updatedAt":             p.UpdatedAt,
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

type StaminaModelMaster struct {
	StaminaModelId           *string `json:"staminaModelId"`
	Name                     *string `json:"name"`
	Metadata                 *string `json:"metadata"`
	Description              *string `json:"description"`
	RecoverIntervalMinutes   *int32  `json:"recoverIntervalMinutes"`
	RecoverValue             *int32  `json:"recoverValue"`
	InitialCapacity          *int32  `json:"initialCapacity"`
	IsOverflow               *bool   `json:"isOverflow"`
	MaxCapacity              *int32  `json:"maxCapacity"`
	MaxStaminaTableName      *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName    *string `json:"recoverValueTableName"`
	CreatedAt                *int64  `json:"createdAt"`
	UpdatedAt                *int64  `json:"updatedAt"`
}

func NewStaminaModelMasterFromDict(data map[string]interface{}) StaminaModelMaster {
	return StaminaModelMaster{
		StaminaModelId:           core.CastString(data["staminaModelId"]),
		Name:                     core.CastString(data["name"]),
		Metadata:                 core.CastString(data["metadata"]),
		Description:              core.CastString(data["description"]),
		RecoverIntervalMinutes:   core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:             core.CastInt32(data["recoverValue"]),
		InitialCapacity:          core.CastInt32(data["initialCapacity"]),
		IsOverflow:               core.CastBool(data["isOverflow"]),
		MaxCapacity:              core.CastInt32(data["maxCapacity"]),
		MaxStaminaTableName:      core.CastString(data["maxStaminaTableName"]),
		RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
		RecoverValueTableName:    core.CastString(data["recoverValueTableName"]),
		CreatedAt:                core.CastInt64(data["createdAt"]),
		UpdatedAt:                core.CastInt64(data["updatedAt"]),
	}
}

func (p StaminaModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"staminaModelId":           p.StaminaModelId,
		"name":                     p.Name,
		"metadata":                 p.Metadata,
		"description":              p.Description,
		"recoverIntervalMinutes":   p.RecoverIntervalMinutes,
		"recoverValue":             p.RecoverValue,
		"initialCapacity":          p.InitialCapacity,
		"isOverflow":               p.IsOverflow,
		"maxCapacity":              p.MaxCapacity,
		"maxStaminaTableName":      p.MaxStaminaTableName,
		"recoverIntervalTableName": p.RecoverIntervalTableName,
		"recoverValueTableName":    p.RecoverValueTableName,
		"createdAt":                p.CreatedAt,
		"updatedAt":                p.UpdatedAt,
	}
}

func (p StaminaModelMaster) Pointer() *StaminaModelMaster {
	return &p
}

func CastStaminaModelMasters(data []interface{}) []StaminaModelMaster {
	v := make([]StaminaModelMaster, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelMastersFromDict(data []StaminaModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MaxStaminaTableMaster struct {
	MaxStaminaTableId *string `json:"maxStaminaTableId"`
	Name              *string `json:"name"`
	Metadata          *string `json:"metadata"`
	Description       *string `json:"description"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values            []int32 `json:"values"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
}

func NewMaxStaminaTableMasterFromDict(data map[string]interface{}) MaxStaminaTableMaster {
	return MaxStaminaTableMaster{
		MaxStaminaTableId: core.CastString(data["maxStaminaTableId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		Description:       core.CastString(data["description"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
	}
}

func (p MaxStaminaTableMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"maxStaminaTableId": p.MaxStaminaTableId,
		"name":              p.Name,
		"metadata":          p.Metadata,
		"description":       p.Description,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p MaxStaminaTableMaster) Pointer() *MaxStaminaTableMaster {
	return &p
}

func CastMaxStaminaTableMasters(data []interface{}) []MaxStaminaTableMaster {
	v := make([]MaxStaminaTableMaster, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTableMastersFromDict(data []MaxStaminaTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverIntervalTableMaster struct {
	RecoverIntervalTableId *string `json:"recoverIntervalTableId"`
	Name                   *string `json:"name"`
	Metadata               *string `json:"metadata"`
	Description            *string `json:"description"`
	ExperienceModelId      *string `json:"experienceModelId"`
	Values                 []int32 `json:"values"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
}

func NewRecoverIntervalTableMasterFromDict(data map[string]interface{}) RecoverIntervalTableMaster {
	return RecoverIntervalTableMaster{
		RecoverIntervalTableId: core.CastString(data["recoverIntervalTableId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		Values:                 core.CastInt32s(core.CastArray(data["values"])),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
	}
}

func (p RecoverIntervalTableMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"recoverIntervalTableId": p.RecoverIntervalTableId,
		"name":                   p.Name,
		"metadata":               p.Metadata,
		"description":            p.Description,
		"experienceModelId":      p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p RecoverIntervalTableMaster) Pointer() *RecoverIntervalTableMaster {
	return &p
}

func CastRecoverIntervalTableMasters(data []interface{}) []RecoverIntervalTableMaster {
	v := make([]RecoverIntervalTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTableMastersFromDict(data []RecoverIntervalTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverValueTableMaster struct {
	RecoverValueTableId *string `json:"recoverValueTableId"`
	Name                *string `json:"name"`
	Metadata            *string `json:"metadata"`
	Description         *string `json:"description"`
	ExperienceModelId   *string `json:"experienceModelId"`
	Values              []int32 `json:"values"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
}

func NewRecoverValueTableMasterFromDict(data map[string]interface{}) RecoverValueTableMaster {
	return RecoverValueTableMaster{
		RecoverValueTableId: core.CastString(data["recoverValueTableId"]),
		Name:                core.CastString(data["name"]),
		Metadata:            core.CastString(data["metadata"]),
		Description:         core.CastString(data["description"]),
		ExperienceModelId:   core.CastString(data["experienceModelId"]),
		Values:              core.CastInt32s(core.CastArray(data["values"])),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
	}
}

func (p RecoverValueTableMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"recoverValueTableId": p.RecoverValueTableId,
		"name":                p.Name,
		"metadata":            p.Metadata,
		"description":         p.Description,
		"experienceModelId":   p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p RecoverValueTableMaster) Pointer() *RecoverValueTableMaster {
	return &p
}

func CastRecoverValueTableMasters(data []interface{}) []RecoverValueTableMaster {
	v := make([]RecoverValueTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTableMastersFromDict(data []RecoverValueTableMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentStaminaMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentStaminaMasterFromDict(data map[string]interface{}) CurrentStaminaMaster {
	return CurrentStaminaMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentStaminaMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentStaminaMaster) Pointer() *CurrentStaminaMaster {
	return &p
}

func CastCurrentStaminaMasters(data []interface{}) []CurrentStaminaMaster {
	v := make([]CurrentStaminaMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentStaminaMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentStaminaMastersFromDict(data []CurrentStaminaMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StaminaModel struct {
	StaminaModelId         *string               `json:"staminaModelId"`
	Name                   *string               `json:"name"`
	Metadata               *string               `json:"metadata"`
	RecoverIntervalMinutes *int32                `json:"recoverIntervalMinutes"`
	RecoverValue           *int32                `json:"recoverValue"`
	InitialCapacity        *int32                `json:"initialCapacity"`
	IsOverflow             *bool                 `json:"isOverflow"`
	MaxCapacity            *int32                `json:"maxCapacity"`
	MaxStaminaTable        *MaxStaminaTable      `json:"maxStaminaTable"`
	RecoverIntervalTable   *RecoverIntervalTable `json:"recoverIntervalTable"`
	RecoverValueTable      *RecoverValueTable    `json:"recoverValueTable"`
}

func NewStaminaModelFromDict(data map[string]interface{}) StaminaModel {
	return StaminaModel{
		StaminaModelId:         core.CastString(data["staminaModelId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:           core.CastInt32(data["recoverValue"]),
		InitialCapacity:        core.CastInt32(data["initialCapacity"]),
		IsOverflow:             core.CastBool(data["isOverflow"]),
		MaxCapacity:            core.CastInt32(data["maxCapacity"]),
		MaxStaminaTable:        NewMaxStaminaTableFromDict(core.CastMap(data["maxStaminaTable"])).Pointer(),
		RecoverIntervalTable:   NewRecoverIntervalTableFromDict(core.CastMap(data["recoverIntervalTable"])).Pointer(),
		RecoverValueTable:      NewRecoverValueTableFromDict(core.CastMap(data["recoverValueTable"])).Pointer(),
	}
}

func (p StaminaModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"staminaModelId":         p.StaminaModelId,
		"name":                   p.Name,
		"metadata":               p.Metadata,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
		"recoverValue":           p.RecoverValue,
		"initialCapacity":        p.InitialCapacity,
		"isOverflow":             p.IsOverflow,
		"maxCapacity":            p.MaxCapacity,
		"maxStaminaTable":        p.MaxStaminaTable.ToDict(),
		"recoverIntervalTable":   p.RecoverIntervalTable.ToDict(),
		"recoverValueTable":      p.RecoverValueTable.ToDict(),
	}
}

func (p StaminaModel) Pointer() *StaminaModel {
	return &p
}

func CastStaminaModels(data []interface{}) []StaminaModel {
	v := make([]StaminaModel, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelsFromDict(data []StaminaModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MaxStaminaTable struct {
	MaxStaminaTableId *string `json:"maxStaminaTableId"`
	Name              *string `json:"name"`
	Metadata          *string `json:"metadata"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values            []int32 `json:"values"`
}

func NewMaxStaminaTableFromDict(data map[string]interface{}) MaxStaminaTable {
	return MaxStaminaTable{
		MaxStaminaTableId: core.CastString(data["maxStaminaTableId"]),
		Name:              core.CastString(data["name"]),
		Metadata:          core.CastString(data["metadata"]),
		ExperienceModelId: core.CastString(data["experienceModelId"]),
		Values:            core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p MaxStaminaTable) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"maxStaminaTableId": p.MaxStaminaTableId,
		"name":              p.Name,
		"metadata":          p.Metadata,
		"experienceModelId": p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p MaxStaminaTable) Pointer() *MaxStaminaTable {
	return &p
}

func CastMaxStaminaTables(data []interface{}) []MaxStaminaTable {
	v := make([]MaxStaminaTable, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTablesFromDict(data []MaxStaminaTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverIntervalTable struct {
	RecoverIntervalTableId *string `json:"recoverIntervalTableId"`
	Name                   *string `json:"name"`
	Metadata               *string `json:"metadata"`
	ExperienceModelId      *string `json:"experienceModelId"`
	Values                 []int32 `json:"values"`
}

func NewRecoverIntervalTableFromDict(data map[string]interface{}) RecoverIntervalTable {
	return RecoverIntervalTable{
		RecoverIntervalTableId: core.CastString(data["recoverIntervalTableId"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		ExperienceModelId:      core.CastString(data["experienceModelId"]),
		Values:                 core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p RecoverIntervalTable) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"recoverIntervalTableId": p.RecoverIntervalTableId,
		"name":                   p.Name,
		"metadata":               p.Metadata,
		"experienceModelId":      p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p RecoverIntervalTable) Pointer() *RecoverIntervalTable {
	return &p
}

func CastRecoverIntervalTables(data []interface{}) []RecoverIntervalTable {
	v := make([]RecoverIntervalTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTablesFromDict(data []RecoverIntervalTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RecoverValueTable struct {
	RecoverValueTableId *string `json:"recoverValueTableId"`
	Name                *string `json:"name"`
	Metadata            *string `json:"metadata"`
	ExperienceModelId   *string `json:"experienceModelId"`
	Values              []int32 `json:"values"`
}

func NewRecoverValueTableFromDict(data map[string]interface{}) RecoverValueTable {
	return RecoverValueTable{
		RecoverValueTableId: core.CastString(data["recoverValueTableId"]),
		Name:                core.CastString(data["name"]),
		Metadata:            core.CastString(data["metadata"]),
		ExperienceModelId:   core.CastString(data["experienceModelId"]),
		Values:              core.CastInt32s(core.CastArray(data["values"])),
	}
}

func (p RecoverValueTable) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"recoverValueTableId": p.RecoverValueTableId,
		"name":                p.Name,
		"metadata":            p.Metadata,
		"experienceModelId":   p.ExperienceModelId,
		"values": core.CastInt32sFromDict(
			p.Values,
		),
	}
}

func (p RecoverValueTable) Pointer() *RecoverValueTable {
	return &p
}

func CastRecoverValueTables(data []interface{}) []RecoverValueTable {
	v := make([]RecoverValueTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTablesFromDict(data []RecoverValueTable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Stamina struct {
	StaminaId              *string `json:"staminaId"`
	StaminaName            *string `json:"staminaName"`
	UserId                 *string `json:"userId"`
	Value                  *int32  `json:"value"`
	MaxValue               *int32  `json:"maxValue"`
	RecoverIntervalMinutes *int32  `json:"recoverIntervalMinutes"`
	RecoverValue           *int32  `json:"recoverValue"`
	OverflowValue          *int32  `json:"overflowValue"`
	NextRecoverAt          *int64  `json:"nextRecoverAt"`
	LastRecoveredAt        *int64  `json:"lastRecoveredAt"`
	CreatedAt              *int64  `json:"createdAt"`
	UpdatedAt              *int64  `json:"updatedAt"`
}

func NewStaminaFromDict(data map[string]interface{}) Stamina {
	return Stamina{
		StaminaId:              core.CastString(data["staminaId"]),
		StaminaName:            core.CastString(data["staminaName"]),
		UserId:                 core.CastString(data["userId"]),
		Value:                  core.CastInt32(data["value"]),
		MaxValue:               core.CastInt32(data["maxValue"]),
		RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
		RecoverValue:           core.CastInt32(data["recoverValue"]),
		OverflowValue:          core.CastInt32(data["overflowValue"]),
		NextRecoverAt:          core.CastInt64(data["nextRecoverAt"]),
		LastRecoveredAt:        core.CastInt64(data["lastRecoveredAt"]),
		CreatedAt:              core.CastInt64(data["createdAt"]),
		UpdatedAt:              core.CastInt64(data["updatedAt"]),
	}
}

func (p Stamina) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"staminaId":              p.StaminaId,
		"staminaName":            p.StaminaName,
		"userId":                 p.UserId,
		"value":                  p.Value,
		"maxValue":               p.MaxValue,
		"recoverIntervalMinutes": p.RecoverIntervalMinutes,
		"recoverValue":           p.RecoverValue,
		"overflowValue":          p.OverflowValue,
		"nextRecoverAt":          p.NextRecoverAt,
		"lastRecoveredAt":        p.LastRecoveredAt,
		"createdAt":              p.CreatedAt,
		"updatedAt":              p.UpdatedAt,
	}
}

func (p Stamina) Pointer() *Stamina {
	return &p
}

func CastStaminas(data []interface{}) []Stamina {
	v := make([]Stamina, 0)
	for _, d := range data {
		v = append(v, NewStaminaFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminasFromDict(data []Stamina) []interface{} {
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
