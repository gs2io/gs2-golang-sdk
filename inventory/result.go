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

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromJson(data string) DescribeNamespacesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesResultFromDict(dict)
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
	return DescribeNamespacesResult{
		Items:         CastNamespaces(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
	return &p
}

type CreateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceResultFromDict(dict)
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
	return CreateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
	return &p
}

type GetNamespaceStatusResult struct {
	Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusResultFromDict(dict)
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
	return GetNamespaceStatusResult{
		Status: core.CastString(data["status"]),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
	}
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
	return &p
}

type GetNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceResultFromDict(dict)
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
	return GetNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
	return &p
}

type UpdateNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceResultFromDict(dict)
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
	return UpdateNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
	return &p
}

type DeleteNamespaceResult struct {
	Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceResultFromDict(dict)
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DescribeInventoryModelMastersResult struct {
	Items         []InventoryModelMaster `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
}

type DescribeInventoryModelMastersAsyncResult struct {
	result *DescribeInventoryModelMastersResult
	err    error
}

func NewDescribeInventoryModelMastersResultFromJson(data string) DescribeInventoryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryModelMastersResultFromDict(dict)
}

func NewDescribeInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeInventoryModelMastersResult {
	return DescribeInventoryModelMastersResult{
		Items:         CastInventoryModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeInventoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeInventoryModelMastersResult) Pointer() *DescribeInventoryModelMastersResult {
	return &p
}

type CreateInventoryModelMasterResult struct {
	Item *InventoryModelMaster `json:"item"`
}

type CreateInventoryModelMasterAsyncResult struct {
	result *CreateInventoryModelMasterResult
	err    error
}

func NewCreateInventoryModelMasterResultFromJson(data string) CreateInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateInventoryModelMasterResultFromDict(dict)
}

func NewCreateInventoryModelMasterResultFromDict(data map[string]interface{}) CreateInventoryModelMasterResult {
	return CreateInventoryModelMasterResult{
		Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateInventoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateInventoryModelMasterResult) Pointer() *CreateInventoryModelMasterResult {
	return &p
}

type GetInventoryModelMasterResult struct {
	Item *InventoryModelMaster `json:"item"`
}

type GetInventoryModelMasterAsyncResult struct {
	result *GetInventoryModelMasterResult
	err    error
}

func NewGetInventoryModelMasterResultFromJson(data string) GetInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryModelMasterResultFromDict(dict)
}

func NewGetInventoryModelMasterResultFromDict(data map[string]interface{}) GetInventoryModelMasterResult {
	return GetInventoryModelMasterResult{
		Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInventoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInventoryModelMasterResult) Pointer() *GetInventoryModelMasterResult {
	return &p
}

type UpdateInventoryModelMasterResult struct {
	Item *InventoryModelMaster `json:"item"`
}

type UpdateInventoryModelMasterAsyncResult struct {
	result *UpdateInventoryModelMasterResult
	err    error
}

func NewUpdateInventoryModelMasterResultFromJson(data string) UpdateInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateInventoryModelMasterResultFromDict(dict)
}

func NewUpdateInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateInventoryModelMasterResult {
	return UpdateInventoryModelMasterResult{
		Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateInventoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateInventoryModelMasterResult) Pointer() *UpdateInventoryModelMasterResult {
	return &p
}

type DeleteInventoryModelMasterResult struct {
	Item *InventoryModelMaster `json:"item"`
}

type DeleteInventoryModelMasterAsyncResult struct {
	result *DeleteInventoryModelMasterResult
	err    error
}

func NewDeleteInventoryModelMasterResultFromJson(data string) DeleteInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInventoryModelMasterResultFromDict(dict)
}

func NewDeleteInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteInventoryModelMasterResult {
	return DeleteInventoryModelMasterResult{
		Item: NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteInventoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteInventoryModelMasterResult) Pointer() *DeleteInventoryModelMasterResult {
	return &p
}

type DescribeInventoryModelsResult struct {
	Items []InventoryModel `json:"items"`
}

type DescribeInventoryModelsAsyncResult struct {
	result *DescribeInventoryModelsResult
	err    error
}

func NewDescribeInventoryModelsResultFromJson(data string) DescribeInventoryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoryModelsResultFromDict(dict)
}

func NewDescribeInventoryModelsResultFromDict(data map[string]interface{}) DescribeInventoryModelsResult {
	return DescribeInventoryModelsResult{
		Items: CastInventoryModels(core.CastArray(data["items"])),
	}
}

func (p DescribeInventoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeInventoryModelsResult) Pointer() *DescribeInventoryModelsResult {
	return &p
}

type GetInventoryModelResult struct {
	Item *InventoryModel `json:"item"`
}

type GetInventoryModelAsyncResult struct {
	result *GetInventoryModelResult
	err    error
}

func NewGetInventoryModelResultFromJson(data string) GetInventoryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryModelResultFromDict(dict)
}

func NewGetInventoryModelResultFromDict(data map[string]interface{}) GetInventoryModelResult {
	return GetInventoryModelResult{
		Item: NewInventoryModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInventoryModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInventoryModelResult) Pointer() *GetInventoryModelResult {
	return &p
}

type DescribeItemModelMastersResult struct {
	Items         []ItemModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeItemModelMastersAsyncResult struct {
	result *DescribeItemModelMastersResult
	err    error
}

func NewDescribeItemModelMastersResultFromJson(data string) DescribeItemModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemModelMastersResultFromDict(dict)
}

func NewDescribeItemModelMastersResultFromDict(data map[string]interface{}) DescribeItemModelMastersResult {
	return DescribeItemModelMastersResult{
		Items:         CastItemModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeItemModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeItemModelMastersResult) Pointer() *DescribeItemModelMastersResult {
	return &p
}

type CreateItemModelMasterResult struct {
	Item *ItemModelMaster `json:"item"`
}

type CreateItemModelMasterAsyncResult struct {
	result *CreateItemModelMasterResult
	err    error
}

func NewCreateItemModelMasterResultFromJson(data string) CreateItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateItemModelMasterResultFromDict(dict)
}

func NewCreateItemModelMasterResultFromDict(data map[string]interface{}) CreateItemModelMasterResult {
	return CreateItemModelMasterResult{
		Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateItemModelMasterResult) Pointer() *CreateItemModelMasterResult {
	return &p
}

type GetItemModelMasterResult struct {
	Item *ItemModelMaster `json:"item"`
}

type GetItemModelMasterAsyncResult struct {
	result *GetItemModelMasterResult
	err    error
}

func NewGetItemModelMasterResultFromJson(data string) GetItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemModelMasterResultFromDict(dict)
}

func NewGetItemModelMasterResultFromDict(data map[string]interface{}) GetItemModelMasterResult {
	return GetItemModelMasterResult{
		Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetItemModelMasterResult) Pointer() *GetItemModelMasterResult {
	return &p
}

type UpdateItemModelMasterResult struct {
	Item *ItemModelMaster `json:"item"`
}

type UpdateItemModelMasterAsyncResult struct {
	result *UpdateItemModelMasterResult
	err    error
}

func NewUpdateItemModelMasterResultFromJson(data string) UpdateItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateItemModelMasterResultFromDict(dict)
}

func NewUpdateItemModelMasterResultFromDict(data map[string]interface{}) UpdateItemModelMasterResult {
	return UpdateItemModelMasterResult{
		Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateItemModelMasterResult) Pointer() *UpdateItemModelMasterResult {
	return &p
}

type DeleteItemModelMasterResult struct {
	Item *ItemModelMaster `json:"item"`
}

type DeleteItemModelMasterAsyncResult struct {
	result *DeleteItemModelMasterResult
	err    error
}

func NewDeleteItemModelMasterResultFromJson(data string) DeleteItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteItemModelMasterResultFromDict(dict)
}

func NewDeleteItemModelMasterResultFromDict(data map[string]interface{}) DeleteItemModelMasterResult {
	return DeleteItemModelMasterResult{
		Item: NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteItemModelMasterResult) Pointer() *DeleteItemModelMasterResult {
	return &p
}

type DescribeItemModelsResult struct {
	Items []ItemModel `json:"items"`
}

type DescribeItemModelsAsyncResult struct {
	result *DescribeItemModelsResult
	err    error
}

func NewDescribeItemModelsResultFromJson(data string) DescribeItemModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemModelsResultFromDict(dict)
}

func NewDescribeItemModelsResultFromDict(data map[string]interface{}) DescribeItemModelsResult {
	return DescribeItemModelsResult{
		Items: CastItemModels(core.CastArray(data["items"])),
	}
}

func (p DescribeItemModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeItemModelsResult) Pointer() *DescribeItemModelsResult {
	return &p
}

type GetItemModelResult struct {
	Item *ItemModel `json:"item"`
}

type GetItemModelAsyncResult struct {
	result *GetItemModelResult
	err    error
}

func NewGetItemModelResultFromJson(data string) GetItemModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemModelResultFromDict(dict)
}

func NewGetItemModelResultFromDict(data map[string]interface{}) GetItemModelResult {
	return GetItemModelResult{
		Item: NewItemModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetItemModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetItemModelResult) Pointer() *GetItemModelResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentItemModelMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromJson(data string) ExportMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExportMasterResultFromDict(dict)
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentItemModelMasterResult struct {
	Item *CurrentItemModelMaster `json:"item"`
}

type GetCurrentItemModelMasterAsyncResult struct {
	result *GetCurrentItemModelMasterResult
	err    error
}

func NewGetCurrentItemModelMasterResultFromJson(data string) GetCurrentItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentItemModelMasterResultFromDict(dict)
}

func NewGetCurrentItemModelMasterResultFromDict(data map[string]interface{}) GetCurrentItemModelMasterResult {
	return GetCurrentItemModelMasterResult{
		Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentItemModelMasterResult) Pointer() *GetCurrentItemModelMasterResult {
	return &p
}

type UpdateCurrentItemModelMasterResult struct {
	Item *CurrentItemModelMaster `json:"item"`
}

type UpdateCurrentItemModelMasterAsyncResult struct {
	result *UpdateCurrentItemModelMasterResult
	err    error
}

func NewUpdateCurrentItemModelMasterResultFromJson(data string) UpdateCurrentItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentItemModelMasterResultFromDict(dict)
}

func NewUpdateCurrentItemModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterResult {
	return UpdateCurrentItemModelMasterResult{
		Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentItemModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentItemModelMasterResult) Pointer() *UpdateCurrentItemModelMasterResult {
	return &p
}

type UpdateCurrentItemModelMasterFromGitHubResult struct {
	Item *CurrentItemModelMaster `json:"item"`
}

type UpdateCurrentItemModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentItemModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentItemModelMasterFromGitHubResultFromJson(data string) UpdateCurrentItemModelMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentItemModelMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentItemModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentItemModelMasterFromGitHubResult {
	return UpdateCurrentItemModelMasterFromGitHubResult{
		Item: NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentItemModelMasterFromGitHubResult) Pointer() *UpdateCurrentItemModelMasterFromGitHubResult {
	return &p
}

type DescribeInventoriesResult struct {
	Items         []Inventory `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeInventoriesAsyncResult struct {
	result *DescribeInventoriesResult
	err    error
}

func NewDescribeInventoriesResultFromJson(data string) DescribeInventoriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoriesResultFromDict(dict)
}

func NewDescribeInventoriesResultFromDict(data map[string]interface{}) DescribeInventoriesResult {
	return DescribeInventoriesResult{
		Items:         CastInventories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeInventoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeInventoriesResult) Pointer() *DescribeInventoriesResult {
	return &p
}

type DescribeInventoriesByUserIdResult struct {
	Items         []Inventory `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeInventoriesByUserIdAsyncResult struct {
	result *DescribeInventoriesByUserIdResult
	err    error
}

func NewDescribeInventoriesByUserIdResultFromJson(data string) DescribeInventoriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeInventoriesByUserIdResultFromDict(dict)
}

func NewDescribeInventoriesByUserIdResultFromDict(data map[string]interface{}) DescribeInventoriesByUserIdResult {
	return DescribeInventoriesByUserIdResult{
		Items:         CastInventories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeInventoriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeInventoriesByUserIdResult) Pointer() *DescribeInventoriesByUserIdResult {
	return &p
}

type GetInventoryResult struct {
	Item *Inventory `json:"item"`
}

type GetInventoryAsyncResult struct {
	result *GetInventoryResult
	err    error
}

func NewGetInventoryResultFromJson(data string) GetInventoryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryResultFromDict(dict)
}

func NewGetInventoryResultFromDict(data map[string]interface{}) GetInventoryResult {
	return GetInventoryResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInventoryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInventoryResult) Pointer() *GetInventoryResult {
	return &p
}

type GetInventoryByUserIdResult struct {
	Item *Inventory `json:"item"`
}

type GetInventoryByUserIdAsyncResult struct {
	result *GetInventoryByUserIdResult
	err    error
}

func NewGetInventoryByUserIdResultFromJson(data string) GetInventoryByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetInventoryByUserIdResultFromDict(dict)
}

func NewGetInventoryByUserIdResultFromDict(data map[string]interface{}) GetInventoryByUserIdResult {
	return GetInventoryByUserIdResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetInventoryByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetInventoryByUserIdResult) Pointer() *GetInventoryByUserIdResult {
	return &p
}

type AddCapacityByUserIdResult struct {
	Item *Inventory `json:"item"`
}

type AddCapacityByUserIdAsyncResult struct {
	result *AddCapacityByUserIdResult
	err    error
}

func NewAddCapacityByUserIdResultFromJson(data string) AddCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddCapacityByUserIdResultFromDict(dict)
}

func NewAddCapacityByUserIdResultFromDict(data map[string]interface{}) AddCapacityByUserIdResult {
	return AddCapacityByUserIdResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p AddCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p AddCapacityByUserIdResult) Pointer() *AddCapacityByUserIdResult {
	return &p
}

type SetCapacityByUserIdResult struct {
	Item *Inventory `json:"item"`
}

type SetCapacityByUserIdAsyncResult struct {
	result *SetCapacityByUserIdResult
	err    error
}

func NewSetCapacityByUserIdResultFromJson(data string) SetCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetCapacityByUserIdResultFromDict(dict)
}

func NewSetCapacityByUserIdResultFromDict(data map[string]interface{}) SetCapacityByUserIdResult {
	return SetCapacityByUserIdResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SetCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p SetCapacityByUserIdResult) Pointer() *SetCapacityByUserIdResult {
	return &p
}

type DeleteInventoryByUserIdResult struct {
	Item *Inventory `json:"item"`
}

type DeleteInventoryByUserIdAsyncResult struct {
	result *DeleteInventoryByUserIdResult
	err    error
}

func NewDeleteInventoryByUserIdResultFromJson(data string) DeleteInventoryByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteInventoryByUserIdResultFromDict(dict)
}

func NewDeleteInventoryByUserIdResultFromDict(data map[string]interface{}) DeleteInventoryByUserIdResult {
	return DeleteInventoryByUserIdResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteInventoryByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteInventoryByUserIdResult) Pointer() *DeleteInventoryByUserIdResult {
	return &p
}

type AddCapacityByStampSheetResult struct {
	Item *Inventory `json:"item"`
}

type AddCapacityByStampSheetAsyncResult struct {
	result *AddCapacityByStampSheetResult
	err    error
}

func NewAddCapacityByStampSheetResultFromJson(data string) AddCapacityByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddCapacityByStampSheetResultFromDict(dict)
}

func NewAddCapacityByStampSheetResultFromDict(data map[string]interface{}) AddCapacityByStampSheetResult {
	return AddCapacityByStampSheetResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p AddCapacityByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p AddCapacityByStampSheetResult) Pointer() *AddCapacityByStampSheetResult {
	return &p
}

type SetCapacityByStampSheetResult struct {
	Item *Inventory `json:"item"`
}

type SetCapacityByStampSheetAsyncResult struct {
	result *SetCapacityByStampSheetResult
	err    error
}

func NewSetCapacityByStampSheetResultFromJson(data string) SetCapacityByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetCapacityByStampSheetResultFromDict(dict)
}

func NewSetCapacityByStampSheetResultFromDict(data map[string]interface{}) SetCapacityByStampSheetResult {
	return SetCapacityByStampSheetResult{
		Item: NewInventoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SetCapacityByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p SetCapacityByStampSheetResult) Pointer() *SetCapacityByStampSheetResult {
	return &p
}

type DescribeItemSetsResult struct {
	Items         []ItemSet `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeItemSetsAsyncResult struct {
	result *DescribeItemSetsResult
	err    error
}

func NewDescribeItemSetsResultFromJson(data string) DescribeItemSetsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemSetsResultFromDict(dict)
}

func NewDescribeItemSetsResultFromDict(data map[string]interface{}) DescribeItemSetsResult {
	return DescribeItemSetsResult{
		Items:         CastItemSets(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeItemSetsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeItemSetsResult) Pointer() *DescribeItemSetsResult {
	return &p
}

type DescribeItemSetsByUserIdResult struct {
	Items         []ItemSet `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeItemSetsByUserIdAsyncResult struct {
	result *DescribeItemSetsByUserIdResult
	err    error
}

func NewDescribeItemSetsByUserIdResultFromJson(data string) DescribeItemSetsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeItemSetsByUserIdResultFromDict(dict)
}

func NewDescribeItemSetsByUserIdResultFromDict(data map[string]interface{}) DescribeItemSetsByUserIdResult {
	return DescribeItemSetsByUserIdResult{
		Items:         CastItemSets(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeItemSetsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeItemSetsByUserIdResult) Pointer() *DescribeItemSetsByUserIdResult {
	return &p
}

type GetItemSetResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type GetItemSetAsyncResult struct {
	result *GetItemSetResult
	err    error
}

func NewGetItemSetResultFromJson(data string) GetItemSetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemSetResultFromDict(dict)
}

func NewGetItemSetResultFromDict(data map[string]interface{}) GetItemSetResult {
	return GetItemSetResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p GetItemSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p GetItemSetResult) Pointer() *GetItemSetResult {
	return &p
}

type GetItemSetByUserIdResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type GetItemSetByUserIdAsyncResult struct {
	result *GetItemSetByUserIdResult
	err    error
}

func NewGetItemSetByUserIdResultFromJson(data string) GetItemSetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemSetByUserIdResultFromDict(dict)
}

func NewGetItemSetByUserIdResultFromDict(data map[string]interface{}) GetItemSetByUserIdResult {
	return GetItemSetByUserIdResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p GetItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p GetItemSetByUserIdResult) Pointer() *GetItemSetByUserIdResult {
	return &p
}

type GetItemWithSignatureResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
	Body      *string    `json:"body"`
	Signature *string    `json:"signature"`
}

type GetItemWithSignatureAsyncResult struct {
	result *GetItemWithSignatureResult
	err    error
}

func NewGetItemWithSignatureResultFromJson(data string) GetItemWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemWithSignatureResultFromDict(dict)
}

func NewGetItemWithSignatureResultFromDict(data map[string]interface{}) GetItemWithSignatureResult {
	return GetItemWithSignatureResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetItemWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetItemWithSignatureResult) Pointer() *GetItemWithSignatureResult {
	return &p
}

type GetItemWithSignatureByUserIdResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
	Body      *string    `json:"body"`
	Signature *string    `json:"signature"`
}

type GetItemWithSignatureByUserIdAsyncResult struct {
	result *GetItemWithSignatureByUserIdResult
	err    error
}

func NewGetItemWithSignatureByUserIdResultFromJson(data string) GetItemWithSignatureByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetItemWithSignatureByUserIdResultFromDict(dict)
}

func NewGetItemWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetItemWithSignatureByUserIdResult {
	return GetItemWithSignatureByUserIdResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetItemWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetItemWithSignatureByUserIdResult) Pointer() *GetItemWithSignatureByUserIdResult {
	return &p
}

type AcquireItemSetByUserIdResult struct {
	Items         []ItemSet  `json:"items"`
	ItemModel     *ItemModel `json:"itemModel"`
	Inventory     *Inventory `json:"inventory"`
	OverflowCount *int64     `json:"overflowCount"`
}

type AcquireItemSetByUserIdAsyncResult struct {
	result *AcquireItemSetByUserIdResult
	err    error
}

func NewAcquireItemSetByUserIdResultFromJson(data string) AcquireItemSetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetByUserIdResultFromDict(dict)
}

func NewAcquireItemSetByUserIdResultFromDict(data map[string]interface{}) AcquireItemSetByUserIdResult {
	return AcquireItemSetByUserIdResult{
		Items:         CastItemSets(core.CastArray(data["items"])),
		ItemModel:     NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory:     NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		OverflowCount: core.CastInt64(data["overflowCount"]),
	}
}

func (p AcquireItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel":     p.ItemModel.ToDict(),
		"inventory":     p.Inventory.ToDict(),
		"overflowCount": p.OverflowCount,
	}
}

func (p AcquireItemSetByUserIdResult) Pointer() *AcquireItemSetByUserIdResult {
	return &p
}

type ConsumeItemSetResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type ConsumeItemSetAsyncResult struct {
	result *ConsumeItemSetResult
	err    error
}

func NewConsumeItemSetResultFromJson(data string) ConsumeItemSetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetResultFromDict(dict)
}

func NewConsumeItemSetResultFromDict(data map[string]interface{}) ConsumeItemSetResult {
	return ConsumeItemSetResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p ConsumeItemSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p ConsumeItemSetResult) Pointer() *ConsumeItemSetResult {
	return &p
}

type ConsumeItemSetByUserIdResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type ConsumeItemSetByUserIdAsyncResult struct {
	result *ConsumeItemSetByUserIdResult
	err    error
}

func NewConsumeItemSetByUserIdResultFromJson(data string) ConsumeItemSetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetByUserIdResultFromDict(dict)
}

func NewConsumeItemSetByUserIdResultFromDict(data map[string]interface{}) ConsumeItemSetByUserIdResult {
	return ConsumeItemSetByUserIdResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p ConsumeItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p ConsumeItemSetByUserIdResult) Pointer() *ConsumeItemSetByUserIdResult {
	return &p
}

type DescribeReferenceOfResult struct {
	Items     []string   `json:"items"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DescribeReferenceOfAsyncResult struct {
	result *DescribeReferenceOfResult
	err    error
}

func NewDescribeReferenceOfResultFromJson(data string) DescribeReferenceOfResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReferenceOfResultFromDict(dict)
}

func NewDescribeReferenceOfResultFromDict(data map[string]interface{}) DescribeReferenceOfResult {
	return DescribeReferenceOfResult{
		Items:     core.CastStrings(core.CastArray(data["items"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DescribeReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
			p.Items,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DescribeReferenceOfResult) Pointer() *DescribeReferenceOfResult {
	return &p
}

type DescribeReferenceOfByUserIdResult struct {
	Items     []string   `json:"items"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DescribeReferenceOfByUserIdAsyncResult struct {
	result *DescribeReferenceOfByUserIdResult
	err    error
}

func NewDescribeReferenceOfByUserIdResultFromJson(data string) DescribeReferenceOfByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReferenceOfByUserIdResultFromDict(dict)
}

func NewDescribeReferenceOfByUserIdResultFromDict(data map[string]interface{}) DescribeReferenceOfByUserIdResult {
	return DescribeReferenceOfByUserIdResult{
		Items:     core.CastStrings(core.CastArray(data["items"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DescribeReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
			p.Items,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DescribeReferenceOfByUserIdResult) Pointer() *DescribeReferenceOfByUserIdResult {
	return &p
}

type GetReferenceOfResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type GetReferenceOfAsyncResult struct {
	result *GetReferenceOfResult
	err    error
}

func NewGetReferenceOfResultFromJson(data string) GetReferenceOfResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReferenceOfResultFromDict(dict)
}

func NewGetReferenceOfResultFromDict(data map[string]interface{}) GetReferenceOfResult {
	return GetReferenceOfResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p GetReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p GetReferenceOfResult) Pointer() *GetReferenceOfResult {
	return &p
}

type GetReferenceOfByUserIdResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type GetReferenceOfByUserIdAsyncResult struct {
	result *GetReferenceOfByUserIdResult
	err    error
}

func NewGetReferenceOfByUserIdResultFromJson(data string) GetReferenceOfByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReferenceOfByUserIdResultFromDict(dict)
}

func NewGetReferenceOfByUserIdResultFromDict(data map[string]interface{}) GetReferenceOfByUserIdResult {
	return GetReferenceOfByUserIdResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p GetReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p GetReferenceOfByUserIdResult) Pointer() *GetReferenceOfByUserIdResult {
	return &p
}

type VerifyReferenceOfResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type VerifyReferenceOfAsyncResult struct {
	result *VerifyReferenceOfResult
	err    error
}

func NewVerifyReferenceOfResultFromJson(data string) VerifyReferenceOfResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfResultFromDict(dict)
}

func NewVerifyReferenceOfResultFromDict(data map[string]interface{}) VerifyReferenceOfResult {
	return VerifyReferenceOfResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p VerifyReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p VerifyReferenceOfResult) Pointer() *VerifyReferenceOfResult {
	return &p
}

type VerifyReferenceOfByUserIdResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type VerifyReferenceOfByUserIdAsyncResult struct {
	result *VerifyReferenceOfByUserIdResult
	err    error
}

func NewVerifyReferenceOfByUserIdResultFromJson(data string) VerifyReferenceOfByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfByUserIdResultFromDict(dict)
}

func NewVerifyReferenceOfByUserIdResultFromDict(data map[string]interface{}) VerifyReferenceOfByUserIdResult {
	return VerifyReferenceOfByUserIdResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p VerifyReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p VerifyReferenceOfByUserIdResult) Pointer() *VerifyReferenceOfByUserIdResult {
	return &p
}

type AddReferenceOfResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfAsyncResult struct {
	result *AddReferenceOfResult
	err    error
}

func NewAddReferenceOfResultFromJson(data string) AddReferenceOfResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfResultFromDict(dict)
}

func NewAddReferenceOfResultFromDict(data map[string]interface{}) AddReferenceOfResult {
	return AddReferenceOfResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p AddReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p AddReferenceOfResult) Pointer() *AddReferenceOfResult {
	return &p
}

type AddReferenceOfByUserIdResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfByUserIdAsyncResult struct {
	result *AddReferenceOfByUserIdResult
	err    error
}

func NewAddReferenceOfByUserIdResultFromJson(data string) AddReferenceOfByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfByUserIdResultFromDict(dict)
}

func NewAddReferenceOfByUserIdResultFromDict(data map[string]interface{}) AddReferenceOfByUserIdResult {
	return AddReferenceOfByUserIdResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p AddReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p AddReferenceOfByUserIdResult) Pointer() *AddReferenceOfByUserIdResult {
	return &p
}

type DeleteReferenceOfResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfAsyncResult struct {
	result *DeleteReferenceOfResult
	err    error
}

func NewDeleteReferenceOfResultFromJson(data string) DeleteReferenceOfResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfResultFromDict(dict)
}

func NewDeleteReferenceOfResultFromDict(data map[string]interface{}) DeleteReferenceOfResult {
	return DeleteReferenceOfResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DeleteReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DeleteReferenceOfResult) Pointer() *DeleteReferenceOfResult {
	return &p
}

type DeleteReferenceOfByUserIdResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfByUserIdAsyncResult struct {
	result *DeleteReferenceOfByUserIdResult
	err    error
}

func NewDeleteReferenceOfByUserIdResultFromJson(data string) DeleteReferenceOfByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfByUserIdResultFromDict(dict)
}

func NewDeleteReferenceOfByUserIdResultFromDict(data map[string]interface{}) DeleteReferenceOfByUserIdResult {
	return DeleteReferenceOfByUserIdResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DeleteReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DeleteReferenceOfByUserIdResult) Pointer() *DeleteReferenceOfByUserIdResult {
	return &p
}

type DeleteItemSetByUserIdResult struct {
	Items     []ItemSet  `json:"items"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DeleteItemSetByUserIdAsyncResult struct {
	result *DeleteItemSetByUserIdResult
	err    error
}

func NewDeleteItemSetByUserIdResultFromJson(data string) DeleteItemSetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteItemSetByUserIdResultFromDict(dict)
}

func NewDeleteItemSetByUserIdResultFromDict(data map[string]interface{}) DeleteItemSetByUserIdResult {
	return DeleteItemSetByUserIdResult{
		Items:     CastItemSets(core.CastArray(data["items"])),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DeleteItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DeleteItemSetByUserIdResult) Pointer() *DeleteItemSetByUserIdResult {
	return &p
}

type AcquireItemSetByStampSheetResult struct {
	Items         []ItemSet  `json:"items"`
	ItemModel     *ItemModel `json:"itemModel"`
	Inventory     *Inventory `json:"inventory"`
	OverflowCount *int64     `json:"overflowCount"`
}

type AcquireItemSetByStampSheetAsyncResult struct {
	result *AcquireItemSetByStampSheetResult
	err    error
}

func NewAcquireItemSetByStampSheetResultFromJson(data string) AcquireItemSetByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetByStampSheetResultFromDict(dict)
}

func NewAcquireItemSetByStampSheetResultFromDict(data map[string]interface{}) AcquireItemSetByStampSheetResult {
	return AcquireItemSetByStampSheetResult{
		Items:         CastItemSets(core.CastArray(data["items"])),
		ItemModel:     NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory:     NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		OverflowCount: core.CastInt64(data["overflowCount"]),
	}
}

func (p AcquireItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel":     p.ItemModel.ToDict(),
		"inventory":     p.Inventory.ToDict(),
		"overflowCount": p.OverflowCount,
	}
}

func (p AcquireItemSetByStampSheetResult) Pointer() *AcquireItemSetByStampSheetResult {
	return &p
}

type AddReferenceOfItemSetByStampSheetResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type AddReferenceOfItemSetByStampSheetAsyncResult struct {
	result *AddReferenceOfItemSetByStampSheetResult
	err    error
}

func NewAddReferenceOfItemSetByStampSheetResultFromJson(data string) AddReferenceOfItemSetByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddReferenceOfItemSetByStampSheetResultFromDict(dict)
}

func NewAddReferenceOfItemSetByStampSheetResultFromDict(data map[string]interface{}) AddReferenceOfItemSetByStampSheetResult {
	return AddReferenceOfItemSetByStampSheetResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p AddReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p AddReferenceOfItemSetByStampSheetResult) Pointer() *AddReferenceOfItemSetByStampSheetResult {
	return &p
}

type DeleteReferenceOfItemSetByStampSheetResult struct {
	Item      []string   `json:"item"`
	ItemSet   *ItemSet   `json:"itemSet"`
	ItemModel *ItemModel `json:"itemModel"`
	Inventory *Inventory `json:"inventory"`
}

type DeleteReferenceOfItemSetByStampSheetAsyncResult struct {
	result *DeleteReferenceOfItemSetByStampSheetResult
	err    error
}

func NewDeleteReferenceOfItemSetByStampSheetResultFromJson(data string) DeleteReferenceOfItemSetByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReferenceOfItemSetByStampSheetResultFromDict(dict)
}

func NewDeleteReferenceOfItemSetByStampSheetResultFromDict(data map[string]interface{}) DeleteReferenceOfItemSetByStampSheetResult {
	return DeleteReferenceOfItemSetByStampSheetResult{
		Item:      core.CastStrings(core.CastArray(data["item"])),
		ItemSet:   NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel: NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory: NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
	}
}

func (p DeleteReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":   p.ItemSet.ToDict(),
		"itemModel": p.ItemModel.ToDict(),
		"inventory": p.Inventory.ToDict(),
	}
}

func (p DeleteReferenceOfItemSetByStampSheetResult) Pointer() *DeleteReferenceOfItemSetByStampSheetResult {
	return &p
}

type ConsumeItemSetByStampTaskResult struct {
	Items           []ItemSet  `json:"items"`
	ItemModel       *ItemModel `json:"itemModel"`
	Inventory       *Inventory `json:"inventory"`
	NewContextStack *string    `json:"newContextStack"`
}

type ConsumeItemSetByStampTaskAsyncResult struct {
	result *ConsumeItemSetByStampTaskResult
	err    error
}

func NewConsumeItemSetByStampTaskResultFromJson(data string) ConsumeItemSetByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeItemSetByStampTaskResultFromDict(dict)
}

func NewConsumeItemSetByStampTaskResultFromDict(data map[string]interface{}) ConsumeItemSetByStampTaskResult {
	return ConsumeItemSetByStampTaskResult{
		Items:           CastItemSets(core.CastArray(data["items"])),
		ItemModel:       NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory:       NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p ConsumeItemSetByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel":       p.ItemModel.ToDict(),
		"inventory":       p.Inventory.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p ConsumeItemSetByStampTaskResult) Pointer() *ConsumeItemSetByStampTaskResult {
	return &p
}

type VerifyReferenceOfByStampTaskResult struct {
	Item            []string   `json:"item"`
	ItemSet         *ItemSet   `json:"itemSet"`
	ItemModel       *ItemModel `json:"itemModel"`
	Inventory       *Inventory `json:"inventory"`
	NewContextStack *string    `json:"newContextStack"`
}

type VerifyReferenceOfByStampTaskAsyncResult struct {
	result *VerifyReferenceOfByStampTaskResult
	err    error
}

func NewVerifyReferenceOfByStampTaskResultFromJson(data string) VerifyReferenceOfByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReferenceOfByStampTaskResultFromDict(dict)
}

func NewVerifyReferenceOfByStampTaskResultFromDict(data map[string]interface{}) VerifyReferenceOfByStampTaskResult {
	return VerifyReferenceOfByStampTaskResult{
		Item:            core.CastStrings(core.CastArray(data["item"])),
		ItemSet:         NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer(),
		ItemModel:       NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer(),
		Inventory:       NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyReferenceOfByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": core.CastStringsFromDict(
			p.Item,
		),
		"itemSet":         p.ItemSet.ToDict(),
		"itemModel":       p.ItemModel.ToDict(),
		"inventory":       p.Inventory.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyReferenceOfByStampTaskResult) Pointer() *VerifyReferenceOfByStampTaskResult {
	return &p
}
