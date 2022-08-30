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

package megaField

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
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

type AreaModel struct {
	AreaModelId *string `json:"areaModelId"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
}

func NewAreaModelFromJson(data string) AreaModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAreaModelFromDict(dict)
}

func NewAreaModelFromDict(data map[string]interface{}) AreaModel {
    return AreaModel {
        AreaModelId: core.CastString(data["areaModelId"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
    }
}

func (p AreaModel) ToDict() map[string]interface{} {
    
    var areaModelId *string
    if p.AreaModelId != nil {
        areaModelId = p.AreaModelId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    return map[string]interface{} {
        "areaModelId": areaModelId,
        "name": name,
        "metadata": metadata,
    }
}

func (p AreaModel) Pointer() *AreaModel {
    return &p
}

func CastAreaModels(data []interface{}) []AreaModel {
	v := make([]AreaModel, 0)
	for _, d := range data {
		v = append(v, NewAreaModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAreaModelsFromDict(data []AreaModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type AreaModelMaster struct {
	AreaModelMasterId *string `json:"areaModelMasterId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewAreaModelMasterFromJson(data string) AreaModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAreaModelMasterFromDict(dict)
}

func NewAreaModelMasterFromDict(data map[string]interface{}) AreaModelMaster {
    return AreaModelMaster {
        AreaModelMasterId: core.CastString(data["areaModelMasterId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p AreaModelMaster) ToDict() map[string]interface{} {
    
    var areaModelMasterId *string
    if p.AreaModelMasterId != nil {
        areaModelMasterId = p.AreaModelMasterId
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
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "areaModelMasterId": areaModelMasterId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p AreaModelMaster) Pointer() *AreaModelMaster {
    return &p
}

func CastAreaModelMasters(data []interface{}) []AreaModelMaster {
	v := make([]AreaModelMaster, 0)
	for _, d := range data {
		v = append(v, NewAreaModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAreaModelMastersFromDict(data []AreaModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LayerModel struct {
	LayerModelId *string `json:"layerModelId"`
	AreaModelName *string `json:"areaModelName"`
	Name *string `json:"name"`
	Metadata *string `json:"metadata"`
}

func NewLayerModelFromJson(data string) LayerModel {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLayerModelFromDict(dict)
}

func NewLayerModelFromDict(data map[string]interface{}) LayerModel {
    return LayerModel {
        LayerModelId: core.CastString(data["layerModelId"]),
        AreaModelName: core.CastString(data["areaModelName"]),
        Name: core.CastString(data["name"]),
        Metadata: core.CastString(data["metadata"]),
    }
}

func (p LayerModel) ToDict() map[string]interface{} {
    
    var layerModelId *string
    if p.LayerModelId != nil {
        layerModelId = p.LayerModelId
    }
    var areaModelName *string
    if p.AreaModelName != nil {
        areaModelName = p.AreaModelName
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    return map[string]interface{} {
        "layerModelId": layerModelId,
        "areaModelName": areaModelName,
        "name": name,
        "metadata": metadata,
    }
}

func (p LayerModel) Pointer() *LayerModel {
    return &p
}

func CastLayerModels(data []interface{}) []LayerModel {
	v := make([]LayerModel, 0)
	for _, d := range data {
		v = append(v, NewLayerModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLayerModelsFromDict(data []LayerModel) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type LayerModelMaster struct {
	LayerModelMasterId *string `json:"layerModelMasterId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
}

func NewLayerModelMasterFromJson(data string) LayerModelMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLayerModelMasterFromDict(dict)
}

func NewLayerModelMasterFromDict(data map[string]interface{}) LayerModelMaster {
    return LayerModelMaster {
        LayerModelMasterId: core.CastString(data["layerModelMasterId"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
    }
}

func (p LayerModelMaster) ToDict() map[string]interface{} {
    
    var layerModelMasterId *string
    if p.LayerModelMasterId != nil {
        layerModelMasterId = p.LayerModelMasterId
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
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var updatedAt *int64
    if p.UpdatedAt != nil {
        updatedAt = p.UpdatedAt
    }
    return map[string]interface{} {
        "layerModelMasterId": layerModelMasterId,
        "name": name,
        "description": description,
        "metadata": metadata,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
    }
}

func (p LayerModelMaster) Pointer() *LayerModelMaster {
    return &p
}

func CastLayerModelMasters(data []interface{}) []LayerModelMaster {
	v := make([]LayerModelMaster, 0)
	for _, d := range data {
		v = append(v, NewLayerModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLayerModelMastersFromDict(data []LayerModelMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type CurrentFieldMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings *string `json:"settings"`
}

func NewCurrentFieldMasterFromJson(data string) CurrentFieldMaster {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCurrentFieldMasterFromDict(dict)
}

func NewCurrentFieldMasterFromDict(data map[string]interface{}) CurrentFieldMaster {
    return CurrentFieldMaster {
        NamespaceId: core.CastString(data["namespaceId"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p CurrentFieldMaster) ToDict() map[string]interface{} {
    
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

func (p CurrentFieldMaster) Pointer() *CurrentFieldMaster {
    return &p
}

func CastCurrentFieldMasters(data []interface{}) []CurrentFieldMaster {
	v := make([]CurrentFieldMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentFieldMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentFieldMastersFromDict(data []CurrentFieldMaster) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Layer struct {
	LayerId *string `json:"layerId"`
	AreaModelName *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
	NumberOfMinEntries *int32 `json:"numberOfMinEntries"`
	NumberOfMaxEntries *int32 `json:"numberOfMaxEntries"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewLayerFromJson(data string) Layer {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLayerFromDict(dict)
}

func NewLayerFromDict(data map[string]interface{}) Layer {
    return Layer {
        LayerId: core.CastString(data["layerId"]),
        AreaModelName: core.CastString(data["areaModelName"]),
        LayerModelName: core.CastString(data["layerModelName"]),
        NumberOfMinEntries: core.CastInt32(data["numberOfMinEntries"]),
        NumberOfMaxEntries: core.CastInt32(data["numberOfMaxEntries"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Layer) ToDict() map[string]interface{} {
    
    var layerId *string
    if p.LayerId != nil {
        layerId = p.LayerId
    }
    var areaModelName *string
    if p.AreaModelName != nil {
        areaModelName = p.AreaModelName
    }
    var layerModelName *string
    if p.LayerModelName != nil {
        layerModelName = p.LayerModelName
    }
    var numberOfMinEntries *int32
    if p.NumberOfMinEntries != nil {
        numberOfMinEntries = p.NumberOfMinEntries
    }
    var numberOfMaxEntries *int32
    if p.NumberOfMaxEntries != nil {
        numberOfMaxEntries = p.NumberOfMaxEntries
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "layerId": layerId,
        "areaModelName": areaModelName,
        "layerModelName": layerModelName,
        "numberOfMinEntries": numberOfMinEntries,
        "numberOfMaxEntries": numberOfMaxEntries,
        "createdAt": createdAt,
    }
}

func (p Layer) Pointer() *Layer {
    return &p
}

func CastLayers(data []interface{}) []Layer {
	v := make([]Layer, 0)
	for _, d := range data {
		v = append(v, NewLayerFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastLayersFromDict(data []Layer) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Spatial struct {
	SpatialId *string `json:"spatialId"`
	UserId *string `json:"userId"`
	AreaModelName *string `json:"areaModelName"`
	LayerModelName *string `json:"layerModelName"`
	Position *Position `json:"position"`
	Vector *Vector `json:"vector"`
	R *float32 `json:"r"`
	LastSyncAt *int64 `json:"lastSyncAt"`
	CreatedAt *int64 `json:"createdAt"`
}

func NewSpatialFromJson(data string) Spatial {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSpatialFromDict(dict)
}

func NewSpatialFromDict(data map[string]interface{}) Spatial {
    return Spatial {
        SpatialId: core.CastString(data["spatialId"]),
        UserId: core.CastString(data["userId"]),
        AreaModelName: core.CastString(data["areaModelName"]),
        LayerModelName: core.CastString(data["layerModelName"]),
        Position: NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
        Vector: NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
        R: core.CastFloat32(data["r"]),
        LastSyncAt: core.CastInt64(data["lastSyncAt"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
    }
}

func (p Spatial) ToDict() map[string]interface{} {
    
    var spatialId *string
    if p.SpatialId != nil {
        spatialId = p.SpatialId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var areaModelName *string
    if p.AreaModelName != nil {
        areaModelName = p.AreaModelName
    }
    var layerModelName *string
    if p.LayerModelName != nil {
        layerModelName = p.LayerModelName
    }
    var position map[string]interface{}
    if p.Position != nil {
        position = p.Position.ToDict()
    }
    var vector map[string]interface{}
    if p.Vector != nil {
        vector = p.Vector.ToDict()
    }
    var r *float32
    if p.R != nil {
        r = p.R
    }
    var lastSyncAt *int64
    if p.LastSyncAt != nil {
        lastSyncAt = p.LastSyncAt
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    return map[string]interface{} {
        "spatialId": spatialId,
        "userId": userId,
        "areaModelName": areaModelName,
        "layerModelName": layerModelName,
        "position": position,
        "vector": vector,
        "r": r,
        "lastSyncAt": lastSyncAt,
        "createdAt": createdAt,
    }
}

func (p Spatial) Pointer() *Spatial {
    return &p
}

func CastSpatials(data []interface{}) []Spatial {
	v := make([]Spatial, 0)
	for _, d := range data {
		v = append(v, NewSpatialFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSpatialsFromDict(data []Spatial) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Position struct {
	X *float32 `json:"x"`
	Y *float32 `json:"y"`
	Z *float32 `json:"z"`
}

func NewPositionFromJson(data string) Position {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewPositionFromDict(dict)
}

func NewPositionFromDict(data map[string]interface{}) Position {
    return Position {
        X: core.CastFloat32(data["x"]),
        Y: core.CastFloat32(data["y"]),
        Z: core.CastFloat32(data["z"]),
    }
}

func (p Position) ToDict() map[string]interface{} {
    
    var x *float32
    if p.X != nil {
        x = p.X
    }
    var y *float32
    if p.Y != nil {
        y = p.Y
    }
    var z *float32
    if p.Z != nil {
        z = p.Z
    }
    return map[string]interface{} {
        "x": x,
        "y": y,
        "z": z,
    }
}

func (p Position) Pointer() *Position {
    return &p
}

func CastPositions(data []interface{}) []Position {
	v := make([]Position, 0)
	for _, d := range data {
		v = append(v, NewPositionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPositionsFromDict(data []Position) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type MyPosition struct {
	Position *Position `json:"position"`
	Vector *Vector `json:"vector"`
	R *float32 `json:"r"`
}

func NewMyPositionFromJson(data string) MyPosition {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMyPositionFromDict(dict)
}

func NewMyPositionFromDict(data map[string]interface{}) MyPosition {
    return MyPosition {
        Position: NewPositionFromDict(core.CastMap(data["position"])).Pointer(),
        Vector: NewVectorFromDict(core.CastMap(data["vector"])).Pointer(),
        R: core.CastFloat32(data["r"]),
    }
}

func (p MyPosition) ToDict() map[string]interface{} {
    
    var position map[string]interface{}
    if p.Position != nil {
        position = p.Position.ToDict()
    }
    var vector map[string]interface{}
    if p.Vector != nil {
        vector = p.Vector.ToDict()
    }
    var r *float32
    if p.R != nil {
        r = p.R
    }
    return map[string]interface{} {
        "position": position,
        "vector": vector,
        "r": r,
    }
}

func (p MyPosition) Pointer() *MyPosition {
    return &p
}

func CastMyPositions(data []interface{}) []MyPosition {
	v := make([]MyPosition, 0)
	for _, d := range data {
		v = append(v, NewMyPositionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMyPositionsFromDict(data []MyPosition) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Scope struct {
	R *float32 `json:"r"`
	Limit *int32 `json:"limit"`
}

func NewScopeFromJson(data string) Scope {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewScopeFromDict(dict)
}

func NewScopeFromDict(data map[string]interface{}) Scope {
    return Scope {
        R: core.CastFloat32(data["r"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p Scope) ToDict() map[string]interface{} {
    
    var r *float32
    if p.R != nil {
        r = p.R
    }
    var limit *int32
    if p.Limit != nil {
        limit = p.Limit
    }
    return map[string]interface{} {
        "r": r,
        "limit": limit,
    }
}

func (p Scope) Pointer() *Scope {
    return &p
}

func CastScopes(data []interface{}) []Scope {
	v := make([]Scope, 0)
	for _, d := range data {
		v = append(v, NewScopeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScopesFromDict(data []Scope) []interface{} {
    v := make([]interface{}, 0)
    for _, d := range data {
        v = append(v, d.ToDict())
    }
    return v
}

type Vector struct {
	X *float32 `json:"x"`
	Y *float32 `json:"y"`
	Z *float32 `json:"z"`
}

func NewVectorFromJson(data string) Vector {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVectorFromDict(dict)
}

func NewVectorFromDict(data map[string]interface{}) Vector {
    return Vector {
        X: core.CastFloat32(data["x"]),
        Y: core.CastFloat32(data["y"]),
        Z: core.CastFloat32(data["z"]),
    }
}

func (p Vector) ToDict() map[string]interface{} {
    
    var x *float32
    if p.X != nil {
        x = p.X
    }
    var y *float32
    if p.Y != nil {
        y = p.Y
    }
    var z *float32
    if p.Z != nil {
        z = p.Z
    }
    return map[string]interface{} {
        "x": x,
        "y": y,
        "z": z,
    }
}

func (p Vector) Pointer() *Vector {
    return &p
}

func CastVectors(data []interface{}) []Vector {
	v := make([]Vector, 0)
	for _, d := range data {
		v = append(v, NewVectorFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVectorsFromDict(data []Vector) []interface{} {
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