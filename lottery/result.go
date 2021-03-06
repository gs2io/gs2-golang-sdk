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

import "core"

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
    return DescribeNamespacesResult {
        Items: CastNamespaces(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return CreateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetNamespaceStatusResult {
        Status: core.CastString(data["status"]),
    }
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return GetNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return UpdateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    return DeleteNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeLotteryModelMastersResult struct {
    Items []LotteryModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeLotteryModelMastersAsyncResult struct {
	result *DescribeLotteryModelMastersResult
	err    error
}

func NewDescribeLotteryModelMastersResultFromDict(data map[string]interface{}) DescribeLotteryModelMastersResult {
    return DescribeLotteryModelMastersResult {
        Items: CastLotteryModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeLotteryModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Item *LotteryModelMaster `json:"item"`
}

type CreateLotteryModelMasterAsyncResult struct {
	result *CreateLotteryModelMasterResult
	err    error
}

func NewCreateLotteryModelMasterResultFromDict(data map[string]interface{}) CreateLotteryModelMasterResult {
    return CreateLotteryModelMasterResult {
        Item: NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateLotteryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateLotteryModelMasterResult) Pointer() *CreateLotteryModelMasterResult {
    return &p
}

type GetLotteryModelMasterResult struct {
    Item *LotteryModelMaster `json:"item"`
}

type GetLotteryModelMasterAsyncResult struct {
	result *GetLotteryModelMasterResult
	err    error
}

func NewGetLotteryModelMasterResultFromDict(data map[string]interface{}) GetLotteryModelMasterResult {
    return GetLotteryModelMasterResult {
        Item: NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLotteryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLotteryModelMasterResult) Pointer() *GetLotteryModelMasterResult {
    return &p
}

type UpdateLotteryModelMasterResult struct {
    Item *LotteryModelMaster `json:"item"`
}

type UpdateLotteryModelMasterAsyncResult struct {
	result *UpdateLotteryModelMasterResult
	err    error
}

func NewUpdateLotteryModelMasterResultFromDict(data map[string]interface{}) UpdateLotteryModelMasterResult {
    return UpdateLotteryModelMasterResult {
        Item: NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateLotteryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateLotteryModelMasterResult) Pointer() *UpdateLotteryModelMasterResult {
    return &p
}

type DeleteLotteryModelMasterResult struct {
    Item *LotteryModelMaster `json:"item"`
}

type DeleteLotteryModelMasterAsyncResult struct {
	result *DeleteLotteryModelMasterResult
	err    error
}

func NewDeleteLotteryModelMasterResultFromDict(data map[string]interface{}) DeleteLotteryModelMasterResult {
    return DeleteLotteryModelMasterResult {
        Item: NewLotteryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteLotteryModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteLotteryModelMasterResult) Pointer() *DeleteLotteryModelMasterResult {
    return &p
}

type DescribePrizeTableMastersResult struct {
    Items []PrizeTableMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribePrizeTableMastersAsyncResult struct {
	result *DescribePrizeTableMastersResult
	err    error
}

func NewDescribePrizeTableMastersResultFromDict(data map[string]interface{}) DescribePrizeTableMastersResult {
    return DescribePrizeTableMastersResult {
        Items: CastPrizeTableMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribePrizeTableMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
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
    Item *PrizeTableMaster `json:"item"`
}

type CreatePrizeTableMasterAsyncResult struct {
	result *CreatePrizeTableMasterResult
	err    error
}

func NewCreatePrizeTableMasterResultFromDict(data map[string]interface{}) CreatePrizeTableMasterResult {
    return CreatePrizeTableMasterResult {
        Item: NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreatePrizeTableMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreatePrizeTableMasterResult) Pointer() *CreatePrizeTableMasterResult {
    return &p
}

type GetPrizeTableMasterResult struct {
    Item *PrizeTableMaster `json:"item"`
}

type GetPrizeTableMasterAsyncResult struct {
	result *GetPrizeTableMasterResult
	err    error
}

func NewGetPrizeTableMasterResultFromDict(data map[string]interface{}) GetPrizeTableMasterResult {
    return GetPrizeTableMasterResult {
        Item: NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetPrizeTableMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetPrizeTableMasterResult) Pointer() *GetPrizeTableMasterResult {
    return &p
}

type UpdatePrizeTableMasterResult struct {
    Item *PrizeTableMaster `json:"item"`
}

type UpdatePrizeTableMasterAsyncResult struct {
	result *UpdatePrizeTableMasterResult
	err    error
}

func NewUpdatePrizeTableMasterResultFromDict(data map[string]interface{}) UpdatePrizeTableMasterResult {
    return UpdatePrizeTableMasterResult {
        Item: NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdatePrizeTableMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdatePrizeTableMasterResult) Pointer() *UpdatePrizeTableMasterResult {
    return &p
}

type DeletePrizeTableMasterResult struct {
    Item *PrizeTableMaster `json:"item"`
}

type DeletePrizeTableMasterAsyncResult struct {
	result *DeletePrizeTableMasterResult
	err    error
}

func NewDeletePrizeTableMasterResultFromDict(data map[string]interface{}) DeletePrizeTableMasterResult {
    return DeletePrizeTableMasterResult {
        Item: NewPrizeTableMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeletePrizeTableMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeletePrizeTableMasterResult) Pointer() *DeletePrizeTableMasterResult {
    return &p
}

type DescribeBoxesResult struct {
    Items []Box `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBoxesAsyncResult struct {
	result *DescribeBoxesResult
	err    error
}

func NewDescribeBoxesResultFromDict(data map[string]interface{}) DescribeBoxesResult {
    return DescribeBoxesResult {
        Items: CastBoxes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBoxesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBoxesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBoxesResult) Pointer() *DescribeBoxesResult {
    return &p
}

type DescribeBoxesByUserIdResult struct {
    Items []Box `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBoxesByUserIdAsyncResult struct {
	result *DescribeBoxesByUserIdResult
	err    error
}

func NewDescribeBoxesByUserIdResultFromDict(data map[string]interface{}) DescribeBoxesByUserIdResult {
    return DescribeBoxesByUserIdResult {
        Items: CastBoxes(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBoxesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBoxesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBoxesByUserIdResult) Pointer() *DescribeBoxesByUserIdResult {
    return &p
}

type GetBoxResult struct {
    Item *BoxItems `json:"item"`
}

type GetBoxAsyncResult struct {
	result *GetBoxResult
	err    error
}

func NewGetBoxResultFromDict(data map[string]interface{}) GetBoxResult {
    return GetBoxResult {
        Item: NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBoxResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBoxResult) Pointer() *GetBoxResult {
    return &p
}

type GetBoxByUserIdResult struct {
    Item *BoxItems `json:"item"`
}

type GetBoxByUserIdAsyncResult struct {
	result *GetBoxByUserIdResult
	err    error
}

func NewGetBoxByUserIdResultFromDict(data map[string]interface{}) GetBoxByUserIdResult {
    return GetBoxByUserIdResult {
        Item: NewBoxItemsFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBoxByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBoxByUserIdResult) Pointer() *GetBoxByUserIdResult {
    return &p
}

type GetRawBoxByUserIdResult struct {
    Item *Box `json:"item"`
}

type GetRawBoxByUserIdAsyncResult struct {
	result *GetRawBoxByUserIdResult
	err    error
}

func NewGetRawBoxByUserIdResultFromDict(data map[string]interface{}) GetRawBoxByUserIdResult {
    return GetRawBoxByUserIdResult {
        Item: NewBoxFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRawBoxByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRawBoxByUserIdResult) Pointer() *GetRawBoxByUserIdResult {
    return &p
}

type ResetBoxResult struct {
}

type ResetBoxAsyncResult struct {
	result *ResetBoxResult
	err    error
}

func NewResetBoxResultFromDict(data map[string]interface{}) ResetBoxResult {
    return ResetBoxResult {
    }
}

func (p ResetBoxResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p ResetBoxResult) Pointer() *ResetBoxResult {
    return &p
}

type ResetBoxByUserIdResult struct {
}

type ResetBoxByUserIdAsyncResult struct {
	result *ResetBoxByUserIdResult
	err    error
}

func NewResetBoxByUserIdResultFromDict(data map[string]interface{}) ResetBoxByUserIdResult {
    return ResetBoxByUserIdResult {
    }
}

func (p ResetBoxByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p ResetBoxByUserIdResult) Pointer() *ResetBoxByUserIdResult {
    return &p
}

type DescribeLotteryModelsResult struct {
    Items []LotteryModel `json:"items"`
}

type DescribeLotteryModelsAsyncResult struct {
	result *DescribeLotteryModelsResult
	err    error
}

func NewDescribeLotteryModelsResultFromDict(data map[string]interface{}) DescribeLotteryModelsResult {
    return DescribeLotteryModelsResult {
        Items: CastLotteryModels(core.CastArray(data["items"])),
    }
}

func (p DescribeLotteryModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastLotteryModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeLotteryModelsResult) Pointer() *DescribeLotteryModelsResult {
    return &p
}

type GetLotteryModelResult struct {
    Item *LotteryModel `json:"item"`
}

type GetLotteryModelAsyncResult struct {
	result *GetLotteryModelResult
	err    error
}

func NewGetLotteryModelResultFromDict(data map[string]interface{}) GetLotteryModelResult {
    return GetLotteryModelResult {
        Item: NewLotteryModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetLotteryModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetLotteryModelResult) Pointer() *GetLotteryModelResult {
    return &p
}

type DescribePrizeTablesResult struct {
    Items []PrizeTable `json:"items"`
}

type DescribePrizeTablesAsyncResult struct {
	result *DescribePrizeTablesResult
	err    error
}

func NewDescribePrizeTablesResultFromDict(data map[string]interface{}) DescribePrizeTablesResult {
    return DescribePrizeTablesResult {
        Items: CastPrizeTables(core.CastArray(data["items"])),
    }
}

func (p DescribePrizeTablesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastPrizeTablesFromDict(
            p.Items,
        ),
    }
}

func (p DescribePrizeTablesResult) Pointer() *DescribePrizeTablesResult {
    return &p
}

type GetPrizeTableResult struct {
    Item *PrizeTable `json:"item"`
}

type GetPrizeTableAsyncResult struct {
	result *GetPrizeTableResult
	err    error
}

func NewGetPrizeTableResultFromDict(data map[string]interface{}) GetPrizeTableResult {
    return GetPrizeTableResult {
        Item: NewPrizeTableFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetPrizeTableResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetPrizeTableResult) Pointer() *GetPrizeTableResult {
    return &p
}

type DrawByUserIdResult struct {
    Items []DrawnPrize `json:"items"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    BoxItems *BoxItems `json:"boxItems"`
}

type DrawByUserIdAsyncResult struct {
	result *DrawByUserIdResult
	err    error
}

func NewDrawByUserIdResultFromDict(data map[string]interface{}) DrawByUserIdResult {
    return DrawByUserIdResult {
        Items: CastDrawnPrizes(core.CastArray(data["items"])),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        BoxItems: NewBoxItemsFromDict(core.CastMap(data["boxItems"])).Pointer(),
    }
}

func (p DrawByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastDrawnPrizesFromDict(
            p.Items,
        ),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "boxItems": p.BoxItems.ToDict(),
    }
}

func (p DrawByUserIdResult) Pointer() *DrawByUserIdResult {
    return &p
}

type DescribeProbabilitiesResult struct {
    Items []Probability `json:"items"`
}

type DescribeProbabilitiesAsyncResult struct {
	result *DescribeProbabilitiesResult
	err    error
}

func NewDescribeProbabilitiesResultFromDict(data map[string]interface{}) DescribeProbabilitiesResult {
    return DescribeProbabilitiesResult {
        Items: CastProbabilities(core.CastArray(data["items"])),
    }
}

func (p DescribeProbabilitiesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastProbabilitiesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeProbabilitiesResult) Pointer() *DescribeProbabilitiesResult {
    return &p
}

type DescribeProbabilitiesByUserIdResult struct {
    Items []Probability `json:"items"`
}

type DescribeProbabilitiesByUserIdAsyncResult struct {
	result *DescribeProbabilitiesByUserIdResult
	err    error
}

func NewDescribeProbabilitiesByUserIdResultFromDict(data map[string]interface{}) DescribeProbabilitiesByUserIdResult {
    return DescribeProbabilitiesByUserIdResult {
        Items: CastProbabilities(core.CastArray(data["items"])),
    }
}

func (p DescribeProbabilitiesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastProbabilitiesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeProbabilitiesByUserIdResult) Pointer() *DescribeProbabilitiesByUserIdResult {
    return &p
}

type DrawByStampSheetResult struct {
    Items []DrawnPrize `json:"items"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    BoxItems *BoxItems `json:"boxItems"`
}

type DrawByStampSheetAsyncResult struct {
	result *DrawByStampSheetResult
	err    error
}

func NewDrawByStampSheetResultFromDict(data map[string]interface{}) DrawByStampSheetResult {
    return DrawByStampSheetResult {
        Items: CastDrawnPrizes(core.CastArray(data["items"])),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        BoxItems: NewBoxItemsFromDict(core.CastMap(data["boxItems"])).Pointer(),
    }
}

func (p DrawByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastDrawnPrizesFromDict(
            p.Items,
        ),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "boxItems": p.BoxItems.ToDict(),
    }
}

func (p DrawByStampSheetResult) Pointer() *DrawByStampSheetResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentLotteryMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
    return &p
}

type GetCurrentLotteryMasterResult struct {
    Item *CurrentLotteryMaster `json:"item"`
}

type GetCurrentLotteryMasterAsyncResult struct {
	result *GetCurrentLotteryMasterResult
	err    error
}

func NewGetCurrentLotteryMasterResultFromDict(data map[string]interface{}) GetCurrentLotteryMasterResult {
    return GetCurrentLotteryMasterResult {
        Item: NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentLotteryMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentLotteryMasterResult) Pointer() *GetCurrentLotteryMasterResult {
    return &p
}

type UpdateCurrentLotteryMasterResult struct {
    Item *CurrentLotteryMaster `json:"item"`
}

type UpdateCurrentLotteryMasterAsyncResult struct {
	result *UpdateCurrentLotteryMasterResult
	err    error
}

func NewUpdateCurrentLotteryMasterResultFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterResult {
    return UpdateCurrentLotteryMasterResult {
        Item: NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentLotteryMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentLotteryMasterResult) Pointer() *UpdateCurrentLotteryMasterResult {
    return &p
}

type UpdateCurrentLotteryMasterFromGitHubResult struct {
    Item *CurrentLotteryMaster `json:"item"`
}

type UpdateCurrentLotteryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLotteryMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentLotteryMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentLotteryMasterFromGitHubResult {
    return UpdateCurrentLotteryMasterFromGitHubResult {
        Item: NewCurrentLotteryMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentLotteryMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentLotteryMasterFromGitHubResult) Pointer() *UpdateCurrentLotteryMasterFromGitHubResult {
    return &p
}