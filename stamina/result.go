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

package stamina

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

type DescribeStaminaModelMastersResult struct {
	Items         []StaminaModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStaminaModelMastersAsyncResult struct {
	result *DescribeStaminaModelMastersResult
	err    error
}

func NewDescribeStaminaModelMastersResultFromJson(data string) DescribeStaminaModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaModelMastersResultFromDict(dict)
}

func NewDescribeStaminaModelMastersResultFromDict(data map[string]interface{}) DescribeStaminaModelMastersResult {
	return DescribeStaminaModelMastersResult{
		Items: func() []StaminaModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastStaminaModelMasters(core.CastArray(data["items"]))
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

func (p DescribeStaminaModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaModelMastersFromDict(
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

func (p DescribeStaminaModelMastersResult) Pointer() *DescribeStaminaModelMastersResult {
	return &p
}

type CreateStaminaModelMasterResult struct {
	Item     *StaminaModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateStaminaModelMasterAsyncResult struct {
	result *CreateStaminaModelMasterResult
	err    error
}

func NewCreateStaminaModelMasterResultFromJson(data string) CreateStaminaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStaminaModelMasterResultFromDict(dict)
}

func NewCreateStaminaModelMasterResultFromDict(data map[string]interface{}) CreateStaminaModelMasterResult {
	return CreateStaminaModelMasterResult{
		Item: func() *StaminaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateStaminaModelMasterResult) Pointer() *CreateStaminaModelMasterResult {
	return &p
}

type GetStaminaModelMasterResult struct {
	Item     *StaminaModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStaminaModelMasterAsyncResult struct {
	result *GetStaminaModelMasterResult
	err    error
}

func NewGetStaminaModelMasterResultFromJson(data string) GetStaminaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaModelMasterResultFromDict(dict)
}

func NewGetStaminaModelMasterResultFromDict(data map[string]interface{}) GetStaminaModelMasterResult {
	return GetStaminaModelMasterResult{
		Item: func() *StaminaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetStaminaModelMasterResult) Pointer() *GetStaminaModelMasterResult {
	return &p
}

type UpdateStaminaModelMasterResult struct {
	Item     *StaminaModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateStaminaModelMasterAsyncResult struct {
	result *UpdateStaminaModelMasterResult
	err    error
}

func NewUpdateStaminaModelMasterResultFromJson(data string) UpdateStaminaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStaminaModelMasterResultFromDict(dict)
}

func NewUpdateStaminaModelMasterResultFromDict(data map[string]interface{}) UpdateStaminaModelMasterResult {
	return UpdateStaminaModelMasterResult{
		Item: func() *StaminaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateStaminaModelMasterResult) Pointer() *UpdateStaminaModelMasterResult {
	return &p
}

type DeleteStaminaModelMasterResult struct {
	Item     *StaminaModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStaminaModelMasterAsyncResult struct {
	result *DeleteStaminaModelMasterResult
	err    error
}

func NewDeleteStaminaModelMasterResultFromJson(data string) DeleteStaminaModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStaminaModelMasterResultFromDict(dict)
}

func NewDeleteStaminaModelMasterResultFromDict(data map[string]interface{}) DeleteStaminaModelMasterResult {
	return DeleteStaminaModelMasterResult{
		Item: func() *StaminaModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteStaminaModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteStaminaModelMasterResult) Pointer() *DeleteStaminaModelMasterResult {
	return &p
}

type DescribeMaxStaminaTableMastersResult struct {
	Items         []MaxStaminaTableMaster `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeMaxStaminaTableMastersAsyncResult struct {
	result *DescribeMaxStaminaTableMastersResult
	err    error
}

func NewDescribeMaxStaminaTableMastersResultFromJson(data string) DescribeMaxStaminaTableMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMaxStaminaTableMastersResultFromDict(dict)
}

func NewDescribeMaxStaminaTableMastersResultFromDict(data map[string]interface{}) DescribeMaxStaminaTableMastersResult {
	return DescribeMaxStaminaTableMastersResult{
		Items: func() []MaxStaminaTableMaster {
			if data["items"] == nil {
				return nil
			}
			return CastMaxStaminaTableMasters(core.CastArray(data["items"]))
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

func (p DescribeMaxStaminaTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastMaxStaminaTableMastersFromDict(
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

func (p DescribeMaxStaminaTableMastersResult) Pointer() *DescribeMaxStaminaTableMastersResult {
	return &p
}

type CreateMaxStaminaTableMasterResult struct {
	Item     *MaxStaminaTableMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type CreateMaxStaminaTableMasterAsyncResult struct {
	result *CreateMaxStaminaTableMasterResult
	err    error
}

func NewCreateMaxStaminaTableMasterResultFromJson(data string) CreateMaxStaminaTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMaxStaminaTableMasterResultFromDict(dict)
}

func NewCreateMaxStaminaTableMasterResultFromDict(data map[string]interface{}) CreateMaxStaminaTableMasterResult {
	return CreateMaxStaminaTableMasterResult{
		Item: func() *MaxStaminaTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateMaxStaminaTableMasterResult) Pointer() *CreateMaxStaminaTableMasterResult {
	return &p
}

type GetMaxStaminaTableMasterResult struct {
	Item     *MaxStaminaTableMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetMaxStaminaTableMasterAsyncResult struct {
	result *GetMaxStaminaTableMasterResult
	err    error
}

func NewGetMaxStaminaTableMasterResultFromJson(data string) GetMaxStaminaTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMaxStaminaTableMasterResultFromDict(dict)
}

func NewGetMaxStaminaTableMasterResultFromDict(data map[string]interface{}) GetMaxStaminaTableMasterResult {
	return GetMaxStaminaTableMasterResult{
		Item: func() *MaxStaminaTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetMaxStaminaTableMasterResult) Pointer() *GetMaxStaminaTableMasterResult {
	return &p
}

type UpdateMaxStaminaTableMasterResult struct {
	Item     *MaxStaminaTableMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateMaxStaminaTableMasterAsyncResult struct {
	result *UpdateMaxStaminaTableMasterResult
	err    error
}

func NewUpdateMaxStaminaTableMasterResultFromJson(data string) UpdateMaxStaminaTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMaxStaminaTableMasterResultFromDict(dict)
}

func NewUpdateMaxStaminaTableMasterResultFromDict(data map[string]interface{}) UpdateMaxStaminaTableMasterResult {
	return UpdateMaxStaminaTableMasterResult{
		Item: func() *MaxStaminaTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateMaxStaminaTableMasterResult) Pointer() *UpdateMaxStaminaTableMasterResult {
	return &p
}

type DeleteMaxStaminaTableMasterResult struct {
	Item     *MaxStaminaTableMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DeleteMaxStaminaTableMasterAsyncResult struct {
	result *DeleteMaxStaminaTableMasterResult
	err    error
}

func NewDeleteMaxStaminaTableMasterResultFromJson(data string) DeleteMaxStaminaTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMaxStaminaTableMasterResultFromDict(dict)
}

func NewDeleteMaxStaminaTableMasterResultFromDict(data map[string]interface{}) DeleteMaxStaminaTableMasterResult {
	return DeleteMaxStaminaTableMasterResult{
		Item: func() *MaxStaminaTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewMaxStaminaTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteMaxStaminaTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteMaxStaminaTableMasterResult) Pointer() *DeleteMaxStaminaTableMasterResult {
	return &p
}

type DescribeRecoverIntervalTableMastersResult struct {
	Items         []RecoverIntervalTableMaster `json:"items"`
	NextPageToken *string                      `json:"nextPageToken"`
	Metadata      *core.ResultMetadata         `json:"metadata"`
}

type DescribeRecoverIntervalTableMastersAsyncResult struct {
	result *DescribeRecoverIntervalTableMastersResult
	err    error
}

func NewDescribeRecoverIntervalTableMastersResultFromJson(data string) DescribeRecoverIntervalTableMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRecoverIntervalTableMastersResultFromDict(dict)
}

func NewDescribeRecoverIntervalTableMastersResultFromDict(data map[string]interface{}) DescribeRecoverIntervalTableMastersResult {
	return DescribeRecoverIntervalTableMastersResult{
		Items: func() []RecoverIntervalTableMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRecoverIntervalTableMasters(core.CastArray(data["items"]))
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

func (p DescribeRecoverIntervalTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRecoverIntervalTableMastersFromDict(
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

func (p DescribeRecoverIntervalTableMastersResult) Pointer() *DescribeRecoverIntervalTableMastersResult {
	return &p
}

type CreateRecoverIntervalTableMasterResult struct {
	Item     *RecoverIntervalTableMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type CreateRecoverIntervalTableMasterAsyncResult struct {
	result *CreateRecoverIntervalTableMasterResult
	err    error
}

func NewCreateRecoverIntervalTableMasterResultFromJson(data string) CreateRecoverIntervalTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRecoverIntervalTableMasterResultFromDict(dict)
}

func NewCreateRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) CreateRecoverIntervalTableMasterResult {
	return CreateRecoverIntervalTableMasterResult{
		Item: func() *RecoverIntervalTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateRecoverIntervalTableMasterResult) Pointer() *CreateRecoverIntervalTableMasterResult {
	return &p
}

type GetRecoverIntervalTableMasterResult struct {
	Item     *RecoverIntervalTableMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type GetRecoverIntervalTableMasterAsyncResult struct {
	result *GetRecoverIntervalTableMasterResult
	err    error
}

func NewGetRecoverIntervalTableMasterResultFromJson(data string) GetRecoverIntervalTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRecoverIntervalTableMasterResultFromDict(dict)
}

func NewGetRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) GetRecoverIntervalTableMasterResult {
	return GetRecoverIntervalTableMasterResult{
		Item: func() *RecoverIntervalTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRecoverIntervalTableMasterResult) Pointer() *GetRecoverIntervalTableMasterResult {
	return &p
}

type UpdateRecoverIntervalTableMasterResult struct {
	Item     *RecoverIntervalTableMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type UpdateRecoverIntervalTableMasterAsyncResult struct {
	result *UpdateRecoverIntervalTableMasterResult
	err    error
}

func NewUpdateRecoverIntervalTableMasterResultFromJson(data string) UpdateRecoverIntervalTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRecoverIntervalTableMasterResultFromDict(dict)
}

func NewUpdateRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) UpdateRecoverIntervalTableMasterResult {
	return UpdateRecoverIntervalTableMasterResult{
		Item: func() *RecoverIntervalTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateRecoverIntervalTableMasterResult) Pointer() *UpdateRecoverIntervalTableMasterResult {
	return &p
}

type DeleteRecoverIntervalTableMasterResult struct {
	Item     *RecoverIntervalTableMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type DeleteRecoverIntervalTableMasterAsyncResult struct {
	result *DeleteRecoverIntervalTableMasterResult
	err    error
}

func NewDeleteRecoverIntervalTableMasterResultFromJson(data string) DeleteRecoverIntervalTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRecoverIntervalTableMasterResultFromDict(dict)
}

func NewDeleteRecoverIntervalTableMasterResultFromDict(data map[string]interface{}) DeleteRecoverIntervalTableMasterResult {
	return DeleteRecoverIntervalTableMasterResult{
		Item: func() *RecoverIntervalTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverIntervalTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRecoverIntervalTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteRecoverIntervalTableMasterResult) Pointer() *DeleteRecoverIntervalTableMasterResult {
	return &p
}

type DescribeRecoverValueTableMastersResult struct {
	Items         []RecoverValueTableMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
	Metadata      *core.ResultMetadata      `json:"metadata"`
}

type DescribeRecoverValueTableMastersAsyncResult struct {
	result *DescribeRecoverValueTableMastersResult
	err    error
}

func NewDescribeRecoverValueTableMastersResultFromJson(data string) DescribeRecoverValueTableMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRecoverValueTableMastersResultFromDict(dict)
}

func NewDescribeRecoverValueTableMastersResultFromDict(data map[string]interface{}) DescribeRecoverValueTableMastersResult {
	return DescribeRecoverValueTableMastersResult{
		Items: func() []RecoverValueTableMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRecoverValueTableMasters(core.CastArray(data["items"]))
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

func (p DescribeRecoverValueTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRecoverValueTableMastersFromDict(
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

func (p DescribeRecoverValueTableMastersResult) Pointer() *DescribeRecoverValueTableMastersResult {
	return &p
}

type CreateRecoverValueTableMasterResult struct {
	Item     *RecoverValueTableMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type CreateRecoverValueTableMasterAsyncResult struct {
	result *CreateRecoverValueTableMasterResult
	err    error
}

func NewCreateRecoverValueTableMasterResultFromJson(data string) CreateRecoverValueTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRecoverValueTableMasterResultFromDict(dict)
}

func NewCreateRecoverValueTableMasterResultFromDict(data map[string]interface{}) CreateRecoverValueTableMasterResult {
	return CreateRecoverValueTableMasterResult{
		Item: func() *RecoverValueTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateRecoverValueTableMasterResult) Pointer() *CreateRecoverValueTableMasterResult {
	return &p
}

type GetRecoverValueTableMasterResult struct {
	Item     *RecoverValueTableMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type GetRecoverValueTableMasterAsyncResult struct {
	result *GetRecoverValueTableMasterResult
	err    error
}

func NewGetRecoverValueTableMasterResultFromJson(data string) GetRecoverValueTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRecoverValueTableMasterResultFromDict(dict)
}

func NewGetRecoverValueTableMasterResultFromDict(data map[string]interface{}) GetRecoverValueTableMasterResult {
	return GetRecoverValueTableMasterResult{
		Item: func() *RecoverValueTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRecoverValueTableMasterResult) Pointer() *GetRecoverValueTableMasterResult {
	return &p
}

type UpdateRecoverValueTableMasterResult struct {
	Item     *RecoverValueTableMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type UpdateRecoverValueTableMasterAsyncResult struct {
	result *UpdateRecoverValueTableMasterResult
	err    error
}

func NewUpdateRecoverValueTableMasterResultFromJson(data string) UpdateRecoverValueTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRecoverValueTableMasterResultFromDict(dict)
}

func NewUpdateRecoverValueTableMasterResultFromDict(data map[string]interface{}) UpdateRecoverValueTableMasterResult {
	return UpdateRecoverValueTableMasterResult{
		Item: func() *RecoverValueTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateRecoverValueTableMasterResult) Pointer() *UpdateRecoverValueTableMasterResult {
	return &p
}

type DeleteRecoverValueTableMasterResult struct {
	Item     *RecoverValueTableMaster `json:"item"`
	Metadata *core.ResultMetadata     `json:"metadata"`
}

type DeleteRecoverValueTableMasterAsyncResult struct {
	result *DeleteRecoverValueTableMasterResult
	err    error
}

func NewDeleteRecoverValueTableMasterResultFromJson(data string) DeleteRecoverValueTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRecoverValueTableMasterResultFromDict(dict)
}

func NewDeleteRecoverValueTableMasterResultFromDict(data map[string]interface{}) DeleteRecoverValueTableMasterResult {
	return DeleteRecoverValueTableMasterResult{
		Item: func() *RecoverValueTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRecoverValueTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRecoverValueTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteRecoverValueTableMasterResult) Pointer() *DeleteRecoverValueTableMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentStaminaMaster `json:"item"`
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
		Item: func() *CurrentStaminaMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentStaminaMasterResult struct {
	Item     *CurrentStaminaMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetCurrentStaminaMasterAsyncResult struct {
	result *GetCurrentStaminaMasterResult
	err    error
}

func NewGetCurrentStaminaMasterResultFromJson(data string) GetCurrentStaminaMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentStaminaMasterResultFromDict(dict)
}

func NewGetCurrentStaminaMasterResultFromDict(data map[string]interface{}) GetCurrentStaminaMasterResult {
	return GetCurrentStaminaMasterResult{
		Item: func() *CurrentStaminaMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentStaminaMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetCurrentStaminaMasterResult) Pointer() *GetCurrentStaminaMasterResult {
	return &p
}

type UpdateCurrentStaminaMasterResult struct {
	Item     *CurrentStaminaMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentStaminaMasterAsyncResult struct {
	result *UpdateCurrentStaminaMasterResult
	err    error
}

func NewUpdateCurrentStaminaMasterResultFromJson(data string) UpdateCurrentStaminaMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStaminaMasterResultFromDict(dict)
}

func NewUpdateCurrentStaminaMasterResultFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterResult {
	return UpdateCurrentStaminaMasterResult{
		Item: func() *CurrentStaminaMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentStaminaMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentStaminaMasterResult) Pointer() *UpdateCurrentStaminaMasterResult {
	return &p
}

type UpdateCurrentStaminaMasterFromGitHubResult struct {
	Item     *CurrentStaminaMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentStaminaMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentStaminaMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentStaminaMasterFromGitHubResultFromJson(data string) UpdateCurrentStaminaMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentStaminaMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentStaminaMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentStaminaMasterFromGitHubResult {
	return UpdateCurrentStaminaMasterFromGitHubResult{
		Item: func() *CurrentStaminaMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentStaminaMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentStaminaMasterFromGitHubResult) Pointer() *UpdateCurrentStaminaMasterFromGitHubResult {
	return &p
}

type DescribeStaminaModelsResult struct {
	Items    []StaminaModel       `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeStaminaModelsAsyncResult struct {
	result *DescribeStaminaModelsResult
	err    error
}

func NewDescribeStaminaModelsResultFromJson(data string) DescribeStaminaModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminaModelsResultFromDict(dict)
}

func NewDescribeStaminaModelsResultFromDict(data map[string]interface{}) DescribeStaminaModelsResult {
	return DescribeStaminaModelsResult{
		Items: func() []StaminaModel {
			if data["items"] == nil {
				return nil
			}
			return CastStaminaModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeStaminaModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminaModelsFromDict(
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

func (p DescribeStaminaModelsResult) Pointer() *DescribeStaminaModelsResult {
	return &p
}

type GetStaminaModelResult struct {
	Item     *StaminaModel        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStaminaModelAsyncResult struct {
	result *GetStaminaModelResult
	err    error
}

func NewGetStaminaModelResultFromJson(data string) GetStaminaModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaModelResultFromDict(dict)
}

func NewGetStaminaModelResultFromDict(data map[string]interface{}) GetStaminaModelResult {
	return GetStaminaModelResult{
		Item: func() *StaminaModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStaminaModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetStaminaModelResult) Pointer() *GetStaminaModelResult {
	return &p
}

type DescribeStaminasResult struct {
	Items         []Stamina            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStaminasAsyncResult struct {
	result *DescribeStaminasResult
	err    error
}

func NewDescribeStaminasResultFromJson(data string) DescribeStaminasResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminasResultFromDict(dict)
}

func NewDescribeStaminasResultFromDict(data map[string]interface{}) DescribeStaminasResult {
	return DescribeStaminasResult{
		Items: func() []Stamina {
			if data["items"] == nil {
				return nil
			}
			return CastStaminas(core.CastArray(data["items"]))
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

func (p DescribeStaminasResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminasFromDict(
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

func (p DescribeStaminasResult) Pointer() *DescribeStaminasResult {
	return &p
}

type DescribeStaminasByUserIdResult struct {
	Items         []Stamina            `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStaminasByUserIdAsyncResult struct {
	result *DescribeStaminasByUserIdResult
	err    error
}

func NewDescribeStaminasByUserIdResultFromJson(data string) DescribeStaminasByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStaminasByUserIdResultFromDict(dict)
}

func NewDescribeStaminasByUserIdResultFromDict(data map[string]interface{}) DescribeStaminasByUserIdResult {
	return DescribeStaminasByUserIdResult{
		Items: func() []Stamina {
			if data["items"] == nil {
				return nil
			}
			return CastStaminas(core.CastArray(data["items"]))
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

func (p DescribeStaminasByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStaminasFromDict(
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

func (p DescribeStaminasByUserIdResult) Pointer() *DescribeStaminasByUserIdResult {
	return &p
}

type GetStaminaResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type GetStaminaAsyncResult struct {
	result *GetStaminaResult
	err    error
}

func NewGetStaminaResultFromJson(data string) GetStaminaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaResultFromDict(dict)
}

func NewGetStaminaResultFromDict(data map[string]interface{}) GetStaminaResult {
	return GetStaminaResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStaminaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetStaminaResult) Pointer() *GetStaminaResult {
	return &p
}

type GetStaminaByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type GetStaminaByUserIdAsyncResult struct {
	result *GetStaminaByUserIdResult
	err    error
}

func NewGetStaminaByUserIdResultFromJson(data string) GetStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStaminaByUserIdResultFromDict(dict)
}

func NewGetStaminaByUserIdResultFromDict(data map[string]interface{}) GetStaminaByUserIdResult {
	return GetStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetStaminaByUserIdResult) Pointer() *GetStaminaByUserIdResult {
	return &p
}

type UpdateStaminaByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type UpdateStaminaByUserIdAsyncResult struct {
	result *UpdateStaminaByUserIdResult
	err    error
}

func NewUpdateStaminaByUserIdResultFromJson(data string) UpdateStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStaminaByUserIdResultFromDict(dict)
}

func NewUpdateStaminaByUserIdResultFromDict(data map[string]interface{}) UpdateStaminaByUserIdResult {
	return UpdateStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateStaminaByUserIdResult) Pointer() *UpdateStaminaByUserIdResult {
	return &p
}

type ConsumeStaminaResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type ConsumeStaminaAsyncResult struct {
	result *ConsumeStaminaResult
	err    error
}

func NewConsumeStaminaResultFromJson(data string) ConsumeStaminaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaResultFromDict(dict)
}

func NewConsumeStaminaResultFromDict(data map[string]interface{}) ConsumeStaminaResult {
	return ConsumeStaminaResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ConsumeStaminaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ConsumeStaminaResult) Pointer() *ConsumeStaminaResult {
	return &p
}

type ConsumeStaminaByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type ConsumeStaminaByUserIdAsyncResult struct {
	result *ConsumeStaminaByUserIdResult
	err    error
}

func NewConsumeStaminaByUserIdResultFromJson(data string) ConsumeStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaByUserIdResultFromDict(dict)
}

func NewConsumeStaminaByUserIdResultFromDict(data map[string]interface{}) ConsumeStaminaByUserIdResult {
	return ConsumeStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ConsumeStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ConsumeStaminaByUserIdResult) Pointer() *ConsumeStaminaByUserIdResult {
	return &p
}

type ApplyStaminaResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type ApplyStaminaAsyncResult struct {
	result *ApplyStaminaResult
	err    error
}

func NewApplyStaminaResultFromJson(data string) ApplyStaminaResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyStaminaResultFromDict(dict)
}

func NewApplyStaminaResultFromDict(data map[string]interface{}) ApplyStaminaResult {
	return ApplyStaminaResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ApplyStaminaResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ApplyStaminaResult) Pointer() *ApplyStaminaResult {
	return &p
}

type ApplyStaminaByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type ApplyStaminaByUserIdAsyncResult struct {
	result *ApplyStaminaByUserIdResult
	err    error
}

func NewApplyStaminaByUserIdResultFromJson(data string) ApplyStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyStaminaByUserIdResultFromDict(dict)
}

func NewApplyStaminaByUserIdResultFromDict(data map[string]interface{}) ApplyStaminaByUserIdResult {
	return ApplyStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ApplyStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ApplyStaminaByUserIdResult) Pointer() *ApplyStaminaByUserIdResult {
	return &p
}

type RecoverStaminaByUserIdResult struct {
	Item          *Stamina             `json:"item"`
	StaminaModel  *StaminaModel        `json:"staminaModel"`
	OverflowValue *int32               `json:"overflowValue"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type RecoverStaminaByUserIdAsyncResult struct {
	result *RecoverStaminaByUserIdResult
	err    error
}

func NewRecoverStaminaByUserIdResultFromJson(data string) RecoverStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecoverStaminaByUserIdResultFromDict(dict)
}

func NewRecoverStaminaByUserIdResultFromDict(data map[string]interface{}) RecoverStaminaByUserIdResult {
	return RecoverStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		OverflowValue: func() *int32 {
			v, ok := data["overflowValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["overflowValue"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RecoverStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"overflowValue": p.OverflowValue,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RecoverStaminaByUserIdResult) Pointer() *RecoverStaminaByUserIdResult {
	return &p
}

type RaiseMaxValueByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RaiseMaxValueByUserIdAsyncResult struct {
	result *RaiseMaxValueByUserIdResult
	err    error
}

func NewRaiseMaxValueByUserIdResultFromJson(data string) RaiseMaxValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRaiseMaxValueByUserIdResultFromDict(dict)
}

func NewRaiseMaxValueByUserIdResultFromDict(data map[string]interface{}) RaiseMaxValueByUserIdResult {
	return RaiseMaxValueByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RaiseMaxValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RaiseMaxValueByUserIdResult) Pointer() *RaiseMaxValueByUserIdResult {
	return &p
}

type DecreaseMaxValueResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaxValueAsyncResult struct {
	result *DecreaseMaxValueResult
	err    error
}

func NewDecreaseMaxValueResultFromJson(data string) DecreaseMaxValueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaxValueResultFromDict(dict)
}

func NewDecreaseMaxValueResultFromDict(data map[string]interface{}) DecreaseMaxValueResult {
	return DecreaseMaxValueResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DecreaseMaxValueResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DecreaseMaxValueResult) Pointer() *DecreaseMaxValueResult {
	return &p
}

type DecreaseMaxValueByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaxValueByUserIdAsyncResult struct {
	result *DecreaseMaxValueByUserIdResult
	err    error
}

func NewDecreaseMaxValueByUserIdResultFromJson(data string) DecreaseMaxValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaxValueByUserIdResultFromDict(dict)
}

func NewDecreaseMaxValueByUserIdResultFromDict(data map[string]interface{}) DecreaseMaxValueByUserIdResult {
	return DecreaseMaxValueByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DecreaseMaxValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DecreaseMaxValueByUserIdResult) Pointer() *DecreaseMaxValueByUserIdResult {
	return &p
}

type SetMaxValueByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetMaxValueByUserIdAsyncResult struct {
	result *SetMaxValueByUserIdResult
	err    error
}

func NewSetMaxValueByUserIdResultFromJson(data string) SetMaxValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByUserIdResultFromDict(dict)
}

func NewSetMaxValueByUserIdResultFromDict(data map[string]interface{}) SetMaxValueByUserIdResult {
	return SetMaxValueByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetMaxValueByUserIdResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetMaxValueByUserIdResult) Pointer() *SetMaxValueByUserIdResult {
	return &p
}

type SetRecoverIntervalByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverIntervalByUserIdAsyncResult struct {
	result *SetRecoverIntervalByUserIdResult
	err    error
}

func NewSetRecoverIntervalByUserIdResultFromJson(data string) SetRecoverIntervalByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByUserIdResultFromDict(dict)
}

func NewSetRecoverIntervalByUserIdResultFromDict(data map[string]interface{}) SetRecoverIntervalByUserIdResult {
	return SetRecoverIntervalByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverIntervalByUserIdResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverIntervalByUserIdResult) Pointer() *SetRecoverIntervalByUserIdResult {
	return &p
}

type SetRecoverValueByUserIdResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverValueByUserIdAsyncResult struct {
	result *SetRecoverValueByUserIdResult
	err    error
}

func NewSetRecoverValueByUserIdResultFromJson(data string) SetRecoverValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByUserIdResultFromDict(dict)
}

func NewSetRecoverValueByUserIdResultFromDict(data map[string]interface{}) SetRecoverValueByUserIdResult {
	return SetRecoverValueByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverValueByUserIdResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverValueByUserIdResult) Pointer() *SetRecoverValueByUserIdResult {
	return &p
}

type SetMaxValueByStatusResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetMaxValueByStatusAsyncResult struct {
	result *SetMaxValueByStatusResult
	err    error
}

func NewSetMaxValueByStatusResultFromJson(data string) SetMaxValueByStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByStatusResultFromDict(dict)
}

func NewSetMaxValueByStatusResultFromDict(data map[string]interface{}) SetMaxValueByStatusResult {
	return SetMaxValueByStatusResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetMaxValueByStatusResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetMaxValueByStatusResult) Pointer() *SetMaxValueByStatusResult {
	return &p
}

type SetRecoverIntervalByStatusResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverIntervalByStatusAsyncResult struct {
	result *SetRecoverIntervalByStatusResult
	err    error
}

func NewSetRecoverIntervalByStatusResultFromJson(data string) SetRecoverIntervalByStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByStatusResultFromDict(dict)
}

func NewSetRecoverIntervalByStatusResultFromDict(data map[string]interface{}) SetRecoverIntervalByStatusResult {
	return SetRecoverIntervalByStatusResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverIntervalByStatusResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverIntervalByStatusResult) Pointer() *SetRecoverIntervalByStatusResult {
	return &p
}

type SetRecoverValueByStatusResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverValueByStatusAsyncResult struct {
	result *SetRecoverValueByStatusResult
	err    error
}

func NewSetRecoverValueByStatusResultFromJson(data string) SetRecoverValueByStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByStatusResultFromDict(dict)
}

func NewSetRecoverValueByStatusResultFromDict(data map[string]interface{}) SetRecoverValueByStatusResult {
	return SetRecoverValueByStatusResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverValueByStatusResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverValueByStatusResult) Pointer() *SetRecoverValueByStatusResult {
	return &p
}

type DeleteStaminaByUserIdResult struct {
	Item     *Stamina             `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStaminaByUserIdAsyncResult struct {
	result *DeleteStaminaByUserIdResult
	err    error
}

func NewDeleteStaminaByUserIdResultFromJson(data string) DeleteStaminaByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStaminaByUserIdResultFromDict(dict)
}

func NewDeleteStaminaByUserIdResultFromDict(data map[string]interface{}) DeleteStaminaByUserIdResult {
	return DeleteStaminaByUserIdResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteStaminaByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteStaminaByUserIdResult) Pointer() *DeleteStaminaByUserIdResult {
	return &p
}

type VerifyStaminaValueResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaValueAsyncResult struct {
	result *VerifyStaminaValueResult
	err    error
}

func NewVerifyStaminaValueResultFromJson(data string) VerifyStaminaValueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaValueResultFromDict(dict)
}

func NewVerifyStaminaValueResultFromDict(data map[string]interface{}) VerifyStaminaValueResult {
	return VerifyStaminaValueResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaValueResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaValueResult) Pointer() *VerifyStaminaValueResult {
	return &p
}

type VerifyStaminaValueByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaValueByUserIdAsyncResult struct {
	result *VerifyStaminaValueByUserIdResult
	err    error
}

func NewVerifyStaminaValueByUserIdResultFromJson(data string) VerifyStaminaValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaValueByUserIdResultFromDict(dict)
}

func NewVerifyStaminaValueByUserIdResultFromDict(data map[string]interface{}) VerifyStaminaValueByUserIdResult {
	return VerifyStaminaValueByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaValueByUserIdResult) Pointer() *VerifyStaminaValueByUserIdResult {
	return &p
}

type VerifyStaminaMaxValueResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaMaxValueAsyncResult struct {
	result *VerifyStaminaMaxValueResult
	err    error
}

func NewVerifyStaminaMaxValueResultFromJson(data string) VerifyStaminaMaxValueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaMaxValueResultFromDict(dict)
}

func NewVerifyStaminaMaxValueResultFromDict(data map[string]interface{}) VerifyStaminaMaxValueResult {
	return VerifyStaminaMaxValueResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaMaxValueResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaMaxValueResult) Pointer() *VerifyStaminaMaxValueResult {
	return &p
}

type VerifyStaminaMaxValueByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaMaxValueByUserIdAsyncResult struct {
	result *VerifyStaminaMaxValueByUserIdResult
	err    error
}

func NewVerifyStaminaMaxValueByUserIdResultFromJson(data string) VerifyStaminaMaxValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaMaxValueByUserIdResultFromDict(dict)
}

func NewVerifyStaminaMaxValueByUserIdResultFromDict(data map[string]interface{}) VerifyStaminaMaxValueByUserIdResult {
	return VerifyStaminaMaxValueByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaMaxValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaMaxValueByUserIdResult) Pointer() *VerifyStaminaMaxValueByUserIdResult {
	return &p
}

type VerifyStaminaRecoverIntervalMinutesResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverIntervalMinutesAsyncResult struct {
	result *VerifyStaminaRecoverIntervalMinutesResult
	err    error
}

func NewVerifyStaminaRecoverIntervalMinutesResultFromJson(data string) VerifyStaminaRecoverIntervalMinutesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverIntervalMinutesResultFromDict(dict)
}

func NewVerifyStaminaRecoverIntervalMinutesResultFromDict(data map[string]interface{}) VerifyStaminaRecoverIntervalMinutesResult {
	return VerifyStaminaRecoverIntervalMinutesResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaRecoverIntervalMinutesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaRecoverIntervalMinutesResult) Pointer() *VerifyStaminaRecoverIntervalMinutesResult {
	return &p
}

type VerifyStaminaRecoverIntervalMinutesByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult struct {
	result *VerifyStaminaRecoverIntervalMinutesByUserIdResult
	err    error
}

func NewVerifyStaminaRecoverIntervalMinutesByUserIdResultFromJson(data string) VerifyStaminaRecoverIntervalMinutesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverIntervalMinutesByUserIdResultFromDict(dict)
}

func NewVerifyStaminaRecoverIntervalMinutesByUserIdResultFromDict(data map[string]interface{}) VerifyStaminaRecoverIntervalMinutesByUserIdResult {
	return VerifyStaminaRecoverIntervalMinutesByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaRecoverIntervalMinutesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaRecoverIntervalMinutesByUserIdResult) Pointer() *VerifyStaminaRecoverIntervalMinutesByUserIdResult {
	return &p
}

type VerifyStaminaRecoverValueResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverValueAsyncResult struct {
	result *VerifyStaminaRecoverValueResult
	err    error
}

func NewVerifyStaminaRecoverValueResultFromJson(data string) VerifyStaminaRecoverValueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverValueResultFromDict(dict)
}

func NewVerifyStaminaRecoverValueResultFromDict(data map[string]interface{}) VerifyStaminaRecoverValueResult {
	return VerifyStaminaRecoverValueResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaRecoverValueResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaRecoverValueResult) Pointer() *VerifyStaminaRecoverValueResult {
	return &p
}

type VerifyStaminaRecoverValueByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverValueByUserIdAsyncResult struct {
	result *VerifyStaminaRecoverValueByUserIdResult
	err    error
}

func NewVerifyStaminaRecoverValueByUserIdResultFromJson(data string) VerifyStaminaRecoverValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverValueByUserIdResultFromDict(dict)
}

func NewVerifyStaminaRecoverValueByUserIdResultFromDict(data map[string]interface{}) VerifyStaminaRecoverValueByUserIdResult {
	return VerifyStaminaRecoverValueByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaRecoverValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaRecoverValueByUserIdResult) Pointer() *VerifyStaminaRecoverValueByUserIdResult {
	return &p
}

type VerifyStaminaOverflowValueResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaOverflowValueAsyncResult struct {
	result *VerifyStaminaOverflowValueResult
	err    error
}

func NewVerifyStaminaOverflowValueResultFromJson(data string) VerifyStaminaOverflowValueResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaOverflowValueResultFromDict(dict)
}

func NewVerifyStaminaOverflowValueResultFromDict(data map[string]interface{}) VerifyStaminaOverflowValueResult {
	return VerifyStaminaOverflowValueResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaOverflowValueResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaOverflowValueResult) Pointer() *VerifyStaminaOverflowValueResult {
	return &p
}

type VerifyStaminaOverflowValueByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaOverflowValueByUserIdAsyncResult struct {
	result *VerifyStaminaOverflowValueByUserIdResult
	err    error
}

func NewVerifyStaminaOverflowValueByUserIdResultFromJson(data string) VerifyStaminaOverflowValueByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaOverflowValueByUserIdResultFromDict(dict)
}

func NewVerifyStaminaOverflowValueByUserIdResultFromDict(data map[string]interface{}) VerifyStaminaOverflowValueByUserIdResult {
	return VerifyStaminaOverflowValueByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyStaminaOverflowValueByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyStaminaOverflowValueByUserIdResult) Pointer() *VerifyStaminaOverflowValueByUserIdResult {
	return &p
}

type RecoverStaminaByStampSheetResult struct {
	Item          *Stamina             `json:"item"`
	StaminaModel  *StaminaModel        `json:"staminaModel"`
	OverflowValue *int32               `json:"overflowValue"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type RecoverStaminaByStampSheetAsyncResult struct {
	result *RecoverStaminaByStampSheetResult
	err    error
}

func NewRecoverStaminaByStampSheetResultFromJson(data string) RecoverStaminaByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecoverStaminaByStampSheetResultFromDict(dict)
}

func NewRecoverStaminaByStampSheetResultFromDict(data map[string]interface{}) RecoverStaminaByStampSheetResult {
	return RecoverStaminaByStampSheetResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		OverflowValue: func() *int32 {
			v, ok := data["overflowValue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["overflowValue"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RecoverStaminaByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"overflowValue": p.OverflowValue,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RecoverStaminaByStampSheetResult) Pointer() *RecoverStaminaByStampSheetResult {
	return &p
}

type RaiseMaxValueByStampSheetResult struct {
	Item         *Stamina             `json:"item"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type RaiseMaxValueByStampSheetAsyncResult struct {
	result *RaiseMaxValueByStampSheetResult
	err    error
}

func NewRaiseMaxValueByStampSheetResultFromJson(data string) RaiseMaxValueByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRaiseMaxValueByStampSheetResultFromDict(dict)
}

func NewRaiseMaxValueByStampSheetResultFromDict(data map[string]interface{}) RaiseMaxValueByStampSheetResult {
	return RaiseMaxValueByStampSheetResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p RaiseMaxValueByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p RaiseMaxValueByStampSheetResult) Pointer() *RaiseMaxValueByStampSheetResult {
	return &p
}

type DecreaseMaxValueByStampTaskResult struct {
	Item            *Stamina             `json:"item"`
	StaminaModel    *StaminaModel        `json:"staminaModel"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type DecreaseMaxValueByStampTaskAsyncResult struct {
	result *DecreaseMaxValueByStampTaskResult
	err    error
}

func NewDecreaseMaxValueByStampTaskResultFromJson(data string) DecreaseMaxValueByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaxValueByStampTaskResultFromDict(dict)
}

func NewDecreaseMaxValueByStampTaskResultFromDict(data map[string]interface{}) DecreaseMaxValueByStampTaskResult {
	return DecreaseMaxValueByStampTaskResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
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

func (p DecreaseMaxValueByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
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

func (p DecreaseMaxValueByStampTaskResult) Pointer() *DecreaseMaxValueByStampTaskResult {
	return &p
}

type SetMaxValueByStampSheetResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetMaxValueByStampSheetAsyncResult struct {
	result *SetMaxValueByStampSheetResult
	err    error
}

func NewSetMaxValueByStampSheetResultFromJson(data string) SetMaxValueByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaxValueByStampSheetResultFromDict(dict)
}

func NewSetMaxValueByStampSheetResultFromDict(data map[string]interface{}) SetMaxValueByStampSheetResult {
	return SetMaxValueByStampSheetResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetMaxValueByStampSheetResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetMaxValueByStampSheetResult) Pointer() *SetMaxValueByStampSheetResult {
	return &p
}

type SetRecoverIntervalByStampSheetResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverIntervalByStampSheetAsyncResult struct {
	result *SetRecoverIntervalByStampSheetResult
	err    error
}

func NewSetRecoverIntervalByStampSheetResultFromJson(data string) SetRecoverIntervalByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverIntervalByStampSheetResultFromDict(dict)
}

func NewSetRecoverIntervalByStampSheetResultFromDict(data map[string]interface{}) SetRecoverIntervalByStampSheetResult {
	return SetRecoverIntervalByStampSheetResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverIntervalByStampSheetResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverIntervalByStampSheetResult) Pointer() *SetRecoverIntervalByStampSheetResult {
	return &p
}

type SetRecoverValueByStampSheetResult struct {
	Item         *Stamina             `json:"item"`
	Old          *Stamina             `json:"old"`
	StaminaModel *StaminaModel        `json:"staminaModel"`
	Metadata     *core.ResultMetadata `json:"metadata"`
}

type SetRecoverValueByStampSheetAsyncResult struct {
	result *SetRecoverValueByStampSheetResult
	err    error
}

func NewSetRecoverValueByStampSheetResultFromJson(data string) SetRecoverValueByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRecoverValueByStampSheetResultFromDict(dict)
}

func NewSetRecoverValueByStampSheetResultFromDict(data map[string]interface{}) SetRecoverValueByStampSheetResult {
	return SetRecoverValueByStampSheetResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Stamina {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRecoverValueByStampSheetResult) ToDict() map[string]interface{} {
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
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SetRecoverValueByStampSheetResult) Pointer() *SetRecoverValueByStampSheetResult {
	return &p
}

type ConsumeStaminaByStampTaskResult struct {
	Item            *Stamina             `json:"item"`
	StaminaModel    *StaminaModel        `json:"staminaModel"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type ConsumeStaminaByStampTaskAsyncResult struct {
	result *ConsumeStaminaByStampTaskResult
	err    error
}

func NewConsumeStaminaByStampTaskResultFromJson(data string) ConsumeStaminaByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewConsumeStaminaByStampTaskResultFromDict(dict)
}

func NewConsumeStaminaByStampTaskResultFromDict(data map[string]interface{}) ConsumeStaminaByStampTaskResult {
	return ConsumeStaminaByStampTaskResult{
		Item: func() *Stamina {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		StaminaModel: func() *StaminaModel {
			v, ok := data["staminaModel"]
			if !ok || v == nil {
				return nil
			}
			return NewStaminaModelFromDict(core.CastMap(data["staminaModel"])).Pointer()
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

func (p ConsumeStaminaByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"staminaModel": func() map[string]interface{} {
			if p.StaminaModel == nil {
				return nil
			}
			return p.StaminaModel.ToDict()
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

func (p ConsumeStaminaByStampTaskResult) Pointer() *ConsumeStaminaByStampTaskResult {
	return &p
}

type VerifyStaminaValueByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaValueByStampTaskAsyncResult struct {
	result *VerifyStaminaValueByStampTaskResult
	err    error
}

func NewVerifyStaminaValueByStampTaskResultFromJson(data string) VerifyStaminaValueByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaValueByStampTaskResultFromDict(dict)
}

func NewVerifyStaminaValueByStampTaskResultFromDict(data map[string]interface{}) VerifyStaminaValueByStampTaskResult {
	return VerifyStaminaValueByStampTaskResult{
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

func (p VerifyStaminaValueByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyStaminaValueByStampTaskResult) Pointer() *VerifyStaminaValueByStampTaskResult {
	return &p
}

type VerifyStaminaMaxValueByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaMaxValueByStampTaskAsyncResult struct {
	result *VerifyStaminaMaxValueByStampTaskResult
	err    error
}

func NewVerifyStaminaMaxValueByStampTaskResultFromJson(data string) VerifyStaminaMaxValueByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaMaxValueByStampTaskResultFromDict(dict)
}

func NewVerifyStaminaMaxValueByStampTaskResultFromDict(data map[string]interface{}) VerifyStaminaMaxValueByStampTaskResult {
	return VerifyStaminaMaxValueByStampTaskResult{
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

func (p VerifyStaminaMaxValueByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyStaminaMaxValueByStampTaskResult) Pointer() *VerifyStaminaMaxValueByStampTaskResult {
	return &p
}

type VerifyStaminaRecoverIntervalMinutesByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult struct {
	result *VerifyStaminaRecoverIntervalMinutesByStampTaskResult
	err    error
}

func NewVerifyStaminaRecoverIntervalMinutesByStampTaskResultFromJson(data string) VerifyStaminaRecoverIntervalMinutesByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverIntervalMinutesByStampTaskResultFromDict(dict)
}

func NewVerifyStaminaRecoverIntervalMinutesByStampTaskResultFromDict(data map[string]interface{}) VerifyStaminaRecoverIntervalMinutesByStampTaskResult {
	return VerifyStaminaRecoverIntervalMinutesByStampTaskResult{
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

func (p VerifyStaminaRecoverIntervalMinutesByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyStaminaRecoverIntervalMinutesByStampTaskResult) Pointer() *VerifyStaminaRecoverIntervalMinutesByStampTaskResult {
	return &p
}

type VerifyStaminaRecoverValueByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaRecoverValueByStampTaskAsyncResult struct {
	result *VerifyStaminaRecoverValueByStampTaskResult
	err    error
}

func NewVerifyStaminaRecoverValueByStampTaskResultFromJson(data string) VerifyStaminaRecoverValueByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaRecoverValueByStampTaskResultFromDict(dict)
}

func NewVerifyStaminaRecoverValueByStampTaskResultFromDict(data map[string]interface{}) VerifyStaminaRecoverValueByStampTaskResult {
	return VerifyStaminaRecoverValueByStampTaskResult{
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

func (p VerifyStaminaRecoverValueByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyStaminaRecoverValueByStampTaskResult) Pointer() *VerifyStaminaRecoverValueByStampTaskResult {
	return &p
}

type VerifyStaminaOverflowValueByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyStaminaOverflowValueByStampTaskAsyncResult struct {
	result *VerifyStaminaOverflowValueByStampTaskResult
	err    error
}

func NewVerifyStaminaOverflowValueByStampTaskResultFromJson(data string) VerifyStaminaOverflowValueByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyStaminaOverflowValueByStampTaskResultFromDict(dict)
}

func NewVerifyStaminaOverflowValueByStampTaskResultFromDict(data map[string]interface{}) VerifyStaminaOverflowValueByStampTaskResult {
	return VerifyStaminaOverflowValueByStampTaskResult{
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

func (p VerifyStaminaOverflowValueByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyStaminaOverflowValueByStampTaskResult) Pointer() *VerifyStaminaOverflowValueByStampTaskResult {
	return &p
}
