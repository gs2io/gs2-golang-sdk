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

package script

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** 説明文 */
	Description *string `json:"description"`
	/** ログの出力設定 */
	LogSetting *LogSetting `json:"logSetting"`
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
	if p.LogSetting != nil {
		data["logSetting"] = *p.LogSetting.ToDict()
	}
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Script struct {
	/** スクリプト */
	ScriptId *string `json:"scriptId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** スクリプト名 */
	Name *string `json:"name"`
	/** 説明文 */
	Description *string `json:"description"`
	/** Luaスクリプト */
	Script *string `json:"script"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Script) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["scriptId"] = p.ScriptId
	data["ownerId"] = p.OwnerId
	data["name"] = p.Name
	data["description"] = p.Description
	data["script"] = p.Script
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
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

type LogSetting struct {
	/** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["loggingNamespaceId"] = p.LoggingNamespaceId
	return &data
}
