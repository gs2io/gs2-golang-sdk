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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         *[]*Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Namespace, 0)
    	for _, item := range *p.Items {
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

type DescribeCategoryModelsResult struct {
    /** カテゴリのリスト */
	Items         *[]*CategoryModel	`json:"items"`
}

func (p *DescribeCategoryModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*CategoryModel, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeCategoryModelsAsyncResult struct {
	result *DescribeCategoryModelsResult
	err    error
}

type GetCategoryModelResult struct {
    /** カテゴリ */
	Item         *CategoryModel	`json:"item"`
}

func (p *GetCategoryModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCategoryModelAsyncResult struct {
	result *GetCategoryModelResult
	err    error
}

type DescribeCategoryModelMastersResult struct {
    /** カテゴリマスターのリスト */
	Items         *[]*CategoryModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCategoryModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*CategoryModelMaster, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeCategoryModelMastersAsyncResult struct {
	result *DescribeCategoryModelMastersResult
	err    error
}

type CreateCategoryModelMasterResult struct {
    /** 作成したカテゴリマスター */
	Item         *CategoryModelMaster	`json:"item"`
}

func (p *CreateCategoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateCategoryModelMasterAsyncResult struct {
	result *CreateCategoryModelMasterResult
	err    error
}

type GetCategoryModelMasterResult struct {
    /** カテゴリマスター */
	Item         *CategoryModelMaster	`json:"item"`
}

func (p *GetCategoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCategoryModelMasterAsyncResult struct {
	result *GetCategoryModelMasterResult
	err    error
}

type UpdateCategoryModelMasterResult struct {
    /** 更新したカテゴリマスター */
	Item         *CategoryModelMaster	`json:"item"`
}

func (p *UpdateCategoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCategoryModelMasterAsyncResult struct {
	result *UpdateCategoryModelMasterResult
	err    error
}

type DeleteCategoryModelMasterResult struct {
    /** 削除したカテゴリマスター */
	Item         *CategoryModelMaster	`json:"item"`
}

func (p *DeleteCategoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteCategoryModelMasterAsyncResult struct {
	result *DeleteCategoryModelMasterResult
	err    error
}

type DescribeSubscribesByCategoryNameResult struct {
    /** 購読対象のリスト */
	Items         *[]*SubscribeUser	`json:"items"`
}

func (p *DescribeSubscribesByCategoryNameResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SubscribeUser, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeSubscribesByCategoryNameAsyncResult struct {
	result *DescribeSubscribesByCategoryNameResult
	err    error
}

type DescribeSubscribesByCategoryNameAndUserIdResult struct {
    /** 購読対象のリスト */
	Items         *[]*SubscribeUser	`json:"items"`
}

func (p *DescribeSubscribesByCategoryNameAndUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SubscribeUser, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeSubscribesByCategoryNameAndUserIdAsyncResult struct {
	result *DescribeSubscribesByCategoryNameAndUserIdResult
	err    error
}

type SubscribeResult struct {
    /** 購読した購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *SubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SubscribeAsyncResult struct {
	result *SubscribeResult
	err    error
}

type SubscribeByUserIdResult struct {
    /** 購読した購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *SubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SubscribeByUserIdAsyncResult struct {
	result *SubscribeByUserIdResult
	err    error
}

type GetSubscribeResult struct {
    /** 購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *GetSubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

type GetSubscribeByUserIdResult struct {
    /** 購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *GetSubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

type UnsubscribeResult struct {
    /** 解除した購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *UnsubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnsubscribeAsyncResult struct {
	result *UnsubscribeResult
	err    error
}

type UnsubscribeByUserIdResult struct {
    /** 解除した購読対象 */
	Item         *SubscribeUser	`json:"item"`
}

func (p *UnsubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnsubscribeByUserIdAsyncResult struct {
	result *UnsubscribeByUserIdResult
	err    error
}

type DescribeScoresResult struct {
    /** スコアのリスト */
	Items         *[]*Score	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeScoresResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Score, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeScoresAsyncResult struct {
	result *DescribeScoresResult
	err    error
}

type DescribeScoresByUserIdResult struct {
    /** スコアのリスト */
	Items         *[]*Score	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeScoresByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Score, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeScoresByUserIdAsyncResult struct {
	result *DescribeScoresByUserIdResult
	err    error
}

type GetScoreResult struct {
    /** スコア */
	Item         *Score	`json:"item"`
}

func (p *GetScoreResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetScoreAsyncResult struct {
	result *GetScoreResult
	err    error
}

type GetScoreByUserIdResult struct {
    /** スコア */
	Item         *Score	`json:"item"`
}

func (p *GetScoreByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetScoreByUserIdAsyncResult struct {
	result *GetScoreByUserIdResult
	err    error
}

type DescribeRankingsResult struct {
    /** ランキングのリスト */
	Items         *[]*Ranking	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRankingsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Ranking, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRankingsAsyncResult struct {
	result *DescribeRankingsResult
	err    error
}

type DescribeRankingssByUserIdResult struct {
    /** ランキングのリスト */
	Items         *[]*Ranking	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRankingssByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Ranking, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRankingssByUserIdAsyncResult struct {
	result *DescribeRankingssByUserIdResult
	err    error
}

type DescribeNearRankingsResult struct {
    /** ランキングのリスト */
	Items         *[]*Ranking	`json:"items"`
}

func (p *DescribeNearRankingsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Ranking, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeNearRankingsAsyncResult struct {
	result *DescribeNearRankingsResult
	err    error
}

type GetRankingResult struct {
    /** ランキング */
	Item         *Ranking	`json:"item"`
}

func (p *GetRankingResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRankingAsyncResult struct {
	result *GetRankingResult
	err    error
}

type GetRankingByUserIdResult struct {
    /** ランキング */
	Item         *Ranking	`json:"item"`
}

func (p *GetRankingByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRankingByUserIdAsyncResult struct {
	result *GetRankingByUserIdResult
	err    error
}

type PutScoreResult struct {
    /** 登録したスコア */
	Item         *Score	`json:"item"`
}

func (p *PutScoreResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type PutScoreAsyncResult struct {
	result *PutScoreResult
	err    error
}

type PutScoreByUserIdResult struct {
    /** 登録したスコア */
	Item         *Score	`json:"item"`
}

func (p *PutScoreByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type PutScoreByUserIdAsyncResult struct {
	result *PutScoreByUserIdResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なランキング設定 */
	Item         *CurrentRankingMaster	`json:"item"`
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

type GetCurrentRankingMasterResult struct {
    /** 現在有効なランキング設定 */
	Item         *CurrentRankingMaster	`json:"item"`
}

func (p *GetCurrentRankingMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentRankingMasterAsyncResult struct {
	result *GetCurrentRankingMasterResult
	err    error
}

type UpdateCurrentRankingMasterResult struct {
    /** 更新した現在有効なランキング設定 */
	Item         *CurrentRankingMaster	`json:"item"`
}

func (p *UpdateCurrentRankingMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRankingMasterAsyncResult struct {
	result *UpdateCurrentRankingMasterResult
	err    error
}

type UpdateCurrentRankingMasterFromGitHubResult struct {
    /** 更新した現在有効なランキング設定 */
	Item         *CurrentRankingMaster	`json:"item"`
}

func (p *UpdateCurrentRankingMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRankingMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRankingMasterFromGitHubResult
	err    error
}
