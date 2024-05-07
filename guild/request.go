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

package guild

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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                       core.CastString(data["name"]),
		Description:                core.CastString(data["description"]),
		JoinNotification:           NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:          NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		ChangeMemberNotification:   NewNotificationSettingFromDict(core.CastMap(data["changeMemberNotification"])).Pointer(),
		ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
		RemoveRequestNotification:  NewNotificationSettingFromDict(core.CastMap(data["removeRequestNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                       p.Name,
		"description":                p.Description,
		"joinNotification":           p.JoinNotification.ToDict(),
		"leaveNotification":          p.LeaveNotification.ToDict(),
		"changeMemberNotification":   p.ChangeMemberNotification.ToDict(),
		"receiveRequestNotification": p.ReceiveRequestNotification.ToDict(),
		"removeRequestNotification":  p.RemoveRequestNotification.ToDict(),
		"logSetting":                 p.LogSetting.ToDict(),
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
	SourceRequestId            *string              `json:"sourceRequestId"`
	RequestId                  *string              `json:"requestId"`
	ContextStack               *string              `json:"contextStack"`
	NamespaceName              *string              `json:"namespaceName"`
	Description                *string              `json:"description"`
	JoinNotification           *NotificationSetting `json:"joinNotification"`
	LeaveNotification          *NotificationSetting `json:"leaveNotification"`
	ChangeMemberNotification   *NotificationSetting `json:"changeMemberNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	RemoveRequestNotification  *NotificationSetting `json:"removeRequestNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:              core.CastString(data["namespaceName"]),
		Description:                core.CastString(data["description"]),
		JoinNotification:           NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:          NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		ChangeMemberNotification:   NewNotificationSettingFromDict(core.CastMap(data["changeMemberNotification"])).Pointer(),
		ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
		RemoveRequestNotification:  NewNotificationSettingFromDict(core.CastMap(data["removeRequestNotification"])).Pointer(),
		LogSetting:                 NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":              p.NamespaceName,
		"description":                p.Description,
		"joinNotification":           p.JoinNotification.ToDict(),
		"leaveNotification":          p.LeaveNotification.ToDict(),
		"changeMemberNotification":   p.ChangeMemberNotification.ToDict(),
		"receiveRequestNotification": p.ReceiveRequestNotification.ToDict(),
		"removeRequestNotification":  p.RemoveRequestNotification.ToDict(),
		"logSetting":                 p.LogSetting.ToDict(),
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDumpUserDataByUserIdRequestFromJson(data string) DumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDumpUserDataByUserIdRequestFromDict(dict)
}

func NewDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) DumpUserDataByUserIdRequest {
	return DumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckDumpUserDataByUserIdRequestFromJson(data string) CheckDumpUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckDumpUserDataByUserIdRequestFromDict(dict)
}

func NewCheckDumpUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckDumpUserDataByUserIdRequest {
	return CheckDumpUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckDumpUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCleanUserDataByUserIdRequestFromJson(data string) CleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CleanUserDataByUserIdRequest {
	return CleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckCleanUserDataByUserIdRequestFromJson(data string) CheckCleanUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckCleanUserDataByUserIdRequestFromDict(dict)
}

func NewCheckCleanUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckCleanUserDataByUserIdRequest {
	return CheckCleanUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckCleanUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewPrepareImportUserDataByUserIdRequestFromJson(data string) PrepareImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPrepareImportUserDataByUserIdRequestFromDict(dict)
}

func NewPrepareImportUserDataByUserIdRequestFromDict(data map[string]interface{}) PrepareImportUserDataByUserIdRequest {
	return PrepareImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PrepareImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewImportUserDataByUserIdRequestFromJson(data string) ImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewImportUserDataByUserIdRequestFromDict(dict)
}

func NewImportUserDataByUserIdRequestFromDict(data map[string]interface{}) ImportUserDataByUserIdRequest {
	return ImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p ImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
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
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCheckImportUserDataByUserIdRequestFromJson(data string) CheckImportUserDataByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCheckImportUserDataByUserIdRequestFromDict(dict)
}

func NewCheckImportUserDataByUserIdRequestFromDict(data map[string]interface{}) CheckImportUserDataByUserIdRequest {
	return CheckImportUserDataByUserIdRequest{
		UserId:          core.CastString(data["userId"]),
		UploadToken:     core.CastString(data["uploadToken"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CheckImportUserDataByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"userId":          p.UserId,
		"uploadToken":     p.UploadToken,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CheckImportUserDataByUserIdRequest) Pointer() *CheckImportUserDataByUserIdRequest {
	return &p
}

type DescribeGuildModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeGuildModelMastersRequestFromJson(data string) DescribeGuildModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGuildModelMastersRequestFromDict(dict)
}

func NewDescribeGuildModelMastersRequestFromDict(data map[string]interface{}) DescribeGuildModelMastersRequest {
	return DescribeGuildModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeGuildModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGuildModelMastersRequest) Pointer() *DescribeGuildModelMastersRequest {
	return &p
}

type CreateGuildModelMasterRequest struct {
	SourceRequestId           *string     `json:"sourceRequestId"`
	RequestId                 *string     `json:"requestId"`
	ContextStack              *string     `json:"contextStack"`
	NamespaceName             *string     `json:"namespaceName"`
	Name                      *string     `json:"name"`
	Description               *string     `json:"description"`
	Metadata                  *string     `json:"metadata"`
	DefaultMaximumMemberCount *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount        *int32      `json:"maximumMemberCount"`
	Roles                     []RoleModel `json:"roles"`
	GuildMasterRole           *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes     *int32      `json:"rejoinCoolTimeMinutes"`
}

func NewCreateGuildModelMasterRequestFromJson(data string) CreateGuildModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildModelMasterRequestFromDict(dict)
}

func NewCreateGuildModelMasterRequestFromDict(data map[string]interface{}) CreateGuildModelMasterRequest {
	return CreateGuildModelMasterRequest{
		NamespaceName:             core.CastString(data["namespaceName"]),
		Name:                      core.CastString(data["name"]),
		Description:               core.CastString(data["description"]),
		Metadata:                  core.CastString(data["metadata"]),
		DefaultMaximumMemberCount: core.CastInt32(data["defaultMaximumMemberCount"]),
		MaximumMemberCount:        core.CastInt32(data["maximumMemberCount"]),
		Roles:                     CastRoleModels(core.CastArray(data["roles"])),
		GuildMasterRole:           core.CastString(data["guildMasterRole"]),
		GuildMemberDefaultRole:    core.CastString(data["guildMemberDefaultRole"]),
		RejoinCoolTimeMinutes:     core.CastInt32(data["rejoinCoolTimeMinutes"]),
	}
}

func (p CreateGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"name":                      p.Name,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"defaultMaximumMemberCount": p.DefaultMaximumMemberCount,
		"maximumMemberCount":        p.MaximumMemberCount,
		"roles": CastRoleModelsFromDict(
			p.Roles,
		),
		"guildMasterRole":        p.GuildMasterRole,
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
		"rejoinCoolTimeMinutes":  p.RejoinCoolTimeMinutes,
	}
}

func (p CreateGuildModelMasterRequest) Pointer() *CreateGuildModelMasterRequest {
	return &p
}

type GetGuildModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func NewGetGuildModelMasterRequestFromJson(data string) GetGuildModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildModelMasterRequestFromDict(dict)
}

func NewGetGuildModelMasterRequestFromDict(data map[string]interface{}) GetGuildModelMasterRequest {
	return GetGuildModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
	}
}

