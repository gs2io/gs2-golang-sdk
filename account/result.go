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

package account

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

type DescribeAccountsResult struct {
	Items         []Account            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAccountsAsyncResult struct {
	result *DescribeAccountsResult
	err    error
}

func NewDescribeAccountsResultFromJson(data string) DescribeAccountsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAccountsResultFromDict(dict)
}

func NewDescribeAccountsResultFromDict(data map[string]interface{}) DescribeAccountsResult {
	return DescribeAccountsResult{
		Items: func() []Account {
			if data["items"] == nil {
				return nil
			}
			return CastAccounts(core.CastArray(data["items"]))
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

func (p DescribeAccountsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccountsFromDict(
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

func (p DescribeAccountsResult) Pointer() *DescribeAccountsResult {
	return &p
}

type CreateAccountResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateAccountAsyncResult struct {
	result *CreateAccountResult
	err    error
}

func NewCreateAccountResultFromJson(data string) CreateAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAccountResultFromDict(dict)
}

func NewCreateAccountResultFromDict(data map[string]interface{}) CreateAccountResult {
	return CreateAccountResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateAccountResult) ToDict() map[string]interface{} {
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

func (p CreateAccountResult) Pointer() *CreateAccountResult {
	return &p
}

type UpdateTimeOffsetResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateTimeOffsetAsyncResult struct {
	result *UpdateTimeOffsetResult
	err    error
}

func NewUpdateTimeOffsetResultFromJson(data string) UpdateTimeOffsetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTimeOffsetResultFromDict(dict)
}

func NewUpdateTimeOffsetResultFromDict(data map[string]interface{}) UpdateTimeOffsetResult {
	return UpdateTimeOffsetResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateTimeOffsetResult) ToDict() map[string]interface{} {
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

func (p UpdateTimeOffsetResult) Pointer() *UpdateTimeOffsetResult {
	return &p
}

type UpdateBannedResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateBannedAsyncResult struct {
	result *UpdateBannedResult
	err    error
}

func NewUpdateBannedResultFromJson(data string) UpdateBannedResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBannedResultFromDict(dict)
}

func NewUpdateBannedResultFromDict(data map[string]interface{}) UpdateBannedResult {
	return UpdateBannedResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateBannedResult) ToDict() map[string]interface{} {
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

func (p UpdateBannedResult) Pointer() *UpdateBannedResult {
	return &p
}

type AddBanResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddBanAsyncResult struct {
	result *AddBanResult
	err    error
}

func NewAddBanResultFromJson(data string) AddBanResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddBanResultFromDict(dict)
}

func NewAddBanResultFromDict(data map[string]interface{}) AddBanResult {
	return AddBanResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AddBanResult) ToDict() map[string]interface{} {
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

func (p AddBanResult) Pointer() *AddBanResult {
	return &p
}

type RemoveBanResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RemoveBanAsyncResult struct {
	result *RemoveBanResult
	err    error
}

func NewRemoveBanResultFromJson(data string) RemoveBanResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRemoveBanResultFromDict(dict)
}

func NewRemoveBanResultFromDict(data map[string]interface{}) RemoveBanResult {
	return RemoveBanResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RemoveBanResult) ToDict() map[string]interface{} {
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

func (p RemoveBanResult) Pointer() *RemoveBanResult {
	return &p
}

type GetAccountResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAccountAsyncResult struct {
	result *GetAccountResult
	err    error
}

func NewGetAccountResultFromJson(data string) GetAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAccountResultFromDict(dict)
}

func NewGetAccountResultFromDict(data map[string]interface{}) GetAccountResult {
	return GetAccountResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetAccountResult) ToDict() map[string]interface{} {
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

func (p GetAccountResult) Pointer() *GetAccountResult {
	return &p
}

type DeleteAccountResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAccountAsyncResult struct {
	result *DeleteAccountResult
	err    error
}

func NewDeleteAccountResultFromJson(data string) DeleteAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAccountResultFromDict(dict)
}

func NewDeleteAccountResultFromDict(data map[string]interface{}) DeleteAccountResult {
	return DeleteAccountResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteAccountResult) ToDict() map[string]interface{} {
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

func (p DeleteAccountResult) Pointer() *DeleteAccountResult {
	return &p
}

type AuthenticationResult struct {
	Item        *Account             `json:"item"`
	BanStatuses []BanStatus          `json:"banStatuses"`
	Body        *string              `json:"body"`
	Signature   *string              `json:"signature"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type AuthenticationAsyncResult struct {
	result *AuthenticationResult
	err    error
}

func NewAuthenticationResultFromJson(data string) AuthenticationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAuthenticationResultFromDict(dict)
}

func NewAuthenticationResultFromDict(data map[string]interface{}) AuthenticationResult {
	return AuthenticationResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		BanStatuses: func() []BanStatus {
			if data["banStatuses"] == nil {
				return nil
			}
			return CastBanStatuses(core.CastArray(data["banStatuses"]))
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

func (p AuthenticationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"banStatuses": CastBanStatusesFromDict(
			p.BanStatuses,
		),
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

func (p AuthenticationResult) Pointer() *AuthenticationResult {
	return &p
}

type DescribeTakeOversResult struct {
	Items         []TakeOver           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeTakeOversAsyncResult struct {
	result *DescribeTakeOversResult
	err    error
}

func NewDescribeTakeOversResultFromJson(data string) DescribeTakeOversResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversResultFromDict(dict)
}

func NewDescribeTakeOversResultFromDict(data map[string]interface{}) DescribeTakeOversResult {
	return DescribeTakeOversResult{
		Items: func() []TakeOver {
			if data["items"] == nil {
				return nil
			}
			return CastTakeOvers(core.CastArray(data["items"]))
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

func (p DescribeTakeOversResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOversFromDict(
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

func (p DescribeTakeOversResult) Pointer() *DescribeTakeOversResult {
	return &p
}

type DescribeTakeOversByUserIdResult struct {
	Items         []TakeOver           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeTakeOversByUserIdAsyncResult struct {
	result *DescribeTakeOversByUserIdResult
	err    error
}

func NewDescribeTakeOversByUserIdResultFromJson(data string) DescribeTakeOversByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversByUserIdResultFromDict(dict)
}

func NewDescribeTakeOversByUserIdResultFromDict(data map[string]interface{}) DescribeTakeOversByUserIdResult {
	return DescribeTakeOversByUserIdResult{
		Items: func() []TakeOver {
			if data["items"] == nil {
				return nil
			}
			return CastTakeOvers(core.CastArray(data["items"]))
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

func (p DescribeTakeOversByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOversFromDict(
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

func (p DescribeTakeOversByUserIdResult) Pointer() *DescribeTakeOversByUserIdResult {
	return &p
}

type CreateTakeOverResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateTakeOverAsyncResult struct {
	result *CreateTakeOverResult
	err    error
}

func NewCreateTakeOverResultFromJson(data string) CreateTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverResultFromDict(dict)
}

func NewCreateTakeOverResultFromDict(data map[string]interface{}) CreateTakeOverResult {
	return CreateTakeOverResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateTakeOverResult) ToDict() map[string]interface{} {
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

func (p CreateTakeOverResult) Pointer() *CreateTakeOverResult {
	return &p
}

type CreateTakeOverByUserIdResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateTakeOverByUserIdAsyncResult struct {
	result *CreateTakeOverByUserIdResult
	err    error
}

func NewCreateTakeOverByUserIdResultFromJson(data string) CreateTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverByUserIdResultFromDict(dict)
}

func NewCreateTakeOverByUserIdResultFromDict(data map[string]interface{}) CreateTakeOverByUserIdResult {
	return CreateTakeOverByUserIdResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateTakeOverByUserIdResult) ToDict() map[string]interface{} {
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

func (p CreateTakeOverByUserIdResult) Pointer() *CreateTakeOverByUserIdResult {
	return &p
}

type CreateTakeOverOpenIdConnectResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateTakeOverOpenIdConnectAsyncResult struct {
	result *CreateTakeOverOpenIdConnectResult
	err    error
}

func NewCreateTakeOverOpenIdConnectResultFromJson(data string) CreateTakeOverOpenIdConnectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverOpenIdConnectResultFromDict(dict)
}

func NewCreateTakeOverOpenIdConnectResultFromDict(data map[string]interface{}) CreateTakeOverOpenIdConnectResult {
	return CreateTakeOverOpenIdConnectResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateTakeOverOpenIdConnectResult) ToDict() map[string]interface{} {
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

func (p CreateTakeOverOpenIdConnectResult) Pointer() *CreateTakeOverOpenIdConnectResult {
	return &p
}

type CreateTakeOverOpenIdConnectAndByUserIdResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateTakeOverOpenIdConnectAndByUserIdAsyncResult struct {
	result *CreateTakeOverOpenIdConnectAndByUserIdResult
	err    error
}

func NewCreateTakeOverOpenIdConnectAndByUserIdResultFromJson(data string) CreateTakeOverOpenIdConnectAndByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverOpenIdConnectAndByUserIdResultFromDict(dict)
}

func NewCreateTakeOverOpenIdConnectAndByUserIdResultFromDict(data map[string]interface{}) CreateTakeOverOpenIdConnectAndByUserIdResult {
	return CreateTakeOverOpenIdConnectAndByUserIdResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateTakeOverOpenIdConnectAndByUserIdResult) ToDict() map[string]interface{} {
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

func (p CreateTakeOverOpenIdConnectAndByUserIdResult) Pointer() *CreateTakeOverOpenIdConnectAndByUserIdResult {
	return &p
}

type GetTakeOverResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetTakeOverAsyncResult struct {
	result *GetTakeOverResult
	err    error
}

func NewGetTakeOverResultFromJson(data string) GetTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverResultFromDict(dict)
}

func NewGetTakeOverResultFromDict(data map[string]interface{}) GetTakeOverResult {
	return GetTakeOverResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetTakeOverResult) ToDict() map[string]interface{} {
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

func (p GetTakeOverResult) Pointer() *GetTakeOverResult {
	return &p
}

type GetTakeOverByUserIdResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetTakeOverByUserIdAsyncResult struct {
	result *GetTakeOverByUserIdResult
	err    error
}

func NewGetTakeOverByUserIdResultFromJson(data string) GetTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverByUserIdResultFromDict(dict)
}

func NewGetTakeOverByUserIdResultFromDict(data map[string]interface{}) GetTakeOverByUserIdResult {
	return GetTakeOverByUserIdResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetTakeOverByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetTakeOverByUserIdResult) Pointer() *GetTakeOverByUserIdResult {
	return &p
}

type UpdateTakeOverResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateTakeOverAsyncResult struct {
	result *UpdateTakeOverResult
	err    error
}

func NewUpdateTakeOverResultFromJson(data string) UpdateTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverResultFromDict(dict)
}

func NewUpdateTakeOverResultFromDict(data map[string]interface{}) UpdateTakeOverResult {
	return UpdateTakeOverResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateTakeOverResult) ToDict() map[string]interface{} {
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

func (p UpdateTakeOverResult) Pointer() *UpdateTakeOverResult {
	return &p
}

type UpdateTakeOverByUserIdResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateTakeOverByUserIdAsyncResult struct {
	result *UpdateTakeOverByUserIdResult
	err    error
}

func NewUpdateTakeOverByUserIdResultFromJson(data string) UpdateTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverByUserIdResultFromDict(dict)
}

func NewUpdateTakeOverByUserIdResultFromDict(data map[string]interface{}) UpdateTakeOverByUserIdResult {
	return UpdateTakeOverByUserIdResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateTakeOverByUserIdResult) ToDict() map[string]interface{} {
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

func (p UpdateTakeOverByUserIdResult) Pointer() *UpdateTakeOverByUserIdResult {
	return &p
}

type DeleteTakeOverResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteTakeOverAsyncResult struct {
	result *DeleteTakeOverResult
	err    error
}

func NewDeleteTakeOverResultFromJson(data string) DeleteTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverResultFromDict(dict)
}

func NewDeleteTakeOverResultFromDict(data map[string]interface{}) DeleteTakeOverResult {
	return DeleteTakeOverResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteTakeOverResult) ToDict() map[string]interface{} {
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

func (p DeleteTakeOverResult) Pointer() *DeleteTakeOverResult {
	return &p
}

type DeleteTakeOverByUserIdentifierResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteTakeOverByUserIdentifierAsyncResult struct {
	result *DeleteTakeOverByUserIdentifierResult
	err    error
}

func NewDeleteTakeOverByUserIdentifierResultFromJson(data string) DeleteTakeOverByUserIdentifierResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdentifierResultFromDict(dict)
}

func NewDeleteTakeOverByUserIdentifierResultFromDict(data map[string]interface{}) DeleteTakeOverByUserIdentifierResult {
	return DeleteTakeOverByUserIdentifierResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteTakeOverByUserIdentifierResult) ToDict() map[string]interface{} {
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

func (p DeleteTakeOverByUserIdentifierResult) Pointer() *DeleteTakeOverByUserIdentifierResult {
	return &p
}

type DeleteTakeOverByUserIdResult struct {
	Item     *TakeOver            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteTakeOverByUserIdAsyncResult struct {
	result *DeleteTakeOverByUserIdResult
	err    error
}

func NewDeleteTakeOverByUserIdResultFromJson(data string) DeleteTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdResultFromDict(dict)
}

func NewDeleteTakeOverByUserIdResultFromDict(data map[string]interface{}) DeleteTakeOverByUserIdResult {
	return DeleteTakeOverByUserIdResult{
		Item: func() *TakeOver {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteTakeOverByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteTakeOverByUserIdResult) Pointer() *DeleteTakeOverByUserIdResult {
	return &p
}

type DoTakeOverResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DoTakeOverAsyncResult struct {
	result *DoTakeOverResult
	err    error
}

func NewDoTakeOverResultFromJson(data string) DoTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoTakeOverResultFromDict(dict)
}

func NewDoTakeOverResultFromDict(data map[string]interface{}) DoTakeOverResult {
	return DoTakeOverResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoTakeOverResult) ToDict() map[string]interface{} {
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

func (p DoTakeOverResult) Pointer() *DoTakeOverResult {
	return &p
}

type DoTakeOverOpenIdConnectResult struct {
	Item     *Account             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DoTakeOverOpenIdConnectAsyncResult struct {
	result *DoTakeOverOpenIdConnectResult
	err    error
}

func NewDoTakeOverOpenIdConnectResultFromJson(data string) DoTakeOverOpenIdConnectResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoTakeOverOpenIdConnectResultFromDict(dict)
}

func NewDoTakeOverOpenIdConnectResultFromDict(data map[string]interface{}) DoTakeOverOpenIdConnectResult {
	return DoTakeOverOpenIdConnectResult{
		Item: func() *Account {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAccountFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DoTakeOverOpenIdConnectResult) ToDict() map[string]interface{} {
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

func (p DoTakeOverOpenIdConnectResult) Pointer() *DoTakeOverOpenIdConnectResult {
	return &p
}

type GetAuthorizationUrlResult struct {
	AuthorizationUrl *string              `json:"authorizationUrl"`
	Metadata         *core.ResultMetadata `json:"metadata"`
}

type GetAuthorizationUrlAsyncResult struct {
	result *GetAuthorizationUrlResult
	err    error
}

func NewGetAuthorizationUrlResultFromJson(data string) GetAuthorizationUrlResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAuthorizationUrlResultFromDict(dict)
}

func NewGetAuthorizationUrlResultFromDict(data map[string]interface{}) GetAuthorizationUrlResult {
	return GetAuthorizationUrlResult{
		AuthorizationUrl: func() *string {
			v, ok := data["authorizationUrl"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["authorizationUrl"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetAuthorizationUrlResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"authorizationUrl": p.AuthorizationUrl,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetAuthorizationUrlResult) Pointer() *GetAuthorizationUrlResult {
	return &p
}

type DescribePlatformIdsResult struct {
	Items         []PlatformId         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribePlatformIdsAsyncResult struct {
	result *DescribePlatformIdsResult
	err    error
}

func NewDescribePlatformIdsResultFromJson(data string) DescribePlatformIdsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePlatformIdsResultFromDict(dict)
}

func NewDescribePlatformIdsResultFromDict(data map[string]interface{}) DescribePlatformIdsResult {
	return DescribePlatformIdsResult{
		Items: func() []PlatformId {
			if data["items"] == nil {
				return nil
			}
			return CastPlatformIds(core.CastArray(data["items"]))
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

func (p DescribePlatformIdsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPlatformIdsFromDict(
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

func (p DescribePlatformIdsResult) Pointer() *DescribePlatformIdsResult {
	return &p
}

type DescribePlatformIdsByUserIdResult struct {
	Items         []PlatformId         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribePlatformIdsByUserIdAsyncResult struct {
	result *DescribePlatformIdsByUserIdResult
	err    error
}

func NewDescribePlatformIdsByUserIdResultFromJson(data string) DescribePlatformIdsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePlatformIdsByUserIdResultFromDict(dict)
}

func NewDescribePlatformIdsByUserIdResultFromDict(data map[string]interface{}) DescribePlatformIdsByUserIdResult {
	return DescribePlatformIdsByUserIdResult{
		Items: func() []PlatformId {
			if data["items"] == nil {
				return nil
			}
			return CastPlatformIds(core.CastArray(data["items"]))
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

func (p DescribePlatformIdsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPlatformIdsFromDict(
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

func (p DescribePlatformIdsByUserIdResult) Pointer() *DescribePlatformIdsByUserIdResult {
	return &p
}

type CreatePlatformIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreatePlatformIdAsyncResult struct {
	result *CreatePlatformIdResult
	err    error
}

func NewCreatePlatformIdResultFromJson(data string) CreatePlatformIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePlatformIdResultFromDict(dict)
}

func NewCreatePlatformIdResultFromDict(data map[string]interface{}) CreatePlatformIdResult {
	return CreatePlatformIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreatePlatformIdResult) ToDict() map[string]interface{} {
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

func (p CreatePlatformIdResult) Pointer() *CreatePlatformIdResult {
	return &p
}

type CreatePlatformIdByUserIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreatePlatformIdByUserIdAsyncResult struct {
	result *CreatePlatformIdByUserIdResult
	err    error
}

func NewCreatePlatformIdByUserIdResultFromJson(data string) CreatePlatformIdByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePlatformIdByUserIdResultFromDict(dict)
}

func NewCreatePlatformIdByUserIdResultFromDict(data map[string]interface{}) CreatePlatformIdByUserIdResult {
	return CreatePlatformIdByUserIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreatePlatformIdByUserIdResult) ToDict() map[string]interface{} {
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

func (p CreatePlatformIdByUserIdResult) Pointer() *CreatePlatformIdByUserIdResult {
	return &p
}

type GetPlatformIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPlatformIdAsyncResult struct {
	result *GetPlatformIdResult
	err    error
}

func NewGetPlatformIdResultFromJson(data string) GetPlatformIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPlatformIdResultFromDict(dict)
}

func NewGetPlatformIdResultFromDict(data map[string]interface{}) GetPlatformIdResult {
	return GetPlatformIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetPlatformIdResult) ToDict() map[string]interface{} {
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

func (p GetPlatformIdResult) Pointer() *GetPlatformIdResult {
	return &p
}

type GetPlatformIdByUserIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPlatformIdByUserIdAsyncResult struct {
	result *GetPlatformIdByUserIdResult
	err    error
}

func NewGetPlatformIdByUserIdResultFromJson(data string) GetPlatformIdByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPlatformIdByUserIdResultFromDict(dict)
}

func NewGetPlatformIdByUserIdResultFromDict(data map[string]interface{}) GetPlatformIdByUserIdResult {
	return GetPlatformIdByUserIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetPlatformIdByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetPlatformIdByUserIdResult) Pointer() *GetPlatformIdByUserIdResult {
	return &p
}

type FindPlatformIdResult struct {
	Item     *PlatformUser        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FindPlatformIdAsyncResult struct {
	result *FindPlatformIdResult
	err    error
}

func NewFindPlatformIdResultFromJson(data string) FindPlatformIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFindPlatformIdResultFromDict(dict)
}

func NewFindPlatformIdResultFromDict(data map[string]interface{}) FindPlatformIdResult {
	return FindPlatformIdResult{
		Item: func() *PlatformUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p FindPlatformIdResult) ToDict() map[string]interface{} {
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

func (p FindPlatformIdResult) Pointer() *FindPlatformIdResult {
	return &p
}

type FindPlatformIdByUserIdResult struct {
	Item     *PlatformUser        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FindPlatformIdByUserIdAsyncResult struct {
	result *FindPlatformIdByUserIdResult
	err    error
}

func NewFindPlatformIdByUserIdResultFromJson(data string) FindPlatformIdByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFindPlatformIdByUserIdResultFromDict(dict)
}

func NewFindPlatformIdByUserIdResultFromDict(data map[string]interface{}) FindPlatformIdByUserIdResult {
	return FindPlatformIdByUserIdResult{
		Item: func() *PlatformUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p FindPlatformIdByUserIdResult) ToDict() map[string]interface{} {
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

func (p FindPlatformIdByUserIdResult) Pointer() *FindPlatformIdByUserIdResult {
	return &p
}

type DeletePlatformIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeletePlatformIdAsyncResult struct {
	result *DeletePlatformIdResult
	err    error
}

func NewDeletePlatformIdResultFromJson(data string) DeletePlatformIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePlatformIdResultFromDict(dict)
}

func NewDeletePlatformIdResultFromDict(data map[string]interface{}) DeletePlatformIdResult {
	return DeletePlatformIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeletePlatformIdResult) ToDict() map[string]interface{} {
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

func (p DeletePlatformIdResult) Pointer() *DeletePlatformIdResult {
	return &p
}

type DeletePlatformIdByUserIdentifierResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeletePlatformIdByUserIdentifierAsyncResult struct {
	result *DeletePlatformIdByUserIdentifierResult
	err    error
}

func NewDeletePlatformIdByUserIdentifierResultFromJson(data string) DeletePlatformIdByUserIdentifierResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePlatformIdByUserIdentifierResultFromDict(dict)
}

func NewDeletePlatformIdByUserIdentifierResultFromDict(data map[string]interface{}) DeletePlatformIdByUserIdentifierResult {
	return DeletePlatformIdByUserIdentifierResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeletePlatformIdByUserIdentifierResult) ToDict() map[string]interface{} {
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

func (p DeletePlatformIdByUserIdentifierResult) Pointer() *DeletePlatformIdByUserIdentifierResult {
	return &p
}

type DeletePlatformIdByUserIdResult struct {
	Item     *PlatformId          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeletePlatformIdByUserIdAsyncResult struct {
	result *DeletePlatformIdByUserIdResult
	err    error
}

func NewDeletePlatformIdByUserIdResultFromJson(data string) DeletePlatformIdByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePlatformIdByUserIdResultFromDict(dict)
}

func NewDeletePlatformIdByUserIdResultFromDict(data map[string]interface{}) DeletePlatformIdByUserIdResult {
	return DeletePlatformIdByUserIdResult{
		Item: func() *PlatformId {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPlatformIdFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeletePlatformIdByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeletePlatformIdByUserIdResult) Pointer() *DeletePlatformIdByUserIdResult {
	return &p
}

type GetDataOwnerByUserIdResult struct {
	Item     *DataOwner           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetDataOwnerByUserIdAsyncResult struct {
	result *GetDataOwnerByUserIdResult
	err    error
}

func NewGetDataOwnerByUserIdResultFromJson(data string) GetDataOwnerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDataOwnerByUserIdResultFromDict(dict)
}

func NewGetDataOwnerByUserIdResultFromDict(data map[string]interface{}) GetDataOwnerByUserIdResult {
	return GetDataOwnerByUserIdResult{
		Item: func() *DataOwner {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataOwnerFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetDataOwnerByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetDataOwnerByUserIdResult) Pointer() *GetDataOwnerByUserIdResult {
	return &p
}

type UpdateDataOwnerByUserIdResult struct {
	Item     *DataOwner           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateDataOwnerByUserIdAsyncResult struct {
	result *UpdateDataOwnerByUserIdResult
	err    error
}

func NewUpdateDataOwnerByUserIdResultFromJson(data string) UpdateDataOwnerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateDataOwnerByUserIdResultFromDict(dict)
}

func NewUpdateDataOwnerByUserIdResultFromDict(data map[string]interface{}) UpdateDataOwnerByUserIdResult {
	return UpdateDataOwnerByUserIdResult{
		Item: func() *DataOwner {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataOwnerFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateDataOwnerByUserIdResult) ToDict() map[string]interface{} {
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

func (p UpdateDataOwnerByUserIdResult) Pointer() *UpdateDataOwnerByUserIdResult {
	return &p
}

type DeleteDataOwnerByUserIdResult struct {
	Item     *DataOwner           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteDataOwnerByUserIdAsyncResult struct {
	result *DeleteDataOwnerByUserIdResult
	err    error
}

func NewDeleteDataOwnerByUserIdResultFromJson(data string) DeleteDataOwnerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDataOwnerByUserIdResultFromDict(dict)
}

func NewDeleteDataOwnerByUserIdResultFromDict(data map[string]interface{}) DeleteDataOwnerByUserIdResult {
	return DeleteDataOwnerByUserIdResult{
		Item: func() *DataOwner {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDataOwnerFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteDataOwnerByUserIdResult) ToDict() map[string]interface{} {
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

func (p DeleteDataOwnerByUserIdResult) Pointer() *DeleteDataOwnerByUserIdResult {
	return &p
}

type DescribeTakeOverTypeModelsResult struct {
	Items    []TakeOverTypeModel  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeTakeOverTypeModelsAsyncResult struct {
	result *DescribeTakeOverTypeModelsResult
	err    error
}

func NewDescribeTakeOverTypeModelsResultFromJson(data string) DescribeTakeOverTypeModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOverTypeModelsResultFromDict(dict)
}

func NewDescribeTakeOverTypeModelsResultFromDict(data map[string]interface{}) DescribeTakeOverTypeModelsResult {
	return DescribeTakeOverTypeModelsResult{
		Items: func() []TakeOverTypeModel {
			if data["items"] == nil {
				return nil
			}
			return CastTakeOverTypeModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeTakeOverTypeModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOverTypeModelsFromDict(
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

func (p DescribeTakeOverTypeModelsResult) Pointer() *DescribeTakeOverTypeModelsResult {
	return &p
}

type GetTakeOverTypeModelResult struct {
	Item     *TakeOverTypeModel   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetTakeOverTypeModelAsyncResult struct {
	result *GetTakeOverTypeModelResult
	err    error
}

func NewGetTakeOverTypeModelResultFromJson(data string) GetTakeOverTypeModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverTypeModelResultFromDict(dict)
}

func NewGetTakeOverTypeModelResultFromDict(data map[string]interface{}) GetTakeOverTypeModelResult {
	return GetTakeOverTypeModelResult{
		Item: func() *TakeOverTypeModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverTypeModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetTakeOverTypeModelResult) ToDict() map[string]interface{} {
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

func (p GetTakeOverTypeModelResult) Pointer() *GetTakeOverTypeModelResult {
	return &p
}

type DescribeTakeOverTypeModelMastersResult struct {
	Items         []TakeOverTypeModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeTakeOverTypeModelMastersAsyncResult struct {
	result *DescribeTakeOverTypeModelMastersResult
	err    error
}

func NewDescribeTakeOverTypeModelMastersResultFromJson(data string) DescribeTakeOverTypeModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOverTypeModelMastersResultFromDict(dict)
}

func NewDescribeTakeOverTypeModelMastersResultFromDict(data map[string]interface{}) DescribeTakeOverTypeModelMastersResult {
	return DescribeTakeOverTypeModelMastersResult{
		Items: func() []TakeOverTypeModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastTakeOverTypeModelMasters(core.CastArray(data["items"]))
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

func (p DescribeTakeOverTypeModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOverTypeModelMastersFromDict(
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

func (p DescribeTakeOverTypeModelMastersResult) Pointer() *DescribeTakeOverTypeModelMastersResult {
	return &p
}

type CreateTakeOverTypeModelMasterResult struct {
	Item     *TakeOverTypeModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type CreateTakeOverTypeModelMasterAsyncResult struct {
	result *CreateTakeOverTypeModelMasterResult
	err    error
}

func NewCreateTakeOverTypeModelMasterResultFromJson(data string) CreateTakeOverTypeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverTypeModelMasterResultFromDict(dict)
}

func NewCreateTakeOverTypeModelMasterResultFromDict(data map[string]interface{}) CreateTakeOverTypeModelMasterResult {
	return CreateTakeOverTypeModelMasterResult{
		Item: func() *TakeOverTypeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverTypeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateTakeOverTypeModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateTakeOverTypeModelMasterResult) Pointer() *CreateTakeOverTypeModelMasterResult {
	return &p
}

type GetTakeOverTypeModelMasterResult struct {
	Item     *TakeOverTypeModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetTakeOverTypeModelMasterAsyncResult struct {
	result *GetTakeOverTypeModelMasterResult
	err    error
}

func NewGetTakeOverTypeModelMasterResultFromJson(data string) GetTakeOverTypeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverTypeModelMasterResultFromDict(dict)
}

func NewGetTakeOverTypeModelMasterResultFromDict(data map[string]interface{}) GetTakeOverTypeModelMasterResult {
	return GetTakeOverTypeModelMasterResult{
		Item: func() *TakeOverTypeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverTypeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetTakeOverTypeModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetTakeOverTypeModelMasterResult) Pointer() *GetTakeOverTypeModelMasterResult {
	return &p
}

type UpdateTakeOverTypeModelMasterResult struct {
	Item     *TakeOverTypeModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type UpdateTakeOverTypeModelMasterAsyncResult struct {
	result *UpdateTakeOverTypeModelMasterResult
	err    error
}

func NewUpdateTakeOverTypeModelMasterResultFromJson(data string) UpdateTakeOverTypeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverTypeModelMasterResultFromDict(dict)
}

func NewUpdateTakeOverTypeModelMasterResultFromDict(data map[string]interface{}) UpdateTakeOverTypeModelMasterResult {
	return UpdateTakeOverTypeModelMasterResult{
		Item: func() *TakeOverTypeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverTypeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateTakeOverTypeModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateTakeOverTypeModelMasterResult) Pointer() *UpdateTakeOverTypeModelMasterResult {
	return &p
}

type DeleteTakeOverTypeModelMasterResult struct {
	Item     *TakeOverTypeModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type DeleteTakeOverTypeModelMasterAsyncResult struct {
	result *DeleteTakeOverTypeModelMasterResult
	err    error
}

func NewDeleteTakeOverTypeModelMasterResultFromJson(data string) DeleteTakeOverTypeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverTypeModelMasterResultFromDict(dict)
}

func NewDeleteTakeOverTypeModelMasterResultFromDict(data map[string]interface{}) DeleteTakeOverTypeModelMasterResult {
	return DeleteTakeOverTypeModelMasterResult{
		Item: func() *TakeOverTypeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewTakeOverTypeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteTakeOverTypeModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteTakeOverTypeModelMasterResult) Pointer() *DeleteTakeOverTypeModelMasterResult {
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
