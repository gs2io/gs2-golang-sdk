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

package chat

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** ゲームプレイヤーによるルームの作成を許可するか */
	AllowCreateRoom *bool   `json:"allowCreateRoom"`
    /** メッセージを投稿したときに実行するスクリプト */
	PostMessageScript *ScriptSetting   `json:"postMessageScript"`
    /** ルームを作成したときに実行するスクリプト */
	CreateRoomScript *ScriptSetting   `json:"createRoomScript"`
    /** ルームを削除したときに実行するスクリプト */
	DeleteRoomScript *ScriptSetting   `json:"deleteRoomScript"`
    /** ルームを購読したときに実行するスクリプト */
	SubscribeRoomScript *ScriptSetting   `json:"subscribeRoomScript"`
    /** ルームの購読を解除したときに実行するスクリプト */
	UnsubscribeRoomScript *ScriptSetting   `json:"unsubscribeRoomScript"`
    /** 購読しているルームに新しい投稿がきたときのプッシュ通知 */
	PostNotification *NotificationSetting   `json:"postNotification"`
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
    data["allowCreateRoom"] = p.AllowCreateRoom
    if p.PostMessageScript != nil {
        data["postMessageScript"] = *p.PostMessageScript.ToDict()
    }
    if p.CreateRoomScript != nil {
        data["createRoomScript"] = *p.CreateRoomScript.ToDict()
    }
    if p.DeleteRoomScript != nil {
        data["deleteRoomScript"] = *p.DeleteRoomScript.ToDict()
    }
    if p.SubscribeRoomScript != nil {
        data["subscribeRoomScript"] = *p.SubscribeRoomScript.ToDict()
    }
    if p.UnsubscribeRoomScript != nil {
        data["unsubscribeRoomScript"] = *p.UnsubscribeRoomScript.ToDict()
    }
    if p.PostNotification != nil {
        data["postNotification"] = *p.PostNotification.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Room struct {
    /** ルーム */
	RoomId *string   `json:"roomId"`
    /** ルーム名 */
	Name *string   `json:"name"`
    /** ルームを作成したユーザID */
	UserId *string   `json:"userId"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** メッセージを投稿するために必要となるパスワード */
	Password *string   `json:"password"`
    /** ルームに参加可能なユーザIDリスト */
	WhiteListUserIds []string   `json:"whiteListUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Room) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["roomId"] = p.RoomId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["metadata"] = p.Metadata
    data["password"] = p.Password
    if p.WhiteListUserIds != nil {
        var _whiteListUserIds []string
        for _, item := range p.WhiteListUserIds {
            _whiteListUserIds = append(_whiteListUserIds, item)
        }
        data["whiteListUserIds"] = &_whiteListUserIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Message struct {
    /** メッセージ */
	MessageId *string   `json:"messageId"`
    /** ルーム名 */
	RoomName *string   `json:"roomName"`
    /** メッセージ名 */
	Name *string   `json:"name"`
    /** 発言したユーザID */
	UserId *string   `json:"userId"`
    /** メッセージの種類を分類したい時の種類番号 */
	Category *int32   `json:"category"`
    /** メタデータ */
	Metadata *string   `json:"metadata"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Message) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["messageId"] = p.MessageId
    data["roomName"] = p.RoomName
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["category"] = p.Category
    data["metadata"] = p.Metadata
    data["createdAt"] = p.CreatedAt
    return &data
}

type Subscribe struct {
    /** 購読 */
	SubscribeId *string   `json:"subscribeId"`
    /** 購読するユーザID */
	UserId *string   `json:"userId"`
    /** 購読するルーム名 */
	RoomName *string   `json:"roomName"`
    /** 新着メッセージ通知を受け取るカテゴリリスト */
	NotificationTypes []NotificationType   `json:"notificationTypes"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Subscribe) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["subscribeId"] = p.SubscribeId
    data["userId"] = p.UserId
    data["roomName"] = p.RoomName
    if p.NotificationTypes != nil {
        var _notificationTypes []*map[string]interface {}
        for _, item := range p.NotificationTypes {
            _notificationTypes = append(_notificationTypes, item.ToDict())
        }
        data["notificationTypes"] = &_notificationTypes
    }
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

type NotificationType struct {
    /** 新着メッセージ通知を受け取るカテゴリ */
	Category *int32   `json:"category"`
    /** オフラインだった時にモバイルプッシュ通知に転送するか */
	EnableTransferMobilePushNotification *bool   `json:"enableTransferMobilePushNotification"`
}

func (p *NotificationType) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["category"] = p.Category
    data["enableTransferMobilePushNotification"] = p.EnableTransferMobilePushNotification
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
