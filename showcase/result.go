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

package showcase

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

type DescribeSalesItemMastersResult struct {
    /** 商品マスターのリスト */
	Items         []SalesItemMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeSalesItemMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]SalesItemMaster, 0)
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

type DescribeSalesItemMastersAsyncResult struct {
	result *DescribeSalesItemMastersResult
	err    error
}

type CreateSalesItemMasterResult struct {
    /** 作成した商品マスター */
	Item         *SalesItemMaster	`json:"item"`
}

func (p *CreateSalesItemMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateSalesItemMasterAsyncResult struct {
	result *CreateSalesItemMasterResult
	err    error
}

type GetSalesItemMasterResult struct {
    /** 商品マスター */
	Item         *SalesItemMaster	`json:"item"`
}

func (p *GetSalesItemMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSalesItemMasterAsyncResult struct {
	result *GetSalesItemMasterResult
	err    error
}

type UpdateSalesItemMasterResult struct {
    /** 更新した商品マスター */
	Item         *SalesItemMaster	`json:"item"`
}

func (p *UpdateSalesItemMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateSalesItemMasterAsyncResult struct {
	result *UpdateSalesItemMasterResult
	err    error
}

type DeleteSalesItemMasterResult struct {
    /** 削除した商品マスター */
	Item         *SalesItemMaster	`json:"item"`
}

func (p *DeleteSalesItemMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteSalesItemMasterAsyncResult struct {
	result *DeleteSalesItemMasterResult
	err    error
}

type DescribeSalesItemGroupMastersResult struct {
    /** 商品グループマスターのリスト */
	Items         []SalesItemGroupMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeSalesItemGroupMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]SalesItemGroupMaster, 0)
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

type DescribeSalesItemGroupMastersAsyncResult struct {
	result *DescribeSalesItemGroupMastersResult
	err    error
}

type CreateSalesItemGroupMasterResult struct {
    /** 作成した商品グループマスター */
	Item         *SalesItemGroupMaster	`json:"item"`
}

func (p *CreateSalesItemGroupMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateSalesItemGroupMasterAsyncResult struct {
	result *CreateSalesItemGroupMasterResult
	err    error
}

type GetSalesItemGroupMasterResult struct {
    /** 商品グループマスター */
	Item         *SalesItemGroupMaster	`json:"item"`
}

func (p *GetSalesItemGroupMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSalesItemGroupMasterAsyncResult struct {
	result *GetSalesItemGroupMasterResult
	err    error
}

type UpdateSalesItemGroupMasterResult struct {
    /** 更新した商品グループマスター */
	Item         *SalesItemGroupMaster	`json:"item"`
}

func (p *UpdateSalesItemGroupMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateSalesItemGroupMasterAsyncResult struct {
	result *UpdateSalesItemGroupMasterResult
	err    error
}

type DeleteSalesItemGroupMasterResult struct {
    /** 削除した商品グループマスター */
	Item         *SalesItemGroupMaster	`json:"item"`
}

func (p *DeleteSalesItemGroupMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteSalesItemGroupMasterAsyncResult struct {
	result *DeleteSalesItemGroupMasterResult
	err    error
}

type DescribeShowcaseMastersResult struct {
    /** 陳列棚マスターのリスト */
	Items         []ShowcaseMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeShowcaseMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]ShowcaseMaster, 0)
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

type DescribeShowcaseMastersAsyncResult struct {
	result *DescribeShowcaseMastersResult
	err    error
}

type CreateShowcaseMasterResult struct {
    /** 作成した陳列棚マスター */
	Item         *ShowcaseMaster	`json:"item"`
}

func (p *CreateShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateShowcaseMasterAsyncResult struct {
	result *CreateShowcaseMasterResult
	err    error
}

type GetShowcaseMasterResult struct {
    /** 陳列棚マスター */
	Item         *ShowcaseMaster	`json:"item"`
}

func (p *GetShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetShowcaseMasterAsyncResult struct {
	result *GetShowcaseMasterResult
	err    error
}

type UpdateShowcaseMasterResult struct {
    /** 更新した陳列棚マスター */
	Item         *ShowcaseMaster	`json:"item"`
}

func (p *UpdateShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateShowcaseMasterAsyncResult struct {
	result *UpdateShowcaseMasterResult
	err    error
}

type DeleteShowcaseMasterResult struct {
    /** 削除した陳列棚マスター */
	Item         *ShowcaseMaster	`json:"item"`
}

func (p *DeleteShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteShowcaseMasterAsyncResult struct {
	result *DeleteShowcaseMasterResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な陳列棚マスター */
	Item         *CurrentShowcaseMaster	`json:"item"`
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

type GetCurrentShowcaseMasterResult struct {
    /** 現在有効な陳列棚マスター */
	Item         *CurrentShowcaseMaster	`json:"item"`
}

func (p *GetCurrentShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentShowcaseMasterAsyncResult struct {
	result *GetCurrentShowcaseMasterResult
	err    error
}

type UpdateCurrentShowcaseMasterResult struct {
    /** 更新した現在有効な陳列棚マスター */
	Item         *CurrentShowcaseMaster	`json:"item"`
}

func (p *UpdateCurrentShowcaseMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentShowcaseMasterAsyncResult struct {
	result *UpdateCurrentShowcaseMasterResult
	err    error
}

type UpdateCurrentShowcaseMasterFromGitHubResult struct {
    /** 更新した現在有効な陳列棚マスター */
	Item         *CurrentShowcaseMaster	`json:"item"`
}

func (p *UpdateCurrentShowcaseMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentShowcaseMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentShowcaseMasterFromGitHubResult
	err    error
}

type DescribeShowcasesResult struct {
    /** 陳列棚のリスト */
	Items         []Showcase	`json:"items"`
}

func (p *DescribeShowcasesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Showcase, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeShowcasesAsyncResult struct {
	result *DescribeShowcasesResult
	err    error
}

type DescribeShowcasesByUserIdResult struct {
    /** 陳列棚のリスト */
	Items         []Showcase	`json:"items"`
}

func (p *DescribeShowcasesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Showcase, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeShowcasesByUserIdAsyncResult struct {
	result *DescribeShowcasesByUserIdResult
	err    error
}

type GetShowcaseResult struct {
    /** 陳列棚 */
	Item         *Showcase	`json:"item"`
}

func (p *GetShowcaseResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetShowcaseAsyncResult struct {
	result *GetShowcaseResult
	err    error
}

type GetShowcaseByUserIdResult struct {
    /** 陳列棚 */
	Item         *Showcase	`json:"item"`
}

func (p *GetShowcaseByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetShowcaseByUserIdAsyncResult struct {
	result *GetShowcaseByUserIdResult
	err    error
}

type BuyResult struct {
    /** 商品 */
	Item         *SalesItem	`json:"item"`
    /** 購入処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *BuyResult) ToDict() *map[string]interface{} {
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

type BuyAsyncResult struct {
	result *BuyResult
	err    error
}

type BuyByUserIdResult struct {
    /** 商品 */
	Item         *SalesItem	`json:"item"`
    /** 購入処理の実行に使用するスタンプシート */
	StampSheet         *string	`json:"stampSheet"`
    /** スタンプシートの署名計算に使用した暗号鍵GRN */
	StampSheetEncryptionKeyId         *string	`json:"stampSheetEncryptionKeyId"`
}

func (p *BuyByUserIdResult) ToDict() *map[string]interface{} {
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

type BuyByUserIdAsyncResult struct {
	result *BuyByUserIdResult
	err    error
}
