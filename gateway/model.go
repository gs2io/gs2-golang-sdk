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

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** 説明文 */
	Description *string `json:"description"`
	/** Firebase の通知送信に使用するシークレットトークン */
	FirebaseSecret *string `json:"firebaseSecret"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
	/** None */
	Status *string `json:"status"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
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
	data["status"] = p.Status
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type WebSocketSession struct {
	/** コネクションID */
	ConnectionId *string `json:"connectionId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	NamespaceName *string `json:"namespaceName"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *WebSocketSession) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["connectionId"] = p.ConnectionId
	data["ownerId"] = p.OwnerId
	data["namespaceName"] = p.NamespaceName
	data["userId"] = p.UserId
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type FirebaseToken struct {
	/** Firebaseデバイストークン のGRN */
	FirebaseTokenId *string `json:"firebaseTokenId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** Firebase Cloud Messaging のデバイストークン */
	Token *string `json:"token"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
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

type ResponseCache struct {
	/** None */
	Region *string `json:"region"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** レスポンスキャッシュ のGRN */
	ResponseCacheId *string `json:"responseCacheId"`
	/** None */
	RequestHash *string `json:"requestHash"`
	/** APIの応答内容 */
	Result *string `json:"result"`
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
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["loggingNamespaceId"] = p.LoggingNamespaceId
	return &data
}
