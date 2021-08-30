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

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
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

type DescribeStaminaModelMastersResult struct {
	Items         []StaminaModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
}

type DescribeStaminaModelMastersAsyncResult struct {
	result *DescribeStaminaModelMastersResult
	err    error
}

func NewDescribeStaminaModelMastersResultFromDict(data map[string]interface{}) DescribeStaminaModelMastersResult {
	return DescribeStaminaModelMastersResult{
		Items:         CastStaminaModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStaminaModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStaminaModelMastersResult) Pointer() *DescribeStaminaModelMastersResult {
	return &p
}

type CreateStaminaModelMasterResult struct {
	Item *StaminaModelMaster `json:"item"`
}

type CreateStaminaModelMasterAsyncResult struct {
	result *CreateStaminaModelMasterResult
	err    error
}

func NewCreateStaminaModelMasterResultFromDict(data map[string]interface{}) CreateStaminaModelMasterResult {
	return CreateStaminaModelMasterResult{
		Item: NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateStaminaModelMasterResult) Pointer() *CreateStaminaModelMasterResult {
	return &p
}

type GetStaminaModelMasterResult struct {
	Item *StaminaModelMaster `json:"item"`
}

type GetStaminaModelMasterAsyncResult struct {
	result *GetStaminaModelMasterResult
	err    error
}

func NewGetStaminaModelMasterResultFromDict(data map[string]interface{}) GetStaminaModelMasterResult {
	return GetStaminaModelMasterResult{
		Item: NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStaminaModelMasterResult) Pointer() *GetStaminaModelMasterResult {
	return &p
}

type UpdateStaminaModelMasterResult struct {
	Item *StaminaModelMaster `json:"item"`
}

type UpdateStaminaModelMasterAsyncResult struct {
	result *UpdateStaminaModelMasterResult
	err    error
}

func NewUpdateStaminaModelMasterResultFromDict(data map[string]interface{}) UpdateStaminaModelMasterResult {
	return UpdateStaminaModelMasterResult{
		Item: NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateStaminaModelMasterResult) Pointer() *UpdateStaminaModelMasterResult {
	return &p
}

type DeleteStaminaModelMasterResult struct {
	Item *StaminaModelMaster `json:"item"`
}

type DeleteStaminaModelMasterAsyncResult struct {
	result *DeleteStaminaModelMasterResult
	err    error
}

func NewDeleteStaminaModelMasterResultFromDict(data map[string]interface{}) DeleteStaminaModelMasterResult {
	return DeleteStaminaModelMasterResult{
		Item: NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteStaminaModelMasterResult) Pointer() *DeleteStaminaModelMasterResult {
	return &p
}

type DescribeMaxStaminaTableMastersResult struct {
	Items         []MaxStaminaTableMaster `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
}

type DescribeMaxStaminaTableMastersAsyncResult struct {
	result *DescribeMaxStaminaTableMastersResult
	err    error
}

func NewDescribeMaxStaminaTableMastersResultFromDict(data map[string]interface{}) DescribeMaxStaminaTableMastersResult {
	return DescribeMaxStaminaTableMastersResult{
		Items:         CastMaxStaminaTableMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMaxStaminaTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMaxStaminaTableMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMaxStaminaTableMastersResult) Pointer() *DescribeMaxStaminaTableMastersResult {
	return &p
}

type CreateMaxStaminaTableMasterResult struct {
	Item *MaxStaminaTableMaster `json:"item"`
}

type CreateMaxStaminaTableMasterAsyncResult struct {
	result *CreateMaxStaminaTableMasterResult
	err    error
}

func NewCreateMaxStaminaTableMasterResultFromDict(data map[string]interface{}) CreateMaxStaminaTableMasterResult {
	return CreateMaxStaminaTableMasterResult{
		Item: NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateMaxStaminaTableMasterResult) Pointer() *CreateMaxStaminaTableMasterResult {
	return &p
}

type GetMaxStaminaTableMasterResult struct {
	Item *MaxStaminaTableMaster `json:"item"`
}

type GetMaxStaminaTableMasterAsyncResult struct {
	result *GetMaxStaminaTableMasterResult
	err    error
}

func NewGetMaxStaminaTableMasterResultFromDict(data map[string]interface{}) GetMaxStaminaTableMasterResult {
	return GetMaxStaminaTableMasterResult{
		Item: NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMaxStaminaTableMasterResult) Pointer() *GetMaxStaminaTableMasterResult {
	return &p
}

type UpdateMaxStaminaTableMasterResult struct {
	Item *MaxStaminaTableMaster `json:"item"`
}

type UpdateMaxStaminaTableMasterAsyncResult struct {
	result *UpdateMaxStaminaTableMasterResult
	err    error
}

func NewUpdateMaxStaminaTableMasterResultFromDict(data map[string]interface{}) UpdateMaxStaminaTableMasterResult {
	return UpdateMaxStaminaTableMasterResult{
		Item: NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateMaxStaminaTableMasterResult) Pointer() *UpdateMaxStaminaTableMasterResult {
	return &p
}

type DeleteMaxStaminaTableMasterResult struct {
	Item *MaxStaminaTableMaster `json:"item"`
}

type DeleteMaxStaminaTableMasterAsyncResult struct {
	result *DeleteMaxStaminaTableMasterResult
	err    error
}

func NewDeleteMaxStaminaTableMasterResultFromDict(data map[string]interface{}) DeleteMaxStaminaTableMasterResult {
	return DeleteMaxStaminaTableMasterResult{
		Item: NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMaxStaminaTableMasterResult) Pointer() *DeleteMaxStaminaTableMasterResult {
	return &p
}

type DescribeRecoverIntervalTableMastersResult struct {
	Items         []RecoverIntervalTableMaster `json:"items"`
	NextPageToken *string                      `json:"nextPageToken"`
}

type DescribeRecoverIntervalTableMastersAsyncResult struct {
	result *DescribeRecoverIntervalTableMastersResult
	err    error
}

func NewDescribeRecoverIntervalTableMastersResultFromDict(data map[string]interface{}) DescribeRecoverIntervalTableMastersResult {
	return DescribeRecoverIntervalTableMastersResult{
		Items:         CastRecoverIntervalTableMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRecoverIntervalTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRecoverIntervalTableMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRecoverIntervalTableMastersResult) Pointer() *DescribeRecoverIntervalTableMastersResult {
	return &p
}

type CreateRecoverIntervalTableMasterResult struct {
	Item *RecoverIntervalTableMaster `json:"item"`
}

type CreateRecoverIntervalTableMasterAsyncResult struct {
	result *CreateRecoverIntervalTableMasterResult
	err    error
}

func NewCreateRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) CreateRecoverIntervalTableMasterResult {
	return CreateRecoverIntervalTableMasterResult{
		Item: NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateRecoverIntervalTableMasterResult) Pointer() *CreateRecoverIntervalTableMasterResult {
	return &p
}

type GetRecoverIntervalTableMasterResult struct {
	Item *RecoverIntervalTableMaster `json:"item"`
}

type GetRecoverIntervalTableMasterAsyncResult struct {
	result *GetRecoverIntervalTableMasterResult
	err    error
}

func NewGetRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) GetRecoverIntervalTableMasterResult {
	return GetRecoverIntervalTableMasterResult{
		Item: NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRecoverIntervalTableMasterResult) Pointer() *GetRecoverIntervalTableMasterResult {
	return &p
}

type UpdateRecoverIntervalTableMasterResult struct {
	Item *RecoverIntervalTableMaster `json:"item"`
}

type UpdateRecoverIntervalTableMasterAsyncResult struct {
	result *UpdateRecoverIntervalTableMasterResult
	err    error
}

func NewUpdateRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) UpdateRecoverIntervalTableMasterResult {
	return UpdateRecoverIntervalTableMasterResult{
		Item: NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRecoverIntervalTableMasterResult) Pointer() *UpdateRecoverIntervalTableMasterResult {
	return &p
}

type DeleteRecoverIntervalTableMasterResult struct {
	Item *RecoverIntervalTableMaster `json:"item"`
}

type DeleteRecoverIntervalTableMasterAsyncResult struct {
	result *DeleteRecoverIntervalTableMasterResult
	err    error
}

func NewDeleteRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) DeleteRecoverIntervalTableMasterResult {
	return DeleteRecoverIntervalTableMasterResult{
		Item: NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRecoverIntervalTableMasterResult) Pointer() *DeleteRecoverIntervalTableMasterResult {
	return &p
}

type DescribeRecoverValueTableMastersResult struct {
	Items         []RecoverValueTableMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribeRecoverValueTableMastersAsyncResult struct {
	result *DescribeRecoverValueTableMastersResult
	err    error
}

func NewDescribeRecoverValueTableMastersResultFromDict(data map[string]interface{}) DescribeRecoverValueTableMastersResult {
	return DescribeRecoverValueTableMastersResult{
		Items:         CastRecoverValueTableMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRecoverValueTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRecoverValueTableMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRecoverValueTableMastersResult) Pointer() *DescribeRecoverValueTableMastersResult {
	return &p
}

type CreateRecoverValueTableMasterResult struct {
	Item *RecoverValueTableMaster `json:"item"`
}

type CreateRecoverValueTableMasterAsyncResult struct {
	result *CreateRecoverValueTableMasterResult
	err    error
}

func NewCreateRecoverValueTableMasterResultFromDict(data map[string]interface{}) CreateRecoverValueTableMasterResult {
	return CreateRecoverValueTableMasterResult{
		Item: NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateRecoverValueTableMasterResult) Pointer() *CreateRecoverValueTableMasterResult {
	return &p
}

type GetRecoverValueTableMasterResult struct {
	Item *RecoverValueTableMaster `json:"item"`
}

type GetRecoverValueTableMasterAsyncResult struct {
	result *GetRecoverValueTableMasterResult
	err    error
}

func NewGetRecoverValueTableMasterResultFromDict(data map[string]interface{}) GetRecoverValueTableMasterResult {
	return GetRecoverValueTableMasterResult{
		Item: NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRecoverValueTableMasterResult) Pointer() *GetRecoverValueTableMasterResult {
	return &p
}

type UpdateRecoverValueTableMasterResult struct {
	Item *RecoverValueTableMaster `json:"item"`
}

type UpdateRecoverValueTableMasterAsyncResult struct {
	result *UpdateRecoverValueTableMasterResult
	err    error
}

func NewUpdateRecoverValueTableMasterResultFromDict(data map[string]interface{}) UpdateRecoverValueTableMasterResult {
	return UpdateRecoverValueTableMasterResult{
		Item: NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRecoverValueTableMasterResult) Pointer() *UpdateRecoverValueTableMasterResult {
	return &p
}

type DeleteRecoverValueTableMasterResult struct {
	Item *RecoverValueTableMaster `json:"item"`
}

type DeleteRecoverValueTableMasterAsyncResult struct {
	result *DeleteRecoverValueTableMasterResult
	err    error
}

func NewDeleteRecoverValueTableMasterResultFromDict(data map[string]interface{}) DeleteRecoverValueTableMasterResult {
	return DeleteRecoverValueTableMasterResult{
		Item: NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRecoverValueTableMasterResult) Pointer() *DeleteRecoverValueTableMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentStaminaMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentStaminaMasterResult struct {
	Item *CurrentStaminaMaster `json:"item"`
}

type GetCurrentStaminaMasterAsyncResult struct {
	result *GetCurrentStaminaMasterResult
	err    error
}

func NewGetCurrentStaminaMasterResultFromDict(data map[string]interface{}) GetCurrentStaminaMasterResult {
	return GetCurrentStaminaMasterResult{
		Item: NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentStaminaMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentStaminaMasterResult) Pointer() *GetCurrentStaminaMasterResult {
	return &p
}

type UpdateCurrentStaminaMasterResult struct {
	Item *CurrentStaminaMaster `json:"item"`
}

type UpdateCurrentStaminaMasterAsyncResult struct {
	result *UpdateCurrentStaminaMasterResult
	err    error
}

func NewUpdateCurrentStaminaMasterResultFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterResult {
	return UpdateCurrentStaminaMasterResult{
		Item: NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentStaminaMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentStaminaMasterResult) Pointer() *UpdateCurrentStaminaMasterResult {
	return &p
}

type UpdateCurrentStaminaMasterFromGitHubResult struct {
	Item *CurrentStaminaMaster `json:"item"`
}

type UpdateCurrentStaminaMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentStaminaMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentStaminaMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterFromGitHubResult {
	return UpdateCurrentStaminaMasterFromGitHubResult{
		Item: NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubResult) Pointer() *UpdateCurrentStaminaMasterFromGitHubResult {
	return &p
}

type DescribeStaminaModelsResult struct {
	Items []StaminaModel `json:"items"`
}

type DescribeStaminaModelsAsyncResult struct {
	result *DescribeStaminaModelsResult
	err    error
}

func NewDescribeStaminaModelsResultFromDict(data map[string]interface{}) DescribeStaminaModelsResult {
	return DescribeStaminaModelsResult{
		Items: CastStaminaModels(core.CastArray(data["items"])),
	}
}

func (p DescribeStaminaModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeStaminaModelsResult) Pointer() *DescribeStaminaModelsResult {
	return &p
}

type GetStaminaModelResult struct {
	Item *StaminaModel `json:"item"`
}

type GetStaminaModelAsyncResult struct {
	result *GetStaminaModelResult
	err    error
}

func NewGetStaminaModelResultFromDict(data map[string]interface{}) GetStaminaModelResult {
	return GetStaminaModelResult{
		Item: NewStaminaModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStaminaModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetStaminaModelResult) Pointer() *GetStaminaModelResult {
	return &p
}

type DescribeStaminasResult struct {
	Items         []Stamina `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeStaminasAsyncResult struct {
	result *DescribeStaminasResult
	err    error
}

func NewDescribeStaminasResultFromDict(data map[string]interface{}) DescribeStaminasResult {
	return DescribeStaminasResult{
		Items:         CastStaminas(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStaminasResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStaminasResult) Pointer() *DescribeStaminasResult {
	return &p
}

type DescribeStaminasByUserIdResult struct {
	Items         []Stamina `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeStaminasByUserIdAsyncResult struct {
	result *DescribeStaminasByUserIdResult
	err    error
}

func NewDescribeStaminasByUserIdResultFromDict(data map[string]interface{}) DescribeStaminasByUserIdResult {
	return DescribeStaminasByUserIdResult{
		Items:         CastStaminas(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStaminasByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStaminasByUserIdResult) Pointer() *DescribeStaminasByUserIdResult {
	return &p
}

type GetStaminaResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type GetStaminaAsyncResult struct {
	result *GetStaminaResult
	err    error
}

func NewGetStaminaResultFromDict(data map[string]interface{}) GetStaminaResult {
	return GetStaminaResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p GetStaminaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p GetStaminaResult) Pointer() *GetStaminaResult {
	return &p
}

type GetStaminaByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type GetStaminaByUserIdAsyncResult struct {
	result *GetStaminaByUserIdResult
	err    error
}

func NewGetStaminaByUserIdResultFromDict(data map[string]interface{}) GetStaminaByUserIdResult {
	return GetStaminaByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p GetStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p GetStaminaByUserIdResult) Pointer() *GetStaminaByUserIdResult {
	return &p
}

type UpdateStaminaByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type UpdateStaminaByUserIdAsyncResult struct {
	result *UpdateStaminaByUserIdResult
	err    error
}

func NewUpdateStaminaByUserIdResultFromDict(data map[string]interface{}) UpdateStaminaByUserIdResult {
	return UpdateStaminaByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p UpdateStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p UpdateStaminaByUserIdResult) Pointer() *UpdateStaminaByUserIdResult {
	return &p
}

type ConsumeStaminaResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type ConsumeStaminaAsyncResult struct {
	result *ConsumeStaminaResult
	err    error
}

func NewConsumeStaminaResultFromDict(data map[string]interface{}) ConsumeStaminaResult {
	return ConsumeStaminaResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p ConsumeStaminaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p ConsumeStaminaResult) Pointer() *ConsumeStaminaResult {
	return &p
}

type ConsumeStaminaByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type ConsumeStaminaByUserIdAsyncResult struct {
	result *ConsumeStaminaByUserIdResult
	err    error
}

func NewConsumeStaminaByUserIdResultFromDict(data map[string]interface{}) ConsumeStaminaByUserIdResult {
	return ConsumeStaminaByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p ConsumeStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p ConsumeStaminaByUserIdResult) Pointer() *ConsumeStaminaByUserIdResult {
	return &p
}

type RecoverStaminaByUserIdResult struct {
	Item          *Stamina      `json:"item"`
	StaminaModel  *StaminaModel `json:"staminaModel"`
	OverflowValue *int64        `json:"overflowValue"`
}

type RecoverStaminaByUserIdAsyncResult struct {
	result *RecoverStaminaByUserIdResult
	err    error
}

func NewRecoverStaminaByUserIdResultFromDict(data map[string]interface{}) RecoverStaminaByUserIdResult {
	return RecoverStaminaByUserIdResult{
		Item:          NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel:  NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
		OverflowValue: core.CastInt64(data["overflowValue"]),
	}
}

func (p RecoverStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"staminaModel":  p.StaminaModel.ToDict(),
		"overflowValue": p.OverflowValue,
	}
}

func (p RecoverStaminaByUserIdResult) Pointer() *RecoverStaminaByUserIdResult {
	return &p
}

type RaiseMaxValueByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type RaiseMaxValueByUserIdAsyncResult struct {
	result *RaiseMaxValueByUserIdResult
	err    error
}

func NewRaiseMaxValueByUserIdResultFromDict(data map[string]interface{}) RaiseMaxValueByUserIdResult {
	return RaiseMaxValueByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p RaiseMaxValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p RaiseMaxValueByUserIdResult) Pointer() *RaiseMaxValueByUserIdResult {
	return &p
}

type SetMaxValueByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetMaxValueByUserIdAsyncResult struct {
	result *SetMaxValueByUserIdResult
	err    error
}

func NewSetMaxValueByUserIdResultFromDict(data map[string]interface{}) SetMaxValueByUserIdResult {
	return SetMaxValueByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetMaxValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetMaxValueByUserIdResult) Pointer() *SetMaxValueByUserIdResult {
	return &p
}

type SetRecoverIntervalByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverIntervalByUserIdAsyncResult struct {
	result *SetRecoverIntervalByUserIdResult
	err    error
}

func NewSetRecoverIntervalByUserIdResultFromDict(data map[string]interface{}) SetRecoverIntervalByUserIdResult {
	return SetRecoverIntervalByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverIntervalByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverIntervalByUserIdResult) Pointer() *SetRecoverIntervalByUserIdResult {
	return &p
}

type SetRecoverValueByUserIdResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverValueByUserIdAsyncResult struct {
	result *SetRecoverValueByUserIdResult
	err    error
}

func NewSetRecoverValueByUserIdResultFromDict(data map[string]interface{}) SetRecoverValueByUserIdResult {
	return SetRecoverValueByUserIdResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverValueByUserIdResult) Pointer() *SetRecoverValueByUserIdResult {
	return &p
}

type SetMaxValueByStatusResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetMaxValueByStatusAsyncResult struct {
	result *SetMaxValueByStatusResult
	err    error
}

func NewSetMaxValueByStatusResultFromDict(data map[string]interface{}) SetMaxValueByStatusResult {
	return SetMaxValueByStatusResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetMaxValueByStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetMaxValueByStatusResult) Pointer() *SetMaxValueByStatusResult {
	return &p
}

type SetRecoverIntervalByStatusResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverIntervalByStatusAsyncResult struct {
	result *SetRecoverIntervalByStatusResult
	err    error
}

func NewSetRecoverIntervalByStatusResultFromDict(data map[string]interface{}) SetRecoverIntervalByStatusResult {
	return SetRecoverIntervalByStatusResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverIntervalByStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverIntervalByStatusResult) Pointer() *SetRecoverIntervalByStatusResult {
	return &p
}

type SetRecoverValueByStatusResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverValueByStatusAsyncResult struct {
	result *SetRecoverValueByStatusResult
	err    error
}

func NewSetRecoverValueByStatusResultFromDict(data map[string]interface{}) SetRecoverValueByStatusResult {
	return SetRecoverValueByStatusResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverValueByStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverValueByStatusResult) Pointer() *SetRecoverValueByStatusResult {
	return &p
}

type DeleteStaminaByUserIdResult struct {
}

type DeleteStaminaByUserIdAsyncResult struct {
	result *DeleteStaminaByUserIdResult
	err    error
}

func NewDeleteStaminaByUserIdResultFromDict(data map[string]interface{}) DeleteStaminaByUserIdResult {
	return DeleteStaminaByUserIdResult{}
}

func (p DeleteStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteStaminaByUserIdResult) Pointer() *DeleteStaminaByUserIdResult {
	return &p
}

type RecoverStaminaByStampSheetResult struct {
	Item          *Stamina      `json:"item"`
	StaminaModel  *StaminaModel `json:"staminaModel"`
	OverflowValue *int64        `json:"overflowValue"`
}

type RecoverStaminaByStampSheetAsyncResult struct {
	result *RecoverStaminaByStampSheetResult
	err    error
}

func NewRecoverStaminaByStampSheetResultFromDict(data map[string]interface{}) RecoverStaminaByStampSheetResult {
	return RecoverStaminaByStampSheetResult{
		Item:          NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel:  NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
		OverflowValue: core.CastInt64(data["overflowValue"]),
	}
}

func (p RecoverStaminaByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":          p.Item.ToDict(),
		"staminaModel":  p.StaminaModel.ToDict(),
		"overflowValue": p.OverflowValue,
	}
}

func (p RecoverStaminaByStampSheetResult) Pointer() *RecoverStaminaByStampSheetResult {
	return &p
}

type RaiseMaxValueByStampSheetResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type RaiseMaxValueByStampSheetAsyncResult struct {
	result *RaiseMaxValueByStampSheetResult
	err    error
}

func NewRaiseMaxValueByStampSheetResultFromDict(data map[string]interface{}) RaiseMaxValueByStampSheetResult {
	return RaiseMaxValueByStampSheetResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p RaiseMaxValueByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p RaiseMaxValueByStampSheetResult) Pointer() *RaiseMaxValueByStampSheetResult {
	return &p
}

type SetMaxValueByStampSheetResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetMaxValueByStampSheetAsyncResult struct {
	result *SetMaxValueByStampSheetResult
	err    error
}

func NewSetMaxValueByStampSheetResultFromDict(data map[string]interface{}) SetMaxValueByStampSheetResult {
	return SetMaxValueByStampSheetResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetMaxValueByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetMaxValueByStampSheetResult) Pointer() *SetMaxValueByStampSheetResult {
	return &p
}

type SetRecoverIntervalByStampSheetResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverIntervalByStampSheetAsyncResult struct {
	result *SetRecoverIntervalByStampSheetResult
	err    error
}

func NewSetRecoverIntervalByStampSheetResultFromDict(data map[string]interface{}) SetRecoverIntervalByStampSheetResult {
	return SetRecoverIntervalByStampSheetResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverIntervalByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverIntervalByStampSheetResult) Pointer() *SetRecoverIntervalByStampSheetResult {
	return &p
}

type SetRecoverValueByStampSheetResult struct {
	Item         *Stamina      `json:"item"`
	StaminaModel *StaminaModel `json:"staminaModel"`
}

type SetRecoverValueByStampSheetAsyncResult struct {
	result *SetRecoverValueByStampSheetResult
	err    error
}

func NewSetRecoverValueByStampSheetResultFromDict(data map[string]interface{}) SetRecoverValueByStampSheetResult {
	return SetRecoverValueByStampSheetResult{
		Item:         NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel: NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
	}
}

func (p SetRecoverValueByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":         p.Item.ToDict(),
		"staminaModel": p.StaminaModel.ToDict(),
	}
}

func (p SetRecoverValueByStampSheetResult) Pointer() *SetRecoverValueByStampSheetResult {
	return &p
}

type ConsumeStaminaByStampTaskResult struct {
	Item            *Stamina      `json:"item"`
	StaminaModel    *StaminaModel `json:"staminaModel"`
	NewContextStack *string       `json:"newContextStack"`
}

type ConsumeStaminaByStampTaskAsyncResult struct {
	result *ConsumeStaminaByStampTaskResult
	err    error
}

func NewConsumeStaminaByStampTaskResultFromDict(data map[string]interface{}) ConsumeStaminaByStampTaskResult {
	return ConsumeStaminaByStampTaskResult{
		Item:            NewStaminaFromDict(core.CastMap(data["item"])).Pointer(),
		StaminaModel:    NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p ConsumeStaminaByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"staminaModel":    p.StaminaModel.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p ConsumeStaminaByStampTaskResult) Pointer() *ConsumeStaminaByStampTaskResult {
	return &p
}
