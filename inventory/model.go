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
	NamespaceId             *string        `json:"namespaceId"`
	Name                    *string        `json:"name"`
	Description             *string        `json:"description"`
	AcquireScript           *ScriptSetting `json:"acquireScript"`
	OverflowScript          *ScriptSetting `json:"overflowScript"`
	ConsumeScript           *ScriptSetting `json:"consumeScript"`
	SimpleItemAcquireScript *ScriptSetting `json:"simpleItemAcquireScript"`
	SimpleItemConsumeScript *ScriptSetting `json:"simpleItemConsumeScript"`
	LogSetting              *LogSetting    `json:"logSetting"`
	CreatedAt               *int64         `json:"createdAt"`
	UpdatedAt               *int64         `json:"updatedAt"`
	Revision                *int64         `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:             core.CastString(data["namespaceId"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		AcquireScript:           NewScriptSettingFromDict(core.CastMap(data["acquireScript"])).Pointer(),
		OverflowScript:          NewScriptSettingFromDict(core.CastMap(data["overflowScript"])).Pointer(),
		ConsumeScript:           NewScriptSettingFromDict(core.CastMap(data["consumeScript"])).Pointer(),
		SimpleItemAcquireScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemAcquireScript"])).Pointer(),
		SimpleItemConsumeScript: NewScriptSettingFromDict(core.CastMap(data["simpleItemConsumeScript"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
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
	var acquireScript map[string]interface{}
	if p.AcquireScript != nil {
		acquireScript = p.AcquireScript.ToDict()
	}
	var overflowScript map[string]interface{}
	if p.OverflowScript != nil {
		overflowScript = p.OverflowScript.ToDict()
	}
	var consumeScript map[string]interface{}
	if p.ConsumeScript != nil {
		consumeScript = p.ConsumeScript.ToDict()
	}
	var simpleItemAcquireScript map[string]interface{}
	if p.SimpleItemAcquireScript != nil {
		simpleItemAcquireScript = p.SimpleItemAcquireScript.ToDict()
	}
	var simpleItemConsumeScript map[string]interface{}
	if p.SimpleItemConsumeScript != nil {
		simpleItemConsumeScript = p.SimpleItemConsumeScript.ToDict()
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
		"namespaceId":             namespaceId,
		"name":                    name,
		"description":             description,
		"acquireScript":           acquireScript,
		"overflowScript":          overflowScript,
		"consumeScript":           consumeScript,
		"simpleItemAcquireScript": simpleItemAcquireScript,
		"simpleItemConsumeScript": simpleItemConsumeScript,
		"logSetting":              logSetting,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
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
	Revision              *int64  `json:"revision"`
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
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p InventoryModelMaster) ToDict() map[string]interface{} {

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var initialCapacity *int32
	if p.InitialCapacity != nil {
		initialCapacity = p.InitialCapacity
	}
	var maxCapacity *int32
	if p.MaxCapacity != nil {
		maxCapacity = p.MaxCapacity
	}
	var protectReferencedItem *bool
	if p.ProtectReferencedItem != nil {
		protectReferencedItem = p.ProtectReferencedItem
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
		"inventoryModelId":      inventoryModelId,
		"name":                  name,
		"metadata":              metadata,
		"description":           description,
		"initialCapacity":       initialCapacity,
		"maxCapacity":           maxCapacity,
		"protectReferencedItem": protectReferencedItem,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
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

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var initialCapacity *int32
	if p.InitialCapacity != nil {
		initialCapacity = p.InitialCapacity
	}
	var maxCapacity *int32
	if p.MaxCapacity != nil {
		maxCapacity = p.MaxCapacity
	}
	var protectReferencedItem *bool
	if p.ProtectReferencedItem != nil {
		protectReferencedItem = p.ProtectReferencedItem
	}
	var itemModels []interface{}
	if p.ItemModels != nil {
		itemModels = CastItemModelsFromDict(
			p.ItemModels,
		)
	}
	return map[string]interface{}{
		"inventoryModelId":      inventoryModelId,
		"name":                  name,
		"metadata":              metadata,
		"initialCapacity":       initialCapacity,
		"maxCapacity":           maxCapacity,
		"protectReferencedItem": protectReferencedItem,
		"itemModels":            itemModels,
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
	Revision            *int64  `json:"revision"`
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
		Revision:            core.CastInt64(data["revision"]),
	}
}

func (p ItemModelMaster) ToDict() map[string]interface{} {

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
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
	var stackingLimit *int64
	if p.StackingLimit != nil {
		stackingLimit = p.StackingLimit
	}
	var allowMultipleStacks *bool
	if p.AllowMultipleStacks != nil {
		allowMultipleStacks = p.AllowMultipleStacks
	}
	var sortValue *int32
	if p.SortValue != nil {
		sortValue = p.SortValue
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
		"itemModelId":         itemModelId,
		"inventoryName":       inventoryName,
		"name":                name,
		"description":         description,
		"metadata":            metadata,
		"stackingLimit":       stackingLimit,
		"allowMultipleStacks": allowMultipleStacks,
		"sortValue":           sortValue,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"revision":            revision,
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

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var stackingLimit *int64
	if p.StackingLimit != nil {
		stackingLimit = p.StackingLimit
	}
	var allowMultipleStacks *bool
	if p.AllowMultipleStacks != nil {
		allowMultipleStacks = p.AllowMultipleStacks
	}
	var sortValue *int32
	if p.SortValue != nil {
		sortValue = p.SortValue
	}
	return map[string]interface{}{
		"itemModelId":         itemModelId,
		"name":                name,
		"metadata":            metadata,
		"stackingLimit":       stackingLimit,
		"allowMultipleStacks": allowMultipleStacks,
		"sortValue":           sortValue,
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

type SimpleInventoryModelMaster struct {
	InventoryModelId *string `json:"inventoryModelId"`
	Name             *string `json:"name"`
	Metadata         *string `json:"metadata"`
	Description      *string `json:"description"`
	CreatedAt        *int64  `json:"createdAt"`
	UpdatedAt        *int64  `json:"updatedAt"`
	Revision         *int64  `json:"revision"`
}

func NewSimpleInventoryModelMasterFromJson(data string) SimpleInventoryModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleInventoryModelMasterFromDict(dict)
}

func NewSimpleInventoryModelMasterFromDict(data map[string]interface{}) SimpleInventoryModelMaster {
	return SimpleInventoryModelMaster{
		InventoryModelId: core.CastString(data["inventoryModelId"]),
		Name:             core.CastString(data["name"]),
		Metadata:         core.CastString(data["metadata"]),
		Description:      core.CastString(data["description"]),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
	}
}

func (p SimpleInventoryModelMaster) ToDict() map[string]interface{} {

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
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
		"inventoryModelId": inventoryModelId,
		"name":             name,
		"metadata":         metadata,
		"description":      description,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
	}
}

func (p SimpleInventoryModelMaster) Pointer() *SimpleInventoryModelMaster {
	return &p
}

func CastSimpleInventoryModelMasters(data []interface{}) []SimpleInventoryModelMaster {
	v := make([]SimpleInventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoryModelMastersFromDict(data []SimpleInventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleInventoryModel struct {
	InventoryModelId *string           `json:"inventoryModelId"`
	Name             *string           `json:"name"`
	Metadata         *string           `json:"metadata"`
	SimpleItemModels []SimpleItemModel `json:"simpleItemModels"`
}

func NewSimpleInventoryModelFromJson(data string) SimpleInventoryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleInventoryModelFromDict(dict)
}

func NewSimpleInventoryModelFromDict(data map[string]interface{}) SimpleInventoryModel {
	return SimpleInventoryModel{
		InventoryModelId: core.CastString(data["inventoryModelId"]),
		Name:             core.CastString(data["name"]),
		Metadata:         core.CastString(data["metadata"]),
		SimpleItemModels: CastSimpleItemModels(core.CastArray(data["simpleItemModels"])),
	}
}

func (p SimpleInventoryModel) ToDict() map[string]interface{} {

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var simpleItemModels []interface{}
	if p.SimpleItemModels != nil {
		simpleItemModels = CastSimpleItemModelsFromDict(
			p.SimpleItemModels,
		)
	}
	return map[string]interface{}{
		"inventoryModelId": inventoryModelId,
		"name":             name,
		"metadata":         metadata,
		"simpleItemModels": simpleItemModels,
	}
}

func (p SimpleInventoryModel) Pointer() *SimpleInventoryModel {
	return &p
}

func CastSimpleInventoryModels(data []interface{}) []SimpleInventoryModel {
	v := make([]SimpleInventoryModel, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoryModelsFromDict(data []SimpleInventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItemModelMaster struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Metadata    *string `json:"metadata"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func NewSimpleItemModelMasterFromJson(data string) SimpleItemModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleItemModelMasterFromDict(dict)
}

