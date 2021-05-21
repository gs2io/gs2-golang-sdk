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

package mission

type CounterModelMaster struct {
    /** カウンターの種類マスター */
	CounterId *string   `json:"counterId"`
    /** カウンター名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** カウンターの種類マスターの説明 */
	Description *string   `json:"description"`
    /** カウンターのリセットタイミング */
	Scopes []CounterScopeModel   `json:"scopes"`
    /** カウントアップ可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *CounterModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["counterId"] = p.CounterId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    if p.Scopes != nil {
        var _scopes []*map[string]interface {}
        for _, item := range p.Scopes {
            _scopes = append(_scopes, item.ToDict())
        }
        data["scopes"] = &_scopes
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentMissionMaster struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentMissionMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
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

type Counter struct {
    /** カウンター */
	CounterId *string   `json:"counterId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** カウンター名 */
	Name *string   `json:"name"`
    /** 値 */
	Values []ScopedValue   `json:"values"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Counter) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["counterId"] = p.CounterId
    data["userId"] = p.UserId
    data["name"] = p.Name
    if p.Values != nil {
        var _values []*map[string]interface {}
        for _, item := range p.Values {
            _values = append(_values, item.ToDict())
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
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

type MissionTaskModelMaster struct {
    /** ミッションタスクマスター */
	MissionTaskId *string   `json:"missionTaskId"`
    /** タスク名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** ミッションタスクの説明 */
	Description *string   `json:"description"`
    /** カウンター名 */
	CounterName *string   `json:"counterName"`
    /** 目標値 */
	TargetValue *int64   `json:"targetValue"`
    /** ミッション達成時の報酬 */
	CompleteAcquireActions []AcquireAction   `json:"completeAcquireActions"`
    /** 達成報酬の受け取り可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
    /** このタスクに挑戦するために達成しておく必要のあるタスクの名前 */
	PremiseMissionTaskName *string   `json:"premiseMissionTaskName"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *MissionTaskModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["missionTaskId"] = p.MissionTaskId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["counterName"] = p.CounterName
    data["targetValue"] = p.TargetValue
    if p.CompleteAcquireActions != nil {
        var _completeAcquireActions []*map[string]interface {}
        for _, item := range p.CompleteAcquireActions {
            _completeAcquireActions = append(_completeAcquireActions, item.ToDict())
        }
        data["completeAcquireActions"] = &_completeAcquireActions
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["premiseMissionTaskName"] = p.PremiseMissionTaskName
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
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

type MissionGroupModelMaster struct {
    /** ミッショングループマスター */
	MissionGroupId *string   `json:"missionGroupId"`
    /** ミッショングループ名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** ミッショングループの説明 */
	Description *string   `json:"description"`
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *string   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
    /** ミッションを達成したときの通知先ネームスペース のGRN */
	CompleteNotificationNamespaceId *string   `json:"completeNotificationNamespaceId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *MissionGroupModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["missionGroupId"] = p.MissionGroupId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["resetType"] = p.ResetType
    data["resetDayOfMonth"] = p.ResetDayOfMonth
    data["resetDayOfWeek"] = p.ResetDayOfWeek
    data["resetHour"] = p.ResetHour
    data["completeNotificationNamespaceId"] = p.CompleteNotificationNamespaceId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type MissionGroupModel struct {
    /** ミッショングループ */
	MissionGroupId *string   `json:"missionGroupId"`
    /** グループ名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** タスクリスト */
	Tasks []MissionTaskModel   `json:"tasks"`
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *string   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
    /** ミッションを達成したときの通知先ネームスペース のGRN */
	CompleteNotificationNamespaceId *string   `json:"completeNotificationNamespaceId"`
}

func (p *MissionGroupModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["missionGroupId"] = p.MissionGroupId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Tasks != nil {
        var _tasks []*map[string]interface {}
        for _, item := range p.Tasks {
            _tasks = append(_tasks, item.ToDict())
        }
        data["tasks"] = &_tasks
    }
    data["resetType"] = p.ResetType
    data["resetDayOfMonth"] = p.ResetDayOfMonth
    data["resetDayOfWeek"] = p.ResetDayOfWeek
    data["resetHour"] = p.ResetHour
    data["completeNotificationNamespaceId"] = p.CompleteNotificationNamespaceId
    return &data
}

type Complete struct {
    /** 達成状況 */
	CompleteId *string   `json:"completeId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ミッショングループ名 */
	MissionGroupName *string   `json:"missionGroupName"`
    /** 達成済みのタスク名リスト */
	CompletedMissionTaskNames []string   `json:"completedMissionTaskNames"`
    /** 報酬の受け取り済みのタスク名リスト */
	ReceivedMissionTaskNames []string   `json:"receivedMissionTaskNames"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Complete) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["completeId"] = p.CompleteId
    data["userId"] = p.UserId
    data["missionGroupName"] = p.MissionGroupName
    if p.CompletedMissionTaskNames != nil {
        var _completedMissionTaskNames []string
        for _, item := range p.CompletedMissionTaskNames {
            _completedMissionTaskNames = append(_completedMissionTaskNames, item)
        }
        data["completedMissionTaskNames"] = &_completedMissionTaskNames
    }
    if p.ReceivedMissionTaskNames != nil {
        var _receivedMissionTaskNames []string
        for _, item := range p.ReceivedMissionTaskNames {
            _receivedMissionTaskNames = append(_receivedMissionTaskNames, item)
        }
        data["receivedMissionTaskNames"] = &_receivedMissionTaskNames
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type MissionTaskModel struct {
    /** ミッションタスク */
	MissionTaskId *string   `json:"missionTaskId"`
    /** タスク名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** カウンター名 */
	CounterName *string   `json:"counterName"`
    /** 目標値 */
	TargetValue *int64   `json:"targetValue"`
    /** ミッション達成時の報酬 */
	CompleteAcquireActions []AcquireAction   `json:"completeAcquireActions"`
    /** 達成報酬の受け取り可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
    /** このタスクに挑戦するために達成しておく必要のあるタスクの名前 */
	PremiseMissionTaskName *string   `json:"premiseMissionTaskName"`
}

func (p *MissionTaskModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["missionTaskId"] = p.MissionTaskId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["counterName"] = p.CounterName
    data["targetValue"] = p.TargetValue
    if p.CompleteAcquireActions != nil {
        var _completeAcquireActions []*map[string]interface {}
        for _, item := range p.CompleteAcquireActions {
            _completeAcquireActions = append(_completeAcquireActions, item.ToDict())
        }
        data["completeAcquireActions"] = &_completeAcquireActions
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["premiseMissionTaskName"] = p.PremiseMissionTaskName
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

type CounterModel struct {
    /** カウンターの種類 */
	CounterId *string   `json:"counterId"`
    /** カウンター名 */
	Name *string   `json:"name"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** カウンターのリセットタイミング */
	Scopes []CounterScopeModel   `json:"scopes"`
    /** カウントアップ可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *string   `json:"challengePeriodEventId"`
}

func (p *CounterModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["counterId"] = p.CounterId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Scopes != nil {
        var _scopes []*map[string]interface {}
        for _, item := range p.Scopes {
            _scopes = append(_scopes, item.ToDict())
        }
        data["scopes"] = &_scopes
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    return &data
}

type ScopedValue struct {
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** カウント */
	Value *int64   `json:"value"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ScopedValue) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["resetType"] = p.ResetType
    data["value"] = p.Value
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** ミッションを達成したときに実行するスクリプト */
	MissionCompleteScript *ScriptSetting   `json:"missionCompleteScript"`
    /** カウンターを上昇したときに実行するスクリプト */
	CounterIncrementScript *ScriptSetting   `json:"counterIncrementScript"`
    /** 報酬を受け取ったときに実行するスクリプト */
	ReceiveRewardsScript *ScriptSetting   `json:"receiveRewardsScript"`
    /** 報酬付与処理をジョブとして追加するキューネームスペース のGRN */
	QueueNamespaceId *string   `json:"queueNamespaceId"`
    /** 報酬付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *string   `json:"keyId"`
    /** ミッションのタスクを達成したときのプッシュ通知 */
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
    if p.MissionCompleteScript != nil {
        data["missionCompleteScript"] = *p.MissionCompleteScript.ToDict()
    }
    if p.CounterIncrementScript != nil {
        data["counterIncrementScript"] = *p.CounterIncrementScript.ToDict()
    }
    if p.ReceiveRewardsScript != nil {
        data["receiveRewardsScript"] = *p.ReceiveRewardsScript.ToDict()
    }
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
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

type CounterScopeModel struct {
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *string   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
}

func (p *CounterScopeModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["resetType"] = p.ResetType
    data["resetDayOfMonth"] = p.ResetDayOfMonth
    data["resetDayOfWeek"] = p.ResetDayOfWeek
    data["resetHour"] = p.ResetHour
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
