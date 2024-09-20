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
		NamespaceId:                 core.CastString(data["namespaceId"]),
		Name:                        core.CastString(data["name"]),
		Description:                 core.CastString(data["description"]),
		SupportSpeculativeExecution: core.CastString(data["supportSpeculativeExecution"]),
		TransactionSetting:          NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		StartScript:                 NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
		PassScript:                  NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
		ErrorScript:                 NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
		LowestStateMachineVersion:   core.CastInt64(data["lowestStateMachineVersion"]),
		LogSetting:                  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:                   core.CastInt64(data["createdAt"]),
		UpdatedAt:                   core.CastInt64(data["updatedAt"]),
		Revision:                    core.CastInt64(data["revision"]),
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
	var supportSpeculativeExecution *string
	if p.SupportSpeculativeExecution != nil {
		supportSpeculativeExecution = p.SupportSpeculativeExecution
	}
	var transactionSetting map[string]interface{}
	if p.TransactionSetting != nil {
		transactionSetting = func() map[string]interface{} {
			if p.TransactionSetting == nil {
				return nil
			}
			return p.TransactionSetting.ToDict()
		}()
	}
	var startScript map[string]interface{}
	if p.StartScript != nil {
		startScript = func() map[string]interface{} {
			if p.StartScript == nil {
				return nil
			}
			return p.StartScript.ToDict()
		}()
	}
	var passScript map[string]interface{}
	if p.PassScript != nil {
		passScript = func() map[string]interface{} {
			if p.PassScript == nil {
				return nil
			}
			return p.PassScript.ToDict()
		}()
	}
	var errorScript map[string]interface{}
	if p.ErrorScript != nil {
		errorScript = func() map[string]interface{} {
			if p.ErrorScript == nil {
				return nil
			}
			return p.ErrorScript.ToDict()
		}()
	}
	var lowestStateMachineVersion *int64
	if p.LowestStateMachineVersion != nil {
		lowestStateMachineVersion = p.LowestStateMachineVersion
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
		"namespaceId":                 namespaceId,
		"name":                        name,
		"description":                 description,
		"supportSpeculativeExecution": supportSpeculativeExecution,
		"transactionSetting":          transactionSetting,
		"startScript":                 startScript,
		"passScript":                  passScript,
		"errorScript":                 errorScript,
		"lowestStateMachineVersion":   lowestStateMachineVersion,
		"logSetting":                  logSetting,
		"createdAt":                   createdAt,
		"updatedAt":                   updatedAt,
		"revision":                    revision,
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
		StateMachineId:       core.CastString(data["stateMachineId"]),
		MainStateMachineName: core.CastString(data["mainStateMachineName"]),
		Payload:              core.CastString(data["payload"]),
		Version:              core.CastInt64(data["version"]),
		CreatedAt:            core.CastInt64(data["createdAt"]),
		UpdatedAt:            core.CastInt64(data["updatedAt"]),
		Revision:             core.CastInt64(data["revision"]),
	}
}

func (p StateMachineMaster) ToDict() map[string]interface{} {

	var stateMachineId *string
	if p.StateMachineId != nil {
		stateMachineId = p.StateMachineId
	}
	var mainStateMachineName *string
	if p.MainStateMachineName != nil {
		mainStateMachineName = p.MainStateMachineName
	}
	var payload *string
	if p.Payload != nil {
		payload = p.Payload
	}
	var version *int64
	if p.Version != nil {
		version = p.Version
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
		"stateMachineId":       stateMachineId,
		"mainStateMachineName": mainStateMachineName,
		"payload":              payload,
		"version":              version,
		"createdAt":            createdAt,
		"updatedAt":            updatedAt,
		"revision":             revision,
	}
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
		StatusId:                   core.CastString(data["statusId"]),
		UserId:                     core.CastString(data["userId"]),
		Name:                       core.CastString(data["name"]),
		StateMachineVersion:        core.CastInt64(data["stateMachineVersion"]),
		EnableSpeculativeExecution: core.CastString(data["enableSpeculativeExecution"]),
		StateMachineDefinition:     core.CastString(data["stateMachineDefinition"]),
		RandomStatus:               NewRandomStatusFromDict(core.CastMap(data["randomStatus"])).Pointer(),
		Stacks:                     CastStackEntries(core.CastArray(data["stacks"])),
		Variables:                  CastVariables(core.CastArray(data["variables"])),
		Status:                     core.CastString(data["status"]),
		LastError:                  core.CastString(data["lastError"]),
		TransitionCount:            core.CastInt32(data["transitionCount"]),
		CreatedAt:                  core.CastInt64(data["createdAt"]),
		UpdatedAt:                  core.CastInt64(data["updatedAt"]),
	}
}

