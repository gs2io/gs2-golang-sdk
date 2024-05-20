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

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeNamespacesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeNamespacesRequest{}
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
		*p = DescribeNamespacesRequest{}
	} else {
		*p = DescribeNamespacesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewDescribeNamespacesRequestFromJson(data string) (DescribeNamespacesRequest, error) {
	req := DescribeNamespacesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeNamespacesRequest{}, err
	}
	return req, nil
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
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
}

func (p *CreateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateNamespaceRequest{}
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
		*p = CreateNamespaceRequest{}
	} else {
		*p = CreateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["name"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Name); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Type); err != nil {
					return err
				}
			}
		}
		if v, ok := d["gcpCredentialJson"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.GcpCredentialJson); err != nil {
					return err
				}
			}
		}
		if v, ok := d["bigQueryDatasetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.BigQueryDatasetName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["logExpireDays"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LogExpireDays); err != nil {
				return err
			}
		}
		if v, ok := d["awsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsRegion); err != nil {
					return err
				}
			}
		}
		if v, ok := d["awsAccessKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsAccessKeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["awsSecretAccessKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsSecretAccessKey); err != nil {
					return err
				}
			}
		}
		if v, ok := d["firehoseStreamName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.FirehoseStreamName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCreateNamespaceRequestFromJson(data string) (CreateNamespaceRequest, error) {
	req := CreateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateNamespaceRequest{}, err
	}
	return req, nil
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
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
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceStatusRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceStatusRequest{}
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
		*p = GetNamespaceStatusRequest{}
	} else {
		*p = GetNamespaceStatusRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceStatusRequestFromJson(data string) (GetNamespaceStatusRequest, error) {
	req := GetNamespaceStatusRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceStatusRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *GetNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetNamespaceRequest{}
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
		*p = GetNamespaceRequest{}
	} else {
		*p = GetNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetNamespaceRequestFromJson(data string) (GetNamespaceRequest, error) {
	req := GetNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetNamespaceRequest{}, err
	}
	return req, nil
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	SourceRequestId     *string `json:"sourceRequestId"`
	RequestId           *string `json:"requestId"`
	ContextStack        *string `json:"contextStack"`
	NamespaceName       *string `json:"namespaceName"`
	Description         *string `json:"description"`
	Type                *string `json:"type"`
	GcpCredentialJson   *string `json:"gcpCredentialJson"`
	BigQueryDatasetName *string `json:"bigQueryDatasetName"`
	LogExpireDays       *int32  `json:"logExpireDays"`
	AwsRegion           *string `json:"awsRegion"`
	AwsAccessKeyId      *string `json:"awsAccessKeyId"`
	AwsSecretAccessKey  *string `json:"awsSecretAccessKey"`
	FirehoseStreamName  *string `json:"firehoseStreamName"`
}

func (p *UpdateNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UpdateNamespaceRequest{}
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
		*p = UpdateNamespaceRequest{}
	} else {
		*p = UpdateNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["description"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Description); err != nil {
					return err
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Type); err != nil {
					return err
				}
			}
		}
		if v, ok := d["gcpCredentialJson"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.GcpCredentialJson); err != nil {
					return err
				}
			}
		}
		if v, ok := d["bigQueryDatasetName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.BigQueryDatasetName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["logExpireDays"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LogExpireDays); err != nil {
				return err
			}
		}
		if v, ok := d["awsRegion"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsRegion); err != nil {
					return err
				}
			}
		}
		if v, ok := d["awsAccessKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsAccessKeyId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["awsSecretAccessKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.AwsSecretAccessKey); err != nil {
					return err
				}
			}
		}
		if v, ok := d["firehoseStreamName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.FirehoseStreamName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewUpdateNamespaceRequestFromJson(data string) (UpdateNamespaceRequest, error) {
	req := UpdateNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return UpdateNamespaceRequest{}, err
	}
	return req, nil
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		Description:         core.CastString(data["description"]),
		Type:                core.CastString(data["type"]),
		GcpCredentialJson:   core.CastString(data["gcpCredentialJson"]),
		BigQueryDatasetName: core.CastString(data["bigQueryDatasetName"]),
		LogExpireDays:       core.CastInt32(data["logExpireDays"]),
		AwsRegion:           core.CastString(data["awsRegion"]),
		AwsAccessKeyId:      core.CastString(data["awsAccessKeyId"]),
		AwsSecretAccessKey:  core.CastString(data["awsSecretAccessKey"]),
		FirehoseStreamName:  core.CastString(data["firehoseStreamName"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":       p.NamespaceName,
		"description":         p.Description,
		"type":                p.Type,
		"gcpCredentialJson":   p.GcpCredentialJson,
		"bigQueryDatasetName": p.BigQueryDatasetName,
		"logExpireDays":       p.LogExpireDays,
		"awsRegion":           p.AwsRegion,
		"awsAccessKeyId":      p.AwsAccessKeyId,
		"awsSecretAccessKey":  p.AwsSecretAccessKey,
		"firehoseStreamName":  p.FirehoseStreamName,
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *DeleteNamespaceRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteNamespaceRequest{}
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
		*p = DeleteNamespaceRequest{}
	} else {
		*p = DeleteNamespaceRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteNamespaceRequestFromJson(data string) (DeleteNamespaceRequest, error) {
	req := DeleteNamespaceRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteNamespaceRequest{}, err
	}
	return req, nil
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type QueryAccessLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *QueryAccessLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QueryAccessLogRequest{}
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
		*p = QueryAccessLogRequest{}
	} else {
		*p = QueryAccessLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Service); err != nil {
					return err
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Method); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewQueryAccessLogRequestFromJson(data string) (QueryAccessLogRequest, error) {
	req := QueryAccessLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return QueryAccessLogRequest{}, err
	}
	return req, nil
}

