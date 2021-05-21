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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** 最終集計日時リスト */
	LastCalculatedAts []CalculatedAt   `json:"lastCalculatedAts"`
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
        for _, item := range p.LastCalculatedAts {
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
	CategoryModelId *string   `json:"categoryModelId"`
    /** カテゴリ名 */
	Name *string   `json:"name"`
    /** カテゴリのメタデータ */
	Metadata *string   `json:"metadata"`
    /** スコアの最小値 */
	MinimumValue *int64   `json:"minimumValue"`
    /** スコアの最大値 */
	MaximumValue *int64   `json:"maximumValue"`
    /** スコアのソート方向 */
	OrderDirection *string   `json:"orderDirection"`
    /** ランキングの種類 */
	Scope *string   `json:"scope"`
    /** ユーザID毎にスコアを1つしか登録されないようにする */
	UniqueByUserId *bool   `json:"uniqueByUserId"`
    /** スコアの固定集計開始時刻(時) */
	CalculateFixedTimingHour *int32   `json:"calculateFixedTimingHour"`
    /** スコアの固定集計開始時刻(分) */
	CalculateFixedTimingMinute *int32   `json:"calculateFixedTimingMinute"`
    /** スコアの集計間隔(分) */
	CalculateIntervalMinutes *int32   `json:"calculateIntervalMinutes"`
    /** スコアの登録可能期間とするイベントマスター のGRN */
	EntryPeriodEventId *string   `json:"entryPeriodEventId"`
    /** アクセス可能期間とするイベントマスター のGRN */
	AccessPeriodEventId *string   `json:"accessPeriodEventId"`
    /** ランキングの世代 */
	Generation *string   `json:"generation"`
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
	CategoryModelId *string   `json:"categoryModelId"`
    /** カテゴリモデル名 */
	Name *string   `json:"name"`
    /** カテゴリマスターの説明 */
	Description *string   `json:"description"`
    /** カテゴリマスターのメタデータ */
	Metadata *string   `json:"metadata"`
    /** スコアの最小値 */
	MinimumValue *int64   `json:"minimumValue"`
    /** スコアの最大値 */
	MaximumValue *int64   `json:"maximumValue"`
    /** スコアのソート方向 */
	OrderDirection *string   `json:"orderDirection"`
    /** ランキングの種類 */
	Scope *string   `json:"scope"`
    /** ユーザID毎にスコアを1つしか登録されないようにする */
	UniqueByUserId *bool   `json:"uniqueByUserId"`
    /** スコアの固定集計開始時刻(時) */
	CalculateFixedTimingHour *int32   `json:"calculateFixedTimingHour"`
    /** スコアの固定集計開始時刻(分) */
	CalculateFixedTimingMinute *int32   `json:"calculateFixedTimingMinute"`
    /** スコアの集計間隔(分) */
	CalculateIntervalMinutes *int32   `json:"calculateIntervalMinutes"`
    /** スコアの登録可能期間とするイベントマスター のGRN */
	EntryPeriodEventId *string   `json:"entryPeriodEventId"`
    /** アクセス可能期間とするイベントマスター のGRN */
	AccessPeriodEventId *string   `json:"accessPeriodEventId"`
    /** ランキングの世代 */
	Generation *string   `json:"generation"`
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
	SubscribeId *string   `json:"subscribeId"`
    /** カテゴリ名 */
	CategoryName *string   `json:"categoryName"`
    /** 購読するユーザID */
	UserId *string   `json:"userId"`
    /** 購読されるユーザIDリスト */
	TargetUserIds []string   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Subscribe) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["subscribeId"] = p.SubscribeId
    data["categoryName"] = p.CategoryName
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []string
        for _, item := range p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
    }
    data["createdAt"] = p.CreatedAt
    return &data
}

type Score struct {
    /** スコア */
	ScoreId *string   `json:"scoreId"`
    /** カテゴリ名 */
	CategoryName *string   `json:"categoryName"`
    /** ユーザID */
	UserId *string   `json:"userId"`
    /** スコアのユニークID */
	UniqueId *string   `json:"uniqueId"`
    /** スコアを獲得したユーザID */
	ScorerUserId *string   `json:"scorerUserId"`
    /** スコア */
	Score *int64   `json:"score"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
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
	UserId *string   `json:"userId"`
    /** スコア */
	Score *int64   `json:"score"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
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
	Region *string   `json:"region"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *string   `json:"responseCacheId"`
    /** None */
	RequestHash *string   `json:"requestHash"`
    /** APIの応答内容 */
	Result *string   `json:"result"`
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
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentRankingMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
    return &data
}

type CalculatedAt struct {
    /** カテゴリ名 */
	CategoryName *string   `json:"categoryName"`
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
	CategoryName *string   `json:"categoryName"`
    /** 購読するユーザID */
	UserId *string   `json:"userId"`
    /** 購読されるユーザID */
	TargetUserId *string   `json:"targetUserId"`
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
	ApiKeyId *string   `json:"apiKeyId"`
    /** リポジトリ名 */
	RepositoryName *string   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *string   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *string   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *string   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *string   `json:"branchName"`
    /** タグ名 */
	TagName *string   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["apiKeyId"] = p.ApiKeyId
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
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
