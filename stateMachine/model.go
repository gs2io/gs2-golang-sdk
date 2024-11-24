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

package stateMachine

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                 *string             `json:"namespaceId"`
	Name                        *string             `json:"name"`
	Description                 *string             `json:"description"`
	SupportSpeculativeExecution *string             `json:"supportSpeculativeExecution"`
	TransactionSetting          *TransactionSetting `json:"transactionSetting"`
	StartScript                 *ScriptSetting      `json:"startScript"`
	PassScript                  *ScriptSetting      `json:"passScript"`
	ErrorScript                 *ScriptSetting      `json:"errorScript"`
	LowestStateMachineVersion   *int64              `json:"lowestStateMachineVersion"`
	LogSetting                  *LogSetting         `json:"logSetting"`
	CreatedAt                   *int64              `json:"createdAt"`
	UpdatedAt                   *int64              `json:"updatedAt"`
	Revision                    *int64              `json:"revision"`
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
		if v, ok := d["supportSpeculativeExecution"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SupportSpeculativeExecution = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SupportSpeculativeExecution = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SupportSpeculativeExecution = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SupportSpeculativeExecution = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SupportSpeculativeExecution = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SupportSpeculativeExecution)
				}
			}
		}
		if v, ok := d["transactionSetting"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransactionSetting)
		}
		if v, ok := d["startScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StartScript)
		}
		if v, ok := d["passScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.PassScript)
		}
		if v, ok := d["errorScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorScript)
		}
		if v, ok := d["lowestStateMachineVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LowestStateMachineVersion)
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
		SupportSpeculativeExecution: func() *string {
			v, ok := data["supportSpeculativeExecution"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["supportSpeculativeExecution"])
		}(),
		TransactionSetting: func() *TransactionSetting {
			v, ok := data["transactionSetting"]
			if !ok || v == nil {
				return nil
			}
			return NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer()
		}(),
		StartScript: func() *ScriptSetting {
			v, ok := data["startScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer()
		}(),
		PassScript: func() *ScriptSetting {
			v, ok := data["passScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer()
		}(),
		ErrorScript: func() *ScriptSetting {
			v, ok := data["errorScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer()
		}(),
		LowestStateMachineVersion: func() *int64 {
			v, ok := data["lowestStateMachineVersion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["lowestStateMachineVersion"])
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
	if p.SupportSpeculativeExecution != nil {
		m["supportSpeculativeExecution"] = p.SupportSpeculativeExecution
	}
	if p.TransactionSetting != nil {
		m["transactionSetting"] = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	if p.StartScript != nil {
		m["startScript"] = func() map[string]interface{} {
			if p.StartScript == nil {
				return nil
			}
			return p.StartScript.ToDict()
		}()
	}
	if p.PassScript != nil {
		m["passScript"] = func() map[string]interface{} {
			if p.PassScript == nil {
				return nil
			}
			return p.PassScript.ToDict()
		}()
	}
	if p.ErrorScript != nil {
		m["errorScript"] = func() map[string]interface{} {
			if p.ErrorScript == nil {
				return nil
			}
			return p.ErrorScript.ToDict()
		}()
	}
	if p.LowestStateMachineVersion != nil {
		m["lowestStateMachineVersion"] = p.LowestStateMachineVersion
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

type StateMachineMaster struct {
	StateMachineId       *string `json:"stateMachineId"`
	MainStateMachineName *string `json:"mainStateMachineName"`
	Payload              *string `json:"payload"`
	Version              *int64  `json:"version"`
	CreatedAt            *int64  `json:"createdAt"`
	UpdatedAt            *int64  `json:"updatedAt"`
	Revision             *int64  `json:"revision"`
}

func (p *StateMachineMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StateMachineMaster{}
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
		*p = StateMachineMaster{}
	} else {
		*p = StateMachineMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stateMachineId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StateMachineId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StateMachineId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StateMachineId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StateMachineId)
				}
			}
		}
		if v, ok := d["mainStateMachineName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.MainStateMachineName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.MainStateMachineName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.MainStateMachineName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.MainStateMachineName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.MainStateMachineName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.MainStateMachineName)
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
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
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

func NewStateMachineMasterFromJson(data string) StateMachineMaster {
	req := StateMachineMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStateMachineMasterFromDict(data map[string]interface{}) StateMachineMaster {
	return StateMachineMaster{
		StateMachineId: func() *string {
			v, ok := data["stateMachineId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stateMachineId"])
		}(),
		MainStateMachineName: func() *string {
			v, ok := data["mainStateMachineName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["mainStateMachineName"])
		}(),
		Payload: func() *string {
			v, ok := data["payload"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["payload"])
		}(),
		Version: func() *int64 {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["version"])
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

func (p StateMachineMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StateMachineId != nil {
		m["stateMachineId"] = p.StateMachineId
	}
	if p.MainStateMachineName != nil {
		m["mainStateMachineName"] = p.MainStateMachineName
	}
	if p.Payload != nil {
		m["payload"] = p.Payload
	}
	if p.Version != nil {
		m["version"] = p.Version
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

func (p StateMachineMaster) Pointer() *StateMachineMaster {
	return &p
}

func CastStateMachineMasters(data []interface{}) []StateMachineMaster {
	v := make([]StateMachineMaster, 0)
	for _, d := range data {
		v = append(v, NewStateMachineMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStateMachineMastersFromDict(data []StateMachineMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	StatusId                   *string       `json:"statusId"`
	UserId                     *string       `json:"userId"`
	Name                       *string       `json:"name"`
	StateMachineVersion        *int64        `json:"stateMachineVersion"`
	EnableSpeculativeExecution *string       `json:"enableSpeculativeExecution"`
	StateMachineDefinition     *string       `json:"stateMachineDefinition"`
	RandomStatus               *RandomStatus `json:"randomStatus"`
	Stacks                     []StackEntry  `json:"stacks"`
	Variables                  []Variable    `json:"variables"`
	Status                     *string       `json:"status"`
	LastError                  *string       `json:"lastError"`
	TransitionCount            *int32        `json:"transitionCount"`
	CreatedAt                  *int64        `json:"createdAt"`
	UpdatedAt                  *int64        `json:"updatedAt"`
}

func (p *Status) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Status{}
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
		*p = Status{}
	} else {
		*p = Status{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["statusId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StatusId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StatusId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StatusId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StatusId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StatusId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StatusId)
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
		if v, ok := d["stateMachineVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StateMachineVersion)
		}
		if v, ok := d["enableSpeculativeExecution"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EnableSpeculativeExecution = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EnableSpeculativeExecution = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EnableSpeculativeExecution = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EnableSpeculativeExecution = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EnableSpeculativeExecution = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EnableSpeculativeExecution)
				}
			}
		}
		if v, ok := d["stateMachineDefinition"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StateMachineDefinition = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StateMachineDefinition = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StateMachineDefinition = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineDefinition = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineDefinition = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StateMachineDefinition)
				}
			}
		}
		if v, ok := d["randomStatus"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.RandomStatus)
		}
		if v, ok := d["stacks"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Stacks)
		}
		if v, ok := d["variables"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Variables)
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
		if v, ok := d["lastError"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LastError = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LastError = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LastError = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LastError = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LastError = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LastError)
				}
			}
		}
		if v, ok := d["transitionCount"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.TransitionCount)
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

func NewStatusFromJson(data string) Status {
	req := Status{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStatusFromDict(data map[string]interface{}) Status {
	return Status{
		StatusId: func() *string {
			v, ok := data["statusId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["statusId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Name: func() *string {
			v, ok := data["name"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["name"])
		}(),
		StateMachineVersion: func() *int64 {
			v, ok := data["stateMachineVersion"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["stateMachineVersion"])
		}(),
		EnableSpeculativeExecution: func() *string {
			v, ok := data["enableSpeculativeExecution"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["enableSpeculativeExecution"])
		}(),
		StateMachineDefinition: func() *string {
			v, ok := data["stateMachineDefinition"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stateMachineDefinition"])
		}(),
		RandomStatus: func() *RandomStatus {
			v, ok := data["randomStatus"]
			if !ok || v == nil {
				return nil
			}
			return NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer()
		}(),
		Stacks: func() []StackEntry {
			if data["stacks"] == nil {
				return nil
			}
			return CastStackEntries(core.CastArray(data["stacks"]))
		}(),
		Variables: func() []Variable {
			if data["variables"] == nil {
				return nil
			}
			return CastVariables(core.CastArray(data["variables"]))
		}(),
		Status: func() *string {
			v, ok := data["status"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["status"])
		}(),
		LastError: func() *string {
			v, ok := data["lastError"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["lastError"])
		}(),
		TransitionCount: func() *int32 {
			v, ok := data["transitionCount"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["transitionCount"])
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

func (p Status) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StatusId != nil {
		m["statusId"] = p.StatusId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.StateMachineVersion != nil {
		m["stateMachineVersion"] = p.StateMachineVersion
	}
	if p.EnableSpeculativeExecution != nil {
		m["enableSpeculativeExecution"] = p.EnableSpeculativeExecution
	}
	if p.StateMachineDefinition != nil {
		m["stateMachineDefinition"] = p.StateMachineDefinition
	}
	if p.RandomStatus != nil {
		m["randomStatus"] = func() map[string]interface{} {
			if p.RandomStatus == nil {
				return nil
			}
			return p.RandomStatus.ToDict()
		}()
	}
	if p.Stacks != nil {
		m["stacks"] = CastStackEntriesFromDict(
			p.Stacks,
		)
	}
	if p.Variables != nil {
		m["variables"] = CastVariablesFromDict(
			p.Variables,
		)
	}
	if p.Status != nil {
		m["status"] = p.Status
	}
	if p.LastError != nil {
		m["lastError"] = p.LastError
	}
	if p.TransitionCount != nil {
		m["transitionCount"] = p.TransitionCount
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	if p.UpdatedAt != nil {
		m["updatedAt"] = p.UpdatedAt
	}
	return m
}

func (p Status) Pointer() *Status {
	return &p
}

func CastStatuses(data []interface{}) []Status {
	v := make([]Status, 0)
	for _, d := range data {
		v = append(v, NewStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStatusesFromDict(data []Status) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type StackEntry struct {
	StateMachineName *string `json:"stateMachineName"`
	TaskName         *string `json:"taskName"`
}

func (p *StackEntry) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = StackEntry{}
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
		*p = StackEntry{}
	} else {
		*p = StackEntry{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stateMachineName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StateMachineName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StateMachineName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StateMachineName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StateMachineName)
				}
			}
		}
		if v, ok := d["taskName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TaskName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TaskName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TaskName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TaskName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TaskName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TaskName)
				}
			}
		}
	}
	return nil
}

func NewStackEntryFromJson(data string) StackEntry {
	req := StackEntry{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewStackEntryFromDict(data map[string]interface{}) StackEntry {
	return StackEntry{
		StateMachineName: func() *string {
			v, ok := data["stateMachineName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stateMachineName"])
		}(),
		TaskName: func() *string {
			v, ok := data["taskName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["taskName"])
		}(),
	}
}

func (p StackEntry) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StateMachineName != nil {
		m["stateMachineName"] = p.StateMachineName
	}
	if p.TaskName != nil {
		m["taskName"] = p.TaskName
	}
	return m
}

func (p StackEntry) Pointer() *StackEntry {
	return &p
}

func CastStackEntries(data []interface{}) []StackEntry {
	v := make([]StackEntry, 0)
	for _, d := range data {
		v = append(v, NewStackEntryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStackEntriesFromDict(data []StackEntry) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Variable struct {
	StateMachineName *string `json:"stateMachineName"`
	Value            *string `json:"value"`
}

func (p *Variable) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Variable{}
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
		*p = Variable{}
	} else {
		*p = Variable{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["stateMachineName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.StateMachineName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.StateMachineName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.StateMachineName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.StateMachineName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.StateMachineName)
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

func NewVariableFromJson(data string) Variable {
	req := Variable{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVariableFromDict(data map[string]interface{}) Variable {
	return Variable{
		StateMachineName: func() *string {
			v, ok := data["stateMachineName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["stateMachineName"])
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

func (p Variable) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.StateMachineName != nil {
		m["stateMachineName"] = p.StateMachineName
	}
	if p.Value != nil {
		m["value"] = p.Value
	}
	return m
}

func (p Variable) Pointer() *Variable {
	return &p
}

func CastVariables(data []interface{}) []Variable {
	v := make([]Variable, 0)
	for _, d := range data {
		v = append(v, NewVariableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVariablesFromDict(data []Variable) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionSetting struct {
	EnableAutoRun             *bool   `json:"enableAutoRun"`
	EnableAtomicCommit        *bool   `json:"enableAtomicCommit"`
	TransactionUseDistributor *bool   `json:"transactionUseDistributor"`
	AcquireActionUseJobQueue  *bool   `json:"acquireActionUseJobQueue"`
	DistributorNamespaceId    *string `json:"distributorNamespaceId"`
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
		if v, ok := d["acquireActionUseJobQueue"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireActionUseJobQueue)
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
		AcquireActionUseJobQueue: func() *bool {
			v, ok := data["acquireActionUseJobQueue"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["acquireActionUseJobQueue"])
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
	if p.AcquireActionUseJobQueue != nil {
		m["acquireActionUseJobQueue"] = p.AcquireActionUseJobQueue
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

type Event struct {
	EventType        *string           `json:"eventType"`
	ChangeStateEvent *ChangeStateEvent `json:"changeStateEvent"`
	EmitEvent        *EmitEvent        `json:"emitEvent"`
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
		if v, ok := d["eventType"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventType = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventType = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventType = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventType = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventType = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventType)
				}
			}
		}
		if v, ok := d["changeStateEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangeStateEvent)
		}
		if v, ok := d["emitEvent"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EmitEvent)
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
		EventType: func() *string {
			v, ok := data["eventType"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["eventType"])
		}(),
		ChangeStateEvent: func() *ChangeStateEvent {
			v, ok := data["changeStateEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewChangeStateEventFromDict(core.CastMap(data["changeStateEvent"])).Pointer()
		}(),
		EmitEvent: func() *EmitEvent {
			v, ok := data["emitEvent"]
			if !ok || v == nil {
				return nil
			}
			return NewEmitEventFromDict(core.CastMap(data["emitEvent"])).Pointer()
		}(),
	}
}

func (p Event) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.EventType != nil {
		m["eventType"] = p.EventType
	}
	if p.ChangeStateEvent != nil {
		m["changeStateEvent"] = func() map[string]interface{} {
			if p.ChangeStateEvent == nil {
				return nil
			}
			return p.ChangeStateEvent.ToDict()
		}()
	}
	if p.EmitEvent != nil {
		m["emitEvent"] = func() map[string]interface{} {
			if p.EmitEvent == nil {
				return nil
			}
			return p.EmitEvent.ToDict()
		}()
	}
	return m
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

type ChangeStateEvent struct {
	TaskName  *string `json:"taskName"`
	Hash      *string `json:"hash"`
	Timestamp *int64  `json:"timestamp"`
}

func (p *ChangeStateEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ChangeStateEvent{}
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
		*p = ChangeStateEvent{}
	} else {
		*p = ChangeStateEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["taskName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.TaskName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.TaskName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.TaskName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.TaskName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.TaskName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.TaskName)
				}
			}
		}
		if v, ok := d["hash"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Hash = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Hash = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Hash = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Hash = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Hash = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Hash)
				}
			}
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
	}
	return nil
}

func NewChangeStateEventFromJson(data string) ChangeStateEvent {
	req := ChangeStateEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewChangeStateEventFromDict(data map[string]interface{}) ChangeStateEvent {
	return ChangeStateEvent{
		TaskName: func() *string {
			v, ok := data["taskName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["taskName"])
		}(),
		Hash: func() *string {
			v, ok := data["hash"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["hash"])
		}(),
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
	}
}

func (p ChangeStateEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TaskName != nil {
		m["taskName"] = p.TaskName
	}
	if p.Hash != nil {
		m["hash"] = p.Hash
	}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	return m
}

func (p ChangeStateEvent) Pointer() *ChangeStateEvent {
	return &p
}

func CastChangeStateEvents(data []interface{}) []ChangeStateEvent {
	v := make([]ChangeStateEvent, 0)
	for _, d := range data {
		v = append(v, NewChangeStateEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastChangeStateEventsFromDict(data []ChangeStateEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type EmitEvent struct {
	Event      *string `json:"event"`
	Parameters *string `json:"parameters"`
	Timestamp  *int64  `json:"timestamp"`
}

func (p *EmitEvent) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = EmitEvent{}
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
		*p = EmitEvent{}
	} else {
		*p = EmitEvent{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["event"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Event = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Event = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Event = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Event = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Event = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Event)
				}
			}
		}
		if v, ok := d["parameters"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Parameters = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Parameters = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Parameters = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Parameters = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Parameters = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Parameters)
				}
			}
		}
		if v, ok := d["timestamp"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Timestamp)
		}
	}
	return nil
}

func NewEmitEventFromJson(data string) EmitEvent {
	req := EmitEvent{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewEmitEventFromDict(data map[string]interface{}) EmitEvent {
	return EmitEvent{
		Event: func() *string {
			v, ok := data["event"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["event"])
		}(),
		Parameters: func() *string {
			v, ok := data["parameters"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["parameters"])
		}(),
		Timestamp: func() *int64 {
			v, ok := data["timestamp"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["timestamp"])
		}(),
	}
}

func (p EmitEvent) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Event != nil {
		m["event"] = p.Event
	}
	if p.Parameters != nil {
		m["parameters"] = p.Parameters
	}
	if p.Timestamp != nil {
		m["timestamp"] = p.Timestamp
	}
	return m
}

func (p EmitEvent) Pointer() *EmitEvent {
	return &p
}

func CastEmitEvents(data []interface{}) []EmitEvent {
	v := make([]EmitEvent, 0)
	for _, d := range data {
		v = append(v, NewEmitEventFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastEmitEventsFromDict(data []EmitEvent) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomStatus struct {
	Seed *int64       `json:"seed"`
	Used []RandomUsed `json:"used"`
}

func (p *RandomStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomStatus{}
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
		*p = RandomStatus{}
	} else {
		*p = RandomStatus{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["seed"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Seed)
		}
		if v, ok := d["used"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Used)
		}
	}
	return nil
}

func NewRandomStatusFromJson(data string) RandomStatus {
	req := RandomStatus{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomStatusFromDict(data map[string]interface{}) RandomStatus {
	return RandomStatus{
		Seed: func() *int64 {
			v, ok := data["seed"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["seed"])
		}(),
		Used: func() []RandomUsed {
			if data["used"] == nil {
				return nil
			}
			return CastRandomUseds(core.CastArray(data["used"]))
		}(),
	}
}

func (p RandomStatus) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Seed != nil {
		m["seed"] = p.Seed
	}
	if p.Used != nil {
		m["used"] = CastRandomUsedsFromDict(
			p.Used,
		)
	}
	return m
}

func (p RandomStatus) Pointer() *RandomStatus {
	return &p
}

func CastRandomStatuses(data []interface{}) []RandomStatus {
	v := make([]RandomStatus, 0)
	for _, d := range data {
		v = append(v, NewRandomStatusFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomStatusesFromDict(data []RandomStatus) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type RandomUsed struct {
	Category *int64 `json:"category"`
	Used     *int64 `json:"used"`
}

func (p *RandomUsed) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = RandomUsed{}
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
		*p = RandomUsed{}
	} else {
		*p = RandomUsed{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["category"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Category)
		}
		if v, ok := d["used"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Used)
		}
	}
	return nil
}

func NewRandomUsedFromJson(data string) RandomUsed {
	req := RandomUsed{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewRandomUsedFromDict(data map[string]interface{}) RandomUsed {
	return RandomUsed{
		Category: func() *int64 {
			v, ok := data["category"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["category"])
		}(),
		Used: func() *int64 {
			v, ok := data["used"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["used"])
		}(),
	}
}

func (p RandomUsed) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Category != nil {
		m["category"] = p.Category
	}
	if p.Used != nil {
		m["used"] = p.Used
	}
	return m
}

func (p RandomUsed) Pointer() *RandomUsed {
	return &p
}

func CastRandomUseds(data []interface{}) []RandomUsed {
	v := make([]RandomUsed, 0)
	for _, d := range data {
		v = append(v, NewRandomUsedFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRandomUsedsFromDict(data []RandomUsed) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VerifyActionResult struct {
	Action        *string `json:"action"`
	VerifyRequest *string `json:"verifyRequest"`
	StatusCode    *int32  `json:"statusCode"`
	VerifyResult  *string `json:"verifyResult"`
}

func (p *VerifyActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VerifyActionResult{}
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
		*p = VerifyActionResult{}
	} else {
		*p = VerifyActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["verifyRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["verifyResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VerifyResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VerifyResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VerifyResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VerifyResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VerifyResult)
				}
			}
		}
	}
	return nil
}

func NewVerifyActionResultFromJson(data string) VerifyActionResult {
	req := VerifyActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVerifyActionResultFromDict(data map[string]interface{}) VerifyActionResult {
	return VerifyActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		VerifyRequest: func() *string {
			v, ok := data["verifyRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		VerifyResult: func() *string {
			v, ok := data["verifyResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["verifyResult"])
		}(),
	}
}

func (p VerifyActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.VerifyRequest != nil {
		m["verifyRequest"] = p.VerifyRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.VerifyResult != nil {
		m["verifyResult"] = p.VerifyResult
	}
	return m
}

func (p VerifyActionResult) Pointer() *VerifyActionResult {
	return &p
}

func CastVerifyActionResults(data []interface{}) []VerifyActionResult {
	v := make([]VerifyActionResult, 0)
	for _, d := range data {
		v = append(v, NewVerifyActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVerifyActionResultsFromDict(data []VerifyActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ConsumeActionResult struct {
	Action         *string `json:"action"`
	ConsumeRequest *string `json:"consumeRequest"`
	StatusCode     *int32  `json:"statusCode"`
	ConsumeResult  *string `json:"consumeResult"`
}

func (p *ConsumeActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ConsumeActionResult{}
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
		*p = ConsumeActionResult{}
	} else {
		*p = ConsumeActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["consumeRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["consumeResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ConsumeResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ConsumeResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ConsumeResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ConsumeResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ConsumeResult)
				}
			}
		}
	}
	return nil
}

func NewConsumeActionResultFromJson(data string) ConsumeActionResult {
	req := ConsumeActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewConsumeActionResultFromDict(data map[string]interface{}) ConsumeActionResult {
	return ConsumeActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		ConsumeRequest: func() *string {
			v, ok := data["consumeRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		ConsumeResult: func() *string {
			v, ok := data["consumeResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["consumeResult"])
		}(),
	}
}

func (p ConsumeActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.ConsumeRequest != nil {
		m["consumeRequest"] = p.ConsumeRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.ConsumeResult != nil {
		m["consumeResult"] = p.ConsumeResult
	}
	return m
}

func (p ConsumeActionResult) Pointer() *ConsumeActionResult {
	return &p
}

func CastConsumeActionResults(data []interface{}) []ConsumeActionResult {
	v := make([]ConsumeActionResult, 0)
	for _, d := range data {
		v = append(v, NewConsumeActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastConsumeActionResultsFromDict(data []ConsumeActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcquireActionResult struct {
	Action         *string `json:"action"`
	AcquireRequest *string `json:"acquireRequest"`
	StatusCode     *int32  `json:"statusCode"`
	AcquireResult  *string `json:"acquireResult"`
}

func (p *AcquireActionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcquireActionResult{}
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
		*p = AcquireActionResult{}
	} else {
		*p = AcquireActionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["acquireRequest"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireRequest = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireRequest = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireRequest = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireRequest = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireRequest)
				}
			}
		}
		if v, ok := d["statusCode"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.StatusCode)
		}
		if v, ok := d["acquireResult"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcquireResult = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcquireResult = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcquireResult = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcquireResult = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcquireResult)
				}
			}
		}
	}
	return nil
}

func NewAcquireActionResultFromJson(data string) AcquireActionResult {
	req := AcquireActionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcquireActionResultFromDict(data map[string]interface{}) AcquireActionResult {
	return AcquireActionResult{
		Action: func() *string {
			v, ok := data["action"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["action"])
		}(),
		AcquireRequest: func() *string {
			v, ok := data["acquireRequest"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireRequest"])
		}(),
		StatusCode: func() *int32 {
			v, ok := data["statusCode"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["statusCode"])
		}(),
		AcquireResult: func() *string {
			v, ok := data["acquireResult"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acquireResult"])
		}(),
	}
}

func (p AcquireActionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Action != nil {
		m["action"] = p.Action
	}
	if p.AcquireRequest != nil {
		m["acquireRequest"] = p.AcquireRequest
	}
	if p.StatusCode != nil {
		m["statusCode"] = p.StatusCode
	}
	if p.AcquireResult != nil {
		m["acquireResult"] = p.AcquireResult
	}
	return m
}

func (p AcquireActionResult) Pointer() *AcquireActionResult {
	return &p
}

func CastAcquireActionResults(data []interface{}) []AcquireActionResult {
	v := make([]AcquireActionResult, 0)
	for _, d := range data {
		v = append(v, NewAcquireActionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcquireActionResultsFromDict(data []AcquireActionResult) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type TransactionResult struct {
	TransactionId  *string               `json:"transactionId"`
	VerifyResults  []VerifyActionResult  `json:"verifyResults"`
	ConsumeResults []ConsumeActionResult `json:"consumeResults"`
	AcquireResults []AcquireActionResult `json:"acquireResults"`
}

func (p *TransactionResult) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TransactionResult{}
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
		*p = TransactionResult{}
	} else {
		*p = TransactionResult{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
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
		if v, ok := d["verifyResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VerifyResults)
		}
		if v, ok := d["consumeResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumeResults)
		}
		if v, ok := d["acquireResults"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquireResults)
		}
	}
	return nil
}

func NewTransactionResultFromJson(data string) TransactionResult {
	req := TransactionResult{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTransactionResultFromDict(data map[string]interface{}) TransactionResult {
	return TransactionResult{
		TransactionId: func() *string {
			v, ok := data["transactionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["transactionId"])
		}(),
		VerifyResults: func() []VerifyActionResult {
			if data["verifyResults"] == nil {
				return nil
			}
			return CastVerifyActionResults(core.CastArray(data["verifyResults"]))
		}(),
		ConsumeResults: func() []ConsumeActionResult {
			if data["consumeResults"] == nil {
				return nil
			}
			return CastConsumeActionResults(core.CastArray(data["consumeResults"]))
		}(),
		AcquireResults: func() []AcquireActionResult {
			if data["acquireResults"] == nil {
				return nil
			}
			return CastAcquireActionResults(core.CastArray(data["acquireResults"]))
		}(),
	}
}

func (p TransactionResult) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.TransactionId != nil {
		m["transactionId"] = p.TransactionId
	}
	if p.VerifyResults != nil {
		m["verifyResults"] = CastVerifyActionResultsFromDict(
			p.VerifyResults,
		)
	}
	if p.ConsumeResults != nil {
		m["consumeResults"] = CastConsumeActionResultsFromDict(
			p.ConsumeResults,
		)
	}
	if p.AcquireResults != nil {
		m["acquireResults"] = CastAcquireActionResultsFromDict(
			p.AcquireResults,
		)
	}
	return m
}

func (p TransactionResult) Pointer() *TransactionResult {
	return &p
}

func CastTransactionResults(data []interface{}) []TransactionResult {
	v := make([]TransactionResult, 0)
	for _, d := range data {
		v = append(v, NewTransactionResultFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTransactionResultsFromDict(data []TransactionResult) []interface{} {
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
	m := map[string]interface{}{}
	if p.TriggerScriptId != nil {
		m["triggerScriptId"] = p.TriggerScriptId
	}
	if p.DoneTriggerTargetType != nil {
		m["doneTriggerTargetType"] = p.DoneTriggerTargetType
	}
	if p.DoneTriggerScriptId != nil {
		m["doneTriggerScriptId"] = p.DoneTriggerScriptId
	}
	if p.DoneTriggerQueueNamespaceId != nil {
		m["doneTriggerQueueNamespaceId"] = p.DoneTriggerQueueNamespaceId
	}
	return m
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
