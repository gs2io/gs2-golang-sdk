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

type DescribeSalesItemMastersResult struct {
    Items []SalesItemMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSalesItemMastersAsyncResult struct {
	result *DescribeSalesItemMastersResult
	err    error
}

func NewDescribeSalesItemMastersResultFromDict(data map[string]interface{}) DescribeSalesItemMastersResult {
    return DescribeSalesItemMastersResult {
        Items: CastSalesItemMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSalesItemMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSalesItemMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSalesItemMastersResult) Pointer() *DescribeSalesItemMastersResult {
    return &p
}

type CreateSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type CreateSalesItemMasterAsyncResult struct {
	result *CreateSalesItemMasterResult
	err    error
}

func NewCreateSalesItemMasterResultFromDict(data map[string]interface{}) CreateSalesItemMasterResult {
    return CreateSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSalesItemMasterResult) Pointer() *CreateSalesItemMasterResult {
    return &p
}

type GetSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type GetSalesItemMasterAsyncResult struct {
	result *GetSalesItemMasterResult
	err    error
}

func NewGetSalesItemMasterResultFromDict(data map[string]interface{}) GetSalesItemMasterResult {
    return GetSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSalesItemMasterResult) Pointer() *GetSalesItemMasterResult {
    return &p
}

type UpdateSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type UpdateSalesItemMasterAsyncResult struct {
	result *UpdateSalesItemMasterResult
	err    error
}

func NewUpdateSalesItemMasterResultFromDict(data map[string]interface{}) UpdateSalesItemMasterResult {
    return UpdateSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSalesItemMasterResult) Pointer() *UpdateSalesItemMasterResult {
    return &p
}

type DeleteSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type DeleteSalesItemMasterAsyncResult struct {
	result *DeleteSalesItemMasterResult
	err    error
}

func NewDeleteSalesItemMasterResultFromDict(data map[string]interface{}) DeleteSalesItemMasterResult {
    return DeleteSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSalesItemMasterResult) Pointer() *DeleteSalesItemMasterResult {
    return &p
}

type DescribeSalesItemGroupMastersResult struct {
    Items []SalesItemGroupMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSalesItemGroupMastersAsyncResult struct {
	result *DescribeSalesItemGroupMastersResult
	err    error
}

func NewDescribeSalesItemGroupMastersResultFromDict(data map[string]interface{}) DescribeSalesItemGroupMastersResult {
    return DescribeSalesItemGroupMastersResult {
        Items: CastSalesItemGroupMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSalesItemGroupMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSalesItemGroupMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSalesItemGroupMastersResult) Pointer() *DescribeSalesItemGroupMastersResult {
    return &p
}

type CreateSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type CreateSalesItemGroupMasterAsyncResult struct {
	result *CreateSalesItemGroupMasterResult
	err    error
}

func NewCreateSalesItemGroupMasterResultFromDict(data map[string]interface{}) CreateSalesItemGroupMasterResult {
    return CreateSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSalesItemGroupMasterResult) Pointer() *CreateSalesItemGroupMasterResult {
    return &p
}

type GetSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type GetSalesItemGroupMasterAsyncResult struct {
	result *GetSalesItemGroupMasterResult
	err    error
}

func NewGetSalesItemGroupMasterResultFromDict(data map[string]interface{}) GetSalesItemGroupMasterResult {
    return GetSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSalesItemGroupMasterResult) Pointer() *GetSalesItemGroupMasterResult {
    return &p
}

type UpdateSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type UpdateSalesItemGroupMasterAsyncResult struct {
	result *UpdateSalesItemGroupMasterResult
	err    error
}

func NewUpdateSalesItemGroupMasterResultFromDict(data map[string]interface{}) UpdateSalesItemGroupMasterResult {
    return UpdateSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSalesItemGroupMasterResult) Pointer() *UpdateSalesItemGroupMasterResult {
    return &p
}

type DeleteSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type DeleteSalesItemGroupMasterAsyncResult struct {
	result *DeleteSalesItemGroupMasterResult
	err    error
}

func NewDeleteSalesItemGroupMasterResultFromDict(data map[string]interface{}) DeleteSalesItemGroupMasterResult {
    return DeleteSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSalesItemGroupMasterResult) Pointer() *DeleteSalesItemGroupMasterResult {
    return &p
}

type DescribeShowcaseMastersResult struct {
    Items []ShowcaseMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeShowcaseMastersAsyncResult struct {
	result *DescribeShowcaseMastersResult
	err    error
}

func NewDescribeShowcaseMastersResultFromDict(data map[string]interface{}) DescribeShowcaseMastersResult {
    return DescribeShowcaseMastersResult {
        Items: CastShowcaseMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeShowcaseMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcaseMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeShowcaseMastersResult) Pointer() *DescribeShowcaseMastersResult {
    return &p
}

type CreateShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type CreateShowcaseMasterAsyncResult struct {
	result *CreateShowcaseMasterResult
	err    error
}

func NewCreateShowcaseMasterResultFromDict(data map[string]interface{}) CreateShowcaseMasterResult {
    return CreateShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateShowcaseMasterResult) Pointer() *CreateShowcaseMasterResult {
    return &p
}

type GetShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type GetShowcaseMasterAsyncResult struct {
	result *GetShowcaseMasterResult
	err    error
}

func NewGetShowcaseMasterResultFromDict(data map[string]interface{}) GetShowcaseMasterResult {
    return GetShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseMasterResult) Pointer() *GetShowcaseMasterResult {
    return &p
}

type UpdateShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type UpdateShowcaseMasterAsyncResult struct {
	result *UpdateShowcaseMasterResult
	err    error
}

func NewUpdateShowcaseMasterResultFromDict(data map[string]interface{}) UpdateShowcaseMasterResult {
    return UpdateShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateShowcaseMasterResult) Pointer() *UpdateShowcaseMasterResult {
    return &p
}

type DeleteShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type DeleteShowcaseMasterAsyncResult struct {
	result *DeleteShowcaseMasterResult
	err    error
}

func NewDeleteShowcaseMasterResultFromDict(data map[string]interface{}) DeleteShowcaseMasterResult {
    return DeleteShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteShowcaseMasterResult) Pointer() *DeleteShowcaseMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentShowcaseMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type GetCurrentShowcaseMasterAsyncResult struct {
	result *GetCurrentShowcaseMasterResult
	err    error
}

func NewGetCurrentShowcaseMasterResultFromDict(data map[string]interface{}) GetCurrentShowcaseMasterResult {
    return GetCurrentShowcaseMasterResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentShowcaseMasterResult) Pointer() *GetCurrentShowcaseMasterResult {
    return &p
}

type UpdateCurrentShowcaseMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type UpdateCurrentShowcaseMasterAsyncResult struct {
	result *UpdateCurrentShowcaseMasterResult
	err    error
}

func NewUpdateCurrentShowcaseMasterResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterResult {
    return UpdateCurrentShowcaseMasterResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentShowcaseMasterResult) Pointer() *UpdateCurrentShowcaseMasterResult {
    return &p
}

type UpdateCurrentShowcaseMasterFromGitHubResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type UpdateCurrentShowcaseMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentShowcaseMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentShowcaseMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterFromGitHubResult {
    return UpdateCurrentShowcaseMasterFromGitHubResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) Pointer() *UpdateCurrentShowcaseMasterFromGitHubResult {
    return &p
}

type DescribeShowcasesResult struct {
    Items []Showcase `json:"items"`
}

type DescribeShowcasesAsyncResult struct {
	result *DescribeShowcasesResult
	err    error
}

func NewDescribeShowcasesResultFromDict(data map[string]interface{}) DescribeShowcasesResult {
    return DescribeShowcasesResult {
        Items: CastShowcases(core.CastArray(data["items"])),
    }
}

func (p DescribeShowcasesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcasesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeShowcasesResult) Pointer() *DescribeShowcasesResult {
    return &p
}

type DescribeShowcasesByUserIdResult struct {
    Items []Showcase `json:"items"`
}

type DescribeShowcasesByUserIdAsyncResult struct {
	result *DescribeShowcasesByUserIdResult
	err    error
}

func NewDescribeShowcasesByUserIdResultFromDict(data map[string]interface{}) DescribeShowcasesByUserIdResult {
    return DescribeShowcasesByUserIdResult {
        Items: CastShowcases(core.CastArray(data["items"])),
    }
}

func (p DescribeShowcasesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcasesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeShowcasesByUserIdResult) Pointer() *DescribeShowcasesByUserIdResult {
    return &p
}

type GetShowcaseResult struct {
    Item *Showcase `json:"item"`
}

type GetShowcaseAsyncResult struct {
	result *GetShowcaseResult
	err    error
}

func NewGetShowcaseResultFromDict(data map[string]interface{}) GetShowcaseResult {
    return GetShowcaseResult {
        Item: NewShowcaseFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseResult) Pointer() *GetShowcaseResult {
    return &p
}

type GetShowcaseByUserIdResult struct {
    Item *Showcase `json:"item"`
}

type GetShowcaseByUserIdAsyncResult struct {
	result *GetShowcaseByUserIdResult
	err    error
}

func NewGetShowcaseByUserIdResultFromDict(data map[string]interface{}) GetShowcaseByUserIdResult {
    return GetShowcaseByUserIdResult {
        Item: NewShowcaseFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseByUserIdResult) Pointer() *GetShowcaseByUserIdResult {
    return &p
}

type BuyResult struct {
    Item *SalesItem `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type BuyAsyncResult struct {
	result *BuyResult
	err    error
}

func NewBuyResultFromDict(data map[string]interface{}) BuyResult {
    return BuyResult {
        Item: NewSalesItemFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p BuyResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p BuyResult) Pointer() *BuyResult {
    return &p
}

type BuyByUserIdResult struct {
    Item *SalesItem `json:"item"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
}

type BuyByUserIdAsyncResult struct {
	result *BuyByUserIdResult
	err    error
}

func NewBuyByUserIdResultFromDict(data map[string]interface{}) BuyByUserIdResult {
    return BuyByUserIdResult {
        Item: NewSalesItemFromDict(core.CastMap(data["item"])).Pointer(),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
    }
}

func (p BuyByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
    }
}

func (p BuyByUserIdResult) Pointer() *BuyByUserIdResult {
    return &p
}