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

package datastore

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	DoneUploadScript *ScriptSetting `json:"doneUploadScript"`
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
        DoneUploadScript: NewScriptSettingFromDict(core.CastMap(data["doneUploadScript"])).Pointer(),
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
    var doneUploadScript map[string]interface{}
    if p.DoneUploadScript != nil {
        doneUploadScript = p.DoneUploadScript.ToDict()
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
        "doneUploadScript": doneUploadScript,
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

type DataObject struct {
	DataObjectId *string `json:"dataObjectId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	Scope *string `json:"scope"`
	AllowUserIds []*string `json:"allowUserIds"`
	Status *string `json:"status"`
	Generation *string `json:"generation"`
	PreviousGeneration *string `json:"previousGeneration"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewDataObjectFromJson(data string) DataObject {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDataObjectFromDict(dict)
}

func NewDataObjectFromDict(data map[string]interface{}) DataObject {
    return DataObject {
        DataObjectId: core.CastString(data["dataObjectId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        Scope: core.CastString(data["scope"]),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        Status: core.CastString(data["status"]),
        Generation: core.CastString(data["generation"]),
        PreviousGeneration: core.CastString(data["previousGeneration"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p DataObject) ToDict() map[string]interface{} {
    
    var dataObjectId *string
    if p.DataObjectId != nil {
        dataObjectId = p.DataObjectId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var scope *string
    if p.Scope != nil {
        scope = p.Scope
    }
    var allowUserIds []interface{}
    if p.AllowUserIds != nil {
        allowUserIds = core.CastStringsFromDict(
            p.AllowUserIds,
        )
    }
    var status *string
    if p.Status != nil {
        status = p.Status
    }
    var generation *string
    if p.Generation != nil {
        generation = p.Generation
    }
    var previousGeneration *string
    if p.PreviousGeneration != nil {
        previousGeneration = p.PreviousGeneration
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
        "dataObjectId": dataObjectId,
        "name": name,
        "userId": userId,
        "scope": scope,
        "allowUserIds": allowUserIds,
        "status": status,
        "generation": generation,
        "previousGeneration": previousGeneration,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p DataObject) Pointer() *DataObject {
    return &p
}

func CastDataObjects(data []interface{}) []DataObject {
	v := make([]DataObject, 0)
	for _, d := range data {
		v = append(v, NewDataObjectFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDataObjectsFromDict(data []DataObject) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type DataObjectHistory struct {
	DataObjectHistoryId *string `json:"dataObjectHistoryId"`
	DataObjectName *string `json:"dataObjectName"`
	Generation *string `json:"generation"`
	ContentLength *int64 `json:"contentLength"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewDataObjectHistoryFromJson(data string) DataObjectHistory {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDataObjectHistoryFromDict(dict)
}

func NewDataObjectHistoryFromDict(data map[string]interface{}) DataObjectHistory {
    return DataObjectHistory {
        DataObjectHistoryId: core.CastString(data["dataObjectHistoryId"]),
        DataObjectName: core.CastString(data["dataObjectName"]),
        Generation: core.CastString(data["generation"]),
        ContentLength: core.CastInt64(data["contentLength"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p DataObjectHistory) ToDict() map[string]interface{} {
    
    var dataObjectHistoryId *string
    if p.DataObjectHistoryId != nil {
        dataObjectHistoryId = p.DataObjectHistoryId
    }
    var dataObjectName *string
    if p.DataObjectName != nil {
        dataObjectName = p.DataObjectName
    }
    var generation *string
    if p.Generation != nil {
        generation = p.Generation
    }
    var contentLength *int64
    if p.ContentLength != nil {
        contentLength = p.ContentLength
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "dataObjectHistoryId": dataObjectHistoryId,
        "dataObjectName": dataObjectName,
        "generation": generation,
        "contentLength": contentLength,
        "createdAt": createdAt,
    }
}

func (p DataObjectHistory) Pointer() *DataObjectHistory {
    return &p
}

func CastDataObjectHistories(data []interface{}) []DataObjectHistory {
	v := make([]DataObjectHistory, 0)
	for _, d := range data {
		v = append(v, NewDataObjectHistoryFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastDataObjectHistoriesFromDict(data []DataObjectHistory) []interface{} {
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