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

package deploy

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Stack struct {
    /** スタック */
	StackId *core.String   `json:"stackId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** スタック名 */
	Name *core.String   `json:"name"`
    /** スタックの説明 */
	Description *core.String   `json:"description"`
    /** テンプレートデータ */
	Template *core.String   `json:"template"`
    /** 実行状態 */
	Status *core.String   `json:"status"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Stack) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["stackId"] = p.StackId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["description"] = p.Description
    data["template"] = p.Template
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Resource struct {
    /** 作成されたのリソース */
	ResourceId *core.String   `json:"resourceId"`
    /** 操作対象のリソース */
	Type *core.String   `json:"type"`
    /** 作成中のリソース名 */
	Name *core.String   `json:"name"`
    /** リクエストパラメータ */
	Request *core.String   `json:"request"`
    /** リソースの作成・更新のレスポンス */
	Response *core.String   `json:"response"`
    /** ロールバック操作の種類 */
	RollbackContext *core.String   `json:"rollbackContext"`
    /** ロールバック用のリクエストパラメータ */
	RollbackRequest *core.String   `json:"rollbackRequest"`
    /** ロールバック時に依存しているリソースの名前 */
	RollbackAfter *[]core.String   `json:"rollbackAfter"`
    /** リソースを作成したときに Output に記録するフィールド */
	OutputFields *[]*OutputField   `json:"outputFields"`
    /** このリソースが作成された時の実行 ID */
	WorkId *core.String   `json:"workId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Resource) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["resourceId"] = p.ResourceId
    data["type"] = p.Type
    data["name"] = p.Name
    data["request"] = p.Request
    data["response"] = p.Response
    data["rollbackContext"] = p.RollbackContext
    data["rollbackRequest"] = p.RollbackRequest
    if p.RollbackAfter != nil {
        var _rollbackAfter []core.String
        for _, item := range *p.RollbackAfter {
            _rollbackAfter = append(_rollbackAfter, item)
        }
        data["rollbackAfter"] = &_rollbackAfter
    }
    if p.OutputFields != nil {
        var _outputFields []*map[string]interface {}
        for _, item := range *p.OutputFields {
            _outputFields = append(_outputFields, item.ToDict())
        }
        data["outputFields"] = &_outputFields
    }
    data["workId"] = p.WorkId
    data["createdAt"] = p.CreatedAt
    return &data
}

type WorkingStack struct {
    /** 実行中のスタック */
	StackId *core.String   `json:"stackId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** 実行中のスタック名 */
	Name *core.String   `json:"name"`
    /** 実行に対して割り振られる一意な ID */
	WorkId *core.String   `json:"workId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *WorkingStack) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["stackId"] = p.StackId
    data["ownerId"] = p.OwnerId
    data["name"] = p.Name
    data["workId"] = p.WorkId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type WorkingResource struct {
    /** 作成中のリソース */
	ResourceId *core.String   `json:"resourceId"`
    /** 操作の種類 */
	Context *core.String   `json:"context"`
    /** 操作対象のリソース */
	Type *core.String   `json:"type"`
    /** 作成中のリソース名 */
	Name *core.String   `json:"name"`
    /** リクエストパラメータ */
	Request *core.String   `json:"request"`
    /** 依存しているリソースの名前 */
	After *[]core.String   `json:"after"`
    /** ロールバック操作の種類 */
	RollbackContext *core.String   `json:"rollbackContext"`
    /** ロールバック用のリクエストパラメータ */
	RollbackRequest *core.String   `json:"rollbackRequest"`
    /** ロールバック時に依存しているリソースの名前 */
	RollbackAfter *[]core.String   `json:"rollbackAfter"`
    /** リソースを作成したときに Output に記録するフィールド */
	OutputFields *[]*OutputField   `json:"outputFields"`
    /** 実行に対して割り振られる一意な ID */
	WorkId *core.String   `json:"workId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *WorkingResource) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["resourceId"] = p.ResourceId
    data["context"] = p.Context
    data["type"] = p.Type
    data["name"] = p.Name
    data["request"] = p.Request
    if p.After != nil {
        var _after []core.String
        for _, item := range *p.After {
            _after = append(_after, item)
        }
        data["after"] = &_after
    }
    data["rollbackContext"] = p.RollbackContext
    data["rollbackRequest"] = p.RollbackRequest
    if p.RollbackAfter != nil {
        var _rollbackAfter []core.String
        for _, item := range *p.RollbackAfter {
            _rollbackAfter = append(_rollbackAfter, item)
        }
        data["rollbackAfter"] = &_rollbackAfter
    }
    if p.OutputFields != nil {
        var _outputFields []*map[string]interface {}
        for _, item := range *p.OutputFields {
            _outputFields = append(_outputFields, item.ToDict())
        }
        data["outputFields"] = &_outputFields
    }
    data["workId"] = p.WorkId
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Event struct {
    /** 発生したイベント */
	EventId *core.String   `json:"eventId"`
    /** イベント名 */
	Name *core.String   `json:"name"`
    /** イベントの種類 */
	ResourceName *core.String   `json:"resourceName"`
    /** イベントの種類 */
	Type *core.String   `json:"type"`
    /** メッセージ */
	Message *core.String   `json:"message"`
    /** 日時 */
	EventAt *int64   `json:"eventAt"`
}

func (p *Event) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["eventId"] = p.EventId
    data["name"] = p.Name
    data["resourceName"] = p.ResourceName
    data["type"] = p.Type
    data["message"] = p.Message
    data["eventAt"] = p.EventAt
    return &data
}

type Output struct {
    /** アウトプット */
	OutputId *core.String   `json:"outputId"`
    /** アウトプット名 */
	Name *core.String   `json:"name"`
    /** 値 */
	Value *core.String   `json:"value"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *Output) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["outputId"] = p.OutputId
    data["name"] = p.Name
    data["value"] = p.Value
    data["createdAt"] = p.CreatedAt
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

type OutputField struct {
    /** 名前 */
	Name *core.String   `json:"name"`
    /** フィールド名 */
	FieldName *core.String   `json:"fieldName"`
}

func (p *OutputField) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["name"] = p.Name
    data["fieldName"] = p.FieldName
    return &data
}
