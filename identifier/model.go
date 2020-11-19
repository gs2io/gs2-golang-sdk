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

package identifier

type User struct {
	/** ユーザ */
	UserId *string `json:"userId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ユーザー名 */
	Name *string `json:"name"`
	/** ユーザの説明 */
	Description *string `json:"description"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *User) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["userId"] = p.UserId
	data["ownerId"] = p.OwnerId
	data["name"] = p.Name
	data["description"] = p.Description
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type SecurityPolicy struct {
	/** セキュリティポリシー */
	SecurityPolicyId *string `json:"securityPolicyId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** セキュリティポリシー名 */
	Name *string `json:"name"`
	/** セキュリティポリシーの説明 */
	Description *string `json:"description"`
	/** ポリシードキュメント */
	Policy *string `json:"policy"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *SecurityPolicy) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["securityPolicyId"] = p.SecurityPolicyId
	data["ownerId"] = p.OwnerId
	data["name"] = p.Name
	data["description"] = p.Description
	data["policy"] = p.Policy
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Identifier struct {
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** クライアントID */
	ClientId *string `json:"clientId"`
	/** ユーザー名 */
	UserName *string `json:"userName"`
	/** クライアントシークレット */
	ClientSecret *string `json:"clientSecret"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
}

func (p *Identifier) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["ownerId"] = p.OwnerId
	data["clientId"] = p.ClientId
	data["userName"] = p.UserName
	data["clientSecret"] = p.ClientSecret
	data["createdAt"] = p.CreatedAt
	return &data
}

type Password struct {
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ユーザ */
	UserId *string `json:"userId"`
	/** ユーザー名 */
	UserName *string `json:"userName"`
	/** パスワード */
	Password *string `json:"password"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
}

func (p *Password) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["ownerId"] = p.OwnerId
	data["userId"] = p.UserId
	data["userName"] = p.UserName
	data["password"] = p.Password
	data["createdAt"] = p.CreatedAt
	return &data
}

type AttachSecurityPolicy struct {
	/** ユーザ のGRN */
	UserId *string `json:"userId"`
	/** セキュリティポリシー のGRNのリスト */
	SecurityPolicyIds *[]string `json:"securityPolicyIds"`
	/** 作成日時 */
	AttachedAt *int64 `json:"attachedAt"`
}

func (p *AttachSecurityPolicy) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["userId"] = p.UserId
	if p.SecurityPolicyIds != nil {
		var _securityPolicyIds []string
		for _, item := range *p.SecurityPolicyIds {
			_securityPolicyIds = append(_securityPolicyIds, item)
		}
		data["securityPolicyIds"] = &_securityPolicyIds
	}
	data["attachedAt"] = p.AttachedAt
	return &data
}

type ProjectToken struct {
	/** プロジェクトトークン */
	Token *string `json:"token"`
}

func (p *ProjectToken) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["token"] = p.Token
	return &data
}
