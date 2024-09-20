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

package money2

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
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
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
		Item:                 NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		WithdrawTransactions: CastDepositTransactions(core.CastArray(data["withdrawTransactions"])),
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
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
		),
	}
}

func (p WithdrawResult) Pointer() *WithdrawResult {
	return &p
}

type WithdrawByUserIdResult struct {
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
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
		Item:                 NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		WithdrawTransactions: CastDepositTransactions(core.CastArray(data["withdrawTransactions"])),
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
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
		),
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
	Item                 *Wallet              `json:"item"`
	WithdrawTransactions []DepositTransaction `json:"withdrawTransactions"`
	NewContextStack      *string              `json:"newContextStack"`
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
		Item:                 NewWalletFromDict(core.CastMap(data["item"])).Pointer(),
		WithdrawTransactions: CastDepositTransactions(core.CastArray(data["withdrawTransactions"])),
		NewContextStack:      core.CastString(data["newContextStack"]),
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
		"withdrawTransactions": CastDepositTransactionsFromDict(
			p.WithdrawTransactions,
		),
		"newContextStack": p.NewContextStack,
	}
}

func (p WithdrawByStampTaskResult) Pointer() *WithdrawByStampTaskResult {
	return &p
}

type DescribeEventsByUserIdResult struct {
	Items         []Event `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeEventsByUserIdAsyncResult struct {
	result *DescribeEventsByUserIdResult
	err    error
}

func NewDescribeEventsByUserIdResultFromJson(data string) DescribeEventsByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeEventsByUserIdResultFromDict(dict)
}

func NewDescribeEventsByUserIdResultFromDict(data map[string]interface{}) DescribeEventsByUserIdResult {
	return DescribeEventsByUserIdResult{
		Items:         CastEvents(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEventsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEventsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEventsByUserIdResult) Pointer() *DescribeEventsByUserIdResult {
	return &p
}

type GetEventByTransactionIdResult struct {
	Item *Event `json:"item"`
}

type GetEventByTransactionIdAsyncResult struct {
	result *GetEventByTransactionIdResult
	err    error
}

func NewGetEventByTransactionIdResultFromJson(data string) GetEventByTransactionIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetEventByTransactionIdResultFromDict(dict)
}

func NewGetEventByTransactionIdResultFromDict(data map[string]interface{}) GetEventByTransactionIdResult {
	return GetEventByTransactionIdResult{
		Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEventByTransactionIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetEventByTransactionIdResult) Pointer() *GetEventByTransactionIdResult {
	return &p
}

type VerifyReceiptResult struct {
	Item *Event `json:"item"`
}

type VerifyReceiptAsyncResult struct {
	result *VerifyReceiptResult
	err    error
}

func NewVerifyReceiptResultFromJson(data string) VerifyReceiptResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptResultFromDict(dict)
}

func NewVerifyReceiptResultFromDict(data map[string]interface{}) VerifyReceiptResult {
	return VerifyReceiptResult{
		Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VerifyReceiptResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyReceiptResult) Pointer() *VerifyReceiptResult {
	return &p
}

type VerifyReceiptByUserIdResult struct {
	Item *Event `json:"item"`
}

type VerifyReceiptByUserIdAsyncResult struct {
	result *VerifyReceiptByUserIdResult
	err    error
}

func NewVerifyReceiptByUserIdResultFromJson(data string) VerifyReceiptByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptByUserIdResultFromDict(dict)
}

func NewVerifyReceiptByUserIdResultFromDict(data map[string]interface{}) VerifyReceiptByUserIdResult {
	return VerifyReceiptByUserIdResult{
		Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VerifyReceiptByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p VerifyReceiptByUserIdResult) Pointer() *VerifyReceiptByUserIdResult {
	return &p
}

type VerifyReceiptByStampTaskResult struct {
	Item            *Event  `json:"item"`
	NewContextStack *string `json:"newContextStack"`
}

type VerifyReceiptByStampTaskAsyncResult struct {
	result *VerifyReceiptByStampTaskResult
	err    error
}

func NewVerifyReceiptByStampTaskResultFromJson(data string) VerifyReceiptByStampTaskResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVerifyReceiptByStampTaskResultFromDict(dict)
}

func NewVerifyReceiptByStampTaskResultFromDict(data map[string]interface{}) VerifyReceiptByStampTaskResult {
	return VerifyReceiptByStampTaskResult{
		Item:            NewEventFromDict(core.CastMap(data["item"])).Pointer(),
		NewContextStack: core.CastString(data["newContextStack"]),
	}
}

func (p VerifyReceiptByStampTaskResult) ToDict() map[string]interface{} {
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

func (p VerifyReceiptByStampTaskResult) Pointer() *VerifyReceiptByStampTaskResult {
	return &p
}

type DescribeStoreContentModelsResult struct {
	Items []StoreContentModel `json:"items"`
}

type DescribeStoreContentModelsAsyncResult struct {
	result *DescribeStoreContentModelsResult
	err    error
}

func NewDescribeStoreContentModelsResultFromJson(data string) DescribeStoreContentModelsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreContentModelsResultFromDict(dict)
}

func NewDescribeStoreContentModelsResultFromDict(data map[string]interface{}) DescribeStoreContentModelsResult {
	return DescribeStoreContentModelsResult{
		Items: CastStoreContentModels(core.CastArray(data["items"])),
	}
}

func (p DescribeStoreContentModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreContentModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeStoreContentModelsResult) Pointer() *DescribeStoreContentModelsResult {
	return &p
}

type GetStoreContentModelResult struct {
	Item *StoreContentModel `json:"item"`
}

type GetStoreContentModelAsyncResult struct {
	result *GetStoreContentModelResult
	err    error
}

func NewGetStoreContentModelResultFromJson(data string) GetStoreContentModelResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreContentModelResultFromDict(dict)
}

func NewGetStoreContentModelResultFromDict(data map[string]interface{}) GetStoreContentModelResult {
	return GetStoreContentModelResult{
		Item: NewStoreContentModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStoreContentModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStoreContentModelResult) Pointer() *GetStoreContentModelResult {
	return &p
}

type DescribeStoreContentModelMastersResult struct {
	Items         []StoreContentModelMaster `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribeStoreContentModelMastersAsyncResult struct {
	result *DescribeStoreContentModelMastersResult
	err    error
}

func NewDescribeStoreContentModelMastersResultFromJson(data string) DescribeStoreContentModelMastersResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeStoreContentModelMastersResultFromDict(dict)
}

