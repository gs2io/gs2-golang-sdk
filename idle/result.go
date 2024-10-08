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

package idle

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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DumpUserDataByUserIdResult struct {
}

type DumpUserDataByUserIdAsyncResult struct {
	result *DumpUserDataByUserIdResult
	err    error
}

func NewDumpUserDataByUserIdResultFromJson(data string) DumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdResultFromDict(dict)
}

func NewDumpUserDataByUserIdResultFromDict(data map[string]interface{}) DumpUserDataByUserIdResult {
	return DumpUserDataByUserIdResult{}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DumpUserDataByUserIdResult) Pointer() *DumpUserDataByUserIdResult {
	return &p
}

type CheckDumpUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckDumpUserDataByUserIdAsyncResult struct {
	result *CheckDumpUserDataByUserIdResult
	err    error
}

func NewCheckDumpUserDataByUserIdResultFromJson(data string) CheckDumpUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdResultFromDict(dict)
}

func NewCheckDumpUserDataByUserIdResultFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdResult {
	return CheckDumpUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckDumpUserDataByUserIdResult) Pointer() *CheckDumpUserDataByUserIdResult {
	return &p
}

type CleanUserDataByUserIdResult struct {
}

type CleanUserDataByUserIdAsyncResult struct {
	result *CleanUserDataByUserIdResult
	err    error
}

func NewCleanUserDataByUserIdResultFromJson(data string) CleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdResultFromDict(dict)
}

func NewCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CleanUserDataByUserIdResult {
	return CleanUserDataByUserIdResult{}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CleanUserDataByUserIdResult) Pointer() *CleanUserDataByUserIdResult {
	return &p
}

type CheckCleanUserDataByUserIdResult struct {
}

type CheckCleanUserDataByUserIdAsyncResult struct {
	result *CheckCleanUserDataByUserIdResult
	err    error
}

func NewCheckCleanUserDataByUserIdResultFromJson(data string) CheckCleanUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdResultFromDict(dict)
}

func NewCheckCleanUserDataByUserIdResultFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdResult {
	return CheckCleanUserDataByUserIdResult{}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CheckCleanUserDataByUserIdResult) Pointer() *CheckCleanUserDataByUserIdResult {
	return &p
}

type PrepareImportUserDataByUserIdResult struct {
	UploadToken *string `json:"uploadToken"`
	UploadUrl   *string `json:"uploadUrl"`
}

type PrepareImportUserDataByUserIdAsyncResult struct {
	result *PrepareImportUserDataByUserIdResult
	err    error
}

func NewPrepareImportUserDataByUserIdResultFromJson(data string) PrepareImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdResultFromDict(dict)
}

func NewPrepareImportUserDataByUserIdResultFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdResult {
	return PrepareImportUserDataByUserIdResult{
		UploadToken: core.CastString(data["uploadToken"]),
		UploadUrl:   core.CastString(data["uploadUrl"]),
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
	}
}

func (p PrepareImportUserDataByUserIdResult) Pointer() *PrepareImportUserDataByUserIdResult {
	return &p
}

type ImportUserDataByUserIdResult struct {
}

type ImportUserDataByUserIdAsyncResult struct {
	result *ImportUserDataByUserIdResult
	err    error
}

func NewImportUserDataByUserIdResultFromJson(data string) ImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdResultFromDict(dict)
}

func NewImportUserDataByUserIdResultFromDict(data map[string]interface{}) ImportUserDataByUserIdResult {
	return ImportUserDataByUserIdResult{}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ImportUserDataByUserIdResult) Pointer() *ImportUserDataByUserIdResult {
	return &p
}

type CheckImportUserDataByUserIdResult struct {
	Url *string `json:"url"`
}

type CheckImportUserDataByUserIdAsyncResult struct {
	result *CheckImportUserDataByUserIdResult
	err    error
}

func NewCheckImportUserDataByUserIdResultFromJson(data string) CheckImportUserDataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdResultFromDict(dict)
}

