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

package exchange

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
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
	Status         *string	`json:"status"`
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
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeRateModelsResult struct {
    /** 交換レートモデルのリスト */
	Items         []RateModel	`json:"items"`
}

func (p *DescribeRateModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]RateModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeRateModelsAsyncResult struct {
	result *DescribeRateModelsResult
	err    error
}

type GetRateModelResult struct {
    /** 交換レートモデル */
	Item         *RateModel	`json:"item"`
}

func (p *GetRateModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRateModelAsyncResult struct {
	result *GetRateModelResult
	err    error
}

type DescribeRateModelMastersResult struct {
    /** 交換レートマスターのリスト */
	Items         []RateModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeRateModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]RateModelMaster, 0)
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

type DescribeRateModelMastersAsyncResult struct {
	result *DescribeRateModelMastersResult
	err    error
}

type CreateRateModelMasterResult struct {
    /** 作成した交換レートマスター */
	Item         *RateModelMaster	`json:"item"`
}

func (p *CreateRateModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRateModelMasterAsyncResult struct {
	result *CreateRateModelMasterResult
	err    error
}

type GetRateModelMasterResult struct {
    /** 交換レートマスター */
	Item         *RateModelMaster	`json:"item"`
}

func (p *GetRateModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRateModelMasterAsyncResult struct {
	result *GetRateModelMasterResult
	err    error
}

type UpdateRateModelMasterResult struct {
    /** 更新した交換レートマスター */
	Item         *RateModelMaster	`json:"item"`
}

func (p *UpdateRateModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateRateModelMasterAsyncResult struct {
	result *UpdateRateModelMasterResult
	err    error
}

type DeleteRateModelMasterResult struct {
    /** 削除した交換レートマスター */
	Item         *RateModelMaster	`json:"item"`
}

func (p *DeleteRateModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRateModelMasterAsyncResult struct {
	result *DeleteRateModelMasterResult
	err    error
}

type ExchangeResult struct {
    /** 交換レートモデル */
	Item         *RateModel	`json:"item"`
    /** 交換処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *ExchangeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type ExchangeAsyncResult struct {
	result *ExchangeResult
	err    error
}

type ExchangeByUserIdResult struct {
    /** 交換レートモデル */
	Item         *RateModel	`json:"item"`
    /** 交換処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *ExchangeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type ExchangeByUserIdAsyncResult struct {
	result *ExchangeByUserIdResult
	err    error
}

type ExchangeByStampSheetResult struct {
    /** 交換レートモデル */
	Item         *RateModel	`json:"item"`
    /** 交換処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *ExchangeByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type ExchangeByStampSheetAsyncResult struct {
	result *ExchangeByStampSheetResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な交換レート設定 */
	Item         *CurrentRateMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentRateMasterResult struct {
    /** 現在有効な交換レート設定 */
	Item         *CurrentRateMaster	`json:"item"`
}

func (p *GetCurrentRateMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentRateMasterAsyncResult struct {
	result *GetCurrentRateMasterResult
	err    error
}

type UpdateCurrentRateMasterResult struct {
    /** 更新した現在有効な交換レート設定 */
	Item         *CurrentRateMaster	`json:"item"`
}

func (p *UpdateCurrentRateMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRateMasterAsyncResult struct {
	result *UpdateCurrentRateMasterResult
	err    error
}

type UpdateCurrentRateMasterFromGitHubResult struct {
    /** 更新した現在有効な交換レート設定 */
	Item         *CurrentRateMaster	`json:"item"`
}

func (p *UpdateCurrentRateMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentRateMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentRateMasterFromGitHubResult
	err    error
}

type CreateAwaitByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 取得できるようになる日時 */
	UnlockAt         *int64	`json:"unlockAt"`
}

func (p *CreateAwaitByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.UnlockAt != nil {
        data["unlockAt"] = p.UnlockAt
    }
    return &data
}

type CreateAwaitByUserIdAsyncResult struct {
	result *CreateAwaitByUserIdResult
	err    error
}

type DescribeAwaitsResult struct {
    /** 交換待機のリスト */
	Items         []Await	`json:"items"`
    /** 次のページを取得するためのトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeAwaitsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Await, 0)
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

type DescribeAwaitsAsyncResult struct {
	result *DescribeAwaitsResult
	err    error
}

type DescribeAwaitsByUserIdResult struct {
    /** 交換待機のリスト */
	Items         []Await	`json:"items"`
    /** 次のページを取得するためのトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeAwaitsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Await, 0)
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

type DescribeAwaitsByUserIdAsyncResult struct {
	result *DescribeAwaitsByUserIdResult
	err    error
}

type GetAwaitResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
}

func (p *GetAwaitResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetAwaitAsyncResult struct {
	result *GetAwaitResult
	err    error
}

type GetAwaitByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
}

func (p *GetAwaitByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetAwaitByUserIdAsyncResult struct {
	result *GetAwaitByUserIdResult
	err    error
}

type AcquireResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 報酬取得処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *AcquireResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type AcquireAsyncResult struct {
	result *AcquireResult
	err    error
}

type AcquireByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 報酬取得処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *AcquireByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type AcquireByUserIdAsyncResult struct {
	result *AcquireByUserIdResult
	err    error
}

type AcquireForceByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 報酬取得処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *AcquireForceByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type AcquireForceByUserIdAsyncResult struct {
	result *AcquireForceByUserIdResult
	err    error
}

type SkipResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 報酬取得処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *SkipResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type SkipAsyncResult struct {
	result *SkipResult
	err    error
}

type SkipByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** 報酬取得処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *SkipByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    return &data
}

type SkipByUserIdAsyncResult struct {
	result *SkipByUserIdResult
	err    error
}

type DeleteAwaitResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
}

func (p *DeleteAwaitResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteAwaitAsyncResult struct {
	result *DeleteAwaitResult
	err    error
}

type DeleteAwaitByUserIdResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
}

func (p *DeleteAwaitByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteAwaitByUserIdAsyncResult struct {
	result *DeleteAwaitByUserIdResult
	err    error
}

type CreateAwaitByStampSheetResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
}

func (p *CreateAwaitByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateAwaitByStampSheetAsyncResult struct {
	result *CreateAwaitByStampSheetResult
	err    error
}

type DeleteAwaitByStampTaskResult struct {
    /** 交換待機 */
	Item         *Await	`json:"item"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *string	`json:"newContextStack"`
}

func (p *DeleteAwaitByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    if p.NewContextStack != nil {
        data["newContextStack"] = p.NewContextStack
    }
    return &data
}

type DeleteAwaitByStampTaskAsyncResult struct {
	result *DeleteAwaitByStampTaskResult
	err    error
}
