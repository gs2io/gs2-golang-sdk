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
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
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
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Description *core.String	`json:"description"`
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
	NamespaceName *core.String	`json:"namespaceName"`
}

type DescribeRoomsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Metadata *core.String	`json:"metadata"`
	Password *core.String	`json:"password"`
	WhiteListUserIds *[]core.String	`json:"whiteListUserIds"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CreateRoomFromBackendRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	UserId *core.String	`json:"userId"`
	Metadata *core.String	`json:"metadata"`
	Password *core.String	`json:"password"`
	WhiteListUserIds *[]core.String	`json:"whiteListUserIds"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
}

type UpdateRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	Metadata *core.String	`json:"metadata"`
	Password *core.String	`json:"password"`
	WhiteListUserIds *[]core.String	`json:"whiteListUserIds"`
}

type DeleteRoomRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DeleteRoomFromBackendRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DescribeMessagesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	Password *core.String	`json:"password"`
	StartAt *int64	`json:"startAt"`
	Limit *int64	`json:"limit"`
}

type PostRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	Category *int32	`json:"category"`
	Metadata *core.String	`json:"metadata"`
	Password *core.String	`json:"password"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type PostByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	Category *int32	`json:"category"`
	Metadata *core.String	`json:"metadata"`
	Password *core.String	`json:"password"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetMessageRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	MessageName *core.String	`json:"messageName"`
}

type DeleteMessageRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	MessageName *core.String	`json:"messageName"`
}

type DescribeSubscribesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeSubscribesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DescribeSubscribesByRoomNameRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type SubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	NotificationTypes *[]*NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type SubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	NotificationTypes *[]*NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetSubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetSubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type UpdateNotificationTypeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	NotificationTypes *[]*NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type UpdateNotificationTypeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	NotificationTypes *[]*NotificationType	`json:"notificationTypes"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type UnsubscribeRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type UnsubscribeByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RoomName *core.String	`json:"roomName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}