func NewQueryAccessLogRequestFromDict(data map[string]interface{}) QueryAccessLogRequest {
	return QueryAccessLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryAccessLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryAccessLogRequest) Pointer() *QueryAccessLogRequest {
	return &p
}

type CountAccessLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CountAccessLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CountAccessLogRequest{}
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
		*p = CountAccessLogRequest{}
	} else {
		*p = CountAccessLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Service); err != nil {
				return err
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Method); err != nil {
				return err
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.UserId); err != nil {
				return err
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCountAccessLogRequestFromJson(data string) (CountAccessLogRequest, error) {
	req := CountAccessLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CountAccessLogRequest{}, err
	}
	return req, nil
}

func NewCountAccessLogRequestFromDict(data map[string]interface{}) CountAccessLogRequest {
	return CountAccessLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountAccessLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountAccessLogRequest) Pointer() *CountAccessLogRequest {
	return &p
}

type QueryIssueStampSheetLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *QueryIssueStampSheetLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QueryIssueStampSheetLogRequest{}
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
		*p = QueryIssueStampSheetLogRequest{}
	} else {
		*p = QueryIssueStampSheetLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Service); err != nil {
					return err
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Method); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Action); err != nil {
					return err
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewQueryIssueStampSheetLogRequestFromJson(data string) (QueryIssueStampSheetLogRequest, error) {
	req := QueryIssueStampSheetLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return QueryIssueStampSheetLogRequest{}, err
	}
	return req, nil
}

func NewQueryIssueStampSheetLogRequestFromDict(data map[string]interface{}) QueryIssueStampSheetLogRequest {
	return QueryIssueStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryIssueStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryIssueStampSheetLogRequest) Pointer() *QueryIssueStampSheetLogRequest {
	return &p
}

type CountIssueStampSheetLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CountIssueStampSheetLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CountIssueStampSheetLogRequest{}
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
		*p = CountIssueStampSheetLogRequest{}
	} else {
		*p = CountIssueStampSheetLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Service); err != nil {
				return err
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Method); err != nil {
				return err
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.UserId); err != nil {
				return err
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Action); err != nil {
				return err
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCountIssueStampSheetLogRequestFromJson(data string) (CountIssueStampSheetLogRequest, error) {
	req := CountIssueStampSheetLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CountIssueStampSheetLogRequest{}, err
	}
	return req, nil
}

