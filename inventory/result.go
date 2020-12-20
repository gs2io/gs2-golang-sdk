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

package inventory

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
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeInventoryModelMastersResult struct {
    /** インベントリモデルマスターのリスト */
	Items         *[]*InventoryModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeInventoryModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeInventoryModelMastersAsyncResult struct {
	result *DescribeInventoryModelMastersResult
	err    error
}

type CreateInventoryModelMasterResult struct {
    /** 作成したインベントリモデルマスター */
	Item         *InventoryModelMaster	`json:"item"`
}

func (p *CreateInventoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateInventoryModelMasterAsyncResult struct {
	result *CreateInventoryModelMasterResult
	err    error
}

type GetInventoryModelMasterResult struct {
    /** インベントリモデルマスター */
	Item         *InventoryModelMaster	`json:"item"`
}

func (p *GetInventoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetInventoryModelMasterAsyncResult struct {
	result *GetInventoryModelMasterResult
	err    error
}

type UpdateInventoryModelMasterResult struct {
    /** 更新したインベントリモデルマスター */
	Item         *InventoryModelMaster	`json:"item"`
}

func (p *UpdateInventoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateInventoryModelMasterAsyncResult struct {
	result *UpdateInventoryModelMasterResult
	err    error
}

type DeleteInventoryModelMasterResult struct {
    /** 削除したインベントリモデルマスター */
	Item         *InventoryModelMaster	`json:"item"`
}

func (p *DeleteInventoryModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteInventoryModelMasterAsyncResult struct {
	result *DeleteInventoryModelMasterResult
	err    error
}

type DescribeInventoryModelsResult struct {
    /** インベントリモデルのリスト */
	Items         *[]*InventoryModel	`json:"items"`
}

func (p *DescribeInventoryModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeInventoryModelsAsyncResult struct {
	result *DescribeInventoryModelsResult
	err    error
}

type GetInventoryModelResult struct {
    /** インベントリモデル */
	Item         *InventoryModel	`json:"item"`
}

func (p *GetInventoryModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetInventoryModelAsyncResult struct {
	result *GetInventoryModelResult
	err    error
}

type DescribeItemModelMastersResult struct {
    /** アイテムモデルマスターのリスト */
	Items         *[]*ItemModelMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeItemModelMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeItemModelMastersAsyncResult struct {
	result *DescribeItemModelMastersResult
	err    error
}

type CreateItemModelMasterResult struct {
    /** 作成したアイテムモデルマスター */
	Item         *ItemModelMaster	`json:"item"`
}

func (p *CreateItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateItemModelMasterAsyncResult struct {
	result *CreateItemModelMasterResult
	err    error
}

type GetItemModelMasterResult struct {
    /** アイテムモデルマスター */
	Item         *ItemModelMaster	`json:"item"`
}

func (p *GetItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetItemModelMasterAsyncResult struct {
	result *GetItemModelMasterResult
	err    error
}

type UpdateItemModelMasterResult struct {
    /** 更新したアイテムモデルマスター */
	Item         *ItemModelMaster	`json:"item"`
}

func (p *UpdateItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateItemModelMasterAsyncResult struct {
	result *UpdateItemModelMasterResult
	err    error
}

type DeleteItemModelMasterResult struct {
    /** 削除したアイテムモデルマスター */
	Item         *ItemModelMaster	`json:"item"`
}

func (p *DeleteItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteItemModelMasterAsyncResult struct {
	result *DeleteItemModelMasterResult
	err    error
}

type DescribeItemModelsResult struct {
    /** アイテムモデルのリスト */
	Items         *[]*ItemModel	`json:"items"`
}

func (p *DescribeItemModelsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeItemModelsAsyncResult struct {
	result *DescribeItemModelsResult
	err    error
}

type GetItemModelResult struct {
    /** None */
	Item         *ItemModel	`json:"item"`
}

func (p *GetItemModelResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetItemModelAsyncResult struct {
	result *GetItemModelResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効な所持品マスター */
	Item         *CurrentItemModelMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentItemModelMasterResult struct {
    /** 現在有効な所持品マスター */
	Item         *CurrentItemModelMaster	`json:"item"`
}

func (p *GetCurrentItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentItemModelMasterAsyncResult struct {
	result *GetCurrentItemModelMasterResult
	err    error
}

type UpdateCurrentItemModelMasterResult struct {
    /** 更新した現在有効な所持品マスター */
	Item         *CurrentItemModelMaster	`json:"item"`
}

func (p *UpdateCurrentItemModelMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentItemModelMasterAsyncResult struct {
	result *UpdateCurrentItemModelMasterResult
	err    error
}

type UpdateCurrentItemModelMasterFromGitHubResult struct {
    /** 更新した現在有効な所持品マスター */
	Item         *CurrentItemModelMaster	`json:"item"`
}

func (p *UpdateCurrentItemModelMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentItemModelMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentItemModelMasterFromGitHubResult
	err    error
}

type DescribeInventoriesResult struct {
    /** インベントリのリスト */
	Items         *[]*Inventory	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeInventoriesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeInventoriesAsyncResult struct {
	result *DescribeInventoriesResult
	err    error
}

type DescribeInventoriesByUserIdResult struct {
    /** インベントリのリスト */
	Items         *[]*Inventory	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeInventoriesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeInventoriesByUserIdAsyncResult struct {
	result *DescribeInventoriesByUserIdResult
	err    error
}

type GetInventoryResult struct {
    /** インベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *GetInventoryResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetInventoryAsyncResult struct {
	result *GetInventoryResult
	err    error
}

type GetInventoryByUserIdResult struct {
    /** インベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *GetInventoryByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetInventoryByUserIdAsyncResult struct {
	result *GetInventoryByUserIdResult
	err    error
}

type AddCapacityByUserIdResult struct {
    /** キャパシティ加算後のインベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *AddCapacityByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type AddCapacityByUserIdAsyncResult struct {
	result *AddCapacityByUserIdResult
	err    error
}

type SetCapacityByUserIdResult struct {
    /** 更新後のインベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *SetCapacityByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type SetCapacityByUserIdAsyncResult struct {
	result *SetCapacityByUserIdResult
	err    error
}

type DeleteInventoryByUserIdResult struct {
    /** インベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *DeleteInventoryByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteInventoryByUserIdAsyncResult struct {
	result *DeleteInventoryByUserIdResult
	err    error
}

type AddCapacityByStampSheetResult struct {
    /** キャパシティ加算後のインベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *AddCapacityByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type AddCapacityByStampSheetAsyncResult struct {
	result *AddCapacityByStampSheetResult
	err    error
}

type SetCapacityByStampSheetResult struct {
    /** 更新後のインベントリ */
	Item         *Inventory	`json:"item"`
}

func (p *SetCapacityByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type SetCapacityByStampSheetAsyncResult struct {
	result *SetCapacityByStampSheetResult
	err    error
}

type DescribeItemSetsResult struct {
    /** 有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeItemSetsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeItemSetsAsyncResult struct {
	result *DescribeItemSetsResult
	err    error
}

type DescribeItemSetsByUserIdResult struct {
    /** 有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeItemSetsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeItemSetsByUserIdAsyncResult struct {
	result *DescribeItemSetsByUserIdResult
	err    error
}

type GetItemSetResult struct {
    /** 有効期限毎の{model_name} */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *GetItemSetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type GetItemSetAsyncResult struct {
	result *GetItemSetResult
	err    error
}

type GetItemSetByUserIdResult struct {
    /** 有効期限毎の{model_name} */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *GetItemSetByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type GetItemSetByUserIdAsyncResult struct {
	result *GetItemSetByUserIdResult
	err    error
}

type GetItemWithSignatureResult struct {
    /** 有効期限毎の{model_name} */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** 署名対象のアイテムセット情報 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *GetItemWithSignatureResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.Body != nil {
     data["Body"] = p.Body
    }
    if p.Signature != nil {
     data["Signature"] = p.Signature
    }
    return &data
}

type GetItemWithSignatureAsyncResult struct {
	result *GetItemWithSignatureResult
	err    error
}

type GetItemWithSignatureByUserIdResult struct {
    /** 有効期限毎の{model_name} */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** 署名対象のアイテムセット情報 */
	Body         *core.String	`json:"body"`
    /** 署名 */
	Signature         *core.String	`json:"signature"`
}

func (p *GetItemWithSignatureByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.Body != nil {
     data["Body"] = p.Body
    }
    if p.Signature != nil {
     data["Signature"] = p.Signature
    }
    return &data
}

type GetItemWithSignatureByUserIdAsyncResult struct {
	result *GetItemWithSignatureByUserIdResult
	err    error
}

type AcquireItemSetByUserIdResult struct {
    /** 加算後の有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** 所持数量の上限を超えて受け取れずに GS2-Inbox に転送したアイテムの数量 */
	OverflowCount         *int64	`json:"overflowCount"`
}

func (p *AcquireItemSetByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.OverflowCount != nil {
     data["OverflowCount"] = p.OverflowCount
    }
    return &data
}

type AcquireItemSetByUserIdAsyncResult struct {
	result *AcquireItemSetByUserIdResult
	err    error
}

type ConsumeItemSetResult struct {
    /** 消費後の有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *ConsumeItemSetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type ConsumeItemSetAsyncResult struct {
	result *ConsumeItemSetResult
	err    error
}

type ConsumeItemSetByUserIdResult struct {
    /** 消費後の有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *ConsumeItemSetByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type ConsumeItemSetByUserIdAsyncResult struct {
	result *ConsumeItemSetByUserIdResult
	err    error
}

type DescribeReferenceOfResult struct {
    /** この所持品の参照元リスト */
	Items         *[]core.String	`json:"items"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DescribeReferenceOfResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DescribeReferenceOfAsyncResult struct {
	result *DescribeReferenceOfResult
	err    error
}

type DescribeReferenceOfByUserIdResult struct {
    /** この所持品の参照元リスト */
	Items         *[]core.String	`json:"items"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DescribeReferenceOfByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DescribeReferenceOfByUserIdAsyncResult struct {
	result *DescribeReferenceOfByUserIdResult
	err    error
}

type GetReferenceOfResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *GetReferenceOfResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type GetReferenceOfAsyncResult struct {
	result *GetReferenceOfResult
	err    error
}

type GetReferenceOfByUserIdResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *GetReferenceOfByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type GetReferenceOfByUserIdAsyncResult struct {
	result *GetReferenceOfByUserIdResult
	err    error
}

type VerifyReferenceOfResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *VerifyReferenceOfResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type VerifyReferenceOfAsyncResult struct {
	result *VerifyReferenceOfResult
	err    error
}

type VerifyReferenceOfByUserIdResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *VerifyReferenceOfByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type VerifyReferenceOfByUserIdAsyncResult struct {
	result *VerifyReferenceOfByUserIdResult
	err    error
}

type AddReferenceOfResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元追加後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *AddReferenceOfResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type AddReferenceOfAsyncResult struct {
	result *AddReferenceOfResult
	err    error
}

type AddReferenceOfByUserIdResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元追加後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *AddReferenceOfByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type AddReferenceOfByUserIdAsyncResult struct {
	result *AddReferenceOfByUserIdResult
	err    error
}

type DeleteReferenceOfResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元削除後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DeleteReferenceOfResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DeleteReferenceOfAsyncResult struct {
	result *DeleteReferenceOfResult
	err    error
}

type DeleteReferenceOfByUserIdResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元削除後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DeleteReferenceOfByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DeleteReferenceOfByUserIdAsyncResult struct {
	result *DeleteReferenceOfByUserIdResult
	err    error
}

type DeleteItemSetByUserIdResult struct {
    /** 削除した有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DeleteItemSetByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DeleteItemSetByUserIdAsyncResult struct {
	result *DeleteItemSetByUserIdResult
	err    error
}

type AcquireItemSetByStampSheetResult struct {
    /** 加算後の有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** 所持数量の上限を超えて受け取れずに GS2-Inbox に転送したアイテムの数量 */
	OverflowCount         *int64	`json:"overflowCount"`
}

func (p *AcquireItemSetByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.OverflowCount != nil {
     data["OverflowCount"] = p.OverflowCount
    }
    return &data
}

type AcquireItemSetByStampSheetAsyncResult struct {
	result *AcquireItemSetByStampSheetResult
	err    error
}

type AddReferenceOfItemSetByStampSheetResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元追加後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *AddReferenceOfItemSetByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type AddReferenceOfItemSetByStampSheetAsyncResult struct {
	result *AddReferenceOfItemSetByStampSheetResult
	err    error
}

type DeleteReferenceOfItemSetByStampSheetResult struct {
    /** この所持品の参照元リスト */
	Item         *[]core.String	`json:"item"`
    /** 参照元削除後の有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
}

func (p *DeleteReferenceOfItemSetByStampSheetResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    return &data
}

type DeleteReferenceOfItemSetByStampSheetAsyncResult struct {
	result *DeleteReferenceOfItemSetByStampSheetResult
	err    error
}

type ConsumeItemSetByStampTaskResult struct {
    /** 消費後の有効期限ごとのアイテム所持数量のリスト */
	Items         *[]*ItemSet	`json:"items"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *ConsumeItemSetByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.NewContextStack != nil {
     data["NewContextStack"] = p.NewContextStack
    }
    return &data
}

type ConsumeItemSetByStampTaskAsyncResult struct {
	result *ConsumeItemSetByStampTaskResult
	err    error
}

type VerifyReferenceOfByStampTaskResult struct {
    /** この所持品の参照元のリスト */
	Item         *[]core.String	`json:"item"`
    /** 有効期限ごとのアイテム所持数量 */
	ItemSet         *ItemSet	`json:"itemSet"`
    /** アイテムモデル */
	ItemModel         *ItemModel	`json:"itemModel"`
    /** インベントリ */
	Inventory         *Inventory	`json:"inventory"`
    /** スタンプタスクの実行結果を記録したコンテキスト */
	NewContextStack         *core.String	`json:"newContextStack"`
}

func (p *VerifyReferenceOfByStampTaskResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
    }
    if p.ItemSet != nil {
     data["ItemSet"] = p.ItemSet.ToDict()
    }
    if p.ItemModel != nil {
     data["ItemModel"] = p.ItemModel.ToDict()
    }
    if p.Inventory != nil {
     data["Inventory"] = p.Inventory.ToDict()
    }
    if p.NewContextStack != nil {
     data["NewContextStack"] = p.NewContextStack
    }
    return &data
}

type VerifyReferenceOfByStampTaskAsyncResult struct {
	result *VerifyReferenceOfByStampTaskResult
	err    error
}
