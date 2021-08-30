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

type DescribeNamespacesRequest struct {
	RequestId    *string `json:"requestId"`
	ContextStack *string `json:"contextStack"`
	PageToken    *string `json:"pageToken"`
	Limit        *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
	return DescribeNamespacesRequest{
		PageToken: core.CastString(data["pageToken"]),
		Limit:     core.CastInt32(data["limit"]),
	}
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"pageToken": p.PageToken,
		"limit":     p.Limit,
	}
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
	return &p
}

type CreateNamespaceRequest struct {
	RequestId             *string              `json:"requestId"`
	ContextStack          *string              `json:"contextStack"`
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
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
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
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
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
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
	return GetNamespaceStatusRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
	return &p
}

type GetNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
	return GetNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
	return &p
}

type UpdateNamespaceRequest struct {
	RequestId             *string              `json:"requestId"`
	ContextStack          *string              `json:"contextStack"`
	NamespaceName         *string              `json:"namespaceName"`
	Description           *string              `json:"description"`
	AllowCreateRoom       *bool                `json:"allowCreateRoom"`
	PostMessageScript     *ScriptSetting       `json:"postMessageScript"`
	CreateRoomScript      *ScriptSetting       `json:"createRoomScript"`
	DeleteRoomScript      *ScriptSetting       `json:"deleteRoomScript"`
	SubscribeRoomScript   *ScriptSetting       `json:"subscribeRoomScript"`
	UnsubscribeRoomScript *ScriptSetting       `json:"unsubscribeRoomScript"`
	PostNotification      *NotificationSetting `json:"postNotification"`
	LogSetting            *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:         core.CastString(data["namespaceName"]),
		Description:           core.CastString(data["description"]),
		AllowCreateRoom:       core.CastBool(data["allowCreateRoom"]),
		PostMessageScript:     NewScriptSettingFromDict(core.CastMap(data["postMessageScript"])).Pointer(),
		CreateRoomScript:      NewScriptSettingFromDict(core.CastMap(data["createRoomScript"])).Pointer(),
		DeleteRoomScript:      NewScriptSettingFromDict(core.CastMap(data["deleteRoomScript"])).Pointer(),
		SubscribeRoomScript:   NewScriptSettingFromDict(core.CastMap(data["subscribeRoomScript"])).Pointer(),
		UnsubscribeRoomScript: NewScriptSettingFromDict(core.CastMap(data["unsubscribeRoomScript"])).Pointer(),
		PostNotification:      NewNotificationSettingFromDict(core.CastMap(data["postNotification"])).Pointer(),
		LogSetting:            NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":         p.NamespaceName,
		"description":           p.Description,
		"allowCreateRoom":       p.AllowCreateRoom,
		"postMessageScript":     p.PostMessageScript.ToDict(),
		"createRoomScript":      p.CreateRoomScript.ToDict(),
		"deleteRoomScript":      p.DeleteRoomScript.ToDict(),
		"subscribeRoomScript":   p.SubscribeRoomScript.ToDict(),
		"unsubscribeRoomScript": p.UnsubscribeRoomScript.ToDict(),
		"postNotification":      p.PostNotification.ToDict(),
		"logSetting":            p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
	return DeleteNamespaceRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
	return &p
}

type DescribeRoomsRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeRoomsRequestFromDict(data map[string]interface{}) DescribeRoomsRequest {
	return DescribeRoomsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRoomsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRoomsRequest) Pointer() *DescribeRoomsRequest {
	return &p
}

type CreateRoomRequest struct {
	RequestId        *string  `json:"requestId"`
	ContextStack     *string  `json:"contextStack"`
	NamespaceName    *string  `json:"namespaceName"`
	AccessToken      *string  `json:"accessToken"`
	Name             *string  `json:"name"`
	Metadata         *string  `json:"metadata"`
	Password         *string  `json:"password"`
	WhiteListUserIds []string `json:"whiteListUserIds"`
}

