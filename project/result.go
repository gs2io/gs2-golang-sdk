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

package project

type CreateAccountResult struct {
    /** 作成したGS2アカウント */
	Item         *Account	`json:"item"`
}

func (p *CreateAccountResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateAccountAsyncResult struct {
	result *CreateAccountResult
	err    error
}

type VerifyResult struct {
    /** 有効化したGS2アカウント */
	Item         *Account	`json:"item"`
}

func (p *VerifyResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type VerifyAsyncResult struct {
	result *VerifyResult
	err    error
}

type SignInResult struct {
    /** サインインしたGS2アカウント */
	Item         *Account	`json:"item"`
    /** GS2-Console にアクセスするのに使用するトークン */
	AccountToken         *string	`json:"accountToken"`
}

func (p *SignInResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.AccountToken != nil {
        data["accountToken"] = p.AccountToken
    }
    return &data
}

type SignInAsyncResult struct {
	result *SignInResult
	err    error
}

type IssueAccountTokenResult struct {
    /** GS2-Console にアクセスするのに使用するトークン */
	AccountToken         *string	`json:"accountToken"`
}

func (p *IssueAccountTokenResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.AccountToken != nil {
        data["accountToken"] = p.AccountToken
    }
    return &data
}

type IssueAccountTokenAsyncResult struct {
	result *IssueAccountTokenResult
	err    error
}

type ForgetResult struct {
    /** パスワードを再発行するために必要なトークン */
	IssuePasswordToken         *string	`json:"issuePasswordToken"`
}

func (p *ForgetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.IssuePasswordToken != nil {
        data["issuePasswordToken"] = p.IssuePasswordToken
    }
    return &data
}

type ForgetAsyncResult struct {
	result *ForgetResult
	err    error
}

type IssuePasswordResult struct {
    /** 新しいパスワード */
	NewPassword         *string	`json:"newPassword"`
}

func (p *IssuePasswordResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.NewPassword != nil {
        data["newPassword"] = p.NewPassword
    }
    return &data
}

type IssuePasswordAsyncResult struct {
	result *IssuePasswordResult
	err    error
}

type UpdateAccountResult struct {
    /** 更新したGS2アカウント */
	Item         *Account	`json:"item"`
}

func (p *UpdateAccountResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateAccountAsyncResult struct {
	result *UpdateAccountResult
	err    error
}

type DeleteAccountResult struct {
}

func (p *DeleteAccountResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteAccountAsyncResult struct {
	result *DeleteAccountResult
	err    error
}

type DescribeProjectsResult struct {
    /** プロジェクトのリスト */
	Items         []Project	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeProjectsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Project, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeProjectsAsyncResult struct {
	result *DescribeProjectsResult
	err    error
}

type CreateProjectResult struct {
    /** 作成したプロジェクト */
	Item         *Project	`json:"item"`
}

func (p *CreateProjectResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateProjectAsyncResult struct {
	result *CreateProjectResult
	err    error
}

type GetProjectResult struct {
    /** プロジェクト */
	Item         *Project	`json:"item"`
}

func (p *GetProjectResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetProjectAsyncResult struct {
	result *GetProjectResult
	err    error
}

type GetProjectTokenResult struct {
    /** サインインしたプロジェクト */
	Item         *Project	`json:"item"`
    /** オーナーID */
	OwnerId         *string	`json:"ownerId"`
    /** プロジェクトトークン */
	ProjectToken         *string	`json:"projectToken"`
}

func (p *GetProjectTokenResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.OwnerId != nil {
        data["ownerId"] = p.OwnerId
    }
    if p.ProjectToken != nil {
        data["projectToken"] = p.ProjectToken
    }
    return &data
}

type GetProjectTokenAsyncResult struct {
	result *GetProjectTokenResult
	err    error
}

type GetProjectTokenByIdentifierResult struct {
    /** サインインしたプロジェクト */
	Item         *Project	`json:"item"`
    /** オーナーID */
	OwnerId         *string	`json:"ownerId"`
    /** プロジェクトトークン */
	ProjectToken         *string	`json:"projectToken"`
}

func (p *GetProjectTokenByIdentifierResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.OwnerId != nil {
        data["ownerId"] = p.OwnerId
    }
    if p.ProjectToken != nil {
        data["projectToken"] = p.ProjectToken
    }
    return &data
}

type GetProjectTokenByIdentifierAsyncResult struct {
	result *GetProjectTokenByIdentifierResult
	err    error
}

type UpdateProjectResult struct {
    /** 更新したプロジェクト */
	Item         *Project	`json:"item"`
}

func (p *UpdateProjectResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateProjectAsyncResult struct {
	result *UpdateProjectResult
	err    error
}

type DeleteProjectResult struct {
    /** 削除したプロジェクト */
	Item         *Project	`json:"item"`
}

func (p *DeleteProjectResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteProjectAsyncResult struct {
	result *DeleteProjectResult
	err    error
}

type DescribeBillingMethodsResult struct {
    /** 支払い方法のリスト */
	Items         []BillingMethod	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeBillingMethodsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]BillingMethod, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeBillingMethodsAsyncResult struct {
	result *DescribeBillingMethodsResult
	err    error
}

type CreateBillingMethodResult struct {
    /** 作成した支払い方法 */
	Item         *BillingMethod	`json:"item"`
}

func (p *CreateBillingMethodResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateBillingMethodAsyncResult struct {
	result *CreateBillingMethodResult
	err    error
}

type GetBillingMethodResult struct {
    /** 支払い方法 */
	Item         *BillingMethod	`json:"item"`
}

func (p *GetBillingMethodResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetBillingMethodAsyncResult struct {
	result *GetBillingMethodResult
	err    error
}

type UpdateBillingMethodResult struct {
    /** 更新した支払い方法 */
	Item         *BillingMethod	`json:"item"`
}

func (p *UpdateBillingMethodResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateBillingMethodAsyncResult struct {
	result *UpdateBillingMethodResult
	err    error
}

type DeleteBillingMethodResult struct {
    /** 削除した支払い方法 */
	Item         *BillingMethod	`json:"item"`
}

func (p *DeleteBillingMethodResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteBillingMethodAsyncResult struct {
	result *DeleteBillingMethodResult
	err    error
}

type DescribeReceiptsResult struct {
    /** 領収書のリスト */
	Items         []Receipt	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeReceiptsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Receipt, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeReceiptsAsyncResult struct {
	result *DescribeReceiptsResult
	err    error
}

type DescribeBillingsResult struct {
    /** 利用状況のリスト */
	Items         []Billing	`json:"items"`
}

func (p *DescribeBillingsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Billing, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeBillingsAsyncResult struct {
	result *DescribeBillingsResult
	err    error
}
