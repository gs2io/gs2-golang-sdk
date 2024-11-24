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

package ranking2

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
	}
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastNamespacesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
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
	}
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"status": p.Status,
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
	return DumpUserDataByUserIdResult{}
}

func (p DumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
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
	}
}

func (p CheckDumpUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
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
	return CleanUserDataByUserIdResult{}
}

func (p CleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
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
	return CheckCleanUserDataByUserIdResult{}
}

func (p CheckCleanUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
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
	}
}

func (p PrepareImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"uploadToken": p.UploadToken,
		"uploadUrl":   p.UploadUrl,
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
	return ImportUserDataByUserIdResult{}
}

func (p ImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
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
	}
}

func (p CheckImportUserDataByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"url": p.Url,
	}
}

func (p CheckImportUserDataByUserIdResult) Pointer() *CheckImportUserDataByUserIdResult {
	return &p
}

type DescribeGlobalRankingModelsResult struct {
	Items    []GlobalRankingModel `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeGlobalRankingModelsAsyncResult struct {
	result *DescribeGlobalRankingModelsResult
	err    error
}

func NewDescribeGlobalRankingModelsResultFromJson(data string) DescribeGlobalRankingModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingModelsResultFromDict(dict)
}

func NewDescribeGlobalRankingModelsResultFromDict(data map[string]interface{}) DescribeGlobalRankingModelsResult {
	return DescribeGlobalRankingModelsResult{
		Items: func() []GlobalRankingModel {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeGlobalRankingModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeGlobalRankingModelsResult) Pointer() *DescribeGlobalRankingModelsResult {
	return &p
}

type GetGlobalRankingModelResult struct {
	Item     *GlobalRankingModel  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGlobalRankingModelAsyncResult struct {
	result *GetGlobalRankingModelResult
	err    error
}

func NewGetGlobalRankingModelResultFromJson(data string) GetGlobalRankingModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingModelResultFromDict(dict)
}

func NewGetGlobalRankingModelResultFromDict(data map[string]interface{}) GetGlobalRankingModelResult {
	return GetGlobalRankingModelResult{
		Item: func() *GlobalRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingModelResult) Pointer() *GetGlobalRankingModelResult {
	return &p
}

type DescribeGlobalRankingModelMastersResult struct {
	Items         []GlobalRankingModelMaster `json:"items"`
	NextPageToken *string                    `json:"nextPageToken"`
	Metadata      *core.ResultMetadata       `json:"metadata"`
}

type DescribeGlobalRankingModelMastersAsyncResult struct {
	result *DescribeGlobalRankingModelMastersResult
	err    error
}

func NewDescribeGlobalRankingModelMastersResultFromJson(data string) DescribeGlobalRankingModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingModelMastersResultFromDict(dict)
}

func NewDescribeGlobalRankingModelMastersResultFromDict(data map[string]interface{}) DescribeGlobalRankingModelMastersResult {
	return DescribeGlobalRankingModelMastersResult{
		Items: func() []GlobalRankingModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingModelMastersResult) Pointer() *DescribeGlobalRankingModelMastersResult {
	return &p
}

type CreateGlobalRankingModelMasterResult struct {
	Item     *GlobalRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type CreateGlobalRankingModelMasterAsyncResult struct {
	result *CreateGlobalRankingModelMasterResult
	err    error
}

func NewCreateGlobalRankingModelMasterResultFromJson(data string) CreateGlobalRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalRankingModelMasterResultFromDict(dict)
}

func NewCreateGlobalRankingModelMasterResultFromDict(data map[string]interface{}) CreateGlobalRankingModelMasterResult {
	return CreateGlobalRankingModelMasterResult{
		Item: func() *GlobalRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateGlobalRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateGlobalRankingModelMasterResult) Pointer() *CreateGlobalRankingModelMasterResult {
	return &p
}

type GetGlobalRankingModelMasterResult struct {
	Item     *GlobalRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type GetGlobalRankingModelMasterAsyncResult struct {
	result *GetGlobalRankingModelMasterResult
	err    error
}

func NewGetGlobalRankingModelMasterResultFromJson(data string) GetGlobalRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingModelMasterResultFromDict(dict)
}

func NewGetGlobalRankingModelMasterResultFromDict(data map[string]interface{}) GetGlobalRankingModelMasterResult {
	return GetGlobalRankingModelMasterResult{
		Item: func() *GlobalRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingModelMasterResult) Pointer() *GetGlobalRankingModelMasterResult {
	return &p
}

type UpdateGlobalRankingModelMasterResult struct {
	Item     *GlobalRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type UpdateGlobalRankingModelMasterAsyncResult struct {
	result *UpdateGlobalRankingModelMasterResult
	err    error
}

func NewUpdateGlobalRankingModelMasterResultFromJson(data string) UpdateGlobalRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGlobalRankingModelMasterResultFromDict(dict)
}

func NewUpdateGlobalRankingModelMasterResultFromDict(data map[string]interface{}) UpdateGlobalRankingModelMasterResult {
	return UpdateGlobalRankingModelMasterResult{
		Item: func() *GlobalRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateGlobalRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateGlobalRankingModelMasterResult) Pointer() *UpdateGlobalRankingModelMasterResult {
	return &p
}

type DeleteGlobalRankingModelMasterResult struct {
	Item     *GlobalRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata      `json:"metadata"`
}

type DeleteGlobalRankingModelMasterAsyncResult struct {
	result *DeleteGlobalRankingModelMasterResult
	err    error
}

func NewDeleteGlobalRankingModelMasterResultFromJson(data string) DeleteGlobalRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGlobalRankingModelMasterResultFromDict(dict)
}

func NewDeleteGlobalRankingModelMasterResultFromDict(data map[string]interface{}) DeleteGlobalRankingModelMasterResult {
	return DeleteGlobalRankingModelMasterResult{
		Item: func() *GlobalRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteGlobalRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteGlobalRankingModelMasterResult) Pointer() *DeleteGlobalRankingModelMasterResult {
	return &p
}

type DescribeGlobalRankingScoresResult struct {
	Items         []GlobalRankingScore `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGlobalRankingScoresAsyncResult struct {
	result *DescribeGlobalRankingScoresResult
	err    error
}

func NewDescribeGlobalRankingScoresResultFromJson(data string) DescribeGlobalRankingScoresResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingScoresResultFromDict(dict)
}

