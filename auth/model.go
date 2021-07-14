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

import "core"

type AccessToken struct {
	Token  *string `json:"token"`
	UserId *string `json:"userId"`
	Expire *int64  `json:"expire"`
}

func NewAccessTokenFromDict(data map[string]interface{}) AccessToken {
	return AccessToken{
		Token:  core.CastString(data["token"]),
		UserId: core.CastString(data["userId"]),
		Expire: core.CastInt64(data["expire"]),
	}
}

func (p AccessToken) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"token":  p.Token,
		"userId": p.UserId,
		"expire": p.Expire,
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
