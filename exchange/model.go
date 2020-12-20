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
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** 直接交換APIの呼び出しを許可する。許可しない場合はスタンプシート経由でしか交換できない */
	EnableDirectExchange *bool   `json:"enableDirectExchange"`
    /** 交換結果の受け取りに待ち時間の発生する交換機能を利用するか */
	EnableAwaitExchange *bool   `json:"enableAwaitExchange"`
    /** 交換処理をジョブとして追加するキューのネームスペース のGRN */
	QueueNamespaceId *core.String   `json:"queueNamespaceId"`
    /** 交換処理のスタンプシートで使用する暗号鍵GRN */
	KeyId *core.String   `json:"keyId"`
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
    data["enableDirectExchange"] = p.EnableDirectExchange
    data["enableAwaitExchange"] = p.EnableAwaitExchange
    data["queueNamespaceId"] = p.QueueNamespaceId
    data["keyId"] = p.KeyId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type RateModel struct {
    /** 交換レートマスター */
	RateModelId *core.String   `json:"rateModelId"`
    /** 交換レートの種類名 */
	Name *core.String   `json:"name"`
    /** 交換レートの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 消費アクションリスト */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** 交換の種類 */
	TimingType *core.String   `json:"timingType"`
    /** 交換実行から実際に報酬を受け取れるようになるまでの待ち時間（分） */
	LockTime *int32   `json:"lockTime"`
    /** スキップをすることができるか */
	EnableSkip *bool   `json:"enableSkip"`
    /** 時短消費アクションリスト */
	SkipConsumeActions *[]*ConsumeAction   `json:"skipConsumeActions"`
    /** 入手アクションリスト */
	AcquireActions *[]*AcquireAction   `json:"acquireActions"`
}

func (p *RateModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["rateModelId"] = p.RateModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.ConsumeActions != nil {
        var _consumeActions []*map[string]interface {}
        for _, item := range *p.ConsumeActions {
            _consumeActions = append(_consumeActions, item.ToDict())
        }
        data["consumeActions"] = &_consumeActions
    }
    data["timingType"] = p.TimingType
    data["lockTime"] = p.LockTime
    data["enableSkip"] = p.EnableSkip
    if p.SkipConsumeActions != nil {
        var _skipConsumeActions []*map[string]interface {}
        for _, item := range *p.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item.ToDict())
        }
        data["skipConsumeActions"] = &_skipConsumeActions
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

type RateModelMaster struct {
    /** 交換レートマスター */
	RateModelId *core.String   `json:"rateModelId"`
    /** 交換レート名 */
	Name *core.String   `json:"name"`
    /** 交換レートマスターの説明 */
	Description *core.String   `json:"description"`
    /** 交換レートのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 消費アクションリスト */
	ConsumeActions *[]*ConsumeAction   `json:"consumeActions"`
    /** 交換の種類 */
	TimingType *core.String   `json:"timingType"`
    /** 交換実行から実際に報酬を受け取れるようになるまでの待ち時間（分） */
	LockTime *int32   `json:"lockTime"`
    /** スキップをすることができるか */
	EnableSkip *bool   `json:"enableSkip"`
    /** 時短消費アクションリスト */
	SkipConsumeActions *[]*ConsumeAction   `json:"skipConsumeActions"`
    /** 入手アクションリスト */
	AcquireActions *[]*AcquireAction   `json:"acquireActions"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *RateModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["rateModelId"] = p.RateModelId
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
    data["timingType"] = p.TimingType
    data["lockTime"] = p.LockTime
    data["enableSkip"] = p.EnableSkip
    if p.SkipConsumeActions != nil {
        var _skipConsumeActions []*map[string]interface {}
        for _, item := range *p.SkipConsumeActions {
            _skipConsumeActions = append(_skipConsumeActions, item.ToDict())
        }
        data["skipConsumeActions"] = &_skipConsumeActions
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

type CurrentRateMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentRateMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type Await struct {
    /** 交換待機 */
	AwaitId *core.String   `json:"awaitId"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 交換レート名 */
	RateName *core.String   `json:"rateName"`
    /** 交換待機の名前 */
	Name *core.String   `json:"name"`
    /** 交換数 */
	Count *int32   `json:"count"`
    /** 作成日時 */
	ExchangedAt *int64   `json:"exchangedAt"`
}

func (p *Await) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["awaitId"] = p.AwaitId
    data["userId"] = p.UserId
    data["rateName"] = p.RateName
    data["name"] = p.Name
    data["count"] = p.Count
    data["exchangedAt"] = p.ExchangedAt
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
