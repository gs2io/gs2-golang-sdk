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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** アイテム入手したときに実行するスクリプト */
	AcquireScript *ScriptSetting   `json:"acquireScript"`
    /** 入手上限に当たって入手できなかったときに実行するスクリプト */
	OverflowScript *ScriptSetting   `json:"overflowScript"`
    /** アイテム消費するときに実行するスクリプト */
	ConsumeScript *ScriptSetting   `json:"consumeScript"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Namespace) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["description"] = p.Description
    if p.AcquireScript != nil {
        data["acquireScript"] = *p.AcquireScript.ToDict()
    }
    if p.OverflowScript != nil {
        data["overflowScript"] = *p.OverflowScript.ToDict()
    }
    if p.ConsumeScript != nil {
        data["consumeScript"] = *p.ConsumeScript.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type InventoryModelMaster struct {
    /** インベントリモデルマスター */
	InventoryModelId *string   `json:"inventoryModelId"`
    /** インベントリの種類名 */
	Name *string   `json:"name"`
    /** インベントリの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** インベントリモデルマスターの説明 */
	Description *string   `json:"description"`
    /** インベントリの初期サイズ */
	InitialCapacity *int32   `json:"initialCapacity"`
    /** インベントリの最大サイズ */
	MaxCapacity *int32   `json:"maxCapacity"`
    /** 参照元が登録されているアイテムセットは削除できなくする */
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *InventoryModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["inventoryModelId"] = p.InventoryModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["initialCapacity"] = p.InitialCapacity
    data["maxCapacity"] = p.MaxCapacity
    data["protectReferencedItem"] = p.ProtectReferencedItem
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type InventoryModel struct {
    /** インベントリモデル */
	InventoryModelId *string   `json:"inventoryModelId"`
    /** インベントリの種類名 */
	Name *string   `json:"name"`
    /** インベントリの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** インベントリの初期サイズ */
	InitialCapacity *int32   `json:"initialCapacity"`
    /** インベントリの最大サイズ */
	MaxCapacity *int32   `json:"maxCapacity"`
    /** 参照元が登録されているアイテムセットは削除できなくする */
	ProtectReferencedItem *bool   `json:"protectReferencedItem"`
    /** インベントリに格納可能なアイテムモデル一覧 */
	ItemModels []ItemModel   `json:"itemModels"`
}

func (p *InventoryModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["inventoryModelId"] = p.InventoryModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["initialCapacity"] = p.InitialCapacity
    data["maxCapacity"] = p.MaxCapacity
    data["protectReferencedItem"] = p.ProtectReferencedItem
    if p.ItemModels != nil {
        var _itemModels []*map[string]interface {}
        for _, item := range p.ItemModels {
            _itemModels = append(_itemModels, item.ToDict())
        }
        data["itemModels"] = &_itemModels
    }
    return &data
}

type ItemModelMaster struct {
    /** アイテムモデルマスター */
	ItemModelId *string   `json:"itemModelId"`
    /** アイテムの種類名 */
	InventoryName *string   `json:"inventoryName"`
    /** アイテムモデルの種類名 */
	Name *string   `json:"name"`
    /** アイテムモデルマスターの説明 */
	Description *string   `json:"description"`
    /** アイテムモデルの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** スタック可能な最大数量 */
	StackingLimit *int64   `json:"stackingLimit"`
    /** スタック可能な最大数量を超えた時複数枠にアイテムを保管することを許すか */
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
    /** 表示順番 */
	SortValue *int32   `json:"sortValue"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ItemModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["itemModelId"] = p.ItemModelId
    data["inventoryName"] = p.InventoryName
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["stackingLimit"] = p.StackingLimit
    data["allowMultipleStacks"] = p.AllowMultipleStacks
    data["sortValue"] = p.SortValue
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ItemModel struct {
    /** アイテムモデルマスター */
	ItemModelId *string   `json:"itemModelId"`
    /** アイテムモデルの種類名 */
	Name *string   `json:"name"`
    /** アイテムモデルの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** スタック可能な最大数量 */
	StackingLimit *int64   `json:"stackingLimit"`
    /** スタック可能な最大数量を超えた時複数枠にアイテムを保管することを許すか */
	AllowMultipleStacks *bool   `json:"allowMultipleStacks"`
    /** 表示順番 */
	SortValue *int32   `json:"sortValue"`
}

