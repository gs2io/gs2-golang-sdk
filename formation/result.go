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

type DescribeFormModelMastersResult struct {
	Items         []FormModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeFormModelMastersAsyncResult struct {
	result *DescribeFormModelMastersResult
	err    error
}

func NewDescribeFormModelMastersResultFromDict(data map[string]interface{}) DescribeFormModelMastersResult {
	return DescribeFormModelMastersResult{
		Items:         CastFormModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeFormModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeFormModelMastersResult) Pointer() *DescribeFormModelMastersResult {
	return &p
}

type CreateFormModelMasterResult struct {
	Item *FormModelMaster `json:"item"`
}

type CreateFormModelMasterAsyncResult struct {
	result *CreateFormModelMasterResult
	err    error
}

func NewCreateFormModelMasterResultFromDict(data map[string]interface{}) CreateFormModelMasterResult {
	return CreateFormModelMasterResult{
		Item: NewFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateFormModelMasterResult) Pointer() *CreateFormModelMasterResult {
	return &p
}

type GetFormModelMasterResult struct {
	Item *FormModelMaster `json:"item"`
}

type GetFormModelMasterAsyncResult struct {
	result *GetFormModelMasterResult
	err    error
}

func NewGetFormModelMasterResultFromDict(data map[string]interface{}) GetFormModelMasterResult {
	return GetFormModelMasterResult{
		Item: NewFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetFormModelMasterResult) Pointer() *GetFormModelMasterResult {
	return &p
}

type UpdateFormModelMasterResult struct {
	Item *FormModelMaster `json:"item"`
}

type UpdateFormModelMasterAsyncResult struct {
	result *UpdateFormModelMasterResult
	err    error
}

func NewUpdateFormModelMasterResultFromDict(data map[string]interface{}) UpdateFormModelMasterResult {
	return UpdateFormModelMasterResult{
		Item: NewFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateFormModelMasterResult) Pointer() *UpdateFormModelMasterResult {
	return &p
}

type DeleteFormModelMasterResult struct {
	Item *FormModelMaster `json:"item"`
}

type DeleteFormModelMasterAsyncResult struct {
	result *DeleteFormModelMasterResult
	err    error
}

func NewDeleteFormModelMasterResultFromDict(data map[string]interface{}) DeleteFormModelMasterResult {
	return DeleteFormModelMasterResult{
		Item: NewFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteFormModelMasterResult) Pointer() *DeleteFormModelMasterResult {
	return &p
}

type DescribeMoldModelsResult struct {
	Items []MoldModel `json:"items"`
}

type DescribeMoldModelsAsyncResult struct {
	result *DescribeMoldModelsResult
	err    error
}

func NewDescribeMoldModelsResultFromDict(data map[string]interface{}) DescribeMoldModelsResult {
	return DescribeMoldModelsResult{
		Items: CastMoldModels(core.CastArray(data["items"])),
	}
}

func (p DescribeMoldModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoldModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeMoldModelsResult) Pointer() *DescribeMoldModelsResult {
	return &p
}

type GetMoldModelResult struct {
	Item *MoldModel `json:"item"`
}

type GetMoldModelAsyncResult struct {
	result *GetMoldModelResult
	err    error
}

func NewGetMoldModelResultFromDict(data map[string]interface{}) GetMoldModelResult {
	return GetMoldModelResult{
		Item: NewMoldModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMoldModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMoldModelResult) Pointer() *GetMoldModelResult {
	return &p
}

type DescribeMoldModelMastersResult struct {
	Items         []MoldModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
}

type DescribeMoldModelMastersAsyncResult struct {
	result *DescribeMoldModelMastersResult
	err    error
}

func NewDescribeMoldModelMastersResultFromDict(data map[string]interface{}) DescribeMoldModelMastersResult {
	return DescribeMoldModelMastersResult{
		Items:         CastMoldModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMoldModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoldModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMoldModelMastersResult) Pointer() *DescribeMoldModelMastersResult {
	return &p
}

type CreateMoldModelMasterResult struct {
	Item *MoldModelMaster `json:"item"`
}

type CreateMoldModelMasterAsyncResult struct {
	result *CreateMoldModelMasterResult
	err    error
}

func NewCreateMoldModelMasterResultFromDict(data map[string]interface{}) CreateMoldModelMasterResult {
	return CreateMoldModelMasterResult{
		Item: NewMoldModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateMoldModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateMoldModelMasterResult) Pointer() *CreateMoldModelMasterResult {
	return &p
}

type GetMoldModelMasterResult struct {
	Item *MoldModelMaster `json:"item"`
}

type GetMoldModelMasterAsyncResult struct {
	result *GetMoldModelMasterResult
	err    error
}

func NewGetMoldModelMasterResultFromDict(data map[string]interface{}) GetMoldModelMasterResult {
	return GetMoldModelMasterResult{
		Item: NewMoldModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetMoldModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetMoldModelMasterResult) Pointer() *GetMoldModelMasterResult {
	return &p
}

type UpdateMoldModelMasterResult struct {
	Item *MoldModelMaster `json:"item"`
}

type UpdateMoldModelMasterAsyncResult struct {
	result *UpdateMoldModelMasterResult
	err    error
}

func NewUpdateMoldModelMasterResultFromDict(data map[string]interface{}) UpdateMoldModelMasterResult {
	return UpdateMoldModelMasterResult{
		Item: NewMoldModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateMoldModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateMoldModelMasterResult) Pointer() *UpdateMoldModelMasterResult {
	return &p
}

type DeleteMoldModelMasterResult struct {
	Item *MoldModelMaster `json:"item"`
}

type DeleteMoldModelMasterAsyncResult struct {
	result *DeleteMoldModelMasterResult
	err    error
}

func NewDeleteMoldModelMasterResultFromDict(data map[string]interface{}) DeleteMoldModelMasterResult {
	return DeleteMoldModelMasterResult{
		Item: NewMoldModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMoldModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMoldModelMasterResult) Pointer() *DeleteMoldModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentFormMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentFormMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentFormMasterResult struct {
	Item *CurrentFormMaster `json:"item"`
}

type GetCurrentFormMasterAsyncResult struct {
	result *GetCurrentFormMasterResult
	err    error
}

func NewGetCurrentFormMasterResultFromDict(data map[string]interface{}) GetCurrentFormMasterResult {
	return GetCurrentFormMasterResult{
		Item: NewCurrentFormMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentFormMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentFormMasterResult) Pointer() *GetCurrentFormMasterResult {
	return &p
}

type UpdateCurrentFormMasterResult struct {
	Item *CurrentFormMaster `json:"item"`
}

type UpdateCurrentFormMasterAsyncResult struct {
	result *UpdateCurrentFormMasterResult
	err    error
}

func NewUpdateCurrentFormMasterResultFromDict(data map[string]interface{}) UpdateCurrentFormMasterResult {
	return UpdateCurrentFormMasterResult{
		Item: NewCurrentFormMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentFormMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentFormMasterResult) Pointer() *UpdateCurrentFormMasterResult {
	return &p
}

type UpdateCurrentFormMasterFromGitHubResult struct {
	Item *CurrentFormMaster `json:"item"`
}

type UpdateCurrentFormMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentFormMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentFormMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentFormMasterFromGitHubResult {
	return UpdateCurrentFormMasterFromGitHubResult{
		Item: NewCurrentFormMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentFormMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentFormMasterFromGitHubResult) Pointer() *UpdateCurrentFormMasterFromGitHubResult {
	return &p
}

type DescribeMoldsResult struct {
	Items         []Mold  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeMoldsAsyncResult struct {
	result *DescribeMoldsResult
	err    error
}

func NewDescribeMoldsResultFromDict(data map[string]interface{}) DescribeMoldsResult {
	return DescribeMoldsResult{
		Items:         CastMolds(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMoldsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoldsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMoldsResult) Pointer() *DescribeMoldsResult {
	return &p
}

type DescribeMoldsByUserIdResult struct {
	Items         []Mold  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeMoldsByUserIdAsyncResult struct {
	result *DescribeMoldsByUserIdResult
	err    error
}

func NewDescribeMoldsByUserIdResultFromDict(data map[string]interface{}) DescribeMoldsByUserIdResult {
	return DescribeMoldsByUserIdResult{
		Items:         CastMolds(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeMoldsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMoldsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeMoldsByUserIdResult) Pointer() *DescribeMoldsByUserIdResult {
	return &p
}

type GetMoldResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type GetMoldAsyncResult struct {
	result *GetMoldResult
	err    error
}

func NewGetMoldResultFromDict(data map[string]interface{}) GetMoldResult {
	return GetMoldResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p GetMoldResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p GetMoldResult) Pointer() *GetMoldResult {
	return &p
}

type GetMoldByUserIdResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type GetMoldByUserIdAsyncResult struct {
	result *GetMoldByUserIdResult
	err    error
}

func NewGetMoldByUserIdResultFromDict(data map[string]interface{}) GetMoldByUserIdResult {
	return GetMoldByUserIdResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p GetMoldByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p GetMoldByUserIdResult) Pointer() *GetMoldByUserIdResult {
	return &p
}

type SetMoldCapacityByUserIdResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type SetMoldCapacityByUserIdAsyncResult struct {
	result *SetMoldCapacityByUserIdResult
	err    error
}

func NewSetMoldCapacityByUserIdResultFromDict(data map[string]interface{}) SetMoldCapacityByUserIdResult {
	return SetMoldCapacityByUserIdResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p SetMoldCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p SetMoldCapacityByUserIdResult) Pointer() *SetMoldCapacityByUserIdResult {
	return &p
}

type AddMoldCapacityByUserIdResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type AddMoldCapacityByUserIdAsyncResult struct {
	result *AddMoldCapacityByUserIdResult
	err    error
}

func NewAddMoldCapacityByUserIdResultFromDict(data map[string]interface{}) AddMoldCapacityByUserIdResult {
	return AddMoldCapacityByUserIdResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p AddMoldCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p AddMoldCapacityByUserIdResult) Pointer() *AddMoldCapacityByUserIdResult {
	return &p
}

type DeleteMoldResult struct {
	Item *Mold `json:"item"`
}

type DeleteMoldAsyncResult struct {
	result *DeleteMoldResult
	err    error
}

func NewDeleteMoldResultFromDict(data map[string]interface{}) DeleteMoldResult {
	return DeleteMoldResult{
		Item: NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMoldResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMoldResult) Pointer() *DeleteMoldResult {
	return &p
}

type DeleteMoldByUserIdResult struct {
	Item *Mold `json:"item"`
}

type DeleteMoldByUserIdAsyncResult struct {
	result *DeleteMoldByUserIdResult
	err    error
}

func NewDeleteMoldByUserIdResultFromDict(data map[string]interface{}) DeleteMoldByUserIdResult {
	return DeleteMoldByUserIdResult{
		Item: NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteMoldByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteMoldByUserIdResult) Pointer() *DeleteMoldByUserIdResult {
	return &p
}

type AddCapacityByStampSheetResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type AddCapacityByStampSheetAsyncResult struct {
	result *AddCapacityByStampSheetResult
	err    error
}

func NewAddCapacityByStampSheetResultFromDict(data map[string]interface{}) AddCapacityByStampSheetResult {
	return AddCapacityByStampSheetResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p AddCapacityByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p AddCapacityByStampSheetResult) Pointer() *AddCapacityByStampSheetResult {
	return &p
}

type SetCapacityByStampSheetResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type SetCapacityByStampSheetAsyncResult struct {
	result *SetCapacityByStampSheetResult
	err    error
}

func NewSetCapacityByStampSheetResultFromDict(data map[string]interface{}) SetCapacityByStampSheetResult {
	return SetCapacityByStampSheetResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p SetCapacityByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p SetCapacityByStampSheetResult) Pointer() *SetCapacityByStampSheetResult {
	return &p
}

type DescribeFormsResult struct {
	Items         []Form  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeFormsAsyncResult struct {
	result *DescribeFormsResult
	err    error
}

func NewDescribeFormsResultFromDict(data map[string]interface{}) DescribeFormsResult {
	return DescribeFormsResult{
		Items:         CastForms(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeFormsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeFormsResult) Pointer() *DescribeFormsResult {
	return &p
}

type DescribeFormsByUserIdResult struct {
	Items         []Form  `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeFormsByUserIdAsyncResult struct {
	result *DescribeFormsByUserIdResult
	err    error
}

func NewDescribeFormsByUserIdResultFromDict(data map[string]interface{}) DescribeFormsByUserIdResult {
	return DescribeFormsByUserIdResult{
		Items:         CastForms(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeFormsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFormsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeFormsByUserIdResult) Pointer() *DescribeFormsByUserIdResult {
	return &p
}

type GetFormResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type GetFormAsyncResult struct {
	result *GetFormResult
	err    error
}

func NewGetFormResultFromDict(data map[string]interface{}) GetFormResult {
	return GetFormResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p GetFormResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p GetFormResult) Pointer() *GetFormResult {
	return &p
}

type GetFormByUserIdResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type GetFormByUserIdAsyncResult struct {
	result *GetFormByUserIdResult
	err    error
}

func NewGetFormByUserIdResultFromDict(data map[string]interface{}) GetFormByUserIdResult {
	return GetFormByUserIdResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p GetFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p GetFormByUserIdResult) Pointer() *GetFormByUserIdResult {
	return &p
}

type GetFormWithSignatureResult struct {
	Item      *Form      `json:"item"`
	Body      *string    `json:"body"`
	Signature *string    `json:"signature"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type GetFormWithSignatureAsyncResult struct {
	result *GetFormWithSignatureResult
	err    error
}

func NewGetFormWithSignatureResultFromDict(data map[string]interface{}) GetFormWithSignatureResult {
	return GetFormWithSignatureResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p GetFormWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p GetFormWithSignatureResult) Pointer() *GetFormWithSignatureResult {
	return &p
}

type GetFormWithSignatureByUserIdResult struct {
	Item      *Form      `json:"item"`
	Body      *string    `json:"body"`
	Signature *string    `json:"signature"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type GetFormWithSignatureByUserIdAsyncResult struct {
	result *GetFormWithSignatureByUserIdResult
	err    error
}

func NewGetFormWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetFormWithSignatureByUserIdResult {
	return GetFormWithSignatureByUserIdResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p GetFormWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p GetFormWithSignatureByUserIdResult) Pointer() *GetFormWithSignatureByUserIdResult {
	return &p
}

type SetFormByUserIdResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type SetFormByUserIdAsyncResult struct {
	result *SetFormByUserIdResult
	err    error
}

func NewSetFormByUserIdResultFromDict(data map[string]interface{}) SetFormByUserIdResult {
	return SetFormByUserIdResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p SetFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p SetFormByUserIdResult) Pointer() *SetFormByUserIdResult {
	return &p
}

type SetFormWithSignatureResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type SetFormWithSignatureAsyncResult struct {
	result *SetFormWithSignatureResult
	err    error
}

func NewSetFormWithSignatureResultFromDict(data map[string]interface{}) SetFormWithSignatureResult {
	return SetFormWithSignatureResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p SetFormWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p SetFormWithSignatureResult) Pointer() *SetFormWithSignatureResult {
	return &p
}

type AcquireActionsToFormPropertiesResult struct {
	Item                      *Form   `json:"item"`
	Mold                      *Mold   `json:"mold"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type AcquireActionsToFormPropertiesAsyncResult struct {
	result *AcquireActionsToFormPropertiesResult
	err    error
}

func NewAcquireActionsToFormPropertiesResultFromDict(data map[string]interface{}) AcquireActionsToFormPropertiesResult {
	return AcquireActionsToFormPropertiesResult{
		Item:                      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:                      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p AcquireActionsToFormPropertiesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"mold":                      p.Mold.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p AcquireActionsToFormPropertiesResult) Pointer() *AcquireActionsToFormPropertiesResult {
	return &p
}

type DeleteFormResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type DeleteFormAsyncResult struct {
	result *DeleteFormResult
	err    error
}

func NewDeleteFormResultFromDict(data map[string]interface{}) DeleteFormResult {
	return DeleteFormResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p DeleteFormResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p DeleteFormResult) Pointer() *DeleteFormResult {
	return &p
}

type DeleteFormByUserIdResult struct {
	Item      *Form      `json:"item"`
	Mold      *Mold      `json:"mold"`
	MoldModel *MoldModel `json:"moldModel"`
	FormModel *FormModel `json:"formModel"`
}

type DeleteFormByUserIdAsyncResult struct {
	result *DeleteFormByUserIdResult
	err    error
}

func NewDeleteFormByUserIdResultFromDict(data map[string]interface{}) DeleteFormByUserIdResult {
	return DeleteFormByUserIdResult{
		Item:      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		FormModel: NewFormModelFromDict(core.CastMap(data["formModel"])).Pointer(),
	}
}

func (p DeleteFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"mold":      p.Mold.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
		"formModel": p.FormModel.ToDict(),
	}
}

func (p DeleteFormByUserIdResult) Pointer() *DeleteFormByUserIdResult {
	return &p
}

type AcquireActionToFormPropertiesByStampSheetResult struct {
	Item                      *Form   `json:"item"`
	Mold                      *Mold   `json:"mold"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type AcquireActionToFormPropertiesByStampSheetAsyncResult struct {
	result *AcquireActionToFormPropertiesByStampSheetResult
	err    error
}

func NewAcquireActionToFormPropertiesByStampSheetResultFromDict(data map[string]interface{}) AcquireActionToFormPropertiesByStampSheetResult {
	return AcquireActionToFormPropertiesByStampSheetResult{
		Item:                      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:                      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p AcquireActionToFormPropertiesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"mold":                      p.Mold.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p AcquireActionToFormPropertiesByStampSheetResult) Pointer() *AcquireActionToFormPropertiesByStampSheetResult {
	return &p
}
