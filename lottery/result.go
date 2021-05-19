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

package lottery

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

type DescribeLotteryModelMastersResult struct {
    /** 抽選の種類マスターのリスト */
	Items         []LotteryModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeLotteryModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]LotteryModelMaster, 0)
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

type DescribeLotteryModelMastersAsyncResult struct {
	result *DescribeLotteryModelMastersResult
	err    error
}

type CreateLotteryModelMasterResult struct {
    /** 作成した抽選の種類マスター */
	Item         *LotteryModelMaster	`json:"item"`
}

func (p *CreateLotteryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateLotteryModelMasterAsyncResult struct {
	result *CreateLotteryModelMasterResult
	err    error
}

type GetLotteryModelMasterResult struct {
    /** 抽選の種類マスター */
	Item         *LotteryModelMaster	`json:"item"`
}

func (p *GetLotteryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetLotteryModelMasterAsyncResult struct {
	result *GetLotteryModelMasterResult
	err    error
}

type UpdateLotteryModelMasterResult struct {
    /** 更新した抽選の種類マスター */
	Item         *LotteryModelMaster	`json:"item"`
}

func (p *UpdateLotteryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateLotteryModelMasterAsyncResult struct {
	result *UpdateLotteryModelMasterResult
	err    error
}

type DeleteLotteryModelMasterResult struct {
    /** 削除した抽選の種類マスター */
	Item         *LotteryModelMaster	`json:"item"`
}

func (p *DeleteLotteryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteLotteryModelMasterAsyncResult struct {
	result *DeleteLotteryModelMasterResult
	err    error
}

type DescribePrizeTableMastersResult struct {
    /** 排出確率テーブルマスターのリスト */
	Items         []PrizeTableMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribePrizeTableMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]PrizeTableMaster, 0)
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

type DescribePrizeTableMastersAsyncResult struct {
	result *DescribePrizeTableMastersResult
	err    error
}

type CreatePrizeTableMasterResult struct {
    /** 作成した排出確率テーブルマスター */
	Item         *PrizeTableMaster	`json:"item"`
}

func (p *CreatePrizeTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreatePrizeTableMasterAsyncResult struct {
	result *CreatePrizeTableMasterResult
	err    error
}

type GetPrizeTableMasterResult struct {
    /** 排出確率テーブルマスター */
	Item         *PrizeTableMaster	`json:"item"`
}

func (p *GetPrizeTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetPrizeTableMasterAsyncResult struct {
	result *GetPrizeTableMasterResult
	err    error
}

type UpdatePrizeTableMasterResult struct {
    /** 更新した排出確率テーブルマスター */
	Item         *PrizeTableMaster	`json:"item"`
}

func (p *UpdatePrizeTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdatePrizeTableMasterAsyncResult struct {
	result *UpdatePrizeTableMasterResult
	err    error
}

type DeletePrizeTableMasterResult struct {
    /** 削除した排出確率テーブルマスター */
	Item         *PrizeTableMaster	`json:"item"`
}

func (p *DeletePrizeTableMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeletePrizeTableMasterAsyncResult struct {
	result *DeletePrizeTableMasterResult
	err    error
}

type DescribeBoxesResult struct {
    /** ボックスのリスト */
	Items         []Box	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeBoxesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Box, 0)
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

type DescribeBoxesAsyncResult struct {
	result *DescribeBoxesResult
	err    error
}

type DescribeBoxesByUserIdResult struct {
    /** ボックスのリスト */
	Items         []Box	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeBoxesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Box, 0)
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

type DescribeBoxesByUserIdAsyncResult struct {
	result *DescribeBoxesByUserIdResult
	err    error
}

type GetBoxResult struct {
    /** ボックスから取り出したアイテムのリスト */
	Item         *BoxItems	`json:"item"`
}

func (p *GetBoxResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetBoxAsyncResult struct {
	result *GetBoxResult
	err    error
}

type GetBoxByUserIdResult struct {
    /** ボックスから取り出したアイテムのリスト */
	Item         *BoxItems	`json:"item"`
}

func (p *GetBoxByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetBoxByUserIdAsyncResult struct {
	result *GetBoxByUserIdResult
	err    error
}

type GetRawBoxByUserIdResult struct {
    /** ボックス */
	Item         *Box	`json:"item"`
}

func (p *GetRawBoxByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRawBoxByUserIdAsyncResult struct {
	result *GetRawBoxByUserIdResult
	err    error
}

type ResetBoxResult struct {
}

func (p *ResetBoxResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type ResetBoxAsyncResult struct {
	result *ResetBoxResult
	err    error
}

type ResetBoxByUserIdResult struct {
}

func (p *ResetBoxByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type ResetBoxByUserIdAsyncResult struct {
	result *ResetBoxByUserIdResult
	err    error
}

type DescribeLotteryModelsResult struct {
    /** 抽選の種類のリスト */
	Items         []LotteryModel	`json:"items"`
}

func (p *DescribeLotteryModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]LotteryModel, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeLotteryModelsAsyncResult struct {
	result *DescribeLotteryModelsResult
	err    error
}

type GetLotteryModelResult struct {
    /** 抽選の種類 */
	Item         *LotteryModel	`json:"item"`
}

func (p *GetLotteryModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetLotteryModelAsyncResult struct {
	result *GetLotteryModelResult
	err    error
}

type DescribePrizeTablesResult struct {
    /** 排出確率テーブルのリスト */
	Items         []PrizeTable	`json:"items"`
}

func (p *DescribePrizeTablesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]PrizeTable, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribePrizeTablesAsyncResult struct {
	result *DescribePrizeTablesResult
	err    error
}

type GetPrizeTableResult struct {
    /** 排出確率テーブル */
	Item         *PrizeTable	`json:"item"`
}

func (p *GetPrizeTableResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetPrizeTableAsyncResult struct {
	result *GetPrizeTableResult
	err    error
}

type DrawByUserIdResult struct {
    /** 抽選結果の景品リスト */
	Items         []DrawnPrize	`json:"items"`
    /** 排出された景品を入手するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
    /** ボックスから取り出したアイテムのリスト */
	BoxItems         *BoxItems	`json:"boxItems"`
}

func (p *DrawByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]DrawnPrize, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    if p.BoxItems != nil {
        data["boxItems"] = p.BoxItems.ToDict()
    }
    return &data
}

type DrawByUserIdAsyncResult struct {
	result *DrawByUserIdResult
	err    error
}

type DescribeProbabilitiesResult struct {
    /** 景品の当選確率リスト */
	Items         []Probability	`json:"items"`
}

func (p *DescribeProbabilitiesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Probability, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeProbabilitiesAsyncResult struct {
	result *DescribeProbabilitiesResult
	err    error
}

type DescribeProbabilitiesByUserIdResult struct {
    /** 景品の当選確率リスト */
	Items         []Probability	`json:"items"`
}

func (p *DescribeProbabilitiesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Probability, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeProbabilitiesByUserIdAsyncResult struct {
	result *DescribeProbabilitiesByUserIdResult
	err    error
}

type DrawByStampSheetResult struct {
    /** 抽選結果の景品リスト */
	Items         []DrawnPrize	`json:"items"`
    /** 排出された景品を入手するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
    /** ボックスから取り出したアイテムのリスト */
	BoxItems         *BoxItems	`json:"boxItems"`
}

func (p *DrawByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]DrawnPrize, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.StampSheet != nil {
        data["stampSheet"] = p.StampSheet
    }
    if p.StampSheetEncryptionKeyId != nil {
        data["stampSheetEncryptionKeyId"] = p.StampSheetEncryptionKeyId
    }
    if p.BoxItems != nil {
        data["boxItems"] = p.BoxItems.ToDict()
    }
    return &data
}

type DrawByStampSheetAsyncResult struct {
	result *DrawByStampSheetResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な抽選設定 */
	Item         *CurrentLotteryMaster	`json:"item"`
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

type GetCurrentLotteryMasterResult struct {
    /** 現在有効な抽選設定 */
	Item         *CurrentLotteryMaster	`json:"item"`
}

func (p *GetCurrentLotteryMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentLotteryMasterAsyncResult struct {
	result *GetCurrentLotteryMasterResult
	err    error
}

type UpdateCurrentLotteryMasterResult struct {
    /** 更新した現在有効な抽選設定 */
	Item         *CurrentLotteryMaster	`json:"item"`
}

func (p *UpdateCurrentLotteryMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentLotteryMasterAsyncResult struct {
	result *UpdateCurrentLotteryMasterResult
	err    error
}

type UpdateCurrentLotteryMasterFromGitHubResult struct {
    /** 更新した現在有効な抽選設定 */
	Item         *CurrentLotteryMaster	`json:"item"`
}

func (p *UpdateCurrentLotteryMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentLotteryMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentLotteryMasterFromGitHubResult
	err    error
}
