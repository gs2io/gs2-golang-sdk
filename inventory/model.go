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

package inventory

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId    *string        `json:"namespaceId"`
	Name           *string        `json:"name"`
	Description    *string        `json:"description"`
	AcquireScript  *ScriptSetting `json:"acquireScript"`
	OverflowScript *ScriptSetting `json:"overflowScript"`
	ConsumeScript  *ScriptSetting `json:"consumeScript"`
	LogSetting     *LogSetting    `json:"logSetting"`
	CreatedAt      *int64         `json:"createdAt"`
	UpdatedAt      *int64         `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:    core.CastString(data["namespaceId"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		AcquireScript:  NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript: NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:  NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		LogSetting:     NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":    p.NamespaceId,
		"name":           p.Name,
		"description":    p.Description,
		"acquireScript":  p.AcquireScript.ToDict(),
		"overflowScript": p.OverflowScript.ToDict(),
		"consumeScript":  p.ConsumeScript.ToDict(),
		"logSetting":     p.LogSetting.ToDict(),
		"createdAt":      p.CreatedAt,
		"updatedAt":      p.UpdatedAt,
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

type InventoryModelMaster struct {
	InventoryModelId      *string `json:"inventoryModelId"`
	Name                  *string `json:"name"`
	Metadata              *string `json:"metadata"`
	Description           *string `json:"description"`
	InitialCapacity       *int32  `json:"initialCapacity"`
	MaxCapacity           *int32  `json:"maxCapacity"`
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
	CreatedAt             *int64  `json:"createdAt"`
	UpdatedAt             *int64  `json:"updatedAt"`
}

func NewInventoryModelMasterFromJson(data string) InventoryModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryModelMasterFromDict(dict)
}

func NewInventoryModelMasterFromDict(data map[string]interface{}) InventoryModelMaster {
	return InventoryModelMaster{
		InventoryModelId:      core.CastString(data["inventoryModelId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		Description:           core.CastString(data["description"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
	}
}

func (p InventoryModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"inventoryModelId":      p.InventoryModelId,
		"name":                  p.Name,
		"metadata":              p.Metadata,
		"description":           p.Description,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
		"createdAt":             p.CreatedAt,
		"updatedAt":             p.UpdatedAt,
	}
}

func (p InventoryModelMaster) Pointer() *InventoryModelMaster {
	return &p
}

func CastInventoryModelMasters(data []interface{}) []InventoryModelMaster {
	v := make([]InventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryModelMastersFromDict(data []InventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InventoryModel struct {
	InventoryModelId      *string     `json:"inventoryModelId"`
	Name                  *string     `json:"name"`
	Metadata              *string     `json:"metadata"`
	InitialCapacity       *int32      `json:"initialCapacity"`
	MaxCapacity           *int32      `json:"maxCapacity"`
	ProtectReferencedItem *bool       `json:"protectReferencedItem"`
	ItemModels            []ItemModel `json:"itemModels"`
}

func NewInventoryModelFromJson(data string) InventoryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryModelFromDict(dict)
}

func NewInventoryModelFromDict(data map[string]interface{}) InventoryModel {
	return InventoryModel{
		InventoryModelId:      core.CastString(data["inventoryModelId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		InitialCapacity:       core.CastInt32(data["initialCapacity"]),
		MaxCapacity:           core.CastInt32(data["maxCapacity"]),
		ProtectReferencedItem: core.CastBool(data["protectReferencedItem"]),
		ItemModels:            CastItemModels(core.CastArray(data["itemModels"])),
	}
}

func (p InventoryModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"inventoryModelId":      p.InventoryModelId,
		"name":                  p.Name,
		"metadata":              p.Metadata,
		"initialCapacity":       p.InitialCapacity,
		"maxCapacity":           p.MaxCapacity,
		"protectReferencedItem": p.ProtectReferencedItem,
		"itemModels": CastItemModelsFromDict(
			p.ItemModels,
		),
	}
}

func (p InventoryModel) Pointer() *InventoryModel {
	return &p
}

func CastInventoryModels(data []interface{}) []InventoryModel {
	v := make([]InventoryModel, 0)
	for _, d := range data {
		v = append(v, NewInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoryModelsFromDict(data []InventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemModelMaster struct {
	ItemModelId         *string `json:"itemModelId"`
	InventoryName       *string `json:"inventoryName"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
}

func NewItemModelMasterFromJson(data string) ItemModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewItemModelMasterFromDict(dict)
}

