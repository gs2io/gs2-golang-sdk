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

package loginReward

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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
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
	return DumpUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
	return CleanUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
	return CheckCleanUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
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
	return ImportUserDataByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
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
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeBonusModelMastersResult struct {
	Items         []BonusModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBonusModelMastersAsyncResult struct {
	result *DescribeBonusModelMastersResult
	err    error
}

func NewDescribeBonusModelMastersResultFromJson(data string) DescribeBonusModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBonusModelMastersResultFromDict(dict)
}

func NewDescribeBonusModelMastersResultFromDict(data map[string]interface{}) DescribeBonusModelMastersResult {
	return DescribeBonusModelMastersResult{
		Items: func() []BonusModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastBonusModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeBonusModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBonusModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeBonusModelMastersResult) Pointer() *DescribeBonusModelMastersResult {
	return &p
}

type CreateBonusModelMasterResult struct {
	Item     *BonusModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateBonusModelMasterAsyncResult struct {
	result *CreateBonusModelMasterResult
	err    error
}

func NewCreateBonusModelMasterResultFromJson(data string) CreateBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBonusModelMasterResultFromDict(dict)
}

func NewCreateBonusModelMasterResultFromDict(data map[string]interface{}) CreateBonusModelMasterResult {
	return CreateBonusModelMasterResult{
		Item: func() *BonusModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateBonusModelMasterResult) Pointer() *CreateBonusModelMasterResult {
	return &p
}

type GetBonusModelMasterResult struct {
	Item     *BonusModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBonusModelMasterAsyncResult struct {
	result *GetBonusModelMasterResult
	err    error
}

func NewGetBonusModelMasterResultFromJson(data string) GetBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBonusModelMasterResultFromDict(dict)
}

func NewGetBonusModelMasterResultFromDict(data map[string]interface{}) GetBonusModelMasterResult {
	return GetBonusModelMasterResult{
		Item: func() *BonusModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBonusModelMasterResult) Pointer() *GetBonusModelMasterResult {
	return &p
}

type UpdateBonusModelMasterResult struct {
	Item     *BonusModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateBonusModelMasterAsyncResult struct {
	result *UpdateBonusModelMasterResult
	err    error
}

func NewUpdateBonusModelMasterResultFromJson(data string) UpdateBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBonusModelMasterResultFromDict(dict)
}

func NewUpdateBonusModelMasterResultFromDict(data map[string]interface{}) UpdateBonusModelMasterResult {
	return UpdateBonusModelMasterResult{
		Item: func() *BonusModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateBonusModelMasterResult) Pointer() *UpdateBonusModelMasterResult {
	return &p
}

type DeleteBonusModelMasterResult struct {
	Item     *BonusModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteBonusModelMasterAsyncResult struct {
	result *DeleteBonusModelMasterResult
	err    error
}

func NewDeleteBonusModelMasterResultFromJson(data string) DeleteBonusModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBonusModelMasterResultFromDict(dict)
}

func NewDeleteBonusModelMasterResultFromDict(data map[string]interface{}) DeleteBonusModelMasterResult {
	return DeleteBonusModelMasterResult{
		Item: func() *BonusModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteBonusModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteBonusModelMasterResult) Pointer() *DeleteBonusModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentBonusMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *CurrentBonusMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentBonusMasterResult struct {
	Item     *CurrentBonusMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentBonusMasterAsyncResult struct {
	result *GetCurrentBonusMasterResult
	err    error
}

func NewGetCurrentBonusMasterResultFromJson(data string) GetCurrentBonusMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentBonusMasterResultFromDict(dict)
}

func NewGetCurrentBonusMasterResultFromDict(data map[string]interface{}) GetCurrentBonusMasterResult {
	return GetCurrentBonusMasterResult{
		Item: func() *CurrentBonusMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentBonusMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetCurrentBonusMasterResult) Pointer() *GetCurrentBonusMasterResult {
	return &p
}

type UpdateCurrentBonusMasterResult struct {
	Item     *CurrentBonusMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentBonusMasterAsyncResult struct {
	result *UpdateCurrentBonusMasterResult
	err    error
}

func NewUpdateCurrentBonusMasterResultFromJson(data string) UpdateCurrentBonusMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentBonusMasterResultFromDict(dict)
}

func NewUpdateCurrentBonusMasterResultFromDict(data map[string]interface{}) UpdateCurrentBonusMasterResult {
	return UpdateCurrentBonusMasterResult{
		Item: func() *CurrentBonusMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentBonusMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentBonusMasterResult) Pointer() *UpdateCurrentBonusMasterResult {
	return &p
}

type UpdateCurrentBonusMasterFromGitHubResult struct {
	Item     *CurrentBonusMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentBonusMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentBonusMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentBonusMasterFromGitHubResultFromJson(data string) UpdateCurrentBonusMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentBonusMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentBonusMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentBonusMasterFromGitHubResult {
	return UpdateCurrentBonusMasterFromGitHubResult{
		Item: func() *CurrentBonusMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentBonusMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentBonusMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentBonusMasterFromGitHubResult) Pointer() *UpdateCurrentBonusMasterFromGitHubResult {
	return &p
}

type DescribeBonusModelsResult struct {
	Items    []BonusModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeBonusModelsAsyncResult struct {
	result *DescribeBonusModelsResult
	err    error
}

func NewDescribeBonusModelsResultFromJson(data string) DescribeBonusModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBonusModelsResultFromDict(dict)
}

func NewDescribeBonusModelsResultFromDict(data map[string]interface{}) DescribeBonusModelsResult {
	return DescribeBonusModelsResult{
		Items: func() []BonusModel {
			if data["items"] == nil {
				return nil
			}
			return CastBonusModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeBonusModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBonusModelsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeBonusModelsResult) Pointer() *DescribeBonusModelsResult {
	return &p
}

type GetBonusModelResult struct {
	Item     *BonusModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBonusModelAsyncResult struct {
	result *GetBonusModelResult
	err    error
}

func NewGetBonusModelResultFromJson(data string) GetBonusModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBonusModelResultFromDict(dict)
}

func NewGetBonusModelResultFromDict(data map[string]interface{}) GetBonusModelResult {
	return GetBonusModelResult{
		Item: func() *BonusModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBonusModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBonusModelResult) Pointer() *GetBonusModelResult {
	return &p
}

type ReceiveResult struct {
	Item                      *ReceiveStatus          `json:"item"`
	BonusModel                *BonusModel             `json:"bonusModel"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReceiveResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ReceiveResult) Pointer() *ReceiveResult {
	return &p
}

type ReceiveByUserIdResult struct {
	Item                      *ReceiveStatus          `json:"item"`
	BonusModel                *BonusModel             `json:"bonusModel"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ReceiveByUserIdResult) Pointer() *ReceiveByUserIdResult {
	return &p
}

type MissedReceiveResult struct {
	Item                      *ReceiveStatus          `json:"item"`
	BonusModel                *BonusModel             `json:"bonusModel"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type MissedReceiveAsyncResult struct {
	result *MissedReceiveResult
	err    error
}

func NewMissedReceiveResultFromJson(data string) MissedReceiveResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissedReceiveResultFromDict(dict)
}

func NewMissedReceiveResultFromDict(data map[string]interface{}) MissedReceiveResult {
	return MissedReceiveResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p MissedReceiveResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p MissedReceiveResult) Pointer() *MissedReceiveResult {
	return &p
}

type MissedReceiveByUserIdResult struct {
	Item                      *ReceiveStatus          `json:"item"`
	BonusModel                *BonusModel             `json:"bonusModel"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type MissedReceiveByUserIdAsyncResult struct {
	result *MissedReceiveByUserIdResult
	err    error
}

func NewMissedReceiveByUserIdResultFromJson(data string) MissedReceiveByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMissedReceiveByUserIdResultFromDict(dict)
}

func NewMissedReceiveByUserIdResultFromDict(data map[string]interface{}) MissedReceiveByUserIdResult {
	return MissedReceiveByUserIdResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p MissedReceiveByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p MissedReceiveByUserIdResult) Pointer() *MissedReceiveByUserIdResult {
	return &p
}

type DescribeReceiveStatusesResult struct {
	Items         []ReceiveStatus      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeReceiveStatusesAsyncResult struct {
	result *DescribeReceiveStatusesResult
	err    error
}

func NewDescribeReceiveStatusesResultFromJson(data string) DescribeReceiveStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveStatusesResultFromDict(dict)
}

func NewDescribeReceiveStatusesResultFromDict(data map[string]interface{}) DescribeReceiveStatusesResult {
	return DescribeReceiveStatusesResult{
		Items: func() []ReceiveStatus {
			if data["items"] == nil {
				return nil
			}
			return CastReceiveStatuses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeReceiveStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeReceiveStatusesResult) Pointer() *DescribeReceiveStatusesResult {
	return &p
}

type DescribeReceiveStatusesByUserIdResult struct {
	Items         []ReceiveStatus      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeReceiveStatusesByUserIdAsyncResult struct {
	result *DescribeReceiveStatusesByUserIdResult
	err    error
}

func NewDescribeReceiveStatusesByUserIdResultFromJson(data string) DescribeReceiveStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveStatusesByUserIdResultFromDict(dict)
}

func NewDescribeReceiveStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeReceiveStatusesByUserIdResult {
	return DescribeReceiveStatusesByUserIdResult{
		Items: func() []ReceiveStatus {
			if data["items"] == nil {
				return nil
			}
			return CastReceiveStatuses(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeReceiveStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveStatusesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeReceiveStatusesByUserIdResult) Pointer() *DescribeReceiveStatusesByUserIdResult {
	return &p
}

type GetReceiveStatusResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type GetReceiveStatusAsyncResult struct {
	result *GetReceiveStatusResult
	err    error
}

func NewGetReceiveStatusResultFromJson(data string) GetReceiveStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveStatusResultFromDict(dict)
}

func NewGetReceiveStatusResultFromDict(data map[string]interface{}) GetReceiveStatusResult {
	return GetReceiveStatusResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetReceiveStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetReceiveStatusResult) Pointer() *GetReceiveStatusResult {
	return &p
}

type GetReceiveStatusByUserIdResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type GetReceiveStatusByUserIdAsyncResult struct {
	result *GetReceiveStatusByUserIdResult
	err    error
}

func NewGetReceiveStatusByUserIdResultFromJson(data string) GetReceiveStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveStatusByUserIdResultFromDict(dict)
}

func NewGetReceiveStatusByUserIdResultFromDict(data map[string]interface{}) GetReceiveStatusByUserIdResult {
	return GetReceiveStatusByUserIdResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetReceiveStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetReceiveStatusByUserIdResult) Pointer() *GetReceiveStatusByUserIdResult {
	return &p
}

type DeleteReceiveStatusByUserIdResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type DeleteReceiveStatusByUserIdAsyncResult struct {
	result *DeleteReceiveStatusByUserIdResult
	err    error
}

func NewDeleteReceiveStatusByUserIdResultFromJson(data string) DeleteReceiveStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceiveStatusByUserIdResultFromDict(dict)
}

func NewDeleteReceiveStatusByUserIdResultFromDict(data map[string]interface{}) DeleteReceiveStatusByUserIdResult {
	return DeleteReceiveStatusByUserIdResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteReceiveStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteReceiveStatusByUserIdResult) Pointer() *DeleteReceiveStatusByUserIdResult {
	return &p
}

type DeleteReceiveStatusByStampSheetResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type DeleteReceiveStatusByStampSheetAsyncResult struct {
	result *DeleteReceiveStatusByStampSheetResult
	err    error
}

func NewDeleteReceiveStatusByStampSheetResultFromJson(data string) DeleteReceiveStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteReceiveStatusByStampSheetResultFromDict(dict)
}

func NewDeleteReceiveStatusByStampSheetResultFromDict(data map[string]interface{}) DeleteReceiveStatusByStampSheetResult {
	return DeleteReceiveStatusByStampSheetResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteReceiveStatusByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteReceiveStatusByStampSheetResult) Pointer() *DeleteReceiveStatusByStampSheetResult {
	return &p
}

type MarkReceivedResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type MarkReceivedAsyncResult struct {
	result *MarkReceivedResult
	err    error
}

func NewMarkReceivedResultFromJson(data string) MarkReceivedResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedResultFromDict(dict)
}

func NewMarkReceivedResultFromDict(data map[string]interface{}) MarkReceivedResult {
	return MarkReceivedResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p MarkReceivedResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p MarkReceivedResult) Pointer() *MarkReceivedResult {
	return &p
}

type MarkReceivedByUserIdResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type MarkReceivedByUserIdAsyncResult struct {
	result *MarkReceivedByUserIdResult
	err    error
}

func NewMarkReceivedByUserIdResultFromJson(data string) MarkReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedByUserIdResultFromDict(dict)
}

func NewMarkReceivedByUserIdResultFromDict(data map[string]interface{}) MarkReceivedByUserIdResult {
	return MarkReceivedByUserIdResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p MarkReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p MarkReceivedByUserIdResult) Pointer() *MarkReceivedByUserIdResult {
	return &p
}

type UnmarkReceivedByUserIdResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type UnmarkReceivedByUserIdAsyncResult struct {
	result *UnmarkReceivedByUserIdResult
	err    error
}

func NewUnmarkReceivedByUserIdResultFromJson(data string) UnmarkReceivedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnmarkReceivedByUserIdResultFromDict(dict)
}

func NewUnmarkReceivedByUserIdResultFromDict(data map[string]interface{}) UnmarkReceivedByUserIdResult {
	return UnmarkReceivedByUserIdResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnmarkReceivedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnmarkReceivedByUserIdResult) Pointer() *UnmarkReceivedByUserIdResult {
	return &p
}

type MarkReceivedByStampTaskResult struct {
	Item            *ReceiveStatus       `json:"item"`
	BonusModel      *BonusModel          `json:"bonusModel"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type MarkReceivedByStampTaskAsyncResult struct {
	result *MarkReceivedByStampTaskResult
	err    error
}

func NewMarkReceivedByStampTaskResultFromJson(data string) MarkReceivedByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMarkReceivedByStampTaskResultFromDict(dict)
}

func NewMarkReceivedByStampTaskResultFromDict(data map[string]interface{}) MarkReceivedByStampTaskResult {
	return MarkReceivedByStampTaskResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p MarkReceivedByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p MarkReceivedByStampTaskResult) Pointer() *MarkReceivedByStampTaskResult {
	return &p
}

type UnmarkReceivedByStampSheetResult struct {
	Item       *ReceiveStatus       `json:"item"`
	BonusModel *BonusModel          `json:"bonusModel"`
	Metadata   *core.ResultMetadata `json:"metadata"`
}

type UnmarkReceivedByStampSheetAsyncResult struct {
	result *UnmarkReceivedByStampSheetResult
	err    error
}

func NewUnmarkReceivedByStampSheetResultFromJson(data string) UnmarkReceivedByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnmarkReceivedByStampSheetResultFromDict(dict)
}

func NewUnmarkReceivedByStampSheetResultFromDict(data map[string]interface{}) UnmarkReceivedByStampSheetResult {
	return UnmarkReceivedByStampSheetResult{
		Item: func() *ReceiveStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BonusModel: func() *BonusModel {
			v, ok := data["bonusModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBonusModelFromDict(core.CastMap(data["bonusModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnmarkReceivedByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"bonusModel": func() map[string]interface{} {
			if p.BonusModel == nil {
				return nil
			}
			return p.BonusModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnmarkReceivedByStampSheetResult) Pointer() *UnmarkReceivedByStampSheetResult {
	return &p
}
