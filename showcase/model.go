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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	Name *core.String   `json:"name"`
    /** 購入処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *core.String   `json:"queueNamespaceId"`
    /** 購入処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *core.String   `json:"keyId"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
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
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
    data["description"] = p.Description
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type SalesItemMaster struct {
    /** 商品マスター */
	SalesItemId *core.String   `json:"salesItemId"`
    /** 商品名 */
	Name *core.String   `json:"name"`
    /** 商品マスターの説明 */
	Description *core.String   `json:"description"`
    /** 商品のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 消費アクションリスト */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** 入手アクションリスト */
	AcquireActions *[]*AcquireAction   `json:"acquireActions"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *SalesItemMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["salesItemId"] = p.SalesItemId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range *p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range *p.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        data["acquireActions"] = &_acquireActions
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type SalesItemGroupMaster struct {
    /** 商品グループマスター */
	SalesItemGroupId *core.String   `json:"salesItemGroupId"`
    /** 商品名 */
	Name *core.String   `json:"name"`
    /** 商品グループマスターの説明 */
	Description *core.String   `json:"description"`
    /** 商品のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 商品グループに含める商品リスト */
	SalesItemNames *[]core.String   `json:"salesItemNames"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *SalesItemGroupMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["salesItemGroupId"] = p.SalesItemGroupId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.SalesItemNames != nil {
        var _salesItemNames []core.String
        for _, item := range *p.SalesItemNames {
            _salesItemNames = append(_salesItemNames, item)
        }
        data["salesItemNames"] = &_salesItemNames
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type ShowcaseMaster struct {
    /** 陳列棚マスター */
	ShowcaseId *core.String   `json:"showcaseId"`
    /** 陳列棚名 */
	Name *core.String   `json:"name"`
    /** 陳列棚マスターの説明 */
	Description *core.String   `json:"description"`
    /** 商品のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 陳列する商品モデル一覧 */
	DisplayItems *[]*DisplayItemMaster   `json:"displayItems"`
    /** 販売期間とするイベントマスター のGRN */
	SalesPeriodEventId *core.String   `json:"salesPeriodEventId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *ShowcaseMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["showcaseId"] = p.ShowcaseId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.DisplayItems != nil {
        var _displayItems []*map[string]interface {}
        for _, item := range *p.DisplayItems {
            _displayItems = append(_displayItems, item.ToDict())
        }
        data["displayItems"] = &_displayItems
    }
    data["salesPeriodEventId"] = p.SalesPeriodEventId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentShowcaseMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentShowcaseMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type ResponseCache struct {
    /** None */
	Region *core.String   `json:"region"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** レスポンスキャッシュ のGRN */
	ResponseCacheId *core.String   `json:"responseCacheId"`
    /** None */
	RequestHash *core.String   `json:"requestHash"`
    /** APIの応答内容 */
	Result *core.String   `json:"result"`
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
	GitHubApiKeyId *core.String   `json:"gitHubApiKeyId"`
    /** リポジトリ名 */
	RepositoryName *core.String   `json:"repositoryName"`
    /** ソースコードのファイルパス */
	SourcePath *core.String   `json:"sourcePath"`
    /** コードの取得元 */
	ReferenceType *core.String   `json:"referenceType"`
    /** コミットハッシュ */
	CommitHash *core.String   `json:"commitHash"`
    /** ブランチ名 */
	BranchName *core.String   `json:"branchName"`
    /** タグ名 */
	TagName *core.String   `json:"tagName"`
}

func (p *GitHubCheckoutSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gitHubApiKeyId"] = p.GitHubApiKeyId
    data["repositoryName"] = p.RepositoryName
    data["sourcePath"] = p.SourcePath
    data["referenceType"] = p.ReferenceType
    data["commitHash"] = p.CommitHash
    data["branchName"] = p.BranchName
    data["tagName"] = p.TagName
    return &data
}

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}

