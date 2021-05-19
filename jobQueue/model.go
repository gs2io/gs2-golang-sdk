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

package jobQueue

type Namespace struct {
    /** ネームスペース */
	NamespaceId *string   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *string   `json:"ownerId"`
    /** ネームスペース名 */
	Name *string   `json:"name"`
    /** ネームスペースの説明 */
	Description *string   `json:"description"`
    /** ジョブキューにジョブが登録されたときののプッシュ通知 */
	PushNotification *NotificationSetting   `json:"pushNotification"`
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
    if p.PushNotification != nil {
        data["pushNotification"] = *p.PushNotification.ToDict()
    }
    if p.LogSetting != nil {
        data["logSetting"] = *p.LogSetting.ToDict()
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type Job struct {
    /** ジョブ */
	JobId *string   `json:"jobId"`
    /** ジョブの名前 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ジョブの実行に使用するスクリプト のGRN */
	ScriptId *string   `json:"scriptId"`
    /** 引数 */
	Args *string   `json:"args"`
    /** 現在のリトライ回数 */
	CurrentRetryCount *int32   `json:"currentRetryCount"`
    /** 最大試行回数 */
	MaxTryCount *int32   `json:"maxTryCount"`
    /** ソート用インデックス(現在時刻(ミリ秒).登録時のインデックス) */
	Index *float64   `json:"index"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *Job) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["jobId"] = p.JobId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["scriptId"] = p.ScriptId
    data["args"] = p.Args
    data["currentRetryCount"] = p.CurrentRetryCount
    data["maxTryCount"] = p.MaxTryCount
    data["index"] = p.Index
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type JobResult struct {
    /** ジョブ実行結果 */
	JobResultId *string   `json:"jobResultId"`
    /** ジョブ */
	JobId *string   `json:"jobId"`
    /** None */
	TryNumber *int32   `json:"tryNumber"`
    /** None */
	StatusCode *int32   `json:"statusCode"`
    /** レスポンスの内容 */
	Result *string   `json:"result"`
    /** 作成日時 */
	TryAt *int64   `json:"tryAt"`
}

func (p *JobResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["jobResultId"] = p.JobResultId
    data["jobId"] = p.JobId
    data["tryNumber"] = p.TryNumber
    data["statusCode"] = p.StatusCode
    data["result"] = p.Result
    data["tryAt"] = p.TryAt
    return &data
}

type DeadLetterJob struct {
    /** デッドレタージョブ */
	DeadLetterJobId *string   `json:"deadLetterJobId"`
    /** ジョブの名前 */
	Name *string   `json:"name"`
    /** ユーザーID */
	UserId *string   `json:"userId"`
    /** ジョブの実行に使用するスクリプト のGRN */
	ScriptId *string   `json:"scriptId"`
    /** 引数 */
	Args *string   `json:"args"`
    /** ジョブ実行結果 */
	Result []JobResultBody   `json:"result"`
    /** 作成日時 */
	CreatedAt *int64   `json:"createdAt"`
    /** 最終更新日時 */
	UpdatedAt *int64   `json:"updatedAt"`
}

func (p *DeadLetterJob) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["deadLetterJobId"] = p.DeadLetterJobId
    data["name"] = p.Name
    data["userId"] = p.UserId
    data["scriptId"] = p.ScriptId
    data["args"] = p.Args
    if p.Result != nil {
        var _result []*map[string]interface {}
        for _, item := range p.Result {
            _result = append(_result, item.ToDict())
        }
        data["result"] = &_result
    }
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
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

type NotificationSetting struct {
    /** プッシュ通知に使用する GS2-Gateway のネームスペース のGRN */
	GatewayNamespaceId *string   `json:"gatewayNamespaceId"`
    /** モバイルプッシュ通知へ転送するか */
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
    /** モバイルプッシュ通知で使用するサウンドファイル名 */
	Sound *string   `json:"sound"`
}

func (p *NotificationSetting) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["gatewayNamespaceId"] = p.GatewayNamespaceId
    data["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
    data["sound"] = p.Sound
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

type JobEntry struct {
    /** スクリプト のGRN */
	ScriptId *string   `json:"scriptId"`
    /** 引数 */
	Args *string   `json:"args"`
    /** 最大試行回数 */
	MaxTryCount *int32   `json:"maxTryCount"`
}

func (p *JobEntry) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["scriptId"] = p.ScriptId
    data["args"] = p.Args
    data["maxTryCount"] = p.MaxTryCount
    return &data
}

type JobResultBody struct {
    /** 試行回数 */
	TryNumber *int32   `json:"tryNumber"`
    /** ステータスコード */
	StatusCode *int32   `json:"statusCode"`
    /** レスポンスの内容 */
	Result *string   `json:"result"`
    /** 実行日時 */
	TryAt *int64   `json:"tryAt"`
}

func (p *JobResultBody) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["tryNumber"] = p.TryNumber
    data["statusCode"] = p.StatusCode
    data["result"] = p.Result
    data["tryAt"] = p.TryAt
    return &data
}