func NewCreateRoomRequestFromDict(data map[string]interface{}) CreateRoomRequest {
	return CreateRoomRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		AccessToken:      core.CastString(data["accessToken"]),
		Name:             core.CastString(data["name"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
	}
}

func (p CreateRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"name":          p.Name,
		"metadata":      p.Metadata,
		"password":      p.Password,
		"whiteListUserIds": core.CastStringsFromDict(
			p.WhiteListUserIds,
		),
	}
}

func (p CreateRoomRequest) Pointer() *CreateRoomRequest {
	return &p
}

type CreateRoomFromBackendRequest struct {
	RequestId        *string  `json:"requestId"`
	ContextStack     *string  `json:"contextStack"`
	NamespaceName    *string  `json:"namespaceName"`
	Name             *string  `json:"name"`
	UserId           *string  `json:"userId"`
	Metadata         *string  `json:"metadata"`
	Password         *string  `json:"password"`
	WhiteListUserIds []string `json:"whiteListUserIds"`
}

func NewCreateRoomFromBackendRequestFromDict(data map[string]interface{}) CreateRoomFromBackendRequest {
	return CreateRoomFromBackendRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		Name:             core.CastString(data["name"]),
		UserId:           core.CastString(data["userId"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
	}
}

func (p CreateRoomFromBackendRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"userId":        p.UserId,
		"metadata":      p.Metadata,
		"password":      p.Password,
		"whiteListUserIds": core.CastStringsFromDict(
			p.WhiteListUserIds,
		),
	}
}

func (p CreateRoomFromBackendRequest) Pointer() *CreateRoomFromBackendRequest {
	return &p
}

type GetRoomRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
}

func NewGetRoomRequestFromDict(data map[string]interface{}) GetRoomRequest {
	return GetRoomRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
	}
}

func (p GetRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
	}
}

func (p GetRoomRequest) Pointer() *GetRoomRequest {
	return &p
}

type UpdateRoomRequest struct {
	RequestId        *string  `json:"requestId"`
	ContextStack     *string  `json:"contextStack"`
	NamespaceName    *string  `json:"namespaceName"`
	RoomName         *string  `json:"roomName"`
	Metadata         *string  `json:"metadata"`
	Password         *string  `json:"password"`
	WhiteListUserIds []string `json:"whiteListUserIds"`
}

func NewUpdateRoomRequestFromDict(data map[string]interface{}) UpdateRoomRequest {
	return UpdateRoomRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		RoomName:         core.CastString(data["roomName"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
	}
}

func (p UpdateRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"metadata":      p.Metadata,
		"password":      p.Password,
		"whiteListUserIds": core.CastStringsFromDict(
			p.WhiteListUserIds,
		),
	}
}

func (p UpdateRoomRequest) Pointer() *UpdateRoomRequest {
	return &p
}

type DeleteRoomRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	AccessToken   *string `json:"accessToken"`
}

func NewDeleteRoomRequestFromDict(data map[string]interface{}) DeleteRoomRequest {
	return DeleteRoomRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p DeleteRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
	}
}

func (p DeleteRoomRequest) Pointer() *DeleteRoomRequest {
	return &p
}

type DeleteRoomFromBackendRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	UserId        *string `json:"userId"`
}

func NewDeleteRoomFromBackendRequestFromDict(data map[string]interface{}) DeleteRoomFromBackendRequest {
	return DeleteRoomFromBackendRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p DeleteRoomFromBackendRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
	}
}

func (p DeleteRoomFromBackendRequest) Pointer() *DeleteRoomFromBackendRequest {
	return &p
}

type DescribeMessagesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	Password      *string `json:"password"`
	StartAt       *int64  `json:"startAt"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeMessagesRequestFromDict(data map[string]interface{}) DescribeMessagesRequest {
	return DescribeMessagesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		Password:      core.CastString(data["password"]),
		StartAt:       core.CastInt64(data["startAt"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMessagesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"password":      p.Password,
		"startAt":       p.StartAt,
		"limit":         p.Limit,
	}
}

func (p DescribeMessagesRequest) Pointer() *DescribeMessagesRequest {
	return &p
}

type PostRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	AccessToken   *string `json:"accessToken"`
	Category      *int32  `json:"category"`
	Metadata      *string `json:"metadata"`
	Password      *string `json:"password"`
}

func NewPostRequestFromDict(data map[string]interface{}) PostRequest {
	return PostRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		Category:      core.CastInt32(data["category"]),
		Metadata:      core.CastString(data["metadata"]),
		Password:      core.CastString(data["password"]),
	}
}

func (p PostRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
		"category":      p.Category,
		"metadata":      p.Metadata,
		"password":      p.Password,
	}
}

func (p PostRequest) Pointer() *PostRequest {
	return &p
}

type PostByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	UserId        *string `json:"userId"`
	Category      *int32  `json:"category"`
	Metadata      *string `json:"metadata"`
	Password      *string `json:"password"`
}

func NewPostByUserIdRequestFromDict(data map[string]interface{}) PostByUserIdRequest {
	return PostByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		UserId:        core.CastString(data["userId"]),
		Category:      core.CastInt32(data["category"]),
		Metadata:      core.CastString(data["metadata"]),
		Password:      core.CastString(data["password"]),
	}
}

func (p PostByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
		"category":      p.Category,
		"metadata":      p.Metadata,
		"password":      p.Password,
	}
}

func (p PostByUserIdRequest) Pointer() *PostByUserIdRequest {
	return &p
}

type GetMessageRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	MessageName   *string `json:"messageName"`
}

func NewGetMessageRequestFromDict(data map[string]interface{}) GetMessageRequest {
	return GetMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p GetMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"messageName":   p.MessageName,
	}
}

func (p GetMessageRequest) Pointer() *GetMessageRequest {
	return &p
}

type DeleteMessageRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	MessageName   *string `json:"messageName"`
}

func NewDeleteMessageRequestFromDict(data map[string]interface{}) DeleteMessageRequest {
	return DeleteMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p DeleteMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"messageName":   p.MessageName,
	}
}

func (p DeleteMessageRequest) Pointer() *DeleteMessageRequest {
	return &p
}

type DescribeSubscribesRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	AccessToken   *string `json:"accessToken"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSubscribesRequestFromDict(data map[string]interface{}) DescribeSubscribesRequest {
	return DescribeSubscribesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSubscribesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribesRequest) Pointer() *DescribeSubscribesRequest {
	return &p
}

type DescribeSubscribesByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	UserId        *string `json:"userId"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSubscribesByUserIdRequestFromDict(data map[string]interface{}) DescribeSubscribesByUserIdRequest {
	return DescribeSubscribesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSubscribesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribesByUserIdRequest) Pointer() *DescribeSubscribesByUserIdRequest {
	return &p
}

type DescribeSubscribesByRoomNameRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	PageToken     *string `json:"pageToken"`
	Limit         *int32  `json:"limit"`
}

func NewDescribeSubscribesByRoomNameRequestFromDict(data map[string]interface{}) DescribeSubscribesByRoomNameRequest {
	return DescribeSubscribesByRoomNameRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeSubscribesByRoomNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeSubscribesByRoomNameRequest) Pointer() *DescribeSubscribesByRoomNameRequest {
	return &p
}

type SubscribeRequest struct {
	RequestId         *string            `json:"requestId"`
	ContextStack      *string            `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	RoomName          *string            `json:"roomName"`
	AccessToken       *string            `json:"accessToken"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
}

func NewSubscribeRequestFromDict(data map[string]interface{}) SubscribeRequest {
	return SubscribeRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		RoomName:          core.CastString(data["roomName"]),
		AccessToken:       core.CastString(data["accessToken"]),
		NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
	}
}

func (p SubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
		"notificationTypes": CastNotificationTypesFromDict(
			p.NotificationTypes,
		),
	}
}

func (p SubscribeRequest) Pointer() *SubscribeRequest {
	return &p
}

type SubscribeByUserIdRequest struct {
	RequestId         *string            `json:"requestId"`
	ContextStack      *string            `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	RoomName          *string            `json:"roomName"`
	UserId            *string            `json:"userId"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
}

func NewSubscribeByUserIdRequestFromDict(data map[string]interface{}) SubscribeByUserIdRequest {
	return SubscribeByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		RoomName:          core.CastString(data["roomName"]),
		UserId:            core.CastString(data["userId"]),
		NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
	}
}

func (p SubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
		"notificationTypes": CastNotificationTypesFromDict(
			p.NotificationTypes,
		),
	}
}

func (p SubscribeByUserIdRequest) Pointer() *SubscribeByUserIdRequest {
	return &p
}

type GetSubscribeRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	AccessToken   *string `json:"accessToken"`
}

func NewGetSubscribeRequestFromDict(data map[string]interface{}) GetSubscribeRequest {
	return GetSubscribeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetSubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetSubscribeRequest) Pointer() *GetSubscribeRequest {
	return &p
}

type GetSubscribeByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	UserId        *string `json:"userId"`
}

func NewGetSubscribeByUserIdRequestFromDict(data map[string]interface{}) GetSubscribeByUserIdRequest {
	return GetSubscribeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetSubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
	}
}

func (p GetSubscribeByUserIdRequest) Pointer() *GetSubscribeByUserIdRequest {
	return &p
}

type UpdateNotificationTypeRequest struct {
	RequestId         *string            `json:"requestId"`
	ContextStack      *string            `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	RoomName          *string            `json:"roomName"`
	AccessToken       *string            `json:"accessToken"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
}

func NewUpdateNotificationTypeRequestFromDict(data map[string]interface{}) UpdateNotificationTypeRequest {
	return UpdateNotificationTypeRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		RoomName:          core.CastString(data["roomName"]),
		AccessToken:       core.CastString(data["accessToken"]),
		NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
	}
}

func (p UpdateNotificationTypeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
		"notificationTypes": CastNotificationTypesFromDict(
			p.NotificationTypes,
		),
	}
}

func (p UpdateNotificationTypeRequest) Pointer() *UpdateNotificationTypeRequest {
	return &p
}

type UpdateNotificationTypeByUserIdRequest struct {
	RequestId         *string            `json:"requestId"`
	ContextStack      *string            `json:"contextStack"`
	NamespaceName     *string            `json:"namespaceName"`
	RoomName          *string            `json:"roomName"`
	UserId            *string            `json:"userId"`
	NotificationTypes []NotificationType `json:"notificationTypes"`
}

func NewUpdateNotificationTypeByUserIdRequestFromDict(data map[string]interface{}) UpdateNotificationTypeByUserIdRequest {
	return UpdateNotificationTypeByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		RoomName:          core.CastString(data["roomName"]),
		UserId:            core.CastString(data["userId"]),
		NotificationTypes: CastNotificationTypes(core.CastArray(data["notificationTypes"])),
	}
}

func (p UpdateNotificationTypeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
		"notificationTypes": CastNotificationTypesFromDict(
			p.NotificationTypes,
		),
	}
}

func (p UpdateNotificationTypeByUserIdRequest) Pointer() *UpdateNotificationTypeByUserIdRequest {
	return &p
}

type UnsubscribeRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	AccessToken   *string `json:"accessToken"`
}

func NewUnsubscribeRequestFromDict(data map[string]interface{}) UnsubscribeRequest {
	return UnsubscribeRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p UnsubscribeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"accessToken":   p.AccessToken,
	}
}

func (p UnsubscribeRequest) Pointer() *UnsubscribeRequest {
	return &p
}

type UnsubscribeByUserIdRequest struct {
	RequestId     *string `json:"requestId"`
	ContextStack  *string `json:"contextStack"`
	NamespaceName *string `json:"namespaceName"`
	RoomName      *string `json:"roomName"`
	UserId        *string `json:"userId"`
}

func NewUnsubscribeByUserIdRequestFromDict(data map[string]interface{}) UnsubscribeByUserIdRequest {
	return UnsubscribeByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p UnsubscribeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
	}
}

func (p UnsubscribeByUserIdRequest) Pointer() *UnsubscribeByUserIdRequest {
	return &p
}
