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

package dictionary

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

type DescribeEntryModelsResult struct {
	Items []EntryModel `json:"items"`
}

type DescribeEntryModelsAsyncResult struct {
	result *DescribeEntryModelsResult
	err    error
}

func NewDescribeEntryModelsResultFromDict(data map[string]interface{}) DescribeEntryModelsResult {
	return DescribeEntryModelsResult{
		Items: CastEntryModels(core.CastArray(data["items"])),
	}
}

func (p DescribeEntryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeEntryModelsResult) Pointer() *DescribeEntryModelsResult {
	return &p
}

type GetEntryModelResult struct {
	Item *EntryModel `json:"item"`
}

type GetEntryModelAsyncResult struct {
	result *GetEntryModelResult
	err    error
}

func NewGetEntryModelResultFromDict(data map[string]interface{}) GetEntryModelResult {
	return GetEntryModelResult{
		Item: NewEntryModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEntryModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetEntryModelResult) Pointer() *GetEntryModelResult {
	return &p
}

type DescribeEntryModelMastersResult struct {
	Items         []EntryModelMaster `json:"items"`
	NextPageToken *string            `json:"nextPageToken"`
}

type DescribeEntryModelMastersAsyncResult struct {
	result *DescribeEntryModelMastersResult
	err    error
}

func NewDescribeEntryModelMastersResultFromDict(data map[string]interface{}) DescribeEntryModelMastersResult {
	return DescribeEntryModelMastersResult{
		Items:         CastEntryModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEntryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntryModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEntryModelMastersResult) Pointer() *DescribeEntryModelMastersResult {
	return &p
}

type CreateEntryModelMasterResult struct {
	Item *EntryModelMaster `json:"item"`
}

type CreateEntryModelMasterAsyncResult struct {
	result *CreateEntryModelMasterResult
	err    error
}

func NewCreateEntryModelMasterResultFromDict(data map[string]interface{}) CreateEntryModelMasterResult {
	return CreateEntryModelMasterResult{
		Item: NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateEntryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateEntryModelMasterResult) Pointer() *CreateEntryModelMasterResult {
	return &p
}

type GetEntryModelMasterResult struct {
	Item *EntryModelMaster `json:"item"`
}

type GetEntryModelMasterAsyncResult struct {
	result *GetEntryModelMasterResult
	err    error
}

func NewGetEntryModelMasterResultFromDict(data map[string]interface{}) GetEntryModelMasterResult {
	return GetEntryModelMasterResult{
		Item: NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEntryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetEntryModelMasterResult) Pointer() *GetEntryModelMasterResult {
	return &p
}

type UpdateEntryModelMasterResult struct {
	Item *EntryModelMaster `json:"item"`
}

type UpdateEntryModelMasterAsyncResult struct {
	result *UpdateEntryModelMasterResult
	err    error
}

func NewUpdateEntryModelMasterResultFromDict(data map[string]interface{}) UpdateEntryModelMasterResult {
	return UpdateEntryModelMasterResult{
		Item: NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateEntryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateEntryModelMasterResult) Pointer() *UpdateEntryModelMasterResult {
	return &p
}

type DeleteEntryModelMasterResult struct {
	Item *EntryModelMaster `json:"item"`
}

type DeleteEntryModelMasterAsyncResult struct {
	result *DeleteEntryModelMasterResult
	err    error
}

func NewDeleteEntryModelMasterResultFromDict(data map[string]interface{}) DeleteEntryModelMasterResult {
	return DeleteEntryModelMasterResult{
		Item: NewEntryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteEntryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteEntryModelMasterResult) Pointer() *DeleteEntryModelMasterResult {
	return &p
}

type DescribeEntriesResult struct {
	Items         []Entry `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeEntriesAsyncResult struct {
	result *DescribeEntriesResult
	err    error
}

func NewDescribeEntriesResultFromDict(data map[string]interface{}) DescribeEntriesResult {
	return DescribeEntriesResult{
		Items:         CastEntries(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEntriesResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEntriesResult) Pointer() *DescribeEntriesResult {
	return &p
}

type DescribeEntriesByUserIdResult struct {
	Items         []Entry `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeEntriesByUserIdAsyncResult struct {
	result *DescribeEntriesByUserIdResult
	err    error
}

func NewDescribeEntriesByUserIdResultFromDict(data map[string]interface{}) DescribeEntriesByUserIdResult {
	return DescribeEntriesByUserIdResult{
		Items:         CastEntries(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeEntriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeEntriesByUserIdResult) Pointer() *DescribeEntriesByUserIdResult {
	return &p
}

type AddEntriesByUserIdResult struct {
	Items []Entry `json:"items"`
}

type AddEntriesByUserIdAsyncResult struct {
	result *AddEntriesByUserIdResult
	err    error
}

func NewAddEntriesByUserIdResultFromDict(data map[string]interface{}) AddEntriesByUserIdResult {
	return AddEntriesByUserIdResult{
		Items: CastEntries(core.CastArray(data["items"])),
	}
}

func (p AddEntriesByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
			p.Items,
		),
	}
}

func (p AddEntriesByUserIdResult) Pointer() *AddEntriesByUserIdResult {
	return &p
}

type GetEntryResult struct {
	Item *Entry `json:"item"`
}

type GetEntryAsyncResult struct {
	result *GetEntryResult
	err    error
}

func NewGetEntryResultFromDict(data map[string]interface{}) GetEntryResult {
	return GetEntryResult{
		Item: NewEntryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEntryResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetEntryResult) Pointer() *GetEntryResult {
	return &p
}

type GetEntryByUserIdResult struct {
	Item *Entry `json:"item"`
}

type GetEntryByUserIdAsyncResult struct {
	result *GetEntryByUserIdResult
	err    error
}

func NewGetEntryByUserIdResultFromDict(data map[string]interface{}) GetEntryByUserIdResult {
	return GetEntryByUserIdResult{
		Item: NewEntryFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetEntryByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetEntryByUserIdResult) Pointer() *GetEntryByUserIdResult {
	return &p
}

type GetEntryWithSignatureResult struct {
	Item      *Entry  `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetEntryWithSignatureAsyncResult struct {
	result *GetEntryWithSignatureResult
	err    error
}

func NewGetEntryWithSignatureResultFromDict(data map[string]interface{}) GetEntryWithSignatureResult {
	return GetEntryWithSignatureResult{
		Item:      NewEntryFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetEntryWithSignatureResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetEntryWithSignatureResult) Pointer() *GetEntryWithSignatureResult {
	return &p
}

type GetEntryWithSignatureByUserIdResult struct {
	Item      *Entry  `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetEntryWithSignatureByUserIdAsyncResult struct {
	result *GetEntryWithSignatureByUserIdResult
	err    error
}

func NewGetEntryWithSignatureByUserIdResultFromDict(data map[string]interface{}) GetEntryWithSignatureByUserIdResult {
	return GetEntryWithSignatureByUserIdResult{
		Item:      NewEntryFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetEntryWithSignatureByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetEntryWithSignatureByUserIdResult) Pointer() *GetEntryWithSignatureByUserIdResult {
	return &p
}

type ResetByUserIdResult struct {
}

type ResetByUserIdAsyncResult struct {
	result *ResetByUserIdResult
	err    error
}

func NewResetByUserIdResultFromDict(data map[string]interface{}) ResetByUserIdResult {
	return ResetByUserIdResult{}
}

func (p ResetByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p ResetByUserIdResult) Pointer() *ResetByUserIdResult {
	return &p
}

type AddEntriesByStampSheetResult struct {
	Items []Entry `json:"items"`
}

type AddEntriesByStampSheetAsyncResult struct {
	result *AddEntriesByStampSheetResult
	err    error
}

func NewAddEntriesByStampSheetResultFromDict(data map[string]interface{}) AddEntriesByStampSheetResult {
	return AddEntriesByStampSheetResult{
		Items: CastEntries(core.CastArray(data["items"])),
	}
}

func (p AddEntriesByStampSheetResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastEntriesFromDict(
			p.Items,
		),
	}
}

func (p AddEntriesByStampSheetResult) Pointer() *AddEntriesByStampSheetResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentEntryMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentEntryMasterResult struct {
	Item *CurrentEntryMaster `json:"item"`
}

type GetCurrentEntryMasterAsyncResult struct {
	result *GetCurrentEntryMasterResult
	err    error
}

func NewGetCurrentEntryMasterResultFromDict(data map[string]interface{}) GetCurrentEntryMasterResult {
	return GetCurrentEntryMasterResult{
		Item: NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentEntryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentEntryMasterResult) Pointer() *GetCurrentEntryMasterResult {
	return &p
}

type UpdateCurrentEntryMasterResult struct {
	Item *CurrentEntryMaster `json:"item"`
}

type UpdateCurrentEntryMasterAsyncResult struct {
	result *UpdateCurrentEntryMasterResult
	err    error
}

func NewUpdateCurrentEntryMasterResultFromDict(data map[string]interface{}) UpdateCurrentEntryMasterResult {
	return UpdateCurrentEntryMasterResult{
		Item: NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentEntryMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentEntryMasterResult) Pointer() *UpdateCurrentEntryMasterResult {
	return &p
}

type UpdateCurrentEntryMasterFromGitHubResult struct {
	Item *CurrentEntryMaster `json:"item"`
}

type UpdateCurrentEntryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEntryMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentEntryMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentEntryMasterFromGitHubResult {
	return UpdateCurrentEntryMasterFromGitHubResult{
		Item: NewCurrentEntryMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentEntryMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentEntryMasterFromGitHubResult) Pointer() *UpdateCurrentEntryMasterFromGitHubResult {
	return &p
}
