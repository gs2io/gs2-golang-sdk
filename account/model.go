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

package account

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
    /** アカウント引き継ぎ時にパスワードを変更するか */
	ChangePasswordIfTakeOver *bool   `json:"changePasswordIfTakeOver"`
    /** アカウント新規作成したときに実行するスクリプト */
	CreateAccountScript *ScriptSetting   `json:"createAccountScript"`
    /** 認証したときに実行するスクリプト */
	AuthenticationScript *ScriptSetting   `json:"authenticationScript"`
    /** 引き継ぎ情報登録したときに実行するスクリプト */
	CreateTakeOverScript *ScriptSetting   `json:"createTakeOverScript"`
    /** 引き継ぎ実行したときに実行するスクリプト */
	DoTakeOverScript *ScriptSetting   `json:"doTakeOverScript"`
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
    data["changePasswordIfTakeOver"] = p.ChangePasswordIfTakeOver
    if p.CreateAccountScript != nil {
        data["createAccountScript"] = *p.CreateAccountScript.ToDict()
    }
    if p.AuthenticationScript != nil {
        data["authenticationScript"] = *p.AuthenticationScript.ToDict()
    }
    if p.CreateTakeOverScript != nil {
        data["createTakeOverScript"] = *p.CreateTakeOverScript.ToDict()
    }
    if p.DoTakeOverScript != nil {
        data["doTakeOverScript"] = *p.DoTakeOverScript.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Account struct {
    /** ゲームプレイヤーアカウント */
	AccountId *core.String   `json:"accountId"`
    /** アカウントID */
	UserId *core.String   `json:"userId"`
    /** パスワード */
	Password *core.String   `json:"password"`
    /** 現在時刻に対する補正値（現在時刻を起点とした秒数） */
	TimeOffset *int32   `json:"timeOffset"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Account) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["accountId"] = p.AccountId
    data["userId"] = p.UserId
    data["password"] = p.Password
    data["timeOffset"] = p.TimeOffset
    data["createdAt"] = p.CreatedAt
    return &data
}

type TakeOver struct {
    /** 引き継ぎ設定 */
	TakeOverId *core.String   `json:"takeOverId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** スロット番号 */
	Type *int32   `json:"type"`
    /** 引き継ぎ用ユーザーID */
	UserIdentifier *core.String   `json:"userIdentifier"`
    /** パスワード */
	Password *core.String   `json:"password"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *TakeOver) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["takeOverId"] = p.TakeOverId
    data["userId"] = p.UserId
    data["type"] = p.Type
    data["userIdentifier"] = p.UserIdentifier
    data["password"] = p.Password
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
