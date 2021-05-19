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

type Namespace struct {
    /** クエストを分類するカテゴリー */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** カテゴリ名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** クエスト開始したときに実行するスクリプト */
	StartQuestScript *ScriptSetting   `json:"startQuestScript"`
    /** クエストクリアしたときに実行するスクリプト */
	CompleteQuestScript *ScriptSetting   `json:"completeQuestScript"`
    /** クエスト失敗したときに実行するスクリプト */
	FailedQuestScript *ScriptSetting   `json:"failedQuestScript"`
    /** 報酬付与処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *string   `json:"queueNamespaceId"`
    /** 報酬付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *string   `json:"keyId"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** None */
	Status *string   `json:"status"`
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
	QuestGroupModelId *string   `json:"questGroupModelId"`
    /** クエストグループモデル名 */
	Name *string   `json:"name"`
    /** クエストグループマスターの説明 */
	Description *string   `json:"description"`
    /** クエストグループのメタデータ */
	Metadata *string   `json:"metadata"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
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
	QuestModelId *string   `json:"questModelId"`
    /** クエストモデル名 */
	QuestGroupName *string   `json:"questGroupName"`
    /** クエスト名 */
	Name *string   `json:"name"`
    /** クエストモデルの説明 */
	Description *string   `json:"description"`
    /** クエストのメタデータ */
	Metadata *string   `json:"metadata"`
    /** クエストの内容 */
	Contents []Contents   `json:"contents"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
    /** クエストの参加料 */
	ConsumeActions []ConsumeAction   `json:"consumeActions"`
    /** クエスト失敗時の報酬 */
	FailedAcquireActions []AcquireAction   `json:"failedAcquireActions"`
    /** クエストに挑戦するためにクリアしておく必要のあるクエスト名 */
	PremiseQuestNames []string   `json:"premiseQuestNames"`
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
        for _, item := range p.Contents {
            _contents = append(_contents, item.ToDict())
        }
        data["contents"] = &_contents
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.FailedAcquireActions != nil {
        var _failedAcquireActions []*map[string]interface {}
        for _, item := range p.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item.ToDict())
        }
        data["failedAcquireActions"] = &_failedAcquireActions
    }
    if p.PremiseQuestNames != nil {
        var _premiseQuestNames []string
        for _, item := range p.PremiseQuestNames {
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
	NamespaceName *string   `json:"namespaceName"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentQuestMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
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

type Contents struct {
    /** クエストモデルのメタデータ */
	Metadata *string   `json:"metadata"`
    /** クエストクリア時の報酬 */
	CompleteAcquireActions []AcquireAction   `json:"completeAcquireActions"`
    /** 抽選する重み */
	Weight *int32   `json:"weight"`
}

func (p *Contents) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["metadata"] = p.Metadata
    if p.CompleteAcquireActions != nil {
        var _completeAcquireActions []*map[string]interface {}
        for _, item := range p.CompleteAcquireActions {
            _completeAcquireActions = append(_completeAcquireActions, item.ToDict())
        }
        data["completeAcquireActions"] = &_completeAcquireActions
    }
    data["weight"] = p.Weight
    return &data
}

type ConsumeAction struct {
    /** スタンプタスクで実行するアクションの種類 */
	Action *string   `json:"action"`
    /** 消費リクエストのJSON */
	Request *string   `json:"request"`
}

func (p *ConsumeAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type AcquireAction struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *string   `json:"action"`
    /** 入手リクエストのJSON */
	Request *string   `json:"request"`
}

func (p *AcquireAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type Reward struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *string   `json:"action"`
    /** リクエストモデル */
	Request *string   `json:"request"`
    /** 入手するリソースGRN */
	ItemId *string   `json:"itemId"`
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
	Key *string   `json:"key"`
    /** 値 */
	Value *string   `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["key"] = p.Key
    data["value"] = p.Value
    return &data
}

type GitHubCheckoutSetting struct {
    /** リソースの取得に使用するGitHub のAPIキー のGRN */
	GitHubApiKeyId *string   `json:"gitHubApiKeyId"`
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
	TriggerScriptId *string   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *string   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *string   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *string   `json:"doneTriggerQueueNamespaceId"`
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
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}

type Progress struct {
    /** クエスト挑戦 */
	ProgressId *string   `json:"progressId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** トランザクションID */
	TransactionId *string   `json:"transactionId"`
    /** クエストモデル */
	QuestModelId *string   `json:"questModelId"`
    /** 乱数シード */
	RandomSeed *int64   `json:"randomSeed"`
    /** クエストで得られる報酬の上限 */
	Rewards []Reward   `json:"rewards"`
    /** クエストモデルのメタデータ */
	Metadata *string   `json:"metadata"`
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
        for _, item := range p.Rewards {
            _rewards = append(_rewards, item.ToDict())
        }
        data["rewards"] = &_rewards
    }
    data["metadata"] = p.Metadata
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CompletedQuestList struct {
    /** クエスト進行 */
	CompletedQuestListId *string   `json:"completedQuestListId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** クエストグループ名 */
	QuestGroupName *string   `json:"questGroupName"`
    /** 攻略済みのクエスト名一覧のリスト */
	CompleteQuestNames []string   `json:"completeQuestNames"`
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
        var _completeQuestNames []string
        for _, item := range p.CompleteQuestNames {
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
	QuestGroupModelId *string   `json:"questGroupModelId"`
    /** クエストグループ名 */
	Name *string   `json:"name"`
    /** クエストグループのメタデータ */
	Metadata *string   `json:"metadata"`
    /** グループに属するクエスト */
	Quests []QuestModel   `json:"quests"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
}

func (p *QuestGroupModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questGroupModelId"] = p.QuestGroupModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Quests != nil {
        var _quests []*map[string]interface {}
        for _, item := range p.Quests {
            _quests = append(_quests, item.ToDict())
        }
        data["quests"] = &_quests
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    return &data
}

type QuestModel struct {
    /** クエストモデル */
	QuestModelId *string   `json:"questModelId"`
    /** クエストモデル名 */
	Name *string   `json:"name"`
    /** クエストモデルのメタデータ */
	Metadata *string   `json:"metadata"`
    /** クエストの内容 */
	Contents []Contents   `json:"contents"`
    /** 挑戦可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
    /** クエストの参加料 */
	ConsumeActions []ConsumeAction   `json:"consumeActions"`
    /** クエスト失敗時の報酬 */
	FailedAcquireActions []AcquireAction   `json:"failedAcquireActions"`
    /** クエストに挑戦するためにクリアしておく必要のあるクエスト名 */
	PremiseQuestNames []string   `json:"premiseQuestNames"`
}

func (p *QuestModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["questModelId"] = p.QuestModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Contents != nil {
        var _contents []*map[string]interface {}
        for _, item := range p.Contents {
            _contents = append(_contents, item.ToDict())
        }
        data["contents"] = &_contents
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.FailedAcquireActions != nil {
        var _failedAcquireActions []*map[string]interface {}
        for _, item := range p.FailedAcquireActions {
            _failedAcquireActions = append(_failedAcquireActions, item.ToDict())
        }
        data["failedAcquireActions"] = &_failedAcquireActions
    }
    if p.PremiseQuestNames != nil {
        var _premiseQuestNames []string
        for _, item := range p.PremiseQuestNames {
            _premiseQuestNames = append(_premiseQuestNames, item)
        }
        data["premiseQuestNames"] = &_premiseQuestNames
    }
    return &data
}
