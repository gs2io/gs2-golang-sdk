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

package buff

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId     *string        `json:"namespaceId"`
	Name            *string        `json:"name"`
	Description     *string        `json:"description"`
	ApplyBuffScript *ScriptSetting `json:"applyBuffScript"`
	LogSetting      *LogSetting    `json:"logSetting"`
	CreatedAt       *int64         `json:"createdAt"`
	UpdatedAt       *int64         `json:"updatedAt"`
	Revision        *int64         `json:"revision"`
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
		if v, ok := d["applyBuffScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ApplyBuffScript)
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
		ApplyBuffScript: func() *ScriptSetting {
			v, ok := data["applyBuffScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["applyBuffScript"])).Pointer()
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
	var applyBuffScript map[string]interface{}
	if p.ApplyBuffScript != nil {
		applyBuffScript = func() map[string]interface{} {
			if p.ApplyBuffScript == nil {
				return nil
			}
			return p.ApplyBuffScript.ToDict()
		}()
	}
	var logSetting map[string]interface{}
	if p.LogSetting != nil {
		logSetting = func() map[string]interface{} {
			if p.LogSetting == nil {
				return nil
			}
			return p.LogSetting.ToDict()
		}()
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
		"namespaceId":     namespaceId,
		"name":            name,
		"description":     description,
		"applyBuffScript": applyBuffScript,
		"logSetting":      logSetting,
		"createdAt":       createdAt,
		"updatedAt":       updatedAt,
		"revision":        revision,
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

type BuffTargetModel struct {
	TargetModelName *string         `json:"targetModelName"`
	TargetFieldName *string         `json:"targetFieldName"`
	ConditionGrns   []BuffTargetGrn `json:"conditionGrns"`
	Rate            *float32        `json:"rate"`
}

func (p *BuffTargetModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BuffTargetModel{}
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
		*p = BuffTargetModel{}
	} else {
		*p = BuffTargetModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["targetModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetModelName)
				}
			}
		}
		if v, ok := d["targetFieldName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetFieldName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetFieldName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetFieldName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetFieldName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetFieldName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetFieldName)
				}
			}
		}
		if v, ok := d["conditionGrns"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConditionGrns)
		}
		if v, ok := d["rate"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rate)
		}
	}
	return nil
}

func NewBuffTargetModelFromJson(data string) BuffTargetModel {
	req := BuffTargetModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBuffTargetModelFromDict(data map[string]interface{}) BuffTargetModel {
	return BuffTargetModel{
		TargetModelName: func() *string {
			v, ok := data["targetModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetModelName"])
		}(),
		TargetFieldName: func() *string {
			v, ok := data["targetFieldName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetFieldName"])
		}(),
		ConditionGrns: func() []BuffTargetGrn {
			if data["conditionGrns"] == nil {
				return nil
			}
			return CastBuffTargetGrns(core.CastArray(data["conditionGrns"]))
		}(),
		Rate: func() *float32 {
			v, ok := data["rate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["rate"])
		}(),
	}
}

