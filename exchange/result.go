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

import "github.com/gs2io/gs2-golang-sdk/core"

type DescribeNamespacesResult struct {
	Items         []Namespace `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
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
