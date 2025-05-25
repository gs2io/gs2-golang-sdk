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

type Stage struct {
	StageId         *string `json:"stageId"`
	Name            *string `json:"name"`
	SourceStageName *string `json:"sourceStageName"`
	SortNumber      *int32  `json:"sortNumber"`
	Status          *string `json:"status"`
	CreatedAt       *int64  `json:"createdAt"`
	UpdatedAt       *int64  `json:"updatedAt"`
	Revision        *int64  `json:"revision"`
}

func (p *Stage) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Stage{}
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
		*p = Stage{}
	} else {
		*p = Stage{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stageId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StageId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StageId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StageId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StageId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StageId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StageId)
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
		if v, ok := d["sourceStageName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourceStageName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourceStageName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourceStageName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourceStageName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourceStageName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourceStageName)
				}
			}
		}
		if v, ok := d["sortNumber"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.SortNumber)
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

func NewStageFromJson(data string) Stage {
	req := Stage{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStageFromDict(data map[string]interface{}) Stage {
	return Stage{
		StageId: func() *string {
			v, ok := data["stageId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stageId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		SourceStageName: func() *string {
			v, ok := data["sourceStageName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourceStageName"])
		}(),
		SortNumber: func() *int32 {
			v, ok := data["sortNumber"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["sortNumber"])
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

func (p Stage) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StageId != nil {
		m["stageId"] = p.StageId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.SourceStageName != nil {
		m["sourceStageName"] = p.SourceStageName
	}
	if p.SortNumber != nil {
		m["sortNumber"] = p.SortNumber
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

func (p Stage) Pointer() *Stage {
	return &p
}

func CastStages(data []interface{}) []Stage {
	v := make([]Stage, 0)
	for _, d := range data {
		v = append(v, NewStageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStagesFromDict(data []Stage) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Output struct {
	OutputId  *string `json:"outputId"`
	Name      *string `json:"name"`
	Text      *string `json:"text"`
	CreatedAt *int64  `json:"createdAt"`
	Revision  *int64  `json:"revision"`
}

func (p *Output) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Output{}
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
		*p = Output{}
	} else {
		*p = Output{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["outputId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.OutputId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.OutputId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.OutputId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.OutputId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.OutputId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.OutputId)
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
		if v, ok := d["text"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Text = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Text = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Text = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Text = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Text = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Text)
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

func NewOutputFromJson(data string) Output {
	req := Output{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewOutputFromDict(data map[string]interface{}) Output {
	return Output{
		OutputId: func() *string {
			v, ok := data["outputId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["outputId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Text: func() *string {
			v, ok := data["text"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["text"])
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

func (p Output) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.OutputId != nil {
		m["outputId"] = p.OutputId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Text != nil {
		m["text"] = p.Text
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.Revision != nil {
		m["revision"] = p.Revision
	}
	return m
}

func (p Output) Pointer() *Output {
	return &p
}

func CastOutputs(data []interface{}) []Output {
	v := make([]Output, 0)
	for _, d := range data {
		v = append(v, NewOutputFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOutputsFromDict(data []Output) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Microservice struct {
	Name    *string `json:"name"`
	Version *string `json:"version"`
}

func (p *Microservice) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Microservice{}
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
		*p = Microservice{}
	} else {
		*p = Microservice{}
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
		if v, ok := d["version"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Version = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Version = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Version = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Version = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Version = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Version)
				}
			}
		}
	}
	return nil
}

func NewMicroserviceFromJson(data string) Microservice {
	req := Microservice{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMicroserviceFromDict(data map[string]interface{}) Microservice {
	return Microservice{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Version: func() *string {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["version"])
		}(),
	}
}

func (p Microservice) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Version != nil {
		m["version"] = p.Version
	}
	return m
}

func (p Microservice) Pointer() *Microservice {
	return &p
}

func CastMicroservices(data []interface{}) []Microservice {
	v := make([]Microservice, 0)
	for _, d := range data {
		v = append(v, NewMicroserviceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMicroservicesFromDict(data []Microservice) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
