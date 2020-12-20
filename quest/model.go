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

type Namespace struct {
    /** クエストを分類するカテゴリー */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** カテゴリ名 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** クエスト開始したときに実行するスクリプト */
	StartQuestScript *ScriptSetting   `json:"startQuestScript"`
    /** クエストクリアしたときに実行するスクリプト */
	CompleteQuestScript *ScriptSetting   `json:"completeQuestScript"`
    /** クエスト失敗したときに実行するスクリプト */
	FailedQuestScript *ScriptSetting   `json:"failedQuestScript"`
    /** 報酬付与処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *core.String   `json:"queueNamespaceId"`
    /** 報酬付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *core.String   `json:"keyId"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** None */
	Status *core.String   `json:"status"`
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
    if p.StartQuestScript != nil {
        data["startQuestScript"] = *p.StartQuestScript.ToDict()
    }
    if p.CompleteQuestScript != nil {
        data["completeQuestScript"] = *p.CompleteQuestScript.ToDict()
    }
    if p.FailedQuestScript != nil {
        data["failedQuestScript"] = *p.FailedQuestScript.ToDict()
    }
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type QuestGroupModelMaster struct {
    /** クエストグループマスター */
	QuestGroupModelId *core.String   `json:"questGroupModelId"`
    /** クエストグループモデル名 */
	Name *core.String   `json:"name"`
    /** クエストグループマスターの説明 */
	Description *core.String   `json:"description"`
    /** クエストグループのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *QuestGroupModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questGroupModelId"] = p.QuestGroupModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type QuestModelMaster struct {
    /** クエストモデルマスター */
	QuestModelId *core.String   `json:"questModelId"`
    /** クエストモデル名 */
	QuestGroupName *core.String   `json:"questGroupName"`
    /** クエスト名 */
	Name *core.String   `json:"name"`
    /** クエストモデルの説明 */
	Description *core.String   `json:"description"`
    /** クエストのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** クエストの内容 */
	Contents *[]*Contents   `json:"contents"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
    /** クエストの参加料 */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** クエスト失敗時の報酬 */
	FailedAcquireActions *[]*AcquireAction   `json:"failedAcquireActions"`
    /** クエストに挑戦するためにクリアしておく必要のあるクエスト名 */
	PremiseQuestNames *[]core.String   `json:"premiseQuestNames"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *QuestModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questModelId"] = p.QuestModelId
    data["questGroupName"] = p.QuestGroupName
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.Contents != nil {
        var _contents []*map[string]interface {}
        for _, item := range *p.Contents {
            _contents = append(_contents, item.ToDict())
        }
        data["contents"] = &_contents
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range *p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.FailedAcquireActions != nil {
        var _failedAcquireActions []*map[string]interface {}
        for _, item := range *p.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item.ToDict())
        }
        data["failedAcquireActions"] = &_failedAcquireActions
    }
    if p.PremiseQuestNames != nil {
        var _premiseQuestNames []core.String
        for _, item := range *p.PremiseQuestNames {
            _premiseQuestNames = append(_premiseQuestNames, item)
        }
        data["premiseQuestNames"] = &_premiseQuestNames
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentQuestMaster struct {
    /** カテゴリ名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentQuestMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
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

type Contents struct {
    /** クエストモデルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** クエストクリア時の報酬 */
	CompleteAcquireActions *[]*AcquireAction   `json:"completeAcquireActions"`
    /** 抽選する重み */
	Weight *int32   `json:"weight"`
}

func (p *Contents) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["metadata"] = p.Metadata
    if p.CompleteAcquireActions != nil {
        var _completeAcquireActions []*map[string]interface {}
        for _, item := range *p.CompleteAcquireActions {
            _completeAcquireActions = append(_completeAcquireActions, item.ToDict())
        }
        data["completeAcquireActions"] = &_completeAcquireActions
    }
    data["weight"] = p.Weight
    return &data
}

type ConsumeAction struct {
    /** スタンプタスクで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** 消費リクエストのJSON */
	Request *core.String   `json:"request"`
}

func (p *ConsumeAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type AcquireAction struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** 入手リクエストのJSON */
	Request *core.String   `json:"request"`
}

func (p *AcquireAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type Reward struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** リクエストモデル */
	Request *core.String   `json:"request"`
    /** 入手するリソースGRN */
	ItemId *core.String   `json:"itemId"`
    /** 入手する数量 */
	Value *int32   `json:"value"`
}

func (p *Reward) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    data["itemId"] = p.ItemId
    data["value"] = p.Value
    return &data
}

