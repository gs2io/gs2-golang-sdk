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
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

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
	Status              *string `json:"status"`
	CreatedAt           *int64  `json:"createdAt"`
	UpdatedAt           *int64  `json:"updatedAt"`
	Revision            *int64  `json:"revision"`
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
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["gcpCredentialJson"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GcpCredentialJson = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GcpCredentialJson = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GcpCredentialJson = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GcpCredentialJson = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GcpCredentialJson = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GcpCredentialJson)
				}
			}
		}
		if v, ok := d["bigQueryDatasetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BigQueryDatasetName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BigQueryDatasetName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BigQueryDatasetName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BigQueryDatasetName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BigQueryDatasetName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BigQueryDatasetName)
				}
			}
		}
		if v, ok := d["logExpireDays"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LogExpireDays)
		}
		if v, ok := d["awsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AwsRegion = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AwsRegion = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AwsRegion = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AwsRegion = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AwsRegion = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AwsRegion)
				}
			}
		}
		if v, ok := d["awsAccessKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AwsAccessKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AwsAccessKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AwsAccessKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AwsAccessKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AwsAccessKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AwsAccessKeyId)
				}
			}
		}
		if v, ok := d["awsSecretAccessKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AwsSecretAccessKey = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AwsSecretAccessKey = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AwsSecretAccessKey = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AwsSecretAccessKey = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AwsSecretAccessKey = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AwsSecretAccessKey)
				}
			}
		}
		if v, ok := d["firehoseStreamName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FirehoseStreamName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FirehoseStreamName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FirehoseStreamName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FirehoseStreamName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FirehoseStreamName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FirehoseStreamName)
				}
			}
		}
		if v, ok := d["status"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Status = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Status = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Status = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Status)
				}
			}
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
		Status:              core.CastString(data["status"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
		Revision:            core.CastInt64(data["revision"]),
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
	return map[string]interface{}{
		"namespaceId":         namespaceId,
		"name":                name,
		"description":         description,
		"type":                _type,
		"gcpCredentialJson":   gcpCredentialJson,
		"bigQueryDatasetName": bigQueryDatasetName,
		"logExpireDays":       logExpireDays,
		"awsRegion":           awsRegion,
		"awsAccessKeyId":      awsAccessKeyId,
		"awsSecretAccessKey":  awsSecretAccessKey,
		"firehoseStreamName":  firehoseStreamName,
		"status":              status,
		"createdAt":           createdAt,
		"updatedAt":           updatedAt,
		"revision":            revision,
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

func (p *AccessLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AccessLog{}
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
		*p = AccessLog{}
	} else {
		*p = AccessLog{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["requestId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RequestId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RequestId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RequestId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RequestId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RequestId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RequestId)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
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
	}
	return nil
}

func NewAccessLogFromJson(data string) AccessLog {
	req := AccessLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"timestamp": timestamp,
		"requestId": requestId,
		"service":   service,
		"method":    method,
		"userId":    userId,
		"request":   request,
		"result":    result,
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

func (p *AccessLogCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AccessLogCount{}
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
		*p = AccessLogCount{}
	} else {
		*p = AccessLogCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewAccessLogCountFromJson(data string) AccessLogCount {
	req := AccessLogCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"service": service,
		"method":  method,
		"userId":  userId,
		"count":   count,
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
	Timestamp     *int64    `json:"timestamp"`
	TransactionId *string   `json:"transactionId"`
	Service       *string   `json:"service"`
	Method        *string   `json:"method"`
	UserId        *string   `json:"userId"`
	Action        *string   `json:"action"`
	Args          *string   `json:"args"`
	Tasks         []*string `json:"tasks"`
}

func (p *IssueStampSheetLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IssueStampSheetLog{}
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
		*p = IssueStampSheetLog{}
	} else {
		*p = IssueStampSheetLog{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
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
		if v, ok := d["tasks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tasks)
		}
	}
	return nil
}

func NewIssueStampSheetLogFromJson(data string) IssueStampSheetLog {
	req := IssueStampSheetLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
		Tasks:         core.CastStrings(core.CastArray(data["tasks"])),
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
	return map[string]interface{}{
		"timestamp":     timestamp,
		"transactionId": transactionId,
		"service":       service,
		"method":        method,
		"userId":        userId,
		"action":        action,
		"args":          args,
		"tasks":         tasks,
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

func (p *IssueStampSheetLogCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = IssueStampSheetLogCount{}
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
		*p = IssueStampSheetLogCount{}
	} else {
		*p = IssueStampSheetLogCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewIssueStampSheetLogCountFromJson(data string) IssueStampSheetLogCount {
	req := IssueStampSheetLogCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"service": service,
		"method":  method,
		"userId":  userId,
		"action":  action,
		"count":   count,
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

func (p *ExecuteStampSheetLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExecuteStampSheetLog{}
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
		*p = ExecuteStampSheetLog{}
	} else {
		*p = ExecuteStampSheetLog{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["transactionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TransactionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TransactionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TransactionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TransactionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TransactionId)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
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
	}
	return nil
}

func NewExecuteStampSheetLogFromJson(data string) ExecuteStampSheetLog {
	req := ExecuteStampSheetLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"timestamp":     timestamp,
		"transactionId": transactionId,
		"service":       service,
		"method":        method,
		"userId":        userId,
		"action":        action,
		"args":          args,
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

func (p *ExecuteStampSheetLogCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExecuteStampSheetLogCount{}
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
		*p = ExecuteStampSheetLogCount{}
	} else {
		*p = ExecuteStampSheetLogCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewExecuteStampSheetLogCountFromJson(data string) ExecuteStampSheetLogCount {
	req := ExecuteStampSheetLogCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"service": service,
		"method":  method,
		"userId":  userId,
		"action":  action,
		"count":   count,
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

func (p *ExecuteStampTaskLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExecuteStampTaskLog{}
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
		*p = ExecuteStampTaskLog{}
	} else {
		*p = ExecuteStampTaskLog{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["taskId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TaskId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TaskId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TaskId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TaskId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TaskId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TaskId)
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
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
	}
	return nil
}

func NewExecuteStampTaskLogFromJson(data string) ExecuteStampTaskLog {
	req := ExecuteStampTaskLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"timestamp": timestamp,
		"taskId":    taskId,
		"service":   service,
		"method":    method,
		"userId":    userId,
		"action":    action,
		"args":      args,
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

func (p *ExecuteStampTaskLogCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ExecuteStampTaskLogCount{}
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
		*p = ExecuteStampTaskLogCount{}
	} else {
		*p = ExecuteStampTaskLogCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Action = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Action = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Action = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Action = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Action)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewExecuteStampTaskLogCountFromJson(data string) ExecuteStampTaskLogCount {
	req := ExecuteStampTaskLogCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
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
	return map[string]interface{}{
		"service": service,
		"method":  method,
		"userId":  userId,
		"action":  action,
		"count":   count,
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

type AccessLogWithTelemetry struct {
	Timestamp       *int64  `json:"timestamp"`
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	Duration        *int64  `json:"duration"`
	Service         *string `json:"service"`
	Method          *string `json:"method"`
	UserId          *string `json:"userId"`
	Request         *string `json:"request"`
	Result          *string `json:"result"`
	Status          *string `json:"status"`
}

func (p *AccessLogWithTelemetry) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AccessLogWithTelemetry{}
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
		*p = AccessLogWithTelemetry{}
	} else {
		*p = AccessLogWithTelemetry{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["sourceRequestId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourceRequestId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourceRequestId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourceRequestId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourceRequestId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourceRequestId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourceRequestId)
				}
			}
		}
		if v, ok := d["requestId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RequestId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RequestId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RequestId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RequestId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RequestId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RequestId)
				}
			}
		}
		if v, ok := d["duration"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Duration)
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Service = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Service = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Service = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Service = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Service)
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Method = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Method = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Method = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Method = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Method)
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
		if v, ok := d["request"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Request = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Request = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Request = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Request = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Request)
				}
			}
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
		if v, ok := d["status"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Status = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Status = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Status = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Status)
				}
			}
		}
	}
	return nil
}

func NewAccessLogWithTelemetryFromJson(data string) AccessLogWithTelemetry {
	req := AccessLogWithTelemetry{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAccessLogWithTelemetryFromDict(data map[string]interface{}) AccessLogWithTelemetry {
	return AccessLogWithTelemetry{
		Timestamp:       core.CastInt64(data["timestamp"]),
		SourceRequestId: core.CastString(data["sourceRequestId"]),
		RequestId:       core.CastString(data["requestId"]),
		Duration:        core.CastInt64(data["duration"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Request:         core.CastString(data["request"]),
		Result:          core.CastString(data["result"]),
		Status:          core.CastString(data["status"]),
	}
}

func (p AccessLogWithTelemetry) ToDict() map[string]interface{} {

	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	var sourceRequestId *string
	if p.SourceRequestId != nil {
		sourceRequestId = p.SourceRequestId
	}
	var requestId *string
	if p.RequestId != nil {
		requestId = p.RequestId
	}
	var duration *int64
	if p.Duration != nil {
		duration = p.Duration
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
	var status *string
	if p.Status != nil {
		status = p.Status
	}
	return map[string]interface{}{
		"timestamp":       timestamp,
		"sourceRequestId": sourceRequestId,
		"requestId":       requestId,
		"duration":        duration,
		"service":         service,
		"method":          method,
		"userId":          userId,
		"request":         request,
		"result":          result,
		"status":          status,
	}
}

func (p AccessLogWithTelemetry) Pointer() *AccessLogWithTelemetry {
	return &p
}

func CastAccessLogWithTelemetries(data []interface{}) []AccessLogWithTelemetry {
	v := make([]AccessLogWithTelemetry, 0)
	for _, d := range data {
		v = append(v, NewAccessLogWithTelemetryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAccessLogWithTelemetriesFromDict(data []AccessLogWithTelemetry) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Insight struct {
	InsightId *string `json:"insightId"`
	Name      *string `json:"name"`
	TaskId    *string `json:"taskId"`
	Host      *string `json:"host"`
	Password  *string `json:"password"`
	Status    *string `json:"status"`
	CreatedAt *int64  `json:"createdAt"`
	Revision  *int64  `json:"revision"`
}

func (p *Insight) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Insight{}
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
		*p = Insight{}
	} else {
		*p = Insight{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["insightId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.InsightId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.InsightId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.InsightId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.InsightId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.InsightId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.InsightId)
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
		if v, ok := d["taskId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TaskId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TaskId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TaskId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TaskId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TaskId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TaskId)
				}
			}
		}
		if v, ok := d["host"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Host = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Host = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Host = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Host = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Host = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Host)
				}
			}
		}
		if v, ok := d["password"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Password = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Password = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Password = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Password = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Password)
				}
			}
		}
		if v, ok := d["status"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Status = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Status = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Status = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Status = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Status)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewInsightFromJson(data string) Insight {
	req := Insight{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInsightFromDict(data map[string]interface{}) Insight {
	return Insight{
		InsightId: core.CastString(data["insightId"]),
		Name:      core.CastString(data["name"]),
		TaskId:    core.CastString(data["taskId"]),
		Host:      core.CastString(data["host"]),
		Password:  core.CastString(data["password"]),
		Status:    core.CastString(data["status"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		Revision:  core.CastInt64(data["revision"]),
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
	return map[string]interface{}{
		"insightId": insightId,
		"name":      name,
		"taskId":    taskId,
		"host":      host,
		"password":  password,
		"status":    status,
		"createdAt": createdAt,
		"revision":  revision,
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
