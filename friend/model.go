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

package friend

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** フォローされたときに実行するスクリプト */
	FollowScript *ScriptSetting   `json:"followScript"`
    /** アンフォローされたときに実行するスクリプト */
	UnfollowScript *ScriptSetting   `json:"unfollowScript"`
    /** フレンドリクエストを発行したときに実行するスクリプト */
	SendRequestScript *ScriptSetting   `json:"sendRequestScript"`
    /** フレンドリクエストをキャンセルしたときに実行するスクリプト */
	CancelRequestScript *ScriptSetting   `json:"cancelRequestScript"`
    /** フレンドリクエストを承諾したときに実行するスクリプト */
	AcceptRequestScript *ScriptSetting   `json:"acceptRequestScript"`
    /** フレンドリクエストを拒否したときに実行するスクリプト */
	RejectRequestScript *ScriptSetting   `json:"rejectRequestScript"`
    /** フレンドを削除したときに実行するスクリプト */
	DeleteFriendScript *ScriptSetting   `json:"deleteFriendScript"`
    /** プロフィールを更新したときに実行するスクリプト */
	UpdateProfileScript *ScriptSetting   `json:"updateProfileScript"`
    /** フォローされたときのプッシュ通知 */
	FollowNotification *NotificationSetting   `json:"followNotification"`
    /** フレンドリクエストが届いたときのプッシュ通知 */
	ReceiveRequestNotification *NotificationSetting   `json:"receiveRequestNotification"`
    /** フレンドリクエストが承認されたときのプッシュ通知 */
	AcceptRequestNotification *NotificationSetting   `json:"acceptRequestNotification"`
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
    if p.FollowScript != nil {
        data["followScript"] = *p.FollowScript.ToDict()
    }
    if p.UnfollowScript != nil {
        data["unfollowScript"] = *p.UnfollowScript.ToDict()
    }
    if p.SendRequestScript != nil {
        data["sendRequestScript"] = *p.SendRequestScript.ToDict()
    }
    if p.CancelRequestScript != nil {
        data["cancelRequestScript"] = *p.CancelRequestScript.ToDict()
    }
    if p.AcceptRequestScript != nil {
        data["acceptRequestScript"] = *p.AcceptRequestScript.ToDict()
    }
    if p.RejectRequestScript != nil {
        data["rejectRequestScript"] = *p.RejectRequestScript.ToDict()
    }
    if p.DeleteFriendScript != nil {
        data["deleteFriendScript"] = *p.DeleteFriendScript.ToDict()
    }
    if p.UpdateProfileScript != nil {
        data["updateProfileScript"] = *p.UpdateProfileScript.ToDict()
    }
    if p.FollowNotification != nil {
        data["followNotification"] = *p.FollowNotification.ToDict()
    }
    if p.ReceiveRequestNotification != nil {
        data["receiveRequestNotification"] = *p.ReceiveRequestNotification.ToDict()
    }
    if p.AcceptRequestNotification != nil {
        data["acceptRequestNotification"] = *p.AcceptRequestNotification.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Profile struct {
    /** プロフィール */
	ProfileId *string   `json:"profileId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 公開されるプロフィール */
	PublicProfile *string   `json:"publicProfile"`
    /** フォロワー向けに公開されるプロフィール */
	FollowerProfile *string   `json:"followerProfile"`
    /** フレンド向けに公開されるプロフィール */
	FriendProfile *string   `json:"friendProfile"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Profile) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["profileId"] = p.ProfileId
    data["userId"] = p.UserId
    data["publicProfile"] = p.PublicProfile
    data["followerProfile"] = p.FollowerProfile
    data["friendProfile"] = p.FriendProfile
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Follow struct {
    /** フォロー */
	FollowId *string   `json:"followId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** フォローしているユーザーIDリスト */
	TargetUserIds []string   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Follow) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["followId"] = p.FollowId
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []string
        for _, item := range p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Friend struct {
    /** フレンド */
	FriendId *string   `json:"friendId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** フレンドのユーザーIDリスト */
	TargetUserIds []string   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Friend) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["friendId"] = p.FriendId
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []string
        for _, item := range p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type SendBox struct {
    /** フレンドリクエストの受信ボックス */
	SendBoxId *string   `json:"sendBoxId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** フレンドリクエストの宛先ユーザーIDリスト */
	TargetUserIds []string   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *SendBox) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["sendBoxId"] = p.SendBoxId
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []string
        for _, item := range p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Inbox struct {
    /** フレンドリクエストの受信ボックス */
	InboxId *string   `json:"inboxId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** フレンドリクエストを送ってきたユーザーIDリスト */
	FromUserIds []string   `json:"fromUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Inbox) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["inboxId"] = p.InboxId
    data["userId"] = p.UserId
    if p.FromUserIds != nil {
        var _fromUserIds []string
        for _, item := range p.FromUserIds {
            _fromUserIds = append(_fromUserIds, item)
        }
        data["fromUserIds"] = &_fromUserIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type BlackList struct {
    /** ブラックリスト */
	BlackListId *string   `json:"blackListId"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ブラックリストのユーザーIDリスト */
	TargetUserIds []string   `json:"targetUserIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *BlackList) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["blackListId"] = p.BlackListId
    data["userId"] = p.UserId
    if p.TargetUserIds != nil {
        var _targetUserIds []string
        for _, item := range p.TargetUserIds {
            _targetUserIds = append(_targetUserIds, item)
        }
        data["targetUserIds"] = &_targetUserIds
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

type FollowUser struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 公開されるプロフィール */
	PublicProfile *string   `json:"publicProfile"`
    /** フォロワー向けに公開されるプロフィール */
	FollowerProfile *string   `json:"followerProfile"`
}

func (p *FollowUser) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    data["publicProfile"] = p.PublicProfile
    data["followerProfile"] = p.FollowerProfile
    return &data
}

type FriendUser struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 公開されるプロフィール */
	PublicProfile *string   `json:"publicProfile"`
    /** フレンド向けに公開されるプロフィール */
	FriendProfile *string   `json:"friendProfile"`
}

func (p *FriendUser) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    data["publicProfile"] = p.PublicProfile
    data["friendProfile"] = p.FriendProfile
    return &data
}

type FriendRequest struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ユーザーID */
	TargetUserId *string   `json:"targetUserId"`
}

func (p *FriendRequest) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    data["targetUserId"] = p.TargetUserId
    return &data
}

type PublicProfile struct {
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 公開されるプロフィール */
	PublicProfile *string   `json:"publicProfile"`
}

func (p *PublicProfile) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["userId"] = p.UserId
    data["publicProfile"] = p.PublicProfile
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
