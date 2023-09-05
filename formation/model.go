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

package formation

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	UpdateMoldScript   *ScriptSetting      `json:"updateMoldScript"`
	UpdateFormScript   *ScriptSetting      `json:"updateFormScript"`
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
		UpdateMoldScript:   NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
		UpdateFormScript:   NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
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
	var updateMoldScript map[string]interface{}
	if p.UpdateMoldScript != nil {
		updateMoldScript = p.UpdateMoldScript.ToDict()
	}
	var updateFormScript map[string]interface{}
	if p.UpdateFormScript != nil {
		updateFormScript = p.UpdateFormScript.ToDict()
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
		"updateMoldScript":   updateMoldScript,
		"updateFormScript":   updateFormScript,
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

type FormModel struct {
	FormModelId *string     `json:"formModelId"`
	Name        *string     `json:"name"`
	Metadata    *string     `json:"metadata"`
	Slots       []SlotModel `json:"slots"`
}

func NewFormModelFromJson(data string) FormModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormModelFromDict(dict)
}

func NewFormModelFromDict(data map[string]interface{}) FormModel {
	return FormModel{
		FormModelId: core.CastString(data["formModelId"]),
		Name:        core.CastString(data["name"]),
		Metadata:    core.CastString(data["metadata"]),
		Slots:       CastSlotModels(core.CastArray(data["slots"])),
	}
}

