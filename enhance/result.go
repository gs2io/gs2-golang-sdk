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

package enhance

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

type DescribeUnleashRateModelsResult struct {
	Items    []UnleashRateModel   `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeUnleashRateModelsAsyncResult struct {
	result *DescribeUnleashRateModelsResult
	err    error
}

func NewDescribeUnleashRateModelsResultFromJson(data string) DescribeUnleashRateModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeUnleashRateModelsResultFromDict(dict)
}

func NewDescribeUnleashRateModelsResultFromDict(data map[string]interface{}) DescribeUnleashRateModelsResult {
	return DescribeUnleashRateModelsResult{
		Items: func() []UnleashRateModel {
			if data["items"] == nil {
				return nil
			}
			return CastUnleashRateModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeUnleashRateModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastUnleashRateModelsFromDict(
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

func (p DescribeUnleashRateModelsResult) Pointer() *DescribeUnleashRateModelsResult {
	return &p
}

type GetUnleashRateModelResult struct {
	Item     *UnleashRateModel    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetUnleashRateModelAsyncResult struct {
	result *GetUnleashRateModelResult
	err    error
}

func NewGetUnleashRateModelResultFromJson(data string) GetUnleashRateModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetUnleashRateModelResultFromDict(dict)
}

func NewGetUnleashRateModelResultFromDict(data map[string]interface{}) GetUnleashRateModelResult {
	return GetUnleashRateModelResult{
		Item: func() *UnleashRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetUnleashRateModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetUnleashRateModelResult) Pointer() *GetUnleashRateModelResult {
	return &p
}

type DescribeUnleashRateModelMastersResult struct {
	Items         []UnleashRateModelMaster `json:"items"`
	NextPageToken *string                  `json:"nextPageToken"`
	Metadata      *core.ResultMetadata     `json:"metadata"`
}

type DescribeUnleashRateModelMastersAsyncResult struct {
	result *DescribeUnleashRateModelMastersResult
	err    error
}

func NewDescribeUnleashRateModelMastersResultFromJson(data string) DescribeUnleashRateModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeUnleashRateModelMastersResultFromDict(dict)
}

func NewDescribeUnleashRateModelMastersResultFromDict(data map[string]interface{}) DescribeUnleashRateModelMastersResult {
	return DescribeUnleashRateModelMastersResult{
		Items: func() []UnleashRateModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastUnleashRateModelMasters(core.CastArray(data["items"]))
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

func (p DescribeUnleashRateModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastUnleashRateModelMastersFromDict(
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

func (p DescribeUnleashRateModelMastersResult) Pointer() *DescribeUnleashRateModelMastersResult {
	return &p
}

type CreateUnleashRateModelMasterResult struct {
	Item     *UnleashRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type CreateUnleashRateModelMasterAsyncResult struct {
	result *CreateUnleashRateModelMasterResult
	err    error
}

func NewCreateUnleashRateModelMasterResultFromJson(data string) CreateUnleashRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateUnleashRateModelMasterResultFromDict(dict)
}

func NewCreateUnleashRateModelMasterResultFromDict(data map[string]interface{}) CreateUnleashRateModelMasterResult {
	return CreateUnleashRateModelMasterResult{
		Item: func() *UnleashRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateUnleashRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateUnleashRateModelMasterResult) Pointer() *CreateUnleashRateModelMasterResult {
	return &p
}

type GetUnleashRateModelMasterResult struct {
	Item     *UnleashRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type GetUnleashRateModelMasterAsyncResult struct {
	result *GetUnleashRateModelMasterResult
	err    error
}

func NewGetUnleashRateModelMasterResultFromJson(data string) GetUnleashRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetUnleashRateModelMasterResultFromDict(dict)
}

func NewGetUnleashRateModelMasterResultFromDict(data map[string]interface{}) GetUnleashRateModelMasterResult {
	return GetUnleashRateModelMasterResult{
		Item: func() *UnleashRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetUnleashRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetUnleashRateModelMasterResult) Pointer() *GetUnleashRateModelMasterResult {
	return &p
}

type UpdateUnleashRateModelMasterResult struct {
	Item     *UnleashRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type UpdateUnleashRateModelMasterAsyncResult struct {
	result *UpdateUnleashRateModelMasterResult
	err    error
}

func NewUpdateUnleashRateModelMasterResultFromJson(data string) UpdateUnleashRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateUnleashRateModelMasterResultFromDict(dict)
}

func NewUpdateUnleashRateModelMasterResultFromDict(data map[string]interface{}) UpdateUnleashRateModelMasterResult {
	return UpdateUnleashRateModelMasterResult{
		Item: func() *UnleashRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateUnleashRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p UpdateUnleashRateModelMasterResult) Pointer() *UpdateUnleashRateModelMasterResult {
	return &p
}

type DeleteUnleashRateModelMasterResult struct {
	Item     *UnleashRateModelMaster `json:"item"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type DeleteUnleashRateModelMasterAsyncResult struct {
	result *DeleteUnleashRateModelMasterResult
	err    error
}

func NewDeleteUnleashRateModelMasterResultFromJson(data string) DeleteUnleashRateModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteUnleashRateModelMasterResultFromDict(dict)
}

