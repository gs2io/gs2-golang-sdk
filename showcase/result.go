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

package showcase

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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
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
		},
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeSalesItemMastersResult struct {
	Items         []SalesItemMaster    `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSalesItemMastersAsyncResult struct {
	result *DescribeSalesItemMastersResult
	err    error
}

func NewDescribeSalesItemMastersResultFromJson(data string) DescribeSalesItemMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSalesItemMastersResultFromDict(dict)
}

func NewDescribeSalesItemMastersResultFromDict(data map[string]interface{}) DescribeSalesItemMastersResult {
	return DescribeSalesItemMastersResult{
		Items: func() []SalesItemMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSalesItemMasters(core.CastArray(data["items"]))
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

func (p DescribeSalesItemMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSalesItemMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeSalesItemMastersResult) Pointer() *DescribeSalesItemMastersResult {
	return &p
}

type CreateSalesItemMasterResult struct {
	Item     *SalesItemMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateSalesItemMasterAsyncResult struct {
	result *CreateSalesItemMasterResult
	err    error
}

func NewCreateSalesItemMasterResultFromJson(data string) CreateSalesItemMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSalesItemMasterResultFromDict(dict)
}

func NewCreateSalesItemMasterResultFromDict(data map[string]interface{}) CreateSalesItemMasterResult {
	return CreateSalesItemMasterResult{
		Item: func() *SalesItemMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateSalesItemMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p CreateSalesItemMasterResult) Pointer() *CreateSalesItemMasterResult {
	return &p
}

type GetSalesItemMasterResult struct {
	Item     *SalesItemMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSalesItemMasterAsyncResult struct {
	result *GetSalesItemMasterResult
	err    error
}

func NewGetSalesItemMasterResultFromJson(data string) GetSalesItemMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSalesItemMasterResultFromDict(dict)
}

func NewGetSalesItemMasterResultFromDict(data map[string]interface{}) GetSalesItemMasterResult {
	return GetSalesItemMasterResult{
		Item: func() *SalesItemMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSalesItemMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetSalesItemMasterResult) Pointer() *GetSalesItemMasterResult {
	return &p
}

type UpdateSalesItemMasterResult struct {
	Item     *SalesItemMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateSalesItemMasterAsyncResult struct {
	result *UpdateSalesItemMasterResult
	err    error
}

func NewUpdateSalesItemMasterResultFromJson(data string) UpdateSalesItemMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSalesItemMasterResultFromDict(dict)
}

func NewUpdateSalesItemMasterResultFromDict(data map[string]interface{}) UpdateSalesItemMasterResult {
	return UpdateSalesItemMasterResult{
		Item: func() *SalesItemMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateSalesItemMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateSalesItemMasterResult) Pointer() *UpdateSalesItemMasterResult {
	return &p
}

type DeleteSalesItemMasterResult struct {
	Item     *SalesItemMaster     `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSalesItemMasterAsyncResult struct {
	result *DeleteSalesItemMasterResult
	err    error
}

func NewDeleteSalesItemMasterResultFromJson(data string) DeleteSalesItemMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSalesItemMasterResultFromDict(dict)
}

