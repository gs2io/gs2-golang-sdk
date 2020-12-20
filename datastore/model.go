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

package datastore

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
    /** アップロード完了報告時に実行するスクリプト */
	DoneUploadScript *ScriptSetting   `json:"doneUploadScript"`
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
    if p.DoneUploadScript != nil {
        data["doneUploadScript"] = *p.DoneUploadScript.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
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

type DataObject struct {
    /** データオブジェクト */
	DataObjectId *core.String   `json:"dataObjectId"`
    /** データの名前 */
	Name *core.String   `json:"name"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** ファイルのアクセス権 */
	Scope *core.String   `json:"scope"`
    /** 公開するユーザIDリスト */
	AllowUserIds *[]core.String   `json:"allowUserIds"`
    /** プラットフォーム */
	Platform *core.String   `json:"platform"`
    /** 状態 */
	Status *core.String   `json:"status"`
    /** データの世代 */
	Generation *core.String   `json:"generation"`
    /** 以前有効だったデータの世代 */
	PreviousGeneration *core.String   `json:"previousGeneration"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *DataObject) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["dataObjectId"] = p.DataObjectId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["scope"] = p.Scope
    if p.AllowUserIds != nil {
        var _allowUserIds []core.String
        for _, item := range *p.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        data["allowUserIds"] = &_allowUserIds
    }
    data["platform"] = p.Platform
    data["status"] = p.Status
    data["generation"] = p.Generation
    data["previousGeneration"] = p.PreviousGeneration
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type DataObjectHistory struct {
    /** データオブジェクト履歴 */
	DataObjectHistoryId *core.String   `json:"dataObjectHistoryId"`
    /** データオブジェクト名 */
	DataObjectName *core.String   `json:"dataObjectName"`
    /** 世代ID */
	Generation *core.String   `json:"generation"`
    /** データサイズ */
	ContentLength *int64   `json:"contentLength"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
}

func (p *DataObjectHistory) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["dataObjectHistoryId"] = p.DataObjectHistoryId
    data["dataObjectName"] = p.DataObjectName
    data["generation"] = p.Generation
    data["contentLength"] = p.ContentLength
    data["createdAt"] = p.CreatedAt
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
