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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** 最終集計日時リスト */
	LastCalculatedAts *[]*CalculatedAt   `json:"lastCalculatedAts"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["description"] = p.Description
    if p.LastCalculatedAts != nil {
        var _lastCalculatedAts []*map[string]interface {}
        for _, item := range *p.LastCalculatedAts {
            _lastCalculatedAts = append(_lastCalculatedAts, item.ToDict())
        }
        data["lastCalculatedAts"] = &_lastCalculatedAts
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CategoryModel struct {
    /** カテゴリ */
	CategoryModelId *core.String   `json:"categoryModelId"`
    /** カテゴリ名 */
	Name *core.String   `json:"name"`
    /** カテゴリのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スコアの最小値 */
	MinimumValue *int64   `json:"minimumValue"`
    /** スコアの最大値 */
	MaximumValue *int64   `json:"maximumValue"`
    /** スコアのソート方向 */
	OrderDirection *core.String   `json:"orderDirection"`
    /** ランキングの種類 */
	Scope *core.String   `json:"scope"`
    /** ユーザID毎にスコアを1つしか登録されないようにする */
	UniqueByUserId *bool   `json:"uniqueByUserId"`
    /** スコアの固定集計開始時刻(時) */
	CalculateFixedTimingHour *int32   `json:"calculateFixedTimingHour"`
    /** スコアの固定集計開始時刻(分) */
	CalculateFixedTimingMinute *int32   `json:"calculateFixedTimingMinute"`
    /** スコアの集計間隔(分) */
	CalculateIntervalMinutes *int32   `json:"calculateIntervalMinutes"`
    /** スコアの登録可能期間とするイベントマスター のGRN */
	EntryPeriodEventId *core.String   `json:"entryPeriodEventId"`
    /** アクセス可能期間とするイベントマスター のGRN */
	AccessPeriodEventId *core.String   `json:"accessPeriodEventId"`
    /** ランキングの世代 */
	Generation *core.String   `json:"generation"`
}

func (p *CategoryModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["categoryModelId"] = p.CategoryModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["minimumValue"] = p.MinimumValue
    data["maximumValue"] = p.MaximumValue
    data["orderDirection"] = p.OrderDirection
    data["scope"] = p.Scope
    data["uniqueByUserId"] = p.UniqueByUserId
    data["calculateFixedTimingHour"] = p.CalculateFixedTimingHour
    data["calculateFixedTimingMinute"] = p.CalculateFixedTimingMinute
    data["calculateIntervalMinutes"] = p.CalculateIntervalMinutes
    data["entryPeriodEventId"] = p.EntryPeriodEventId
    data["accessPeriodEventId"] = p.AccessPeriodEventId
    data["generation"] = p.Generation
    return &data
}

type CategoryModelMaster struct {
    /** カテゴリマスター */
	CategoryModelId *core.String   `json:"categoryModelId"`
    /** カテゴリモデル名 */
	Name *core.String   `json:"name"`
    /** カテゴリマスターの説明 */
	Description *core.String   `json:"description"`
    /** カテゴリマスターのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スコアの最小値 */
	MinimumValue *int64   `json:"minimumValue"`
    /** スコアの最大値 */
	MaximumValue *int64   `json:"maximumValue"`
    /** スコアのソート方向 */
	OrderDirection *core.String   `json:"orderDirection"`
    /** ランキングの種類 */
	Scope *core.String   `json:"scope"`
    /** ユーザID毎にスコアを1つしか登録されないようにする */
	UniqueByUserId *bool   `json:"uniqueByUserId"`
    /** スコアの固定集計開始時刻(時) */
	CalculateFixedTimingHour *int32   `json:"calculateFixedTimingHour"`
    /** スコアの固定集計開始時刻(分) */
	CalculateFixedTimingMinute *int32   `json:"calculateFixedTimingMinute"`
    /** スコアの集計間隔(分) */
	CalculateIntervalMinutes *int32   `json:"calculateIntervalMinutes"`
    /** スコアの登録可能期間とするイベントマスター のGRN */
	EntryPeriodEventId *core.String   `json:"entryPeriodEventId"`
    /** アクセス可能期間とするイベントマスター のGRN */
	AccessPeriodEventId *core.String   `json:"accessPeriodEventId"`
    /** ランキングの世代 */
	Generation *core.String   `json:"generation"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *CategoryModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["categoryModelId"] = p.CategoryModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["minimumValue"] = p.MinimumValue
    data["maximumValue"] = p.MaximumValue
    data["orderDirection"] = p.OrderDirection
    data["scope"] = p.Scope
    data["uniqueByUserId"] = p.UniqueByUserId
    data["calculateFixedTimingHour"] = p.CalculateFixedTimingHour
    data["calculateFixedTimingMinute"] = p.CalculateFixedTimingMinute
    data["calculateIntervalMinutes"] = p.CalculateIntervalMinutes
    data["entryPeriodEventId"] = p.EntryPeriodEventId
    data["accessPeriodEventId"] = p.AccessPeriodEventId
    data["generation"] = p.Generation
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Subscribe struct {
    /** 購読 */
	SubscribeId *core.String   `json:"subscribeId"`
    /** カテゴリ名 */
	CategoryName *core.String   `json:"categoryName"`
    /** 購読するユーザID */
	UserId *core.String   `json:"userId"`
    /** 購読されるユーザIDリスト */
	TargetUserIds *[]core.String   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Subscribe) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["subscribeId"] = p.SubscribeId
    data["categoryName"] = p.CategoryName
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []core.String
        for _, item := range *p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
    }
    data["createdAt"] = p.CreatedAt
    return &data
}