func NewDescribeStoreContentModelMastersResultFromDict(data map[string]interface{}) DescribeStoreContentModelMastersResult {
	return DescribeStoreContentModelMastersResult{
		Items:         CastStoreContentModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeStoreContentModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastStoreContentModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeStoreContentModelMastersResult) Pointer() *DescribeStoreContentModelMastersResult {
	return &p
}

type CreateStoreContentModelMasterResult struct {
	Item *StoreContentModelMaster `json:"item"`
}

type CreateStoreContentModelMasterAsyncResult struct {
	result *CreateStoreContentModelMasterResult
	err    error
}

func NewCreateStoreContentModelMasterResultFromJson(data string) CreateStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateStoreContentModelMasterResultFromDict(dict)
}

func NewCreateStoreContentModelMasterResultFromDict(data map[string]interface{}) CreateStoreContentModelMasterResult {
	return CreateStoreContentModelMasterResult{
		Item: NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateStoreContentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p CreateStoreContentModelMasterResult) Pointer() *CreateStoreContentModelMasterResult {
	return &p
}

type GetStoreContentModelMasterResult struct {
	Item *StoreContentModelMaster `json:"item"`
}

type GetStoreContentModelMasterAsyncResult struct {
	result *GetStoreContentModelMasterResult
	err    error
}

func NewGetStoreContentModelMasterResultFromJson(data string) GetStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetStoreContentModelMasterResultFromDict(dict)
}

