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

package lock

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string     `json:"namespaceId"`
	Name        *string     `json:"name"`
	Description *string     `json:"description"`
	LogSetting  *LogSetting `json:"logSetting"`
	CreatedAt   *int64      `json:"createdAt"`
	UpdatedAt   *int64      `json:"updatedAt"`
	Revision    *int64      `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId: core.CastString(data["namespaceId"]),
		Name:        core.CastString(data["name"]),
		Description: core.CastString(data["description"]),
		LogSetting:  NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:   core.CastInt64(data["createdAt"]),
		UpdatedAt:   core.CastInt64(data["updatedAt"]),
		Revision:    core.CastInt64(data["revision"]),
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
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"namespaceId": namespaceId,
		"name":        name,
		"description": description,
		"logSetting":  logSetting,
		"createdAt":   createdAt,
		"updatedAt":   updatedAt,
		"revision":    revision,
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

type Mutex struct {
	MutexId       *string `json:"mutexId"`
	UserId        *string `json:"userId"`
	PropertyId    *string `json:"propertyId"`
	TransactionId *string `json:"transactionId"`
	CreatedAt     *int64  `json:"createdAt"`
	Revision      *int64  `json:"revision"`
}

func NewMutexFromJson(data string) Mutex {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewMutexFromDict(dict)
}

func NewMutexFromDict(data map[string]interface{}) Mutex {
	return Mutex{
		MutexId:       core.CastString(data["mutexId"]),
		UserId:        core.CastString(data["userId"]),
		PropertyId:    core.CastString(data["propertyId"]),
		TransactionId: core.CastString(data["transactionId"]),
		CreatedAt:     core.CastInt64(data["createdAt"]),
		Revision:      core.CastInt64(data["revision"]),
	}
}

func (p Mutex) ToDict() map[string]interface{} {

	var mutexId *string
	if p.MutexId != nil {
		mutexId = p.MutexId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var propertyId *string
	if p.PropertyId != nil {
		propertyId = p.PropertyId
	}
	var transactionId *string
	if p.TransactionId != nil {
		transactionId = p.TransactionId
	}
	var createdAt *int64
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt
	}
	var revision *int64
	if p.Revision != nil {
		revision = p.Revision
	}
	return map[string]interface{}{
		"mutexId":       mutexId,
		"userId":        userId,
		"propertyId":    propertyId,
		"transactionId": transactionId,
		"createdAt":     createdAt,
		"revision":      revision,
	}
}

func (p Mutex) Pointer() *Mutex {
	return &p
}

func CastMutexes(data []interface{}) []Mutex {
	v := make([]Mutex, 0)
	for _, d := range data {
		v = append(v, NewMutexFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMutexesFromDict(data []Mutex) []interface{} {
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
