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
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
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
		Items:         CastNamespaces(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
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
	Item *Namespace `json:"item"`
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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
	return &p
}

type GetNamespaceStatusResult struct {
	Status *string `json:"status"`
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
		Status: core.CastString(data["status"]),
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
	Item *Namespace `json:"item"`
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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
	return &p
}

type UpdateNamespaceResult struct {
	Item *Namespace `json:"item"`
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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
	return &p
}

type DeleteNamespaceResult struct {
	Item *Namespace `json:"item"`
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
		Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
	return &p
}

type DescribeRateModelsResult struct {
	Items []RateModel `json:"items"`
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
		Items: CastRateModels(core.CastArray(data["items"])),
	}
}

func (p DescribeRateModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeRateModelsResult) Pointer() *DescribeRateModelsResult {
	return &p
}

type GetRateModelResult struct {
	Item *RateModel `json:"item"`
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
		Item: NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRateModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRateModelResult) Pointer() *GetRateModelResult {
	return &p
}

type DescribeRateModelMastersResult struct {
	Items         []RateModelMaster `json:"items"`
	NextPageToken *string           `json:"nextPageToken"`
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
		Items:         CastRateModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRateModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRateModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRateModelMastersResult) Pointer() *DescribeRateModelMastersResult {
	return &p
}

type CreateRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
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
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateRateModelMasterResult) Pointer() *CreateRateModelMasterResult {
	return &p
}

type GetRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
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
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRateModelMasterResult) Pointer() *GetRateModelMasterResult {
	return &p
}

type UpdateRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
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
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRateModelMasterResult) Pointer() *UpdateRateModelMasterResult {
	return &p
}

type DeleteRateModelMasterResult struct {
	Item *RateModelMaster `json:"item"`
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
		Item: NewRateModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRateModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRateModelMasterResult) Pointer() *DeleteRateModelMasterResult {
	return &p
}

type ExchangeResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p ExchangeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p ExchangeResult) Pointer() *ExchangeResult {
	return &p
}

type ExchangeByUserIdResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p ExchangeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p ExchangeByUserIdResult) Pointer() *ExchangeByUserIdResult {
	return &p
}

type ExchangeByStampSheetResult struct {
	Item                      *RateModel `json:"item"`
	StampSheet                *string    `json:"stampSheet"`
	StampSheetEncryptionKeyId *string    `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewRateModelFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p ExchangeByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p ExchangeByStampSheetResult) Pointer() *ExchangeByStampSheetResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
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
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
	return &p
}

type GetCurrentRateMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
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
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentRateMasterResult) Pointer() *GetCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterResult struct {
	Item *CurrentRateMaster `json:"item"`
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
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRateMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRateMasterResult) Pointer() *UpdateCurrentRateMasterResult {
	return &p
}

type UpdateCurrentRateMasterFromGitHubResult struct {
	Item *CurrentRateMaster `json:"item"`
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
		Item: NewCurrentRateMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRateMasterFromGitHubResult) Pointer() *UpdateCurrentRateMasterFromGitHubResult {
	return &p
}

type CreateAwaitByUserIdResult struct {
	Item     *Await `json:"item"`
	UnlockAt *int64 `json:"unlockAt"`
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
		Item:     NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		UnlockAt: core.CastInt64(data["unlockAt"]),
	}
}

func (p CreateAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":     p.Item.ToDict(),
		"unlockAt": p.UnlockAt,
	}
}

func (p CreateAwaitByUserIdResult) Pointer() *CreateAwaitByUserIdResult {
	return &p
}

type DescribeAwaitsResult struct {
	Items         []Await `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
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
		Items:         CastAwaits(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeAwaitsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAwaitsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeAwaitsResult) Pointer() *DescribeAwaitsResult {
	return &p
}

type DescribeAwaitsByUserIdResult struct {
	Items         []Await `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
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
		Items:         CastAwaits(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeAwaitsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAwaitsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeAwaitsByUserIdResult) Pointer() *DescribeAwaitsByUserIdResult {
	return &p
}

type GetAwaitResult struct {
	Item *Await `json:"item"`
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
		Item: NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetAwaitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetAwaitResult) Pointer() *GetAwaitResult {
	return &p
}

type GetAwaitByUserIdResult struct {
	Item *Await `json:"item"`
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
		Item: NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetAwaitByUserIdResult) Pointer() *GetAwaitByUserIdResult {
	return &p
}

type AcquireResult struct {
	Item                      *Await  `json:"item"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p AcquireResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p AcquireResult) Pointer() *AcquireResult {
	return &p
}

type AcquireByUserIdResult struct {
	Item                      *Await  `json:"item"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p AcquireByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p AcquireByUserIdResult) Pointer() *AcquireByUserIdResult {
	return &p
}

type AcquireForceByUserIdResult struct {
	Item                      *Await  `json:"item"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p AcquireForceByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p AcquireForceByUserIdResult) Pointer() *AcquireForceByUserIdResult {
	return &p
}

type SkipResult struct {
	Item                      *Await  `json:"item"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type SkipAsyncResult struct {
	result *SkipResult
	err    error
}

func NewSkipResultFromJson(data string) SkipResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSkipResultFromDict(dict)
}

func NewSkipResultFromDict(data map[string]interface{}) SkipResult {
	return SkipResult{
		Item:                      NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p SkipResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p SkipResult) Pointer() *SkipResult {
	return &p
}

type SkipByUserIdResult struct {
	Item                      *Await  `json:"item"`
	StampSheet                *string `json:"stampSheet"`
	StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
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
		Item:                      NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		StampSheet:                core.CastString(data["stampSheet"]),
		StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
	}
}

func (p SkipByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                      p.Item.ToDict(),
		"stampSheet":                p.StampSheet,
		"stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
	}
}

func (p SkipByUserIdResult) Pointer() *SkipByUserIdResult {
	return &p
}

type DeleteAwaitResult struct {
	Item *Await `json:"item"`
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
		Item: NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteAwaitResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteAwaitResult) Pointer() *DeleteAwaitResult {
	return &p
}

type DeleteAwaitByUserIdResult struct {
	Item *Await `json:"item"`
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
		Item: NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteAwaitByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteAwaitByUserIdResult) Pointer() *DeleteAwaitByUserIdResult {
	return &p
}

type CreateAwaitByStampSheetResult struct {
	Item     *Await `json:"item"`
	UnlockAt *int64 `json:"unlockAt"`
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
		Item:     NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		UnlockAt: core.CastInt64(data["unlockAt"]),
	}
}

func (p CreateAwaitByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":     p.Item.ToDict(),
		"unlockAt": p.UnlockAt,
	}
}

func (p CreateAwaitByStampSheetResult) Pointer() *CreateAwaitByStampSheetResult {
	return &p
}

type DeleteAwaitByStampTaskResult struct {
	Item            *Await  `json:"item"`
	NewContextStack *string `json:"newContextStack"`
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
		Item:            NewAwaitFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p DeleteAwaitByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":            p.Item.ToDict(),
		"newContextStack": p.NewContextStack,
	}
}

func (p DeleteAwaitByStampTaskResult) Pointer() *DeleteAwaitByStampTaskResult {
	return &p
}