func (p GetGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p GetGuildModelMasterRequest) Pointer() *GetGuildModelMasterRequest {
	return &p
}

type UpdateGuildModelMasterRequest struct {
	SourceRequestId           *string     `json:"sourceRequestId"`
	RequestId                 *string     `json:"requestId"`
	ContextStack              *string     `json:"contextStack"`
	NamespaceName             *string     `json:"namespaceName"`
	GuildModelName            *string     `json:"guildModelName"`
	Description               *string     `json:"description"`
	Metadata                  *string     `json:"metadata"`
	DefaultMaximumMemberCount *int32      `json:"defaultMaximumMemberCount"`
	MaximumMemberCount        *int32      `json:"maximumMemberCount"`
	Roles                     []RoleModel `json:"roles"`
	GuildMasterRole           *string     `json:"guildMasterRole"`
	GuildMemberDefaultRole    *string     `json:"guildMemberDefaultRole"`
	RejoinCoolTimeMinutes     *int32      `json:"rejoinCoolTimeMinutes"`
}

func NewUpdateGuildModelMasterRequestFromJson(data string) UpdateGuildModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildModelMasterRequestFromDict(dict)
}

func NewUpdateGuildModelMasterRequestFromDict(data map[string]interface{}) UpdateGuildModelMasterRequest {
	return UpdateGuildModelMasterRequest{
		NamespaceName:             core.CastString(data["namespaceName"]),
		GuildModelName:            core.CastString(data["guildModelName"]),
		Description:               core.CastString(data["description"]),
		Metadata:                  core.CastString(data["metadata"]),
		DefaultMaximumMemberCount: core.CastInt32(data["defaultMaximumMemberCount"]),
		MaximumMemberCount:        core.CastInt32(data["maximumMemberCount"]),
		Roles:                     CastRoleModels(core.CastArray(data["roles"])),
		GuildMasterRole:           core.CastString(data["guildMasterRole"]),
		GuildMemberDefaultRole:    core.CastString(data["guildMemberDefaultRole"]),
		RejoinCoolTimeMinutes:     core.CastInt32(data["rejoinCoolTimeMinutes"]),
	}
}

func (p UpdateGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":             p.NamespaceName,
		"guildModelName":            p.GuildModelName,
		"description":               p.Description,
		"metadata":                  p.Metadata,
		"defaultMaximumMemberCount": p.DefaultMaximumMemberCount,
		"maximumMemberCount":        p.MaximumMemberCount,
		"roles": CastRoleModelsFromDict(
			p.Roles,
		),
		"guildMasterRole":        p.GuildMasterRole,
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
		"rejoinCoolTimeMinutes":  p.RejoinCoolTimeMinutes,
	}
}

func (p UpdateGuildModelMasterRequest) Pointer() *UpdateGuildModelMasterRequest {
	return &p
}

type DeleteGuildModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func NewDeleteGuildModelMasterRequestFromJson(data string) DeleteGuildModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildModelMasterRequestFromDict(dict)
}

func NewDeleteGuildModelMasterRequestFromDict(data map[string]interface{}) DeleteGuildModelMasterRequest {
	return DeleteGuildModelMasterRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
	}
}

