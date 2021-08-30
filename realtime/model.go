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

package realtime

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId        *string              `json:"namespaceId"`
	Name               *string              `json:"name"`
	Description        *string              `json:"description"`
	ServerType         *string              `json:"serverType"`
	ServerSpec         *string              `json:"serverSpec"`
	CreateNotification *NotificationSetting `json:"createNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
	CreatedAt          *int64               `json:"createdAt"`
	UpdatedAt          *int64               `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:        core.CastString(data["namespaceId"]),
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		ServerType:         core.CastString(data["serverType"]),
		ServerSpec:         core.CastString(data["serverSpec"]),
		CreateNotification: NewNotificationSettingFromDict(core.CastMap(data["createNotification"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:          core.CastInt64(data["createdAt"]),
		UpdatedAt:          core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":        p.NamespaceId,
		"name":               p.Name,
		"description":        p.Description,
		"serverType":         p.ServerType,
		"serverSpec":         p.ServerSpec,
		"createNotification": p.CreateNotification.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
		"createdAt":          p.CreatedAt,
		"updatedAt":          p.UpdatedAt,
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

type Room struct {
	RoomId              *string  `json:"roomId"`
	Name                *string  `json:"name"`
	IpAddress           *string  `json:"ipAddress"`
	Port                *int32   `json:"port"`
	EncryptionKey       *string  `json:"encryptionKey"`
	NotificationUserIds []string `json:"notificationUserIds"`
	CreatedAt           *int64   `json:"createdAt"`
	UpdatedAt           *int64   `json:"updatedAt"`
}

func NewRoomFromDict(data map[string]interface{}) Room {
	return Room{
		RoomId:              core.CastString(data["roomId"]),
		Name:                core.CastString(data["name"]),
		IpAddress:           core.CastString(data["ipAddress"]),
		Port:                core.CastInt32(data["port"]),
		EncryptionKey:       core.CastString(data["encryptionKey"]),
		NotificationUserIds: core.CastStrings(core.CastArray(data["notificationUserIds"])),
		CreatedAt:           core.CastInt64(data["createdAt"]),
		UpdatedAt:           core.CastInt64(data["updatedAt"]),
	}
}

func (p Room) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"roomId":        p.RoomId,
		"name":          p.Name,
		"ipAddress":     p.IpAddress,
		"port":          p.Port,
		"encryptionKey": p.EncryptionKey,
		"notificationUserIds": core.CastStringsFromDict(
			p.NotificationUserIds,
		),
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}
}

func (p Room) Pointer() *Room {
	return &p
}

func CastRooms(data []interface{}) []Room {
	v := make([]Room, 0)
	for _, d := range data {
		v = append(v, NewRoomFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastRoomsFromDict(data []Room) []interface{} {
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

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
	return NotificationSetting{
		GatewayNamespaceId:               core.CastString(data["gatewayNamespaceId"]),
		EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
		Sound:                            core.CastString(data["sound"]),
	}
}

func (p NotificationSetting) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"gatewayNamespaceId":               p.GatewayNamespaceId,
		"enableTransferMobileNotification": p.EnableTransferMobileNotification,
		"sound":                            p.Sound,
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
