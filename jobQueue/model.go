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
	NamespaceId *string `json:"namespaceId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	// Deprecated: should not be used
	TransactionSetting   *TransactionSetting   `json:"transactionSetting"`
	TransactionSettingV2 *TransactionSettingV2 `json:"transactionSettingV2"`
	EnableAutoRun        *bool                 `json:"enableAutoRun"`
	RunNotification      *NotificationSetting  `json:"runNotification"`
	PushNotification     *NotificationSetting  `json:"pushNotification"`
	LogSetting           *LogSetting           `json:"logSetting"`
	CreatedAt            *int64                `json:"createdAt"`
	UpdatedAt            *int64                `json:"updatedAt"`
	Revision             *int64                `json:"revision"`
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
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["transactionSettingV2"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSettingV2)
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
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		TransactionSettingV2: func() *TransactionSettingV2 {
			v, ok := data["transactionSettingV2"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingV2FromDict(core.CastMap(data["transactionSettingV2"])).Pointer()
		}(),
		EnableAutoRun: func() *bool {
			v, ok := data["enableAutoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAutoRun"])
		}(),
		RunNotification: func() *NotificationSetting {
			v, ok := data["runNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["runNotification"])).Pointer()
		}(),
		PushNotification: func() *NotificationSetting {
			v, ok := data["pushNotification"]
			if !ok || v == nil {
				return nil
			}
			return NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer()
		}(),
		LogSetting: func() *LogSetting {
			v, ok := data["logSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer()
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
		Revision: func() *int64 {
			v, ok := data["revision"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["revision"])
		}(),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.TransactionSetting != nil {
		m["transactionSetting"] = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	if p.TransactionSettingV2 != nil {
		m["transactionSettingV2"] = func() map[string]interface{} {
			if p.TransactionSettingV2 == nil {
				return nil
			}
			return p.TransactionSettingV2.ToDict()
		}()
	}
	if p.EnableAutoRun != nil {
		m["enableAutoRun"] = p.EnableAutoRun
	}
	if p.RunNotification != nil {
		m["runNotification"] = func() map[string]interface{} {
			if p.RunNotification == nil {
				return nil
			}
			return p.RunNotification.ToDict()
		}()
	}
	if p.PushNotification != nil {
		m["pushNotification"] = func() map[string]interface{} {
			if p.PushNotification == nil {
				return nil
			}
			return p.PushNotification.ToDict()
		}()
	}
	if p.LogSetting != nil {
		m["logSetting"] = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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
		JobId: func() *string {
			v, ok := data["jobId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["jobId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		ScriptId: func() *string {
			v, ok := data["scriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scriptId"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
		CurrentRetryCount: func() *int32 {
			v, ok := data["currentRetryCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["currentRetryCount"])
		}(),
		MaxTryCount: func() *int32 {
			v, ok := data["maxTryCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxTryCount"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
		UpdatedAt: func() *int64 {
			v, ok := data["updatedAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["updatedAt"])
		}(),
	}
}

func (p Job) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.JobId != nil {
		m["jobId"] = p.JobId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.ScriptId != nil {
		m["scriptId"] = p.ScriptId
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	if p.CurrentRetryCount != nil {
		m["currentRetryCount"] = p.CurrentRetryCount
	}
	if p.MaxTryCount != nil {
		m["maxTryCount"] = p.MaxTryCount
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
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
		JobResultId: func() *string {
			v, ok := data["jobResultId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["jobResultId"])
		}(),
		JobId: func() *string {
			v, ok := data["jobId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["jobId"])
		}(),
		ScriptId: func() *string {
			v, ok := data["scriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scriptId"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
		TryNumber: func() *int32 {
			v, ok := data["tryNumber"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["tryNumber"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		TryAt: func() *int64 {
			v, ok := data["tryAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tryAt"])
		}(),
	}
}

func (p JobResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.JobResultId != nil {
		m["jobResultId"] = p.JobResultId
	}
	if p.JobId != nil {
		m["jobId"] = p.JobId
	}
	if p.ScriptId != nil {
		m["scriptId"] = p.ScriptId
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	if p.TryNumber != nil {
		m["tryNumber"] = p.TryNumber
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.Result != nil {
		m["result"] = p.Result
	}
	if p.TryAt != nil {
		m["tryAt"] = p.TryAt
	}
	return m
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

type NotificationSetting struct {
	GatewayNamespaceId               *string                     `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool                       `json:"enableTransferMobileNotification"`
	Sound                            *string                     `json:"sound"`
	MobileNotificationMessages       []MobileNotificationMessage `json:"mobileNotificationMessages"`
	Enable                           *string                     `json:"enable"`
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
		if v, ok := d["mobileNotificationMessages"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.MobileNotificationMessages)
		}
		if v, ok := d["enable"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Enable = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Enable = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Enable = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Enable = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Enable = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Enable)
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
		GatewayNamespaceId: func() *string {
			v, ok := data["gatewayNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gatewayNamespaceId"])
		}(),
		EnableTransferMobileNotification: func() *bool {
			v, ok := data["enableTransferMobileNotification"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableTransferMobileNotification"])
		}(),
		Sound: func() *string {
			v, ok := data["sound"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sound"])
		}(),
		MobileNotificationMessages: func() []MobileNotificationMessage {
			if data["mobileNotificationMessages"] == nil {
				return nil
			}
			return CastMobileNotificationMessages(core.CastArray(data["mobileNotificationMessages"]))
		}(),
		Enable: func() *string {
			v, ok := data["enable"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enable"])
		}(),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.GatewayNamespaceId != nil {
		m["gatewayNamespaceId"] = p.GatewayNamespaceId
	}
	if p.EnableTransferMobileNotification != nil {
		m["enableTransferMobileNotification"] = p.EnableTransferMobileNotification
	}
	if p.Sound != nil {
		m["sound"] = p.Sound
	}
	if p.MobileNotificationMessages != nil {
		m["mobileNotificationMessages"] = CastMobileNotificationMessagesFromDict(
			p.MobileNotificationMessages,
		)
	}
	if p.Enable != nil {
		m["enable"] = p.Enable
	}
	return m
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

type MobileNotificationMessage struct {
	Locale  *string `json:"locale"`
	Title   *string `json:"title"`
	Message *string `json:"message"`
}

func (p *MobileNotificationMessage) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MobileNotificationMessage{}
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
		*p = MobileNotificationMessage{}
	} else {
		*p = MobileNotificationMessage{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["locale"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Locale = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Locale = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Locale = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Locale = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Locale = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Locale)
				}
			}
		}
		if v, ok := d["title"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Title = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Title = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Title = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Title = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Title = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Title)
				}
			}
		}
		if v, ok := d["message"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Message = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Message = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Message = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Message = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Message = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Message)
				}
			}
		}
	}
	return nil
}

