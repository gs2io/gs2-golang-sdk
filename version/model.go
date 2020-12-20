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
    /** 説明文 */
	Description *core.String   `json:"description"`
    /** バージョンチェック通過後に改めて発行するプロジェクトトークンの権限判定に使用する ユーザ のGRN */
	AssumeUserId *core.String   `json:"assumeUserId"`
    /** バージョンを承認したときに実行するスクリプト */
	AcceptVersionScript *ScriptSetting   `json:"acceptVersionScript"`
    /** バージョンチェック時 に実行されるスクリプト のGRN */
	CheckVersionTriggerScriptId *core.String   `json:"checkVersionTriggerScriptId"`
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
    data["assumeUserId"] = p.AssumeUserId
    if p.AcceptVersionScript != nil {
        data["acceptVersionScript"] = *p.AcceptVersionScript.ToDict()
    }
    data["checkVersionTriggerScriptId"] = p.CheckVersionTriggerScriptId
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type VersionModelMaster struct {
    /** バージョンマスター */
	VersionModelId *core.String   `json:"versionModelId"`
    /** バージョン名 */
	Name *core.String   `json:"name"`
    /** バージョンマスターの説明 */
	Description *core.String   `json:"description"`
    /** バージョンのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** バージョンアップを促すバージョン */
	WarningVersion *Version   `json:"warningVersion"`
    /** バージョンチェックを蹴るバージョン */
	ErrorVersion *Version   `json:"errorVersion"`
    /** 判定に使用するバージョン値の種類 */
	Scope *core.String   `json:"scope"`
    /** 現在のバージョン */
	CurrentVersion *Version   `json:"currentVersion"`
    /** 判定するバージョン値に署名検証を必要とするか */
	NeedSignature *bool   `json:"needSignature"`
    /** 署名検証に使用する暗号鍵 のGRN */
	SignatureKeyId *core.String   `json:"signatureKeyId"`
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
	VersionModelId *core.String   `json:"versionModelId"`
    /** バージョンの種類名 */
	Name *core.String   `json:"name"`
    /** バージョンの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** バージョンアップを促すバージョン */
	WarningVersion *Version   `json:"warningVersion"`
    /** バージョンチェックを蹴るバージョン */
	ErrorVersion *Version   `json:"errorVersion"`
    /** 判定に使用するバージョン値の種類 */
	Scope *core.String   `json:"scope"`
    /** 現在のバージョン */
	CurrentVersion *Version   `json:"currentVersion"`
    /** 判定するバージョン値に署名検証を必要とするか */
	NeedSignature *bool   `json:"needSignature"`
    /** 署名検証に使用する暗号鍵 のGRN */
	SignatureKeyId *core.String   `json:"signatureKeyId"`
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
	AcceptVersionId *core.String   `json:"acceptVersionId"`
    /** 承認したバージョン名 */
	VersionName *core.String   `json:"versionName"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
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
	VersionName *core.String   `json:"versionName"`
    /** バージョン */
	Version *Version   `json:"version"`
    /** ボディ */
	Body *core.String   `json:"body"`
    /** 署名 */
	Signature *core.String   `json:"signature"`
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
	Region *core.String   `json:"region"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** バージョンの種類名 */
	VersionName *core.String   `json:"versionName"`
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

type CurrentVersionMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentVersionMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type ScriptSetting struct {
    /** 実行前に使用する GS2-Script のスクリプト のGRN */
	TriggerScriptId *core.String   `json:"triggerScriptId"`
    /** 完了通知の通知先 */
	DoneTriggerTargetType *core.String   `json:"doneTriggerTargetType"`
    /** 完了時に使用する GS2-Script のスクリプト のGRN */
	DoneTriggerScriptId *core.String   `json:"doneTriggerScriptId"`
    /** 完了時に使用する GS2-JobQueue のネームスペース のGRN */
	DoneTriggerQueueNamespaceId *core.String   `json:"doneTriggerQueueNamespaceId"`
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