func NewSimpleItemModelMasterFromDict(data map[string]interface{}) SimpleItemModelMaster {
	return SimpleItemModelMaster{
		ItemModelId: core.CastString(data["itemModelId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Metadata:    core.CastString(data["metadata"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p SimpleItemModelMaster) ToDict() map[string]interface{} {

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
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
		"itemModelId": itemModelId,
		"name":        name,
		"description": description,
		"metadata":    metadata,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p SimpleItemModelMaster) Pointer() *SimpleItemModelMaster {
	return &p
}

func CastSimpleItemModelMasters(data []interface{}) []SimpleItemModelMaster {
	v := make([]SimpleItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemModelMastersFromDict(data []SimpleItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItemModel struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Metadata    *string `json:"metadata"`
}

func NewSimpleItemModelFromJson(data string) SimpleItemModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleItemModelFromDict(dict)
}

func NewSimpleItemModelFromDict(data map[string]interface{}) SimpleItemModel {
	return SimpleItemModel{
		ItemModelId: core.CastString(data["itemModelId"]),
		Name:        core.CastString(data["name"]),
		Metadata:    core.CastString(data["metadata"]),
	}
}

func (p SimpleItemModel) ToDict() map[string]interface{} {

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"itemModelId": itemModelId,
		"name":        name,
		"metadata":    metadata,
	}
}

func (p SimpleItemModel) Pointer() *SimpleItemModel {
	return &p
}

func CastSimpleItemModels(data []interface{}) []SimpleItemModel {
	v := make([]SimpleItemModel, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemModelsFromDict(data []SimpleItemModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventoryModelMaster struct {
	InventoryModelId *string `json:"inventoryModelId"`
	Name             *string `json:"name"`
	Metadata         *string `json:"metadata"`
	Description      *string `json:"description"`
	CreatedAt        *int64  `json:"createdAt"`
	UpdatedAt        *int64  `json:"updatedAt"`
	Revision         *int64  `json:"revision"`
}

func NewBigInventoryModelMasterFromJson(data string) BigInventoryModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigInventoryModelMasterFromDict(dict)
}

func NewBigInventoryModelMasterFromDict(data map[string]interface{}) BigInventoryModelMaster {
	return BigInventoryModelMaster{
		InventoryModelId: core.CastString(data["inventoryModelId"]),
		Name:             core.CastString(data["name"]),
		Metadata:         core.CastString(data["metadata"]),
		Description:      core.CastString(data["description"]),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
	}
}

func (p BigInventoryModelMaster) ToDict() map[string]interface{} {

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var description *string
	if p.Description != nil {
		description = p.Description
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
		"inventoryModelId": inventoryModelId,
		"name":             name,
		"metadata":         metadata,
		"description":      description,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
	}
}

func (p BigInventoryModelMaster) Pointer() *BigInventoryModelMaster {
	return &p
}

func CastBigInventoryModelMasters(data []interface{}) []BigInventoryModelMaster {
	v := make([]BigInventoryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoryModelMastersFromDict(data []BigInventoryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventoryModel struct {
	InventoryModelId *string        `json:"inventoryModelId"`
	Name             *string        `json:"name"`
	Metadata         *string        `json:"metadata"`
	BigItemModels    []BigItemModel `json:"bigItemModels"`
}

func NewBigInventoryModelFromJson(data string) BigInventoryModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigInventoryModelFromDict(dict)
}

func NewBigInventoryModelFromDict(data map[string]interface{}) BigInventoryModel {
	return BigInventoryModel{
		InventoryModelId: core.CastString(data["inventoryModelId"]),
		Name:             core.CastString(data["name"]),
		Metadata:         core.CastString(data["metadata"]),
		BigItemModels:    CastBigItemModels(core.CastArray(data["bigItemModels"])),
	}
}

func (p BigInventoryModel) ToDict() map[string]interface{} {

	var inventoryModelId *string
	if p.InventoryModelId != nil {
		inventoryModelId = p.InventoryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var bigItemModels []interface{}
	if p.BigItemModels != nil {
		bigItemModels = CastBigItemModelsFromDict(
			p.BigItemModels,
		)
	}
	return map[string]interface{}{
		"inventoryModelId": inventoryModelId,
		"name":             name,
		"metadata":         metadata,
		"bigItemModels":    bigItemModels,
	}
}

func (p BigInventoryModel) Pointer() *BigInventoryModel {
	return &p
}

func CastBigInventoryModels(data []interface{}) []BigInventoryModel {
	v := make([]BigInventoryModel, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoryModelsFromDict(data []BigInventoryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItemModelMaster struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Metadata    *string `json:"metadata"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func NewBigItemModelMasterFromJson(data string) BigItemModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigItemModelMasterFromDict(dict)
}

func NewBigItemModelMasterFromDict(data map[string]interface{}) BigItemModelMaster {
	return BigItemModelMaster{
		ItemModelId: core.CastString(data["itemModelId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Metadata:    core.CastString(data["metadata"]),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p BigItemModelMaster) ToDict() map[string]interface{} {

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
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
		"itemModelId": itemModelId,
		"name":        name,
		"description": description,
		"metadata":    metadata,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p BigItemModelMaster) Pointer() *BigItemModelMaster {
	return &p
}

func CastBigItemModelMasters(data []interface{}) []BigItemModelMaster {
	v := make([]BigItemModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBigItemModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemModelMastersFromDict(data []BigItemModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItemModel struct {
	ItemModelId *string `json:"itemModelId"`
	Name        *string `json:"name"`
	Metadata    *string `json:"metadata"`
}

func NewBigItemModelFromJson(data string) BigItemModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigItemModelFromDict(dict)
}

func NewBigItemModelFromDict(data map[string]interface{}) BigItemModel {
	return BigItemModel{
		ItemModelId: core.CastString(data["itemModelId"]),
		Name:        core.CastString(data["name"]),
		Metadata:    core.CastString(data["metadata"]),
	}
}

func (p BigItemModel) ToDict() map[string]interface{} {

	var itemModelId *string
	if p.ItemModelId != nil {
		itemModelId = p.ItemModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"itemModelId": itemModelId,
		"name":        name,
		"metadata":    metadata,
	}
}

func (p BigItemModel) Pointer() *BigItemModel {
	return &p
}

func CastBigItemModels(data []interface{}) []BigItemModel {
	v := make([]BigItemModel, 0)
	for _, d := range data {
		v = append(v, NewBigItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemModelsFromDict(data []BigItemModel) []interface{} {
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
	Revision                      *int64  `json:"revision"`
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
		Revision:                      core.CastInt64(data["revision"]),
	}
}

func (p Inventory) ToDict() map[string]interface{} {

	var inventoryId *string
	if p.InventoryId != nil {
		inventoryId = p.InventoryId
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var currentInventoryCapacityUsage *int32
	if p.CurrentInventoryCapacityUsage != nil {
		currentInventoryCapacityUsage = p.CurrentInventoryCapacityUsage
	}
	var currentInventoryMaxCapacity *int32
	if p.CurrentInventoryMaxCapacity != nil {
		currentInventoryMaxCapacity = p.CurrentInventoryMaxCapacity
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
		"inventoryId":                   inventoryId,
		"inventoryName":                 inventoryName,
		"userId":                        userId,
		"currentInventoryCapacityUsage": currentInventoryCapacityUsage,
		"currentInventoryMaxCapacity":   currentInventoryMaxCapacity,
		"createdAt":                     createdAt,
		"updatedAt":                     updatedAt,
		"revision":                      revision,
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
	ItemSetId     *string   `json:"itemSetId"`
	Name          *string   `json:"name"`
	InventoryName *string   `json:"inventoryName"`
	UserId        *string   `json:"userId"`
	ItemName      *string   `json:"itemName"`
	Count         *int64    `json:"count"`
	ReferenceOf   []*string `json:"referenceOf"`
	SortValue     *int32    `json:"sortValue"`
	ExpiresAt     *int64    `json:"expiresAt"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
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

	var itemSetId *string
	if p.ItemSetId != nil {
		itemSetId = p.ItemSetId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var referenceOf []interface{}
	if p.ReferenceOf != nil {
		referenceOf = core.CastStringsFromDict(
			p.ReferenceOf,
		)
	}
	var sortValue *int32
	if p.SortValue != nil {
		sortValue = p.SortValue
	}
	var expiresAt *int64
	if p.ExpiresAt != nil {
		expiresAt = p.ExpiresAt
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var updatedAt *int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt
	}
	return map[string]interface{}{
		"itemSetId":     itemSetId,
		"name":          name,
		"inventoryName": inventoryName,
		"userId":        userId,
		"itemName":      itemName,
		"count":         count,
		"referenceOf":   referenceOf,
		"sortValue":     sortValue,
		"expiresAt":     expiresAt,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
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

type ReferenceOf struct {
	ReferenceOfId *string `json:"referenceOfId"`
	Name          *string `json:"name"`
}

func NewReferenceOfFromJson(data string) ReferenceOf {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReferenceOfFromDict(dict)
}

func NewReferenceOfFromDict(data map[string]interface{}) ReferenceOf {
	return ReferenceOf{
		ReferenceOfId: core.CastString(data["referenceOfId"]),
		Name:          core.CastString(data["name"]),
	}
}

func (p ReferenceOf) ToDict() map[string]interface{} {

	var referenceOfId *string
	if p.ReferenceOfId != nil {
		referenceOfId = p.ReferenceOfId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	return map[string]interface{}{
		"referenceOfId": referenceOfId,
		"name":          name,
	}
}

func (p ReferenceOf) Pointer() *ReferenceOf {
	return &p
}

func CastReferenceOves(data []interface{}) []ReferenceOf {
	v := make([]ReferenceOf, 0)
	for _, d := range data {
		v = append(v, NewReferenceOfFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastReferenceOvesFromDict(data []ReferenceOf) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleInventory struct {
	InventoryId   *string      `json:"inventoryId"`
	InventoryName *string      `json:"inventoryName"`
	UserId        *string      `json:"userId"`
	SimpleItems   []SimpleItem `json:"simpleItems"`
	CreatedAt     *int64       `json:"createdAt"`
	UpdatedAt     *int64       `json:"updatedAt"`
	Revision      *int64       `json:"revision"`
}

func NewSimpleInventoryFromJson(data string) SimpleInventory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleInventoryFromDict(dict)
}

func NewSimpleInventoryFromDict(data map[string]interface{}) SimpleInventory {
	return SimpleInventory{
		InventoryId:   core.CastString(data["inventoryId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		SimpleItems:   CastSimpleItems(core.CastArray(data["simpleItems"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p SimpleInventory) ToDict() map[string]interface{} {

	var inventoryId *string
	if p.InventoryId != nil {
		inventoryId = p.InventoryId
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var simpleItems []interface{}
	if p.SimpleItems != nil {
		simpleItems = CastSimpleItemsFromDict(
			p.SimpleItems,
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
		"inventoryId":   inventoryId,
		"inventoryName": inventoryName,
		"userId":        userId,
		"simpleItems":   simpleItems,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
		"revision":      revision,
	}
}

func (p SimpleInventory) Pointer() *SimpleInventory {
	return &p
}

func CastSimpleInventories(data []interface{}) []SimpleInventory {
	v := make([]SimpleInventory, 0)
	for _, d := range data {
		v = append(v, NewSimpleInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleInventoriesFromDict(data []SimpleInventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SimpleItem struct {
	ItemId   *string `json:"itemId"`
	UserId   *string `json:"userId"`
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
	Revision *int64  `json:"revision"`
}

func NewSimpleItemFromJson(data string) SimpleItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSimpleItemFromDict(dict)
}

func NewSimpleItemFromDict(data map[string]interface{}) SimpleItem {
	return SimpleItem{
		ItemId:   core.CastString(data["itemId"]),
		UserId:   core.CastString(data["userId"]),
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
		Revision: core.CastInt64(data["revision"]),
	}
}

func (p SimpleItem) ToDict() map[string]interface{} {

	var itemId *string
	if p.ItemId != nil {
		itemId = p.ItemId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"itemId":   itemId,
		"userId":   userId,
		"itemName": itemName,
		"count":    count,
		"revision": revision,
	}
}

func (p SimpleItem) Pointer() *SimpleItem {
	return &p
}

func CastSimpleItems(data []interface{}) []SimpleItem {
	v := make([]SimpleItem, 0)
	for _, d := range data {
		v = append(v, NewSimpleItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSimpleItemsFromDict(data []SimpleItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigInventory struct {
	InventoryId   *string   `json:"inventoryId"`
	InventoryName *string   `json:"inventoryName"`
	UserId        *string   `json:"userId"`
	BigItems      []BigItem `json:"bigItems"`
	CreatedAt     *int64    `json:"createdAt"`
	UpdatedAt     *int64    `json:"updatedAt"`
}

func NewBigInventoryFromJson(data string) BigInventory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigInventoryFromDict(dict)
}

func NewBigInventoryFromDict(data map[string]interface{}) BigInventory {
	return BigInventory{
		InventoryId:   core.CastString(data["inventoryId"]),
		InventoryName: core.CastString(data["inventoryName"]),
		UserId:        core.CastString(data["userId"]),
		BigItems:      CastBigItems(core.CastArray(data["bigItems"])),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		UpdatedAt:     core.CastInt64(data["updatedAt"]),
	}
}

func (p BigInventory) ToDict() map[string]interface{} {

	var inventoryId *string
	if p.InventoryId != nil {
		inventoryId = p.InventoryId
	}
	var inventoryName *string
	if p.InventoryName != nil {
		inventoryName = p.InventoryName
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var bigItems []interface{}
	if p.BigItems != nil {
		bigItems = CastBigItemsFromDict(
			p.BigItems,
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
	return map[string]interface{}{
		"inventoryId":   inventoryId,
		"inventoryName": inventoryName,
		"userId":        userId,
		"bigItems":      bigItems,
		"createdAt":     createdAt,
		"updatedAt":     updatedAt,
	}
}

func (p BigInventory) Pointer() *BigInventory {
	return &p
}

func CastBigInventories(data []interface{}) []BigInventory {
	v := make([]BigInventory, 0)
	for _, d := range data {
		v = append(v, NewBigInventoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigInventoriesFromDict(data []BigInventory) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BigItem struct {
	ItemId    *string `json:"itemId"`
	UserId    *string `json:"userId"`
	ItemName  *string `json:"itemName"`
	Count     *string `json:"count"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
	Revision  *int64  `json:"revision"`
}

func NewBigItemFromJson(data string) BigItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBigItemFromDict(dict)
}

func NewBigItemFromDict(data map[string]interface{}) BigItem {
	return BigItem{
		ItemId:    core.CastString(data["itemId"]),
		UserId:    core.CastString(data["userId"]),
		ItemName:  core.CastString(data["itemName"]),
		Count:     core.CastString(data["count"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p BigItem) ToDict() map[string]interface{} {

	var itemId *string
	if p.ItemId != nil {
		itemId = p.ItemId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *string
	if p.Count != nil {
		count = p.Count
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
		"itemId":    itemId,
		"userId":    userId,
		"itemName":  itemName,
		"count":     count,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
}

func (p BigItem) Pointer() *BigItem {
	return &p
}

func CastBigItems(data []interface{}) []BigItem {
	v := make([]BigItem, 0)
	for _, d := range data {
		v = append(v, NewBigItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBigItemsFromDict(data []BigItem) []interface{} {
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

type AcquireCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func NewAcquireCountFromJson(data string) AcquireCount {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireCountFromDict(dict)
}

func NewAcquireCountFromDict(data map[string]interface{}) AcquireCount {
	return AcquireCount{
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p AcquireCount) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"itemName": itemName,
		"count":    count,
	}
}

func (p AcquireCount) Pointer() *AcquireCount {
	return &p
}

func CastAcquireCounts(data []interface{}) []AcquireCount {
	v := make([]AcquireCount, 0)
	for _, d := range data {
		v = append(v, NewAcquireCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireCountsFromDict(data []AcquireCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func NewConsumeCountFromJson(data string) ConsumeCount {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeCountFromDict(dict)
}

func NewConsumeCountFromDict(data map[string]interface{}) ConsumeCount {
	return ConsumeCount{
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p ConsumeCount) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"itemName": itemName,
		"count":    count,
	}
}

func (p ConsumeCount) Pointer() *ConsumeCount {
	return &p
}

func CastConsumeCounts(data []interface{}) []ConsumeCount {
	v := make([]ConsumeCount, 0)
	for _, d := range data {
		v = append(v, NewConsumeCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeCountsFromDict(data []ConsumeCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type HeldCount struct {
	ItemName *string `json:"itemName"`
	Count    *int64  `json:"count"`
}

func NewHeldCountFromJson(data string) HeldCount {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewHeldCountFromDict(dict)
}

func NewHeldCountFromDict(data map[string]interface{}) HeldCount {
	return HeldCount{
		ItemName: core.CastString(data["itemName"]),
		Count:    core.CastInt64(data["count"]),
	}
}

func (p HeldCount) ToDict() map[string]interface{} {

	var itemName *string
	if p.ItemName != nil {
		itemName = p.ItemName
	}
	var count *int64
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"itemName": itemName,
		"count":    count,
	}
}

func (p HeldCount) Pointer() *HeldCount {
	return &p
}

func CastHeldCounts(data []interface{}) []HeldCount {
	v := make([]HeldCount, 0)
	for _, d := range data {
		v = append(v, NewHeldCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastHeldCountsFromDict(data []HeldCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
