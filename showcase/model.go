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

package showcase

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId        *string             `json:"namespaceId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	TransactionSetting *TransactionSetting `json:"transactionSetting"`
	BuyScript          *ScriptSetting      `json:"buyScript"`
	LogSetting         *LogSetting         `json:"logSetting"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
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
		NamespaceId:        core.CastString(data["namespaceId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		TransactionSetting: NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		BuyScript:          NewScriptSettingFromDict(core.CastMap(data["buyScript"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		QueueNamespaceId:   core.CastString(data["queueNamespaceId"]),
		KeyId:              core.CastString(data["keyId"]),
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
	var buyScript map[string]interface{}
	if p.BuyScript != nil {
		buyScript = p.BuyScript.ToDict()
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
		"namespaceId":        namespaceId,
		"name":               name,
		"description":        description,
		"transactionSetting": transactionSetting,
		"buyScript":          buyScript,
		"logSetting":         logSetting,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"queueNamespaceId":   queueNamespaceId,
		"keyId":              keyId,
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

type SalesItemMaster struct {
	SalesItemId    *string         `json:"salesItemId"`
	Name           *string         `json:"name"`
	Description    *string         `json:"description"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	CreatedAt      *int64          `json:"createdAt"`
	UpdatedAt      *int64          `json:"updatedAt"`
	Revision       *int64          `json:"revision"`
}

func NewSalesItemMasterFromJson(data string) SalesItemMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSalesItemMasterFromDict(dict)
}

func NewSalesItemMasterFromDict(data map[string]interface{}) SalesItemMaster {
	return SalesItemMaster{
		SalesItemId:    core.CastString(data["salesItemId"]),
		Name:           core.CastString(data["name"]),
		Description:    core.CastString(data["description"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
		CreatedAt:      core.CastInt64(data["createdAt"]),
		UpdatedAt:      core.CastInt64(data["updatedAt"]),
		Revision:       core.CastInt64(data["revision"]),
	}
}

func (p SalesItemMaster) ToDict() map[string]interface{} {

	var salesItemId *string
	if p.SalesItemId != nil {
		salesItemId = p.SalesItemId
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
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
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
		"salesItemId":    salesItemId,
		"name":           name,
		"description":    description,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
		"createdAt":      createdAt,
		"updatedAt":      updatedAt,
		"revision":       revision,
	}
}

func (p SalesItemMaster) Pointer() *SalesItemMaster {
	return &p
}

func CastSalesItemMasters(data []interface{}) []SalesItemMaster {
	v := make([]SalesItemMaster, 0)
	for _, d := range data {
		v = append(v, NewSalesItemMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemMastersFromDict(data []SalesItemMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItemGroupMaster struct {
	SalesItemGroupId *string   `json:"salesItemGroupId"`
	Name             *string   `json:"name"`
	Description      *string   `json:"description"`
	Metadata         *string   `json:"metadata"`
	SalesItemNames   []*string `json:"salesItemNames"`
	CreatedAt        *int64    `json:"createdAt"`
	UpdatedAt        *int64    `json:"updatedAt"`
	Revision         *int64    `json:"revision"`
}

func NewSalesItemGroupMasterFromJson(data string) SalesItemGroupMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSalesItemGroupMasterFromDict(dict)
}

func NewSalesItemGroupMasterFromDict(data map[string]interface{}) SalesItemGroupMaster {
	return SalesItemGroupMaster{
		SalesItemGroupId: core.CastString(data["salesItemGroupId"]),
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		Metadata:         core.CastString(data["metadata"]),
		SalesItemNames:   core.CastStrings(core.CastArray(data["salesItemNames"])),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
	}
}

func (p SalesItemGroupMaster) ToDict() map[string]interface{} {

	var salesItemGroupId *string
	if p.SalesItemGroupId != nil {
		salesItemGroupId = p.SalesItemGroupId
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
	var salesItemNames []interface{}
	if p.SalesItemNames != nil {
		salesItemNames = core.CastStringsFromDict(
			p.SalesItemNames,
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
		"salesItemGroupId": salesItemGroupId,
		"name":             name,
		"description":      description,
		"metadata":         metadata,
		"salesItemNames":   salesItemNames,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
	}
}

func (p SalesItemGroupMaster) Pointer() *SalesItemGroupMaster {
	return &p
}

func CastSalesItemGroupMasters(data []interface{}) []SalesItemGroupMaster {
	v := make([]SalesItemGroupMaster, 0)
	for _, d := range data {
		v = append(v, NewSalesItemGroupMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemGroupMastersFromDict(data []SalesItemGroupMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ShowcaseMaster struct {
	ShowcaseId         *string             `json:"showcaseId"`
	Name               *string             `json:"name"`
	Description        *string             `json:"description"`
	Metadata           *string             `json:"metadata"`
	SalesPeriodEventId *string             `json:"salesPeriodEventId"`
	DisplayItems       []DisplayItemMaster `json:"displayItems"`
	CreatedAt          *int64              `json:"createdAt"`
	UpdatedAt          *int64              `json:"updatedAt"`
	Revision           *int64              `json:"revision"`
}

func NewShowcaseMasterFromJson(data string) ShowcaseMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseMasterFromDict(dict)
}

func NewShowcaseMasterFromDict(data map[string]interface{}) ShowcaseMaster {
	return ShowcaseMaster{
		ShowcaseId:         core.CastString(data["showcaseId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		Metadata:           core.CastString(data["metadata"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		DisplayItems:       CastDisplayItemMasters(core.CastArray(data["displayItems"])),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p ShowcaseMaster) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
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
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastDisplayItemMastersFromDict(
			p.DisplayItems,
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
		"showcaseId":         showcaseId,
		"name":               name,
		"description":        description,
		"metadata":           metadata,
		"salesPeriodEventId": salesPeriodEventId,
		"displayItems":       displayItems,
		"createdAt":          createdAt,
		"updatedAt":          updatedAt,
		"revision":           revision,
	}
}

func (p ShowcaseMaster) Pointer() *ShowcaseMaster {
	return &p
}

func CastShowcaseMasters(data []interface{}) []ShowcaseMaster {
	v := make([]ShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcaseMastersFromDict(data []ShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentShowcaseMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func NewCurrentShowcaseMasterFromJson(data string) CurrentShowcaseMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCurrentShowcaseMasterFromDict(dict)
}

func NewCurrentShowcaseMasterFromDict(data map[string]interface{}) CurrentShowcaseMaster {
	return CurrentShowcaseMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentShowcaseMaster) ToDict() map[string]interface{} {

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

func (p CurrentShowcaseMaster) Pointer() *CurrentShowcaseMaster {
	return &p
}

func CastCurrentShowcaseMasters(data []interface{}) []CurrentShowcaseMaster {
	v := make([]CurrentShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentShowcaseMastersFromDict(data []CurrentShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItem struct {
	Name           *string         `json:"name"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
}

func NewSalesItemFromJson(data string) SalesItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSalesItemFromDict(dict)
}

func NewSalesItemFromDict(data map[string]interface{}) SalesItem {
	return SalesItem{
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
	}
}

func (p SalesItem) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	return map[string]interface{}{
		"name":           name,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
	}
}

func (p SalesItem) Pointer() *SalesItem {
	return &p
}

func CastSalesItems(data []interface{}) []SalesItem {
	v := make([]SalesItem, 0)
	for _, d := range data {
		v = append(v, NewSalesItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemsFromDict(data []SalesItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SalesItemGroup struct {
	Name       *string     `json:"name"`
	Metadata   *string     `json:"metadata"`
	SalesItems []SalesItem `json:"salesItems"`
}

func NewSalesItemGroupFromJson(data string) SalesItemGroup {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSalesItemGroupFromDict(dict)
}

func NewSalesItemGroupFromDict(data map[string]interface{}) SalesItemGroup {
	return SalesItemGroup{
		Name:       core.CastString(data["name"]),
		Metadata:   core.CastString(data["metadata"]),
		SalesItems: CastSalesItems(core.CastArray(data["salesItems"])),
	}
}

func (p SalesItemGroup) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var salesItems []interface{}
	if p.SalesItems != nil {
		salesItems = CastSalesItemsFromDict(
			p.SalesItems,
		)
	}
	return map[string]interface{}{
		"name":       name,
		"metadata":   metadata,
		"salesItems": salesItems,
	}
}

func (p SalesItemGroup) Pointer() *SalesItemGroup {
	return &p
}

func CastSalesItemGroups(data []interface{}) []SalesItemGroup {
	v := make([]SalesItemGroup, 0)
	for _, d := range data {
		v = append(v, NewSalesItemGroupFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSalesItemGroupsFromDict(data []SalesItemGroup) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Showcase struct {
	ShowcaseId         *string       `json:"showcaseId"`
	Name               *string       `json:"name"`
	Metadata           *string       `json:"metadata"`
	SalesPeriodEventId *string       `json:"salesPeriodEventId"`
	DisplayItems       []DisplayItem `json:"displayItems"`
}

func NewShowcaseFromJson(data string) Showcase {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewShowcaseFromDict(dict)
}

func NewShowcaseFromDict(data map[string]interface{}) Showcase {
	return Showcase{
		ShowcaseId:         core.CastString(data["showcaseId"]),
		Name:               core.CastString(data["name"]),
		Metadata:           core.CastString(data["metadata"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		DisplayItems:       CastDisplayItems(core.CastArray(data["displayItems"])),
	}
}

func (p Showcase) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastDisplayItemsFromDict(
			p.DisplayItems,
		)
	}
	return map[string]interface{}{
		"showcaseId":         showcaseId,
		"name":               name,
		"metadata":           metadata,
		"salesPeriodEventId": salesPeriodEventId,
		"displayItems":       displayItems,
	}
}

func (p Showcase) Pointer() *Showcase {
	return &p
}

func CastShowcases(data []interface{}) []Showcase {
	v := make([]Showcase, 0)
	for _, d := range data {
		v = append(v, NewShowcaseFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastShowcasesFromDict(data []Showcase) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DisplayItem struct {
	DisplayItemId      *string         `json:"displayItemId"`
	Type               *string         `json:"type"`
	SalesItem          *SalesItem      `json:"salesItem"`
	SalesItemGroup     *SalesItemGroup `json:"salesItemGroup"`
	SalesPeriodEventId *string         `json:"salesPeriodEventId"`
}

func NewDisplayItemFromJson(data string) DisplayItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDisplayItemFromDict(dict)
}

func NewDisplayItemFromDict(data map[string]interface{}) DisplayItem {
	return DisplayItem{
		DisplayItemId:      core.CastString(data["displayItemId"]),
		Type:               core.CastString(data["type"]),
		SalesItem:          NewSalesItemFromDict(core.CastMap(data["salesItem"])).Pointer(),
		SalesItemGroup:     NewSalesItemGroupFromDict(core.CastMap(data["salesItemGroup"])).Pointer(),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
	}
}

func (p DisplayItem) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var salesItem map[string]interface{}
	if p.SalesItem != nil {
		salesItem = p.SalesItem.ToDict()
	}
	var salesItemGroup map[string]interface{}
	if p.SalesItemGroup != nil {
		salesItemGroup = p.SalesItemGroup.ToDict()
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	return map[string]interface{}{
		"displayItemId":      displayItemId,
		"type":               _type,
		"salesItem":          salesItem,
		"salesItemGroup":     salesItemGroup,
		"salesPeriodEventId": salesPeriodEventId,
	}
}

func (p DisplayItem) Pointer() *DisplayItem {
	return &p
}

func CastDisplayItems(data []interface{}) []DisplayItem {
	v := make([]DisplayItem, 0)
	for _, d := range data {
		v = append(v, NewDisplayItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDisplayItemsFromDict(data []DisplayItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type DisplayItemMaster struct {
	DisplayItemId      *string `json:"displayItemId"`
	Type               *string `json:"type"`
	SalesItemName      *string `json:"salesItemName"`
	SalesItemGroupName *string `json:"salesItemGroupName"`
	SalesPeriodEventId *string `json:"salesPeriodEventId"`
	Revision           *int64  `json:"revision"`
}

func NewDisplayItemMasterFromJson(data string) DisplayItemMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDisplayItemMasterFromDict(dict)
}

func NewDisplayItemMasterFromDict(data map[string]interface{}) DisplayItemMaster {
	return DisplayItemMaster{
		DisplayItemId:      core.CastString(data["displayItemId"]),
		Type:               core.CastString(data["type"]),
		SalesItemName:      core.CastString(data["salesItemName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
		Revision:           core.CastInt64(data["revision"]),
	}
}

func (p DisplayItemMaster) ToDict() map[string]interface{} {

	var displayItemId *string
	if p.DisplayItemId != nil {
		displayItemId = p.DisplayItemId
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var salesItemName *string
	if p.SalesItemName != nil {
		salesItemName = p.SalesItemName
	}
	var salesItemGroupName *string
	if p.SalesItemGroupName != nil {
		salesItemGroupName = p.SalesItemGroupName
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"displayItemId":      displayItemId,
		"type":               _type,
		"salesItemName":      salesItemName,
		"salesItemGroupName": salesItemGroupName,
		"salesPeriodEventId": salesPeriodEventId,
		"revision":           revision,
	}
}

func (p DisplayItemMaster) Pointer() *DisplayItemMaster {
	return &p
}

func CastDisplayItemMasters(data []interface{}) []DisplayItemMaster {
	v := make([]DisplayItemMaster, 0)
	for _, d := range data {
		v = append(v, NewDisplayItemMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDisplayItemMastersFromDict(data []DisplayItemMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomShowcaseMaster struct {
	ShowcaseId            *string                  `json:"showcaseId"`
	Name                  *string                  `json:"name"`
	Description           *string                  `json:"description"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
	CreatedAt             *int64                   `json:"createdAt"`
	UpdatedAt             *int64                   `json:"updatedAt"`
	Revision              *int64                   `json:"revision"`
}

func NewRandomShowcaseMasterFromJson(data string) RandomShowcaseMaster {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseMasterFromDict(dict)
}

func NewRandomShowcaseMasterFromDict(data map[string]interface{}) RandomShowcaseMaster {
	return RandomShowcaseMaster{
		ShowcaseId:            core.CastString(data["showcaseId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
		Revision:              core.CastInt64(data["revision"]),
	}
}

func (p RandomShowcaseMaster) ToDict() map[string]interface{} {

	var showcaseId *string
	if p.ShowcaseId != nil {
		showcaseId = p.ShowcaseId
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
	var maximumNumberOfChoice *int32
	if p.MaximumNumberOfChoice != nil {
		maximumNumberOfChoice = p.MaximumNumberOfChoice
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		)
	}
	var baseTimestamp *int64
	if p.BaseTimestamp != nil {
		baseTimestamp = p.BaseTimestamp
	}
	var resetIntervalHours *int32
	if p.ResetIntervalHours != nil {
		resetIntervalHours = p.ResetIntervalHours
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
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
		"showcaseId":            showcaseId,
		"name":                  name,
		"description":           description,
		"metadata":              metadata,
		"maximumNumberOfChoice": maximumNumberOfChoice,
		"displayItems":          displayItems,
		"baseTimestamp":         baseTimestamp,
		"resetIntervalHours":    resetIntervalHours,
		"salesPeriodEventId":    salesPeriodEventId,
		"createdAt":             createdAt,
		"updatedAt":             updatedAt,
		"revision":              revision,
	}
}

func (p RandomShowcaseMaster) Pointer() *RandomShowcaseMaster {
	return &p
}

func CastRandomShowcaseMasters(data []interface{}) []RandomShowcaseMaster {
	v := make([]RandomShowcaseMaster, 0)
	for _, d := range data {
		v = append(v, NewRandomShowcaseMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomShowcaseMastersFromDict(data []RandomShowcaseMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomShowcase struct {
	RandomShowcaseId      *string                  `json:"randomShowcaseId"`
	Name                  *string                  `json:"name"`
	Metadata              *string                  `json:"metadata"`
	MaximumNumberOfChoice *int32                   `json:"maximumNumberOfChoice"`
	DisplayItems          []RandomDisplayItemModel `json:"displayItems"`
	BaseTimestamp         *int64                   `json:"baseTimestamp"`
	ResetIntervalHours    *int32                   `json:"resetIntervalHours"`
	SalesPeriodEventId    *string                  `json:"salesPeriodEventId"`
}

func NewRandomShowcaseFromJson(data string) RandomShowcase {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseFromDict(dict)
}

func NewRandomShowcaseFromDict(data map[string]interface{}) RandomShowcase {
	return RandomShowcase{
		RandomShowcaseId:      core.CastString(data["randomShowcaseId"]),
		Name:                  core.CastString(data["name"]),
		Metadata:              core.CastString(data["metadata"]),
		MaximumNumberOfChoice: core.CastInt32(data["maximumNumberOfChoice"]),
		DisplayItems:          CastRandomDisplayItemModels(core.CastArray(data["displayItems"])),
		BaseTimestamp:         core.CastInt64(data["baseTimestamp"]),
		ResetIntervalHours:    core.CastInt32(data["resetIntervalHours"]),
		SalesPeriodEventId:    core.CastString(data["salesPeriodEventId"]),
	}
}

func (p RandomShowcase) ToDict() map[string]interface{} {

	var randomShowcaseId *string
	if p.RandomShowcaseId != nil {
		randomShowcaseId = p.RandomShowcaseId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var maximumNumberOfChoice *int32
	if p.MaximumNumberOfChoice != nil {
		maximumNumberOfChoice = p.MaximumNumberOfChoice
	}
	var displayItems []interface{}
	if p.DisplayItems != nil {
		displayItems = CastRandomDisplayItemModelsFromDict(
			p.DisplayItems,
		)
	}
	var baseTimestamp *int64
	if p.BaseTimestamp != nil {
		baseTimestamp = p.BaseTimestamp
	}
	var resetIntervalHours *int32
	if p.ResetIntervalHours != nil {
		resetIntervalHours = p.ResetIntervalHours
	}
	var salesPeriodEventId *string
	if p.SalesPeriodEventId != nil {
		salesPeriodEventId = p.SalesPeriodEventId
	}
	return map[string]interface{}{
		"randomShowcaseId":      randomShowcaseId,
		"name":                  name,
		"metadata":              metadata,
		"maximumNumberOfChoice": maximumNumberOfChoice,
		"displayItems":          displayItems,
		"baseTimestamp":         baseTimestamp,
		"resetIntervalHours":    resetIntervalHours,
		"salesPeriodEventId":    salesPeriodEventId,
	}
}

func (p RandomShowcase) Pointer() *RandomShowcase {
	return &p
}

func CastRandomShowcases(data []interface{}) []RandomShowcase {
	v := make([]RandomShowcase, 0)
	for _, d := range data {
		v = append(v, NewRandomShowcaseFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomShowcasesFromDict(data []RandomShowcase) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type PurchaseCount struct {
	Name  *string `json:"name"`
	Count *int32  `json:"count"`
}

func NewPurchaseCountFromJson(data string) PurchaseCount {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPurchaseCountFromDict(dict)
}

func NewPurchaseCountFromDict(data map[string]interface{}) PurchaseCount {
	return PurchaseCount{
		Name:  core.CastString(data["name"]),
		Count: core.CastInt32(data["count"]),
	}
}

func (p PurchaseCount) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var count *int32
	if p.Count != nil {
		count = p.Count
	}
	return map[string]interface{}{
		"name":  name,
		"count": count,
	}
}

func (p PurchaseCount) Pointer() *PurchaseCount {
	return &p
}

func CastPurchaseCounts(data []interface{}) []PurchaseCount {
	v := make([]PurchaseCount, 0)
	for _, d := range data {
		v = append(v, NewPurchaseCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPurchaseCountsFromDict(data []PurchaseCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomDisplayItem struct {
	ShowcaseName         *string         `json:"showcaseName"`
	Name                 *string         `json:"name"`
	Metadata             *string         `json:"metadata"`
	ConsumeActions       []ConsumeAction `json:"consumeActions"`
	AcquireActions       []AcquireAction `json:"acquireActions"`
	CurrentPurchaseCount *int32          `json:"currentPurchaseCount"`
	MaximumPurchaseCount *int32          `json:"maximumPurchaseCount"`
}

func NewRandomDisplayItemFromJson(data string) RandomDisplayItem {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomDisplayItemFromDict(dict)
}

func NewRandomDisplayItemFromDict(data map[string]interface{}) RandomDisplayItem {
	return RandomDisplayItem{
		ShowcaseName:         core.CastString(data["showcaseName"]),
		Name:                 core.CastString(data["name"]),
		Metadata:             core.CastString(data["metadata"]),
		ConsumeActions:       CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions:       CastAcquireActions(core.CastArray(data["acquireActions"])),
		CurrentPurchaseCount: core.CastInt32(data["currentPurchaseCount"]),
		MaximumPurchaseCount: core.CastInt32(data["maximumPurchaseCount"]),
	}
}

func (p RandomDisplayItem) ToDict() map[string]interface{} {

	var showcaseName *string
	if p.ShowcaseName != nil {
		showcaseName = p.ShowcaseName
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var currentPurchaseCount *int32
	if p.CurrentPurchaseCount != nil {
		currentPurchaseCount = p.CurrentPurchaseCount
	}
	var maximumPurchaseCount *int32
	if p.MaximumPurchaseCount != nil {
		maximumPurchaseCount = p.MaximumPurchaseCount
	}
	return map[string]interface{}{
		"showcaseName":         showcaseName,
		"name":                 name,
		"metadata":             metadata,
		"consumeActions":       consumeActions,
		"acquireActions":       acquireActions,
		"currentPurchaseCount": currentPurchaseCount,
		"maximumPurchaseCount": maximumPurchaseCount,
	}
}

func (p RandomDisplayItem) Pointer() *RandomDisplayItem {
	return &p
}

func CastRandomDisplayItems(data []interface{}) []RandomDisplayItem {
	v := make([]RandomDisplayItem, 0)
	for _, d := range data {
		v = append(v, NewRandomDisplayItemFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomDisplayItemsFromDict(data []RandomDisplayItem) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomDisplayItemModel struct {
	Name           *string         `json:"name"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
	Stock          *int32          `json:"stock"`
	Weight         *int32          `json:"weight"`
}

func NewRandomDisplayItemModelFromJson(data string) RandomDisplayItemModel {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomDisplayItemModelFromDict(dict)
}

func NewRandomDisplayItemModelFromDict(data map[string]interface{}) RandomDisplayItemModel {
	return RandomDisplayItemModel{
		Name:           core.CastString(data["name"]),
		Metadata:       core.CastString(data["metadata"]),
		ConsumeActions: CastConsumeActions(core.CastArray(data["consumeActions"])),
		AcquireActions: CastAcquireActions(core.CastArray(data["acquireActions"])),
		Stock:          core.CastInt32(data["stock"]),
		Weight:         core.CastInt32(data["weight"]),
	}
}

func (p RandomDisplayItemModel) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var consumeActions []interface{}
	if p.ConsumeActions != nil {
		consumeActions = CastConsumeActionsFromDict(
			p.ConsumeActions,
		)
	}
	var acquireActions []interface{}
	if p.AcquireActions != nil {
		acquireActions = CastAcquireActionsFromDict(
			p.AcquireActions,
		)
	}
	var stock *int32
	if p.Stock != nil {
		stock = p.Stock
	}
	var weight *int32
	if p.Weight != nil {
		weight = p.Weight
	}
	return map[string]interface{}{
		"name":           name,
		"metadata":       metadata,
		"consumeActions": consumeActions,
		"acquireActions": acquireActions,
		"stock":          stock,
		"weight":         weight,
	}
}

func (p RandomDisplayItemModel) Pointer() *RandomDisplayItemModel {
	return &p
}

func CastRandomDisplayItemModels(data []interface{}) []RandomDisplayItemModel {
	v := make([]RandomDisplayItemModel, 0)
	for _, d := range data {
		v = append(v, NewRandomDisplayItemModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomDisplayItemModelsFromDict(data []RandomDisplayItemModel) []interface{} {
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
