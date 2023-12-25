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

package grade

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	ChangeGradeScript  *ScriptSetting      `json:"changeGradeScript"`
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
		ChangeGradeScript:  NewScriptSettingFromDict(core.CastMap(data["changeGradeScript"])).Pointer(),
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
	var changeGradeScript map[string]interface{}
	if p.ChangeGradeScript != nil {
		changeGradeScript = p.ChangeGradeScript.ToDict()
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
		"changeGradeScript":  changeGradeScript,
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

type GradeModelMaster struct {
	GradeModelId       *string             `json:"gradeModelId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
}

func NewGradeModelMasterFromJson(data string) GradeModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGradeModelMasterFromDict(dict)
}

func NewGradeModelMasterFromDict(data map[string]interface{}) GradeModelMaster {
	return GradeModelMaster{
		GradeModelId:       core.CastString(data["gradeModelId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultGrades:      CastDefaultGradeModels(core.CastArray(data["defaultGrades"])),
		ExperienceModelId:  core.CastString(data["experienceModelId"]),
		GradeEntries:       CastGradeEntryModels(core.CastArray(data["gradeEntries"])),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p GradeModelMaster) ToDict() map[string]interface{} {

	var gradeModelId *string
	if p.GradeModelId != nil {
		gradeModelId = p.GradeModelId
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
	var defaultGrades []interface{}
	if p.DefaultGrades != nil {
		defaultGrades = CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var gradeEntries []interface{}
	if p.GradeEntries != nil {
		gradeEntries = CastGradeEntryModelsFromDict(
			p.GradeEntries,
		)
	}
	var acquireActionRates []interface{}
	if p.AcquireActionRates != nil {
		acquireActionRates = CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
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
		"gradeModelId":       gradeModelId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"defaultGrades":      defaultGrades,
		"experienceModelId":  experienceModelId,
		"gradeEntries":       gradeEntries,
		"acquireActionRates": acquireActionRates,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p GradeModelMaster) Pointer() *GradeModelMaster {
	return &p
}

func CastGradeModelMasters(data []interface{}) []GradeModelMaster {
	v := make([]GradeModelMaster, 0)
	for _, d := range data {
		v = append(v, NewGradeModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGradeModelMastersFromDict(data []GradeModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GradeModel struct {
	GradeModelId       *string             `json:"gradeModelId"`
	Name               *string             `json:"name"`
	Metadata           *string             `json:"metadata"`
	DefaultGrades      []DefaultGradeModel `json:"defaultGrades"`
	ExperienceModelId  *string             `json:"experienceModelId"`
	GradeEntries       []GradeEntryModel   `json:"gradeEntries"`
	AcquireActionRates []AcquireActionRate `json:"acquireActionRates"`
}

func NewGradeModelFromJson(data string) GradeModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGradeModelFromDict(dict)
}

func NewGradeModelFromDict(data map[string]interface{}) GradeModel {
	return GradeModel{
		GradeModelId:       core.CastString(data["gradeModelId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		DefaultGrades:      CastDefaultGradeModels(core.CastArray(data["defaultGrades"])),
		ExperienceModelId:  core.CastString(data["experienceModelId"]),
		GradeEntries:       CastGradeEntryModels(core.CastArray(data["gradeEntries"])),
		AcquireActionRates: CastAcquireActionRates(core.CastArray(data["acquireActionRates"])),
	}
}

func (p GradeModel) ToDict() map[string]interface{} {

	var gradeModelId *string
	if p.GradeModelId != nil {
		gradeModelId = p.GradeModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var defaultGrades []interface{}
	if p.DefaultGrades != nil {
		defaultGrades = CastDefaultGradeModelsFromDict(
			p.DefaultGrades,
		)
	}
	var experienceModelId *string
	if p.ExperienceModelId != nil {
		experienceModelId = p.ExperienceModelId
	}
	var gradeEntries []interface{}
	if p.GradeEntries != nil {
		gradeEntries = CastGradeEntryModelsFromDict(
			p.GradeEntries,
		)
	}
	var acquireActionRates []interface{}
	if p.AcquireActionRates != nil {
		acquireActionRates = CastAcquireActionRatesFromDict(
			p.AcquireActionRates,
		)
	}
	return map[string]interface{}{
		"gradeModelId":       gradeModelId,
		"name":               name,
		"metadata":           metadata,
		"defaultGrades":      defaultGrades,
		"experienceModelId":  experienceModelId,
		"gradeEntries":       gradeEntries,
		"acquireActionRates": acquireActionRates,
	}
}

func (p GradeModel) Pointer() *GradeModel {
	return &p
}

func CastGradeModels(data []interface{}) []GradeModel {
	v := make([]GradeModel, 0)
	for _, d := range data {
		v = append(v, NewGradeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGradeModelsFromDict(data []GradeModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	StatusId   *string `json:"statusId"`
	GradeName  *string `json:"gradeName"`
	UserId     *string `json:"userId"`
	PropertyId *string `json:"propertyId"`
	GradeValue *int64  `json:"gradeValue"`
	CreatedAt  *int64  `json:"createdAt"`
	UpdatedAt  *int64  `json:"updatedAt"`
	Revision   *int64  `json:"revision"`
}

func NewStatusFromJson(data string) Status {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		StatusId:   core.CastString(data["statusId"]),
		GradeName:  core.CastString(data["gradeName"]),
		UserId:     core.CastString(data["userId"]),
		PropertyId: core.CastString(data["propertyId"]),
		GradeValue: core.CastInt64(data["gradeValue"]),
		CreatedAt:  core.CastInt64(data["createdAt"]),
		UpdatedAt:  core.CastInt64(data["updatedAt"]),
		Revision:   core.CastInt64(data["revision"]),
	}
}

func (p Status) ToDict() map[string]interface{} {

	var statusId *string
	if p.StatusId != nil {
		statusId = p.StatusId
	}
	var gradeName *string
	if p.GradeName != nil {
		gradeName = p.GradeName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var gradeValue *int64
	if p.GradeValue != nil {
		gradeValue = p.GradeValue
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
		"statusId":   statusId,
		"gradeName":  gradeName,
		"userId":     userId,
		"propertyId": propertyId,
		"gradeValue": gradeValue,
		"createdAt":  createdAt,
		"updatedAt":  updatedAt,
		"revision":   revision,
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

type DefaultGradeModel struct {
	PropertyIdRegex   *string `json:"propertyIdRegex"`
	DefaultGradeValue *int64  `json:"defaultGradeValue"`
}

func NewDefaultGradeModelFromJson(data string) DefaultGradeModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDefaultGradeModelFromDict(dict)
}

func NewDefaultGradeModelFromDict(data map[string]interface{}) DefaultGradeModel {
	return DefaultGradeModel{
		PropertyIdRegex:   core.CastString(data["propertyIdRegex"]),
		DefaultGradeValue: core.CastInt64(data["defaultGradeValue"]),
	}
}

func (p DefaultGradeModel) ToDict() map[string]interface{} {

	var propertyIdRegex *string
	if p.PropertyIdRegex != nil {
		propertyIdRegex = p.PropertyIdRegex
	}
	var defaultGradeValue *int64
	if p.DefaultGradeValue != nil {
		defaultGradeValue = p.DefaultGradeValue
	}
	return map[string]interface{}{
		"propertyIdRegex":   propertyIdRegex,
		"defaultGradeValue": defaultGradeValue,
	}
}

func (p DefaultGradeModel) Pointer() *DefaultGradeModel {
	return &p
}

func CastDefaultGradeModels(data []interface{}) []DefaultGradeModel {
	v := make([]DefaultGradeModel, 0)
	for _, d := range data {
		v = append(v, NewDefaultGradeModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDefaultGradeModelsFromDict(data []DefaultGradeModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GradeEntryModel struct {
	Metadata               *string `json:"metadata"`
	RankCapValue           *int64  `json:"rankCapValue"`
	PropertyIdRegex        *string `json:"propertyIdRegex"`
	GradeUpPropertyIdRegex *string `json:"gradeUpPropertyIdRegex"`
}

func NewGradeEntryModelFromJson(data string) GradeEntryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGradeEntryModelFromDict(dict)
}

func NewGradeEntryModelFromDict(data map[string]interface{}) GradeEntryModel {
	return GradeEntryModel{
		Metadata:               core.CastString(data["metadata"]),
		RankCapValue:           core.CastInt64(data["rankCapValue"]),
		PropertyIdRegex:        core.CastString(data["propertyIdRegex"]),
		GradeUpPropertyIdRegex: core.CastString(data["gradeUpPropertyIdRegex"]),
	}
}

func (p GradeEntryModel) ToDict() map[string]interface{} {

	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var rankCapValue *int64
	if p.RankCapValue != nil {
		rankCapValue = p.RankCapValue
	}
	var propertyIdRegex *string
	if p.PropertyIdRegex != nil {
		propertyIdRegex = p.PropertyIdRegex
	}
	var gradeUpPropertyIdRegex *string
	if p.GradeUpPropertyIdRegex != nil {
		gradeUpPropertyIdRegex = p.GradeUpPropertyIdRegex
	}
	return map[string]interface{}{
		"metadata":               metadata,
		"rankCapValue":           rankCapValue,
		"propertyIdRegex":        propertyIdRegex,
		"gradeUpPropertyIdRegex": gradeUpPropertyIdRegex,
	}
}

func (p GradeEntryModel) Pointer() *GradeEntryModel {
	return &p
}

func CastGradeEntryModels(data []interface{}) []GradeEntryModel {
	v := make([]GradeEntryModel, 0)
	for _, d := range data {
		v = append(v, NewGradeEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGradeEntryModelsFromDict(data []GradeEntryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireActionRate struct {
	Name     *string    `json:"name"`
	Mode     *string    `json:"mode"`
	Rates    []*float64 `json:"rates"`
	BigRates []*string  `json:"bigRates"`
}

func NewAcquireActionRateFromJson(data string) AcquireActionRate {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionRateFromDict(dict)
}

func NewAcquireActionRateFromDict(data map[string]interface{}) AcquireActionRate {
	return AcquireActionRate{
		Name:     core.CastString(data["name"]),
		Mode:     core.CastString(data["mode"]),
		Rates:    core.CastFloat64s(core.CastArray(data["rates"])),
		BigRates: core.CastStrings(core.CastArray(data["bigRates"])),
	}
}

func (p AcquireActionRate) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var mode *string
	if p.Mode != nil {
		mode = p.Mode
	}
	var rates []interface{}
	if p.Rates != nil {
		rates = core.CastFloat64sFromDict(
			p.Rates,
		)
	}
	var bigRates []interface{}
	if p.BigRates != nil {
		bigRates = core.CastStringsFromDict(
			p.BigRates,
		)
	}
	return map[string]interface{}{
		"name":     name,
		"mode":     mode,
		"rates":    rates,
		"bigRates": bigRates,
	}
}

func (p AcquireActionRate) Pointer() *AcquireActionRate {
	return &p
}

func CastAcquireActionRates(data []interface{}) []AcquireActionRate {
	v := make([]AcquireActionRate, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionRateFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionRatesFromDict(data []AcquireActionRate) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentGradeMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentGradeMasterFromJson(data string) CurrentGradeMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentGradeMasterFromDict(dict)
}

func NewCurrentGradeMasterFromDict(data map[string]interface{}) CurrentGradeMaster {
	return CurrentGradeMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentGradeMaster) ToDict() map[string]interface{} {

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

func (p CurrentGradeMaster) Pointer() *CurrentGradeMaster {
	return &p
}

func CastCurrentGradeMasters(data []interface{}) []CurrentGradeMaster {
	v := make([]CurrentGradeMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentGradeMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentGradeMastersFromDict(data []CurrentGradeMaster) []interface{} {
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
