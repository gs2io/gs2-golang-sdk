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

package freeze

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeStagesRequest struct {
	ContextStack *string `json:"contextStack"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *DescribeStagesRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeStagesRequest{}
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
		*p = DescribeStagesRequest{}
	} else {
		*p = DescribeStagesRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
	}
	return nil
}

func NewDescribeStagesRequestFromJson(data string) (DescribeStagesRequest, error) {
	req := DescribeStagesRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeStagesRequest{}, err
	}
	return req, nil
}

func NewDescribeStagesRequestFromDict(data map[string]interface{}) DescribeStagesRequest {
	return DescribeStagesRequest{}
}

func (p DescribeStagesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{}
}

func (p DescribeStagesRequest) Pointer() *DescribeStagesRequest {
	return &p
}

type GetStageRequest struct {
	ContextStack *string `json:"contextStack"`
	StageName    *string `json:"stageName"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *GetStageRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetStageRequest{}
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
		*p = GetStageRequest{}
	} else {
		*p = GetStageRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageName)
				}
			}
		}
	}
	return nil
}

func NewGetStageRequestFromJson(data string) (GetStageRequest, error) {
	req := GetStageRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetStageRequest{}, err
	}
	return req, nil
}

func NewGetStageRequestFromDict(data map[string]interface{}) GetStageRequest {
	return GetStageRequest{
		StageName: func() *string {
			v, ok := data["stageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageName"])
		}(),
	}
}

func (p GetStageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stageName": p.StageName,
	}
}

func (p GetStageRequest) Pointer() *GetStageRequest {
	return &p
}

type PromoteStageRequest struct {
	ContextStack *string `json:"contextStack"`
	StageName    *string `json:"stageName"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *PromoteStageRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = PromoteStageRequest{}
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
		*p = PromoteStageRequest{}
	} else {
		*p = PromoteStageRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageName)
				}
			}
		}
	}
	return nil
}

func NewPromoteStageRequestFromJson(data string) (PromoteStageRequest, error) {
	req := PromoteStageRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return PromoteStageRequest{}, err
	}
	return req, nil
}

func NewPromoteStageRequestFromDict(data map[string]interface{}) PromoteStageRequest {
	return PromoteStageRequest{
		StageName: func() *string {
			v, ok := data["stageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageName"])
		}(),
	}
}

func (p PromoteStageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stageName": p.StageName,
	}
}

func (p PromoteStageRequest) Pointer() *PromoteStageRequest {
	return &p
}

type RollbackStageRequest struct {
	ContextStack *string `json:"contextStack"`
	StageName    *string `json:"stageName"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *RollbackStageRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RollbackStageRequest{}
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
		*p = RollbackStageRequest{}
	} else {
		*p = RollbackStageRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageName)
				}
			}
		}
	}
	return nil
}

func NewRollbackStageRequestFromJson(data string) (RollbackStageRequest, error) {
	req := RollbackStageRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return RollbackStageRequest{}, err
	}
	return req, nil
}

func NewRollbackStageRequestFromDict(data map[string]interface{}) RollbackStageRequest {
	return RollbackStageRequest{
		StageName: func() *string {
			v, ok := data["stageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageName"])
		}(),
	}
}

func (p RollbackStageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stageName": p.StageName,
	}
}

func (p RollbackStageRequest) Pointer() *RollbackStageRequest {
	return &p
}

type DescribeOutputsRequest struct {
	ContextStack *string `json:"contextStack"`
	StageName    *string `json:"stageName"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *DescribeOutputsRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = DescribeOutputsRequest{}
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
		*p = DescribeOutputsRequest{}
	} else {
		*p = DescribeOutputsRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageName)
				}
			}
		}
		if v, ok := d["pageToken"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
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
					_ = json.Unmarshal(*v, &p.PageToken)
				}
			}
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewDescribeOutputsRequestFromJson(data string) (DescribeOutputsRequest, error) {
	req := DescribeOutputsRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return DescribeOutputsRequest{}, err
	}
	return req, nil
}

func NewDescribeOutputsRequestFromDict(data map[string]interface{}) DescribeOutputsRequest {
	return DescribeOutputsRequest{
		StageName: func() *string {
			v, ok := data["stageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageName"])
		}(),
		PageToken: func() *string {
			v, ok := data["pageToken"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["pageToken"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p DescribeOutputsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stageName": p.StageName,
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeOutputsRequest) Pointer() *DescribeOutputsRequest {
	return &p
}

type GetOutputRequest struct {
	ContextStack *string `json:"contextStack"`
	StageName    *string `json:"stageName"`
	OutputName   *string `json:"outputName"`
	DryRun       *bool   `json:"dryRun"`
}

func (p *GetOutputRequest) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GetOutputRequest{}
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
		*p = GetOutputRequest{}
	} else {
		*p = GetOutputRequest{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageName)
				}
			}
		}
		if v, ok := d["outputName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OutputName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OutputName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OutputName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OutputName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OutputName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OutputName)
				}
			}
		}
	}
	return nil
}

func NewGetOutputRequestFromJson(data string) (GetOutputRequest, error) {
	req := GetOutputRequest{}
	err := json.Unmarshal([]byte(data), &req)
	if err != nil {
		return GetOutputRequest{}, err
	}
	return req, nil
}

func NewGetOutputRequestFromDict(data map[string]interface{}) GetOutputRequest {
	return GetOutputRequest{
		StageName: func() *string {
			v, ok := data["stageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageName"])
		}(),
		OutputName: func() *string {
			v, ok := data["outputName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["outputName"])
		}(),
	}
}

func (p GetOutputRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stageName":  p.StageName,
		"outputName": p.OutputName,
	}
}

func (p GetOutputRequest) Pointer() *GetOutputRequest {
	return &p
}
