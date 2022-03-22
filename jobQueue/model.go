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

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	PushNotification *NotificationSetting `json:"pushNotification"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
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
        PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    var pushNotification map[string]interface{}
    if p.PushNotification != nil {
        pushNotification = p.PushNotification.ToDict()
    }
    var logSetting map[string]interface{}
    if p.LogSetting != nil {
        logSetting = p.LogSetting.ToDict()
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "pushNotification": pushNotification,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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

type Job struct {
	JobId *string `json:"jobId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	ScriptId *string `json:"scriptId"`
	Args *string `json:"args"`
	CurrentRetryCount *int32 `json:"currentRetryCount"`
	MaxTryCount *int32 `json:"maxTryCount"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewJobFromJson(data string) Job {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewJobFromDict(dict)
}

func NewJobFromDict(data map[string]interface{}) Job {
    return Job {
        JobId: core.CastString(data["jobId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        ScriptId: core.CastString(data["scriptId"]),
        Args: core.CastString(data["args"]),
        CurrentRetryCount: core.CastInt32(data["currentRetryCount"]),
        MaxTryCount: core.CastInt32(data["maxTryCount"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p Job) ToDict() map[string]interface{} {
    
    var jobId *string
    if p.JobId != nil {
        jobId = p.JobId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var scriptId *string
    if p.ScriptId != nil {
        scriptId = p.ScriptId
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    var currentRetryCount *int32
    if p.CurrentRetryCount != nil {
        currentRetryCount = p.CurrentRetryCount
    }
    var maxTryCount *int32
    if p.MaxTryCount != nil {
        maxTryCount = p.MaxTryCount
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "jobId": jobId,
        "name": name,
        "userId": userId,
        "scriptId": scriptId,
        "args": args,
        "currentRetryCount": currentRetryCount,
        "maxTryCount": maxTryCount,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p Job) Pointer() *Job {
    return &p
}

func CastJobs(data []interface{}) []Job {
	v := make([]Job, 0)
	for _, d := range data {
		v = append(v, NewJobFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJobsFromDict(data []Job) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type JobResult struct {
	JobResultId *string `json:"jobResultId"`
	JobId *string `json:"jobId"`
	TryNumber *int32 `json:"tryNumber"`
	StatusCode *int32 `json:"statusCode"`
	Result *string `json:"result"`
	TryAt *int64 `json:"tryAt"`
}

func NewJobResultFromJson(data string) JobResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewJobResultFromDict(dict)
}

func NewJobResultFromDict(data map[string]interface{}) JobResult {
    return JobResult {
        JobResultId: core.CastString(data["jobResultId"]),
        JobId: core.CastString(data["jobId"]),
        TryNumber: core.CastInt32(data["tryNumber"]),
        StatusCode: core.CastInt32(data["statusCode"]),
        Result: core.CastString(data["result"]),
        TryAt: core.CastInt64(data["tryAt"]),
    }
}

func (p JobResult) ToDict() map[string]interface{} {
    
    var jobResultId *string
    if p.JobResultId != nil {
        jobResultId = p.JobResultId
    }
    var jobId *string
    if p.JobId != nil {
        jobId = p.JobId
    }
    var tryNumber *int32
    if p.TryNumber != nil {
        tryNumber = p.TryNumber
    }
    var statusCode *int32
    if p.StatusCode != nil {
        statusCode = p.StatusCode
    }
    var result *string
    if p.Result != nil {
        result = p.Result
    }
    var tryAt *int64
    if p.TryAt != nil {
        tryAt = p.TryAt
    }
    return map[string]interface{} {
        "jobResultId": jobResultId,
        "jobId": jobId,
        "tryNumber": tryNumber,
        "statusCode": statusCode,
        "result": result,
        "tryAt": tryAt,
    }
}

func (p JobResult) Pointer() *JobResult {
    return &p
}

func CastJobResults(data []interface{}) []JobResult {
	v := make([]JobResult, 0)
	for _, d := range data {
		v = append(v, NewJobResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJobResultsFromDict(data []JobResult) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type DeadLetterJob struct {
	DeadLetterJobId *string `json:"deadLetterJobId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	ScriptId *string `json:"scriptId"`
	Args *string `json:"args"`
	Result []JobResultBody `json:"result"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewDeadLetterJobFromJson(data string) DeadLetterJob {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeadLetterJobFromDict(dict)
}

func NewDeadLetterJobFromDict(data map[string]interface{}) DeadLetterJob {
    return DeadLetterJob {
        DeadLetterJobId: core.CastString(data["deadLetterJobId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        ScriptId: core.CastString(data["scriptId"]),
        Args: core.CastString(data["args"]),
        Result: CastJobResultBodies(core.CastArray(data["result"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p DeadLetterJob) ToDict() map[string]interface{} {
    
    var deadLetterJobId *string
    if p.DeadLetterJobId != nil {
        deadLetterJobId = p.DeadLetterJobId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var scriptId *string
    if p.ScriptId != nil {
        scriptId = p.ScriptId
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    var result []interface{}
    if p.Result != nil {
        result = CastJobResultBodiesFromDict(
            p.Result,
        )
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "deadLetterJobId": deadLetterJobId,
        "name": name,
        "userId": userId,
        "scriptId": scriptId,
        "args": args,
        "result": result,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p DeadLetterJob) Pointer() *DeadLetterJob {
    return &p
}

func CastDeadLetterJobs(data []interface{}) []DeadLetterJob {
	v := make([]DeadLetterJob, 0)
	for _, d := range data {
		v = append(v, NewDeadLetterJobFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDeadLetterJobsFromDict(data []DeadLetterJob) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type NotificationSetting struct {
	GatewayNamespaceId *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool `json:"enableTransferMobileNotification"`
	Sound *string `json:"sound"`
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNotificationSettingFromDict(dict)
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
    return NotificationSetting {
        GatewayNamespaceId: core.CastString(data["gatewayNamespaceId"]),
        EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
        Sound: core.CastString(data["sound"]),
    }
}

func (p NotificationSetting) ToDict() map[string]interface{} {
    
    var gatewayNamespaceId *string
    if p.GatewayNamespaceId != nil {
        gatewayNamespaceId = p.GatewayNamespaceId
    }
    var enableTransferMobileNotification *bool
    if p.EnableTransferMobileNotification != nil {
        enableTransferMobileNotification = p.EnableTransferMobileNotification
    }
    var sound *string
    if p.Sound != nil {
        sound = p.Sound
    }
    return map[string]interface{} {
        "gatewayNamespaceId": gatewayNamespaceId,
        "enableTransferMobileNotification": enableTransferMobileNotification,
        "sound": sound,
    }
}

func (p NotificationSetting) Pointer() *NotificationSetting {
    return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func NewLogSettingFromJson(data string) LogSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    
    var loggingNamespaceId *string
    if p.LoggingNamespaceId != nil {
        loggingNamespaceId = p.LoggingNamespaceId
    }
    return map[string]interface{} {
        "loggingNamespaceId": loggingNamespaceId,
    }
}

func (p LogSetting) Pointer() *LogSetting {
    return &p
}

func CastLogSettings(data []interface{}) []LogSetting {
	v := make([]LogSetting, 0)
	for _, d := range data {
		v = append(v, NewLogSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogSettingsFromDict(data []LogSetting) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type JobEntry struct {
	ScriptId *string `json:"scriptId"`
	Args *string `json:"args"`
	MaxTryCount *int32 `json:"maxTryCount"`
}

func NewJobEntryFromJson(data string) JobEntry {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewJobEntryFromDict(dict)
}

func NewJobEntryFromDict(data map[string]interface{}) JobEntry {
    return JobEntry {
        ScriptId: core.CastString(data["scriptId"]),
        Args: core.CastString(data["args"]),
        MaxTryCount: core.CastInt32(data["maxTryCount"]),
    }
}

func (p JobEntry) ToDict() map[string]interface{} {
    
    var scriptId *string
    if p.ScriptId != nil {
        scriptId = p.ScriptId
    }
    var args *string
    if p.Args != nil {
        args = p.Args
    }
    var maxTryCount *int32
    if p.MaxTryCount != nil {
        maxTryCount = p.MaxTryCount
    }
    return map[string]interface{} {
        "scriptId": scriptId,
        "args": args,
        "maxTryCount": maxTryCount,
    }
}

func (p JobEntry) Pointer() *JobEntry {
    return &p
}

func CastJobEntries(data []interface{}) []JobEntry {
	v := make([]JobEntry, 0)
	for _, d := range data {
		v = append(v, NewJobEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJobEntriesFromDict(data []JobEntry) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type JobResultBody struct {
	TryNumber *int32 `json:"tryNumber"`
	StatusCode *int32 `json:"statusCode"`
	Result *string `json:"result"`
	TryAt *int64 `json:"tryAt"`
}

func NewJobResultBodyFromJson(data string) JobResultBody {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewJobResultBodyFromDict(dict)
}

func NewJobResultBodyFromDict(data map[string]interface{}) JobResultBody {
    return JobResultBody {
        TryNumber: core.CastInt32(data["tryNumber"]),
        StatusCode: core.CastInt32(data["statusCode"]),
        Result: core.CastString(data["result"]),
        TryAt: core.CastInt64(data["tryAt"]),
    }
}

func (p JobResultBody) ToDict() map[string]interface{} {
    
    var tryNumber *int32
    if p.TryNumber != nil {
        tryNumber = p.TryNumber
    }
    var statusCode *int32
    if p.StatusCode != nil {
        statusCode = p.StatusCode
    }
    var result *string
    if p.Result != nil {
        result = p.Result
    }
    var tryAt *int64
    if p.TryAt != nil {
        tryAt = p.TryAt
    }
    return map[string]interface{} {
        "tryNumber": tryNumber,
        "statusCode": statusCode,
        "result": result,
        "tryAt": tryAt,
    }
}

func (p JobResultBody) Pointer() *JobResultBody {
    return &p
}

func CastJobResultBodies(data []interface{}) []JobResultBody {
	v := make([]JobResultBody, 0)
	for _, d := range data {
		v = append(v, NewJobResultBodyFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastJobResultBodiesFromDict(data []JobResultBody) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}