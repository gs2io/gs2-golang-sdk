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

package datastore

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

type DescribeDataObjectsResult struct {
	Items         []DataObject         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeDataObjectsAsyncResult struct {
	result *DescribeDataObjectsResult
	err    error
}

func NewDescribeDataObjectsResultFromJson(data string) DescribeDataObjectsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDataObjectsResultFromDict(dict)
}

func NewDescribeDataObjectsResultFromDict(data map[string]interface{}) DescribeDataObjectsResult {
	return DescribeDataObjectsResult{
		Items: func() []DataObject {
			if data["items"] == nil {
				return nil
			}
			return CastDataObjects(core.CastArray(data["items"]))
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

func (p DescribeDataObjectsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectsFromDict(
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

func (p DescribeDataObjectsResult) Pointer() *DescribeDataObjectsResult {
	return &p
}

type DescribeDataObjectsByUserIdResult struct {
	Items         []DataObject         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeDataObjectsByUserIdAsyncResult struct {
	result *DescribeDataObjectsByUserIdResult
	err    error
}

func NewDescribeDataObjectsByUserIdResultFromJson(data string) DescribeDataObjectsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDataObjectsByUserIdResultFromDict(dict)
}

func NewDescribeDataObjectsByUserIdResultFromDict(data map[string]interface{}) DescribeDataObjectsByUserIdResult {
	return DescribeDataObjectsByUserIdResult{
		Items: func() []DataObject {
			if data["items"] == nil {
				return nil
			}
			return CastDataObjects(core.CastArray(data["items"]))
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

func (p DescribeDataObjectsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectsFromDict(
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

func (p DescribeDataObjectsByUserIdResult) Pointer() *DescribeDataObjectsByUserIdResult {
	return &p
}

type PrepareUploadResult struct {
	Item      *DataObject          `json:"item"`
	UploadUrl *string              `json:"uploadUrl"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type PrepareUploadAsyncResult struct {
	result *PrepareUploadResult
	err    error
}

func NewPrepareUploadResultFromJson(data string) PrepareUploadResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareUploadResultFromDict(dict)
}

func NewPrepareUploadResultFromDict(data map[string]interface{}) PrepareUploadResult {
	return PrepareUploadResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PrepareUploadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"uploadUrl": p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareUploadResult) Pointer() *PrepareUploadResult {
	return &p
}

type PrepareUploadByUserIdResult struct {
	Item      *DataObject          `json:"item"`
	UploadUrl *string              `json:"uploadUrl"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type PrepareUploadByUserIdAsyncResult struct {
	result *PrepareUploadByUserIdResult
	err    error
}

func NewPrepareUploadByUserIdResultFromJson(data string) PrepareUploadByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareUploadByUserIdResultFromDict(dict)
}

func NewPrepareUploadByUserIdResultFromDict(data map[string]interface{}) PrepareUploadByUserIdResult {
	return PrepareUploadByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PrepareUploadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"uploadUrl": p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareUploadByUserIdResult) Pointer() *PrepareUploadByUserIdResult {
	return &p
}

type UpdateDataObjectResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateDataObjectAsyncResult struct {
	result *UpdateDataObjectResult
	err    error
}

func NewUpdateDataObjectResultFromJson(data string) UpdateDataObjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateDataObjectResultFromDict(dict)
}

func NewUpdateDataObjectResultFromDict(data map[string]interface{}) UpdateDataObjectResult {
	return UpdateDataObjectResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateDataObjectResult) ToDict() map[string]interface{} {
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

func (p UpdateDataObjectResult) Pointer() *UpdateDataObjectResult {
	return &p
}

type UpdateDataObjectByUserIdResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateDataObjectByUserIdAsyncResult struct {
	result *UpdateDataObjectByUserIdResult
	err    error
}

func NewUpdateDataObjectByUserIdResultFromJson(data string) UpdateDataObjectByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateDataObjectByUserIdResultFromDict(dict)
}

func NewUpdateDataObjectByUserIdResultFromDict(data map[string]interface{}) UpdateDataObjectByUserIdResult {
	return UpdateDataObjectByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateDataObjectByUserIdResult) ToDict() map[string]interface{} {
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

func (p UpdateDataObjectByUserIdResult) Pointer() *UpdateDataObjectByUserIdResult {
	return &p
}

type PrepareReUploadResult struct {
	Item      *DataObject          `json:"item"`
	UploadUrl *string              `json:"uploadUrl"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type PrepareReUploadAsyncResult struct {
	result *PrepareReUploadResult
	err    error
}

func NewPrepareReUploadResultFromJson(data string) PrepareReUploadResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareReUploadResultFromDict(dict)
}

func NewPrepareReUploadResultFromDict(data map[string]interface{}) PrepareReUploadResult {
	return PrepareReUploadResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PrepareReUploadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"uploadUrl": p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareReUploadResult) Pointer() *PrepareReUploadResult {
	return &p
}

type PrepareReUploadByUserIdResult struct {
	Item      *DataObject          `json:"item"`
	UploadUrl *string              `json:"uploadUrl"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type PrepareReUploadByUserIdAsyncResult struct {
	result *PrepareReUploadByUserIdResult
	err    error
}

func NewPrepareReUploadByUserIdResultFromJson(data string) PrepareReUploadByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareReUploadByUserIdResultFromDict(dict)
}

func NewPrepareReUploadByUserIdResultFromDict(data map[string]interface{}) PrepareReUploadByUserIdResult {
	return PrepareReUploadByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
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

func (p PrepareReUploadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"uploadUrl": p.UploadUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareReUploadByUserIdResult) Pointer() *PrepareReUploadByUserIdResult {
	return &p
}

type DoneUploadResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DoneUploadAsyncResult struct {
	result *DoneUploadResult
	err    error
}

func NewDoneUploadResultFromJson(data string) DoneUploadResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoneUploadResultFromDict(dict)
}

func NewDoneUploadResultFromDict(data map[string]interface{}) DoneUploadResult {
	return DoneUploadResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoneUploadResult) ToDict() map[string]interface{} {
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

func (p DoneUploadResult) Pointer() *DoneUploadResult {
	return &p
}

type DoneUploadByUserIdResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DoneUploadByUserIdAsyncResult struct {
	result *DoneUploadByUserIdResult
	err    error
}

func NewDoneUploadByUserIdResultFromJson(data string) DoneUploadByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoneUploadByUserIdResultFromDict(dict)
}

func NewDoneUploadByUserIdResultFromDict(data map[string]interface{}) DoneUploadByUserIdResult {
	return DoneUploadByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoneUploadByUserIdResult) ToDict() map[string]interface{} {
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

func (p DoneUploadByUserIdResult) Pointer() *DoneUploadByUserIdResult {
	return &p
}

type DeleteDataObjectResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteDataObjectAsyncResult struct {
	result *DeleteDataObjectResult
	err    error
}

func NewDeleteDataObjectResultFromJson(data string) DeleteDataObjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDataObjectResultFromDict(dict)
}

func NewDeleteDataObjectResultFromDict(data map[string]interface{}) DeleteDataObjectResult {
	return DeleteDataObjectResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteDataObjectResult) ToDict() map[string]interface{} {
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

func (p DeleteDataObjectResult) Pointer() *DeleteDataObjectResult {
	return &p
}

type DeleteDataObjectByUserIdResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteDataObjectByUserIdAsyncResult struct {
	result *DeleteDataObjectByUserIdResult
	err    error
}

func NewDeleteDataObjectByUserIdResultFromJson(data string) DeleteDataObjectByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDataObjectByUserIdResultFromDict(dict)
}

func NewDeleteDataObjectByUserIdResultFromDict(data map[string]interface{}) DeleteDataObjectByUserIdResult {
	return DeleteDataObjectByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteDataObjectByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteDataObjectByUserIdResult) Pointer() *DeleteDataObjectByUserIdResult {
	return &p
}

type PrepareDownloadResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadAsyncResult struct {
	result *PrepareDownloadResult
	err    error
}

func NewPrepareDownloadResultFromJson(data string) PrepareDownloadResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadResultFromDict(dict)
}

func NewPrepareDownloadResultFromDict(data map[string]interface{}) PrepareDownloadResult {
	return PrepareDownloadResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadResult) Pointer() *PrepareDownloadResult {
	return &p
}

type PrepareDownloadByUserIdResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadByUserIdAsyncResult struct {
	result *PrepareDownloadByUserIdResult
	err    error
}

func NewPrepareDownloadByUserIdResultFromJson(data string) PrepareDownloadByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadByUserIdResultFromDict(dict)
}

func NewPrepareDownloadByUserIdResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdResult {
	return PrepareDownloadByUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadByUserIdResult) Pointer() *PrepareDownloadByUserIdResult {
	return &p
}

type PrepareDownloadByGenerationResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadByGenerationAsyncResult struct {
	result *PrepareDownloadByGenerationResult
	err    error
}

func NewPrepareDownloadByGenerationResultFromJson(data string) PrepareDownloadByGenerationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadByGenerationResultFromDict(dict)
}

func NewPrepareDownloadByGenerationResultFromDict(data map[string]interface{}) PrepareDownloadByGenerationResult {
	return PrepareDownloadByGenerationResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadByGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadByGenerationResult) Pointer() *PrepareDownloadByGenerationResult {
	return &p
}

type PrepareDownloadByGenerationAndUserIdResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadByGenerationAndUserIdAsyncResult struct {
	result *PrepareDownloadByGenerationAndUserIdResult
	err    error
}

func NewPrepareDownloadByGenerationAndUserIdResultFromJson(data string) PrepareDownloadByGenerationAndUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadByGenerationAndUserIdResultFromDict(dict)
}

func NewPrepareDownloadByGenerationAndUserIdResultFromDict(data map[string]interface{}) PrepareDownloadByGenerationAndUserIdResult {
	return PrepareDownloadByGenerationAndUserIdResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadByGenerationAndUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadByGenerationAndUserIdResult) Pointer() *PrepareDownloadByGenerationAndUserIdResult {
	return &p
}

type PrepareDownloadOwnDataResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadOwnDataAsyncResult struct {
	result *PrepareDownloadOwnDataResult
	err    error
}

func NewPrepareDownloadOwnDataResultFromJson(data string) PrepareDownloadOwnDataResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadOwnDataResultFromDict(dict)
}

func NewPrepareDownloadOwnDataResultFromDict(data map[string]interface{}) PrepareDownloadOwnDataResult {
	return PrepareDownloadOwnDataResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadOwnDataResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadOwnDataResult) Pointer() *PrepareDownloadOwnDataResult {
	return &p
}

type PrepareDownloadByUserIdAndDataObjectNameResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadByUserIdAndDataObjectNameAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameResult
	err    error
}

func NewPrepareDownloadByUserIdAndDataObjectNameResultFromJson(data string) PrepareDownloadByUserIdAndDataObjectNameResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadByUserIdAndDataObjectNameResultFromDict(dict)
}

func NewPrepareDownloadByUserIdAndDataObjectNameResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameResult {
	return PrepareDownloadByUserIdAndDataObjectNameResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameResult) Pointer() *PrepareDownloadByUserIdAndDataObjectNameResult {
	return &p
}

type PrepareDownloadOwnDataByGenerationResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadOwnDataByGenerationAsyncResult struct {
	result *PrepareDownloadOwnDataByGenerationResult
	err    error
}

func NewPrepareDownloadOwnDataByGenerationResultFromJson(data string) PrepareDownloadOwnDataByGenerationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadOwnDataByGenerationResultFromDict(dict)
}

func NewPrepareDownloadOwnDataByGenerationResultFromDict(data map[string]interface{}) PrepareDownloadOwnDataByGenerationResult {
	return PrepareDownloadOwnDataByGenerationResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadOwnDataByGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadOwnDataByGenerationResult) Pointer() *PrepareDownloadOwnDataByGenerationResult {
	return &p
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult struct {
	Item          *DataObject          `json:"item"`
	FileUrl       *string              `json:"fileUrl"`
	ContentLength *int64               `json:"contentLength"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult struct {
	result *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult
	err    error
}

func NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationResultFromJson(data string) PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationResultFromDict(dict)
}

func NewPrepareDownloadByUserIdAndDataObjectNameAndGenerationResultFromDict(data map[string]interface{}) PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult {
	return PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		FileUrl: func() *string {
			v, ok := data["fileUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fileUrl"])
		}(),
		ContentLength: func() *int64 {
			v, ok := data["contentLength"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["contentLength"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"fileUrl":       p.FileUrl,
		"contentLength": p.ContentLength,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult) Pointer() *PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult {
	return &p
}

type RestoreDataObjectResult struct {
	Item     *DataObject          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RestoreDataObjectAsyncResult struct {
	result *RestoreDataObjectResult
	err    error
}

func NewRestoreDataObjectResultFromJson(data string) RestoreDataObjectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRestoreDataObjectResultFromDict(dict)
}

func NewRestoreDataObjectResultFromDict(data map[string]interface{}) RestoreDataObjectResult {
	return RestoreDataObjectResult{
		Item: func() *DataObject {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RestoreDataObjectResult) ToDict() map[string]interface{} {
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

func (p RestoreDataObjectResult) Pointer() *RestoreDataObjectResult {
	return &p
}

type DescribeDataObjectHistoriesResult struct {
	Items         []DataObjectHistory  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeDataObjectHistoriesAsyncResult struct {
	result *DescribeDataObjectHistoriesResult
	err    error
}

func NewDescribeDataObjectHistoriesResultFromJson(data string) DescribeDataObjectHistoriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDataObjectHistoriesResultFromDict(dict)
}

func NewDescribeDataObjectHistoriesResultFromDict(data map[string]interface{}) DescribeDataObjectHistoriesResult {
	return DescribeDataObjectHistoriesResult{
		Items: func() []DataObjectHistory {
			if data["items"] == nil {
				return nil
			}
			return CastDataObjectHistories(core.CastArray(data["items"]))
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

func (p DescribeDataObjectHistoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectHistoriesFromDict(
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

func (p DescribeDataObjectHistoriesResult) Pointer() *DescribeDataObjectHistoriesResult {
	return &p
}

type DescribeDataObjectHistoriesByUserIdResult struct {
	Items         []DataObjectHistory  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeDataObjectHistoriesByUserIdAsyncResult struct {
	result *DescribeDataObjectHistoriesByUserIdResult
	err    error
}

func NewDescribeDataObjectHistoriesByUserIdResultFromJson(data string) DescribeDataObjectHistoriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDataObjectHistoriesByUserIdResultFromDict(dict)
}

func NewDescribeDataObjectHistoriesByUserIdResultFromDict(data map[string]interface{}) DescribeDataObjectHistoriesByUserIdResult {
	return DescribeDataObjectHistoriesByUserIdResult{
		Items: func() []DataObjectHistory {
			if data["items"] == nil {
				return nil
			}
			return CastDataObjectHistories(core.CastArray(data["items"]))
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

func (p DescribeDataObjectHistoriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDataObjectHistoriesFromDict(
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

func (p DescribeDataObjectHistoriesByUserIdResult) Pointer() *DescribeDataObjectHistoriesByUserIdResult {
	return &p
}

type GetDataObjectHistoryResult struct {
	Item     *DataObjectHistory   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDataObjectHistoryAsyncResult struct {
	result *GetDataObjectHistoryResult
	err    error
}

func NewGetDataObjectHistoryResultFromJson(data string) GetDataObjectHistoryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDataObjectHistoryResultFromDict(dict)
}

func NewGetDataObjectHistoryResultFromDict(data map[string]interface{}) GetDataObjectHistoryResult {
	return GetDataObjectHistoryResult{
		Item: func() *DataObjectHistory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectHistoryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetDataObjectHistoryResult) ToDict() map[string]interface{} {
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

func (p GetDataObjectHistoryResult) Pointer() *GetDataObjectHistoryResult {
	return &p
}

type GetDataObjectHistoryByUserIdResult struct {
	Item     *DataObjectHistory   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDataObjectHistoryByUserIdAsyncResult struct {
	result *GetDataObjectHistoryByUserIdResult
	err    error
}

func NewGetDataObjectHistoryByUserIdResultFromJson(data string) GetDataObjectHistoryByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDataObjectHistoryByUserIdResultFromDict(dict)
}

func NewGetDataObjectHistoryByUserIdResultFromDict(data map[string]interface{}) GetDataObjectHistoryByUserIdResult {
	return GetDataObjectHistoryByUserIdResult{
		Item: func() *DataObjectHistory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataObjectHistoryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetDataObjectHistoryByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetDataObjectHistoryByUserIdResult) Pointer() *GetDataObjectHistoryByUserIdResult {
	return &p
}
