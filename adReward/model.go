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

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId             *string              `json:"namespaceId"`
	Name                    *string              `json:"name"`
	Description             *string              `json:"description"`
	Admob                   *AdMob               `json:"admob"`
	UnityAd                 *UnityAd             `json:"unityAd"`
	AppLovinMaxes           []AppLovinMax        `json:"appLovinMaxes"`
	ChangePointNotification *NotificationSetting `json:"changePointNotification"`
	LogSetting              *LogSetting          `json:"logSetting"`
	CreatedAt               *int64               `json:"createdAt"`
	UpdatedAt               *int64               `json:"updatedAt"`
	Revision                *int64               `json:"revision"`
}

func NewNamespaceFromJson(data string) Namespace {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNamespaceFromDict(dict)
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:             core.CastString(data["namespaceId"]),
		Name:                    core.CastString(data["name"]),
		Description:             core.CastString(data["description"]),
		Admob:                   NewAdMobFromDict(core.CastMap(data["admob"])).Pointer(),
		UnityAd:                 NewUnityAdFromDict(core.CastMap(data["unityAd"])).Pointer(),
		AppLovinMaxes:           CastAppLovinMaxes(core.CastArray(data["appLovinMaxes"])),
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

func NewPointFromJson(data string) Point {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPointFromDict(dict)
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

func NewAdMobFromJson(data string) AdMob {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAdMobFromDict(dict)
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

func NewUnityAdFromJson(data string) UnityAd {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnityAdFromDict(dict)
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

func NewAppLovinMaxFromJson(data string) AppLovinMax {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAppLovinMaxFromDict(dict)
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

type NotificationSetting struct {
	GatewayNamespaceId               *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool   `json:"enableTransferMobileNotification"`
	Sound                            *string `json:"sound"`
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewNotificationSettingFromDict(dict)
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
