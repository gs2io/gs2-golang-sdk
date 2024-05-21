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
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId      *string              `json:"namespaceId"`
	Name             *string              `json:"name"`
	Description      *string              `json:"description"`
	EnableAutoRun    *bool                `json:"enableAutoRun"`
	RunNotification  *NotificationSetting `json:"runNotification"`
	PushNotification *NotificationSetting `json:"pushNotification"`
	LogSetting       *LogSetting          `json:"logSetting"`
	CreatedAt        *int64               `json:"createdAt"`
	UpdatedAt        *int64               `json:"updatedAt"`
	Revision         *int64               `json:"revision"`
}

func (p *Namespace) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Namespace{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Namespace{}
	} else {
		*p = Namespace{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Description = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Description = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Description = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Description = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Description)
				}
			}
		}
		if v, ok := d["enableAutoRun"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAutoRun)
		}
		if v, ok := d["runNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RunNotification)
		}
		if v, ok := d["pushNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PushNotification)
		}
		if v, ok := d["logSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogSetting)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewNamespaceFromJson(data string) Namespace {
	req := Namespace{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:      core.CastString(data["namespaceId"]),
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		EnableAutoRun:    core.CastBool(data["enableAutoRun"]),
		RunNotification:  NewNotificationSettingFromDict(core.CastMap(data["runNotification"])).Pointer(),
		PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
		Revision:         core.CastInt64(data["revision"]),
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
	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var runNotification map[string]interface{}
	if p.RunNotification != nil {
		runNotification = p.RunNotification.ToDict()
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId":      namespaceId,
		"name":             name,
		"description":      description,
		"enableAutoRun":    enableAutoRun,
		"runNotification":  runNotification,
		"pushNotification": pushNotification,
		"logSetting":       logSetting,
		"createdAt":        createdAt,
		"updatedAt":        updatedAt,
		"revision":         revision,
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
	JobId             *string `json:"jobId"`
	Name              *string `json:"name"`
	UserId            *string `json:"userId"`
	ScriptId          *string `json:"scriptId"`
	Args              *string `json:"args"`
	CurrentRetryCount *int32  `json:"currentRetryCount"`
	MaxTryCount       *int32  `json:"maxTryCount"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
}

func (p *Job) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Job{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = Job{}
	} else {
		*p = Job{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["jobId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JobId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JobId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JobId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JobId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JobId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JobId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["scriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScriptId)
				}
			}
		}
		if v, ok := d["args"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Args = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Args = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Args = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Args)
				}
			}
		}
		if v, ok := d["currentRetryCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentRetryCount)
		}
		if v, ok := d["maxTryCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxTryCount)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewJobFromJson(data string) Job {
	req := Job{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJobFromDict(data map[string]interface{}) Job {
	return Job{
		JobId:             core.CastString(data["jobId"]),
		Name:              core.CastString(data["name"]),
		UserId:            core.CastString(data["userId"]),
		ScriptId:          core.CastString(data["scriptId"]),
		Args:              core.CastString(data["args"]),
		CurrentRetryCount: core.CastInt32(data["currentRetryCount"]),
		MaxTryCount:       core.CastInt32(data["maxTryCount"]),
		CreatedAt:         core.CastInt64(data["createdAt"]),
		UpdatedAt:         core.CastInt64(data["updatedAt"]),
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
	return map[string]interface{}{
		"jobId":             jobId,
		"name":              name,
		"userId":            userId,
		"scriptId":          scriptId,
		"args":              args,
		"currentRetryCount": currentRetryCount,
		"maxTryCount":       maxTryCount,
		"createdAt":         createdAt,
		"updatedAt":         updatedAt,
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
	JobId       *string `json:"jobId"`
	ScriptId    *string `json:"scriptId"`
	Args        *string `json:"args"`
	TryNumber   *int32  `json:"tryNumber"`
	StatusCode  *int32  `json:"statusCode"`
	Result      *string `json:"result"`
	TryAt       *int64  `json:"tryAt"`
}

func (p *JobResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = JobResult{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = JobResult{}
	} else {
		*p = JobResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["jobResultId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JobResultId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JobResultId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JobResultId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JobResultId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JobResultId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JobResultId)
				}
			}
		}
		if v, ok := d["jobId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.JobId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.JobId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.JobId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.JobId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.JobId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.JobId)
				}
			}
		}
		if v, ok := d["scriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScriptId)
				}
			}
		}
		if v, ok := d["args"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Args = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Args = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Args = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Args)
				}
			}
		}
		if v, ok := d["tryNumber"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TryNumber)
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["result"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Result = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Result = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Result = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Result = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Result = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Result)
				}
			}
		}
		if v, ok := d["tryAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TryAt)
		}
	}
	return nil
}

func NewJobResultFromJson(data string) JobResult {
	req := JobResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJobResultFromDict(data map[string]interface{}) JobResult {
	return JobResult{
		JobResultId: core.CastString(data["jobResultId"]),
		JobId:       core.CastString(data["jobId"]),
		ScriptId:    core.CastString(data["scriptId"]),
		Args:        core.CastString(data["args"]),
		TryNumber:   core.CastInt32(data["tryNumber"]),
		StatusCode:  core.CastInt32(data["statusCode"]),
		Result:      core.CastString(data["result"]),
		TryAt:       core.CastInt64(data["tryAt"]),
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
	var scriptId *string
	if p.ScriptId != nil {
		scriptId = p.ScriptId
	}
	var args *string
	if p.Args != nil {
		args = p.Args
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
	return map[string]interface{}{
		"jobResultId": jobResultId,
		"jobId":       jobId,
		"scriptId":    scriptId,
		"args":        args,
		"tryNumber":   tryNumber,
		"statusCode":  statusCode,
		"result":      result,
		"tryAt":       tryAt,
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
	DeadLetterJobId *string         `json:"deadLetterJobId"`
	Name            *string         `json:"name"`
	UserId          *string         `json:"userId"`
	ScriptId        *string         `json:"scriptId"`
	Args            *string         `json:"args"`
	Result          []JobResultBody `json:"result"`
	CreatedAt       *int64          `json:"createdAt"`
	UpdatedAt       *int64          `json:"updatedAt"`
}

func (p *DeadLetterJob) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeadLetterJob{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = DeadLetterJob{}
	} else {
		*p = DeadLetterJob{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["deadLetterJobId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DeadLetterJobId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DeadLetterJobId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DeadLetterJobId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DeadLetterJobId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DeadLetterJobId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DeadLetterJobId)
				}
			}
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Name = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Name = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Name = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Name = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Name)
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.UserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.UserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.UserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.UserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.UserId)
				}
			}
		}
		if v, ok := d["scriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScriptId)
				}
			}
		}
		if v, ok := d["args"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Args = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Args = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Args = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Args)
				}
			}
		}
		if v, ok := d["result"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Result)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["updatedAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UpdatedAt)
		}
	}
	return nil
}

func NewDeadLetterJobFromJson(data string) DeadLetterJob {
	req := DeadLetterJob{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDeadLetterJobFromDict(data map[string]interface{}) DeadLetterJob {
	return DeadLetterJob{
		DeadLetterJobId: core.CastString(data["deadLetterJobId"]),
		Name:            core.CastString(data["name"]),
		UserId:          core.CastString(data["userId"]),
		ScriptId:        core.CastString(data["scriptId"]),
		Args:            core.CastString(data["args"]),
		Result:          CastJobResultBodies(core.CastArray(data["result"])),
		CreatedAt:       core.CastInt64(data["createdAt"]),
		UpdatedAt:       core.CastInt64(data["updatedAt"]),
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
	return map[string]interface{}{
		"deadLetterJobId": deadLetterJobId,
		"name":            name,
		"userId":          userId,
		"scriptId":        scriptId,
		"args":            args,
		"result":          result,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
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
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func (p *NotificationSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NotificationSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = NotificationSetting{}
	} else {
		*p = NotificationSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["gatewayNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatewayNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatewayNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatewayNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatewayNamespaceId)
				}
			}
		}
		if v, ok := d["enableTransferMobileNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableTransferMobileNotification)
		}
		if v, ok := d["sound"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Sound = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Sound = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Sound = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Sound)
				}
			}
		}
	}
	return nil
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	req := NotificationSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
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
	return map[string]interface{}{
		"gatewayNamespaceId":               gatewayNamespaceId,
		"enableTransferMobileNotification": enableTransferMobileNotification,
		"sound":                            sound,
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

func (p *LogSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LogSetting{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = LogSetting{}
	} else {
		*p = LogSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["loggingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LoggingNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LoggingNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LoggingNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LoggingNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LoggingNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewLogSettingFromJson(data string) LogSetting {
	req := LogSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {

	var loggingNamespaceId *string
	if p.LoggingNamespaceId != nil {
		loggingNamespaceId = p.LoggingNamespaceId
	}
	return map[string]interface{}{
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
	ScriptId    *string `json:"scriptId"`
	Args        *string `json:"args"`
	MaxTryCount *int32  `json:"maxTryCount"`
}

func (p *JobEntry) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = JobEntry{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = JobEntry{}
	} else {
		*p = JobEntry{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["scriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScriptId)
				}
			}
		}
		if v, ok := d["args"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Args = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Args = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Args = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Args = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Args)
				}
			}
		}
		if v, ok := d["maxTryCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MaxTryCount)
		}
	}
	return nil
}

func NewJobEntryFromJson(data string) JobEntry {
	req := JobEntry{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJobEntryFromDict(data map[string]interface{}) JobEntry {
	return JobEntry{
		ScriptId:    core.CastString(data["scriptId"]),
		Args:        core.CastString(data["args"]),
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
	return map[string]interface{}{
		"scriptId":    scriptId,
		"args":        args,
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
	TryNumber  *int32  `json:"tryNumber"`
	StatusCode *int32  `json:"statusCode"`
	Result     *string `json:"result"`
	TryAt      *int64  `json:"tryAt"`
}

func (p *JobResultBody) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = JobResultBody{}
		return nil
	}
	if str[0] == '"' {
		var strVal string
		err := json.Unmarshal(data, &strVal)
		if err != nil {
			return err
		}
		str = strVal
	}
	if str == "null" {
		*p = JobResultBody{}
	} else {
		*p = JobResultBody{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["tryNumber"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TryNumber)
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["result"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Result = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Result = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Result = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Result = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Result = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Result)
				}
			}
		}
		if v, ok := d["tryAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TryAt)
		}
	}
	return nil
}

func NewJobResultBodyFromJson(data string) JobResultBody {
	req := JobResultBody{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewJobResultBodyFromDict(data map[string]interface{}) JobResultBody {
	return JobResultBody{
		TryNumber:  core.CastInt32(data["tryNumber"]),
		StatusCode: core.CastInt32(data["statusCode"]),
		Result:     core.CastString(data["result"]),
		TryAt:      core.CastInt64(data["tryAt"]),
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
	return map[string]interface{}{
		"tryNumber":  tryNumber,
		"statusCode": statusCode,
		"result":     result,
		"tryAt":      tryAt,
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
