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

package ranking

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
	Items         []Namespace          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Namespace {
			if data["items"] == nil {
				return nil
			}
			return CastNamespaces(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Status   *string              `json:"status"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *Namespace           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Namespace {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewNamespaceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
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
		UploadToken: func() *string {
			v, ok := data["uploadToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadToken"])
		}(),
		UploadUrl: func() *string {
			v, ok := data["uploadUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["uploadUrl"])
		}(),
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
	Metadata *core.ResultMetadata `json:"metadata"`
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
	Url      *string              `json:"url"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Url: func() *string {
			v, ok := data["url"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["url"])
		}(),
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

type DescribeCategoryModelsResult struct {
	Items    []CategoryModel      `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Items: func() []CategoryModel {
			if data["items"] == nil {
				return nil
			}
			return CastCategoryModels(core.CastArray(data["items"]))
		}(),
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
	Item     *CategoryModel       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CategoryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCategoryModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type DescribeCategoryModelMastersResult struct {
	Items         []CategoryModelMaster `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
	Metadata      *core.ResultMetadata  `json:"metadata"`
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
		Items: func() []CategoryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastCategoryModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
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
	Item     *CategoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CategoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *CategoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CategoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *CategoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CategoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *CategoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CategoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type SubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SubscribeAsyncResult struct {
	result *SubscribeResult
	err    error
}

func NewSubscribeResultFromJson(data string) SubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeResultFromDict(dict)
}

func NewSubscribeResultFromDict(data map[string]interface{}) SubscribeResult {
	return SubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SubscribeResult) Pointer() *SubscribeResult {
	return &p
}

type SubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SubscribeByUserIdAsyncResult struct {
	result *SubscribeByUserIdResult
	err    error
}

func NewSubscribeByUserIdResultFromJson(data string) SubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeByUserIdResultFromDict(dict)
}

func NewSubscribeByUserIdResultFromDict(data map[string]interface{}) SubscribeByUserIdResult {
	return SubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p SubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p SubscribeByUserIdResult) Pointer() *SubscribeByUserIdResult {
	return &p
}

type DescribeScoresResult struct {
	Items         []Score              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeScoresAsyncResult struct {
	result *DescribeScoresResult
	err    error
}

func NewDescribeScoresResultFromJson(data string) DescribeScoresResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScoresResultFromDict(dict)
}

func NewDescribeScoresResultFromDict(data map[string]interface{}) DescribeScoresResult {
	return DescribeScoresResult{
		Items: func() []Score {
			if data["items"] == nil {
				return nil
			}
			return CastScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeScoresResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeScoresResult) Pointer() *DescribeScoresResult {
	return &p
}

type DescribeScoresByUserIdResult struct {
	Items         []Score              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeScoresByUserIdAsyncResult struct {
	result *DescribeScoresByUserIdResult
	err    error
}

func NewDescribeScoresByUserIdResultFromJson(data string) DescribeScoresByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeScoresByUserIdResultFromDict(dict)
}

func NewDescribeScoresByUserIdResultFromDict(data map[string]interface{}) DescribeScoresByUserIdResult {
	return DescribeScoresByUserIdResult{
		Items: func() []Score {
			if data["items"] == nil {
				return nil
			}
			return CastScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeScoresByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeScoresByUserIdResult) Pointer() *DescribeScoresByUserIdResult {
	return &p
}

type GetScoreResult struct {
	Item     *Score               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetScoreAsyncResult struct {
	result *GetScoreResult
	err    error
}

func NewGetScoreResultFromJson(data string) GetScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScoreResultFromDict(dict)
}

func NewGetScoreResultFromDict(data map[string]interface{}) GetScoreResult {
	return GetScoreResult{
		Item: func() *Score {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetScoreResult) Pointer() *GetScoreResult {
	return &p
}

type GetScoreByUserIdResult struct {
	Item     *Score               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetScoreByUserIdAsyncResult struct {
	result *GetScoreByUserIdResult
	err    error
}

func NewGetScoreByUserIdResultFromJson(data string) GetScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetScoreByUserIdResultFromDict(dict)
}

func NewGetScoreByUserIdResultFromDict(data map[string]interface{}) GetScoreByUserIdResult {
	return GetScoreByUserIdResult{
		Item: func() *Score {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetScoreByUserIdResult) Pointer() *GetScoreByUserIdResult {
	return &p
}

type DescribeRankingsResult struct {
	Items         []Ranking            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRankingsAsyncResult struct {
	result *DescribeRankingsResult
	err    error
}

func NewDescribeRankingsResultFromJson(data string) DescribeRankingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingsResultFromDict(dict)
}

func NewDescribeRankingsResultFromDict(data map[string]interface{}) DescribeRankingsResult {
	return DescribeRankingsResult{
		Items: func() []Ranking {
			if data["items"] == nil {
				return nil
			}
			return CastRankings(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRankingsResult) Pointer() *DescribeRankingsResult {
	return &p
}

type DescribeRankingssByUserIdResult struct {
	Items         []Ranking            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRankingssByUserIdAsyncResult struct {
	result *DescribeRankingssByUserIdResult
	err    error
}

func NewDescribeRankingssByUserIdResultFromJson(data string) DescribeRankingssByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRankingssByUserIdResultFromDict(dict)
}

func NewDescribeRankingssByUserIdResultFromDict(data map[string]interface{}) DescribeRankingssByUserIdResult {
	return DescribeRankingssByUserIdResult{
		Items: func() []Ranking {
			if data["items"] == nil {
				return nil
			}
			return CastRankings(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeRankingssByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRankingssByUserIdResult) Pointer() *DescribeRankingssByUserIdResult {
	return &p
}

type DescribeNearRankingsResult struct {
	Items    []Ranking            `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeNearRankingsAsyncResult struct {
	result *DescribeNearRankingsResult
	err    error
}

func NewDescribeNearRankingsResultFromJson(data string) DescribeNearRankingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNearRankingsResultFromDict(dict)
}

func NewDescribeNearRankingsResultFromDict(data map[string]interface{}) DescribeNearRankingsResult {
	return DescribeNearRankingsResult{
		Items: func() []Ranking {
			if data["items"] == nil {
				return nil
			}
			return CastRankings(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeNearRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
	}
}

func (p DescribeNearRankingsResult) Pointer() *DescribeNearRankingsResult {
	return &p
}

type GetRankingResult struct {
	Item     *Ranking             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRankingAsyncResult struct {
	result *GetRankingResult
	err    error
}

func NewGetRankingResultFromJson(data string) GetRankingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingResultFromDict(dict)
}

func NewGetRankingResultFromDict(data map[string]interface{}) GetRankingResult {
	return GetRankingResult{
		Item: func() *Ranking {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRankingFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetRankingResult) Pointer() *GetRankingResult {
	return &p
}

type GetRankingByUserIdResult struct {
	Item     *Ranking             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRankingByUserIdAsyncResult struct {
	result *GetRankingByUserIdResult
	err    error
}

func NewGetRankingByUserIdResultFromJson(data string) GetRankingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRankingByUserIdResultFromDict(dict)
}

func NewGetRankingByUserIdResultFromDict(data map[string]interface{}) GetRankingByUserIdResult {
	return GetRankingByUserIdResult{
		Item: func() *Ranking {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRankingFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetRankingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetRankingByUserIdResult) Pointer() *GetRankingByUserIdResult {
	return &p
}

type PutScoreResult struct {
	Item     *Score               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutScoreAsyncResult struct {
	result *PutScoreResult
	err    error
}

func NewPutScoreResultFromJson(data string) PutScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutScoreResultFromDict(dict)
}

func NewPutScoreResultFromDict(data map[string]interface{}) PutScoreResult {
	return PutScoreResult{
		Item: func() *Score {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutScoreResult) Pointer() *PutScoreResult {
	return &p
}

type PutScoreByUserIdResult struct {
	Item     *Score               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutScoreByUserIdAsyncResult struct {
	result *PutScoreByUserIdResult
	err    error
}

func NewPutScoreByUserIdResultFromJson(data string) PutScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutScoreByUserIdResultFromDict(dict)
}

func NewPutScoreByUserIdResultFromDict(data map[string]interface{}) PutScoreByUserIdResult {
	return PutScoreByUserIdResult{
		Item: func() *Score {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutScoreByUserIdResult) Pointer() *PutScoreByUserIdResult {
	return &p
}

type CalcRankingResult struct {
	Processing *bool                `json:"processing"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type CalcRankingAsyncResult struct {
	result *CalcRankingResult
	err    error
}

func NewCalcRankingResultFromJson(data string) CalcRankingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCalcRankingResultFromDict(dict)
}

func NewCalcRankingResultFromDict(data map[string]interface{}) CalcRankingResult {
	return CalcRankingResult{
		Processing: func() *bool {
			v, ok := data["processing"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["processing"])
		}(),
	}
}

func (p CalcRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"processing": p.Processing,
	}
}

func (p CalcRankingResult) Pointer() *CalcRankingResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
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
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentRankingMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetCurrentRankingMasterAsyncResult struct {
	result *GetCurrentRankingMasterResult
	err    error
}

func NewGetCurrentRankingMasterResultFromJson(data string) GetCurrentRankingMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRankingMasterResultFromDict(dict)
}

func NewGetCurrentRankingMasterResultFromDict(data map[string]interface{}) GetCurrentRankingMasterResult {
	return GetCurrentRankingMasterResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentRankingMasterResult) Pointer() *GetCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentRankingMasterAsyncResult struct {
	result *UpdateCurrentRankingMasterResult
	err    error
}

func NewUpdateCurrentRankingMasterResultFromJson(data string) UpdateCurrentRankingMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterResultFromDict(dict)
}

func NewUpdateCurrentRankingMasterResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterResult {
	return UpdateCurrentRankingMasterResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentRankingMasterResult) Pointer() *UpdateCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterFromGitHubResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentRankingMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRankingMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRankingMasterFromGitHubResultFromJson(data string) UpdateCurrentRankingMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentRankingMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubResult {
	return UpdateCurrentRankingMasterFromGitHubResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) Pointer() *UpdateCurrentRankingMasterFromGitHubResult {
	return &p
}

type GetSubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

func NewGetSubscribeResultFromJson(data string) GetSubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeResultFromDict(dict)
}

func NewGetSubscribeResultFromDict(data map[string]interface{}) GetSubscribeResult {
	return GetSubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeResult) Pointer() *GetSubscribeResult {
	return &p
}

type GetSubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

func NewGetSubscribeByUserIdResultFromJson(data string) GetSubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeByUserIdResultFromDict(dict)
}

func NewGetSubscribeByUserIdResultFromDict(data map[string]interface{}) GetSubscribeByUserIdResult {
	return GetSubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeByUserIdResult) Pointer() *GetSubscribeByUserIdResult {
	return &p
}

type UnsubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnsubscribeAsyncResult struct {
	result *UnsubscribeResult
	err    error
}

func NewUnsubscribeResultFromJson(data string) UnsubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeResultFromDict(dict)
}

func NewUnsubscribeResultFromDict(data map[string]interface{}) UnsubscribeResult {
	return UnsubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UnsubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UnsubscribeResult) Pointer() *UnsubscribeResult {
	return &p
}

type UnsubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnsubscribeByUserIdAsyncResult struct {
	result *UnsubscribeByUserIdResult
	err    error
}

func NewUnsubscribeByUserIdResultFromJson(data string) UnsubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeByUserIdResultFromDict(dict)
}

func NewUnsubscribeByUserIdResultFromDict(data map[string]interface{}) UnsubscribeByUserIdResult {
	return UnsubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UnsubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UnsubscribeByUserIdResult) Pointer() *UnsubscribeByUserIdResult {
	return &p
}

type DescribeSubscribesByCategoryNameResult struct {
	Items    []SubscribeUser      `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscribesByCategoryNameAsyncResult struct {
	result *DescribeSubscribesByCategoryNameResult
	err    error
}

func NewDescribeSubscribesByCategoryNameResultFromJson(data string) DescribeSubscribesByCategoryNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByCategoryNameResultFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameResultFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameResult {
	return DescribeSubscribesByCategoryNameResult{
		Items: func() []SubscribeUser {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeUsers(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeSubscribesByCategoryNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
	}
}

func (p DescribeSubscribesByCategoryNameResult) Pointer() *DescribeSubscribesByCategoryNameResult {
	return &p
}

type DescribeSubscribesByCategoryNameAndUserIdResult struct {
	Items    []SubscribeUser      `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscribesByCategoryNameAndUserIdAsyncResult struct {
	result *DescribeSubscribesByCategoryNameAndUserIdResult
	err    error
}

func NewDescribeSubscribesByCategoryNameAndUserIdResultFromJson(data string) DescribeSubscribesByCategoryNameAndUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByCategoryNameAndUserIdResultFromDict(dict)
}

func NewDescribeSubscribesByCategoryNameAndUserIdResultFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameAndUserIdResult {
	return DescribeSubscribesByCategoryNameAndUserIdResult{
		Items: func() []SubscribeUser {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeUsers(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdResult) Pointer() *DescribeSubscribesByCategoryNameAndUserIdResult {
	return &p
}
