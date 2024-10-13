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

package deploy

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Stack struct {
	StackId     *string `json:"stackId"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Template    *string `json:"template"`
	Status      *string `json:"status"`
	CreatedAt   *int64  `json:"createdAt"`
	UpdatedAt   *int64  `json:"updatedAt"`
	Revision    *int64  `json:"revision"`
}

func (p *Stack) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Stack{}
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
		*p = Stack{}
	} else {
		*p = Stack{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stackId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StackId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StackId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StackId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StackId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StackId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StackId)
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
		if v, ok := d["template"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Template = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Template = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Template = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Template = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Template)
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

func NewStackFromJson(data string) Stack {
	req := Stack{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStackFromDict(data map[string]interface{}) Stack {
	return Stack{
		StackId: func() *string {
			v, ok := data["stackId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stackId"])
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
		Template: func() *string {
			v, ok := data["template"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["template"])
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

func (p Stack) ToDict() map[string]interface{} {

	var stackId *string
	if p.StackId != nil {
		stackId = p.StackId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var template *string
	if p.Template != nil {
		template = p.Template
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
		"stackId":     stackId,
		"name":        name,
		"description": description,
		"template":    template,
		"status":      status,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
	}
}

func (p Stack) Pointer() *Stack {
	return &p
}

func CastStacks(data []interface{}) []Stack {
	v := make([]Stack, 0)
	for _, d := range data {
		v = append(v, NewStackFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStacksFromDict(data []Stack) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Resource struct {
	ResourceId      *string       `json:"resourceId"`
	Type            *string       `json:"type"`
	Name            *string       `json:"name"`
	Request         *string       `json:"request"`
	Response        *string       `json:"response"`
	RollbackContext *string       `json:"rollbackContext"`
	RollbackRequest *string       `json:"rollbackRequest"`
	RollbackAfter   []*string     `json:"rollbackAfter"`
	OutputFields    []OutputField `json:"outputFields"`
	WorkId          *string       `json:"workId"`
	CreatedAt       *int64        `json:"createdAt"`
}

func (p *Resource) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Resource{}
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
		*p = Resource{}
	} else {
		*p = Resource{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["resourceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceId)
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
		if v, ok := d["response"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Response = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Response = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Response = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Response = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Response = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Response)
				}
			}
		}
		if v, ok := d["rollbackContext"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RollbackContext = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RollbackContext = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RollbackContext = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RollbackContext = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RollbackContext = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RollbackContext)
				}
			}
		}
		if v, ok := d["rollbackRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RollbackRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RollbackRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RollbackRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RollbackRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RollbackRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RollbackRequest)
				}
			}
		}
		if v, ok := d["rollbackAfter"]; ok && v != nil {
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
				p.RollbackAfter = l
			}
		}
		if v, ok := d["outputFields"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.OutputFields)
		}
		if v, ok := d["workId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.WorkId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.WorkId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.WorkId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.WorkId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.WorkId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.WorkId)
				}
			}
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewResourceFromJson(data string) Resource {
	req := Resource{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewResourceFromDict(data map[string]interface{}) Resource {
	return Resource{
		ResourceId: func() *string {
			v, ok := data["resourceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceId"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Request: func() *string {
			v, ok := data["request"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["request"])
		}(),
		Response: func() *string {
			v, ok := data["response"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["response"])
		}(),
		RollbackContext: func() *string {
			v, ok := data["rollbackContext"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rollbackContext"])
		}(),
		RollbackRequest: func() *string {
			v, ok := data["rollbackRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["rollbackRequest"])
		}(),
		RollbackAfter: func() []*string {
			v, ok := data["rollbackAfter"]
			if !ok || v == nil {
				return nil
			}
			return core.CastStrings(core.CastArray(v))
		}(),
		OutputFields: func() []OutputField {
			if data["outputFields"] == nil {
				return nil
			}
			return CastOutputFields(core.CastArray(data["outputFields"]))
		}(),
		WorkId: func() *string {
			v, ok := data["workId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["workId"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p Resource) ToDict() map[string]interface{} {

	var resourceId *string
	if p.ResourceId != nil {
		resourceId = p.ResourceId
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var request *string
	if p.Request != nil {
		request = p.Request
	}
	var response *string
	if p.Response != nil {
		response = p.Response
	}
	var rollbackContext *string
	if p.RollbackContext != nil {
		rollbackContext = p.RollbackContext
	}
	var rollbackRequest *string
	if p.RollbackRequest != nil {
		rollbackRequest = p.RollbackRequest
	}
	var rollbackAfter []interface{}
	if p.RollbackAfter != nil {
		rollbackAfter = core.CastStringsFromDict(
			p.RollbackAfter,
		)
	}
	var outputFields []interface{}
	if p.OutputFields != nil {
		outputFields = CastOutputFieldsFromDict(
			p.OutputFields,
		)
	}
	var workId *string
	if p.WorkId != nil {
		workId = p.WorkId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"resourceId":      resourceId,
		"type":            _type,
		"name":            name,
		"request":         request,
		"response":        response,
		"rollbackContext": rollbackContext,
		"rollbackRequest": rollbackRequest,
		"rollbackAfter":   rollbackAfter,
		"outputFields":    outputFields,
		"workId":          workId,
		"createdAt":       createdAt,
	}
}

func (p Resource) Pointer() *Resource {
	return &p
}

func CastResources(data []interface{}) []Resource {
	v := make([]Resource, 0)
	for _, d := range data {
		v = append(v, NewResourceFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastResourcesFromDict(data []Resource) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Event struct {
	EventId      *string `json:"eventId"`
	Name         *string `json:"name"`
	ResourceName *string `json:"resourceName"`
	Type         *string `json:"type"`
	Message      *string `json:"message"`
	EventAt      *int64  `json:"eventAt"`
	Revision     *int64  `json:"revision"`
}

func (p *Event) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Event{}
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
		*p = Event{}
	} else {
		*p = Event{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["eventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventId)
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
		if v, ok := d["resourceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceName)
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
		if v, ok := d["eventAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EventAt)
		}
		if v, ok := d["revision"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Revision)
		}
	}
	return nil
}

func NewEventFromJson(data string) Event {
	req := Event{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewEventFromDict(data map[string]interface{}) Event {
	return Event{
		EventId: func() *string {
			v, ok := data["eventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		ResourceName: func() *string {
			v, ok := data["resourceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceName"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		Message: func() *string {
			v, ok := data["message"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["message"])
		}(),
		EventAt: func() *int64 {
			v, ok := data["eventAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["eventAt"])
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

func (p Event) ToDict() map[string]interface{} {

	var eventId *string
	if p.EventId != nil {
		eventId = p.EventId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var resourceName *string
	if p.ResourceName != nil {
		resourceName = p.ResourceName
	}
	var _type *string
	if p.Type != nil {
		_type = p.Type
	}
	var message *string
	if p.Message != nil {
		message = p.Message
	}
	var eventAt *int64
	if p.EventAt != nil {
		eventAt = p.EventAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"eventId":      eventId,
		"name":         name,
		"resourceName": resourceName,
		"type":         _type,
		"message":      message,
		"eventAt":      eventAt,
		"revision":     revision,
	}
}

func (p Event) Pointer() *Event {
	return &p
}

func CastEvents(data []interface{}) []Event {
	v := make([]Event, 0)
	for _, d := range data {
		v = append(v, NewEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEventsFromDict(data []Event) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Output struct {
	OutputId  *string `json:"outputId"`
	Name      *string `json:"name"`
	Value     *string `json:"value"`
	CreatedAt *int64  `json:"createdAt"`
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
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
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
		Value: func() *string {
			v, ok := data["value"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["value"])
		}(),
		CreatedAt: func() *int64 {
			v, ok := data["createdAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["createdAt"])
		}(),
	}
}

func (p Output) ToDict() map[string]interface{} {

	var outputId *string
	if p.OutputId != nil {
		outputId = p.OutputId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	return map[string]interface{}{
		"outputId":  outputId,
		"name":      name,
		"value":     value,
		"createdAt": createdAt,
	}
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

type OutputField struct {
	Name      *string `json:"name"`
	FieldName *string `json:"fieldName"`
}

func (p *OutputField) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = OutputField{}
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
		*p = OutputField{}
	} else {
		*p = OutputField{}
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
		if v, ok := d["fieldName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.FieldName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.FieldName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.FieldName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.FieldName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.FieldName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.FieldName)
				}
			}
		}
	}
	return nil
}

func NewOutputFieldFromJson(data string) OutputField {
	req := OutputField{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewOutputFieldFromDict(data map[string]interface{}) OutputField {
	return OutputField{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		FieldName: func() *string {
			v, ok := data["fieldName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["fieldName"])
		}(),
	}
}

func (p OutputField) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var fieldName *string
	if p.FieldName != nil {
		fieldName = p.FieldName
	}
	return map[string]interface{}{
		"name":      name,
		"fieldName": fieldName,
	}
}

func (p OutputField) Pointer() *OutputField {
	return &p
}

func CastOutputFields(data []interface{}) []OutputField {
	v := make([]OutputField, 0)
	for _, d := range data {
		v = append(v, NewOutputFieldFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOutputFieldsFromDict(data []OutputField) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ChangeSet struct {
	ResourceName *string `json:"resourceName"`
	ResourceType *string `json:"resourceType"`
	Operation    *string `json:"operation"`
}

func (p *ChangeSet) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ChangeSet{}
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
		*p = ChangeSet{}
	} else {
		*p = ChangeSet{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["resourceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceName)
				}
			}
		}
		if v, ok := d["resourceType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ResourceType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ResourceType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ResourceType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ResourceType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ResourceType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ResourceType)
				}
			}
		}
		if v, ok := d["operation"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Operation = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Operation = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Operation = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Operation = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Operation = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Operation)
				}
			}
		}
	}
	return nil
}

func NewChangeSetFromJson(data string) ChangeSet {
	req := ChangeSet{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewChangeSetFromDict(data map[string]interface{}) ChangeSet {
	return ChangeSet{
		ResourceName: func() *string {
			v, ok := data["resourceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceName"])
		}(),
		ResourceType: func() *string {
			v, ok := data["resourceType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["resourceType"])
		}(),
		Operation: func() *string {
			v, ok := data["operation"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["operation"])
		}(),
	}
}

func (p ChangeSet) ToDict() map[string]interface{} {

	var resourceName *string
	if p.ResourceName != nil {
		resourceName = p.ResourceName
	}
	var resourceType *string
	if p.ResourceType != nil {
		resourceType = p.ResourceType
	}
	var operation *string
	if p.Operation != nil {
		operation = p.Operation
	}
	return map[string]interface{}{
		"resourceName": resourceName,
		"resourceType": resourceType,
		"operation":    operation,
	}
}

func (p ChangeSet) Pointer() *ChangeSet {
	return &p
}

func CastChangeSets(data []interface{}) []ChangeSet {
	v := make([]ChangeSet, 0)
	for _, d := range data {
		v = append(v, NewChangeSetFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChangeSetsFromDict(data []ChangeSet) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId       *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath     *string `json:"sourcePath"`
	ReferenceType  *string `json:"referenceType"`
	CommitHash     *string `json:"commitHash"`
	BranchName     *string `json:"branchName"`
	TagName        *string `json:"tagName"`
}

func (p *GitHubCheckoutSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = GitHubCheckoutSetting{}
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
		*p = GitHubCheckoutSetting{}
	} else {
		*p = GitHubCheckoutSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["apiKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApiKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApiKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApiKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApiKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApiKeyId)
				}
			}
		}
		if v, ok := d["repositoryName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.RepositoryName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.RepositoryName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.RepositoryName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.RepositoryName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.RepositoryName)
				}
			}
		}
		if v, ok := d["sourcePath"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SourcePath = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SourcePath = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SourcePath = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SourcePath = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SourcePath)
				}
			}
		}
		if v, ok := d["referenceType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ReferenceType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ReferenceType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ReferenceType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ReferenceType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ReferenceType)
				}
			}
		}
		if v, ok := d["commitHash"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CommitHash = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CommitHash = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CommitHash = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CommitHash = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CommitHash)
				}
			}
		}
		if v, ok := d["branchName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BranchName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BranchName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BranchName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BranchName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BranchName)
				}
			}
		}
		if v, ok := d["tagName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TagName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TagName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TagName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TagName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TagName)
				}
			}
		}
	}
	return nil
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
	req := GitHubCheckoutSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
	return GitHubCheckoutSetting{
		ApiKeyId: func() *string {
			v, ok := data["apiKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["apiKeyId"])
		}(),
		RepositoryName: func() *string {
			v, ok := data["repositoryName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["repositoryName"])
		}(),
		SourcePath: func() *string {
			v, ok := data["sourcePath"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["sourcePath"])
		}(),
		ReferenceType: func() *string {
			v, ok := data["referenceType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["referenceType"])
		}(),
		CommitHash: func() *string {
			v, ok := data["commitHash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["commitHash"])
		}(),
		BranchName: func() *string {
			v, ok := data["branchName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["branchName"])
		}(),
		TagName: func() *string {
			v, ok := data["tagName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["tagName"])
		}(),
	}
}

func (p GitHubCheckoutSetting) ToDict() map[string]interface{} {

	var apiKeyId *string
	if p.ApiKeyId != nil {
		apiKeyId = p.ApiKeyId
	}
	var repositoryName *string
	if p.RepositoryName != nil {
		repositoryName = p.RepositoryName
	}
	var sourcePath *string
	if p.SourcePath != nil {
		sourcePath = p.SourcePath
	}
	var referenceType *string
	if p.ReferenceType != nil {
		referenceType = p.ReferenceType
	}
	var commitHash *string
	if p.CommitHash != nil {
		commitHash = p.CommitHash
	}
	var branchName *string
	if p.BranchName != nil {
		branchName = p.BranchName
	}
	var tagName *string
	if p.TagName != nil {
		tagName = p.TagName
	}
	return map[string]interface{}{
		"apiKeyId":       apiKeyId,
		"repositoryName": repositoryName,
		"sourcePath":     sourcePath,
		"referenceType":  referenceType,
		"commitHash":     commitHash,
		"branchName":     branchName,
		"tagName":        tagName,
	}
}

func (p GitHubCheckoutSetting) Pointer() *GitHubCheckoutSetting {
	return &p
}

func CastGitHubCheckoutSettings(data []interface{}) []GitHubCheckoutSetting {
	v := make([]GitHubCheckoutSetting, 0)
	for _, d := range data {
		v = append(v, NewGitHubCheckoutSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastGitHubCheckoutSettingsFromDict(data []GitHubCheckoutSetting) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
