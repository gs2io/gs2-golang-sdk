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

package quest

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** クエストを分類するカテゴリーのリスト */
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
    /** 作成したクエストを分類するカテゴリー */
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
    /** クエストを分類するカテゴリー */
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
    /** 更新したクエストを分類するカテゴリー */
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
    /** 削除したクエストを分類するカテゴリー */
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

type DescribeQuestGroupModelMastersResult struct {
    /** クエストグループマスターのリスト */
	Items         *[]*QuestGroupModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeQuestGroupModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*QuestGroupModelMaster, 0)
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

type DescribeQuestGroupModelMastersAsyncResult struct {
	result *DescribeQuestGroupModelMastersResult
	err    error
}

type CreateQuestGroupModelMasterResult struct {
    /** 作成したクエストグループマスター */
	Item         *QuestGroupModelMaster	`json:"item"`
}

func (p *CreateQuestGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateQuestGroupModelMasterAsyncResult struct {
	result *CreateQuestGroupModelMasterResult
	err    error
}

type GetQuestGroupModelMasterResult struct {
    /** クエストグループマスター */
	Item         *QuestGroupModelMaster	`json:"item"`
}

func (p *GetQuestGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetQuestGroupModelMasterAsyncResult struct {
	result *GetQuestGroupModelMasterResult
	err    error
}

type UpdateQuestGroupModelMasterResult struct {
    /** 更新したクエストグループマスター */
	Item         *QuestGroupModelMaster	`json:"item"`
}

func (p *UpdateQuestGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateQuestGroupModelMasterAsyncResult struct {
	result *UpdateQuestGroupModelMasterResult
	err    error
}

type DeleteQuestGroupModelMasterResult struct {
    /** 削除したクエストグループマスター */
	Item         *QuestGroupModelMaster	`json:"item"`
}

func (p *DeleteQuestGroupModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteQuestGroupModelMasterAsyncResult struct {
	result *DeleteQuestGroupModelMasterResult
	err    error
}

type DescribeQuestModelMastersResult struct {
    /** クエストモデルマスターのリスト */
	Items         *[]*QuestModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeQuestModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*QuestModelMaster, 0)
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

type DescribeQuestModelMastersAsyncResult struct {
	result *DescribeQuestModelMastersResult
	err    error
}

type CreateQuestModelMasterResult struct {
    /** 作成したクエストモデルマスター */
	Item         *QuestModelMaster	`json:"item"`
}

func (p *CreateQuestModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateQuestModelMasterAsyncResult struct {
	result *CreateQuestModelMasterResult
	err    error
}

type GetQuestModelMasterResult struct {
    /** クエストモデルマスター */
	Item         *QuestModelMaster	`json:"item"`
}

func (p *GetQuestModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetQuestModelMasterAsyncResult struct {
	result *GetQuestModelMasterResult
	err    error
}

type UpdateQuestModelMasterResult struct {
    /** 更新したクエストモデルマスター */
	Item         *QuestModelMaster	`json:"item"`
}

func (p *UpdateQuestModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateQuestModelMasterAsyncResult struct {
	result *UpdateQuestModelMasterResult
	err    error
}

type DeleteQuestModelMasterResult struct {
    /** 削除したクエストモデルマスター */
	Item         *QuestModelMaster	`json:"item"`
}

func (p *DeleteQuestModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteQuestModelMasterAsyncResult struct {
	result *DeleteQuestModelMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なクエストマスター */
	Item         *CurrentQuestMaster	`json:"item"`
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

type GetCurrentQuestMasterResult struct {
    /** 現在有効なクエストマスター */
	Item         *CurrentQuestMaster	`json:"item"`
}

func (p *GetCurrentQuestMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentQuestMasterAsyncResult struct {
	result *GetCurrentQuestMasterResult
	err    error
}

type UpdateCurrentQuestMasterResult struct {
    /** 更新した現在有効なクエストマスター */
	Item         *CurrentQuestMaster	`json:"item"`
}

func (p *UpdateCurrentQuestMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentQuestMasterAsyncResult struct {
	result *UpdateCurrentQuestMasterResult
	err    error
}

type UpdateCurrentQuestMasterFromGitHubResult struct {
    /** 更新した現在有効なクエストマスター */
	Item         *CurrentQuestMaster	`json:"item"`
}

func (p *UpdateCurrentQuestMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentQuestMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentQuestMasterFromGitHubResult
	err    error
}

type DescribeProgressesByUserIdResult struct {
    /** クエスト挑戦のリスト */
	Items         *[]*Progress	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeProgressesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Progress, 0)
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

type DescribeProgressesByUserIdAsyncResult struct {
	result *DescribeProgressesByUserIdResult
	err    error
}

type CreateProgressByUserIdResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
}

func (p *CreateProgressByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateProgressByUserIdAsyncResult struct {
	result *CreateProgressByUserIdResult
	err    error
}

type GetProgressResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
    /** クエストグループ */
	QuestGroup         *QuestGroupModel	`json:"questGroup"`
    /** クエストモデル */
	Quest         *QuestModel	`json:"quest"`
}

func (p *GetProgressResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.QuestGroup != nil {
        data["questGroup"] = p.QuestGroup.ToDict()
    }
    if p.Quest != nil {
        data["quest"] = p.Quest.ToDict()
    }
    return &data
}

type GetProgressAsyncResult struct {
	result *GetProgressResult
	err    error
}

type GetProgressByUserIdResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
    /** クエストグループ */
	QuestGroup         *QuestGroupModel	`json:"questGroup"`
    /** クエストモデル */
	Quest         *QuestModel	`json:"quest"`
}

func (p *GetProgressByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.QuestGroup != nil {
        data["questGroup"] = p.QuestGroup.ToDict()
    }
    if p.Quest != nil {
        data["quest"] = p.Quest.ToDict()
    }
    return &data
}

type GetProgressByUserIdAsyncResult struct {
	result *GetProgressByUserIdResult
	err    error
}

type StartResult struct {
    /** クエストの開始処理の実行に使用するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *StartResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type StartAsyncResult struct {
	result *StartResult
	err    error
}

type StartByUserIdResult struct {
    /** クエストの開始処理の実行に使用するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *StartByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type StartByUserIdAsyncResult struct {
	result *StartByUserIdResult
	err    error
}

type EndResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
    /** 報酬付与処理の実行に使用するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *EndResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type EndAsyncResult struct {
	result *EndResult
	err    error
}

type EndByUserIdResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
    /** 報酬付与処理の実行に使用するスタンプシート */
	StampSheet         *core.String	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *core.String	`json:"stampSheetEncryptionKeyId"`
}

func (p *EndByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type EndByUserIdAsyncResult struct {
	result *EndByUserIdResult
	err    error
}

type DeleteProgressResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
}

func (p *DeleteProgressResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteProgressAsyncResult struct {
	result *DeleteProgressResult
	err    error
}

type DeleteProgressByUserIdResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
}

func (p *DeleteProgressByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteProgressByUserIdAsyncResult struct {
	result *DeleteProgressByUserIdResult
	err    error
}

type CreateProgressByStampSheetResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
}

func (p *CreateProgressByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateProgressByStampSheetAsyncResult struct {
	result *CreateProgressByStampSheetResult
	err    error
}

type DeleteProgressByStampTaskResult struct {
    /** クエスト挑戦 */
	Item         *Progress	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *DeleteProgressByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type DeleteProgressByStampTaskAsyncResult struct {
	result *DeleteProgressByStampTaskResult
	err    error
}

type DescribeCompletedQuestListsResult struct {
    /** クエスト進行のリスト */
	Items         *[]*CompletedQuestList	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCompletedQuestListsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*CompletedQuestList, 0)
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

type DescribeCompletedQuestListsAsyncResult struct {
	result *DescribeCompletedQuestListsResult
	err    error
}

type DescribeCompletedQuestListsByUserIdResult struct {
    /** クエスト進行のリスト */
	Items         *[]*CompletedQuestList	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCompletedQuestListsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*CompletedQuestList, 0)
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

type DescribeCompletedQuestListsByUserIdAsyncResult struct {
	result *DescribeCompletedQuestListsByUserIdResult
	err    error
}

type GetCompletedQuestListResult struct {
    /** クエスト進行 */
	Item         *CompletedQuestList	`json:"item"`
}

func (p *GetCompletedQuestListResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCompletedQuestListAsyncResult struct {
	result *GetCompletedQuestListResult
	err    error
}

type GetCompletedQuestListByUserIdResult struct {
    /** クエスト進行 */
	Item         *CompletedQuestList	`json:"item"`
}

func (p *GetCompletedQuestListByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCompletedQuestListByUserIdAsyncResult struct {
	result *GetCompletedQuestListByUserIdResult
	err    error
}

type DeleteCompletedQuestListByUserIdResult struct {
    /** クエスト進行 */
	Item         *CompletedQuestList	`json:"item"`
}

func (p *DeleteCompletedQuestListByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteCompletedQuestListByUserIdAsyncResult struct {
	result *DeleteCompletedQuestListByUserIdResult
	err    error
}

type DescribeQuestGroupModelsResult struct {
    /** クエストグループのリスト */
	Items         *[]*QuestGroupModel	`json:"items"`
}

func (p *DescribeQuestGroupModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*QuestGroupModel, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeQuestGroupModelsAsyncResult struct {
	result *DescribeQuestGroupModelsResult
	err    error
}

type GetQuestGroupModelResult struct {
    /** クエストグループ */
	Item         *QuestGroupModel	`json:"item"`
}

func (p *GetQuestGroupModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetQuestGroupModelAsyncResult struct {
	result *GetQuestGroupModelResult
	err    error
}

type DescribeQuestModelsResult struct {
    /** Noneのリスト */
	Items         *[]*QuestModel	`json:"items"`
}

func (p *DescribeQuestModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*QuestModel, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeQuestModelsAsyncResult struct {
	result *DescribeQuestModelsResult
	err    error
}

type GetQuestModelResult struct {
    /** None */
	Item         *QuestModel	`json:"item"`
}

func (p *GetQuestModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetQuestModelAsyncResult struct {
	result *GetQuestModelResult
	err    error
}
