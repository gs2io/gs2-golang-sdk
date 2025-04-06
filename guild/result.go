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

package guild

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

type DescribeGuildModelMastersResult struct {
	Items         []GuildModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGuildModelMastersAsyncResult struct {
	result *DescribeGuildModelMastersResult
	err    error
}

func NewDescribeGuildModelMastersResultFromJson(data string) DescribeGuildModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGuildModelMastersResultFromDict(dict)
}

func NewDescribeGuildModelMastersResultFromDict(data map[string]interface{}) DescribeGuildModelMastersResult {
	return DescribeGuildModelMastersResult{
		Items: func() []GuildModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastGuildModelMasters(core.CastArray(data["items"]))
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

func (p DescribeGuildModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGuildModelMastersFromDict(
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

func (p DescribeGuildModelMastersResult) Pointer() *DescribeGuildModelMastersResult {
	return &p
}

type CreateGuildModelMasterResult struct {
	Item     *GuildModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGuildModelMasterAsyncResult struct {
	result *CreateGuildModelMasterResult
	err    error
}

func NewCreateGuildModelMasterResultFromJson(data string) CreateGuildModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildModelMasterResultFromDict(dict)
}

func NewCreateGuildModelMasterResultFromDict(data map[string]interface{}) CreateGuildModelMasterResult {
	return CreateGuildModelMasterResult{
		Item: func() *GuildModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateGuildModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateGuildModelMasterResult) Pointer() *CreateGuildModelMasterResult {
	return &p
}

type GetGuildModelMasterResult struct {
	Item     *GuildModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGuildModelMasterAsyncResult struct {
	result *GetGuildModelMasterResult
	err    error
}

func NewGetGuildModelMasterResultFromJson(data string) GetGuildModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildModelMasterResultFromDict(dict)
}

func NewGetGuildModelMasterResultFromDict(data map[string]interface{}) GetGuildModelMasterResult {
	return GetGuildModelMasterResult{
		Item: func() *GuildModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetGuildModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetGuildModelMasterResult) Pointer() *GetGuildModelMasterResult {
	return &p
}

type UpdateGuildModelMasterResult struct {
	Item     *GuildModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGuildModelMasterAsyncResult struct {
	result *UpdateGuildModelMasterResult
	err    error
}

func NewUpdateGuildModelMasterResultFromJson(data string) UpdateGuildModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildModelMasterResultFromDict(dict)
}

func NewUpdateGuildModelMasterResultFromDict(data map[string]interface{}) UpdateGuildModelMasterResult {
	return UpdateGuildModelMasterResult{
		Item: func() *GuildModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateGuildModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateGuildModelMasterResult) Pointer() *UpdateGuildModelMasterResult {
	return &p
}

type DeleteGuildModelMasterResult struct {
	Item     *GuildModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGuildModelMasterAsyncResult struct {
	result *DeleteGuildModelMasterResult
	err    error
}

func NewDeleteGuildModelMasterResultFromJson(data string) DeleteGuildModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildModelMasterResultFromDict(dict)
}

func NewDeleteGuildModelMasterResultFromDict(data map[string]interface{}) DeleteGuildModelMasterResult {
	return DeleteGuildModelMasterResult{
		Item: func() *GuildModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteGuildModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteGuildModelMasterResult) Pointer() *DeleteGuildModelMasterResult {
	return &p
}

type DescribeGuildModelsResult struct {
	Items    []GuildModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeGuildModelsAsyncResult struct {
	result *DescribeGuildModelsResult
	err    error
}

func NewDescribeGuildModelsResultFromJson(data string) DescribeGuildModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGuildModelsResultFromDict(dict)
}

func NewDescribeGuildModelsResultFromDict(data map[string]interface{}) DescribeGuildModelsResult {
	return DescribeGuildModelsResult{
		Items: func() []GuildModel {
			if data["items"] == nil {
				return nil
			}
			return CastGuildModels(core.CastArray(data["items"]))
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

func (p DescribeGuildModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGuildModelsFromDict(
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

func (p DescribeGuildModelsResult) Pointer() *DescribeGuildModelsResult {
	return &p
}

type GetGuildModelResult struct {
	Item     *GuildModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGuildModelAsyncResult struct {
	result *GetGuildModelResult
	err    error
}

func NewGetGuildModelResultFromJson(data string) GetGuildModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildModelResultFromDict(dict)
}

func NewGetGuildModelResultFromDict(data map[string]interface{}) GetGuildModelResult {
	return GetGuildModelResult{
		Item: func() *GuildModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetGuildModelResult) ToDict() map[string]interface{} {
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

func (p GetGuildModelResult) Pointer() *GetGuildModelResult {
	return &p
}

type SearchGuildsResult struct {
	Items         []Guild              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type SearchGuildsAsyncResult struct {
	result *SearchGuildsResult
	err    error
}

func NewSearchGuildsResultFromJson(data string) SearchGuildsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSearchGuildsResultFromDict(dict)
}

func NewSearchGuildsResultFromDict(data map[string]interface{}) SearchGuildsResult {
	return SearchGuildsResult{
		Items: func() []Guild {
			if data["items"] == nil {
				return nil
			}
			return CastGuilds(core.CastArray(data["items"]))
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

func (p SearchGuildsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGuildsFromDict(
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

func (p SearchGuildsResult) Pointer() *SearchGuildsResult {
	return &p
}

type SearchGuildsByUserIdResult struct {
	Items         []Guild              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type SearchGuildsByUserIdAsyncResult struct {
	result *SearchGuildsByUserIdResult
	err    error
}

func NewSearchGuildsByUserIdResultFromJson(data string) SearchGuildsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSearchGuildsByUserIdResultFromDict(dict)
}

func NewSearchGuildsByUserIdResultFromDict(data map[string]interface{}) SearchGuildsByUserIdResult {
	return SearchGuildsByUserIdResult{
		Items: func() []Guild {
			if data["items"] == nil {
				return nil
			}
			return CastGuilds(core.CastArray(data["items"]))
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

func (p SearchGuildsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGuildsFromDict(
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

func (p SearchGuildsByUserIdResult) Pointer() *SearchGuildsByUserIdResult {
	return &p
}

type CreateGuildResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGuildAsyncResult struct {
	result *CreateGuildResult
	err    error
}

func NewCreateGuildResultFromJson(data string) CreateGuildResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildResultFromDict(dict)
}

func NewCreateGuildResultFromDict(data map[string]interface{}) CreateGuildResult {
	return CreateGuildResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateGuildResult) ToDict() map[string]interface{} {
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

func (p CreateGuildResult) Pointer() *CreateGuildResult {
	return &p
}

type CreateGuildByUserIdResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGuildByUserIdAsyncResult struct {
	result *CreateGuildByUserIdResult
	err    error
}

func NewCreateGuildByUserIdResultFromJson(data string) CreateGuildByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildByUserIdResultFromDict(dict)
}

func NewCreateGuildByUserIdResultFromDict(data map[string]interface{}) CreateGuildByUserIdResult {
	return CreateGuildByUserIdResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateGuildByUserIdResult) ToDict() map[string]interface{} {
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

func (p CreateGuildByUserIdResult) Pointer() *CreateGuildByUserIdResult {
	return &p
}

type GetGuildResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGuildAsyncResult struct {
	result *GetGuildResult
	err    error
}

func NewGetGuildResultFromJson(data string) GetGuildResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildResultFromDict(dict)
}

func NewGetGuildResultFromDict(data map[string]interface{}) GetGuildResult {
	return GetGuildResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetGuildResult) ToDict() map[string]interface{} {
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

func (p GetGuildResult) Pointer() *GetGuildResult {
	return &p
}

type GetGuildByUserIdResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGuildByUserIdAsyncResult struct {
	result *GetGuildByUserIdResult
	err    error
}

func NewGetGuildByUserIdResultFromJson(data string) GetGuildByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildByUserIdResultFromDict(dict)
}

func NewGetGuildByUserIdResultFromDict(data map[string]interface{}) GetGuildByUserIdResult {
	return GetGuildByUserIdResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetGuildByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetGuildByUserIdResult) Pointer() *GetGuildByUserIdResult {
	return &p
}

type UpdateGuildResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGuildAsyncResult struct {
	result *UpdateGuildResult
	err    error
}

func NewUpdateGuildResultFromJson(data string) UpdateGuildResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildResultFromDict(dict)
}

func NewUpdateGuildResultFromDict(data map[string]interface{}) UpdateGuildResult {
	return UpdateGuildResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateGuildResult) ToDict() map[string]interface{} {
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

func (p UpdateGuildResult) Pointer() *UpdateGuildResult {
	return &p
}

type UpdateGuildByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGuildByGuildNameAsyncResult struct {
	result *UpdateGuildByGuildNameResult
	err    error
}

func NewUpdateGuildByGuildNameResultFromJson(data string) UpdateGuildByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildByGuildNameResultFromDict(dict)
}

func NewUpdateGuildByGuildNameResultFromDict(data map[string]interface{}) UpdateGuildByGuildNameResult {
	return UpdateGuildByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateGuildByGuildNameResult) ToDict() map[string]interface{} {
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

func (p UpdateGuildByGuildNameResult) Pointer() *UpdateGuildByGuildNameResult {
	return &p
}

type DeleteMemberResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteMemberAsyncResult struct {
	result *DeleteMemberResult
	err    error
}

func NewDeleteMemberResultFromJson(data string) DeleteMemberResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMemberResultFromDict(dict)
}

func NewDeleteMemberResultFromDict(data map[string]interface{}) DeleteMemberResult {
	return DeleteMemberResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteMemberResult) ToDict() map[string]interface{} {
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

func (p DeleteMemberResult) Pointer() *DeleteMemberResult {
	return &p
}

type DeleteMemberByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteMemberByGuildNameAsyncResult struct {
	result *DeleteMemberByGuildNameResult
	err    error
}

func NewDeleteMemberByGuildNameResultFromJson(data string) DeleteMemberByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMemberByGuildNameResultFromDict(dict)
}

func NewDeleteMemberByGuildNameResultFromDict(data map[string]interface{}) DeleteMemberByGuildNameResult {
	return DeleteMemberByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteMemberByGuildNameResult) ToDict() map[string]interface{} {
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

func (p DeleteMemberByGuildNameResult) Pointer() *DeleteMemberByGuildNameResult {
	return &p
}

type UpdateMemberRoleResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateMemberRoleAsyncResult struct {
	result *UpdateMemberRoleResult
	err    error
}

func NewUpdateMemberRoleResultFromJson(data string) UpdateMemberRoleResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberRoleResultFromDict(dict)
}

func NewUpdateMemberRoleResultFromDict(data map[string]interface{}) UpdateMemberRoleResult {
	return UpdateMemberRoleResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateMemberRoleResult) ToDict() map[string]interface{} {
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

func (p UpdateMemberRoleResult) Pointer() *UpdateMemberRoleResult {
	return &p
}

type UpdateMemberRoleByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateMemberRoleByGuildNameAsyncResult struct {
	result *UpdateMemberRoleByGuildNameResult
	err    error
}

func NewUpdateMemberRoleByGuildNameResultFromJson(data string) UpdateMemberRoleByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberRoleByGuildNameResultFromDict(dict)
}

func NewUpdateMemberRoleByGuildNameResultFromDict(data map[string]interface{}) UpdateMemberRoleByGuildNameResult {
	return UpdateMemberRoleByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateMemberRoleByGuildNameResult) ToDict() map[string]interface{} {
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

func (p UpdateMemberRoleByGuildNameResult) Pointer() *UpdateMemberRoleByGuildNameResult {
	return &p
}

type BatchUpdateMemberRoleResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type BatchUpdateMemberRoleAsyncResult struct {
	result *BatchUpdateMemberRoleResult
	err    error
}

func NewBatchUpdateMemberRoleResultFromJson(data string) BatchUpdateMemberRoleResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBatchUpdateMemberRoleResultFromDict(dict)
}

func NewBatchUpdateMemberRoleResultFromDict(data map[string]interface{}) BatchUpdateMemberRoleResult {
	return BatchUpdateMemberRoleResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p BatchUpdateMemberRoleResult) ToDict() map[string]interface{} {
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

func (p BatchUpdateMemberRoleResult) Pointer() *BatchUpdateMemberRoleResult {
	return &p
}

type BatchUpdateMemberRoleByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type BatchUpdateMemberRoleByGuildNameAsyncResult struct {
	result *BatchUpdateMemberRoleByGuildNameResult
	err    error
}

func NewBatchUpdateMemberRoleByGuildNameResultFromJson(data string) BatchUpdateMemberRoleByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBatchUpdateMemberRoleByGuildNameResultFromDict(dict)
}

func NewBatchUpdateMemberRoleByGuildNameResultFromDict(data map[string]interface{}) BatchUpdateMemberRoleByGuildNameResult {
	return BatchUpdateMemberRoleByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p BatchUpdateMemberRoleByGuildNameResult) ToDict() map[string]interface{} {
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

func (p BatchUpdateMemberRoleByGuildNameResult) Pointer() *BatchUpdateMemberRoleByGuildNameResult {
	return &p
}

type DeleteGuildResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGuildAsyncResult struct {
	result *DeleteGuildResult
	err    error
}

func NewDeleteGuildResultFromJson(data string) DeleteGuildResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildResultFromDict(dict)
}

func NewDeleteGuildResultFromDict(data map[string]interface{}) DeleteGuildResult {
	return DeleteGuildResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteGuildResult) ToDict() map[string]interface{} {
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

func (p DeleteGuildResult) Pointer() *DeleteGuildResult {
	return &p
}

type DeleteGuildByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGuildByGuildNameAsyncResult struct {
	result *DeleteGuildByGuildNameResult
	err    error
}

func NewDeleteGuildByGuildNameResultFromJson(data string) DeleteGuildByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildByGuildNameResultFromDict(dict)
}

func NewDeleteGuildByGuildNameResultFromDict(data map[string]interface{}) DeleteGuildByGuildNameResult {
	return DeleteGuildByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteGuildByGuildNameResult) ToDict() map[string]interface{} {
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

func (p DeleteGuildByGuildNameResult) Pointer() *DeleteGuildByGuildNameResult {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult struct {
	result *IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	err    error
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromJson(data string) IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(dict)
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	return IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult) ToDict() map[string]interface{} {
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

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaximumCurrentMaximumMemberCountAsyncResult struct {
	result *DecreaseMaximumCurrentMaximumMemberCountResult
	err    error
}

func NewDecreaseMaximumCurrentMaximumMemberCountResultFromJson(data string) DecreaseMaximumCurrentMaximumMemberCountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumCurrentMaximumMemberCountResultFromDict(dict)
}

func NewDecreaseMaximumCurrentMaximumMemberCountResultFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountResult {
	return DecreaseMaximumCurrentMaximumMemberCountResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DecreaseMaximumCurrentMaximumMemberCountResult) ToDict() map[string]interface{} {
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

func (p DecreaseMaximumCurrentMaximumMemberCountResult) Pointer() *DecreaseMaximumCurrentMaximumMemberCountResult {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult struct {
	result *DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	err    error
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromJson(data string) DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(dict)
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	return DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult) ToDict() map[string]interface{} {
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

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult {
	return &p
}

type VerifyCurrentMaximumMemberCountResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyCurrentMaximumMemberCountAsyncResult struct {
	result *VerifyCurrentMaximumMemberCountResult
	err    error
}

func NewVerifyCurrentMaximumMemberCountResultFromJson(data string) VerifyCurrentMaximumMemberCountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCurrentMaximumMemberCountResultFromDict(dict)
}

func NewVerifyCurrentMaximumMemberCountResultFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountResult {
	return VerifyCurrentMaximumMemberCountResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyCurrentMaximumMemberCountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyCurrentMaximumMemberCountResult) Pointer() *VerifyCurrentMaximumMemberCountResult {
	return &p
}

type VerifyCurrentMaximumMemberCountByGuildNameResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyCurrentMaximumMemberCountByGuildNameAsyncResult struct {
	result *VerifyCurrentMaximumMemberCountByGuildNameResult
	err    error
}

func NewVerifyCurrentMaximumMemberCountByGuildNameResultFromJson(data string) VerifyCurrentMaximumMemberCountByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCurrentMaximumMemberCountByGuildNameResultFromDict(dict)
}

func NewVerifyCurrentMaximumMemberCountByGuildNameResultFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountByGuildNameResult {
	return VerifyCurrentMaximumMemberCountByGuildNameResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyCurrentMaximumMemberCountByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyCurrentMaximumMemberCountByGuildNameResult) Pointer() *VerifyCurrentMaximumMemberCountByGuildNameResult {
	return &p
}

type VerifyIncludeMemberResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeMemberAsyncResult struct {
	result *VerifyIncludeMemberResult
	err    error
}

func NewVerifyIncludeMemberResultFromJson(data string) VerifyIncludeMemberResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeMemberResultFromDict(dict)
}

func NewVerifyIncludeMemberResultFromDict(data map[string]interface{}) VerifyIncludeMemberResult {
	return VerifyIncludeMemberResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyIncludeMemberResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyIncludeMemberResult) Pointer() *VerifyIncludeMemberResult {
	return &p
}

type VerifyIncludeMemberByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeMemberByUserIdAsyncResult struct {
	result *VerifyIncludeMemberByUserIdResult
	err    error
}

func NewVerifyIncludeMemberByUserIdResultFromJson(data string) VerifyIncludeMemberByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeMemberByUserIdResultFromDict(dict)
}

func NewVerifyIncludeMemberByUserIdResultFromDict(data map[string]interface{}) VerifyIncludeMemberByUserIdResult {
	return VerifyIncludeMemberByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyIncludeMemberByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyIncludeMemberByUserIdResult) Pointer() *VerifyIncludeMemberByUserIdResult {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByGuildNameResult struct {
	Item     *Guild               `json:"item"`
	Old      *Guild               `json:"old"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult struct {
	result *SetMaximumCurrentMaximumMemberCountByGuildNameResult
	err    error
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameResultFromJson(data string) SetMaximumCurrentMaximumMemberCountByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(dict)
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameResultFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByGuildNameResult {
	return SetMaximumCurrentMaximumMemberCountByGuildNameResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Guild {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["old"])).Pointer()
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

func (p SetMaximumCurrentMaximumMemberCountByGuildNameResult) ToDict() map[string]interface{} {
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetMaximumCurrentMaximumMemberCountByGuildNameResult) Pointer() *SetMaximumCurrentMaximumMemberCountByGuildNameResult {
	return &p
}

type AssumeResult struct {
	Token    *string              `json:"token"`
	UserId   *string              `json:"userId"`
	Expire   *int64               `json:"expire"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AssumeAsyncResult struct {
	result *AssumeResult
	err    error
}

func NewAssumeResultFromJson(data string) AssumeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAssumeResultFromDict(dict)
}

func NewAssumeResultFromDict(data map[string]interface{}) AssumeResult {
	return AssumeResult{
		Token: func() *string {
			v, ok := data["token"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["token"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Expire: func() *int64 {
			v, ok := data["expire"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expire"])
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

func (p AssumeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"token":  p.Token,
		"userId": p.UserId,
		"expire": p.Expire,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AssumeResult) Pointer() *AssumeResult {
	return &p
}

type AssumeByUserIdResult struct {
	Token    *string              `json:"token"`
	UserId   *string              `json:"userId"`
	Expire   *int64               `json:"expire"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AssumeByUserIdAsyncResult struct {
	result *AssumeByUserIdResult
	err    error
}

func NewAssumeByUserIdResultFromJson(data string) AssumeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAssumeByUserIdResultFromDict(dict)
}

func NewAssumeByUserIdResultFromDict(data map[string]interface{}) AssumeByUserIdResult {
	return AssumeByUserIdResult{
		Token: func() *string {
			v, ok := data["token"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["token"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Expire: func() *int64 {
			v, ok := data["expire"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["expire"])
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

func (p AssumeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"token":  p.Token,
		"userId": p.UserId,
		"expire": p.Expire,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AssumeByUserIdResult) Pointer() *AssumeByUserIdResult {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult struct {
	result *IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult
	err    error
}

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetResultFromJson(data string) IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetResultFromDict(dict)
}

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetResultFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult {
	return IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult) ToDict() map[string]interface{} {
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

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult struct {
	Item            *Guild               `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult struct {
	result *DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult
	err    error
}

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskResultFromJson(data string) DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskResultFromDict(dict)
}

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskResultFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult {
	return DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
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

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByStampSheetResult struct {
	Item     *Guild               `json:"item"`
	Old      *Guild               `json:"old"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult struct {
	result *SetMaximumCurrentMaximumMemberCountByStampSheetResult
	err    error
}

func NewSetMaximumCurrentMaximumMemberCountByStampSheetResultFromJson(data string) SetMaximumCurrentMaximumMemberCountByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumCurrentMaximumMemberCountByStampSheetResultFromDict(dict)
}

func NewSetMaximumCurrentMaximumMemberCountByStampSheetResultFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByStampSheetResult {
	return SetMaximumCurrentMaximumMemberCountByStampSheetResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Guild {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["old"])).Pointer()
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

func (p SetMaximumCurrentMaximumMemberCountByStampSheetResult) ToDict() map[string]interface{} {
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
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetMaximumCurrentMaximumMemberCountByStampSheetResult) Pointer() *SetMaximumCurrentMaximumMemberCountByStampSheetResult {
	return &p
}

type VerifyCurrentMaximumMemberCountByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyCurrentMaximumMemberCountByStampTaskAsyncResult struct {
	result *VerifyCurrentMaximumMemberCountByStampTaskResult
	err    error
}

func NewVerifyCurrentMaximumMemberCountByStampTaskResultFromJson(data string) VerifyCurrentMaximumMemberCountByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyCurrentMaximumMemberCountByStampTaskResultFromDict(dict)
}

func NewVerifyCurrentMaximumMemberCountByStampTaskResultFromDict(data map[string]interface{}) VerifyCurrentMaximumMemberCountByStampTaskResult {
	return VerifyCurrentMaximumMemberCountByStampTaskResult{
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

func (p VerifyCurrentMaximumMemberCountByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyCurrentMaximumMemberCountByStampTaskResult) Pointer() *VerifyCurrentMaximumMemberCountByStampTaskResult {
	return &p
}

type VerifyIncludeMemberByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeMemberByStampTaskAsyncResult struct {
	result *VerifyIncludeMemberByStampTaskResult
	err    error
}

func NewVerifyIncludeMemberByStampTaskResultFromJson(data string) VerifyIncludeMemberByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeMemberByStampTaskResultFromDict(dict)
}

func NewVerifyIncludeMemberByStampTaskResultFromDict(data map[string]interface{}) VerifyIncludeMemberByStampTaskResult {
	return VerifyIncludeMemberByStampTaskResult{
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

func (p VerifyIncludeMemberByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyIncludeMemberByStampTaskResult) Pointer() *VerifyIncludeMemberByStampTaskResult {
	return &p
}

type DescribeJoinedGuildsResult struct {
	Items         []JoinedGuild        `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeJoinedGuildsAsyncResult struct {
	result *DescribeJoinedGuildsResult
	err    error
}

func NewDescribeJoinedGuildsResultFromJson(data string) DescribeJoinedGuildsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedGuildsResultFromDict(dict)
}

func NewDescribeJoinedGuildsResultFromDict(data map[string]interface{}) DescribeJoinedGuildsResult {
	return DescribeJoinedGuildsResult{
		Items: func() []JoinedGuild {
			if data["items"] == nil {
				return nil
			}
			return CastJoinedGuilds(core.CastArray(data["items"]))
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

func (p DescribeJoinedGuildsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJoinedGuildsFromDict(
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

func (p DescribeJoinedGuildsResult) Pointer() *DescribeJoinedGuildsResult {
	return &p
}

type DescribeJoinedGuildsByUserIdResult struct {
	Items         []JoinedGuild        `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeJoinedGuildsByUserIdAsyncResult struct {
	result *DescribeJoinedGuildsByUserIdResult
	err    error
}

func NewDescribeJoinedGuildsByUserIdResultFromJson(data string) DescribeJoinedGuildsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedGuildsByUserIdResultFromDict(dict)
}

func NewDescribeJoinedGuildsByUserIdResultFromDict(data map[string]interface{}) DescribeJoinedGuildsByUserIdResult {
	return DescribeJoinedGuildsByUserIdResult{
		Items: func() []JoinedGuild {
			if data["items"] == nil {
				return nil
			}
			return CastJoinedGuilds(core.CastArray(data["items"]))
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

func (p DescribeJoinedGuildsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJoinedGuildsFromDict(
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

func (p DescribeJoinedGuildsByUserIdResult) Pointer() *DescribeJoinedGuildsByUserIdResult {
	return &p
}

type GetJoinedGuildResult struct {
	Item     *JoinedGuild         `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetJoinedGuildAsyncResult struct {
	result *GetJoinedGuildResult
	err    error
}

func NewGetJoinedGuildResultFromJson(data string) GetJoinedGuildResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedGuildResultFromDict(dict)
}

func NewGetJoinedGuildResultFromDict(data map[string]interface{}) GetJoinedGuildResult {
	return GetJoinedGuildResult{
		Item: func() *JoinedGuild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetJoinedGuildResult) ToDict() map[string]interface{} {
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

func (p GetJoinedGuildResult) Pointer() *GetJoinedGuildResult {
	return &p
}

type GetJoinedGuildByUserIdResult struct {
	Item     *JoinedGuild         `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetJoinedGuildByUserIdAsyncResult struct {
	result *GetJoinedGuildByUserIdResult
	err    error
}

func NewGetJoinedGuildByUserIdResultFromJson(data string) GetJoinedGuildByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedGuildByUserIdResultFromDict(dict)
}

func NewGetJoinedGuildByUserIdResultFromDict(data map[string]interface{}) GetJoinedGuildByUserIdResult {
	return GetJoinedGuildByUserIdResult{
		Item: func() *JoinedGuild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetJoinedGuildByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetJoinedGuildByUserIdResult) Pointer() *GetJoinedGuildByUserIdResult {
	return &p
}

type UpdateMemberMetadataResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateMemberMetadataAsyncResult struct {
	result *UpdateMemberMetadataResult
	err    error
}

func NewUpdateMemberMetadataResultFromJson(data string) UpdateMemberMetadataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberMetadataResultFromDict(dict)
}

func NewUpdateMemberMetadataResultFromDict(data map[string]interface{}) UpdateMemberMetadataResult {
	return UpdateMemberMetadataResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateMemberMetadataResult) ToDict() map[string]interface{} {
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

func (p UpdateMemberMetadataResult) Pointer() *UpdateMemberMetadataResult {
	return &p
}

type UpdateMemberMetadataByUserIdResult struct {
	Item     *Guild               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateMemberMetadataByUserIdAsyncResult struct {
	result *UpdateMemberMetadataByUserIdResult
	err    error
}

func NewUpdateMemberMetadataByUserIdResultFromJson(data string) UpdateMemberMetadataByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberMetadataByUserIdResultFromDict(dict)
}

func NewUpdateMemberMetadataByUserIdResultFromDict(data map[string]interface{}) UpdateMemberMetadataByUserIdResult {
	return UpdateMemberMetadataByUserIdResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateMemberMetadataByUserIdResult) ToDict() map[string]interface{} {
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

func (p UpdateMemberMetadataByUserIdResult) Pointer() *UpdateMemberMetadataByUserIdResult {
	return &p
}

type WithdrawalResult struct {
	Item     *JoinedGuild         `json:"item"`
	Guild    *Guild               `json:"guild"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WithdrawalAsyncResult struct {
	result *WithdrawalResult
	err    error
}

func NewWithdrawalResultFromJson(data string) WithdrawalResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawalResultFromDict(dict)
}

func NewWithdrawalResultFromDict(data map[string]interface{}) WithdrawalResult {
	return WithdrawalResult{
		Item: func() *JoinedGuild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p WithdrawalResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WithdrawalResult) Pointer() *WithdrawalResult {
	return &p
}

type WithdrawalByUserIdResult struct {
	Item     *JoinedGuild         `json:"item"`
	Guild    *Guild               `json:"guild"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type WithdrawalByUserIdAsyncResult struct {
	result *WithdrawalByUserIdResult
	err    error
}

func NewWithdrawalByUserIdResultFromJson(data string) WithdrawalByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawalByUserIdResultFromDict(dict)
}

func NewWithdrawalByUserIdResultFromDict(data map[string]interface{}) WithdrawalByUserIdResult {
	return WithdrawalByUserIdResult{
		Item: func() *JoinedGuild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p WithdrawalByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WithdrawalByUserIdResult) Pointer() *WithdrawalByUserIdResult {
	return &p
}

type GetLastGuildMasterActivityResult struct {
	Item     *LastGuildMasterActivity `json:"item"`
	Guild    *Guild                   `json:"guild"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetLastGuildMasterActivityAsyncResult struct {
	result *GetLastGuildMasterActivityResult
	err    error
}

func NewGetLastGuildMasterActivityResultFromJson(data string) GetLastGuildMasterActivityResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLastGuildMasterActivityResultFromDict(dict)
}

func NewGetLastGuildMasterActivityResultFromDict(data map[string]interface{}) GetLastGuildMasterActivityResult {
	return GetLastGuildMasterActivityResult{
		Item: func() *LastGuildMasterActivity {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLastGuildMasterActivityFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p GetLastGuildMasterActivityResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetLastGuildMasterActivityResult) Pointer() *GetLastGuildMasterActivityResult {
	return &p
}

type GetLastGuildMasterActivityByGuildNameResult struct {
	Item     *LastGuildMasterActivity `json:"item"`
	Guild    *Guild                   `json:"guild"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetLastGuildMasterActivityByGuildNameAsyncResult struct {
	result *GetLastGuildMasterActivityByGuildNameResult
	err    error
}

func NewGetLastGuildMasterActivityByGuildNameResultFromJson(data string) GetLastGuildMasterActivityByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLastGuildMasterActivityByGuildNameResultFromDict(dict)
}

func NewGetLastGuildMasterActivityByGuildNameResultFromDict(data map[string]interface{}) GetLastGuildMasterActivityByGuildNameResult {
	return GetLastGuildMasterActivityByGuildNameResult{
		Item: func() *LastGuildMasterActivity {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLastGuildMasterActivityFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p GetLastGuildMasterActivityByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetLastGuildMasterActivityByGuildNameResult) Pointer() *GetLastGuildMasterActivityByGuildNameResult {
	return &p
}

type PromoteSeniorMemberResult struct {
	Item     *LastGuildMasterActivity `json:"item"`
	Guild    *Guild                   `json:"guild"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type PromoteSeniorMemberAsyncResult struct {
	result *PromoteSeniorMemberResult
	err    error
}

func NewPromoteSeniorMemberResultFromJson(data string) PromoteSeniorMemberResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPromoteSeniorMemberResultFromDict(dict)
}

func NewPromoteSeniorMemberResultFromDict(data map[string]interface{}) PromoteSeniorMemberResult {
	return PromoteSeniorMemberResult{
		Item: func() *LastGuildMasterActivity {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLastGuildMasterActivityFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p PromoteSeniorMemberResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PromoteSeniorMemberResult) Pointer() *PromoteSeniorMemberResult {
	return &p
}

type PromoteSeniorMemberByGuildNameResult struct {
	Item     *LastGuildMasterActivity `json:"item"`
	Guild    *Guild                   `json:"guild"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type PromoteSeniorMemberByGuildNameAsyncResult struct {
	result *PromoteSeniorMemberByGuildNameResult
	err    error
}

func NewPromoteSeniorMemberByGuildNameResultFromJson(data string) PromoteSeniorMemberByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPromoteSeniorMemberByGuildNameResultFromDict(dict)
}

func NewPromoteSeniorMemberByGuildNameResultFromDict(data map[string]interface{}) PromoteSeniorMemberByGuildNameResult {
	return PromoteSeniorMemberByGuildNameResult{
		Item: func() *LastGuildMasterActivity {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLastGuildMasterActivityFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p PromoteSeniorMemberByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PromoteSeniorMemberByGuildNameResult) Pointer() *PromoteSeniorMemberByGuildNameResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentGuildMaster  `json:"item"`
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
		Item: func() *CurrentGuildMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGuildMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentGuildMasterResult struct {
	Item     *CurrentGuildMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentGuildMasterAsyncResult struct {
	result *GetCurrentGuildMasterResult
	err    error
}

func NewGetCurrentGuildMasterResultFromJson(data string) GetCurrentGuildMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentGuildMasterResultFromDict(dict)
}

func NewGetCurrentGuildMasterResultFromDict(data map[string]interface{}) GetCurrentGuildMasterResult {
	return GetCurrentGuildMasterResult{
		Item: func() *CurrentGuildMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGuildMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentGuildMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentGuildMasterResult) Pointer() *GetCurrentGuildMasterResult {
	return &p
}

type UpdateCurrentGuildMasterResult struct {
	Item     *CurrentGuildMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentGuildMasterAsyncResult struct {
	result *UpdateCurrentGuildMasterResult
	err    error
}

func NewUpdateCurrentGuildMasterResultFromJson(data string) UpdateCurrentGuildMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGuildMasterResultFromDict(dict)
}

func NewUpdateCurrentGuildMasterResultFromDict(data map[string]interface{}) UpdateCurrentGuildMasterResult {
	return UpdateCurrentGuildMasterResult{
		Item: func() *CurrentGuildMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGuildMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentGuildMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentGuildMasterResult) Pointer() *UpdateCurrentGuildMasterResult {
	return &p
}

type UpdateCurrentGuildMasterFromGitHubResult struct {
	Item     *CurrentGuildMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentGuildMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentGuildMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentGuildMasterFromGitHubResultFromJson(data string) UpdateCurrentGuildMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGuildMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentGuildMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentGuildMasterFromGitHubResult {
	return UpdateCurrentGuildMasterFromGitHubResult{
		Item: func() *CurrentGuildMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGuildMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentGuildMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentGuildMasterFromGitHubResult) Pointer() *UpdateCurrentGuildMasterFromGitHubResult {
	return &p
}

type DescribeReceiveRequestsResult struct {
	Items         []ReceiveMemberRequest `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeReceiveRequestsAsyncResult struct {
	result *DescribeReceiveRequestsResult
	err    error
}

func NewDescribeReceiveRequestsResultFromJson(data string) DescribeReceiveRequestsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveRequestsResultFromDict(dict)
}

func NewDescribeReceiveRequestsResultFromDict(data map[string]interface{}) DescribeReceiveRequestsResult {
	return DescribeReceiveRequestsResult{
		Items: func() []ReceiveMemberRequest {
			if data["items"] == nil {
				return nil
			}
			return CastReceiveMemberRequests(core.CastArray(data["items"]))
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

func (p DescribeReceiveRequestsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveMemberRequestsFromDict(
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

func (p DescribeReceiveRequestsResult) Pointer() *DescribeReceiveRequestsResult {
	return &p
}

type DescribeReceiveRequestsByGuildNameResult struct {
	Items         []ReceiveMemberRequest `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeReceiveRequestsByGuildNameAsyncResult struct {
	result *DescribeReceiveRequestsByGuildNameResult
	err    error
}

func NewDescribeReceiveRequestsByGuildNameResultFromJson(data string) DescribeReceiveRequestsByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveRequestsByGuildNameResultFromDict(dict)
}

func NewDescribeReceiveRequestsByGuildNameResultFromDict(data map[string]interface{}) DescribeReceiveRequestsByGuildNameResult {
	return DescribeReceiveRequestsByGuildNameResult{
		Items: func() []ReceiveMemberRequest {
			if data["items"] == nil {
				return nil
			}
			return CastReceiveMemberRequests(core.CastArray(data["items"]))
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

func (p DescribeReceiveRequestsByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastReceiveMemberRequestsFromDict(
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

func (p DescribeReceiveRequestsByGuildNameResult) Pointer() *DescribeReceiveRequestsByGuildNameResult {
	return &p
}

type GetReceiveRequestResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetReceiveRequestAsyncResult struct {
	result *GetReceiveRequestResult
	err    error
}

func NewGetReceiveRequestResultFromJson(data string) GetReceiveRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveRequestResultFromDict(dict)
}

func NewGetReceiveRequestResultFromDict(data map[string]interface{}) GetReceiveRequestResult {
	return GetReceiveRequestResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetReceiveRequestResult) ToDict() map[string]interface{} {
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

func (p GetReceiveRequestResult) Pointer() *GetReceiveRequestResult {
	return &p
}

type GetReceiveRequestByGuildNameResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetReceiveRequestByGuildNameAsyncResult struct {
	result *GetReceiveRequestByGuildNameResult
	err    error
}

func NewGetReceiveRequestByGuildNameResultFromJson(data string) GetReceiveRequestByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveRequestByGuildNameResultFromDict(dict)
}

func NewGetReceiveRequestByGuildNameResultFromDict(data map[string]interface{}) GetReceiveRequestByGuildNameResult {
	return GetReceiveRequestByGuildNameResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetReceiveRequestByGuildNameResult) ToDict() map[string]interface{} {
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

func (p GetReceiveRequestByGuildNameResult) Pointer() *GetReceiveRequestByGuildNameResult {
	return &p
}

type AcceptRequestResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Guild    *Guild                `json:"guild"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type AcceptRequestAsyncResult struct {
	result *AcceptRequestResult
	err    error
}

func NewAcceptRequestResultFromJson(data string) AcceptRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestResultFromDict(dict)
}

func NewAcceptRequestResultFromDict(data map[string]interface{}) AcceptRequestResult {
	return AcceptRequestResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p AcceptRequestResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcceptRequestResult) Pointer() *AcceptRequestResult {
	return &p
}

type AcceptRequestByGuildNameResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Guild    *Guild                `json:"guild"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type AcceptRequestByGuildNameAsyncResult struct {
	result *AcceptRequestByGuildNameResult
	err    error
}

func NewAcceptRequestByGuildNameResultFromJson(data string) AcceptRequestByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestByGuildNameResultFromDict(dict)
}

func NewAcceptRequestByGuildNameResultFromDict(data map[string]interface{}) AcceptRequestByGuildNameResult {
	return AcceptRequestByGuildNameResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p AcceptRequestByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcceptRequestByGuildNameResult) Pointer() *AcceptRequestByGuildNameResult {
	return &p
}

type RejectRequestResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type RejectRequestAsyncResult struct {
	result *RejectRequestResult
	err    error
}

func NewRejectRequestResultFromJson(data string) RejectRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRejectRequestResultFromDict(dict)
}

func NewRejectRequestResultFromDict(data map[string]interface{}) RejectRequestResult {
	return RejectRequestResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p RejectRequestResult) ToDict() map[string]interface{} {
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

func (p RejectRequestResult) Pointer() *RejectRequestResult {
	return &p
}

type RejectRequestByGuildNameResult struct {
	Item     *ReceiveMemberRequest `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type RejectRequestByGuildNameAsyncResult struct {
	result *RejectRequestByGuildNameResult
	err    error
}

func NewRejectRequestByGuildNameResultFromJson(data string) RejectRequestByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRejectRequestByGuildNameResultFromDict(dict)
}

func NewRejectRequestByGuildNameResultFromDict(data map[string]interface{}) RejectRequestByGuildNameResult {
	return RejectRequestByGuildNameResult{
		Item: func() *ReceiveMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewReceiveMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p RejectRequestByGuildNameResult) ToDict() map[string]interface{} {
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

func (p RejectRequestByGuildNameResult) Pointer() *RejectRequestByGuildNameResult {
	return &p
}

type DescribeSendRequestsResult struct {
	Items         []SendMemberRequest  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSendRequestsAsyncResult struct {
	result *DescribeSendRequestsResult
	err    error
}

func NewDescribeSendRequestsResultFromJson(data string) DescribeSendRequestsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSendRequestsResultFromDict(dict)
}

func NewDescribeSendRequestsResultFromDict(data map[string]interface{}) DescribeSendRequestsResult {
	return DescribeSendRequestsResult{
		Items: func() []SendMemberRequest {
			if data["items"] == nil {
				return nil
			}
			return CastSendMemberRequests(core.CastArray(data["items"]))
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

func (p DescribeSendRequestsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSendMemberRequestsFromDict(
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

func (p DescribeSendRequestsResult) Pointer() *DescribeSendRequestsResult {
	return &p
}

type DescribeSendRequestsByUserIdResult struct {
	Items         []SendMemberRequest  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSendRequestsByUserIdAsyncResult struct {
	result *DescribeSendRequestsByUserIdResult
	err    error
}

func NewDescribeSendRequestsByUserIdResultFromJson(data string) DescribeSendRequestsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSendRequestsByUserIdResultFromDict(dict)
}

func NewDescribeSendRequestsByUserIdResultFromDict(data map[string]interface{}) DescribeSendRequestsByUserIdResult {
	return DescribeSendRequestsByUserIdResult{
		Items: func() []SendMemberRequest {
			if data["items"] == nil {
				return nil
			}
			return CastSendMemberRequests(core.CastArray(data["items"]))
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

func (p DescribeSendRequestsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSendMemberRequestsFromDict(
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

func (p DescribeSendRequestsByUserIdResult) Pointer() *DescribeSendRequestsByUserIdResult {
	return &p
}

type GetSendRequestResult struct {
	Item     *SendMemberRequest   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSendRequestAsyncResult struct {
	result *GetSendRequestResult
	err    error
}

func NewGetSendRequestResultFromJson(data string) GetSendRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSendRequestResultFromDict(dict)
}

func NewGetSendRequestResultFromDict(data map[string]interface{}) GetSendRequestResult {
	return GetSendRequestResult{
		Item: func() *SendMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSendRequestResult) ToDict() map[string]interface{} {
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

func (p GetSendRequestResult) Pointer() *GetSendRequestResult {
	return &p
}

type GetSendRequestByUserIdResult struct {
	Item     *SendMemberRequest   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSendRequestByUserIdAsyncResult struct {
	result *GetSendRequestByUserIdResult
	err    error
}

func NewGetSendRequestByUserIdResultFromJson(data string) GetSendRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSendRequestByUserIdResultFromDict(dict)
}

func NewGetSendRequestByUserIdResultFromDict(data map[string]interface{}) GetSendRequestByUserIdResult {
	return GetSendRequestByUserIdResult{
		Item: func() *SendMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSendRequestByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetSendRequestByUserIdResult) Pointer() *GetSendRequestByUserIdResult {
	return &p
}

type SendRequestResult struct {
	Item              *Guild               `json:"item"`
	SendMemberRequest *SendMemberRequest   `json:"sendMemberRequest"`
	Metadata          *core.ResultMetadata `json:"metadata"`
}

type SendRequestAsyncResult struct {
	result *SendRequestResult
	err    error
}

func NewSendRequestResultFromJson(data string) SendRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendRequestResultFromDict(dict)
}

func NewSendRequestResultFromDict(data map[string]interface{}) SendRequestResult {
	return SendRequestResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		SendMemberRequest: func() *SendMemberRequest {
			v, ok := data["sendMemberRequest"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["sendMemberRequest"])).Pointer()
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

func (p SendRequestResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"sendMemberRequest": func() map[string]interface{} {
			if p.SendMemberRequest == nil {
				return nil
			}
			return p.SendMemberRequest.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SendRequestResult) Pointer() *SendRequestResult {
	return &p
}

type SendRequestByUserIdResult struct {
	Item              *Guild               `json:"item"`
	SendMemberRequest *SendMemberRequest   `json:"sendMemberRequest"`
	Metadata          *core.ResultMetadata `json:"metadata"`
}

type SendRequestByUserIdAsyncResult struct {
	result *SendRequestByUserIdResult
	err    error
}

func NewSendRequestByUserIdResultFromJson(data string) SendRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendRequestByUserIdResultFromDict(dict)
}

func NewSendRequestByUserIdResultFromDict(data map[string]interface{}) SendRequestByUserIdResult {
	return SendRequestByUserIdResult{
		Item: func() *Guild {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		SendMemberRequest: func() *SendMemberRequest {
			v, ok := data["sendMemberRequest"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["sendMemberRequest"])).Pointer()
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

func (p SendRequestByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"sendMemberRequest": func() map[string]interface{} {
			if p.SendMemberRequest == nil {
				return nil
			}
			return p.SendMemberRequest.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SendRequestByUserIdResult) Pointer() *SendRequestByUserIdResult {
	return &p
}

type DeleteRequestResult struct {
	Item     *SendMemberRequest   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteRequestAsyncResult struct {
	result *DeleteRequestResult
	err    error
}

func NewDeleteRequestResultFromJson(data string) DeleteRequestResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRequestResultFromDict(dict)
}

func NewDeleteRequestResultFromDict(data map[string]interface{}) DeleteRequestResult {
	return DeleteRequestResult{
		Item: func() *SendMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteRequestResult) ToDict() map[string]interface{} {
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

func (p DeleteRequestResult) Pointer() *DeleteRequestResult {
	return &p
}

type DeleteRequestByUserIdResult struct {
	Item     *SendMemberRequest   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteRequestByUserIdAsyncResult struct {
	result *DeleteRequestByUserIdResult
	err    error
}

func NewDeleteRequestByUserIdResultFromJson(data string) DeleteRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRequestByUserIdResultFromDict(dict)
}

func NewDeleteRequestByUserIdResultFromDict(data map[string]interface{}) DeleteRequestByUserIdResult {
	return DeleteRequestByUserIdResult{
		Item: func() *SendMemberRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSendMemberRequestFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteRequestByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteRequestByUserIdResult) Pointer() *DeleteRequestByUserIdResult {
	return &p
}

type DescribeIgnoreUsersResult struct {
	Items         []IgnoreUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeIgnoreUsersAsyncResult struct {
	result *DescribeIgnoreUsersResult
	err    error
}

func NewDescribeIgnoreUsersResultFromJson(data string) DescribeIgnoreUsersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIgnoreUsersResultFromDict(dict)
}

func NewDescribeIgnoreUsersResultFromDict(data map[string]interface{}) DescribeIgnoreUsersResult {
	return DescribeIgnoreUsersResult{
		Items: func() []IgnoreUser {
			if data["items"] == nil {
				return nil
			}
			return CastIgnoreUsers(core.CastArray(data["items"]))
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

func (p DescribeIgnoreUsersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIgnoreUsersFromDict(
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

func (p DescribeIgnoreUsersResult) Pointer() *DescribeIgnoreUsersResult {
	return &p
}

type DescribeIgnoreUsersByGuildNameResult struct {
	Items         []IgnoreUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeIgnoreUsersByGuildNameAsyncResult struct {
	result *DescribeIgnoreUsersByGuildNameResult
	err    error
}

func NewDescribeIgnoreUsersByGuildNameResultFromJson(data string) DescribeIgnoreUsersByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIgnoreUsersByGuildNameResultFromDict(dict)
}

func NewDescribeIgnoreUsersByGuildNameResultFromDict(data map[string]interface{}) DescribeIgnoreUsersByGuildNameResult {
	return DescribeIgnoreUsersByGuildNameResult{
		Items: func() []IgnoreUser {
			if data["items"] == nil {
				return nil
			}
			return CastIgnoreUsers(core.CastArray(data["items"]))
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

func (p DescribeIgnoreUsersByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIgnoreUsersFromDict(
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

func (p DescribeIgnoreUsersByGuildNameResult) Pointer() *DescribeIgnoreUsersByGuildNameResult {
	return &p
}

type GetIgnoreUserResult struct {
	Item     *IgnoreUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetIgnoreUserAsyncResult struct {
	result *GetIgnoreUserResult
	err    error
}

func NewGetIgnoreUserResultFromJson(data string) GetIgnoreUserResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIgnoreUserResultFromDict(dict)
}

func NewGetIgnoreUserResultFromDict(data map[string]interface{}) GetIgnoreUserResult {
	return GetIgnoreUserResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetIgnoreUserResult) ToDict() map[string]interface{} {
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

func (p GetIgnoreUserResult) Pointer() *GetIgnoreUserResult {
	return &p
}

type GetIgnoreUserByGuildNameResult struct {
	Item     *IgnoreUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetIgnoreUserByGuildNameAsyncResult struct {
	result *GetIgnoreUserByGuildNameResult
	err    error
}

func NewGetIgnoreUserByGuildNameResultFromJson(data string) GetIgnoreUserByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIgnoreUserByGuildNameResultFromDict(dict)
}

func NewGetIgnoreUserByGuildNameResultFromDict(data map[string]interface{}) GetIgnoreUserByGuildNameResult {
	return GetIgnoreUserByGuildNameResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetIgnoreUserByGuildNameResult) ToDict() map[string]interface{} {
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

func (p GetIgnoreUserByGuildNameResult) Pointer() *GetIgnoreUserByGuildNameResult {
	return &p
}

type AddIgnoreUserResult struct {
	Item     *IgnoreUser          `json:"item"`
	Guild    *Guild               `json:"guild"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddIgnoreUserAsyncResult struct {
	result *AddIgnoreUserResult
	err    error
}

func NewAddIgnoreUserResultFromJson(data string) AddIgnoreUserResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddIgnoreUserResultFromDict(dict)
}

func NewAddIgnoreUserResultFromDict(data map[string]interface{}) AddIgnoreUserResult {
	return AddIgnoreUserResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p AddIgnoreUserResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AddIgnoreUserResult) Pointer() *AddIgnoreUserResult {
	return &p
}

type AddIgnoreUserByGuildNameResult struct {
	Item     *IgnoreUser          `json:"item"`
	Guild    *Guild               `json:"guild"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddIgnoreUserByGuildNameAsyncResult struct {
	result *AddIgnoreUserByGuildNameResult
	err    error
}

func NewAddIgnoreUserByGuildNameResultFromJson(data string) AddIgnoreUserByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddIgnoreUserByGuildNameResultFromDict(dict)
}

func NewAddIgnoreUserByGuildNameResultFromDict(data map[string]interface{}) AddIgnoreUserByGuildNameResult {
	return AddIgnoreUserByGuildNameResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Guild: func() *Guild {
			v, ok := data["guild"]
			if !ok || v == nil {
				return nil
			}
			return NewGuildFromDict(core.CastMap(data["guild"])).Pointer()
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

func (p AddIgnoreUserByGuildNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"guild": func() map[string]interface{} {
			if p.Guild == nil {
				return nil
			}
			return p.Guild.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AddIgnoreUserByGuildNameResult) Pointer() *AddIgnoreUserByGuildNameResult {
	return &p
}

type DeleteIgnoreUserResult struct {
	Item     *IgnoreUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteIgnoreUserAsyncResult struct {
	result *DeleteIgnoreUserResult
	err    error
}

func NewDeleteIgnoreUserResultFromJson(data string) DeleteIgnoreUserResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteIgnoreUserResultFromDict(dict)
}

func NewDeleteIgnoreUserResultFromDict(data map[string]interface{}) DeleteIgnoreUserResult {
	return DeleteIgnoreUserResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteIgnoreUserResult) ToDict() map[string]interface{} {
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

func (p DeleteIgnoreUserResult) Pointer() *DeleteIgnoreUserResult {
	return &p
}

type DeleteIgnoreUserByGuildNameResult struct {
	Item     *IgnoreUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteIgnoreUserByGuildNameAsyncResult struct {
	result *DeleteIgnoreUserByGuildNameResult
	err    error
}

func NewDeleteIgnoreUserByGuildNameResultFromJson(data string) DeleteIgnoreUserByGuildNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteIgnoreUserByGuildNameResultFromDict(dict)
}

func NewDeleteIgnoreUserByGuildNameResultFromDict(data map[string]interface{}) DeleteIgnoreUserByGuildNameResult {
	return DeleteIgnoreUserByGuildNameResult{
		Item: func() *IgnoreUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIgnoreUserFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteIgnoreUserByGuildNameResult) ToDict() map[string]interface{} {
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

func (p DeleteIgnoreUserByGuildNameResult) Pointer() *DeleteIgnoreUserByGuildNameResult {
	return &p
}
