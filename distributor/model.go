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

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** ネームスペースの説明 */
	Description *string `json:"description"`
	/** リソース溢れ処理に使用する ユーザ のGRN */
	AssumeUserId *string `json:"assumeUserId"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
	/** None */
	Status *string `json:"status"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
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
	data["status"] = p.Status
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type DistributorModelMaster struct {
	/** 配信設定マスター */
	DistributorModelId *string `json:"distributorModelId"`
	/** 配信設定名 */
	Name *string `json:"name"`
	/** 配信設定マスターの説明 */
	Description *string `json:"description"`
	/** 配信設定のメタデータ */
	Metadata *string `json:"metadata"`
	/** 所持品がキャパシティをオーバーしたときに転送するプレゼントボックスのネームスペース のGRN */
	InboxNamespaceId *string `json:"inboxNamespaceId"`
	/** ディストリビューターを通して処理出来る対象のリソースGRNのホワイトリスト */
	WhiteListTargetIds *[]string `json:"whiteListTargetIds"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *DistributorModelMaster) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["distributorModelId"] = p.DistributorModelId
	data["name"] = p.Name
	data["description"] = p.Description
	data["metadata"] = p.Metadata
	data["inboxNamespaceId"] = p.InboxNamespaceId
	if p.WhiteListTargetIds != nil {
		var _whiteListTargetIds []string
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
	DistributorModelId *string `json:"distributorModelId"`
	/** ディストリビューターの種類名 */
	Name *string `json:"name"`
	/** ディストリビューターの種類のメタデータ */
	Metadata *string `json:"metadata"`
	/** 所持品がキャパシティをオーバーしたときに転送するプレゼントボックスのネームスペース のGRN */
	InboxNamespaceId *string `json:"inboxNamespaceId"`
	/** ディストリビューターを通して処理出来る対象のリソースGRNのホワイトリスト */
	WhiteListTargetIds *[]string `json:"whiteListTargetIds"`
}

func (p *DistributorModel) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["distributorModelId"] = p.DistributorModelId
	data["name"] = p.Name
	data["metadata"] = p.Metadata
	data["inboxNamespaceId"] = p.InboxNamespaceId
	if p.WhiteListTargetIds != nil {
		var _whiteListTargetIds []string
		for _, item := range *p.WhiteListTargetIds {
			_whiteListTargetIds = append(_whiteListTargetIds, item)
		}
		data["whiteListTargetIds"] = &_whiteListTargetIds
	}
	return &data
}

type CurrentDistributorMaster struct {
	/** ネームスペース名 */
	NamespaceName *string `json:"namespaceName"`
	/** マスターデータ */
	Settings *string `json:"settings"`
}

func (p *CurrentDistributorMaster) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["namespaceName"] = p.NamespaceName
	data["settings"] = p.Settings
	return &data
}

type ResponseCache struct {
	/** None */
	Region *string `json:"region"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** レスポンスキャッシュ のGRN */
	ResponseCacheId *string `json:"responseCacheId"`
	/** None */
	RequestHash *string `json:"requestHash"`
	/** APIの応答内容 */
	Result *string `json:"result"`
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
	GitHubApiKeyId *string `json:"gitHubApiKeyId"`
	/** リポジトリ名 */
	RepositoryName *string `json:"repositoryName"`
	/** ソースコードのファイルパス */
	SourcePath *string `json:"sourcePath"`
	/** コードの取得元 */
	ReferenceType *string `json:"referenceType"`
	/** コミットハッシュ */
	CommitHash *string `json:"commitHash"`
	/** ブランチ名 */
	BranchName *string `json:"branchName"`
	/** タグ名 */
	TagName *string `json:"tagName"`
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
	Action *string `json:"action"`
	/** 加算リクエストのJSON */
	Request *string `json:"request"`
}

func (p *DistributeResource) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["action"] = p.Action
	data["request"] = p.Request
	return &data
}

type LogSetting struct {
	/** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["loggingNamespaceId"] = p.LoggingNamespaceId
	return &data
}
