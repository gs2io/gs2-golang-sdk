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

package gateway

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
    /** Firebase の通知送信に使用するシークレットトークン */
	FirebaseSecret *core.String   `json:"firebaseSecret"`
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
    data["firebaseSecret"] = p.FirebaseSecret
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type WebSocketSession struct {
    /** コネクションID */
	ConnectionId *core.String   `json:"connectionId"`
    /** API ID */
	ApiId *core.String   `json:"apiId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *WebSocketSession) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["connectionId"] = p.ConnectionId
    data["apiId"] = p.ApiId
    data["ownerId"] = p.OwnerId
    data["namespaceName"] = p.NamespaceName
    data["userId"] = p.UserId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type FirebaseToken struct {
    /** Firebaseデバイストークン のGRN */
	FirebaseTokenId *core.String   `json:"firebaseTokenId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** Firebase Cloud Messaging のデバイストークン */
	Token *core.String   `json:"token"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *FirebaseToken) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["firebaseTokenId"] = p.FirebaseTokenId
    data["ownerId"] = p.OwnerId
    data["userId"] = p.UserId
    data["token"] = p.Token
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Session struct {
    /** WebSocketセッション のGRN */
	SessionId *core.String   `json:"sessionId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** WebSocketセッション名 */
	SessionName *core.String   `json:"sessionName"`
    /** API Gateway の APIID */
	ApiId *core.String   `json:"apiId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Session) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["sessionId"] = p.SessionId
    data["ownerId"] = p.OwnerId
    data["userId"] = p.UserId
    data["sessionName"] = p.SessionName
    data["apiId"] = p.ApiId
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

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
