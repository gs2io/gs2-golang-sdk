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

package ranking

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

type DescribeCategoryModelsResult struct {
	Items []CategoryModel `json:"items"`
}

type DescribeCategoryModelsAsyncResult struct {
	result *DescribeCategoryModelsResult
	err    error
}

func NewDescribeCategoryModelsResultFromDict(data map[string]interface{}) DescribeCategoryModelsResult {
	return DescribeCategoryModelsResult{
		Items: CastCategoryModels(core.CastArray(data["items"])),
	}
}

func (p DescribeCategoryModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCategoryModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeCategoryModelsResult) Pointer() *DescribeCategoryModelsResult {
	return &p
}

type GetCategoryModelResult struct {
	Item *CategoryModel `json:"item"`
}

type GetCategoryModelAsyncResult struct {
	result *GetCategoryModelResult
	err    error
}

func NewGetCategoryModelResultFromDict(data map[string]interface{}) GetCategoryModelResult {
	return GetCategoryModelResult{
		Item: NewCategoryModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCategoryModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCategoryModelResult) Pointer() *GetCategoryModelResult {
	return &p
}

type DescribeCategoryModelMastersResult struct {
	Items         []CategoryModelMaster `json:"items"`
	NextPageToken *string               `json:"nextPageToken"`
}

type DescribeCategoryModelMastersAsyncResult struct {
	result *DescribeCategoryModelMastersResult
	err    error
}

func NewDescribeCategoryModelMastersResultFromDict(data map[string]interface{}) DescribeCategoryModelMastersResult {
	return DescribeCategoryModelMastersResult{
		Items:         CastCategoryModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeCategoryModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastCategoryModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeCategoryModelMastersResult) Pointer() *DescribeCategoryModelMastersResult {
	return &p
}

type CreateCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type CreateCategoryModelMasterAsyncResult struct {
	result *CreateCategoryModelMasterResult
	err    error
}

func NewCreateCategoryModelMasterResultFromDict(data map[string]interface{}) CreateCategoryModelMasterResult {
	return CreateCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateCategoryModelMasterResult) Pointer() *CreateCategoryModelMasterResult {
	return &p
}

type GetCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type GetCategoryModelMasterAsyncResult struct {
	result *GetCategoryModelMasterResult
	err    error
}

func NewGetCategoryModelMasterResultFromDict(data map[string]interface{}) GetCategoryModelMasterResult {
	return GetCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCategoryModelMasterResult) Pointer() *GetCategoryModelMasterResult {
	return &p
}

type UpdateCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type UpdateCategoryModelMasterAsyncResult struct {
	result *UpdateCategoryModelMasterResult
	err    error
}

func NewUpdateCategoryModelMasterResultFromDict(data map[string]interface{}) UpdateCategoryModelMasterResult {
	return UpdateCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCategoryModelMasterResult) Pointer() *UpdateCategoryModelMasterResult {
	return &p
}

type DeleteCategoryModelMasterResult struct {
	Item *CategoryModelMaster `json:"item"`
}

type DeleteCategoryModelMasterAsyncResult struct {
	result *DeleteCategoryModelMasterResult
	err    error
}

func NewDeleteCategoryModelMasterResultFromDict(data map[string]interface{}) DeleteCategoryModelMasterResult {
	return DeleteCategoryModelMasterResult{
		Item: NewCategoryModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteCategoryModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteCategoryModelMasterResult) Pointer() *DeleteCategoryModelMasterResult {
	return &p
}

type DescribeSubscribesByCategoryNameResult struct {
	Items []SubscribeUser `json:"items"`
}

type DescribeSubscribesByCategoryNameAsyncResult struct {
	result *DescribeSubscribesByCategoryNameResult
	err    error
}

func NewDescribeSubscribesByCategoryNameResultFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameResult {
	return DescribeSubscribesByCategoryNameResult{
		Items: CastSubscribeUsers(core.CastArray(data["items"])),
	}
}

func (p DescribeSubscribesByCategoryNameResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
	}
}

func (p DescribeSubscribesByCategoryNameResult) Pointer() *DescribeSubscribesByCategoryNameResult {
	return &p
}

type DescribeSubscribesByCategoryNameAndUserIdResult struct {
	Items []SubscribeUser `json:"items"`
}

type DescribeSubscribesByCategoryNameAndUserIdAsyncResult struct {
	result *DescribeSubscribesByCategoryNameAndUserIdResult
	err    error
}

func NewDescribeSubscribesByCategoryNameAndUserIdResultFromDict(data map[string]interface{}) DescribeSubscribesByCategoryNameAndUserIdResult {
	return DescribeSubscribesByCategoryNameAndUserIdResult{
		Items: CastSubscribeUsers(core.CastArray(data["items"])),
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastSubscribeUsersFromDict(
			p.Items,
		),
	}
}

func (p DescribeSubscribesByCategoryNameAndUserIdResult) Pointer() *DescribeSubscribesByCategoryNameAndUserIdResult {
	return &p
}

type SubscribeResult struct {
	Item *SubscribeUser `json:"item"`
}

type SubscribeAsyncResult struct {
	result *SubscribeResult
	err    error
}

func NewSubscribeResultFromDict(data map[string]interface{}) SubscribeResult {
	return SubscribeResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p SubscribeResult) Pointer() *SubscribeResult {
	return &p
}

type SubscribeByUserIdResult struct {
	Item *SubscribeUser `json:"item"`
}

type SubscribeByUserIdAsyncResult struct {
	result *SubscribeByUserIdResult
	err    error
}

func NewSubscribeByUserIdResultFromDict(data map[string]interface{}) SubscribeByUserIdResult {
	return SubscribeByUserIdResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p SubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p SubscribeByUserIdResult) Pointer() *SubscribeByUserIdResult {
	return &p
}

type GetSubscribeResult struct {
	Item *SubscribeUser `json:"item"`
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

func NewGetSubscribeResultFromDict(data map[string]interface{}) GetSubscribeResult {
	return GetSubscribeResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetSubscribeResult) Pointer() *GetSubscribeResult {
	return &p
}

type GetSubscribeByUserIdResult struct {
	Item *SubscribeUser `json:"item"`
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

func NewGetSubscribeByUserIdResultFromDict(data map[string]interface{}) GetSubscribeByUserIdResult {
	return GetSubscribeByUserIdResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetSubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetSubscribeByUserIdResult) Pointer() *GetSubscribeByUserIdResult {
	return &p
}

type UnsubscribeResult struct {
	Item *SubscribeUser `json:"item"`
}

type UnsubscribeAsyncResult struct {
	result *UnsubscribeResult
	err    error
}

func NewUnsubscribeResultFromDict(data map[string]interface{}) UnsubscribeResult {
	return UnsubscribeResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UnsubscribeResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UnsubscribeResult) Pointer() *UnsubscribeResult {
	return &p
}

type UnsubscribeByUserIdResult struct {
	Item *SubscribeUser `json:"item"`
}

type UnsubscribeByUserIdAsyncResult struct {
	result *UnsubscribeByUserIdResult
	err    error
}

func NewUnsubscribeByUserIdResultFromDict(data map[string]interface{}) UnsubscribeByUserIdResult {
	return UnsubscribeByUserIdResult{
		Item: NewSubscribeUserFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UnsubscribeByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UnsubscribeByUserIdResult) Pointer() *UnsubscribeByUserIdResult {
	return &p
}

type DescribeScoresResult struct {
	Items         []Score `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeScoresAsyncResult struct {
	result *DescribeScoresResult
	err    error
}

func NewDescribeScoresResultFromDict(data map[string]interface{}) DescribeScoresResult {
	return DescribeScoresResult{
		Items:         CastScores(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeScoresResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeScoresResult) Pointer() *DescribeScoresResult {
	return &p
}

type DescribeScoresByUserIdResult struct {
	Items         []Score `json:"items"`
	NextPageToken *string `json:"nextPageToken"`
}

type DescribeScoresByUserIdAsyncResult struct {
	result *DescribeScoresByUserIdResult
	err    error
}

func NewDescribeScoresByUserIdResultFromDict(data map[string]interface{}) DescribeScoresByUserIdResult {
	return DescribeScoresByUserIdResult{
		Items:         CastScores(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeScoresByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastScoresFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeScoresByUserIdResult) Pointer() *DescribeScoresByUserIdResult {
	return &p
}

type GetScoreResult struct {
	Item *Score `json:"item"`
}

type GetScoreAsyncResult struct {
	result *GetScoreResult
	err    error
}

func NewGetScoreResultFromDict(data map[string]interface{}) GetScoreResult {
	return GetScoreResult{
		Item: NewScoreFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetScoreResult) Pointer() *GetScoreResult {
	return &p
}

type GetScoreByUserIdResult struct {
	Item *Score `json:"item"`
}

type GetScoreByUserIdAsyncResult struct {
	result *GetScoreByUserIdResult
	err    error
}

func NewGetScoreByUserIdResultFromDict(data map[string]interface{}) GetScoreByUserIdResult {
	return GetScoreByUserIdResult{
		Item: NewScoreFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetScoreByUserIdResult) Pointer() *GetScoreByUserIdResult {
	return &p
}

type DescribeRankingsResult struct {
	Items         []Ranking `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeRankingsAsyncResult struct {
	result *DescribeRankingsResult
	err    error
}

func NewDescribeRankingsResultFromDict(data map[string]interface{}) DescribeRankingsResult {
	return DescribeRankingsResult{
		Items:         CastRankings(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRankingsResult) Pointer() *DescribeRankingsResult {
	return &p
}

type DescribeRankingssByUserIdResult struct {
	Items         []Ranking `json:"items"`
	NextPageToken *string   `json:"nextPageToken"`
}

type DescribeRankingssByUserIdAsyncResult struct {
	result *DescribeRankingssByUserIdResult
	err    error
}

func NewDescribeRankingssByUserIdResultFromDict(data map[string]interface{}) DescribeRankingssByUserIdResult {
	return DescribeRankingssByUserIdResult{
		Items:         CastRankings(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRankingssByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRankingssByUserIdResult) Pointer() *DescribeRankingssByUserIdResult {
	return &p
}

type DescribeNearRankingsResult struct {
	Items []Ranking `json:"items"`
}

type DescribeNearRankingsAsyncResult struct {
	result *DescribeNearRankingsResult
	err    error
}

func NewDescribeNearRankingsResultFromDict(data map[string]interface{}) DescribeNearRankingsResult {
	return DescribeNearRankingsResult{
		Items: CastRankings(core.CastArray(data["items"])),
	}
}

func (p DescribeNearRankingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRankingsFromDict(
			p.Items,
		),
	}
}

func (p DescribeNearRankingsResult) Pointer() *DescribeNearRankingsResult {
	return &p
}

type GetRankingResult struct {
	Item *Ranking `json:"item"`
}

type GetRankingAsyncResult struct {
	result *GetRankingResult
	err    error
}

func NewGetRankingResultFromDict(data map[string]interface{}) GetRankingResult {
	return GetRankingResult{
		Item: NewRankingFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRankingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRankingResult) Pointer() *GetRankingResult {
	return &p
}

type GetRankingByUserIdResult struct {
	Item *Ranking `json:"item"`
}

type GetRankingByUserIdAsyncResult struct {
	result *GetRankingByUserIdResult
	err    error
}

func NewGetRankingByUserIdResultFromDict(data map[string]interface{}) GetRankingByUserIdResult {
	return GetRankingByUserIdResult{
		Item: NewRankingFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRankingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRankingByUserIdResult) Pointer() *GetRankingByUserIdResult {
	return &p
}

type PutScoreResult struct {
	Item *Score `json:"item"`
}

type PutScoreAsyncResult struct {
	result *PutScoreResult
	err    error
}

func NewPutScoreResultFromDict(data map[string]interface{}) PutScoreResult {
	return PutScoreResult{
		Item: NewScoreFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p PutScoreResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p PutScoreResult) Pointer() *PutScoreResult {
	return &p
}

type PutScoreByUserIdResult struct {
	Item *Score `json:"item"`
}

type PutScoreByUserIdAsyncResult struct {
	result *PutScoreByUserIdResult
	err    error
}

func NewPutScoreByUserIdResultFromDict(data map[string]interface{}) PutScoreByUserIdResult {
	return PutScoreByUserIdResult{
		Item: NewScoreFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p PutScoreByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p PutScoreByUserIdResult) Pointer() *PutScoreByUserIdResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentRankingMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentRankingMasterResult struct {
	Item *CurrentRankingMaster `json:"item"`
}

type GetCurrentRankingMasterAsyncResult struct {
	result *GetCurrentRankingMasterResult
	err    error
}

func NewGetCurrentRankingMasterResultFromDict(data map[string]interface{}) GetCurrentRankingMasterResult {
	return GetCurrentRankingMasterResult{
		Item: NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentRankingMasterResult) Pointer() *GetCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterResult struct {
	Item *CurrentRankingMaster `json:"item"`
}

type UpdateCurrentRankingMasterAsyncResult struct {
	result *UpdateCurrentRankingMasterResult
	err    error
}

func NewUpdateCurrentRankingMasterResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterResult {
	return UpdateCurrentRankingMasterResult{
		Item: NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRankingMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRankingMasterResult) Pointer() *UpdateCurrentRankingMasterResult {
	return &p
}

type UpdateCurrentRankingMasterFromGitHubResult struct {
	Item *CurrentRankingMaster `json:"item"`
}

type UpdateCurrentRankingMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRankingMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRankingMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRankingMasterFromGitHubResult {
	return UpdateCurrentRankingMasterFromGitHubResult{
		Item: NewCurrentRankingMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRankingMasterFromGitHubResult) Pointer() *UpdateCurrentRankingMasterFromGitHubResult {
	return &p
}
