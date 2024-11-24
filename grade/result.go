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

package grade

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
	"github.com/gs2io/gs2-golang-sdk/experience"
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

type DescribeGradeModelMastersResult struct {
	Items         []GradeModelMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGradeModelMastersAsyncResult struct {
	result *DescribeGradeModelMastersResult
	err    error
}

func NewDescribeGradeModelMastersResultFromJson(data string) DescribeGradeModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelMastersResultFromDict(dict)
}

func NewDescribeGradeModelMastersResultFromDict(data map[string]interface{}) DescribeGradeModelMastersResult {
	return DescribeGradeModelMastersResult{
		Items: func() []GradeModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastGradeModelMasters(core.CastArray(data["items"]))
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

func (p DescribeGradeModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGradeModelMastersFromDict(
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

func (p DescribeGradeModelMastersResult) Pointer() *DescribeGradeModelMastersResult {
	return &p
}

type CreateGradeModelMasterResult struct {
	Item     *GradeModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateGradeModelMasterAsyncResult struct {
	result *CreateGradeModelMasterResult
	err    error
}

func NewCreateGradeModelMasterResultFromJson(data string) CreateGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGradeModelMasterResultFromDict(dict)
}

func NewCreateGradeModelMasterResultFromDict(data map[string]interface{}) CreateGradeModelMasterResult {
	return CreateGradeModelMasterResult{
		Item: func() *GradeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p CreateGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p CreateGradeModelMasterResult) Pointer() *CreateGradeModelMasterResult {
	return &p
}

type GetGradeModelMasterResult struct {
	Item     *GradeModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGradeModelMasterAsyncResult struct {
	result *GetGradeModelMasterResult
	err    error
}

func NewGetGradeModelMasterResultFromJson(data string) GetGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelMasterResultFromDict(dict)
}

func NewGetGradeModelMasterResultFromDict(data map[string]interface{}) GetGradeModelMasterResult {
	return GetGradeModelMasterResult{
		Item: func() *GradeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p GetGradeModelMasterResult) Pointer() *GetGradeModelMasterResult {
	return &p
}

type UpdateGradeModelMasterResult struct {
	Item     *GradeModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateGradeModelMasterAsyncResult struct {
	result *UpdateGradeModelMasterResult
	err    error
}

func NewUpdateGradeModelMasterResultFromJson(data string) UpdateGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGradeModelMasterResultFromDict(dict)
}

func NewUpdateGradeModelMasterResultFromDict(data map[string]interface{}) UpdateGradeModelMasterResult {
	return UpdateGradeModelMasterResult{
		Item: func() *GradeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p UpdateGradeModelMasterResult) Pointer() *UpdateGradeModelMasterResult {
	return &p
}

type DeleteGradeModelMasterResult struct {
	Item     *GradeModelMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGradeModelMasterAsyncResult struct {
	result *DeleteGradeModelMasterResult
	err    error
}

func NewDeleteGradeModelMasterResultFromJson(data string) DeleteGradeModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGradeModelMasterResultFromDict(dict)
}

func NewDeleteGradeModelMasterResultFromDict(data map[string]interface{}) DeleteGradeModelMasterResult {
	return DeleteGradeModelMasterResult{
		Item: func() *GradeModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGradeModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteGradeModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p DeleteGradeModelMasterResult) Pointer() *DeleteGradeModelMasterResult {
	return &p
}

type DescribeGradeModelsResult struct {
	Items    []GradeModel         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeGradeModelsAsyncResult struct {
	result *DescribeGradeModelsResult
	err    error
}

func NewDescribeGradeModelsResultFromJson(data string) DescribeGradeModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGradeModelsResultFromDict(dict)
}

func NewDescribeGradeModelsResultFromDict(data map[string]interface{}) DescribeGradeModelsResult {
	return DescribeGradeModelsResult{
		Items: func() []GradeModel {
			if data["items"] == nil {
				return nil
			}
			return CastGradeModels(core.CastArray(data["items"]))
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DescribeGradeModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGradeModelsFromDict(
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

func (p DescribeGradeModelsResult) Pointer() *DescribeGradeModelsResult {
	return &p
}

type GetGradeModelResult struct {
	Item     *GradeModel          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGradeModelAsyncResult struct {
	result *GetGradeModelResult
	err    error
}

func NewGetGradeModelResultFromJson(data string) GetGradeModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGradeModelResultFromDict(dict)
}

func NewGetGradeModelResultFromDict(data map[string]interface{}) GetGradeModelResult {
	return GetGradeModelResult{
		Item: func() *GradeModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGradeModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetGradeModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p GetGradeModelResult) Pointer() *GetGradeModelResult {
	return &p
}

type DescribeStatusesResult struct {
	Items         []Status             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStatusesAsyncResult struct {
	result *DescribeStatusesResult
	err    error
}

func NewDescribeStatusesResultFromJson(data string) DescribeStatusesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesResultFromDict(dict)
}

func NewDescribeStatusesResultFromDict(data map[string]interface{}) DescribeStatusesResult {
	return DescribeStatusesResult{
		Items: func() []Status {
			if data["items"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["items"]))
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

func (p DescribeStatusesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
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

func (p DescribeStatusesResult) Pointer() *DescribeStatusesResult {
	return &p
}

type DescribeStatusesByUserIdResult struct {
	Items         []Status             `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeStatusesByUserIdAsyncResult struct {
	result *DescribeStatusesByUserIdResult
	err    error
}

func NewDescribeStatusesByUserIdResultFromJson(data string) DescribeStatusesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStatusesByUserIdResultFromDict(dict)
}

func NewDescribeStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeStatusesByUserIdResult {
	return DescribeStatusesByUserIdResult{
		Items: func() []Status {
			if data["items"] == nil {
				return nil
			}
			return CastStatuses(core.CastArray(data["items"]))
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

func (p DescribeStatusesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStatusesFromDict(
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

func (p DescribeStatusesByUserIdResult) Pointer() *DescribeStatusesByUserIdResult {
	return &p
}

type GetStatusResult struct {
	Item     *Status              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStatusAsyncResult struct {
	result *GetStatusResult
	err    error
}

func NewGetStatusResultFromJson(data string) GetStatusResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusResultFromDict(dict)
}

func NewGetStatusResultFromDict(data map[string]interface{}) GetStatusResult {
	return GetStatusResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p GetStatusResult) Pointer() *GetStatusResult {
	return &p
}

type GetStatusByUserIdResult struct {
	Item     *Status              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetStatusByUserIdAsyncResult struct {
	result *GetStatusByUserIdResult
	err    error
}

func NewGetStatusByUserIdResultFromJson(data string) GetStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStatusByUserIdResultFromDict(dict)
}

func NewGetStatusByUserIdResultFromDict(data map[string]interface{}) GetStatusByUserIdResult {
	return GetStatusByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p GetStatusByUserIdResult) Pointer() *GetStatusByUserIdResult {
	return &p
}

type AddGradeByUserIdResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type AddGradeByUserIdAsyncResult struct {
	result *AddGradeByUserIdResult
	err    error
}

func NewAddGradeByUserIdResultFromJson(data string) AddGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByUserIdResultFromDict(dict)
}

func NewAddGradeByUserIdResultFromDict(data map[string]interface{}) AddGradeByUserIdResult {
	return AddGradeByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AddGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p AddGradeByUserIdResult) Pointer() *AddGradeByUserIdResult {
	return &p
}

type SubGradeResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type SubGradeAsyncResult struct {
	result *SubGradeResult
	err    error
}

func NewSubGradeResultFromJson(data string) SubGradeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeResultFromDict(dict)
}

func NewSubGradeResultFromDict(data map[string]interface{}) SubGradeResult {
	return SubGradeResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SubGradeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p SubGradeResult) Pointer() *SubGradeResult {
	return &p
}

type SubGradeByUserIdResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type SubGradeByUserIdAsyncResult struct {
	result *SubGradeByUserIdResult
	err    error
}

func NewSubGradeByUserIdResultFromJson(data string) SubGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByUserIdResultFromDict(dict)
}

func NewSubGradeByUserIdResultFromDict(data map[string]interface{}) SubGradeByUserIdResult {
	return SubGradeByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SubGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p SubGradeByUserIdResult) Pointer() *SubGradeByUserIdResult {
	return &p
}

type SetGradeByUserIdResult struct {
	Item                    *Status              `json:"item"`
	Old                     *Status              `json:"old"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type SetGradeByUserIdAsyncResult struct {
	result *SetGradeByUserIdResult
	err    error
}

func NewSetGradeByUserIdResultFromJson(data string) SetGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetGradeByUserIdResultFromDict(dict)
}

func NewSetGradeByUserIdResultFromDict(data map[string]interface{}) SetGradeByUserIdResult {
	return SetGradeByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Old: func() *Status {
			v, ok := data["old"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["old"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SetGradeByUserIdResult) ToDict() map[string]interface{} {
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
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p SetGradeByUserIdResult) Pointer() *SetGradeByUserIdResult {
	return &p
}

type ApplyRankCapResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type ApplyRankCapAsyncResult struct {
	result *ApplyRankCapResult
	err    error
}

func NewApplyRankCapResultFromJson(data string) ApplyRankCapResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapResultFromDict(dict)
}

func NewApplyRankCapResultFromDict(data map[string]interface{}) ApplyRankCapResult {
	return ApplyRankCapResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ApplyRankCapResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p ApplyRankCapResult) Pointer() *ApplyRankCapResult {
	return &p
}

type ApplyRankCapByUserIdResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type ApplyRankCapByUserIdAsyncResult struct {
	result *ApplyRankCapByUserIdResult
	err    error
}

func NewApplyRankCapByUserIdResultFromJson(data string) ApplyRankCapByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByUserIdResultFromDict(dict)
}

func NewApplyRankCapByUserIdResultFromDict(data map[string]interface{}) ApplyRankCapByUserIdResult {
	return ApplyRankCapByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ApplyRankCapByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p ApplyRankCapByUserIdResult) Pointer() *ApplyRankCapByUserIdResult {
	return &p
}

type DeleteStatusByUserIdResult struct {
	Item     *Status              `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteStatusByUserIdAsyncResult struct {
	result *DeleteStatusByUserIdResult
	err    error
}

func NewDeleteStatusByUserIdResultFromJson(data string) DeleteStatusByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStatusByUserIdResultFromDict(dict)
}

func NewDeleteStatusByUserIdResultFromDict(data map[string]interface{}) DeleteStatusByUserIdResult {
	return DeleteStatusByUserIdResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p DeleteStatusByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p DeleteStatusByUserIdResult) Pointer() *DeleteStatusByUserIdResult {
	return &p
}

type VerifyGradeResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeAsyncResult struct {
	result *VerifyGradeResult
	err    error
}

func NewVerifyGradeResultFromJson(data string) VerifyGradeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeResultFromDict(dict)
}

func NewVerifyGradeResultFromDict(data map[string]interface{}) VerifyGradeResult {
	return VerifyGradeResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyGradeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeResult) Pointer() *VerifyGradeResult {
	return &p
}

type VerifyGradeByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeByUserIdAsyncResult struct {
	result *VerifyGradeByUserIdResult
	err    error
}

func NewVerifyGradeByUserIdResultFromJson(data string) VerifyGradeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByUserIdResultFromDict(dict)
}

func NewVerifyGradeByUserIdResultFromDict(data map[string]interface{}) VerifyGradeByUserIdResult {
	return VerifyGradeByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyGradeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeByUserIdResult) Pointer() *VerifyGradeByUserIdResult {
	return &p
}

type VerifyGradeUpMaterialResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeUpMaterialAsyncResult struct {
	result *VerifyGradeUpMaterialResult
	err    error
}

func NewVerifyGradeUpMaterialResultFromJson(data string) VerifyGradeUpMaterialResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialResultFromDict(dict)
}

func NewVerifyGradeUpMaterialResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialResult {
	return VerifyGradeUpMaterialResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyGradeUpMaterialResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeUpMaterialResult) Pointer() *VerifyGradeUpMaterialResult {
	return &p
}

type VerifyGradeUpMaterialByUserIdResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeUpMaterialByUserIdAsyncResult struct {
	result *VerifyGradeUpMaterialByUserIdResult
	err    error
}

func NewVerifyGradeUpMaterialByUserIdResultFromJson(data string) VerifyGradeUpMaterialByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByUserIdResultFromDict(dict)
}

func NewVerifyGradeUpMaterialByUserIdResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialByUserIdResult {
	return VerifyGradeUpMaterialByUserIdResult{
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p VerifyGradeUpMaterialByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeUpMaterialByUserIdResult) Pointer() *VerifyGradeUpMaterialByUserIdResult {
	return &p
}

type AddGradeByStampSheetResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type AddGradeByStampSheetAsyncResult struct {
	result *AddGradeByStampSheetResult
	err    error
}

func NewAddGradeByStampSheetResultFromJson(data string) AddGradeByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddGradeByStampSheetResultFromDict(dict)
}

func NewAddGradeByStampSheetResultFromDict(data map[string]interface{}) AddGradeByStampSheetResult {
	return AddGradeByStampSheetResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p AddGradeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p AddGradeByStampSheetResult) Pointer() *AddGradeByStampSheetResult {
	return &p
}

type ApplyRankCapByStampSheetResult struct {
	Item                    *Status              `json:"item"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type ApplyRankCapByStampSheetAsyncResult struct {
	result *ApplyRankCapByStampSheetResult
	err    error
}

func NewApplyRankCapByStampSheetResultFromJson(data string) ApplyRankCapByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewApplyRankCapByStampSheetResultFromDict(dict)
}

func NewApplyRankCapByStampSheetResultFromDict(data map[string]interface{}) ApplyRankCapByStampSheetResult {
	return ApplyRankCapByStampSheetResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p ApplyRankCapByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p ApplyRankCapByStampSheetResult) Pointer() *ApplyRankCapByStampSheetResult {
	return &p
}

type SubGradeByStampTaskResult struct {
	Item                    *Status              `json:"item"`
	NewContextStack         *string              `json:"newContextStack"`
	ExperienceNamespaceName *string              `json:"experienceNamespaceName"`
	ExperienceStatus        *experience.Status   `json:"experienceStatus"`
	Metadata                *core.ResultMetadata `json:"metadata"`
}

type SubGradeByStampTaskAsyncResult struct {
	result *SubGradeByStampTaskResult
	err    error
}

func NewSubGradeByStampTaskResultFromJson(data string) SubGradeByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubGradeByStampTaskResultFromDict(dict)
}

func NewSubGradeByStampTaskResultFromDict(data map[string]interface{}) SubGradeByStampTaskResult {
	return SubGradeByStampTaskResult{
		Item: func() *Status {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewStatusFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
		ExperienceNamespaceName: func() *string {
			v, ok := data["experienceNamespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["experienceNamespaceName"])
		}(),
		ExperienceStatus: func() *experience.Status {
			v, ok := data["experienceStatus"]
			if !ok || v == nil {
				return nil
			}
			return experience.NewStatusFromDict(core.CastMap(data["experienceStatus"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p SubGradeByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack":         p.NewContextStack,
		"experienceNamespaceName": p.ExperienceNamespaceName,
		"experienceStatus": func() map[string]interface{} {
			if p.ExperienceStatus == nil {
				return nil
			}
			return p.ExperienceStatus.ToDict()
		}(),
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p SubGradeByStampTaskResult) Pointer() *SubGradeByStampTaskResult {
	return &p
}

type MultiplyAcquireActionsByUserIdResult struct {
	Items                     []AcquireAction         `json:"items"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type MultiplyAcquireActionsByUserIdAsyncResult struct {
	result *MultiplyAcquireActionsByUserIdResult
	err    error
}

func NewMultiplyAcquireActionsByUserIdResultFromJson(data string) MultiplyAcquireActionsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByUserIdResultFromDict(dict)
}

func NewMultiplyAcquireActionsByUserIdResultFromDict(data map[string]interface{}) MultiplyAcquireActionsByUserIdResult {
	return MultiplyAcquireActionsByUserIdResult{
		Items: func() []AcquireAction {
			if data["items"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["items"]))
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

func (p MultiplyAcquireActionsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
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

func (p MultiplyAcquireActionsByUserIdResult) Pointer() *MultiplyAcquireActionsByUserIdResult {
	return &p
}

type MultiplyAcquireActionsByStampSheetResult struct {
	Items                     []AcquireAction         `json:"items"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type MultiplyAcquireActionsByStampSheetAsyncResult struct {
	result *MultiplyAcquireActionsByStampSheetResult
	err    error
}

func NewMultiplyAcquireActionsByStampSheetResultFromJson(data string) MultiplyAcquireActionsByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMultiplyAcquireActionsByStampSheetResultFromDict(dict)
}

func NewMultiplyAcquireActionsByStampSheetResultFromDict(data map[string]interface{}) MultiplyAcquireActionsByStampSheetResult {
	return MultiplyAcquireActionsByStampSheetResult{
		Items: func() []AcquireAction {
			if data["items"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["items"]))
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

func (p MultiplyAcquireActionsByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAcquireActionsFromDict(
			p.Items,
		),
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

func (p MultiplyAcquireActionsByStampSheetResult) Pointer() *MultiplyAcquireActionsByStampSheetResult {
	return &p
}

type VerifyGradeByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeByStampTaskAsyncResult struct {
	result *VerifyGradeByStampTaskResult
	err    error
}

func NewVerifyGradeByStampTaskResultFromJson(data string) VerifyGradeByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeByStampTaskResultFromDict(dict)
}

func NewVerifyGradeByStampTaskResultFromDict(data map[string]interface{}) VerifyGradeByStampTaskResult {
	return VerifyGradeByStampTaskResult{
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

func (p VerifyGradeByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeByStampTaskResult) Pointer() *VerifyGradeByStampTaskResult {
	return &p
}

type VerifyGradeUpMaterialByStampTaskResult struct {
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyGradeUpMaterialByStampTaskAsyncResult struct {
	result *VerifyGradeUpMaterialByStampTaskResult
	err    error
}

func NewVerifyGradeUpMaterialByStampTaskResultFromJson(data string) VerifyGradeUpMaterialByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGradeUpMaterialByStampTaskResultFromDict(dict)
}

func NewVerifyGradeUpMaterialByStampTaskResultFromDict(data map[string]interface{}) VerifyGradeUpMaterialByStampTaskResult {
	return VerifyGradeUpMaterialByStampTaskResult{
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

func (p VerifyGradeUpMaterialByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"newContextStack": p.NewContextStack,
		"metadata": func() map[string]interface{} {
			if p.Metadata == nil {
				return nil
			}
			return p.Metadata.ToDict()
		},
	}
}

func (p VerifyGradeUpMaterialByStampTaskResult) Pointer() *VerifyGradeUpMaterialByStampTaskResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentGradeMaster  `json:"item"`
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
		Item: func() *CurrentGradeMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
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

type GetCurrentGradeMasterResult struct {
	Item     *CurrentGradeMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetCurrentGradeMasterAsyncResult struct {
	result *GetCurrentGradeMasterResult
	err    error
}

func NewGetCurrentGradeMasterResultFromJson(data string) GetCurrentGradeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentGradeMasterResultFromDict(dict)
}

func NewGetCurrentGradeMasterResultFromDict(data map[string]interface{}) GetCurrentGradeMasterResult {
	return GetCurrentGradeMasterResult{
		Item: func() *CurrentGradeMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p GetCurrentGradeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p GetCurrentGradeMasterResult) Pointer() *GetCurrentGradeMasterResult {
	return &p
}

type UpdateCurrentGradeMasterResult struct {
	Item     *CurrentGradeMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentGradeMasterAsyncResult struct {
	result *UpdateCurrentGradeMasterResult
	err    error
}

func NewUpdateCurrentGradeMasterResultFromJson(data string) UpdateCurrentGradeMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterResultFromDict(dict)
}

func NewUpdateCurrentGradeMasterResultFromDict(data map[string]interface{}) UpdateCurrentGradeMasterResult {
	return UpdateCurrentGradeMasterResult{
		Item: func() *CurrentGradeMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentGradeMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p UpdateCurrentGradeMasterResult) Pointer() *UpdateCurrentGradeMasterResult {
	return &p
}

type UpdateCurrentGradeMasterFromGitHubResult struct {
	Item     *CurrentGradeMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateCurrentGradeMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentGradeMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentGradeMasterFromGitHubResultFromJson(data string) UpdateCurrentGradeMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGradeMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentGradeMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentGradeMasterFromGitHubResult {
	return UpdateCurrentGradeMasterFromGitHubResult{
		Item: func() *CurrentGradeMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentGradeMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		Metadata: func() *core.ResultMetadata {
			if data["metadata"] == nil {
				return nil
			}
			v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
			return &v
		}(),
	}
}

func (p UpdateCurrentGradeMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
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

func (p UpdateCurrentGradeMasterFromGitHubResult) Pointer() *UpdateCurrentGradeMasterFromGitHubResult {
	return &p
}
