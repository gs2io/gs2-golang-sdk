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
	NamespaceId      *string        `json:"namespaceId"`
	Name             *string        `json:"name"`
	Description      *string        `json:"description"`
	DoneUploadScript *ScriptSetting `json:"doneUploadScript"`
	LogSetting       *LogSetting    `json:"logSetting"`
	CreatedAt        *int64         `json:"createdAt"`
	UpdatedAt        *int64         `json:"updatedAt"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:      core.CastString(data["namespaceId"]),
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		DoneUploadScript: NewScriptSettingFromDict(core.CastMap(data["doneUploadScript"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":      p.NamespaceId,
		"name":             p.Name,
		"description":      p.Description,
		"doneUploadScript": p.DoneUploadScript.ToDict(),
		"logSetting":       p.LogSetting.ToDict(),
		"createdAt":        p.CreatedAt,
		"updatedAt":        p.UpdatedAt,
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
	TriggerScriptId             *string `json:"triggerScriptId"`
	DoneTriggerTargetType       *string `json:"doneTriggerTargetType"`
	DoneTriggerScriptId         *string `json:"doneTriggerScriptId"`
	DoneTriggerQueueNamespaceId *string `json:"doneTriggerQueueNamespaceId"`
}

func NewScriptSettingFromJson(data string) ScriptSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewScriptSettingFromDict(dict)
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
	return map[string]interface{}{
		"triggerScriptId":             p.TriggerScriptId,
		"doneTriggerTargetType":       p.DoneTriggerTargetType,
		"doneTriggerScriptId":         p.DoneTriggerScriptId,
		"doneTriggerQueueNamespaceId": p.DoneTriggerQueueNamespaceId,
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
	DataObjectId       *string  `json:"dataObjectId"`
	Name               *string  `json:"name"`
	UserId             *string  `json:"userId"`
	Scope              *string  `json:"scope"`
	AllowUserIds       []string `json:"allowUserIds"`
	Status             *string  `json:"status"`
	Generation         *string  `json:"generation"`
	PreviousGeneration *string  `json:"previousGeneration"`
	CreatedAt          *int64   `json:"createdAt"`
	UpdatedAt          *int64   `json:"updatedAt"`
}

func NewDataObjectFromJson(data string) DataObject {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDataObjectFromDict(dict)
}

func NewDataObjectFromDict(data map[string]interface{}) DataObject {
	return DataObject{
		DataObjectId:       core.CastString(data["dataObjectId"]),
		Name:               core.CastString(data["name"]),
		UserId:             core.CastString(data["userId"]),
		Scope:              core.CastString(data["scope"]),
		AllowUserIds:       core.CastStrings(core.CastArray(data["allowUserIds"])),
		Status:             core.CastString(data["status"]),
		Generation:         core.CastString(data["generation"]),
		PreviousGeneration: core.CastString(data["previousGeneration"]),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
	}
}

func (p DataObject) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"dataObjectId": p.DataObjectId,
		"name":         p.Name,
		"userId":       p.UserId,
		"scope":        p.Scope,
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"status":             p.Status,
		"generation":         p.Generation,
		"previousGeneration": p.PreviousGeneration,
		"createdAt":          p.CreatedAt,
		"updatedAt":          p.UpdatedAt,
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
	DataObjectName      *string `json:"dataObjectName"`
	Generation          *string `json:"generation"`
	ContentLength       *int64  `json:"contentLength"`
	CreatedAt           *int64  `json:"createdAt"`
}

func NewDataObjectHistoryFromJson(data string) DataObjectHistory {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDataObjectHistoryFromDict(dict)
}

func NewDataObjectHistoryFromDict(data map[string]interface{}) DataObjectHistory {
	return DataObjectHistory{
		DataObjectHistoryId: core.CastString(data["dataObjectHistoryId"]),
		DataObjectName:      core.CastString(data["dataObjectName"]),
		Generation:          core.CastString(data["generation"]),
		ContentLength:       core.CastInt64(data["contentLength"]),
		CreatedAt:           core.CastInt64(data["createdAt"]),
	}
}

func (p DataObjectHistory) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"dataObjectHistoryId": p.DataObjectHistoryId,
		"dataObjectName":      p.DataObjectName,
		"generation":          p.Generation,
		"contentLength":       p.ContentLength,
		"createdAt":           p.CreatedAt,
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
	return LogSetting{
		LoggingNamespaceId: core.CastString(data["loggingNamespaceId"]),
	}
}

func (p LogSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"loggingNamespaceId": p.LoggingNamespaceId,
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
