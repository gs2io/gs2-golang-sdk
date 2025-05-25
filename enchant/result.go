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

package enchant

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

type DescribeBalanceParameterModelsResult struct {
	Items    []BalanceParameterModel `json:"items"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type DescribeBalanceParameterModelsAsyncResult struct {
	result *DescribeBalanceParameterModelsResult
	err    error
}

func NewDescribeBalanceParameterModelsResultFromJson(data string) DescribeBalanceParameterModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterModelsResultFromDict(dict)
}

func NewDescribeBalanceParameterModelsResultFromDict(data map[string]interface{}) DescribeBalanceParameterModelsResult {
	return DescribeBalanceParameterModelsResult{
		Items: func() []BalanceParameterModel {
			if data["items"] == nil {
				return nil
			}
			return CastBalanceParameterModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeBalanceParameterModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBalanceParameterModelsFromDict(
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

func (p DescribeBalanceParameterModelsResult) Pointer() *DescribeBalanceParameterModelsResult {
	return &p
}

type GetBalanceParameterModelResult struct {
	Item     *BalanceParameterModel `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetBalanceParameterModelAsyncResult struct {
	result *GetBalanceParameterModelResult
	err    error
}

func NewGetBalanceParameterModelResultFromJson(data string) GetBalanceParameterModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterModelResultFromDict(dict)
}

func NewGetBalanceParameterModelResultFromDict(data map[string]interface{}) GetBalanceParameterModelResult {
	return GetBalanceParameterModelResult{
		Item: func() *BalanceParameterModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBalanceParameterModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBalanceParameterModelResult) Pointer() *GetBalanceParameterModelResult {
	return &p
}

type DescribeBalanceParameterModelMastersResult struct {
	Items         []BalanceParameterModelMaster `json:"items"`
	NextPageToken *string                       `json:"nextPageToken"`
	Metadata      *core.ResultMetadata          `json:"metadata"`
}

type DescribeBalanceParameterModelMastersAsyncResult struct {
	result *DescribeBalanceParameterModelMastersResult
	err    error
}

func NewDescribeBalanceParameterModelMastersResultFromJson(data string) DescribeBalanceParameterModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterModelMastersResultFromDict(dict)
}

