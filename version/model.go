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

package version

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId                 *string        `json:"namespaceId"`
	Name                        *string        `json:"name"`
	Description                 *string        `json:"description"`
	AssumeUserId                *string        `json:"assumeUserId"`
	AcceptVersionScript         *ScriptSetting `json:"acceptVersionScript"`
	CheckVersionTriggerScriptId *string        `json:"checkVersionTriggerScriptId"`
	LogSetting                  *LogSetting    `json:"logSetting"`
	CreatedAt                   *int64         `json:"createdAt"`
	UpdatedAt                   *int64         `json:"updatedAt"`
	Revision                    *int64         `json:"revision"`
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
		if v, ok := d["assumeUserId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AssumeUserId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AssumeUserId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AssumeUserId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AssumeUserId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AssumeUserId)
				}
			}
		}
		if v, ok := d["acceptVersionScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcceptVersionScript)
		}
		if v, ok := d["checkVersionTriggerScriptId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.CheckVersionTriggerScriptId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.CheckVersionTriggerScriptId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.CheckVersionTriggerScriptId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.CheckVersionTriggerScriptId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.CheckVersionTriggerScriptId)
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
		AssumeUserId: func() *string {
			v, ok := data["assumeUserId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["assumeUserId"])
		}(),
		AcceptVersionScript: func() *ScriptSetting {
			v, ok := data["acceptVersionScript"]
			if !ok || v == nil {
				return nil
			}
			return NewScriptSettingFromDict(core.CastMap(data["acceptVersionScript"])).Pointer()
		}(),
		CheckVersionTriggerScriptId: func() *string {
			v, ok := data["checkVersionTriggerScriptId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["checkVersionTriggerScriptId"])
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
	if p.AssumeUserId != nil {
		m["assumeUserId"] = p.AssumeUserId
	}
	if p.AcceptVersionScript != nil {
		m["acceptVersionScript"] = func() map[string]interface{} {
			if p.AcceptVersionScript == nil {
				return nil
			}
			return p.AcceptVersionScript.ToDict()
		}()
	}
	if p.CheckVersionTriggerScriptId != nil {
		m["checkVersionTriggerScriptId"] = p.CheckVersionTriggerScriptId
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

type VersionModelMaster struct {
	VersionModelId   *string           `json:"versionModelId"`
	Name             *string           `json:"name"`
	Description      *string           `json:"description"`
	Metadata         *string           `json:"metadata"`
	Scope            *string           `json:"scope"`
	Type             *string           `json:"type"`
	CurrentVersion   *Version          `json:"currentVersion"`
	WarningVersion   *Version          `json:"warningVersion"`
	ErrorVersion     *Version          `json:"errorVersion"`
	ScheduleVersions []ScheduleVersion `json:"scheduleVersions"`
	NeedSignature    *bool             `json:"needSignature"`
	SignatureKeyId   *string           `json:"signatureKeyId"`
	CreatedAt        *int64            `json:"createdAt"`
	UpdatedAt        *int64            `json:"updatedAt"`
	Revision         *int64            `json:"revision"`
}

func (p *VersionModelMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VersionModelMaster{}
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
		*p = VersionModelMaster{}
	} else {
		*p = VersionModelMaster{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["versionModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionModelId)
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
		}
		if v, ok := d["warningVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WarningVersion)
		}
		if v, ok := d["errorVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorVersion)
		}
		if v, ok := d["scheduleVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScheduleVersions)
		}
		if v, ok := d["needSignature"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NeedSignature)
		}
		if v, ok := d["signatureKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SignatureKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SignatureKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SignatureKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SignatureKeyId)
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

func NewVersionModelMasterFromJson(data string) VersionModelMaster {
	req := VersionModelMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVersionModelMasterFromDict(data map[string]interface{}) VersionModelMaster {
	return VersionModelMaster{
		VersionModelId: func() *string {
			v, ok := data["versionModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionModelId"])
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
		Scope: func() *string {
			v, ok := data["scope"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scope"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
		WarningVersion: func() *Version {
			v, ok := data["warningVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer()
		}(),
		ErrorVersion: func() *Version {
			v, ok := data["errorVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer()
		}(),
		ScheduleVersions: func() []ScheduleVersion {
			if data["scheduleVersions"] == nil {
				return nil
			}
			return CastScheduleVersions(core.CastArray(data["scheduleVersions"]))
		}(),
		NeedSignature: func() *bool {
			v, ok := data["needSignature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["needSignature"])
		}(),
		SignatureKeyId: func() *string {
			v, ok := data["signatureKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signatureKeyId"])
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

func (p VersionModelMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.VersionModelId != nil {
		m["versionModelId"] = p.VersionModelId
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
	if p.Scope != nil {
		m["scope"] = p.Scope
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.CurrentVersion != nil {
		m["currentVersion"] = func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}()
	}
	if p.WarningVersion != nil {
		m["warningVersion"] = func() map[string]interface{} {
			if p.WarningVersion == nil {
				return nil
			}
			return p.WarningVersion.ToDict()
		}()
	}
	if p.ErrorVersion != nil {
		m["errorVersion"] = func() map[string]interface{} {
			if p.ErrorVersion == nil {
				return nil
			}
			return p.ErrorVersion.ToDict()
		}()
	}
	if p.ScheduleVersions != nil {
		m["scheduleVersions"] = CastScheduleVersionsFromDict(
			p.ScheduleVersions,
		)
	}
	if p.NeedSignature != nil {
		m["needSignature"] = p.NeedSignature
	}
	if p.SignatureKeyId != nil {
		m["signatureKeyId"] = p.SignatureKeyId
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

func (p VersionModelMaster) Pointer() *VersionModelMaster {
	return &p
}

func CastVersionModelMasters(data []interface{}) []VersionModelMaster {
	v := make([]VersionModelMaster, 0)
	for _, d := range data {
		v = append(v, NewVersionModelMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionModelMastersFromDict(data []VersionModelMaster) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type VersionModel struct {
	VersionModelId   *string           `json:"versionModelId"`
	Name             *string           `json:"name"`
	Metadata         *string           `json:"metadata"`
	Scope            *string           `json:"scope"`
	Type             *string           `json:"type"`
	CurrentVersion   *Version          `json:"currentVersion"`
	WarningVersion   *Version          `json:"warningVersion"`
	ErrorVersion     *Version          `json:"errorVersion"`
	ScheduleVersions []ScheduleVersion `json:"scheduleVersions"`
	NeedSignature    *bool             `json:"needSignature"`
	SignatureKeyId   *string           `json:"signatureKeyId"`
}

func (p *VersionModel) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = VersionModel{}
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
		*p = VersionModel{}
	} else {
		*p = VersionModel{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["versionModelId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionModelId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionModelId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionModelId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionModelId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionModelId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionModelId)
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
		if v, ok := d["scope"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Scope = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Scope = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Scope = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Scope = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Scope)
				}
			}
		}
		if v, ok := d["type"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Type = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Type = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Type = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Type = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Type)
				}
			}
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
		}
		if v, ok := d["warningVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WarningVersion)
		}
		if v, ok := d["errorVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorVersion)
		}
		if v, ok := d["scheduleVersions"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ScheduleVersions)
		}
		if v, ok := d["needSignature"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.NeedSignature)
		}
		if v, ok := d["signatureKeyId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.SignatureKeyId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.SignatureKeyId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.SignatureKeyId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.SignatureKeyId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.SignatureKeyId)
				}
			}
		}
	}
	return nil
}

func NewVersionModelFromJson(data string) VersionModel {
	req := VersionModel{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVersionModelFromDict(data map[string]interface{}) VersionModel {
	return VersionModel{
		VersionModelId: func() *string {
			v, ok := data["versionModelId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionModelId"])
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
		Scope: func() *string {
			v, ok := data["scope"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scope"])
		}(),
		Type: func() *string {
			v, ok := data["type"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["type"])
		}(),
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
		WarningVersion: func() *Version {
			v, ok := data["warningVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer()
		}(),
		ErrorVersion: func() *Version {
			v, ok := data["errorVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer()
		}(),
		ScheduleVersions: func() []ScheduleVersion {
			if data["scheduleVersions"] == nil {
				return nil
			}
			return CastScheduleVersions(core.CastArray(data["scheduleVersions"]))
		}(),
		NeedSignature: func() *bool {
			v, ok := data["needSignature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastBool(data["needSignature"])
		}(),
		SignatureKeyId: func() *string {
			v, ok := data["signatureKeyId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signatureKeyId"])
		}(),
	}
}

func (p VersionModel) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.VersionModelId != nil {
		m["versionModelId"] = p.VersionModelId
	}
	if p.Name != nil {
		m["name"] = p.Name
	}
	if p.Metadata != nil {
		m["metadata"] = p.Metadata
	}
	if p.Scope != nil {
		m["scope"] = p.Scope
	}
	if p.Type != nil {
		m["type"] = p.Type
	}
	if p.CurrentVersion != nil {
		m["currentVersion"] = func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}()
	}
	if p.WarningVersion != nil {
		m["warningVersion"] = func() map[string]interface{} {
			if p.WarningVersion == nil {
				return nil
			}
			return p.WarningVersion.ToDict()
		}()
	}
	if p.ErrorVersion != nil {
		m["errorVersion"] = func() map[string]interface{} {
			if p.ErrorVersion == nil {
				return nil
			}
			return p.ErrorVersion.ToDict()
		}()
	}
	if p.ScheduleVersions != nil {
		m["scheduleVersions"] = CastScheduleVersionsFromDict(
			p.ScheduleVersions,
		)
	}
	if p.NeedSignature != nil {
		m["needSignature"] = p.NeedSignature
	}
	if p.SignatureKeyId != nil {
		m["signatureKeyId"] = p.SignatureKeyId
	}
	return m
}

func (p VersionModel) Pointer() *VersionModel {
	return &p
}

func CastVersionModels(data []interface{}) []VersionModel {
	v := make([]VersionModel, 0)
	for _, d := range data {
		v = append(v, NewVersionModelFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionModelsFromDict(data []VersionModel) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AcceptVersion struct {
	AcceptVersionId *string  `json:"acceptVersionId"`
	VersionName     *string  `json:"versionName"`
	UserId          *string  `json:"userId"`
	Version         *Version `json:"version"`
	CreatedAt       *int64   `json:"createdAt"`
	UpdatedAt       *int64   `json:"updatedAt"`
	Revision        *int64   `json:"revision"`
}

func (p *AcceptVersion) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AcceptVersion{}
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
		*p = AcceptVersion{}
	} else {
		*p = AcceptVersion{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["acceptVersionId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AcceptVersionId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AcceptVersionId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AcceptVersionId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AcceptVersionId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AcceptVersionId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AcceptVersionId)
				}
			}
		}
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
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

func NewAcceptVersionFromJson(data string) AcceptVersion {
	req := AcceptVersion{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAcceptVersionFromDict(data map[string]interface{}) AcceptVersion {
	return AcceptVersion{
		AcceptVersionId: func() *string {
			v, ok := data["acceptVersionId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["acceptVersionId"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		UserId: func() *string {
			v, ok := data["userId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["userId"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
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

func (p AcceptVersion) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.AcceptVersionId != nil {
		m["acceptVersionId"] = p.AcceptVersionId
	}
	if p.VersionName != nil {
		m["versionName"] = p.VersionName
	}
	if p.UserId != nil {
		m["userId"] = p.UserId
	}
	if p.Version != nil {
		m["version"] = func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
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

func (p AcceptVersion) Pointer() *AcceptVersion {
	return &p
}

func CastAcceptVersions(data []interface{}) []AcceptVersion {
	v := make([]AcceptVersion, 0)
	for _, d := range data {
		v = append(v, NewAcceptVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAcceptVersionsFromDict(data []AcceptVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Status struct {
	VersionModel   *VersionModel `json:"versionModel"`
	CurrentVersion *Version      `json:"currentVersion"`
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
		if v, ok := d["versionModel"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.VersionModel)
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
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
		VersionModel: func() *VersionModel {
			v, ok := data["versionModel"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionModelFromDict(core.CastMap(data["versionModel"])).Pointer()
		}(),
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
	}
}

func (p Status) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.VersionModel != nil {
		m["versionModel"] = func() map[string]interface{} {
			if p.VersionModel == nil {
				return nil
			}
			return p.VersionModel.ToDict()
		}()
	}
	if p.CurrentVersion != nil {
		m["currentVersion"] = func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}()
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

type TargetVersion struct {
	VersionName *string  `json:"versionName"`
	Body        *string  `json:"body"`
	Signature   *string  `json:"signature"`
	Version     *Version `json:"version"`
}

func (p *TargetVersion) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = TargetVersion{}
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
		*p = TargetVersion{}
	} else {
		*p = TargetVersion{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
		if v, ok := d["body"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Body = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Body = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Body = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Body = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Body = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Body)
				}
			}
		}
		if v, ok := d["signature"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Signature = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Signature = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Signature = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Signature = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Signature = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Signature)
				}
			}
		}
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
		}
	}
	return nil
}

func NewTargetVersionFromJson(data string) TargetVersion {
	req := TargetVersion{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewTargetVersionFromDict(data map[string]interface{}) TargetVersion {
	return TargetVersion{
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		Body: func() *string {
			v, ok := data["body"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["body"])
		}(),
		Signature: func() *string {
			v, ok := data["signature"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["signature"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
		}(),
	}
}

func (p TargetVersion) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.VersionName != nil {
		m["versionName"] = p.VersionName
	}
	if p.Body != nil {
		m["body"] = p.Body
	}
	if p.Signature != nil {
		m["signature"] = p.Signature
	}
	if p.Version != nil {
		m["version"] = func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}()
	}
	return m
}

func (p TargetVersion) Pointer() *TargetVersion {
	return &p
}

func CastTargetVersions(data []interface{}) []TargetVersion {
	v := make([]TargetVersion, 0)
	for _, d := range data {
		v = append(v, NewTargetVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastTargetVersionsFromDict(data []TargetVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type SignTargetVersion struct {
	Region        *string  `json:"region"`
	NamespaceName *string  `json:"namespaceName"`
	VersionName   *string  `json:"versionName"`
	Version       *Version `json:"version"`
}

func (p *SignTargetVersion) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = SignTargetVersion{}
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
		*p = SignTargetVersion{}
	} else {
		*p = SignTargetVersion{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["region"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Region = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Region = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Region = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Region = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Region)
				}
			}
		}
		if v, ok := d["namespaceName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.NamespaceName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.NamespaceName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.NamespaceName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.NamespaceName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.NamespaceName)
				}
			}
		}
		if v, ok := d["versionName"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.VersionName = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.VersionName = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.VersionName = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.VersionName = &strValue
				default:
					_ = json.Unmarshal(*v, &p.VersionName)
				}
			}
		}
		if v, ok := d["version"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Version)
		}
	}
	return nil
}

func NewSignTargetVersionFromJson(data string) SignTargetVersion {
	req := SignTargetVersion{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewSignTargetVersionFromDict(data map[string]interface{}) SignTargetVersion {
	return SignTargetVersion{
		Region: func() *string {
			v, ok := data["region"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["region"])
		}(),
		NamespaceName: func() *string {
			v, ok := data["namespaceName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["namespaceName"])
		}(),
		VersionName: func() *string {
			v, ok := data["versionName"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["versionName"])
		}(),
		Version: func() *Version {
			v, ok := data["version"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["version"])).Pointer()
		}(),
	}
}

func (p SignTargetVersion) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Region != nil {
		m["region"] = p.Region
	}
	if p.NamespaceName != nil {
		m["namespaceName"] = p.NamespaceName
	}
	if p.VersionName != nil {
		m["versionName"] = p.VersionName
	}
	if p.Version != nil {
		m["version"] = func() map[string]interface{} {
			if p.Version == nil {
				return nil
			}
			return p.Version.ToDict()
		}()
	}
	return m
}

func (p SignTargetVersion) Pointer() *SignTargetVersion {
	return &p
}

func CastSignTargetVersions(data []interface{}) []SignTargetVersion {
	v := make([]SignTargetVersion, 0)
	for _, d := range data {
		v = append(v, NewSignTargetVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSignTargetVersionsFromDict(data []SignTargetVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type CurrentVersionMaster struct {
	NamespaceId *string `json:"namespaceId"`
	Settings    *string `json:"settings"`
}

func (p *CurrentVersionMaster) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = CurrentVersionMaster{}
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
		*p = CurrentVersionMaster{}
	} else {
		*p = CurrentVersionMaster{}
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

func NewCurrentVersionMasterFromJson(data string) CurrentVersionMaster {
	req := CurrentVersionMaster{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewCurrentVersionMasterFromDict(data map[string]interface{}) CurrentVersionMaster {
	return CurrentVersionMaster{
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

func (p CurrentVersionMaster) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.NamespaceId != nil {
		m["namespaceId"] = p.NamespaceId
	}
	if p.Settings != nil {
		m["settings"] = p.Settings
	}
	return m
}

func (p CurrentVersionMaster) Pointer() *CurrentVersionMaster {
	return &p
}

func CastCurrentVersionMasters(data []interface{}) []CurrentVersionMaster {
	v := make([]CurrentVersionMaster, 0)
	for _, d := range data {
		v = append(v, NewCurrentVersionMasterFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastCurrentVersionMastersFromDict(data []CurrentVersionMaster) []interface{} {
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

type Version struct {
	Major *int32 `json:"major"`
	Minor *int32 `json:"minor"`
	Micro *int32 `json:"micro"`
}

func (p *Version) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Version{}
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
		*p = Version{}
	} else {
		*p = Version{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["major"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Major)
		}
		if v, ok := d["minor"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Minor)
		}
		if v, ok := d["micro"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Micro)
		}
	}
	return nil
}

func NewVersionFromJson(data string) Version {
	req := Version{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewVersionFromDict(data map[string]interface{}) Version {
	return Version{
		Major: func() *int32 {
			v, ok := data["major"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["major"])
		}(),
		Minor: func() *int32 {
			v, ok := data["minor"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["minor"])
		}(),
		Micro: func() *int32 {
			v, ok := data["micro"]
			if !ok || v == nil {
				return nil
			}
			return core.CastInt32(data["micro"])
		}(),
	}
}

func (p Version) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.Major != nil {
		m["major"] = p.Major
	}
	if p.Minor != nil {
		m["minor"] = p.Minor
	}
	if p.Micro != nil {
		m["micro"] = p.Micro
	}
	return m
}

func (p Version) Pointer() *Version {
	return &p
}

func CastVersions(data []interface{}) []Version {
	v := make([]Version, 0)
	for _, d := range data {
		v = append(v, NewVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastVersionsFromDict(data []Version) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ScheduleVersion struct {
	CurrentVersion  *Version `json:"currentVersion"`
	WarningVersion  *Version `json:"warningVersion"`
	ErrorVersion    *Version `json:"errorVersion"`
	ScheduleEventId *string  `json:"scheduleEventId"`
}

func (p *ScheduleVersion) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = ScheduleVersion{}
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
		*p = ScheduleVersion{}
	} else {
		*p = ScheduleVersion{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["currentVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.CurrentVersion)
		}
		if v, ok := d["warningVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.WarningVersion)
		}
		if v, ok := d["errorVersion"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ErrorVersion)
		}
		if v, ok := d["scheduleEventId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.ScheduleEventId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.ScheduleEventId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.ScheduleEventId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleEventId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.ScheduleEventId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.ScheduleEventId)
				}
			}
		}
	}
	return nil
}

func NewScheduleVersionFromJson(data string) ScheduleVersion {
	req := ScheduleVersion{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewScheduleVersionFromDict(data map[string]interface{}) ScheduleVersion {
	return ScheduleVersion{
		CurrentVersion: func() *Version {
			v, ok := data["currentVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["currentVersion"])).Pointer()
		}(),
		WarningVersion: func() *Version {
			v, ok := data["warningVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["warningVersion"])).Pointer()
		}(),
		ErrorVersion: func() *Version {
			v, ok := data["errorVersion"]
			if !ok || v == nil {
				return nil
			}
			return NewVersionFromDict(core.CastMap(data["errorVersion"])).Pointer()
		}(),
		ScheduleEventId: func() *string {
			v, ok := data["scheduleEventId"]
			if !ok || v == nil {
				return nil
			}
			return core.CastString(data["scheduleEventId"])
		}(),
	}
}

func (p ScheduleVersion) ToDict() map[string]interface{} {
	m := map[string]interface{}{}
	if p.CurrentVersion != nil {
		m["currentVersion"] = func() map[string]interface{} {
			if p.CurrentVersion == nil {
				return nil
			}
			return p.CurrentVersion.ToDict()
		}()
	}
	if p.WarningVersion != nil {
		m["warningVersion"] = func() map[string]interface{} {
			if p.WarningVersion == nil {
				return nil
			}
			return p.WarningVersion.ToDict()
		}()
	}
	if p.ErrorVersion != nil {
		m["errorVersion"] = func() map[string]interface{} {
			if p.ErrorVersion == nil {
				return nil
			}
			return p.ErrorVersion.ToDict()
		}()
	}
	if p.ScheduleEventId != nil {
		m["scheduleEventId"] = p.ScheduleEventId
	}
	return m
}

func (p ScheduleVersion) Pointer() *ScheduleVersion {
	return &p
}

func CastScheduleVersions(data []interface{}) []ScheduleVersion {
	v := make([]ScheduleVersion, 0)
	for _, d := range data {
		v = append(v, NewScheduleVersionFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastScheduleVersionsFromDict(data []ScheduleVersion) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}