func NewCountIssueStampSheetLogRequestFromDict(data map[string]interface{}) CountIssueStampSheetLogRequest {
	return CountIssueStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountIssueStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountIssueStampSheetLogRequest) Pointer() *CountIssueStampSheetLogRequest {
	return &p
}

type QueryExecuteStampSheetLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *QueryExecuteStampSheetLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QueryExecuteStampSheetLogRequest{}
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
		*p = QueryExecuteStampSheetLogRequest{}
	} else {
		*p = QueryExecuteStampSheetLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Service); err != nil {
					return err
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Method); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Action); err != nil {
					return err
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewQueryExecuteStampSheetLogRequestFromJson(data string) (QueryExecuteStampSheetLogRequest, error) {
	req := QueryExecuteStampSheetLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return QueryExecuteStampSheetLogRequest{}, err
	}
	return req, nil
}

func NewQueryExecuteStampSheetLogRequestFromDict(data map[string]interface{}) QueryExecuteStampSheetLogRequest {
	return QueryExecuteStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryExecuteStampSheetLogRequest) Pointer() *QueryExecuteStampSheetLogRequest {
	return &p
}

type CountExecuteStampSheetLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CountExecuteStampSheetLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CountExecuteStampSheetLogRequest{}
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
		*p = CountExecuteStampSheetLogRequest{}
	} else {
		*p = CountExecuteStampSheetLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Service); err != nil {
				return err
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Method); err != nil {
				return err
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.UserId); err != nil {
				return err
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Action); err != nil {
				return err
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCountExecuteStampSheetLogRequestFromJson(data string) (CountExecuteStampSheetLogRequest, error) {
	req := CountExecuteStampSheetLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CountExecuteStampSheetLogRequest{}, err
	}
	return req, nil
}

func NewCountExecuteStampSheetLogRequestFromDict(data map[string]interface{}) CountExecuteStampSheetLogRequest {
	return CountExecuteStampSheetLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountExecuteStampSheetLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountExecuteStampSheetLogRequest) Pointer() *CountExecuteStampSheetLogRequest {
	return &p
}

type QueryExecuteStampTaskLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Service            *string `json:"service"`
	Method             *string `json:"method"`
	UserId             *string `json:"userId"`
	Action             *string `json:"action"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *QueryExecuteStampTaskLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QueryExecuteStampTaskLogRequest{}
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
		*p = QueryExecuteStampTaskLogRequest{}
	} else {
		*p = QueryExecuteStampTaskLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Service); err != nil {
					return err
				}
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Method); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Action); err != nil {
					return err
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewQueryExecuteStampTaskLogRequestFromJson(data string) (QueryExecuteStampTaskLogRequest, error) {
	req := QueryExecuteStampTaskLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return QueryExecuteStampTaskLogRequest{}, err
	}
	return req, nil
}

func NewQueryExecuteStampTaskLogRequestFromDict(data map[string]interface{}) QueryExecuteStampTaskLogRequest {
	return QueryExecuteStampTaskLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastString(data["service"]),
		Method:          core.CastString(data["method"]),
		UserId:          core.CastString(data["userId"]),
		Action:          core.CastString(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryExecuteStampTaskLogRequest) Pointer() *QueryExecuteStampTaskLogRequest {
	return &p
}

type CountExecuteStampTaskLogRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Service         *bool   `json:"service"`
	Method          *bool   `json:"method"`
	UserId          *bool   `json:"userId"`
	Action          *bool   `json:"action"`
	Begin           *int64  `json:"begin"`
	End             *int64  `json:"end"`
	LongTerm        *bool   `json:"longTerm"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func (p *CountExecuteStampTaskLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CountExecuteStampTaskLogRequest{}
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
		*p = CountExecuteStampTaskLogRequest{}
	} else {
		*p = CountExecuteStampTaskLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["service"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Service); err != nil {
				return err
			}
		}
		if v, ok := d["method"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Method); err != nil {
				return err
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.UserId); err != nil {
				return err
			}
		}
		if v, ok := d["action"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Action); err != nil {
				return err
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCountExecuteStampTaskLogRequestFromJson(data string) (CountExecuteStampTaskLogRequest, error) {
	req := CountExecuteStampTaskLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CountExecuteStampTaskLogRequest{}, err
	}
	return req, nil
}

