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

type LoginResult struct {
	/** アクセストークン */
	Token *string `json:"token"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** 有効期限 */
	Expire *int64 `json:"expire"`
}

func (p *LoginResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Token != nil {
		data["Token"] = p.Token
	}
	if p.UserId != nil {
		data["UserId"] = p.UserId
	}
	if p.Expire != nil {
		data["Expire"] = p.Expire
	}
	return &data
}

type LoginAsyncResult struct {
	result *LoginResult
	err    error
}

type LoginBySignatureResult struct {
	/** アクセストークン */
	Token *string `json:"token"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** 有効期限 */
	Expire *int64 `json:"expire"`
}

func (p *LoginBySignatureResult) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	if p.Token != nil {
		data["Token"] = p.Token
	}
	if p.UserId != nil {
		data["UserId"] = p.UserId
	}
	if p.Expire != nil {
		data["Expire"] = p.Expire
	}
	return &data
}

type LoginBySignatureAsyncResult struct {
	result *LoginBySignatureResult
	err    error
}
