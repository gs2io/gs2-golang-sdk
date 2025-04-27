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

package dictionary

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

type DescribeEntryModelsResult struct {
	Items    []EntryModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeEntryModelsAsyncResult struct {
	result *DescribeEntryModelsResult
	err    error
}

func NewDescribeEntryModelsResultFromJson(data string) DescribeEntryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntryModelsResultFromDict(dict)
}

func NewDescribeEntryModelsResultFromDict(data map[string]interface{}) DescribeEntryModelsResult {
	return DescribeEntryModelsResult{
		Items: func() []EntryModel {
			if data["items"] == nil {
				return nil
			}
			return CastEntryModels(core.CastArray(data["items"]))
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

func (p DescribeEntryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntryModelsFromDict(
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

func (p DescribeEntryModelsResult) Pointer() *DescribeEntryModelsResult {
	return &p
}

type GetEntryModelResult struct {
	Item     *EntryModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEntryModelAsyncResult struct {
	result *GetEntryModelResult
	err    error
}

func NewGetEntryModelResultFromJson(data string) GetEntryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryModelResultFromDict(dict)
}

func NewGetEntryModelResultFromDict(data map[string]interface{}) GetEntryModelResult {
	return GetEntryModelResult{
		Item: func() *EntryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetEntryModelResult) ToDict() map[string]interface{} {
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

func (p GetEntryModelResult) Pointer() *GetEntryModelResult {
	return &p
}

type DescribeEntryModelMastersResult struct {
	Items         []EntryModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeEntryModelMastersAsyncResult struct {
	result *DescribeEntryModelMastersResult
	err    error
}

func NewDescribeEntryModelMastersResultFromJson(data string) DescribeEntryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntryModelMastersResultFromDict(dict)
}

func NewDescribeEntryModelMastersResultFromDict(data map[string]interface{}) DescribeEntryModelMastersResult {
	return DescribeEntryModelMastersResult{
		Items: func() []EntryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastEntryModelMasters(core.CastArray(data["items"]))
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

func (p DescribeEntryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntryModelMastersFromDict(
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

func (p DescribeEntryModelMastersResult) Pointer() *DescribeEntryModelMastersResult {
	return &p
}

type CreateEntryModelMasterResult struct {
	Item     *EntryModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateEntryModelMasterAsyncResult struct {
	result *CreateEntryModelMasterResult
	err    error
}

func NewCreateEntryModelMasterResultFromJson(data string) CreateEntryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateEntryModelMasterResultFromDict(dict)
}

func NewCreateEntryModelMasterResultFromDict(data map[string]interface{}) CreateEntryModelMasterResult {
	return CreateEntryModelMasterResult{
		Item: func() *EntryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateEntryModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateEntryModelMasterResult) Pointer() *CreateEntryModelMasterResult {
	return &p
}

type GetEntryModelMasterResult struct {
	Item     *EntryModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEntryModelMasterAsyncResult struct {
	result *GetEntryModelMasterResult
	err    error
}

func NewGetEntryModelMasterResultFromJson(data string) GetEntryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryModelMasterResultFromDict(dict)
}

func NewGetEntryModelMasterResultFromDict(data map[string]interface{}) GetEntryModelMasterResult {
	return GetEntryModelMasterResult{
		Item: func() *EntryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetEntryModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetEntryModelMasterResult) Pointer() *GetEntryModelMasterResult {
	return &p
}

type UpdateEntryModelMasterResult struct {
	Item     *EntryModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateEntryModelMasterAsyncResult struct {
	result *UpdateEntryModelMasterResult
	err    error
}

func NewUpdateEntryModelMasterResultFromJson(data string) UpdateEntryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateEntryModelMasterResultFromDict(dict)
}

func NewUpdateEntryModelMasterResultFromDict(data map[string]interface{}) UpdateEntryModelMasterResult {
	return UpdateEntryModelMasterResult{
		Item: func() *EntryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateEntryModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateEntryModelMasterResult) Pointer() *UpdateEntryModelMasterResult {
	return &p
}

type DeleteEntryModelMasterResult struct {
	Item     *EntryModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteEntryModelMasterAsyncResult struct {
	result *DeleteEntryModelMasterResult
	err    error
}

func NewDeleteEntryModelMasterResultFromJson(data string) DeleteEntryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEntryModelMasterResultFromDict(dict)
}

func NewDeleteEntryModelMasterResultFromDict(data map[string]interface{}) DeleteEntryModelMasterResult {
	return DeleteEntryModelMasterResult{
		Item: func() *EntryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteEntryModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteEntryModelMasterResult) Pointer() *DeleteEntryModelMasterResult {
	return &p
}

type DescribeEntriesResult struct {
	Items         []Entry              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeEntriesAsyncResult struct {
	result *DescribeEntriesResult
	err    error
}

func NewDescribeEntriesResultFromJson(data string) DescribeEntriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntriesResultFromDict(dict)
}

func NewDescribeEntriesResultFromDict(data map[string]interface{}) DescribeEntriesResult {
	return DescribeEntriesResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p DescribeEntriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p DescribeEntriesResult) Pointer() *DescribeEntriesResult {
	return &p
}

type DescribeEntriesByUserIdResult struct {
	Items         []Entry              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeEntriesByUserIdAsyncResult struct {
	result *DescribeEntriesByUserIdResult
	err    error
}

func NewDescribeEntriesByUserIdResultFromJson(data string) DescribeEntriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEntriesByUserIdResultFromDict(dict)
}

func NewDescribeEntriesByUserIdResultFromDict(data map[string]interface{}) DescribeEntriesByUserIdResult {
	return DescribeEntriesByUserIdResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p DescribeEntriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p DescribeEntriesByUserIdResult) Pointer() *DescribeEntriesByUserIdResult {
	return &p
}

type AddEntriesByUserIdResult struct {
	Items    []Entry              `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddEntriesByUserIdAsyncResult struct {
	result *AddEntriesByUserIdResult
	err    error
}

func NewAddEntriesByUserIdResultFromJson(data string) AddEntriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddEntriesByUserIdResultFromDict(dict)
}

func NewAddEntriesByUserIdResultFromDict(data map[string]interface{}) AddEntriesByUserIdResult {
	return AddEntriesByUserIdResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p AddEntriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p AddEntriesByUserIdResult) Pointer() *AddEntriesByUserIdResult {
	return &p
}

type GetEntryResult struct {
	Item     *Entry               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEntryAsyncResult struct {
	result *GetEntryResult
	err    error
}

func NewGetEntryResultFromJson(data string) GetEntryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryResultFromDict(dict)
}

func NewGetEntryResultFromDict(data map[string]interface{}) GetEntryResult {
	return GetEntryResult{
		Item: func() *Entry {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetEntryResult) ToDict() map[string]interface{} {
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

func (p GetEntryResult) Pointer() *GetEntryResult {
	return &p
}

type GetEntryByUserIdResult struct {
	Item     *Entry               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEntryByUserIdAsyncResult struct {
	result *GetEntryByUserIdResult
	err    error
}

func NewGetEntryByUserIdResultFromJson(data string) GetEntryByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryByUserIdResultFromDict(dict)
}

func NewGetEntryByUserIdResultFromDict(data map[string]interface{}) GetEntryByUserIdResult {
	return GetEntryByUserIdResult{
		Item: func() *Entry {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetEntryByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetEntryByUserIdResult) Pointer() *GetEntryByUserIdResult {
	return &p
}

type GetEntryWithSignatureResult struct {
	Item      *Entry               `json:"item"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetEntryWithSignatureAsyncResult struct {
	result *GetEntryWithSignatureResult
	err    error
}

func NewGetEntryWithSignatureResultFromJson(data string) GetEntryWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryWithSignatureResultFromDict(dict)
}

func NewGetEntryWithSignatureResultFromDict(data map[string]interface{}) GetEntryWithSignatureResult {
	return GetEntryWithSignatureResult{
		Item: func() *Entry {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Body: func() *string {
			v, ok := data["body"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["body"])
		}(),
		Signature: func() *string {
			v, ok := data["signature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signature"])
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

func (p GetEntryWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"body":      p.Body,
		"signature": p.Signature,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetEntryWithSignatureResult) Pointer() *GetEntryWithSignatureResult {
	return &p
}

type GetEntryWithSignatureByUserIdResult struct {
	Item      *Entry               `json:"item"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetEntryWithSignatureByUserIdAsyncResult struct {
	result *GetEntryWithSignatureByUserIdResult
	err    error
}

func NewGetEntryWithSignatureByUserIdResultFromJson(data string) GetEntryWithSignatureByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEntryWithSignatureByUserIdResultFromDict(dict)
}

func NewGetEntryWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetEntryWithSignatureByUserIdResult {
	return GetEntryWithSignatureByUserIdResult{
		Item: func() *Entry {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEntryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Body: func() *string {
			v, ok := data["body"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["body"])
		}(),
		Signature: func() *string {
			v, ok := data["signature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signature"])
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

func (p GetEntryWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"body":      p.Body,
		"signature": p.Signature,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetEntryWithSignatureByUserIdResult) Pointer() *GetEntryWithSignatureByUserIdResult {
	return &p
}

type ResetByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetByUserIdAsyncResult struct {
	result *ResetByUserIdResult
	err    error
}

func NewResetByUserIdResultFromJson(data string) ResetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetByUserIdResultFromDict(dict)
}

func NewResetByUserIdResultFromDict(data map[string]interface{}) ResetByUserIdResult {
	return ResetByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ResetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ResetByUserIdResult) Pointer() *ResetByUserIdResult {
	return &p
}

type VerifyEntryResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyEntryAsyncResult struct {
	result *VerifyEntryResult
	err    error
}

func NewVerifyEntryResultFromJson(data string) VerifyEntryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEntryResultFromDict(dict)
}

func NewVerifyEntryResultFromDict(data map[string]interface{}) VerifyEntryResult {
	return VerifyEntryResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyEntryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyEntryResult) Pointer() *VerifyEntryResult {
	return &p
}

type VerifyEntryByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyEntryByUserIdAsyncResult struct {
	result *VerifyEntryByUserIdResult
	err    error
}

func NewVerifyEntryByUserIdResultFromJson(data string) VerifyEntryByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEntryByUserIdResultFromDict(dict)
}

func NewVerifyEntryByUserIdResultFromDict(data map[string]interface{}) VerifyEntryByUserIdResult {
	return VerifyEntryByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyEntryByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyEntryByUserIdResult) Pointer() *VerifyEntryByUserIdResult {
	return &p
}

type DeleteEntriesResult struct {
	Items    []Entry              `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteEntriesAsyncResult struct {
	result *DeleteEntriesResult
	err    error
}

func NewDeleteEntriesResultFromJson(data string) DeleteEntriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEntriesResultFromDict(dict)
}

func NewDeleteEntriesResultFromDict(data map[string]interface{}) DeleteEntriesResult {
	return DeleteEntriesResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p DeleteEntriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p DeleteEntriesResult) Pointer() *DeleteEntriesResult {
	return &p
}

type DeleteEntriesByUserIdResult struct {
	Items    []Entry              `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteEntriesByUserIdAsyncResult struct {
	result *DeleteEntriesByUserIdResult
	err    error
}

func NewDeleteEntriesByUserIdResultFromJson(data string) DeleteEntriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEntriesByUserIdResultFromDict(dict)
}

func NewDeleteEntriesByUserIdResultFromDict(data map[string]interface{}) DeleteEntriesByUserIdResult {
	return DeleteEntriesByUserIdResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p DeleteEntriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p DeleteEntriesByUserIdResult) Pointer() *DeleteEntriesByUserIdResult {
	return &p
}

type AddEntriesByStampSheetResult struct {
	Items    []Entry              `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddEntriesByStampSheetAsyncResult struct {
	result *AddEntriesByStampSheetResult
	err    error
}

func NewAddEntriesByStampSheetResultFromJson(data string) AddEntriesByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddEntriesByStampSheetResultFromDict(dict)
}

func NewAddEntriesByStampSheetResultFromDict(data map[string]interface{}) AddEntriesByStampSheetResult {
	return AddEntriesByStampSheetResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p AddEntriesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
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

func (p AddEntriesByStampSheetResult) Pointer() *AddEntriesByStampSheetResult {
	return &p
}

type DeleteEntriesByStampTaskResult struct {
	Items           []Entry              `json:"items"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type DeleteEntriesByStampTaskAsyncResult struct {
	result *DeleteEntriesByStampTaskResult
	err    error
}

func NewDeleteEntriesByStampTaskResultFromJson(data string) DeleteEntriesByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteEntriesByStampTaskResultFromDict(dict)
}

func NewDeleteEntriesByStampTaskResultFromDict(data map[string]interface{}) DeleteEntriesByStampTaskResult {
	return DeleteEntriesByStampTaskResult{
		Items: func() []Entry {
			if data["items"] == nil {
				return nil
			}
			return CastEntries(core.CastArray(data["items"]))
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

func (p DeleteEntriesByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
			p.Items,
		),
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteEntriesByStampTaskResult) Pointer() *DeleteEntriesByStampTaskResult {
	return &p
}

type VerifyEntryByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyEntryByStampTaskAsyncResult struct {
	result *VerifyEntryByStampTaskResult
	err    error
}

func NewVerifyEntryByStampTaskResultFromJson(data string) VerifyEntryByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyEntryByStampTaskResultFromDict(dict)
}

func NewVerifyEntryByStampTaskResultFromDict(data map[string]interface{}) VerifyEntryByStampTaskResult {
	return VerifyEntryByStampTaskResult{
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

func (p VerifyEntryByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyEntryByStampTaskResult) Pointer() *VerifyEntryByStampTaskResult {
	return &p
}

type DescribeLikesResult struct {
	Items         []Like               `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeLikesAsyncResult struct {
	result *DescribeLikesResult
	err    error
}

func NewDescribeLikesResultFromJson(data string) DescribeLikesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLikesResultFromDict(dict)
}

func NewDescribeLikesResultFromDict(data map[string]interface{}) DescribeLikesResult {
	return DescribeLikesResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p DescribeLikesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p DescribeLikesResult) Pointer() *DescribeLikesResult {
	return &p
}

type DescribeLikesByUserIdResult struct {
	Items         []Like               `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeLikesByUserIdAsyncResult struct {
	result *DescribeLikesByUserIdResult
	err    error
}

func NewDescribeLikesByUserIdResultFromJson(data string) DescribeLikesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLikesByUserIdResultFromDict(dict)
}

func NewDescribeLikesByUserIdResultFromDict(data map[string]interface{}) DescribeLikesByUserIdResult {
	return DescribeLikesByUserIdResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p DescribeLikesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p DescribeLikesByUserIdResult) Pointer() *DescribeLikesByUserIdResult {
	return &p
}

type AddLikesResult struct {
	Items    []Like               `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddLikesAsyncResult struct {
	result *AddLikesResult
	err    error
}

func NewAddLikesResultFromJson(data string) AddLikesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddLikesResultFromDict(dict)
}

func NewAddLikesResultFromDict(data map[string]interface{}) AddLikesResult {
	return AddLikesResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p AddLikesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p AddLikesResult) Pointer() *AddLikesResult {
	return &p
}

type AddLikesByUserIdResult struct {
	Items    []Like               `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddLikesByUserIdAsyncResult struct {
	result *AddLikesByUserIdResult
	err    error
}

func NewAddLikesByUserIdResultFromJson(data string) AddLikesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddLikesByUserIdResultFromDict(dict)
}

func NewAddLikesByUserIdResultFromDict(data map[string]interface{}) AddLikesByUserIdResult {
	return AddLikesByUserIdResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p AddLikesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p AddLikesByUserIdResult) Pointer() *AddLikesByUserIdResult {
	return &p
}

type GetLikeResult struct {
	Item     *Like                `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLikeAsyncResult struct {
	result *GetLikeResult
	err    error
}

func NewGetLikeResultFromJson(data string) GetLikeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLikeResultFromDict(dict)
}

func NewGetLikeResultFromDict(data map[string]interface{}) GetLikeResult {
	return GetLikeResult{
		Item: func() *Like {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLikeFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetLikeResult) ToDict() map[string]interface{} {
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

func (p GetLikeResult) Pointer() *GetLikeResult {
	return &p
}

type GetLikeByUserIdResult struct {
	Item     *Like                `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLikeByUserIdAsyncResult struct {
	result *GetLikeByUserIdResult
	err    error
}

func NewGetLikeByUserIdResultFromJson(data string) GetLikeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLikeByUserIdResultFromDict(dict)
}

func NewGetLikeByUserIdResultFromDict(data map[string]interface{}) GetLikeByUserIdResult {
	return GetLikeByUserIdResult{
		Item: func() *Like {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLikeFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetLikeByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetLikeByUserIdResult) Pointer() *GetLikeByUserIdResult {
	return &p
}

type ResetLikesResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetLikesAsyncResult struct {
	result *ResetLikesResult
	err    error
}

func NewResetLikesResultFromJson(data string) ResetLikesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetLikesResultFromDict(dict)
}

func NewResetLikesResultFromDict(data map[string]interface{}) ResetLikesResult {
	return ResetLikesResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ResetLikesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ResetLikesResult) Pointer() *ResetLikesResult {
	return &p
}

type ResetLikesByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetLikesByUserIdAsyncResult struct {
	result *ResetLikesByUserIdResult
	err    error
}

func NewResetLikesByUserIdResultFromJson(data string) ResetLikesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetLikesByUserIdResultFromDict(dict)
}

func NewResetLikesByUserIdResultFromDict(data map[string]interface{}) ResetLikesByUserIdResult {
	return ResetLikesByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ResetLikesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ResetLikesByUserIdResult) Pointer() *ResetLikesByUserIdResult {
	return &p
}

type DeleteLikesResult struct {
	Items    []Like               `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteLikesAsyncResult struct {
	result *DeleteLikesResult
	err    error
}

func NewDeleteLikesResultFromJson(data string) DeleteLikesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLikesResultFromDict(dict)
}

func NewDeleteLikesResultFromDict(data map[string]interface{}) DeleteLikesResult {
	return DeleteLikesResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p DeleteLikesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p DeleteLikesResult) Pointer() *DeleteLikesResult {
	return &p
}

type DeleteLikesByUserIdResult struct {
	Items    []Like               `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteLikesByUserIdAsyncResult struct {
	result *DeleteLikesByUserIdResult
	err    error
}

func NewDeleteLikesByUserIdResultFromJson(data string) DeleteLikesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLikesByUserIdResultFromDict(dict)
}

func NewDeleteLikesByUserIdResultFromDict(data map[string]interface{}) DeleteLikesByUserIdResult {
	return DeleteLikesByUserIdResult{
		Items: func() []Like {
			if data["items"] == nil {
				return nil
			}
			return CastLikes(core.CastArray(data["items"]))
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

func (p DeleteLikesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLikesFromDict(
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

func (p DeleteLikesByUserIdResult) Pointer() *DeleteLikesByUserIdResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentEntryMaster  `json:"item"`
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
		Item: func() *CurrentEntryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentEntryMasterResult struct {
	Item     *CurrentEntryMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentEntryMasterAsyncResult struct {
	result *GetCurrentEntryMasterResult
	err    error
}

func NewGetCurrentEntryMasterResultFromJson(data string) GetCurrentEntryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentEntryMasterResultFromDict(dict)
}

func NewGetCurrentEntryMasterResultFromDict(data map[string]interface{}) GetCurrentEntryMasterResult {
	return GetCurrentEntryMasterResult{
		Item: func() *CurrentEntryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentEntryMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentEntryMasterResult) Pointer() *GetCurrentEntryMasterResult {
	return &p
}

type PreUpdateCurrentEntryMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentEntryMasterAsyncResult struct {
	result *PreUpdateCurrentEntryMasterResult
	err    error
}

func NewPreUpdateCurrentEntryMasterResultFromJson(data string) PreUpdateCurrentEntryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentEntryMasterResultFromDict(dict)
}

func NewPreUpdateCurrentEntryMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentEntryMasterResult {
	return PreUpdateCurrentEntryMasterResult{
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

func (p PreUpdateCurrentEntryMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentEntryMasterResult) Pointer() *PreUpdateCurrentEntryMasterResult {
	return &p
}

type UpdateCurrentEntryMasterResult struct {
	Item     *CurrentEntryMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentEntryMasterAsyncResult struct {
	result *UpdateCurrentEntryMasterResult
	err    error
}

func NewUpdateCurrentEntryMasterResultFromJson(data string) UpdateCurrentEntryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEntryMasterResultFromDict(dict)
}

func NewUpdateCurrentEntryMasterResultFromDict(data map[string]interface{}) UpdateCurrentEntryMasterResult {
	return UpdateCurrentEntryMasterResult{
		Item: func() *CurrentEntryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentEntryMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentEntryMasterResult) Pointer() *UpdateCurrentEntryMasterResult {
	return &p
}

type UpdateCurrentEntryMasterFromGitHubResult struct {
	Item     *CurrentEntryMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentEntryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEntryMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentEntryMasterFromGitHubResultFromJson(data string) UpdateCurrentEntryMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentEntryMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentEntryMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentEntryMasterFromGitHubResult {
	return UpdateCurrentEntryMasterFromGitHubResult{
		Item: func() *CurrentEntryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentEntryMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentEntryMasterFromGitHubResult) Pointer() *UpdateCurrentEntryMasterFromGitHubResult {
	return &p
}
