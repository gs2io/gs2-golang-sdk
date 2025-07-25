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

package money2

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

type DescribeWalletsResult struct {
	Items         []Wallet             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeWalletsAsyncResult struct {
	result *DescribeWalletsResult
	err    error
}

func NewDescribeWalletsResultFromJson(data string) DescribeWalletsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsResultFromDict(dict)
}

func NewDescribeWalletsResultFromDict(data map[string]interface{}) DescribeWalletsResult {
	return DescribeWalletsResult{
		Items: func() []Wallet {
			if data["items"] == nil {
				return nil
			}
			return CastWallets(core.CastArray(data["items"]))
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

func (p DescribeWalletsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastWalletsFromDict(
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

func (p DescribeWalletsResult) Pointer() *DescribeWalletsResult {
	return &p
}

type DescribeWalletsByUserIdResult struct {
	Items         []Wallet             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeWalletsByUserIdAsyncResult struct {
	result *DescribeWalletsByUserIdResult
	err    error
}

func NewDescribeWalletsByUserIdResultFromJson(data string) DescribeWalletsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsByUserIdResultFromDict(dict)
}

func NewDescribeWalletsByUserIdResultFromDict(data map[string]interface{}) DescribeWalletsByUserIdResult {
	return DescribeWalletsByUserIdResult{
		Items: func() []Wallet {
			if data["items"] == nil {
				return nil
			}
			return CastWallets(core.CastArray(data["items"]))
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

func (p DescribeWalletsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastWalletsFromDict(
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

func (p DescribeWalletsByUserIdResult) Pointer() *DescribeWalletsByUserIdResult {
	return &p
}

type GetWalletResult struct {
	Item     *Wallet              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetWalletAsyncResult struct {
	result *GetWalletResult
	err    error
}

func NewGetWalletResultFromJson(data string) GetWalletResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletResultFromDict(dict)
}

func NewGetWalletResultFromDict(data map[string]interface{}) GetWalletResult {
	return GetWalletResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetWalletResult) ToDict() map[string]interface{} {
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

func (p GetWalletResult) Pointer() *GetWalletResult {
	return &p
}

type GetWalletByUserIdResult struct {
	Item     *Wallet              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetWalletByUserIdAsyncResult struct {
	result *GetWalletByUserIdResult
	err    error
}

func NewGetWalletByUserIdResultFromJson(data string) GetWalletByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletByUserIdResultFromDict(dict)
}

func NewGetWalletByUserIdResultFromDict(data map[string]interface{}) GetWalletByUserIdResult {
	return GetWalletByUserIdResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetWalletByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetWalletByUserIdResult) Pointer() *GetWalletByUserIdResult {
	return &p
}

type DepositByUserIdResult struct {
	Item     *Wallet              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DepositByUserIdAsyncResult struct {
	result *DepositByUserIdResult
	err    error
}

func NewDepositByUserIdResultFromJson(data string) DepositByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByUserIdResultFromDict(dict)
}

func NewDepositByUserIdResultFromDict(data map[string]interface{}) DepositByUserIdResult {
	return DepositByUserIdResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DepositByUserIdResult) ToDict() map[string]interface{} {
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

func (p DepositByUserIdResult) Pointer() *DepositByUserIdResult {
	return &p
}

type WithdrawResult struct {
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
	Metadata             *core.ResultMetadata `json:"metadata"`
}

type WithdrawAsyncResult struct {
	result *WithdrawResult
	err    error
}

func NewWithdrawResultFromJson(data string) WithdrawResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawResultFromDict(dict)
}

func NewWithdrawResultFromDict(data map[string]interface{}) WithdrawResult {
	return WithdrawResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		WithdrawTransactions: func() []DepositTransaction {
			if data["withdrawTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["withdrawTransactions"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p WithdrawResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WithdrawResult) Pointer() *WithdrawResult {
	return &p
}

type WithdrawByUserIdResult struct {
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
	Metadata             *core.ResultMetadata `json:"metadata"`
}

type WithdrawByUserIdAsyncResult struct {
	result *WithdrawByUserIdResult
	err    error
}

func NewWithdrawByUserIdResultFromJson(data string) WithdrawByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByUserIdResultFromDict(dict)
}

func NewWithdrawByUserIdResultFromDict(data map[string]interface{}) WithdrawByUserIdResult {
	return WithdrawByUserIdResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		WithdrawTransactions: func() []DepositTransaction {
			if data["withdrawTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["withdrawTransactions"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p WithdrawByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p WithdrawByUserIdResult) Pointer() *WithdrawByUserIdResult {
	return &p
}

type DepositByStampSheetResult struct {
	Item     *Wallet              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DepositByStampSheetAsyncResult struct {
	result *DepositByStampSheetResult
	err    error
}

func NewDepositByStampSheetResultFromJson(data string) DepositByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByStampSheetResultFromDict(dict)
}

func NewDepositByStampSheetResultFromDict(data map[string]interface{}) DepositByStampSheetResult {
	return DepositByStampSheetResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DepositByStampSheetResult) ToDict() map[string]interface{} {
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

func (p DepositByStampSheetResult) Pointer() *DepositByStampSheetResult {
	return &p
}

type WithdrawByStampTaskResult struct {
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
	NewContextStack      *string              `json:"newContextStack"`
	Metadata             *core.ResultMetadata `json:"metadata"`
}

type WithdrawByStampTaskAsyncResult struct {
	result *WithdrawByStampTaskResult
	err    error
}

func NewWithdrawByStampTaskResultFromJson(data string) WithdrawByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByStampTaskResultFromDict(dict)
}

func NewWithdrawByStampTaskResultFromDict(data map[string]interface{}) WithdrawByStampTaskResult {
	return WithdrawByStampTaskResult{
		Item: func() *Wallet {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewWalletFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		WithdrawTransactions: func() []DepositTransaction {
			if data["withdrawTransactions"] == nil {
				return nil
			}
			return CastDepositTransactions(core.CastArray(data["withdrawTransactions"]))
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

func (p WithdrawByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
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

func (p WithdrawByStampTaskResult) Pointer() *WithdrawByStampTaskResult {
	return &p
}

type DescribeEventsByUserIdResult struct {
	Items         []Event              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeEventsByUserIdAsyncResult struct {
	result *DescribeEventsByUserIdResult
	err    error
}

func NewDescribeEventsByUserIdResultFromJson(data string) DescribeEventsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsByUserIdResultFromDict(dict)
}

func NewDescribeEventsByUserIdResultFromDict(data map[string]interface{}) DescribeEventsByUserIdResult {
	return DescribeEventsByUserIdResult{
		Items: func() []Event {
			if data["items"] == nil {
				return nil
			}
			return CastEvents(core.CastArray(data["items"]))
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

func (p DescribeEventsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
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

func (p DescribeEventsByUserIdResult) Pointer() *DescribeEventsByUserIdResult {
	return &p
}

type GetEventByTransactionIdResult struct {
	Item     *Event               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetEventByTransactionIdAsyncResult struct {
	result *GetEventByTransactionIdResult
	err    error
}

func NewGetEventByTransactionIdResultFromJson(data string) GetEventByTransactionIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventByTransactionIdResultFromDict(dict)
}

func NewGetEventByTransactionIdResultFromDict(data map[string]interface{}) GetEventByTransactionIdResult {
	return GetEventByTransactionIdResult{
		Item: func() *Event {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEventFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetEventByTransactionIdResult) ToDict() map[string]interface{} {
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

func (p GetEventByTransactionIdResult) Pointer() *GetEventByTransactionIdResult {
	return &p
}

type VerifyReceiptResult struct {
	Item     *Event               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyReceiptAsyncResult struct {
	result *VerifyReceiptResult
	err    error
}

func NewVerifyReceiptResultFromJson(data string) VerifyReceiptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptResultFromDict(dict)
}

func NewVerifyReceiptResultFromDict(data map[string]interface{}) VerifyReceiptResult {
	return VerifyReceiptResult{
		Item: func() *Event {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEventFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyReceiptResult) ToDict() map[string]interface{} {
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

func (p VerifyReceiptResult) Pointer() *VerifyReceiptResult {
	return &p
}

type VerifyReceiptByUserIdResult struct {
	Item     *Event               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyReceiptByUserIdAsyncResult struct {
	result *VerifyReceiptByUserIdResult
	err    error
}

func NewVerifyReceiptByUserIdResultFromJson(data string) VerifyReceiptByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptByUserIdResultFromDict(dict)
}

func NewVerifyReceiptByUserIdResultFromDict(data map[string]interface{}) VerifyReceiptByUserIdResult {
	return VerifyReceiptByUserIdResult{
		Item: func() *Event {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEventFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyReceiptByUserIdResult) ToDict() map[string]interface{} {
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

func (p VerifyReceiptByUserIdResult) Pointer() *VerifyReceiptByUserIdResult {
	return &p
}

type VerifyReceiptByStampTaskResult struct {
	Item            *Event               `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyReceiptByStampTaskAsyncResult struct {
	result *VerifyReceiptByStampTaskResult
	err    error
}

func NewVerifyReceiptByStampTaskResultFromJson(data string) VerifyReceiptByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptByStampTaskResultFromDict(dict)
}

func NewVerifyReceiptByStampTaskResultFromDict(data map[string]interface{}) VerifyReceiptByStampTaskResult {
	return VerifyReceiptByStampTaskResult{
		Item: func() *Event {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewEventFromDict(core.CastMap(data["item"])).Pointer()
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

func (p VerifyReceiptByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyReceiptByStampTaskResult) Pointer() *VerifyReceiptByStampTaskResult {
	return &p
}

type DescribeSubscriptionStatusesResult struct {
	Items    []SubscriptionStatus `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscriptionStatusesAsyncResult struct {
	result *DescribeSubscriptionStatusesResult
	err    error
}

func NewDescribeSubscriptionStatusesResultFromJson(data string) DescribeSubscriptionStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscriptionStatusesResultFromDict(dict)
}

func NewDescribeSubscriptionStatusesResultFromDict(data map[string]interface{}) DescribeSubscriptionStatusesResult {
	return DescribeSubscriptionStatusesResult{
		Items: func() []SubscriptionStatus {
			if data["items"] == nil {
				return nil
			}
			return CastSubscriptionStatuses(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeSubscriptionStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscriptionStatusesFromDict(
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

func (p DescribeSubscriptionStatusesResult) Pointer() *DescribeSubscriptionStatusesResult {
	return &p
}

type DescribeSubscriptionStatusesByUserIdResult struct {
	Items    []SubscriptionStatus `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscriptionStatusesByUserIdAsyncResult struct {
	result *DescribeSubscriptionStatusesByUserIdResult
	err    error
}

func NewDescribeSubscriptionStatusesByUserIdResultFromJson(data string) DescribeSubscriptionStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscriptionStatusesByUserIdResultFromDict(dict)
}

func NewDescribeSubscriptionStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeSubscriptionStatusesByUserIdResult {
	return DescribeSubscriptionStatusesByUserIdResult{
		Items: func() []SubscriptionStatus {
			if data["items"] == nil {
				return nil
			}
			return CastSubscriptionStatuses(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeSubscriptionStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscriptionStatusesFromDict(
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

func (p DescribeSubscriptionStatusesByUserIdResult) Pointer() *DescribeSubscriptionStatusesByUserIdResult {
	return &p
}

type GetSubscriptionStatusResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscriptionStatusAsyncResult struct {
	result *GetSubscriptionStatusResult
	err    error
}

func NewGetSubscriptionStatusResultFromJson(data string) GetSubscriptionStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscriptionStatusResultFromDict(dict)
}

func NewGetSubscriptionStatusResultFromDict(data map[string]interface{}) GetSubscriptionStatusResult {
	return GetSubscriptionStatusResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSubscriptionStatusResult) ToDict() map[string]interface{} {
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

func (p GetSubscriptionStatusResult) Pointer() *GetSubscriptionStatusResult {
	return &p
}

type GetSubscriptionStatusByUserIdResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscriptionStatusByUserIdAsyncResult struct {
	result *GetSubscriptionStatusByUserIdResult
	err    error
}

func NewGetSubscriptionStatusByUserIdResultFromJson(data string) GetSubscriptionStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscriptionStatusByUserIdResultFromDict(dict)
}

func NewGetSubscriptionStatusByUserIdResultFromDict(data map[string]interface{}) GetSubscriptionStatusByUserIdResult {
	return GetSubscriptionStatusByUserIdResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSubscriptionStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p GetSubscriptionStatusByUserIdResult) Pointer() *GetSubscriptionStatusByUserIdResult {
	return &p
}

type AllocateSubscriptionStatusResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AllocateSubscriptionStatusAsyncResult struct {
	result *AllocateSubscriptionStatusResult
	err    error
}

func NewAllocateSubscriptionStatusResultFromJson(data string) AllocateSubscriptionStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAllocateSubscriptionStatusResultFromDict(dict)
}

func NewAllocateSubscriptionStatusResultFromDict(data map[string]interface{}) AllocateSubscriptionStatusResult {
	return AllocateSubscriptionStatusResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AllocateSubscriptionStatusResult) ToDict() map[string]interface{} {
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

func (p AllocateSubscriptionStatusResult) Pointer() *AllocateSubscriptionStatusResult {
	return &p
}

type AllocateSubscriptionStatusByUserIdResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AllocateSubscriptionStatusByUserIdAsyncResult struct {
	result *AllocateSubscriptionStatusByUserIdResult
	err    error
}

func NewAllocateSubscriptionStatusByUserIdResultFromJson(data string) AllocateSubscriptionStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAllocateSubscriptionStatusByUserIdResultFromDict(dict)
}

func NewAllocateSubscriptionStatusByUserIdResultFromDict(data map[string]interface{}) AllocateSubscriptionStatusByUserIdResult {
	return AllocateSubscriptionStatusByUserIdResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AllocateSubscriptionStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p AllocateSubscriptionStatusByUserIdResult) Pointer() *AllocateSubscriptionStatusByUserIdResult {
	return &p
}

type TakeoverSubscriptionStatusResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type TakeoverSubscriptionStatusAsyncResult struct {
	result *TakeoverSubscriptionStatusResult
	err    error
}

func NewTakeoverSubscriptionStatusResultFromJson(data string) TakeoverSubscriptionStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTakeoverSubscriptionStatusResultFromDict(dict)
}

func NewTakeoverSubscriptionStatusResultFromDict(data map[string]interface{}) TakeoverSubscriptionStatusResult {
	return TakeoverSubscriptionStatusResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p TakeoverSubscriptionStatusResult) ToDict() map[string]interface{} {
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

func (p TakeoverSubscriptionStatusResult) Pointer() *TakeoverSubscriptionStatusResult {
	return &p
}

type TakeoverSubscriptionStatusByUserIdResult struct {
	Item     *SubscriptionStatus  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type TakeoverSubscriptionStatusByUserIdAsyncResult struct {
	result *TakeoverSubscriptionStatusByUserIdResult
	err    error
}

func NewTakeoverSubscriptionStatusByUserIdResultFromJson(data string) TakeoverSubscriptionStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewTakeoverSubscriptionStatusByUserIdResultFromDict(dict)
}

func NewTakeoverSubscriptionStatusByUserIdResultFromDict(data map[string]interface{}) TakeoverSubscriptionStatusByUserIdResult {
	return TakeoverSubscriptionStatusByUserIdResult{
		Item: func() *SubscriptionStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscriptionStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p TakeoverSubscriptionStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p TakeoverSubscriptionStatusByUserIdResult) Pointer() *TakeoverSubscriptionStatusByUserIdResult {
	return &p
}

type DescribeRefundHistoriesByUserIdResult struct {
	Items         []RefundHistory      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRefundHistoriesByUserIdAsyncResult struct {
	result *DescribeRefundHistoriesByUserIdResult
	err    error
}

func NewDescribeRefundHistoriesByUserIdResultFromJson(data string) DescribeRefundHistoriesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRefundHistoriesByUserIdResultFromDict(dict)
}

func NewDescribeRefundHistoriesByUserIdResultFromDict(data map[string]interface{}) DescribeRefundHistoriesByUserIdResult {
	return DescribeRefundHistoriesByUserIdResult{
		Items: func() []RefundHistory {
			if data["items"] == nil {
				return nil
			}
			return CastRefundHistories(core.CastArray(data["items"]))
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

func (p DescribeRefundHistoriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRefundHistoriesFromDict(
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

func (p DescribeRefundHistoriesByUserIdResult) Pointer() *DescribeRefundHistoriesByUserIdResult {
	return &p
}

type DescribeRefundHistoriesByDateResult struct {
	Items         []RefundHistory      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRefundHistoriesByDateAsyncResult struct {
	result *DescribeRefundHistoriesByDateResult
	err    error
}

func NewDescribeRefundHistoriesByDateResultFromJson(data string) DescribeRefundHistoriesByDateResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRefundHistoriesByDateResultFromDict(dict)
}

func NewDescribeRefundHistoriesByDateResultFromDict(data map[string]interface{}) DescribeRefundHistoriesByDateResult {
	return DescribeRefundHistoriesByDateResult{
		Items: func() []RefundHistory {
			if data["items"] == nil {
				return nil
			}
			return CastRefundHistories(core.CastArray(data["items"]))
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

func (p DescribeRefundHistoriesByDateResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRefundHistoriesFromDict(
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

func (p DescribeRefundHistoriesByDateResult) Pointer() *DescribeRefundHistoriesByDateResult {
	return &p
}

type GetRefundHistoryResult struct {
	Item     *RefundHistory       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRefundHistoryAsyncResult struct {
	result *GetRefundHistoryResult
	err    error
}

func NewGetRefundHistoryResultFromJson(data string) GetRefundHistoryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRefundHistoryResultFromDict(dict)
}

func NewGetRefundHistoryResultFromDict(data map[string]interface{}) GetRefundHistoryResult {
	return GetRefundHistoryResult{
		Item: func() *RefundHistory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRefundHistoryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRefundHistoryResult) ToDict() map[string]interface{} {
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

func (p GetRefundHistoryResult) Pointer() *GetRefundHistoryResult {
	return &p
}

type DescribeStoreContentModelsResult struct {
	Items    []StoreContentModel  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeStoreContentModelsAsyncResult struct {
	result *DescribeStoreContentModelsResult
	err    error
}

func NewDescribeStoreContentModelsResultFromJson(data string) DescribeStoreContentModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreContentModelsResultFromDict(dict)
}

func NewDescribeStoreContentModelsResultFromDict(data map[string]interface{}) DescribeStoreContentModelsResult {
	return DescribeStoreContentModelsResult{
		Items: func() []StoreContentModel {
			if data["items"] == nil {
				return nil
			}
			return CastStoreContentModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeStoreContentModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreContentModelsFromDict(
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

func (p DescribeStoreContentModelsResult) Pointer() *DescribeStoreContentModelsResult {
	return &p
}

type GetStoreContentModelResult struct {
	Item     *StoreContentModel   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStoreContentModelAsyncResult struct {
	result *GetStoreContentModelResult
	err    error
}

func NewGetStoreContentModelResultFromJson(data string) GetStoreContentModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreContentModelResultFromDict(dict)
}

func NewGetStoreContentModelResultFromDict(data map[string]interface{}) GetStoreContentModelResult {
	return GetStoreContentModelResult{
		Item: func() *StoreContentModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreContentModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStoreContentModelResult) ToDict() map[string]interface{} {
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

func (p GetStoreContentModelResult) Pointer() *GetStoreContentModelResult {
	return &p
}

type DescribeStoreContentModelMastersResult struct {
	Items         []StoreContentModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeStoreContentModelMastersAsyncResult struct {
	result *DescribeStoreContentModelMastersResult
	err    error
}

func NewDescribeStoreContentModelMastersResultFromJson(data string) DescribeStoreContentModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreContentModelMastersResultFromDict(dict)
}

func NewDescribeStoreContentModelMastersResultFromDict(data map[string]interface{}) DescribeStoreContentModelMastersResult {
	return DescribeStoreContentModelMastersResult{
		Items: func() []StoreContentModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastStoreContentModelMasters(core.CastArray(data["items"]))
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

func (p DescribeStoreContentModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreContentModelMastersFromDict(
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

func (p DescribeStoreContentModelMastersResult) Pointer() *DescribeStoreContentModelMastersResult {
	return &p
}

type CreateStoreContentModelMasterResult struct {
	Item     *StoreContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type CreateStoreContentModelMasterAsyncResult struct {
	result *CreateStoreContentModelMasterResult
	err    error
}

func NewCreateStoreContentModelMasterResultFromJson(data string) CreateStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStoreContentModelMasterResultFromDict(dict)
}

func NewCreateStoreContentModelMasterResultFromDict(data map[string]interface{}) CreateStoreContentModelMasterResult {
	return CreateStoreContentModelMasterResult{
		Item: func() *StoreContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateStoreContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateStoreContentModelMasterResult) Pointer() *CreateStoreContentModelMasterResult {
	return &p
}

type GetStoreContentModelMasterResult struct {
	Item     *StoreContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetStoreContentModelMasterAsyncResult struct {
	result *GetStoreContentModelMasterResult
	err    error
}

func NewGetStoreContentModelMasterResultFromJson(data string) GetStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreContentModelMasterResultFromDict(dict)
}

func NewGetStoreContentModelMasterResultFromDict(data map[string]interface{}) GetStoreContentModelMasterResult {
	return GetStoreContentModelMasterResult{
		Item: func() *StoreContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStoreContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetStoreContentModelMasterResult) Pointer() *GetStoreContentModelMasterResult {
	return &p
}

type UpdateStoreContentModelMasterResult struct {
	Item     *StoreContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type UpdateStoreContentModelMasterAsyncResult struct {
	result *UpdateStoreContentModelMasterResult
	err    error
}

func NewUpdateStoreContentModelMasterResultFromJson(data string) UpdateStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStoreContentModelMasterResultFromDict(dict)
}

func NewUpdateStoreContentModelMasterResultFromDict(data map[string]interface{}) UpdateStoreContentModelMasterResult {
	return UpdateStoreContentModelMasterResult{
		Item: func() *StoreContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateStoreContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateStoreContentModelMasterResult) Pointer() *UpdateStoreContentModelMasterResult {
	return &p
}

type DeleteStoreContentModelMasterResult struct {
	Item     *StoreContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type DeleteStoreContentModelMasterAsyncResult struct {
	result *DeleteStoreContentModelMasterResult
	err    error
}

func NewDeleteStoreContentModelMasterResultFromJson(data string) DeleteStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStoreContentModelMasterResultFromDict(dict)
}

func NewDeleteStoreContentModelMasterResultFromDict(data map[string]interface{}) DeleteStoreContentModelMasterResult {
	return DeleteStoreContentModelMasterResult{
		Item: func() *StoreContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteStoreContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteStoreContentModelMasterResult) Pointer() *DeleteStoreContentModelMasterResult {
	return &p
}

type DescribeStoreSubscriptionContentModelsResult struct {
	Items    []StoreSubscriptionContentModel `json:"items"`
	Metadata *core.ResultMetadata            `json:"metadata"`
}

type DescribeStoreSubscriptionContentModelsAsyncResult struct {
	result *DescribeStoreSubscriptionContentModelsResult
	err    error
}

func NewDescribeStoreSubscriptionContentModelsResultFromJson(data string) DescribeStoreSubscriptionContentModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreSubscriptionContentModelsResultFromDict(dict)
}

func NewDescribeStoreSubscriptionContentModelsResultFromDict(data map[string]interface{}) DescribeStoreSubscriptionContentModelsResult {
	return DescribeStoreSubscriptionContentModelsResult{
		Items: func() []StoreSubscriptionContentModel {
			if data["items"] == nil {
				return nil
			}
			return CastStoreSubscriptionContentModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeStoreSubscriptionContentModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreSubscriptionContentModelsFromDict(
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

func (p DescribeStoreSubscriptionContentModelsResult) Pointer() *DescribeStoreSubscriptionContentModelsResult {
	return &p
}

type GetStoreSubscriptionContentModelResult struct {
	Item     *StoreSubscriptionContentModel `json:"item"`
	Metadata *core.ResultMetadata           `json:"metadata"`
}

type GetStoreSubscriptionContentModelAsyncResult struct {
	result *GetStoreSubscriptionContentModelResult
	err    error
}

func NewGetStoreSubscriptionContentModelResultFromJson(data string) GetStoreSubscriptionContentModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreSubscriptionContentModelResultFromDict(dict)
}

func NewGetStoreSubscriptionContentModelResultFromDict(data map[string]interface{}) GetStoreSubscriptionContentModelResult {
	return GetStoreSubscriptionContentModelResult{
		Item: func() *StoreSubscriptionContentModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreSubscriptionContentModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStoreSubscriptionContentModelResult) ToDict() map[string]interface{} {
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

func (p GetStoreSubscriptionContentModelResult) Pointer() *GetStoreSubscriptionContentModelResult {
	return &p
}

type DescribeStoreSubscriptionContentModelMastersResult struct {
	Items         []StoreSubscriptionContentModelMaster `json:"items"`
	NextPageToken *string                               `json:"nextPageToken"`
	Metadata      *core.ResultMetadata                  `json:"metadata"`
}

type DescribeStoreSubscriptionContentModelMastersAsyncResult struct {
	result *DescribeStoreSubscriptionContentModelMastersResult
	err    error
}

func NewDescribeStoreSubscriptionContentModelMastersResultFromJson(data string) DescribeStoreSubscriptionContentModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreSubscriptionContentModelMastersResultFromDict(dict)
}

func NewDescribeStoreSubscriptionContentModelMastersResultFromDict(data map[string]interface{}) DescribeStoreSubscriptionContentModelMastersResult {
	return DescribeStoreSubscriptionContentModelMastersResult{
		Items: func() []StoreSubscriptionContentModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastStoreSubscriptionContentModelMasters(core.CastArray(data["items"]))
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

func (p DescribeStoreSubscriptionContentModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreSubscriptionContentModelMastersFromDict(
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

func (p DescribeStoreSubscriptionContentModelMastersResult) Pointer() *DescribeStoreSubscriptionContentModelMastersResult {
	return &p
}

type CreateStoreSubscriptionContentModelMasterResult struct {
	Item     *StoreSubscriptionContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata                 `json:"metadata"`
}

type CreateStoreSubscriptionContentModelMasterAsyncResult struct {
	result *CreateStoreSubscriptionContentModelMasterResult
	err    error
}

func NewCreateStoreSubscriptionContentModelMasterResultFromJson(data string) CreateStoreSubscriptionContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStoreSubscriptionContentModelMasterResultFromDict(dict)
}

func NewCreateStoreSubscriptionContentModelMasterResultFromDict(data map[string]interface{}) CreateStoreSubscriptionContentModelMasterResult {
	return CreateStoreSubscriptionContentModelMasterResult{
		Item: func() *StoreSubscriptionContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreSubscriptionContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateStoreSubscriptionContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p CreateStoreSubscriptionContentModelMasterResult) Pointer() *CreateStoreSubscriptionContentModelMasterResult {
	return &p
}

type GetStoreSubscriptionContentModelMasterResult struct {
	Item     *StoreSubscriptionContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata                 `json:"metadata"`
}

type GetStoreSubscriptionContentModelMasterAsyncResult struct {
	result *GetStoreSubscriptionContentModelMasterResult
	err    error
}

func NewGetStoreSubscriptionContentModelMasterResultFromJson(data string) GetStoreSubscriptionContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreSubscriptionContentModelMasterResultFromDict(dict)
}

func NewGetStoreSubscriptionContentModelMasterResultFromDict(data map[string]interface{}) GetStoreSubscriptionContentModelMasterResult {
	return GetStoreSubscriptionContentModelMasterResult{
		Item: func() *StoreSubscriptionContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreSubscriptionContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStoreSubscriptionContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p GetStoreSubscriptionContentModelMasterResult) Pointer() *GetStoreSubscriptionContentModelMasterResult {
	return &p
}

type UpdateStoreSubscriptionContentModelMasterResult struct {
	Item     *StoreSubscriptionContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata                 `json:"metadata"`
}

type UpdateStoreSubscriptionContentModelMasterAsyncResult struct {
	result *UpdateStoreSubscriptionContentModelMasterResult
	err    error
}

func NewUpdateStoreSubscriptionContentModelMasterResultFromJson(data string) UpdateStoreSubscriptionContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStoreSubscriptionContentModelMasterResultFromDict(dict)
}

func NewUpdateStoreSubscriptionContentModelMasterResultFromDict(data map[string]interface{}) UpdateStoreSubscriptionContentModelMasterResult {
	return UpdateStoreSubscriptionContentModelMasterResult{
		Item: func() *StoreSubscriptionContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreSubscriptionContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateStoreSubscriptionContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p UpdateStoreSubscriptionContentModelMasterResult) Pointer() *UpdateStoreSubscriptionContentModelMasterResult {
	return &p
}

type DeleteStoreSubscriptionContentModelMasterResult struct {
	Item     *StoreSubscriptionContentModelMaster `json:"item"`
	Metadata *core.ResultMetadata                 `json:"metadata"`
}

type DeleteStoreSubscriptionContentModelMasterAsyncResult struct {
	result *DeleteStoreSubscriptionContentModelMasterResult
	err    error
}

func NewDeleteStoreSubscriptionContentModelMasterResultFromJson(data string) DeleteStoreSubscriptionContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStoreSubscriptionContentModelMasterResultFromDict(dict)
}

func NewDeleteStoreSubscriptionContentModelMasterResultFromDict(data map[string]interface{}) DeleteStoreSubscriptionContentModelMasterResult {
	return DeleteStoreSubscriptionContentModelMasterResult{
		Item: func() *StoreSubscriptionContentModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStoreSubscriptionContentModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteStoreSubscriptionContentModelMasterResult) ToDict() map[string]interface{} {
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

func (p DeleteStoreSubscriptionContentModelMasterResult) Pointer() *DeleteStoreSubscriptionContentModelMasterResult {
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

type DescribeDailyTransactionHistoriesByCurrencyResult struct {
	Items         []DailyTransactionHistory `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeDailyTransactionHistoriesByCurrencyAsyncResult struct {
	result *DescribeDailyTransactionHistoriesByCurrencyResult
	err    error
}

func NewDescribeDailyTransactionHistoriesByCurrencyResultFromJson(data string) DescribeDailyTransactionHistoriesByCurrencyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDailyTransactionHistoriesByCurrencyResultFromDict(dict)
}

func NewDescribeDailyTransactionHistoriesByCurrencyResultFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesByCurrencyResult {
	return DescribeDailyTransactionHistoriesByCurrencyResult{
		Items: func() []DailyTransactionHistory {
			if data["items"] == nil {
				return nil
			}
			return CastDailyTransactionHistories(core.CastArray(data["items"]))
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

func (p DescribeDailyTransactionHistoriesByCurrencyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDailyTransactionHistoriesFromDict(
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

func (p DescribeDailyTransactionHistoriesByCurrencyResult) Pointer() *DescribeDailyTransactionHistoriesByCurrencyResult {
	return &p
}

type DescribeDailyTransactionHistoriesResult struct {
	Items         []DailyTransactionHistory `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeDailyTransactionHistoriesAsyncResult struct {
	result *DescribeDailyTransactionHistoriesResult
	err    error
}

func NewDescribeDailyTransactionHistoriesResultFromJson(data string) DescribeDailyTransactionHistoriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDailyTransactionHistoriesResultFromDict(dict)
}

func NewDescribeDailyTransactionHistoriesResultFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesResult {
	return DescribeDailyTransactionHistoriesResult{
		Items: func() []DailyTransactionHistory {
			if data["items"] == nil {
				return nil
			}
			return CastDailyTransactionHistories(core.CastArray(data["items"]))
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

func (p DescribeDailyTransactionHistoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDailyTransactionHistoriesFromDict(
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

func (p DescribeDailyTransactionHistoriesResult) Pointer() *DescribeDailyTransactionHistoriesResult {
	return &p
}

type GetDailyTransactionHistoryResult struct {
	Item     *DailyTransactionHistory `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetDailyTransactionHistoryAsyncResult struct {
	result *GetDailyTransactionHistoryResult
	err    error
}

func NewGetDailyTransactionHistoryResultFromJson(data string) GetDailyTransactionHistoryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDailyTransactionHistoryResultFromDict(dict)
}

func NewGetDailyTransactionHistoryResultFromDict(data map[string]interface{}) GetDailyTransactionHistoryResult {
	return GetDailyTransactionHistoryResult{
		Item: func() *DailyTransactionHistory {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewDailyTransactionHistoryFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetDailyTransactionHistoryResult) ToDict() map[string]interface{} {
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

func (p GetDailyTransactionHistoryResult) Pointer() *GetDailyTransactionHistoryResult {
	return &p
}

type DescribeUnusedBalancesResult struct {
	Items         []UnusedBalance      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeUnusedBalancesAsyncResult struct {
	result *DescribeUnusedBalancesResult
	err    error
}

func NewDescribeUnusedBalancesResultFromJson(data string) DescribeUnusedBalancesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeUnusedBalancesResultFromDict(dict)
}

func NewDescribeUnusedBalancesResultFromDict(data map[string]interface{}) DescribeUnusedBalancesResult {
	return DescribeUnusedBalancesResult{
		Items: func() []UnusedBalance {
			if data["items"] == nil {
				return nil
			}
			return CastUnusedBalances(core.CastArray(data["items"]))
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

func (p DescribeUnusedBalancesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastUnusedBalancesFromDict(
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

func (p DescribeUnusedBalancesResult) Pointer() *DescribeUnusedBalancesResult {
	return &p
}

type GetUnusedBalanceResult struct {
	Item     *UnusedBalance       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetUnusedBalanceAsyncResult struct {
	result *GetUnusedBalanceResult
	err    error
}

func NewGetUnusedBalanceResultFromJson(data string) GetUnusedBalanceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetUnusedBalanceResultFromDict(dict)
}

func NewGetUnusedBalanceResultFromDict(data map[string]interface{}) GetUnusedBalanceResult {
	return GetUnusedBalanceResult{
		Item: func() *UnusedBalance {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnusedBalanceFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetUnusedBalanceResult) ToDict() map[string]interface{} {
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

func (p GetUnusedBalanceResult) Pointer() *GetUnusedBalanceResult {
	return &p
}
