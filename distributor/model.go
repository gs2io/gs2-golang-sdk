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

package distributor

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
    /** リソース溢れ処理に使用する ユーザ のGRN */
	AssumeUserId *core.String   `json:"assumeUserId"`
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
    data["assumeUserId"] = p.AssumeUserId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type DistributorModelMaster struct {
    /** 配信設定マスター */
	DistributorModelId *core.String   `json:"distributorModelId"`
    /** 配信設定名 */
	Name *core.String   `json:"name"`
    /** 配信設定マスターの説明 */
	Description *core.String   `json:"description"`
    /** 配信設定のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 所持品がキャパシティをオーバーしたときに転送するプレゼントボックスのネームスペース のGRN */
	InboxNamespaceId *core.String   `json:"inboxNamespaceId"`
    /** ディストリビューターを通して処理出来る対象のリソースGRNのホワイトリスト */
	WhiteListTargetIds *[]core.String   `json:"whiteListTargetIds"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *DistributorModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["distributorModelId"] = p.DistributorModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["inboxNamespaceId"] = p.InboxNamespaceId
    if p.WhiteListTargetIds != nil {
        var _whiteListTargetIds []core.String
        for _, item := range *p.WhiteListTargetIds {
            _whiteListTargetIds = append(_whiteListTargetIds, item)
        }
        data["whiteListTargetIds"] = &_whiteListTargetIds
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type DistributorModel struct {
    /** 配信設定 */
	DistributorModelId *core.String   `json:"distributorModelId"`
    /** ディストリビューターの種類名 */
	Name *core.String   `json:"name"`
    /** ディストリビューターの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** 所持品がキャパシティをオーバーしたときに転送するプレゼントボックスのネームスペース のGRN */
	InboxNamespaceId *core.String   `json:"inboxNamespaceId"`
    /** ディストリビューターを通して処理出来る対象のリソースGRNのホワイトリスト */
	WhiteListTargetIds *[]core.String   `json:"whiteListTargetIds"`
}

func (p *DistributorModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["distributorModelId"] = p.DistributorModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["inboxNamespaceId"] = p.InboxNamespaceId
    if p.WhiteListTargetIds != nil {
        var _whiteListTargetIds []core.String
        for _, item := range *p.WhiteListTargetIds {
            _whiteListTargetIds = append(_whiteListTargetIds, item)
        }
        data["whiteListTargetIds"] = &_whiteListTargetIds
    }
    return &data
}

type CurrentDistributorMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentDistributorMaster) ToDict() *map[string]interface{} {
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

type DistributeResource struct {
    /** スタンプシートで実行するアクションの種類 */
	Action *core.String   `json:"action"`
    /** 加算リクエストのJSON */
	Request *core.String   `json:"request"`
}

func (p *DistributeResource) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["action"] = p.Action
    data["request"] = p.Request
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
