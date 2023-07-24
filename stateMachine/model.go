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
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	StartScript *ScriptSetting `json:"startScript"`
	PassScript *ScriptSetting `json:"passScript"`
	ErrorScript *ScriptSetting `json:"errorScript"`
	LowestStateMachineVersion *int64 `json:"lowestStateMachineVersion"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
    return Namespace {
        NamespaceId: core.CastString(data["namespaceId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        StartScript: NewScriptSettingFromDict(core.CastMap(data["startScript"])).Pointer(),
        PassScript: NewScriptSettingFromDict(core.CastMap(data["passScript"])).Pointer(),
        ErrorScript: NewScriptSettingFromDict(core.CastMap(data["errorScript"])).Pointer(),
        LowestStateMachineVersion: core.CastInt64(data["lowestStateMachineVersion"]),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    var startScript map[string]interface{}
    if p.StartScript != nil {
        startScript = p.StartScript.ToDict()
    }
    var passScript map[string]interface{}
    if p.PassScript != nil {
        passScript = p.PassScript.ToDict()
    }
    var errorScript map[string]interface{}
    if p.ErrorScript != nil {
        errorScript = p.ErrorScript.ToDict()
    }
    var lowestStateMachineVersion *int64
    if p.LowestStateMachineVersion != nil {
        lowestStateMachineVersion = p.LowestStateMachineVersion
    }
    var logSetting map[string]interface{}
    if p.LogSetting != nil {
        logSetting = p.LogSetting.ToDict()
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "startScript": startScript,
        "passScript": passScript,
        "errorScript": errorScript,
        "lowestStateMachineVersion": lowestStateMachineVersion,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	StateMachineId *string `json:"stateMachineId"`
	MainStateMachineName *string `json:"mainStateMachineName"`
	Payload *string `json:"payload"`
	Version *int64 `json:"version"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewStateMachineMasterFromJson(data string) StateMachineMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStateMachineMasterFromDict(dict)
}

func NewStateMachineMasterFromDict(data map[string]interface{}) StateMachineMaster {
    return StateMachineMaster {
        StateMachineId: core.CastString(data["stateMachineId"]),
        MainStateMachineName: core.CastString(data["mainStateMachineName"]),
        Payload: core.CastString(data["payload"]),
        Version: core.CastInt64(data["version"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    return map[string]interface{} {
        "stateMachineId": stateMachineId,
        "mainStateMachineName": mainStateMachineName,
        "payload": payload,
        "version": version,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	StatusId *string `json:"statusId"`
	UserId *string `json:"userId"`
	Name *string `json:"name"`
	StateMachineVersion *int64 `json:"stateMachineVersion"`
	Stacks []StackEntry `json:"stacks"`
	Variables []Variable `json:"variables"`
	Status *string `json:"status"`
	LastError *string `json:"lastError"`
	TransitionCount *int32 `json:"transitionCount"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewStatusFromJson(data string) Status {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStatusFromDict(dict)
}

func NewStatusFromDict(data map[string]interface{}) Status {
    return Status {
        StatusId: core.CastString(data["statusId"]),
        UserId: core.CastString(data["userId"]),
        Name: core.CastString(data["name"]),
        StateMachineVersion: core.CastInt64(data["stateMachineVersion"]),
        Stacks: CastStackEntries(core.CastArray(data["stacks"])),
        Variables: CastVariables(core.CastArray(data["variables"])),
        Status: core.CastString(data["status"]),
        LastError: core.CastString(data["lastError"]),
        TransitionCount: core.CastInt32(data["transitionCount"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
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
    return map[string]interface{} {
        "statusId": statusId,
        "userId": userId,
        "name": name,
        "stateMachineVersion": stateMachineVersion,
        "stacks": stacks,
        "variables": variables,
        "status": status,
        "lastError": lastError,
        "transitionCount": transitionCount,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
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
	TaskName *string `json:"taskName"`
}

func NewStackEntryFromJson(data string) StackEntry {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStackEntryFromDict(dict)
}

func NewStackEntryFromDict(data map[string]interface{}) StackEntry {
    return StackEntry {
        StateMachineName: core.CastString(data["stateMachineName"]),
        TaskName: core.CastString(data["taskName"]),
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
    return map[string]interface{} {
        "stateMachineName": stateMachineName,
        "taskName": taskName,
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
	Value *string `json:"value"`
}

func NewVariableFromJson(data string) Variable {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVariableFromDict(dict)
}

func NewVariableFromDict(data map[string]interface{}) Variable {
    return Variable {
        StateMachineName: core.CastString(data["stateMachineName"]),
        Value: core.CastString(data["value"]),
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
    return map[string]interface{} {
        "stateMachineName": stateMachineName,
        "value": value,
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

type LogSetting struct {
	LoggingNamespaceId *string `json:"loggingNamespaceId"`
}

func NewLogSettingFromJson(data string) LogSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLogSettingFromDict(dict)
}

func NewLogSettingFromDict(data map[string]interface{}) LogSetting {
    return LogSetting {
        LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
    }
}

func (p LogSetting) ToDict() map[string]interface{} {
    
    var loggingNamespaceId *string
    if p.LoggingNamespaceId != nil {
        loggingNamespaceId = p.LoggingNamespaceId
    }
    return map[string]interface{} {
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
	TriggerScriptId *string `json:"triggerScriptId"`
	DoneTriggerTargetType *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScriptSettingFromDict(dict)
}

func NewScriptSettingFromDict(data map[string]interface{}) ScriptSetting {
    return ScriptSetting {
        TriggerScriptId: core.CastString(data["triggerScriptId"]),
        DoneTriggerTargetType: core.CastString(data["doneTriggerTargetType"]),
        DoneTriggerScriptId: core.CastString(data["doneTriggerScriptId"]),
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
    return map[string]interface{} {
        "triggerScriptId": triggerScriptId,
        "doneTriggerTargetType": doneTriggerTargetType,
        "doneTriggerScriptId": doneTriggerScriptId,
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