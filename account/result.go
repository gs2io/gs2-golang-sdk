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
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
	return DeleteNamespaceResult{}
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
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

type GetAccountResult struct {
	Item *Account `json:"item"`
}

type GetAccountAsyncResult struct {
	result *GetAccountResult
	err    error
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
}

type DeleteAccountAsyncResult struct {
	result *DeleteAccountResult
	err    error
}

func NewDeleteAccountResultFromDict(data map[string]interface{}) DeleteAccountResult {
	return DeleteAccountResult{}
}

func (p DeleteAccountResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteAccountResult) Pointer() *DeleteAccountResult {
	return &p
}

type AuthenticationResult struct {
	Item      *Account `json:"item"`
	Body      *string  `json:"body"`
	Signature *string  `json:"signature"`
}

type AuthenticationAsyncResult struct {
	result *AuthenticationResult
	err    error
}

func NewAuthenticationResultFromDict(data map[string]interface{}) AuthenticationResult {
	return AuthenticationResult{
		Item:      NewAccountFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p AuthenticationResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
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
}

type DeleteTakeOverAsyncResult struct {
	result *DeleteTakeOverResult
	err    error
}

func NewDeleteTakeOverResultFromDict(data map[string]interface{}) DeleteTakeOverResult {
	return DeleteTakeOverResult{}
}

func (p DeleteTakeOverResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteTakeOverResult) Pointer() *DeleteTakeOverResult {
	return &p
}

type DeleteTakeOverByUserIdentifierResult struct {
}

type DeleteTakeOverByUserIdentifierAsyncResult struct {
	result *DeleteTakeOverByUserIdentifierResult
	err    error
}

func NewDeleteTakeOverByUserIdentifierResultFromDict(data map[string]interface{}) DeleteTakeOverByUserIdentifierResult {
	return DeleteTakeOverByUserIdentifierResult{}
}

func (p DeleteTakeOverByUserIdentifierResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DeleteTakeOverByUserIdentifierResult) Pointer() *DeleteTakeOverByUserIdentifierResult {
	return &p
}

type DoTakeOverResult struct {
	Item *Account `json:"item"`
}

type DoTakeOverAsyncResult struct {
	result *DoTakeOverResult
	err    error
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
