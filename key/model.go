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

package key

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
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Key struct {
    /** 暗号鍵 */
	KeyId *core.String   `json:"keyId"`
    /** 暗号鍵名 */
	Name *core.String   `json:"name"`
    /** 説明文 */
	Description *core.String   `json:"description"`
    /** 暗号鍵 */
	Secret *core.String   `json:"secret"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Key) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["keyId"] = p.KeyId
    data["name"] = p.Name
    data["description"] = p.Description
    data["secret"] = p.Secret
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type GitHubApiKey struct {
    /** GitHub のAPIキー */
	ApiKeyId *core.String   `json:"apiKeyId"`
    /** GitHub APIキー名 */
	Name *core.String   `json:"name"`
    /** 説明文 */
	Description *core.String   `json:"description"`
    /** APIキー */
	ApiKey *core.String   `json:"apiKey"`
    /** APIキーの暗号化に使用する暗号鍵名 */
	EncryptionKeyName *core.String   `json:"encryptionKeyName"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *GitHubApiKey) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["apiKeyId"] = p.ApiKeyId
    data["name"] = p.Name
    data["description"] = p.Description
    data["apiKey"] = p.ApiKey
    data["encryptionKeyName"] = p.EncryptionKeyName
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
