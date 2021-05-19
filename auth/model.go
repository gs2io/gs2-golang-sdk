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

package auth

type AccessToken struct {
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** アクセストークン */
	Token *string   `json:"token"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 有効期限 */
	Expire *int64   `json:"expire"`
}

func (p *AccessToken) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["ownerId"] = p.OwnerId
    data["token"] = p.Token
    data["userId"] = p.UserId
    data["expire"] = p.Expire
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