type SalesItem struct {
    /** 商品名 */
	Name *core.String   `json:"name"`
    /** 商品のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 消費アクションリスト */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** 入手アクションリスト */
	AcquireActions *[]*AcquireAction   `json:"acquireActions"`
}

func (p *SalesItem) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range *p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    if p.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range *p.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        data["acquireActions"] = &_acquireActions
    }
    return &data
}

type SalesItemGroup struct {
    /** 商品グループ名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 商品リスト */
	SalesItems *[]*SalesItem   `json:"salesItems"`
}

func (p *SalesItemGroup) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.SalesItems != nil {
        var _salesItems []*map[string]interface {}
        for _, item := range *p.SalesItems {
            _salesItems = append(_salesItems, item.ToDict())
        }
        data["salesItems"] = &_salesItems
    }
    return &data
}

type Showcase struct {
    /** 陳列棚 */
	ShowcaseId *core.String   `json:"showcaseId"`
    /** 商品名 */
	Name *core.String   `json:"name"`
    /** 商品のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** インベントリに格納可能なアイテムモデル一覧 */
	DisplayItems *[]*DisplayItem   `json:"displayItems"`
    /** 販売期間とするイベントマスター のGRN */
	SalesPeriodEventId *core.String   `json:"salesPeriodEventId"`
}

func (p *Showcase) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["showcaseId"] = p.ShowcaseId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.DisplayItems != nil {
        var _displayItems []*map[string]interface {}
        for _, item := range *p.DisplayItems {
            _displayItems = append(_displayItems, item.ToDict())
        }
        data["displayItems"] = &_displayItems
    }
    data["salesPeriodEventId"] = p.SalesPeriodEventId
    return &data
}

type DisplayItem struct {
    /** 陳列商品ID */
	DisplayItemId *core.String   `json:"displayItemId"`
    /** 種類 */
	Type *core.String   `json:"type"`
    /** 陳列する商品 */
	SalesItem *SalesItem   `json:"salesItem"`
    /** 陳列する商品グループ */
	SalesItemGroup *SalesItemGroup   `json:"salesItemGroup"`
    /** 販売期間とするイベントマスター のGRN */
	SalesPeriodEventId *core.String   `json:"salesPeriodEventId"`
}

func (p *DisplayItem) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["displayItemId"] = p.DisplayItemId
    data["type"] = p.Type
    if p.SalesItem != nil {
        data["salesItem"] = *p.SalesItem.ToDict()
    }
    if p.SalesItemGroup != nil {
        data["salesItemGroup"] = *p.SalesItemGroup.ToDict()
    }
    data["salesPeriodEventId"] = p.SalesPeriodEventId
    return &data
}

type Config struct {
    /** 名前 */
	Key *core.String   `json:"key"`
    /** 値 */
	Value *core.String   `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["key"] = p.Key
    data["value"] = p.Value
    return &data
}

type ConsumeAction struct {
    /** スタンプタスクで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** 消費リクエストのJSON */
	Request *core.String   `json:"request"`
}

func (p *ConsumeAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type AcquireAction struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** 入手リクエストのJSON */
	Request *core.String   `json:"request"`
}

func (p *AcquireAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type DisplayItemMaster struct {
    /** 陳列商品ID */
	DisplayItemId *core.String   `json:"displayItemId"`
    /** 種類 */
	Type *core.String   `json:"type"`
    /** 陳列する商品の名前 */
	SalesItemName *core.String   `json:"salesItemName"`
    /** 陳列する商品グループの名前 */
	SalesItemGroupName *core.String   `json:"salesItemGroupName"`
    /** 販売期間とするイベントマスター のGRN */
	SalesPeriodEventId *core.String   `json:"salesPeriodEventId"`
}

func (p *DisplayItemMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["displayItemId"] = p.DisplayItemId
    data["type"] = p.Type
    data["salesItemName"] = p.SalesItemName
    data["salesItemGroupName"] = p.SalesItemGroupName
    data["salesPeriodEventId"] = p.SalesPeriodEventId
    return &data
}