func (p DeleteGuildModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p DeleteGuildModelMasterRequest) Pointer() *DeleteGuildModelMasterRequest {
	return &p
}

type DescribeGuildModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeGuildModelsRequestFromJson(data string) DescribeGuildModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGuildModelsRequestFromDict(dict)
}

func NewDescribeGuildModelsRequestFromDict(data map[string]interface{}) DescribeGuildModelsRequest {
	return DescribeGuildModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeGuildModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeGuildModelsRequest) Pointer() *DescribeGuildModelsRequest {
	return &p
}

type GetGuildModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
}

func NewGetGuildModelRequestFromJson(data string) GetGuildModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildModelRequestFromDict(dict)
}

func NewGetGuildModelRequestFromDict(data map[string]interface{}) GetGuildModelRequest {
	return GetGuildModelRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
	}
}

func (p GetGuildModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
	}
}

func (p GetGuildModelRequest) Pointer() *GetGuildModelRequest {
	return &p
}

type SearchGuildsRequest struct {
	SourceRequestId         *string   `json:"sourceRequestId"`
	RequestId               *string   `json:"requestId"`
	ContextStack            *string   `json:"contextStack"`
	DuplicationAvoider      *string   `json:"duplicationAvoider"`
	NamespaceName           *string   `json:"namespaceName"`
	GuildModelName          *string   `json:"guildModelName"`
	AccessToken             *string   `json:"accessToken"`
	DisplayName             *string   `json:"displayName"`
	OperationPolicies       []*int32  `json:"operationPolicies"`
	JoinPolicies            []*string `json:"joinPolicies"`
	IncludeFullMembersGuild *bool     `json:"includeFullMembersGuild"`
	PageToken               *string   `json:"pageToken"`
	Limit                   *int32    `json:"limit"`
}

func NewSearchGuildsRequestFromJson(data string) SearchGuildsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSearchGuildsRequestFromDict(dict)
}

func NewSearchGuildsRequestFromDict(data map[string]interface{}) SearchGuildsRequest {
	return SearchGuildsRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		GuildModelName:          core.CastString(data["guildModelName"]),
		AccessToken:             core.CastString(data["accessToken"]),
		DisplayName:             core.CastString(data["displayName"]),
		OperationPolicies:       core.CastInt32s(core.CastArray(data["operationPolicies"])),
		JoinPolicies:            core.CastStrings(core.CastArray(data["joinPolicies"])),
		IncludeFullMembersGuild: core.CastBool(data["includeFullMembersGuild"]),
		PageToken:               core.CastString(data["pageToken"]),
		Limit:                   core.CastInt32(data["limit"]),
	}
}

func (p SearchGuildsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"displayName":    p.DisplayName,
		"operationPolicies": core.CastInt32sFromDict(
			p.OperationPolicies,
		),
		"joinPolicies": core.CastStringsFromDict(
			p.JoinPolicies,
		),
		"includeFullMembersGuild": p.IncludeFullMembersGuild,
		"pageToken":               p.PageToken,
		"limit":                   p.Limit,
	}
}

func (p SearchGuildsRequest) Pointer() *SearchGuildsRequest {
	return &p
}

type SearchGuildsByUserIdRequest struct {
	SourceRequestId         *string   `json:"sourceRequestId"`
	RequestId               *string   `json:"requestId"`
	ContextStack            *string   `json:"contextStack"`
	DuplicationAvoider      *string   `json:"duplicationAvoider"`
	NamespaceName           *string   `json:"namespaceName"`
	GuildModelName          *string   `json:"guildModelName"`
	UserId                  *string   `json:"userId"`
	DisplayName             *string   `json:"displayName"`
	OperationPolicies       []*int32  `json:"operationPolicies"`
	JoinPolicies            []*string `json:"joinPolicies"`
	IncludeFullMembersGuild *bool     `json:"includeFullMembersGuild"`
	PageToken               *string   `json:"pageToken"`
	Limit                   *int32    `json:"limit"`
	TimeOffsetToken         *string   `json:"timeOffsetToken"`
}

func NewSearchGuildsByUserIdRequestFromJson(data string) SearchGuildsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSearchGuildsByUserIdRequestFromDict(dict)
}

func NewSearchGuildsByUserIdRequestFromDict(data map[string]interface{}) SearchGuildsByUserIdRequest {
	return SearchGuildsByUserIdRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		GuildModelName:          core.CastString(data["guildModelName"]),
		UserId:                  core.CastString(data["userId"]),
		DisplayName:             core.CastString(data["displayName"]),
		OperationPolicies:       core.CastInt32s(core.CastArray(data["operationPolicies"])),
		JoinPolicies:            core.CastStrings(core.CastArray(data["joinPolicies"])),
		IncludeFullMembersGuild: core.CastBool(data["includeFullMembersGuild"]),
		PageToken:               core.CastString(data["pageToken"]),
		Limit:                   core.CastInt32(data["limit"]),
		TimeOffsetToken:         core.CastString(data["timeOffsetToken"]),
	}
}

func (p SearchGuildsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"userId":         p.UserId,
		"displayName":    p.DisplayName,
		"operationPolicies": core.CastInt32sFromDict(
			p.OperationPolicies,
		),
		"joinPolicies": core.CastStringsFromDict(
			p.JoinPolicies,
		),
		"includeFullMembersGuild": p.IncludeFullMembersGuild,
		"pageToken":               p.PageToken,
		"limit":                   p.Limit,
		"timeOffsetToken":         p.TimeOffsetToken,
	}
}