func (p BuffTargetModel) ToDict() map[string]interface{} {

	var targetModelName *string
	if p.TargetModelName != nil {
		targetModelName = p.TargetModelName
	}
	var targetFieldName *string
	if p.TargetFieldName != nil {
		targetFieldName = p.TargetFieldName
	}
	var conditionGrns []interface{}
	if p.ConditionGrns != nil {
		conditionGrns = CastBuffTargetGrnsFromDict(
			p.ConditionGrns,
		)
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	return map[string]interface{}{
		"targetModelName": targetModelName,
		"targetFieldName": targetFieldName,
		"conditionGrns":   conditionGrns,
		"rate":            rate,
	}
}

func (p BuffTargetModel) Pointer() *BuffTargetModel {
	return &p
}

func CastBuffTargetModels(data []interface{}) []BuffTargetModel {
	v := make([]BuffTargetModel, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetModelsFromDict(data []BuffTargetModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffTargetAction struct {
	TargetActionName *string         `json:"targetActionName"`
	TargetFieldName  *string         `json:"targetFieldName"`
	ConditionGrns    []BuffTargetGrn `json:"conditionGrns"`
	Rate             *float32        `json:"rate"`
}

func (p *BuffTargetAction) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BuffTargetAction{}
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
		*p = BuffTargetAction{}
	} else {
		*p = BuffTargetAction{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["targetActionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetActionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetActionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetActionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetActionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetActionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetActionName)
				}
			}
		}
		if v, ok := d["targetFieldName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetFieldName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetFieldName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetFieldName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetFieldName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetFieldName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetFieldName)
				}
			}
		}
		if v, ok := d["conditionGrns"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConditionGrns)
		}
		if v, ok := d["rate"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rate)
		}
	}
	return nil
}

func NewBuffTargetActionFromJson(data string) BuffTargetAction {
	req := BuffTargetAction{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBuffTargetActionFromDict(data map[string]interface{}) BuffTargetAction {
	return BuffTargetAction{
		TargetActionName: func() *string {
			v, ok := data["targetActionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetActionName"])
		}(),
		TargetFieldName: func() *string {
			v, ok := data["targetFieldName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetFieldName"])
		}(),
		ConditionGrns: func() []BuffTargetGrn {
			if data["conditionGrns"] == nil {
				return nil
			}
			return CastBuffTargetGrns(core.CastArray(data["conditionGrns"]))
		}(),
		Rate: func() *float32 {
			v, ok := data["rate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["rate"])
		}(),
	}
}

func (p BuffTargetAction) ToDict() map[string]interface{} {

	var targetActionName *string
	if p.TargetActionName != nil {
		targetActionName = p.TargetActionName
	}
	var targetFieldName *string
	if p.TargetFieldName != nil {
		targetFieldName = p.TargetFieldName
	}
	var conditionGrns []interface{}
	if p.ConditionGrns != nil {
		conditionGrns = CastBuffTargetGrnsFromDict(
			p.ConditionGrns,
		)
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	return map[string]interface{}{
		"targetActionName": targetActionName,
		"targetFieldName":  targetFieldName,
		"conditionGrns":    conditionGrns,
		"rate":             rate,
	}
}

func (p BuffTargetAction) Pointer() *BuffTargetAction {
	return &p
}

func CastBuffTargetActions(data []interface{}) []BuffTargetAction {
	v := make([]BuffTargetAction, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetActionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetActionsFromDict(data []BuffTargetAction) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffTargetGrn struct {
	TargetModelName *string `json:"targetModelName"`
	TargetGrn       *string `json:"targetGrn"`
}

func (p *BuffTargetGrn) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BuffTargetGrn{}
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
		*p = BuffTargetGrn{}
	} else {
		*p = BuffTargetGrn{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["targetModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetModelName)
				}
			}
		}
		if v, ok := d["targetGrn"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetGrn = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetGrn = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetGrn = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetGrn = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetGrn = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetGrn)
				}
			}
		}
	}
	return nil
}

