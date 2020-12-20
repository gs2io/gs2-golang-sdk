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

package log

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
    /** ネームスペース */
	NamespaceId *core.String   `json:"namespaceId"`
    /** オーナーID */
	OwnerId *core.String   `json:"ownerId"`
    /** カテゴリー名 */
	Name *core.String   `json:"name"`
    /** ネームスペースの説明 */
	Description *core.String   `json:"description"`
    /** ログの書き出し方法 */
	Type *core.String   `json:"type"`
    /** GCPのクレデンシャル */
	GcpCredentialJson *core.String   `json:"gcpCredentialJson"`
    /** BigQueryのデータセット名 */
	BigQueryDatasetName *core.String   `json:"bigQueryDatasetName"`
    /** ログの保存期間(日) */
	LogExpireDays *int32   `json:"logExpireDays"`
    /** AWSのリージョン */
	AwsRegion *core.String   `json:"awsRegion"`
    /** AWSのアクセスキーID */
	AwsAccessKeyId *core.String   `json:"awsAccessKeyId"`
    /** AWSのシークレットアクセスキー */
	AwsSecretAccessKey *core.String   `json:"awsSecretAccessKey"`
    /** Kinesis Firehose のストリーム名 */
	FirehoseStreamName *core.String   `json:"firehoseStreamName"`
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
    data["type"] = p.Type
    data["gcpCredentialJson"] = p.GcpCredentialJson
    data["bigQueryDatasetName"] = p.BigQueryDatasetName
    data["logExpireDays"] = p.LogExpireDays
    data["awsRegion"] = p.AwsRegion
    data["awsAccessKeyId"] = p.AwsAccessKeyId
    data["awsSecretAccessKey"] = p.AwsSecretAccessKey
    data["firehoseStreamName"] = p.FirehoseStreamName
    data["status"] = p.Status
    data["createdAt"] = p.CreatedAt
    data["updatedAt"] = p.UpdatedAt
    return &data
}

type AccessLog struct {
    /** 日時 */
	Timestamp *int64   `json:"timestamp"`
    /** リクエストID */
	RequestId *core.String   `json:"requestId"`
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** リクエストパラメータ */
	Request *core.String   `json:"request"`
    /** 応答内容 */
	Result *core.String   `json:"result"`
}

func (p *AccessLog) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["timestamp"] = p.Timestamp
    data["requestId"] = p.RequestId
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["request"] = p.Request
    data["result"] = p.Result
    return &data
}

type AccessLogCount struct {
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 回数 */
	Count *int64   `json:"count"`
}

func (p *AccessLogCount) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["count"] = p.Count
    return &data
}

type IssueStampSheetLog struct {
    /** 日時 */
	Timestamp *int64   `json:"timestamp"`
    /** トランザクションID */
	TransactionId *core.String   `json:"transactionId"`
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 引数 */
	Args *core.String   `json:"args"`
    /** スタンプタスク */
	Tasks *core.String   `json:"tasks"`
}

func (p *IssueStampSheetLog) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["timestamp"] = p.Timestamp
    data["transactionId"] = p.TransactionId
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["args"] = p.Args
    data["tasks"] = p.Tasks
    return &data
}

type IssueStampSheetLogCount struct {
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 回数 */
	Count *int64   `json:"count"`
}

func (p *IssueStampSheetLogCount) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["count"] = p.Count
    return &data
}

type ExecuteStampSheetLog struct {
    /** 日時 */
	Timestamp *int64   `json:"timestamp"`
    /** トランザクションID */
	TransactionId *core.String   `json:"transactionId"`
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 引数 */
	Args *core.String   `json:"args"`
}

func (p *ExecuteStampSheetLog) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["timestamp"] = p.Timestamp
    data["transactionId"] = p.TransactionId
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["args"] = p.Args
    return &data
}

type ExecuteStampSheetLogCount struct {
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 回数 */
	Count *int64   `json:"count"`
}

func (p *ExecuteStampSheetLogCount) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["count"] = p.Count
    return &data
}

type ExecuteStampTaskLog struct {
    /** 日時 */
	Timestamp *int64   `json:"timestamp"`
    /** タスクID */
	TaskId *core.String   `json:"taskId"`
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 引数 */
	Args *core.String   `json:"args"`
}

func (p *ExecuteStampTaskLog) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["timestamp"] = p.Timestamp
    data["taskId"] = p.TaskId
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["args"] = p.Args
    return &data
}

type ExecuteStampTaskLogCount struct {
    /** マイクロサービスの種類 */
	Service *core.String   `json:"service"`
    /** マイクロサービスのメソッド */
	Method *core.String   `json:"method"`
    /** ユーザーID */
	UserId *core.String   `json:"userId"`
    /** 報酬アクション */
	Action *core.String   `json:"action"`
    /** 回数 */
	Count *int64   `json:"count"`
}

func (p *ExecuteStampTaskLogCount) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    data["service"] = p.Service
    data["method"] = p.Method
    data["userId"] = p.UserId
    data["action"] = p.Action
    data["count"] = p.Count
    return &data
}
