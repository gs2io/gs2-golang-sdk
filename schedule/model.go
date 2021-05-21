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

type EventMaster struct {
    /** イベントマスター */
	EventId *string   `json:"eventId"`
    /** イベントの種類名 */
	Name *string   `json:"name"`
    /** イベントマスターの説明 */
	Description *string   `json:"description"`
    /** イベントの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** イベント期間の種類 */
	ScheduleType *string   `json:"scheduleType"`
    /** 繰り返しの種類 */
	RepeatType *string   `json:"repeatType"`
    /** イベントの開始日時 */
	AbsoluteBegin *int64   `json:"absoluteBegin"`
    /** イベントの終了日時 */
	AbsoluteEnd *int64   `json:"absoluteEnd"`
    /** イベントの繰り返し開始日 */
	RepeatBeginDayOfMonth *int32   `json:"repeatBeginDayOfMonth"`
    /** イベントの繰り返し終了日 */
	RepeatEndDayOfMonth *int32   `json:"repeatEndDayOfMonth"`
    /** イベントの繰り返し開始曜日 */
	RepeatBeginDayOfWeek *string   `json:"repeatBeginDayOfWeek"`
    /** イベントの繰り返し終了曜日 */
	RepeatEndDayOfWeek *string   `json:"repeatEndDayOfWeek"`
    /** イベントの繰り返し開始時間 */
	RepeatBeginHour *int32   `json:"repeatBeginHour"`
    /** イベントの繰り返し終了時間 */
	RepeatEndHour *int32   `json:"repeatEndHour"`
    /** イベントの開始トリガー名 */
	RelativeTriggerName *string   `json:"relativeTriggerName"`
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
	TriggerId *string   `json:"triggerId"`
    /** トリガーの名前 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
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
	EventId *string   `json:"eventId"`
    /** イベントの種類名 */
	Name *string   `json:"name"`
    /** イベントの種類のメタデータ */
	Metadata *string   `json:"metadata"`
    /** イベント期間の種類 */
	ScheduleType *string   `json:"scheduleType"`
    /** 繰り返しの種類 */
	RepeatType *string   `json:"repeatType"`
    /** イベントの開始日時 */
	AbsoluteBegin *int64   `json:"absoluteBegin"`
    /** イベントの終了日時 */
	AbsoluteEnd *int64   `json:"absoluteEnd"`
    /** イベントの繰り返し開始日 */
	RepeatBeginDayOfMonth *int32   `json:"repeatBeginDayOfMonth"`
    /** イベントの繰り返し終了日 */
	RepeatEndDayOfMonth *int32   `json:"repeatEndDayOfMonth"`
    /** イベントの繰り返し開始曜日 */
	RepeatBeginDayOfWeek *string   `json:"repeatBeginDayOfWeek"`
    /** イベントの繰り返し終了曜日 */
	RepeatEndDayOfWeek *string   `json:"repeatEndDayOfWeek"`
    /** イベントの繰り返し開始時間 */
	RepeatBeginHour *int32   `json:"repeatBeginHour"`
    /** イベントの繰り返し終了時間 */
	RepeatEndHour *int32   `json:"repeatEndHour"`
    /** イベントの開始トリガー */
	RelativeTriggerName *string   `json:"relativeTriggerName"`
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
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** マスターデータ */
	Settings *string   `json:"settings"`
}

func (p *CurrentEventMaster) ToDict() *map[string]interface{} {
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