type Config struct {
    /** 名前 */
	Key *core.String   `json:"key"`
    /** 値 */
	Value *core.String   `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["key"] = p.Key
    data["value"] = p.Value
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

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *core.String   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *core.String   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *core.String   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *core.String   `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerScriptId"] = p.TriggerScriptId
    data["doneTriggerTargetType"] = p.DoneTriggerTargetType
    data["doneTriggerScriptId"] = p.DoneTriggerScriptId
    data["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
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

type Progress struct {
    /** クエスト挑戦 */
	ProgressId *core.String   `json:"progressId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** トランザクションID */
	TransactionId *core.String   `json:"transactionId"`
    /** クエストモデル */
	QuestModelId *core.String   `json:"questModelId"`
    /** 乱数シード */
	RandomSeed *int64   `json:"randomSeed"`
    /** クエストで得られる報酬の上限 */
	Rewards *[]*Reward   `json:"rewards"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Progress) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["progressId"] = p.ProgressId
    data["userId"] = p.UserId
    data["transactionId"] = p.TransactionId
    data["questModelId"] = p.QuestModelId
    data["randomSeed"] = p.RandomSeed
    if p.Rewards != nil {
        var _rewards []*map[string]interface {}
        for _, item := range *p.Rewards {
            _rewards = append(_rewards, item.ToDict())
        }
        data["rewards"] = &_rewards
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CompletedQuestList struct {
    /** クエスト進行 */
	CompletedQuestListId *core.String   `json:"completedQuestListId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** クエストグループ名 */
	QuestGroupName *core.String   `json:"questGroupName"`
    /** 攻略済みのクエスト名一覧のリスト */
	CompleteQuestNames *[]core.String   `json:"completeQuestNames"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *CompletedQuestList) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["completedQuestListId"] = p.CompletedQuestListId
    data["userId"] = p.UserId
    data["questGroupName"] = p.QuestGroupName
    if p.CompleteQuestNames != nil {
        var _completeQuestNames []core.String
        for _, item := range *p.CompleteQuestNames {
            _completeQuestNames = append(_completeQuestNames, item)
        }
        data["completeQuestNames"] = &_completeQuestNames
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type QuestGroupModel struct {
    /** クエストグループ */
	QuestGroupModelId *core.String   `json:"questGroupModelId"`
    /** クエストグループ名 */
	Name *core.String   `json:"name"`
    /** クエストグループのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** グループに属するクエスト */
	Quests *[]*QuestModel   `json:"quests"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
}

func (p *QuestGroupModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questGroupModelId"] = p.QuestGroupModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Quests != nil {
        var _quests []*map[string]interface {}
        for _, item := range *p.Quests {
            _quests = append(_quests, item.ToDict())
        }
        data["quests"] = &_quests
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    return &data
}

type QuestModel struct {
    /** クエストモデル */
	QuestModelId *core.String   `json:"questModelId"`
    /** クエストモデル名 */
	Name *core.String   `json:"name"`
    /** クエストモデルのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** クエストの内容 */
	Contents *[]*Contents   `json:"contents"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
    /** クエストの参加料 */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** クエスト失敗時の報酬 */
	FailedAcquireActions *[]*AcquireAction   `json:"failedAcquireActions"`
    /** クエストに挑戦するためにクリアしておく必要のあるクエスト名 */
	PremiseQuestNames *[]core.String   `json:"premiseQuestNames"`
}

func (p *QuestModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questModelId"] = p.QuestModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Contents != nil {
        var _contents []*map[string]interface {}
        for _, item := range *p.Contents {
            _contents = append(_contents, item.ToDict())
        }
        data["contents"] = &_contents
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range *p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.FailedAcquireActions != nil {
        var _failedAcquireActions []*map[string]interface {}
        for _, item := range *p.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item.ToDict())
        }
        data["failedAcquireActions"] = &_failedAcquireActions
    }
    if p.PremiseQuestNames != nil {
        var _premiseQuestNames []core.String
        for _, item := range *p.PremiseQuestNames {
            _premiseQuestNames = append(_premiseQuestNames, item)
        }
        data["premiseQuestNames"] = &_premiseQuestNames
    }
    return &data
}
