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

package money

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
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeWalletsResult struct {
    Items []Wallet `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeWalletsAsyncResult struct {
	result *DescribeWalletsResult
	err    error
}

func NewDescribeWalletsResultFromDict(data map[string]interface{}) DescribeWalletsResult {
    return DescribeWalletsResult {
        Items: CastWallets(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeWalletsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastWalletsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeWalletsResult) Pointer() *DescribeWalletsResult {
    return &p
}

type DescribeWalletsByUserIdResult struct {
    Items []Wallet `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeWalletsByUserIdAsyncResult struct {
	result *DescribeWalletsByUserIdResult
	err    error
}

func NewDescribeWalletsByUserIdResultFromDict(data map[string]interface{}) DescribeWalletsByUserIdResult {
    return DescribeWalletsByUserIdResult {
        Items: CastWallets(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeWalletsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastWalletsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeWalletsByUserIdResult) Pointer() *DescribeWalletsByUserIdResult {
    return &p
}

type GetWalletResult struct {
    Item *Wallet `json:"item"`
}

type GetWalletAsyncResult struct {
	result *GetWalletResult
	err    error
}

func NewGetWalletResultFromDict(data map[string]interface{}) GetWalletResult {
    return GetWalletResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetWalletResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetWalletResult) Pointer() *GetWalletResult {
    return &p
}

type GetWalletByUserIdResult struct {
    Item *Wallet `json:"item"`
}

type GetWalletByUserIdAsyncResult struct {
	result *GetWalletByUserIdResult
	err    error
}

func NewGetWalletByUserIdResultFromDict(data map[string]interface{}) GetWalletByUserIdResult {
    return GetWalletByUserIdResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetWalletByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetWalletByUserIdResult) Pointer() *GetWalletByUserIdResult {
    return &p
}

type DepositByUserIdResult struct {
    Item *Wallet `json:"item"`
}

type DepositByUserIdAsyncResult struct {
	result *DepositByUserIdResult
	err    error
}

func NewDepositByUserIdResultFromDict(data map[string]interface{}) DepositByUserIdResult {
    return DepositByUserIdResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DepositByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DepositByUserIdResult) Pointer() *DepositByUserIdResult {
    return &p
}

type WithdrawResult struct {
    Item *Wallet `json:"item"`
    Price *float32 `json:"price"`
}

type WithdrawAsyncResult struct {
	result *WithdrawResult
	err    error
}

func NewWithdrawResultFromDict(data map[string]interface{}) WithdrawResult {
    return WithdrawResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
        Price: core.CastFloat32(data["price"]),
    }
}

func (p WithdrawResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "price": p.Price,
    }
}

func (p WithdrawResult) Pointer() *WithdrawResult {
    return &p
}

type WithdrawByUserIdResult struct {
    Item *Wallet `json:"item"`
    Price *float32 `json:"price"`
}

type WithdrawByUserIdAsyncResult struct {
	result *WithdrawByUserIdResult
	err    error
}

func NewWithdrawByUserIdResultFromDict(data map[string]interface{}) WithdrawByUserIdResult {
    return WithdrawByUserIdResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
        Price: core.CastFloat32(data["price"]),
    }
}

func (p WithdrawByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "price": p.Price,
    }
}

func (p WithdrawByUserIdResult) Pointer() *WithdrawByUserIdResult {
    return &p
}

type DepositByStampSheetResult struct {
    Item *Wallet `json:"item"`
}

type DepositByStampSheetAsyncResult struct {
	result *DepositByStampSheetResult
	err    error
}

func NewDepositByStampSheetResultFromDict(data map[string]interface{}) DepositByStampSheetResult {
    return DepositByStampSheetResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DepositByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DepositByStampSheetResult) Pointer() *DepositByStampSheetResult {
    return &p
}

type WithdrawByStampTaskResult struct {
    Item *Wallet `json:"item"`
    Price *float32 `json:"price"`
    NewContextStack *string `json:"newContextStack"`
}

type WithdrawByStampTaskAsyncResult struct {
	result *WithdrawByStampTaskResult
	err    error
}

func NewWithdrawByStampTaskResultFromDict(data map[string]interface{}) WithdrawByStampTaskResult {
    return WithdrawByStampTaskResult {
        Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
        Price: core.CastFloat32(data["price"]),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p WithdrawByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "price": p.Price,
        "newContextStack": p.NewContextStack,
    }
}

func (p WithdrawByStampTaskResult) Pointer() *WithdrawByStampTaskResult {
    return &p
}

type DescribeReceiptsResult struct {
    Items []Receipt `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeReceiptsAsyncResult struct {
	result *DescribeReceiptsResult
	err    error
}

func NewDescribeReceiptsResultFromDict(data map[string]interface{}) DescribeReceiptsResult {
    return DescribeReceiptsResult {
        Items: CastReceipts(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeReceiptsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastReceiptsFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeReceiptsResult) Pointer() *DescribeReceiptsResult {
    return &p
}

type GetByUserIdAndTransactionIdResult struct {
    Item *Receipt `json:"item"`
}

type GetByUserIdAndTransactionIdAsyncResult struct {
	result *GetByUserIdAndTransactionIdResult
	err    error
}

func NewGetByUserIdAndTransactionIdResultFromDict(data map[string]interface{}) GetByUserIdAndTransactionIdResult {
    return GetByUserIdAndTransactionIdResult {
        Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetByUserIdAndTransactionIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetByUserIdAndTransactionIdResult) Pointer() *GetByUserIdAndTransactionIdResult {
    return &p
}

type RecordReceiptResult struct {
    Item *Receipt `json:"item"`
}

type RecordReceiptAsyncResult struct {
	result *RecordReceiptResult
	err    error
}

func NewRecordReceiptResultFromDict(data map[string]interface{}) RecordReceiptResult {
    return RecordReceiptResult {
        Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p RecordReceiptResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p RecordReceiptResult) Pointer() *RecordReceiptResult {
    return &p
}

type RecordReceiptByStampTaskResult struct {
    Item *Receipt `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type RecordReceiptByStampTaskAsyncResult struct {
	result *RecordReceiptByStampTaskResult
	err    error
}

func NewRecordReceiptByStampTaskResultFromDict(data map[string]interface{}) RecordReceiptByStampTaskResult {
    return RecordReceiptByStampTaskResult {
        Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p RecordReceiptByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p RecordReceiptByStampTaskResult) Pointer() *RecordReceiptByStampTaskResult {
    return &p
}