func NewDescribeBalanceParameterModelMastersResultFromDict(data map[string]interface{}) DescribeBalanceParameterModelMastersResult {
	return DescribeBalanceParameterModelMastersResult{
		Items: func() []BalanceParameterModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastBalanceParameterModelMasters(core.CastArray(data["items"]))
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

func (p DescribeBalanceParameterModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBalanceParameterModelMastersFromDict(
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

func (p DescribeBalanceParameterModelMastersResult) Pointer() *DescribeBalanceParameterModelMastersResult {
	return &p
}

type CreateBalanceParameterModelMasterResult struct {
	Item     *BalanceParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type CreateBalanceParameterModelMasterAsyncResult struct {
	result *CreateBalanceParameterModelMasterResult
	err    error
}

func NewCreateBalanceParameterModelMasterResultFromJson(data string) CreateBalanceParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateBalanceParameterModelMasterResultFromDict(dict)
}

func NewCreateBalanceParameterModelMasterResultFromDict(data map[string]interface{}) CreateBalanceParameterModelMasterResult {
	return CreateBalanceParameterModelMasterResult{
		Item: func() *BalanceParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateBalanceParameterModelMasterResult) Pointer() *CreateBalanceParameterModelMasterResult {
	return &p
}

type GetBalanceParameterModelMasterResult struct {
	Item     *BalanceParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type GetBalanceParameterModelMasterAsyncResult struct {
	result *GetBalanceParameterModelMasterResult
	err    error
}

func NewGetBalanceParameterModelMasterResultFromJson(data string) GetBalanceParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterModelMasterResultFromDict(dict)
}

func NewGetBalanceParameterModelMasterResultFromDict(data map[string]interface{}) GetBalanceParameterModelMasterResult {
	return GetBalanceParameterModelMasterResult{
		Item: func() *BalanceParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBalanceParameterModelMasterResult) Pointer() *GetBalanceParameterModelMasterResult {
	return &p
}

type UpdateBalanceParameterModelMasterResult struct {
	Item     *BalanceParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type UpdateBalanceParameterModelMasterAsyncResult struct {
	result *UpdateBalanceParameterModelMasterResult
	err    error
}

func NewUpdateBalanceParameterModelMasterResultFromJson(data string) UpdateBalanceParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBalanceParameterModelMasterResultFromDict(dict)
}

func NewUpdateBalanceParameterModelMasterResultFromDict(data map[string]interface{}) UpdateBalanceParameterModelMasterResult {
	return UpdateBalanceParameterModelMasterResult{
		Item: func() *BalanceParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateBalanceParameterModelMasterResult) Pointer() *UpdateBalanceParameterModelMasterResult {
	return &p
}

type DeleteBalanceParameterModelMasterResult struct {
	Item     *BalanceParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type DeleteBalanceParameterModelMasterAsyncResult struct {
	result *DeleteBalanceParameterModelMasterResult
	err    error
}

func NewDeleteBalanceParameterModelMasterResultFromJson(data string) DeleteBalanceParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBalanceParameterModelMasterResultFromDict(dict)
}

func NewDeleteBalanceParameterModelMasterResultFromDict(data map[string]interface{}) DeleteBalanceParameterModelMasterResult {
	return DeleteBalanceParameterModelMasterResult{
		Item: func() *BalanceParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteBalanceParameterModelMasterResult) Pointer() *DeleteBalanceParameterModelMasterResult {
	return &p
}

type DescribeRarityParameterModelsResult struct {
	Items    []RarityParameterModel `json:"items"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DescribeRarityParameterModelsAsyncResult struct {
	result *DescribeRarityParameterModelsResult
	err    error
}

func NewDescribeRarityParameterModelsResultFromJson(data string) DescribeRarityParameterModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterModelsResultFromDict(dict)
}

func NewDescribeRarityParameterModelsResultFromDict(data map[string]interface{}) DescribeRarityParameterModelsResult {
	return DescribeRarityParameterModelsResult{
		Items: func() []RarityParameterModel {
			if data["items"] == nil {
				return nil
			}
			return CastRarityParameterModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeRarityParameterModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRarityParameterModelsFromDict(
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

func (p DescribeRarityParameterModelsResult) Pointer() *DescribeRarityParameterModelsResult {
	return &p
}

type GetRarityParameterModelResult struct {
	Item     *RarityParameterModel `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetRarityParameterModelAsyncResult struct {
	result *GetRarityParameterModelResult
	err    error
}

func NewGetRarityParameterModelResultFromJson(data string) GetRarityParameterModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterModelResultFromDict(dict)
}

func NewGetRarityParameterModelResultFromDict(data map[string]interface{}) GetRarityParameterModelResult {
	return GetRarityParameterModelResult{
		Item: func() *RarityParameterModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRarityParameterModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRarityParameterModelResult) Pointer() *GetRarityParameterModelResult {
	return &p
}

type DescribeRarityParameterModelMastersResult struct {
	Items         []RarityParameterModelMaster `json:"items"`
	NextPageToken *string                      `json:"nextPageToken"`
	Metadata      *core.ResultMetadata         `json:"metadata"`
}

type DescribeRarityParameterModelMastersAsyncResult struct {
	result *DescribeRarityParameterModelMastersResult
	err    error
}

func NewDescribeRarityParameterModelMastersResultFromJson(data string) DescribeRarityParameterModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterModelMastersResultFromDict(dict)
}

func NewDescribeRarityParameterModelMastersResultFromDict(data map[string]interface{}) DescribeRarityParameterModelMastersResult {
	return DescribeRarityParameterModelMastersResult{
		Items: func() []RarityParameterModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRarityParameterModelMasters(core.CastArray(data["items"]))
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

func (p DescribeRarityParameterModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRarityParameterModelMastersFromDict(
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

func (p DescribeRarityParameterModelMastersResult) Pointer() *DescribeRarityParameterModelMastersResult {
	return &p
}

type CreateRarityParameterModelMasterResult struct {
	Item     *RarityParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type CreateRarityParameterModelMasterAsyncResult struct {
	result *CreateRarityParameterModelMasterResult
	err    error
}

func NewCreateRarityParameterModelMasterResultFromJson(data string) CreateRarityParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRarityParameterModelMasterResultFromDict(dict)
}

func NewCreateRarityParameterModelMasterResultFromDict(data map[string]interface{}) CreateRarityParameterModelMasterResult {
	return CreateRarityParameterModelMasterResult{
		Item: func() *RarityParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRarityParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateRarityParameterModelMasterResult) Pointer() *CreateRarityParameterModelMasterResult {
	return &p
}

type GetRarityParameterModelMasterResult struct {
	Item     *RarityParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type GetRarityParameterModelMasterAsyncResult struct {
	result *GetRarityParameterModelMasterResult
	err    error
}

func NewGetRarityParameterModelMasterResultFromJson(data string) GetRarityParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterModelMasterResultFromDict(dict)
}

func NewGetRarityParameterModelMasterResultFromDict(data map[string]interface{}) GetRarityParameterModelMasterResult {
	return GetRarityParameterModelMasterResult{
		Item: func() *RarityParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRarityParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRarityParameterModelMasterResult) Pointer() *GetRarityParameterModelMasterResult {
	return &p
}

type UpdateRarityParameterModelMasterResult struct {
	Item     *RarityParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type UpdateRarityParameterModelMasterAsyncResult struct {
	result *UpdateRarityParameterModelMasterResult
	err    error
}

func NewUpdateRarityParameterModelMasterResultFromJson(data string) UpdateRarityParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRarityParameterModelMasterResultFromDict(dict)
}

func NewUpdateRarityParameterModelMasterResultFromDict(data map[string]interface{}) UpdateRarityParameterModelMasterResult {
	return UpdateRarityParameterModelMasterResult{
		Item: func() *RarityParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRarityParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateRarityParameterModelMasterResult) Pointer() *UpdateRarityParameterModelMasterResult {
	return &p
}

type DeleteRarityParameterModelMasterResult struct {
	Item     *RarityParameterModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type DeleteRarityParameterModelMasterAsyncResult struct {
	result *DeleteRarityParameterModelMasterResult
	err    error
}

func NewDeleteRarityParameterModelMasterResultFromJson(data string) DeleteRarityParameterModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRarityParameterModelMasterResultFromDict(dict)
}

func NewDeleteRarityParameterModelMasterResultFromDict(data map[string]interface{}) DeleteRarityParameterModelMasterResult {
	return DeleteRarityParameterModelMasterResult{
		Item: func() *RarityParameterModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRarityParameterModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteRarityParameterModelMasterResult) Pointer() *DeleteRarityParameterModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentParameterMaster `json:"item"`
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
		Item: func() *CurrentParameterMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentParameterMasterResult struct {
	Item     *CurrentParameterMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type GetCurrentParameterMasterAsyncResult struct {
	result *GetCurrentParameterMasterResult
	err    error
}

func NewGetCurrentParameterMasterResultFromJson(data string) GetCurrentParameterMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentParameterMasterResultFromDict(dict)
}

func NewGetCurrentParameterMasterResultFromDict(data map[string]interface{}) GetCurrentParameterMasterResult {
	return GetCurrentParameterMasterResult{
		Item: func() *CurrentParameterMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentParameterMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetCurrentParameterMasterResult) Pointer() *GetCurrentParameterMasterResult {
	return &p
}

type PreUpdateCurrentParameterMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentParameterMasterAsyncResult struct {
	result *PreUpdateCurrentParameterMasterResult
	err    error
}

func NewPreUpdateCurrentParameterMasterResultFromJson(data string) PreUpdateCurrentParameterMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentParameterMasterResultFromDict(dict)
}

func NewPreUpdateCurrentParameterMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentParameterMasterResult {
	return PreUpdateCurrentParameterMasterResult{
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

func (p PreUpdateCurrentParameterMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentParameterMasterResult) Pointer() *PreUpdateCurrentParameterMasterResult {
	return &p
}

type UpdateCurrentParameterMasterResult struct {
	Item     *CurrentParameterMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type UpdateCurrentParameterMasterAsyncResult struct {
	result *UpdateCurrentParameterMasterResult
	err    error
}

func NewUpdateCurrentParameterMasterResultFromJson(data string) UpdateCurrentParameterMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentParameterMasterResultFromDict(dict)
}

func NewUpdateCurrentParameterMasterResultFromDict(data map[string]interface{}) UpdateCurrentParameterMasterResult {
	return UpdateCurrentParameterMasterResult{
		Item: func() *CurrentParameterMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentParameterMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentParameterMasterResult) Pointer() *UpdateCurrentParameterMasterResult {
	return &p
}

type UpdateCurrentParameterMasterFromGitHubResult struct {
	Item     *CurrentParameterMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type UpdateCurrentParameterMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentParameterMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentParameterMasterFromGitHubResultFromJson(data string) UpdateCurrentParameterMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentParameterMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentParameterMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentParameterMasterFromGitHubResult {
	return UpdateCurrentParameterMasterFromGitHubResult{
		Item: func() *CurrentParameterMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentParameterMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentParameterMasterFromGitHubResult) Pointer() *UpdateCurrentParameterMasterFromGitHubResult {
	return &p
}

type DescribeBalanceParameterStatusesResult struct {
	Items         []BalanceParameterStatus `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
	Metadata      *core.ResultMetadata     `json:"metadata"`
}

type DescribeBalanceParameterStatusesAsyncResult struct {
	result *DescribeBalanceParameterStatusesResult
	err    error
}

func NewDescribeBalanceParameterStatusesResultFromJson(data string) DescribeBalanceParameterStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterStatusesResultFromDict(dict)
}

func NewDescribeBalanceParameterStatusesResultFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesResult {
	return DescribeBalanceParameterStatusesResult{
		Items: func() []BalanceParameterStatus {
			if data["items"] == nil {
				return nil
			}
			return CastBalanceParameterStatuses(core.CastArray(data["items"]))
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

func (p DescribeBalanceParameterStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBalanceParameterStatusesFromDict(
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

func (p DescribeBalanceParameterStatusesResult) Pointer() *DescribeBalanceParameterStatusesResult {
	return &p
}

type DescribeBalanceParameterStatusesByUserIdResult struct {
	Items         []BalanceParameterStatus `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
	Metadata      *core.ResultMetadata     `json:"metadata"`
}

type DescribeBalanceParameterStatusesByUserIdAsyncResult struct {
	result *DescribeBalanceParameterStatusesByUserIdResult
	err    error
}

func NewDescribeBalanceParameterStatusesByUserIdResultFromJson(data string) DescribeBalanceParameterStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBalanceParameterStatusesByUserIdResultFromDict(dict)
}

func NewDescribeBalanceParameterStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesByUserIdResult {
	return DescribeBalanceParameterStatusesByUserIdResult{
		Items: func() []BalanceParameterStatus {
			if data["items"] == nil {
				return nil
			}
			return CastBalanceParameterStatuses(core.CastArray(data["items"]))
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

func (p DescribeBalanceParameterStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBalanceParameterStatusesFromDict(
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

func (p DescribeBalanceParameterStatusesByUserIdResult) Pointer() *DescribeBalanceParameterStatusesByUserIdResult {
	return &p
}

type GetBalanceParameterStatusResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type GetBalanceParameterStatusAsyncResult struct {
	result *GetBalanceParameterStatusResult
	err    error
}

func NewGetBalanceParameterStatusResultFromJson(data string) GetBalanceParameterStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterStatusResultFromDict(dict)
}

func NewGetBalanceParameterStatusResultFromDict(data map[string]interface{}) GetBalanceParameterStatusResult {
	return GetBalanceParameterStatusResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBalanceParameterStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBalanceParameterStatusResult) Pointer() *GetBalanceParameterStatusResult {
	return &p
}

type GetBalanceParameterStatusByUserIdResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type GetBalanceParameterStatusByUserIdAsyncResult struct {
	result *GetBalanceParameterStatusByUserIdResult
	err    error
}

func NewGetBalanceParameterStatusByUserIdResultFromJson(data string) GetBalanceParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewGetBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) GetBalanceParameterStatusByUserIdResult {
	return GetBalanceParameterStatusByUserIdResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetBalanceParameterStatusByUserIdResult) Pointer() *GetBalanceParameterStatusByUserIdResult {
	return &p
}

type DeleteBalanceParameterStatusByUserIdResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type DeleteBalanceParameterStatusByUserIdAsyncResult struct {
	result *DeleteBalanceParameterStatusByUserIdResult
	err    error
}

func NewDeleteBalanceParameterStatusByUserIdResultFromJson(data string) DeleteBalanceParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewDeleteBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) DeleteBalanceParameterStatusByUserIdResult {
	return DeleteBalanceParameterStatusByUserIdResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteBalanceParameterStatusByUserIdResult) Pointer() *DeleteBalanceParameterStatusByUserIdResult {
	return &p
}

type ReDrawBalanceParameterStatusByUserIdResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Old      *BalanceParameterStatus `json:"old"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type ReDrawBalanceParameterStatusByUserIdAsyncResult struct {
	result *ReDrawBalanceParameterStatusByUserIdResult
	err    error
}

func NewReDrawBalanceParameterStatusByUserIdResultFromJson(data string) ReDrawBalanceParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewReDrawBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByUserIdResult {
	return ReDrawBalanceParameterStatusByUserIdResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *BalanceParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReDrawBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p ReDrawBalanceParameterStatusByUserIdResult) Pointer() *ReDrawBalanceParameterStatusByUserIdResult {
	return &p
}

type ReDrawBalanceParameterStatusByStampSheetResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Old      *BalanceParameterStatus `json:"old"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type ReDrawBalanceParameterStatusByStampSheetAsyncResult struct {
	result *ReDrawBalanceParameterStatusByStampSheetResult
	err    error
}

func NewReDrawBalanceParameterStatusByStampSheetResultFromJson(data string) ReDrawBalanceParameterStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawBalanceParameterStatusByStampSheetResultFromDict(dict)
}

func NewReDrawBalanceParameterStatusByStampSheetResultFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByStampSheetResult {
	return ReDrawBalanceParameterStatusByStampSheetResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *BalanceParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReDrawBalanceParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
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

func (p ReDrawBalanceParameterStatusByStampSheetResult) Pointer() *ReDrawBalanceParameterStatusByStampSheetResult {
	return &p
}

type SetBalanceParameterStatusByUserIdResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Old      *BalanceParameterStatus `json:"old"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type SetBalanceParameterStatusByUserIdAsyncResult struct {
	result *SetBalanceParameterStatusByUserIdResult
	err    error
}

func NewSetBalanceParameterStatusByUserIdResultFromJson(data string) SetBalanceParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewSetBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) SetBalanceParameterStatusByUserIdResult {
	return SetBalanceParameterStatusByUserIdResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *BalanceParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p SetBalanceParameterStatusByUserIdResult) Pointer() *SetBalanceParameterStatusByUserIdResult {
	return &p
}

type SetBalanceParameterStatusByStampSheetResult struct {
	Item     *BalanceParameterStatus `json:"item"`
	Old      *BalanceParameterStatus `json:"old"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type SetBalanceParameterStatusByStampSheetAsyncResult struct {
	result *SetBalanceParameterStatusByStampSheetResult
	err    error
}

func NewSetBalanceParameterStatusByStampSheetResultFromJson(data string) SetBalanceParameterStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetBalanceParameterStatusByStampSheetResultFromDict(dict)
}

func NewSetBalanceParameterStatusByStampSheetResultFromDict(data map[string]interface{}) SetBalanceParameterStatusByStampSheetResult {
	return SetBalanceParameterStatusByStampSheetResult{
		Item: func() *BalanceParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *BalanceParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetBalanceParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
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

func (p SetBalanceParameterStatusByStampSheetResult) Pointer() *SetBalanceParameterStatusByStampSheetResult {
	return &p
}

type DescribeRarityParameterStatusesResult struct {
	Items         []RarityParameterStatus `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeRarityParameterStatusesAsyncResult struct {
	result *DescribeRarityParameterStatusesResult
	err    error
}

func NewDescribeRarityParameterStatusesResultFromJson(data string) DescribeRarityParameterStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterStatusesResultFromDict(dict)
}

func NewDescribeRarityParameterStatusesResultFromDict(data map[string]interface{}) DescribeRarityParameterStatusesResult {
	return DescribeRarityParameterStatusesResult{
		Items: func() []RarityParameterStatus {
			if data["items"] == nil {
				return nil
			}
			return CastRarityParameterStatuses(core.CastArray(data["items"]))
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

func (p DescribeRarityParameterStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRarityParameterStatusesFromDict(
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

func (p DescribeRarityParameterStatusesResult) Pointer() *DescribeRarityParameterStatusesResult {
	return &p
}

type DescribeRarityParameterStatusesByUserIdResult struct {
	Items         []RarityParameterStatus `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeRarityParameterStatusesByUserIdAsyncResult struct {
	result *DescribeRarityParameterStatusesByUserIdResult
	err    error
}

func NewDescribeRarityParameterStatusesByUserIdResultFromJson(data string) DescribeRarityParameterStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRarityParameterStatusesByUserIdResultFromDict(dict)
}

func NewDescribeRarityParameterStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeRarityParameterStatusesByUserIdResult {
	return DescribeRarityParameterStatusesByUserIdResult{
		Items: func() []RarityParameterStatus {
			if data["items"] == nil {
				return nil
			}
			return CastRarityParameterStatuses(core.CastArray(data["items"]))
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

func (p DescribeRarityParameterStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRarityParameterStatusesFromDict(
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

func (p DescribeRarityParameterStatusesByUserIdResult) Pointer() *DescribeRarityParameterStatusesByUserIdResult {
	return &p
}

type GetRarityParameterStatusResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetRarityParameterStatusAsyncResult struct {
	result *GetRarityParameterStatusResult
	err    error
}

func NewGetRarityParameterStatusResultFromJson(data string) GetRarityParameterStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterStatusResultFromDict(dict)
}

func NewGetRarityParameterStatusResultFromDict(data map[string]interface{}) GetRarityParameterStatusResult {
	return GetRarityParameterStatusResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRarityParameterStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRarityParameterStatusResult) Pointer() *GetRarityParameterStatusResult {
	return &p
}

type GetRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetRarityParameterStatusByUserIdAsyncResult struct {
	result *GetRarityParameterStatusByUserIdResult
	err    error
}

func NewGetRarityParameterStatusByUserIdResultFromJson(data string) GetRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewGetRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) GetRarityParameterStatusByUserIdResult {
	return GetRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRarityParameterStatusByUserIdResult) Pointer() *GetRarityParameterStatusByUserIdResult {
	return &p
}

type DeleteRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DeleteRarityParameterStatusByUserIdAsyncResult struct {
	result *DeleteRarityParameterStatusByUserIdResult
	err    error
}

func NewDeleteRarityParameterStatusByUserIdResultFromJson(data string) DeleteRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewDeleteRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) DeleteRarityParameterStatusByUserIdResult {
	return DeleteRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteRarityParameterStatusByUserIdResult) Pointer() *DeleteRarityParameterStatusByUserIdResult {
	return &p
}

type ReDrawRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type ReDrawRarityParameterStatusByUserIdAsyncResult struct {
	result *ReDrawRarityParameterStatusByUserIdResult
	err    error
}

func NewReDrawRarityParameterStatusByUserIdResultFromJson(data string) ReDrawRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewReDrawRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByUserIdResult {
	return ReDrawRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReDrawRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p ReDrawRarityParameterStatusByUserIdResult) Pointer() *ReDrawRarityParameterStatusByUserIdResult {
	return &p
}

type ReDrawRarityParameterStatusByStampSheetResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type ReDrawRarityParameterStatusByStampSheetAsyncResult struct {
	result *ReDrawRarityParameterStatusByStampSheetResult
	err    error
}

func NewReDrawRarityParameterStatusByStampSheetResultFromJson(data string) ReDrawRarityParameterStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReDrawRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewReDrawRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByStampSheetResult {
	return ReDrawRarityParameterStatusByStampSheetResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ReDrawRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
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

func (p ReDrawRarityParameterStatusByStampSheetResult) Pointer() *ReDrawRarityParameterStatusByStampSheetResult {
	return &p
}

type AddRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type AddRarityParameterStatusByUserIdAsyncResult struct {
	result *AddRarityParameterStatusByUserIdResult
	err    error
}

func NewAddRarityParameterStatusByUserIdResultFromJson(data string) AddRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewAddRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) AddRarityParameterStatusByUserIdResult {
	return AddRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AddRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p AddRarityParameterStatusByUserIdResult) Pointer() *AddRarityParameterStatusByUserIdResult {
	return &p
}

type AddRarityParameterStatusByStampSheetResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type AddRarityParameterStatusByStampSheetAsyncResult struct {
	result *AddRarityParameterStatusByStampSheetResult
	err    error
}

func NewAddRarityParameterStatusByStampSheetResultFromJson(data string) AddRarityParameterStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewAddRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) AddRarityParameterStatusByStampSheetResult {
	return AddRarityParameterStatusByStampSheetResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AddRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
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

func (p AddRarityParameterStatusByStampSheetResult) Pointer() *AddRarityParameterStatusByStampSheetResult {
	return &p
}

type VerifyRarityParameterStatusResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type VerifyRarityParameterStatusAsyncResult struct {
	result *VerifyRarityParameterStatusResult
	err    error
}

func NewVerifyRarityParameterStatusResultFromJson(data string) VerifyRarityParameterStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusResultFromDict(dict)
}

func NewVerifyRarityParameterStatusResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusResult {
	return VerifyRarityParameterStatusResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyRarityParameterStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyRarityParameterStatusResult) Pointer() *VerifyRarityParameterStatusResult {
	return &p
}

type VerifyRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type VerifyRarityParameterStatusByUserIdAsyncResult struct {
	result *VerifyRarityParameterStatusByUserIdResult
	err    error
}

func NewVerifyRarityParameterStatusByUserIdResultFromJson(data string) VerifyRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewVerifyRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusByUserIdResult {
	return VerifyRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p VerifyRarityParameterStatusByUserIdResult) Pointer() *VerifyRarityParameterStatusByUserIdResult {
	return &p
}

type VerifyRarityParameterStatusByStampTaskResult struct {
	Item            *RarityParameterStatus `json:"item"`
	NewContextStack *string                `json:"newContextStack"`
	Metadata        *core.ResultMetadata   `json:"metadata"`
}

type VerifyRarityParameterStatusByStampTaskAsyncResult struct {
	result *VerifyRarityParameterStatusByStampTaskResult
	err    error
}

func NewVerifyRarityParameterStatusByStampTaskResultFromJson(data string) VerifyRarityParameterStatusByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyRarityParameterStatusByStampTaskResultFromDict(dict)
}

func NewVerifyRarityParameterStatusByStampTaskResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusByStampTaskResult {
	return VerifyRarityParameterStatusByStampTaskResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
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

func (p VerifyRarityParameterStatusByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyRarityParameterStatusByStampTaskResult) Pointer() *VerifyRarityParameterStatusByStampTaskResult {
	return &p
}

type SetRarityParameterStatusByUserIdResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type SetRarityParameterStatusByUserIdAsyncResult struct {
	result *SetRarityParameterStatusByUserIdResult
	err    error
}

func NewSetRarityParameterStatusByUserIdResultFromJson(data string) SetRarityParameterStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewSetRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) SetRarityParameterStatusByUserIdResult {
	return SetRarityParameterStatusByUserIdResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
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

func (p SetRarityParameterStatusByUserIdResult) Pointer() *SetRarityParameterStatusByUserIdResult {
	return &p
}

type SetRarityParameterStatusByStampSheetResult struct {
	Item     *RarityParameterStatus `json:"item"`
	Old      *RarityParameterStatus `json:"old"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type SetRarityParameterStatusByStampSheetAsyncResult struct {
	result *SetRarityParameterStatusByStampSheetResult
	err    error
}

func NewSetRarityParameterStatusByStampSheetResultFromJson(data string) SetRarityParameterStatusByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewSetRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) SetRarityParameterStatusByStampSheetResult {
	return SetRarityParameterStatusByStampSheetResult{
		Item: func() *RarityParameterStatus {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *RarityParameterStatus {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
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

func (p SetRarityParameterStatusByStampSheetResult) Pointer() *SetRarityParameterStatusByStampSheetResult {
	return &p
}