func NewItemModelMasterFromDict(data map[string]interface{}) ItemModelMaster {
	return ItemModelMaster{
		ItemModelId:         core.CastString(data["itemModelId"]),
		InventoryName:       core.CastString(data["inventoryName"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
	}
}

func (p ItemModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"itemModelId":         p.ItemModelId,
		"inventoryName":       p.InventoryName,
		"name":                p.Name,
		"description":         p.Description,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
		"createdAt":           p.CreatedAt,
		"updatedAt":           p.UpdatedAt,
	}
}

func (p ItemModelMaster) Pointer() *ItemModelMaster {
	return &p
}

func CastItemModelMasters(data []interface{}) []ItemModelMaster {
	v := make([]ItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemModelMastersFromDict(data []ItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemModel struct {
	ItemModelId         *string `json:"itemModelId"`
	Name                *string `json:"name"`
	Metadata            *string `json:"metadata"`
	StackingLimit       *int64  `json:"stackingLimit"`
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
	SortValue           *int32  `json:"sortValue"`
}

func NewItemModelFromJson(data string) ItemModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewItemModelFromDict(dict)
}

func NewItemModelFromDict(data map[string]interface{}) ItemModel {
	return ItemModel{
		ItemModelId:         core.CastString(data["itemModelId"]),
		Name:                core.CastString(data["name"]),
		Metadata:            core.CastString(data["metadata"]),
		StackingLimit:       core.CastInt64(data["stackingLimit"]),
		AllowMultipleStacks: core.CastBool(data["allowMultipleStacks"]),
		SortValue:           core.CastInt32(data["sortValue"]),
	}
}

func (p ItemModel) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"itemModelId":         p.ItemModelId,
		"name":                p.Name,
		"metadata":            p.Metadata,
		"stackingLimit":       p.StackingLimit,
		"allowMultipleStacks": p.AllowMultipleStacks,
		"sortValue":           p.SortValue,
	}
}

func (p ItemModel) Pointer() *ItemModel {
	return &p
}

func CastItemModels(data []interface{}) []ItemModel {
	v := make([]ItemModel, 0)
	for _, d := range data {
		v = append(v, NewItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemModelsFromDict(data []ItemModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentItemModelMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentItemModelMasterFromJson(data string) CurrentItemModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentItemModelMasterFromDict(dict)
}

func NewCurrentItemModelMasterFromDict(data map[string]interface{}) CurrentItemModelMaster {
	return CurrentItemModelMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentItemModelMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
	}
}

func (p CurrentItemModelMaster) Pointer() *CurrentItemModelMaster {
	return &p
}

func CastCurrentItemModelMasters(data []interface{}) []CurrentItemModelMaster {
	v := make([]CurrentItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentItemModelMastersFromDict(data []CurrentItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Inventory struct {
	InventoryId                   *string `json:"inventoryId"`
	InventoryName                 *string `json:"inventoryName"`
	UserId                        *string `json:"userId"`
	CurrentInventoryCapacityUsage *int32  `json:"currentInventoryCapacityUsage"`
	CurrentInventoryMaxCapacity   *int32  `json:"currentInventoryMaxCapacity"`
	CreatedAt                     *int64  `json:"createdAt"`
	UpdatedAt                     *int64  `json:"updatedAt"`
}

func NewInventoryFromJson(data string) Inventory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewInventoryFromDict(dict)
}

func NewInventoryFromDict(data map[string]interface{}) Inventory {
	return Inventory{
		InventoryId:                   core.CastString(data["inventoryId"]),
		InventoryName:                 core.CastString(data["inventoryName"]),
		UserId:                        core.CastString(data["userId"]),
		CurrentInventoryCapacityUsage: core.CastInt32(data["currentInventoryCapacityUsage"]),
		CurrentInventoryMaxCapacity:   core.CastInt32(data["currentInventoryMaxCapacity"]),
		CreatedAt:                     core.CastInt64(data["createdAt"]),
		UpdatedAt:                     core.CastInt64(data["updatedAt"]),
	}
}

func (p Inventory) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"inventoryId":                   p.InventoryId,
		"inventoryName":                 p.InventoryName,
		"userId":                        p.UserId,
		"currentInventoryCapacityUsage": p.CurrentInventoryCapacityUsage,
		"currentInventoryMaxCapacity":   p.CurrentInventoryMaxCapacity,
		"createdAt":                     p.CreatedAt,
		"updatedAt":                     p.UpdatedAt,
	}
}

func (p Inventory) Pointer() *Inventory {
	return &p
}

func CastInventories(data []interface{}) []Inventory {
	v := make([]Inventory, 0)
	for _, d := range data {
		v = append(v, NewInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInventoriesFromDict(data []Inventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ItemSet struct {
	ItemSetId     *string  `json:"itemSetId"`
	Name          *string  `json:"name"`
	InventoryName *string  `json:"inventoryName"`
	UserId        *string  `json:"userId"`
	ItemName      *string  `json:"itemName"`
	Count         *int64   `json:"count"`
	ReferenceOf   []string `json:"referenceOf"`
	SortValue     *int32   `json:"sortValue"`
	ExpiresAt     *int64   `json:"expiresAt"`
	CreatedAt     *int64   `json:"createdAt"`
	UpdatedAt     *int64   `json:"updatedAt"`
}

func NewItemSetFromJson(data string) ItemSet {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewItemSetFromDict(dict)
}

func NewItemSetFromDict(data map[string]interface{}) ItemSet {
	return ItemSet{
		ItemSetId:     core.CastString(data["itemSetId"]),
		Name:          core.CastString(data["name"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		ItemName:      core.CastString(data["itemName"]),
		Count:         core.CastInt64(data["count"]),
		ReferenceOf:   core.CastStrings(core.CastArray(data["referenceOf"])),
		SortValue:     core.CastInt32(data["sortValue"]),
		ExpiresAt:     core.CastInt64(data["expiresAt"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
	}
}

func (p ItemSet) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"itemSetId":     p.ItemSetId,
		"name":          p.Name,
		"inventoryName": p.InventoryName,
		"userId":        p.UserId,
		"itemName":      p.ItemName,
		"count":         p.Count,
		"referenceOf": core.CastStringsFromDict(
			p.ReferenceOf,
		),
		"sortValue": p.SortValue,
		"expiresAt": p.ExpiresAt,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p ItemSet) Pointer() *ItemSet {
	return &p
}

func CastItemSets(data []interface{}) []ItemSet {
	v := make([]ItemSet, 0)
	for _, d := range data {
		v = append(v, NewItemSetFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastItemSetsFromDict(data []ItemSet) []interface{} {
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
