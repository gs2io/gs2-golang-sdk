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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** レーティング計算機能を使用するか */
	EnableRating *bool   `json:"enableRating"`
    /** ギャザリング新規作成時のアクション */
	CreateGatheringTriggerType *string   `json:"createGatheringTriggerType"`
    /** ギャザリング新規作成時 にルームを作成するネームスペース のGRN */
	CreateGatheringTriggerRealtimeNamespaceId *string   `json:"createGatheringTriggerRealtimeNamespaceId"`
    /** ギャザリング新規作成時 に実行されるスクリプト のGRN */
	CreateGatheringTriggerScriptId *string   `json:"createGatheringTriggerScriptId"`
    /** マッチメイキング完了時のアクション */
	CompleteMatchmakingTriggerType *string   `json:"completeMatchmakingTriggerType"`
    /** マッチメイキング完了時 にルームを作成するネームスペース のGRN */
	CompleteMatchmakingTriggerRealtimeNamespaceId *string   `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
    /** マッチメイキング完了時 に実行されるスクリプト のGRN */
	CompleteMatchmakingTriggerScriptId *string   `json:"completeMatchmakingTriggerScriptId"`
    /** ギャザリングに新規プレイヤーが参加したときのプッシュ通知 */
	JoinNotification *NotificationSetting   `json:"joinNotification"`
    /** ギャザリングからプレイヤーが離脱したときのプッシュ通知 */
	LeaveNotification *NotificationSetting   `json:"leaveNotification"`
    /** マッチメイキングが完了したときのプッシュ通知 */
	CompleteNotification *NotificationSetting   `json:"completeNotification"`
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
    data["enableRating"] = p.EnableRating
    data["createGatheringTriggerType"] = p.CreateGatheringTriggerType
    data["createGatheringTriggerRealtimeNamespaceId"] = p.CreateGatheringTriggerRealtimeNamespaceId
    data["createGatheringTriggerScriptId"] = p.CreateGatheringTriggerScriptId
    data["completeMatchmakingTriggerType"] = p.CompleteMatchmakingTriggerType
    data["completeMatchmakingTriggerRealtimeNamespaceId"] = p.CompleteMatchmakingTriggerRealtimeNamespaceId
    data["completeMatchmakingTriggerScriptId"] = p.CompleteMatchmakingTriggerScriptId
    if p.JoinNotification != nil {
        data["joinNotification"] = *p.JoinNotification.ToDict()
    }
    if p.LeaveNotification != nil {
        data["leaveNotification"] = *p.LeaveNotification.ToDict()
    }
    if p.CompleteNotification != nil {
        data["completeNotification"] = *p.CompleteNotification.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Gathering struct {
    /** ギャザリング */
	GatheringId *string   `json:"gatheringId"`
    /** ギャザリング名 */
	Name *string   `json:"name"`
    /** 募集条件 */
	AttributeRanges []AttributeRange   `json:"attributeRanges"`
    /** 参加者 */
	CapacityOfRoles []CapacityOfRole   `json:"capacityOfRoles"`
    /** 参加を許可するユーザIDリスト */
	AllowUserIds []string   `json:"allowUserIds"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** ギャザリングの有効期限 */
	ExpiresAt *int64   `json:"expiresAt"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Gathering) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gatheringId"] = p.GatheringId
    data["name"] = p.Name
    if p.AttributeRanges != nil {
        var _attributeRanges []*map[string]interface {}
        for _, item := range p.AttributeRanges {
            _attributeRanges = append(_attributeRanges, item.ToDict())
        }
        data["attributeRanges"] = &_attributeRanges
    }
    if p.CapacityOfRoles != nil {
        var _capacityOfRoles []*map[string]interface {}
        for _, item := range p.CapacityOfRoles {
            _capacityOfRoles = append(_capacityOfRoles, item.ToDict())
        }
        data["capacityOfRoles"] = &_capacityOfRoles
    }
    if p.AllowUserIds != nil {
        var _allowUserIds []string
        for _, item := range p.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        data["allowUserIds"] = &_allowUserIds
    }
    data["metadata"] = p.Metadata
    data["expiresAt"] = p.ExpiresAt
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type RatingModelMaster struct {
    /** レーティングモデルマスター */
	RatingModelId *string   `json:"ratingModelId"`
    /** レーティングの種類名 */
	Name *string   `json:"name"`
    /** レーティングの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** レーティングモデルマスターの説明 */
	Description *string   `json:"description"`
    /** レート値の変動の大きさ */
	Volatility *int32   `json:"volatility"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *RatingModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["ratingModelId"] = p.RatingModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["volatility"] = p.Volatility
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type RatingModel struct {
    /** レーティングモデル */
	RatingModelId *string   `json:"ratingModelId"`
    /** レーティングの種類名 */
	Name *string   `json:"name"`
    /** レーティングの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** レート値の変動の大きさ */
	Volatility *int32   `json:"volatility"`
}

func (p *RatingModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["ratingModelId"] = p.RatingModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["volatility"] = p.Volatility
    return &data
}

type CurrentRatingModelMaster struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentRatingModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
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

type NotificationSetting struct {
    /** プッシュ通知に使用する GS2-Gateway のネームスペース のGRN */
	GatewayNamespaceId *string   `json:"gatewayNamespaceId"`
    /** モバイルプッシュ通知へ転送するか */
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
    /** モバイルプッシュ通知で使用するサウンドファイル名 */
	Sound *string   `json:"sound"`
}