func (p Status) ToDict() map[string]interface{} {

	var statusId *string
	if p.StatusId != nil {
		statusId = p.StatusId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var name *string
	if p.Name != nil {
		name = p.Name
	}
	var stateMachineVersion *int64
	if p.StateMachineVersion != nil {
		stateMachineVersion = p.StateMachineVersion
	}
	var enableSpeculativeExecution *string
	if p.EnableSpeculativeExecution != nil {
		enableSpeculativeExecution = p.EnableSpeculativeExecution
	}
	var stateMachineDefinition *string
	if p.StateMachineDefinition != nil {
		stateMachineDefinition = p.StateMachineDefinition
	}
	var randomStatus map[string]interface{}
	if p.RandomStatus != nil {
		randomStatus = func() map[string]interface{} {
			if p.RandomStatus == nil {
				return nil
			}
			return p.RandomStatus.ToDict()
		}()
	}
	var stacks []interface{}
	if p.Stacks != nil {
		stacks = CastStackEntriesFromDict(
			p.Stacks,
		)
	}
	var variables []interface{}
	if p.Variables != nil {
		variables = CastVariablesFromDict(
			p.Variables,
		)
	}
	var status *string
	if p.Status != nil {
		status = p.Status
	}
	var lastError *string
	if p.LastError != nil {
		lastError = p.LastError
	}
	var transitionCount *int32
	if p.TransitionCount != nil {
		transitionCount = p.TransitionCount
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
		"statusId":                   statusId,
		"userId":                     userId,
		"name":                       name,
		"stateMachineVersion":        stateMachineVersion,
		"enableSpeculativeExecution": enableSpeculativeExecution,
		"stateMachineDefinition":     stateMachineDefinition,
		"randomStatus":               randomStatus,
		"stacks":                     stacks,
		"variables":                  variables,
		"status":                     status,
		"lastError":                  lastError,
		"transitionCount":            transitionCount,
		"createdAt":                  createdAt,
		"updatedAt":                  updatedAt,
	}
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
		StateMachineName: core.CastString(data["stateMachineName"]),
		TaskName:         core.CastString(data["taskName"]),
	}
}

func (p StackEntry) ToDict() map[string]interface{} {

	var stateMachineName *string
	if p.StateMachineName != nil {
		stateMachineName = p.StateMachineName
	}
	var taskName *string
	if p.TaskName != nil {
		taskName = p.TaskName
	}
	return map[string]interface{}{
		"stateMachineName": stateMachineName,
		"taskName":         taskName,
	}
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
		StateMachineName: core.CastString(data["stateMachineName"]),
		Value:            core.CastString(data["value"]),
	}
}

func (p Variable) ToDict() map[string]interface{} {

	var stateMachineName *string
	if p.StateMachineName != nil {
		stateMachineName = p.StateMachineName
	}
	var value *string
	if p.Value != nil {
		value = p.Value
	}
	return map[string]interface{}{
		"stateMachineName": stateMachineName,
		"value":            value,
	}
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
	EnableAutoRun          *bool   `json:"enableAutoRun"`
	DistributorNamespaceId *string `json:"distributorNamespaceId"`
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
		EnableAutoRun:          core.CastBool(data["enableAutoRun"]),
		DistributorNamespaceId: core.CastString(data["distributorNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
	}
}

func (p TransactionSetting) ToDict() map[string]interface{} {

	var enableAutoRun *bool
	if p.EnableAutoRun != nil {
		enableAutoRun = p.EnableAutoRun
	}
	var distributorNamespaceId *string
	if p.DistributorNamespaceId != nil {
		distributorNamespaceId = p.DistributorNamespaceId
	}
	var keyId *string
	if p.KeyId != nil {
		keyId = p.KeyId
	}
	var queueNamespaceId *string
	if p.QueueNamespaceId != nil {
		queueNamespaceId = p.QueueNamespaceId
	}
	return map[string]interface{}{
		"enableAutoRun":          enableAutoRun,
		"distributorNamespaceId": distributorNamespaceId,
		"keyId":                  keyId,
		"queueNamespaceId":       queueNamespaceId,
	}
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
		EventType:        core.CastString(data["eventType"]),
		ChangeStateEvent: NewChangeStateEventFromDict(core.CastMap(data["changeStateEvent"])).Pointer(),
		EmitEvent:        NewEmitEventFromDict(core.CastMap(data["emitEvent"])).Pointer(),
	}
}

