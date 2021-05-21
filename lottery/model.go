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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** 景品付与処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *string   `json:"queueNamespaceId"`
    /** 景品付与処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *string   `json:"keyId"`
    /** 抽選処理時 に実行されるスクリプト のGRN */
	LotteryTriggerScriptId *string   `json:"lotteryTriggerScriptId"`
    /** 排出テーブル選択時 に実行されるスクリプト のGRN */
	ChoicePrizeTableScriptId *string   `json:"choicePrizeTableScriptId"`
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
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
    data["lotteryTriggerScriptId"] = p.LotteryTriggerScriptId
    data["choicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type LotteryModelMaster struct {
    /** 抽選の種類マスター */
	LotteryModelId *string   `json:"lotteryModelId"`
    /** 抽選モデルの種類名 */
	Name *string   `json:"name"`
    /** 抽選モデルの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** 抽選の種類マスターの説明 */
	Description *string   `json:"description"`
    /** 抽選モード */
	Mode *string   `json:"mode"`
    /** 抽選方法 */
	Method *string   `json:"method"`
    /** 景品テーブルの名前 */
	PrizeTableName *string   `json:"prizeTableName"`
    /** 抽選テーブルを確定するスクリプト のGRN */
	ChoicePrizeTableScriptId *string   `json:"choicePrizeTableScriptId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *LotteryModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["lotteryModelId"] = p.LotteryModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    data["mode"] = p.Mode
    data["method"] = p.Method
    data["prizeTableName"] = p.PrizeTableName
    data["choicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type PrizeTableMaster struct {
    /** 排出確率テーブルマスター */
	PrizeTableId *string   `json:"prizeTableId"`
    /** 排出確率テーブル名 */
	Name *string   `json:"name"`
    /** 排出確率テーブルのメタデータ */
	Metadata *string   `json:"metadata"`
    /** 排出確率テーブルマスターの説明 */
	Description *string   `json:"description"`
    /** 景品リスト */
	Prizes []Prize   `json:"prizes"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *PrizeTableMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["prizeTableId"] = p.PrizeTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["description"] = p.Description
    if p.Prizes != nil {
        var _prizes []*map[string]interface {}
        for _, item := range p.Prizes {
            _prizes = append(_prizes, item.ToDict())
        }
        data["prizes"] = &_prizes
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Box struct {
    /** ボックス */
	BoxId *string   `json:"boxId"`
    /** 排出確率テーブル名 */
	PrizeTableName *string   `json:"prizeTableName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 排出済み景品のインデックスのリスト */
	DrawnIndexes []int32   `json:"drawnIndexes"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Box) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["boxId"] = p.BoxId
    data["prizeTableName"] = p.PrizeTableName
    data["userId"] = p.UserId
    if p.DrawnIndexes != nil {
        var _drawnIndexes []int32
        for _, item := range p.DrawnIndexes {
            _drawnIndexes = append(_drawnIndexes, item)
        }
        data["drawnIndexes"] = &_drawnIndexes
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type LotteryModel struct {
    /** 抽選の種類マスター */
	LotteryModelId *string   `json:"lotteryModelId"`
    /** 抽選モデルの種類名 */
	Name *string   `json:"name"`
    /** 抽選モデルの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** 抽選モード */
	Mode *string   `json:"mode"`
    /** 抽選方法 */
	Method *string   `json:"method"`
    /** 景品テーブルの名前 */
	PrizeTableName *string   `json:"prizeTableName"`
    /** 抽選テーブルを確定するスクリプト のGRN */
	ChoicePrizeTableScriptId *string   `json:"choicePrizeTableScriptId"`
}

func (p *LotteryModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["lotteryModelId"] = p.LotteryModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["mode"] = p.Mode
    data["method"] = p.Method
    data["prizeTableName"] = p.PrizeTableName
    data["choicePrizeTableScriptId"] = p.ChoicePrizeTableScriptId
    return &data
}

type PrizeTable struct {
    /** 排出確率テーブルマスター */
	PrizeTableId *string   `json:"prizeTableId"`
    /** 景品テーブル名 */
	Name *string   `json:"name"`
    /** 景品テーブルのメタデータ */
	Metadata *string   `json:"metadata"`
    /** 景品リスト */
	Prizes []Prize   `json:"prizes"`
}

func (p *PrizeTable) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["prizeTableId"] = p.PrizeTableId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Prizes != nil {
        var _prizes []*map[string]interface {}
        for _, item := range p.Prizes {
            _prizes = append(_prizes, item.ToDict())
        }
        data["prizes"] = &_prizes
    }
    return &data
}

type Probability struct {
    /** 景品の種類 */
	Prize *DrawnPrize   `json:"prize"`
    /** 排出確率(0.0〜1.0) */
	Rate *float32   `json:"rate"`
}

func (p *Probability) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Prize != nil {
        data["prize"] = *p.Prize.ToDict()
    }
    data["rate"] = p.Rate
    return &data
}

type CurrentLotteryMaster struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentLotteryMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
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

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}

type Prize struct {
    /** 景品ID */
	PrizeId *string   `json:"prizeId"`
    /** 景品の種類 */
	Type *string   `json:"type"`
    /** 景品の入手アクションリスト */
	AcquireActions []AcquireAction   `json:"acquireActions"`
    /** 排出確率テーブルの名前 */
	PrizeTableName *string   `json:"prizeTableName"`
    /** 排出重み */
	Weight *int32   `json:"weight"`
}

func (p *Prize) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["prizeId"] = p.PrizeId
    data["type"] = p.Type
    if p.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range p.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        data["acquireActions"] = &_acquireActions
    }
    data["prizeTableName"] = p.PrizeTableName
    data["weight"] = p.Weight
    return &data
}

type AcquireAction struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *string   `json:"action"`
    /** 入手リクエストのJSON */
	Request *string   `json:"request"`
}

func (p *AcquireAction) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
    return &data
}

type DrawnPrize struct {
    /** 入手アクションのリスト */
	AcquireActions []AcquireAction   `json:"acquireActions"`
}

func (p *DrawnPrize) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range p.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        data["acquireActions"] = &_acquireActions
    }
    return &data
}

type BoxItem struct {
    /** 入手アクションのリスト */
	AcquireActions []AcquireAction   `json:"acquireActions"`
    /** 残り数量 */
	Remaining *int32   `json:"remaining"`
    /** 初期数量 */
	Initial *int32   `json:"initial"`
}

func (p *BoxItem) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.AcquireActions != nil {
        var _acquireActions []*map[string]interface {}
        for _, item := range p.AcquireActions {
            _acquireActions = append(_acquireActions, item.ToDict())
        }
        data["acquireActions"] = &_acquireActions
    }
    data["remaining"] = p.Remaining
    data["initial"] = p.Initial
    return &data
}

type BoxItems struct {
    /** ボックス */
	BoxId *string   `json:"boxId"`
    /** 排出確率テーブル名 */
	PrizeTableName *string   `json:"prizeTableName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ボックスから取り出したアイテムのリスト */
	Items []BoxItem   `json:"items"`
}

func (p *BoxItems) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["boxId"] = p.BoxId
    data["prizeTableName"] = p.PrizeTableName
    data["userId"] = p.UserId
    if p.Items != nil {
        var _items []*map[string]interface {}
        for _, item := range p.Items {
            _items = append(_items, item.ToDict())
        }
        data["items"] = &_items
    }
    return &data
}

type Config struct {
    /** 名前 */
	Key *string   `json:"key"`
    /** 値 */
	Value *string   `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["key"] = p.Key
    data["value"] = p.Value
    return &data
}
