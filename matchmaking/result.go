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

package matchmaking

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

type DescribeGatheringsResult struct {
	Items         []Gathering          `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGatheringsAsyncResult struct {
	result *DescribeGatheringsResult
	err    error
}

func NewDescribeGatheringsResultFromJson(data string) DescribeGatheringsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGatheringsResultFromDict(dict)
}

func NewDescribeGatheringsResultFromDict(data map[string]interface{}) DescribeGatheringsResult {
	return DescribeGatheringsResult{
		Items: func() []Gathering {
			if data["items"] == nil {
				return nil
			}
			return CastGatherings(core.CastArray(data["items"]))
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

func (p DescribeGatheringsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGatheringsFromDict(
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

func (p DescribeGatheringsResult) Pointer() *DescribeGatheringsResult {
	return &p
}

type CreateGatheringResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGatheringAsyncResult struct {
	result *CreateGatheringResult
	err    error
}

func NewCreateGatheringResultFromJson(data string) CreateGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGatheringResultFromDict(dict)
}

func NewCreateGatheringResultFromDict(data map[string]interface{}) CreateGatheringResult {
	return CreateGatheringResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateGatheringResult) ToDict() map[string]interface{} {
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

func (p CreateGatheringResult) Pointer() *CreateGatheringResult {
	return &p
}

type CreateGatheringByUserIdResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGatheringByUserIdAsyncResult struct {
	result *CreateGatheringByUserIdResult
	err    error
}

func NewCreateGatheringByUserIdResultFromJson(data string) CreateGatheringByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGatheringByUserIdResultFromDict(dict)
}

func NewCreateGatheringByUserIdResultFromDict(data map[string]interface{}) CreateGatheringByUserIdResult {
	return CreateGatheringByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateGatheringByUserIdResult) ToDict() map[string]interface{} {
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

func (p CreateGatheringByUserIdResult) Pointer() *CreateGatheringByUserIdResult {
	return &p
}

type UpdateGatheringResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGatheringAsyncResult struct {
	result *UpdateGatheringResult
	err    error
}

func NewUpdateGatheringResultFromJson(data string) UpdateGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGatheringResultFromDict(dict)
}

func NewUpdateGatheringResultFromDict(data map[string]interface{}) UpdateGatheringResult {
	return UpdateGatheringResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateGatheringResult) ToDict() map[string]interface{} {
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

func (p UpdateGatheringResult) Pointer() *UpdateGatheringResult {
	return &p
}

type UpdateGatheringByUserIdResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGatheringByUserIdAsyncResult struct {
	result *UpdateGatheringByUserIdResult
	err    error
}

func NewUpdateGatheringByUserIdResultFromJson(data string) UpdateGatheringByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGatheringByUserIdResultFromDict(dict)
}

func NewUpdateGatheringByUserIdResultFromDict(data map[string]interface{}) UpdateGatheringByUserIdResult {
	return UpdateGatheringByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateGatheringByUserIdResult) ToDict() map[string]interface{} {
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

func (p UpdateGatheringByUserIdResult) Pointer() *UpdateGatheringByUserIdResult {
	return &p
}

type DoMatchmakingByPlayerResult struct {
	Item                    *Gathering           `json:"item"`
	MatchmakingContextToken *string              `json:"matchmakingContextToken"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type DoMatchmakingByPlayerAsyncResult struct {
	result *DoMatchmakingByPlayerResult
	err    error
}

func NewDoMatchmakingByPlayerResultFromJson(data string) DoMatchmakingByPlayerResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingByPlayerResultFromDict(dict)
}

func NewDoMatchmakingByPlayerResultFromDict(data map[string]interface{}) DoMatchmakingByPlayerResult {
	return DoMatchmakingByPlayerResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoMatchmakingByPlayerResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DoMatchmakingByPlayerResult) Pointer() *DoMatchmakingByPlayerResult {
	return &p
}

type DoMatchmakingResult struct {
	Item                    *Gathering           `json:"item"`
	MatchmakingContextToken *string              `json:"matchmakingContextToken"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type DoMatchmakingAsyncResult struct {
	result *DoMatchmakingResult
	err    error
}

func NewDoMatchmakingResultFromJson(data string) DoMatchmakingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingResultFromDict(dict)
}

func NewDoMatchmakingResultFromDict(data map[string]interface{}) DoMatchmakingResult {
	return DoMatchmakingResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoMatchmakingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DoMatchmakingResult) Pointer() *DoMatchmakingResult {
	return &p
}

type DoMatchmakingByUserIdResult struct {
	Item                    *Gathering           `json:"item"`
	MatchmakingContextToken *string              `json:"matchmakingContextToken"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type DoMatchmakingByUserIdAsyncResult struct {
	result *DoMatchmakingByUserIdResult
	err    error
}

func NewDoMatchmakingByUserIdResultFromJson(data string) DoMatchmakingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingByUserIdResultFromDict(dict)
}