func NewDeleteUnleashRateModelMasterResultFromDict(data map[string]interface{}) DeleteUnleashRateModelMasterResult {
	return DeleteUnleashRateModelMasterResult{
		Item: func() *UnleashRateModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteUnleashRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteUnleashRateModelMasterResult) Pointer() *DeleteUnleashRateModelMasterResult {
	return &p
}

type DirectEnhanceResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	AcquireExperience         *int64                  `json:"acquireExperience"`
	BonusRate                 *float32                `json:"bonusRate"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DirectEnhanceAsyncResult struct {
	result *DirectEnhanceResult
	err    error
}

func NewDirectEnhanceResultFromJson(data string) DirectEnhanceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceResultFromDict(dict)
}

func NewDirectEnhanceResultFromDict(data map[string]interface{}) DirectEnhanceResult {
	return DirectEnhanceResult{
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
		AcquireExperience: func() *int64 {
			v, ok := data["acquireExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["acquireExperience"])
		}(),
		BonusRate: func() *float32 {
			v, ok := data["bonusRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["bonusRate"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DirectEnhanceResult) ToDict() map[string]interface{} {
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
		"acquireExperience": p.AcquireExperience,
		"bonusRate":         p.BonusRate,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DirectEnhanceResult) Pointer() *DirectEnhanceResult {
	return &p
}

type DirectEnhanceByUserIdResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	AcquireExperience         *int64                  `json:"acquireExperience"`
	BonusRate                 *float32                `json:"bonusRate"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DirectEnhanceByUserIdAsyncResult struct {
	result *DirectEnhanceByUserIdResult
	err    error
}

func NewDirectEnhanceByUserIdResultFromJson(data string) DirectEnhanceByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceByUserIdResultFromDict(dict)
}

