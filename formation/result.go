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

type GetFormModelResult struct {
	Item *FormModel `json:"item"`
}

type GetFormModelAsyncResult struct {
	result *GetFormModelResult
	err    error
}

func NewGetFormModelResultFromJson(data string) GetFormModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormModelResultFromDict(dict)
}

func NewGetFormModelResultFromDict(data map[string]interface{}) GetFormModelResult {
	return GetFormModelResult{
		Item: NewFormModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetFormModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetFormModelResult) Pointer() *GetFormModelResult {
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

func NewDescribeFormModelMastersResultFromJson(data string) DescribeFormModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormModelMastersResultFromDict(dict)
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

func NewCreateFormModelMasterResultFromJson(data string) CreateFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateFormModelMasterResultFromDict(dict)
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

func NewGetFormModelMasterResultFromJson(data string) GetFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormModelMasterResultFromDict(dict)
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

func NewUpdateFormModelMasterResultFromJson(data string) UpdateFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateFormModelMasterResultFromDict(dict)
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

func NewDeleteFormModelMasterResultFromJson(data string) DeleteFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormModelMasterResultFromDict(dict)
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

func NewDescribeMoldModelsResultFromJson(data string) DescribeMoldModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldModelsResultFromDict(dict)
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

func NewGetMoldModelResultFromJson(data string) GetMoldModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldModelResultFromDict(dict)
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

func NewDescribeMoldModelMastersResultFromJson(data string) DescribeMoldModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldModelMastersResultFromDict(dict)
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

func NewCreateMoldModelMasterResultFromJson(data string) CreateMoldModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMoldModelMasterResultFromDict(dict)
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

func NewGetMoldModelMasterResultFromJson(data string) GetMoldModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldModelMasterResultFromDict(dict)
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

func NewUpdateMoldModelMasterResultFromJson(data string) UpdateMoldModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMoldModelMasterResultFromDict(dict)
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

func NewDeleteMoldModelMasterResultFromJson(data string) DeleteMoldModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldModelMasterResultFromDict(dict)
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

type DescribePropertyFormModelsResult struct {
	Items []PropertyFormModel `json:"items"`
}

type DescribePropertyFormModelsAsyncResult struct {
	result *DescribePropertyFormModelsResult
	err    error
}

func NewDescribePropertyFormModelsResultFromJson(data string) DescribePropertyFormModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormModelsResultFromDict(dict)
}

func NewDescribePropertyFormModelsResultFromDict(data map[string]interface{}) DescribePropertyFormModelsResult {
	return DescribePropertyFormModelsResult{
		Items: CastPropertyFormModels(core.CastArray(data["items"])),
	}
}

func (p DescribePropertyFormModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPropertyFormModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribePropertyFormModelsResult) Pointer() *DescribePropertyFormModelsResult {
	return &p
}

type GetPropertyFormModelResult struct {
	Item *PropertyFormModel `json:"item"`
}

type GetPropertyFormModelAsyncResult struct {
	result *GetPropertyFormModelResult
	err    error
}

func NewGetPropertyFormModelResultFromJson(data string) GetPropertyFormModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormModelResultFromDict(dict)
}