type Score struct {
    /** スコア */
	ScoreId *core.String   `json:"scoreId"`
    /** カテゴリ名 */
	CategoryName *core.String   `json:"categoryName"`
    /** ユーザID */
	UserId *core.String   `json:"userId"`
    /** スコアのユニークID */
	UniqueId *core.String   `json:"uniqueId"`
    /** スコアを獲得したユーザID */
	ScorerUserId *core.String   `json:"scorerUserId"`
    /** スコア */
	Score *int64   `json:"score"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Score) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["scoreId"] = p.ScoreId
    data["categoryName"] = p.CategoryName
    data["userId"] = p.UserId
    data["uniqueId"] = p.UniqueId
    data["scorerUserId"] = p.ScorerUserId
    data["score"] = p.Score
    data["metadata"] = p.Metadata
    data["createdAt"] = p.CreatedAt
    return &data
}

type Ranking struct {
    /** 順位 */
	Rank *int64   `json:"rank"`
    /** 1位からのインデックス */
	Index *int64   `json:"index"`
    /** ユーザID */
	UserId *core.String   `json:"userId"`
    /** スコア */
	Score *int64   `json:"score"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Ranking) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["rank"] = p.Rank
    data["index"] = p.Index
    data["userId"] = p.UserId
    data["score"] = p.Score
    data["metadata"] = p.Metadata
    data["createdAt"] = p.CreatedAt
    return &data
}

type ResponseCache struct {
    /** None */
	Region *core.String   `json:"region"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *core.String   `json:"responseCacheId"`
    /** None */
	RequestHash *core.String   `json:"requestHash"`
    /** APIの応答内容 */
	Result *core.String   `json:"result"`
}

func (p *ResponseCache) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["region"] = p.Region
    data["ownerId"] = p.OwnerId
    data["responseCacheId"] = p.ResponseCacheId
    data["requestHash"] = p.RequestHash
    data["result"] = p.Result
    return &data
}

type CurrentRankingMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentRankingMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type CalculatedAt struct {
    /** カテゴリ名 */
	CategoryName *core.String   `json:"categoryName"`
    /** 集計日時 */
	CalculatedAt *int64   `json:"calculatedAt"`
}

func (p *CalculatedAt) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["categoryName"] = p.CategoryName
    data["calculatedAt"] = p.CalculatedAt
    return &data
}

type SubscribeUser struct {
    /** カテゴリ名 */
	CategoryName *core.String   `json:"categoryName"`
    /** 購読するユーザID */
	UserId *core.String   `json:"userId"`
    /** 購読されるユーザID */
	TargetUserId *core.String   `json:"targetUserId"`
}

func (p *SubscribeUser) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["categoryName"] = p.CategoryName
    data["userId"] = p.UserId
    data["targetUserId"] = p.TargetUserId
    return &data
}

type GitHubCheckoutSetting struct {
    /** リソースの取得に使用するGitHub のAPIキー のGRN */
	GitHubApiKeyId *core.String   `json:"gitHubApiKeyId"`
    /** リポジトリ名 */
	RepositoryName *core.String   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *core.String   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *core.String   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *core.String   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *core.String   `json:"branchName"`
    /** タグ名 */
	TagName *core.String   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gitHubApiKeyId"] = p.GitHubApiKeyId
    data["repositoryName"] = p.RepositoryName
    data["sourcePath"] = p.SourcePath
    data["referenceType"] = p.ReferenceType
    data["commitHash"] = p.CommitHash
    data["branchName"] = p.BranchName
    data["tagName"] = p.TagName
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
