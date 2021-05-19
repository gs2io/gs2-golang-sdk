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

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
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

type DataObject struct {
    /** データオブジェクト */
	DataObjectId *string   `json:"dataObjectId"`
    /** データの名前 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ファイルのアクセス権 */
	Scope *string   `json:"scope"`
    /** 公開するユーザIDリスト */
	AllowUserIds []string   `json:"allowUserIds"`
    /** プラットフォーム */
	Platform *string   `json:"platform"`
    /** 状態 */
	Status *string   `json:"status"`
    /** データの世代 */
	Generation *string   `json:"generation"`
    /** 以前有効だったデータの世代 */
	PreviousGeneration *string   `json:"previousGeneration"`
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
        var _allowUserIds []string
        for _, item := range p.AllowUserIds {
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
	DataObjectHistoryId *string   `json:"dataObjectHistoryId"`
    /** データオブジェクト名 */
	DataObjectName *string   `json:"dataObjectName"`
    /** 世代ID */
	Generation *string   `json:"generation"`
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

type LogSetting struct {
    /** ログの記録に使用する GS2-Log のネームスペース のGRN */
	LoggingNamespaceId *string   `json:"loggingNamespaceId"`
}

func (p *LogSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["loggingNamespaceId"] = p.LoggingNamespaceId
    return &data
}
