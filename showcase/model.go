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

import "core"

type Namespace struct {
	NamespaceId      *string     `json:"namespaceId"`
	Name             *string     `json:"name"`
	QueueNamespaceId *string     `json:"queueNamespaceId"`
	KeyId            *string     `json:"keyId"`
	Description      *string     `json:"description"`
	LogSetting       *LogSetting `json:"logSetting"`
	CreatedAt        *int64      `json:"createdAt"`
	UpdatedAt        *int64      `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:      core.CastString(data["namespaceId"]),
		Name:             core.CastString(data["name"]),
		QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
		KeyId:            core.CastString(data["keyId"]),
		Description:      core.CastString(data["description"]),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":      p.NamespaceId,
		"name":             p.Name,
		"queueNamespaceId": p.QueueNamespaceId,
		"keyId":            p.KeyId,
		"description":      p.Description,
		"logSetting":       p.LogSetting.ToDict(),
		"createdAt":        p.CreatedAt,
		"updatedAt":        p.UpdatedAt,
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
	}
}

func (p SalesItemMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"salesItemId": p.SalesItemId,
		"name":        p.Name,
		"description": p.Description,
		"metadata":    p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
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
	SalesItemGroupId *string  `json:"salesItemGroupId"`
	Name             *string  `json:"name"`
	Description      *string  `json:"description"`
	Metadata         *string  `json:"metadata"`
	SalesItemNames   []string `json:"salesItemNames"`
	CreatedAt        *int64   `json:"createdAt"`
	UpdatedAt        *int64   `json:"updatedAt"`
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
	}
}

func (p SalesItemGroupMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"salesItemGroupId": p.SalesItemGroupId,
		"name":             p.Name,
		"description":      p.Description,
		"metadata":         p.Metadata,
		"salesItemNames": core.CastStringsFromDict(
			p.SalesItemNames,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
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
	}
}

func (p ShowcaseMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"showcaseId":         p.ShowcaseId,
		"name":               p.Name,
		"description":        p.Description,
		"metadata":           p.Metadata,
		"salesPeriodEventId": p.SalesPeriodEventId,
		"displayItems": CastDisplayItemMastersFromDict(
			p.DisplayItems,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
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

func NewCurrentShowcaseMasterFromDict(data map[string]interface{}) CurrentShowcaseMaster {
	return CurrentShowcaseMaster{
		NamespaceId: core.CastString(data["namespaceId"]),
		Settings:    core.CastString(data["settings"]),
	}
}

func (p CurrentShowcaseMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId": p.NamespaceId,
		"settings":    p.Settings,
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

type SalesItem struct {
	Name           *string         `json:"name"`
	Metadata       *string         `json:"metadata"`
	ConsumeActions []ConsumeAction `json:"consumeActions"`
	AcquireActions []AcquireAction `json:"acquireActions"`
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
	return map[string]interface{}{
		"name":     p.Name,
		"metadata": p.Metadata,
		"consumeActions": CastConsumeActionsFromDict(
			p.ConsumeActions,
		),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
		),
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

func NewSalesItemGroupFromDict(data map[string]interface{}) SalesItemGroup {
	return SalesItemGroup{
		Name:       core.CastString(data["name"]),
		Metadata:   core.CastString(data["metadata"]),
		SalesItems: CastSalesItems(core.CastArray(data["salesItems"])),
	}
}

func (p SalesItemGroup) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"metadata": p.Metadata,
		"salesItems": CastSalesItemsFromDict(
			p.SalesItems,
		),
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
	return map[string]interface{}{
		"showcaseId":         p.ShowcaseId,
		"name":               p.Name,
		"metadata":           p.Metadata,
		"salesPeriodEventId": p.SalesPeriodEventId,
		"displayItems": CastDisplayItemsFromDict(
			p.DisplayItems,
		),
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
	return map[string]interface{}{
		"displayItemId":      p.DisplayItemId,
		"type":               p.Type,
		"salesItem":          p.SalesItem.ToDict(),
		"salesItemGroup":     p.SalesItemGroup.ToDict(),
		"salesPeriodEventId": p.SalesPeriodEventId,
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

type ConsumeAction struct {
	Action  *string `json:"action"`
	Request *string `json:"request"`
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

type DisplayItemMaster struct {
	DisplayItemId      *string `json:"displayItemId"`
	Type               *string `json:"type"`
	SalesItemName      *string `json:"salesItemName"`
	SalesItemGroupName *string `json:"salesItemGroupName"`
	SalesPeriodEventId *string `json:"salesPeriodEventId"`
}

func NewDisplayItemMasterFromDict(data map[string]interface{}) DisplayItemMaster {
	return DisplayItemMaster{
		DisplayItemId:      core.CastString(data["displayItemId"]),
		Type:               core.CastString(data["type"]),
		SalesItemName:      core.CastString(data["salesItemName"]),
		SalesItemGroupName: core.CastString(data["salesItemGroupName"]),
		SalesPeriodEventId: core.CastString(data["salesPeriodEventId"]),
	}
}

func (p DisplayItemMaster) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"displayItemId":      p.DisplayItemId,
		"type":               p.Type,
		"salesItemName":      p.SalesItemName,
		"salesItemGroupName": p.SalesItemGroupName,
		"salesPeriodEventId": p.SalesPeriodEventId,
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