func NewGetStoreContentModelMasterResultFromDict(data map[string]interface{}) GetStoreContentModelMasterResult {
	return GetStoreContentModelMasterResult{
		Item: NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetStoreContentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetStoreContentModelMasterResult) Pointer() *GetStoreContentModelMasterResult {
	return &p
}

type UpdateStoreContentModelMasterResult struct {
	Item *StoreContentModelMaster `json:"item"`
}

type UpdateStoreContentModelMasterAsyncResult struct {
	result *UpdateStoreContentModelMasterResult
	err    error
}

func NewUpdateStoreContentModelMasterResultFromJson(data string) UpdateStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateStoreContentModelMasterResultFromDict(dict)
}

func NewUpdateStoreContentModelMasterResultFromDict(data map[string]interface{}) UpdateStoreContentModelMasterResult {
	return UpdateStoreContentModelMasterResult{
		Item: NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateStoreContentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateStoreContentModelMasterResult) Pointer() *UpdateStoreContentModelMasterResult {
	return &p
}

type DeleteStoreContentModelMasterResult struct {
	Item *StoreContentModelMaster `json:"item"`
}

type DeleteStoreContentModelMasterAsyncResult struct {
	result *DeleteStoreContentModelMasterResult
	err    error
}

func NewDeleteStoreContentModelMasterResultFromJson(data string) DeleteStoreContentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteStoreContentModelMasterResultFromDict(dict)
}

func NewDeleteStoreContentModelMasterResultFromDict(data map[string]interface{}) DeleteStoreContentModelMasterResult {
	return DeleteStoreContentModelMasterResult{
		Item: NewStoreContentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteStoreContentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p DeleteStoreContentModelMasterResult) Pointer() *DeleteStoreContentModelMasterResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentModelMaster `json:"item"`
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
		Item: NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentModelMasterResult struct {
	Item *CurrentModelMaster `json:"item"`
}

type GetCurrentModelMasterAsyncResult struct {
	result *GetCurrentModelMasterResult
	err    error
}

func NewGetCurrentModelMasterResultFromJson(data string) GetCurrentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentModelMasterResultFromDict(dict)
}

func NewGetCurrentModelMasterResultFromDict(data map[string]interface{}) GetCurrentModelMasterResult {
	return GetCurrentModelMasterResult{
		Item: NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetCurrentModelMasterResult) Pointer() *GetCurrentModelMasterResult {
	return &p
}

type UpdateCurrentModelMasterResult struct {
	Item *CurrentModelMaster `json:"item"`
}

type UpdateCurrentModelMasterAsyncResult struct {
	result *UpdateCurrentModelMasterResult
	err    error
}

func NewUpdateCurrentModelMasterResultFromJson(data string) UpdateCurrentModelMasterResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentModelMasterResultFromDict(dict)
}

func NewUpdateCurrentModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentModelMasterResult {
	return UpdateCurrentModelMasterResult{
		Item: NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentModelMasterResult) Pointer() *UpdateCurrentModelMasterResult {
	return &p
}

type UpdateCurrentModelMasterFromGitHubResult struct {
	Item *CurrentModelMaster `json:"item"`
}

type UpdateCurrentModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentModelMasterFromGitHubResultFromJson(data string) UpdateCurrentModelMasterFromGitHubResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentModelMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentModelMasterFromGitHubResult {
	return UpdateCurrentModelMasterFromGitHubResult{
		Item: NewCurrentModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentModelMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p UpdateCurrentModelMasterFromGitHubResult) Pointer() *UpdateCurrentModelMasterFromGitHubResult {
	return &p
}

type DescribeDailyTransactionHistoriesByCurrencyResult struct {
	Items         []DailyTransactionHistory `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribeDailyTransactionHistoriesByCurrencyAsyncResult struct {
	result *DescribeDailyTransactionHistoriesByCurrencyResult
	err    error
}

func NewDescribeDailyTransactionHistoriesByCurrencyResultFromJson(data string) DescribeDailyTransactionHistoriesByCurrencyResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDailyTransactionHistoriesByCurrencyResultFromDict(dict)
}

func NewDescribeDailyTransactionHistoriesByCurrencyResultFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesByCurrencyResult {
	return DescribeDailyTransactionHistoriesByCurrencyResult{
		Items:         CastDailyTransactionHistories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDailyTransactionHistoriesByCurrencyResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDailyTransactionHistoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDailyTransactionHistoriesByCurrencyResult) Pointer() *DescribeDailyTransactionHistoriesByCurrencyResult {
	return &p
}

type DescribeDailyTransactionHistoriesResult struct {
	Items         []DailyTransactionHistory `json:"items"`
	NextPageToken *string                   `json:"nextPageToken"`
}

type DescribeDailyTransactionHistoriesAsyncResult struct {
	result *DescribeDailyTransactionHistoriesResult
	err    error
}

func NewDescribeDailyTransactionHistoriesResultFromJson(data string) DescribeDailyTransactionHistoriesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDailyTransactionHistoriesResultFromDict(dict)
}

func NewDescribeDailyTransactionHistoriesResultFromDict(data map[string]interface{}) DescribeDailyTransactionHistoriesResult {
	return DescribeDailyTransactionHistoriesResult{
		Items:         CastDailyTransactionHistories(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeDailyTransactionHistoriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastDailyTransactionHistoriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeDailyTransactionHistoriesResult) Pointer() *DescribeDailyTransactionHistoriesResult {
	return &p
}

type GetDailyTransactionHistoryResult struct {
	Item *DailyTransactionHistory `json:"item"`
}

type GetDailyTransactionHistoryAsyncResult struct {
	result *GetDailyTransactionHistoryResult
	err    error
}

func NewGetDailyTransactionHistoryResultFromJson(data string) GetDailyTransactionHistoryResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDailyTransactionHistoryResultFromDict(dict)
}

func NewGetDailyTransactionHistoryResultFromDict(data map[string]interface{}) GetDailyTransactionHistoryResult {
	return GetDailyTransactionHistoryResult{
		Item: NewDailyTransactionHistoryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDailyTransactionHistoryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetDailyTransactionHistoryResult) Pointer() *GetDailyTransactionHistoryResult {
	return &p
}

type DescribeUnusedBalancesResult struct {
	Items         []UnusedBalance `json:"items"`
	NextPageToken *string         `json:"nextPageToken"`
}

type DescribeUnusedBalancesAsyncResult struct {
	result *DescribeUnusedBalancesResult
	err    error
}

func NewDescribeUnusedBalancesResultFromJson(data string) DescribeUnusedBalancesResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeUnusedBalancesResultFromDict(dict)
}

func NewDescribeUnusedBalancesResultFromDict(data map[string]interface{}) DescribeUnusedBalancesResult {
	return DescribeUnusedBalancesResult{
		Items:         CastUnusedBalances(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeUnusedBalancesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastUnusedBalancesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeUnusedBalancesResult) Pointer() *DescribeUnusedBalancesResult {
	return &p
}

type GetUnusedBalanceResult struct {
	Item *UnusedBalance `json:"item"`
}

type GetUnusedBalanceAsyncResult struct {
	result *GetUnusedBalanceResult
	err    error
}

func NewGetUnusedBalanceResultFromJson(data string) GetUnusedBalanceResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetUnusedBalanceResultFromDict(dict)
}

func NewGetUnusedBalanceResultFromDict(data map[string]interface{}) GetUnusedBalanceResult {
	return GetUnusedBalanceResult{
		Item: NewUnusedBalanceFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetUnusedBalanceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": func() map[string]interface{} {
			if p.Item == nil {
				return nil
			}
			return p.Item.ToDict()
		}(),
	}
}

func (p GetUnusedBalanceResult) Pointer() *GetUnusedBalanceResult {
	return &p
}
