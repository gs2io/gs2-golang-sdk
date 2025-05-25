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

package exchange

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

type DescribeRateModelsResult struct {
	Items    []RateModel          `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeRateModelsAsyncResult struct {
	result *DescribeRateModelsResult
	err    error
}

func NewDescribeRateModelsResultFromJson(data string) DescribeRateModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelsResultFromDict(dict)
}

func NewDescribeRateModelsResultFromDict(data map[string]interface{}) DescribeRateModelsResult {
	return DescribeRateModelsResult{
		Items: func() []RateModel {
			if data["items"] == nil {
				return nil
			}
			return CastRateModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeRateModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelsFromDict(
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

func (p DescribeRateModelsResult) Pointer() *DescribeRateModelsResult {
	return &p
}

type GetRateModelResult struct {
	Item     *RateModel           `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRateModelAsyncResult struct {
	result *GetRateModelResult
	err    error
}

func NewGetRateModelResultFromJson(data string) GetRateModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelResultFromDict(dict)
}

func NewGetRateModelResultFromDict(data map[string]interface{}) GetRateModelResult {
	return GetRateModelResult{
		Item: func() *RateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRateModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRateModelResult) Pointer() *GetRateModelResult {
	return &p
}

type DescribeRateModelMastersResult struct {
	Items         []RateModelMaster    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeRateModelMastersAsyncResult struct {
	result *DescribeRateModelMastersResult
	err    error
}

func NewDescribeRateModelMastersResultFromJson(data string) DescribeRateModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRateModelMastersResultFromDict(dict)
}

func NewDescribeRateModelMastersResultFromDict(data map[string]interface{}) DescribeRateModelMastersResult {
	return DescribeRateModelMastersResult{
		Items: func() []RateModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRateModelMasters(core.CastArray(data["items"]))
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

func (p DescribeRateModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelMastersFromDict(
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

func (p DescribeRateModelMastersResult) Pointer() *DescribeRateModelMastersResult {
	return &p
}

type CreateRateModelMasterResult struct {
	Item     *RateModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateRateModelMasterAsyncResult struct {
	result *CreateRateModelMasterResult
	err    error
}

func NewCreateRateModelMasterResultFromJson(data string) CreateRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRateModelMasterResultFromDict(dict)
}

func NewCreateRateModelMasterResultFromDict(data map[string]interface{}) CreateRateModelMasterResult {
	return CreateRateModelMasterResult{
		Item: func() *RateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateRateModelMasterResult) Pointer() *CreateRateModelMasterResult {
	return &p
}

type GetRateModelMasterResult struct {
	Item     *RateModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRateModelMasterAsyncResult struct {
	result *GetRateModelMasterResult
	err    error
}

func NewGetRateModelMasterResultFromJson(data string) GetRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRateModelMasterResultFromDict(dict)
}

func NewGetRateModelMasterResultFromDict(data map[string]interface{}) GetRateModelMasterResult {
	return GetRateModelMasterResult{
		Item: func() *RateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetRateModelMasterResult) Pointer() *GetRateModelMasterResult {
	return &p
}

type UpdateRateModelMasterResult struct {
	Item     *RateModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateRateModelMasterAsyncResult struct {
	result *UpdateRateModelMasterResult
	err    error
}

func NewUpdateRateModelMasterResultFromJson(data string) UpdateRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRateModelMasterResultFromDict(dict)
}

func NewUpdateRateModelMasterResultFromDict(data map[string]interface{}) UpdateRateModelMasterResult {
	return UpdateRateModelMasterResult{
		Item: func() *RateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateRateModelMasterResult) Pointer() *UpdateRateModelMasterResult {
	return &p
}

type DeleteRateModelMasterResult struct {
	Item     *RateModelMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteRateModelMasterAsyncResult struct {
	result *DeleteRateModelMasterResult
	err    error
}

func NewDeleteRateModelMasterResultFromJson(data string) DeleteRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRateModelMasterResultFromDict(dict)
}

func NewDeleteRateModelMasterResultFromDict(data map[string]interface{}) DeleteRateModelMasterResult {
	return DeleteRateModelMasterResult{
		Item: func() *RateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteRateModelMasterResult) Pointer() *DeleteRateModelMasterResult {
	return &p
}

type DescribeIncrementalRateModelsResult struct {
	Items    []IncrementalRateModel `json:"items"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DescribeIncrementalRateModelsAsyncResult struct {
	result *DescribeIncrementalRateModelsResult
	err    error
}

func NewDescribeIncrementalRateModelsResultFromJson(data string) DescribeIncrementalRateModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIncrementalRateModelsResultFromDict(dict)
}

func NewDescribeIncrementalRateModelsResultFromDict(data map[string]interface{}) DescribeIncrementalRateModelsResult {
	return DescribeIncrementalRateModelsResult{
		Items: func() []IncrementalRateModel {
			if data["items"] == nil {
				return nil
			}
			return CastIncrementalRateModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeIncrementalRateModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIncrementalRateModelsFromDict(
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

func (p DescribeIncrementalRateModelsResult) Pointer() *DescribeIncrementalRateModelsResult {
	return &p
}

type GetIncrementalRateModelResult struct {
	Item     *IncrementalRateModel `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetIncrementalRateModelAsyncResult struct {
	result *GetIncrementalRateModelResult
	err    error
}

func NewGetIncrementalRateModelResultFromJson(data string) GetIncrementalRateModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIncrementalRateModelResultFromDict(dict)
}

func NewGetIncrementalRateModelResultFromDict(data map[string]interface{}) GetIncrementalRateModelResult {
	return GetIncrementalRateModelResult{
		Item: func() *IncrementalRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetIncrementalRateModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetIncrementalRateModelResult) Pointer() *GetIncrementalRateModelResult {
	return &p
}

type DescribeIncrementalRateModelMastersResult struct {
	Items         []IncrementalRateModelMaster `json:"items"`
	NextPageToken *string                      `json:"nextPageToken"`
	Metadata      *core.ResultMetadata         `json:"metadata"`
}

type DescribeIncrementalRateModelMastersAsyncResult struct {
	result *DescribeIncrementalRateModelMastersResult
	err    error
}

func NewDescribeIncrementalRateModelMastersResultFromJson(data string) DescribeIncrementalRateModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeIncrementalRateModelMastersResultFromDict(dict)
}

func NewDescribeIncrementalRateModelMastersResultFromDict(data map[string]interface{}) DescribeIncrementalRateModelMastersResult {
	return DescribeIncrementalRateModelMastersResult{
		Items: func() []IncrementalRateModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastIncrementalRateModelMasters(core.CastArray(data["items"]))
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

func (p DescribeIncrementalRateModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastIncrementalRateModelMastersFromDict(
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

func (p DescribeIncrementalRateModelMastersResult) Pointer() *DescribeIncrementalRateModelMastersResult {
	return &p
}

type CreateIncrementalRateModelMasterResult struct {
	Item     *IncrementalRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type CreateIncrementalRateModelMasterAsyncResult struct {
	result *CreateIncrementalRateModelMasterResult
	err    error
}

func NewCreateIncrementalRateModelMasterResultFromJson(data string) CreateIncrementalRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateIncrementalRateModelMasterResultFromDict(dict)
}

func NewCreateIncrementalRateModelMasterResultFromDict(data map[string]interface{}) CreateIncrementalRateModelMasterResult {
	return CreateIncrementalRateModelMasterResult{
		Item: func() *IncrementalRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateIncrementalRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateIncrementalRateModelMasterResult) Pointer() *CreateIncrementalRateModelMasterResult {
	return &p
}

type GetIncrementalRateModelMasterResult struct {
	Item     *IncrementalRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type GetIncrementalRateModelMasterAsyncResult struct {
	result *GetIncrementalRateModelMasterResult
	err    error
}

func NewGetIncrementalRateModelMasterResultFromJson(data string) GetIncrementalRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetIncrementalRateModelMasterResultFromDict(dict)
}

func NewGetIncrementalRateModelMasterResultFromDict(data map[string]interface{}) GetIncrementalRateModelMasterResult {
	return GetIncrementalRateModelMasterResult{
		Item: func() *IncrementalRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetIncrementalRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetIncrementalRateModelMasterResult) Pointer() *GetIncrementalRateModelMasterResult {
	return &p
}

type UpdateIncrementalRateModelMasterResult struct {
	Item     *IncrementalRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type UpdateIncrementalRateModelMasterAsyncResult struct {
	result *UpdateIncrementalRateModelMasterResult
	err    error
}

func NewUpdateIncrementalRateModelMasterResultFromJson(data string) UpdateIncrementalRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateIncrementalRateModelMasterResultFromDict(dict)
}

func NewUpdateIncrementalRateModelMasterResultFromDict(data map[string]interface{}) UpdateIncrementalRateModelMasterResult {
	return UpdateIncrementalRateModelMasterResult{
		Item: func() *IncrementalRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateIncrementalRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateIncrementalRateModelMasterResult) Pointer() *UpdateIncrementalRateModelMasterResult {
	return &p
}

type DeleteIncrementalRateModelMasterResult struct {
	Item     *IncrementalRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata        `json:"metadata"`
}

type DeleteIncrementalRateModelMasterAsyncResult struct {
	result *DeleteIncrementalRateModelMasterResult
	err    error
}

func NewDeleteIncrementalRateModelMasterResultFromJson(data string) DeleteIncrementalRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteIncrementalRateModelMasterResultFromDict(dict)
}

func NewDeleteIncrementalRateModelMasterResultFromDict(data map[string]interface{}) DeleteIncrementalRateModelMasterResult {
	return DeleteIncrementalRateModelMasterResult{
		Item: func() *IncrementalRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteIncrementalRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteIncrementalRateModelMasterResult) Pointer() *DeleteIncrementalRateModelMasterResult {
	return &p
}

type ExchangeResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ExchangeAsyncResult struct {
	result *ExchangeResult
	err    error
}

func NewExchangeResultFromJson(data string) ExchangeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeResultFromDict(dict)
}

func NewExchangeResultFromDict(data map[string]interface{}) ExchangeResult {
	return ExchangeResult{
		Item: func() *RateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ExchangeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ExchangeResult) Pointer() *ExchangeResult {
	return &p
}

type ExchangeByUserIdResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ExchangeByUserIdAsyncResult struct {
	result *ExchangeByUserIdResult
	err    error
}

func NewExchangeByUserIdResultFromJson(data string) ExchangeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeByUserIdResultFromDict(dict)
}

func NewExchangeByUserIdResultFromDict(data map[string]interface{}) ExchangeByUserIdResult {
	return ExchangeByUserIdResult{
		Item: func() *RateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ExchangeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ExchangeByUserIdResult) Pointer() *ExchangeByUserIdResult {
	return &p
}

type ExchangeByStampSheetResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ExchangeByStampSheetAsyncResult struct {
	result *ExchangeByStampSheetResult
	err    error
}

func NewExchangeByStampSheetResultFromJson(data string) ExchangeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExchangeByStampSheetResultFromDict(dict)
}

func NewExchangeByStampSheetResultFromDict(data map[string]interface{}) ExchangeByStampSheetResult {
	return ExchangeByStampSheetResult{
		Item: func() *RateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ExchangeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p ExchangeByStampSheetResult) Pointer() *ExchangeByStampSheetResult {
	return &p
}

type IncrementalExchangeResult struct {
	Item                      *IncrementalRateModel   `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type IncrementalExchangeAsyncResult struct {
	result *IncrementalExchangeResult
	err    error
}

func NewIncrementalExchangeResultFromJson(data string) IncrementalExchangeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeResultFromDict(dict)
}

func NewIncrementalExchangeResultFromDict(data map[string]interface{}) IncrementalExchangeResult {
	return IncrementalExchangeResult{
		Item: func() *IncrementalRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p IncrementalExchangeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p IncrementalExchangeResult) Pointer() *IncrementalExchangeResult {
	return &p
}

type IncrementalExchangeByUserIdResult struct {
	Item                      *IncrementalRateModel   `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type IncrementalExchangeByUserIdAsyncResult struct {
	result *IncrementalExchangeByUserIdResult
	err    error
}

func NewIncrementalExchangeByUserIdResultFromJson(data string) IncrementalExchangeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeByUserIdResultFromDict(dict)
}

func NewIncrementalExchangeByUserIdResultFromDict(data map[string]interface{}) IncrementalExchangeByUserIdResult {
	return IncrementalExchangeByUserIdResult{
		Item: func() *IncrementalRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p IncrementalExchangeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p IncrementalExchangeByUserIdResult) Pointer() *IncrementalExchangeByUserIdResult {
	return &p
}

type IncrementalExchangeByStampSheetResult struct {
	Item                      *IncrementalRateModel   `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type IncrementalExchangeByStampSheetAsyncResult struct {
	result *IncrementalExchangeByStampSheetResult
	err    error
}

func NewIncrementalExchangeByStampSheetResultFromJson(data string) IncrementalExchangeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementalExchangeByStampSheetResultFromDict(dict)
}

func NewIncrementalExchangeByStampSheetResultFromDict(data map[string]interface{}) IncrementalExchangeByStampSheetResult {
	return IncrementalExchangeByStampSheetResult{
		Item: func() *IncrementalRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewIncrementalRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p IncrementalExchangeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p IncrementalExchangeByStampSheetResult) Pointer() *IncrementalExchangeByStampSheetResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentRateMaster   `json:"item"`
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
		Item: func() *CurrentRateMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentRateMasterResult struct {
	Item     *CurrentRateMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentRateMasterAsyncResult struct {
	result *GetCurrentRateMasterResult
	err    error
}

func NewGetCurrentRateMasterResultFromJson(data string) GetCurrentRateMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRateMasterResultFromDict(dict)
}

func NewGetCurrentRateMasterResultFromDict(data map[string]interface{}) GetCurrentRateMasterResult {
	return GetCurrentRateMasterResult{
		Item: func() *CurrentRateMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetCurrentRateMasterResult) Pointer() *GetCurrentRateMasterResult {
	return &p
}

type PreUpdateCurrentRateMasterResult struct {
	UploadToken *string              `json:"uploadToken"`
	UploadUrl   *string              `json:"uploadUrl"`
	Metadata    *core.ResultMetadata `json:"metadata"`
}

type PreUpdateCurrentRateMasterAsyncResult struct {
	result *PreUpdateCurrentRateMasterResult
	err    error
}

func NewPreUpdateCurrentRateMasterResultFromJson(data string) PreUpdateCurrentRateMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPreUpdateCurrentRateMasterResultFromDict(dict)
}

func NewPreUpdateCurrentRateMasterResultFromDict(data map[string]interface{}) PreUpdateCurrentRateMasterResult {
	return PreUpdateCurrentRateMasterResult{
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

func (p PreUpdateCurrentRateMasterResult) ToDict() map[string]interface{} {
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

func (p PreUpdateCurrentRateMasterResult) Pointer() *PreUpdateCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterResult struct {
	Item     *CurrentRateMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentRateMasterAsyncResult struct {
	result *UpdateCurrentRateMasterResult
	err    error
}

func NewUpdateCurrentRateMasterResultFromJson(data string) UpdateCurrentRateMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterResultFromDict(dict)
}

func NewUpdateCurrentRateMasterResultFromDict(data map[string]interface{}) UpdateCurrentRateMasterResult {
	return UpdateCurrentRateMasterResult{
		Item: func() *CurrentRateMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentRateMasterResult) Pointer() *UpdateCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterFromGitHubResult struct {
	Item     *CurrentRateMaster   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentRateMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRateMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRateMasterFromGitHubResultFromJson(data string) UpdateCurrentRateMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRateMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentRateMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRateMasterFromGitHubResult {
	return UpdateCurrentRateMasterFromGitHubResult{
		Item: func() *CurrentRateMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) Pointer() *UpdateCurrentRateMasterFromGitHubResult {
	return &p
}

type CreateAwaitByUserIdResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateAwaitByUserIdAsyncResult struct {
	result *CreateAwaitByUserIdResult
	err    error
}

func NewCreateAwaitByUserIdResultFromJson(data string) CreateAwaitByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAwaitByUserIdResultFromDict(dict)
}

func NewCreateAwaitByUserIdResultFromDict(data map[string]interface{}) CreateAwaitByUserIdResult {
	return CreateAwaitByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateAwaitByUserIdResult) Pointer() *CreateAwaitByUserIdResult {
	return &p
}

type DescribeAwaitsResult struct {
	Items         []Await              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAwaitsAsyncResult struct {
	result *DescribeAwaitsResult
	err    error
}

func NewDescribeAwaitsResultFromJson(data string) DescribeAwaitsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAwaitsResultFromDict(dict)
}

func NewDescribeAwaitsResultFromDict(data map[string]interface{}) DescribeAwaitsResult {
	return DescribeAwaitsResult{
		Items: func() []Await {
			if data["items"] == nil {
				return nil
			}
			return CastAwaits(core.CastArray(data["items"]))
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

func (p DescribeAwaitsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAwaitsFromDict(
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

func (p DescribeAwaitsResult) Pointer() *DescribeAwaitsResult {
	return &p
}

type DescribeAwaitsByUserIdResult struct {
	Items         []Await              `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeAwaitsByUserIdAsyncResult struct {
	result *DescribeAwaitsByUserIdResult
	err    error
}

func NewDescribeAwaitsByUserIdResultFromJson(data string) DescribeAwaitsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAwaitsByUserIdResultFromDict(dict)
}

func NewDescribeAwaitsByUserIdResultFromDict(data map[string]interface{}) DescribeAwaitsByUserIdResult {
	return DescribeAwaitsByUserIdResult{
		Items: func() []Await {
			if data["items"] == nil {
				return nil
			}
			return CastAwaits(core.CastArray(data["items"]))
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

func (p DescribeAwaitsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAwaitsFromDict(
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

func (p DescribeAwaitsByUserIdResult) Pointer() *DescribeAwaitsByUserIdResult {
	return &p
}

type GetAwaitResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAwaitAsyncResult struct {
	result *GetAwaitResult
	err    error
}

func NewGetAwaitResultFromJson(data string) GetAwaitResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAwaitResultFromDict(dict)
}

func NewGetAwaitResultFromDict(data map[string]interface{}) GetAwaitResult {
	return GetAwaitResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetAwaitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetAwaitResult) Pointer() *GetAwaitResult {
	return &p
}

type GetAwaitByUserIdResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetAwaitByUserIdAsyncResult struct {
	result *GetAwaitByUserIdResult
	err    error
}

func NewGetAwaitByUserIdResultFromJson(data string) GetAwaitByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAwaitByUserIdResultFromDict(dict)
}

func NewGetAwaitByUserIdResultFromDict(data map[string]interface{}) GetAwaitByUserIdResult {
	return GetAwaitByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetAwaitByUserIdResult) Pointer() *GetAwaitByUserIdResult {
	return &p
}

type AcquireResult struct {
	Item                      *Await                  `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type AcquireAsyncResult struct {
	result *AcquireResult
	err    error
}

func NewAcquireResultFromJson(data string) AcquireResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireResultFromDict(dict)
}

func NewAcquireResultFromDict(data map[string]interface{}) AcquireResult {
	return AcquireResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AcquireResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireResult) Pointer() *AcquireResult {
	return &p
}

type AcquireByUserIdResult struct {
	Item                      *Await                  `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type AcquireByUserIdAsyncResult struct {
	result *AcquireByUserIdResult
	err    error
}

func NewAcquireByUserIdResultFromJson(data string) AcquireByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireByUserIdResultFromDict(dict)
}

func NewAcquireByUserIdResultFromDict(data map[string]interface{}) AcquireByUserIdResult {
	return AcquireByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AcquireByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireByUserIdResult) Pointer() *AcquireByUserIdResult {
	return &p
}

type AcquireForceByUserIdResult struct {
	Item                      *Await                  `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type AcquireForceByUserIdAsyncResult struct {
	result *AcquireForceByUserIdResult
	err    error
}

func NewAcquireForceByUserIdResultFromJson(data string) AcquireForceByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireForceByUserIdResultFromDict(dict)
}

func NewAcquireForceByUserIdResultFromDict(data map[string]interface{}) AcquireForceByUserIdResult {
	return AcquireForceByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AcquireForceByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireForceByUserIdResult) Pointer() *AcquireForceByUserIdResult {
	return &p
}

type SkipByUserIdResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SkipByUserIdAsyncResult struct {
	result *SkipByUserIdResult
	err    error
}

func NewSkipByUserIdResultFromJson(data string) SkipByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSkipByUserIdResultFromDict(dict)
}

func NewSkipByUserIdResultFromDict(data map[string]interface{}) SkipByUserIdResult {
	return SkipByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SkipByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SkipByUserIdResult) Pointer() *SkipByUserIdResult {
	return &p
}

type DeleteAwaitResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAwaitAsyncResult struct {
	result *DeleteAwaitResult
	err    error
}

func NewDeleteAwaitResultFromJson(data string) DeleteAwaitResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitResultFromDict(dict)
}

func NewDeleteAwaitResultFromDict(data map[string]interface{}) DeleteAwaitResult {
	return DeleteAwaitResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteAwaitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteAwaitResult) Pointer() *DeleteAwaitResult {
	return &p
}

type DeleteAwaitByUserIdResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteAwaitByUserIdAsyncResult struct {
	result *DeleteAwaitByUserIdResult
	err    error
}

func NewDeleteAwaitByUserIdResultFromJson(data string) DeleteAwaitByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitByUserIdResultFromDict(dict)
}

func NewDeleteAwaitByUserIdResultFromDict(data map[string]interface{}) DeleteAwaitByUserIdResult {
	return DeleteAwaitByUserIdResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteAwaitByUserIdResult) Pointer() *DeleteAwaitByUserIdResult {
	return &p
}

type CreateAwaitByStampSheetResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateAwaitByStampSheetAsyncResult struct {
	result *CreateAwaitByStampSheetResult
	err    error
}

func NewCreateAwaitByStampSheetResultFromJson(data string) CreateAwaitByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAwaitByStampSheetResultFromDict(dict)
}

func NewCreateAwaitByStampSheetResultFromDict(data map[string]interface{}) CreateAwaitByStampSheetResult {
	return CreateAwaitByStampSheetResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateAwaitByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateAwaitByStampSheetResult) Pointer() *CreateAwaitByStampSheetResult {
	return &p
}

type AcquireForceByStampSheetResult struct {
	Item                      *Await                  `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type AcquireForceByStampSheetAsyncResult struct {
	result *AcquireForceByStampSheetResult
	err    error
}

func NewAcquireForceByStampSheetResultFromJson(data string) AcquireForceByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcquireForceByStampSheetResultFromDict(dict)
}

func NewAcquireForceByStampSheetResultFromDict(data map[string]interface{}) AcquireForceByStampSheetResult {
	return AcquireForceByStampSheetResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		StampSheet: func() *string {
			v, ok := data["stampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheet"])
		}(),
		StampSheetEncryptionKeyId: func() *string {
			v, ok := data["stampSheetEncryptionKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stampSheetEncryptionKeyId"])
		}(),
		AutoRunStampSheet: func() *bool {
			v, ok := data["autoRunStampSheet"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["autoRunStampSheet"])
		}(),
		AtomicCommit: func() *bool {
			v, ok := data["atomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["atomicCommit"])
		}(),
		Transaction: func() *string {
			v, ok := data["transaction"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transaction"])
		}(),
		TransactionResult: func() *core.TransactionResult {
			v, ok := data["transactionResult"]
			if !ok || v == nil {
				return nil
			}
			return core.NewTransactionResultFromDict(core.CastMap(data["transactionResult"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AcquireForceByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"transactionId":             p.TransactionId,
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
		"autoRunStampSheet":         p.AutoRunStampSheet,
		"atomicCommit":              p.AtomicCommit,
		"transaction":               p.Transaction,
		"transactionResult": func() map[string]interface{} {
			if p.TransactionResult == nil {
				return nil
			}
			return p.TransactionResult.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p AcquireForceByStampSheetResult) Pointer() *AcquireForceByStampSheetResult {
	return &p
}

type SkipByStampSheetResult struct {
	Item     *Await               `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type SkipByStampSheetAsyncResult struct {
	result *SkipByStampSheetResult
	err    error
}

func NewSkipByStampSheetResultFromJson(data string) SkipByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSkipByStampSheetResultFromDict(dict)
}

func NewSkipByStampSheetResultFromDict(data map[string]interface{}) SkipByStampSheetResult {
	return SkipByStampSheetResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SkipByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p SkipByStampSheetResult) Pointer() *SkipByStampSheetResult {
	return &p
}

type DeleteAwaitByStampTaskResult struct {
	Item            *Await               `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type DeleteAwaitByStampTaskAsyncResult struct {
	result *DeleteAwaitByStampTaskResult
	err    error
}

func NewDeleteAwaitByStampTaskResultFromJson(data string) DeleteAwaitByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAwaitByStampTaskResultFromDict(dict)
}

func NewDeleteAwaitByStampTaskResultFromDict(data map[string]interface{}) DeleteAwaitByStampTaskResult {
	return DeleteAwaitByStampTaskResult{
		Item: func() *Await {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewAwaitFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteAwaitByStampTaskResult) ToDict() map[string]interface{} {
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

func (p DeleteAwaitByStampTaskResult) Pointer() *DeleteAwaitByStampTaskResult {
	return &p
}
