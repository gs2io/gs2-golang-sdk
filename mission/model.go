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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type MissionTaskModelMaster struct {
    /** ミッションタスクマスター */
	MissionTaskId *core.String   `json:"missionTaskId"`
    /** タスク名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** ミッションタスクの説明 */
	Description *core.String   `json:"description"`
    /** カウンター名 */
	CounterName *core.String   `json:"counterName"`
    /** 目標値 */
	TargetValue *int64   `json:"targetValue"`
    /** ミッション達成時の報酬 */
	CompleteAcquireActions *[]*AcquireAction   `json:"completeAcquireActions"`
    /** 達成報酬の受け取り可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
    /** このタスクに挑戦するために達成しておく必要のあるタスクの名前 */
	PremiseMissionTaskName *core.String   `json:"premiseMissionTaskName"`
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
        for _, item := range *p.CompleteAcquireActions {
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

type ScopedValue struct {
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
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

type CounterModel struct {
    /** カウンターの種類 */
	CounterId *core.String   `json:"counterId"`
    /** カウンター名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** カウンターのリセットタイミング */
	Scopes *[]*CounterScopeModel   `json:"scopes"`
    /** カウントアップ可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
}

func (p *CounterModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["counterId"] = p.CounterId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Scopes != nil {
        var _scopes []*map[string]interface {}
        for _, item := range *p.Scopes {
            _scopes = append(_scopes, item.ToDict())
        }
        data["scopes"] = &_scopes
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    return &data
}

type Counter struct {
    /** カウンター */
	CounterId *core.String   `json:"counterId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** カウンター名 */
	Name *core.String   `json:"name"`
    /** 値 */
	Values *[]*ScopedValue   `json:"values"`
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
        for _, item := range *p.Values {
            _values = append(_values, item.ToDict())
        }
        data["values"] = &_values
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CounterModelMaster struct {
    /** カウンターの種類マスター */
	CounterId *core.String   `json:"counterId"`
    /** カウンター名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** カウンターの種類マスターの説明 */
	Description *core.String   `json:"description"`
    /** カウンターのリセットタイミング */
	Scopes *[]*CounterScopeModel   `json:"scopes"`
    /** カウントアップ可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
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
        for _, item := range *p.Scopes {
            _scopes = append(_scopes, item.ToDict())
        }
        data["scopes"] = &_scopes
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
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

type CounterScopeModel struct {
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *core.String   `json:"resetDayOfWeek"`
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

type MissionGroupModel struct {
    /** ミッショングループ */
	MissionGroupId *core.String   `json:"missionGroupId"`
    /** グループ名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** タスクリスト */
	Tasks *[]*MissionTaskModel   `json:"tasks"`
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *core.String   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
    /** ミッションを達成したときの通知先ネームスペース のGRN */
	CompleteNotificationNamespaceId *core.String   `json:"completeNotificationNamespaceId"`
}

func (p *MissionGroupModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["missionGroupId"] = p.MissionGroupId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Tasks != nil {
        var _tasks []*map[string]interface {}
        for _, item := range *p.Tasks {
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

type CurrentMissionMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentMissionMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type NotificationSetting struct {
    /** プッシュ通知に使用する GS2-Gateway のネームスペース のGRN */
	GatewayNamespaceId *core.String   `json:"gatewayNamespaceId"`
    /** モバイルプッシュ通知へ転送するか */
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
    /** モバイルプッシュ通知で使用するサウンドファイル名 */
	Sound *core.String   `json:"sound"`
}

func (p *NotificationSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gatewayNamespaceId"] = p.GatewayNamespaceId
    data["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
    data["sound"] = p.Sound
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

type MissionGroupModelMaster struct {
    /** ミッショングループマスター */
	MissionGroupId *core.String   `json:"missionGroupId"`
    /** ミッショングループ名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** ミッショングループの説明 */
	Description *core.String   `json:"description"`
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *core.String   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
    /** ミッションを達成したときの通知先ネームスペース のGRN */
	CompleteNotificationNamespaceId *core.String   `json:"completeNotificationNamespaceId"`
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

type Complete struct {
    /** 達成状況 */
	CompleteId *core.String   `json:"completeId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** ミッショングループ名 */
	MissionGroupName *core.String   `json:"missionGroupName"`
    /** 達成済みのタスク名リスト */
	CompletedMissionTaskNames *[]core.String   `json:"completedMissionTaskNames"`
    /** 報酬の受け取り済みのタスク名リスト */
	ReceivedMissionTaskNames *[]core.String   `json:"receivedMissionTaskNames"`
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
        var _completedMissionTaskNames []core.String
        for _, item := range *p.CompletedMissionTaskNames {
            _completedMissionTaskNames = append(_completedMissionTaskNames, item)
        }
        data["completedMissionTaskNames"] = &_completedMissionTaskNames
    }
    if p.ReceivedMissionTaskNames != nil {
        var _receivedMissionTaskNames []core.String
        for _, item := range *p.ReceivedMissionTaskNames {
            _receivedMissionTaskNames = append(_receivedMissionTaskNames, item)
        }
        data["receivedMissionTaskNames"] = &_receivedMissionTaskNames
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** ミッションを達成したときに実行するスクリプト */
	MissionCompleteScript *ScriptSetting   `json:"missionCompleteScript"`
    /** カウンターを上昇したときに実行するスクリプト */
	CounterIncrementScript *ScriptSetting   `json:"counterIncrementScript"`
    /** 報酬を受け取ったときに実行するスクリプト */
	ReceiveRewardsScript *ScriptSetting   `json:"receiveRewardsScript"`
    /** 報酬付与処理をジョブとして追加するキューネームスペース のGRN */
	QueueNamespaceId *core.String   `json:"queueNamespaceId"`
    /** 報酬付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *core.String   `json:"keyId"`
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

type MissionTaskModel struct {
    /** ミッションタスク */
	MissionTaskId *core.String   `json:"missionTaskId"`
    /** タスク名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** カウンター名 */
	CounterName *core.String   `json:"counterName"`
    /** 目標値 */
	TargetValue *int64   `json:"targetValue"`
    /** ミッション達成時の報酬 */
	CompleteAcquireActions *[]*AcquireAction   `json:"completeAcquireActions"`
    /** 達成報酬の受け取り可能な期間を指定するイベントマスター のGRN */
	ChallengePeriodEventId *core.String   `json:"challengePeriodEventId"`
    /** このタスクに挑戦するために達成しておく必要のあるタスクの名前 */
	PremiseMissionTaskName *core.String   `json:"premiseMissionTaskName"`
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
        for _, item := range *p.CompleteAcquireActions {
            _completeAcquireActions = append(_completeAcquireActions, item.ToDict())
        }
        data["completeAcquireActions"] = &_completeAcquireActions
    }
    data["challengePeriodEventId"] = p.ChallengePeriodEventId
    data["premiseMissionTaskName"] = p.PremiseMissionTaskName
    return &data
}