func (p Event) ToDict() map[string]interface{} {

	var eventType *string
	if p.EventType != nil {
		eventType = p.EventType
	}
	var changeStateEvent map[string]interface{}
	if p.ChangeStateEvent != nil {
		changeStateEvent = func() map[string]interface{} {
			if p.ChangeStateEvent == nil {
				return nil
			}
			return p.ChangeStateEvent.ToDict()
		}()
	}
	var emitEvent map[string]interface{}
	if p.EmitEvent != nil {
		emitEvent = func() map[string]interface{} {
			if p.EmitEvent == nil {
				return nil
			}
			return p.EmitEvent.ToDict()
		}()
	}
	return map[string]interface{}{
		"eventType":        eventType,
		"changeStateEvent": changeStateEvent,
		"emitEvent":        emitEvent,
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
		TaskName:  core.CastString(data["taskName"]),
		Hash:      core.CastString(data["hash"]),
		Timestamp: core.CastInt64(data["timestamp"]),
	}
}

func (p ChangeStateEvent) ToDict() map[string]interface{} {

	var taskName *string
	if p.TaskName != nil {
		taskName = p.TaskName
	}
	var hash *string
	if p.Hash != nil {
		hash = p.Hash
	}
	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	return map[string]interface{}{
		"taskName":  taskName,
		"hash":      hash,
		"timestamp": timestamp,
	}
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
		Event:      core.CastString(data["event"]),
		Parameters: core.CastString(data["parameters"]),
		Timestamp:  core.CastInt64(data["timestamp"]),
	}
}

func (p EmitEvent) ToDict() map[string]interface{} {

	var event *string
	if p.Event != nil {
		event = p.Event
	}
	var parameters *string
	if p.Parameters != nil {
		parameters = p.Parameters
	}
	var timestamp *int64
	if p.Timestamp != nil {
		timestamp = p.Timestamp
	}
	return map[string]interface{}{
		"event":      event,
		"parameters": parameters,
		"timestamp":  timestamp,
	}
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
		Seed: core.CastInt64(data["seed"]),
		Used: CastRandomUseds(core.CastArray(data["used"])),
	}
}

func (p RandomStatus) ToDict() map[string]interface{} {

	var seed *int64
	if p.Seed != nil {
		seed = p.Seed
	}
	var used []interface{}
	if p.Used != nil {
		used = CastRandomUsedsFromDict(
			p.Used,
		)
	}
	return map[string]interface{}{
		"seed": seed,
		"used": used,
	}
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
		Category: core.CastInt64(data["category"]),
		Used:     core.CastInt64(data["used"]),
	}
}

func (p RandomUsed) ToDict() map[string]interface{} {

	var category *int64
	if p.Category != nil {
		category = p.Category
	}
	var used *int64
	if p.Used != nil {
		used = p.Used
	}
	return map[string]interface{}{
		"category": category,
		"used":     used,
	}
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
		TriggerScriptId:             core.CastString(data["triggerScriptId"]),
		DoneTriggerTargetType:       core.CastString(data["doneTriggerTargetType"]),
		DoneTriggerScriptId:         core.CastString(data["doneTriggerScriptId"]),
		DoneTriggerQueueNamespaceId: core.CastString(data["doneTriggerQueueNamespaceId"]),
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
