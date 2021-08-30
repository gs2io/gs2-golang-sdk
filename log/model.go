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

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId         *string `json:"namespaceId"`
	Name                *string `json:"name"`
	Description         *string `json:"description"`
	Type                *string `json:"type"`
	GcpCredentialJson   *string `json:"gcpCredentialJson"`
	BigQueryDatasetName *string `json:"bigQueryDatasetName"`
	LogExpireDays       *int32  `json:"logExpireDays"`
	AwsRegion           *string `json:"awsRegion"`
	AwsAccessKeyId      *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey  *string `json:"awsSecretAccessKey"`
	FirehoseStreamName  *string `json:"firehoseStreamName"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:         core.CastString(data["namespaceId"]),
		Name:                core.CastString(data["name"]),
		Description:         core.CastString(data["description"]),
		Type:                core.CastString(data["type"]),
		GcpCredentialJson:   core.CastString(data["gcpCredentialJson"]),
		BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
		LogExpireDays:       core.CastInt32(data["logExpireDays"]),
		AwsRegion:           core.CastString(data["awsRegion"]),
		AwsAccessKeyId:      core.CastString(data["awsAccessKeyId"]),
		AwsSecretAccessKey:  core.CastString(data["awsSecretAccessKey"]),
		FirehoseStreamName:  core.CastString(data["firehoseStreamName"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":         p.NamespaceId,
		"name":                p.Name,
		"description":         p.Description,
		"type":                p.Type,
		"gcpCredentialJson":   p.GcpCredentialJson,
		"bigQueryDatasetName": p.BigQueryDatasetName,
		"logExpireDays":       p.LogExpireDays,
		"awsRegion":           p.AwsRegion,
		"awsAccessKeyId":      p.AwsAccessKeyId,
		"awsSecretAccessKey":  p.AwsSecretAccessKey,
		"firehoseStreamName":  p.FirehoseStreamName,
		"createdAt":           p.CreatedAt,
		"updatedAt":           p.UpdatedAt,
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
	Timestamp *int64  `json:"timestamp"`
	RequestId *string `json:"requestId"`
	Service   *string `json:"service"`
	Method    *string `json:"method"`
	UserId    *string `json:"userId"`
	Request   *string `json:"request"`
	Result    *string `json:"result"`
}

func NewAccessLogFromDict(data map[string]interface{}) AccessLog {
	return AccessLog{
		Timestamp: core.CastInt64(data["timestamp"]),
		RequestId: core.CastString(data["requestId"]),
		Service:   core.CastString(data["service"]),
		Method:    core.CastString(data["method"]),
		UserId:    core.CastString(data["userId"]),
		Request:   core.CastString(data["request"]),
		Result:    core.CastString(data["result"]),
	}
}

func (p AccessLog) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"timestamp": p.Timestamp,
		"requestId": p.RequestId,
		"service":   p.Service,
		"method":    p.Method,
		"userId":    p.UserId,
		"request":   p.Request,
		"result":    p.Result,
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
	Method  *string `json:"method"`
	UserId  *string `json:"userId"`
	Count   *int64  `json:"count"`
}

func NewAccessLogCountFromDict(data map[string]interface{}) AccessLogCount {
	return AccessLogCount{
		Service: core.CastString(data["service"]),
		Method:  core.CastString(data["method"]),
		UserId:  core.CastString(data["userId"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p AccessLogCount) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"service": p.Service,
		"method":  p.Method,
		"userId":  p.UserId,
		"count":   p.Count,
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
	Timestamp     *int64  `json:"timestamp"`
	TransactionId *string `json:"transactionId"`
	Service       *string `json:"service"`
	Method        *string `json:"method"`
	UserId        *string `json:"userId"`
	Action        *string `json:"action"`
	Args          *string `json:"args"`
	Tasks         *string `json:"tasks"`
}

func NewIssueStampSheetLogFromDict(data map[string]interface{}) IssueStampSheetLog {
	return IssueStampSheetLog{
		Timestamp:     core.CastInt64(data["timestamp"]),
		TransactionId: core.CastString(data["transactionId"]),
		Service:       core.CastString(data["service"]),
		Method:        core.CastString(data["method"]),
		UserId:        core.CastString(data["userId"]),
		Action:        core.CastString(data["action"]),
		Args:          core.CastString(data["args"]),
		Tasks:         core.CastString(data["tasks"]),
	}
}

func (p IssueStampSheetLog) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"timestamp":     p.Timestamp,
		"transactionId": p.TransactionId,
		"service":       p.Service,
		"method":        p.Method,
		"userId":        p.UserId,
		"action":        p.Action,
		"args":          p.Args,
		"tasks":         p.Tasks,
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
	Method  *string `json:"method"`
	UserId  *string `json:"userId"`
	Action  *string `json:"action"`
	Count   *int64  `json:"count"`
}

func NewIssueStampSheetLogCountFromDict(data map[string]interface{}) IssueStampSheetLogCount {
	return IssueStampSheetLogCount{
		Service: core.CastString(data["service"]),
		Method:  core.CastString(data["method"]),
		UserId:  core.CastString(data["userId"]),
		Action:  core.CastString(data["action"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p IssueStampSheetLogCount) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"service": p.Service,
		"method":  p.Method,
		"userId":  p.UserId,
		"action":  p.Action,
		"count":   p.Count,
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
	Timestamp     *int64  `json:"timestamp"`
	TransactionId *string `json:"transactionId"`
	Service       *string `json:"service"`
	Method        *string `json:"method"`
	UserId        *string `json:"userId"`
	Action        *string `json:"action"`
	Args          *string `json:"args"`
}

func NewExecuteStampSheetLogFromDict(data map[string]interface{}) ExecuteStampSheetLog {
	return ExecuteStampSheetLog{
		Timestamp:     core.CastInt64(data["timestamp"]),
		TransactionId: core.CastString(data["transactionId"]),
		Service:       core.CastString(data["service"]),
		Method:        core.CastString(data["method"]),
		UserId:        core.CastString(data["userId"]),
		Action:        core.CastString(data["action"]),
		Args:          core.CastString(data["args"]),
	}
}

func (p ExecuteStampSheetLog) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"timestamp":     p.Timestamp,
		"transactionId": p.TransactionId,
		"service":       p.Service,
		"method":        p.Method,
		"userId":        p.UserId,
		"action":        p.Action,
		"args":          p.Args,
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
	Method  *string `json:"method"`
	UserId  *string `json:"userId"`
	Action  *string `json:"action"`
	Count   *int64  `json:"count"`
}

func NewExecuteStampSheetLogCountFromDict(data map[string]interface{}) ExecuteStampSheetLogCount {
	return ExecuteStampSheetLogCount{
		Service: core.CastString(data["service"]),
		Method:  core.CastString(data["method"]),
		UserId:  core.CastString(data["userId"]),
		Action:  core.CastString(data["action"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p ExecuteStampSheetLogCount) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"service": p.Service,
		"method":  p.Method,
		"userId":  p.UserId,
		"action":  p.Action,
		"count":   p.Count,
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
	Timestamp *int64  `json:"timestamp"`
	TaskId    *string `json:"taskId"`
	Service   *string `json:"service"`
	Method    *string `json:"method"`
	UserId    *string `json:"userId"`
	Action    *string `json:"action"`
	Args      *string `json:"args"`
}

func NewExecuteStampTaskLogFromDict(data map[string]interface{}) ExecuteStampTaskLog {
	return ExecuteStampTaskLog{
		Timestamp: core.CastInt64(data["timestamp"]),
		TaskId:    core.CastString(data["taskId"]),
		Service:   core.CastString(data["service"]),
		Method:    core.CastString(data["method"]),
		UserId:    core.CastString(data["userId"]),
		Action:    core.CastString(data["action"]),
		Args:      core.CastString(data["args"]),
	}
}

func (p ExecuteStampTaskLog) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"timestamp": p.Timestamp,
		"taskId":    p.TaskId,
		"service":   p.Service,
		"method":    p.Method,
		"userId":    p.UserId,
		"action":    p.Action,
		"args":      p.Args,
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
	Method  *string `json:"method"`
	UserId  *string `json:"userId"`
	Action  *string `json:"action"`
	Count   *int64  `json:"count"`
}

func NewExecuteStampTaskLogCountFromDict(data map[string]interface{}) ExecuteStampTaskLogCount {
	return ExecuteStampTaskLogCount{
		Service: core.CastString(data["service"]),
		Method:  core.CastString(data["method"]),
		UserId:  core.CastString(data["userId"]),
		Action:  core.CastString(data["action"]),
		Count:   core.CastInt64(data["count"]),
	}
}

func (p ExecuteStampTaskLogCount) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"service": p.Service,
		"method":  p.Method,
		"userId":  p.UserId,
		"action":  p.Action,
		"count":   p.Count,
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