func NewGetPropertyFormModelResultFromDict(data map[string]interface{}) GetPropertyFormModelResult {
	return GetPropertyFormModelResult{
		Item: NewPropertyFormModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetPropertyFormModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetPropertyFormModelResult) Pointer() *GetPropertyFormModelResult {
	return &p
}

type DescribePropertyFormModelMastersResult struct {
	Items         []PropertyFormModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribePropertyFormModelMastersAsyncResult struct {
	result *DescribePropertyFormModelMastersResult
	err    error
}

func NewDescribePropertyFormModelMastersResultFromJson(data string) DescribePropertyFormModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormModelMastersResultFromDict(dict)
}

func NewDescribePropertyFormModelMastersResultFromDict(data map[string]interface{}) DescribePropertyFormModelMastersResult {
	return DescribePropertyFormModelMastersResult{
		Items:         CastPropertyFormModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribePropertyFormModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPropertyFormModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePropertyFormModelMastersResult) Pointer() *DescribePropertyFormModelMastersResult {
	return &p
}

type CreatePropertyFormModelMasterResult struct {
	Item *PropertyFormModelMaster `json:"item"`
}

type CreatePropertyFormModelMasterAsyncResult struct {
	result *CreatePropertyFormModelMasterResult
	err    error
}

func NewCreatePropertyFormModelMasterResultFromJson(data string) CreatePropertyFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePropertyFormModelMasterResultFromDict(dict)
}

func NewCreatePropertyFormModelMasterResultFromDict(data map[string]interface{}) CreatePropertyFormModelMasterResult {
	return CreatePropertyFormModelMasterResult{
		Item: NewPropertyFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreatePropertyFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreatePropertyFormModelMasterResult) Pointer() *CreatePropertyFormModelMasterResult {
	return &p
}

type GetPropertyFormModelMasterResult struct {
	Item *PropertyFormModelMaster `json:"item"`
}

type GetPropertyFormModelMasterAsyncResult struct {
	result *GetPropertyFormModelMasterResult
	err    error
}

func NewGetPropertyFormModelMasterResultFromJson(data string) GetPropertyFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormModelMasterResultFromDict(dict)
}

func NewGetPropertyFormModelMasterResultFromDict(data map[string]interface{}) GetPropertyFormModelMasterResult {
	return GetPropertyFormModelMasterResult{
		Item: NewPropertyFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetPropertyFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetPropertyFormModelMasterResult) Pointer() *GetPropertyFormModelMasterResult {
	return &p
}

type UpdatePropertyFormModelMasterResult struct {
	Item *PropertyFormModelMaster `json:"item"`
}

type UpdatePropertyFormModelMasterAsyncResult struct {
	result *UpdatePropertyFormModelMasterResult
	err    error
}

func NewUpdatePropertyFormModelMasterResultFromJson(data string) UpdatePropertyFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdatePropertyFormModelMasterResultFromDict(dict)
}

func NewUpdatePropertyFormModelMasterResultFromDict(data map[string]interface{}) UpdatePropertyFormModelMasterResult {
	return UpdatePropertyFormModelMasterResult{
		Item: NewPropertyFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdatePropertyFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdatePropertyFormModelMasterResult) Pointer() *UpdatePropertyFormModelMasterResult {
	return &p
}

type DeletePropertyFormModelMasterResult struct {
	Item *PropertyFormModelMaster `json:"item"`
}

type DeletePropertyFormModelMasterAsyncResult struct {
	result *DeletePropertyFormModelMasterResult
	err    error
}

func NewDeletePropertyFormModelMasterResultFromJson(data string) DeletePropertyFormModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormModelMasterResultFromDict(dict)
}

func NewDeletePropertyFormModelMasterResultFromDict(data map[string]interface{}) DeletePropertyFormModelMasterResult {
	return DeletePropertyFormModelMasterResult{
		Item: NewPropertyFormModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeletePropertyFormModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeletePropertyFormModelMasterResult) Pointer() *DeletePropertyFormModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentFormMaster `json:"item"`
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

func NewGetCurrentFormMasterResultFromJson(data string) GetCurrentFormMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentFormMasterResultFromDict(dict)
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

func NewUpdateCurrentFormMasterResultFromJson(data string) UpdateCurrentFormMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFormMasterResultFromDict(dict)
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

func NewUpdateCurrentFormMasterFromGitHubResultFromJson(data string) UpdateCurrentFormMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentFormMasterFromGitHubResultFromDict(dict)
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

func NewDescribeMoldsResultFromJson(data string) DescribeMoldsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldsResultFromDict(dict)
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

func NewDescribeMoldsByUserIdResultFromJson(data string) DescribeMoldsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMoldsByUserIdResultFromDict(dict)
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

func NewGetMoldResultFromJson(data string) GetMoldResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldResultFromDict(dict)
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

func NewGetMoldByUserIdResultFromJson(data string) GetMoldByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMoldByUserIdResultFromDict(dict)
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
	Old       *Mold      `json:"old"`
	MoldModel *MoldModel `json:"moldModel"`
}

type SetMoldCapacityByUserIdAsyncResult struct {
	result *SetMoldCapacityByUserIdResult
	err    error
}

func NewSetMoldCapacityByUserIdResultFromJson(data string) SetMoldCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMoldCapacityByUserIdResultFromDict(dict)
}