func (p *ItemModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["itemModelId"] = p.ItemModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["stackingLimit"] = p.StackingLimit
    data["allowMultipleStacks"] = p.AllowMultipleStacks
    data["sortValue"] = p.SortValue
    return &data
}

type CurrentItemModelMaster struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentItemModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
    return &data
}

type Inventory struct {
    /** インベントリ */
	InventoryId *string   `json:"inventoryId"`
    /** インベントリモデル名 */
	InventoryName *string   `json:"inventoryName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 現在のインベントリのキャパシティ使用量 */
	CurrentInventoryCapacityUsage *int32   `json:"currentInventoryCapacityUsage"`
    /** 現在のインベントリの最大キャパシティ */
	CurrentInventoryMaxCapacity *int32   `json:"currentInventoryMaxCapacity"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Inventory) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["inventoryId"] = p.InventoryId
    data["inventoryName"] = p.InventoryName
    data["userId"] = p.UserId
    data["currentInventoryCapacityUsage"] = p.CurrentInventoryCapacityUsage
    data["currentInventoryMaxCapacity"] = p.CurrentInventoryMaxCapacity
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ItemSet struct {
    /** 有効期限ごとのアイテム所持数量 */
	ItemSetId *string   `json:"itemSetId"`
    /** アイテムセットを識別する名前 */
	Name *string   `json:"name"`
    /** インベントリの名前 */
	InventoryName *string   `json:"inventoryName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** アイテムマスターの名前 */
	ItemName *string   `json:"itemName"`
    /** 所持数量 */
	Count *int64   `json:"count"`
    /** この所持品の参照元リスト */
	ReferenceOf []string   `json:"referenceOf"`
    /** 表示順番 */
	SortValue *int32   `json:"sortValue"`
    /** 有効期限 */
	ExpiresAt *int64   `json:"expiresAt"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ItemSet) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["itemSetId"] = p.ItemSetId
    data["name"] = p.Name
    data["inventoryName"] = p.InventoryName
    data["userId"] = p.UserId
    data["itemName"] = p.ItemName
    data["count"] = p.Count
    if p.ReferenceOf != nil {
        var _referenceOf []string
        for _, item := range p.ReferenceOf {
            _referenceOf = append(_referenceOf, item)
        }
        data["referenceOf"] = &_referenceOf
    }
    data["sortValue"] = p.SortValue
    data["expiresAt"] = p.ExpiresAt
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ItemSetGroup struct {
    /** 有効期限ごとのアイテム所持数量 (このモデルは SDK では使用されません) */
	ItemSetGroupId *string   `json:"itemSetGroupId"`
    /** インベントリの名前 */
	InventoryName *string   `json:"inventoryName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** アイテムマスターの名前 */
	ItemName *string   `json:"itemName"`
    /** 表示順番 */
	SortValue *int32   `json:"sortValue"`
    /** アイテムセットIDのリスト */
	ItemSetItemSetIdList []string   `json:"itemSetItemSetIdList"`
    /** アイテムセットを識別する名前のリスト */
	ItemSetNameList []string   `json:"itemSetNameList"`
    /** 所持数量のリスト */
	ItemSetCountList []int64   `json:"itemSetCountList"`
    /** 参照元のリストのリスト */
	ItemSetReferenceOfList [][]string   `json:"itemSetReferenceOfList"`
    /** 有効期限のリスト */
	ItemSetExpiresAtList []int64   `json:"itemSetExpiresAtList"`
    /** 作成日時のリスト */
	ItemSetCreatedAtList []int64   `json:"itemSetCreatedAtList"`
    /** 更新日時のリスト */
	ItemSetUpdatedAtList []int64   `json:"itemSetUpdatedAtList"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ItemSetGroup) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["itemSetGroupId"] = p.ItemSetGroupId
    data["inventoryName"] = p.InventoryName
    data["userId"] = p.UserId
    data["itemName"] = p.ItemName
    data["sortValue"] = p.SortValue
    if p.ItemSetItemSetIdList != nil {
        var _itemSetItemSetIdList []string
        for _, item := range p.ItemSetItemSetIdList {
            _itemSetItemSetIdList = append(_itemSetItemSetIdList, item)
        }
        data["itemSetItemSetIdList"] = &_itemSetItemSetIdList
    }
    if p.ItemSetNameList != nil {
        var _itemSetNameList []string
        for _, item := range p.ItemSetNameList {
            _itemSetNameList = append(_itemSetNameList, item)
        }
        data["itemSetNameList"] = &_itemSetNameList
    }
    if p.ItemSetCountList != nil {
        var _itemSetCountList []int64
        for _, item := range p.ItemSetCountList {
            _itemSetCountList = append(_itemSetCountList, item)
        }
        data["itemSetCountList"] = &_itemSetCountList
    }
    if p.ItemSetReferenceOfList != nil {
        var _itemSetReferenceOfList [][]string
        for _, item := range p.ItemSetReferenceOfList {
            _itemSetReferenceOfList = append(_itemSetReferenceOfList, item)
        }
        data["itemSetReferenceOfList"] = &_itemSetReferenceOfList
    }
    if p.ItemSetExpiresAtList != nil {
        var _itemSetExpiresAtList []int64
        for _, item := range p.ItemSetExpiresAtList {
            _itemSetExpiresAtList = append(_itemSetExpiresAtList, item)
        }
        data["itemSetExpiresAtList"] = &_itemSetExpiresAtList
    }
    if p.ItemSetCreatedAtList != nil {
        var _itemSetCreatedAtList []int64
        for _, item := range p.ItemSetCreatedAtList {
            _itemSetCreatedAtList = append(_itemSetCreatedAtList, item)
        }
        data["itemSetCreatedAtList"] = &_itemSetCreatedAtList
    }
    if p.ItemSetUpdatedAtList != nil {
        var _itemSetUpdatedAtList []int64
        for _, item := range p.ItemSetUpdatedAtList {
            _itemSetUpdatedAtList = append(_itemSetUpdatedAtList, item)
        }
        data["itemSetUpdatedAtList"] = &_itemSetUpdatedAtList
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ResponseCache struct {
    /** None */
	Region *string   `json:"region"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *string   `json:"responseCacheId"`
    /** None */
	RequestHash *string   `json:"requestHash"`
    /** APIの応答内容 */
	Result *string   `json:"result"`
}

func (p *ResponseCache) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["region"] = p.Region
    data["ownerId"] = p.OwnerId
    data["responseCacheId"] = p.ResponseCacheId
    data["requestHash"] = p.RequestHash
    data["result"] = p.Result
    return &data
}

type GitHubCheckoutSetting struct {
    /** リソースの取得に使用するGitHub のAPIキー のGRN */
	ApiKeyId *string   `json:"apiKeyId"`
    /** リポジトリ名 */
	RepositoryName *string   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *string   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *string   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *string   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *string   `json:"branchName"`
    /** タグ名 */
	TagName *string   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["apiKeyId"] = p.ApiKeyId
    data["repositoryName"] = p.RepositoryName
    data["sourcePath"] = p.SourcePath
    data["referenceType"] = p.ReferenceType
    data["commitHash"] = p.CommitHash
    data["branchName"] = p.BranchName
    data["tagName"] = p.TagName
    return &data
}

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *string   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *string   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *string   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *string   `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerScriptId"] = p.TriggerScriptId
    data["doneTriggerTargetType"] = p.DoneTriggerTargetType
    data["doneTriggerScriptId"] = p.DoneTriggerScriptId
    data["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
