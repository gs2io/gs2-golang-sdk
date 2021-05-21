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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
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
	CounterId *string   `json:"counterId"`
    /** 回数制限の種類の名前 */
	LimitName *string   `json:"limitName"`
    /** カウンターの名前 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
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
	LimitModelId *string   `json:"limitModelId"`
    /** 回数制限の種類名 */
	Name *string   `json:"name"`
    /** 回数制限の種類マスターの説明 */
	Description *string   `json:"description"`
    /** 回数制限の種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *string   `json:"resetDayOfWeek"`
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
    /** 現在有効な回数制限設定 */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentLimitMaster) ToDict() *map[string]interface{} {
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

type LimitModel struct {
    /** 回数制限の種類 */
	LimitModelId *string   `json:"limitModelId"`
    /** 回数制限の種類名 */
	Name *string   `json:"name"`
    /** 回数制限の種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** リセットタイミング */
	ResetType *string   `json:"resetType"`
    /** リセットをする日にち */
	ResetDayOfMonth *int32   `json:"resetDayOfMonth"`
    /** リセットする曜日 */
	ResetDayOfWeek *string   `json:"resetDayOfWeek"`
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
