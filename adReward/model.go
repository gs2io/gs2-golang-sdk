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

package adReward

import (
	"encoding/json"
	"strconv"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId             *string              `json:"namespaceId"`
	Name                    *string              `json:"name"`
	Description             *string              `json:"description"`
	Admob                   *AdMob               `json:"admob"`
	UnityAd                 *UnityAd             `json:"unityAd"`
	AppLovinMaxes           []AppLovinMax        `json:"appLovinMaxes"`
	AcquirePointScript      *ScriptSetting       `json:"acquirePointScript"`
	ConsumePointScript      *ScriptSetting       `json:"consumePointScript"`
	ChangePointNotification *NotificationSetting `json:"changePointNotification"`
	LogSetting              *LogSetting          `json:"logSetting"`
	CreatedAt               *int64               `json:"createdAt"`
	UpdatedAt               *int64               `json:"updatedAt"`
	Revision                *int64               `json:"revision"`
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
		if v, ok := d["admob"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Admob)
		}
		if v, ok := d["unityAd"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.UnityAd)
		}
		if v, ok := d["appLovinMaxes"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AppLovinMaxes)
		}
		if v, ok := d["acquirePointScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.AcquirePointScript)
		}
		if v, ok := d["consumePointScript"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ConsumePointScript)
		}
		if v, ok := d["changePointNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.ChangePointNotification)
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
		NamespaceId:             core.CastString(data["namespaceId"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Admob:                   NewAdMobFromDict(core.CastMap(data["admob"])).Pointer(),
		UnityAd:                 NewUnityAdFromDict(core.CastMap(data["unityAd"])).Pointer(),
		AppLovinMaxes:           CastAppLovinMaxes(core.CastArray(data["appLovinMaxes"])),
		AcquirePointScript:      NewScriptSettingFromDict(core.CastMap(data["acquirePointScript"])).Pointer(),
		ConsumePointScript:      NewScriptSettingFromDict(core.CastMap(data["consumePointScript"])).Pointer(),
		ChangePointNotification: NewNotificationSettingFromDict(core.CastMap(data["changePointNotification"])).Pointer(),
		LogSetting:              NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:               core.CastInt64(data["createdAt"]),
		UpdatedAt:               core.CastInt64(data["updatedAt"]),
		Revision:                core.CastInt64(data["revision"]),
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
	var admob map[string]interface{}
	if p.Admob != nil {
		admob = p.Admob.ToDict()
	}
	var unityAd map[string]interface{}
	if p.UnityAd != nil {
		unityAd = p.UnityAd.ToDict()
	}
	var appLovinMaxes []interface{}
	if p.AppLovinMaxes != nil {
		appLovinMaxes = CastAppLovinMaxesFromDict(
			p.AppLovinMaxes,
		)
	}
	var acquirePointScript map[string]interface{}
	if p.AcquirePointScript != nil {
		acquirePointScript = p.AcquirePointScript.ToDict()
	}
	var consumePointScript map[string]interface{}
	if p.ConsumePointScript != nil {
		consumePointScript = p.ConsumePointScript.ToDict()
	}
	var changePointNotification map[string]interface{}
	if p.ChangePointNotification != nil {
		changePointNotification = p.ChangePointNotification.ToDict()
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
		"namespaceId":             namespaceId,
		"name":                    name,
		"description":             description,
		"admob":                   admob,
		"unityAd":                 unityAd,
		"appLovinMaxes":           appLovinMaxes,
		"acquirePointScript":      acquirePointScript,
		"consumePointScript":      consumePointScript,
		"changePointNotification": changePointNotification,
		"logSetting":              logSetting,
		"createdAt":               createdAt,
		"updatedAt":               updatedAt,
		"revision":                revision,
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

type Point struct {
	PointId   *string `json:"pointId"`
	UserId    *string `json:"userId"`
	Point     *int64  `json:"point"`
	CreatedAt *int64  `json:"createdAt"`
	UpdatedAt *int64  `json:"updatedAt"`
	Revision  *int64  `json:"revision"`
}

func (p *Point) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = Point{}
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
		*p = Point{}
	} else {
		*p = Point{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["pointId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.PointId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.PointId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.PointId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.PointId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.PointId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.PointId)
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
		if v, ok := d["point"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.Point)
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

func NewPointFromJson(data string) Point {
	req := Point{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewPointFromDict(data map[string]interface{}) Point {
	return Point{
		PointId:   core.CastString(data["pointId"]),
		UserId:    core.CastString(data["userId"]),
		Point:     core.CastInt64(data["point"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
		UpdatedAt: core.CastInt64(data["updatedAt"]),
		Revision:  core.CastInt64(data["revision"]),
	}
}

func (p Point) ToDict() map[string]interface{} {

	var pointId *string
	if p.PointId != nil {
		pointId = p.PointId
	}
	var userId *string
	if p.UserId != nil {
		userId = p.UserId
	}
	var point *int64
	if p.Point != nil {
		point = p.Point
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
		"pointId":   pointId,
		"userId":    userId,
		"point":     point,
		"createdAt": createdAt,
		"updatedAt": updatedAt,
		"revision":  revision,
	}
}

func (p Point) Pointer() *Point {
	return &p
}

func CastPoints(data []interface{}) []Point {
	v := make([]Point, 0)
	for _, d := range data {
		v = append(v, NewPointFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastPointsFromDict(data []Point) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AdMob struct {
	AllowAdUnitIds []*string `json:"allowAdUnitIds"`
}

func (p *AdMob) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AdMob{}
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
		*p = AdMob{}
	} else {
		*p = AdMob{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["allowAdUnitIds"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.AllowAdUnitIds = l
			}
		}
	}
	return nil
}

func NewAdMobFromJson(data string) AdMob {
	req := AdMob{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAdMobFromDict(data map[string]interface{}) AdMob {
	return AdMob{
		AllowAdUnitIds: core.CastStrings(core.CastArray(data["allowAdUnitIds"])),
	}
}

func (p AdMob) ToDict() map[string]interface{} {

	var allowAdUnitIds []interface{}
	if p.AllowAdUnitIds != nil {
		allowAdUnitIds = core.CastStringsFromDict(
			p.AllowAdUnitIds,
		)
	}
	return map[string]interface{}{
		"allowAdUnitIds": allowAdUnitIds,
	}
}

func (p AdMob) Pointer() *AdMob {
	return &p
}

func CastAdMobs(data []interface{}) []AdMob {
	v := make([]AdMob, 0)
	for _, d := range data {
		v = append(v, NewAdMobFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAdMobsFromDict(data []AdMob) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type UnityAd struct {
	Keys []*string `json:"keys"`
}

func (p *UnityAd) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = UnityAd{}
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
		*p = UnityAd{}
	} else {
		*p = UnityAd{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["keys"]; ok && v != nil {
			var v2 []interface{}
			if err := json.Unmarshal(*v, &v2); err == nil {
				l := make([]*string, len(v2))
				for i, v3 := range v2 {
					switch v4 := v3.(type) {
					case string:
						l[i] = &v4
					case float64:
						strValue := strconv.FormatFloat(v4, 'f', -1, 64)
						l[i] = &strValue
					case int:
						strValue := strconv.Itoa(v4)
						l[i] = &strValue
					case int32:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					case int64:
						strValue := strconv.Itoa(int(v4))
						l[i] = &strValue
					default:
					}
				}
				p.Keys = l
			}
		}
	}
	return nil
}

func NewUnityAdFromJson(data string) UnityAd {
	req := UnityAd{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewUnityAdFromDict(data map[string]interface{}) UnityAd {
	return UnityAd{
		Keys: core.CastStrings(core.CastArray(data["keys"])),
	}
}

func (p UnityAd) ToDict() map[string]interface{} {

	var keys []interface{}
	if p.Keys != nil {
		keys = core.CastStringsFromDict(
			p.Keys,
		)
	}
	return map[string]interface{}{
		"keys": keys,
	}
}

func (p UnityAd) Pointer() *UnityAd {
	return &p
}

func CastUnityAds(data []interface{}) []UnityAd {
	v := make([]UnityAd, 0)
	for _, d := range data {
		v = append(v, NewUnityAdFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastUnityAdsFromDict(data []UnityAd) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type AppLovinMax struct {
	AllowAdUnitId *string `json:"allowAdUnitId"`
	EventKey      *string `json:"eventKey"`
}

func (p *AppLovinMax) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = AppLovinMax{}
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
		*p = AppLovinMax{}
	} else {
		*p = AppLovinMax{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["allowAdUnitId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.AllowAdUnitId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.AllowAdUnitId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.AllowAdUnitId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.AllowAdUnitId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.AllowAdUnitId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.AllowAdUnitId)
				}
			}
		}
		if v, ok := d["eventKey"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.EventKey = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.EventKey = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.EventKey = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.EventKey = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.EventKey = &strValue
				default:
					_ = json.Unmarshal(*v, &p.EventKey)
				}
			}
		}
	}
	return nil
}

func NewAppLovinMaxFromJson(data string) AppLovinMax {
	req := AppLovinMax{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewAppLovinMaxFromDict(data map[string]interface{}) AppLovinMax {
	return AppLovinMax{
		AllowAdUnitId: core.CastString(data["allowAdUnitId"]),
		EventKey:      core.CastString(data["eventKey"]),
	}
}

func (p AppLovinMax) ToDict() map[string]interface{} {

	var allowAdUnitId *string
	if p.AllowAdUnitId != nil {
		allowAdUnitId = p.AllowAdUnitId
	}
	var eventKey *string
	if p.EventKey != nil {
		eventKey = p.EventKey
	}
	return map[string]interface{}{
		"allowAdUnitId": allowAdUnitId,
		"eventKey":      eventKey,
	}
}

func (p AppLovinMax) Pointer() *AppLovinMax {
	return &p
}

func CastAppLovinMaxes(data []interface{}) []AppLovinMax {
	v := make([]AppLovinMax, 0)
	for _, d := range data {
		v = append(v, NewAppLovinMaxFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastAppLovinMaxesFromDict(data []AppLovinMax) []interface{} {
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

type NotificationSetting struct {
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func (p *NotificationSetting) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) == 0 {
		*p = NotificationSetting{}
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
		*p = NotificationSetting{}
	} else {
		*p = NotificationSetting{}
		d := map[string]*json.RawMessage{}
		if err := json.Unmarshal([]byte(str), &d); err != nil {
			return err
		}
		if v, ok := d["gatewayNamespaceId"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.GatewayNamespaceId = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.GatewayNamespaceId = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.GatewayNamespaceId = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.GatewayNamespaceId = &strValue
				default:
					_ = json.Unmarshal(*v, &p.GatewayNamespaceId)
				}
			}
		}
		if v, ok := d["enableTransferMobileNotification"]; ok && v != nil {
			_ = json.Unmarshal(*v, &p.EnableTransferMobileNotification)
		}
		if v, ok := d["sound"]; ok && v != nil {
			var temp interface{}
			if err := json.Unmarshal(*v, &temp); err == nil {
				switch v2 := temp.(type) {
				case string:
					p.Sound = &v2
				case float64:
					strValue := strconv.FormatFloat(v2, 'f', -1, 64)
					p.Sound = &strValue
				case int:
					strValue := strconv.Itoa(v2)
					p.Sound = &strValue
				case int32:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				case int64:
					strValue := strconv.Itoa(int(v2))
					p.Sound = &strValue
				default:
					_ = json.Unmarshal(*v, &p.Sound)
				}
			}
		}
	}
	return nil
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	req := NotificationSetting{}
	_ = json.Unmarshal([]byte(data), &req)
	return req
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {

	var gatewayNamespaceId *string
	if p.GatewayNamespaceId != nil {
		gatewayNamespaceId = p.GatewayNamespaceId
	}
	var enableTransferMobileNotification *bool
	if p.EnableTransferMobileNotification != nil {
		enableTransferMobileNotification = p.EnableTransferMobileNotification
	}
	var sound *string
	if p.Sound != nil {
		sound = p.Sound
	}
	return map[string]interface{}{
		"gatewayNamespaceId":               gatewayNamespaceId,
		"enableTransferMobileNotification": enableTransferMobileNotification,
		"sound":                            sound,
	}
}

func (p NotificationSetting) Pointer() *NotificationSetting {
	return &p
}

func CastNotificationSettings(data []interface{}) []NotificationSetting {
	v := make([]NotificationSetting, 0)
	for _, d := range data {
		v = append(v, NewNotificationSettingFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationSettingsFromDict(data []NotificationSetting) []interface{} {
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
