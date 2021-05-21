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

package version

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** 説明文 */
	Description *string   `json:"description"`
    /** バージョンチェック通過後に改めて発行するプロジェクトトークンの権限判定に使用する ユーザ のGRN */
	AssumeUserId *string   `json:"assumeUserId"`
    /** バージョンを承認したときに実行するスクリプト */
	AcceptVersionScript *ScriptSetting   `json:"acceptVersionScript"`
    /** バージョンチェック時 に実行されるスクリプト のGRN */
	CheckVersionTriggerScriptId *string   `json:"checkVersionTriggerScriptId"`
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
    if p.AcceptVersionScript != nil {
        data["acceptVersionScript"] = *p.AcceptVersionScript.ToDict()
    }
    data["checkVersionTriggerScriptId"] = p.CheckVersionTriggerScriptId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type VersionModelMaster struct {
    /** バージョンマスター */
	VersionModelId *string   `json:"versionModelId"`
    /** バージョン名 */
	Name *string   `json:"name"`
    /** バージョンマスターの説明 */
	Description *string   `json:"description"`
    /** バージョンのメタデータ */
	Metadata *string   `json:"metadata"`
    /** バージョンアップを促すバージョン */
	WarningVersion *Version   `json:"warningVersion"`
    /** バージョンチェックを蹴るバージョン */
	ErrorVersion *Version   `json:"errorVersion"`
    /** 判定に使用するバージョン値の種類 */
	Scope *string   `json:"scope"`
    /** 現在のバージョン */
	CurrentVersion *Version   `json:"currentVersion"`
    /** 判定するバージョン値に署名検証を必要とするか */
	NeedSignature *bool   `json:"needSignature"`
    /** 署名検証に使用する暗号鍵 のGRN */
	SignatureKeyId *string   `json:"signatureKeyId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *VersionModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["versionModelId"] = p.VersionModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.WarningVersion != nil {
        data["warningVersion"] = *p.WarningVersion.ToDict()
    }
    if p.ErrorVersion != nil {
        data["errorVersion"] = *p.ErrorVersion.ToDict()
    }
    data["scope"] = p.Scope
    if p.CurrentVersion != nil {
        data["currentVersion"] = *p.CurrentVersion.ToDict()
    }
    data["needSignature"] = p.NeedSignature
    data["signatureKeyId"] = p.SignatureKeyId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Version struct {
    /** メジャーバージョン */
	Major *int32   `json:"major"`
    /** マイナーバージョン */
	Minor *int32   `json:"minor"`
    /** マイクロバージョン */
	Micro *int32   `json:"micro"`
}

func (p *Version) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["major"] = p.Major
    data["minor"] = p.Minor
    data["micro"] = p.Micro
    return &data
}

type VersionModel struct {
    /** バージョン設定 */
	VersionModelId *string   `json:"versionModelId"`
    /** バージョンの種類名 */
	Name *string   `json:"name"`
    /** バージョンの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** バージョンアップを促すバージョン */
	WarningVersion *Version   `json:"warningVersion"`
    /** バージョンチェックを蹴るバージョン */
	ErrorVersion *Version   `json:"errorVersion"`
    /** 判定に使用するバージョン値の種類 */
	Scope *string   `json:"scope"`
    /** 現在のバージョン */
	CurrentVersion *Version   `json:"currentVersion"`
    /** 判定するバージョン値に署名検証を必要とするか */
	NeedSignature *bool   `json:"needSignature"`
    /** 署名検証に使用する暗号鍵 のGRN */
	SignatureKeyId *string   `json:"signatureKeyId"`
}

func (p *VersionModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["versionModelId"] = p.VersionModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.WarningVersion != nil {
        data["warningVersion"] = *p.WarningVersion.ToDict()
    }
    if p.ErrorVersion != nil {
        data["errorVersion"] = *p.ErrorVersion.ToDict()
    }
    data["scope"] = p.Scope
    if p.CurrentVersion != nil {
        data["currentVersion"] = *p.CurrentVersion.ToDict()
    }
    data["needSignature"] = p.NeedSignature
    data["signatureKeyId"] = p.SignatureKeyId
    return &data
}

type AcceptVersion struct {
    /** 承認したバージョン */
	AcceptVersionId *string   `json:"acceptVersionId"`
    /** 承認したバージョン名 */
	VersionName *string   `json:"versionName"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** 承認したバージョン */
	Version *Version   `json:"version"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *AcceptVersion) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["acceptVersionId"] = p.AcceptVersionId
    data["versionName"] = p.VersionName
    data["userId"] = p.UserId
    if p.Version != nil {
        data["version"] = *p.Version.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Status struct {
    /** バージョン設定 */
	VersionModel *VersionModel   `json:"versionModel"`
    /** 現在のバージョン */
	CurrentVersion *Version   `json:"currentVersion"`
}

func (p *Status) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.VersionModel != nil {
        data["versionModel"] = *p.VersionModel.ToDict()
    }
    if p.CurrentVersion != nil {
        data["currentVersion"] = *p.CurrentVersion.ToDict()
    }
    return &data
}

type TargetVersion struct {
    /** バージョンの名前 */
	VersionName *string   `json:"versionName"`
    /** バージョン */
	Version *Version   `json:"version"`
    /** ボディ */
	Body *string   `json:"body"`
    /** 署名 */
	Signature *string   `json:"signature"`
}

func (p *TargetVersion) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["versionName"] = p.VersionName
    if p.Version != nil {
        data["version"] = *p.Version.ToDict()
    }
    data["body"] = p.Body
    data["signature"] = p.Signature
    return &data
}

type SignTargetVersion struct {
    /** None */
	Region *string   `json:"region"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	NamespaceName *string   `json:"namespaceName"`
    /** バージョンの種類名 */
	VersionName *string   `json:"versionName"`
    /** バージョン */
	Version *Version   `json:"version"`
}

func (p *SignTargetVersion) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["region"] = p.Region
    data["ownerId"] = p.OwnerId
    data["namespaceName"] = p.NamespaceName
    data["versionName"] = p.VersionName
    if p.Version != nil {
        data["version"] = *p.Version.ToDict()
    }
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

type CurrentVersionMaster struct {
    /** 現在有効なバージョン */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentVersionMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceId"] = p.NamespaceId
    data["settings"] = p.Settings
    return &data
}

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *string   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *string   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *string   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *string   `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerScriptId"] = p.TriggerScriptId
    data["doneTriggerTargetType"] = p.DoneTriggerTargetType
    data["doneTriggerScriptId"] = p.DoneTriggerScriptId
    data["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
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