func (p SearchGuildsByUserIdRequest) Pointer() *SearchGuildsByUserIdRequest {
	return &p
}

type CreateGuildRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	AccessToken            *string     `json:"accessToken"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	OperationPolicy        *int32      `json:"operationPolicy"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func NewCreateGuildRequestFromJson(data string) CreateGuildRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildRequestFromDict(dict)
}

func NewCreateGuildRequestFromDict(data map[string]interface{}) CreateGuildRequest {
	return CreateGuildRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		AccessToken:            core.CastString(data["accessToken"]),
		GuildModelName:         core.CastString(data["guildModelName"]),
		DisplayName:            core.CastString(data["displayName"]),
		OperationPolicy:        core.CastInt32(data["operationPolicy"]),
		JoinPolicy:             core.CastString(data["joinPolicy"]),
		CustomRoles:            CastRoleModels(core.CastArray(data["customRoles"])),
		GuildMemberDefaultRole: core.CastString(data["guildMemberDefaultRole"]),
	}
}

func (p CreateGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"displayName":     p.DisplayName,
		"operationPolicy": p.OperationPolicy,
		"joinPolicy":      p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p CreateGuildRequest) Pointer() *CreateGuildRequest {
	return &p
}

type CreateGuildByUserIdRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	UserId                 *string     `json:"userId"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	OperationPolicy        *int32      `json:"operationPolicy"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
	TimeOffsetToken        *string     `json:"timeOffsetToken"`
}

func NewCreateGuildByUserIdRequestFromJson(data string) CreateGuildByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGuildByUserIdRequestFromDict(dict)
}

func NewCreateGuildByUserIdRequestFromDict(data map[string]interface{}) CreateGuildByUserIdRequest {
	return CreateGuildByUserIdRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		UserId:                 core.CastString(data["userId"]),
		GuildModelName:         core.CastString(data["guildModelName"]),
		DisplayName:            core.CastString(data["displayName"]),
		OperationPolicy:        core.CastInt32(data["operationPolicy"]),
		JoinPolicy:             core.CastString(data["joinPolicy"]),
		CustomRoles:            CastRoleModels(core.CastArray(data["customRoles"])),
		GuildMemberDefaultRole: core.CastString(data["guildMemberDefaultRole"]),
		TimeOffsetToken:        core.CastString(data["timeOffsetToken"]),
	}
}

func (p CreateGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"displayName":     p.DisplayName,
		"operationPolicy": p.OperationPolicy,
		"joinPolicy":      p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
		"timeOffsetToken":        p.TimeOffsetToken,
	}
}

func (p CreateGuildByUserIdRequest) Pointer() *CreateGuildByUserIdRequest {
	return &p
}

type GetGuildRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
}

func NewGetGuildRequestFromJson(data string) GetGuildRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildRequestFromDict(dict)
}

func NewGetGuildRequestFromDict(data map[string]interface{}) GetGuildRequest {
	return GetGuildRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
	}
}

func (p GetGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p GetGuildRequest) Pointer() *GetGuildRequest {
	return &p
}

type GetGuildByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetGuildByUserIdRequestFromJson(data string) GetGuildByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGuildByUserIdRequestFromDict(dict)
}

func NewGetGuildByUserIdRequestFromDict(data map[string]interface{}) GetGuildByUserIdRequest {
	return GetGuildByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		GuildName:       core.CastString(data["guildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetGuildByUserIdRequest) Pointer() *GetGuildByUserIdRequest {
	return &p
}

type UpdateGuildRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	AccessToken            *string     `json:"accessToken"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	OperationPolicy        *int32      `json:"operationPolicy"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func NewUpdateGuildRequestFromJson(data string) UpdateGuildRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildRequestFromDict(dict)
}

func NewUpdateGuildRequestFromDict(data map[string]interface{}) UpdateGuildRequest {
	return UpdateGuildRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		AccessToken:            core.CastString(data["accessToken"]),
		GuildModelName:         core.CastString(data["guildModelName"]),
		DisplayName:            core.CastString(data["displayName"]),
		OperationPolicy:        core.CastInt32(data["operationPolicy"]),
		JoinPolicy:             core.CastString(data["joinPolicy"]),
		CustomRoles:            CastRoleModels(core.CastArray(data["customRoles"])),
		GuildMemberDefaultRole: core.CastString(data["guildMemberDefaultRole"]),
	}
}

func (p UpdateGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"displayName":     p.DisplayName,
		"operationPolicy": p.OperationPolicy,
		"joinPolicy":      p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p UpdateGuildRequest) Pointer() *UpdateGuildRequest {
	return &p
}

type UpdateGuildByGuildNameRequest struct {
	SourceRequestId        *string     `json:"sourceRequestId"`
	RequestId              *string     `json:"requestId"`
	ContextStack           *string     `json:"contextStack"`
	DuplicationAvoider     *string     `json:"duplicationAvoider"`
	NamespaceName          *string     `json:"namespaceName"`
	GuildName              *string     `json:"guildName"`
	GuildModelName         *string     `json:"guildModelName"`
	DisplayName            *string     `json:"displayName"`
	OperationPolicy        *int32      `json:"operationPolicy"`
	JoinPolicy             *string     `json:"joinPolicy"`
	CustomRoles            []RoleModel `json:"customRoles"`
	GuildMemberDefaultRole *string     `json:"guildMemberDefaultRole"`
}

