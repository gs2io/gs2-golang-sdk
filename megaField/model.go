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
	"strconv"

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

type AreaModel struct {
	AreaModelId *string      `json:"areaModelId"`
	Name        *string      `json:"name"`
	Metadata    *string      `json:"metadata"`
	LayerModels []LayerModel `json:"layerModels"`
}

func (p *AreaModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AreaModel{}
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
		*p = AreaModel{}
	} else {
		*p = AreaModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["areaModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelId)
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
		if v, ok := d["layerModels"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LayerModels)
		}
	}
	return nil
}

func NewAreaModelFromJson(data string) AreaModel {
	req := AreaModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAreaModelFromDict(data map[string]interface{}) AreaModel {
	return AreaModel{
		AreaModelId: func() *string {
			v, ok := data["areaModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["areaModelId"])
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
		LayerModels: func() []LayerModel {
			if data["layerModels"] == nil {
				return nil
			}
			return CastLayerModels(core.CastArray(data["layerModels"]))
		}(),
	}
}

func (p AreaModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AreaModelId != nil {
		m["areaModelId"] = p.AreaModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.LayerModels != nil {
		m["layerModels"] = CastLayerModelsFromDict(
			p.LayerModels,
		)
	}
	return m
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
	Name              *string `json:"name"`
	Description       *string `json:"description"`
	Metadata          *string `json:"metadata"`
	CreatedAt         *int64  `json:"createdAt"`
	UpdatedAt         *int64  `json:"updatedAt"`
	Revision          *int64  `json:"revision"`
}

func (p *AreaModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AreaModelMaster{}
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
		*p = AreaModelMaster{}
	} else {
		*p = AreaModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["areaModelMasterId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelMasterId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelMasterId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelMasterId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelMasterId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelMasterId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelMasterId)
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

func NewAreaModelMasterFromJson(data string) AreaModelMaster {
	req := AreaModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAreaModelMasterFromDict(data map[string]interface{}) AreaModelMaster {
	return AreaModelMaster{
		AreaModelMasterId: func() *string {
			v, ok := data["areaModelMasterId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["areaModelMasterId"])
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

func (p AreaModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AreaModelMasterId != nil {
		m["areaModelMasterId"] = p.AreaModelMasterId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
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
	Name         *string `json:"name"`
	Metadata     *string `json:"metadata"`
}

func (p *LayerModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LayerModel{}
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
		*p = LayerModel{}
	} else {
		*p = LayerModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["layerModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelId)
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
	}
	return nil
}

func NewLayerModelFromJson(data string) LayerModel {
	req := LayerModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLayerModelFromDict(data map[string]interface{}) LayerModel {
	return LayerModel{
		LayerModelId: func() *string {
			v, ok := data["layerModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerModelId"])
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
	}
}

func (p LayerModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LayerModelId != nil {
		m["layerModelId"] = p.LayerModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	return m
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
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	Metadata           *string `json:"metadata"`
	CreatedAt          *int64  `json:"createdAt"`
	UpdatedAt          *int64  `json:"updatedAt"`
	Revision           *int64  `json:"revision"`
}

func (p *LayerModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = LayerModelMaster{}
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
		*p = LayerModelMaster{}
	} else {
		*p = LayerModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["layerModelMasterId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelMasterId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelMasterId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelMasterId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelMasterId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelMasterId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelMasterId)
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

func NewLayerModelMasterFromJson(data string) LayerModelMaster {
	req := LayerModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLayerModelMasterFromDict(data map[string]interface{}) LayerModelMaster {
	return LayerModelMaster{
		LayerModelMasterId: func() *string {
			v, ok := data["layerModelMasterId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerModelMasterId"])
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

func (p LayerModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LayerModelMasterId != nil {
		m["layerModelMasterId"] = p.LayerModelMasterId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Description != nil {
		m["description"] = p.Description
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
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
	Settings    *string `json:"settings"`
}

func (p *CurrentFieldMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentFieldMaster{}
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
		*p = CurrentFieldMaster{}
	} else {
		*p = CurrentFieldMaster{}
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

func NewCurrentFieldMasterFromJson(data string) CurrentFieldMaster {
	req := CurrentFieldMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentFieldMasterFromDict(data map[string]interface{}) CurrentFieldMaster {
	return CurrentFieldMaster{
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

func (p CurrentFieldMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
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
	LayerId            *string `json:"layerId"`
	AreaModelName      *string `json:"areaModelName"`
	LayerModelName     *string `json:"layerModelName"`
	NumberOfMinEntries *int32  `json:"numberOfMinEntries"`
	NumberOfMaxEntries *int32  `json:"numberOfMaxEntries"`
	CreatedAt          *int64  `json:"createdAt"`
}

func (p *Layer) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Layer{}
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
		*p = Layer{}
	} else {
		*p = Layer{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["layerId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerId)
				}
			}
		}
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["numberOfMinEntries"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NumberOfMinEntries)
		}
		if v, ok := d["numberOfMaxEntries"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NumberOfMaxEntries)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewLayerFromJson(data string) Layer {
	req := Layer{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewLayerFromDict(data map[string]interface{}) Layer {
	return Layer{
		LayerId: func() *string {
			v, ok := data["layerId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerId"])
		}(),
		AreaModelName: func() *string {
			v, ok := data["areaModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["areaModelName"])
		}(),
		LayerModelName: func() *string {
			v, ok := data["layerModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerModelName"])
		}(),
		NumberOfMinEntries: func() *int32 {
			v, ok := data["numberOfMinEntries"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["numberOfMinEntries"])
		}(),
		NumberOfMaxEntries: func() *int32 {
			v, ok := data["numberOfMaxEntries"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["numberOfMaxEntries"])
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

func (p Layer) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LayerId != nil {
		m["layerId"] = p.LayerId
	}
	if p.AreaModelName != nil {
		m["areaModelName"] = p.AreaModelName
	}
	if p.LayerModelName != nil {
		m["layerModelName"] = p.LayerModelName
	}
	if p.NumberOfMinEntries != nil {
		m["numberOfMinEntries"] = p.NumberOfMinEntries
	}
	if p.NumberOfMaxEntries != nil {
		m["numberOfMaxEntries"] = p.NumberOfMaxEntries
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
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
	SpatialId      *string   `json:"spatialId"`
	UserId         *string   `json:"userId"`
	AreaModelName  *string   `json:"areaModelName"`
	LayerModelName *string   `json:"layerModelName"`
	Position       *Position `json:"position"`
	Vector         *Vector   `json:"vector"`
	R              *float32  `json:"r"`
	LastSyncAt     *int64    `json:"lastSyncAt"`
	CreatedAt      *int64    `json:"createdAt"`
}

func (p *Spatial) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Spatial{}
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
		*p = Spatial{}
	} else {
		*p = Spatial{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["spatialId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SpatialId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SpatialId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SpatialId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SpatialId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SpatialId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SpatialId)
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
		if v, ok := d["areaModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AreaModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AreaModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AreaModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AreaModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AreaModelName)
				}
			}
		}
		if v, ok := d["layerModelName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerModelName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerModelName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerModelName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerModelName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerModelName)
				}
			}
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["vector"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Vector)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
		if v, ok := d["lastSyncAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.LastSyncAt)
		}
		if v, ok := d["createdAt"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CreatedAt)
		}
	}
	return nil
}

func NewSpatialFromJson(data string) Spatial {
	req := Spatial{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSpatialFromDict(data map[string]interface{}) Spatial {
	return Spatial{
		SpatialId: func() *string {
			v, ok := data["spatialId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["spatialId"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		AreaModelName: func() *string {
			v, ok := data["areaModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["areaModelName"])
		}(),
		LayerModelName: func() *string {
			v, ok := data["layerModelName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerModelName"])
		}(),
		Position: func() *Position {
			v, ok := data["position"]
			if !ok || v == nil {
				return nil
			}
			return NewPositionFromDict(core.CastMap(data["position"])).Pointer()
		}(),
		Vector: func() *Vector {
			v, ok := data["vector"]
			if !ok || v == nil {
				return nil
			}
			return NewVectorFromDict(core.CastMap(data["vector"])).Pointer()
		}(),
		R: func() *float32 {
			v, ok := data["r"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["r"])
		}(),
		LastSyncAt: func() *int64 {
			v, ok := data["lastSyncAt"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt64(data["lastSyncAt"])
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

func (p Spatial) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.SpatialId != nil {
		m["spatialId"] = p.SpatialId
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.AreaModelName != nil {
		m["areaModelName"] = p.AreaModelName
	}
	if p.LayerModelName != nil {
		m["layerModelName"] = p.LayerModelName
	}
	if p.Position != nil {
		m["position"] = func() map[string]interface{} {
			if p.Position == nil {
				return nil
			}
			return p.Position.ToDict()
		}()
	}
	if p.Vector != nil {
		m["vector"] = func() map[string]interface{} {
			if p.Vector == nil {
				return nil
			}
			return p.Vector.ToDict()
		}()
	}
	if p.R != nil {
		m["r"] = p.R
	}
	if p.LastSyncAt != nil {
		m["lastSyncAt"] = p.LastSyncAt
	}
	if p.CreatedAt != nil {
		m["createdAt"] = p.CreatedAt
	}
	return m
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

func (p *Position) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Position{}
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
		*p = Position{}
	} else {
		*p = Position{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["x"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.X)
		}
		if v, ok := d["y"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Y)
		}
		if v, ok := d["z"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Z)
		}
	}
	return nil
}

func NewPositionFromJson(data string) Position {
	req := Position{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPositionFromDict(data map[string]interface{}) Position {
	return Position{
		X: func() *float32 {
			v, ok := data["x"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["x"])
		}(),
		Y: func() *float32 {
			v, ok := data["y"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["y"])
		}(),
		Z: func() *float32 {
			v, ok := data["z"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["z"])
		}(),
	}
}

func (p Position) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.X != nil {
		m["x"] = p.X
	}
	if p.Y != nil {
		m["y"] = p.Y
	}
	if p.Z != nil {
		m["z"] = p.Z
	}
	return m
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
	Vector   *Vector   `json:"vector"`
	R        *float32  `json:"r"`
}

func (p *MyPosition) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = MyPosition{}
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
		*p = MyPosition{}
	} else {
		*p = MyPosition{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["position"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Position)
		}
		if v, ok := d["vector"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Vector)
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
	}
	return nil
}

func NewMyPositionFromJson(data string) MyPosition {
	req := MyPosition{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewMyPositionFromDict(data map[string]interface{}) MyPosition {
	return MyPosition{
		Position: func() *Position {
			v, ok := data["position"]
			if !ok || v == nil {
				return nil
			}
			return NewPositionFromDict(core.CastMap(data["position"])).Pointer()
		}(),
		Vector: func() *Vector {
			v, ok := data["vector"]
			if !ok || v == nil {
				return nil
			}
			return NewVectorFromDict(core.CastMap(data["vector"])).Pointer()
		}(),
		R: func() *float32 {
			v, ok := data["r"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["r"])
		}(),
	}
}

func (p MyPosition) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Position != nil {
		m["position"] = func() map[string]interface{} {
			if p.Position == nil {
				return nil
			}
			return p.Position.ToDict()
		}()
	}
	if p.Vector != nil {
		m["vector"] = func() map[string]interface{} {
			if p.Vector == nil {
				return nil
			}
			return p.Vector.ToDict()
		}()
	}
	if p.R != nil {
		m["r"] = p.R
	}
	return m
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
	LayerName *string  `json:"layerName"`
	R         *float32 `json:"r"`
	Limit     *int32   `json:"limit"`
}

func (p *Scope) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Scope{}
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
		*p = Scope{}
	} else {
		*p = Scope{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["layerName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.LayerName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.LayerName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.LayerName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.LayerName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.LayerName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.LayerName)
				}
			}
		}
		if v, ok := d["r"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.R)
		}
		if v, ok := d["limit"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Limit)
		}
	}
	return nil
}

func NewScopeFromJson(data string) Scope {
	req := Scope{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScopeFromDict(data map[string]interface{}) Scope {
	return Scope{
		LayerName: func() *string {
			v, ok := data["layerName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["layerName"])
		}(),
		R: func() *float32 {
			v, ok := data["r"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["r"])
		}(),
		Limit: func() *int32 {
			v, ok := data["limit"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["limit"])
		}(),
	}
}

func (p Scope) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.LayerName != nil {
		m["layerName"] = p.LayerName
	}
	if p.R != nil {
		m["r"] = p.R
	}
	if p.Limit != nil {
		m["limit"] = p.Limit
	}
	return m
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

func (p *Vector) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Vector{}
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
		*p = Vector{}
	} else {
		*p = Vector{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["x"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.X)
		}
		if v, ok := d["y"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Y)
		}
		if v, ok := d["z"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Z)
		}
	}
	return nil
}

func NewVectorFromJson(data string) Vector {
	req := Vector{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVectorFromDict(data map[string]interface{}) Vector {
	return Vector{
		X: func() *float32 {
			v, ok := data["x"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["x"])
		}(),
		Y: func() *float32 {
			v, ok := data["y"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["y"])
		}(),
		Z: func() *float32 {
			v, ok := data["z"]
			if !ok || v == nil {
				return nil
			}
			return core.CastFloat32(data["z"])
		}(),
	}
}

func (p Vector) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.X != nil {
		m["x"] = p.X
	}
	if p.Y != nil {
		m["y"] = p.Y
	}
	if p.Z != nil {
		m["z"] = p.Z
	}
	return m
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
	m := map[string]interface{}{}
	if p.ApiKeyId != nil {
		m["apiKeyId"] = p.ApiKeyId
	}
	if p.RepositoryName != nil {
		m["repositoryName"] = p.RepositoryName
	}
	if p.SourcePath != nil {
		m["sourcePath"] = p.SourcePath
	}
	if p.ReferenceType != nil {
		m["referenceType"] = p.ReferenceType
	}
	if p.CommitHash != nil {
		m["commitHash"] = p.CommitHash
	}
	if p.BranchName != nil {
		m["branchName"] = p.BranchName
	}
	if p.TagName != nil {
		m["tagName"] = p.TagName
	}
	return m
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
