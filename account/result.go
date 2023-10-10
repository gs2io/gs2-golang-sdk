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

package account

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

type DescribeAccountsResult struct {
	Items         []Account `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeAccountsAsyncResult struct {
	result *DescribeAccountsResult
	err    error
}

func NewDescribeAccountsResultFromJson(data string) DescribeAccountsResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeAccountsResultFromDict(dict)
}

func NewDescribeAccountsResultFromDict(data map[string]interface{}) DescribeAccountsResult {
	return DescribeAccountsResult{
		Items:         CastAccounts(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeAccountsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastAccountsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeAccountsResult) Pointer() *DescribeAccountsResult {
	return &p
}

type CreateAccountResult struct {
	Item *Account `json:"item"`
}

type CreateAccountAsyncResult struct {
	result *CreateAccountResult
	err    error
}

func NewCreateAccountResultFromJson(data string) CreateAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateAccountResultFromDict(dict)
}

func NewCreateAccountResultFromDict(data map[string]interface{}) CreateAccountResult {
	return CreateAccountResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateAccountResult) Pointer() *CreateAccountResult {
	return &p
}

type UpdateTimeOffsetResult struct {
	Item *Account `json:"item"`
}

type UpdateTimeOffsetAsyncResult struct {
	result *UpdateTimeOffsetResult
	err    error
}

func NewUpdateTimeOffsetResultFromJson(data string) UpdateTimeOffsetResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTimeOffsetResultFromDict(dict)
}

func NewUpdateTimeOffsetResultFromDict(data map[string]interface{}) UpdateTimeOffsetResult {
	return UpdateTimeOffsetResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateTimeOffsetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateTimeOffsetResult) Pointer() *UpdateTimeOffsetResult {
	return &p
}

type UpdateBannedResult struct {
	Item *Account `json:"item"`
}

type UpdateBannedAsyncResult struct {
	result *UpdateBannedResult
	err    error
}

func NewUpdateBannedResultFromJson(data string) UpdateBannedResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateBannedResultFromDict(dict)
}

func NewUpdateBannedResultFromDict(data map[string]interface{}) UpdateBannedResult {
	return UpdateBannedResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateBannedResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateBannedResult) Pointer() *UpdateBannedResult {
	return &p
}

type AddBanResult struct {
	Item *Account `json:"item"`
}

type AddBanAsyncResult struct {
	result *AddBanResult
	err    error
}

func NewAddBanResultFromJson(data string) AddBanResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAddBanResultFromDict(dict)
}

func NewAddBanResultFromDict(data map[string]interface{}) AddBanResult {
	return AddBanResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p AddBanResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p AddBanResult) Pointer() *AddBanResult {
	return &p
}

type RemoveBanResult struct {
	Item *Account `json:"item"`
}

type RemoveBanAsyncResult struct {
	result *RemoveBanResult
	err    error
}

func NewRemoveBanResultFromJson(data string) RemoveBanResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRemoveBanResultFromDict(dict)
}

func NewRemoveBanResultFromDict(data map[string]interface{}) RemoveBanResult {
	return RemoveBanResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p RemoveBanResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p RemoveBanResult) Pointer() *RemoveBanResult {
	return &p
}

type GetAccountResult struct {
	Item *Account `json:"item"`
}

type GetAccountAsyncResult struct {
	result *GetAccountResult
	err    error
}

func NewGetAccountResultFromJson(data string) GetAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetAccountResultFromDict(dict)
}

func NewGetAccountResultFromDict(data map[string]interface{}) GetAccountResult {
	return GetAccountResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetAccountResult) Pointer() *GetAccountResult {
	return &p
}

type DeleteAccountResult struct {
	Item *Account `json:"item"`
}

type DeleteAccountAsyncResult struct {
	result *DeleteAccountResult
	err    error
}

func NewDeleteAccountResultFromJson(data string) DeleteAccountResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteAccountResultFromDict(dict)
}

func NewDeleteAccountResultFromDict(data map[string]interface{}) DeleteAccountResult {
	return DeleteAccountResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteAccountResult) Pointer() *DeleteAccountResult {
	return &p
}

type AuthenticationResult struct {
	Item        *Account    `json:"item"`
	BanStatuses []BanStatus `json:"banStatuses"`
	Body        *string     `json:"body"`
	Signature   *string     `json:"signature"`
}

type AuthenticationAsyncResult struct {
	result *AuthenticationResult
	err    error
}

func NewAuthenticationResultFromJson(data string) AuthenticationResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAuthenticationResultFromDict(dict)
}

func NewAuthenticationResultFromDict(data map[string]interface{}) AuthenticationResult {
	return AuthenticationResult{
		Item:        NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
		BanStatuses: CastBanStatuses(core.CastArray(data["banStatuses"])),
		Body:        core.CastString(data["body"]),
		Signature:   core.CastString(data["signature"]),
	}
}

func (p AuthenticationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
		"banStatuses": CastBanStatusesFromDict(
			p.BanStatuses,
		),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p AuthenticationResult) Pointer() *AuthenticationResult {
	return &p
}

type DescribeTakeOversResult struct {
	Items         []TakeOver `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeTakeOversAsyncResult struct {
	result *DescribeTakeOversResult
	err    error
}

func NewDescribeTakeOversResultFromJson(data string) DescribeTakeOversResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversResultFromDict(dict)
}

