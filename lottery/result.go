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

package lottery

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

type DescribeLotteryModelMastersResult struct {
	Items         []LotteryModelMaster `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeLotteryModelMastersAsyncResult struct {
	result *DescribeLotteryModelMastersResult
	err    error
}

func NewDescribeLotteryModelMastersResultFromJson(data string) DescribeLotteryModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryModelMastersResultFromDict(dict)
}

func NewDescribeLotteryModelMastersResultFromDict(data map[string]interface{}) DescribeLotteryModelMastersResult {
	return DescribeLotteryModelMastersResult{
		Items: func() []LotteryModelMaster {
			if data["items"] == nil {
				return nil
			}
			return CastLotteryModelMasters(core.CastArray(data["items"]))
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

func (p DescribeLotteryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLotteryModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeLotteryModelMastersResult) Pointer() *DescribeLotteryModelMastersResult {
	return &p
}

type CreateLotteryModelMasterResult struct {
	Item     *LotteryModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreateLotteryModelMasterAsyncResult struct {
	result *CreateLotteryModelMasterResult
	err    error
}

func NewCreateLotteryModelMasterResultFromJson(data string) CreateLotteryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateLotteryModelMasterResultFromDict(dict)
}

func NewCreateLotteryModelMasterResultFromDict(data map[string]interface{}) CreateLotteryModelMasterResult {
	return CreateLotteryModelMasterResult{
		Item: func() *LotteryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreateLotteryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateLotteryModelMasterResult) Pointer() *CreateLotteryModelMasterResult {
	return &p
}

type GetLotteryModelMasterResult struct {
	Item     *LotteryModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLotteryModelMasterAsyncResult struct {
	result *GetLotteryModelMasterResult
	err    error
}

func NewGetLotteryModelMasterResultFromJson(data string) GetLotteryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryModelMasterResultFromDict(dict)
}

func NewGetLotteryModelMasterResultFromDict(data map[string]interface{}) GetLotteryModelMasterResult {
	return GetLotteryModelMasterResult{
		Item: func() *LotteryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetLotteryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetLotteryModelMasterResult) Pointer() *GetLotteryModelMasterResult {
	return &p
}

type UpdateLotteryModelMasterResult struct {
	Item     *LotteryModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdateLotteryModelMasterAsyncResult struct {
	result *UpdateLotteryModelMasterResult
	err    error
}

func NewUpdateLotteryModelMasterResultFromJson(data string) UpdateLotteryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateLotteryModelMasterResultFromDict(dict)
}

func NewUpdateLotteryModelMasterResultFromDict(data map[string]interface{}) UpdateLotteryModelMasterResult {
	return UpdateLotteryModelMasterResult{
		Item: func() *LotteryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateLotteryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateLotteryModelMasterResult) Pointer() *UpdateLotteryModelMasterResult {
	return &p
}

type DeleteLotteryModelMasterResult struct {
	Item     *LotteryModelMaster  `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeleteLotteryModelMasterAsyncResult struct {
	result *DeleteLotteryModelMasterResult
	err    error
}

func NewDeleteLotteryModelMasterResultFromJson(data string) DeleteLotteryModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteLotteryModelMasterResultFromDict(dict)
}