func NewUpdateGuildByGuildNameRequestFromJson(data string) UpdateGuildByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGuildByGuildNameRequestFromDict(dict)
}

func NewUpdateGuildByGuildNameRequestFromDict(data map[string]interface{}) UpdateGuildByGuildNameRequest {
	return UpdateGuildByGuildNameRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		GuildName:              core.CastString(data["guildName"]),
		GuildModelName:         core.CastString(data["guildModelName"]),
		DisplayName:            core.CastString(data["displayName"]),
		OperationPolicy:        core.CastInt32(data["operationPolicy"]),
		JoinPolicy:             core.CastString(data["joinPolicy"]),
		CustomRoles:            CastRoleModels(core.CastArray(data["customRoles"])),
		GuildMemberDefaultRole: core.CastString(data["guildMemberDefaultRole"]),
	}
}

func (p UpdateGuildByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"guildName":       p.GuildName,
		"guildModelName":  p.GuildModelName,
		"displayName":     p.DisplayName,
		"operationPolicy": p.OperationPolicy,
		"joinPolicy":      p.JoinPolicy,
		"customRoles": CastRoleModelsFromDict(
			p.CustomRoles,
		),
		"guildMemberDefaultRole": p.GuildMemberDefaultRole,
	}
}

func (p UpdateGuildByGuildNameRequest) Pointer() *UpdateGuildByGuildNameRequest {
	return &p
}

type DeleteMemberRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
}

func NewDeleteMemberRequestFromJson(data string) DeleteMemberRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMemberRequestFromDict(dict)
}

func NewDeleteMemberRequestFromDict(data map[string]interface{}) DeleteMemberRequest {
	return DeleteMemberRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		TargetUserId:   core.CastString(data["targetUserId"]),
	}
}

func (p DeleteMemberRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"targetUserId":   p.TargetUserId,
	}
}

func (p DeleteMemberRequest) Pointer() *DeleteMemberRequest {
	return &p
}

type DeleteMemberByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TargetUserId       *string `json:"targetUserId"`
}

func NewDeleteMemberByGuildNameRequestFromJson(data string) DeleteMemberByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMemberByGuildNameRequestFromDict(dict)
}

func NewDeleteMemberByGuildNameRequestFromDict(data map[string]interface{}) DeleteMemberByGuildNameRequest {
	return DeleteMemberByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		TargetUserId:   core.CastString(data["targetUserId"]),
	}
}

func (p DeleteMemberByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"targetUserId":   p.TargetUserId,
	}
}

func (p DeleteMemberByGuildNameRequest) Pointer() *DeleteMemberByGuildNameRequest {
	return &p
}

type UpdateMemberRoleRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	TargetUserId       *string `json:"targetUserId"`
	RoleName           *string `json:"roleName"`
}

func NewUpdateMemberRoleRequestFromJson(data string) UpdateMemberRoleRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberRoleRequestFromDict(dict)
}

func NewUpdateMemberRoleRequestFromDict(data map[string]interface{}) UpdateMemberRoleRequest {
	return UpdateMemberRoleRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		TargetUserId:   core.CastString(data["targetUserId"]),
		RoleName:       core.CastString(data["roleName"]),
	}
}

func (p UpdateMemberRoleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"targetUserId":   p.TargetUserId,
		"roleName":       p.RoleName,
	}
}

func (p UpdateMemberRoleRequest) Pointer() *UpdateMemberRoleRequest {
	return &p
}

type UpdateMemberRoleByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TargetUserId       *string `json:"targetUserId"`
	RoleName           *string `json:"roleName"`
}

func NewUpdateMemberRoleByGuildNameRequestFromJson(data string) UpdateMemberRoleByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMemberRoleByGuildNameRequestFromDict(dict)
}

func NewUpdateMemberRoleByGuildNameRequestFromDict(data map[string]interface{}) UpdateMemberRoleByGuildNameRequest {
	return UpdateMemberRoleByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		TargetUserId:   core.CastString(data["targetUserId"]),
		RoleName:       core.CastString(data["roleName"]),
	}
}

func (p UpdateMemberRoleByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"targetUserId":   p.TargetUserId,
		"roleName":       p.RoleName,
	}
}

func (p UpdateMemberRoleByGuildNameRequest) Pointer() *UpdateMemberRoleByGuildNameRequest {
	return &p
}

type DeleteGuildRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
}

func NewDeleteGuildRequestFromJson(data string) DeleteGuildRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildRequestFromDict(dict)
}

func NewDeleteGuildRequestFromDict(data map[string]interface{}) DeleteGuildRequest {
	return DeleteGuildRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
	}
}

func (p DeleteGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
	}
}

func (p DeleteGuildRequest) Pointer() *DeleteGuildRequest {
	return &p
}

type DeleteGuildByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func NewDeleteGuildByGuildNameRequestFromJson(data string) DeleteGuildByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGuildByGuildNameRequestFromDict(dict)
}

func NewDeleteGuildByGuildNameRequestFromDict(data map[string]interface{}) DeleteGuildByGuildNameRequest {
	return DeleteGuildByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
	}
}

func (p DeleteGuildByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p DeleteGuildByGuildNameRequest) Pointer() *DeleteGuildByGuildNameRequest {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	Value              *int32  `json:"value"`
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(dict)
}

func NewIncreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		Value:          core.CastInt32(data["value"]),
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"value":          p.Value,
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	Value              *int32  `json:"value"`
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(dict)
}

func NewDecreaseMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		Value:          core.CastInt32(data["value"]),
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"value":          p.Value,
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildName          *string `json:"guildName"`
	GuildModelName     *string `json:"guildModelName"`
	Value              *int32  `json:"value"`
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameRequestFromJson(data string) SetMaximumCurrentMaximumMemberCountByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(dict)
}

func NewSetMaximumCurrentMaximumMemberCountByGuildNameRequestFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return SetMaximumCurrentMaximumMemberCountByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildName:      core.CastString(data["guildName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		Value:          core.CastInt32(data["value"]),
	}
}

func (p SetMaximumCurrentMaximumMemberCountByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildName":      p.GuildName,
		"guildModelName": p.GuildModelName,
		"value":          p.Value,
	}
}

func (p SetMaximumCurrentMaximumMemberCountByGuildNameRequest) Pointer() *SetMaximumCurrentMaximumMemberCountByGuildNameRequest {
	return &p
}

type AssumeRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func NewAssumeRequestFromJson(data string) AssumeRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAssumeRequestFromDict(dict)
}

func NewAssumeRequestFromDict(data map[string]interface{}) AssumeRequest {
	return AssumeRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
	}
}

func (p AssumeRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p AssumeRequest) Pointer() *AssumeRequest {
	return &p
}

type AssumeByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewAssumeByUserIdRequestFromJson(data string) AssumeByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAssumeByUserIdRequestFromDict(dict)
}

func NewAssumeByUserIdRequestFromDict(data map[string]interface{}) AssumeByUserIdRequest {
	return AssumeByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		GuildName:       core.CastString(data["guildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p AssumeByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p AssumeByUserIdRequest) Pointer() *AssumeByUserIdRequest {
	return &p
}

type IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetRequestFromJson(data string) IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(dict)
}

func NewIncreaseMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(data map[string]interface{}) IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest) Pointer() *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return &p
}

type DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskRequestFromJson(data string) DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskRequestFromDict(dict)
}

func NewDecreaseMaximumCurrentMaximumMemberCountByStampTaskRequestFromDict(data map[string]interface{}) DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest {
	return DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest) Pointer() *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest {
	return &p
}

type SetMaximumCurrentMaximumMemberCountByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewSetMaximumCurrentMaximumMemberCountByStampSheetRequestFromJson(data string) SetMaximumCurrentMaximumMemberCountByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSetMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(dict)
}

func NewSetMaximumCurrentMaximumMemberCountByStampSheetRequestFromDict(data map[string]interface{}) SetMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return SetMaximumCurrentMaximumMemberCountByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p SetMaximumCurrentMaximumMemberCountByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p SetMaximumCurrentMaximumMemberCountByStampSheetRequest) Pointer() *SetMaximumCurrentMaximumMemberCountByStampSheetRequest {
	return &p
}

type DescribeJoinedGuildsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeJoinedGuildsRequestFromJson(data string) DescribeJoinedGuildsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedGuildsRequestFromDict(dict)
}

func NewDescribeJoinedGuildsRequestFromDict(data map[string]interface{}) DescribeJoinedGuildsRequest {
	return DescribeJoinedGuildsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeJoinedGuildsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeJoinedGuildsRequest) Pointer() *DescribeJoinedGuildsRequest {
	return &p
}

type DescribeJoinedGuildsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeJoinedGuildsByUserIdRequestFromJson(data string) DescribeJoinedGuildsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJoinedGuildsByUserIdRequestFromDict(dict)
}

func NewDescribeJoinedGuildsByUserIdRequestFromDict(data map[string]interface{}) DescribeJoinedGuildsByUserIdRequest {
	return DescribeJoinedGuildsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeJoinedGuildsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeJoinedGuildsByUserIdRequest) Pointer() *DescribeJoinedGuildsByUserIdRequest {
	return &p
}

type GetJoinedGuildRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
}

func NewGetJoinedGuildRequestFromJson(data string) GetJoinedGuildRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedGuildRequestFromDict(dict)
}

func NewGetJoinedGuildRequestFromDict(data map[string]interface{}) GetJoinedGuildRequest {
	return GetJoinedGuildRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
	}
}

func (p GetJoinedGuildRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p GetJoinedGuildRequest) Pointer() *GetJoinedGuildRequest {
	return &p
}

type GetJoinedGuildByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetJoinedGuildByUserIdRequestFromJson(data string) GetJoinedGuildByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJoinedGuildByUserIdRequestFromDict(dict)
}

func NewGetJoinedGuildByUserIdRequestFromDict(data map[string]interface{}) GetJoinedGuildByUserIdRequest {
	return GetJoinedGuildByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		GuildName:       core.CastString(data["guildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetJoinedGuildByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetJoinedGuildByUserIdRequest) Pointer() *GetJoinedGuildByUserIdRequest {
	return &p
}

type WithdrawalRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
}

func NewWithdrawalRequestFromJson(data string) WithdrawalRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawalRequestFromDict(dict)
}

func NewWithdrawalRequestFromDict(data map[string]interface{}) WithdrawalRequest {
	return WithdrawalRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
	}
}

func (p WithdrawalRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
	}
}

func (p WithdrawalRequest) Pointer() *WithdrawalRequest {
	return &p
}

type WithdrawalByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewWithdrawalByUserIdRequestFromJson(data string) WithdrawalByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewWithdrawalByUserIdRequestFromDict(dict)
}