func NewBuffTargetGrnFromJson(data string) BuffTargetGrn {
	req := BuffTargetGrn{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBuffTargetGrnFromDict(data map[string]interface{}) BuffTargetGrn {
	return BuffTargetGrn{
		TargetModelName: func() *string {
			v, ok := data["targetModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetModelName"])
		}(),
		TargetGrn: func() *string {
			v, ok := data["targetGrn"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetGrn"])
		}(),
	}
}

func (p BuffTargetGrn) ToDict() map[string]interface{} {

	var targetModelName *string
	if p.TargetModelName != nil {
		targetModelName = p.TargetModelName
	}
	var targetGrn *string
	if p.TargetGrn != nil {
		targetGrn = p.TargetGrn
	}
	return map[string]interface{}{
		"targetModelName": targetModelName,
		"targetGrn":       targetGrn,
	}
}

func (p BuffTargetGrn) Pointer() *BuffTargetGrn {
	return &p
}

func CastBuffTargetGrns(data []interface{}) []BuffTargetGrn {
	v := make([]BuffTargetGrn, 0)
	for _, d := range data {
		v = append(v, NewBuffTargetGrnFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffTargetGrnsFromDict(data []BuffTargetGrn) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffEntryModel struct {
	BuffEntryModelId           *string           `json:"buffEntryModelId"`
	Name                       *string           `json:"name"`
	Metadata                   *string           `json:"metadata"`
	Expression                 *string           `json:"expression"`
	TargetType                 *string           `json:"targetType"`
	TargetModel                *BuffTargetModel  `json:"targetModel"`
	TargetAction               *BuffTargetAction `json:"targetAction"`
	Priority                   *int32            `json:"priority"`
	ApplyPeriodScheduleEventId *string           `json:"applyPeriodScheduleEventId"`
}

func (p *BuffEntryModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BuffEntryModel{}
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
		*p = BuffEntryModel{}
	} else {
		*p = BuffEntryModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["buffEntryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BuffEntryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BuffEntryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BuffEntryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BuffEntryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BuffEntryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BuffEntryModelId)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["expression"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Expression = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Expression = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Expression = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Expression = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Expression = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Expression)
				}
			}
		}
		if v, ok := d["targetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetType)
				}
			}
		}
		if v, ok := d["targetModel"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetModel)
		}
		if v, ok := d["targetAction"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetAction)
		}
		if v, ok := d["priority"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Priority)
		}
		if v, ok := d["applyPeriodScheduleEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApplyPeriodScheduleEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApplyPeriodScheduleEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApplyPeriodScheduleEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApplyPeriodScheduleEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApplyPeriodScheduleEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApplyPeriodScheduleEventId)
				}
			}
		}
	}
	return nil
}

