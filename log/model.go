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
	NamespaceId          *string `json:"namespaceId"`
	Name                 *string `json:"name"`
	Description          *string `json:"description"`
	Type                 *string `json:"type"`
	GcpCredentialJson    *string `json:"gcpCredentialJson"`
	BigQueryDatasetName  *string `json:"bigQueryDatasetName"`
	LogExpireDays        *int32  `json:"logExpireDays"`
	AwsRegion            *string `json:"awsRegion"`
	AwsAccessKeyId       *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey   *string `json:"awsSecretAccessKey"`
	FirehoseStreamName   *string `json:"firehoseStreamName"`
	FirehoseCompressData *string `json:"firehoseCompressData"`
	Status               *string `json:"status"`
	CreatedAt            *int64  `json:"createdAt"`
	UpdatedAt            *int64  `json:"updatedAt"`
	Revision             *int64  `json:"revision"`
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
		if v, ok := d["firehoseCompressData"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FirehoseCompressData = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FirehoseCompressData = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FirehoseCompressData = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FirehoseCompressData = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FirehoseCompressData = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FirehoseCompressData)
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
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		GcpCredentialJson: func() *string {
			v, ok := data["gcpCredentialJson"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["gcpCredentialJson"])
		}(),
		BigQueryDatasetName: func() *string {
			v, ok := data["bigQueryDatasetName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["bigQueryDatasetName"])
		}(),
		LogExpireDays: func() *int32 {
			v, ok := data["logExpireDays"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["logExpireDays"])
		}(),
		AwsRegion: func() *string {
			v, ok := data["awsRegion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["awsRegion"])
		}(),
		AwsAccessKeyId: func() *string {
			v, ok := data["awsAccessKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["awsAccessKeyId"])
		}(),
		AwsSecretAccessKey: func() *string {
			v, ok := data["awsSecretAccessKey"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["awsSecretAccessKey"])
		}(),
		FirehoseStreamName: func() *string {
			v, ok := data["firehoseStreamName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["firehoseStreamName"])
		}(),
		FirehoseCompressData: func() *string {
			v, ok := data["firehoseCompressData"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["firehoseCompressData"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
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
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.GcpCredentialJson != nil {
		m["gcpCredentialJson"] = p.GcpCredentialJson
	}
	if p.BigQueryDatasetName != nil {
		m["bigQueryDatasetName"] = p.BigQueryDatasetName
	}
	if p.LogExpireDays != nil {
		m["logExpireDays"] = p.LogExpireDays
	}
	if p.AwsRegion != nil {
		m["awsRegion"] = p.AwsRegion
	}
	if p.AwsAccessKeyId != nil {
		m["awsAccessKeyId"] = p.AwsAccessKeyId
	}
	if p.AwsSecretAccessKey != nil {
		m["awsSecretAccessKey"] = p.AwsSecretAccessKey
	}
	if p.FirehoseStreamName != nil {
		m["firehoseStreamName"] = p.FirehoseStreamName
	}
	if p.FirehoseCompressData != nil {
		m["firehoseCompressData"] = p.FirehoseCompressData
	}
	if p.Status != nil {
		m["status"] = p.Status
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
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		RequestId: func() *string {
			v, ok := data["requestId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["requestId"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Request: func() *string {
			v, ok := data["request"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["request"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
	}
}

func (p AccessLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.RequestId != nil {
		m["requestId"] = p.RequestId
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Request != nil {
		m["request"] = p.Request
	}
	if p.Result != nil {
		m["result"] = p.Result
	}
	return m
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
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p AccessLogCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
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
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.Tasks = l
			}
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
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
		Tasks: func() []*string {
			v, ok := data["tasks"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
	}
}

func (p IssueStampSheetLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	if p.Tasks != nil {
		m["tasks"] = core.CastStringsFromDict(
			p.Tasks,
		)
	}
	return m
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
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p IssueStampSheetLogCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
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
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
	}
}

func (p ExecuteStampSheetLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	return m
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
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p ExecuteStampSheetLogCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
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
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		TaskId: func() *string {
			v, ok := data["taskId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["taskId"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Args: func() *string {
			v, ok := data["args"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["args"])
		}(),
	}
}

func (p ExecuteStampTaskLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.TaskId != nil {
		m["taskId"] = p.TaskId
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Args != nil {
		m["args"] = p.Args
	}
	return m
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
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p ExecuteStampTaskLogCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
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

type InGameLog struct {
	Timestamp *int64         `json:"timestamp"`
	RequestId *string        `json:"requestId"`
	UserId    *string        `json:"userId"`
	Tags      []InGameLogTag `json:"tags"`
	Payload   *string        `json:"payload"`
}

func (p *InGameLog) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = InGameLog{}
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
		*p = InGameLog{}
	} else {
		*p = InGameLog{}
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
		if v, ok := d["tags"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Tags)
		}
		if v, ok := d["payload"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Payload = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Payload = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Payload = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Payload)
				}
			}
		}
	}
	return nil
}

func NewInGameLogFromJson(data string) InGameLog {
	req := InGameLog{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInGameLogFromDict(data map[string]interface{}) InGameLog {
	return InGameLog{
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		RequestId: func() *string {
			v, ok := data["requestId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["requestId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Tags: func() []InGameLogTag {
			if data["tags"] == nil {
				return nil
			}
			return CastInGameLogTags(core.CastArray(data["tags"]))
		}(),
		Payload: func() *string {
			v, ok := data["payload"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["payload"])
		}(),
	}
}

func (p InGameLog) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.RequestId != nil {
		m["requestId"] = p.RequestId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Tags != nil {
		m["tags"] = CastInGameLogTagsFromDict(
			p.Tags,
		)
	}
	if p.Payload != nil {
		m["payload"] = p.Payload
	}
	return m
}

func (p InGameLog) Pointer() *InGameLog {
	return &p
}

func CastInGameLogs(data []interface{}) []InGameLog {
	v := make([]InGameLog, 0)
	for _, d := range data {
		v = append(v, NewInGameLogFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInGameLogsFromDict(data []InGameLog) []interface{} {
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
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		SourceRequestId: func() *string {
			v, ok := data["sourceRequestId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourceRequestId"])
		}(),
		RequestId: func() *string {
			v, ok := data["requestId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["requestId"])
		}(),
		Duration: func() *int64 {
			v, ok := data["duration"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["duration"])
		}(),
		Service: func() *string {
			v, ok := data["service"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["service"])
		}(),
		Method: func() *string {
			v, ok := data["method"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["method"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Request: func() *string {
			v, ok := data["request"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["request"])
		}(),
		Result: func() *string {
			v, ok := data["result"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["result"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
	}
}

func (p AccessLogWithTelemetry) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.SourceRequestId != nil {
		m["sourceRequestId"] = p.SourceRequestId
	}
	if p.RequestId != nil {
		m["requestId"] = p.RequestId
	}
	if p.Duration != nil {
		m["duration"] = p.Duration
	}
	if p.Service != nil {
		m["service"] = p.Service
	}
	if p.Method != nil {
		m["method"] = p.Method
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Request != nil {
		m["request"] = p.Request
	}
	if p.Result != nil {
		m["result"] = p.Result
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	return m
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
		InsightId: func() *string {
			v, ok := data["insightId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["insightId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		TaskId: func() *string {
			v, ok := data["taskId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["taskId"])
		}(),
		Host: func() *string {
			v, ok := data["host"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["host"])
		}(),
		Password: func() *string {
			v, ok := data["password"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["password"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
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

func (p Insight) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.InsightId != nil {
		m["insightId"] = p.InsightId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.TaskId != nil {
		m["taskId"] = p.TaskId
	}
	if p.Host != nil {
		m["host"] = p.Host
	}
	if p.Password != nil {
		m["password"] = p.Password
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
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

type FacetModel struct {
	FacetModelId *string `json:"facetModelId"`
	Field        *string `json:"field"`
	Type         *string `json:"type"`
	DisplayName  *string `json:"displayName"`
	Order        *int32  `json:"order"`
}

func (p *FacetModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FacetModel{}
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
		*p = FacetModel{}
	} else {
		*p = FacetModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["facetModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FacetModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FacetModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FacetModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FacetModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FacetModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FacetModelId)
				}
			}
		}
		if v, ok := d["field"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Field = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Field = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Field = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Field)
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
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
				}
			}
		}
		if v, ok := d["order"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Order)
		}
	}
	return nil
}

func NewFacetModelFromJson(data string) FacetModel {
	req := FacetModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFacetModelFromDict(data map[string]interface{}) FacetModel {
	return FacetModel{
		FacetModelId: func() *string {
			v, ok := data["facetModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["facetModelId"])
		}(),
		Field: func() *string {
			v, ok := data["field"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["field"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Order: func() *int32 {
			v, ok := data["order"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["order"])
		}(),
	}
}

func (p FacetModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.FacetModelId != nil {
		m["facetModelId"] = p.FacetModelId
	}
	if p.Field != nil {
		m["field"] = p.Field
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.DisplayName != nil {
		m["displayName"] = p.DisplayName
	}
	if p.Order != nil {
		m["order"] = p.Order
	}
	return m
}

func (p FacetModel) Pointer() *FacetModel {
	return &p
}

func CastFacetModels(data []interface{}) []FacetModel {
	v := make([]FacetModel, 0)
	for _, d := range data {
		v = append(v, NewFacetModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFacetModelsFromDict(data []FacetModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Dashboard struct {
	DashboardId *string `json:"dashboardId"`
	Name        *string `json:"name"`
	DisplayName *string `json:"displayName"`
	Description *string `json:"description"`
	Payload     *string `json:"payload"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
}

func (p *Dashboard) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Dashboard{}
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
		*p = Dashboard{}
	} else {
		*p = Dashboard{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["dashboardId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DashboardId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DashboardId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DashboardId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DashboardId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DashboardId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DashboardId)
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
		if v, ok := d["displayName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DisplayName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DisplayName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DisplayName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DisplayName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DisplayName)
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
		if v, ok := d["payload"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Payload = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Payload = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Payload = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Payload = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Payload)
				}
			}
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

func NewDashboardFromJson(data string) Dashboard {
	req := Dashboard{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewDashboardFromDict(data map[string]interface{}) Dashboard {
	return Dashboard{
		DashboardId: func() *string {
			v, ok := data["dashboardId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["dashboardId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		DisplayName: func() *string {
			v, ok := data["displayName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["displayName"])
		}(),
		Description: func() *string {
			v, ok := data["description"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["description"])
		}(),
		Payload: func() *string {
			v, ok := data["payload"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["payload"])
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

func (p Dashboard) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.DashboardId != nil {
		m["dashboardId"] = p.DashboardId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.DisplayName != nil {
		m["displayName"] = p.DisplayName
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Payload != nil {
		m["payload"] = p.Payload
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Dashboard) Pointer() *Dashboard {
	return &p
}

func CastDashboards(data []interface{}) []Dashboard {
	v := make([]Dashboard, 0)
	for _, d := range data {
		v = append(v, NewDashboardFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDashboardsFromDict(data []Dashboard) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AggregationConfig struct {
	Type  *string `json:"type"`
	Field *string `json:"field"`
}

func (p *AggregationConfig) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AggregationConfig{}
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
		*p = AggregationConfig{}
	} else {
		*p = AggregationConfig{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["field"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Field = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Field = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Field = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Field)
				}
			}
		}
	}
	return nil
}

func NewAggregationConfigFromJson(data string) AggregationConfig {
	req := AggregationConfig{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAggregationConfigFromDict(data map[string]interface{}) AggregationConfig {
	return AggregationConfig{
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		Field: func() *string {
			v, ok := data["field"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["field"])
		}(),
	}
}

func (p AggregationConfig) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.Field != nil {
		m["field"] = p.Field
	}
	return m
}

func (p AggregationConfig) Pointer() *AggregationConfig {
	return &p
}

func CastAggregationConfigs(data []interface{}) []AggregationConfig {
	v := make([]AggregationConfig, 0)
	for _, d := range data {
		v = append(v, NewAggregationConfigFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAggregationConfigsFromDict(data []AggregationConfig) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Facet struct {
	Field       *string           `json:"field"`
	Values      []FacetValueCount `json:"values"`
	Range       *NumericRange     `json:"range"`
	GlobalRange *NumericRange     `json:"globalRange"`
}

func (p *Facet) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Facet{}
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
		*p = Facet{}
	} else {
		*p = Facet{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["field"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Field = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Field = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Field = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Field = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Field)
				}
			}
		}
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
		if v, ok := d["range"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Range)
		}
		if v, ok := d["globalRange"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.GlobalRange)
		}
	}
	return nil
}

func NewFacetFromJson(data string) Facet {
	req := Facet{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFacetFromDict(data map[string]interface{}) Facet {
	return Facet{
		Field: func() *string {
			v, ok := data["field"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["field"])
		}(),
		Values: func() []FacetValueCount {
			if data["values"] == nil {
				return nil
			}
			return CastFacetValueCounts(core.CastArray(data["values"]))
		}(),
		Range: func() *NumericRange {
			v, ok := data["range"]
			if !ok || v == nil {
				return nil
			}
			return NewNumericRangeFromDict(core.CastMap(data["range"])).Pointer()
		}(),
		GlobalRange: func() *NumericRange {
			v, ok := data["globalRange"]
			if !ok || v == nil {
				return nil
			}
			return NewNumericRangeFromDict(core.CastMap(data["globalRange"])).Pointer()
		}(),
	}
}

func (p Facet) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Field != nil {
		m["field"] = p.Field
	}
	if p.Values != nil {
		m["values"] = CastFacetValueCountsFromDict(
			p.Values,
		)
	}
	if p.Range != nil {
		m["range"] = func() map[string]interface{} {
			if p.Range == nil {
				return nil
			}
			return p.Range.ToDict()
		}()
	}
	if p.GlobalRange != nil {
		m["globalRange"] = func() map[string]interface{} {
			if p.GlobalRange == nil {
				return nil
			}
			return p.GlobalRange.ToDict()
		}()
	}
	return m
}

func (p Facet) Pointer() *Facet {
	return &p
}

func CastFacets(data []interface{}) []Facet {
	v := make([]Facet, 0)
	for _, d := range data {
		v = append(v, NewFacetFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFacetsFromDict(data []Facet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type FacetValueCount struct {
	Value *string `json:"value"`
	Count *int64  `json:"count"`
}

func (p *FacetValueCount) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = FacetValueCount{}
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
		*p = FacetValueCount{}
	} else {
		*p = FacetValueCount{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
		if v, ok := d["count"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Count)
		}
	}
	return nil
}

func NewFacetValueCountFromJson(data string) FacetValueCount {
	req := FacetValueCount{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewFacetValueCountFromDict(data map[string]interface{}) FacetValueCount {
	return FacetValueCount{
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
		Count: func() *int64 {
			v, ok := data["count"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["count"])
		}(),
	}
}

func (p FacetValueCount) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Value != nil {
		m["value"] = p.Value
	}
	if p.Count != nil {
		m["count"] = p.Count
	}
	return m
}

func (p FacetValueCount) Pointer() *FacetValueCount {
	return &p
}

func CastFacetValueCounts(data []interface{}) []FacetValueCount {
	v := make([]FacetValueCount, 0)
	for _, d := range data {
		v = append(v, NewFacetValueCountFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastFacetValueCountsFromDict(data []FacetValueCount) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Label struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *Label) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Label{}
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
		*p = Label{}
	} else {
		*p = Label{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
	}
	return nil
}

func NewLabelFromJson(data string) Label {
	req := Label{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLabelFromDict(data map[string]interface{}) Label {
	return Label{
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
	}
}

func (p Label) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p Label) Pointer() *Label {
	return &p
}

func CastLabels(data []interface{}) []Label {
	v := make([]Label, 0)
	for _, d := range data {
		v = append(v, NewLabelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLabelsFromDict(data []Label) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type LogEntry struct {
	Timestamp *int64  `json:"timestamp"`
	Status    *string `json:"status"`
	Duration  *int64  `json:"duration"`
	Line      *string `json:"line"`
	Labels    []Label `json:"labels"`
}

func (p *LogEntry) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LogEntry{}
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
		*p = LogEntry{}
	} else {
		*p = LogEntry{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
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
		if v, ok := d["duration"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Duration)
		}
		if v, ok := d["line"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Line = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Line = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Line = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Line = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Line = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Line)
				}
			}
		}
		if v, ok := d["labels"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Labels)
		}
	}
	return nil
}

func NewLogEntryFromJson(data string) LogEntry {
	req := LogEntry{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLogEntryFromDict(data map[string]interface{}) LogEntry {
	return LogEntry{
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
		Duration: func() *int64 {
			v, ok := data["duration"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["duration"])
		}(),
		Line: func() *string {
			v, ok := data["line"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["line"])
		}(),
		Labels: func() []Label {
			if data["labels"] == nil {
				return nil
			}
			return CastLabels(core.CastArray(data["labels"]))
		}(),
	}
}

func (p LogEntry) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	if p.Duration != nil {
		m["duration"] = p.Duration
	}
	if p.Line != nil {
		m["line"] = p.Line
	}
	if p.Labels != nil {
		m["labels"] = CastLabelsFromDict(
			p.Labels,
		)
	}
	return m
}

func (p LogEntry) Pointer() *LogEntry {
	return &p
}

func CastLogEntries(data []interface{}) []LogEntry {
	v := make([]LogEntry, 0)
	for _, d := range data {
		v = append(v, NewLogEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLogEntriesFromDict(data []LogEntry) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NumericRange struct {
	Min *float64 `json:"min"`
	Max *float64 `json:"max"`
}

func (p *NumericRange) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NumericRange{}
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
		*p = NumericRange{}
	} else {
		*p = NumericRange{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["min"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Min)
		}
		if v, ok := d["max"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Max)
		}
	}
	return nil
}

func NewNumericRangeFromJson(data string) NumericRange {
	req := NumericRange{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNumericRangeFromDict(data map[string]interface{}) NumericRange {
	return NumericRange{
		Min: func() *float64 {
			v, ok := data["min"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["min"])
		}(),
		Max: func() *float64 {
			v, ok := data["max"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["max"])
		}(),
	}
}

func (p NumericRange) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Min != nil {
		m["min"] = p.Min
	}
	if p.Max != nil {
		m["max"] = p.Max
	}
	return m
}

func (p NumericRange) Pointer() *NumericRange {
	return &p
}

func CastNumericRanges(data []interface{}) []NumericRange {
	v := make([]NumericRange, 0)
	for _, d := range data {
		v = append(v, NewNumericRangeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNumericRangesFromDict(data []NumericRange) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TimeseriesMetadata struct {
	Keys    []*string `json:"keys"`
	GroupBy []*string `json:"groupBy"`
}

func (p *TimeseriesMetadata) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TimeseriesMetadata{}
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
		*p = TimeseriesMetadata{}
	} else {
		*p = TimeseriesMetadata{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["keys"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.Keys = l
			}
		}
		if v, ok := d["groupBy"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.GroupBy = l
			}
		}
	}
	return nil
}

func NewTimeseriesMetadataFromJson(data string) TimeseriesMetadata {
	req := TimeseriesMetadata{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTimeseriesMetadataFromDict(data map[string]interface{}) TimeseriesMetadata {
	return TimeseriesMetadata{
		Keys: func() []*string {
			v, ok := data["keys"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		GroupBy: func() []*string {
			v, ok := data["groupBy"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
	}
}

func (p TimeseriesMetadata) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Keys != nil {
		m["keys"] = core.CastStringsFromDict(
			p.Keys,
		)
	}
	if p.GroupBy != nil {
		m["groupBy"] = core.CastStringsFromDict(
			p.GroupBy,
		)
	}
	return m
}

func (p TimeseriesMetadata) Pointer() *TimeseriesMetadata {
	return &p
}

func CastTimeseriesMetadatas(data []interface{}) []TimeseriesMetadata {
	v := make([]TimeseriesMetadata, 0)
	for _, d := range data {
		v = append(v, NewTimeseriesMetadataFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeseriesMetadatasFromDict(data []TimeseriesMetadata) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TimeseriesPoint struct {
	Timestamp *int64            `json:"timestamp"`
	Values    []TimeseriesValue `json:"values"`
}

func (p *TimeseriesPoint) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TimeseriesPoint{}
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
		*p = TimeseriesPoint{}
	} else {
		*p = TimeseriesPoint{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
		if v, ok := d["values"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Values)
		}
	}
	return nil
}

func NewTimeseriesPointFromJson(data string) TimeseriesPoint {
	req := TimeseriesPoint{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTimeseriesPointFromDict(data map[string]interface{}) TimeseriesPoint {
	return TimeseriesPoint{
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
		Values: func() []TimeseriesValue {
			if data["values"] == nil {
				return nil
			}
			return CastTimeseriesValues(core.CastArray(data["values"]))
		}(),
	}
}

func (p TimeseriesPoint) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	if p.Values != nil {
		m["values"] = CastTimeseriesValuesFromDict(
			p.Values,
		)
	}
	return m
}

func (p TimeseriesPoint) Pointer() *TimeseriesPoint {
	return &p
}

func CastTimeseriesPoints(data []interface{}) []TimeseriesPoint {
	v := make([]TimeseriesPoint, 0)
	for _, d := range data {
		v = append(v, NewTimeseriesPointFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeseriesPointsFromDict(data []TimeseriesPoint) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TimeseriesValue struct {
	Key   *string  `json:"key"`
	Value *float64 `json:"value"`
}

func (p *TimeseriesValue) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TimeseriesValue{}
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
		*p = TimeseriesValue{}
	} else {
		*p = TimeseriesValue{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Value)
		}
	}
	return nil
}

func NewTimeseriesValueFromJson(data string) TimeseriesValue {
	req := TimeseriesValue{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTimeseriesValueFromDict(data map[string]interface{}) TimeseriesValue {
	return TimeseriesValue{
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *float64 {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat64(data["value"])
		}(),
	}
}

func (p TimeseriesValue) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p TimeseriesValue) Pointer() *TimeseriesValue {
	return &p
}

func CastTimeseriesValues(data []interface{}) []TimeseriesValue {
	v := make([]TimeseriesValue, 0)
	for _, d := range data {
		v = append(v, NewTimeseriesValueFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTimeseriesValuesFromDict(data []TimeseriesValue) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Trace struct {
	TraceId   *string    `json:"traceId"`
	Spans     []LogEntry `json:"spans"`
	Truncated *bool      `json:"truncated"`
}

func (p *Trace) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Trace{}
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
		*p = Trace{}
	} else {
		*p = Trace{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["traceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TraceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TraceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TraceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TraceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TraceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TraceId)
				}
			}
		}
		if v, ok := d["spans"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Spans)
		}
		if v, ok := d["truncated"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Truncated)
		}
	}
	return nil
}

func NewTraceFromJson(data string) Trace {
	req := Trace{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTraceFromDict(data map[string]interface{}) Trace {
	return Trace{
		TraceId: func() *string {
			v, ok := data["traceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["traceId"])
		}(),
		Spans: func() []LogEntry {
			if data["spans"] == nil {
				return nil
			}
			return CastLogEntries(core.CastArray(data["spans"]))
		}(),
		Truncated: func() *bool {
			v, ok := data["truncated"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["truncated"])
		}(),
	}
}

func (p Trace) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TraceId != nil {
		m["traceId"] = p.TraceId
	}
	if p.Spans != nil {
		m["spans"] = CastLogEntriesFromDict(
			p.Spans,
		)
	}
	if p.Truncated != nil {
		m["truncated"] = p.Truncated
	}
	return m
}

func (p Trace) Pointer() *Trace {
	return &p
}

func CastTraces(data []interface{}) []Trace {
	v := make([]Trace, 0)
	for _, d := range data {
		v = append(v, NewTraceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTracesFromDict(data []Trace) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type MetricModel struct {
	Name   *string   `json:"name"`
	Type   *string   `json:"type"`
	Labels []*string `json:"labels"`
}

func (p *MetricModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MetricModel{}
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
		*p = MetricModel{}
	} else {
		*p = MetricModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["labels"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.Labels = l
			}
		}
	}
	return nil
}

func NewMetricModelFromJson(data string) MetricModel {
	req := MetricModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMetricModelFromDict(data map[string]interface{}) MetricModel {
	return MetricModel{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		Labels: func() []*string {
			v, ok := data["labels"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
	}
}

func (p MetricModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.Labels != nil {
		m["labels"] = core.CastStringsFromDict(
			p.Labels,
		)
	}
	return m
}

func (p MetricModel) Pointer() *MetricModel {
	return &p
}

func CastMetricModels(data []interface{}) []MetricModel {
	v := make([]MetricModel, 0)
	for _, d := range data {
		v = append(v, NewMetricModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMetricModelsFromDict(data []MetricModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type InGameLogTag struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

func (p *InGameLogTag) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = InGameLogTag{}
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
		*p = InGameLogTag{}
	} else {
		*p = InGameLogTag{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["key"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Key = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Key = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Key = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Key = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Key)
				}
			}
		}
		if v, ok := d["value"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Value = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Value = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Value = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Value = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Value)
				}
			}
		}
	}
	return nil
}

func NewInGameLogTagFromJson(data string) InGameLogTag {
	req := InGameLogTag{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewInGameLogTagFromDict(data map[string]interface{}) InGameLogTag {
	return InGameLogTag{
		Key: func() *string {
			v, ok := data["key"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["key"])
		}(),
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
	}
}

func (p InGameLogTag) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Key != nil {
		m["key"] = p.Key
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p InGameLogTag) Pointer() *InGameLogTag {
	return &p
}

func CastInGameLogTags(data []interface{}) []InGameLogTag {
	v := make([]InGameLogTag, 0)
	for _, d := range data {
		v = append(v, NewInGameLogTagFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastInGameLogTagsFromDict(data []InGameLogTag) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