func NewDirectEnhanceByUserIdResultFromDict(data map[string]interface{}) DirectEnhanceByUserIdResult {
	return DirectEnhanceByUserIdResult{
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
		AcquireExperience: func() *int64 {
			v, ok := data["acquireExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["acquireExperience"])
		}(),
		BonusRate: func() *float32 {
			v, ok := data["bonusRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["bonusRate"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DirectEnhanceByUserIdResult) ToDict() map[string]interface{} {
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
		"acquireExperience": p.AcquireExperience,
		"bonusRate":         p.BonusRate,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DirectEnhanceByUserIdResult) Pointer() *DirectEnhanceByUserIdResult {
	return &p
}

type DirectEnhanceByStampSheetResult struct {
	Item                      *RateModel              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	AcquireExperience         *int64                  `json:"acquireExperience"`
	BonusRate                 *float32                `json:"bonusRate"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DirectEnhanceByStampSheetAsyncResult struct {
	result *DirectEnhanceByStampSheetResult
	err    error
}

func NewDirectEnhanceByStampSheetResultFromJson(data string) DirectEnhanceByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDirectEnhanceByStampSheetResultFromDict(dict)
}

func NewDirectEnhanceByStampSheetResultFromDict(data map[string]interface{}) DirectEnhanceByStampSheetResult {
	return DirectEnhanceByStampSheetResult{
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
		AcquireExperience: func() *int64 {
			v, ok := data["acquireExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["acquireExperience"])
		}(),
		BonusRate: func() *float32 {
			v, ok := data["bonusRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["bonusRate"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DirectEnhanceByStampSheetResult) ToDict() map[string]interface{} {
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
		"acquireExperience": p.AcquireExperience,
		"bonusRate":         p.BonusRate,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DirectEnhanceByStampSheetResult) Pointer() *DirectEnhanceByStampSheetResult {
	return &p
}

type UnleashResult struct {
	Item                      *UnleashRateModel       `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type UnleashAsyncResult struct {
	result *UnleashResult
	err    error
}

func NewUnleashResultFromJson(data string) UnleashResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashResultFromDict(dict)
}

func NewUnleashResultFromDict(data map[string]interface{}) UnleashResult {
	return UnleashResult{
		Item: func() *UnleashRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UnleashResult) ToDict() map[string]interface{} {
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

func (p UnleashResult) Pointer() *UnleashResult {
	return &p
}

type UnleashByUserIdResult struct {
	Item                      *UnleashRateModel       `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type UnleashByUserIdAsyncResult struct {
	result *UnleashByUserIdResult
	err    error
}

func NewUnleashByUserIdResultFromJson(data string) UnleashByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashByUserIdResultFromDict(dict)
}

func NewUnleashByUserIdResultFromDict(data map[string]interface{}) UnleashByUserIdResult {
	return UnleashByUserIdResult{
		Item: func() *UnleashRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UnleashByUserIdResult) ToDict() map[string]interface{} {
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

func (p UnleashByUserIdResult) Pointer() *UnleashByUserIdResult {
	return &p
}

type UnleashByStampSheetResult struct {
	Item                      *UnleashRateModel       `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type UnleashByStampSheetAsyncResult struct {
	result *UnleashByStampSheetResult
	err    error
}

func NewUnleashByStampSheetResultFromJson(data string) UnleashByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnleashByStampSheetResultFromDict(dict)
}

func NewUnleashByStampSheetResultFromDict(data map[string]interface{}) UnleashByStampSheetResult {
	return UnleashByStampSheetResult{
		Item: func() *UnleashRateModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewUnleashRateModelFromDict(core.CastMap(data["item"])).Pointer()
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

func (p UnleashByStampSheetResult) ToDict() map[string]interface{} {
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

func (p UnleashByStampSheetResult) Pointer() *UnleashByStampSheetResult {
	return &p
}

type CreateProgressByUserIdResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateProgressByUserIdAsyncResult struct {
	result *CreateProgressByUserIdResult
	err    error
}

func NewCreateProgressByUserIdResultFromJson(data string) CreateProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProgressByUserIdResultFromDict(dict)
}

func NewCreateProgressByUserIdResultFromDict(data map[string]interface{}) CreateProgressByUserIdResult {
	return CreateProgressByUserIdResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateProgressByUserIdResult) Pointer() *CreateProgressByUserIdResult {
	return &p
}

type GetProgressResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

func NewGetProgressResultFromJson(data string) GetProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProgressResultFromDict(dict)
}

func NewGetProgressResultFromDict(data map[string]interface{}) GetProgressResult {
	return GetProgressResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetProgressResult) Pointer() *GetProgressResult {
	return &p
}

type GetProgressByUserIdResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetProgressByUserIdAsyncResult struct {
	result *GetProgressByUserIdResult
	err    error
}

func NewGetProgressByUserIdResultFromJson(data string) GetProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetProgressByUserIdResultFromDict(dict)
}

func NewGetProgressByUserIdResultFromDict(data map[string]interface{}) GetProgressByUserIdResult {
	return GetProgressByUserIdResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p GetProgressByUserIdResult) Pointer() *GetProgressByUserIdResult {
	return &p
}

type StartResult struct {
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type StartAsyncResult struct {
	result *StartResult
	err    error
}

func NewStartResultFromJson(data string) StartResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartResultFromDict(dict)
}

func NewStartResultFromDict(data map[string]interface{}) StartResult {
	return StartResult{
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

func (p StartResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func (p StartResult) Pointer() *StartResult {
	return &p
}

type StartByUserIdResult struct {
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type StartByUserIdAsyncResult struct {
	result *StartByUserIdResult
	err    error
}

func NewStartByUserIdResultFromJson(data string) StartByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewStartByUserIdResultFromDict(dict)
}

func NewStartByUserIdResultFromDict(data map[string]interface{}) StartByUserIdResult {
	return StartByUserIdResult{
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

func (p StartByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func (p StartByUserIdResult) Pointer() *StartByUserIdResult {
	return &p
}

type EndResult struct {
	Item                      *Progress               `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	AcquireExperience         *int64                  `json:"acquireExperience"`
	BonusRate                 *float32                `json:"bonusRate"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type EndAsyncResult struct {
	result *EndResult
	err    error
}

func NewEndResultFromJson(data string) EndResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEndResultFromDict(dict)
}

func NewEndResultFromDict(data map[string]interface{}) EndResult {
	return EndResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
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
		AcquireExperience: func() *int64 {
			v, ok := data["acquireExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["acquireExperience"])
		}(),
		BonusRate: func() *float32 {
			v, ok := data["bonusRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["bonusRate"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p EndResult) ToDict() map[string]interface{} {
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
		"acquireExperience": p.AcquireExperience,
		"bonusRate":         p.BonusRate,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p EndResult) Pointer() *EndResult {
	return &p
}

type EndByUserIdResult struct {
	Item                      *Progress               `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	AcquireExperience         *int64                  `json:"acquireExperience"`
	BonusRate                 *float32                `json:"bonusRate"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type EndByUserIdAsyncResult struct {
	result *EndByUserIdResult
	err    error
}

func NewEndByUserIdResultFromJson(data string) EndByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEndByUserIdResultFromDict(dict)
}

func NewEndByUserIdResultFromDict(data map[string]interface{}) EndByUserIdResult {
	return EndByUserIdResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
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
		AcquireExperience: func() *int64 {
			v, ok := data["acquireExperience"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["acquireExperience"])
		}(),
		BonusRate: func() *float32 {
			v, ok := data["bonusRate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["bonusRate"])
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p EndByUserIdResult) ToDict() map[string]interface{} {
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
		"acquireExperience": p.AcquireExperience,
		"bonusRate":         p.BonusRate,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p EndByUserIdResult) Pointer() *EndByUserIdResult {
	return &p
}

type DeleteProgressResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteProgressAsyncResult struct {
	result *DeleteProgressResult
	err    error
}

func NewDeleteProgressResultFromJson(data string) DeleteProgressResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressResultFromDict(dict)
}

func NewDeleteProgressResultFromDict(data map[string]interface{}) DeleteProgressResult {
	return DeleteProgressResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteProgressResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteProgressResult) Pointer() *DeleteProgressResult {
	return &p
}

type DeleteProgressByUserIdResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteProgressByUserIdAsyncResult struct {
	result *DeleteProgressByUserIdResult
	err    error
}

func NewDeleteProgressByUserIdResultFromJson(data string) DeleteProgressByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressByUserIdResultFromDict(dict)
}

func NewDeleteProgressByUserIdResultFromDict(data map[string]interface{}) DeleteProgressByUserIdResult {
	return DeleteProgressByUserIdResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteProgressByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p DeleteProgressByUserIdResult) Pointer() *DeleteProgressByUserIdResult {
	return &p
}

type CreateProgressByStampSheetResult struct {
	Item     *Progress            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateProgressByStampSheetAsyncResult struct {
	result *CreateProgressByStampSheetResult
	err    error
}

func NewCreateProgressByStampSheetResultFromJson(data string) CreateProgressByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateProgressByStampSheetResultFromDict(dict)
}

func NewCreateProgressByStampSheetResultFromDict(data map[string]interface{}) CreateProgressByStampSheetResult {
	return CreateProgressByStampSheetResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateProgressByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		}(),
	}
}

func (p CreateProgressByStampSheetResult) Pointer() *CreateProgressByStampSheetResult {
	return &p
}

type DeleteProgressByStampTaskResult struct {
	Item            *Progress            `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type DeleteProgressByStampTaskAsyncResult struct {
	result *DeleteProgressByStampTaskResult
	err    error
}

func NewDeleteProgressByStampTaskResultFromJson(data string) DeleteProgressByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteProgressByStampTaskResultFromDict(dict)
}

func NewDeleteProgressByStampTaskResultFromDict(data map[string]interface{}) DeleteProgressByStampTaskResult {
	return DeleteProgressByStampTaskResult{
		Item: func() *Progress {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewProgressFromDict(core.CastMap(data["item"])).Pointer()
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

func (p DeleteProgressByStampTaskResult) ToDict() map[string]interface{} {
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

func (p DeleteProgressByStampTaskResult) Pointer() *DeleteProgressByStampTaskResult {
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
