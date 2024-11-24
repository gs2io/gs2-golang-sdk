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

package friend

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

type GetProfileResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetProfileAsyncResult struct {
	result *GetProfileResult
	err    error
}

func NewGetProfileResultFromJson(data string) GetProfileResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProfileResultFromDict(dict)
}

func NewGetProfileResultFromDict(data map[string]interface{}) GetProfileResult {
	return GetProfileResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetProfileResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetProfileResult) Pointer() *GetProfileResult {
	return &p
}

type GetProfileByUserIdResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetProfileByUserIdAsyncResult struct {
	result *GetProfileByUserIdResult
	err    error
}

func NewGetProfileByUserIdResultFromJson(data string) GetProfileByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProfileByUserIdResultFromDict(dict)
}

func NewGetProfileByUserIdResultFromDict(data map[string]interface{}) GetProfileByUserIdResult {
	return GetProfileByUserIdResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetProfileByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetProfileByUserIdResult) Pointer() *GetProfileByUserIdResult {
	return &p
}

type UpdateProfileResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateProfileAsyncResult struct {
	result *UpdateProfileResult
	err    error
}

func NewUpdateProfileResultFromJson(data string) UpdateProfileResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateProfileResultFromDict(dict)
}

func NewUpdateProfileResultFromDict(data map[string]interface{}) UpdateProfileResult {
	return UpdateProfileResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateProfileResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateProfileResult) Pointer() *UpdateProfileResult {
	return &p
}

type UpdateProfileByUserIdResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateProfileByUserIdAsyncResult struct {
	result *UpdateProfileByUserIdResult
	err    error
}

func NewUpdateProfileByUserIdResultFromJson(data string) UpdateProfileByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateProfileByUserIdResultFromDict(dict)
}

func NewUpdateProfileByUserIdResultFromDict(data map[string]interface{}) UpdateProfileByUserIdResult {
	return UpdateProfileByUserIdResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateProfileByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateProfileByUserIdResult) Pointer() *UpdateProfileByUserIdResult {
	return &p
}

type DeleteProfileByUserIdResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteProfileByUserIdAsyncResult struct {
	result *DeleteProfileByUserIdResult
	err    error
}

func NewDeleteProfileByUserIdResultFromJson(data string) DeleteProfileByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProfileByUserIdResultFromDict(dict)
}

func NewDeleteProfileByUserIdResultFromDict(data map[string]interface{}) DeleteProfileByUserIdResult {
	return DeleteProfileByUserIdResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteProfileByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteProfileByUserIdResult) Pointer() *DeleteProfileByUserIdResult {
	return &p
}

type UpdateProfileByStampSheetResult struct {
	Item     *Profile             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateProfileByStampSheetAsyncResult struct {
	result *UpdateProfileByStampSheetResult
	err    error
}

func NewUpdateProfileByStampSheetResultFromJson(data string) UpdateProfileByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateProfileByStampSheetResultFromDict(dict)
}

