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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeUsersResult struct {
    /** ユーザのリスト */
	Items         *[]*User	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeUsersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*User, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeUsersAsyncResult struct {
	result *DescribeUsersResult
	err    error
}

type CreateUserResult struct {
    /** 作成したユーザ */
	Item         *User	`json:"item"`
}

func (p *CreateUserResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateUserAsyncResult struct {
	result *CreateUserResult
	err    error
}

type UpdateUserResult struct {
    /** 更新後のユーザ */
	Item         *User	`json:"item"`
}

func (p *UpdateUserResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateUserAsyncResult struct {
	result *UpdateUserResult
	err    error
}

type GetUserResult struct {
    /** ユーザ */
	Item         *User	`json:"item"`
}

func (p *GetUserResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetUserAsyncResult struct {
	result *GetUserResult
	err    error
}

type DeleteUserResult struct {
}

func (p *DeleteUserResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteUserAsyncResult struct {
	result *DeleteUserResult
	err    error
}

type DescribeSecurityPoliciesResult struct {
    /** セキュリティポリシーのリスト */
	Items         *[]*SecurityPolicy	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeSecurityPoliciesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SecurityPolicy, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeSecurityPoliciesAsyncResult struct {
	result *DescribeSecurityPoliciesResult
	err    error
}

type DescribeCommonSecurityPoliciesResult struct {
    /** セキュリティポリシーのリスト */
	Items         *[]*SecurityPolicy	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeCommonSecurityPoliciesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SecurityPolicy, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeCommonSecurityPoliciesAsyncResult struct {
	result *DescribeCommonSecurityPoliciesResult
	err    error
}

type CreateSecurityPolicyResult struct {
    /** 作成したセキュリティポリシー */
	Item         *SecurityPolicy	`json:"item"`
}

func (p *CreateSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateSecurityPolicyAsyncResult struct {
	result *CreateSecurityPolicyResult
	err    error
}

type UpdateSecurityPolicyResult struct {
    /** 更新後のセキュリティポリシー */
	Item         *SecurityPolicy	`json:"item"`
}

func (p *UpdateSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateSecurityPolicyAsyncResult struct {
	result *UpdateSecurityPolicyResult
	err    error
}

type GetSecurityPolicyResult struct {
    /** セキュリティポリシー */
	Item         *SecurityPolicy	`json:"item"`
}

func (p *GetSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSecurityPolicyAsyncResult struct {
	result *GetSecurityPolicyResult
	err    error
}

type DeleteSecurityPolicyResult struct {
}

func (p *DeleteSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteSecurityPolicyAsyncResult struct {
	result *DeleteSecurityPolicyResult
	err    error
}

type DescribeIdentifiersResult struct {
    /** クレデンシャルのリスト */
	Items         *[]*Identifier	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeIdentifiersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Identifier, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeIdentifiersAsyncResult struct {
	result *DescribeIdentifiersResult
	err    error
}

type CreateIdentifierResult struct {
    /** 作成したクレデンシャル */
	Item         *Identifier	`json:"item"`
    /** クライアントシークレット */
	ClientSecret         *core.String	`json:"clientSecret"`
}

func (p *CreateIdentifierResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.ClientSecret != nil {
        data["clientSecret"] = p.ClientSecret
    }
    return &data
}

type CreateIdentifierAsyncResult struct {
	result *CreateIdentifierResult
	err    error
}

type GetIdentifierResult struct {
    /** クレデンシャル */
	Item         *Identifier	`json:"item"`
}

func (p *GetIdentifierResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetIdentifierAsyncResult struct {
	result *GetIdentifierResult
	err    error
}

type DeleteIdentifierResult struct {
}

func (p *DeleteIdentifierResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteIdentifierAsyncResult struct {
	result *DeleteIdentifierResult
	err    error
}

type DescribePasswordsResult struct {
    /** パスワードのリスト */
	Items         *[]*Password	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribePasswordsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Password, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribePasswordsAsyncResult struct {
	result *DescribePasswordsResult
	err    error
}

type CreatePasswordResult struct {
    /** 作成したパスワード */
	Item         *Password	`json:"item"`
}

func (p *CreatePasswordResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreatePasswordAsyncResult struct {
	result *CreatePasswordResult
	err    error
}

type GetPasswordResult struct {
    /** パスワード */
	Item         *Password	`json:"item"`
}

func (p *GetPasswordResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetPasswordAsyncResult struct {
	result *GetPasswordResult
	err    error
}

type DeletePasswordResult struct {
}

func (p *DeletePasswordResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeletePasswordAsyncResult struct {
	result *DeletePasswordResult
	err    error
}

type GetHasSecurityPolicyResult struct {
    /** セキュリティポリシーのリスト */
	Items         *[]*SecurityPolicy	`json:"items"`
}

func (p *GetHasSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SecurityPolicy, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type GetHasSecurityPolicyAsyncResult struct {
	result *GetHasSecurityPolicyResult
	err    error
}

type AttachSecurityPolicyResult struct {
    /** 新しくユーザーに割り当てたセキュリティポリシーのリスト */
	Items         *[]*SecurityPolicy	`json:"items"`
}

func (p *AttachSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SecurityPolicy, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type AttachSecurityPolicyAsyncResult struct {
	result *AttachSecurityPolicyResult
	err    error
}

type DetachSecurityPolicyResult struct {
    /** 剥奪したあとユーザーに引き続き割り当てられているセキュリティポリシーのリスト */
	Items         *[]*SecurityPolicy	`json:"items"`
}

func (p *DetachSecurityPolicyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*SecurityPolicy, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DetachSecurityPolicyAsyncResult struct {
	result *DetachSecurityPolicyResult
	err    error
}

type LoginResult struct {
    /** プロジェクトトークン */
	AccessToken         *core.String	`json:"accessToken"`
    /** Bearer */
	TokenType         *core.String	`json:"tokenType"`
    /** 有効期間(秒) */
	ExpiresIn         *int32	`json:"expiresIn"`
}

func (p *LoginResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.AccessToken != nil {
        data["accessToken"] = p.AccessToken
    }
    if p.TokenType != nil {
        data["tokenType"] = p.TokenType
    }
    if p.ExpiresIn != nil {
        data["expiresIn"] = p.ExpiresIn
    }
    return &data
}

type LoginAsyncResult struct {
	result *LoginResult
	err    error
}

type LoginByUserResult struct {
    /** プロジェクトトークン */
	Item         *ProjectToken	`json:"item"`
}

func (p *LoginByUserResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type LoginByUserAsyncResult struct {
	result *LoginByUserResult
	err    error
}
