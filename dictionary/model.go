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

package dictionary

type Namespace struct {
	/** ネームスペース */
	NamespaceId *string `json:"namespaceId"`
	/** オーナーID */
	OwnerId *string `json:"ownerId"`
	/** ネームスペース名 */
	Name *string `json:"name"`
	/** ネームスペースの説明 */
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

type EntryModel struct {
	/** エントリーモデルマスター */
	EntryModelId *string `json:"entryModelId"`
	/** エントリーの種類名 */
	Name *string `json:"name"`
	/** エントリーの種類のメタデータ */
	Metadata *string `json:"metadata"`
}

func (p *EntryModel) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["entryModelId"] = p.EntryModelId
	data["name"] = p.Name
	data["metadata"] = p.Metadata
	return &data
}

type EntryModelMaster struct {
	/** エントリーモデルマスター */
	EntryModelId *string `json:"entryModelId"`
	/** エントリーモデル名 */
	Name *string `json:"name"`
	/** エントリーモデルマスターの説明 */
	Description *string `json:"description"`
	/** エントリーモデルのメタデータ */
	Metadata *string `json:"metadata"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *EntryModelMaster) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["entryModelId"] = p.EntryModelId
	data["name"] = p.Name
	data["description"] = p.Description
	data["metadata"] = p.Metadata
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type Entry struct {
	/** エントリー のGRN */
	EntryId *string `json:"entryId"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** エントリーの種類名 */
	Name *string `json:"name"`
	/** None */
	AcquiredAt *int64 `json:"acquiredAt"`
}

func (p *Entry) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["entryId"] = p.EntryId
	data["userId"] = p.UserId
	data["name"] = p.Name
	data["acquiredAt"] = p.AcquiredAt
	return &data
}

type Toc struct {
	/** 見出し */
	TocId *string `json:"tocId"`
	/** ユーザーID */
	UserId *string `json:"userId"`
	/** インデックス */
	Index *int32 `json:"index"`
	/** エントリーのリスト */
	Entries *[]*Entry `json:"entries"`
	/** 作成日時 */
	CreatedAt *int64 `json:"createdAt"`
	/** 最終更新日時 */
	UpdatedAt *int64 `json:"updatedAt"`
}

func (p *Toc) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["tocId"] = p.TocId
	data["userId"] = p.UserId
	data["index"] = p.Index
	if p.Entries != nil {
		var _entries []*map[string]interface{}
		for _, item := range *p.Entries {
			_entries = append(_entries, item.ToDict())
		}
		data["entries"] = &_entries
	}
	data["createdAt"] = p.CreatedAt
	data["updatedAt"] = p.UpdatedAt
	return &data
}

type CurrentEntryMaster struct {
	/** ネームスペース名 */
	NamespaceName *string `json:"namespaceName"`
	/** マスターデータ */
	Settings *string `json:"settings"`
}

func (p *CurrentEntryMaster) ToDict() *map[string]interface{} {
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

type Config struct {
	/** 名前 */
	Key *string `json:"key"`
	/** 値 */
	Value *string `json:"value"`
}

func (p *Config) ToDict() *map[string]interface{} {
	var data = map[string]interface{}{}
	data["key"] = p.Key
	data["value"] = p.Value
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
