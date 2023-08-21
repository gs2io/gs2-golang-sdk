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

package stamina

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	OverflowTriggerScript *ScriptSetting `json:"overflowTriggerScript"`
	LogSetting *LogSetting `json:"logSetting"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
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
        OverflowTriggerScript: NewScriptSettingFromDict(core.CastMap(data["overflowTriggerScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
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
    var overflowTriggerScript map[string]interface{}
    if p.OverflowTriggerScript != nil {
        overflowTriggerScript = p.OverflowTriggerScript.ToDict()
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
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "name": name,
        "description": description,
        "overflowTriggerScript": overflowTriggerScript,
        "logSetting": logSetting,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
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

type StaminaModelMaster struct {
	StaminaModelId *string `json:"staminaModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	RecoverIntervalMinutes *int32 `json:"recoverIntervalMinutes"`
	RecoverValue *int32 `json:"recoverValue"`
	InitialCapacity *int32 `json:"initialCapacity"`
	IsOverflow *bool `json:"isOverflow"`
	MaxCapacity *int32 `json:"maxCapacity"`
	MaxStaminaTableName *string `json:"maxStaminaTableName"`
	RecoverIntervalTableName *string `json:"recoverIntervalTableName"`
	RecoverValueTableName *string `json:"recoverValueTableName"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewStaminaModelMasterFromJson(data string) StaminaModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStaminaModelMasterFromDict(dict)
}

func NewStaminaModelMasterFromDict(data map[string]interface{}) StaminaModelMaster {
    return StaminaModelMaster {
        StaminaModelId: core.CastString(data["staminaModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
        RecoverValue: core.CastInt32(data["recoverValue"]),
        InitialCapacity: core.CastInt32(data["initialCapacity"]),
        IsOverflow: core.CastBool(data["isOverflow"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
        MaxStaminaTableName: core.CastString(data["maxStaminaTableName"]),
        RecoverIntervalTableName: core.CastString(data["recoverIntervalTableName"]),
        RecoverValueTableName: core.CastString(data["recoverValueTableName"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p StaminaModelMaster) ToDict() map[string]interface{} {
    
    var staminaModelId *string
    if p.StaminaModelId != nil {
        staminaModelId = p.StaminaModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var recoverIntervalMinutes *int32
    if p.RecoverIntervalMinutes != nil {
        recoverIntervalMinutes = p.RecoverIntervalMinutes
    }
    var recoverValue *int32
    if p.RecoverValue != nil {
        recoverValue = p.RecoverValue
    }
    var initialCapacity *int32
    if p.InitialCapacity != nil {
        initialCapacity = p.InitialCapacity
    }
    var isOverflow *bool
    if p.IsOverflow != nil {
        isOverflow = p.IsOverflow
    }
    var maxCapacity *int32
    if p.MaxCapacity != nil {
        maxCapacity = p.MaxCapacity
    }
    var maxStaminaTableName *string
    if p.MaxStaminaTableName != nil {
        maxStaminaTableName = p.MaxStaminaTableName
    }
    var recoverIntervalTableName *string
    if p.RecoverIntervalTableName != nil {
        recoverIntervalTableName = p.RecoverIntervalTableName
    }
    var recoverValueTableName *string
    if p.RecoverValueTableName != nil {
        recoverValueTableName = p.RecoverValueTableName
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
    return map[string]interface{} {
        "staminaModelId": staminaModelId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "recoverIntervalMinutes": recoverIntervalMinutes,
        "recoverValue": recoverValue,
        "initialCapacity": initialCapacity,
        "isOverflow": isOverflow,
        "maxCapacity": maxCapacity,
        "maxStaminaTableName": maxStaminaTableName,
        "recoverIntervalTableName": recoverIntervalTableName,
        "recoverValueTableName": recoverValueTableName,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p StaminaModelMaster) Pointer() *StaminaModelMaster {
    return &p
}

func CastStaminaModelMasters(data []interface{}) []StaminaModelMaster {
	v := make([]StaminaModelMaster, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelMastersFromDict(data []StaminaModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MaxStaminaTableMaster struct {
	MaxStaminaTableId *string `json:"maxStaminaTableId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewMaxStaminaTableMasterFromJson(data string) MaxStaminaTableMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMaxStaminaTableMasterFromDict(dict)
}

