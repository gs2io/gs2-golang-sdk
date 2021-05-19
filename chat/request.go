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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *string	`json:"name"`
	Description *string	`json:"description"`
	AllowCreateRoom *bool	`json:"allowCreateRoom"`
	PostMessageScript *ScriptSetting	`json:"postMessageScript"`
	CreateRoomScript *ScriptSetting	`json:"createRoomScript"`
	DeleteRoomScript *ScriptSetting	`json:"deleteRoomScript"`
	SubscribeRoomScript *ScriptSetting	`json:"subscribeRoomScript"`
	UnsubscribeRoomScript *ScriptSetting	`json:"unsubscribeRoomScript"`
	PostNotification *NotificationSetting	`json:"postNotification"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Description *string	`json:"description"`
	AllowCreateRoom *bool	`json:"allowCreateRoom"`
	PostMessageScript *ScriptSetting	`json:"postMessageScript"`
	CreateRoomScript *ScriptSetting	`json:"createRoomScript"`
	DeleteRoomScript *ScriptSetting	`json:"deleteRoomScript"`
	SubscribeRoomScript *ScriptSetting	`json:"subscribeRoomScript"`
	UnsubscribeRoomScript *ScriptSetting	`json:"unsubscribeRoomScript"`
	PostNotification *NotificationSetting	`json:"postNotification"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
}

type DescribeRoomsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	Metadata *string	`json:"metadata"`
	Password *string	`json:"password"`
	WhiteListUserIds []string	`json:"whiteListUserIds"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CreateRoomFromBackendRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	Name *string	`json:"name"`
	UserId *string	`json:"userId"`
	Metadata *string	`json:"metadata"`
	Password *string	`json:"password"`
	WhiteListUserIds []string	`json:"whiteListUserIds"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
}

type UpdateRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	Metadata *string	`json:"metadata"`
	Password *string	`json:"password"`
	WhiteListUserIds []string	`json:"whiteListUserIds"`
}

type DeleteRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteRoomFromBackendRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DescribeMessagesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	Password *string	`json:"password"`
	StartAt *int64	`json:"startAt"`
	Limit *int64	`json:"limit"`
}

type PostRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	Category *int32	`json:"category"`
	Metadata *string	`json:"metadata"`
	Password *string	`json:"password"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type PostByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	Category *int32	`json:"category"`
	Metadata *string	`json:"metadata"`
	Password *string	`json:"password"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetMessageRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	MessageName *string	`json:"messageName"`
}

type DeleteMessageRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	MessageName *string	`json:"messageName"`
}

type DescribeSubscribesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeSubscribesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	UserId *string	`json:"userId"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type DescribeSubscribesByRoomNameRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	PageToken *string	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type SubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	NotificationTypes []NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type SubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	NotificationTypes []NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type GetSubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetSubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type UpdateNotificationTypeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	NotificationTypes []NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type UpdateNotificationTypeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	NotificationTypes []NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}

type UnsubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type UnsubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *string	`json:"namespaceName"`
	RoomName *string	`json:"roomName"`
	UserId *string	`json:"userId"`
	XGs2DuplicationAvoider *string	`json:"xGs2DuplicationAvoider"`
}