func NewDescribeTakeOversResultFromDict(data map[string]interface{}) DescribeTakeOversResult {
	return DescribeTakeOversResult{
		Items:         CastTakeOvers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeTakeOversResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOversFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeTakeOversResult) Pointer() *DescribeTakeOversResult {
	return &p
}

type DescribeTakeOversByUserIdResult struct {
	Items         []TakeOver `json:"items"`
	NextPageToken *string    `json:"nextPageToken"`
}

type DescribeTakeOversByUserIdAsyncResult struct {
	result *DescribeTakeOversByUserIdResult
	err    error
}

func NewDescribeTakeOversByUserIdResultFromJson(data string) DescribeTakeOversByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeTakeOversByUserIdResultFromDict(dict)
}

func NewDescribeTakeOversByUserIdResultFromDict(data map[string]interface{}) DescribeTakeOversByUserIdResult {
	return DescribeTakeOversByUserIdResult{
		Items:         CastTakeOvers(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeTakeOversByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastTakeOversFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeTakeOversByUserIdResult) Pointer() *DescribeTakeOversByUserIdResult {
	return &p
}

type CreateTakeOverResult struct {
	Item *TakeOver `json:"item"`
}

type CreateTakeOverAsyncResult struct {
	result *CreateTakeOverResult
	err    error
}

func NewCreateTakeOverResultFromJson(data string) CreateTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverResultFromDict(dict)
}

func NewCreateTakeOverResultFromDict(data map[string]interface{}) CreateTakeOverResult {
	return CreateTakeOverResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateTakeOverResult) Pointer() *CreateTakeOverResult {
	return &p
}

type CreateTakeOverByUserIdResult struct {
	Item *TakeOver `json:"item"`
}

type CreateTakeOverByUserIdAsyncResult struct {
	result *CreateTakeOverByUserIdResult
	err    error
}

func NewCreateTakeOverByUserIdResultFromJson(data string) CreateTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateTakeOverByUserIdResultFromDict(dict)
}

func NewCreateTakeOverByUserIdResultFromDict(data map[string]interface{}) CreateTakeOverByUserIdResult {
	return CreateTakeOverByUserIdResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateTakeOverByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateTakeOverByUserIdResult) Pointer() *CreateTakeOverByUserIdResult {
	return &p
}

type GetTakeOverResult struct {
	Item *TakeOver `json:"item"`
}

type GetTakeOverAsyncResult struct {
	result *GetTakeOverResult
	err    error
}

func NewGetTakeOverResultFromJson(data string) GetTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverResultFromDict(dict)
}

func NewGetTakeOverResultFromDict(data map[string]interface{}) GetTakeOverResult {
	return GetTakeOverResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetTakeOverResult) Pointer() *GetTakeOverResult {
	return &p
}

type GetTakeOverByUserIdResult struct {
	Item *TakeOver `json:"item"`
}

type GetTakeOverByUserIdAsyncResult struct {
	result *GetTakeOverByUserIdResult
	err    error
}

func NewGetTakeOverByUserIdResultFromJson(data string) GetTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetTakeOverByUserIdResultFromDict(dict)
}

func NewGetTakeOverByUserIdResultFromDict(data map[string]interface{}) GetTakeOverByUserIdResult {
	return GetTakeOverByUserIdResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetTakeOverByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetTakeOverByUserIdResult) Pointer() *GetTakeOverByUserIdResult {
	return &p
}

type UpdateTakeOverResult struct {
	Item *TakeOver `json:"item"`
}

type UpdateTakeOverAsyncResult struct {
	result *UpdateTakeOverResult
	err    error
}

func NewUpdateTakeOverResultFromJson(data string) UpdateTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverResultFromDict(dict)
}

func NewUpdateTakeOverResultFromDict(data map[string]interface{}) UpdateTakeOverResult {
	return UpdateTakeOverResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateTakeOverResult) Pointer() *UpdateTakeOverResult {
	return &p
}

type UpdateTakeOverByUserIdResult struct {
	Item *TakeOver `json:"item"`
}