func (p FormModel) ToDict() map[string]interface{} {

	var formModelId *string
	if p.FormModelId != nil {
		formModelId = p.FormModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var slots []interface{}
	if p.Slots != nil {
		slots = CastSlotModelsFromDict(
			p.Slots,
		)
	}
	return map[string]interface{}{
		"formModelId": formModelId,
		"name":        name,
		"metadata":    metadata,
		"slots":       slots,
	}
}

func (p FormModel) Pointer() *FormModel {
	return &p
}

func CastFormModels(data []interface{}) []FormModel {
	v := make([]FormModel, 0)
	for _, d := range data {
		v = append(v, NewFormModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormModelsFromDict(data []FormModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FormModelMaster struct {
	FormModelId *string     `json:"formModelId"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	Metadata    *string     `json:"metadata"`
	Slots       []SlotModel `json:"slots"`
	CreatedAt   *int64      `json:"createdAt"`
	UpdatedAt   *int64      `json:"updatedAt"`
	Revision    *int64      `json:"revision"`
}

func NewFormModelMasterFromJson(data string) FormModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormModelMasterFromDict(dict)
}

func NewFormModelMasterFromDict(data map[string]interface{}) FormModelMaster {
	return FormModelMaster{
		FormModelId: core.CastString(data["formModelId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		Metadata:    core.CastString(data["metadata"]),
		Slots:       CastSlotModels(core.CastArray(data["slots"])),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
	}
}

func (p FormModelMaster) ToDict() map[string]interface{} {

	var formModelId *string
	if p.FormModelId != nil {
		formModelId = p.FormModelId
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
	var slots []interface{}
	if p.Slots != nil {
		slots = CastSlotModelsFromDict(
			p.Slots,
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
		"formModelId": formModelId,
		"name":        name,
		"description": description,
		"metadata":    metadata,
		"slots":       slots,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p FormModelMaster) Pointer() *FormModelMaster {
	return &p
}

func CastFormModelMasters(data []interface{}) []FormModelMaster {
	v := make([]FormModelMaster, 0)
	for _, d := range data {
		v = append(v, NewFormModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormModelMastersFromDict(data []FormModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoldModel struct {
	MoldModelId        *string    `json:"moldModelId"`
	Name               *string    `json:"name"`
	Metadata           *string    `json:"metadata"`
	InitialMaxCapacity *int32     `json:"initialMaxCapacity"`
	MaxCapacity        *int32     `json:"maxCapacity"`
	FormModel          *FormModel `json:"formModel"`
}

func NewMoldModelFromJson(data string) MoldModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoldModelFromDict(dict)
}

func NewMoldModelFromDict(data map[string]interface{}) MoldModel {
	return MoldModel{
		MoldModelId:        core.CastString(data["moldModelId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
		MaxCapacity:        core.CastInt32(data["maxCapacity"]),
		FormModel:          NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p MoldModel) ToDict() map[string]interface{} {

	var moldModelId *string
	if p.MoldModelId != nil {
		moldModelId = p.MoldModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var initialMaxCapacity *int32
	if p.InitialMaxCapacity != nil {
		initialMaxCapacity = p.InitialMaxCapacity
	}
	var maxCapacity *int32
	if p.MaxCapacity != nil {
		maxCapacity = p.MaxCapacity
	}
	var formModel map[string]interface{}
	if p.FormModel != nil {
		formModel = p.FormModel.ToDict()
	}
	return map[string]interface{}{
		"moldModelId":        moldModelId,
		"name":               name,
		"metadata":           metadata,
		"initialMaxCapacity": initialMaxCapacity,
		"maxCapacity":        maxCapacity,
		"formModel":          formModel,
	}
}

func (p MoldModel) Pointer() *MoldModel {
	return &p
}

func CastMoldModels(data []interface{}) []MoldModel {
	v := make([]MoldModel, 0)
	for _, d := range data {
		v = append(v, NewMoldModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoldModelsFromDict(data []MoldModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MoldModelMaster struct {
	MoldModelId        *string `json:"moldModelId"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	InitialMaxCapacity *int32  `json:"initialMaxCapacity"`
	MaxCapacity        *int32  `json:"maxCapacity"`
	FormModelName      *string `json:"formModelName"`
	CreatedAt          *int64  `json:"createdAt"`
	UpdatedAt          *int64  `json:"updatedAt"`
	Revision           *int64  `json:"revision"`
}

func NewMoldModelMasterFromJson(data string) MoldModelMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoldModelMasterFromDict(dict)
}

func NewMoldModelMasterFromDict(data map[string]interface{}) MoldModelMaster {
	return MoldModelMaster{
		MoldModelId:        core.CastString(data["moldModelId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
		MaxCapacity:        core.CastInt32(data["maxCapacity"]),
		FormModelName:      core.CastString(data["formModelName"]),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p MoldModelMaster) ToDict() map[string]interface{} {

	var moldModelId *string
	if p.MoldModelId != nil {
		moldModelId = p.MoldModelId
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
	var initialMaxCapacity *int32
	if p.InitialMaxCapacity != nil {
		initialMaxCapacity = p.InitialMaxCapacity
	}
	var maxCapacity *int32
	if p.MaxCapacity != nil {
		maxCapacity = p.MaxCapacity
	}
	var formModelName *string
	if p.FormModelName != nil {
		formModelName = p.FormModelName
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
		"moldModelId":        moldModelId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"initialMaxCapacity": initialMaxCapacity,
		"maxCapacity":        maxCapacity,
		"formModelName":      formModelName,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p MoldModelMaster) Pointer() *MoldModelMaster {
	return &p
}

func CastMoldModelMasters(data []interface{}) []MoldModelMaster {
	v := make([]MoldModelMaster, 0)
	for _, d := range data {
		v = append(v, NewMoldModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoldModelMastersFromDict(data []MoldModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentFormMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentFormMasterFromJson(data string) CurrentFormMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentFormMasterFromDict(dict)
}

func NewCurrentFormMasterFromDict(data map[string]interface{}) CurrentFormMaster {
	return CurrentFormMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentFormMaster) ToDict() map[string]interface{} {

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

func (p CurrentFormMaster) Pointer() *CurrentFormMaster {
	return &p
}

func CastCurrentFormMasters(data []interface{}) []CurrentFormMaster {
	v := make([]CurrentFormMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentFormMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentFormMastersFromDict(data []CurrentFormMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Mold struct {
	MoldId    *string `json:"moldId"`
	Name      *string `json:"name"`
	UserId    *string `json:"userId"`
	Capacity  *int32  `json:"capacity"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
	Revision  *int64  `json:"revision"`
}

func NewMoldFromJson(data string) Mold {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMoldFromDict(dict)
}

func NewMoldFromDict(data map[string]interface{}) Mold {
	return Mold{
		MoldId:    core.CastString(data["moldId"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		Capacity:  core.CastInt32(data["capacity"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Mold) ToDict() map[string]interface{} {

	var moldId *string
	if p.MoldId != nil {
		moldId = p.MoldId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var capacity *int32
	if p.Capacity != nil {
		capacity = p.Capacity
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
		"moldId":    moldId,
		"name":      name,
		"userId":    userId,
		"capacity":  capacity,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
}

func (p Mold) Pointer() *Mold {
	return &p
}

func CastMolds(data []interface{}) []Mold {
	v := make([]Mold, 0)
	for _, d := range data {
		v = append(v, NewMoldFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMoldsFromDict(data []Mold) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Form struct {
	FormId    *string `json:"formId"`
	Name      *string `json:"name"`
	Index     *int32  `json:"index"`
	Slots     []Slot  `json:"slots"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
	Revision  *int64  `json:"revision"`
}

func NewFormFromJson(data string) Form {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFormFromDict(dict)
}

func NewFormFromDict(data map[string]interface{}) Form {
	return Form{
		FormId:    core.CastString(data["formId"]),
		Name:      core.CastString(data["name"]),
		Index:     core.CastInt32(data["index"]),
		Slots:     CastSlots(core.CastArray(data["slots"])),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Form) ToDict() map[string]interface{} {

	var formId *string
	if p.FormId != nil {
		formId = p.FormId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var index *int32
	if p.Index != nil {
		index = p.Index
	}
	var slots []interface{}
	if p.Slots != nil {
		slots = CastSlotsFromDict(
			p.Slots,
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
		"formId":    formId,
		"name":      name,
		"index":     index,
		"slots":     slots,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
}

func (p Form) Pointer() *Form {
	return &p
}

func CastForms(data []interface{}) []Form {
	v := make([]Form, 0)
	for _, d := range data {
		v = append(v, NewFormFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFormsFromDict(data []Form) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PropertyForm struct {
	FormId     *string `json:"formId"`
	UserId     *string `json:"userId"`
	Name       *string `json:"name"`
	PropertyId *string `json:"propertyId"`
	Slots      []Slot  `json:"slots"`
	CreatedAt  *int64  `json:"createdAt"`
	UpdatedAt  *int64  `json:"updatedAt"`
	Revision   *int64  `json:"revision"`
}

func NewPropertyFormFromJson(data string) PropertyForm {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPropertyFormFromDict(dict)
}

func NewPropertyFormFromDict(data map[string]interface{}) PropertyForm {
	return PropertyForm{
		FormId:     core.CastString(data["formId"]),
		UserId:     core.CastString(data["userId"]),
		Name:       core.CastString(data["name"]),
		PropertyId: core.CastString(data["propertyId"]),
		Slots:      CastSlots(core.CastArray(data["slots"])),
		CreatedAt:  core.CastInt64(data["createdAt"]),
		UpdatedAt:  core.CastInt64(data["updatedAt"]),
		Revision:   core.CastInt64(data["revision"]),
	}
}

func (p PropertyForm) ToDict() map[string]interface{} {

	var formId *string
	if p.FormId != nil {
		formId = p.FormId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var slots []interface{}
	if p.Slots != nil {
		slots = CastSlotsFromDict(
			p.Slots,
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
		"formId":     formId,
		"userId":     userId,
		"name":       name,
		"propertyId": propertyId,
		"slots":      slots,
		"createdAt":  createdAt,
		"updatedAt":  updatedAt,
		"revision":   revision,
	}
}

func (p PropertyForm) Pointer() *PropertyForm {
	return &p
}

func CastPropertyForms(data []interface{}) []PropertyForm {
	v := make([]PropertyForm, 0)
	for _, d := range data {
		v = append(v, NewPropertyFormFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPropertyFormsFromDict(data []PropertyForm) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Slot struct {
	Name       *string `json:"name"`
	PropertyId *string `json:"propertyId"`
	Metadata   *string `json:"metadata"`
}

func NewSlotFromJson(data string) Slot {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSlotFromDict(dict)
}

func NewSlotFromDict(data map[string]interface{}) Slot {
	return Slot{
		Name:       core.CastString(data["name"]),
		PropertyId: core.CastString(data["propertyId"]),
		Metadata:   core.CastString(data["metadata"]),
	}
}

func (p Slot) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"name":       name,
		"propertyId": propertyId,
		"metadata":   metadata,
	}
}

func (p Slot) Pointer() *Slot {
	return &p
}

func CastSlots(data []interface{}) []Slot {
	v := make([]Slot, 0)
	for _, d := range data {
		v = append(v, NewSlotFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSlotsFromDict(data []Slot) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SlotModel struct {
	Name          *string `json:"name"`
	PropertyRegex *string `json:"propertyRegex"`
	Metadata      *string `json:"metadata"`
}

func NewSlotModelFromJson(data string) SlotModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSlotModelFromDict(dict)
}

func NewSlotModelFromDict(data map[string]interface{}) SlotModel {
	return SlotModel{
		Name:          core.CastString(data["name"]),
		PropertyRegex: core.CastString(data["propertyRegex"]),
		Metadata:      core.CastString(data["metadata"]),
	}
}

func (p SlotModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var propertyRegex *string
	if p.PropertyRegex != nil {
		propertyRegex = p.PropertyRegex
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"name":          name,
		"propertyRegex": propertyRegex,
		"metadata":      metadata,
	}
}

func (p SlotModel) Pointer() *SlotModel {
	return &p
}

func CastSlotModels(data []interface{}) []SlotModel {
	v := make([]SlotModel, 0)
	for _, d := range data {
		v = append(v, NewSlotModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSlotModelsFromDict(data []SlotModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SlotWithSignature struct {
	Name         *string `json:"name"`
	PropertyType *string `json:"propertyType"`
	Body         *string `json:"body"`
	Signature    *string `json:"signature"`
	Metadata     *string `json:"metadata"`
}

func NewSlotWithSignatureFromJson(data string) SlotWithSignature {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSlotWithSignatureFromDict(dict)
}

func NewSlotWithSignatureFromDict(data map[string]interface{}) SlotWithSignature {
	return SlotWithSignature{
		Name:         core.CastString(data["name"]),
		PropertyType: core.CastString(data["propertyType"]),
		Body:         core.CastString(data["body"]),
		Signature:    core.CastString(data["signature"]),
		Metadata:     core.CastString(data["metadata"]),
	}
}

func (p SlotWithSignature) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var propertyType *string
	if p.PropertyType != nil {
		propertyType = p.PropertyType
	}
	var body *string
	if p.Body != nil {
		body = p.Body
	}
	var signature *string
	if p.Signature != nil {
		signature = p.Signature
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	return map[string]interface{}{
		"name":         name,
		"propertyType": propertyType,
		"body":         body,
		"signature":    signature,
		"metadata":     metadata,
	}
}

func (p SlotWithSignature) Pointer() *SlotWithSignature {
	return &p
}

func CastSlotWithSignatures(data []interface{}) []SlotWithSignature {
	v := make([]SlotWithSignature, 0)
	for _, d := range data {
		v = append(v, NewSlotWithSignatureFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSlotWithSignaturesFromDict(data []SlotWithSignature) []interface{} {
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

type AcquireActionConfig struct {
	Name   *string  `json:"name"`
	Config []Config `json:"config"`
}

func NewAcquireActionConfigFromJson(data string) AcquireActionConfig {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionConfigFromDict(dict)
}

func NewAcquireActionConfigFromDict(data map[string]interface{}) AcquireActionConfig {
	return AcquireActionConfig{
		Name:   core.CastString(data["name"]),
		Config: CastConfigs(core.CastArray(data["config"])),
	}
}

func (p AcquireActionConfig) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var config []interface{}
	if p.Config != nil {
		config = CastConfigsFromDict(
			p.Config,
		)
	}
	return map[string]interface{}{
		"name":   name,
		"config": config,
	}
}

func (p AcquireActionConfig) Pointer() *AcquireActionConfig {
	return &p
}

func CastAcquireActionConfigs(data []interface{}) []AcquireActionConfig {
	v := make([]AcquireActionConfig, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionConfigsFromDict(data []AcquireActionConfig) []interface{} {
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
