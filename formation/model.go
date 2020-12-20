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

package formation

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
    /** キャパシティを更新するときに実行するスクリプト */
	UpdateMoldScript *ScriptSetting   `json:"updateMoldScript"`
    /** フォームを更新するときに実行するスクリプト */
	UpdateFormScript *ScriptSetting   `json:"updateFormScript"`
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
    if p.UpdateMoldScript != nil {
        data["updateMoldScript"] = *p.UpdateMoldScript.ToDict()
    }
    if p.UpdateFormScript != nil {
        data["updateFormScript"] = *p.UpdateFormScript.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type FormModel struct {
    /** フォームマスター */
	FormModelId *core.String   `json:"formModelId"`
    /** フォームの種類名 */
	Name *core.String   `json:"name"`
    /** フォームの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スリットリスト */
	Slots *[]*SlotModel   `json:"slots"`
}

func (p *FormModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["formModelId"] = p.FormModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.Slots != nil {
        var _slots []*map[string]interface {}
        for _, item := range *p.Slots {
            _slots = append(_slots, item.ToDict())
        }
        data["slots"] = &_slots
    }
    return &data
}

type FormModelMaster struct {
    /** フォームマスター */
	FormModelId *core.String   `json:"formModelId"`
    /** フォーム名 */
	Name *core.String   `json:"name"`
    /** フォームマスターの説明 */
	Description *core.String   `json:"description"`
    /** フォームのメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** スロットリスト */
	Slots *[]*SlotModel   `json:"slots"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *FormModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["formModelId"] = p.FormModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    if p.Slots != nil {
        var _slots []*map[string]interface {}
        for _, item := range *p.Slots {
            _slots = append(_slots, item.ToDict())
        }
        data["slots"] = &_slots
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type MoldModel struct {
    /** フォームの保存領域マスター */
	MoldModelId *core.String   `json:"moldModelId"`
    /** フォームの保存領域名 */
	Name *core.String   `json:"name"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
    /** None */
	FormModel *FormModel   `json:"formModel"`
    /** フォームを保存できる初期キャパシティ */
	InitialMaxCapacity *int32   `json:"initialMaxCapacity"`
    /** フォームを保存できるキャパシティ */
	MaxCapacity *int32   `json:"maxCapacity"`
}

func (p *MoldModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["moldModelId"] = p.MoldModelId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    if p.FormModel != nil {
        data["formModel"] = *p.FormModel.ToDict()
    }
    data["initialMaxCapacity"] = p.InitialMaxCapacity
    data["maxCapacity"] = p.MaxCapacity
    return &data
}

type MoldModelMaster struct {
    /** フォームの保存領域マスター */
	MoldModelId *core.String   `json:"moldModelId"`
    /** フォームの保存領域名 */
	Name *core.String   `json:"name"`
    /** フォームの保存領域マスターの説明 */
	Description *core.String   `json:"description"`
    /** フォームの保存領域のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** フォーム名 */
	FormModelName *core.String   `json:"formModelName"`
    /** フォームを保存できる初期キャパシティ */
	InitialMaxCapacity *int32   `json:"initialMaxCapacity"`
    /** フォームを保存できるキャパシティ */
	MaxCapacity *int32   `json:"maxCapacity"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *MoldModelMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["moldModelId"] = p.MoldModelId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["formModelName"] = p.FormModelName
    data["initialMaxCapacity"] = p.InitialMaxCapacity
    data["maxCapacity"] = p.MaxCapacity
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type CurrentFormMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentFormMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["namespaceName"] = p.NamespaceName
    data["settings"] = p.Settings
    return &data
}

type Mold struct {
    /** 保存したフォーム */
	MoldId *core.String   `json:"moldId"`
    /** フォームの保存領域の名前 */
	Name *core.String   `json:"name"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 現在のキャパシティ */
	Capacity *int32   `json:"capacity"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Mold) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["moldId"] = p.MoldId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["capacity"] = p.Capacity
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Form struct {
    /** フォーム */
	FormId *core.String   `json:"formId"`
    /** フォームの保存領域の名前 */
	Name *core.String   `json:"name"`
    /** 保存領域のインデックス */
	Index *int32   `json:"index"`
    /** スロットリスト */
	Slots *[]*Slot   `json:"slots"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Form) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["formId"] = p.FormId
    data["name"] = p.Name
    data["index"] = p.Index
    if p.Slots != nil {
        var _slots []*map[string]interface {}
        for _, item := range *p.Slots {
            _slots = append(_slots, item.ToDict())
        }
        data["slots"] = &_slots
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Slot struct {
    /** スロットモデル名 */
	Name *core.String   `json:"name"`
    /** プロパティID */
	PropertyId *core.String   `json:"propertyId"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
}

func (p *Slot) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["propertyId"] = p.PropertyId
    data["metadata"] = p.Metadata
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

type SlotModel struct {
    /** スロットモデル名 */
	Name *core.String   `json:"name"`
    /** プロパティとして設定可能な値の正規表現 */
	PropertyRegex *core.String   `json:"propertyRegex"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
}

func (p *SlotModel) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["propertyRegex"] = p.PropertyRegex
    data["metadata"] = p.Metadata
    return &data
}

type SlotWithSignature struct {
    /** スロットモデル名 */
	Name *core.String   `json:"name"`
    /** プロパティの種類 */
	PropertyType *core.String   `json:"propertyType"`
    /** ペイロード */
	Body *core.String   `json:"body"`
    /** プロパティIDのリソースを所有していることを証明する署名 */
	Signature *core.String   `json:"signature"`
    /** メタデータ */
	Metadata *core.String   `json:"metadata"`
}

func (p *SlotWithSignature) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["propertyType"] = p.PropertyType
    data["body"] = p.Body
    data["signature"] = p.Signature
    data["metadata"] = p.Metadata
    return &data
}

type AcquireActionConfig struct {
    /** スロット名 */
	Name *core.String   `json:"name"`
    /** スタンプシートに使用するコンフィグ */
	Config *[]*Config   `json:"config"`
}

func (p *AcquireActionConfig) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    if p.Config != nil {
        var _config []*map[string]interface {}
        for _, item := range *p.Config {
            _config = append(_config, item.ToDict())
        }
        data["config"] = &_config
    }
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
