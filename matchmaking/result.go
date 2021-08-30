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

package matchmaking

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

type DescribeGatheringsResult struct {
	Items         []Gathering `json:"items"`
	NextPageToken *string     `json:"nextPageToken"`
}

type DescribeGatheringsAsyncResult struct {
	result *DescribeGatheringsResult
	err    error
}

func NewDescribeGatheringsResultFromDict(data map[string]interface{}) DescribeGatheringsResult {
	return DescribeGatheringsResult{
		Items:         CastGatherings(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeGatheringsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastGatheringsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeGatheringsResult) Pointer() *DescribeGatheringsResult {
	return &p
}

type CreateGatheringResult struct {
	Item *Gathering `json:"item"`
}

type CreateGatheringAsyncResult struct {
	result *CreateGatheringResult
	err    error
}

func NewCreateGatheringResultFromDict(data map[string]interface{}) CreateGatheringResult {
	return CreateGatheringResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateGatheringResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateGatheringResult) Pointer() *CreateGatheringResult {
	return &p
}

type CreateGatheringByUserIdResult struct {
	Item *Gathering `json:"item"`
}

type CreateGatheringByUserIdAsyncResult struct {
	result *CreateGatheringByUserIdResult
	err    error
}

func NewCreateGatheringByUserIdResultFromDict(data map[string]interface{}) CreateGatheringByUserIdResult {
	return CreateGatheringByUserIdResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateGatheringByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateGatheringByUserIdResult) Pointer() *CreateGatheringByUserIdResult {
	return &p
}

type UpdateGatheringResult struct {
	Item *Gathering `json:"item"`
}

type UpdateGatheringAsyncResult struct {
	result *UpdateGatheringResult
	err    error
}

func NewUpdateGatheringResultFromDict(data map[string]interface{}) UpdateGatheringResult {
	return UpdateGatheringResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateGatheringResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateGatheringResult) Pointer() *UpdateGatheringResult {
	return &p
}

type UpdateGatheringByUserIdResult struct {
	Item *Gathering `json:"item"`
}

type UpdateGatheringByUserIdAsyncResult struct {
	result *UpdateGatheringByUserIdResult
	err    error
}

func NewUpdateGatheringByUserIdResultFromDict(data map[string]interface{}) UpdateGatheringByUserIdResult {
	return UpdateGatheringByUserIdResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateGatheringByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateGatheringByUserIdResult) Pointer() *UpdateGatheringByUserIdResult {
	return &p
}

type DoMatchmakingByPlayerResult struct {
	Item                    *Gathering `json:"item"`
	MatchmakingContextToken *string    `json:"matchmakingContextToken"`
}

type DoMatchmakingByPlayerAsyncResult struct {
	result *DoMatchmakingByPlayerResult
	err    error
}

func NewDoMatchmakingByPlayerResultFromDict(data map[string]interface{}) DoMatchmakingByPlayerResult {
	return DoMatchmakingByPlayerResult{
		Item:                    NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingByPlayerResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingByPlayerResult) Pointer() *DoMatchmakingByPlayerResult {
	return &p
}

type DoMatchmakingResult struct {
	Item                    *Gathering `json:"item"`
	MatchmakingContextToken *string    `json:"matchmakingContextToken"`
}

type DoMatchmakingAsyncResult struct {
	result *DoMatchmakingResult
	err    error
}

func NewDoMatchmakingResultFromDict(data map[string]interface{}) DoMatchmakingResult {
	return DoMatchmakingResult{
		Item:                    NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingResult) Pointer() *DoMatchmakingResult {
	return &p
}

type DoMatchmakingByUserIdResult struct {
	Item                    *Gathering `json:"item"`
	MatchmakingContextToken *string    `json:"matchmakingContextToken"`
}

type DoMatchmakingByUserIdAsyncResult struct {
	result *DoMatchmakingByUserIdResult
	err    error
}

func NewDoMatchmakingByUserIdResultFromDict(data map[string]interface{}) DoMatchmakingByUserIdResult {
	return DoMatchmakingByUserIdResult{
		Item:                    NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":                    p.Item.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingByUserIdResult) Pointer() *DoMatchmakingByUserIdResult {
	return &p
}

type GetGatheringResult struct {
	Item *Gathering `json:"item"`
}

type GetGatheringAsyncResult struct {
	result *GetGatheringResult
	err    error
}

func NewGetGatheringResultFromDict(data map[string]interface{}) GetGatheringResult {
	return GetGatheringResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetGatheringResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetGatheringResult) Pointer() *GetGatheringResult {
	return &p
}

type CancelMatchmakingResult struct {
	Item *Gathering `json:"item"`
}

type CancelMatchmakingAsyncResult struct {
	result *CancelMatchmakingResult
	err    error
}

func NewCancelMatchmakingResultFromDict(data map[string]interface{}) CancelMatchmakingResult {
	return CancelMatchmakingResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CancelMatchmakingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CancelMatchmakingResult) Pointer() *CancelMatchmakingResult {
	return &p
}

type CancelMatchmakingByUserIdResult struct {
	Item *Gathering `json:"item"`
}

type CancelMatchmakingByUserIdAsyncResult struct {
	result *CancelMatchmakingByUserIdResult
	err    error
}

func NewCancelMatchmakingByUserIdResultFromDict(data map[string]interface{}) CancelMatchmakingByUserIdResult {
	return CancelMatchmakingByUserIdResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CancelMatchmakingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CancelMatchmakingByUserIdResult) Pointer() *CancelMatchmakingByUserIdResult {
	return &p
}

type DeleteGatheringResult struct {
	Item *Gathering `json:"item"`
}

type DeleteGatheringAsyncResult struct {
	result *DeleteGatheringResult
	err    error
}

func NewDeleteGatheringResultFromDict(data map[string]interface{}) DeleteGatheringResult {
	return DeleteGatheringResult{
		Item: NewGatheringFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteGatheringResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteGatheringResult) Pointer() *DeleteGatheringResult {
	return &p
}

type DescribeRatingModelMastersResult struct {
	Items         []RatingModelMaster `json:"items"`
	NextPageToken *string             `json:"nextPageToken"`
}

type DescribeRatingModelMastersAsyncResult struct {
	result *DescribeRatingModelMastersResult
	err    error
}

func NewDescribeRatingModelMastersResultFromDict(data map[string]interface{}) DescribeRatingModelMastersResult {
	return DescribeRatingModelMastersResult{
		Items:         CastRatingModelMasters(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRatingModelMastersResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingModelMastersFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRatingModelMastersResult) Pointer() *DescribeRatingModelMastersResult {
	return &p
}

type CreateRatingModelMasterResult struct {
	Item *RatingModelMaster `json:"item"`
}

type CreateRatingModelMasterAsyncResult struct {
	result *CreateRatingModelMasterResult
	err    error
}

func NewCreateRatingModelMasterResultFromDict(data map[string]interface{}) CreateRatingModelMasterResult {
	return CreateRatingModelMasterResult{
		Item: NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p CreateRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p CreateRatingModelMasterResult) Pointer() *CreateRatingModelMasterResult {
	return &p
}

type GetRatingModelMasterResult struct {
	Item *RatingModelMaster `json:"item"`
}

type GetRatingModelMasterAsyncResult struct {
	result *GetRatingModelMasterResult
	err    error
}

func NewGetRatingModelMasterResultFromDict(data map[string]interface{}) GetRatingModelMasterResult {
	return GetRatingModelMasterResult{
		Item: NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRatingModelMasterResult) Pointer() *GetRatingModelMasterResult {
	return &p
}

type UpdateRatingModelMasterResult struct {
	Item *RatingModelMaster `json:"item"`
}

type UpdateRatingModelMasterAsyncResult struct {
	result *UpdateRatingModelMasterResult
	err    error
}

func NewUpdateRatingModelMasterResultFromDict(data map[string]interface{}) UpdateRatingModelMasterResult {
	return UpdateRatingModelMasterResult{
		Item: NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateRatingModelMasterResult) Pointer() *UpdateRatingModelMasterResult {
	return &p
}

type DeleteRatingModelMasterResult struct {
	Item *RatingModelMaster `json:"item"`
}

type DeleteRatingModelMasterAsyncResult struct {
	result *DeleteRatingModelMasterResult
	err    error
}

func NewDeleteRatingModelMasterResultFromDict(data map[string]interface{}) DeleteRatingModelMasterResult {
	return DeleteRatingModelMasterResult{
		Item: NewRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRatingModelMasterResult) Pointer() *DeleteRatingModelMasterResult {
	return &p
}

type DescribeRatingModelsResult struct {
	Items []RatingModel `json:"items"`
}

type DescribeRatingModelsAsyncResult struct {
	result *DescribeRatingModelsResult
	err    error
}

func NewDescribeRatingModelsResultFromDict(data map[string]interface{}) DescribeRatingModelsResult {
	return DescribeRatingModelsResult{
		Items: CastRatingModels(core.CastArray(data["items"])),
	}
}

func (p DescribeRatingModelsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingModelsFromDict(
			p.Items,
		),
	}
}

func (p DescribeRatingModelsResult) Pointer() *DescribeRatingModelsResult {
	return &p
}

type GetRatingModelResult struct {
	Item *RatingModel `json:"item"`
}

type GetRatingModelAsyncResult struct {
	result *GetRatingModelResult
	err    error
}

func NewGetRatingModelResultFromDict(data map[string]interface{}) GetRatingModelResult {
	return GetRatingModelResult{
		Item: NewRatingModelFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRatingModelResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRatingModelResult) Pointer() *GetRatingModelResult {
	return &p
}

type ExportMasterResult struct {
	Item *CurrentRatingModelMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
	return ExportMasterResult{
		Item: NewCurrentRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentRatingModelMasterResult struct {
	Item *CurrentRatingModelMaster `json:"item"`
}

type GetCurrentRatingModelMasterAsyncResult struct {
	result *GetCurrentRatingModelMasterResult
	err    error
}

func NewGetCurrentRatingModelMasterResultFromDict(data map[string]interface{}) GetCurrentRatingModelMasterResult {
	return GetCurrentRatingModelMasterResult{
		Item: NewCurrentRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetCurrentRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetCurrentRatingModelMasterResult) Pointer() *GetCurrentRatingModelMasterResult {
	return &p
}

type UpdateCurrentRatingModelMasterResult struct {
	Item *CurrentRatingModelMaster `json:"item"`
}

type UpdateCurrentRatingModelMasterAsyncResult struct {
	result *UpdateCurrentRatingModelMasterResult
	err    error
}

func NewUpdateCurrentRatingModelMasterResultFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterResult {
	return UpdateCurrentRatingModelMasterResult{
		Item: NewCurrentRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRatingModelMasterResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRatingModelMasterResult) Pointer() *UpdateCurrentRatingModelMasterResult {
	return &p
}

type UpdateCurrentRatingModelMasterFromGitHubResult struct {
	Item *CurrentRatingModelMaster `json:"item"`
}

type UpdateCurrentRatingModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRatingModelMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentRatingModelMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterFromGitHubResult {
	return UpdateCurrentRatingModelMasterFromGitHubResult{
		Item: NewCurrentRatingModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubResult) Pointer() *UpdateCurrentRatingModelMasterFromGitHubResult {
	return &p
}

type DescribeRatingsResult struct {
	Items         []Rating `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeRatingsAsyncResult struct {
	result *DescribeRatingsResult
	err    error
}

func NewDescribeRatingsResultFromDict(data map[string]interface{}) DescribeRatingsResult {
	return DescribeRatingsResult{
		Items:         CastRatings(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRatingsResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRatingsResult) Pointer() *DescribeRatingsResult {
	return &p
}

type DescribeRatingsByUserIdResult struct {
	Items         []Rating `json:"items"`
	NextPageToken *string  `json:"nextPageToken"`
}

type DescribeRatingsByUserIdAsyncResult struct {
	result *DescribeRatingsByUserIdResult
	err    error
}

func NewDescribeRatingsByUserIdResultFromDict(data map[string]interface{}) DescribeRatingsByUserIdResult {
	return DescribeRatingsByUserIdResult{
		Items:         CastRatings(core.CastArray(data["items"])),
		NextPageToken: core.CastString(data["nextPageToken"]),
	}
}

func (p DescribeRatingsByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
			p.Items,
		),
		"nextPageToken": p.NextPageToken,
	}
}

func (p DescribeRatingsByUserIdResult) Pointer() *DescribeRatingsByUserIdResult {
	return &p
}

type GetRatingResult struct {
	Item *Rating `json:"item"`
}

type GetRatingAsyncResult struct {
	result *GetRatingResult
	err    error
}

func NewGetRatingResultFromDict(data map[string]interface{}) GetRatingResult {
	return GetRatingResult{
		Item: NewRatingFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRatingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRatingResult) Pointer() *GetRatingResult {
	return &p
}

type GetRatingByUserIdResult struct {
	Item *Rating `json:"item"`
}

type GetRatingByUserIdAsyncResult struct {
	result *GetRatingByUserIdResult
	err    error
}

func NewGetRatingByUserIdResultFromDict(data map[string]interface{}) GetRatingByUserIdResult {
	return GetRatingByUserIdResult{
		Item: NewRatingFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p GetRatingByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p GetRatingByUserIdResult) Pointer() *GetRatingByUserIdResult {
	return &p
}

type PutResultResult struct {
	Items []Rating `json:"items"`
}

type PutResultAsyncResult struct {
	result *PutResultResult
	err    error
}

func NewPutResultResultFromDict(data map[string]interface{}) PutResultResult {
	return PutResultResult{
		Items: CastRatings(core.CastArray(data["items"])),
	}
}

func (p PutResultResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"items": CastRatingsFromDict(
			p.Items,
		),
	}
}

func (p PutResultResult) Pointer() *PutResultResult {
	return &p
}

type DeleteRatingResult struct {
	Item *Rating `json:"item"`
}

type DeleteRatingAsyncResult struct {
	result *DeleteRatingResult
	err    error
}

func NewDeleteRatingResultFromDict(data map[string]interface{}) DeleteRatingResult {
	return DeleteRatingResult{
		Item: NewRatingFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p DeleteRatingResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p DeleteRatingResult) Pointer() *DeleteRatingResult {
	return &p
}

type GetBallotResult struct {
	Item      *Ballot `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetBallotAsyncResult struct {
	result *GetBallotResult
	err    error
}

func NewGetBallotResultFromDict(data map[string]interface{}) GetBallotResult {
	return GetBallotResult{
		Item:      NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetBallotResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetBallotResult) Pointer() *GetBallotResult {
	return &p
}

type GetBallotByUserIdResult struct {
	Item      *Ballot `json:"item"`
	Body      *string `json:"body"`
	Signature *string `json:"signature"`
}

type GetBallotByUserIdAsyncResult struct {
	result *GetBallotByUserIdResult
	err    error
}

func NewGetBallotByUserIdResultFromDict(data map[string]interface{}) GetBallotByUserIdResult {
	return GetBallotByUserIdResult{
		Item:      NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
		Body:      core.CastString(data["body"]),
		Signature: core.CastString(data["signature"]),
	}
}

func (p GetBallotByUserIdResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item":      p.Item.ToDict(),
		"body":      p.Body,
		"signature": p.Signature,
	}
}

func (p GetBallotByUserIdResult) Pointer() *GetBallotByUserIdResult {
	return &p
}

type VoteResult struct {
	Item *Ballot `json:"item"`
}

type VoteAsyncResult struct {
	result *VoteResult
	err    error
}

func NewVoteResultFromDict(data map[string]interface{}) VoteResult {
	return VoteResult{
		Item: NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VoteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p VoteResult) Pointer() *VoteResult {
	return &p
}

type VoteMultipleResult struct {
	Item *Ballot `json:"item"`
}

type VoteMultipleAsyncResult struct {
	result *VoteMultipleResult
	err    error
}

func NewVoteMultipleResultFromDict(data map[string]interface{}) VoteMultipleResult {
	return VoteMultipleResult{
		Item: NewBallotFromDict(core.CastMap(data["item"])).Pointer(),
	}
}

func (p VoteMultipleResult) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"item": p.Item.ToDict(),
	}
}

func (p VoteMultipleResult) Pointer() *VoteMultipleResult {
	return &p
}

type CommitVoteResult struct {
}

type CommitVoteAsyncResult struct {
	result *CommitVoteResult
	err    error
}

func NewCommitVoteResultFromDict(data map[string]interface{}) CommitVoteResult {
	return CommitVoteResult{}
}

func (p CommitVoteResult) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p CommitVoteResult) Pointer() *CommitVoteResult {
	return &p
}