func NewUpdateProfileByStampSheetResultFromDict(data map[string]interface{}) UpdateProfileByStampSheetResult {
	return UpdateProfileByStampSheetResult{
		Item: func() *Profile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateProfileByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateProfileByStampSheetResult) Pointer() *UpdateProfileByStampSheetResult {
	return &p
}

type DescribeFriendsResult struct {
	Items         []FriendUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeFriendsAsyncResult struct {
	result *DescribeFriendsResult
	err    error
}

func NewDescribeFriendsResultFromJson(data string) DescribeFriendsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFriendsResultFromDict(dict)
}

func NewDescribeFriendsResultFromDict(data map[string]interface{}) DescribeFriendsResult {
	return DescribeFriendsResult{
		Items: func() []FriendUser {
			if data["items"] == nil {
				return nil
			}
			return CastFriendUsers(core.CastArray(data["items"]))
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

func (p DescribeFriendsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFriendUsersFromDict(
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

func (p DescribeFriendsResult) Pointer() *DescribeFriendsResult {
	return &p
}

type DescribeFriendsByUserIdResult struct {
	Items         []FriendUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeFriendsByUserIdAsyncResult struct {
	result *DescribeFriendsByUserIdResult
	err    error
}

func NewDescribeFriendsByUserIdResultFromJson(data string) DescribeFriendsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFriendsByUserIdResultFromDict(dict)
}

func NewDescribeFriendsByUserIdResultFromDict(data map[string]interface{}) DescribeFriendsByUserIdResult {
	return DescribeFriendsByUserIdResult{
		Items: func() []FriendUser {
			if data["items"] == nil {
				return nil
			}
			return CastFriendUsers(core.CastArray(data["items"]))
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

func (p DescribeFriendsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFriendUsersFromDict(
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

func (p DescribeFriendsByUserIdResult) Pointer() *DescribeFriendsByUserIdResult {
	return &p
}

type DescribeBlackListResult struct {
	Items         []*string            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBlackListAsyncResult struct {
	result *DescribeBlackListResult
	err    error
}

func NewDescribeBlackListResultFromJson(data string) DescribeBlackListResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBlackListResultFromDict(dict)
}

func NewDescribeBlackListResultFromDict(data map[string]interface{}) DescribeBlackListResult {
	return DescribeBlackListResult{
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p DescribeBlackListResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
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

func (p DescribeBlackListResult) Pointer() *DescribeBlackListResult {
	return &p
}

type DescribeBlackListByUserIdResult struct {
	Items         []*string            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBlackListByUserIdAsyncResult struct {
	result *DescribeBlackListByUserIdResult
	err    error
}

func NewDescribeBlackListByUserIdResultFromJson(data string) DescribeBlackListByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBlackListByUserIdResultFromDict(dict)
}

func NewDescribeBlackListByUserIdResultFromDict(data map[string]interface{}) DescribeBlackListByUserIdResult {
	return DescribeBlackListByUserIdResult{
		Items: func() []*string {
			v, ok := data["items"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
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

func (p DescribeBlackListByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": core.CastStringsFromDict(
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

func (p DescribeBlackListByUserIdResult) Pointer() *DescribeBlackListByUserIdResult {
	return &p
}

type RegisterBlackListResult struct {
	Item     *BlackList           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RegisterBlackListAsyncResult struct {
	result *RegisterBlackListResult
	err    error
}

func NewRegisterBlackListResultFromJson(data string) RegisterBlackListResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRegisterBlackListResultFromDict(dict)
}

func NewRegisterBlackListResultFromDict(data map[string]interface{}) RegisterBlackListResult {
	return RegisterBlackListResult{
		Item: func() *BlackList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBlackListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RegisterBlackListResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RegisterBlackListResult) Pointer() *RegisterBlackListResult {
	return &p
}

type RegisterBlackListByUserIdResult struct {
	Item     *BlackList           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RegisterBlackListByUserIdAsyncResult struct {
	result *RegisterBlackListByUserIdResult
	err    error
}

func NewRegisterBlackListByUserIdResultFromJson(data string) RegisterBlackListByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRegisterBlackListByUserIdResultFromDict(dict)
}

func NewRegisterBlackListByUserIdResultFromDict(data map[string]interface{}) RegisterBlackListByUserIdResult {
	return RegisterBlackListByUserIdResult{
		Item: func() *BlackList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBlackListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RegisterBlackListByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RegisterBlackListByUserIdResult) Pointer() *RegisterBlackListByUserIdResult {
	return &p
}

type UnregisterBlackListResult struct {
	Item     *BlackList           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnregisterBlackListAsyncResult struct {
	result *UnregisterBlackListResult
	err    error
}

func NewUnregisterBlackListResultFromJson(data string) UnregisterBlackListResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnregisterBlackListResultFromDict(dict)
}

func NewUnregisterBlackListResultFromDict(data map[string]interface{}) UnregisterBlackListResult {
	return UnregisterBlackListResult{
		Item: func() *BlackList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBlackListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnregisterBlackListResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnregisterBlackListResult) Pointer() *UnregisterBlackListResult {
	return &p
}

type UnregisterBlackListByUserIdResult struct {
	Item     *BlackList           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnregisterBlackListByUserIdAsyncResult struct {
	result *UnregisterBlackListByUserIdResult
	err    error
}

func NewUnregisterBlackListByUserIdResultFromJson(data string) UnregisterBlackListByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnregisterBlackListByUserIdResultFromDict(dict)
}

func NewUnregisterBlackListByUserIdResultFromDict(data map[string]interface{}) UnregisterBlackListByUserIdResult {
	return UnregisterBlackListByUserIdResult{
		Item: func() *BlackList {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBlackListFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnregisterBlackListByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnregisterBlackListByUserIdResult) Pointer() *UnregisterBlackListByUserIdResult {
	return &p
}

type DescribeFollowsResult struct {
	Items         []FollowUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeFollowsAsyncResult struct {
	result *DescribeFollowsResult
	err    error
}

func NewDescribeFollowsResultFromJson(data string) DescribeFollowsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFollowsResultFromDict(dict)
}

func NewDescribeFollowsResultFromDict(data map[string]interface{}) DescribeFollowsResult {
	return DescribeFollowsResult{
		Items: func() []FollowUser {
			if data["items"] == nil {
				return nil
			}
			return CastFollowUsers(core.CastArray(data["items"]))
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

func (p DescribeFollowsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFollowUsersFromDict(
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

func (p DescribeFollowsResult) Pointer() *DescribeFollowsResult {
	return &p
}

type DescribeFollowsByUserIdResult struct {
	Items         []FollowUser         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeFollowsByUserIdAsyncResult struct {
	result *DescribeFollowsByUserIdResult
	err    error
}

func NewDescribeFollowsByUserIdResultFromJson(data string) DescribeFollowsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeFollowsByUserIdResultFromDict(dict)
}

func NewDescribeFollowsByUserIdResultFromDict(data map[string]interface{}) DescribeFollowsByUserIdResult {
	return DescribeFollowsByUserIdResult{
		Items: func() []FollowUser {
			if data["items"] == nil {
				return nil
			}
			return CastFollowUsers(core.CastArray(data["items"]))
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

func (p DescribeFollowsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFollowUsersFromDict(
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

func (p DescribeFollowsByUserIdResult) Pointer() *DescribeFollowsByUserIdResult {
	return &p
}

type GetFollowResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFollowAsyncResult struct {
	result *GetFollowResult
	err    error
}

func NewGetFollowResultFromJson(data string) GetFollowResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFollowResultFromDict(dict)
}

func NewGetFollowResultFromDict(data map[string]interface{}) GetFollowResult {
	return GetFollowResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetFollowResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetFollowResult) Pointer() *GetFollowResult {
	return &p
}

type GetFollowByUserIdResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFollowByUserIdAsyncResult struct {
	result *GetFollowByUserIdResult
	err    error
}

func NewGetFollowByUserIdResultFromJson(data string) GetFollowByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFollowByUserIdResultFromDict(dict)
}

func NewGetFollowByUserIdResultFromDict(data map[string]interface{}) GetFollowByUserIdResult {
	return GetFollowByUserIdResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetFollowByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetFollowByUserIdResult) Pointer() *GetFollowByUserIdResult {
	return &p
}

type FollowResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FollowAsyncResult struct {
	result *FollowResult
	err    error
}

func NewFollowResultFromJson(data string) FollowResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFollowResultFromDict(dict)
}

func NewFollowResultFromDict(data map[string]interface{}) FollowResult {
	return FollowResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p FollowResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p FollowResult) Pointer() *FollowResult {
	return &p
}

type FollowByUserIdResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type FollowByUserIdAsyncResult struct {
	result *FollowByUserIdResult
	err    error
}

func NewFollowByUserIdResultFromJson(data string) FollowByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewFollowByUserIdResultFromDict(dict)
}

func NewFollowByUserIdResultFromDict(data map[string]interface{}) FollowByUserIdResult {
	return FollowByUserIdResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p FollowByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p FollowByUserIdResult) Pointer() *FollowByUserIdResult {
	return &p
}

type UnfollowResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnfollowAsyncResult struct {
	result *UnfollowResult
	err    error
}

func NewUnfollowResultFromJson(data string) UnfollowResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnfollowResultFromDict(dict)
}

func NewUnfollowResultFromDict(data map[string]interface{}) UnfollowResult {
	return UnfollowResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnfollowResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnfollowResult) Pointer() *UnfollowResult {
	return &p
}

type UnfollowByUserIdResult struct {
	Item     *FollowUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UnfollowByUserIdAsyncResult struct {
	result *UnfollowByUserIdResult
	err    error
}

func NewUnfollowByUserIdResultFromJson(data string) UnfollowByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnfollowByUserIdResultFromDict(dict)
}

func NewUnfollowByUserIdResultFromDict(data map[string]interface{}) UnfollowByUserIdResult {
	return UnfollowByUserIdResult{
		Item: func() *FollowUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFollowUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UnfollowByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UnfollowByUserIdResult) Pointer() *UnfollowByUserIdResult {
	return &p
}

type GetFriendResult struct {
	Item     *FriendUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFriendAsyncResult struct {
	result *GetFriendResult
	err    error
}

func NewGetFriendResultFromJson(data string) GetFriendResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFriendResultFromDict(dict)
}

func NewGetFriendResultFromDict(data map[string]interface{}) GetFriendResult {
	return GetFriendResult{
		Item: func() *FriendUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetFriendResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetFriendResult) Pointer() *GetFriendResult {
	return &p
}

type GetFriendByUserIdResult struct {
	Item     *FriendUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetFriendByUserIdAsyncResult struct {
	result *GetFriendByUserIdResult
	err    error
}

func NewGetFriendByUserIdResultFromJson(data string) GetFriendByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetFriendByUserIdResultFromDict(dict)
}

func NewGetFriendByUserIdResultFromDict(data map[string]interface{}) GetFriendByUserIdResult {
	return GetFriendByUserIdResult{
		Item: func() *FriendUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetFriendByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetFriendByUserIdResult) Pointer() *GetFriendByUserIdResult {
	return &p
}

type DeleteFriendResult struct {
	Item     *FriendUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteFriendAsyncResult struct {
	result *DeleteFriendResult
	err    error
}

func NewDeleteFriendResultFromJson(data string) DeleteFriendResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFriendResultFromDict(dict)
}

func NewDeleteFriendResultFromDict(data map[string]interface{}) DeleteFriendResult {
	return DeleteFriendResult{
		Item: func() *FriendUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteFriendResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteFriendResult) Pointer() *DeleteFriendResult {
	return &p
}

type DeleteFriendByUserIdResult struct {
	Item     *FriendUser          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteFriendByUserIdAsyncResult struct {
	result *DeleteFriendByUserIdResult
	err    error
}

func NewDeleteFriendByUserIdResultFromJson(data string) DeleteFriendByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteFriendByUserIdResultFromDict(dict)
}

func NewDeleteFriendByUserIdResultFromDict(data map[string]interface{}) DeleteFriendByUserIdResult {
	return DeleteFriendByUserIdResult{
		Item: func() *FriendUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteFriendByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteFriendByUserIdResult) Pointer() *DeleteFriendByUserIdResult {
	return &p
}

type DescribeSendRequestsResult struct {
	Items         []FriendRequest      `json:"items"`
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
		Items: func() []FriendRequest {
			if data["items"] == nil {
				return nil
			}
			return CastFriendRequests(core.CastArray(data["items"]))
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
		"items": CastFriendRequestsFromDict(
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
	Items         []FriendRequest      `json:"items"`
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
		Items: func() []FriendRequest {
			if data["items"] == nil {
				return nil
			}
			return CastFriendRequests(core.CastArray(data["items"]))
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
		"items": CastFriendRequestsFromDict(
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
	Item     *FriendRequest       `json:"item"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *FriendRequest       `json:"item"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *FriendRequest       `json:"item"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
	Item     *FriendRequest       `json:"item"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type DescribeReceiveRequestsResult struct {
	Items         []FriendRequest      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
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
		Items: func() []FriendRequest {
			if data["items"] == nil {
				return nil
			}
			return CastFriendRequests(core.CastArray(data["items"]))
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
		"items": CastFriendRequestsFromDict(
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

type DescribeReceiveRequestsByUserIdResult struct {
	Items         []FriendRequest      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeReceiveRequestsByUserIdAsyncResult struct {
	result *DescribeReceiveRequestsByUserIdResult
	err    error
}

func NewDescribeReceiveRequestsByUserIdResultFromJson(data string) DescribeReceiveRequestsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveRequestsByUserIdResultFromDict(dict)
}

func NewDescribeReceiveRequestsByUserIdResultFromDict(data map[string]interface{}) DescribeReceiveRequestsByUserIdResult {
	return DescribeReceiveRequestsByUserIdResult{
		Items: func() []FriendRequest {
			if data["items"] == nil {
				return nil
			}
			return CastFriendRequests(core.CastArray(data["items"]))
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

func (p DescribeReceiveRequestsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastFriendRequestsFromDict(
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

func (p DescribeReceiveRequestsByUserIdResult) Pointer() *DescribeReceiveRequestsByUserIdResult {
	return &p
}

type GetReceiveRequestResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetReceiveRequestByUserIdResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetReceiveRequestByUserIdAsyncResult struct {
	result *GetReceiveRequestByUserIdResult
	err    error
}

func NewGetReceiveRequestByUserIdResultFromJson(data string) GetReceiveRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveRequestByUserIdResultFromDict(dict)
}

func NewGetReceiveRequestByUserIdResultFromDict(data map[string]interface{}) GetReceiveRequestByUserIdResult {
	return GetReceiveRequestByUserIdResult{
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetReceiveRequestByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetReceiveRequestByUserIdResult) Pointer() *GetReceiveRequestByUserIdResult {
	return &p
}

type AcceptRequestResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type AcceptRequestByUserIdResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AcceptRequestByUserIdAsyncResult struct {
	result *AcceptRequestByUserIdResult
	err    error
}

func NewAcceptRequestByUserIdResultFromJson(data string) AcceptRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestByUserIdResultFromDict(dict)
}

func NewAcceptRequestByUserIdResultFromDict(data map[string]interface{}) AcceptRequestByUserIdResult {
	return AcceptRequestByUserIdResult{
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AcceptRequestByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcceptRequestByUserIdResult) Pointer() *AcceptRequestByUserIdResult {
	return &p
}

type RejectRequestResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
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
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type RejectRequestByUserIdResult struct {
	Item     *FriendRequest       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type RejectRequestByUserIdAsyncResult struct {
	result *RejectRequestByUserIdResult
	err    error
}

func NewRejectRequestByUserIdResultFromJson(data string) RejectRequestByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRejectRequestByUserIdResultFromDict(dict)
}

func NewRejectRequestByUserIdResultFromDict(data map[string]interface{}) RejectRequestByUserIdResult {
	return RejectRequestByUserIdResult{
		Item: func() *FriendRequest {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RejectRequestByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RejectRequestByUserIdResult) Pointer() *RejectRequestByUserIdResult {
	return &p
}

type GetPublicProfileResult struct {
	Item     *PublicProfile       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPublicProfileAsyncResult struct {
	result *GetPublicProfileResult
	err    error
}

func NewGetPublicProfileResultFromJson(data string) GetPublicProfileResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPublicProfileResultFromDict(dict)
}

func NewGetPublicProfileResultFromDict(data map[string]interface{}) GetPublicProfileResult {
	return GetPublicProfileResult{
		Item: func() *PublicProfile {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPublicProfileFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetPublicProfileResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetPublicProfileResult) Pointer() *GetPublicProfileResult {
	return &p
}