func NewSetMoldCapacityByUserIdResultFromDict(data map[string]interface{}) SetMoldCapacityByUserIdResult {
	return SetMoldCapacityByUserIdResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		Old:       NewMoldFromDict(core.CastMap(data["old"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p SetMoldCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"old":       p.Old.ToDict(),
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

func NewAddMoldCapacityByUserIdResultFromJson(data string) AddMoldCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddMoldCapacityByUserIdResultFromDict(dict)
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

type SubMoldCapacityByUserIdResult struct {
	Item      *Mold      `json:"item"`
	MoldModel *MoldModel `json:"moldModel"`
}

type SubMoldCapacityByUserIdAsyncResult struct {
	result *SubMoldCapacityByUserIdResult
	err    error
}

func NewSubMoldCapacityByUserIdResultFromJson(data string) SubMoldCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubMoldCapacityByUserIdResultFromDict(dict)
}

func NewSubMoldCapacityByUserIdResultFromDict(data map[string]interface{}) SubMoldCapacityByUserIdResult {
	return SubMoldCapacityByUserIdResult{
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p SubMoldCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"moldModel": p.MoldModel.ToDict(),
	}
}

func (p SubMoldCapacityByUserIdResult) Pointer() *SubMoldCapacityByUserIdResult {
	return &p
}

type DeleteMoldResult struct {
	Item *Mold `json:"item"`
}

type DeleteMoldAsyncResult struct {
	result *DeleteMoldResult
	err    error
}

func NewDeleteMoldResultFromJson(data string) DeleteMoldResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldResultFromDict(dict)
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

func NewDeleteMoldByUserIdResultFromJson(data string) DeleteMoldByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMoldByUserIdResultFromDict(dict)
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

func NewAddCapacityByStampSheetResultFromJson(data string) AddCapacityByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddCapacityByStampSheetResultFromDict(dict)
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

type SubCapacityByStampTaskResult struct {
	Item            *Mold      `json:"item"`
	MoldModel       *MoldModel `json:"moldModel"`
	NewContextStack *string    `json:"newContextStack"`
}

type SubCapacityByStampTaskAsyncResult struct {
	result *SubCapacityByStampTaskResult
	err    error
}

func NewSubCapacityByStampTaskResultFromJson(data string) SubCapacityByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubCapacityByStampTaskResultFromDict(dict)
}

func NewSubCapacityByStampTaskResultFromDict(data map[string]interface{}) SubCapacityByStampTaskResult {
	return SubCapacityByStampTaskResult{
		Item:            NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		MoldModel:       NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p SubCapacityByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"moldModel":       p.MoldModel.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p SubCapacityByStampTaskResult) Pointer() *SubCapacityByStampTaskResult {
	return &p
}

type SetCapacityByStampSheetResult struct {
	Item      *Mold      `json:"item"`
	Old       *Mold      `json:"old"`
	MoldModel *MoldModel `json:"moldModel"`
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
		Item:      NewMoldFromDict(core.CastMap(data["item"])).Pointer(),
		Old:       NewMoldFromDict(core.CastMap(data["old"])).Pointer(),
		MoldModel: NewMoldModelFromDict(core.CastMap(data["moldModel"])).Pointer(),
	}
}

func (p SetCapacityByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"old":       p.Old.ToDict(),
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

func NewDescribeFormsResultFromJson(data string) DescribeFormsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormsResultFromDict(dict)
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

func NewDescribeFormsByUserIdResultFromJson(data string) DescribeFormsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFormsByUserIdResultFromDict(dict)
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

func NewGetFormResultFromJson(data string) GetFormResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormResultFromDict(dict)
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

func NewGetFormByUserIdResultFromJson(data string) GetFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormByUserIdResultFromDict(dict)
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

func NewGetFormWithSignatureResultFromJson(data string) GetFormWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormWithSignatureResultFromDict(dict)
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

func NewGetFormWithSignatureByUserIdResultFromJson(data string) GetFormWithSignatureByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFormWithSignatureByUserIdResultFromDict(dict)
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

func NewSetFormByUserIdResultFromJson(data string) SetFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFormByUserIdResultFromDict(dict)
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

func NewSetFormWithSignatureResultFromJson(data string) SetFormWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetFormWithSignatureResultFromDict(dict)
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
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type AcquireActionsToFormPropertiesAsyncResult struct {
	result *AcquireActionsToFormPropertiesResult
	err    error
}

func NewAcquireActionsToFormPropertiesResultFromJson(data string) AcquireActionsToFormPropertiesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionsToFormPropertiesResultFromDict(dict)
}

func NewAcquireActionsToFormPropertiesResultFromDict(data map[string]interface{}) AcquireActionsToFormPropertiesResult {
	return AcquireActionsToFormPropertiesResult{
		Item:                      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:                      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p AcquireActionsToFormPropertiesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"mold":                      p.Mold.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
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

func NewDeleteFormResultFromJson(data string) DeleteFormResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormResultFromDict(dict)
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

func NewDeleteFormByUserIdResultFromJson(data string) DeleteFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFormByUserIdResultFromDict(dict)
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
	TransactionId             *string `json:"transactionId"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool   `json:"autoRunStampSheet"`
}

type AcquireActionToFormPropertiesByStampSheetAsyncResult struct {
	result *AcquireActionToFormPropertiesByStampSheetResult
	err    error
}

func NewAcquireActionToFormPropertiesByStampSheetResultFromJson(data string) AcquireActionToFormPropertiesByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionToFormPropertiesByStampSheetResultFromDict(dict)
}

func NewAcquireActionToFormPropertiesByStampSheetResultFromDict(data map[string]interface{}) AcquireActionToFormPropertiesByStampSheetResult {
	return AcquireActionToFormPropertiesByStampSheetResult{
		Item:                      NewFormFromDict(core.CastMap(data["item"])).Pointer(),
		Mold:                      NewMoldFromDict(core.CastMap(data["mold"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p AcquireActionToFormPropertiesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"mold":                      p.Mold.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p AcquireActionToFormPropertiesByStampSheetResult) Pointer() *AcquireActionToFormPropertiesByStampSheetResult {
	return &p
}

type DescribePropertyFormsResult struct {
	Items         []PropertyForm `json:"items"`
	NextPageToken *string        `json:"nextPageToken"`
}

type DescribePropertyFormsAsyncResult struct {
	result *DescribePropertyFormsResult
	err    error
}

func NewDescribePropertyFormsResultFromJson(data string) DescribePropertyFormsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormsResultFromDict(dict)
}

func NewDescribePropertyFormsResultFromDict(data map[string]interface{}) DescribePropertyFormsResult {
	return DescribePropertyFormsResult{
		Items:         CastPropertyForms(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribePropertyFormsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPropertyFormsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePropertyFormsResult) Pointer() *DescribePropertyFormsResult {
	return &p
}

type DescribePropertyFormsByUserIdResult struct {
	Items         []PropertyForm `json:"items"`
	NextPageToken *string        `json:"nextPageToken"`
}

type DescribePropertyFormsByUserIdAsyncResult struct {
	result *DescribePropertyFormsByUserIdResult
	err    error
}

func NewDescribePropertyFormsByUserIdResultFromJson(data string) DescribePropertyFormsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePropertyFormsByUserIdResultFromDict(dict)
}

func NewDescribePropertyFormsByUserIdResultFromDict(data map[string]interface{}) DescribePropertyFormsByUserIdResult {
	return DescribePropertyFormsByUserIdResult{
		Items:         CastPropertyForms(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribePropertyFormsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPropertyFormsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePropertyFormsByUserIdResult) Pointer() *DescribePropertyFormsByUserIdResult {
	return &p
}

type GetPropertyFormResult struct {
	Item              *PropertyForm      `json:"item"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type GetPropertyFormAsyncResult struct {
	result *GetPropertyFormResult
	err    error
}

func NewGetPropertyFormResultFromJson(data string) GetPropertyFormResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormResultFromDict(dict)
}

func NewGetPropertyFormResultFromDict(data map[string]interface{}) GetPropertyFormResult {
	return GetPropertyFormResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p GetPropertyFormResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p GetPropertyFormResult) Pointer() *GetPropertyFormResult {
	return &p
}

type GetPropertyFormByUserIdResult struct {
	Item              *PropertyForm      `json:"item"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type GetPropertyFormByUserIdAsyncResult struct {
	result *GetPropertyFormByUserIdResult
	err    error
}

func NewGetPropertyFormByUserIdResultFromJson(data string) GetPropertyFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormByUserIdResultFromDict(dict)
}

func NewGetPropertyFormByUserIdResultFromDict(data map[string]interface{}) GetPropertyFormByUserIdResult {
	return GetPropertyFormByUserIdResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p GetPropertyFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p GetPropertyFormByUserIdResult) Pointer() *GetPropertyFormByUserIdResult {
	return &p
}

type GetPropertyFormWithSignatureResult struct {
	Item              *PropertyForm      `json:"item"`
	Body              *string            `json:"body"`
	Signature         *string            `json:"signature"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type GetPropertyFormWithSignatureAsyncResult struct {
	result *GetPropertyFormWithSignatureResult
	err    error
}

func NewGetPropertyFormWithSignatureResultFromJson(data string) GetPropertyFormWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormWithSignatureResultFromDict(dict)
}

func NewGetPropertyFormWithSignatureResultFromDict(data map[string]interface{}) GetPropertyFormWithSignatureResult {
	return GetPropertyFormWithSignatureResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		Body:              core.CastString(data["body"]),
		Signature:         core.CastString(data["signature"]),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p GetPropertyFormWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"body":              p.Body,
		"signature":         p.Signature,
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p GetPropertyFormWithSignatureResult) Pointer() *GetPropertyFormWithSignatureResult {
	return &p
}

type GetPropertyFormWithSignatureByUserIdResult struct {
	Item              *PropertyForm      `json:"item"`
	Body              *string            `json:"body"`
	Signature         *string            `json:"signature"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type GetPropertyFormWithSignatureByUserIdAsyncResult struct {
	result *GetPropertyFormWithSignatureByUserIdResult
	err    error
}

func NewGetPropertyFormWithSignatureByUserIdResultFromJson(data string) GetPropertyFormWithSignatureByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPropertyFormWithSignatureByUserIdResultFromDict(dict)
}

func NewGetPropertyFormWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetPropertyFormWithSignatureByUserIdResult {
	return GetPropertyFormWithSignatureByUserIdResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		Body:              core.CastString(data["body"]),
		Signature:         core.CastString(data["signature"]),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p GetPropertyFormWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"body":              p.Body,
		"signature":         p.Signature,
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p GetPropertyFormWithSignatureByUserIdResult) Pointer() *GetPropertyFormWithSignatureByUserIdResult {
	return &p
}

type SetPropertyFormByUserIdResult struct {
	Item              *PropertyForm      `json:"item"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type SetPropertyFormByUserIdAsyncResult struct {
	result *SetPropertyFormByUserIdResult
	err    error
}

func NewSetPropertyFormByUserIdResultFromJson(data string) SetPropertyFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetPropertyFormByUserIdResultFromDict(dict)
}

func NewSetPropertyFormByUserIdResultFromDict(data map[string]interface{}) SetPropertyFormByUserIdResult {
	return SetPropertyFormByUserIdResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p SetPropertyFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p SetPropertyFormByUserIdResult) Pointer() *SetPropertyFormByUserIdResult {
	return &p
}

type SetPropertyFormWithSignatureResult struct {
	Item              *PropertyForm      `json:"item"`
	ProeprtyFormModel *PropertyFormModel `json:"proeprtyFormModel"`
}

type SetPropertyFormWithSignatureAsyncResult struct {
	result *SetPropertyFormWithSignatureResult
	err    error
}

func NewSetPropertyFormWithSignatureResultFromJson(data string) SetPropertyFormWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetPropertyFormWithSignatureResultFromDict(dict)
}

func NewSetPropertyFormWithSignatureResultFromDict(data map[string]interface{}) SetPropertyFormWithSignatureResult {
	return SetPropertyFormWithSignatureResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		ProeprtyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["proeprtyFormModel"])).Pointer(),
	}
}

