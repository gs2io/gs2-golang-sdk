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

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId          *core.RequestId      `json:"requestId"`
	ContextStack       *core.ContextStack   `json:"contextStack"`
	Name               *string              `json:"name"`
	Description        *string              `json:"description"`
	ServerType         *string              `json:"serverType"`
	ServerSpec         *string              `json:"serverSpec"`
	CreateNotification *NotificationSetting `json:"createNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

type GetNamespaceStatusRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateNamespaceRequest struct {
	RequestId          *core.RequestId      `json:"requestId"`
	ContextStack       *core.ContextStack   `json:"contextStack"`
	NamespaceName      *string              `json:"namespaceName"`
	Description        *string              `json:"description"`
	ServerType         *string              `json:"serverType"`
	ServerSpec         *string              `json:"serverSpec"`
	CreateNotification *NotificationSetting `json:"createNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeRoomsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type WantRoomRequest struct {
	RequestId           *core.RequestId    `json:"requestId"`
	ContextStack        *core.ContextStack `json:"contextStack"`
	NamespaceName       *string            `json:"namespaceName"`
	Name                *string            `json:"name"`
	NotificationUserIds *[]string          `json:"notificationUserIds"`
}

type GetRoomRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RoomName      *string            `json:"roomName"`
}

type DeleteRoomRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RoomName      *string            `json:"roomName"`
}