func NewDoMatchmakingByUserIdResultFromDict(data map[string]interface{}) DoMatchmakingByUserIdResult {
	return DoMatchmakingByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoMatchmakingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DoMatchmakingByUserIdResult) Pointer() *DoMatchmakingByUserIdResult {
	return &p
}

type PingResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PingAsyncResult struct {
	result *PingResult
	err    error
}

func NewPingResultFromJson(data string) PingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPingResultFromDict(dict)
}

func NewPingResultFromDict(data map[string]interface{}) PingResult {
	return PingResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PingResult) ToDict() map[string]interface{} {
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

func (p PingResult) Pointer() *PingResult {
	return &p
}

type PingByUserIdResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PingByUserIdAsyncResult struct {
	result *PingByUserIdResult
	err    error
}

func NewPingByUserIdResultFromJson(data string) PingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPingByUserIdResultFromDict(dict)
}

func NewPingByUserIdResultFromDict(data map[string]interface{}) PingByUserIdResult {
	return PingByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PingByUserIdResult) ToDict() map[string]interface{} {
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

func (p PingByUserIdResult) Pointer() *PingByUserIdResult {
	return &p
}

type GetGatheringResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGatheringAsyncResult struct {
	result *GetGatheringResult
	err    error
}

func NewGetGatheringResultFromJson(data string) GetGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGatheringResultFromDict(dict)
}

func NewGetGatheringResultFromDict(data map[string]interface{}) GetGatheringResult {
	return GetGatheringResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetGatheringResult) ToDict() map[string]interface{} {
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

func (p GetGatheringResult) Pointer() *GetGatheringResult {
	return &p
}

type CancelMatchmakingResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CancelMatchmakingAsyncResult struct {
	result *CancelMatchmakingResult
	err    error
}

func NewCancelMatchmakingResultFromJson(data string) CancelMatchmakingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCancelMatchmakingResultFromDict(dict)
}

func NewCancelMatchmakingResultFromDict(data map[string]interface{}) CancelMatchmakingResult {
	return CancelMatchmakingResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CancelMatchmakingResult) ToDict() map[string]interface{} {
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

func (p CancelMatchmakingResult) Pointer() *CancelMatchmakingResult {
	return &p
}

type CancelMatchmakingByUserIdResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CancelMatchmakingByUserIdAsyncResult struct {
	result *CancelMatchmakingByUserIdResult
	err    error
}

func NewCancelMatchmakingByUserIdResultFromJson(data string) CancelMatchmakingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCancelMatchmakingByUserIdResultFromDict(dict)
}

func NewCancelMatchmakingByUserIdResultFromDict(data map[string]interface{}) CancelMatchmakingByUserIdResult {
	return CancelMatchmakingByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CancelMatchmakingByUserIdResult) ToDict() map[string]interface{} {
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

func (p CancelMatchmakingByUserIdResult) Pointer() *CancelMatchmakingByUserIdResult {
	return &p
}

type EarlyCompleteResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type EarlyCompleteAsyncResult struct {
	result *EarlyCompleteResult
	err    error
}

func NewEarlyCompleteResultFromJson(data string) EarlyCompleteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEarlyCompleteResultFromDict(dict)
}

func NewEarlyCompleteResultFromDict(data map[string]interface{}) EarlyCompleteResult {
	return EarlyCompleteResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p EarlyCompleteResult) ToDict() map[string]interface{} {
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

func (p EarlyCompleteResult) Pointer() *EarlyCompleteResult {
	return &p
}

type EarlyCompleteByUserIdResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type EarlyCompleteByUserIdAsyncResult struct {
	result *EarlyCompleteByUserIdResult
	err    error
}

func NewEarlyCompleteByUserIdResultFromJson(data string) EarlyCompleteByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEarlyCompleteByUserIdResultFromDict(dict)
}