func NewDeleteLotteryModelMasterResultFromDict(data map[string]interface{}) DeleteLotteryModelMasterResult {
	return DeleteLotteryModelMasterResult{
		Item: func() *LotteryModelMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeleteLotteryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteLotteryModelMasterResult) Pointer() *DeleteLotteryModelMasterResult {
	return &p
}

type DescribePrizeTableMastersResult struct {
	Items         []PrizeTableMaster   `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribePrizeTableMastersAsyncResult struct {
	result *DescribePrizeTableMastersResult
	err    error
}

func NewDescribePrizeTableMastersResultFromJson(data string) DescribePrizeTableMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeTableMastersResultFromDict(dict)
}

func NewDescribePrizeTableMastersResultFromDict(data map[string]interface{}) DescribePrizeTableMastersResult {
	return DescribePrizeTableMastersResult{
		Items: func() []PrizeTableMaster {
			if data["items"] == nil {
				return nil
			}
			return CastPrizeTableMasters(core.CastArray(data["items"]))
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

func (p DescribePrizeTableMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPrizeTableMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePrizeTableMastersResult) Pointer() *DescribePrizeTableMastersResult {
	return &p
}

type CreatePrizeTableMasterResult struct {
	Item     *PrizeTableMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type CreatePrizeTableMasterAsyncResult struct {
	result *CreatePrizeTableMasterResult
	err    error
}

func NewCreatePrizeTableMasterResultFromJson(data string) CreatePrizeTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreatePrizeTableMasterResultFromDict(dict)
}

func NewCreatePrizeTableMasterResultFromDict(data map[string]interface{}) CreatePrizeTableMasterResult {
	return CreatePrizeTableMasterResult{
		Item: func() *PrizeTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p CreatePrizeTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreatePrizeTableMasterResult) Pointer() *CreatePrizeTableMasterResult {
	return &p
}

type GetPrizeTableMasterResult struct {
	Item     *PrizeTableMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPrizeTableMasterAsyncResult struct {
	result *GetPrizeTableMasterResult
	err    error
}

func NewGetPrizeTableMasterResultFromJson(data string) GetPrizeTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeTableMasterResultFromDict(dict)
}

func NewGetPrizeTableMasterResultFromDict(data map[string]interface{}) GetPrizeTableMasterResult {
	return GetPrizeTableMasterResult{
		Item: func() *PrizeTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetPrizeTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetPrizeTableMasterResult) Pointer() *GetPrizeTableMasterResult {
	return &p
}

type UpdatePrizeTableMasterResult struct {
	Item     *PrizeTableMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type UpdatePrizeTableMasterAsyncResult struct {
	result *UpdatePrizeTableMasterResult
	err    error
}

func NewUpdatePrizeTableMasterResultFromJson(data string) UpdatePrizeTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdatePrizeTableMasterResultFromDict(dict)
}

func NewUpdatePrizeTableMasterResultFromDict(data map[string]interface{}) UpdatePrizeTableMasterResult {
	return UpdatePrizeTableMasterResult{
		Item: func() *PrizeTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdatePrizeTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdatePrizeTableMasterResult) Pointer() *UpdatePrizeTableMasterResult {
	return &p
}

type DeletePrizeTableMasterResult struct {
	Item     *PrizeTableMaster    `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DeletePrizeTableMasterAsyncResult struct {
	result *DeletePrizeTableMasterResult
	err    error
}

func NewDeletePrizeTableMasterResultFromJson(data string) DeletePrizeTableMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeletePrizeTableMasterResultFromDict(dict)
}

func NewDeletePrizeTableMasterResultFromDict(data map[string]interface{}) DeletePrizeTableMasterResult {
	return DeletePrizeTableMasterResult{
		Item: func() *PrizeTableMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p DeletePrizeTableMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeletePrizeTableMasterResult) Pointer() *DeletePrizeTableMasterResult {
	return &p
}

type DescribeLotteryModelsResult struct {
	Items    []LotteryModel       `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeLotteryModelsAsyncResult struct {
	result *DescribeLotteryModelsResult
	err    error
}

func NewDescribeLotteryModelsResultFromJson(data string) DescribeLotteryModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeLotteryModelsResultFromDict(dict)
}

func NewDescribeLotteryModelsResultFromDict(data map[string]interface{}) DescribeLotteryModelsResult {
	return DescribeLotteryModelsResult{
		Items: func() []LotteryModel {
			if data["items"] == nil {
				return nil
			}
			return CastLotteryModels(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeLotteryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastLotteryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeLotteryModelsResult) Pointer() *DescribeLotteryModelsResult {
	return &p
}

type GetLotteryModelResult struct {
	Item     *LotteryModel        `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetLotteryModelAsyncResult struct {
	result *GetLotteryModelResult
	err    error
}

func NewGetLotteryModelResultFromJson(data string) GetLotteryModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetLotteryModelResultFromDict(dict)
}

func NewGetLotteryModelResultFromDict(data map[string]interface{}) GetLotteryModelResult {
	return GetLotteryModelResult{
		Item: func() *LotteryModel {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewLotteryModelFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetLotteryModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetLotteryModelResult) Pointer() *GetLotteryModelResult {
	return &p
}

type DescribePrizeTablesResult struct {
	Items    []PrizeTable         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribePrizeTablesAsyncResult struct {
	result *DescribePrizeTablesResult
	err    error
}

func NewDescribePrizeTablesResultFromJson(data string) DescribePrizeTablesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeTablesResultFromDict(dict)
}

func NewDescribePrizeTablesResultFromDict(data map[string]interface{}) DescribePrizeTablesResult {
	return DescribePrizeTablesResult{
		Items: func() []PrizeTable {
			if data["items"] == nil {
				return nil
			}
			return CastPrizeTables(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribePrizeTablesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPrizeTablesFromDict(
			p.Items,
		),
	}
}

func (p DescribePrizeTablesResult) Pointer() *DescribePrizeTablesResult {
	return &p
}

type GetPrizeTableResult struct {
	Item     *PrizeTable          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPrizeTableAsyncResult struct {
	result *GetPrizeTableResult
	err    error
}

func NewGetPrizeTableResultFromJson(data string) GetPrizeTableResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeTableResultFromDict(dict)
}

func NewGetPrizeTableResultFromDict(data map[string]interface{}) GetPrizeTableResult {
	return GetPrizeTableResult{
		Item: func() *PrizeTable {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeTableFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetPrizeTableResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetPrizeTableResult) Pointer() *GetPrizeTableResult {
	return &p
}

type DrawByUserIdResult struct {
	Items                     []DrawnPrize            `json:"items"`
	BoxItems                  *BoxItems               `json:"boxItems"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DrawByUserIdAsyncResult struct {
	result *DrawByUserIdResult
	err    error
}

func NewDrawByUserIdResultFromJson(data string) DrawByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawByUserIdResultFromDict(dict)
}

func NewDrawByUserIdResultFromDict(data map[string]interface{}) DrawByUserIdResult {
	return DrawByUserIdResult{
		Items: func() []DrawnPrize {
			if data["items"] == nil {
				return nil
			}
			return CastDrawnPrizes(core.CastArray(data["items"]))
		}(),
		BoxItems: func() *BoxItems {
			v, ok := data["boxItems"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["boxItems"])).Pointer()
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

func (p DrawByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDrawnPrizesFromDict(
			p.Items,
		),
		"boxItems": func() map[string]interface{} {
			if p.BoxItems == nil {
				return nil
			}
			return p.BoxItems.ToDict()
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
	}
}

func (p DrawByUserIdResult) Pointer() *DrawByUserIdResult {
	return &p
}

type PredictionResult struct {
	Items    []DrawnPrize         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PredictionAsyncResult struct {
	result *PredictionResult
	err    error
}

func NewPredictionResultFromJson(data string) PredictionResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionResultFromDict(dict)
}

func NewPredictionResultFromDict(data map[string]interface{}) PredictionResult {
	return PredictionResult{
		Items: func() []DrawnPrize {
			if data["items"] == nil {
				return nil
			}
			return CastDrawnPrizes(core.CastArray(data["items"]))
		}(),
	}
}

func (p PredictionResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDrawnPrizesFromDict(
			p.Items,
		),
	}
}

func (p PredictionResult) Pointer() *PredictionResult {
	return &p
}

type PredictionByUserIdResult struct {
	Items    []DrawnPrize         `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type PredictionByUserIdAsyncResult struct {
	result *PredictionByUserIdResult
	err    error
}

func NewPredictionByUserIdResultFromJson(data string) PredictionByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPredictionByUserIdResultFromDict(dict)
}

func NewPredictionByUserIdResultFromDict(data map[string]interface{}) PredictionByUserIdResult {
	return PredictionByUserIdResult{
		Items: func() []DrawnPrize {
			if data["items"] == nil {
				return nil
			}
			return CastDrawnPrizes(core.CastArray(data["items"]))
		}(),
	}
}

func (p PredictionByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDrawnPrizesFromDict(
			p.Items,
		),
	}
}

func (p PredictionByUserIdResult) Pointer() *PredictionByUserIdResult {
	return &p
}

type DrawWithRandomSeedByUserIdResult struct {
	Items                     []DrawnPrize            `json:"items"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DrawWithRandomSeedByUserIdAsyncResult struct {
	result *DrawWithRandomSeedByUserIdResult
	err    error
}

func NewDrawWithRandomSeedByUserIdResultFromJson(data string) DrawWithRandomSeedByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawWithRandomSeedByUserIdResultFromDict(dict)
}

func NewDrawWithRandomSeedByUserIdResultFromDict(data map[string]interface{}) DrawWithRandomSeedByUserIdResult {
	return DrawWithRandomSeedByUserIdResult{
		Items: func() []DrawnPrize {
			if data["items"] == nil {
				return nil
			}
			return CastDrawnPrizes(core.CastArray(data["items"]))
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

func (p DrawWithRandomSeedByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDrawnPrizesFromDict(
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
	}
}

func (p DrawWithRandomSeedByUserIdResult) Pointer() *DrawWithRandomSeedByUserIdResult {
	return &p
}

type DrawByStampSheetResult struct {
	Items                     []DrawnPrize            `json:"items"`
	BoxItems                  *BoxItems               `json:"boxItems"`
	TransactionId             *string                 `json:"transactionId"`
	StampSheet                *string                 `json:"stampSheet"`
	StampSheetEncryptionKeyId *string                 `json:"stampSheetEncryptionKeyId"`
	AutoRunStampSheet         *bool                   `json:"autoRunStampSheet"`
	AtomicCommit              *bool                   `json:"atomicCommit"`
	Transaction               *string                 `json:"transaction"`
	TransactionResult         *core.TransactionResult `json:"transactionResult"`
	Metadata                  *core.ResultMetadata    `json:"metadata"`
}

type DrawByStampSheetAsyncResult struct {
	result *DrawByStampSheetResult
	err    error
}

func NewDrawByStampSheetResultFromJson(data string) DrawByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDrawByStampSheetResultFromDict(dict)
}

func NewDrawByStampSheetResultFromDict(data map[string]interface{}) DrawByStampSheetResult {
	return DrawByStampSheetResult{
		Items: func() []DrawnPrize {
			if data["items"] == nil {
				return nil
			}
			return CastDrawnPrizes(core.CastArray(data["items"]))
		}(),
		BoxItems: func() *BoxItems {
			v, ok := data["boxItems"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["boxItems"])).Pointer()
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

func (p DrawByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDrawnPrizesFromDict(
			p.Items,
		),
		"boxItems": func() map[string]interface{} {
			if p.BoxItems == nil {
				return nil
			}
			return p.BoxItems.ToDict()
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
	}
}

func (p DrawByStampSheetResult) Pointer() *DrawByStampSheetResult {
	return &p
}

type DescribeProbabilitiesResult struct {
	Items    []Probability        `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeProbabilitiesAsyncResult struct {
	result *DescribeProbabilitiesResult
	err    error
}

func NewDescribeProbabilitiesResultFromJson(data string) DescribeProbabilitiesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProbabilitiesResultFromDict(dict)
}

func NewDescribeProbabilitiesResultFromDict(data map[string]interface{}) DescribeProbabilitiesResult {
	return DescribeProbabilitiesResult{
		Items: func() []Probability {
			if data["items"] == nil {
				return nil
			}
			return CastProbabilities(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeProbabilitiesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProbabilitiesFromDict(
			p.Items,
		),
	}
}

func (p DescribeProbabilitiesResult) Pointer() *DescribeProbabilitiesResult {
	return &p
}

type DescribeProbabilitiesByUserIdResult struct {
	Items    []Probability        `json:"items"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type DescribeProbabilitiesByUserIdAsyncResult struct {
	result *DescribeProbabilitiesByUserIdResult
	err    error
}

func NewDescribeProbabilitiesByUserIdResultFromJson(data string) DescribeProbabilitiesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeProbabilitiesByUserIdResultFromDict(dict)
}

func NewDescribeProbabilitiesByUserIdResultFromDict(data map[string]interface{}) DescribeProbabilitiesByUserIdResult {
	return DescribeProbabilitiesByUserIdResult{
		Items: func() []Probability {
			if data["items"] == nil {
				return nil
			}
			return CastProbabilities(core.CastArray(data["items"]))
		}(),
	}
}

func (p DescribeProbabilitiesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastProbabilitiesFromDict(
			p.Items,
		),
	}
}

func (p DescribeProbabilitiesByUserIdResult) Pointer() *DescribeProbabilitiesByUserIdResult {
	return &p
}

type ExportMasterResult struct {
	Item     *CurrentLotteryMaster `json:"item"`
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
		Item: func() *CurrentLotteryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer()
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

type GetCurrentLotteryMasterResult struct {
	Item     *CurrentLotteryMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type GetCurrentLotteryMasterAsyncResult struct {
	result *GetCurrentLotteryMasterResult
	err    error
}

func NewGetCurrentLotteryMasterResultFromJson(data string) GetCurrentLotteryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentLotteryMasterResultFromDict(dict)
}

func NewGetCurrentLotteryMasterResultFromDict(data map[string]interface{}) GetCurrentLotteryMasterResult {
	return GetCurrentLotteryMasterResult{
		Item: func() *CurrentLotteryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetCurrentLotteryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentLotteryMasterResult) Pointer() *GetCurrentLotteryMasterResult {
	return &p
}

type UpdateCurrentLotteryMasterResult struct {
	Item     *CurrentLotteryMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentLotteryMasterAsyncResult struct {
	result *UpdateCurrentLotteryMasterResult
	err    error
}

func NewUpdateCurrentLotteryMasterResultFromJson(data string) UpdateCurrentLotteryMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLotteryMasterResultFromDict(dict)
}

func NewUpdateCurrentLotteryMasterResultFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterResult {
	return UpdateCurrentLotteryMasterResult{
		Item: func() *CurrentLotteryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentLotteryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentLotteryMasterResult) Pointer() *UpdateCurrentLotteryMasterResult {
	return &p
}

type UpdateCurrentLotteryMasterFromGitHubResult struct {
	Item     *CurrentLotteryMaster `json:"item"`
	Metadata *core.ResultMetadata  `json:"metadata"`
}

type UpdateCurrentLotteryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLotteryMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentLotteryMasterFromGitHubResultFromJson(data string) UpdateCurrentLotteryMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentLotteryMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentLotteryMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterFromGitHubResult {
	return UpdateCurrentLotteryMasterFromGitHubResult{
		Item: func() *CurrentLotteryMaster {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p UpdateCurrentLotteryMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentLotteryMasterFromGitHubResult) Pointer() *UpdateCurrentLotteryMasterFromGitHubResult {
	return &p
}

type DescribePrizeLimitsResult struct {
	Items         []PrizeLimit         `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribePrizeLimitsAsyncResult struct {
	result *DescribePrizeLimitsResult
	err    error
}

func NewDescribePrizeLimitsResultFromJson(data string) DescribePrizeLimitsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribePrizeLimitsResultFromDict(dict)
}

func NewDescribePrizeLimitsResultFromDict(data map[string]interface{}) DescribePrizeLimitsResult {
	return DescribePrizeLimitsResult{
		Items: func() []PrizeLimit {
			if data["items"] == nil {
				return nil
			}
			return CastPrizeLimits(core.CastArray(data["items"]))
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

func (p DescribePrizeLimitsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastPrizeLimitsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribePrizeLimitsResult) Pointer() *DescribePrizeLimitsResult {
	return &p
}

type GetPrizeLimitResult struct {
	Item     *PrizeLimit          `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetPrizeLimitAsyncResult struct {
	result *GetPrizeLimitResult
	err    error
}

func NewGetPrizeLimitResultFromJson(data string) GetPrizeLimitResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetPrizeLimitResultFromDict(dict)
}

func NewGetPrizeLimitResultFromDict(data map[string]interface{}) GetPrizeLimitResult {
	return GetPrizeLimitResult{
		Item: func() *PrizeLimit {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewPrizeLimitFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetPrizeLimitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetPrizeLimitResult) Pointer() *GetPrizeLimitResult {
	return &p
}

type ResetPrizeLimitResult struct {
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetPrizeLimitAsyncResult struct {
	result *ResetPrizeLimitResult
	err    error
}

func NewResetPrizeLimitResultFromJson(data string) ResetPrizeLimitResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetPrizeLimitResultFromDict(dict)
}

func NewResetPrizeLimitResultFromDict(data map[string]interface{}) ResetPrizeLimitResult {
	return ResetPrizeLimitResult{}
}

func (p ResetPrizeLimitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ResetPrizeLimitResult) Pointer() *ResetPrizeLimitResult {
	return &p
}

type DescribeBoxesResult struct {
	Items         []BoxItems           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBoxesAsyncResult struct {
	result *DescribeBoxesResult
	err    error
}

func NewDescribeBoxesResultFromJson(data string) DescribeBoxesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBoxesResultFromDict(dict)
}

func NewDescribeBoxesResultFromDict(data map[string]interface{}) DescribeBoxesResult {
	return DescribeBoxesResult{
		Items: func() []BoxItems {
			if data["items"] == nil {
				return nil
			}
			return CastBoxItemses(core.CastArray(data["items"]))
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

func (p DescribeBoxesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBoxItemsesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeBoxesResult) Pointer() *DescribeBoxesResult {
	return &p
}

type DescribeBoxesByUserIdResult struct {
	Items         []BoxItems           `json:"items"`
	NextPageToken *string              `json:"nextPageToken"`
	Metadata      *core.ResultMetadata `json:"metadata"`
}

type DescribeBoxesByUserIdAsyncResult struct {
	result *DescribeBoxesByUserIdResult
	err    error
}

func NewDescribeBoxesByUserIdResultFromJson(data string) DescribeBoxesByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeBoxesByUserIdResultFromDict(dict)
}

func NewDescribeBoxesByUserIdResultFromDict(data map[string]interface{}) DescribeBoxesByUserIdResult {
	return DescribeBoxesByUserIdResult{
		Items: func() []BoxItems {
			if data["items"] == nil {
				return nil
			}
			return CastBoxItemses(core.CastArray(data["items"]))
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

func (p DescribeBoxesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastBoxItemsesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeBoxesByUserIdResult) Pointer() *DescribeBoxesByUserIdResult {
	return &p
}

type GetBoxResult struct {
	Item     *BoxItems            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBoxAsyncResult struct {
	result *GetBoxResult
	err    error
}

func NewGetBoxResultFromJson(data string) GetBoxResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBoxResultFromDict(dict)
}

func NewGetBoxResultFromDict(data map[string]interface{}) GetBoxResult {
	return GetBoxResult{
		Item: func() *BoxItems {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetBoxResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetBoxResult) Pointer() *GetBoxResult {
	return &p
}

type GetBoxByUserIdResult struct {
	Item     *BoxItems            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type GetBoxByUserIdAsyncResult struct {
	result *GetBoxByUserIdResult
	err    error
}

func NewGetBoxByUserIdResultFromJson(data string) GetBoxByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBoxByUserIdResultFromDict(dict)
}

func NewGetBoxByUserIdResultFromDict(data map[string]interface{}) GetBoxByUserIdResult {
	return GetBoxByUserIdResult{
		Item: func() *BoxItems {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p GetBoxByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetBoxByUserIdResult) Pointer() *GetBoxByUserIdResult {
	return &p
}

type ResetBoxResult struct {
	Item     *BoxItems            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetBoxAsyncResult struct {
	result *ResetBoxResult
	err    error
}

func NewResetBoxResultFromJson(data string) ResetBoxResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetBoxResultFromDict(dict)
}

func NewResetBoxResultFromDict(data map[string]interface{}) ResetBoxResult {
	return ResetBoxResult{
		Item: func() *BoxItems {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ResetBoxResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ResetBoxResult) Pointer() *ResetBoxResult {
	return &p
}

type ResetBoxByUserIdResult struct {
	Item     *BoxItems            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetBoxByUserIdAsyncResult struct {
	result *ResetBoxByUserIdResult
	err    error
}

func NewResetBoxByUserIdResultFromJson(data string) ResetBoxByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetBoxByUserIdResultFromDict(dict)
}

func NewResetBoxByUserIdResultFromDict(data map[string]interface{}) ResetBoxByUserIdResult {
	return ResetBoxByUserIdResult{
		Item: func() *BoxItems {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ResetBoxByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ResetBoxByUserIdResult) Pointer() *ResetBoxByUserIdResult {
	return &p
}

type ResetByStampSheetResult struct {
	Item     *BoxItems            `json:"item"`
	Metadata *core.ResultMetadata `json:"metadata"`
}

type ResetByStampSheetAsyncResult struct {
	result *ResetByStampSheetResult
	err    error
}

func NewResetByStampSheetResultFromJson(data string) ResetByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewResetByStampSheetResultFromDict(dict)
}

func NewResetByStampSheetResultFromDict(data map[string]interface{}) ResetByStampSheetResult {
	return ResetByStampSheetResult{
		Item: func() *BoxItems {
			v, ok := data["item"]
			if !ok || v == nil {
				return nil
			}
			return NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer()
		}(),
	}
}

func (p ResetByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p ResetByStampSheetResult) Pointer() *ResetByStampSheetResult {
	return &p
}
