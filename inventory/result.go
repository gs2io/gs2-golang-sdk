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
	"github.com/gs2io/gs2-golang-sdk/grade"
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

type GetServiceVersionResult struct {
	Item     *string              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetServiceVersionAsyncResult struct {
	result *GetServiceVersionResult
	err    error
}

func NewGetServiceVersionResultFromJson(data string) GetServiceVersionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetServiceVersionResultFromDict(dict)
}

func NewGetServiceVersionResultFromDict(data map[string]interface{}) GetServiceVersionResult {
	return GetServiceVersionResult{
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
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

func (p GetServiceVersionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetServiceVersionResult) Pointer() *GetServiceVersionResult {
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

type DescribeInventoryModelMastersResult struct {
	Items         []InventoryModelMaster `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
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
		Items: func() []InventoryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastInventoryModelMasters(core.CastArray(data["items"]))
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

func (p DescribeInventoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryModelMastersFromDict(
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

func (p DescribeInventoryModelMastersResult) Pointer() *DescribeInventoryModelMastersResult {
	return &p
}

type CreateInventoryModelMasterResult struct {
	Item     *InventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
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
		Item: func() *InventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateInventoryModelMasterResult) Pointer() *CreateInventoryModelMasterResult {
	return &p
}

type GetInventoryModelMasterResult struct {
	Item     *InventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
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
		Item: func() *InventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetInventoryModelMasterResult) Pointer() *GetInventoryModelMasterResult {
	return &p
}

type UpdateInventoryModelMasterResult struct {
	Item     *InventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
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
		Item: func() *InventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateInventoryModelMasterResult) Pointer() *UpdateInventoryModelMasterResult {
	return &p
}

type DeleteInventoryModelMasterResult struct {
	Item     *InventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
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
		Item: func() *InventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteInventoryModelMasterResult) Pointer() *DeleteInventoryModelMasterResult {
	return &p
}

type DescribeInventoryModelsResult struct {
	Items    []InventoryModel     `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Items: func() []InventoryModel {
			if data["items"] == nil {
				return nil
			}
			return CastInventoryModels(core.CastArray(data["items"]))
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

func (p DescribeInventoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoryModelsFromDict(
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

func (p DescribeInventoryModelsResult) Pointer() *DescribeInventoryModelsResult {
	return &p
}

type GetInventoryModelResult struct {
	Item     *InventoryModel      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *InventoryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetInventoryModelResult) ToDict() map[string]interface{} {
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

func (p GetInventoryModelResult) Pointer() *GetInventoryModelResult {
	return &p
}

type DescribeItemModelMastersResult struct {
	Items         []ItemModelMaster    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastItemModelMasters(core.CastArray(data["items"]))
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

func (p DescribeItemModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemModelMastersFromDict(
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

func (p DescribeItemModelMastersResult) Pointer() *DescribeItemModelMastersResult {
	return &p
}

type CreateItemModelMasterResult struct {
	Item     *ItemModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *ItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateItemModelMasterResult) Pointer() *CreateItemModelMasterResult {
	return &p
}

type GetItemModelMasterResult struct {
	Item     *ItemModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *ItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetItemModelMasterResult) Pointer() *GetItemModelMasterResult {
	return &p
}

type UpdateItemModelMasterResult struct {
	Item     *ItemModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *ItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateItemModelMasterResult) Pointer() *UpdateItemModelMasterResult {
	return &p
}

type DeleteItemModelMasterResult struct {
	Item     *ItemModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *ItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteItemModelMasterResult) Pointer() *DeleteItemModelMasterResult {
	return &p
}

type DescribeItemModelsResult struct {
	Items    []ItemModel          `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemModel {
			if data["items"] == nil {
				return nil
			}
			return CastItemModels(core.CastArray(data["items"]))
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

func (p DescribeItemModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemModelsFromDict(
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

func (p DescribeItemModelsResult) Pointer() *DescribeItemModelsResult {
	return &p
}

type GetItemModelResult struct {
	Item     *ItemModel           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *ItemModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetItemModelResult) ToDict() map[string]interface{} {
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

func (p GetItemModelResult) Pointer() *GetItemModelResult {
	return &p
}

type DescribeSimpleInventoryModelMastersResult struct {
	Items         []SimpleInventoryModelMaster `json:"items"`
	NextPageToken *string                      `json:"nextPageToken"`
	Metadata      *core.ResultMetadata         `json:"metadata"`
}

type DescribeSimpleInventoryModelMastersAsyncResult struct {
	result *DescribeSimpleInventoryModelMastersResult
	err    error
}

func NewDescribeSimpleInventoryModelMastersResultFromJson(data string) DescribeSimpleInventoryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleInventoryModelMastersResultFromDict(dict)
}

func NewDescribeSimpleInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeSimpleInventoryModelMastersResult {
	return DescribeSimpleInventoryModelMastersResult{
		Items: func() []SimpleInventoryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleInventoryModelMasters(core.CastArray(data["items"]))
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

func (p DescribeSimpleInventoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleInventoryModelMastersFromDict(
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

func (p DescribeSimpleInventoryModelMastersResult) Pointer() *DescribeSimpleInventoryModelMastersResult {
	return &p
}

type CreateSimpleInventoryModelMasterResult struct {
	Item     *SimpleInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type CreateSimpleInventoryModelMasterAsyncResult struct {
	result *CreateSimpleInventoryModelMasterResult
	err    error
}

func NewCreateSimpleInventoryModelMasterResultFromJson(data string) CreateSimpleInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSimpleInventoryModelMasterResultFromDict(dict)
}

func NewCreateSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) CreateSimpleInventoryModelMasterResult {
	return CreateSimpleInventoryModelMasterResult{
		Item: func() *SimpleInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateSimpleInventoryModelMasterResult) Pointer() *CreateSimpleInventoryModelMasterResult {
	return &p
}

type GetSimpleInventoryModelMasterResult struct {
	Item     *SimpleInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type GetSimpleInventoryModelMasterAsyncResult struct {
	result *GetSimpleInventoryModelMasterResult
	err    error
}

func NewGetSimpleInventoryModelMasterResultFromJson(data string) GetSimpleInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleInventoryModelMasterResultFromDict(dict)
}

func NewGetSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) GetSimpleInventoryModelMasterResult {
	return GetSimpleInventoryModelMasterResult{
		Item: func() *SimpleInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetSimpleInventoryModelMasterResult) Pointer() *GetSimpleInventoryModelMasterResult {
	return &p
}

type UpdateSimpleInventoryModelMasterResult struct {
	Item     *SimpleInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type UpdateSimpleInventoryModelMasterAsyncResult struct {
	result *UpdateSimpleInventoryModelMasterResult
	err    error
}

func NewUpdateSimpleInventoryModelMasterResultFromJson(data string) UpdateSimpleInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSimpleInventoryModelMasterResultFromDict(dict)
}

func NewUpdateSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateSimpleInventoryModelMasterResult {
	return UpdateSimpleInventoryModelMasterResult{
		Item: func() *SimpleInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateSimpleInventoryModelMasterResult) Pointer() *UpdateSimpleInventoryModelMasterResult {
	return &p
}

type DeleteSimpleInventoryModelMasterResult struct {
	Item     *SimpleInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type DeleteSimpleInventoryModelMasterAsyncResult struct {
	result *DeleteSimpleInventoryModelMasterResult
	err    error
}

func NewDeleteSimpleInventoryModelMasterResultFromJson(data string) DeleteSimpleInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleInventoryModelMasterResultFromDict(dict)
}

func NewDeleteSimpleInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteSimpleInventoryModelMasterResult {
	return DeleteSimpleInventoryModelMasterResult{
		Item: func() *SimpleInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteSimpleInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteSimpleInventoryModelMasterResult) Pointer() *DeleteSimpleInventoryModelMasterResult {
	return &p
}

type DescribeSimpleInventoryModelsResult struct {
	Items    []SimpleInventoryModel `json:"items"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DescribeSimpleInventoryModelsAsyncResult struct {
	result *DescribeSimpleInventoryModelsResult
	err    error
}

func NewDescribeSimpleInventoryModelsResultFromJson(data string) DescribeSimpleInventoryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleInventoryModelsResultFromDict(dict)
}

func NewDescribeSimpleInventoryModelsResultFromDict(data map[string]interface{}) DescribeSimpleInventoryModelsResult {
	return DescribeSimpleInventoryModelsResult{
		Items: func() []SimpleInventoryModel {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleInventoryModels(core.CastArray(data["items"]))
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

func (p DescribeSimpleInventoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleInventoryModelsFromDict(
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

func (p DescribeSimpleInventoryModelsResult) Pointer() *DescribeSimpleInventoryModelsResult {
	return &p
}

type GetSimpleInventoryModelResult struct {
	Item     *SimpleInventoryModel `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetSimpleInventoryModelAsyncResult struct {
	result *GetSimpleInventoryModelResult
	err    error
}

func NewGetSimpleInventoryModelResultFromJson(data string) GetSimpleInventoryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleInventoryModelResultFromDict(dict)
}

func NewGetSimpleInventoryModelResultFromDict(data map[string]interface{}) GetSimpleInventoryModelResult {
	return GetSimpleInventoryModelResult{
		Item: func() *SimpleInventoryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleInventoryModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSimpleInventoryModelResult) ToDict() map[string]interface{} {
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

func (p GetSimpleInventoryModelResult) Pointer() *GetSimpleInventoryModelResult {
	return &p
}

type DescribeSimpleItemModelMastersResult struct {
	Items         []SimpleItemModelMaster `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeSimpleItemModelMastersAsyncResult struct {
	result *DescribeSimpleItemModelMastersResult
	err    error
}

func NewDescribeSimpleItemModelMastersResultFromJson(data string) DescribeSimpleItemModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemModelMastersResultFromDict(dict)
}

func NewDescribeSimpleItemModelMastersResultFromDict(data map[string]interface{}) DescribeSimpleItemModelMastersResult {
	return DescribeSimpleItemModelMastersResult{
		Items: func() []SimpleItemModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItemModelMasters(core.CastArray(data["items"]))
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

func (p DescribeSimpleItemModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemModelMastersFromDict(
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

func (p DescribeSimpleItemModelMastersResult) Pointer() *DescribeSimpleItemModelMastersResult {
	return &p
}

type CreateSimpleItemModelMasterResult struct {
	Item     *SimpleItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type CreateSimpleItemModelMasterAsyncResult struct {
	result *CreateSimpleItemModelMasterResult
	err    error
}

func NewCreateSimpleItemModelMasterResultFromJson(data string) CreateSimpleItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSimpleItemModelMasterResultFromDict(dict)
}

func NewCreateSimpleItemModelMasterResultFromDict(data map[string]interface{}) CreateSimpleItemModelMasterResult {
	return CreateSimpleItemModelMasterResult{
		Item: func() *SimpleItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateSimpleItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateSimpleItemModelMasterResult) Pointer() *CreateSimpleItemModelMasterResult {
	return &p
}

type GetSimpleItemModelMasterResult struct {
	Item     *SimpleItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetSimpleItemModelMasterAsyncResult struct {
	result *GetSimpleItemModelMasterResult
	err    error
}

func NewGetSimpleItemModelMasterResultFromJson(data string) GetSimpleItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemModelMasterResultFromDict(dict)
}

func NewGetSimpleItemModelMasterResultFromDict(data map[string]interface{}) GetSimpleItemModelMasterResult {
	return GetSimpleItemModelMasterResult{
		Item: func() *SimpleItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSimpleItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetSimpleItemModelMasterResult) Pointer() *GetSimpleItemModelMasterResult {
	return &p
}

type UpdateSimpleItemModelMasterResult struct {
	Item     *SimpleItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateSimpleItemModelMasterAsyncResult struct {
	result *UpdateSimpleItemModelMasterResult
	err    error
}

func NewUpdateSimpleItemModelMasterResultFromJson(data string) UpdateSimpleItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSimpleItemModelMasterResultFromDict(dict)
}

func NewUpdateSimpleItemModelMasterResultFromDict(data map[string]interface{}) UpdateSimpleItemModelMasterResult {
	return UpdateSimpleItemModelMasterResult{
		Item: func() *SimpleItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateSimpleItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateSimpleItemModelMasterResult) Pointer() *UpdateSimpleItemModelMasterResult {
	return &p
}

type DeleteSimpleItemModelMasterResult struct {
	Item     *SimpleItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DeleteSimpleItemModelMasterAsyncResult struct {
	result *DeleteSimpleItemModelMasterResult
	err    error
}

func NewDeleteSimpleItemModelMasterResultFromJson(data string) DeleteSimpleItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleItemModelMasterResultFromDict(dict)
}

func NewDeleteSimpleItemModelMasterResultFromDict(data map[string]interface{}) DeleteSimpleItemModelMasterResult {
	return DeleteSimpleItemModelMasterResult{
		Item: func() *SimpleItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteSimpleItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteSimpleItemModelMasterResult) Pointer() *DeleteSimpleItemModelMasterResult {
	return &p
}

type DescribeSimpleItemModelsResult struct {
	Items    []SimpleItemModel    `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSimpleItemModelsAsyncResult struct {
	result *DescribeSimpleItemModelsResult
	err    error
}

func NewDescribeSimpleItemModelsResultFromJson(data string) DescribeSimpleItemModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemModelsResultFromDict(dict)
}

func NewDescribeSimpleItemModelsResultFromDict(data map[string]interface{}) DescribeSimpleItemModelsResult {
	return DescribeSimpleItemModelsResult{
		Items: func() []SimpleItemModel {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItemModels(core.CastArray(data["items"]))
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

func (p DescribeSimpleItemModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemModelsFromDict(
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

func (p DescribeSimpleItemModelsResult) Pointer() *DescribeSimpleItemModelsResult {
	return &p
}

type GetSimpleItemModelResult struct {
	Item     *SimpleItemModel     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSimpleItemModelAsyncResult struct {
	result *GetSimpleItemModelResult
	err    error
}

func NewGetSimpleItemModelResultFromJson(data string) GetSimpleItemModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemModelResultFromDict(dict)
}

func NewGetSimpleItemModelResultFromDict(data map[string]interface{}) GetSimpleItemModelResult {
	return GetSimpleItemModelResult{
		Item: func() *SimpleItemModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetSimpleItemModelResult) ToDict() map[string]interface{} {
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

func (p GetSimpleItemModelResult) Pointer() *GetSimpleItemModelResult {
	return &p
}

type DescribeBigInventoryModelMastersResult struct {
	Items         []BigInventoryModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeBigInventoryModelMastersAsyncResult struct {
	result *DescribeBigInventoryModelMastersResult
	err    error
}

func NewDescribeBigInventoryModelMastersResultFromJson(data string) DescribeBigInventoryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigInventoryModelMastersResultFromDict(dict)
}

func NewDescribeBigInventoryModelMastersResultFromDict(data map[string]interface{}) DescribeBigInventoryModelMastersResult {
	return DescribeBigInventoryModelMastersResult{
		Items: func() []BigInventoryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastBigInventoryModelMasters(core.CastArray(data["items"]))
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

func (p DescribeBigInventoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigInventoryModelMastersFromDict(
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

func (p DescribeBigInventoryModelMastersResult) Pointer() *DescribeBigInventoryModelMastersResult {
	return &p
}

type CreateBigInventoryModelMasterResult struct {
	Item     *BigInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type CreateBigInventoryModelMasterAsyncResult struct {
	result *CreateBigInventoryModelMasterResult
	err    error
}

func NewCreateBigInventoryModelMasterResultFromJson(data string) CreateBigInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBigInventoryModelMasterResultFromDict(dict)
}

func NewCreateBigInventoryModelMasterResultFromDict(data map[string]interface{}) CreateBigInventoryModelMasterResult {
	return CreateBigInventoryModelMasterResult{
		Item: func() *BigInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateBigInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateBigInventoryModelMasterResult) Pointer() *CreateBigInventoryModelMasterResult {
	return &p
}

type GetBigInventoryModelMasterResult struct {
	Item     *BigInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetBigInventoryModelMasterAsyncResult struct {
	result *GetBigInventoryModelMasterResult
	err    error
}

func NewGetBigInventoryModelMasterResultFromJson(data string) GetBigInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigInventoryModelMasterResultFromDict(dict)
}

func NewGetBigInventoryModelMasterResultFromDict(data map[string]interface{}) GetBigInventoryModelMasterResult {
	return GetBigInventoryModelMasterResult{
		Item: func() *BigInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBigInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetBigInventoryModelMasterResult) Pointer() *GetBigInventoryModelMasterResult {
	return &p
}

type UpdateBigInventoryModelMasterResult struct {
	Item     *BigInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type UpdateBigInventoryModelMasterAsyncResult struct {
	result *UpdateBigInventoryModelMasterResult
	err    error
}

func NewUpdateBigInventoryModelMasterResultFromJson(data string) UpdateBigInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBigInventoryModelMasterResultFromDict(dict)
}

func NewUpdateBigInventoryModelMasterResultFromDict(data map[string]interface{}) UpdateBigInventoryModelMasterResult {
	return UpdateBigInventoryModelMasterResult{
		Item: func() *BigInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateBigInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateBigInventoryModelMasterResult) Pointer() *UpdateBigInventoryModelMasterResult {
	return &p
}

type DeleteBigInventoryModelMasterResult struct {
	Item     *BigInventoryModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type DeleteBigInventoryModelMasterAsyncResult struct {
	result *DeleteBigInventoryModelMasterResult
	err    error
}

func NewDeleteBigInventoryModelMasterResultFromJson(data string) DeleteBigInventoryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigInventoryModelMasterResultFromDict(dict)
}

func NewDeleteBigInventoryModelMasterResultFromDict(data map[string]interface{}) DeleteBigInventoryModelMasterResult {
	return DeleteBigInventoryModelMasterResult{
		Item: func() *BigInventoryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigInventoryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteBigInventoryModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteBigInventoryModelMasterResult) Pointer() *DeleteBigInventoryModelMasterResult {
	return &p
}

type DescribeBigInventoryModelsResult struct {
	Items    []BigInventoryModel  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeBigInventoryModelsAsyncResult struct {
	result *DescribeBigInventoryModelsResult
	err    error
}

func NewDescribeBigInventoryModelsResultFromJson(data string) DescribeBigInventoryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigInventoryModelsResultFromDict(dict)
}

func NewDescribeBigInventoryModelsResultFromDict(data map[string]interface{}) DescribeBigInventoryModelsResult {
	return DescribeBigInventoryModelsResult{
		Items: func() []BigInventoryModel {
			if data["items"] == nil {
				return nil
			}
			return CastBigInventoryModels(core.CastArray(data["items"]))
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

func (p DescribeBigInventoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigInventoryModelsFromDict(
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

func (p DescribeBigInventoryModelsResult) Pointer() *DescribeBigInventoryModelsResult {
	return &p
}

type GetBigInventoryModelResult struct {
	Item     *BigInventoryModel   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBigInventoryModelAsyncResult struct {
	result *GetBigInventoryModelResult
	err    error
}

func NewGetBigInventoryModelResultFromJson(data string) GetBigInventoryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigInventoryModelResultFromDict(dict)
}

func NewGetBigInventoryModelResultFromDict(data map[string]interface{}) GetBigInventoryModelResult {
	return GetBigInventoryModelResult{
		Item: func() *BigInventoryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigInventoryModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBigInventoryModelResult) ToDict() map[string]interface{} {
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

func (p GetBigInventoryModelResult) Pointer() *GetBigInventoryModelResult {
	return &p
}

type DescribeBigItemModelMastersResult struct {
	Items         []BigItemModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBigItemModelMastersAsyncResult struct {
	result *DescribeBigItemModelMastersResult
	err    error
}

func NewDescribeBigItemModelMastersResultFromJson(data string) DescribeBigItemModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemModelMastersResultFromDict(dict)
}

func NewDescribeBigItemModelMastersResultFromDict(data map[string]interface{}) DescribeBigItemModelMastersResult {
	return DescribeBigItemModelMastersResult{
		Items: func() []BigItemModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastBigItemModelMasters(core.CastArray(data["items"]))
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

func (p DescribeBigItemModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigItemModelMastersFromDict(
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

func (p DescribeBigItemModelMastersResult) Pointer() *DescribeBigItemModelMastersResult {
	return &p
}

type CreateBigItemModelMasterResult struct {
	Item     *BigItemModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateBigItemModelMasterAsyncResult struct {
	result *CreateBigItemModelMasterResult
	err    error
}

func NewCreateBigItemModelMasterResultFromJson(data string) CreateBigItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBigItemModelMasterResultFromDict(dict)
}

func NewCreateBigItemModelMasterResultFromDict(data map[string]interface{}) CreateBigItemModelMasterResult {
	return CreateBigItemModelMasterResult{
		Item: func() *BigItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p CreateBigItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateBigItemModelMasterResult) Pointer() *CreateBigItemModelMasterResult {
	return &p
}

type GetBigItemModelMasterResult struct {
	Item     *BigItemModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBigItemModelMasterAsyncResult struct {
	result *GetBigItemModelMasterResult
	err    error
}

func NewGetBigItemModelMasterResultFromJson(data string) GetBigItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemModelMasterResultFromDict(dict)
}

func NewGetBigItemModelMasterResultFromDict(data map[string]interface{}) GetBigItemModelMasterResult {
	return GetBigItemModelMasterResult{
		Item: func() *BigItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBigItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetBigItemModelMasterResult) Pointer() *GetBigItemModelMasterResult {
	return &p
}

type UpdateBigItemModelMasterResult struct {
	Item     *BigItemModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateBigItemModelMasterAsyncResult struct {
	result *UpdateBigItemModelMasterResult
	err    error
}

func NewUpdateBigItemModelMasterResultFromJson(data string) UpdateBigItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBigItemModelMasterResultFromDict(dict)
}

func NewUpdateBigItemModelMasterResultFromDict(data map[string]interface{}) UpdateBigItemModelMasterResult {
	return UpdateBigItemModelMasterResult{
		Item: func() *BigItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateBigItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateBigItemModelMasterResult) Pointer() *UpdateBigItemModelMasterResult {
	return &p
}

type DeleteBigItemModelMasterResult struct {
	Item     *BigItemModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteBigItemModelMasterAsyncResult struct {
	result *DeleteBigItemModelMasterResult
	err    error
}

func NewDeleteBigItemModelMasterResultFromJson(data string) DeleteBigItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigItemModelMasterResultFromDict(dict)
}

func NewDeleteBigItemModelMasterResultFromDict(data map[string]interface{}) DeleteBigItemModelMasterResult {
	return DeleteBigItemModelMasterResult{
		Item: func() *BigItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteBigItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteBigItemModelMasterResult) Pointer() *DeleteBigItemModelMasterResult {
	return &p
}

type DescribeBigItemModelsResult struct {
	Items    []BigItemModel       `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeBigItemModelsAsyncResult struct {
	result *DescribeBigItemModelsResult
	err    error
}

func NewDescribeBigItemModelsResultFromJson(data string) DescribeBigItemModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemModelsResultFromDict(dict)
}

func NewDescribeBigItemModelsResultFromDict(data map[string]interface{}) DescribeBigItemModelsResult {
	return DescribeBigItemModelsResult{
		Items: func() []BigItemModel {
			if data["items"] == nil {
				return nil
			}
			return CastBigItemModels(core.CastArray(data["items"]))
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

func (p DescribeBigItemModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigItemModelsFromDict(
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

func (p DescribeBigItemModelsResult) Pointer() *DescribeBigItemModelsResult {
	return &p
}

type GetBigItemModelResult struct {
	Item     *BigItemModel        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBigItemModelAsyncResult struct {
	result *GetBigItemModelResult
	err    error
}

func NewGetBigItemModelResultFromJson(data string) GetBigItemModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemModelResultFromDict(dict)
}

func NewGetBigItemModelResultFromDict(data map[string]interface{}) GetBigItemModelResult {
	return GetBigItemModelResult{
		Item: func() *BigItemModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBigItemModelResult) ToDict() map[string]interface{} {
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

func (p GetBigItemModelResult) Pointer() *GetBigItemModelResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *CurrentItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentItemModelMasterResult struct {
	Item     *CurrentItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *CurrentItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetCurrentItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentItemModelMasterResult) Pointer() *GetCurrentItemModelMasterResult {
	return &p
}

type PreUpdateCurrentItemModelMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentItemModelMasterAsyncResult struct {
	result *PreUpdateCurrentItemModelMasterResult
	err    error
}

func NewPreUpdateCurrentItemModelMasterResultFromJson(data string) PreUpdateCurrentItemModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentItemModelMasterResultFromDict(dict)
}

func NewPreUpdateCurrentItemModelMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentItemModelMasterResult {
	return PreUpdateCurrentItemModelMasterResult{
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

func (p PreUpdateCurrentItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentItemModelMasterResult) Pointer() *PreUpdateCurrentItemModelMasterResult {
	return &p
}

type UpdateCurrentItemModelMasterResult struct {
	Item     *CurrentItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *CurrentItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentItemModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentItemModelMasterResult) Pointer() *UpdateCurrentItemModelMasterResult {
	return &p
}

type UpdateCurrentItemModelMasterFromGitHubResult struct {
	Item     *CurrentItemModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
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
		Item: func() *CurrentItemModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentItemModelMasterFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UpdateCurrentItemModelMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentItemModelMasterFromGitHubResult) Pointer() *UpdateCurrentItemModelMasterFromGitHubResult {
	return &p
}

type DescribeInventoriesResult struct {
	Items         []Inventory          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Inventory {
			if data["items"] == nil {
				return nil
			}
			return CastInventories(core.CastArray(data["items"]))
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

func (p DescribeInventoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoriesFromDict(
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

func (p DescribeInventoriesResult) Pointer() *DescribeInventoriesResult {
	return &p
}

type DescribeInventoriesByUserIdResult struct {
	Items         []Inventory          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []Inventory {
			if data["items"] == nil {
				return nil
			}
			return CastInventories(core.CastArray(data["items"]))
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

func (p DescribeInventoriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastInventoriesFromDict(
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

func (p DescribeInventoriesByUserIdResult) Pointer() *DescribeInventoriesByUserIdResult {
	return &p
}

type GetInventoryResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetInventoryResult) ToDict() map[string]interface{} {
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

func (p GetInventoryResult) Pointer() *GetInventoryResult {
	return &p
}

type GetInventoryByUserIdResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetInventoryByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetInventoryByUserIdResult) Pointer() *GetInventoryByUserIdResult {
	return &p
}

type AddCapacityByUserIdResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AddCapacityByUserIdResult) ToDict() map[string]interface{} {
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

func (p AddCapacityByUserIdResult) Pointer() *AddCapacityByUserIdResult {
	return &p
}

type SetCapacityByUserIdResult struct {
	Item     *Inventory           `json:"item"`
	Old      *Inventory           `json:"old"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Inventory {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["old"])).Pointer()
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

func (p SetCapacityByUserIdResult) ToDict() map[string]interface{} {
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

func (p SetCapacityByUserIdResult) Pointer() *SetCapacityByUserIdResult {
	return &p
}

type DeleteInventoryByUserIdResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteInventoryByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteInventoryByUserIdResult) Pointer() *DeleteInventoryByUserIdResult {
	return &p
}

type VerifyInventoryCurrentMaxCapacityResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyInventoryCurrentMaxCapacityAsyncResult struct {
	result *VerifyInventoryCurrentMaxCapacityResult
	err    error
}

func NewVerifyInventoryCurrentMaxCapacityResultFromJson(data string) VerifyInventoryCurrentMaxCapacityResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityResultFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityResultFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityResult {
	return VerifyInventoryCurrentMaxCapacityResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyInventoryCurrentMaxCapacityResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyInventoryCurrentMaxCapacityResult) Pointer() *VerifyInventoryCurrentMaxCapacityResult {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult struct {
	result *VerifyInventoryCurrentMaxCapacityByUserIdResult
	err    error
}

func NewVerifyInventoryCurrentMaxCapacityByUserIdResultFromJson(data string) VerifyInventoryCurrentMaxCapacityByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityByUserIdResultFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityByUserIdResultFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByUserIdResult {
	return VerifyInventoryCurrentMaxCapacityByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyInventoryCurrentMaxCapacityByUserIdResult) Pointer() *VerifyInventoryCurrentMaxCapacityByUserIdResult {
	return &p
}

type VerifyInventoryCurrentMaxCapacityByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult struct {
	result *VerifyInventoryCurrentMaxCapacityByStampTaskResult
	err    error
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskResultFromJson(data string) VerifyInventoryCurrentMaxCapacityByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyInventoryCurrentMaxCapacityByStampTaskResultFromDict(dict)
}

func NewVerifyInventoryCurrentMaxCapacityByStampTaskResultFromDict(data map[string]interface{}) VerifyInventoryCurrentMaxCapacityByStampTaskResult {
	return VerifyInventoryCurrentMaxCapacityByStampTaskResult{
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

func (p VerifyInventoryCurrentMaxCapacityByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyInventoryCurrentMaxCapacityByStampTaskResult) Pointer() *VerifyInventoryCurrentMaxCapacityByStampTaskResult {
	return &p
}

type AddCapacityByStampSheetResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AddCapacityByStampSheetResult) ToDict() map[string]interface{} {
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

func (p AddCapacityByStampSheetResult) Pointer() *AddCapacityByStampSheetResult {
	return &p
}

type SetCapacityByStampSheetResult struct {
	Item     *Inventory           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *Inventory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["item"])).Pointer()
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

func (p SetCapacityByStampSheetResult) ToDict() map[string]interface{} {
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

func (p SetCapacityByStampSheetResult) Pointer() *SetCapacityByStampSheetResult {
	return &p
}

type DescribeItemSetsResult struct {
	Items         []ItemSet            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
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

func (p DescribeItemSetsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
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

func (p DescribeItemSetsResult) Pointer() *DescribeItemSetsResult {
	return &p
}

type DescribeItemSetsByUserIdResult struct {
	Items         []ItemSet            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
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

func (p DescribeItemSetsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
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

func (p DescribeItemSetsByUserIdResult) Pointer() *DescribeItemSetsByUserIdResult {
	return &p
}

type GetItemSetResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetItemSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetItemSetResult) Pointer() *GetItemSetResult {
	return &p
}

type GetItemSetByUserIdResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetItemSetByUserIdResult) Pointer() *GetItemSetByUserIdResult {
	return &p
}

type GetItemWithSignatureResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetItemWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
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

func (p GetItemWithSignatureResult) Pointer() *GetItemWithSignatureResult {
	return &p
}

type GetItemWithSignatureByUserIdResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetItemWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
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

func (p GetItemWithSignatureByUserIdResult) Pointer() *GetItemWithSignatureByUserIdResult {
	return &p
}

type AcquireItemSetByUserIdResult struct {
	Items         []ItemSet            `json:"items"`
	ItemModel     *ItemModel           `json:"itemModel"`
	Inventory     *Inventory           `json:"inventory"`
	OverflowCount *int64               `json:"overflowCount"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
		}(),
		OverflowCount: func() *int64 {
			v, ok := data["overflowCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["overflowCount"])
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

func (p AcquireItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"overflowCount": p.OverflowCount,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireItemSetByUserIdResult) Pointer() *AcquireItemSetByUserIdResult {
	return &p
}

type AcquireItemSetWithGradeByUserIdResult struct {
	Item          *ItemSet             `json:"item"`
	Status        *grade.Status        `json:"status"`
	ItemModel     *ItemModel           `json:"itemModel"`
	Inventory     *Inventory           `json:"inventory"`
	OverflowCount *int64               `json:"overflowCount"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type AcquireItemSetWithGradeByUserIdAsyncResult struct {
	result *AcquireItemSetWithGradeByUserIdResult
	err    error
}

func NewAcquireItemSetWithGradeByUserIdResultFromJson(data string) AcquireItemSetWithGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetWithGradeByUserIdResultFromDict(dict)
}

func NewAcquireItemSetWithGradeByUserIdResultFromDict(data map[string]interface{}) AcquireItemSetWithGradeByUserIdResult {
	return AcquireItemSetWithGradeByUserIdResult{
		Item: func() *ItemSet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Status: func() *grade.Status {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return grade.NewStatusFromDict(core.CastMap(data["status"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
		}(),
		OverflowCount: func() *int64 {
			v, ok := data["overflowCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["overflowCount"])
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

func (p AcquireItemSetWithGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"overflowCount": p.OverflowCount,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireItemSetWithGradeByUserIdResult) Pointer() *AcquireItemSetWithGradeByUserIdResult {
	return &p
}

type ConsumeItemSetResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p ConsumeItemSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ConsumeItemSetResult) Pointer() *ConsumeItemSetResult {
	return &p
}

type ConsumeItemSetByUserIdResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p ConsumeItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ConsumeItemSetByUserIdResult) Pointer() *ConsumeItemSetByUserIdResult {
	return &p
}

type DeleteItemSetByUserIdResult struct {
	Items     []ItemSet            `json:"items"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DeleteItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteItemSetByUserIdResult) Pointer() *DeleteItemSetByUserIdResult {
	return &p
}

type VerifyItemSetResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyItemSetAsyncResult struct {
	result *VerifyItemSetResult
	err    error
}

func NewVerifyItemSetResultFromJson(data string) VerifyItemSetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetResultFromDict(dict)
}

func NewVerifyItemSetResultFromDict(data map[string]interface{}) VerifyItemSetResult {
	return VerifyItemSetResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyItemSetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyItemSetResult) Pointer() *VerifyItemSetResult {
	return &p
}

type VerifyItemSetByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyItemSetByUserIdAsyncResult struct {
	result *VerifyItemSetByUserIdResult
	err    error
}

func NewVerifyItemSetByUserIdResultFromJson(data string) VerifyItemSetByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetByUserIdResultFromDict(dict)
}

func NewVerifyItemSetByUserIdResultFromDict(data map[string]interface{}) VerifyItemSetByUserIdResult {
	return VerifyItemSetByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyItemSetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyItemSetByUserIdResult) Pointer() *VerifyItemSetByUserIdResult {
	return &p
}

type AcquireItemSetByStampSheetResult struct {
	Items         []ItemSet            `json:"items"`
	ItemModel     *ItemModel           `json:"itemModel"`
	Inventory     *Inventory           `json:"inventory"`
	OverflowCount *int64               `json:"overflowCount"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
		}(),
		OverflowCount: func() *int64 {
			v, ok := data["overflowCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["overflowCount"])
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

func (p AcquireItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"overflowCount": p.OverflowCount,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireItemSetByStampSheetResult) Pointer() *AcquireItemSetByStampSheetResult {
	return &p
}

type AcquireItemSetWithGradeByStampSheetResult struct {
	Item          *ItemSet             `json:"item"`
	Status        *grade.Status        `json:"status"`
	ItemModel     *ItemModel           `json:"itemModel"`
	Inventory     *Inventory           `json:"inventory"`
	OverflowCount *int64               `json:"overflowCount"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type AcquireItemSetWithGradeByStampSheetAsyncResult struct {
	result *AcquireItemSetWithGradeByStampSheetResult
	err    error
}

func NewAcquireItemSetWithGradeByStampSheetResultFromJson(data string) AcquireItemSetWithGradeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireItemSetWithGradeByStampSheetResultFromDict(dict)
}

func NewAcquireItemSetWithGradeByStampSheetResultFromDict(data map[string]interface{}) AcquireItemSetWithGradeByStampSheetResult {
	return AcquireItemSetWithGradeByStampSheetResult{
		Item: func() *ItemSet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Status: func() *grade.Status {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return grade.NewStatusFromDict(core.CastMap(data["status"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
		}(),
		OverflowCount: func() *int64 {
			v, ok := data["overflowCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["overflowCount"])
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

func (p AcquireItemSetWithGradeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"status": func() map[string]interface{} {
			if p.Status == nil {
				return nil
			}
			return p.Status.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"overflowCount": p.OverflowCount,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireItemSetWithGradeByStampSheetResult) Pointer() *AcquireItemSetWithGradeByStampSheetResult {
	return &p
}

type ConsumeItemSetByStampTaskResult struct {
	Items           []ItemSet            `json:"items"`
	ItemModel       *ItemModel           `json:"itemModel"`
	Inventory       *Inventory           `json:"inventory"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
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
		Items: func() []ItemSet {
			if data["items"] == nil {
				return nil
			}
			return CastItemSets(core.CastArray(data["items"]))
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p ConsumeItemSetByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastItemSetsFromDict(
			p.Items,
		),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
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

func (p ConsumeItemSetByStampTaskResult) Pointer() *ConsumeItemSetByStampTaskResult {
	return &p
}

type VerifyItemSetByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyItemSetByStampTaskAsyncResult struct {
	result *VerifyItemSetByStampTaskResult
	err    error
}

func NewVerifyItemSetByStampTaskResultFromJson(data string) VerifyItemSetByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyItemSetByStampTaskResultFromDict(dict)
}

func NewVerifyItemSetByStampTaskResultFromDict(data map[string]interface{}) VerifyItemSetByStampTaskResult {
	return VerifyItemSetByStampTaskResult{
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

func (p VerifyItemSetByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyItemSetByStampTaskResult) Pointer() *VerifyItemSetByStampTaskResult {
	return &p
}

type DescribeReferenceOfResult struct {
	Items     []*string            `json:"items"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DescribeReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
			p.Items,
		),
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeReferenceOfResult) Pointer() *DescribeReferenceOfResult {
	return &p
}

type DescribeReferenceOfByUserIdResult struct {
	Items     []*string            `json:"items"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DescribeReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
			p.Items,
		),
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DescribeReferenceOfByUserIdResult) Pointer() *DescribeReferenceOfByUserIdResult {
	return &p
}

type GetReferenceOfResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetReferenceOfResult) Pointer() *GetReferenceOfResult {
	return &p
}

type GetReferenceOfByUserIdResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p GetReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetReferenceOfByUserIdResult) Pointer() *GetReferenceOfByUserIdResult {
	return &p
}

type VerifyReferenceOfResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p VerifyReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyReferenceOfResult) Pointer() *VerifyReferenceOfResult {
	return &p
}

type VerifyReferenceOfByUserIdResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p VerifyReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyReferenceOfByUserIdResult) Pointer() *VerifyReferenceOfByUserIdResult {
	return &p
}

type AddReferenceOfResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p AddReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AddReferenceOfResult) Pointer() *AddReferenceOfResult {
	return &p
}

type AddReferenceOfByUserIdResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p AddReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AddReferenceOfByUserIdResult) Pointer() *AddReferenceOfByUserIdResult {
	return &p
}

type DeleteReferenceOfResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DeleteReferenceOfResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteReferenceOfResult) Pointer() *DeleteReferenceOfResult {
	return &p
}

type DeleteReferenceOfByUserIdResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DeleteReferenceOfByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteReferenceOfByUserIdResult) Pointer() *DeleteReferenceOfByUserIdResult {
	return &p
}

type AddReferenceOfItemSetByStampSheetResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p AddReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AddReferenceOfItemSetByStampSheetResult) Pointer() *AddReferenceOfItemSetByStampSheetResult {
	return &p
}

type DeleteReferenceOfItemSetByStampSheetResult struct {
	Item      *string              `json:"item"`
	ItemSet   *ItemSet             `json:"itemSet"`
	ItemModel *ItemModel           `json:"itemModel"`
	Inventory *Inventory           `json:"inventory"`
	Metadata  *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p DeleteReferenceOfItemSetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteReferenceOfItemSetByStampSheetResult) Pointer() *DeleteReferenceOfItemSetByStampSheetResult {
	return &p
}

type VerifyReferenceOfByStampTaskResult struct {
	Item            *string              `json:"item"`
	ItemSet         *ItemSet             `json:"itemSet"`
	ItemModel       *ItemModel           `json:"itemModel"`
	Inventory       *Inventory           `json:"inventory"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
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
		Item: func() *string {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["item"])
		}(),
		ItemSet: func() *ItemSet {
			v, ok := data["itemSet"]
			if !ok || v == nil {
				return nil
			}
			return NewItemSetFromDict(core.CastMap(data["itemSet"])).Pointer()
		}(),
		ItemModel: func() *ItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
		}(),
		Inventory: func() *Inventory {
			v, ok := data["inventory"]
			if !ok || v == nil {
				return nil
			}
			return NewInventoryFromDict(core.CastMap(data["inventory"])).Pointer()
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

func (p VerifyReferenceOfByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item,
		"itemSet": func() map[string]interface{} {
			if p.ItemSet == nil {
				return nil
			}
			return p.ItemSet.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"inventory": func() map[string]interface{} {
			if p.Inventory == nil {
				return nil
			}
			return p.Inventory.ToDict()
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

func (p VerifyReferenceOfByStampTaskResult) Pointer() *VerifyReferenceOfByStampTaskResult {
	return &p
}

type DescribeSimpleItemsResult struct {
	Items         []SimpleItem         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSimpleItemsAsyncResult struct {
	result *DescribeSimpleItemsResult
	err    error
}

func NewDescribeSimpleItemsResultFromJson(data string) DescribeSimpleItemsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemsResultFromDict(dict)
}

func NewDescribeSimpleItemsResultFromDict(data map[string]interface{}) DescribeSimpleItemsResult {
	return DescribeSimpleItemsResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p DescribeSimpleItemsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p DescribeSimpleItemsResult) Pointer() *DescribeSimpleItemsResult {
	return &p
}

type DescribeSimpleItemsByUserIdResult struct {
	Items         []SimpleItem         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSimpleItemsByUserIdAsyncResult struct {
	result *DescribeSimpleItemsByUserIdResult
	err    error
}

func NewDescribeSimpleItemsByUserIdResultFromJson(data string) DescribeSimpleItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSimpleItemsByUserIdResultFromDict(dict)
}

func NewDescribeSimpleItemsByUserIdResultFromDict(data map[string]interface{}) DescribeSimpleItemsByUserIdResult {
	return DescribeSimpleItemsByUserIdResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p DescribeSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p DescribeSimpleItemsByUserIdResult) Pointer() *DescribeSimpleItemsByUserIdResult {
	return &p
}

type GetSimpleItemResult struct {
	Item      *SimpleItem          `json:"item"`
	ItemModel *SimpleItemModel     `json:"itemModel"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetSimpleItemAsyncResult struct {
	result *GetSimpleItemResult
	err    error
}

func NewGetSimpleItemResultFromJson(data string) GetSimpleItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemResultFromDict(dict)
}

func NewGetSimpleItemResultFromDict(data map[string]interface{}) GetSimpleItemResult {
	return GetSimpleItemResult{
		Item: func() *SimpleItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ItemModel: func() *SimpleItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
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

func (p GetSimpleItemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetSimpleItemResult) Pointer() *GetSimpleItemResult {
	return &p
}

type GetSimpleItemByUserIdResult struct {
	Item      *SimpleItem          `json:"item"`
	ItemModel *SimpleItemModel     `json:"itemModel"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetSimpleItemByUserIdAsyncResult struct {
	result *GetSimpleItemByUserIdResult
	err    error
}

func NewGetSimpleItemByUserIdResultFromJson(data string) GetSimpleItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemByUserIdResultFromDict(dict)
}

func NewGetSimpleItemByUserIdResultFromDict(data map[string]interface{}) GetSimpleItemByUserIdResult {
	return GetSimpleItemByUserIdResult{
		Item: func() *SimpleItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ItemModel: func() *SimpleItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
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

func (p GetSimpleItemByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetSimpleItemByUserIdResult) Pointer() *GetSimpleItemByUserIdResult {
	return &p
}

type GetSimpleItemWithSignatureResult struct {
	Item            *SimpleItem          `json:"item"`
	SimpleItemModel *SimpleItemModel     `json:"simpleItemModel"`
	Body            *string              `json:"body"`
	Signature       *string              `json:"signature"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type GetSimpleItemWithSignatureAsyncResult struct {
	result *GetSimpleItemWithSignatureResult
	err    error
}

func NewGetSimpleItemWithSignatureResultFromJson(data string) GetSimpleItemWithSignatureResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemWithSignatureResultFromDict(dict)
}

func NewGetSimpleItemWithSignatureResultFromDict(data map[string]interface{}) GetSimpleItemWithSignatureResult {
	return GetSimpleItemWithSignatureResult{
		Item: func() *SimpleItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		SimpleItemModel: func() *SimpleItemModel {
			v, ok := data["simpleItemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelFromDict(core.CastMap(data["simpleItemModel"])).Pointer()
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

func (p GetSimpleItemWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"simpleItemModel": func() map[string]interface{} {
			if p.SimpleItemModel == nil {
				return nil
			}
			return p.SimpleItemModel.ToDict()
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

func (p GetSimpleItemWithSignatureResult) Pointer() *GetSimpleItemWithSignatureResult {
	return &p
}

type GetSimpleItemWithSignatureByUserIdResult struct {
	Item            *SimpleItem          `json:"item"`
	SimpleItemModel *SimpleItemModel     `json:"simpleItemModel"`
	Body            *string              `json:"body"`
	Signature       *string              `json:"signature"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type GetSimpleItemWithSignatureByUserIdAsyncResult struct {
	result *GetSimpleItemWithSignatureByUserIdResult
	err    error
}

func NewGetSimpleItemWithSignatureByUserIdResultFromJson(data string) GetSimpleItemWithSignatureByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSimpleItemWithSignatureByUserIdResultFromDict(dict)
}

func NewGetSimpleItemWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetSimpleItemWithSignatureByUserIdResult {
	return GetSimpleItemWithSignatureByUserIdResult{
		Item: func() *SimpleItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		SimpleItemModel: func() *SimpleItemModel {
			v, ok := data["simpleItemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewSimpleItemModelFromDict(core.CastMap(data["simpleItemModel"])).Pointer()
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

func (p GetSimpleItemWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"simpleItemModel": func() map[string]interface{} {
			if p.SimpleItemModel == nil {
				return nil
			}
			return p.SimpleItemModel.ToDict()
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

func (p GetSimpleItemWithSignatureByUserIdResult) Pointer() *GetSimpleItemWithSignatureByUserIdResult {
	return &p
}

type AcquireSimpleItemsByUserIdResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcquireSimpleItemsByUserIdAsyncResult struct {
	result *AcquireSimpleItemsByUserIdResult
	err    error
}

func NewAcquireSimpleItemsByUserIdResultFromJson(data string) AcquireSimpleItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireSimpleItemsByUserIdResultFromDict(dict)
}

func NewAcquireSimpleItemsByUserIdResultFromDict(data map[string]interface{}) AcquireSimpleItemsByUserIdResult {
	return AcquireSimpleItemsByUserIdResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p AcquireSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p AcquireSimpleItemsByUserIdResult) Pointer() *AcquireSimpleItemsByUserIdResult {
	return &p
}

type ConsumeSimpleItemsResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ConsumeSimpleItemsAsyncResult struct {
	result *ConsumeSimpleItemsResult
	err    error
}

func NewConsumeSimpleItemsResultFromJson(data string) ConsumeSimpleItemsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsResultFromDict(dict)
}

func NewConsumeSimpleItemsResultFromDict(data map[string]interface{}) ConsumeSimpleItemsResult {
	return ConsumeSimpleItemsResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p ConsumeSimpleItemsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p ConsumeSimpleItemsResult) Pointer() *ConsumeSimpleItemsResult {
	return &p
}

type ConsumeSimpleItemsByUserIdResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ConsumeSimpleItemsByUserIdAsyncResult struct {
	result *ConsumeSimpleItemsByUserIdResult
	err    error
}

func NewConsumeSimpleItemsByUserIdResultFromJson(data string) ConsumeSimpleItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsByUserIdResultFromDict(dict)
}

func NewConsumeSimpleItemsByUserIdResultFromDict(data map[string]interface{}) ConsumeSimpleItemsByUserIdResult {
	return ConsumeSimpleItemsByUserIdResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p ConsumeSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p ConsumeSimpleItemsByUserIdResult) Pointer() *ConsumeSimpleItemsByUserIdResult {
	return &p
}

type SetSimpleItemsByUserIdResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetSimpleItemsByUserIdAsyncResult struct {
	result *SetSimpleItemsByUserIdResult
	err    error
}

func NewSetSimpleItemsByUserIdResultFromJson(data string) SetSimpleItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetSimpleItemsByUserIdResultFromDict(dict)
}

func NewSetSimpleItemsByUserIdResultFromDict(data map[string]interface{}) SetSimpleItemsByUserIdResult {
	return SetSimpleItemsByUserIdResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p SetSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p SetSimpleItemsByUserIdResult) Pointer() *SetSimpleItemsByUserIdResult {
	return &p
}

type DeleteSimpleItemsByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSimpleItemsByUserIdAsyncResult struct {
	result *DeleteSimpleItemsByUserIdResult
	err    error
}

func NewDeleteSimpleItemsByUserIdResultFromJson(data string) DeleteSimpleItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSimpleItemsByUserIdResultFromDict(dict)
}

func NewDeleteSimpleItemsByUserIdResultFromDict(data map[string]interface{}) DeleteSimpleItemsByUserIdResult {
	return DeleteSimpleItemsByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteSimpleItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteSimpleItemsByUserIdResult) Pointer() *DeleteSimpleItemsByUserIdResult {
	return &p
}

type VerifySimpleItemResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifySimpleItemAsyncResult struct {
	result *VerifySimpleItemResult
	err    error
}

func NewVerifySimpleItemResultFromJson(data string) VerifySimpleItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemResultFromDict(dict)
}

func NewVerifySimpleItemResultFromDict(data map[string]interface{}) VerifySimpleItemResult {
	return VerifySimpleItemResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifySimpleItemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifySimpleItemResult) Pointer() *VerifySimpleItemResult {
	return &p
}

type VerifySimpleItemByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifySimpleItemByUserIdAsyncResult struct {
	result *VerifySimpleItemByUserIdResult
	err    error
}

func NewVerifySimpleItemByUserIdResultFromJson(data string) VerifySimpleItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemByUserIdResultFromDict(dict)
}

func NewVerifySimpleItemByUserIdResultFromDict(data map[string]interface{}) VerifySimpleItemByUserIdResult {
	return VerifySimpleItemByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifySimpleItemByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifySimpleItemByUserIdResult) Pointer() *VerifySimpleItemByUserIdResult {
	return &p
}

type AcquireSimpleItemsByStampSheetResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcquireSimpleItemsByStampSheetAsyncResult struct {
	result *AcquireSimpleItemsByStampSheetResult
	err    error
}

func NewAcquireSimpleItemsByStampSheetResultFromJson(data string) AcquireSimpleItemsByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireSimpleItemsByStampSheetResultFromDict(dict)
}

func NewAcquireSimpleItemsByStampSheetResultFromDict(data map[string]interface{}) AcquireSimpleItemsByStampSheetResult {
	return AcquireSimpleItemsByStampSheetResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p AcquireSimpleItemsByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p AcquireSimpleItemsByStampSheetResult) Pointer() *AcquireSimpleItemsByStampSheetResult {
	return &p
}

type ConsumeSimpleItemsByStampTaskResult struct {
	Items           []SimpleItem         `json:"items"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type ConsumeSimpleItemsByStampTaskAsyncResult struct {
	result *ConsumeSimpleItemsByStampTaskResult
	err    error
}

func NewConsumeSimpleItemsByStampTaskResultFromJson(data string) ConsumeSimpleItemsByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeSimpleItemsByStampTaskResultFromDict(dict)
}

func NewConsumeSimpleItemsByStampTaskResultFromDict(data map[string]interface{}) ConsumeSimpleItemsByStampTaskResult {
	return ConsumeSimpleItemsByStampTaskResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p ConsumeSimpleItemsByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p ConsumeSimpleItemsByStampTaskResult) Pointer() *ConsumeSimpleItemsByStampTaskResult {
	return &p
}

type SetSimpleItemsByStampSheetResult struct {
	Items    []SimpleItem         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetSimpleItemsByStampSheetAsyncResult struct {
	result *SetSimpleItemsByStampSheetResult
	err    error
}

func NewSetSimpleItemsByStampSheetResultFromJson(data string) SetSimpleItemsByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetSimpleItemsByStampSheetResultFromDict(dict)
}

func NewSetSimpleItemsByStampSheetResultFromDict(data map[string]interface{}) SetSimpleItemsByStampSheetResult {
	return SetSimpleItemsByStampSheetResult{
		Items: func() []SimpleItem {
			if data["items"] == nil {
				return nil
			}
			return CastSimpleItems(core.CastArray(data["items"]))
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

func (p SetSimpleItemsByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSimpleItemsFromDict(
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

func (p SetSimpleItemsByStampSheetResult) Pointer() *SetSimpleItemsByStampSheetResult {
	return &p
}

type VerifySimpleItemByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifySimpleItemByStampTaskAsyncResult struct {
	result *VerifySimpleItemByStampTaskResult
	err    error
}

func NewVerifySimpleItemByStampTaskResultFromJson(data string) VerifySimpleItemByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySimpleItemByStampTaskResultFromDict(dict)
}

func NewVerifySimpleItemByStampTaskResultFromDict(data map[string]interface{}) VerifySimpleItemByStampTaskResult {
	return VerifySimpleItemByStampTaskResult{
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

func (p VerifySimpleItemByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifySimpleItemByStampTaskResult) Pointer() *VerifySimpleItemByStampTaskResult {
	return &p
}

type DescribeBigItemsResult struct {
	Items         []BigItem            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBigItemsAsyncResult struct {
	result *DescribeBigItemsResult
	err    error
}

func NewDescribeBigItemsResultFromJson(data string) DescribeBigItemsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemsResultFromDict(dict)
}

func NewDescribeBigItemsResultFromDict(data map[string]interface{}) DescribeBigItemsResult {
	return DescribeBigItemsResult{
		Items: func() []BigItem {
			if data["items"] == nil {
				return nil
			}
			return CastBigItems(core.CastArray(data["items"]))
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

func (p DescribeBigItemsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigItemsFromDict(
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

func (p DescribeBigItemsResult) Pointer() *DescribeBigItemsResult {
	return &p
}

type DescribeBigItemsByUserIdResult struct {
	Items         []BigItem            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBigItemsByUserIdAsyncResult struct {
	result *DescribeBigItemsByUserIdResult
	err    error
}

func NewDescribeBigItemsByUserIdResultFromJson(data string) DescribeBigItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBigItemsByUserIdResultFromDict(dict)
}

func NewDescribeBigItemsByUserIdResultFromDict(data map[string]interface{}) DescribeBigItemsByUserIdResult {
	return DescribeBigItemsByUserIdResult{
		Items: func() []BigItem {
			if data["items"] == nil {
				return nil
			}
			return CastBigItems(core.CastArray(data["items"]))
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

func (p DescribeBigItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBigItemsFromDict(
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

func (p DescribeBigItemsByUserIdResult) Pointer() *DescribeBigItemsByUserIdResult {
	return &p
}

type GetBigItemResult struct {
	Item      *BigItem             `json:"item"`
	ItemModel *BigItemModel        `json:"itemModel"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetBigItemAsyncResult struct {
	result *GetBigItemResult
	err    error
}

func NewGetBigItemResultFromJson(data string) GetBigItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemResultFromDict(dict)
}

func NewGetBigItemResultFromDict(data map[string]interface{}) GetBigItemResult {
	return GetBigItemResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ItemModel: func() *BigItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
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

func (p GetBigItemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBigItemResult) Pointer() *GetBigItemResult {
	return &p
}

type GetBigItemByUserIdResult struct {
	Item      *BigItem             `json:"item"`
	ItemModel *BigItemModel        `json:"itemModel"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetBigItemByUserIdAsyncResult struct {
	result *GetBigItemByUserIdResult
	err    error
}

func NewGetBigItemByUserIdResultFromJson(data string) GetBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBigItemByUserIdResultFromDict(dict)
}

func NewGetBigItemByUserIdResultFromDict(data map[string]interface{}) GetBigItemByUserIdResult {
	return GetBigItemByUserIdResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ItemModel: func() *BigItemModel {
			v, ok := data["itemModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemModelFromDict(core.CastMap(data["itemModel"])).Pointer()
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

func (p GetBigItemByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"itemModel": func() map[string]interface{} {
			if p.ItemModel == nil {
				return nil
			}
			return p.ItemModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBigItemByUserIdResult) Pointer() *GetBigItemByUserIdResult {
	return &p
}

type AcquireBigItemByUserIdResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcquireBigItemByUserIdAsyncResult struct {
	result *AcquireBigItemByUserIdResult
	err    error
}

func NewAcquireBigItemByUserIdResultFromJson(data string) AcquireBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireBigItemByUserIdResultFromDict(dict)
}

func NewAcquireBigItemByUserIdResultFromDict(data map[string]interface{}) AcquireBigItemByUserIdResult {
	return AcquireBigItemByUserIdResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AcquireBigItemByUserIdResult) ToDict() map[string]interface{} {
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

func (p AcquireBigItemByUserIdResult) Pointer() *AcquireBigItemByUserIdResult {
	return &p
}

type ConsumeBigItemResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ConsumeBigItemAsyncResult struct {
	result *ConsumeBigItemResult
	err    error
}

func NewConsumeBigItemResultFromJson(data string) ConsumeBigItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemResultFromDict(dict)
}

func NewConsumeBigItemResultFromDict(data map[string]interface{}) ConsumeBigItemResult {
	return ConsumeBigItemResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p ConsumeBigItemResult) ToDict() map[string]interface{} {
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

func (p ConsumeBigItemResult) Pointer() *ConsumeBigItemResult {
	return &p
}

type ConsumeBigItemByUserIdResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ConsumeBigItemByUserIdAsyncResult struct {
	result *ConsumeBigItemByUserIdResult
	err    error
}

func NewConsumeBigItemByUserIdResultFromJson(data string) ConsumeBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemByUserIdResultFromDict(dict)
}

func NewConsumeBigItemByUserIdResultFromDict(data map[string]interface{}) ConsumeBigItemByUserIdResult {
	return ConsumeBigItemByUserIdResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p ConsumeBigItemByUserIdResult) ToDict() map[string]interface{} {
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

func (p ConsumeBigItemByUserIdResult) Pointer() *ConsumeBigItemByUserIdResult {
	return &p
}

type SetBigItemByUserIdResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetBigItemByUserIdAsyncResult struct {
	result *SetBigItemByUserIdResult
	err    error
}

func NewSetBigItemByUserIdResultFromJson(data string) SetBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBigItemByUserIdResultFromDict(dict)
}

func NewSetBigItemByUserIdResultFromDict(data map[string]interface{}) SetBigItemByUserIdResult {
	return SetBigItemByUserIdResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p SetBigItemByUserIdResult) ToDict() map[string]interface{} {
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

func (p SetBigItemByUserIdResult) Pointer() *SetBigItemByUserIdResult {
	return &p
}

type DeleteBigItemByUserIdResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteBigItemByUserIdAsyncResult struct {
	result *DeleteBigItemByUserIdResult
	err    error
}

func NewDeleteBigItemByUserIdResultFromJson(data string) DeleteBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBigItemByUserIdResultFromDict(dict)
}

func NewDeleteBigItemByUserIdResultFromDict(data map[string]interface{}) DeleteBigItemByUserIdResult {
	return DeleteBigItemByUserIdResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteBigItemByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteBigItemByUserIdResult) Pointer() *DeleteBigItemByUserIdResult {
	return &p
}

type VerifyBigItemResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyBigItemAsyncResult struct {
	result *VerifyBigItemResult
	err    error
}

func NewVerifyBigItemResultFromJson(data string) VerifyBigItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemResultFromDict(dict)
}

func NewVerifyBigItemResultFromDict(data map[string]interface{}) VerifyBigItemResult {
	return VerifyBigItemResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyBigItemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyBigItemResult) Pointer() *VerifyBigItemResult {
	return &p
}

type VerifyBigItemByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyBigItemByUserIdAsyncResult struct {
	result *VerifyBigItemByUserIdResult
	err    error
}

func NewVerifyBigItemByUserIdResultFromJson(data string) VerifyBigItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemByUserIdResultFromDict(dict)
}

func NewVerifyBigItemByUserIdResultFromDict(data map[string]interface{}) VerifyBigItemByUserIdResult {
	return VerifyBigItemByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyBigItemByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyBigItemByUserIdResult) Pointer() *VerifyBigItemByUserIdResult {
	return &p
}

type AcquireBigItemByStampSheetResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcquireBigItemByStampSheetAsyncResult struct {
	result *AcquireBigItemByStampSheetResult
	err    error
}

func NewAcquireBigItemByStampSheetResultFromJson(data string) AcquireBigItemByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireBigItemByStampSheetResultFromDict(dict)
}

func NewAcquireBigItemByStampSheetResultFromDict(data map[string]interface{}) AcquireBigItemByStampSheetResult {
	return AcquireBigItemByStampSheetResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p AcquireBigItemByStampSheetResult) ToDict() map[string]interface{} {
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

func (p AcquireBigItemByStampSheetResult) Pointer() *AcquireBigItemByStampSheetResult {
	return &p
}

type ConsumeBigItemByStampTaskResult struct {
	Item            *BigItem             `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type ConsumeBigItemByStampTaskAsyncResult struct {
	result *ConsumeBigItemByStampTaskResult
	err    error
}

func NewConsumeBigItemByStampTaskResultFromJson(data string) ConsumeBigItemByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeBigItemByStampTaskResultFromDict(dict)
}

func NewConsumeBigItemByStampTaskResultFromDict(data map[string]interface{}) ConsumeBigItemByStampTaskResult {
	return ConsumeBigItemByStampTaskResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p ConsumeBigItemByStampTaskResult) ToDict() map[string]interface{} {
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

func (p ConsumeBigItemByStampTaskResult) Pointer() *ConsumeBigItemByStampTaskResult {
	return &p
}

type SetBigItemByStampSheetResult struct {
	Item     *BigItem             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SetBigItemByStampSheetAsyncResult struct {
	result *SetBigItemByStampSheetResult
	err    error
}

func NewSetBigItemByStampSheetResultFromJson(data string) SetBigItemByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBigItemByStampSheetResultFromDict(dict)
}

func NewSetBigItemByStampSheetResultFromDict(data map[string]interface{}) SetBigItemByStampSheetResult {
	return SetBigItemByStampSheetResult{
		Item: func() *BigItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBigItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p SetBigItemByStampSheetResult) ToDict() map[string]interface{} {
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

func (p SetBigItemByStampSheetResult) Pointer() *SetBigItemByStampSheetResult {
	return &p
}

type VerifyBigItemByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyBigItemByStampTaskAsyncResult struct {
	result *VerifyBigItemByStampTaskResult
	err    error
}

func NewVerifyBigItemByStampTaskResultFromJson(data string) VerifyBigItemByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyBigItemByStampTaskResultFromDict(dict)
}

func NewVerifyBigItemByStampTaskResultFromDict(data map[string]interface{}) VerifyBigItemByStampTaskResult {
	return VerifyBigItemByStampTaskResult{
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

func (p VerifyBigItemByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyBigItemByStampTaskResult) Pointer() *VerifyBigItemByStampTaskResult {
	return &p
}