func NewEarlyCompleteByUserIdResultFromDict(data map[string]interface{}) EarlyCompleteByUserIdResult {
	return EarlyCompleteByUserIdResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p EarlyCompleteByUserIdResult) ToDict() map[string]interface{} {
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

func (p EarlyCompleteByUserIdResult) Pointer() *EarlyCompleteByUserIdResult {
	return &p
}

type DeleteGatheringResult struct {
	Item     *Gathering           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGatheringAsyncResult struct {
	result *DeleteGatheringResult
	err    error
}

func NewDeleteGatheringResultFromJson(data string) DeleteGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGatheringResultFromDict(dict)
}

func NewDeleteGatheringResultFromDict(data map[string]interface{}) DeleteGatheringResult {
	return DeleteGatheringResult{
		Item: func() *Gathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteGatheringResult) ToDict() map[string]interface{} {
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

func (p DeleteGatheringResult) Pointer() *DeleteGatheringResult {
	return &p
}

type DescribeRatingModelMastersResult struct {
	Items         []RatingModelMaster  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRatingModelMastersAsyncResult struct {
	result *DescribeRatingModelMastersResult
	err    error
}

func NewDescribeRatingModelMastersResultFromJson(data string) DescribeRatingModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingModelMastersResultFromDict(dict)
}

func NewDescribeRatingModelMastersResultFromDict(data map[string]interface{}) DescribeRatingModelMastersResult {
	return DescribeRatingModelMastersResult{
		Items: func() []RatingModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRatingModelMasters(core.CastArray(data["items"]))
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

func (p DescribeRatingModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingModelMastersFromDict(
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

func (p DescribeRatingModelMastersResult) Pointer() *DescribeRatingModelMastersResult {
	return &p
}

type CreateRatingModelMasterResult struct {
	Item     *RatingModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateRatingModelMasterAsyncResult struct {
	result *CreateRatingModelMasterResult
	err    error
}

func NewCreateRatingModelMasterResultFromJson(data string) CreateRatingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRatingModelMasterResultFromDict(dict)
}

func NewCreateRatingModelMasterResultFromDict(data map[string]interface{}) CreateRatingModelMasterResult {
	return CreateRatingModelMasterResult{
		Item: func() *RatingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRatingModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateRatingModelMasterResult) Pointer() *CreateRatingModelMasterResult {
	return &p
}

type GetRatingModelMasterResult struct {
	Item     *RatingModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRatingModelMasterAsyncResult struct {
	result *GetRatingModelMasterResult
	err    error
}

func NewGetRatingModelMasterResultFromJson(data string) GetRatingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingModelMasterResultFromDict(dict)
}

func NewGetRatingModelMasterResultFromDict(data map[string]interface{}) GetRatingModelMasterResult {
	return GetRatingModelMasterResult{
		Item: func() *RatingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRatingModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetRatingModelMasterResult) Pointer() *GetRatingModelMasterResult {
	return &p
}

type UpdateRatingModelMasterResult struct {
	Item     *RatingModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateRatingModelMasterAsyncResult struct {
	result *UpdateRatingModelMasterResult
	err    error
}

func NewUpdateRatingModelMasterResultFromJson(data string) UpdateRatingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRatingModelMasterResultFromDict(dict)
}

func NewUpdateRatingModelMasterResultFromDict(data map[string]interface{}) UpdateRatingModelMasterResult {
	return UpdateRatingModelMasterResult{
		Item: func() *RatingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRatingModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateRatingModelMasterResult) Pointer() *UpdateRatingModelMasterResult {
	return &p
}

type DeleteRatingModelMasterResult struct {
	Item     *RatingModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteRatingModelMasterAsyncResult struct {
	result *DeleteRatingModelMasterResult
	err    error
}

func NewDeleteRatingModelMasterResultFromJson(data string) DeleteRatingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRatingModelMasterResultFromDict(dict)
}

func NewDeleteRatingModelMasterResultFromDict(data map[string]interface{}) DeleteRatingModelMasterResult {
	return DeleteRatingModelMasterResult{
		Item: func() *RatingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRatingModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteRatingModelMasterResult) Pointer() *DeleteRatingModelMasterResult {
	return &p
}

type DescribeRatingModelsResult struct {
	Items    []RatingModel        `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeRatingModelsAsyncResult struct {
	result *DescribeRatingModelsResult
	err    error
}

func NewDescribeRatingModelsResultFromJson(data string) DescribeRatingModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingModelsResultFromDict(dict)
}

func NewDescribeRatingModelsResultFromDict(data map[string]interface{}) DescribeRatingModelsResult {
	return DescribeRatingModelsResult{
		Items: func() []RatingModel {
			if data["items"] == nil {
				return nil
			}
			return CastRatingModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeRatingModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingModelsFromDict(
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

func (p DescribeRatingModelsResult) Pointer() *DescribeRatingModelsResult {
	return &p
}

type GetRatingModelResult struct {
	Item     *RatingModel         `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRatingModelAsyncResult struct {
	result *GetRatingModelResult
	err    error
}

func NewGetRatingModelResultFromJson(data string) GetRatingModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingModelResultFromDict(dict)
}

func NewGetRatingModelResultFromDict(data map[string]interface{}) GetRatingModelResult {
	return GetRatingModelResult{
		Item: func() *RatingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRatingModelResult) ToDict() map[string]interface{} {
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

func (p GetRatingModelResult) Pointer() *GetRatingModelResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentModelMaster  `json:"item"`
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
		Item: func() *CurrentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentModelMasterResult struct {
	Item     *CurrentModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentModelMasterAsyncResult struct {
	result *GetCurrentModelMasterResult
	err    error
}

func NewGetCurrentModelMasterResultFromJson(data string) GetCurrentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentModelMasterResultFromDict(dict)
}

func NewGetCurrentModelMasterResultFromDict(data map[string]interface{}) GetCurrentModelMasterResult {
	return GetCurrentModelMasterResult{
		Item: func() *CurrentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetCurrentModelMasterResult) Pointer() *GetCurrentModelMasterResult {
	return &p
}

type PreUpdateCurrentModelMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentModelMasterAsyncResult struct {
	result *PreUpdateCurrentModelMasterResult
	err    error
}

func NewPreUpdateCurrentModelMasterResultFromJson(data string) PreUpdateCurrentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentModelMasterResultFromDict(dict)
}

func NewPreUpdateCurrentModelMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentModelMasterResult {
	return PreUpdateCurrentModelMasterResult{
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

func (p PreUpdateCurrentModelMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentModelMasterResult) Pointer() *PreUpdateCurrentModelMasterResult {
	return &p
}

type UpdateCurrentModelMasterResult struct {
	Item     *CurrentModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentModelMasterAsyncResult struct {
	result *UpdateCurrentModelMasterResult
	err    error
}

func NewUpdateCurrentModelMasterResultFromJson(data string) UpdateCurrentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentModelMasterResultFromDict(dict)
}

func NewUpdateCurrentModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentModelMasterResult {
	return UpdateCurrentModelMasterResult{
		Item: func() *CurrentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentModelMasterResult) Pointer() *UpdateCurrentModelMasterResult {
	return &p
}

type UpdateCurrentModelMasterFromGitHubResult struct {
	Item     *CurrentModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentModelMasterFromGitHubResultFromJson(data string) UpdateCurrentModelMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentModelMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentModelMasterFromGitHubResult {
	return UpdateCurrentModelMasterFromGitHubResult{
		Item: func() *CurrentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentModelMasterFromGitHubResult) ToDict() map[string]interface{} {
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

func (p UpdateCurrentModelMasterFromGitHubResult) Pointer() *UpdateCurrentModelMasterFromGitHubResult {
	return &p
}

type DescribeSeasonModelsResult struct {
	Items    []SeasonModel        `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSeasonModelsAsyncResult struct {
	result *DescribeSeasonModelsResult
	err    error
}

func NewDescribeSeasonModelsResultFromJson(data string) DescribeSeasonModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelsResultFromDict(dict)
}

func NewDescribeSeasonModelsResultFromDict(data map[string]interface{}) DescribeSeasonModelsResult {
	return DescribeSeasonModelsResult{
		Items: func() []SeasonModel {
			if data["items"] == nil {
				return nil
			}
			return CastSeasonModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeSeasonModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonModelsFromDict(
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

func (p DescribeSeasonModelsResult) Pointer() *DescribeSeasonModelsResult {
	return &p
}

type GetSeasonModelResult struct {
	Item     *SeasonModel         `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSeasonModelAsyncResult struct {
	result *GetSeasonModelResult
	err    error
}

func NewGetSeasonModelResultFromJson(data string) GetSeasonModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelResultFromDict(dict)
}

func NewGetSeasonModelResultFromDict(data map[string]interface{}) GetSeasonModelResult {
	return GetSeasonModelResult{
		Item: func() *SeasonModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSeasonModelResult) ToDict() map[string]interface{} {
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

func (p GetSeasonModelResult) Pointer() *GetSeasonModelResult {
	return &p
}

type DescribeSeasonModelMastersResult struct {
	Items         []SeasonModelMaster  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSeasonModelMastersAsyncResult struct {
	result *DescribeSeasonModelMastersResult
	err    error
}

func NewDescribeSeasonModelMastersResultFromJson(data string) DescribeSeasonModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonModelMastersResultFromDict(dict)
}

func NewDescribeSeasonModelMastersResultFromDict(data map[string]interface{}) DescribeSeasonModelMastersResult {
	return DescribeSeasonModelMastersResult{
		Items: func() []SeasonModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSeasonModelMasters(core.CastArray(data["items"]))
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

func (p DescribeSeasonModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonModelMastersFromDict(
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

func (p DescribeSeasonModelMastersResult) Pointer() *DescribeSeasonModelMastersResult {
	return &p
}

type CreateSeasonModelMasterResult struct {
	Item     *SeasonModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateSeasonModelMasterAsyncResult struct {
	result *CreateSeasonModelMasterResult
	err    error
}

func NewCreateSeasonModelMasterResultFromJson(data string) CreateSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSeasonModelMasterResultFromDict(dict)
}

func NewCreateSeasonModelMasterResultFromDict(data map[string]interface{}) CreateSeasonModelMasterResult {
	return CreateSeasonModelMasterResult{
		Item: func() *SeasonModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateSeasonModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateSeasonModelMasterResult) Pointer() *CreateSeasonModelMasterResult {
	return &p
}

type GetSeasonModelMasterResult struct {
	Item     *SeasonModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSeasonModelMasterAsyncResult struct {
	result *GetSeasonModelMasterResult
	err    error
}

func NewGetSeasonModelMasterResultFromJson(data string) GetSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonModelMasterResultFromDict(dict)
}

func NewGetSeasonModelMasterResultFromDict(data map[string]interface{}) GetSeasonModelMasterResult {
	return GetSeasonModelMasterResult{
		Item: func() *SeasonModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSeasonModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetSeasonModelMasterResult) Pointer() *GetSeasonModelMasterResult {
	return &p
}

type UpdateSeasonModelMasterResult struct {
	Item     *SeasonModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateSeasonModelMasterAsyncResult struct {
	result *UpdateSeasonModelMasterResult
	err    error
}

func NewUpdateSeasonModelMasterResultFromJson(data string) UpdateSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSeasonModelMasterResultFromDict(dict)
}

func NewUpdateSeasonModelMasterResultFromDict(data map[string]interface{}) UpdateSeasonModelMasterResult {
	return UpdateSeasonModelMasterResult{
		Item: func() *SeasonModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateSeasonModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateSeasonModelMasterResult) Pointer() *UpdateSeasonModelMasterResult {
	return &p
}

type DeleteSeasonModelMasterResult struct {
	Item     *SeasonModelMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSeasonModelMasterAsyncResult struct {
	result *DeleteSeasonModelMasterResult
	err    error
}

func NewDeleteSeasonModelMasterResultFromJson(data string) DeleteSeasonModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSeasonModelMasterResultFromDict(dict)
}

func NewDeleteSeasonModelMasterResultFromDict(data map[string]interface{}) DeleteSeasonModelMasterResult {
	return DeleteSeasonModelMasterResult{
		Item: func() *SeasonModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteSeasonModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteSeasonModelMasterResult) Pointer() *DeleteSeasonModelMasterResult {
	return &p
}

type DescribeSeasonGatheringsResult struct {
	Items         []SeasonGathering    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSeasonGatheringsAsyncResult struct {
	result *DescribeSeasonGatheringsResult
	err    error
}

func NewDescribeSeasonGatheringsResultFromJson(data string) DescribeSeasonGatheringsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSeasonGatheringsResultFromDict(dict)
}

func NewDescribeSeasonGatheringsResultFromDict(data map[string]interface{}) DescribeSeasonGatheringsResult {
	return DescribeSeasonGatheringsResult{
		Items: func() []SeasonGathering {
			if data["items"] == nil {
				return nil
			}
			return CastSeasonGatherings(core.CastArray(data["items"]))
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

func (p DescribeSeasonGatheringsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonGatheringsFromDict(
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

func (p DescribeSeasonGatheringsResult) Pointer() *DescribeSeasonGatheringsResult {
	return &p
}

type DescribeMatchmakingSeasonGatheringsResult struct {
	Items         []SeasonGathering    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeMatchmakingSeasonGatheringsAsyncResult struct {
	result *DescribeMatchmakingSeasonGatheringsResult
	err    error
}

func NewDescribeMatchmakingSeasonGatheringsResultFromJson(data string) DescribeMatchmakingSeasonGatheringsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMatchmakingSeasonGatheringsResultFromDict(dict)
}

func NewDescribeMatchmakingSeasonGatheringsResultFromDict(data map[string]interface{}) DescribeMatchmakingSeasonGatheringsResult {
	return DescribeMatchmakingSeasonGatheringsResult{
		Items: func() []SeasonGathering {
			if data["items"] == nil {
				return nil
			}
			return CastSeasonGatherings(core.CastArray(data["items"]))
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

func (p DescribeMatchmakingSeasonGatheringsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSeasonGatheringsFromDict(
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

func (p DescribeMatchmakingSeasonGatheringsResult) Pointer() *DescribeMatchmakingSeasonGatheringsResult {
	return &p
}

type DoSeasonMatchmakingResult struct {
	Item                    *SeasonGathering     `json:"item"`
	MatchmakingContextToken *string              `json:"matchmakingContextToken"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type DoSeasonMatchmakingAsyncResult struct {
	result *DoSeasonMatchmakingResult
	err    error
}

func NewDoSeasonMatchmakingResultFromJson(data string) DoSeasonMatchmakingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoSeasonMatchmakingResultFromDict(dict)
}

func NewDoSeasonMatchmakingResultFromDict(data map[string]interface{}) DoSeasonMatchmakingResult {
	return DoSeasonMatchmakingResult{
		Item: func() *SeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoSeasonMatchmakingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DoSeasonMatchmakingResult) Pointer() *DoSeasonMatchmakingResult {
	return &p
}

type DoSeasonMatchmakingByUserIdResult struct {
	Item                    *SeasonGathering     `json:"item"`
	MatchmakingContextToken *string              `json:"matchmakingContextToken"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type DoSeasonMatchmakingByUserIdAsyncResult struct {
	result *DoSeasonMatchmakingByUserIdResult
	err    error
}

func NewDoSeasonMatchmakingByUserIdResultFromJson(data string) DoSeasonMatchmakingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoSeasonMatchmakingByUserIdResultFromDict(dict)
}

func NewDoSeasonMatchmakingByUserIdResultFromDict(data map[string]interface{}) DoSeasonMatchmakingByUserIdResult {
	return DoSeasonMatchmakingByUserIdResult{
		Item: func() *SeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		MatchmakingContextToken: func() *string {
			v, ok := data["matchmakingContextToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["matchmakingContextToken"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoSeasonMatchmakingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DoSeasonMatchmakingByUserIdResult) Pointer() *DoSeasonMatchmakingByUserIdResult {
	return &p
}

type GetSeasonGatheringResult struct {
	Item     *SeasonGathering     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSeasonGatheringAsyncResult struct {
	result *GetSeasonGatheringResult
	err    error
}

func NewGetSeasonGatheringResultFromJson(data string) GetSeasonGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSeasonGatheringResultFromDict(dict)
}

func NewGetSeasonGatheringResultFromDict(data map[string]interface{}) GetSeasonGatheringResult {
	return GetSeasonGatheringResult{
		Item: func() *SeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSeasonGatheringResult) ToDict() map[string]interface{} {
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

func (p GetSeasonGatheringResult) Pointer() *GetSeasonGatheringResult {
	return &p
}

type VerifyIncludeParticipantResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeParticipantAsyncResult struct {
	result *VerifyIncludeParticipantResult
	err    error
}

func NewVerifyIncludeParticipantResultFromJson(data string) VerifyIncludeParticipantResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeParticipantResultFromDict(dict)
}

func NewVerifyIncludeParticipantResultFromDict(data map[string]interface{}) VerifyIncludeParticipantResult {
	return VerifyIncludeParticipantResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyIncludeParticipantResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyIncludeParticipantResult) Pointer() *VerifyIncludeParticipantResult {
	return &p
}

type VerifyIncludeParticipantByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeParticipantByUserIdAsyncResult struct {
	result *VerifyIncludeParticipantByUserIdResult
	err    error
}

func NewVerifyIncludeParticipantByUserIdResultFromJson(data string) VerifyIncludeParticipantByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeParticipantByUserIdResultFromDict(dict)
}

func NewVerifyIncludeParticipantByUserIdResultFromDict(data map[string]interface{}) VerifyIncludeParticipantByUserIdResult {
	return VerifyIncludeParticipantByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyIncludeParticipantByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyIncludeParticipantByUserIdResult) Pointer() *VerifyIncludeParticipantByUserIdResult {
	return &p
}

type DeleteSeasonGatheringResult struct {
	Item     *SeasonGathering     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSeasonGatheringAsyncResult struct {
	result *DeleteSeasonGatheringResult
	err    error
}

func NewDeleteSeasonGatheringResultFromJson(data string) DeleteSeasonGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSeasonGatheringResultFromDict(dict)
}

func NewDeleteSeasonGatheringResultFromDict(data map[string]interface{}) DeleteSeasonGatheringResult {
	return DeleteSeasonGatheringResult{
		Item: func() *SeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteSeasonGatheringResult) ToDict() map[string]interface{} {
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

func (p DeleteSeasonGatheringResult) Pointer() *DeleteSeasonGatheringResult {
	return &p
}

type VerifyIncludeParticipantByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyIncludeParticipantByStampTaskAsyncResult struct {
	result *VerifyIncludeParticipantByStampTaskResult
	err    error
}

func NewVerifyIncludeParticipantByStampTaskResultFromJson(data string) VerifyIncludeParticipantByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyIncludeParticipantByStampTaskResultFromDict(dict)
}

func NewVerifyIncludeParticipantByStampTaskResultFromDict(data map[string]interface{}) VerifyIncludeParticipantByStampTaskResult {
	return VerifyIncludeParticipantByStampTaskResult{
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

func (p VerifyIncludeParticipantByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyIncludeParticipantByStampTaskResult) Pointer() *VerifyIncludeParticipantByStampTaskResult {
	return &p
}

type DescribeJoinedSeasonGatheringsResult struct {
	Items         []JoinedSeasonGathering `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeJoinedSeasonGatheringsAsyncResult struct {
	result *DescribeJoinedSeasonGatheringsResult
	err    error
}

func NewDescribeJoinedSeasonGatheringsResultFromJson(data string) DescribeJoinedSeasonGatheringsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedSeasonGatheringsResultFromDict(dict)
}

func NewDescribeJoinedSeasonGatheringsResultFromDict(data map[string]interface{}) DescribeJoinedSeasonGatheringsResult {
	return DescribeJoinedSeasonGatheringsResult{
		Items: func() []JoinedSeasonGathering {
			if data["items"] == nil {
				return nil
			}
			return CastJoinedSeasonGatherings(core.CastArray(data["items"]))
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

func (p DescribeJoinedSeasonGatheringsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJoinedSeasonGatheringsFromDict(
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

func (p DescribeJoinedSeasonGatheringsResult) Pointer() *DescribeJoinedSeasonGatheringsResult {
	return &p
}

type DescribeJoinedSeasonGatheringsByUserIdResult struct {
	Items         []JoinedSeasonGathering `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeJoinedSeasonGatheringsByUserIdAsyncResult struct {
	result *DescribeJoinedSeasonGatheringsByUserIdResult
	err    error
}

func NewDescribeJoinedSeasonGatheringsByUserIdResultFromJson(data string) DescribeJoinedSeasonGatheringsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedSeasonGatheringsByUserIdResultFromDict(dict)
}

func NewDescribeJoinedSeasonGatheringsByUserIdResultFromDict(data map[string]interface{}) DescribeJoinedSeasonGatheringsByUserIdResult {
	return DescribeJoinedSeasonGatheringsByUserIdResult{
		Items: func() []JoinedSeasonGathering {
			if data["items"] == nil {
				return nil
			}
			return CastJoinedSeasonGatherings(core.CastArray(data["items"]))
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

func (p DescribeJoinedSeasonGatheringsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastJoinedSeasonGatheringsFromDict(
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

func (p DescribeJoinedSeasonGatheringsByUserIdResult) Pointer() *DescribeJoinedSeasonGatheringsByUserIdResult {
	return &p
}

type GetJoinedSeasonGatheringResult struct {
	Item     *JoinedSeasonGathering `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetJoinedSeasonGatheringAsyncResult struct {
	result *GetJoinedSeasonGatheringResult
	err    error
}

func NewGetJoinedSeasonGatheringResultFromJson(data string) GetJoinedSeasonGatheringResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedSeasonGatheringResultFromDict(dict)
}

func NewGetJoinedSeasonGatheringResultFromDict(data map[string]interface{}) GetJoinedSeasonGatheringResult {
	return GetJoinedSeasonGatheringResult{
		Item: func() *JoinedSeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetJoinedSeasonGatheringResult) ToDict() map[string]interface{} {
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

func (p GetJoinedSeasonGatheringResult) Pointer() *GetJoinedSeasonGatheringResult {
	return &p
}

type GetJoinedSeasonGatheringByUserIdResult struct {
	Item     *JoinedSeasonGathering `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetJoinedSeasonGatheringByUserIdAsyncResult struct {
	result *GetJoinedSeasonGatheringByUserIdResult
	err    error
}

func NewGetJoinedSeasonGatheringByUserIdResultFromJson(data string) GetJoinedSeasonGatheringByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedSeasonGatheringByUserIdResultFromDict(dict)
}

func NewGetJoinedSeasonGatheringByUserIdResultFromDict(data map[string]interface{}) GetJoinedSeasonGatheringByUserIdResult {
	return GetJoinedSeasonGatheringByUserIdResult{
		Item: func() *JoinedSeasonGathering {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewJoinedSeasonGatheringFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetJoinedSeasonGatheringByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetJoinedSeasonGatheringByUserIdResult) Pointer() *GetJoinedSeasonGatheringByUserIdResult {
	return &p
}

type DescribeRatingsResult struct {
	Items         []Rating             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRatingsAsyncResult struct {
	result *DescribeRatingsResult
	err    error
}

func NewDescribeRatingsResultFromJson(data string) DescribeRatingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingsResultFromDict(dict)
}

func NewDescribeRatingsResultFromDict(data map[string]interface{}) DescribeRatingsResult {
	return DescribeRatingsResult{
		Items: func() []Rating {
			if data["items"] == nil {
				return nil
			}
			return CastRatings(core.CastArray(data["items"]))
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

func (p DescribeRatingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
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

func (p DescribeRatingsResult) Pointer() *DescribeRatingsResult {
	return &p
}

type DescribeRatingsByUserIdResult struct {
	Items         []Rating             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRatingsByUserIdAsyncResult struct {
	result *DescribeRatingsByUserIdResult
	err    error
}

func NewDescribeRatingsByUserIdResultFromJson(data string) DescribeRatingsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingsByUserIdResultFromDict(dict)
}

func NewDescribeRatingsByUserIdResultFromDict(data map[string]interface{}) DescribeRatingsByUserIdResult {
	return DescribeRatingsByUserIdResult{
		Items: func() []Rating {
			if data["items"] == nil {
				return nil
			}
			return CastRatings(core.CastArray(data["items"]))
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

func (p DescribeRatingsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
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

func (p DescribeRatingsByUserIdResult) Pointer() *DescribeRatingsByUserIdResult {
	return &p
}

type GetRatingResult struct {
	Item     *Rating              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRatingAsyncResult struct {
	result *GetRatingResult
	err    error
}

func NewGetRatingResultFromJson(data string) GetRatingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingResultFromDict(dict)
}

func NewGetRatingResultFromDict(data map[string]interface{}) GetRatingResult {
	return GetRatingResult{
		Item: func() *Rating {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRatingResult) ToDict() map[string]interface{} {
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

func (p GetRatingResult) Pointer() *GetRatingResult {
	return &p
}

type GetRatingByUserIdResult struct {
	Item     *Rating              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRatingByUserIdAsyncResult struct {
	result *GetRatingByUserIdResult
	err    error
}

func NewGetRatingByUserIdResultFromJson(data string) GetRatingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingByUserIdResultFromDict(dict)
}

func NewGetRatingByUserIdResultFromDict(data map[string]interface{}) GetRatingByUserIdResult {
	return GetRatingByUserIdResult{
		Item: func() *Rating {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRatingByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetRatingByUserIdResult) Pointer() *GetRatingByUserIdResult {
	return &p
}

type PutResultResult struct {
	Items    []Rating             `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutResultAsyncResult struct {
	result *PutResultResult
	err    error
}

func NewPutResultResultFromJson(data string) PutResultResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutResultResultFromDict(dict)
}

func NewPutResultResultFromDict(data map[string]interface{}) PutResultResult {
	return PutResultResult{
		Items: func() []Rating {
			if data["items"] == nil {
				return nil
			}
			return CastRatings(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p PutResultResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
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

func (p PutResultResult) Pointer() *PutResultResult {
	return &p
}

type DeleteRatingResult struct {
	Item     *Rating              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteRatingAsyncResult struct {
	result *DeleteRatingResult
	err    error
}

func NewDeleteRatingResultFromJson(data string) DeleteRatingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRatingResultFromDict(dict)
}

func NewDeleteRatingResultFromDict(data map[string]interface{}) DeleteRatingResult {
	return DeleteRatingResult{
		Item: func() *Rating {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRatingFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRatingResult) ToDict() map[string]interface{} {
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

func (p DeleteRatingResult) Pointer() *DeleteRatingResult {
	return &p
}

type GetBallotResult struct {
	Item      *Ballot              `json:"item"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetBallotAsyncResult struct {
	result *GetBallotResult
	err    error
}

func NewGetBallotResultFromJson(data string) GetBallotResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotResultFromDict(dict)
}

func NewGetBallotResultFromDict(data map[string]interface{}) GetBallotResult {
	return GetBallotResult{
		Item: func() *Ballot {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBallotFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBallotResult) ToDict() map[string]interface{} {
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

func (p GetBallotResult) Pointer() *GetBallotResult {
	return &p
}

type GetBallotByUserIdResult struct {
	Item      *Ballot              `json:"item"`
	Body      *string              `json:"body"`
	Signature *string              `json:"signature"`
	Metadata  *core.ResultMetadata `json:"metadata"`
}

type GetBallotByUserIdAsyncResult struct {
	result *GetBallotByUserIdResult
	err    error
}

func NewGetBallotByUserIdResultFromJson(data string) GetBallotByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotByUserIdResultFromDict(dict)
}

func NewGetBallotByUserIdResultFromDict(data map[string]interface{}) GetBallotByUserIdResult {
	return GetBallotByUserIdResult{
		Item: func() *Ballot {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBallotFromDict(core.CastMap(data["item"])).Pointer()
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

func (p GetBallotByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetBallotByUserIdResult) Pointer() *GetBallotByUserIdResult {
	return &p
}

type VoteResult struct {
	Item     *Ballot              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VoteAsyncResult struct {
	result *VoteResult
	err    error
}

func NewVoteResultFromJson(data string) VoteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteResultFromDict(dict)
}

func NewVoteResultFromDict(data map[string]interface{}) VoteResult {
	return VoteResult{
		Item: func() *Ballot {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBallotFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VoteResult) ToDict() map[string]interface{} {
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

func (p VoteResult) Pointer() *VoteResult {
	return &p
}

type VoteMultipleResult struct {
	Item     *Ballot              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VoteMultipleAsyncResult struct {
	result *VoteMultipleResult
	err    error
}

func NewVoteMultipleResultFromJson(data string) VoteMultipleResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteMultipleResultFromDict(dict)
}

func NewVoteMultipleResultFromDict(data map[string]interface{}) VoteMultipleResult {
	return VoteMultipleResult{
		Item: func() *Ballot {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBallotFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VoteMultipleResult) ToDict() map[string]interface{} {
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

func (p VoteMultipleResult) Pointer() *VoteMultipleResult {
	return &p
}

type CommitVoteResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CommitVoteAsyncResult struct {
	result *CommitVoteResult
	err    error
}

func NewCommitVoteResultFromJson(data string) CommitVoteResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCommitVoteResultFromDict(dict)
}

func NewCommitVoteResultFromDict(data map[string]interface{}) CommitVoteResult {
	return CommitVoteResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CommitVoteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CommitVoteResult) Pointer() *CommitVoteResult {
	return &p
}
