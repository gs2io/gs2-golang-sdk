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

package inbox

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
    /** 説明文 */
	Description *core.String   `json:"description"`
    /** 開封したメッセージを自動的に削除するか */
	IsAutomaticDeletingEnabled *bool   `json:"isAutomaticDeletingEnabled"`
    /** メッセージ受信したときに実行するスクリプト */
	ReceiveMessageScript *ScriptSetting   `json:"receiveMessageScript"`
    /** メッセージ開封したときに実行するスクリプト */
	ReadMessageScript *ScriptSetting   `json:"readMessageScript"`
    /** メッセージ削除したときに実行するスクリプト */
	DeleteMessageScript *ScriptSetting   `json:"deleteMessageScript"`
    /** 報酬付与処理をジョブとして追加するキューネームスペース のGRN */
	QueueNamespaceId *core.String   `json:"queueNamespaceId"`
    /** 報酬付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *core.String   `json:"keyId"`
    /** メッセージを受信したときのプッシュ通知 */
	ReceiveNotification *NotificationSetting   `json:"receiveNotification"`
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
    data["isAutomaticDeletingEnabled"] = p.IsAutomaticDeletingEnabled
    if p.ReceiveMessageScript != nil {
        data["receiveMessageScript"] = *p.ReceiveMessageScript.ToDict()
    }
    if p.ReadMessageScript != nil {
        data["readMessageScript"] = *p.ReadMessageScript.ToDict()
    }
    if p.DeleteMessageScript != nil {
        data["deleteMessageScript"] = *p.DeleteMessageScript.ToDict()
    }
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
    if p.ReceiveNotification != nil {
        data["receiveNotification"] = *p.ReceiveNotification.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Message struct {
    /** メッセージ */
	MessageId *core.String   `json:"messageId"`
    /** メッセージID */
	Name *core.String   `json:"name"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** メッセージの内容に相当するメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 既読状態 */
	IsRead *bool   `json:"isRead"`
    /** 開封時に実行する入手アクション */
	ReadAcquireActions *[]*AcquireAction   `json:"readAcquireActions"`
    /** 作成日時 */
	ReceivedAt *int64   `json:"receivedAt"`
    /** 最終更新日時 */
	ReadAt *int64   `json:"readAt"`
    /** メッセージの有効期限 */
	ExpiresAt *int64   `json:"expiresAt"`
}

func (p *Message) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["messageId"] = p.MessageId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["metadata"] = p.Metadata
    data["isRead"] = p.IsRead
    if p.ReadAcquireActions != nil {
        var _readAcquireActions []*map[string]interface {}
        for _, item := range *p.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item.ToDict())
        }
        data["readAcquireActions"] = &_readAcquireActions
    }
    data["receivedAt"] = p.ReceivedAt
    data["readAt"] = p.ReadAt
    data["expiresAt"] = p.ExpiresAt
    return &data
}

type CurrentMessageMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentMessageMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type GlobalMessageMaster struct {
    /** 全ユーザに向けたメッセージ */
	GlobalMessageId *core.String   `json:"globalMessageId"`
    /** 全ユーザに向けたメッセージ名 */
	Name *core.String   `json:"name"`
    /** 全ユーザに向けたメッセージの内容に相当するメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 開封時に実行する入手アクション */
	ReadAcquireActions *[]*AcquireAction   `json:"readAcquireActions"`
    /** メッセージを受信したあとメッセージが削除されるまでの期間 */
	ExpiresTimeSpan *TimeSpan   `json:"expiresTimeSpan"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 全ユーザに向けたメッセージの受信期限 */
	ExpiresAt *int64   `json:"expiresAt"`
}

func (p *GlobalMessageMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["globalMessageId"] = p.GlobalMessageId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.ReadAcquireActions != nil {
        var _readAcquireActions []*map[string]interface {}
        for _, item := range *p.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item.ToDict())
        }
        data["readAcquireActions"] = &_readAcquireActions
    }
    if p.ExpiresTimeSpan != nil {
        data["expiresTimeSpan"] = *p.ExpiresTimeSpan.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["expiresAt"] = p.ExpiresAt
    return &data
}

type GlobalMessage struct {
    /** 全ユーザに向けたメッセージ */
	GlobalMessageId *core.String   `json:"globalMessageId"`
    /** 全ユーザに向けたメッセージ名 */
	Name *core.String   `json:"name"`
    /** 全ユーザに向けたメッセージの内容に相当するメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 開封時に実行する入手アクション */
	ReadAcquireActions *[]*AcquireAction   `json:"readAcquireActions"`
    /** メッセージを受信したあとメッセージが削除されるまでの期間 */
	ExpiresTimeSpan *TimeSpan   `json:"expiresTimeSpan"`
    /** 全ユーザに向けたメッセージの有効期限 */
	ExpiresAt *int64   `json:"expiresAt"`
}

func (p *GlobalMessage) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["globalMessageId"] = p.GlobalMessageId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.ReadAcquireActions != nil {
        var _readAcquireActions []*map[string]interface {}
        for _, item := range *p.ReadAcquireActions {
            _readAcquireActions = append(_readAcquireActions, item.ToDict())
        }
        data["readAcquireActions"] = &_readAcquireActions
    }
    if p.ExpiresTimeSpan != nil {
        data["expiresTimeSpan"] = *p.ExpiresTimeSpan.ToDict()
    }
    data["expiresAt"] = p.ExpiresAt
    return &data
}

type Received struct {
    /** 受信済みグローバルメッセージ名 */
	ReceivedId *core.String   `json:"receivedId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 受信したグローバルメッセージ名 */
	ReceivedGlobalMessageNames *[]core.String   `json:"receivedGlobalMessageNames"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Received) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["receivedId"] = p.ReceivedId
    data["userId"] = p.UserId
    if p.ReceivedGlobalMessageNames != nil {
        var _receivedGlobalMessageNames []core.String
        for _, item := range *p.ReceivedGlobalMessageNames {
            _receivedGlobalMessageNames = append(_receivedGlobalMessageNames, item)
        }
        data["receivedGlobalMessageNames"] = &_receivedGlobalMessageNames
    }
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

type TimeSpan struct {
    /** 現在時刻からの日数 */
	Days *int32   `json:"days"`
    /** 現在時刻からの時間 */
	Hours *int32   `json:"hours"`
    /** 現在時刻からの分 */
	Minutes *int32   `json:"minutes"`
}

func (p *TimeSpan) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["days"] = p.Days
    data["hours"] = p.Hours
    data["minutes"] = p.Minutes
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
