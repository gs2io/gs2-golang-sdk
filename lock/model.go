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

package lock

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** カテゴリー名 */
	Name *string `json:"name"`
	/** ネームスペースの説明 */
	Description *string `json:"description"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
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
	if p.LogSetting != nil {
		data["logSetting"] = *p.LogSetting.ToDict()
	}
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Mutex struct {
	/** ミューテックス */
	MutexId *string `json:"mutexId"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** プロパティID */
	PropertyId *string `json:"propertyId"`
	/** ロックを取得したトランザクションID */
	TransactionId *string `json:"transactionId"`
	/** 参照回数 */
	ReferenceCount *int32 `json:"referenceCount"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** ロックの有効期限 */
	TtlAt *int64 `json:"ttlAt"`
}

func (p *Mutex) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["mutexId"] = p.MutexId
	data["userId"] = p.UserId
	data["propertyId"] = p.PropertyId
	data["transactionId"] = p.TransactionId
	data["referenceCount"] = p.ReferenceCount
	data["createdAt"] = p.CreatedAt
	data["ttlAt"] = p.TtlAt
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