func NewMaxStaminaTableMasterFromDict(data map[string]interface{}) MaxStaminaTableMaster {
    return MaxStaminaTableMaster {
        MaxStaminaTableId: core.CastString(data["maxStaminaTableId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p MaxStaminaTableMaster) ToDict() map[string]interface{} {
    
    var maxStaminaTableId *string
    if p.MaxStaminaTableId != nil {
        maxStaminaTableId = p.MaxStaminaTableId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
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
    return map[string]interface{} {
        "maxStaminaTableId": maxStaminaTableId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "experienceModelId": experienceModelId,
        "values": values,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p MaxStaminaTableMaster) Pointer() *MaxStaminaTableMaster {
    return &p
}

func CastMaxStaminaTableMasters(data []interface{}) []MaxStaminaTableMaster {
	v := make([]MaxStaminaTableMaster, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTableMastersFromDict(data []MaxStaminaTableMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RecoverIntervalTableMaster struct {
	RecoverIntervalTableId *string `json:"recoverIntervalTableId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewRecoverIntervalTableMasterFromJson(data string) RecoverIntervalTableMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRecoverIntervalTableMasterFromDict(dict)
}

func NewRecoverIntervalTableMasterFromDict(data map[string]interface{}) RecoverIntervalTableMaster {
    return RecoverIntervalTableMaster {
        RecoverIntervalTableId: core.CastString(data["recoverIntervalTableId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p RecoverIntervalTableMaster) ToDict() map[string]interface{} {
    
    var recoverIntervalTableId *string
    if p.RecoverIntervalTableId != nil {
        recoverIntervalTableId = p.RecoverIntervalTableId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
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
    return map[string]interface{} {
        "recoverIntervalTableId": recoverIntervalTableId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "experienceModelId": experienceModelId,
        "values": values,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p RecoverIntervalTableMaster) Pointer() *RecoverIntervalTableMaster {
    return &p
}

func CastRecoverIntervalTableMasters(data []interface{}) []RecoverIntervalTableMaster {
	v := make([]RecoverIntervalTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTableMastersFromDict(data []RecoverIntervalTableMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RecoverValueTableMaster struct {
	RecoverValueTableId *string `json:"recoverValueTableId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	Description *string `json:"description"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewRecoverValueTableMasterFromJson(data string) RecoverValueTableMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRecoverValueTableMasterFromDict(dict)
}

func NewRecoverValueTableMasterFromDict(data map[string]interface{}) RecoverValueTableMaster {
    return RecoverValueTableMaster {
        RecoverValueTableId: core.CastString(data["recoverValueTableId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        Description: core.CastString(data["description"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p RecoverValueTableMaster) ToDict() map[string]interface{} {
    
    var recoverValueTableId *string
    if p.RecoverValueTableId != nil {
        recoverValueTableId = p.RecoverValueTableId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var description *string
    if p.Description != nil {
        description = p.Description
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
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
    return map[string]interface{} {
        "recoverValueTableId": recoverValueTableId,
        "name": name,
        "metadata": metadata,
        "description": description,
        "experienceModelId": experienceModelId,
        "values": values,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p RecoverValueTableMaster) Pointer() *RecoverValueTableMaster {
    return &p
}

func CastRecoverValueTableMasters(data []interface{}) []RecoverValueTableMaster {
	v := make([]RecoverValueTableMaster, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTableMastersFromDict(data []RecoverValueTableMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentStaminaMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentStaminaMasterFromJson(data string) CurrentStaminaMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentStaminaMasterFromDict(dict)
}

func NewCurrentStaminaMasterFromDict(data map[string]interface{}) CurrentStaminaMaster {
    return CurrentStaminaMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentStaminaMaster) ToDict() map[string]interface{} {
    
    var namespaceId *string
    if p.NamespaceId != nil {
        namespaceId = p.NamespaceId
    }
    var settings *string
    if p.Settings != nil {
        settings = p.Settings
    }
    return map[string]interface{} {
        "namespaceId": namespaceId,
        "settings": settings,
    }
}

func (p CurrentStaminaMaster) Pointer() *CurrentStaminaMaster {
    return &p
}

func CastCurrentStaminaMasters(data []interface{}) []CurrentStaminaMaster {
	v := make([]CurrentStaminaMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentStaminaMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentStaminaMastersFromDict(data []CurrentStaminaMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type StaminaModel struct {
	StaminaModelId *string `json:"staminaModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	RecoverIntervalMinutes *int32 `json:"recoverIntervalMinutes"`
	RecoverValue *int32 `json:"recoverValue"`
	InitialCapacity *int32 `json:"initialCapacity"`
	IsOverflow *bool `json:"isOverflow"`
	MaxCapacity *int32 `json:"maxCapacity"`
	MaxStaminaTable *MaxStaminaTable `json:"maxStaminaTable"`
	RecoverIntervalTable *RecoverIntervalTable `json:"recoverIntervalTable"`
	RecoverValueTable *RecoverValueTable `json:"recoverValueTable"`
}

func NewStaminaModelFromJson(data string) StaminaModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStaminaModelFromDict(dict)
}

func NewStaminaModelFromDict(data map[string]interface{}) StaminaModel {
    return StaminaModel {
        StaminaModelId: core.CastString(data["staminaModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
        RecoverValue: core.CastInt32(data["recoverValue"]),
        InitialCapacity: core.CastInt32(data["initialCapacity"]),
        IsOverflow: core.CastBool(data["isOverflow"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
        MaxStaminaTable: NewMaxStaminaTableFromDict(core.CastMap(data["maxStaminaTable"])).Pointer(),
        RecoverIntervalTable: NewRecoverIntervalTableFromDict(core.CastMap(data["recoverIntervalTable"])).Pointer(),
        RecoverValueTable: NewRecoverValueTableFromDict(core.CastMap(data["recoverValueTable"])).Pointer(),
    }
}

func (p StaminaModel) ToDict() map[string]interface{} {
    
    var staminaModelId *string
    if p.StaminaModelId != nil {
        staminaModelId = p.StaminaModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var recoverIntervalMinutes *int32
    if p.RecoverIntervalMinutes != nil {
        recoverIntervalMinutes = p.RecoverIntervalMinutes
    }
    var recoverValue *int32
    if p.RecoverValue != nil {
        recoverValue = p.RecoverValue
    }
    var initialCapacity *int32
    if p.InitialCapacity != nil {
        initialCapacity = p.InitialCapacity
    }
    var isOverflow *bool
    if p.IsOverflow != nil {
        isOverflow = p.IsOverflow
    }
    var maxCapacity *int32
    if p.MaxCapacity != nil {
        maxCapacity = p.MaxCapacity
    }
    var maxStaminaTable map[string]interface{}
    if p.MaxStaminaTable != nil {
        maxStaminaTable = p.MaxStaminaTable.ToDict()
    }
    var recoverIntervalTable map[string]interface{}
    if p.RecoverIntervalTable != nil {
        recoverIntervalTable = p.RecoverIntervalTable.ToDict()
    }
    var recoverValueTable map[string]interface{}
    if p.RecoverValueTable != nil {
        recoverValueTable = p.RecoverValueTable.ToDict()
    }
    return map[string]interface{} {
        "staminaModelId": staminaModelId,
        "name": name,
        "metadata": metadata,
        "recoverIntervalMinutes": recoverIntervalMinutes,
        "recoverValue": recoverValue,
        "initialCapacity": initialCapacity,
        "isOverflow": isOverflow,
        "maxCapacity": maxCapacity,
        "maxStaminaTable": maxStaminaTable,
        "recoverIntervalTable": recoverIntervalTable,
        "recoverValueTable": recoverValueTable,
    }
}

func (p StaminaModel) Pointer() *StaminaModel {
    return &p
}

func CastStaminaModels(data []interface{}) []StaminaModel {
	v := make([]StaminaModel, 0)
	for _, d := range data {
		v = append(v, NewStaminaModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminaModelsFromDict(data []StaminaModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MaxStaminaTable struct {
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
}

func NewMaxStaminaTableFromJson(data string) MaxStaminaTable {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMaxStaminaTableFromDict(dict)
}

func NewMaxStaminaTableFromDict(data map[string]interface{}) MaxStaminaTable {
    return MaxStaminaTable {
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
    }
}

func (p MaxStaminaTable) ToDict() map[string]interface{} {
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
    }
    return map[string]interface{} {
        "name": name,
        "metadata": metadata,
        "experienceModelId": experienceModelId,
        "values": values,
    }
}

func (p MaxStaminaTable) Pointer() *MaxStaminaTable {
    return &p
}

func CastMaxStaminaTables(data []interface{}) []MaxStaminaTable {
	v := make([]MaxStaminaTable, 0)
	for _, d := range data {
		v = append(v, NewMaxStaminaTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMaxStaminaTablesFromDict(data []MaxStaminaTable) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RecoverIntervalTable struct {
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
}

func NewRecoverIntervalTableFromJson(data string) RecoverIntervalTable {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRecoverIntervalTableFromDict(dict)
}

func NewRecoverIntervalTableFromDict(data map[string]interface{}) RecoverIntervalTable {
    return RecoverIntervalTable {
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
    }
}

func (p RecoverIntervalTable) ToDict() map[string]interface{} {
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
    }
    return map[string]interface{} {
        "name": name,
        "metadata": metadata,
        "experienceModelId": experienceModelId,
        "values": values,
    }
}

func (p RecoverIntervalTable) Pointer() *RecoverIntervalTable {
    return &p
}

func CastRecoverIntervalTables(data []interface{}) []RecoverIntervalTable {
	v := make([]RecoverIntervalTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverIntervalTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverIntervalTablesFromDict(data []RecoverIntervalTable) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type RecoverValueTable struct {
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
	ExperienceModelId *string `json:"experienceModelId"`
	Values []*int32 `json:"values"`
}

func NewRecoverValueTableFromJson(data string) RecoverValueTable {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRecoverValueTableFromDict(dict)
}

func NewRecoverValueTableFromDict(data map[string]interface{}) RecoverValueTable {
    return RecoverValueTable {
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
        ExperienceModelId: core.CastString(data["experienceModelId"]),
        Values: core.CastInt32s(core.CastArray(data["values"])),
    }
}

func (p RecoverValueTable) ToDict() map[string]interface{} {
    
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var experienceModelId *string
    if p.ExperienceModelId != nil {
        experienceModelId = p.ExperienceModelId
    }
    var values []interface{}
    if p.Values != nil {
        values = core.CastInt32sFromDict(
            p.Values,
        )
    }
    return map[string]interface{} {
        "name": name,
        "metadata": metadata,
        "experienceModelId": experienceModelId,
        "values": values,
    }
}

func (p RecoverValueTable) Pointer() *RecoverValueTable {
    return &p
}

func CastRecoverValueTables(data []interface{}) []RecoverValueTable {
	v := make([]RecoverValueTable, 0)
	for _, d := range data {
		v = append(v, NewRecoverValueTableFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRecoverValueTablesFromDict(data []RecoverValueTable) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Stamina struct {
	StaminaId *string `json:"staminaId"`
	StaminaName *string `json:"staminaName"`
	UserId *string `json:"userId"`
	Value *int32 `json:"value"`
	MaxValue *int32 `json:"maxValue"`
	RecoverIntervalMinutes *int32 `json:"recoverIntervalMinutes"`
	RecoverValue *int32 `json:"recoverValue"`
	OverflowValue *int32 `json:"overflowValue"`
	NextRecoverAt *int64 `json:"nextRecoverAt"`
	LastRecoveredAt *int64 `json:"lastRecoveredAt"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewStaminaFromJson(data string) Stamina {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewStaminaFromDict(dict)
}

func NewStaminaFromDict(data map[string]interface{}) Stamina {
    return Stamina {
        StaminaId: core.CastString(data["staminaId"]),
        StaminaName: core.CastString(data["staminaName"]),
        UserId: core.CastString(data["userId"]),
        Value: core.CastInt32(data["value"]),
        MaxValue: core.CastInt32(data["maxValue"]),
        RecoverIntervalMinutes: core.CastInt32(data["recoverIntervalMinutes"]),
        RecoverValue: core.CastInt32(data["recoverValue"]),
        OverflowValue: core.CastInt32(data["overflowValue"]),
        NextRecoverAt: core.CastInt64(data["nextRecoverAt"]),
        LastRecoveredAt: core.CastInt64(data["lastRecoveredAt"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Stamina) ToDict() map[string]interface{} {
    
    var staminaId *string
    if p.StaminaId != nil {
        staminaId = p.StaminaId
    }
    var staminaName *string
    if p.StaminaName != nil {
        staminaName = p.StaminaName
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var value *int32
    if p.Value != nil {
        value = p.Value
    }
    var maxValue *int32
    if p.MaxValue != nil {
        maxValue = p.MaxValue
    }
    var recoverIntervalMinutes *int32
    if p.RecoverIntervalMinutes != nil {
        recoverIntervalMinutes = p.RecoverIntervalMinutes
    }
    var recoverValue *int32
    if p.RecoverValue != nil {
        recoverValue = p.RecoverValue
    }
    var overflowValue *int32
    if p.OverflowValue != nil {
        overflowValue = p.OverflowValue
    }
    var nextRecoverAt *int64
    if p.NextRecoverAt != nil {
        nextRecoverAt = p.NextRecoverAt
    }
    var lastRecoveredAt *int64
    if p.LastRecoveredAt != nil {
        lastRecoveredAt = p.LastRecoveredAt
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
    return map[string]interface{} {
        "staminaId": staminaId,
        "staminaName": staminaName,
        "userId": userId,
        "value": value,
        "maxValue": maxValue,
        "recoverIntervalMinutes": recoverIntervalMinutes,
        "recoverValue": recoverValue,
        "overflowValue": overflowValue,
        "nextRecoverAt": nextRecoverAt,
        "lastRecoveredAt": lastRecoveredAt,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
    }
}

func (p Stamina) Pointer() *Stamina {
    return &p
}

func CastStaminas(data []interface{}) []Stamina {
	v := make([]Stamina, 0)
	for _, d := range data {
		v = append(v, NewStaminaFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastStaminasFromDict(data []Stamina) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type GitHubCheckoutSetting struct {
	ApiKeyId *string `json:"apiKeyId"`
	RepositoryName *string `json:"repositoryName"`
	SourcePath *string `json:"sourcePath"`
	ReferenceType *string `json:"referenceType"`
	CommitHash *string `json:"commitHash"`
	BranchName *string `json:"branchName"`
	TagName *string `json:"tagName"`
}

func NewGitHubCheckoutSettingFromJson(data string) GitHubCheckoutSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGitHubCheckoutSettingFromDict(dict)
}

func NewGitHubCheckoutSettingFromDict(data map[string]interface{}) GitHubCheckoutSetting {
    return GitHubCheckoutSetting {
        ApiKeyId: core.CastString(data["apiKeyId"]),
        RepositoryName: core.CastString(data["repositoryName"]),
        SourcePath: core.CastString(data["sourcePath"]),
        ReferenceType: core.CastString(data["referenceType"]),
        CommitHash: core.CastString(data["commitHash"]),
        BranchName: core.CastString(data["branchName"]),
        TagName: core.CastString(data["tagName"]),
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
    return map[string]interface{} {
        "apiKeyId": apiKeyId,
        "repositoryName": repositoryName,
        "sourcePath": sourcePath,
        "referenceType": referenceType,
        "commitHash": commitHash,
        "branchName": branchName,
        "tagName": tagName,
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