func NewMobileNotificationMessageFromJson(data string) MobileNotificationMessage {
	req := MobileNotificationMessage{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMobileNotificationMessageFromDict(data map[string]interface{}) MobileNotificationMessage {
	return MobileNotificationMessage{
		Locale: func() *string {
			v, ok := data["locale"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["locale"])
		}(),
		Title: func() *string {
			v, ok := data["title"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["title"])
		}(),
		Message: func() *string {
			v, ok := data["message"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["message"])
		}(),
	}
}

func (p MobileNotificationMessage) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Locale != nil {
		m["locale"] = p.Locale
	}
	if p.Title != nil {
		m["title"] = p.Title
	}
	if p.Message != nil {
		m["message"] = p.Message
	}
	return m
}

func (p MobileNotificationMessage) Pointer() *MobileNotificationMessage {
	return &p
}

func CastMobileNotificationMessages(data []interface{}) []MobileNotificationMessage {
	v := make([]MobileNotificationMessage, 0)
	for _, d := range data {
		v = append(v, NewMobileNotificationMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMobileNotificationMessagesFromDict(data []MobileNotificationMessage) []interface{} {
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
		LoggingNamespaceId: func() *string {
			v, ok := data["loggingNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["loggingNamespaceId"])
		}(),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LoggingNamespaceId != nil {
		m["loggingNamespaceId"] = p.LoggingNamespaceId
	}
	return m
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

type TransactionSetting struct {
	EnableAutoRun                      *bool   `json:"enableAutoRun"`
	EnableAtomicCommit                 *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor          *bool   `json:"transactionUseDistributor"`
	CommitScriptResultInUseDistributor *bool   `json:"commitScriptResultInUseDistributor"`
	AcquireActionUseJobQueue           *bool   `json:"acquireActionUseJobQueue"`
	EnableSequentialExecution          *bool   `json:"enableSequentialExecution"`
	DistributorNamespaceId             *string `json:"distributorNamespaceId"`
	// Deprecated: should not be used
	KeyId            *string `json:"keyId"`
	QueueNamespaceId *string `json:"queueNamespaceId"`
}

func (p *TransactionSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionSetting{}
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
		*p = TransactionSetting{}
	} else {
		*p = TransactionSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["enableAutoRun"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAutoRun)
		}
		if v, ok := d["enableAtomicCommit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableAtomicCommit)
		}
		if v, ok := d["transactionUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionUseDistributor)
		}
		if v, ok := d["commitScriptResultInUseDistributor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CommitScriptResultInUseDistributor)
		}
		if v, ok := d["acquireActionUseJobQueue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionUseJobQueue)
		}
		if v, ok := d["enableSequentialExecution"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableSequentialExecution)
		}
		if v, ok := d["distributorNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DistributorNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DistributorNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DistributorNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DistributorNamespaceId)
				}
			}
		}
		if v, ok := d["keyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.KeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.KeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.KeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.KeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.KeyId)
				}
			}
		}
		if v, ok := d["queueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.QueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.QueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.QueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.QueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.QueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewTransactionSettingFromJson(data string) TransactionSetting {
	req := TransactionSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionSettingFromDict(data map[string]interface{}) TransactionSetting {
	return TransactionSetting{
		EnableAutoRun: func() *bool {
			v, ok := data["enableAutoRun"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAutoRun"])
		}(),
		EnableAtomicCommit: func() *bool {
			v, ok := data["enableAtomicCommit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableAtomicCommit"])
		}(),
		TransactionUseDistributor: func() *bool {
			v, ok := data["transactionUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["transactionUseDistributor"])
		}(),
		CommitScriptResultInUseDistributor: func() *bool {
			v, ok := data["commitScriptResultInUseDistributor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["commitScriptResultInUseDistributor"])
		}(),
		AcquireActionUseJobQueue: func() *bool {
			v, ok := data["acquireActionUseJobQueue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["acquireActionUseJobQueue"])
		}(),
		EnableSequentialExecution: func() *bool {
			v, ok := data["enableSequentialExecution"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableSequentialExecution"])
		}(),
		DistributorNamespaceId: func() *string {
			v, ok := data["distributorNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["distributorNamespaceId"])
		}(),
		KeyId: func() *string {
			v, ok := data["keyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["keyId"])
		}(),
		QueueNamespaceId: func() *string {
			v, ok := data["queueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["queueNamespaceId"])
		}(),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EnableAutoRun != nil {
		m["enableAutoRun"] = p.EnableAutoRun
	}
	if p.EnableAtomicCommit != nil {
		m["enableAtomicCommit"] = p.EnableAtomicCommit
	}
	if p.TransactionUseDistributor != nil {
		m["transactionUseDistributor"] = p.TransactionUseDistributor
	}
	if p.CommitScriptResultInUseDistributor != nil {
		m["commitScriptResultInUseDistributor"] = p.CommitScriptResultInUseDistributor
	}
	if p.AcquireActionUseJobQueue != nil {
		m["acquireActionUseJobQueue"] = p.AcquireActionUseJobQueue
	}
	if p.EnableSequentialExecution != nil {
		m["enableSequentialExecution"] = p.EnableSequentialExecution
	}
	if p.DistributorNamespaceId != nil {
		m["distributorNamespaceId"] = p.DistributorNamespaceId
	}
	if p.KeyId != nil {
		m["keyId"] = p.KeyId
	}
	if p.QueueNamespaceId != nil {
		m["queueNamespaceId"] = p.QueueNamespaceId
	}
	return m
}

func (p TransactionSetting) Pointer() *TransactionSetting {
	return &p
}

func CastTransactionSettings(data []interface{}) []TransactionSetting {
	v := make([]TransactionSetting, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingsFromDict(data []TransactionSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSettingV2 struct {
	DistributorNamespaceId  *string `json:"distributorNamespaceId"`
	EnableParallelExecution *bool   `json:"enableParallelExecution"`
}

func (p *TransactionSettingV2) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionSettingV2{}
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
		*p = TransactionSettingV2{}
	} else {
		*p = TransactionSettingV2{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["distributorNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DistributorNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DistributorNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DistributorNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DistributorNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DistributorNamespaceId)
				}
			}
		}
		if v, ok := d["enableParallelExecution"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableParallelExecution)
		}
	}
	return nil
}

func NewTransactionSettingV2FromJson(data string) TransactionSettingV2 {
	req := TransactionSettingV2{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionSettingV2FromDict(data map[string]interface{}) TransactionSettingV2 {
	return TransactionSettingV2{
		DistributorNamespaceId: func() *string {
			v, ok := data["distributorNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["distributorNamespaceId"])
		}(),
		EnableParallelExecution: func() *bool {
			v, ok := data["enableParallelExecution"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["enableParallelExecution"])
		}(),
	}
}

func (p TransactionSettingV2) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DistributorNamespaceId != nil {
		m["distributorNamespaceId"] = p.DistributorNamespaceId
	}
	if p.EnableParallelExecution != nil {
		m["enableParallelExecution"] = p.EnableParallelExecution
	}
	return m
}

func (p TransactionSettingV2) Pointer() *TransactionSettingV2 {
	return &p
}

func CastTransactionSettingV2s(data []interface{}) []TransactionSettingV2 {
	v := make([]TransactionSettingV2, 0)
	for _, d := range data {
		v = append(v, NewTransactionSettingV2FromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionSettingV2sFromDict(data []TransactionSettingV2) []interface{} {
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
		ScriptId: func() *string {
			v, ok := data["scriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scriptId"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
		MaxTryCount: func() *int32 {
			v, ok := data["maxTryCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["maxTryCount"])
		}(),
	}
}

func (p JobEntry) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.ScriptId != nil {
		m["scriptId"] = p.ScriptId
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	if p.MaxTryCount != nil {
		m["maxTryCount"] = p.MaxTryCount
	}
	return m
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
		TryNumber: func() *int32 {
			v, ok := data["tryNumber"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["tryNumber"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		TryAt: func() *int64 {
			v, ok := data["tryAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["tryAt"])
		}(),
	}
}

func (p JobResultBody) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TryNumber != nil {
		m["tryNumber"] = p.TryNumber
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.Result != nil {
		m["result"] = p.Result
	}
	if p.TryAt != nil {
		m["tryAt"] = p.TryAt
	}
	return m
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
