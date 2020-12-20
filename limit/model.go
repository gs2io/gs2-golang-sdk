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

package limit

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
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Counter struct {
    /** カウンター */
	CounterId *core.String   `json:"counterId"`
    /** 回数制限の種類の名前 */
	LimitName *core.String   `json:"limitName"`
    /** カウンターの名前 */
	Name *core.String   `json:"name"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** カウント値 */
	Count *int32   `json:"count"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Counter) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["counterId"] = p.CounterId
    data["limitName"] = p.LimitName
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["count"] = p.Count
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type LimitModelMaster struct {
    /** 回数制限の種類マスター */
	LimitModelId *core.String   `json:"limitModelId"`
    /** 回数制限の種類名 */
	Name *core.String   `json:"name"`
    /** 回数制限の種類マスターの説明 */
	Description *core.String   `json:"description"`
    /** 回数制限の種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *core.String   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *LimitModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["limitModelId"] = p.LimitModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["resetType"] = p.ResetType
    data["resetDayOfMonth"] = p.ResetDayOfMonth
    data["resetDayOfWeek"] = p.ResetDayOfWeek
    data["resetHour"] = p.ResetHour
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentLimitMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentLimitMaster) ToDict() *map[string]interface{} {
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

type LimitModel struct {
    /** 回数制限の種類 */
	LimitModelId *core.String   `json:"limitModelId"`
    /** 回数制限の種類名 */
	Name *core.String   `json:"name"`
    /** 回数制限の種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** リセットタイミング */
	ResetType *core.String   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *core.String   `json:"resetDayOfWeek"`
    /** リセット時刻 */
	ResetHour *int32   `json:"resetHour"`
}

func (p *LimitModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["limitModelId"] = p.LimitModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["resetType"] = p.ResetType
    data["resetDayOfMonth"] = p.ResetDayOfMonth
    data["resetDayOfWeek"] = p.ResetDayOfWeek
    data["resetHour"] = p.ResetHour
    return &data
}