func (p *NotificationSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gatewayNamespaceId"] = p.GatewayNamespaceId
    data["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
    data["sound"] = p.Sound
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

type AttributeRange struct {
    /** 属性名 */
	Name *string   `json:"name"`
    /** ギャザリング参加可能な属性値の最小値 */
	Min *int32   `json:"min"`
    /** ギャザリング参加可能な属性値の最大値 */
	Max *int32   `json:"max"`
}

func (p *AttributeRange) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["min"] = p.Min
    data["max"] = p.Max
    return &data
}

type CapacityOfRole struct {
    /** ロール名 */
	RoleName *string   `json:"roleName"`
    /** ロール名の別名リスト */
	RoleAliases []string   `json:"roleAliases"`
    /** 募集人数 */
	Capacity *int32   `json:"capacity"`
    /** 参加者のプレイヤー情報リスト */
	Participants []Player   `json:"participants"`
}

func (p *CapacityOfRole) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["roleName"] = p.RoleName
    if p.RoleAliases != nil {
        var _roleAliases []string
        for _, item := range p.RoleAliases {
            _roleAliases = append(_roleAliases, item)
        }
        data["roleAliases"] = &_roleAliases
    }
    data["capacity"] = p.Capacity
    if p.Participants != nil {
        var _participants []*map[string]interface {}
        for _, item := range p.Participants {
            _participants = append(_participants, item.ToDict())
        }
        data["participants"] = &_participants
    }
    return &data
}

type Attribute struct {
    /** 属性名 */
	Name *string   `json:"name"`
    /** 属性値 */
	Value *int32   `json:"value"`
}

func (p *Attribute) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["value"] = p.Value
    return &data
}

type Player struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 属性値のリスト */
	Attributes []Attribute   `json:"attributes"`
    /** ロール名 */
	RoleName *string   `json:"roleName"`
    /** 参加を拒否するユーザIDリスト */
	DenyUserIds []string   `json:"denyUserIds"`
}

func (p *Player) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    if p.Attributes != nil {
        var _attributes []*map[string]interface {}
        for _, item := range p.Attributes {
            _attributes = append(_attributes, item.ToDict())
        }
        data["attributes"] = &_attributes
    }
    data["roleName"] = p.RoleName
    if p.DenyUserIds != nil {
        var _denyUserIds []string
        for _, item := range p.DenyUserIds {
            _denyUserIds = append(_denyUserIds, item)
        }
        data["denyUserIds"] = &_denyUserIds
    }
    return &data
}

type Rating struct {
    /** レーティング */
	RatingId *string   `json:"ratingId"`
    /** レーティング名 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** None */
	RateValue *float32   `json:"rateValue"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Rating) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["ratingId"] = p.RatingId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["rateValue"] = p.RateValue
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type GameResult struct {
    /** 順位 */
	Rank *int32   `json:"rank"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
}

func (p *GameResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["rank"] = p.Rank
    data["userId"] = p.UserId
    return &data
}

type Ballot struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** レーティング計算に使用するレーティング名 */
	RatingName *string   `json:"ratingName"`
    /** 投票対象のギャザリング名 */
	GatheringName *string   `json:"gatheringName"`
    /** 参加人数 */
	NumberOfPlayer *int32   `json:"numberOfPlayer"`
}

func (p *Ballot) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    data["ratingName"] = p.RatingName
    data["gatheringName"] = p.GatheringName
    data["numberOfPlayer"] = p.NumberOfPlayer
    return &data
}

type SignedBallot struct {
    /** 投票用紙の署名対象のデータ */
	Body *string   `json:"body"`
    /** 投票用紙の署名 */
	Signature *string   `json:"signature"`
}

func (p *SignedBallot) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["body"] = p.Body
    data["signature"] = p.Signature
    return &data
}

type WrittenBallot struct {
    /** 投票用紙 */
	Ballot *Ballot   `json:"ballot"`
    /** 投票内容。対戦結果のリスト */
	GameResults []GameResult   `json:"gameResults"`
}

func (p *WrittenBallot) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Ballot != nil {
        data["ballot"] = *p.Ballot.ToDict()
    }
    if p.GameResults != nil {
        var _gameResults []*map[string]interface {}
        for _, item := range p.GameResults {
            _gameResults = append(_gameResults, item.ToDict())
        }
        data["gameResults"] = &_gameResults
    }
    return &data
}

type Vote struct {
    /** 投票状況 */
	VoteId *string   `json:"voteId"`
    /** レーティング名 */
	RatingName *string   `json:"ratingName"`
    /** 投票対象のギャザリング名 */
	GatheringName *string   `json:"gatheringName"`
    /** 投票用紙のリスト */
	WrittenBallots []WrittenBallot   `json:"writtenBallots"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Vote) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["voteId"] = p.VoteId
    data["ratingName"] = p.RatingName
    data["gatheringName"] = p.GatheringName
    if p.WrittenBallots != nil {
        var _writtenBallots []*map[string]interface {}
        for _, item := range p.WrittenBallots {
            _writtenBallots = append(_writtenBallots, item.ToDict())
        }
        data["writtenBallots"] = &_writtenBallots
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}
