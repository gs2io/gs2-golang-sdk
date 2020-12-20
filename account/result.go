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

package account

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         *[]*Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
    /** 作成したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
    /** None */
	Status         *core.String	`json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
     data["Status"] = p.Status
    }
    return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
    /** ネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
    /** 更新したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeAccountsResult struct {
    /** ゲームプレイヤーアカウントのリスト */
	Items         *[]*Account	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeAccountsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeAccountsAsyncResult struct {
	result *DescribeAccountsResult
	err    error
}

type CreateAccountResult struct {
    /** 作成したゲームプレイヤーアカウント */
	Item         *Account	`json:"item"`
}

func (p *CreateAccountResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateAccountAsyncResult struct {
	result *CreateAccountResult
	err    error
}

type UpdateTimeOffsetResult struct {
    /** 更新したゲームプレイヤーアカウント */
	Item         *Account	`json:"item"`
}

func (p *UpdateTimeOffsetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateTimeOffsetAsyncResult struct {
	result *UpdateTimeOffsetResult
	err    error
}

type GetAccountResult struct {
    /** ゲームプレイヤーアカウント */
	Item         *Account	`json:"item"`
}

func (p *GetAccountResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetAccountAsyncResult struct {
	result *GetAccountResult
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

type AuthenticationResult struct {
    /** ゲームプレイヤーアカウント */
	Item         *Account	`json:"item"`
    /** 署名対象のアカウント情報 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *AuthenticationResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    if p.Body != nil {
     data["Body"] = p.Body
    }
    if p.Signature != nil {
     data["Signature"] = p.Signature
    }
    return &data
}

type AuthenticationAsyncResult struct {
	result *AuthenticationResult
	err    error
}

type DescribeTakeOversResult struct {
    /** 引き継ぎ設定のリスト */
	Items         *[]*TakeOver	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeTakeOversResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeTakeOversAsyncResult struct {
	result *DescribeTakeOversResult
	err    error
}

type DescribeTakeOversByUserIdResult struct {
    /** 引き継ぎ設定のリスト */
	Items         *[]*TakeOver	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeTakeOversByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeTakeOversByUserIdAsyncResult struct {
	result *DescribeTakeOversByUserIdResult
	err    error
}

type CreateTakeOverResult struct {
    /** 作成した引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *CreateTakeOverResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateTakeOverAsyncResult struct {
	result *CreateTakeOverResult
	err    error
}

type CreateTakeOverByUserIdResult struct {
    /** 作成した引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *CreateTakeOverByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateTakeOverByUserIdAsyncResult struct {
	result *CreateTakeOverByUserIdResult
	err    error
}

type GetTakeOverResult struct {
    /** 引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *GetTakeOverResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetTakeOverAsyncResult struct {
	result *GetTakeOverResult
	err    error
}

type GetTakeOverByUserIdResult struct {
    /** 引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *GetTakeOverByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetTakeOverByUserIdAsyncResult struct {
	result *GetTakeOverByUserIdResult
	err    error
}

type UpdateTakeOverResult struct {
    /** 引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *UpdateTakeOverResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateTakeOverAsyncResult struct {
	result *UpdateTakeOverResult
	err    error
}

type UpdateTakeOverByUserIdResult struct {
    /** 引き継ぎ設定 */
	Item         *TakeOver	`json:"item"`
}

func (p *UpdateTakeOverByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateTakeOverByUserIdAsyncResult struct {
	result *UpdateTakeOverByUserIdResult
	err    error
}

type DeleteTakeOverResult struct {
}

func (p *DeleteTakeOverResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteTakeOverAsyncResult struct {
	result *DeleteTakeOverResult
	err    error
}

type DeleteTakeOverByUserIdentifierResult struct {
}

func (p *DeleteTakeOverByUserIdentifierResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteTakeOverByUserIdentifierAsyncResult struct {
	result *DeleteTakeOverByUserIdentifierResult
	err    error
}

type DoTakeOverResult struct {
    /** ゲームプレイヤーアカウント */
	Item         *Account	`json:"item"`
}

func (p *DoTakeOverResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DoTakeOverAsyncResult struct {
	result *DoTakeOverResult
	err    error
}