func (p SetPropertyFormWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"proeprtyFormModel": p.ProeprtyFormModel.ToDict(),
	}
}

func (p SetPropertyFormWithSignatureResult) Pointer() *SetPropertyFormWithSignatureResult {
	return &p
}

type AcquireActionsToPropertyFormPropertiesResult struct {
	Item                      *PropertyForm `json:"item"`
	TransactionId             *string       `json:"transactionId"`
	StampSheet                *string       `json:"stampSheet"`
	StampSheetEncryptionKeyId *string       `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool         `json:"autoRunStampSheet"`
}

type AcquireActionsToPropertyFormPropertiesAsyncResult struct {
	result *AcquireActionsToPropertyFormPropertiesResult
	err    error
}

func NewAcquireActionsToPropertyFormPropertiesResultFromJson(data string) AcquireActionsToPropertyFormPropertiesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionsToPropertyFormPropertiesResultFromDict(dict)
}

func NewAcquireActionsToPropertyFormPropertiesResultFromDict(data map[string]interface{}) AcquireActionsToPropertyFormPropertiesResult {
	return AcquireActionsToPropertyFormPropertiesResult{
		Item:                      NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p AcquireActionsToPropertyFormPropertiesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p AcquireActionsToPropertyFormPropertiesResult) Pointer() *AcquireActionsToPropertyFormPropertiesResult {
	return &p
}

type DeletePropertyFormResult struct {
	Item              *PropertyForm      `json:"item"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type DeletePropertyFormAsyncResult struct {
	result *DeletePropertyFormResult
	err    error
}

func NewDeletePropertyFormResultFromJson(data string) DeletePropertyFormResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormResultFromDict(dict)
}