func NewWithdrawalByUserIdRequestFromDict(data map[string]interface{}) WithdrawalByUserIdRequest {
	return WithdrawalByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		GuildName:       core.CastString(data["guildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p WithdrawalByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"guildName":       p.GuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p WithdrawalByUserIdRequest) Pointer() *WithdrawalByUserIdRequest {
	return &p
}

type ExportMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewExportMasterRequestFromJson(data string) ExportMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewExportMasterRequestFromDict(dict)
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
	return ExportMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
	return &p
}

type GetCurrentGuildMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentGuildMasterRequestFromJson(data string) GetCurrentGuildMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentGuildMasterRequestFromDict(dict)
}

func NewGetCurrentGuildMasterRequestFromDict(data map[string]interface{}) GetCurrentGuildMasterRequest {
	return GetCurrentGuildMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentGuildMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentGuildMasterRequest) Pointer() *GetCurrentGuildMasterRequest {
	return &p
}

type UpdateCurrentGuildMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentGuildMasterRequestFromJson(data string) UpdateCurrentGuildMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGuildMasterRequestFromDict(dict)
}

func NewUpdateCurrentGuildMasterRequestFromDict(data map[string]interface{}) UpdateCurrentGuildMasterRequest {
	return UpdateCurrentGuildMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentGuildMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentGuildMasterRequest) Pointer() *UpdateCurrentGuildMasterRequest {
	return &p
}

type UpdateCurrentGuildMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentGuildMasterFromGitHubRequestFromJson(data string) UpdateCurrentGuildMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentGuildMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentGuildMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentGuildMasterFromGitHubRequest {
	return UpdateCurrentGuildMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentGuildMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentGuildMasterFromGitHubRequest) Pointer() *UpdateCurrentGuildMasterFromGitHubRequest {
	return &p
}

type DescribeReceiveRequestsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeReceiveRequestsRequestFromJson(data string) DescribeReceiveRequestsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveRequestsRequestFromDict(dict)
}

func NewDescribeReceiveRequestsRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsRequest {
	return DescribeReceiveRequestsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeReceiveRequestsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeReceiveRequestsRequest) Pointer() *DescribeReceiveRequestsRequest {
	return &p
}

type DescribeReceiveRequestsByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeReceiveRequestsByGuildNameRequestFromJson(data string) DescribeReceiveRequestsByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeReceiveRequestsByGuildNameRequestFromDict(dict)
}

func NewDescribeReceiveRequestsByGuildNameRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsByGuildNameRequest {
	return DescribeReceiveRequestsByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeReceiveRequestsByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeReceiveRequestsByGuildNameRequest) Pointer() *DescribeReceiveRequestsByGuildNameRequest {
	return &p
}

type GetReceiveRequestRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	AccessToken     *string `json:"accessToken"`
	FromUserId      *string `json:"fromUserId"`
}

func NewGetReceiveRequestRequestFromJson(data string) GetReceiveRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveRequestRequestFromDict(dict)
}

func NewGetReceiveRequestRequestFromDict(data map[string]interface{}) GetReceiveRequestRequest {
	return GetReceiveRequestRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p GetReceiveRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p GetReceiveRequestRequest) Pointer() *GetReceiveRequestRequest {
	return &p
}

type GetReceiveRequestByGuildNameRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GuildModelName  *string `json:"guildModelName"`
	GuildName       *string `json:"guildName"`
	FromUserId      *string `json:"fromUserId"`
}

func NewGetReceiveRequestByGuildNameRequestFromJson(data string) GetReceiveRequestByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetReceiveRequestByGuildNameRequestFromDict(dict)
}

func NewGetReceiveRequestByGuildNameRequestFromDict(data map[string]interface{}) GetReceiveRequestByGuildNameRequest {
	return GetReceiveRequestByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p GetReceiveRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p GetReceiveRequestByGuildNameRequest) Pointer() *GetReceiveRequestByGuildNameRequest {
	return &p
}

type AcceptRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	FromUserId         *string `json:"fromUserId"`
}

func NewAcceptRequestRequestFromJson(data string) AcceptRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestRequestFromDict(dict)
}

func NewAcceptRequestRequestFromDict(data map[string]interface{}) AcceptRequestRequest {
	return AcceptRequestRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p AcceptRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p AcceptRequestRequest) Pointer() *AcceptRequestRequest {
	return &p
}

type AcceptRequestByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	FromUserId         *string `json:"fromUserId"`
}

func NewAcceptRequestByGuildNameRequestFromJson(data string) AcceptRequestByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewAcceptRequestByGuildNameRequestFromDict(dict)
}

func NewAcceptRequestByGuildNameRequestFromDict(data map[string]interface{}) AcceptRequestByGuildNameRequest {
	return AcceptRequestByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p AcceptRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p AcceptRequestByGuildNameRequest) Pointer() *AcceptRequestByGuildNameRequest {
	return &p
}

type RejectRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	AccessToken        *string `json:"accessToken"`
	FromUserId         *string `json:"fromUserId"`
}

func NewRejectRequestRequestFromJson(data string) RejectRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRejectRequestRequestFromDict(dict)
}

func NewRejectRequestRequestFromDict(data map[string]interface{}) RejectRequestRequest {
	return RejectRequestRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p RejectRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"accessToken":    p.AccessToken,
		"fromUserId":     p.FromUserId,
	}
}

