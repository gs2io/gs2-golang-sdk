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

package version

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

type DescribeVersionModelMastersResult struct {
	Items         []VersionModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeVersionModelMastersAsyncResult struct {
	result *DescribeVersionModelMastersResult
	err    error
}

func NewDescribeVersionModelMastersResultFromJson(data string) DescribeVersionModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeVersionModelMastersResultFromDict(dict)
}

func NewDescribeVersionModelMastersResultFromDict(data map[string]interface{}) DescribeVersionModelMastersResult {
	return DescribeVersionModelMastersResult{
		Items: func() []VersionModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastVersionModelMasters(core.CastArray(data["items"]))
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

func (p DescribeVersionModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastVersionModelMastersFromDict(
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

func (p DescribeVersionModelMastersResult) Pointer() *DescribeVersionModelMastersResult {
	return &p
}

type CreateVersionModelMasterResult struct {
	Item     *VersionModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateVersionModelMasterAsyncResult struct {
	result *CreateVersionModelMasterResult
	err    error
}

func NewCreateVersionModelMasterResultFromJson(data string) CreateVersionModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateVersionModelMasterResultFromDict(dict)
}

func NewCreateVersionModelMasterResultFromDict(data map[string]interface{}) CreateVersionModelMasterResult {
	return CreateVersionModelMasterResult{
		Item: func() *VersionModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateVersionModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateVersionModelMasterResult) Pointer() *CreateVersionModelMasterResult {
	return &p
}

type GetVersionModelMasterResult struct {
	Item     *VersionModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetVersionModelMasterAsyncResult struct {
	result *GetVersionModelMasterResult
	err    error
}

func NewGetVersionModelMasterResultFromJson(data string) GetVersionModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetVersionModelMasterResultFromDict(dict)
}

func NewGetVersionModelMasterResultFromDict(data map[string]interface{}) GetVersionModelMasterResult {
	return GetVersionModelMasterResult{
		Item: func() *VersionModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetVersionModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetVersionModelMasterResult) Pointer() *GetVersionModelMasterResult {
	return &p
}

type UpdateVersionModelMasterResult struct {
	Item     *VersionModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateVersionModelMasterAsyncResult struct {
	result *UpdateVersionModelMasterResult
	err    error
}

func NewUpdateVersionModelMasterResultFromJson(data string) UpdateVersionModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateVersionModelMasterResultFromDict(dict)
}

func NewUpdateVersionModelMasterResultFromDict(data map[string]interface{}) UpdateVersionModelMasterResult {
	return UpdateVersionModelMasterResult{
		Item: func() *VersionModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateVersionModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateVersionModelMasterResult) Pointer() *UpdateVersionModelMasterResult {
	return &p
}

type DeleteVersionModelMasterResult struct {
	Item     *VersionModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteVersionModelMasterAsyncResult struct {
	result *DeleteVersionModelMasterResult
	err    error
}

func NewDeleteVersionModelMasterResultFromJson(data string) DeleteVersionModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteVersionModelMasterResultFromDict(dict)
}

func NewDeleteVersionModelMasterResultFromDict(data map[string]interface{}) DeleteVersionModelMasterResult {
	return DeleteVersionModelMasterResult{
		Item: func() *VersionModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteVersionModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteVersionModelMasterResult) Pointer() *DeleteVersionModelMasterResult {
	return &p
}

type DescribeVersionModelsResult struct {
	Items    []VersionModel       `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeVersionModelsAsyncResult struct {
	result *DescribeVersionModelsResult
	err    error
}

func NewDescribeVersionModelsResultFromJson(data string) DescribeVersionModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeVersionModelsResultFromDict(dict)
}

func NewDescribeVersionModelsResultFromDict(data map[string]interface{}) DescribeVersionModelsResult {
	return DescribeVersionModelsResult{
		Items: func() []VersionModel {
			if data["items"] == nil {
				return nil
			}
			return CastVersionModels(core.CastArray(data["items"]))
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

func (p DescribeVersionModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastVersionModelsFromDict(
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

func (p DescribeVersionModelsResult) Pointer() *DescribeVersionModelsResult {
	return &p
}

type GetVersionModelResult struct {
	Item     *VersionModel        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetVersionModelAsyncResult struct {
	result *GetVersionModelResult
	err    error
}

func NewGetVersionModelResultFromJson(data string) GetVersionModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetVersionModelResultFromDict(dict)
}

func NewGetVersionModelResultFromDict(data map[string]interface{}) GetVersionModelResult {
	return GetVersionModelResult{
		Item: func() *VersionModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetVersionModelResult) ToDict() map[string]interface{} {
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

func (p GetVersionModelResult) Pointer() *GetVersionModelResult {
	return &p
}

type DescribeAcceptVersionsResult struct {
	Items         []AcceptVersion      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAcceptVersionsAsyncResult struct {
	result *DescribeAcceptVersionsResult
	err    error
}

func NewDescribeAcceptVersionsResultFromJson(data string) DescribeAcceptVersionsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAcceptVersionsResultFromDict(dict)
}

func NewDescribeAcceptVersionsResultFromDict(data map[string]interface{}) DescribeAcceptVersionsResult {
	return DescribeAcceptVersionsResult{
		Items: func() []AcceptVersion {
			if data["items"] == nil {
				return nil
			}
			return CastAcceptVersions(core.CastArray(data["items"]))
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

func (p DescribeAcceptVersionsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcceptVersionsFromDict(
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

func (p DescribeAcceptVersionsResult) Pointer() *DescribeAcceptVersionsResult {
	return &p
}

type DescribeAcceptVersionsByUserIdResult struct {
	Items         []AcceptVersion      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAcceptVersionsByUserIdAsyncResult struct {
	result *DescribeAcceptVersionsByUserIdResult
	err    error
}

func NewDescribeAcceptVersionsByUserIdResultFromJson(data string) DescribeAcceptVersionsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAcceptVersionsByUserIdResultFromDict(dict)
}

func NewDescribeAcceptVersionsByUserIdResultFromDict(data map[string]interface{}) DescribeAcceptVersionsByUserIdResult {
	return DescribeAcceptVersionsByUserIdResult{
		Items: func() []AcceptVersion {
			if data["items"] == nil {
				return nil
			}
			return CastAcceptVersions(core.CastArray(data["items"]))
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

func (p DescribeAcceptVersionsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcceptVersionsFromDict(
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

func (p DescribeAcceptVersionsByUserIdResult) Pointer() *DescribeAcceptVersionsByUserIdResult {
	return &p
}

type AcceptResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcceptAsyncResult struct {
	result *AcceptResult
	err    error
}

func NewAcceptResultFromJson(data string) AcceptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptResultFromDict(dict)
}

func NewAcceptResultFromDict(data map[string]interface{}) AcceptResult {
	return AcceptResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AcceptResult) ToDict() map[string]interface{} {
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

func (p AcceptResult) Pointer() *AcceptResult {
	return &p
}

type AcceptByUserIdResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcceptByUserIdAsyncResult struct {
	result *AcceptByUserIdResult
	err    error
}

func NewAcceptByUserIdResultFromJson(data string) AcceptByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptByUserIdResultFromDict(dict)
}

func NewAcceptByUserIdResultFromDict(data map[string]interface{}) AcceptByUserIdResult {
	return AcceptByUserIdResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AcceptByUserIdResult) ToDict() map[string]interface{} {
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

func (p AcceptByUserIdResult) Pointer() *AcceptByUserIdResult {
	return &p
}

type GetAcceptVersionResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAcceptVersionAsyncResult struct {
	result *GetAcceptVersionResult
	err    error
}

func NewGetAcceptVersionResultFromJson(data string) GetAcceptVersionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAcceptVersionResultFromDict(dict)
}

func NewGetAcceptVersionResultFromDict(data map[string]interface{}) GetAcceptVersionResult {
	return GetAcceptVersionResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetAcceptVersionResult) ToDict() map[string]interface{} {
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

func (p GetAcceptVersionResult) Pointer() *GetAcceptVersionResult {
	return &p
}

type GetAcceptVersionByUserIdResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAcceptVersionByUserIdAsyncResult struct {
	result *GetAcceptVersionByUserIdResult
	err    error
}

func NewGetAcceptVersionByUserIdResultFromJson(data string) GetAcceptVersionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAcceptVersionByUserIdResultFromDict(dict)
}

func NewGetAcceptVersionByUserIdResultFromDict(data map[string]interface{}) GetAcceptVersionByUserIdResult {
	return GetAcceptVersionByUserIdResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetAcceptVersionByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetAcceptVersionByUserIdResult) Pointer() *GetAcceptVersionByUserIdResult {
	return &p
}

type DeleteAcceptVersionResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAcceptVersionAsyncResult struct {
	result *DeleteAcceptVersionResult
	err    error
}

func NewDeleteAcceptVersionResultFromJson(data string) DeleteAcceptVersionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAcceptVersionResultFromDict(dict)
}

func NewDeleteAcceptVersionResultFromDict(data map[string]interface{}) DeleteAcceptVersionResult {
	return DeleteAcceptVersionResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteAcceptVersionResult) ToDict() map[string]interface{} {
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

func (p DeleteAcceptVersionResult) Pointer() *DeleteAcceptVersionResult {
	return &p
}

type DeleteAcceptVersionByUserIdResult struct {
	Item     *AcceptVersion       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAcceptVersionByUserIdAsyncResult struct {
	result *DeleteAcceptVersionByUserIdResult
	err    error
}

func NewDeleteAcceptVersionByUserIdResultFromJson(data string) DeleteAcceptVersionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAcceptVersionByUserIdResultFromDict(dict)
}

func NewDeleteAcceptVersionByUserIdResultFromDict(data map[string]interface{}) DeleteAcceptVersionByUserIdResult {
	return DeleteAcceptVersionByUserIdResult{
		Item: func() *AcceptVersion {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAcceptVersionFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteAcceptVersionByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteAcceptVersionByUserIdResult) Pointer() *DeleteAcceptVersionByUserIdResult {
	return &p
}

type CheckVersionResult struct {
	ProjectToken *string              `json:"projectToken"`
	Warnings     []Status             `json:"warnings"`
	Errors       []Status             `json:"errors"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type CheckVersionAsyncResult struct {
	result *CheckVersionResult
	err    error
}

func NewCheckVersionResultFromJson(data string) CheckVersionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckVersionResultFromDict(dict)
}

func NewCheckVersionResultFromDict(data map[string]interface{}) CheckVersionResult {
	return CheckVersionResult{
		ProjectToken: func() *string {
			v, ok := data["projectToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectToken"])
		}(),
		Warnings: func() []Status {
			if data["warnings"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["warnings"]))
		}(),
		Errors: func() []Status {
			if data["errors"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["errors"]))
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

func (p CheckVersionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectToken": p.ProjectToken,
		"warnings": CastStatusesFromDict(
			p.Warnings,
		),
		"errors": CastStatusesFromDict(
			p.Errors,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CheckVersionResult) Pointer() *CheckVersionResult {
	return &p
}

type CheckVersionByUserIdResult struct {
	ProjectToken *string              `json:"projectToken"`
	Warnings     []Status             `json:"warnings"`
	Errors       []Status             `json:"errors"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type CheckVersionByUserIdAsyncResult struct {
	result *CheckVersionByUserIdResult
	err    error
}

func NewCheckVersionByUserIdResultFromJson(data string) CheckVersionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckVersionByUserIdResultFromDict(dict)
}

func NewCheckVersionByUserIdResultFromDict(data map[string]interface{}) CheckVersionByUserIdResult {
	return CheckVersionByUserIdResult{
		ProjectToken: func() *string {
			v, ok := data["projectToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["projectToken"])
		}(),
		Warnings: func() []Status {
			if data["warnings"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["warnings"]))
		}(),
		Errors: func() []Status {
			if data["errors"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["errors"]))
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

func (p CheckVersionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"projectToken": p.ProjectToken,
		"warnings": CastStatusesFromDict(
			p.Warnings,
		),
		"errors": CastStatusesFromDict(
			p.Errors,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CheckVersionByUserIdResult) Pointer() *CheckVersionByUserIdResult {
	return &p
}

type CalculateSignatureResult struct {
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type CalculateSignatureAsyncResult struct {
	result *CalculateSignatureResult
	err    error
}

func NewCalculateSignatureResultFromJson(data string) CalculateSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCalculateSignatureResultFromDict(dict)
}

func NewCalculateSignatureResultFromDict(data map[string]interface{}) CalculateSignatureResult {
	return CalculateSignatureResult{
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

func (p CalculateSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func (p CalculateSignatureResult) Pointer() *CalculateSignatureResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentVersionMaster `json:"item"`
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
		Item: func() *CurrentVersionMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentVersionMasterResult struct {
	Item     *CurrentVersionMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetCurrentVersionMasterAsyncResult struct {
	result *GetCurrentVersionMasterResult
	err    error
}

func NewGetCurrentVersionMasterResultFromJson(data string) GetCurrentVersionMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentVersionMasterResultFromDict(dict)
}

func NewGetCurrentVersionMasterResultFromDict(data map[string]interface{}) GetCurrentVersionMasterResult {
	return GetCurrentVersionMasterResult{
		Item: func() *CurrentVersionMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentVersionMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentVersionMasterResult) Pointer() *GetCurrentVersionMasterResult {
	return &p
}

type UpdateCurrentVersionMasterResult struct {
	Item     *CurrentVersionMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentVersionMasterAsyncResult struct {
	result *UpdateCurrentVersionMasterResult
	err    error
}

func NewUpdateCurrentVersionMasterResultFromJson(data string) UpdateCurrentVersionMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentVersionMasterResultFromDict(dict)
}

func NewUpdateCurrentVersionMasterResultFromDict(data map[string]interface{}) UpdateCurrentVersionMasterResult {
	return UpdateCurrentVersionMasterResult{
		Item: func() *CurrentVersionMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentVersionMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentVersionMasterResult) Pointer() *UpdateCurrentVersionMasterResult {
	return &p
}

type UpdateCurrentVersionMasterFromGitHubResult struct {
	Item     *CurrentVersionMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentVersionMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentVersionMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentVersionMasterFromGitHubResultFromJson(data string) UpdateCurrentVersionMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentVersionMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentVersionMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentVersionMasterFromGitHubResult {
	return UpdateCurrentVersionMasterFromGitHubResult{
		Item: func() *CurrentVersionMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentVersionMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentVersionMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentVersionMasterFromGitHubResult) Pointer() *UpdateCurrentVersionMasterFromGitHubResult {
	return &p
}