func NewCheckImportUserDataByUserIdResultFromDict(data map[string]interface{}) CheckImportUserDataByUserIdResult {
	return CheckImportUserDataByUserIdResult{
		Url: core.CastString(data["url"]),
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeCategoryModelMastersResult struct {
	Items         []CategoryModelMaster `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
}

type DescribeCategoryModelMastersAsyncResult struct {
	result *DescribeCategoryModelMastersResult
	err    error
}

func NewDescribeCategoryModelMastersResultFromJson(data string) DescribeCategoryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelMastersResultFromDict(dict)
}

func NewDescribeCategoryModelMastersResultFromDict(data map[string]interface{}) DescribeCategoryModelMastersResult {
	return DescribeCategoryModelMastersResult{
		Items:         CastCategoryModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCategoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCategoryModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCategoryModelMastersResult) Pointer() *DescribeCategoryModelMastersResult {
	return &p
}

type CreateCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type CreateCategoryModelMasterAsyncResult struct {
	result *CreateCategoryModelMasterResult
	err    error
}

func NewCreateCategoryModelMasterResultFromJson(data string) CreateCategoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCategoryModelMasterResultFromDict(dict)
}

func NewCreateCategoryModelMasterResultFromDict(data map[string]interface{}) CreateCategoryModelMasterResult {
	return CreateCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateCategoryModelMasterResult) Pointer() *CreateCategoryModelMasterResult {
	return &p
}

type GetCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type GetCategoryModelMasterAsyncResult struct {
	result *GetCategoryModelMasterResult
	err    error
}

func NewGetCategoryModelMasterResultFromJson(data string) GetCategoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelMasterResultFromDict(dict)
}

func NewGetCategoryModelMasterResultFromDict(data map[string]interface{}) GetCategoryModelMasterResult {
	return GetCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCategoryModelMasterResult) Pointer() *GetCategoryModelMasterResult {
	return &p
}

type UpdateCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type UpdateCategoryModelMasterAsyncResult struct {
	result *UpdateCategoryModelMasterResult
	err    error
}

func NewUpdateCategoryModelMasterResultFromJson(data string) UpdateCategoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCategoryModelMasterResultFromDict(dict)
}

func NewUpdateCategoryModelMasterResultFromDict(data map[string]interface{}) UpdateCategoryModelMasterResult {
	return UpdateCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCategoryModelMasterResult) Pointer() *UpdateCategoryModelMasterResult {
	return &p
}

type DeleteCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type DeleteCategoryModelMasterAsyncResult struct {
	result *DeleteCategoryModelMasterResult
	err    error
}

func NewDeleteCategoryModelMasterResultFromJson(data string) DeleteCategoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCategoryModelMasterResultFromDict(dict)
}

func NewDeleteCategoryModelMasterResultFromDict(data map[string]interface{}) DeleteCategoryModelMasterResult {
	return DeleteCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteCategoryModelMasterResult) Pointer() *DeleteCategoryModelMasterResult {
	return &p
}

type DescribeCategoryModelsResult struct {
	Items []CategoryModel `json:"items"`
}

type DescribeCategoryModelsAsyncResult struct {
	result *DescribeCategoryModelsResult
	err    error
}

func NewDescribeCategoryModelsResultFromJson(data string) DescribeCategoryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCategoryModelsResultFromDict(dict)
}

func NewDescribeCategoryModelsResultFromDict(data map[string]interface{}) DescribeCategoryModelsResult {
	return DescribeCategoryModelsResult{
		Items: CastCategoryModels(core.CastArray(data["items"])),
	}
}

func (p DescribeCategoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCategoryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeCategoryModelsResult) Pointer() *DescribeCategoryModelsResult {
	return &p
}

type GetCategoryModelResult struct {
	Item *CategoryModel `json:"item"`
}

type GetCategoryModelAsyncResult struct {
	result *GetCategoryModelResult
	err    error
}

func NewGetCategoryModelResultFromJson(data string) GetCategoryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCategoryModelResultFromDict(dict)
}

func NewGetCategoryModelResultFromDict(data map[string]interface{}) GetCategoryModelResult {
	return GetCategoryModelResult{
		Item: NewCategoryModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCategoryModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCategoryModelResult) Pointer() *GetCategoryModelResult {
	return &p
}

type DescribeStatusesResult struct {
	Items         []Status `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeStatusesAsyncResult struct {
	result *DescribeStatusesResult
	err    error
}

func NewDescribeStatusesResultFromJson(data string) DescribeStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesResultFromDict(dict)
}

func NewDescribeStatusesResultFromDict(data map[string]interface{}) DescribeStatusesResult {
	return DescribeStatusesResult{
		Items:         CastStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStatusesResult) Pointer() *DescribeStatusesResult {
	return &p
}

type DescribeStatusesByUserIdResult struct {
	Items         []Status `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeStatusesByUserIdAsyncResult struct {
	result *DescribeStatusesByUserIdResult
	err    error
}

func NewDescribeStatusesByUserIdResultFromJson(data string) DescribeStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesByUserIdResultFromDict(dict)
}

func NewDescribeStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeStatusesByUserIdResult {
	return DescribeStatusesByUserIdResult{
		Items:         CastStatuses(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStatusesByUserIdResult) Pointer() *DescribeStatusesByUserIdResult {
	return &p
}

type GetStatusResult struct {
	Item *Status `json:"item"`
}

type GetStatusAsyncResult struct {
	result *GetStatusResult
	err    error
}

func NewGetStatusResultFromJson(data string) GetStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusResultFromDict(dict)
}

func NewGetStatusResultFromDict(data map[string]interface{}) GetStatusResult {
	return GetStatusResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStatusResult) Pointer() *GetStatusResult {
	return &p
}

type GetStatusByUserIdResult struct {
	Item *Status `json:"item"`
}

type GetStatusByUserIdAsyncResult struct {
	result *GetStatusByUserIdResult
	err    error
}

func NewGetStatusByUserIdResultFromJson(data string) GetStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusByUserIdResultFromDict(dict)
}

func NewGetStatusByUserIdResultFromDict(data map[string]interface{}) GetStatusByUserIdResult {
	return GetStatusByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStatusByUserIdResult) Pointer() *GetStatusByUserIdResult {
	return &p
}

type PredictionResult struct {
	Items  []AcquireAction `json:"items"`
	Status *Status         `json:"status"`
}

type PredictionAsyncResult struct {
	result *PredictionResult
	err    error
}

func NewPredictionResultFromJson(data string) PredictionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionResultFromDict(dict)
}

func NewPredictionResultFromDict(data map[string]interface{}) PredictionResult {
	return PredictionResult{
		Items:  CastAcquireActions(core.CastArray(data["items"])),
		Status: NewStatusFromDict(core.CastMap(data["status"])).Pointer(),
	}
}

func (p PredictionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
	}
}

func (p PredictionResult) Pointer() *PredictionResult {
	return &p
}

type PredictionByUserIdResult struct {
	Items  []AcquireAction `json:"items"`
	Status *Status         `json:"status"`
}

type PredictionByUserIdAsyncResult struct {
	result *PredictionByUserIdResult
	err    error
}

func NewPredictionByUserIdResultFromJson(data string) PredictionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionByUserIdResultFromDict(dict)
}

func NewPredictionByUserIdResultFromDict(data map[string]interface{}) PredictionByUserIdResult {
	return PredictionByUserIdResult{
		Items:  CastAcquireActions(core.CastArray(data["items"])),
		Status: NewStatusFromDict(core.CastMap(data["status"])).Pointer(),
	}
}

func (p PredictionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
	}
}

func (p PredictionByUserIdResult) Pointer() *PredictionByUserIdResult {
	return &p
}

type ReceiveResult struct {
	Items                     []AcquireAction `json:"items"`
	Status                    *Status         `json:"status"`
	TransactionId             *string         `json:"transactionId"`
	StampSheet                *string         `json:"stampSheet"`
	StampSheetEncryptionKeyId *string         `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool           `json:"autoRunStampSheet"`
}

type ReceiveAsyncResult struct {
	result *ReceiveResult
	err    error
}

func NewReceiveResultFromJson(data string) ReceiveResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveResultFromDict(dict)
}

func NewReceiveResultFromDict(data map[string]interface{}) ReceiveResult {
	return ReceiveResult{
		Items:                     CastAcquireActions(core.CastArray(data["items"])),
		Status:                    NewStatusFromDict(core.CastMap(data["status"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReceiveResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReceiveResult) Pointer() *ReceiveResult {
	return &p
}

type ReceiveByUserIdResult struct {
	Items                     []AcquireAction `json:"items"`
	Status                    *Status         `json:"status"`
	TransactionId             *string         `json:"transactionId"`
	StampSheet                *string         `json:"stampSheet"`
	StampSheetEncryptionKeyId *string         `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool           `json:"autoRunStampSheet"`
}

type ReceiveByUserIdAsyncResult struct {
	result *ReceiveByUserIdResult
	err    error
}

func NewReceiveByUserIdResultFromJson(data string) ReceiveByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByUserIdResultFromDict(dict)
}

func NewReceiveByUserIdResultFromDict(data map[string]interface{}) ReceiveByUserIdResult {
	return ReceiveByUserIdResult{
		Items:                     CastAcquireActions(core.CastArray(data["items"])),
		Status:                    NewStatusFromDict(core.CastMap(data["status"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReceiveByUserIdResult) Pointer() *ReceiveByUserIdResult {
	return &p
}

type IncreaseMaximumIdleMinutesByUserIdResult struct {
	Item *Status `json:"item"`
}

type IncreaseMaximumIdleMinutesByUserIdAsyncResult struct {
	result *IncreaseMaximumIdleMinutesByUserIdResult
	err    error
}

func NewIncreaseMaximumIdleMinutesByUserIdResultFromJson(data string) IncreaseMaximumIdleMinutesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumIdleMinutesByUserIdResultFromDict(dict)
}

func NewIncreaseMaximumIdleMinutesByUserIdResultFromDict(data map[string]interface{}) IncreaseMaximumIdleMinutesByUserIdResult {
	return IncreaseMaximumIdleMinutesByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p IncreaseMaximumIdleMinutesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p IncreaseMaximumIdleMinutesByUserIdResult) Pointer() *IncreaseMaximumIdleMinutesByUserIdResult {
	return &p
}

type DecreaseMaximumIdleMinutesResult struct {
	Item *Status `json:"item"`
}

type DecreaseMaximumIdleMinutesAsyncResult struct {
	result *DecreaseMaximumIdleMinutesResult
	err    error
}

func NewDecreaseMaximumIdleMinutesResultFromJson(data string) DecreaseMaximumIdleMinutesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumIdleMinutesResultFromDict(dict)
}

func NewDecreaseMaximumIdleMinutesResultFromDict(data map[string]interface{}) DecreaseMaximumIdleMinutesResult {
	return DecreaseMaximumIdleMinutesResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DecreaseMaximumIdleMinutesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DecreaseMaximumIdleMinutesResult) Pointer() *DecreaseMaximumIdleMinutesResult {
	return &p
}

type DecreaseMaximumIdleMinutesByUserIdResult struct {
	Item *Status `json:"item"`
}

type DecreaseMaximumIdleMinutesByUserIdAsyncResult struct {
	result *DecreaseMaximumIdleMinutesByUserIdResult
	err    error
}

func NewDecreaseMaximumIdleMinutesByUserIdResultFromJson(data string) DecreaseMaximumIdleMinutesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumIdleMinutesByUserIdResultFromDict(dict)
}

func NewDecreaseMaximumIdleMinutesByUserIdResultFromDict(data map[string]interface{}) DecreaseMaximumIdleMinutesByUserIdResult {
	return DecreaseMaximumIdleMinutesByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DecreaseMaximumIdleMinutesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DecreaseMaximumIdleMinutesByUserIdResult) Pointer() *DecreaseMaximumIdleMinutesByUserIdResult {
	return &p
}

type SetMaximumIdleMinutesByUserIdResult struct {
	Item *Status `json:"item"`
	Old  *Status `json:"old"`
}

type SetMaximumIdleMinutesByUserIdAsyncResult struct {
	result *SetMaximumIdleMinutesByUserIdResult
	err    error
}

func NewSetMaximumIdleMinutesByUserIdResultFromJson(data string) SetMaximumIdleMinutesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumIdleMinutesByUserIdResultFromDict(dict)
}

func NewSetMaximumIdleMinutesByUserIdResultFromDict(data map[string]interface{}) SetMaximumIdleMinutesByUserIdResult {
	return SetMaximumIdleMinutesByUserIdResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		Old:  NewStatusFromDict(core.CastMap(data["old"])).Pointer(),
	}
}

func (p SetMaximumIdleMinutesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"old": func() map[string]interface{} {
			if p.Old == nil {
				return nil
			}
			return p.Old.ToDict()
		}(),
	}
}

func (p SetMaximumIdleMinutesByUserIdResult) Pointer() *SetMaximumIdleMinutesByUserIdResult {
	return &p
}

type IncreaseMaximumIdleMinutesByStampSheetResult struct {
	Item *Status `json:"item"`
}

type IncreaseMaximumIdleMinutesByStampSheetAsyncResult struct {
	result *IncreaseMaximumIdleMinutesByStampSheetResult
	err    error
}

func NewIncreaseMaximumIdleMinutesByStampSheetResultFromJson(data string) IncreaseMaximumIdleMinutesByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumIdleMinutesByStampSheetResultFromDict(dict)
}

func NewIncreaseMaximumIdleMinutesByStampSheetResultFromDict(data map[string]interface{}) IncreaseMaximumIdleMinutesByStampSheetResult {
	return IncreaseMaximumIdleMinutesByStampSheetResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p IncreaseMaximumIdleMinutesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p IncreaseMaximumIdleMinutesByStampSheetResult) Pointer() *IncreaseMaximumIdleMinutesByStampSheetResult {
	return &p
}

type DecreaseMaximumIdleMinutesByStampTaskResult struct {
	Item            *Status `json:"item"`
	NewContextStack *string `json:"newContextStack"`
}

type DecreaseMaximumIdleMinutesByStampTaskAsyncResult struct {
	result *DecreaseMaximumIdleMinutesByStampTaskResult
	err    error
}

func NewDecreaseMaximumIdleMinutesByStampTaskResultFromJson(data string) DecreaseMaximumIdleMinutesByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumIdleMinutesByStampTaskResultFromDict(dict)
}

func NewDecreaseMaximumIdleMinutesByStampTaskResultFromDict(data map[string]interface{}) DecreaseMaximumIdleMinutesByStampTaskResult {
	return DecreaseMaximumIdleMinutesByStampTaskResult{
		Item:            NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p DecreaseMaximumIdleMinutesByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DecreaseMaximumIdleMinutesByStampTaskResult) Pointer() *DecreaseMaximumIdleMinutesByStampTaskResult {
	return &p
}

type SetMaximumIdleMinutesByStampSheetResult struct {
	Item *Status `json:"item"`
}

type SetMaximumIdleMinutesByStampSheetAsyncResult struct {
	result *SetMaximumIdleMinutesByStampSheetResult
	err    error
}

func NewSetMaximumIdleMinutesByStampSheetResultFromJson(data string) SetMaximumIdleMinutesByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumIdleMinutesByStampSheetResultFromDict(dict)
}

func NewSetMaximumIdleMinutesByStampSheetResultFromDict(data map[string]interface{}) SetMaximumIdleMinutesByStampSheetResult {
	return SetMaximumIdleMinutesByStampSheetResult{
		Item: NewStatusFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SetMaximumIdleMinutesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SetMaximumIdleMinutesByStampSheetResult) Pointer() *SetMaximumIdleMinutesByStampSheetResult {
	return &p
}

type ReceiveByStampSheetResult struct {
	Items                     []AcquireAction `json:"items"`
	Status                    *Status         `json:"status"`
	TransactionId             *string         `json:"transactionId"`
	StampSheet                *string         `json:"stampSheet"`
	StampSheetEncryptionKeyId *string         `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool           `json:"autoRunStampSheet"`
}

type ReceiveByStampSheetAsyncResult struct {
	result *ReceiveByStampSheetResult
	err    error
}

func NewReceiveByStampSheetResultFromJson(data string) ReceiveByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByStampSheetResultFromDict(dict)
}

func NewReceiveByStampSheetResultFromDict(data map[string]interface{}) ReceiveByStampSheetResult {
	return ReceiveByStampSheetResult{
		Items:                     CastAcquireActions(core.CastArray(data["items"])),
		Status:                    NewStatusFromDict(core.CastMap(data["status"])).Pointer(),
		TransactionId:             core.CastString(data["transactionId"]),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
		AutoRunStampSheet:         core.CastBool(data["autoRunStampSheet"]),
	}
}

func (p ReceiveByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
	}
}

func (p ReceiveByStampSheetResult) Pointer() *ReceiveByStampSheetResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentCategoryMaster `json:"item"`
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
		Item: NewCurrentCategoryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentCategoryMasterResult struct {
	Item *CurrentCategoryMaster `json:"item"`
}

type GetCurrentCategoryMasterAsyncResult struct {
	result *GetCurrentCategoryMasterResult
	err    error
}

func NewGetCurrentCategoryMasterResultFromJson(data string) GetCurrentCategoryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentCategoryMasterResultFromDict(dict)
}

func NewGetCurrentCategoryMasterResultFromDict(data map[string]interface{}) GetCurrentCategoryMasterResult {
	return GetCurrentCategoryMasterResult{
		Item: NewCurrentCategoryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentCategoryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentCategoryMasterResult) Pointer() *GetCurrentCategoryMasterResult {
	return &p
}

type UpdateCurrentCategoryMasterResult struct {
	Item *CurrentCategoryMaster `json:"item"`
}

type UpdateCurrentCategoryMasterAsyncResult struct {
	result *UpdateCurrentCategoryMasterResult
	err    error
}

func NewUpdateCurrentCategoryMasterResultFromJson(data string) UpdateCurrentCategoryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCategoryMasterResultFromDict(dict)
}

func NewUpdateCurrentCategoryMasterResultFromDict(data map[string]interface{}) UpdateCurrentCategoryMasterResult {
	return UpdateCurrentCategoryMasterResult{
		Item: NewCurrentCategoryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentCategoryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentCategoryMasterResult) Pointer() *UpdateCurrentCategoryMasterResult {
	return &p
}

type UpdateCurrentCategoryMasterFromGitHubResult struct {
	Item *CurrentCategoryMaster `json:"item"`
}

type UpdateCurrentCategoryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentCategoryMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentCategoryMasterFromGitHubResultFromJson(data string) UpdateCurrentCategoryMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentCategoryMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentCategoryMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentCategoryMasterFromGitHubResult {
	return UpdateCurrentCategoryMasterFromGitHubResult{
		Item: NewCurrentCategoryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentCategoryMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentCategoryMasterFromGitHubResult) Pointer() *UpdateCurrentCategoryMasterFromGitHubResult {
	return &p
}
