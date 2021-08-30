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

package chat

import "github.com/gs2io/gs2-golang-sdk/core"

type Namespace struct {
	NamespaceId           *string              `json:"namespaceId"`
	Name                  *string              `json:"name"`
	Description           *string              `json:"description"`
	AllowCreateRoom       *bool                `json:"allowCreateRoom"`
	PostMessageScript     *ScriptSetting       `json:"postMessageScript"`
	CreateRoomScript      *ScriptSetting       `json:"createRoomScript"`
	DeleteRoomScript      *ScriptSetting       `json:"deleteRoomScript"`
	SubscribeRoomScript   *ScriptSetting       `json:"subscribeRoomScript"`
	UnsubscribeRoomScript *ScriptSetting       `json:"unsubscribeRoomScript"`
	PostNotification      *NotificationSetting `json:"postNotification"`
	LogSetting            *LogSetting          `json:"logSetting"`
	CreatedAt             *int64               `json:"createdAt"`
	UpdatedAt             *int64               `json:"updatedAt"`
}

func NewNamespaceFromDict(data map[string]interface{}) Namespace {
	return Namespace{
		NamespaceId:           core.CastString(data["namespaceId"]),
		Name:                  core.CastString(data["name"]),
		Description:           core.CastString(data["description"]),
		AllowCreateRoom:       core.CastBool(data["allowCreateRoom"]),
		PostMessageScript:     NewScriptSettingFromDict(core.CastMap(data["postMessageScript"])).Pointer(),
		CreateRoomScript:      NewScriptSettingFromDict(core.CastMap(data["createRoomScript"])).Pointer(),
		DeleteRoomScript:      NewScriptSettingFromDict(core.CastMap(data["deleteRoomScript"])).Pointer(),
		SubscribeRoomScript:   NewScriptSettingFromDict(core.CastMap(data["subscribeRoomScript"])).Pointer(),
		UnsubscribeRoomScript: NewScriptSettingFromDict(core.CastMap(data["unsubscribeRoomScript"])).Pointer(),
		PostNotification:      NewNotificationSettingFromDict(core.CastMap(data["postNotification"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		CreatedAt:             core.CastInt64(data["createdAt"]),
		UpdatedAt:             core.CastInt64(data["updatedAt"]),
	}
}

func (p Namespace) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceId":           p.NamespaceId,
		"name":                  p.Name,
		"description":           p.Description,
		"allowCreateRoom":       p.AllowCreateRoom,
		"postMessageScript":     p.PostMessageScript.ToDict(),
		"createRoomScript":      p.CreateRoomScript.ToDict(),
		"deleteRoomScript":      p.DeleteRoomScript.ToDict(),
		"subscribeRoomScript":   p.SubscribeRoomScript.ToDict(),
		"unsubscribeRoomScript": p.UnsubscribeRoomScript.ToDict(),
		"postNotification":      p.PostNotification.ToDict(),
		"logSetting":            p.LogSetting.ToDict(),
		"createdAt":             p.CreatedAt,
		"updatedAt":             p.UpdatedAt,
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
	RoomId           *string  `json:"roomId"`
	Name             *string  `json:"name"`
	UserId           *string  `json:"userId"`
	Metadata         *string  `json:"metadata"`
	Password         *string  `json:"password"`
	WhiteListUserIds []string `json:"whiteListUserIds"`
	CreatedAt        *int64   `json:"createdAt"`
	UpdatedAt        *int64   `json:"updatedAt"`
}

func NewRoomFromDict(data map[string]interface{}) Room {
	return Room{
		RoomId:           core.CastString(data["roomId"]),
		Name:             core.CastString(data["name"]),
		UserId:           core.CastString(data["userId"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
		CreatedAt:        core.CastInt64(data["createdAt"]),
		UpdatedAt:        core.CastInt64(data["updatedAt"]),
	}
}

func (p Room) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"roomId":   p.RoomId,
		"name":     p.Name,
		"userId":   p.UserId,
		"metadata": p.Metadata,
		"password": p.Password,
		"whiteListUserIds": core.CastStringsFromDict(
			p.WhiteListUserIds,
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

type Message struct {
	MessageId *string `json:"messageId"`
	RoomName  *string `json:"roomName"`
	Name      *string `json:"name"`
	UserId    *string `json:"userId"`
	Category  *int32  `json:"category"`
	Metadata  *string `json:"metadata"`
	CreatedAt *int64  `json:"createdAt"`
}

func NewMessageFromDict(data map[string]interface{}) Message {
	return Message{
		MessageId: core.CastString(data["messageId"]),
		RoomName:  core.CastString(data["roomName"]),
		Name:      core.CastString(data["name"]),
		UserId:    core.CastString(data["userId"]),
		Category:  core.CastInt32(data["category"]),
		Metadata:  core.CastString(data["metadata"]),
		CreatedAt: core.CastInt64(data["createdAt"]),
	}
}

func (p Message) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"messageId": p.MessageId,
		"roomName":  p.RoomName,
		"name":      p.Name,
		"userId":    p.UserId,
		"category":  p.Category,
		"metadata":  p.Metadata,
		"createdAt": p.CreatedAt,
	}
}

func (p Message) Pointer() *Message {
	return &p
}

func CastMessages(data []interface{}) []Message {
	v := make([]Message, 0)
	for _, d := range data {
		v = append(v, NewMessageFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastMessagesFromDict(data []Message) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type Subscribe struct {
	SubscribeId       *string            `json:"subscribeId"`
	UserId            *string            `json:"userId"`
	RoomName          *string            `json:"roomName"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
	CreatedAt         *int64             `json:"createdAt"`
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
	return Subscribe{
		SubscribeId:       core.CastString(data["subscribeId"]),
		UserId:            core.CastString(data["userId"]),
		RoomName:          core.CastString(data["roomName"]),
		NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
		CreatedAt:         core.CastInt64(data["createdAt"]),
	}
}

func (p Subscribe) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"subscribeId": p.SubscribeId,
		"userId":      p.UserId,
		"roomName":    p.RoomName,
		"notificationTypes": CastNotificationTypesFromDict(
			p.NotificationTypes,
		),
		"createdAt": p.CreatedAt,
	}
}

func (p Subscribe) Pointer() *Subscribe {
	return &p
}

func CastSubscribes(data []interface{}) []Subscribe {
	v := make([]Subscribe, 0)
	for _, d := range data {
		v = append(v, NewSubscribeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastSubscribesFromDict(data []Subscribe) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type ResponseCache struct {
	Region          *string `json:"region"`
	ResponseCacheId *string `json:"responseCacheId"`
	RequestHash     *string `json:"requestHash"`
	Result          *string `json:"result"`
}

func NewResponseCacheFromDict(data map[string]interface{}) ResponseCache {
	return ResponseCache{
		Region:          core.CastString(data["region"]),
		ResponseCacheId: core.CastString(data["responseCacheId"]),
		RequestHash:     core.CastString(data["requestHash"]),
		Result:          core.CastString(data["result"]),
	}
}

func (p ResponseCache) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"region":          p.Region,
		"responseCacheId": p.ResponseCacheId,
		"requestHash":     p.RequestHash,
		"result":          p.Result,
	}
}

func (p ResponseCache) Pointer() *ResponseCache {
	return &p
}

func CastResponseCaches(data []interface{}) []ResponseCache {
	v := make([]ResponseCache, 0)
	for _, d := range data {
		v = append(v, NewResponseCacheFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastResponseCachesFromDict(data []ResponseCache) []interface{} {
	v := make([]interface{}, 0)
	for _, d := range data {
		v = append(v, d.ToDict())
	}
	return v
}

type NotificationType struct {
	Category                             *int32 `json:"category"`
	EnableTransferMobilePushNotification *bool  `json:"enableTransferMobilePushNotification"`
}

func NewNotificationTypeFromDict(data map[string]interface{}) NotificationType {
	return NotificationType{
		Category:                             core.CastInt32(data["category"]),
		EnableTransferMobilePushNotification: core.CastBool(data["enableTransferMobilePushNotification"]),
	}
}

func (p NotificationType) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"category":                             p.Category,
		"enableTransferMobilePushNotification": p.EnableTransferMobilePushNotification,
	}
}

func (p NotificationType) Pointer() *NotificationType {
	return &p
}

func CastNotificationTypes(data []interface{}) []NotificationType {
	v := make([]NotificationType, 0)
	for _, d := range data {
		v = append(v, NewNotificationTypeFromDict(d.(map[string]interface{})))
	}
	return v
}

func CastNotificationTypesFromDict(data []NotificationType) []interface{} {
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
