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

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type AccessToken struct {
	Token *string `json:"token"`
	UserId *string `json:"userId"`
	Expire *int64 `json:"expire"`
}

func NewAccessTokenFromJson(data string) AccessToken {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAccessTokenFromDict(dict)
}

func NewAccessTokenFromDict(data map[string]interface{}) AccessToken {
    return AccessToken {
        Token: core.CastString(data["token"]),
        UserId: core.CastString(data["userId"]),
        Expire: core.CastInt64(data["expire"]),
    }
}

func (p AccessToken) ToDict() map[string]interface{} {
    
    var token *string
    if p.Token != nil {
        token = p.Token
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var expire *int64
    if p.Expire != nil {
        expire = p.Expire
    }
    return map[string]interface{} {
        "token": token,
        "userId": userId,
        "expire": expire,
    }
}

func (p AccessToken) Pointer() *AccessToken {
    return &p
}

func CastAccessTokens(data []interface{}) []AccessToken {
	v := make([]AccessToken, 0)
	for _, d := range data {
		v = append(v, NewAccessTokenFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccessTokensFromDict(data []AccessToken) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}