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

package money

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Namespace, 0)
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
        data["item"] = p.Item.ToDict()
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
        data["status"] = p.Status
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
        data["item"] = p.Item.ToDict()
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
        data["item"] = p.Item.ToDict()
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

type DescribeWalletsResult struct {
    /** ウォレットのリスト */
	Items         []Wallet	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeWalletsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Wallet, 0)
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

type DescribeWalletsAsyncResult struct {
	result *DescribeWalletsResult
	err    error
}

type DescribeWalletsByUserIdResult struct {
    /** ウォレットのリスト */
	Items         []Wallet	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeWalletsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Wallet, 0)
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

type DescribeWalletsByUserIdAsyncResult struct {
	result *DescribeWalletsByUserIdResult
	err    error
}

type QueryWalletsResult struct {
    /** ウォレットのリスト */
	Items         []Wallet	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *QueryWalletsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Wallet, 0)
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

type QueryWalletsAsyncResult struct {
	result *QueryWalletsResult
	err    error
}

type GetWalletResult struct {
    /** ウォレット */
	Item         *Wallet	`json:"item"`
}

func (p *GetWalletResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetWalletAsyncResult struct {
	result *GetWalletResult
	err    error
}

type GetWalletByUserIdResult struct {
    /** ウォレット */
	Item         *Wallet	`json:"item"`
}

func (p *GetWalletByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetWalletByUserIdAsyncResult struct {
	result *GetWalletByUserIdResult
	err    error
}

type DepositByUserIdResult struct {
    /** 加算後のウォレット */
	Item         *Wallet	`json:"item"`
}

func (p *DepositByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DepositByUserIdAsyncResult struct {
	result *DepositByUserIdResult
	err    error
}

type WithdrawResult struct {
    /** 消費後のウォレット */
	Item         *Wallet	`json:"item"`
    /** 消費した通貨の価格 */
	Price         *float32	`json:"price"`
}

func (p *WithdrawResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Price != nil {
        data["price"] = p.Price
    }
    return &data
}

type WithdrawAsyncResult struct {
	result *WithdrawResult
	err    error
}

type WithdrawByUserIdResult struct {
    /** 消費後のウォレット */
	Item         *Wallet	`json:"item"`
    /** 消費した通貨の価格 */
	Price         *float32	`json:"price"`
}

func (p *WithdrawByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.Price != nil {
        data["price"] = p.Price
    }
    return &data
}

type WithdrawByUserIdAsyncResult struct {
	result *WithdrawByUserIdResult
	err    error
}

type DepositByStampSheetResult struct {
    /** 加算後のウォレット */
	Item         *Wallet	`json:"item"`
}

func (p *DepositByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DepositByStampSheetAsyncResult struct {
	result *DepositByStampSheetResult
	err    error
}

type WithdrawByStampTaskResult struct {
    /** 消費後のウォレット */
	Item         *Wallet	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *WithdrawByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type WithdrawByStampTaskAsyncResult struct {
	result *WithdrawByStampTaskResult
	err    error
}

type DescribeReceiptsResult struct {
    /** レシートのリスト */
	Items         []Receipt	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
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

type GetByUserIdAndTransactionIdResult struct {
    /** レシート */
	Item         *Receipt	`json:"item"`
}

func (p *GetByUserIdAndTransactionIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetByUserIdAndTransactionIdAsyncResult struct {
	result *GetByUserIdAndTransactionIdResult
	err    error
}

type RecordReceiptResult struct {
    /** レシート */
	Item         *Receipt	`json:"item"`
}

func (p *RecordReceiptResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type RecordReceiptAsyncResult struct {
	result *RecordReceiptResult
	err    error
}

type RecordReceiptByStampTaskResult struct {
    /** レシート */
	Item         *Receipt	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *RecordReceiptByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type RecordReceiptByStampTaskAsyncResult struct {
	result *RecordReceiptByStampTaskResult
	err    error
}
