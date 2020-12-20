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

package schedule

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

type EventMaster struct {
    /** イベントマスター */
	EventId *core.String   `json:"eventId"`
    /** イベントの種類名 */
	Name *core.String   `json:"name"`
    /** イベントマスターの説明 */
	Description *core.String   `json:"description"`
    /** イベントの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** イベント期間の種類 */
	ScheduleType *core.String   `json:"scheduleType"`
    /** 繰り返しの種類 */
	RepeatType *core.String   `json:"repeatType"`
    /** イベントの開始日時 */
	AbsoluteBegin *int64   `json:"absoluteBegin"`
    /** イベントの終了日時 */
	AbsoluteEnd *int64   `json:"absoluteEnd"`
    /** イベントの繰り返し開始日 */
	RepeatBeginDayOfMonth *int32   `json:"repeatBeginDayOfMonth"`
    /** イベントの繰り返し終了日 */
	RepeatEndDayOfMonth *int32   `json:"repeatEndDayOfMonth"`
    /** イベントの繰り返し開始曜日 */
	RepeatBeginDayOfWeek *core.String   `json:"repeatBeginDayOfWeek"`
    /** イベントの繰り返し終了曜日 */
	RepeatEndDayOfWeek *core.String   `json:"repeatEndDayOfWeek"`
    /** イベントの繰り返し開始時間 */
	RepeatBeginHour *int32   `json:"repeatBeginHour"`
    /** イベントの繰り返し終了時間 */
	RepeatEndHour *int32   `json:"repeatEndHour"`
    /** イベントの開始トリガー名 */
	RelativeTriggerName *core.String   `json:"relativeTriggerName"`
    /** イベントの開催期間(秒) */
	RelativeDuration *int32   `json:"relativeDuration"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *EventMaster) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["eventId"] = p.EventId
    data["name"] = p.Name
    data["description"] = p.Description
    data["metadata"] = p.Metadata
    data["scheduleType"] = p.ScheduleType
    data["repeatType"] = p.RepeatType
    data["absoluteBegin"] = p.AbsoluteBegin
    data["absoluteEnd"] = p.AbsoluteEnd
    data["repeatBeginDayOfMonth"] = p.RepeatBeginDayOfMonth
    data["repeatEndDayOfMonth"] = p.RepeatEndDayOfMonth
    data["repeatBeginDayOfWeek"] = p.RepeatBeginDayOfWeek
    data["repeatEndDayOfWeek"] = p.RepeatEndDayOfWeek
    data["repeatBeginHour"] = p.RepeatBeginHour
    data["repeatEndHour"] = p.RepeatEndHour
    data["relativeTriggerName"] = p.RelativeTriggerName
    data["relativeDuration"] = p.RelativeDuration
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Trigger struct {
    /** トリガー */
	TriggerId *core.String   `json:"triggerId"`
    /** トリガーの名前 */
	Name *core.String   `json:"name"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** トリガーの有効期限 */
	ExpiresAt *int64   `json:"expiresAt"`
}

func (p *Trigger) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["triggerId"] = p.TriggerId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["createdAt"] = p.CreatedAt
    data["expiresAt"] = p.ExpiresAt
    return &data
}

type Event struct {
    /** イベントマスター */
	EventId *core.String   `json:"eventId"`
    /** イベントの種類名 */
	Name *core.String   `json:"name"`
    /** イベントの種類のメタデータ */
	Metadata *core.String   `json:"metadata"`
    /** イベント期間の種類 */
	ScheduleType *core.String   `json:"scheduleType"`
    /** 繰り返しの種類 */
	RepeatType *core.String   `json:"repeatType"`
    /** イベントの開始日時 */
	AbsoluteBegin *int64   `json:"absoluteBegin"`
    /** イベントの終了日時 */
	AbsoluteEnd *int64   `json:"absoluteEnd"`
    /** イベントの繰り返し開始日 */
	RepeatBeginDayOfMonth *int32   `json:"repeatBeginDayOfMonth"`
    /** イベントの繰り返し終了日 */
	RepeatEndDayOfMonth *int32   `json:"repeatEndDayOfMonth"`
    /** イベントの繰り返し開始曜日 */
	RepeatBeginDayOfWeek *core.String   `json:"repeatBeginDayOfWeek"`
    /** イベントの繰り返し終了曜日 */
	RepeatEndDayOfWeek *core.String   `json:"repeatEndDayOfWeek"`
    /** イベントの繰り返し開始時間 */
	RepeatBeginHour *int32   `json:"repeatBeginHour"`
    /** イベントの繰り返し終了時間 */
	RepeatEndHour *int32   `json:"repeatEndHour"`
    /** イベントの開始トリガー */
	RelativeTriggerName *core.String   `json:"relativeTriggerName"`
    /** イベントの開催期間(秒) */
	RelativeDuration *int32   `json:"relativeDuration"`
}

func (p *Event) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["eventId"] = p.EventId
    data["name"] = p.Name
    data["metadata"] = p.Metadata
    data["scheduleType"] = p.ScheduleType
    data["repeatType"] = p.RepeatType
    data["absoluteBegin"] = p.AbsoluteBegin
    data["absoluteEnd"] = p.AbsoluteEnd
    data["repeatBeginDayOfMonth"] = p.RepeatBeginDayOfMonth
    data["repeatEndDayOfMonth"] = p.RepeatEndDayOfMonth
    data["repeatBeginDayOfWeek"] = p.RepeatBeginDayOfWeek
    data["repeatEndDayOfWeek"] = p.RepeatEndDayOfWeek
    data["repeatBeginHour"] = p.RepeatBeginHour
    data["repeatEndHour"] = p.RepeatEndHour
    data["relativeTriggerName"] = p.RelativeTriggerName
    data["relativeDuration"] = p.RelativeDuration
    return &data
}

type CurrentEventMaster struct {
    /** ネームスペース名 */
	NamespaceName *core.String   `json:"namespaceName"`
    /** マスターデータ */
	Settings *core.String   `json:"settings"`
}

func (p *CurrentEventMaster) ToDict() *map[string]interface{} {
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