func NewCountExecuteStampTaskLogRequestFromDict(data map[string]interface{}) CountExecuteStampTaskLogRequest {
	return CountExecuteStampTaskLogRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		Service:         core.CastBool(data["service"]),
		Method:          core.CastBool(data["method"]),
		UserId:          core.CastBool(data["userId"]),
		Action:          core.CastBool(data["action"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CountExecuteStampTaskLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"service":         p.Service,
		"method":          p.Method,
		"userId":          p.UserId,
		"action":          p.Action,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CountExecuteStampTaskLogRequest) Pointer() *CountExecuteStampTaskLogRequest {
	return &p
}

type QueryAccessLogWithTelemetryRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	Begin              *int64  `json:"begin"`
	End                *int64  `json:"end"`
	LongTerm           *bool   `json:"longTerm"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func (p *QueryAccessLogWithTelemetryRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = QueryAccessLogWithTelemetryRequest{}
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
		*p = QueryAccessLogWithTelemetryRequest{}
	} else {
		*p = QueryAccessLogWithTelemetryRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["userId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.UserId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["begin"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Begin); err != nil {
				return err
			}
		}
		if v, ok := d["end"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.End); err != nil {
				return err
			}
		}
		if v, ok := d["longTerm"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.LongTerm); err != nil {
				return err
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
		if v, ok := d["timeOffsetToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.TimeOffsetToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.TimeOffsetToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.TimeOffsetToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.TimeOffsetToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.TimeOffsetToken); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewQueryAccessLogWithTelemetryRequestFromJson(data string) (QueryAccessLogWithTelemetryRequest, error) {
	req := QueryAccessLogWithTelemetryRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return QueryAccessLogWithTelemetryRequest{}, err
	}
	return req, nil
}

func NewQueryAccessLogWithTelemetryRequestFromDict(data map[string]interface{}) QueryAccessLogWithTelemetryRequest {
	return QueryAccessLogWithTelemetryRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Begin:           core.CastInt64(data["begin"]),
		End:             core.CastInt64(data["end"]),
		LongTerm:        core.CastBool(data["longTerm"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p QueryAccessLogWithTelemetryRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"begin":           p.Begin,
		"end":             p.End,
		"longTerm":        p.LongTerm,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p QueryAccessLogWithTelemetryRequest) Pointer() *QueryAccessLogWithTelemetryRequest {
	return &p
}

type PutLogRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
	LogCategory        *string `json:"logCategory"`
	Payload            *string `json:"payload"`
}

func (p *PutLogRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PutLogRequest{}
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
		*p = PutLogRequest{}
	} else {
		*p = PutLogRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["loggingNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.LoggingNamespaceId); err != nil {
					return err
				}
			}
		}
		if v, ok := d["logCategory"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.LogCategory = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.LogCategory = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.LogCategory = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.LogCategory = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.LogCategory = &strValue
			default:
				if err := json.Unmarshal(*v, &p.LogCategory); err != nil {
					return err
				}
			}
		}
		if v, ok := d["payload"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
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
				if err := json.Unmarshal(*v, &p.Payload); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewPutLogRequestFromJson(data string) (PutLogRequest, error) {
	req := PutLogRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PutLogRequest{}, err
	}
	return req, nil
}

func NewPutLogRequestFromDict(data map[string]interface{}) PutLogRequest {
	return PutLogRequest{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
		LogCategory:        core.CastString(data["logCategory"]),
		Payload:            core.CastString(data["payload"]),
	}
}

func (p PutLogRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
		"logCategory":        p.LogCategory,
		"payload":            p.Payload,
	}
}

func (p PutLogRequest) Pointer() *PutLogRequest {
	return &p
}

type DescribeInsightsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func (p *DescribeInsightsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeInsightsRequest{}
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
		*p = DescribeInsightsRequest{}
	} else {
		*p = DescribeInsightsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.PageToken = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.PageToken = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.PageToken = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.PageToken = &strValue
			default:
				if err := json.Unmarshal(*v, &p.PageToken); err != nil {
					return err
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			if err := json.Unmarshal(*v, &p.Limit); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewDescribeInsightsRequestFromJson(data string) (DescribeInsightsRequest, error) {
	req := DescribeInsightsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeInsightsRequest{}, err
	}
	return req, nil
}

func NewDescribeInsightsRequestFromDict(data map[string]interface{}) DescribeInsightsRequest {
	return DescribeInsightsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeInsightsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeInsightsRequest) Pointer() *DescribeInsightsRequest {
	return &p
}

type CreateInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func (p *CreateInsightRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CreateInsightRequest{}
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
		*p = CreateInsightRequest{}
	} else {
		*p = CreateInsightRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewCreateInsightRequestFromJson(data string) (CreateInsightRequest, error) {
	req := CreateInsightRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return CreateInsightRequest{}, err
	}
	return req, nil
}

func NewCreateInsightRequestFromDict(data map[string]interface{}) CreateInsightRequest {
	return CreateInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p CreateInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p CreateInsightRequest) Pointer() *CreateInsightRequest {
	return &p
}

type GetInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InsightName     *string `json:"insightName"`
}

func (p *GetInsightRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetInsightRequest{}
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
		*p = GetInsightRequest{}
	} else {
		*p = GetInsightRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["insightName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.InsightName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.InsightName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.InsightName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.InsightName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.InsightName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.InsightName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewGetInsightRequestFromJson(data string) (GetInsightRequest, error) {
	req := GetInsightRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetInsightRequest{}, err
	}
	return req, nil
}

func NewGetInsightRequestFromDict(data map[string]interface{}) GetInsightRequest {
	return GetInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InsightName:   core.CastString(data["insightName"]),
	}
}

func (p GetInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"insightName":   p.InsightName,
	}
}

func (p GetInsightRequest) Pointer() *GetInsightRequest {
	return &p
}

type DeleteInsightRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	InsightName     *string `json:"insightName"`
}

func (p *DeleteInsightRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DeleteInsightRequest{}
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
		*p = DeleteInsightRequest{}
	} else {
		*p = DeleteInsightRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.NamespaceName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.NamespaceName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.NamespaceName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.NamespaceName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.NamespaceName); err != nil {
					return err
				}
			}
		}
		if v, ok := d["insightName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err != nil {
				return err
			}
			switch v2 := temp.(type) {
			case string:
				p.InsightName = &v2
			case float64:
				strValue := strconv.FormatFloat(v2, 'f', -1, 64)
				p.InsightName = &strValue
			case int:
				strValue := strconv.Itoa(v2)
				p.InsightName = &strValue
			case int32:
				strValue := strconv.Itoa(int(v2))
				p.InsightName = &strValue
			case int64:
				strValue := strconv.Itoa(int(v2))
				p.InsightName = &strValue
			default:
				if err := json.Unmarshal(*v, &p.InsightName); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func NewDeleteInsightRequestFromJson(data string) (DeleteInsightRequest, error) {
	req := DeleteInsightRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DeleteInsightRequest{}, err
	}
	return req, nil
}

func NewDeleteInsightRequestFromDict(data map[string]interface{}) DeleteInsightRequest {
	return DeleteInsightRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		InsightName:   core.CastString(data["insightName"]),
	}
}

func (p DeleteInsightRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"insightName":   p.InsightName,
	}
}

func (p DeleteInsightRequest) Pointer() *DeleteInsightRequest {
	return &p
}