func NewBuffEntryModelFromJson(data string) BuffEntryModel {
	req := BuffEntryModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBuffEntryModelFromDict(data map[string]interface{}) BuffEntryModel {
	return BuffEntryModel{
		BuffEntryModelId: func() *string {
			v, ok := data["buffEntryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["buffEntryModelId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		Expression: func() *string {
			v, ok := data["expression"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["expression"])
		}(),
		TargetType: func() *string {
			v, ok := data["targetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetType"])
		}(),
		TargetModel: func() *BuffTargetModel {
			v, ok := data["targetModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBuffTargetModelFromDict(core.CastMap(data["targetModel"])).Pointer()
		}(),
		TargetAction: func() *BuffTargetAction {
			v, ok := data["targetAction"]
			if !ok || v == nil {
				return nil
			}
			return NewBuffTargetActionFromDict(core.CastMap(data["targetAction"])).Pointer()
		}(),
		Priority: func() *int32 {
			v, ok := data["priority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["priority"])
		}(),
		ApplyPeriodScheduleEventId: func() *string {
			v, ok := data["applyPeriodScheduleEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["applyPeriodScheduleEventId"])
		}(),
	}
}

func (p BuffEntryModel) ToDict() map[string]interface{} {

	var buffEntryModelId *string
	if p.BuffEntryModelId != nil {
		buffEntryModelId = p.BuffEntryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var expression *string
	if p.Expression != nil {
		expression = p.Expression
	}
	var targetType *string
	if p.TargetType != nil {
		targetType = p.TargetType
	}
	var targetModel map[string]interface{}
	if p.TargetModel != nil {
		targetModel = func() map[string]interface{} {
			if p.TargetModel == nil {
				return nil
			}
			return p.TargetModel.ToDict()
		}()
	}
	var targetAction map[string]interface{}
	if p.TargetAction != nil {
		targetAction = func() map[string]interface{} {
			if p.TargetAction == nil {
				return nil
			}
			return p.TargetAction.ToDict()
		}()
	}
	var priority *int32
	if p.Priority != nil {
		priority = p.Priority
	}
	var applyPeriodScheduleEventId *string
	if p.ApplyPeriodScheduleEventId != nil {
		applyPeriodScheduleEventId = p.ApplyPeriodScheduleEventId
	}
	return map[string]interface{}{
		"buffEntryModelId":           buffEntryModelId,
		"name":                       name,
		"metadata":                   metadata,
		"expression":                 expression,
		"targetType":                 targetType,
		"targetModel":                targetModel,
		"targetAction":               targetAction,
		"priority":                   priority,
		"applyPeriodScheduleEventId": applyPeriodScheduleEventId,
	}
}

func (p BuffEntryModel) Pointer() *BuffEntryModel {
	return &p
}

func CastBuffEntryModels(data []interface{}) []BuffEntryModel {
	v := make([]BuffEntryModel, 0)
	for _, d := range data {
		v = append(v, NewBuffEntryModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffEntryModelsFromDict(data []BuffEntryModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type BuffEntryModelMaster struct {
	BuffEntryModelId           *string           `json:"buffEntryModelId"`
	Name                       *string           `json:"name"`
	Description                *string           `json:"description"`
	Metadata                   *string           `json:"metadata"`
	Expression                 *string           `json:"expression"`
	TargetType                 *string           `json:"targetType"`
	TargetModel                *BuffTargetModel  `json:"targetModel"`
	TargetAction               *BuffTargetAction `json:"targetAction"`
	Priority                   *int32            `json:"priority"`
	ApplyPeriodScheduleEventId *string           `json:"applyPeriodScheduleEventId"`
	CreatedAt                  *int64            `json:"createdAt"`
	UpdatedAt                  *int64            `json:"updatedAt"`
	Revision                   *int64            `json:"revision"`
}

func (p *BuffEntryModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = BuffEntryModelMaster{}
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
		*p = BuffEntryModelMaster{}
	} else {
		*p = BuffEntryModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["buffEntryModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.BuffEntryModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.BuffEntryModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.BuffEntryModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.BuffEntryModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.BuffEntryModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.BuffEntryModelId)
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
		if v, ok := d["metadata"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Metadata = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Metadata = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Metadata = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Metadata = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Metadata)
				}
			}
		}
		if v, ok := d["expression"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Expression = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Expression = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Expression = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Expression = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Expression = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Expression)
				}
			}
		}
		if v, ok := d["targetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TargetType)
				}
			}
		}
		if v, ok := d["targetModel"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetModel)
		}
		if v, ok := d["targetAction"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TargetAction)
		}
		if v, ok := d["priority"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Priority)
		}
		if v, ok := d["applyPeriodScheduleEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ApplyPeriodScheduleEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ApplyPeriodScheduleEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ApplyPeriodScheduleEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ApplyPeriodScheduleEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ApplyPeriodScheduleEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ApplyPeriodScheduleEventId)
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

func NewBuffEntryModelMasterFromJson(data string) BuffEntryModelMaster {
	req := BuffEntryModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewBuffEntryModelMasterFromDict(data map[string]interface{}) BuffEntryModelMaster {
	return BuffEntryModelMaster{
		BuffEntryModelId: func() *string {
			v, ok := data["buffEntryModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["buffEntryModelId"])
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
		Metadata: func() *string {
			v, ok := data["metadata"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["metadata"])
		}(),
		Expression: func() *string {
			v, ok := data["expression"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["expression"])
		}(),
		TargetType: func() *string {
			v, ok := data["targetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["targetType"])
		}(),
		TargetModel: func() *BuffTargetModel {
			v, ok := data["targetModel"]
			if !ok || v == nil {
				return nil
			}
			return NewBuffTargetModelFromDict(core.CastMap(data["targetModel"])).Pointer()
		}(),
		TargetAction: func() *BuffTargetAction {
			v, ok := data["targetAction"]
			if !ok || v == nil {
				return nil
			}
			return NewBuffTargetActionFromDict(core.CastMap(data["targetAction"])).Pointer()
		}(),
		Priority: func() *int32 {
			v, ok := data["priority"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["priority"])
		}(),
		ApplyPeriodScheduleEventId: func() *string {
			v, ok := data["applyPeriodScheduleEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["applyPeriodScheduleEventId"])
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

func (p BuffEntryModelMaster) ToDict() map[string]interface{} {

	var buffEntryModelId *string
	if p.BuffEntryModelId != nil {
		buffEntryModelId = p.BuffEntryModelId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var description *string
	if p.Description != nil {
		description = p.Description
	}
	var metadata *string
	if p.Metadata != nil {
		metadata = p.Metadata
	}
	var expression *string
	if p.Expression != nil {
		expression = p.Expression
	}
	var targetType *string
	if p.TargetType != nil {
		targetType = p.TargetType
	}
	var targetModel map[string]interface{}
	if p.TargetModel != nil {
		targetModel = func() map[string]interface{} {
			if p.TargetModel == nil {
				return nil
			}
			return p.TargetModel.ToDict()
		}()
	}
	var targetAction map[string]interface{}
	if p.TargetAction != nil {
		targetAction = func() map[string]interface{} {
			if p.TargetAction == nil {
				return nil
			}
			return p.TargetAction.ToDict()
		}()
	}
	var priority *int32
	if p.Priority != nil {
		priority = p.Priority
	}
	var applyPeriodScheduleEventId *string
	if p.ApplyPeriodScheduleEventId != nil {
		applyPeriodScheduleEventId = p.ApplyPeriodScheduleEventId
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
		"buffEntryModelId":           buffEntryModelId,
		"name":                       name,
		"description":                description,
		"metadata":                   metadata,
		"expression":                 expression,
		"targetType":                 targetType,
		"targetModel":                targetModel,
		"targetAction":               targetAction,
		"priority":                   priority,
		"applyPeriodScheduleEventId": applyPeriodScheduleEventId,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
		"revision":                   revision,
	}
}

func (p BuffEntryModelMaster) Pointer() *BuffEntryModelMaster {
	return &p
}

func CastBuffEntryModelMasters(data []interface{}) []BuffEntryModelMaster {
	v := make([]BuffEntryModelMaster, 0)
	for _, d := range data {
		v = append(v, NewBuffEntryModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastBuffEntryModelMastersFromDict(data []BuffEntryModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentBuffMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentBuffMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentBuffMaster{}
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
		*p = CurrentBuffMaster{}
	} else {
		*p = CurrentBuffMaster{}
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
		if v, ok := d["settings"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Settings = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Settings = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Settings = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Settings = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Settings)
				}
			}
		}
	}
	return nil
}

func NewCurrentBuffMasterFromJson(data string) CurrentBuffMaster {
	req := CurrentBuffMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentBuffMasterFromDict(data map[string]interface{}) CurrentBuffMaster {
	return CurrentBuffMaster{
		NamespaceId: func() *string {
			v, ok := data["namespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceId"])
		}(),
		Settings: func() *string {
			v, ok := data["settings"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["settings"])
		}(),
	}
}

func (p CurrentBuffMaster) ToDict() map[string]interface{} {

	var namespaceId *string
	if p.NamespaceId != nil {
		namespaceId = p.NamespaceId
	}
	var settings *string
	if p.Settings != nil {
		settings = p.Settings
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"settings":    settings,
	}
}

func (p CurrentBuffMaster) Pointer() *CurrentBuffMaster {
	return &p
}

func CastCurrentBuffMasters(data []interface{}) []CurrentBuffMaster {
	v := make([]CurrentBuffMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentBuffMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentBuffMastersFromDict(data []CurrentBuffMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type OverrideBuffRate struct {
	Name *string  `json:"name"`
	Rate *float32 `json:"rate"`
}

func (p *OverrideBuffRate) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = OverrideBuffRate{}
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
		*p = OverrideBuffRate{}
	} else {
		*p = OverrideBuffRate{}
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
		if v, ok := d["rate"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Rate)
		}
	}
	return nil
}

func NewOverrideBuffRateFromJson(data string) OverrideBuffRate {
	req := OverrideBuffRate{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewOverrideBuffRateFromDict(data map[string]interface{}) OverrideBuffRate {
	return OverrideBuffRate{
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		Rate: func() *float32 {
			v, ok := data["rate"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["rate"])
		}(),
	}
}

func (p OverrideBuffRate) ToDict() map[string]interface{} {

	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var rate *float32
	if p.Rate != nil {
		rate = p.Rate
	}
	return map[string]interface{}{
		"name": name,
		"rate": rate,
	}
}

func (p OverrideBuffRate) Pointer() *OverrideBuffRate {
	return &p
}

func CastOverrideBuffRates(data []interface{}) []OverrideBuffRate {
	v := make([]OverrideBuffRate, 0)
	for _, d := range data {
		v = append(v, NewOverrideBuffRateFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastOverrideBuffRatesFromDict(data []OverrideBuffRate) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScriptSetting struct {
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func (p *ScriptSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScriptSetting{}
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
		*p = ScriptSetting{}
	} else {
		*p = ScriptSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["triggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerTargetType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerTargetType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerTargetType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerTargetType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerTargetType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerTargetType)
				}
			}
		}
		if v, ok := d["doneTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerScriptId)
				}
			}
		}
		if v, ok := d["doneTriggerQueueNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.DoneTriggerQueueNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.DoneTriggerQueueNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.DoneTriggerQueueNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.DoneTriggerQueueNamespaceId)
				}
			}
		}
	}
	return nil
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	req := ScriptSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
	return ScriptSetting{
		TriggerScriptId: func() *string {
			v, ok := data["triggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["triggerScriptId"])
		}(),
		DoneTriggerTargetType: func() *string {
			v, ok := data["doneTriggerTargetType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerTargetType"])
		}(),
		DoneTriggerScriptId: func() *string {
			v, ok := data["doneTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerScriptId"])
		}(),
		DoneTriggerQueueNamespaceId: func() *string {
			v, ok := data["doneTriggerQueueNamespaceId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["doneTriggerQueueNamespaceId"])
		}(),
	}
}

func (p ScriptSetting) ToDict() map[string]interface{} {

	var triggerScriptId *string
	if p.TriggerScriptId != nil {
		triggerScriptId = p.TriggerScriptId
	}
	var doneTriggerTargetType *string
	if p.DoneTriggerTargetType != nil {
		doneTriggerTargetType = p.DoneTriggerTargetType
	}
	var doneTriggerScriptId *string
	if p.DoneTriggerScriptId != nil {
		doneTriggerScriptId = p.DoneTriggerScriptId
	}
	var doneTriggerQueueNamespaceId *string
	if p.DoneTriggerQueueNamespaceId != nil {
		doneTriggerQueueNamespaceId = p.DoneTriggerQueueNamespaceId
	}
	return map[string]interface{}{
		"triggerScriptId":             triggerScriptId,
		"doneTriggerTargetType":       doneTriggerTargetType,
		"doneTriggerScriptId":         doneTriggerScriptId,
		"doneTriggerQueueNamespaceId": doneTriggerQueueNamespaceId,
	}
}

func (p ScriptSetting) Pointer() *ScriptSetting {
	return &p
}

func CastScriptSettings(data []interface{}) []ScriptSetting {
	v := make([]ScriptSetting, 0)
	for _, d := range data {
		v = append(v, NewScriptSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScriptSettingsFromDict(data []ScriptSetting) []interface{} {
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
