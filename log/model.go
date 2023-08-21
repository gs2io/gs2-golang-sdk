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
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Type *string `json:"type"`
	GcpCredentialJson *string `json:"gcpCredentialJson"`
	BigQueryDatasetName *string `json:"bigQueryDatasetName"`
	LogExpireDays *int32 `json:"logExpireDays"`
	AwsRegion *string `json:"awsRegion"`
	AwsAccessKeyId *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey *string `json:"awsSecretAccessKey"`
	FirehoseStreamName *string `json:"firehoseStreamName"`
	Status *string `json:"status"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
    return Namespace {
        NamespaceId: core.CastString(data["namespaceId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Type: core.CastString(data["type"]),
        GcpCredentialJson: core.CastString(data["gcpCredentialJson"]),
        BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
        LogExpireDays: core.CastInt32(data["logExpireDays"]),
        AwsRegion: core.CastString(data["awsRegion"]),
        AwsAccessKeyId: core.CastString(data["awsAccessKeyId"]),
        AwsSecretAccessKey: core.CastString(data["awsSecretAccessKey"]),
        FirehoseStreamName: core.CastString(data["firehoseStreamName"]),
        Status: core.CastString(data["status"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Namespace) ToDict() map[string]interface{} {
    
    var namespaceId *string
    if p.NamespaceId != nil {
        namespaceId = p.NamespaceId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var _type *string
    if p.Type != nil {
        _type = p.Type
    }
    var gcpCredentialJson *string
    if p.GcpCredentialJson != nil {
        gcpCredentialJson = p.GcpCredentialJson
    }
    var bigQueryDatasetName *string
    if p.BigQueryDatasetName != nil {
        bigQueryDatasetName = p.BigQueryDatasetName
    }
    var logExpireDays *int32
    if p.LogExpireDays != nil {
        logExpireDays = p.LogExpireDays
    }
    var awsRegion *string
    if p.AwsRegion != nil {
        awsRegion = p.AwsRegion
    }
    var awsAccessKeyId *string
    if p.AwsAccessKeyId != nil {
        awsAccessKeyId = p.AwsAccessKeyId
    }
    var awsSecretAccessKey *string
    if p.AwsSecretAccessKey != nil {
        awsSecretAccessKey = p.AwsSecretAccessKey
    }
    var firehoseStreamName *string
    if p.FirehoseStreamName != nil {
        firehoseStreamName = p.FirehoseStreamName
    }
    var status *string
    if p.Status != nil {
        status = p.Status
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "type": _type,
        "gcpCredentialJson": gcpCredentialJson,
        "bigQueryDatasetName": bigQueryDatasetName,
        "logExpireDays": logExpireDays,
        "awsRegion": awsRegion,
        "awsAccessKeyId": awsAccessKeyId,
        "awsSecretAccessKey": awsSecretAccessKey,
        "firehoseStreamName": firehoseStreamName,
        "status": status,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p Namespace) Pointer() *Namespace {
    return &p
}

func CastNamespaces(data []interface{}) []Namespace {
	v := make([]Namespace, 0)
	for _, d := range data {
		v = append(v, NewNamespaceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNamespacesFromDict(data []Namespace) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type AccessLog struct {
	Timestamp *int64 `json:"timestamp"`
	RequestId *string `json:"requestId"`
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Request *string `json:"request"`
	Result *string `json:"result"`
}

func NewAccessLogFromJson(data string) AccessLog {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAccessLogFromDict(dict)
}

func NewAccessLogFromDict(data map[string]interface{}) AccessLog {
    return AccessLog {
        Timestamp: core.CastInt64(data["timestamp"]),
        RequestId: core.CastString(data["requestId"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Request: core.CastString(data["request"]),
        Result: core.CastString(data["result"]),
    }
}

func (p AccessLog) ToDict() map[string]interface{} {
    
    var timestamp *int64
    if p.Timestamp != nil {
        timestamp = p.Timestamp
    }
    var requestId *string
    if p.RequestId != nil {
        requestId = p.RequestId
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var request *string
    if p.Request != nil {
        request = p.Request
    }
    var result *string
    if p.Result != nil {
        result = p.Result
    }
    return map[string]interface{} {
        "timestamp": timestamp,
        "requestId": requestId,
        "service": service,
        "method": method,
        "userId": userId,
        "request": request,
        "result": result,
    }
}

func (p AccessLog) Pointer() *AccessLog {
    return &p
}

func CastAccessLogs(data []interface{}) []AccessLog {
	v := make([]AccessLog, 0)
	for _, d := range data {
		v = append(v, NewAccessLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccessLogsFromDict(data []AccessLog) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type AccessLogCount struct {
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Count *int64 `json:"count"`
}

func NewAccessLogCountFromJson(data string) AccessLogCount {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAccessLogCountFromDict(dict)
}

func NewAccessLogCountFromDict(data map[string]interface{}) AccessLogCount {
    return AccessLogCount {
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Count: core.CastInt64(data["count"]),
    }
}

func (p AccessLogCount) ToDict() map[string]interface{} {
    
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var count *int64
    if p.Count != nil {
        count = p.Count
    }
    return map[string]interface{} {
        "service": service,
        "method": method,
        "userId": userId,
        "count": count,
    }
}

func (p AccessLogCount) Pointer() *AccessLogCount {
    return &p
}

func CastAccessLogCounts(data []interface{}) []AccessLogCount {
	v := make([]AccessLogCount, 0)
	for _, d := range data {
		v = append(v, NewAccessLogCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccessLogCountsFromDict(data []AccessLogCount) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type IssueStampSheetLog struct {
	Timestamp *int64 `json:"timestamp"`
	TransactionId *string `json:"transactionId"`
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Args *string `json:"args"`
	Tasks []*string `json:"tasks"`
}

func NewIssueStampSheetLogFromJson(data string) IssueStampSheetLog {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssueStampSheetLogFromDict(dict)
}

func NewIssueStampSheetLogFromDict(data map[string]interface{}) IssueStampSheetLog {
    return IssueStampSheetLog {
        Timestamp: core.CastInt64(data["timestamp"]),
        TransactionId: core.CastString(data["transactionId"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Args: core.CastString(data["args"]),
        Tasks: core.CastStrings(core.CastArray(data["tasks"])),
    }
}

func (p IssueStampSheetLog) ToDict() map[string]interface{} {
    
    var timestamp *int64
    if p.Timestamp != nil {
        timestamp = p.Timestamp
    }
    var transactionId *string
    if p.TransactionId != nil {
        transactionId = p.TransactionId
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    var tasks []interface{}
    if p.Tasks != nil {
        tasks = core.CastStringsFromDict(
            p.Tasks,
        )
    }
    return map[string]interface{} {
        "timestamp": timestamp,
        "transactionId": transactionId,
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "args": args,
        "tasks": tasks,
    }
}

func (p IssueStampSheetLog) Pointer() *IssueStampSheetLog {
    return &p
}

func CastIssueStampSheetLogs(data []interface{}) []IssueStampSheetLog {
	v := make([]IssueStampSheetLog, 0)
	for _, d := range data {
		v = append(v, NewIssueStampSheetLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIssueStampSheetLogsFromDict(data []IssueStampSheetLog) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type IssueStampSheetLogCount struct {
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Count *int64 `json:"count"`
}

func NewIssueStampSheetLogCountFromJson(data string) IssueStampSheetLogCount {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssueStampSheetLogCountFromDict(dict)
}

func NewIssueStampSheetLogCountFromDict(data map[string]interface{}) IssueStampSheetLogCount {
    return IssueStampSheetLogCount {
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Count: core.CastInt64(data["count"]),
    }
}

func (p IssueStampSheetLogCount) ToDict() map[string]interface{} {
    
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var count *int64
    if p.Count != nil {
        count = p.Count
    }
    return map[string]interface{} {
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "count": count,
    }
}

func (p IssueStampSheetLogCount) Pointer() *IssueStampSheetLogCount {
    return &p
}

func CastIssueStampSheetLogCounts(data []interface{}) []IssueStampSheetLogCount {
	v := make([]IssueStampSheetLogCount, 0)
	for _, d := range data {
		v = append(v, NewIssueStampSheetLogCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastIssueStampSheetLogCountsFromDict(data []IssueStampSheetLogCount) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ExecuteStampSheetLog struct {
	Timestamp *int64 `json:"timestamp"`
	TransactionId *string `json:"transactionId"`
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Args *string `json:"args"`
}

func NewExecuteStampSheetLogFromJson(data string) ExecuteStampSheetLog {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExecuteStampSheetLogFromDict(dict)
}

func NewExecuteStampSheetLogFromDict(data map[string]interface{}) ExecuteStampSheetLog {
    return ExecuteStampSheetLog {
        Timestamp: core.CastInt64(data["timestamp"]),
        TransactionId: core.CastString(data["transactionId"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Args: core.CastString(data["args"]),
    }
}

func (p ExecuteStampSheetLog) ToDict() map[string]interface{} {
    
    var timestamp *int64
    if p.Timestamp != nil {
        timestamp = p.Timestamp
    }
    var transactionId *string
    if p.TransactionId != nil {
        transactionId = p.TransactionId
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    return map[string]interface{} {
        "timestamp": timestamp,
        "transactionId": transactionId,
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "args": args,
    }
}

func (p ExecuteStampSheetLog) Pointer() *ExecuteStampSheetLog {
    return &p
}

func CastExecuteStampSheetLogs(data []interface{}) []ExecuteStampSheetLog {
	v := make([]ExecuteStampSheetLog, 0)
	for _, d := range data {
		v = append(v, NewExecuteStampSheetLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExecuteStampSheetLogsFromDict(data []ExecuteStampSheetLog) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ExecuteStampSheetLogCount struct {
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Count *int64 `json:"count"`
}

func NewExecuteStampSheetLogCountFromJson(data string) ExecuteStampSheetLogCount {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExecuteStampSheetLogCountFromDict(dict)
}

func NewExecuteStampSheetLogCountFromDict(data map[string]interface{}) ExecuteStampSheetLogCount {
    return ExecuteStampSheetLogCount {
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Count: core.CastInt64(data["count"]),
    }
}

func (p ExecuteStampSheetLogCount) ToDict() map[string]interface{} {
    
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var count *int64
    if p.Count != nil {
        count = p.Count
    }
    return map[string]interface{} {
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "count": count,
    }
}

func (p ExecuteStampSheetLogCount) Pointer() *ExecuteStampSheetLogCount {
    return &p
}

func CastExecuteStampSheetLogCounts(data []interface{}) []ExecuteStampSheetLogCount {
	v := make([]ExecuteStampSheetLogCount, 0)
	for _, d := range data {
		v = append(v, NewExecuteStampSheetLogCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExecuteStampSheetLogCountsFromDict(data []ExecuteStampSheetLogCount) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ExecuteStampTaskLog struct {
	Timestamp *int64 `json:"timestamp"`
	TaskId *string `json:"taskId"`
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Args *string `json:"args"`
}

func NewExecuteStampTaskLogFromJson(data string) ExecuteStampTaskLog {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExecuteStampTaskLogFromDict(dict)
}

func NewExecuteStampTaskLogFromDict(data map[string]interface{}) ExecuteStampTaskLog {
    return ExecuteStampTaskLog {
        Timestamp: core.CastInt64(data["timestamp"]),
        TaskId: core.CastString(data["taskId"]),
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Args: core.CastString(data["args"]),
    }
}

func (p ExecuteStampTaskLog) ToDict() map[string]interface{} {
    
    var timestamp *int64
    if p.Timestamp != nil {
        timestamp = p.Timestamp
    }
    var taskId *string
    if p.TaskId != nil {
        taskId = p.TaskId
    }
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    return map[string]interface{} {
        "timestamp": timestamp,
        "taskId": taskId,
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "args": args,
    }
}

func (p ExecuteStampTaskLog) Pointer() *ExecuteStampTaskLog {
    return &p
}

func CastExecuteStampTaskLogs(data []interface{}) []ExecuteStampTaskLog {
	v := make([]ExecuteStampTaskLog, 0)
	for _, d := range data {
		v = append(v, NewExecuteStampTaskLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExecuteStampTaskLogsFromDict(data []ExecuteStampTaskLog) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type ExecuteStampTaskLogCount struct {
	Service *string `json:"service"`
	Method *string `json:"method"`
	UserId *string `json:"userId"`
	Action *string `json:"action"`
	Count *int64 `json:"count"`
}

func NewExecuteStampTaskLogCountFromJson(data string) ExecuteStampTaskLogCount {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExecuteStampTaskLogCountFromDict(dict)
}

func NewExecuteStampTaskLogCountFromDict(data map[string]interface{}) ExecuteStampTaskLogCount {
    return ExecuteStampTaskLogCount {
        Service: core.CastString(data["service"]),
        Method: core.CastString(data["method"]),
        UserId: core.CastString(data["userId"]),
        Action: core.CastString(data["action"]),
        Count: core.CastInt64(data["count"]),
    }
}

func (p ExecuteStampTaskLogCount) ToDict() map[string]interface{} {
    
    var service *string
    if p.Service != nil {
        service = p.Service
    }
    var method *string
    if p.Method != nil {
        method = p.Method
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var action *string
    if p.Action != nil {
        action = p.Action
    }
    var count *int64
    if p.Count != nil {
        count = p.Count
    }
    return map[string]interface{} {
        "service": service,
        "method": method,
        "userId": userId,
        "action": action,
        "count": count,
    }
}

func (p ExecuteStampTaskLogCount) Pointer() *ExecuteStampTaskLogCount {
    return &p
}

func CastExecuteStampTaskLogCounts(data []interface{}) []ExecuteStampTaskLogCount {
	v := make([]ExecuteStampTaskLogCount, 0)
	for _, d := range data {
		v = append(v, NewExecuteStampTaskLogCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastExecuteStampTaskLogCountsFromDict(data []ExecuteStampTaskLogCount) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Insight struct {
	InsightId *string `json:"insightId"`
	Name *string `json:"name"`
	TaskId *string `json:"taskId"`
	Host *string `json:"host"`
	Password *string `json:"password"`
	Status *string `json:"status"`
	CreatedAt *int64 `json:"createdAt"`
	Revision *int64 `json:"revision"`
}

func NewInsightFromJson(data string) Insight {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewInsightFromDict(dict)
}

func NewInsightFromDict(data map[string]interface{}) Insight {
    return Insight {
        InsightId: core.CastString(data["insightId"]),
        Name: core.CastString(data["name"]),
        TaskId: core.CastString(data["taskId"]),
        Host: core.CastString(data["host"]),
        Password: core.CastString(data["password"]),
        Status: core.CastString(data["status"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Insight) ToDict() map[string]interface{} {
    
    var insightId *string
    if p.InsightId != nil {
        insightId = p.InsightId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var taskId *string
    if p.TaskId != nil {
        taskId = p.TaskId
    }
    var host *string
    if p.Host != nil {
        host = p.Host
    }
    var password *string
    if p.Password != nil {
        password = p.Password
    }
    var status *string
    if p.Status != nil {
        status = p.Status
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "insightId": insightId,
        "name": name,
        "taskId": taskId,
        "host": host,
        "password": password,
        "status": status,
        "createdAt": createdAt,
        "revision": revision,
    }
}

func (p Insight) Pointer() *Insight {
    return &p
}

func CastInsights(data []interface{}) []Insight {
	v := make([]Insight, 0)
	for _, d := range data {
		v = append(v, NewInsightFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInsightsFromDict(data []Insight) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}