func NewDescribeGlobalRankingScoresResultFromDict(data map[string]interface{}) DescribeGlobalRankingScoresResult {
	return DescribeGlobalRankingScoresResult{
		Items: func() []GlobalRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingScoresResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingScoresResult) Pointer() *DescribeGlobalRankingScoresResult {
	return &p
}

type DescribeGlobalRankingScoresByUserIdResult struct {
	Items         []GlobalRankingScore `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGlobalRankingScoresByUserIdAsyncResult struct {
	result *DescribeGlobalRankingScoresByUserIdResult
	err    error
}

func NewDescribeGlobalRankingScoresByUserIdResultFromJson(data string) DescribeGlobalRankingScoresByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingScoresByUserIdResultFromDict(dict)
}

func NewDescribeGlobalRankingScoresByUserIdResultFromDict(data map[string]interface{}) DescribeGlobalRankingScoresByUserIdResult {
	return DescribeGlobalRankingScoresByUserIdResult{
		Items: func() []GlobalRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingScoresByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingScoresByUserIdResult) Pointer() *DescribeGlobalRankingScoresByUserIdResult {
	return &p
}

type PutGlobalRankingScoreResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutGlobalRankingScoreAsyncResult struct {
	result *PutGlobalRankingScoreResult
	err    error
}

func NewPutGlobalRankingScoreResultFromJson(data string) PutGlobalRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutGlobalRankingScoreResultFromDict(dict)
}

func NewPutGlobalRankingScoreResultFromDict(data map[string]interface{}) PutGlobalRankingScoreResult {
	return PutGlobalRankingScoreResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutGlobalRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutGlobalRankingScoreResult) Pointer() *PutGlobalRankingScoreResult {
	return &p
}

type PutGlobalRankingScoreByUserIdResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutGlobalRankingScoreByUserIdAsyncResult struct {
	result *PutGlobalRankingScoreByUserIdResult
	err    error
}

func NewPutGlobalRankingScoreByUserIdResultFromJson(data string) PutGlobalRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutGlobalRankingScoreByUserIdResultFromDict(dict)
}

func NewPutGlobalRankingScoreByUserIdResultFromDict(data map[string]interface{}) PutGlobalRankingScoreByUserIdResult {
	return PutGlobalRankingScoreByUserIdResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutGlobalRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutGlobalRankingScoreByUserIdResult) Pointer() *PutGlobalRankingScoreByUserIdResult {
	return &p
}

type GetGlobalRankingScoreResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGlobalRankingScoreAsyncResult struct {
	result *GetGlobalRankingScoreResult
	err    error
}

func NewGetGlobalRankingScoreResultFromJson(data string) GetGlobalRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingScoreResultFromDict(dict)
}

func NewGetGlobalRankingScoreResultFromDict(data map[string]interface{}) GetGlobalRankingScoreResult {
	return GetGlobalRankingScoreResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingScoreResult) Pointer() *GetGlobalRankingScoreResult {
	return &p
}

type GetGlobalRankingScoreByUserIdResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGlobalRankingScoreByUserIdAsyncResult struct {
	result *GetGlobalRankingScoreByUserIdResult
	err    error
}

func NewGetGlobalRankingScoreByUserIdResultFromJson(data string) GetGlobalRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingScoreByUserIdResultFromDict(dict)
}

func NewGetGlobalRankingScoreByUserIdResultFromDict(data map[string]interface{}) GetGlobalRankingScoreByUserIdResult {
	return GetGlobalRankingScoreByUserIdResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingScoreByUserIdResult) Pointer() *GetGlobalRankingScoreByUserIdResult {
	return &p
}

type DeleteGlobalRankingScoreByUserIdResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteGlobalRankingScoreByUserIdAsyncResult struct {
	result *DeleteGlobalRankingScoreByUserIdResult
	err    error
}

func NewDeleteGlobalRankingScoreByUserIdResultFromJson(data string) DeleteGlobalRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGlobalRankingScoreByUserIdResultFromDict(dict)
}

func NewDeleteGlobalRankingScoreByUserIdResultFromDict(data map[string]interface{}) DeleteGlobalRankingScoreByUserIdResult {
	return DeleteGlobalRankingScoreByUserIdResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteGlobalRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteGlobalRankingScoreByUserIdResult) Pointer() *DeleteGlobalRankingScoreByUserIdResult {
	return &p
}

type VerifyGlobalRankingScoreResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGlobalRankingScoreAsyncResult struct {
	result *VerifyGlobalRankingScoreResult
	err    error
}

func NewVerifyGlobalRankingScoreResultFromJson(data string) VerifyGlobalRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGlobalRankingScoreResultFromDict(dict)
}

func NewVerifyGlobalRankingScoreResultFromDict(data map[string]interface{}) VerifyGlobalRankingScoreResult {
	return VerifyGlobalRankingScoreResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifyGlobalRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyGlobalRankingScoreResult) Pointer() *VerifyGlobalRankingScoreResult {
	return &p
}

type VerifyGlobalRankingScoreByUserIdResult struct {
	Item     *GlobalRankingScore  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyGlobalRankingScoreByUserIdAsyncResult struct {
	result *VerifyGlobalRankingScoreByUserIdResult
	err    error
}

func NewVerifyGlobalRankingScoreByUserIdResultFromJson(data string) VerifyGlobalRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGlobalRankingScoreByUserIdResultFromDict(dict)
}

func NewVerifyGlobalRankingScoreByUserIdResultFromDict(data map[string]interface{}) VerifyGlobalRankingScoreByUserIdResult {
	return VerifyGlobalRankingScoreByUserIdResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifyGlobalRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyGlobalRankingScoreByUserIdResult) Pointer() *VerifyGlobalRankingScoreByUserIdResult {
	return &p
}

type VerifyGlobalRankingScoreByStampTaskResult struct {
	Item            *GlobalRankingScore  `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyGlobalRankingScoreByStampTaskAsyncResult struct {
	result *VerifyGlobalRankingScoreByStampTaskResult
	err    error
}

func NewVerifyGlobalRankingScoreByStampTaskResultFromJson(data string) VerifyGlobalRankingScoreByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyGlobalRankingScoreByStampTaskResultFromDict(dict)
}

func NewVerifyGlobalRankingScoreByStampTaskResultFromDict(data map[string]interface{}) VerifyGlobalRankingScoreByStampTaskResult {
	return VerifyGlobalRankingScoreByStampTaskResult{
		Item: func() *GlobalRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p VerifyGlobalRankingScoreByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyGlobalRankingScoreByStampTaskResult) Pointer() *VerifyGlobalRankingScoreByStampTaskResult {
	return &p
}

type DescribeGlobalRankingReceivedRewardsResult struct {
	Items         []GlobalRankingReceivedReward `json:"items"`
	NextPageToken *string                       `json:"nextPageToken"`
	Metadata      *core.ResultMetadata          `json:"metadata"`
}

type DescribeGlobalRankingReceivedRewardsAsyncResult struct {
	result *DescribeGlobalRankingReceivedRewardsResult
	err    error
}

func NewDescribeGlobalRankingReceivedRewardsResultFromJson(data string) DescribeGlobalRankingReceivedRewardsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingReceivedRewardsResultFromDict(dict)
}

func NewDescribeGlobalRankingReceivedRewardsResultFromDict(data map[string]interface{}) DescribeGlobalRankingReceivedRewardsResult {
	return DescribeGlobalRankingReceivedRewardsResult{
		Items: func() []GlobalRankingReceivedReward {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingReceivedRewards(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingReceivedRewardsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingReceivedRewardsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingReceivedRewardsResult) Pointer() *DescribeGlobalRankingReceivedRewardsResult {
	return &p
}

type DescribeGlobalRankingReceivedRewardsByUserIdResult struct {
	Items         []GlobalRankingReceivedReward `json:"items"`
	NextPageToken *string                       `json:"nextPageToken"`
	Metadata      *core.ResultMetadata          `json:"metadata"`
}

type DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult struct {
	result *DescribeGlobalRankingReceivedRewardsByUserIdResult
	err    error
}

func NewDescribeGlobalRankingReceivedRewardsByUserIdResultFromJson(data string) DescribeGlobalRankingReceivedRewardsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingReceivedRewardsByUserIdResultFromDict(dict)
}

func NewDescribeGlobalRankingReceivedRewardsByUserIdResultFromDict(data map[string]interface{}) DescribeGlobalRankingReceivedRewardsByUserIdResult {
	return DescribeGlobalRankingReceivedRewardsByUserIdResult{
		Items: func() []GlobalRankingReceivedReward {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingReceivedRewards(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingReceivedRewardsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingReceivedRewardsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingReceivedRewardsByUserIdResult) Pointer() *DescribeGlobalRankingReceivedRewardsByUserIdResult {
	return &p
}

type CreateGlobalRankingReceivedRewardResult struct {
	Item     *GlobalRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type CreateGlobalRankingReceivedRewardAsyncResult struct {
	result *CreateGlobalRankingReceivedRewardResult
	err    error
}

func NewCreateGlobalRankingReceivedRewardResultFromJson(data string) CreateGlobalRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalRankingReceivedRewardResultFromDict(dict)
}

func NewCreateGlobalRankingReceivedRewardResultFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardResult {
	return CreateGlobalRankingReceivedRewardResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardResult) Pointer() *CreateGlobalRankingReceivedRewardResult {
	return &p
}

type CreateGlobalRankingReceivedRewardByUserIdResult struct {
	Item     *GlobalRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type CreateGlobalRankingReceivedRewardByUserIdAsyncResult struct {
	result *CreateGlobalRankingReceivedRewardByUserIdResult
	err    error
}

func NewCreateGlobalRankingReceivedRewardByUserIdResultFromJson(data string) CreateGlobalRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewCreateGlobalRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardByUserIdResult {
	return CreateGlobalRankingReceivedRewardByUserIdResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardByUserIdResult) Pointer() *CreateGlobalRankingReceivedRewardByUserIdResult {
	return &p
}

type ReceiveGlobalRankingReceivedRewardResult struct {
	Item                      *GlobalRankingModel     `json:"item"`
	AcquireActions            []AcquireAction         `json:"acquireActions"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ReceiveGlobalRankingReceivedRewardAsyncResult struct {
	result *ReceiveGlobalRankingReceivedRewardResult
	err    error
}

func NewReceiveGlobalRankingReceivedRewardResultFromJson(data string) ReceiveGlobalRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalRankingReceivedRewardResultFromDict(dict)
}

func NewReceiveGlobalRankingReceivedRewardResultFromDict(data map[string]interface{}) ReceiveGlobalRankingReceivedRewardResult {
	return ReceiveGlobalRankingReceivedRewardResult{
		Item: func() *GlobalRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
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
	}
}

func (p ReceiveGlobalRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
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
	}
}

func (p ReceiveGlobalRankingReceivedRewardResult) Pointer() *ReceiveGlobalRankingReceivedRewardResult {
	return &p
}

type ReceiveGlobalRankingReceivedRewardByUserIdResult struct {
	Item                      *GlobalRankingModel     `json:"item"`
	AcquireActions            []AcquireAction         `json:"acquireActions"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult struct {
	result *ReceiveGlobalRankingReceivedRewardByUserIdResult
	err    error
}

func NewReceiveGlobalRankingReceivedRewardByUserIdResultFromJson(data string) ReceiveGlobalRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveGlobalRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewReceiveGlobalRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) ReceiveGlobalRankingReceivedRewardByUserIdResult {
	return ReceiveGlobalRankingReceivedRewardByUserIdResult{
		Item: func() *GlobalRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
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
	}
}

func (p ReceiveGlobalRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
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
	}
}

func (p ReceiveGlobalRankingReceivedRewardByUserIdResult) Pointer() *ReceiveGlobalRankingReceivedRewardByUserIdResult {
	return &p
}

type GetGlobalRankingReceivedRewardResult struct {
	Item     *GlobalRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type GetGlobalRankingReceivedRewardAsyncResult struct {
	result *GetGlobalRankingReceivedRewardResult
	err    error
}

func NewGetGlobalRankingReceivedRewardResultFromJson(data string) GetGlobalRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingReceivedRewardResultFromDict(dict)
}

func NewGetGlobalRankingReceivedRewardResultFromDict(data map[string]interface{}) GetGlobalRankingReceivedRewardResult {
	return GetGlobalRankingReceivedRewardResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingReceivedRewardResult) Pointer() *GetGlobalRankingReceivedRewardResult {
	return &p
}

type GetGlobalRankingReceivedRewardByUserIdResult struct {
	Item     *GlobalRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type GetGlobalRankingReceivedRewardByUserIdAsyncResult struct {
	result *GetGlobalRankingReceivedRewardByUserIdResult
	err    error
}

func NewGetGlobalRankingReceivedRewardByUserIdResultFromJson(data string) GetGlobalRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewGetGlobalRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) GetGlobalRankingReceivedRewardByUserIdResult {
	return GetGlobalRankingReceivedRewardByUserIdResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingReceivedRewardByUserIdResult) Pointer() *GetGlobalRankingReceivedRewardByUserIdResult {
	return &p
}

type DeleteGlobalRankingReceivedRewardByUserIdResult struct {
	Item     *GlobalRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type DeleteGlobalRankingReceivedRewardByUserIdAsyncResult struct {
	result *DeleteGlobalRankingReceivedRewardByUserIdResult
	err    error
}

func NewDeleteGlobalRankingReceivedRewardByUserIdResultFromJson(data string) DeleteGlobalRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGlobalRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewDeleteGlobalRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) DeleteGlobalRankingReceivedRewardByUserIdResult {
	return DeleteGlobalRankingReceivedRewardByUserIdResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteGlobalRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteGlobalRankingReceivedRewardByUserIdResult) Pointer() *DeleteGlobalRankingReceivedRewardByUserIdResult {
	return &p
}

type CreateGlobalRankingReceivedRewardByStampTaskResult struct {
	Item            *GlobalRankingReceivedReward `json:"item"`
	NewContextStack *string                      `json:"newContextStack"`
	Metadata        *core.ResultMetadata         `json:"metadata"`
}

type CreateGlobalRankingReceivedRewardByStampTaskAsyncResult struct {
	result *CreateGlobalRankingReceivedRewardByStampTaskResult
	err    error
}

func NewCreateGlobalRankingReceivedRewardByStampTaskResultFromJson(data string) CreateGlobalRankingReceivedRewardByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGlobalRankingReceivedRewardByStampTaskResultFromDict(dict)
}

func NewCreateGlobalRankingReceivedRewardByStampTaskResultFromDict(data map[string]interface{}) CreateGlobalRankingReceivedRewardByStampTaskResult {
	return CreateGlobalRankingReceivedRewardByStampTaskResult{
		Item: func() *GlobalRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p CreateGlobalRankingReceivedRewardByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p CreateGlobalRankingReceivedRewardByStampTaskResult) Pointer() *CreateGlobalRankingReceivedRewardByStampTaskResult {
	return &p
}

type DescribeGlobalRankingsResult struct {
	Items         []GlobalRankingData  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGlobalRankingsAsyncResult struct {
	result *DescribeGlobalRankingsResult
	err    error
}

func NewDescribeGlobalRankingsResultFromJson(data string) DescribeGlobalRankingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingsResultFromDict(dict)
}

func NewDescribeGlobalRankingsResultFromDict(data map[string]interface{}) DescribeGlobalRankingsResult {
	return DescribeGlobalRankingsResult{
		Items: func() []GlobalRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingsResult) Pointer() *DescribeGlobalRankingsResult {
	return &p
}

type DescribeGlobalRankingsByUserIdResult struct {
	Items         []GlobalRankingData  `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeGlobalRankingsByUserIdAsyncResult struct {
	result *DescribeGlobalRankingsByUserIdResult
	err    error
}

func NewDescribeGlobalRankingsByUserIdResultFromJson(data string) DescribeGlobalRankingsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGlobalRankingsByUserIdResultFromDict(dict)
}

func NewDescribeGlobalRankingsByUserIdResultFromDict(data map[string]interface{}) DescribeGlobalRankingsByUserIdResult {
	return DescribeGlobalRankingsByUserIdResult{
		Items: func() []GlobalRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastGlobalRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeGlobalRankingsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGlobalRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGlobalRankingsByUserIdResult) Pointer() *DescribeGlobalRankingsByUserIdResult {
	return &p
}

type GetGlobalRankingResult struct {
	Item     *GlobalRankingData   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGlobalRankingAsyncResult struct {
	result *GetGlobalRankingResult
	err    error
}

func NewGetGlobalRankingResultFromJson(data string) GetGlobalRankingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingResultFromDict(dict)
}

func NewGetGlobalRankingResultFromDict(data map[string]interface{}) GetGlobalRankingResult {
	return GetGlobalRankingResult{
		Item: func() *GlobalRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingResult) Pointer() *GetGlobalRankingResult {
	return &p
}

type GetGlobalRankingByUserIdResult struct {
	Item     *GlobalRankingData   `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetGlobalRankingByUserIdAsyncResult struct {
	result *GetGlobalRankingByUserIdResult
	err    error
}

func NewGetGlobalRankingByUserIdResultFromJson(data string) GetGlobalRankingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGlobalRankingByUserIdResultFromDict(dict)
}

func NewGetGlobalRankingByUserIdResultFromDict(data map[string]interface{}) GetGlobalRankingByUserIdResult {
	return GetGlobalRankingByUserIdResult{
		Item: func() *GlobalRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewGlobalRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetGlobalRankingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetGlobalRankingByUserIdResult) Pointer() *GetGlobalRankingByUserIdResult {
	return &p
}

type DescribeClusterRankingModelsResult struct {
	Items    []ClusterRankingModel `json:"items"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type DescribeClusterRankingModelsAsyncResult struct {
	result *DescribeClusterRankingModelsResult
	err    error
}

func NewDescribeClusterRankingModelsResultFromJson(data string) DescribeClusterRankingModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingModelsResultFromDict(dict)
}

func NewDescribeClusterRankingModelsResultFromDict(data map[string]interface{}) DescribeClusterRankingModelsResult {
	return DescribeClusterRankingModelsResult{
		Items: func() []ClusterRankingModel {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeClusterRankingModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeClusterRankingModelsResult) Pointer() *DescribeClusterRankingModelsResult {
	return &p
}

type GetClusterRankingModelResult struct {
	Item     *ClusterRankingModel `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetClusterRankingModelAsyncResult struct {
	result *GetClusterRankingModelResult
	err    error
}

func NewGetClusterRankingModelResultFromJson(data string) GetClusterRankingModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingModelResultFromDict(dict)
}

func NewGetClusterRankingModelResultFromDict(data map[string]interface{}) GetClusterRankingModelResult {
	return GetClusterRankingModelResult{
		Item: func() *ClusterRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingModelResult) Pointer() *GetClusterRankingModelResult {
	return &p
}

type DescribeClusterRankingModelMastersResult struct {
	Items         []ClusterRankingModelMaster `json:"items"`
	NextPageToken *string                     `json:"nextPageToken"`
	Metadata      *core.ResultMetadata        `json:"metadata"`
}

type DescribeClusterRankingModelMastersAsyncResult struct {
	result *DescribeClusterRankingModelMastersResult
	err    error
}

func NewDescribeClusterRankingModelMastersResultFromJson(data string) DescribeClusterRankingModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingModelMastersResultFromDict(dict)
}

func NewDescribeClusterRankingModelMastersResultFromDict(data map[string]interface{}) DescribeClusterRankingModelMastersResult {
	return DescribeClusterRankingModelMastersResult{
		Items: func() []ClusterRankingModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingModelMastersResult) Pointer() *DescribeClusterRankingModelMastersResult {
	return &p
}

type CreateClusterRankingModelMasterResult struct {
	Item     *ClusterRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata       `json:"metadata"`
}

type CreateClusterRankingModelMasterAsyncResult struct {
	result *CreateClusterRankingModelMasterResult
	err    error
}

func NewCreateClusterRankingModelMasterResultFromJson(data string) CreateClusterRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateClusterRankingModelMasterResultFromDict(dict)
}

func NewCreateClusterRankingModelMasterResultFromDict(data map[string]interface{}) CreateClusterRankingModelMasterResult {
	return CreateClusterRankingModelMasterResult{
		Item: func() *ClusterRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateClusterRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateClusterRankingModelMasterResult) Pointer() *CreateClusterRankingModelMasterResult {
	return &p
}

type GetClusterRankingModelMasterResult struct {
	Item     *ClusterRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata       `json:"metadata"`
}

type GetClusterRankingModelMasterAsyncResult struct {
	result *GetClusterRankingModelMasterResult
	err    error
}

func NewGetClusterRankingModelMasterResultFromJson(data string) GetClusterRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingModelMasterResultFromDict(dict)
}

func NewGetClusterRankingModelMasterResultFromDict(data map[string]interface{}) GetClusterRankingModelMasterResult {
	return GetClusterRankingModelMasterResult{
		Item: func() *ClusterRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingModelMasterResult) Pointer() *GetClusterRankingModelMasterResult {
	return &p
}

type UpdateClusterRankingModelMasterResult struct {
	Item     *ClusterRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata       `json:"metadata"`
}

type UpdateClusterRankingModelMasterAsyncResult struct {
	result *UpdateClusterRankingModelMasterResult
	err    error
}

func NewUpdateClusterRankingModelMasterResultFromJson(data string) UpdateClusterRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateClusterRankingModelMasterResultFromDict(dict)
}

func NewUpdateClusterRankingModelMasterResultFromDict(data map[string]interface{}) UpdateClusterRankingModelMasterResult {
	return UpdateClusterRankingModelMasterResult{
		Item: func() *ClusterRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateClusterRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateClusterRankingModelMasterResult) Pointer() *UpdateClusterRankingModelMasterResult {
	return &p
}

type DeleteClusterRankingModelMasterResult struct {
	Item     *ClusterRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata       `json:"metadata"`
}

type DeleteClusterRankingModelMasterAsyncResult struct {
	result *DeleteClusterRankingModelMasterResult
	err    error
}

func NewDeleteClusterRankingModelMasterResultFromJson(data string) DeleteClusterRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteClusterRankingModelMasterResultFromDict(dict)
}

func NewDeleteClusterRankingModelMasterResultFromDict(data map[string]interface{}) DeleteClusterRankingModelMasterResult {
	return DeleteClusterRankingModelMasterResult{
		Item: func() *ClusterRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteClusterRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteClusterRankingModelMasterResult) Pointer() *DeleteClusterRankingModelMasterResult {
	return &p
}

type DescribeClusterRankingScoresResult struct {
	Items         []ClusterRankingScore `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
	Metadata      *core.ResultMetadata  `json:"metadata"`
}

type DescribeClusterRankingScoresAsyncResult struct {
	result *DescribeClusterRankingScoresResult
	err    error
}

func NewDescribeClusterRankingScoresResultFromJson(data string) DescribeClusterRankingScoresResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingScoresResultFromDict(dict)
}

func NewDescribeClusterRankingScoresResultFromDict(data map[string]interface{}) DescribeClusterRankingScoresResult {
	return DescribeClusterRankingScoresResult{
		Items: func() []ClusterRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingScoresResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingScoresResult) Pointer() *DescribeClusterRankingScoresResult {
	return &p
}

type DescribeClusterRankingScoresByUserIdResult struct {
	Items         []ClusterRankingScore `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
	Metadata      *core.ResultMetadata  `json:"metadata"`
}

type DescribeClusterRankingScoresByUserIdAsyncResult struct {
	result *DescribeClusterRankingScoresByUserIdResult
	err    error
}

func NewDescribeClusterRankingScoresByUserIdResultFromJson(data string) DescribeClusterRankingScoresByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingScoresByUserIdResultFromDict(dict)
}

func NewDescribeClusterRankingScoresByUserIdResultFromDict(data map[string]interface{}) DescribeClusterRankingScoresByUserIdResult {
	return DescribeClusterRankingScoresByUserIdResult{
		Items: func() []ClusterRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingScoresByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingScoresByUserIdResult) Pointer() *DescribeClusterRankingScoresByUserIdResult {
	return &p
}

type PutClusterRankingScoreResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutClusterRankingScoreAsyncResult struct {
	result *PutClusterRankingScoreResult
	err    error
}

func NewPutClusterRankingScoreResultFromJson(data string) PutClusterRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutClusterRankingScoreResultFromDict(dict)
}

func NewPutClusterRankingScoreResultFromDict(data map[string]interface{}) PutClusterRankingScoreResult {
	return PutClusterRankingScoreResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutClusterRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutClusterRankingScoreResult) Pointer() *PutClusterRankingScoreResult {
	return &p
}

type PutClusterRankingScoreByUserIdResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PutClusterRankingScoreByUserIdAsyncResult struct {
	result *PutClusterRankingScoreByUserIdResult
	err    error
}

func NewPutClusterRankingScoreByUserIdResultFromJson(data string) PutClusterRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutClusterRankingScoreByUserIdResultFromDict(dict)
}

func NewPutClusterRankingScoreByUserIdResultFromDict(data map[string]interface{}) PutClusterRankingScoreByUserIdResult {
	return PutClusterRankingScoreByUserIdResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutClusterRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutClusterRankingScoreByUserIdResult) Pointer() *PutClusterRankingScoreByUserIdResult {
	return &p
}

type GetClusterRankingScoreResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetClusterRankingScoreAsyncResult struct {
	result *GetClusterRankingScoreResult
	err    error
}

func NewGetClusterRankingScoreResultFromJson(data string) GetClusterRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingScoreResultFromDict(dict)
}

func NewGetClusterRankingScoreResultFromDict(data map[string]interface{}) GetClusterRankingScoreResult {
	return GetClusterRankingScoreResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingScoreResult) Pointer() *GetClusterRankingScoreResult {
	return &p
}

type GetClusterRankingScoreByUserIdResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetClusterRankingScoreByUserIdAsyncResult struct {
	result *GetClusterRankingScoreByUserIdResult
	err    error
}

func NewGetClusterRankingScoreByUserIdResultFromJson(data string) GetClusterRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingScoreByUserIdResultFromDict(dict)
}

func NewGetClusterRankingScoreByUserIdResultFromDict(data map[string]interface{}) GetClusterRankingScoreByUserIdResult {
	return GetClusterRankingScoreByUserIdResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingScoreByUserIdResult) Pointer() *GetClusterRankingScoreByUserIdResult {
	return &p
}

type DeleteClusterRankingScoreByUserIdResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteClusterRankingScoreByUserIdAsyncResult struct {
	result *DeleteClusterRankingScoreByUserIdResult
	err    error
}

func NewDeleteClusterRankingScoreByUserIdResultFromJson(data string) DeleteClusterRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteClusterRankingScoreByUserIdResultFromDict(dict)
}

func NewDeleteClusterRankingScoreByUserIdResultFromDict(data map[string]interface{}) DeleteClusterRankingScoreByUserIdResult {
	return DeleteClusterRankingScoreByUserIdResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteClusterRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteClusterRankingScoreByUserIdResult) Pointer() *DeleteClusterRankingScoreByUserIdResult {
	return &p
}

type VerifyClusterRankingScoreResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyClusterRankingScoreAsyncResult struct {
	result *VerifyClusterRankingScoreResult
	err    error
}

func NewVerifyClusterRankingScoreResultFromJson(data string) VerifyClusterRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyClusterRankingScoreResultFromDict(dict)
}

func NewVerifyClusterRankingScoreResultFromDict(data map[string]interface{}) VerifyClusterRankingScoreResult {
	return VerifyClusterRankingScoreResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifyClusterRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyClusterRankingScoreResult) Pointer() *VerifyClusterRankingScoreResult {
	return &p
}

type VerifyClusterRankingScoreByUserIdResult struct {
	Item     *ClusterRankingScore `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type VerifyClusterRankingScoreByUserIdAsyncResult struct {
	result *VerifyClusterRankingScoreByUserIdResult
	err    error
}

func NewVerifyClusterRankingScoreByUserIdResultFromJson(data string) VerifyClusterRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyClusterRankingScoreByUserIdResultFromDict(dict)
}

func NewVerifyClusterRankingScoreByUserIdResultFromDict(data map[string]interface{}) VerifyClusterRankingScoreByUserIdResult {
	return VerifyClusterRankingScoreByUserIdResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifyClusterRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyClusterRankingScoreByUserIdResult) Pointer() *VerifyClusterRankingScoreByUserIdResult {
	return &p
}

type VerifyClusterRankingScoreByStampTaskResult struct {
	Item            *ClusterRankingScore `json:"item"`
	NewContextStack *string              `json:"newContextStack"`
	Metadata        *core.ResultMetadata `json:"metadata"`
}

type VerifyClusterRankingScoreByStampTaskAsyncResult struct {
	result *VerifyClusterRankingScoreByStampTaskResult
	err    error
}

func NewVerifyClusterRankingScoreByStampTaskResultFromJson(data string) VerifyClusterRankingScoreByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyClusterRankingScoreByStampTaskResultFromDict(dict)
}

func NewVerifyClusterRankingScoreByStampTaskResultFromDict(data map[string]interface{}) VerifyClusterRankingScoreByStampTaskResult {
	return VerifyClusterRankingScoreByStampTaskResult{
		Item: func() *ClusterRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p VerifyClusterRankingScoreByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifyClusterRankingScoreByStampTaskResult) Pointer() *VerifyClusterRankingScoreByStampTaskResult {
	return &p
}

type DescribeClusterRankingReceivedRewardsResult struct {
	Items         []ClusterRankingReceivedReward `json:"items"`
	NextPageToken *string                        `json:"nextPageToken"`
	Metadata      *core.ResultMetadata           `json:"metadata"`
}

type DescribeClusterRankingReceivedRewardsAsyncResult struct {
	result *DescribeClusterRankingReceivedRewardsResult
	err    error
}

func NewDescribeClusterRankingReceivedRewardsResultFromJson(data string) DescribeClusterRankingReceivedRewardsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingReceivedRewardsResultFromDict(dict)
}

func NewDescribeClusterRankingReceivedRewardsResultFromDict(data map[string]interface{}) DescribeClusterRankingReceivedRewardsResult {
	return DescribeClusterRankingReceivedRewardsResult{
		Items: func() []ClusterRankingReceivedReward {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingReceivedRewards(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingReceivedRewardsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingReceivedRewardsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingReceivedRewardsResult) Pointer() *DescribeClusterRankingReceivedRewardsResult {
	return &p
}

type DescribeClusterRankingReceivedRewardsByUserIdResult struct {
	Items         []ClusterRankingReceivedReward `json:"items"`
	NextPageToken *string                        `json:"nextPageToken"`
	Metadata      *core.ResultMetadata           `json:"metadata"`
}

type DescribeClusterRankingReceivedRewardsByUserIdAsyncResult struct {
	result *DescribeClusterRankingReceivedRewardsByUserIdResult
	err    error
}

func NewDescribeClusterRankingReceivedRewardsByUserIdResultFromJson(data string) DescribeClusterRankingReceivedRewardsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingReceivedRewardsByUserIdResultFromDict(dict)
}

func NewDescribeClusterRankingReceivedRewardsByUserIdResultFromDict(data map[string]interface{}) DescribeClusterRankingReceivedRewardsByUserIdResult {
	return DescribeClusterRankingReceivedRewardsByUserIdResult{
		Items: func() []ClusterRankingReceivedReward {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingReceivedRewards(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingReceivedRewardsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingReceivedRewardsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingReceivedRewardsByUserIdResult) Pointer() *DescribeClusterRankingReceivedRewardsByUserIdResult {
	return &p
}

type CreateClusterRankingReceivedRewardResult struct {
	Item     *ClusterRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata          `json:"metadata"`
}

type CreateClusterRankingReceivedRewardAsyncResult struct {
	result *CreateClusterRankingReceivedRewardResult
	err    error
}

func NewCreateClusterRankingReceivedRewardResultFromJson(data string) CreateClusterRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateClusterRankingReceivedRewardResultFromDict(dict)
}

func NewCreateClusterRankingReceivedRewardResultFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardResult {
	return CreateClusterRankingReceivedRewardResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardResult) Pointer() *CreateClusterRankingReceivedRewardResult {
	return &p
}

type CreateClusterRankingReceivedRewardByUserIdResult struct {
	Item     *ClusterRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata          `json:"metadata"`
}

type CreateClusterRankingReceivedRewardByUserIdAsyncResult struct {
	result *CreateClusterRankingReceivedRewardByUserIdResult
	err    error
}

func NewCreateClusterRankingReceivedRewardByUserIdResultFromJson(data string) CreateClusterRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateClusterRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewCreateClusterRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardByUserIdResult {
	return CreateClusterRankingReceivedRewardByUserIdResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardByUserIdResult) Pointer() *CreateClusterRankingReceivedRewardByUserIdResult {
	return &p
}

type ReceiveClusterRankingReceivedRewardResult struct {
	Item                      *ClusterRankingModel    `json:"item"`
	AcquireActions            []AcquireAction         `json:"acquireActions"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ReceiveClusterRankingReceivedRewardAsyncResult struct {
	result *ReceiveClusterRankingReceivedRewardResult
	err    error
}

func NewReceiveClusterRankingReceivedRewardResultFromJson(data string) ReceiveClusterRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveClusterRankingReceivedRewardResultFromDict(dict)
}

func NewReceiveClusterRankingReceivedRewardResultFromDict(data map[string]interface{}) ReceiveClusterRankingReceivedRewardResult {
	return ReceiveClusterRankingReceivedRewardResult{
		Item: func() *ClusterRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
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
	}
}

func (p ReceiveClusterRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
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
	}
}

func (p ReceiveClusterRankingReceivedRewardResult) Pointer() *ReceiveClusterRankingReceivedRewardResult {
	return &p
}

type ReceiveClusterRankingReceivedRewardByUserIdResult struct {
	Item                      *ClusterRankingModel    `json:"item"`
	AcquireActions            []AcquireAction         `json:"acquireActions"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type ReceiveClusterRankingReceivedRewardByUserIdAsyncResult struct {
	result *ReceiveClusterRankingReceivedRewardByUserIdResult
	err    error
}

func NewReceiveClusterRankingReceivedRewardByUserIdResultFromJson(data string) ReceiveClusterRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveClusterRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewReceiveClusterRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) ReceiveClusterRankingReceivedRewardByUserIdResult {
	return ReceiveClusterRankingReceivedRewardByUserIdResult{
		Item: func() *ClusterRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		AcquireActions: func() []AcquireAction {
			if data["acquireActions"] == nil {
				return nil
			}
			return CastAcquireActions(core.CastArray(data["acquireActions"]))
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
	}
}

func (p ReceiveClusterRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"acquireActions": CastAcquireActionsFromDict(
			p.AcquireActions,
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
	}
}

func (p ReceiveClusterRankingReceivedRewardByUserIdResult) Pointer() *ReceiveClusterRankingReceivedRewardByUserIdResult {
	return &p
}

type GetClusterRankingReceivedRewardResult struct {
	Item     *ClusterRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata          `json:"metadata"`
}

type GetClusterRankingReceivedRewardAsyncResult struct {
	result *GetClusterRankingReceivedRewardResult
	err    error
}

func NewGetClusterRankingReceivedRewardResultFromJson(data string) GetClusterRankingReceivedRewardResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingReceivedRewardResultFromDict(dict)
}

func NewGetClusterRankingReceivedRewardResultFromDict(data map[string]interface{}) GetClusterRankingReceivedRewardResult {
	return GetClusterRankingReceivedRewardResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingReceivedRewardResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingReceivedRewardResult) Pointer() *GetClusterRankingReceivedRewardResult {
	return &p
}

type GetClusterRankingReceivedRewardByUserIdResult struct {
	Item     *ClusterRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata          `json:"metadata"`
}

type GetClusterRankingReceivedRewardByUserIdAsyncResult struct {
	result *GetClusterRankingReceivedRewardByUserIdResult
	err    error
}

func NewGetClusterRankingReceivedRewardByUserIdResultFromJson(data string) GetClusterRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewGetClusterRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) GetClusterRankingReceivedRewardByUserIdResult {
	return GetClusterRankingReceivedRewardByUserIdResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingReceivedRewardByUserIdResult) Pointer() *GetClusterRankingReceivedRewardByUserIdResult {
	return &p
}

type DeleteClusterRankingReceivedRewardByUserIdResult struct {
	Item     *ClusterRankingReceivedReward `json:"item"`
	Metadata *core.ResultMetadata          `json:"metadata"`
}

type DeleteClusterRankingReceivedRewardByUserIdAsyncResult struct {
	result *DeleteClusterRankingReceivedRewardByUserIdResult
	err    error
}

func NewDeleteClusterRankingReceivedRewardByUserIdResultFromJson(data string) DeleteClusterRankingReceivedRewardByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteClusterRankingReceivedRewardByUserIdResultFromDict(dict)
}

func NewDeleteClusterRankingReceivedRewardByUserIdResultFromDict(data map[string]interface{}) DeleteClusterRankingReceivedRewardByUserIdResult {
	return DeleteClusterRankingReceivedRewardByUserIdResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteClusterRankingReceivedRewardByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteClusterRankingReceivedRewardByUserIdResult) Pointer() *DeleteClusterRankingReceivedRewardByUserIdResult {
	return &p
}

type CreateClusterRankingReceivedRewardByStampTaskResult struct {
	Item            *ClusterRankingReceivedReward `json:"item"`
	NewContextStack *string                       `json:"newContextStack"`
	Metadata        *core.ResultMetadata          `json:"metadata"`
}

type CreateClusterRankingReceivedRewardByStampTaskAsyncResult struct {
	result *CreateClusterRankingReceivedRewardByStampTaskResult
	err    error
}

func NewCreateClusterRankingReceivedRewardByStampTaskResultFromJson(data string) CreateClusterRankingReceivedRewardByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateClusterRankingReceivedRewardByStampTaskResultFromDict(dict)
}

func NewCreateClusterRankingReceivedRewardByStampTaskResultFromDict(data map[string]interface{}) CreateClusterRankingReceivedRewardByStampTaskResult {
	return CreateClusterRankingReceivedRewardByStampTaskResult{
		Item: func() *ClusterRankingReceivedReward {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingReceivedRewardFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p CreateClusterRankingReceivedRewardByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p CreateClusterRankingReceivedRewardByStampTaskResult) Pointer() *CreateClusterRankingReceivedRewardByStampTaskResult {
	return &p
}

type DescribeClusterRankingsResult struct {
	Items         []ClusterRankingData `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeClusterRankingsAsyncResult struct {
	result *DescribeClusterRankingsResult
	err    error
}

func NewDescribeClusterRankingsResultFromJson(data string) DescribeClusterRankingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingsResultFromDict(dict)
}

func NewDescribeClusterRankingsResultFromDict(data map[string]interface{}) DescribeClusterRankingsResult {
	return DescribeClusterRankingsResult{
		Items: func() []ClusterRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingsResult) Pointer() *DescribeClusterRankingsResult {
	return &p
}

type DescribeClusterRankingsByUserIdResult struct {
	Items         []ClusterRankingData `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeClusterRankingsByUserIdAsyncResult struct {
	result *DescribeClusterRankingsByUserIdResult
	err    error
}

func NewDescribeClusterRankingsByUserIdResultFromJson(data string) DescribeClusterRankingsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeClusterRankingsByUserIdResultFromDict(dict)
}

func NewDescribeClusterRankingsByUserIdResultFromDict(data map[string]interface{}) DescribeClusterRankingsByUserIdResult {
	return DescribeClusterRankingsByUserIdResult{
		Items: func() []ClusterRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastClusterRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeClusterRankingsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastClusterRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeClusterRankingsByUserIdResult) Pointer() *DescribeClusterRankingsByUserIdResult {
	return &p
}

type GetClusterRankingResult struct {
	Item     *ClusterRankingData  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetClusterRankingAsyncResult struct {
	result *GetClusterRankingResult
	err    error
}

func NewGetClusterRankingResultFromJson(data string) GetClusterRankingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingResultFromDict(dict)
}

func NewGetClusterRankingResultFromDict(data map[string]interface{}) GetClusterRankingResult {
	return GetClusterRankingResult{
		Item: func() *ClusterRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingResult) Pointer() *GetClusterRankingResult {
	return &p
}

type GetClusterRankingByUserIdResult struct {
	Item     *ClusterRankingData  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetClusterRankingByUserIdAsyncResult struct {
	result *GetClusterRankingByUserIdResult
	err    error
}

func NewGetClusterRankingByUserIdResultFromJson(data string) GetClusterRankingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetClusterRankingByUserIdResultFromDict(dict)
}

func NewGetClusterRankingByUserIdResultFromDict(data map[string]interface{}) GetClusterRankingByUserIdResult {
	return GetClusterRankingByUserIdResult{
		Item: func() *ClusterRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewClusterRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetClusterRankingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetClusterRankingByUserIdResult) Pointer() *GetClusterRankingByUserIdResult {
	return &p
}

type DescribeSubscribeRankingModelsResult struct {
	Items    []SubscribeRankingModel `json:"items"`
	Metadata *core.ResultMetadata    `json:"metadata"`
}

type DescribeSubscribeRankingModelsAsyncResult struct {
	result *DescribeSubscribeRankingModelsResult
	err    error
}

func NewDescribeSubscribeRankingModelsResultFromJson(data string) DescribeSubscribeRankingModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingModelsResultFromDict(dict)
}

func NewDescribeSubscribeRankingModelsResultFromDict(data map[string]interface{}) DescribeSubscribeRankingModelsResult {
	return DescribeSubscribeRankingModelsResult{
		Items: func() []SubscribeRankingModel {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeSubscribeRankingModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeSubscribeRankingModelsResult) Pointer() *DescribeSubscribeRankingModelsResult {
	return &p
}

type GetSubscribeRankingModelResult struct {
	Item     *SubscribeRankingModel `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetSubscribeRankingModelAsyncResult struct {
	result *GetSubscribeRankingModelResult
	err    error
}

func NewGetSubscribeRankingModelResultFromJson(data string) GetSubscribeRankingModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingModelResultFromDict(dict)
}

func NewGetSubscribeRankingModelResultFromDict(data map[string]interface{}) GetSubscribeRankingModelResult {
	return GetSubscribeRankingModelResult{
		Item: func() *SubscribeRankingModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingModelResult) Pointer() *GetSubscribeRankingModelResult {
	return &p
}

type DescribeSubscribeRankingModelMastersResult struct {
	Items         []SubscribeRankingModelMaster `json:"items"`
	NextPageToken *string                       `json:"nextPageToken"`
	Metadata      *core.ResultMetadata          `json:"metadata"`
}

type DescribeSubscribeRankingModelMastersAsyncResult struct {
	result *DescribeSubscribeRankingModelMastersResult
	err    error
}

func NewDescribeSubscribeRankingModelMastersResultFromJson(data string) DescribeSubscribeRankingModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingModelMastersResultFromDict(dict)
}

func NewDescribeSubscribeRankingModelMastersResultFromDict(data map[string]interface{}) DescribeSubscribeRankingModelMastersResult {
	return DescribeSubscribeRankingModelMastersResult{
		Items: func() []SubscribeRankingModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingModelMasters(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribeRankingModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribeRankingModelMastersResult) Pointer() *DescribeSubscribeRankingModelMastersResult {
	return &p
}

type CreateSubscribeRankingModelMasterResult struct {
	Item     *SubscribeRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type CreateSubscribeRankingModelMasterAsyncResult struct {
	result *CreateSubscribeRankingModelMasterResult
	err    error
}

func NewCreateSubscribeRankingModelMasterResultFromJson(data string) CreateSubscribeRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateSubscribeRankingModelMasterResultFromDict(dict)
}

func NewCreateSubscribeRankingModelMasterResultFromDict(data map[string]interface{}) CreateSubscribeRankingModelMasterResult {
	return CreateSubscribeRankingModelMasterResult{
		Item: func() *SubscribeRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateSubscribeRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateSubscribeRankingModelMasterResult) Pointer() *CreateSubscribeRankingModelMasterResult {
	return &p
}

type GetSubscribeRankingModelMasterResult struct {
	Item     *SubscribeRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type GetSubscribeRankingModelMasterAsyncResult struct {
	result *GetSubscribeRankingModelMasterResult
	err    error
}

func NewGetSubscribeRankingModelMasterResultFromJson(data string) GetSubscribeRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingModelMasterResultFromDict(dict)
}

func NewGetSubscribeRankingModelMasterResultFromDict(data map[string]interface{}) GetSubscribeRankingModelMasterResult {
	return GetSubscribeRankingModelMasterResult{
		Item: func() *SubscribeRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingModelMasterResult) Pointer() *GetSubscribeRankingModelMasterResult {
	return &p
}

type UpdateSubscribeRankingModelMasterResult struct {
	Item     *SubscribeRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type UpdateSubscribeRankingModelMasterAsyncResult struct {
	result *UpdateSubscribeRankingModelMasterResult
	err    error
}

func NewUpdateSubscribeRankingModelMasterResultFromJson(data string) UpdateSubscribeRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateSubscribeRankingModelMasterResultFromDict(dict)
}

func NewUpdateSubscribeRankingModelMasterResultFromDict(data map[string]interface{}) UpdateSubscribeRankingModelMasterResult {
	return UpdateSubscribeRankingModelMasterResult{
		Item: func() *SubscribeRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateSubscribeRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateSubscribeRankingModelMasterResult) Pointer() *UpdateSubscribeRankingModelMasterResult {
	return &p
}

type DeleteSubscribeRankingModelMasterResult struct {
	Item     *SubscribeRankingModelMaster `json:"item"`
	Metadata *core.ResultMetadata         `json:"metadata"`
}

type DeleteSubscribeRankingModelMasterAsyncResult struct {
	result *DeleteSubscribeRankingModelMasterResult
	err    error
}

func NewDeleteSubscribeRankingModelMasterResultFromJson(data string) DeleteSubscribeRankingModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSubscribeRankingModelMasterResultFromDict(dict)
}

func NewDeleteSubscribeRankingModelMasterResultFromDict(data map[string]interface{}) DeleteSubscribeRankingModelMasterResult {
	return DeleteSubscribeRankingModelMasterResult{
		Item: func() *SubscribeRankingModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteSubscribeRankingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteSubscribeRankingModelMasterResult) Pointer() *DeleteSubscribeRankingModelMasterResult {
	return &p
}

type DescribeSubscribesResult struct {
	Items         []SubscribeUser      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscribesAsyncResult struct {
	result *DescribeSubscribesResult
	err    error
}

func NewDescribeSubscribesResultFromJson(data string) DescribeSubscribesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesResultFromDict(dict)
}

func NewDescribeSubscribesResultFromDict(data map[string]interface{}) DescribeSubscribesResult {
	return DescribeSubscribesResult{
		Items: func() []SubscribeUser {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeUsers(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribesResult) Pointer() *DescribeSubscribesResult {
	return &p
}

type DescribeSubscribesByUserIdResult struct {
	Items         []SubscribeUser      `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeSubscribesByUserIdAsyncResult struct {
	result *DescribeSubscribesByUserIdResult
	err    error
}

func NewDescribeSubscribesByUserIdResultFromJson(data string) DescribeSubscribesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByUserIdResultFromDict(dict)
}

func NewDescribeSubscribesByUserIdResultFromDict(data map[string]interface{}) DescribeSubscribesByUserIdResult {
	return DescribeSubscribesByUserIdResult{
		Items: func() []SubscribeUser {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeUsers(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribesByUserIdResult) Pointer() *DescribeSubscribesByUserIdResult {
	return &p
}

type AddSubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddSubscribeAsyncResult struct {
	result *AddSubscribeResult
	err    error
}

func NewAddSubscribeResultFromJson(data string) AddSubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddSubscribeResultFromDict(dict)
}

func NewAddSubscribeResultFromDict(data map[string]interface{}) AddSubscribeResult {
	return AddSubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p AddSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p AddSubscribeResult) Pointer() *AddSubscribeResult {
	return &p
}

type AddSubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type AddSubscribeByUserIdAsyncResult struct {
	result *AddSubscribeByUserIdResult
	err    error
}

func NewAddSubscribeByUserIdResultFromJson(data string) AddSubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddSubscribeByUserIdResultFromDict(dict)
}

func NewAddSubscribeByUserIdResultFromDict(data map[string]interface{}) AddSubscribeByUserIdResult {
	return AddSubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p AddSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p AddSubscribeByUserIdResult) Pointer() *AddSubscribeByUserIdResult {
	return &p
}

type DescribeSubscribeRankingScoresResult struct {
	Items         []SubscribeRankingScore `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeSubscribeRankingScoresAsyncResult struct {
	result *DescribeSubscribeRankingScoresResult
	err    error
}

func NewDescribeSubscribeRankingScoresResultFromJson(data string) DescribeSubscribeRankingScoresResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingScoresResultFromDict(dict)
}

func NewDescribeSubscribeRankingScoresResultFromDict(data map[string]interface{}) DescribeSubscribeRankingScoresResult {
	return DescribeSubscribeRankingScoresResult{
		Items: func() []SubscribeRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribeRankingScoresResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribeRankingScoresResult) Pointer() *DescribeSubscribeRankingScoresResult {
	return &p
}

type DescribeSubscribeRankingScoresByUserIdResult struct {
	Items         []SubscribeRankingScore `json:"items"`
	NextPageToken *string                 `json:"nextPageToken"`
	Metadata      *core.ResultMetadata    `json:"metadata"`
}

type DescribeSubscribeRankingScoresByUserIdAsyncResult struct {
	result *DescribeSubscribeRankingScoresByUserIdResult
	err    error
}

func NewDescribeSubscribeRankingScoresByUserIdResultFromJson(data string) DescribeSubscribeRankingScoresByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingScoresByUserIdResultFromDict(dict)
}

func NewDescribeSubscribeRankingScoresByUserIdResultFromDict(data map[string]interface{}) DescribeSubscribeRankingScoresByUserIdResult {
	return DescribeSubscribeRankingScoresByUserIdResult{
		Items: func() []SubscribeRankingScore {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingScores(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribeRankingScoresByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribeRankingScoresByUserIdResult) Pointer() *DescribeSubscribeRankingScoresByUserIdResult {
	return &p
}

type PutSubscribeRankingScoreResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type PutSubscribeRankingScoreAsyncResult struct {
	result *PutSubscribeRankingScoreResult
	err    error
}

func NewPutSubscribeRankingScoreResultFromJson(data string) PutSubscribeRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutSubscribeRankingScoreResultFromDict(dict)
}

func NewPutSubscribeRankingScoreResultFromDict(data map[string]interface{}) PutSubscribeRankingScoreResult {
	return PutSubscribeRankingScoreResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutSubscribeRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutSubscribeRankingScoreResult) Pointer() *PutSubscribeRankingScoreResult {
	return &p
}

type PutSubscribeRankingScoreByUserIdResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type PutSubscribeRankingScoreByUserIdAsyncResult struct {
	result *PutSubscribeRankingScoreByUserIdResult
	err    error
}

func NewPutSubscribeRankingScoreByUserIdResultFromJson(data string) PutSubscribeRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutSubscribeRankingScoreByUserIdResultFromDict(dict)
}

func NewPutSubscribeRankingScoreByUserIdResultFromDict(data map[string]interface{}) PutSubscribeRankingScoreByUserIdResult {
	return PutSubscribeRankingScoreByUserIdResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p PutSubscribeRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p PutSubscribeRankingScoreByUserIdResult) Pointer() *PutSubscribeRankingScoreByUserIdResult {
	return &p
}

type GetSubscribeRankingScoreResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetSubscribeRankingScoreAsyncResult struct {
	result *GetSubscribeRankingScoreResult
	err    error
}

func NewGetSubscribeRankingScoreResultFromJson(data string) GetSubscribeRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingScoreResultFromDict(dict)
}

func NewGetSubscribeRankingScoreResultFromDict(data map[string]interface{}) GetSubscribeRankingScoreResult {
	return GetSubscribeRankingScoreResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingScoreResult) Pointer() *GetSubscribeRankingScoreResult {
	return &p
}

type GetSubscribeRankingScoreByUserIdResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type GetSubscribeRankingScoreByUserIdAsyncResult struct {
	result *GetSubscribeRankingScoreByUserIdResult
	err    error
}

func NewGetSubscribeRankingScoreByUserIdResultFromJson(data string) GetSubscribeRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingScoreByUserIdResultFromDict(dict)
}

func NewGetSubscribeRankingScoreByUserIdResultFromDict(data map[string]interface{}) GetSubscribeRankingScoreByUserIdResult {
	return GetSubscribeRankingScoreByUserIdResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingScoreByUserIdResult) Pointer() *GetSubscribeRankingScoreByUserIdResult {
	return &p
}

type DeleteSubscribeRankingScoreByUserIdResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type DeleteSubscribeRankingScoreByUserIdAsyncResult struct {
	result *DeleteSubscribeRankingScoreByUserIdResult
	err    error
}

func NewDeleteSubscribeRankingScoreByUserIdResultFromJson(data string) DeleteSubscribeRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSubscribeRankingScoreByUserIdResultFromDict(dict)
}

func NewDeleteSubscribeRankingScoreByUserIdResultFromDict(data map[string]interface{}) DeleteSubscribeRankingScoreByUserIdResult {
	return DeleteSubscribeRankingScoreByUserIdResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteSubscribeRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteSubscribeRankingScoreByUserIdResult) Pointer() *DeleteSubscribeRankingScoreByUserIdResult {
	return &p
}

type VerifySubscribeRankingScoreResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type VerifySubscribeRankingScoreAsyncResult struct {
	result *VerifySubscribeRankingScoreResult
	err    error
}

func NewVerifySubscribeRankingScoreResultFromJson(data string) VerifySubscribeRankingScoreResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySubscribeRankingScoreResultFromDict(dict)
}

func NewVerifySubscribeRankingScoreResultFromDict(data map[string]interface{}) VerifySubscribeRankingScoreResult {
	return VerifySubscribeRankingScoreResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifySubscribeRankingScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifySubscribeRankingScoreResult) Pointer() *VerifySubscribeRankingScoreResult {
	return &p
}

type VerifySubscribeRankingScoreByUserIdResult struct {
	Item     *SubscribeRankingScore `json:"item"`
	Metadata *core.ResultMetadata   `json:"metadata"`
}

type VerifySubscribeRankingScoreByUserIdAsyncResult struct {
	result *VerifySubscribeRankingScoreByUserIdResult
	err    error
}

func NewVerifySubscribeRankingScoreByUserIdResultFromJson(data string) VerifySubscribeRankingScoreByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySubscribeRankingScoreByUserIdResultFromDict(dict)
}

func NewVerifySubscribeRankingScoreByUserIdResultFromDict(data map[string]interface{}) VerifySubscribeRankingScoreByUserIdResult {
	return VerifySubscribeRankingScoreByUserIdResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p VerifySubscribeRankingScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifySubscribeRankingScoreByUserIdResult) Pointer() *VerifySubscribeRankingScoreByUserIdResult {
	return &p
}

type VerifySubscribeRankingScoreByStampTaskResult struct {
	Item            *SubscribeRankingScore `json:"item"`
	NewContextStack *string                `json:"newContextStack"`
	Metadata        *core.ResultMetadata   `json:"metadata"`
}

type VerifySubscribeRankingScoreByStampTaskAsyncResult struct {
	result *VerifySubscribeRankingScoreByStampTaskResult
	err    error
}

func NewVerifySubscribeRankingScoreByStampTaskResultFromJson(data string) VerifySubscribeRankingScoreByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifySubscribeRankingScoreByStampTaskResultFromDict(dict)
}

func NewVerifySubscribeRankingScoreByStampTaskResultFromDict(data map[string]interface{}) VerifySubscribeRankingScoreByStampTaskResult {
	return VerifySubscribeRankingScoreByStampTaskResult{
		Item: func() *SubscribeRankingScore {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingScoreFromDict(core.CastMap(data["item"])).Pointer()
		}(),
		NewContextStack: func() *string {
			v, ok := data["newContextStack"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["newContextStack"])
		}(),
	}
}

func (p VerifySubscribeRankingScoreByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"newContextStack": p.NewContextStack,
	}
}

func (p VerifySubscribeRankingScoreByStampTaskResult) Pointer() *VerifySubscribeRankingScoreByStampTaskResult {
	return &p
}

type DescribeSubscribeRankingsResult struct {
	Items         []SubscribeRankingData `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeSubscribeRankingsAsyncResult struct {
	result *DescribeSubscribeRankingsResult
	err    error
}

func NewDescribeSubscribeRankingsResultFromJson(data string) DescribeSubscribeRankingsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingsResultFromDict(dict)
}

func NewDescribeSubscribeRankingsResultFromDict(data map[string]interface{}) DescribeSubscribeRankingsResult {
	return DescribeSubscribeRankingsResult{
		Items: func() []SubscribeRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribeRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribeRankingsResult) Pointer() *DescribeSubscribeRankingsResult {
	return &p
}

type DescribeSubscribeRankingsByUserIdResult struct {
	Items         []SubscribeRankingData `json:"items"`
	NextPageToken *string                `json:"nextPageToken"`
	Metadata      *core.ResultMetadata   `json:"metadata"`
}

type DescribeSubscribeRankingsByUserIdAsyncResult struct {
	result *DescribeSubscribeRankingsByUserIdResult
	err    error
}

func NewDescribeSubscribeRankingsByUserIdResultFromJson(data string) DescribeSubscribeRankingsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribeRankingsByUserIdResultFromDict(dict)
}

func NewDescribeSubscribeRankingsByUserIdResultFromDict(data map[string]interface{}) DescribeSubscribeRankingsByUserIdResult {
	return DescribeSubscribeRankingsByUserIdResult{
		Items: func() []SubscribeRankingData {
			if data["items"] == nil {
				return nil
			}
			return CastSubscribeRankingDatas(core.CastArray(data["items"]))
		}(),
		NextPageToken: func() *string {
			v, ok := data["nextPageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["nextPageToken"])
		}(),
	}
}

func (p DescribeSubscribeRankingsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeRankingDatasFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeSubscribeRankingsByUserIdResult) Pointer() *DescribeSubscribeRankingsByUserIdResult {
	return &p
}

type GetSubscribeRankingResult struct {
	Item     *SubscribeRankingData `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetSubscribeRankingAsyncResult struct {
	result *GetSubscribeRankingResult
	err    error
}

func NewGetSubscribeRankingResultFromJson(data string) GetSubscribeRankingResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingResultFromDict(dict)
}

func NewGetSubscribeRankingResultFromDict(data map[string]interface{}) GetSubscribeRankingResult {
	return GetSubscribeRankingResult{
		Item: func() *SubscribeRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingResult) Pointer() *GetSubscribeRankingResult {
	return &p
}

type GetSubscribeRankingByUserIdResult struct {
	Item     *SubscribeRankingData `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetSubscribeRankingByUserIdAsyncResult struct {
	result *GetSubscribeRankingByUserIdResult
	err    error
}

func NewGetSubscribeRankingByUserIdResultFromJson(data string) GetSubscribeRankingByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRankingByUserIdResultFromDict(dict)
}

func NewGetSubscribeRankingByUserIdResultFromDict(data map[string]interface{}) GetSubscribeRankingByUserIdResult {
	return GetSubscribeRankingByUserIdResult{
		Item: func() *SubscribeRankingData {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeRankingDataFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeRankingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeRankingByUserIdResult) Pointer() *GetSubscribeRankingByUserIdResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
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
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
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
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentRankingMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetCurrentRankingMasterAsyncResult struct {
	result *GetCurrentRankingMasterResult
	err    error
}

func NewGetCurrentRankingMasterResultFromJson(data string) GetCurrentRankingMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRankingMasterResultFromDict(dict)
}

func NewGetCurrentRankingMasterResultFromDict(data map[string]interface{}) GetCurrentRankingMasterResult {
	return GetCurrentRankingMasterResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentRankingMasterResult) Pointer() *GetCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentRankingMasterAsyncResult struct {
	result *UpdateCurrentRankingMasterResult
	err    error
}

func NewUpdateCurrentRankingMasterResultFromJson(data string) UpdateCurrentRankingMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterResultFromDict(dict)
}

func NewUpdateCurrentRankingMasterResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterResult {
	return UpdateCurrentRankingMasterResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentRankingMasterResult) Pointer() *UpdateCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterFromGitHubResult struct {
	Item     *CurrentRankingMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentRankingMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRankingMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRankingMasterFromGitHubResultFromJson(data string) UpdateCurrentRankingMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRankingMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentRankingMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubResult {
	return UpdateCurrentRankingMasterFromGitHubResult{
		Item: func() *CurrentRankingMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) Pointer() *UpdateCurrentRankingMasterFromGitHubResult {
	return &p
}

type GetSubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

func NewGetSubscribeResultFromJson(data string) GetSubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeResultFromDict(dict)
}

func NewGetSubscribeResultFromDict(data map[string]interface{}) GetSubscribeResult {
	return GetSubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeResult) Pointer() *GetSubscribeResult {
	return &p
}

type GetSubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

func NewGetSubscribeByUserIdResultFromJson(data string) GetSubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeByUserIdResultFromDict(dict)
}

func NewGetSubscribeByUserIdResultFromDict(data map[string]interface{}) GetSubscribeByUserIdResult {
	return GetSubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetSubscribeByUserIdResult) Pointer() *GetSubscribeByUserIdResult {
	return &p
}

type DeleteSubscribeResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSubscribeAsyncResult struct {
	result *DeleteSubscribeResult
	err    error
}

func NewDeleteSubscribeResultFromJson(data string) DeleteSubscribeResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSubscribeResultFromDict(dict)
}

func NewDeleteSubscribeResultFromDict(data map[string]interface{}) DeleteSubscribeResult {
	return DeleteSubscribeResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteSubscribeResult) Pointer() *DeleteSubscribeResult {
	return &p
}

type DeleteSubscribeByUserIdResult struct {
	Item     *SubscribeUser       `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteSubscribeByUserIdAsyncResult struct {
	result *DeleteSubscribeByUserIdResult
	err    error
}

func NewDeleteSubscribeByUserIdResultFromJson(data string) DeleteSubscribeByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteSubscribeByUserIdResultFromDict(dict)
}

func NewDeleteSubscribeByUserIdResultFromDict(data map[string]interface{}) DeleteSubscribeByUserIdResult {
	return DeleteSubscribeByUserIdResult{
		Item: func() *SubscribeUser {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteSubscribeByUserIdResult) Pointer() *DeleteSubscribeByUserIdResult {
	return &p
}