func NewDeleteSalesItemMasterResultFromDict(data map[string]interface{}) DeleteSalesItemMasterResult {
	return DeleteSalesItemMasterResult{
		Item: func() *SalesItemMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteSalesItemMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DeleteSalesItemMasterResult) Pointer() *DeleteSalesItemMasterResult {
	return &p
}

type DescribeSalesItemGroupMastersResult struct {
	Items         []SalesItemGroupMaster `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeSalesItemGroupMastersAsyncResult struct {
	result *DescribeSalesItemGroupMastersResult
	err    error
}

func NewDescribeSalesItemGroupMastersResultFromJson(data string) DescribeSalesItemGroupMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSalesItemGroupMastersResultFromDict(dict)
}

func NewDescribeSalesItemGroupMastersResultFromDict(data map[string]interface{}) DescribeSalesItemGroupMastersResult {
	return DescribeSalesItemGroupMastersResult{
		Items: func() []SalesItemGroupMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSalesItemGroupMasters(core.CastArray(data["items"]))
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

func (p DescribeSalesItemGroupMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSalesItemGroupMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeSalesItemGroupMastersResult) Pointer() *DescribeSalesItemGroupMastersResult {
	return &p
}

type CreateSalesItemGroupMasterResult struct {
	Item     *SalesItemGroupMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type CreateSalesItemGroupMasterAsyncResult struct {
	result *CreateSalesItemGroupMasterResult
	err    error
}

func NewCreateSalesItemGroupMasterResultFromJson(data string) CreateSalesItemGroupMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSalesItemGroupMasterResultFromDict(dict)
}

func NewCreateSalesItemGroupMasterResultFromDict(data map[string]interface{}) CreateSalesItemGroupMasterResult {
	return CreateSalesItemGroupMasterResult{
		Item: func() *SalesItemGroupMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p CreateSalesItemGroupMasterResult) Pointer() *CreateSalesItemGroupMasterResult {
	return &p
}

type GetSalesItemGroupMasterResult struct {
	Item     *SalesItemGroupMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetSalesItemGroupMasterAsyncResult struct {
	result *GetSalesItemGroupMasterResult
	err    error
}

func NewGetSalesItemGroupMasterResultFromJson(data string) GetSalesItemGroupMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSalesItemGroupMasterResultFromDict(dict)
}

func NewGetSalesItemGroupMasterResultFromDict(data map[string]interface{}) GetSalesItemGroupMasterResult {
	return GetSalesItemGroupMasterResult{
		Item: func() *SalesItemGroupMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetSalesItemGroupMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetSalesItemGroupMasterResult) Pointer() *GetSalesItemGroupMasterResult {
	return &p
}

type UpdateSalesItemGroupMasterResult struct {
	Item     *SalesItemGroupMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateSalesItemGroupMasterAsyncResult struct {
	result *UpdateSalesItemGroupMasterResult
	err    error
}

func NewUpdateSalesItemGroupMasterResultFromJson(data string) UpdateSalesItemGroupMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSalesItemGroupMasterResultFromDict(dict)
}

func NewUpdateSalesItemGroupMasterResultFromDict(data map[string]interface{}) UpdateSalesItemGroupMasterResult {
	return UpdateSalesItemGroupMasterResult{
		Item: func() *SalesItemGroupMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateSalesItemGroupMasterResult) Pointer() *UpdateSalesItemGroupMasterResult {
	return &p
}

type DeleteSalesItemGroupMasterResult struct {
	Item     *SalesItemGroupMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type DeleteSalesItemGroupMasterAsyncResult struct {
	result *DeleteSalesItemGroupMasterResult
	err    error
}

func NewDeleteSalesItemGroupMasterResultFromJson(data string) DeleteSalesItemGroupMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSalesItemGroupMasterResultFromDict(dict)
}

func NewDeleteSalesItemGroupMasterResultFromDict(data map[string]interface{}) DeleteSalesItemGroupMasterResult {
	return DeleteSalesItemGroupMasterResult{
		Item: func() *SalesItemGroupMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteSalesItemGroupMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DeleteSalesItemGroupMasterResult) Pointer() *DeleteSalesItemGroupMasterResult {
	return &p
}

type DescribeShowcaseMastersResult struct {
	Items         []ShowcaseMaster     `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeShowcaseMastersAsyncResult struct {
	result *DescribeShowcaseMastersResult
	err    error
}

func NewDescribeShowcaseMastersResultFromJson(data string) DescribeShowcaseMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcaseMastersResultFromDict(dict)
}

func NewDescribeShowcaseMastersResultFromDict(data map[string]interface{}) DescribeShowcaseMastersResult {
	return DescribeShowcaseMastersResult{
		Items: func() []ShowcaseMaster {
			if data["items"] == nil {
				return nil
			}
			return CastShowcaseMasters(core.CastArray(data["items"]))
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

func (p DescribeShowcaseMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcaseMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeShowcaseMastersResult) Pointer() *DescribeShowcaseMastersResult {
	return &p
}

type CreateShowcaseMasterResult struct {
	Item     *ShowcaseMaster      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateShowcaseMasterAsyncResult struct {
	result *CreateShowcaseMasterResult
	err    error
}

func NewCreateShowcaseMasterResultFromJson(data string) CreateShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateShowcaseMasterResultFromDict(dict)
}

func NewCreateShowcaseMasterResultFromDict(data map[string]interface{}) CreateShowcaseMasterResult {
	return CreateShowcaseMasterResult{
		Item: func() *ShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p CreateShowcaseMasterResult) Pointer() *CreateShowcaseMasterResult {
	return &p
}

type GetShowcaseMasterResult struct {
	Item     *ShowcaseMaster      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetShowcaseMasterAsyncResult struct {
	result *GetShowcaseMasterResult
	err    error
}

func NewGetShowcaseMasterResultFromJson(data string) GetShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseMasterResultFromDict(dict)
}

func NewGetShowcaseMasterResultFromDict(data map[string]interface{}) GetShowcaseMasterResult {
	return GetShowcaseMasterResult{
		Item: func() *ShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetShowcaseMasterResult) Pointer() *GetShowcaseMasterResult {
	return &p
}

type UpdateShowcaseMasterResult struct {
	Item     *ShowcaseMaster      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateShowcaseMasterAsyncResult struct {
	result *UpdateShowcaseMasterResult
	err    error
}

func NewUpdateShowcaseMasterResultFromJson(data string) UpdateShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateShowcaseMasterResultFromDict(dict)
}

func NewUpdateShowcaseMasterResultFromDict(data map[string]interface{}) UpdateShowcaseMasterResult {
	return UpdateShowcaseMasterResult{
		Item: func() *ShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateShowcaseMasterResult) Pointer() *UpdateShowcaseMasterResult {
	return &p
}

type DeleteShowcaseMasterResult struct {
	Item     *ShowcaseMaster      `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteShowcaseMasterAsyncResult struct {
	result *DeleteShowcaseMasterResult
	err    error
}

func NewDeleteShowcaseMasterResultFromJson(data string) DeleteShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteShowcaseMasterResultFromDict(dict)
}

func NewDeleteShowcaseMasterResultFromDict(data map[string]interface{}) DeleteShowcaseMasterResult {
	return DeleteShowcaseMasterResult{
		Item: func() *ShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DeleteShowcaseMasterResult) Pointer() *DeleteShowcaseMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
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
		Item: func() *CurrentShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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
		},
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentShowcaseMasterResult struct {
	Item     *CurrentShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetCurrentShowcaseMasterAsyncResult struct {
	result *GetCurrentShowcaseMasterResult
	err    error
}

func NewGetCurrentShowcaseMasterResultFromJson(data string) GetCurrentShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentShowcaseMasterResultFromDict(dict)
}

func NewGetCurrentShowcaseMasterResultFromDict(data map[string]interface{}) GetCurrentShowcaseMasterResult {
	return GetCurrentShowcaseMasterResult{
		Item: func() *CurrentShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetCurrentShowcaseMasterResult) Pointer() *GetCurrentShowcaseMasterResult {
	return &p
}

type UpdateCurrentShowcaseMasterResult struct {
	Item     *CurrentShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateCurrentShowcaseMasterAsyncResult struct {
	result *UpdateCurrentShowcaseMasterResult
	err    error
}

func NewUpdateCurrentShowcaseMasterResultFromJson(data string) UpdateCurrentShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentShowcaseMasterResultFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterResult {
	return UpdateCurrentShowcaseMasterResult{
		Item: func() *CurrentShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateCurrentShowcaseMasterResult) Pointer() *UpdateCurrentShowcaseMasterResult {
	return &p
}

type UpdateCurrentShowcaseMasterFromGitHubResult struct {
	Item     *CurrentShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type UpdateCurrentShowcaseMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentShowcaseMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentShowcaseMasterFromGitHubResultFromJson(data string) UpdateCurrentShowcaseMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentShowcaseMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterFromGitHubResult {
	return UpdateCurrentShowcaseMasterFromGitHubResult{
		Item: func() *CurrentShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) Pointer() *UpdateCurrentShowcaseMasterFromGitHubResult {
	return &p
}

type DescribeShowcasesResult struct {
	Items    []Showcase           `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeShowcasesAsyncResult struct {
	result *DescribeShowcasesResult
	err    error
}

func NewDescribeShowcasesResultFromJson(data string) DescribeShowcasesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcasesResultFromDict(dict)
}

func NewDescribeShowcasesResultFromDict(data map[string]interface{}) DescribeShowcasesResult {
	return DescribeShowcasesResult{
		Items: func() []Showcase {
			if data["items"] == nil {
				return nil
			}
			return CastShowcases(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeShowcasesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcasesFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeShowcasesResult) Pointer() *DescribeShowcasesResult {
	return &p
}

type DescribeShowcasesByUserIdResult struct {
	Items    []Showcase           `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeShowcasesByUserIdAsyncResult struct {
	result *DescribeShowcasesByUserIdResult
	err    error
}

func NewDescribeShowcasesByUserIdResultFromJson(data string) DescribeShowcasesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeShowcasesByUserIdResultFromDict(dict)
}

func NewDescribeShowcasesByUserIdResultFromDict(data map[string]interface{}) DescribeShowcasesByUserIdResult {
	return DescribeShowcasesByUserIdResult{
		Items: func() []Showcase {
			if data["items"] == nil {
				return nil
			}
			return CastShowcases(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeShowcasesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastShowcasesFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeShowcasesByUserIdResult) Pointer() *DescribeShowcasesByUserIdResult {
	return &p
}

type GetShowcaseResult struct {
	Item     *Showcase            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetShowcaseAsyncResult struct {
	result *GetShowcaseResult
	err    error
}

func NewGetShowcaseResultFromJson(data string) GetShowcaseResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseResultFromDict(dict)
}

func NewGetShowcaseResultFromDict(data map[string]interface{}) GetShowcaseResult {
	return GetShowcaseResult{
		Item: func() *Showcase {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetShowcaseResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetShowcaseResult) Pointer() *GetShowcaseResult {
	return &p
}

type GetShowcaseByUserIdResult struct {
	Item     *Showcase            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetShowcaseByUserIdAsyncResult struct {
	result *GetShowcaseByUserIdResult
	err    error
}

func NewGetShowcaseByUserIdResultFromJson(data string) GetShowcaseByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetShowcaseByUserIdResultFromDict(dict)
}

func NewGetShowcaseByUserIdResultFromDict(data map[string]interface{}) GetShowcaseByUserIdResult {
	return GetShowcaseByUserIdResult{
		Item: func() *Showcase {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewShowcaseFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetShowcaseByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetShowcaseByUserIdResult) Pointer() *GetShowcaseByUserIdResult {
	return &p
}

type BuyResult struct {
	Item                      *SalesItem              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type BuyAsyncResult struct {
	result *BuyResult
	err    error
}

func NewBuyResultFromJson(data string) BuyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuyResultFromDict(dict)
}

func NewBuyResultFromDict(data map[string]interface{}) BuyResult {
	return BuyResult{
		Item: func() *SalesItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p BuyResult) ToDict() map[string]interface{} {
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
		},
	}
}

func (p BuyResult) Pointer() *BuyResult {
	return &p
}

type BuyByUserIdResult struct {
	Item                      *SalesItem              `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type BuyByUserIdAsyncResult struct {
	result *BuyByUserIdResult
	err    error
}

func NewBuyByUserIdResultFromJson(data string) BuyByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewBuyByUserIdResultFromDict(dict)
}

func NewBuyByUserIdResultFromDict(data map[string]interface{}) BuyByUserIdResult {
	return BuyByUserIdResult{
		Item: func() *SalesItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSalesItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p BuyByUserIdResult) ToDict() map[string]interface{} {
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
		},
	}
}

func (p BuyByUserIdResult) Pointer() *BuyByUserIdResult {
	return &p
}

type DescribeRandomShowcaseMastersResult struct {
	Items         []RandomShowcaseMaster `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeRandomShowcaseMastersAsyncResult struct {
	result *DescribeRandomShowcaseMastersResult
	err    error
}

func NewDescribeRandomShowcaseMastersResultFromJson(data string) DescribeRandomShowcaseMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomShowcaseMastersResultFromDict(dict)
}

func NewDescribeRandomShowcaseMastersResultFromDict(data map[string]interface{}) DescribeRandomShowcaseMastersResult {
	return DescribeRandomShowcaseMastersResult{
		Items: func() []RandomShowcaseMaster {
			if data["items"] == nil {
				return nil
			}
			return CastRandomShowcaseMasters(core.CastArray(data["items"]))
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

func (p DescribeRandomShowcaseMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRandomShowcaseMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeRandomShowcaseMastersResult) Pointer() *DescribeRandomShowcaseMastersResult {
	return &p
}

type CreateRandomShowcaseMasterResult struct {
	Item     *RandomShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type CreateRandomShowcaseMasterAsyncResult struct {
	result *CreateRandomShowcaseMasterResult
	err    error
}

func NewCreateRandomShowcaseMasterResultFromJson(data string) CreateRandomShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRandomShowcaseMasterResultFromDict(dict)
}

func NewCreateRandomShowcaseMasterResultFromDict(data map[string]interface{}) CreateRandomShowcaseMasterResult {
	return CreateRandomShowcaseMasterResult{
		Item: func() *RandomShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateRandomShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p CreateRandomShowcaseMasterResult) Pointer() *CreateRandomShowcaseMasterResult {
	return &p
}

type GetRandomShowcaseMasterResult struct {
	Item     *RandomShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetRandomShowcaseMasterAsyncResult struct {
	result *GetRandomShowcaseMasterResult
	err    error
}

func NewGetRandomShowcaseMasterResultFromJson(data string) GetRandomShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomShowcaseMasterResultFromDict(dict)
}

func NewGetRandomShowcaseMasterResultFromDict(data map[string]interface{}) GetRandomShowcaseMasterResult {
	return GetRandomShowcaseMasterResult{
		Item: func() *RandomShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRandomShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetRandomShowcaseMasterResult) Pointer() *GetRandomShowcaseMasterResult {
	return &p
}

type UpdateRandomShowcaseMasterResult struct {
	Item     *RandomShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateRandomShowcaseMasterAsyncResult struct {
	result *UpdateRandomShowcaseMasterResult
	err    error
}

func NewUpdateRandomShowcaseMasterResultFromJson(data string) UpdateRandomShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRandomShowcaseMasterResultFromDict(dict)
}

func NewUpdateRandomShowcaseMasterResultFromDict(data map[string]interface{}) UpdateRandomShowcaseMasterResult {
	return UpdateRandomShowcaseMasterResult{
		Item: func() *RandomShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateRandomShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p UpdateRandomShowcaseMasterResult) Pointer() *UpdateRandomShowcaseMasterResult {
	return &p
}

type DeleteRandomShowcaseMasterResult struct {
	Item     *RandomShowcaseMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type DeleteRandomShowcaseMasterAsyncResult struct {
	result *DeleteRandomShowcaseMasterResult
	err    error
}

func NewDeleteRandomShowcaseMasterResultFromJson(data string) DeleteRandomShowcaseMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRandomShowcaseMasterResultFromDict(dict)
}

func NewDeleteRandomShowcaseMasterResultFromDict(data map[string]interface{}) DeleteRandomShowcaseMasterResult {
	return DeleteRandomShowcaseMasterResult{
		Item: func() *RandomShowcaseMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteRandomShowcaseMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DeleteRandomShowcaseMasterResult) Pointer() *DeleteRandomShowcaseMasterResult {
	return &p
}

type IncrementPurchaseCountResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IncrementPurchaseCountAsyncResult struct {
	result *IncrementPurchaseCountResult
	err    error
}

func NewIncrementPurchaseCountResultFromJson(data string) IncrementPurchaseCountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementPurchaseCountResultFromDict(dict)
}

func NewIncrementPurchaseCountResultFromDict(data map[string]interface{}) IncrementPurchaseCountResult {
	return IncrementPurchaseCountResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p IncrementPurchaseCountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p IncrementPurchaseCountResult) Pointer() *IncrementPurchaseCountResult {
	return &p
}

type IncrementPurchaseCountByUserIdResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type IncrementPurchaseCountByUserIdAsyncResult struct {
	result *IncrementPurchaseCountByUserIdResult
	err    error
}

func NewIncrementPurchaseCountByUserIdResultFromJson(data string) IncrementPurchaseCountByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementPurchaseCountByUserIdResultFromDict(dict)
}

func NewIncrementPurchaseCountByUserIdResultFromDict(data map[string]interface{}) IncrementPurchaseCountByUserIdResult {
	return IncrementPurchaseCountByUserIdResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p IncrementPurchaseCountByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p IncrementPurchaseCountByUserIdResult) Pointer() *IncrementPurchaseCountByUserIdResult {
	return &p
}

type DecrementPurchaseCountByUserIdResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DecrementPurchaseCountByUserIdAsyncResult struct {
	result *DecrementPurchaseCountByUserIdResult
	err    error
}

func NewDecrementPurchaseCountByUserIdResultFromJson(data string) DecrementPurchaseCountByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecrementPurchaseCountByUserIdResultFromDict(dict)
}

func NewDecrementPurchaseCountByUserIdResultFromDict(data map[string]interface{}) DecrementPurchaseCountByUserIdResult {
	return DecrementPurchaseCountByUserIdResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DecrementPurchaseCountByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DecrementPurchaseCountByUserIdResult) Pointer() *DecrementPurchaseCountByUserIdResult {
	return &p
}

type IncrementPurchaseCountByStampTaskResult struct {
	Item            *RandomDisplayItem   `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type IncrementPurchaseCountByStampTaskAsyncResult struct {
	result *IncrementPurchaseCountByStampTaskResult
	err    error
}

func NewIncrementPurchaseCountByStampTaskResultFromJson(data string) IncrementPurchaseCountByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncrementPurchaseCountByStampTaskResultFromDict(dict)
}

func NewIncrementPurchaseCountByStampTaskResultFromDict(data map[string]interface{}) IncrementPurchaseCountByStampTaskResult {
	return IncrementPurchaseCountByStampTaskResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p IncrementPurchaseCountByStampTaskResult) ToDict() map[string]interface{} {
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
		},
	}
}

func (p IncrementPurchaseCountByStampTaskResult) Pointer() *IncrementPurchaseCountByStampTaskResult {
	return &p
}

type DecrementPurchaseCountByStampSheetResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DecrementPurchaseCountByStampSheetAsyncResult struct {
	result *DecrementPurchaseCountByStampSheetResult
	err    error
}

func NewDecrementPurchaseCountByStampSheetResultFromJson(data string) DecrementPurchaseCountByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecrementPurchaseCountByStampSheetResultFromDict(dict)
}

func NewDecrementPurchaseCountByStampSheetResultFromDict(data map[string]interface{}) DecrementPurchaseCountByStampSheetResult {
	return DecrementPurchaseCountByStampSheetResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DecrementPurchaseCountByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DecrementPurchaseCountByStampSheetResult) Pointer() *DecrementPurchaseCountByStampSheetResult {
	return &p
}

type ForceReDrawByUserIdResult struct {
	Items    []RandomDisplayItem  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ForceReDrawByUserIdAsyncResult struct {
	result *ForceReDrawByUserIdResult
	err    error
}

func NewForceReDrawByUserIdResultFromJson(data string) ForceReDrawByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForceReDrawByUserIdResultFromDict(dict)
}

func NewForceReDrawByUserIdResultFromDict(data map[string]interface{}) ForceReDrawByUserIdResult {
	return ForceReDrawByUserIdResult{
		Items: func() []RandomDisplayItem {
			if data["items"] == nil {
				return nil
			}
			return CastRandomDisplayItems(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ForceReDrawByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRandomDisplayItemsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p ForceReDrawByUserIdResult) Pointer() *ForceReDrawByUserIdResult {
	return &p
}

type ForceReDrawByUserIdByStampSheetResult struct {
	Items    []RandomDisplayItem  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ForceReDrawByUserIdByStampSheetAsyncResult struct {
	result *ForceReDrawByUserIdByStampSheetResult
	err    error
}

func NewForceReDrawByUserIdByStampSheetResultFromJson(data string) ForceReDrawByUserIdByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewForceReDrawByUserIdByStampSheetResultFromDict(dict)
}

func NewForceReDrawByUserIdByStampSheetResultFromDict(data map[string]interface{}) ForceReDrawByUserIdByStampSheetResult {
	return ForceReDrawByUserIdByStampSheetResult{
		Items: func() []RandomDisplayItem {
			if data["items"] == nil {
				return nil
			}
			return CastRandomDisplayItems(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ForceReDrawByUserIdByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRandomDisplayItemsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p ForceReDrawByUserIdByStampSheetResult) Pointer() *ForceReDrawByUserIdByStampSheetResult {
	return &p
}

type DescribeRandomDisplayItemsResult struct {
	Items    []RandomDisplayItem  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeRandomDisplayItemsAsyncResult struct {
	result *DescribeRandomDisplayItemsResult
	err    error
}

func NewDescribeRandomDisplayItemsResultFromJson(data string) DescribeRandomDisplayItemsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomDisplayItemsResultFromDict(dict)
}

func NewDescribeRandomDisplayItemsResultFromDict(data map[string]interface{}) DescribeRandomDisplayItemsResult {
	return DescribeRandomDisplayItemsResult{
		Items: func() []RandomDisplayItem {
			if data["items"] == nil {
				return nil
			}
			return CastRandomDisplayItems(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeRandomDisplayItemsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRandomDisplayItemsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeRandomDisplayItemsResult) Pointer() *DescribeRandomDisplayItemsResult {
	return &p
}

type DescribeRandomDisplayItemsByUserIdResult struct {
	Items    []RandomDisplayItem  `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeRandomDisplayItemsByUserIdAsyncResult struct {
	result *DescribeRandomDisplayItemsByUserIdResult
	err    error
}

func NewDescribeRandomDisplayItemsByUserIdResultFromJson(data string) DescribeRandomDisplayItemsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRandomDisplayItemsByUserIdResultFromDict(dict)
}

func NewDescribeRandomDisplayItemsByUserIdResultFromDict(data map[string]interface{}) DescribeRandomDisplayItemsByUserIdResult {
	return DescribeRandomDisplayItemsByUserIdResult{
		Items: func() []RandomDisplayItem {
			if data["items"] == nil {
				return nil
			}
			return CastRandomDisplayItems(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeRandomDisplayItemsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRandomDisplayItemsFromDict(
			p.Items,
		),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p DescribeRandomDisplayItemsByUserIdResult) Pointer() *DescribeRandomDisplayItemsByUserIdResult {
	return &p
}

type GetRandomDisplayItemResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRandomDisplayItemAsyncResult struct {
	result *GetRandomDisplayItemResult
	err    error
}

func NewGetRandomDisplayItemResultFromJson(data string) GetRandomDisplayItemResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomDisplayItemResultFromDict(dict)
}

func NewGetRandomDisplayItemResultFromDict(data map[string]interface{}) GetRandomDisplayItemResult {
	return GetRandomDisplayItemResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRandomDisplayItemResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetRandomDisplayItemResult) Pointer() *GetRandomDisplayItemResult {
	return &p
}

type GetRandomDisplayItemByUserIdResult struct {
	Item     *RandomDisplayItem   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetRandomDisplayItemByUserIdAsyncResult struct {
	result *GetRandomDisplayItemByUserIdResult
	err    error
}

func NewGetRandomDisplayItemByUserIdResultFromJson(data string) GetRandomDisplayItemByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRandomDisplayItemByUserIdResultFromDict(dict)
}

func NewGetRandomDisplayItemByUserIdResultFromDict(data map[string]interface{}) GetRandomDisplayItemByUserIdResult {
	return GetRandomDisplayItemByUserIdResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetRandomDisplayItemByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p GetRandomDisplayItemByUserIdResult) Pointer() *GetRandomDisplayItemByUserIdResult {
	return &p
}

type RandomShowcaseBuyResult struct {
	Item                      *RandomDisplayItem      `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type RandomShowcaseBuyAsyncResult struct {
	result *RandomShowcaseBuyResult
	err    error
}

func NewRandomShowcaseBuyResultFromJson(data string) RandomShowcaseBuyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseBuyResultFromDict(dict)
}

func NewRandomShowcaseBuyResultFromDict(data map[string]interface{}) RandomShowcaseBuyResult {
	return RandomShowcaseBuyResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p RandomShowcaseBuyResult) ToDict() map[string]interface{} {
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
		},
	}
}

func (p RandomShowcaseBuyResult) Pointer() *RandomShowcaseBuyResult {
	return &p
}

type RandomShowcaseBuyByUserIdResult struct {
	Item                      *RandomDisplayItem      `json:"item"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type RandomShowcaseBuyByUserIdAsyncResult struct {
	result *RandomShowcaseBuyByUserIdResult
	err    error
}

func NewRandomShowcaseBuyByUserIdResultFromJson(data string) RandomShowcaseBuyByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRandomShowcaseBuyByUserIdResultFromDict(dict)
}

func NewRandomShowcaseBuyByUserIdResultFromDict(data map[string]interface{}) RandomShowcaseBuyByUserIdResult {
	return RandomShowcaseBuyByUserIdResult{
		Item: func() *RandomDisplayItem {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer()
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

func (p RandomShowcaseBuyByUserIdResult) ToDict() map[string]interface{} {
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
		},
	}
}

func (p RandomShowcaseBuyByUserIdResult) Pointer() *RandomShowcaseBuyByUserIdResult {
	return &p
}