func (p RejectRequestRequest) Pointer() *RejectRequestRequest {
	return &p
}

type RejectRequestByGuildNameRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GuildModelName     *string `json:"guildModelName"`
	GuildName          *string `json:"guildName"`
	FromUserId         *string `json:"fromUserId"`
}

func NewRejectRequestByGuildNameRequestFromJson(data string) RejectRequestByGuildNameRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRejectRequestByGuildNameRequestFromDict(dict)
}

func NewRejectRequestByGuildNameRequestFromDict(data map[string]interface{}) RejectRequestByGuildNameRequest {
	return RejectRequestByGuildNameRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		GuildName:      core.CastString(data["guildName"]),
		FromUserId:     core.CastString(data["fromUserId"]),
	}
}

func (p RejectRequestByGuildNameRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"guildModelName": p.GuildModelName,
		"guildName":      p.GuildName,
		"fromUserId":     p.FromUserId,
	}
}

func (p RejectRequestByGuildNameRequest) Pointer() *RejectRequestByGuildNameRequest {
	return &p
}

type DescribeSendRequestsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeSendRequestsRequestFromJson(data string) DescribeSendRequestsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSendRequestsRequestFromDict(dict)
}

func NewDescribeSendRequestsRequestFromDict(data map[string]interface{}) DescribeSendRequestsRequest {
	return DescribeSendRequestsRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		GuildModelName: core.CastString(data["guildModelName"]),
		PageToken:      core.CastString(data["pageToken"]),
		Limit:          core.CastInt32(data["limit"]),
	}
}

func (p DescribeSendRequestsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"accessToken":    p.AccessToken,
		"guildModelName": p.GuildModelName,
		"pageToken":      p.PageToken,
		"limit":          p.Limit,
	}
}

func (p DescribeSendRequestsRequest) Pointer() *DescribeSendRequestsRequest {
	return &p
}

type DescribeSendRequestsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeSendRequestsByUserIdRequestFromJson(data string) DescribeSendRequestsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeSendRequestsByUserIdRequestFromDict(dict)
}

func NewDescribeSendRequestsByUserIdRequestFromDict(data map[string]interface{}) DescribeSendRequestsByUserIdRequest {
	return DescribeSendRequestsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeSendRequestsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeSendRequestsByUserIdRequest) Pointer() *DescribeSendRequestsByUserIdRequest {
	return &p
}

type GetSendRequestRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	GuildModelName  *string `json:"guildModelName"`
	TargetGuildName *string `json:"targetGuildName"`
}

func NewGetSendRequestRequestFromJson(data string) GetSendRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSendRequestRequestFromDict(dict)
}

func NewGetSendRequestRequestFromDict(data map[string]interface{}) GetSendRequestRequest {
	return GetSendRequestRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
	}
}

func (p GetSendRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p GetSendRequestRequest) Pointer() *GetSendRequestRequest {
	return &p
}

type GetSendRequestByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	GuildModelName  *string `json:"guildModelName"`
	TargetGuildName *string `json:"targetGuildName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetSendRequestByUserIdRequestFromJson(data string) GetSendRequestByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetSendRequestByUserIdRequestFromDict(dict)
}

func NewGetSendRequestByUserIdRequestFromDict(data map[string]interface{}) GetSendRequestByUserIdRequest {
	return GetSendRequestByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetSendRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetSendRequestByUserIdRequest) Pointer() *GetSendRequestByUserIdRequest {
	return &p
}

type SendRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
}

func NewSendRequestRequestFromJson(data string) SendRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendRequestRequestFromDict(dict)
}

func NewSendRequestRequestFromDict(data map[string]interface{}) SendRequestRequest {
	return SendRequestRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
	}
}

func (p SendRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p SendRequestRequest) Pointer() *SendRequestRequest {
	return &p
}

type SendRequestByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewSendRequestByUserIdRequestFromJson(data string) SendRequestByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewSendRequestByUserIdRequestFromDict(dict)
}

func NewSendRequestByUserIdRequestFromDict(data map[string]interface{}) SendRequestByUserIdRequest {
	return SendRequestByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p SendRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p SendRequestByUserIdRequest) Pointer() *SendRequestByUserIdRequest {
	return &p
}

type DeleteRequestRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
}

func NewDeleteRequestRequestFromJson(data string) DeleteRequestRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRequestRequestFromDict(dict)
}

func NewDeleteRequestRequestFromDict(data map[string]interface{}) DeleteRequestRequest {
	return DeleteRequestRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
	}
}

func (p DeleteRequestRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"accessToken":     p.AccessToken,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
	}
}

func (p DeleteRequestRequest) Pointer() *DeleteRequestRequest {
	return &p
}

type DeleteRequestByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	GuildModelName     *string `json:"guildModelName"`
	TargetGuildName    *string `json:"targetGuildName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteRequestByUserIdRequestFromJson(data string) DeleteRequestByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRequestByUserIdRequestFromDict(dict)
}

func NewDeleteRequestByUserIdRequestFromDict(data map[string]interface{}) DeleteRequestByUserIdRequest {
	return DeleteRequestByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		GuildModelName:  core.CastString(data["guildModelName"]),
		TargetGuildName: core.CastString(data["targetGuildName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteRequestByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"guildModelName":  p.GuildModelName,
		"targetGuildName": p.TargetGuildName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteRequestByUserIdRequest) Pointer() *DeleteRequestByUserIdRequest {
	return &p
}
