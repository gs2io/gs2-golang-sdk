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
	Url *string `json:"url"`
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
		Url: core.CastString(data["url"]),
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
	UploadToken *string `json:"uploadToken"`
	UploadUrl   *string `json:"uploadUrl"`
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
		UploadToken: core.CastString(data["uploadToken"]),
		UploadUrl:   core.CastString(data["uploadUrl"]),
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
	Url *string `json:"url"`
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
		Url: core.CastString(data["url"]),
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

type DescribeWalletsResult struct {
	Items         []Wallet `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeWalletsAsyncResult struct {
	result *DescribeWalletsResult
	err    error
}

func NewDescribeWalletsResultFromJson(data string) DescribeWalletsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsResultFromDict(dict)
}

func NewDescribeWalletsResultFromDict(data map[string]interface{}) DescribeWalletsResult {
	return DescribeWalletsResult{
		Items:         CastWallets(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeWalletsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	Items         []Wallet `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeWalletsByUserIdAsyncResult struct {
	result *DescribeWalletsByUserIdResult
	err    error
}

func NewDescribeWalletsByUserIdResultFromJson(data string) DescribeWalletsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeWalletsByUserIdResultFromDict(dict)
}

func NewDescribeWalletsByUserIdResultFromDict(data map[string]interface{}) DescribeWalletsByUserIdResult {
	return DescribeWalletsByUserIdResult{
		Items:         CastWallets(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeWalletsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetWalletResultFromJson(data string) GetWalletResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletResultFromDict(dict)
}

func NewGetWalletResultFromDict(data map[string]interface{}) GetWalletResult {
	return GetWalletResult{
		Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetWalletResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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

func NewGetWalletByUserIdResultFromJson(data string) GetWalletByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetWalletByUserIdResultFromDict(dict)
}

func NewGetWalletByUserIdResultFromDict(data map[string]interface{}) GetWalletByUserIdResult {
	return GetWalletByUserIdResult{
		Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetWalletByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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

func NewDepositByUserIdResultFromJson(data string) DepositByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByUserIdResultFromDict(dict)
}

func NewDepositByUserIdResultFromDict(data map[string]interface{}) DepositByUserIdResult {
	return DepositByUserIdResult{
		Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DepositByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DepositByUserIdResult) Pointer() *DepositByUserIdResult {
	return &p
}

type WithdrawResult struct {
	Item  *Wallet  `json:"item"`
	Price *float32 `json:"price"`
}

type WithdrawAsyncResult struct {
	result *WithdrawResult
	err    error
}

func NewWithdrawResultFromJson(data string) WithdrawResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawResultFromDict(dict)
}

func NewWithdrawResultFromDict(data map[string]interface{}) WithdrawResult {
	return WithdrawResult{
		Item:  NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		Price: core.CastFloat32(data["price"]),
	}
}

func (p WithdrawResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"price": p.Price,
	}
}

func (p WithdrawResult) Pointer() *WithdrawResult {
	return &p
}

type WithdrawByUserIdResult struct {
	Item  *Wallet  `json:"item"`
	Price *float32 `json:"price"`
}

type WithdrawByUserIdAsyncResult struct {
	result *WithdrawByUserIdResult
	err    error
}

func NewWithdrawByUserIdResultFromJson(data string) WithdrawByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByUserIdResultFromDict(dict)
}

func NewWithdrawByUserIdResultFromDict(data map[string]interface{}) WithdrawByUserIdResult {
	return WithdrawByUserIdResult{
		Item:  NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		Price: core.CastFloat32(data["price"]),
	}
}

func (p WithdrawByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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

func NewDepositByStampSheetResultFromJson(data string) DepositByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDepositByStampSheetResultFromDict(dict)
}

func NewDepositByStampSheetResultFromDict(data map[string]interface{}) DepositByStampSheetResult {
	return DepositByStampSheetResult{
		Item: NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DepositByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DepositByStampSheetResult) Pointer() *DepositByStampSheetResult {
	return &p
}

type WithdrawByStampTaskResult struct {
	Item            *Wallet  `json:"item"`
	Price           *float32 `json:"price"`
	NewContextStack *string  `json:"newContextStack"`
}

type WithdrawByStampTaskAsyncResult struct {
	result *WithdrawByStampTaskResult
	err    error
}

func NewWithdrawByStampTaskResultFromJson(data string) WithdrawByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawByStampTaskResultFromDict(dict)
}

func NewWithdrawByStampTaskResultFromDict(data map[string]interface{}) WithdrawByStampTaskResult {
	return WithdrawByStampTaskResult{
		Item:            NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		Price:           core.CastFloat32(data["price"]),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p WithdrawByStampTaskResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
		"price":           p.Price,
		"newContextStack": p.NewContextStack,
	}
}

func (p WithdrawByStampTaskResult) Pointer() *WithdrawByStampTaskResult {
	return &p
}

type DescribeReceiptsResult struct {
	Items         []Receipt `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeReceiptsAsyncResult struct {
	result *DescribeReceiptsResult
	err    error
}

func NewDescribeReceiptsResultFromJson(data string) DescribeReceiptsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiptsResultFromDict(dict)
}

func NewDescribeReceiptsResultFromDict(data map[string]interface{}) DescribeReceiptsResult {
	return DescribeReceiptsResult{
		Items:         CastReceipts(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeReceiptsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
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

func NewGetByUserIdAndTransactionIdResultFromJson(data string) GetByUserIdAndTransactionIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetByUserIdAndTransactionIdResultFromDict(dict)
}

func NewGetByUserIdAndTransactionIdResultFromDict(data map[string]interface{}) GetByUserIdAndTransactionIdResult {
	return GetByUserIdAndTransactionIdResult{
		Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetByUserIdAndTransactionIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
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

func NewRecordReceiptResultFromJson(data string) RecordReceiptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecordReceiptResultFromDict(dict)
}

func NewRecordReceiptResultFromDict(data map[string]interface{}) RecordReceiptResult {
	return RecordReceiptResult{
		Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RecordReceiptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p RecordReceiptResult) Pointer() *RecordReceiptResult {
	return &p
}

type RevertRecordReceiptResult struct {
	Item *Receipt `json:"item"`
}

type RevertRecordReceiptAsyncResult struct {
	result *RevertRecordReceiptResult
	err    error
}

func NewRevertRecordReceiptResultFromJson(data string) RevertRecordReceiptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertRecordReceiptResultFromDict(dict)
}

func NewRevertRecordReceiptResultFromDict(data map[string]interface{}) RevertRecordReceiptResult {
	return RevertRecordReceiptResult{
		Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RevertRecordReceiptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p RevertRecordReceiptResult) Pointer() *RevertRecordReceiptResult {
	return &p
}

type RecordReceiptByStampTaskResult struct {
	Item            *Receipt `json:"item"`
	NewContextStack *string  `json:"newContextStack"`
}

type RecordReceiptByStampTaskAsyncResult struct {
	result *RecordReceiptByStampTaskResult
	err    error
}

func NewRecordReceiptByStampTaskResultFromJson(data string) RecordReceiptByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRecordReceiptByStampTaskResultFromDict(dict)
}

func NewRecordReceiptByStampTaskResultFromDict(data map[string]interface{}) RecordReceiptByStampTaskResult {
	return RecordReceiptByStampTaskResult{
		Item:            NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p RecordReceiptByStampTaskResult) ToDict() map[string]interface{} {
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

func (p RecordReceiptByStampTaskResult) Pointer() *RecordReceiptByStampTaskResult {
	return &p
}

type RevertRecordReceiptByStampSheetResult struct {
	Item *Receipt `json:"item"`
}

type RevertRecordReceiptByStampSheetAsyncResult struct {
	result *RevertRecordReceiptByStampSheetResult
	err    error
}

func NewRevertRecordReceiptByStampSheetResultFromJson(data string) RevertRecordReceiptByStampSheetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertRecordReceiptByStampSheetResultFromDict(dict)
}

func NewRevertRecordReceiptByStampSheetResultFromDict(data map[string]interface{}) RevertRecordReceiptByStampSheetResult {
	return RevertRecordReceiptByStampSheetResult{
		Item: NewReceiptFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RevertRecordReceiptByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p RevertRecordReceiptByStampSheetResult) Pointer() *RevertRecordReceiptByStampSheetResult {
	return &p
}
