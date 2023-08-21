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

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type Namespace struct {
	NamespaceId *string `json:"namespaceId"`
	Name *string `json:"name"`
	Description *string `json:"description"`
	AllowCreateRoom *bool `json:"allowCreateRoom"`
	PostMessageScript *ScriptSetting `json:"postMessageScript"`
	CreateRoomScript *ScriptSetting `json:"createRoomScript"`
	DeleteRoomScript *ScriptSetting `json:"deleteRoomScript"`
	SubscribeRoomScript *ScriptSetting `json:"subscribeRoomScript"`
	UnsubscribeRoomScript *ScriptSetting `json:"unsubscribeRoomScript"`
	PostNotification *NotificationSetting `json:"postNotification"`
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
        AllowCreateRoom: core.CastBool(data["allowCreateRoom"]),
        PostMessageScript: NewScriptSettingFromDict(core.CastMap(data["postMessageScript"])).Pointer(),
        CreateRoomScript: NewScriptSettingFromDict(core.CastMap(data["createRoomScript"])).Pointer(),
        DeleteRoomScript: NewScriptSettingFromDict(core.CastMap(data["deleteRoomScript"])).Pointer(),
        SubscribeRoomScript: NewScriptSettingFromDict(core.CastMap(data["subscribeRoomScript"])).Pointer(),
        UnsubscribeRoomScript: NewScriptSettingFromDict(core.CastMap(data["unsubscribeRoomScript"])).Pointer(),
        PostNotification: NewNotificationSettingFromDict(core.CastMap(data["postNotification"])).Pointer(),
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
    var allowCreateRoom *bool
    if p.AllowCreateRoom != nil {
        allowCreateRoom = p.AllowCreateRoom
    }
    var postMessageScript map[string]interface{}
    if p.PostMessageScript != nil {
        postMessageScript = p.PostMessageScript.ToDict()
    }
    var createRoomScript map[string]interface{}
    if p.CreateRoomScript != nil {
        createRoomScript = p.CreateRoomScript.ToDict()
    }
    var deleteRoomScript map[string]interface{}
    if p.DeleteRoomScript != nil {
        deleteRoomScript = p.DeleteRoomScript.ToDict()
    }
    var subscribeRoomScript map[string]interface{}
    if p.SubscribeRoomScript != nil {
        subscribeRoomScript = p.SubscribeRoomScript.ToDict()
    }
    var unsubscribeRoomScript map[string]interface{}
    if p.UnsubscribeRoomScript != nil {
        unsubscribeRoomScript = p.UnsubscribeRoomScript.ToDict()
    }
    var postNotification map[string]interface{}
    if p.PostNotification != nil {
        postNotification = p.PostNotification.ToDict()
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
        "allowCreateRoom": allowCreateRoom,
        "postMessageScript": postMessageScript,
        "createRoomScript": createRoomScript,
        "deleteRoomScript": deleteRoomScript,
        "subscribeRoomScript": subscribeRoomScript,
        "unsubscribeRoomScript": unsubscribeRoomScript,
        "postNotification": postNotification,
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

type Room struct {
	RoomId *string `json:"roomId"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	Metadata *string `json:"metadata"`
	Password *string `json:"password"`
	WhiteListUserIds []*string `json:"whiteListUserIds"`
	CreatedAt *int64 `json:"createdAt"`
	UpdatedAt *int64 `json:"updatedAt"`
	Revision *int64 `json:"revision"`
}

func NewRoomFromJson(data string) Room {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRoomFromDict(dict)
}

func NewRoomFromDict(data map[string]interface{}) Room {
    return Room {
        RoomId: core.CastString(data["roomId"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        Metadata: core.CastString(data["metadata"]),
        Password: core.CastString(data["password"]),
        WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        UpdatedAt: core.CastInt64(data["updatedAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Room) ToDict() map[string]interface{} {
    
    var roomId *string
    if p.RoomId != nil {
        roomId = p.RoomId
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var password *string
    if p.Password != nil {
        password = p.Password
    }
    var whiteListUserIds []interface{}
    if p.WhiteListUserIds != nil {
        whiteListUserIds = core.CastStringsFromDict(
            p.WhiteListUserIds,
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
        "roomId": roomId,
        "name": name,
        "userId": userId,
        "metadata": metadata,
        "password": password,
        "whiteListUserIds": whiteListUserIds,
        "createdAt": createdAt,
        "updatedAt": updatedAt,
        "revision": revision,
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
	RoomName *string `json:"roomName"`
	Name *string `json:"name"`
	UserId *string `json:"userId"`
	Category *int32 `json:"category"`
	Metadata *string `json:"metadata"`
	CreatedAt *int64 `json:"createdAt"`
	Revision *int64 `json:"revision"`
}

func NewMessageFromJson(data string) Message {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewMessageFromDict(dict)
}

func NewMessageFromDict(data map[string]interface{}) Message {
    return Message {
        MessageId: core.CastString(data["messageId"]),
        RoomName: core.CastString(data["roomName"]),
        Name: core.CastString(data["name"]),
        UserId: core.CastString(data["userId"]),
        Category: core.CastInt32(data["category"]),
        Metadata: core.CastString(data["metadata"]),
        CreatedAt: core.CastInt64(data["createdAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Message) ToDict() map[string]interface{} {
    
    var messageId *string
    if p.MessageId != nil {
        messageId = p.MessageId
    }
    var roomName *string
    if p.RoomName != nil {
        roomName = p.RoomName
    }
    var name *string
    if p.Name != nil {
        name = p.Name
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var category *int32
    if p.Category != nil {
        category = p.Category
    }
    var metadata *string
    if p.Metadata != nil {
        metadata = p.Metadata
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "messageId": messageId,
        "roomName": roomName,
        "name": name,
        "userId": userId,
        "category": category,
        "metadata": metadata,
        "createdAt": createdAt,
        "revision": revision,
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
	SubscribeId *string `json:"subscribeId"`
	UserId *string `json:"userId"`
	RoomName *string `json:"roomName"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
	CreatedAt *int64 `json:"createdAt"`
	Revision *int64 `json:"revision"`
}

func NewSubscribeFromJson(data string) Subscribe {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSubscribeFromDict(dict)
}

func NewSubscribeFromDict(data map[string]interface{}) Subscribe {
    return Subscribe {
        SubscribeId: core.CastString(data["subscribeId"]),
        UserId: core.CastString(data["userId"]),
        RoomName: core.CastString(data["roomName"]),
        NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
        CreatedAt: core.CastInt64(data["createdAt"]),
        Revision: core.CastInt64(data["revision"]),
    }
}

func (p Subscribe) ToDict() map[string]interface{} {
    
    var subscribeId *string
    if p.SubscribeId != nil {
        subscribeId = p.SubscribeId
    }
    var userId *string
    if p.UserId != nil {
        userId = p.UserId
    }
    var roomName *string
    if p.RoomName != nil {
        roomName = p.RoomName
    }
    var notificationTypes []interface{}
    if p.NotificationTypes != nil {
        notificationTypes = CastNotificationTypesFromDict(
            p.NotificationTypes,
        )
    }
    var createdAt *int64
    if p.CreatedAt != nil {
        createdAt = p.CreatedAt
    }
    var revision *int64
    if p.Revision != nil {
        revision = p.Revision
    }
    return map[string]interface{} {
        "subscribeId": subscribeId,
        "userId": userId,
        "roomName": roomName,
        "notificationTypes": notificationTypes,
        "createdAt": createdAt,
        "revision": revision,
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

type NotificationType struct {
	Category *int32 `json:"category"`
	EnableTransferMobilePushNotification *bool `json:"enableTransferMobilePushNotification"`
}

func NewNotificationTypeFromJson(data string) NotificationType {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNotificationTypeFromDict(dict)
}

func NewNotificationTypeFromDict(data map[string]interface{}) NotificationType {
    return NotificationType {
        Category: core.CastInt32(data["category"]),
        EnableTransferMobilePushNotification: core.CastBool(data["enableTransferMobilePushNotification"]),
    }
}

func (p NotificationType) ToDict() map[string]interface{} {
    
    var category *int32
    if p.Category != nil {
        category = p.Category
    }
    var enableTransferMobilePushNotification *bool
    if p.EnableTransferMobilePushNotification != nil {
        enableTransferMobilePushNotification = p.EnableTransferMobilePushNotification
    }
    return map[string]interface{} {
        "category": category,
        "enableTransferMobilePushNotification": enableTransferMobilePushNotification,
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

type NotificationSetting struct {
	GatewayNamespaceId *string `json:"gatewayNamespaceId"`
	EnableTransferMobileNotification *bool `json:"enableTransferMobileNotification"`
	Sound *string `json:"sound"`
}

func NewNotificationSettingFromJson(data string) NotificationSetting {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewNotificationSettingFromDict(dict)
}

func NewNotificationSettingFromDict(data map[string]interface{}) NotificationSetting {
    return NotificationSetting {
        GatewayNamespaceId: core.CastString(data["gatewayNamespaceId"]),
        EnableTransferMobileNotification: core.CastBool(data["enableTransferMobileNotification"]),
        Sound: core.CastString(data["sound"]),
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
    return map[string]interface{} {
        "gatewayNamespaceId": gatewayNamespaceId,
        "enableTransferMobileNotification": enableTransferMobileNotification,
        "sound": sound,
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