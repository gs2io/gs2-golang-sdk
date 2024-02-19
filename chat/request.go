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

type DescribeNamespacesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
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
	SourceRequestId       *string              `json:"sourceRequestId"`
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

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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
	SourceRequestId       *string              `json:"sourceRequestId"`
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

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
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

type DumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p DumpUserDataByUserIdRequest) Pointer() *DumpUserDataByUserIdRequest {
	return &p
}

type CheckDumpUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckDumpUserDataByUserIdRequest) Pointer() *CheckDumpUserDataByUserIdRequest {
	return &p
}

type CleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CleanUserDataByUserIdRequest) Pointer() *CleanUserDataByUserIdRequest {
	return &p
}

type CheckCleanUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p CheckCleanUserDataByUserIdRequest) Pointer() *CheckCleanUserDataByUserIdRequest {
	return &p
}

type PrepareImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId: core.CastString(data["userId"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId": p.UserId,
	}
}

func (p PrepareImportUserDataByUserIdRequest) Pointer() *PrepareImportUserDataByUserIdRequest {
	return &p
}

type ImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p ImportUserDataByUserIdRequest) Pointer() *ImportUserDataByUserIdRequest {
	return &p
}

type CheckImportUserDataByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	UserId             *string `json:"userId"`
	UploadToken        *string `json:"uploadToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:      core.CastString(data["userId"]),
		UploadToken: core.CastString(data["uploadToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":      p.UserId,
		"uploadToken": p.UploadToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeRoomsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
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

type CreateRoomRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	AccessToken        *string   `json:"accessToken"`
	Name               *string   `json:"name"`
	Metadata           *string   `json:"metadata"`
	Password           *string   `json:"password"`
	WhiteListUserIds   []*string `json:"whiteListUserIds"`
}

func NewCreateRoomRequestFromJson(data string) CreateRoomRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRoomRequestFromDict(dict)
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
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	Name               *string   `json:"name"`
	UserId             *string   `json:"userId"`
	Metadata           *string   `json:"metadata"`
	Password           *string   `json:"password"`
	WhiteListUserIds   []*string `json:"whiteListUserIds"`
}

func NewCreateRoomFromBackendRequestFromJson(data string) CreateRoomFromBackendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRoomFromBackendRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
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

type UpdateRoomRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	RoomName           *string   `json:"roomName"`
	Metadata           *string   `json:"metadata"`
	Password           *string   `json:"password"`
	WhiteListUserIds   []*string `json:"whiteListUserIds"`
	AccessToken        *string   `json:"accessToken"`
}

func NewUpdateRoomRequestFromJson(data string) UpdateRoomRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRoomRequestFromDict(dict)
}

func NewUpdateRoomRequestFromDict(data map[string]interface{}) UpdateRoomRequest {
	return UpdateRoomRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		RoomName:         core.CastString(data["roomName"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
		AccessToken:      core.CastString(data["accessToken"]),
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
		"accessToken": p.AccessToken,
	}
}

func (p UpdateRoomRequest) Pointer() *UpdateRoomRequest {
	return &p
}

type UpdateRoomFromBackendRequest struct {
	SourceRequestId    *string   `json:"sourceRequestId"`
	RequestId          *string   `json:"requestId"`
	ContextStack       *string   `json:"contextStack"`
	DuplicationAvoider *string   `json:"duplicationAvoider"`
	NamespaceName      *string   `json:"namespaceName"`
	RoomName           *string   `json:"roomName"`
	Metadata           *string   `json:"metadata"`
	Password           *string   `json:"password"`
	WhiteListUserIds   []*string `json:"whiteListUserIds"`
	UserId             *string   `json:"userId"`
}

func NewUpdateRoomFromBackendRequestFromJson(data string) UpdateRoomFromBackendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRoomFromBackendRequestFromDict(dict)
}

func NewUpdateRoomFromBackendRequestFromDict(data map[string]interface{}) UpdateRoomFromBackendRequest {
	return UpdateRoomFromBackendRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		RoomName:         core.CastString(data["roomName"]),
		Metadata:         core.CastString(data["metadata"]),
		Password:         core.CastString(data["password"]),
		WhiteListUserIds: core.CastStrings(core.CastArray(data["whiteListUserIds"])),
		UserId:           core.CastString(data["userId"]),
	}
}

func (p UpdateRoomFromBackendRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"metadata":      p.Metadata,
		"password":      p.Password,
		"whiteListUserIds": core.CastStringsFromDict(
			p.WhiteListUserIds,
		),
		"userId": p.UserId,
	}
}

func (p UpdateRoomFromBackendRequest) Pointer() *UpdateRoomFromBackendRequest {
	return &p
}

type DeleteRoomRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	AccessToken        *string `json:"accessToken"`
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	UserId             *string `json:"userId"`
}

func NewDeleteRoomFromBackendRequestFromJson(data string) DeleteRoomFromBackendRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRoomFromBackendRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	Password        *string `json:"password"`
	AccessToken     *string `json:"accessToken"`
	StartAt         *int64  `json:"startAt"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMessagesRequestFromJson(data string) DescribeMessagesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesRequestFromDict(dict)
}

