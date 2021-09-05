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
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeNamespacesRequestFromDict(dict)
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
	RequestId          *string              `json:"requestId"`
	ContextStack       *string              `json:"contextStack"`
	DuplicationAvoider *string              `json:"duplicationAvoider"`
	Name               *string              `json:"name"`
	Description        *string              `json:"description"`
	ServerType         *string              `json:"serverType"`
	ServerSpec         *string              `json:"serverSpec"`
	CreateNotification *NotificationSetting `json:"createNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:               core.CastString(data["name"]),
		Description:        core.CastString(data["description"]),
		ServerType:         core.CastString(data["serverType"]),
		ServerSpec:         core.CastString(data["serverSpec"]),
		CreateNotification: NewNotificationSettingFromDict(core.CastMap(data["createNotification"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":               p.Name,
		"description":        p.Description,
		"serverType":         p.ServerType,
		"serverSpec":         p.ServerSpec,
		"createNotification": p.CreateNotification.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
	return &p
}

type GetNamespaceStatusRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceStatusRequestFromDict(dict)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetNamespaceRequestFromDict(dict)
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
	RequestId          *string              `json:"requestId"`
	ContextStack       *string              `json:"contextStack"`
	DuplicationAvoider *string              `json:"duplicationAvoider"`
	NamespaceName      *string              `json:"namespaceName"`
	Description        *string              `json:"description"`
	ServerType         *string              `json:"serverType"`
	ServerSpec         *string              `json:"serverSpec"`
	CreateNotification *NotificationSetting `json:"createNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:      core.CastString(data["namespaceName"]),
		Description:        core.CastString(data["description"]),
		ServerType:         core.CastString(data["serverType"]),
		ServerSpec:         core.CastString(data["serverSpec"]),
		CreateNotification: NewNotificationSettingFromDict(core.CastMap(data["createNotification"])).Pointer(),
		LogSetting:         NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":      p.NamespaceName,
		"description":        p.Description,
		"serverType":         p.ServerType,
		"serverSpec":         p.ServerSpec,
		"createNotification": p.CreateNotification.ToDict(),
		"logSetting":         p.LogSetting.ToDict(),
	}
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
	return &p
}

type DeleteNamespaceRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteNamespaceRequestFromDict(dict)
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeRoomsRequestFromJson(data string) DescribeRoomsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRoomsRequestFromDict(dict)
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

type WantRoomRequest struct {
	RequestId           *string  `json:"requestId"`
	ContextStack        *string  `json:"contextStack"`
	DuplicationAvoider  *string  `json:"duplicationAvoider"`
	NamespaceName       *string  `json:"namespaceName"`
	Name                *string  `json:"name"`
	NotificationUserIds []string `json:"notificationUserIds"`
}

func NewWantRoomRequestFromJson(data string) WantRoomRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWantRoomRequestFromDict(dict)
}

func NewWantRoomRequestFromDict(data map[string]interface{}) WantRoomRequest {
	return WantRoomRequest{
		NamespaceName:       core.CastString(data["namespaceName"]),
		Name:                core.CastString(data["name"]),
		NotificationUserIds: core.CastStrings(core.CastArray(data["notificationUserIds"])),
	}
}

func (p WantRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"notificationUserIds": core.CastStringsFromDict(
			p.NotificationUserIds,
		),
	}
}

func (p WantRoomRequest) Pointer() *WantRoomRequest {
	return &p
}

type GetRoomRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
}

func NewGetRoomRequestFromJson(data string) GetRoomRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRoomRequestFromDict(dict)
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

type DeleteRoomRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
}

func NewDeleteRoomRequestFromJson(data string) DeleteRoomRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRoomRequestFromDict(dict)
}

func NewDeleteRoomRequestFromDict(data map[string]interface{}) DeleteRoomRequest {
	return DeleteRoomRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
	}
}

func (p DeleteRoomRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
	}
}

func (p DeleteRoomRequest) Pointer() *DeleteRoomRequest {
	return &p
}