func NewDeletePropertyFormResultFromDict(data map[string]interface{}) DeletePropertyFormResult {
	return DeletePropertyFormResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p DeletePropertyFormResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p DeletePropertyFormResult) Pointer() *DeletePropertyFormResult {
	return &p
}

type DeletePropertyFormByUserIdResult struct {
	Item              *PropertyForm      `json:"item"`
	PropertyFormModel *PropertyFormModel `json:"propertyFormModel"`
}

type DeletePropertyFormByUserIdAsyncResult struct {
	result *DeletePropertyFormByUserIdResult
	err    error
}

func NewDeletePropertyFormByUserIdResultFromJson(data string) DeletePropertyFormByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePropertyFormByUserIdResultFromDict(dict)
}

func NewDeletePropertyFormByUserIdResultFromDict(data map[string]interface{}) DeletePropertyFormByUserIdResult {
	return DeletePropertyFormByUserIdResult{
		Item:              NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		PropertyFormModel: NewPropertyFormModelFromDict(core.CastMap(data["propertyFormModel"])).Pointer(),
	}
}

func (p DeletePropertyFormByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":              p.Item.ToDict(),
		"propertyFormModel": p.PropertyFormModel.ToDict(),
	}
}

func (p DeletePropertyFormByUserIdResult) Pointer() *DeletePropertyFormByUserIdResult {
	return &p
}

type AcquireActionToPropertyFormPropertiesByStampSheetResult struct {
	Item                      *PropertyForm `json:"item"`
	TransactionId             *string       `json:"transactionId"`
	StampSheet                *string       `json:"stampSheet"`
	StampSheetEncryptionKeyId *string       `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool         `json:"autoRunStampSheet"`
}

type AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult struct {
	result *AcquireActionToPropertyFormPropertiesByStampSheetResult
	err    error
}

func NewAcquireActionToPropertyFormPropertiesByStampSheetResultFromJson(data string) AcquireActionToPropertyFormPropertiesByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireActionToPropertyFormPropertiesByStampSheetResultFromDict(dict)
}

func NewAcquireActionToPropertyFormPropertiesByStampSheetResultFromDict(data map[string]interface{}) AcquireActionToPropertyFormPropertiesByStampSheetResult {
	return AcquireActionToPropertyFormPropertiesByStampSheetResult{
		Item:                      NewPropertyFormFromDict(core.CastMap(data["item"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p AcquireActionToPropertyFormPropertiesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p AcquireActionToPropertyFormPropertiesByStampSheetResult) Pointer() *AcquireActionToPropertyFormPropertiesByStampSheetResult {
	return &p
}