func NewDescribeMessagesRequestFromDict(data map[string]interface{}) DescribeMessagesRequest {
	return DescribeMessagesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		Password:      core.CastString(data["password"]),
		AccessToken:   core.CastString(data["accessToken"]),
		StartAt:       core.CastInt64(data["startAt"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMessagesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"password":      p.Password,
		"accessToken":   p.AccessToken,
		"startAt":       p.StartAt,
		"limit":         p.Limit,
	}
}

func (p DescribeMessagesRequest) Pointer() *DescribeMessagesRequest {
	return &p
}

type DescribeMessagesByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	Password        *string `json:"password"`
	UserId          *string `json:"userId"`
	StartAt         *int64  `json:"startAt"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMessagesByUserIdRequestFromJson(data string) DescribeMessagesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMessagesByUserIdRequestFromDict(dict)
}

func NewDescribeMessagesByUserIdRequestFromDict(data map[string]interface{}) DescribeMessagesByUserIdRequest {
	return DescribeMessagesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		Password:      core.CastString(data["password"]),
		UserId:        core.CastString(data["userId"]),
		StartAt:       core.CastInt64(data["startAt"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMessagesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"password":      p.Password,
		"userId":        p.UserId,
		"startAt":       p.StartAt,
		"limit":         p.Limit,
	}
}

func (p DescribeMessagesByUserIdRequest) Pointer() *DescribeMessagesByUserIdRequest {
	return &p
}

type PostRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	AccessToken        *string `json:"accessToken"`
	Category           *int32  `json:"category"`
	Metadata           *string `json:"metadata"`
	Password           *string `json:"password"`
}

func NewPostRequestFromJson(data string) PostRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPostRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	UserId             *string `json:"userId"`
	Category           *int32  `json:"category"`
	Metadata           *string `json:"metadata"`
	Password           *string `json:"password"`
}

func NewPostByUserIdRequestFromJson(data string) PostByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPostByUserIdRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	MessageName     *string `json:"messageName"`
	Password        *string `json:"password"`
	AccessToken     *string `json:"accessToken"`
}

func NewGetMessageRequestFromJson(data string) GetMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageRequestFromDict(dict)
}

func NewGetMessageRequestFromDict(data map[string]interface{}) GetMessageRequest {
	return GetMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		MessageName:   core.CastString(data["messageName"]),
		Password:      core.CastString(data["password"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"messageName":   p.MessageName,
		"password":      p.Password,
		"accessToken":   p.AccessToken,
	}
}

func (p GetMessageRequest) Pointer() *GetMessageRequest {
	return &p
}

type GetMessageByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	MessageName     *string `json:"messageName"`
	Password        *string `json:"password"`
	UserId          *string `json:"userId"`
}

func NewGetMessageByUserIdRequestFromJson(data string) GetMessageByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMessageByUserIdRequestFromDict(dict)
}

func NewGetMessageByUserIdRequestFromDict(data map[string]interface{}) GetMessageByUserIdRequest {
	return GetMessageByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		MessageName:   core.CastString(data["messageName"]),
		Password:      core.CastString(data["password"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetMessageByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"messageName":   p.MessageName,
		"password":      p.Password,
		"userId":        p.UserId,
	}
}

func (p GetMessageByUserIdRequest) Pointer() *GetMessageByUserIdRequest {
	return &p
}

type DeleteMessageRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	UserId             *string `json:"userId"`
	MessageName        *string `json:"messageName"`
}

func NewDeleteMessageRequestFromJson(data string) DeleteMessageRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMessageRequestFromDict(dict)
}

func NewDeleteMessageRequestFromDict(data map[string]interface{}) DeleteMessageRequest {
	return DeleteMessageRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RoomName:      core.CastString(data["roomName"]),
		UserId:        core.CastString(data["userId"]),
		MessageName:   core.CastString(data["messageName"]),
	}
}

func (p DeleteMessageRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"roomName":      p.RoomName,
		"userId":        p.UserId,
		"messageName":   p.MessageName,
	}
}

func (p DeleteMessageRequest) Pointer() *DeleteMessageRequest {
	return &p
}

type DescribeSubscribesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSubscribesRequestFromJson(data string) DescribeSubscribesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSubscribesByUserIdRequestFromJson(data string) DescribeSubscribesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByUserIdRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSubscribesByRoomNameRequestFromJson(data string) DescribeSubscribesByRoomNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSubscribesByRoomNameRequestFromDict(dict)
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
	SourceRequestId    *string            `json:"sourceRequestId"`
	RequestId          *string            `json:"requestId"`
	ContextStack       *string            `json:"contextStack"`
	DuplicationAvoider *string            `json:"duplicationAvoider"`
	NamespaceName      *string            `json:"namespaceName"`
	RoomName           *string            `json:"roomName"`
	AccessToken        *string            `json:"accessToken"`
	NotificationTypes  []NotificationType `json:"notificationTypes"`
}

func NewSubscribeRequestFromJson(data string) SubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeRequestFromDict(dict)
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
	SourceRequestId    *string            `json:"sourceRequestId"`
	RequestId          *string            `json:"requestId"`
	ContextStack       *string            `json:"contextStack"`
	DuplicationAvoider *string            `json:"duplicationAvoider"`
	NamespaceName      *string            `json:"namespaceName"`
	RoomName           *string            `json:"roomName"`
	UserId             *string            `json:"userId"`
	NotificationTypes  []NotificationType `json:"notificationTypes"`
}

func NewSubscribeByUserIdRequestFromJson(data string) SubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSubscribeByUserIdRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	AccessToken     *string `json:"accessToken"`
}

func NewGetSubscribeRequestFromJson(data string) GetSubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RoomName        *string `json:"roomName"`
	UserId          *string `json:"userId"`
}

func NewGetSubscribeByUserIdRequestFromJson(data string) GetSubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSubscribeByUserIdRequestFromDict(dict)
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
	SourceRequestId    *string            `json:"sourceRequestId"`
	RequestId          *string            `json:"requestId"`
	ContextStack       *string            `json:"contextStack"`
	DuplicationAvoider *string            `json:"duplicationAvoider"`
	NamespaceName      *string            `json:"namespaceName"`
	RoomName           *string            `json:"roomName"`
	AccessToken        *string            `json:"accessToken"`
	NotificationTypes  []NotificationType `json:"notificationTypes"`
}

func NewUpdateNotificationTypeRequestFromJson(data string) UpdateNotificationTypeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNotificationTypeRequestFromDict(dict)
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
	SourceRequestId    *string            `json:"sourceRequestId"`
	RequestId          *string            `json:"requestId"`
	ContextStack       *string            `json:"contextStack"`
	DuplicationAvoider *string            `json:"duplicationAvoider"`
	NamespaceName      *string            `json:"namespaceName"`
	RoomName           *string            `json:"roomName"`
	UserId             *string            `json:"userId"`
	NotificationTypes  []NotificationType `json:"notificationTypes"`
}

func NewUpdateNotificationTypeByUserIdRequestFromJson(data string) UpdateNotificationTypeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNotificationTypeByUserIdRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	AccessToken        *string `json:"accessToken"`
}

func NewUnsubscribeRequestFromJson(data string) UnsubscribeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeRequestFromDict(dict)
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
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	RoomName           *string `json:"roomName"`
	UserId             *string `json:"userId"`
}

func NewUnsubscribeByUserIdRequestFromJson(data string) UnsubscribeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUnsubscribeByUserIdRequestFromDict(dict)
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
