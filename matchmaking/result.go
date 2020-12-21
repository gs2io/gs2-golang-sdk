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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Namespace, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
    /** 作成したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
    /** None */
	Status         *core.String	`json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
        data["status"] = p.Status
    }
    return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
    /** ネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
    /** 更新したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeGatheringsResult struct {
    /** ギャザリングのリスト */
	Items         []Gathering	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeGatheringsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Gathering, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeGatheringsAsyncResult struct {
	result *DescribeGatheringsResult
	err    error
}

type CreateGatheringResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *CreateGatheringResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateGatheringAsyncResult struct {
	result *CreateGatheringResult
	err    error
}

type CreateGatheringByUserIdResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *CreateGatheringByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateGatheringByUserIdAsyncResult struct {
	result *CreateGatheringByUserIdResult
	err    error
}

type UpdateGatheringResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *UpdateGatheringResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateGatheringAsyncResult struct {
	result *UpdateGatheringResult
	err    error
}

type UpdateGatheringByUserIdResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *UpdateGatheringByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateGatheringByUserIdAsyncResult struct {
	result *UpdateGatheringByUserIdResult
	err    error
}

type DoMatchmakingByPlayerResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
    /** マッチメイキングの状態を保持するトークン */
	MatchmakingContextToken         *core.String	`json:"matchmakingContextToken"`
}

func (p *DoMatchmakingByPlayerResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.MatchmakingContextToken != nil {
        data["matchmakingContextToken"] = p.MatchmakingContextToken
    }
    return &data
}

type DoMatchmakingByPlayerAsyncResult struct {
	result *DoMatchmakingByPlayerResult
	err    error
}

type DoMatchmakingResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
    /** マッチメイキングの状態を保持するトークン */
	MatchmakingContextToken         *core.String	`json:"matchmakingContextToken"`
}

func (p *DoMatchmakingResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.MatchmakingContextToken != nil {
        data["matchmakingContextToken"] = p.MatchmakingContextToken
    }
    return &data
}

type DoMatchmakingAsyncResult struct {
	result *DoMatchmakingResult
	err    error
}

type GetGatheringResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *GetGatheringResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetGatheringAsyncResult struct {
	result *GetGatheringResult
	err    error
}

type CancelMatchmakingResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *CancelMatchmakingResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CancelMatchmakingAsyncResult struct {
	result *CancelMatchmakingResult
	err    error
}

type CancelMatchmakingByUserIdResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *CancelMatchmakingByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CancelMatchmakingByUserIdAsyncResult struct {
	result *CancelMatchmakingByUserIdResult
	err    error
}

type DeleteGatheringResult struct {
    /** ギャザリング */
	Item         *Gathering	`json:"item"`
}

func (p *DeleteGatheringResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteGatheringAsyncResult struct {
	result *DeleteGatheringResult
	err    error
}

type DescribeRatingModelMastersResult struct {
    /** レーティングモデルマスターのリスト */
	Items         []RatingModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRatingModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]RatingModelMaster, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRatingModelMastersAsyncResult struct {
	result *DescribeRatingModelMastersResult
	err    error
}

type CreateRatingModelMasterResult struct {
    /** 作成したレーティングモデルマスター */
	Item         *RatingModelMaster	`json:"item"`
}

func (p *CreateRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRatingModelMasterAsyncResult struct {
	result *CreateRatingModelMasterResult
	err    error
}

type GetRatingModelMasterResult struct {
    /** レーティングモデルマスター */
	Item         *RatingModelMaster	`json:"item"`
}

func (p *GetRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRatingModelMasterAsyncResult struct {
	result *GetRatingModelMasterResult
	err    error
}

type UpdateRatingModelMasterResult struct {
    /** 更新したレーティングモデルマスター */
	Item         *RatingModelMaster	`json:"item"`
}

func (p *UpdateRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateRatingModelMasterAsyncResult struct {
	result *UpdateRatingModelMasterResult
	err    error
}

type DeleteRatingModelMasterResult struct {
    /** 削除したレーティングモデルマスター */
	Item         *RatingModelMaster	`json:"item"`
}

func (p *DeleteRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRatingModelMasterAsyncResult struct {
	result *DeleteRatingModelMasterResult
	err    error
}

type DescribeRatingModelsResult struct {
    /** レーティングモデルのリスト */
	Items         []RatingModel	`json:"items"`
}

func (p *DescribeRatingModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]RatingModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeRatingModelsAsyncResult struct {
	result *DescribeRatingModelsResult
	err    error
}

type GetRatingModelResult struct {
    /** レーティングモデル */
	Item         *RatingModel	`json:"item"`
}

func (p *GetRatingModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRatingModelAsyncResult struct {
	result *GetRatingModelResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なレーティングマスター */
	Item         *CurrentRatingModelMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentRatingModelMasterResult struct {
    /** 現在有効なレーティングマスター */
	Item         *CurrentRatingModelMaster	`json:"item"`
}

func (p *GetCurrentRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentRatingModelMasterAsyncResult struct {
	result *GetCurrentRatingModelMasterResult
	err    error
}

type UpdateCurrentRatingModelMasterResult struct {
    /** 更新した現在有効なレーティングマスター */
	Item         *CurrentRatingModelMaster	`json:"item"`
}

func (p *UpdateCurrentRatingModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRatingModelMasterAsyncResult struct {
	result *UpdateCurrentRatingModelMasterResult
	err    error
}

type UpdateCurrentRatingModelMasterFromGitHubResult struct {
    /** 更新した現在有効なレーティングマスター */
	Item         *CurrentRatingModelMaster	`json:"item"`
}

func (p *UpdateCurrentRatingModelMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRatingModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRatingModelMasterFromGitHubResult
	err    error
}

type DescribeRatingsResult struct {
    /** レーティングのリスト */
	Items         []Rating	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRatingsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Rating, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRatingsAsyncResult struct {
	result *DescribeRatingsResult
	err    error
}

type DescribeRatingsByUserIdResult struct {
    /** レーティングのリスト */
	Items         []Rating	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRatingsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Rating, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRatingsByUserIdAsyncResult struct {
	result *DescribeRatingsByUserIdResult
	err    error
}

type GetRatingResult struct {
    /** レーティング */
	Item         *Rating	`json:"item"`
}

func (p *GetRatingResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRatingAsyncResult struct {
	result *GetRatingResult
	err    error
}

type GetRatingByUserIdResult struct {
    /** レーティング */
	Item         *Rating	`json:"item"`
}

func (p *GetRatingByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRatingByUserIdAsyncResult struct {
	result *GetRatingByUserIdResult
	err    error
}

type PutResultResult struct {
    /** 更新後の{model_name}の一覧 */
	Items         []Rating	`json:"items"`
}

func (p *PutResultResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Rating, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type PutResultAsyncResult struct {
	result *PutResultResult
	err    error
}

type DeleteRatingResult struct {
    /** レーティング */
	Item         *Rating	`json:"item"`
}

func (p *DeleteRatingResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRatingAsyncResult struct {
	result *DeleteRatingResult
	err    error
}

type GetBallotResult struct {
    /** 投票用紙 */
	Item         *Ballot	`json:"item"`
    /** 署名対象のデータ */
	Body         *core.String	`json:"body"`
    /** 署名データ */
	Signature         *core.String	`json:"signature"`
}

func (p *GetBallotResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Body != nil {
        data["body"] = p.Body
    }
    if p.Signature != nil {
        data["signature"] = p.Signature
    }
    return &data
}

type GetBallotAsyncResult struct {
	result *GetBallotResult
	err    error
}

type GetBallotByUserIdResult struct {
    /** 投票用紙 */
	Item         *Ballot	`json:"item"`
    /** 署名対象のデータ */
	Body         *core.String	`json:"body"`
    /** 署名データ */
	Signature         *core.String	`json:"signature"`
}

func (p *GetBallotByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Body != nil {
        data["body"] = p.Body
    }
    if p.Signature != nil {
        data["signature"] = p.Signature
    }
    return &data
}

type GetBallotByUserIdAsyncResult struct {
	result *GetBallotByUserIdResult
	err    error
}

type VoteResult struct {
    /** 投票用紙 */
	Item         *Ballot	`json:"item"`
}

func (p *VoteResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type VoteAsyncResult struct {
	result *VoteResult
	err    error
}

type VoteMultipleResult struct {
    /** 投票用紙 */
	Item         *Ballot	`json:"item"`
}

func (p *VoteMultipleResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type VoteMultipleAsyncResult struct {
	result *VoteMultipleResult
	err    error
}

type CommitVoteResult struct {
}

func (p *CommitVoteResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type CommitVoteAsyncResult struct {
	result *CommitVoteResult
	err    error
}
