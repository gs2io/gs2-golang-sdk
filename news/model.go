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

package news

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペースの名前 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** バージョン */
	Version *core.String   `json:"version"`
    /** ログの出力設定 */
	LogSetting *LogSetting   `json:"logSetting"`
    /** None */
	Status *core.String   `json:"status"`
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
    data["version"] = p.Version
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type News struct {
    /** セクション名 */
	Section *core.String   `json:"section"`
    /** コンテンツ名 */
	Content *core.String   `json:"content"`
    /** 記事見出し */
	Title *core.String   `json:"title"`
    /** 配信期間を決定する GS2-Schedule のイベントID */
	ScheduleEventId *core.String   `json:"scheduleEventId"`
    /** タイムスタンプ */
	Timestamp *int64   `json:"timestamp"`
    /** Front Matter */
	FrontMatter *core.String   `json:"frontMatter"`
}

func (p *News) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["section"] = p.Section
    data["content"] = p.Content
    data["title"] = p.Title
    data["scheduleEventId"] = p.ScheduleEventId
    data["timestamp"] = p.Timestamp
    data["frontMatter"] = p.FrontMatter
    return &data
}

type SetCookieRequestEntry struct {
    /** 記事を閲覧できるようにするために設定してほしい Cookie のキー値 */
	Key *core.String   `json:"key"`
    /** 記事を閲覧できるようにするために設定してほしい Cookie の値 */
	Value *core.String   `json:"value"`
}

func (p *SetCookieRequestEntry) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["key"] = p.Key
    data["value"] = p.Value
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

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *core.String   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
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