type UpdateTakeOverByUserIdAsyncResult struct {
	result *UpdateTakeOverByUserIdResult
	err    error
}

func NewUpdateTakeOverByUserIdResultFromJson(data string) UpdateTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateTakeOverByUserIdResultFromDict(dict)
}

func NewUpdateTakeOverByUserIdResultFromDict(data map[string]interface{}) UpdateTakeOverByUserIdResult {
	return UpdateTakeOverByUserIdResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateTakeOverByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateTakeOverByUserIdResult) Pointer() *UpdateTakeOverByUserIdResult {
	return &p
}

type DeleteTakeOverResult struct {
	Item *TakeOver `json:"item"`
}

type DeleteTakeOverAsyncResult struct {
	result *DeleteTakeOverResult
	err    error
}

func NewDeleteTakeOverResultFromJson(data string) DeleteTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverResultFromDict(dict)
}

func NewDeleteTakeOverResultFromDict(data map[string]interface{}) DeleteTakeOverResult {
	return DeleteTakeOverResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteTakeOverResult) Pointer() *DeleteTakeOverResult {
	return &p
}

type DeleteTakeOverByUserIdentifierResult struct {
	Item *TakeOver `json:"item"`
}

type DeleteTakeOverByUserIdentifierAsyncResult struct {
	result *DeleteTakeOverByUserIdentifierResult
	err    error
}

func NewDeleteTakeOverByUserIdentifierResultFromJson(data string) DeleteTakeOverByUserIdentifierResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdentifierResultFromDict(dict)
}

func NewDeleteTakeOverByUserIdentifierResultFromDict(data map[string]interface{}) DeleteTakeOverByUserIdentifierResult {
	return DeleteTakeOverByUserIdentifierResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteTakeOverByUserIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteTakeOverByUserIdentifierResult) Pointer() *DeleteTakeOverByUserIdentifierResult {
	return &p
}

type DeleteTakeOverByUserIdResult struct {
	Item *TakeOver `json:"item"`
}

type DeleteTakeOverByUserIdAsyncResult struct {
	result *DeleteTakeOverByUserIdResult
	err    error
}

func NewDeleteTakeOverByUserIdResultFromJson(data string) DeleteTakeOverByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteTakeOverByUserIdResultFromDict(dict)
}

func NewDeleteTakeOverByUserIdResultFromDict(data map[string]interface{}) DeleteTakeOverByUserIdResult {
	return DeleteTakeOverByUserIdResult{
		Item: NewTakeOverFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteTakeOverByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteTakeOverByUserIdResult) Pointer() *DeleteTakeOverByUserIdResult {
	return &p
}

type DoTakeOverResult struct {
	Item *Account `json:"item"`
}

type DoTakeOverAsyncResult struct {
	result *DoTakeOverResult
	err    error
}

func NewDoTakeOverResultFromJson(data string) DoTakeOverResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoTakeOverResultFromDict(dict)
}

func NewDoTakeOverResultFromDict(data map[string]interface{}) DoTakeOverResult {
	return DoTakeOverResult{
		Item: NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DoTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DoTakeOverResult) Pointer() *DoTakeOverResult {
	return &p
}

type GetDataOwnerByUserIdResult struct {
	Item *DataOwner `json:"item"`
}

type GetDataOwnerByUserIdAsyncResult struct {
	result *GetDataOwnerByUserIdResult
	err    error
}

func NewGetDataOwnerByUserIdResultFromJson(data string) GetDataOwnerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDataOwnerByUserIdResultFromDict(dict)
}

func NewGetDataOwnerByUserIdResultFromDict(data map[string]interface{}) GetDataOwnerByUserIdResult {
	return GetDataOwnerByUserIdResult{
		Item: NewDataOwnerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetDataOwnerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetDataOwnerByUserIdResult) Pointer() *GetDataOwnerByUserIdResult {
	return &p
}

type DeleteDataOwnerByUserIdResult struct {
	Item *DataOwner `json:"item"`
}

type DeleteDataOwnerByUserIdAsyncResult struct {
	result *DeleteDataOwnerByUserIdResult
	err    error
}

func NewDeleteDataOwnerByUserIdResultFromJson(data string) DeleteDataOwnerByUserIdResult {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDataOwnerByUserIdResultFromDict(dict)
}

func NewDeleteDataOwnerByUserIdResultFromDict(data map[string]interface{}) DeleteDataOwnerByUserIdResult {
	return DeleteDataOwnerByUserIdResult{
		Item: NewDataOwnerFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteDataOwnerByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteDataOwnerByUserIdResult) Pointer() *DeleteDataOwnerByUserIdResult {
	return &p
}
