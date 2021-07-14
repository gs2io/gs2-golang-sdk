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

deny overwrite
*/

package auth

import "core"

type LoginResult struct {
    Token *core.AccessToken `json:"token"`
    UserId *string `json:"userId"`
    Expire *int64 `json:"expire"`
}

type LoginAsyncResult struct {
	result *LoginResult
	err    error
}

func NewLoginResultFromDict(data map[string]interface{}) LoginResult {
    return LoginResult {
        Token: core.CastString(data["token"]),
        UserId: core.CastString(data["userId"]),
        Expire: core.CastInt64(data["expire"]),
    }
}

func (p LoginResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "token": p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
    }
}

func (p LoginResult) Pointer() *LoginResult {
    return &p
}

type LoginBySignatureResult struct {
    Token *core.AccessToken `json:"token"`
    UserId *string `json:"userId"`
    Expire *int64 `json:"expire"`
}

type LoginBySignatureAsyncResult struct {
	result *LoginBySignatureResult
	err    error
}

func NewLoginBySignatureResultFromDict(data map[string]interface{}) LoginBySignatureResult {
    return LoginBySignatureResult {
        Token: core.CastString(data["token"]),
        UserId: core.CastString(data["userId"]),
        Expire: core.CastInt64(data["expire"]),
    }
}

func (p LoginBySignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "token": p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
    }
}

func (p LoginBySignatureResult) Pointer() *LoginBySignatureResult {
    return &p
}