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

package matchmaking

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
	SourceRequestId                               *string              `json:"sourceRequestId"`
	RequestId                                     *string              `json:"requestId"`
	ContextStack                                  *string              `json:"contextStack"`
	Name                                          *string              `json:"name"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	EnableCollaborateSeasonRating                 *string              `json:"enableCollaborateSeasonRating"`
	CollaborateSeasonRatingNamespaceId            *string              `json:"collaborateSeasonRatingNamespaceId"`
	CollaborateSeasonRatingTtl                    *int32               `json:"collaborateSeasonRatingTtl"`
	ChangeRatingScript                            *ScriptSetting       `json:"changeRatingScript"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification                      *NotificationSetting `json:"changeRatingNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
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
		EnableRating:               core.CastBool(data["enableRating"]),
		CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		EnableCollaborateSeasonRating:                 core.CastString(data["enableCollaborateSeasonRating"]),
		CollaborateSeasonRatingNamespaceId:            core.CastString(data["collaborateSeasonRatingNamespaceId"]),
		CollaborateSeasonRatingTtl:                    core.CastInt32(data["collaborateSeasonRatingTtl"]),
		ChangeRatingScript:                            NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		ChangeRatingNotification:                      NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                       p.Name,
		"description":                p.Description,
		"enableRating":               p.EnableRating,
		"createGatheringTriggerType": p.CreateGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     p.CreateGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                p.CreateGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                p.CompleteMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            p.CompleteMatchmakingTriggerScriptId,
		"enableCollaborateSeasonRating":                 p.EnableCollaborateSeasonRating,
		"collaborateSeasonRatingNamespaceId":            p.CollaborateSeasonRatingNamespaceId,
		"collaborateSeasonRatingTtl":                    p.CollaborateSeasonRatingTtl,
		"changeRatingScript":                            p.ChangeRatingScript.ToDict(),
		"joinNotification":                              p.JoinNotification.ToDict(),
		"leaveNotification":                             p.LeaveNotification.ToDict(),
		"completeNotification":                          p.CompleteNotification.ToDict(),
		"changeRatingNotification":                      p.ChangeRatingNotification.ToDict(),
		"logSetting":                                    p.LogSetting.ToDict(),
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
	SourceRequestId                               *string              `json:"sourceRequestId"`
	RequestId                                     *string              `json:"requestId"`
	ContextStack                                  *string              `json:"contextStack"`
	NamespaceName                                 *string              `json:"namespaceName"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	EnableCollaborateSeasonRating                 *string              `json:"enableCollaborateSeasonRating"`
	CollaborateSeasonRatingNamespaceId            *string              `json:"collaborateSeasonRatingNamespaceId"`
	CollaborateSeasonRatingTtl                    *int32               `json:"collaborateSeasonRatingTtl"`
	ChangeRatingScript                            *ScriptSetting       `json:"changeRatingScript"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	ChangeRatingNotification                      *NotificationSetting `json:"changeRatingNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
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
		EnableRating:               core.CastBool(data["enableRating"]),
		CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
		CreateGatheringTriggerRealtimeNamespaceId:     core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
		CreateGatheringTriggerScriptId:                core.CastString(data["createGatheringTriggerScriptId"]),
		CompleteMatchmakingTriggerType:                core.CastString(data["completeMatchmakingTriggerType"]),
		CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
		CompleteMatchmakingTriggerScriptId:            core.CastString(data["completeMatchmakingTriggerScriptId"]),
		EnableCollaborateSeasonRating:                 core.CastString(data["enableCollaborateSeasonRating"]),
		CollaborateSeasonRatingNamespaceId:            core.CastString(data["collaborateSeasonRatingNamespaceId"]),
		CollaborateSeasonRatingTtl:                    core.CastInt32(data["collaborateSeasonRatingTtl"]),
		ChangeRatingScript:                            NewScriptSettingFromDict(core.CastMap(data["changeRatingScript"])).Pointer(),
		JoinNotification:                              NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
		LeaveNotification:                             NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
		CompleteNotification:                          NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		ChangeRatingNotification:                      NewNotificationSettingFromDict(core.CastMap(data["changeRatingNotification"])).Pointer(),
		LogSetting:                                    NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":              p.NamespaceName,
		"description":                p.Description,
		"enableRating":               p.EnableRating,
		"createGatheringTriggerType": p.CreateGatheringTriggerType,
		"createGatheringTriggerRealtimeNamespaceId":     p.CreateGatheringTriggerRealtimeNamespaceId,
		"createGatheringTriggerScriptId":                p.CreateGatheringTriggerScriptId,
		"completeMatchmakingTriggerType":                p.CompleteMatchmakingTriggerType,
		"completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
		"completeMatchmakingTriggerScriptId":            p.CompleteMatchmakingTriggerScriptId,
		"enableCollaborateSeasonRating":                 p.EnableCollaborateSeasonRating,
		"collaborateSeasonRatingNamespaceId":            p.CollaborateSeasonRatingNamespaceId,
		"collaborateSeasonRatingTtl":                    p.CollaborateSeasonRatingTtl,
		"changeRatingScript":                            p.ChangeRatingScript.ToDict(),
		"joinNotification":                              p.JoinNotification.ToDict(),
		"leaveNotification":                             p.LeaveNotification.ToDict(),
		"completeNotification":                          p.CompleteNotification.ToDict(),
		"changeRatingNotification":                      p.ChangeRatingNotification.ToDict(),
		"logSetting":                                    p.LogSetting.ToDict(),
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

type DescribeGatheringsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeGatheringsRequestFromJson(data string) DescribeGatheringsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeGatheringsRequestFromDict(dict)
}

func NewDescribeGatheringsRequestFromDict(data map[string]interface{}) DescribeGatheringsRequest {
	return DescribeGatheringsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeGatheringsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeGatheringsRequest) Pointer() *DescribeGatheringsRequest {
	return &p
}

type CreateGatheringRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	AccessToken        *string          `json:"accessToken"`
	Player             *Player          `json:"player"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles    []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds       []*string        `json:"allowUserIds"`
	ExpiresAt          *int64           `json:"expiresAt"`
	ExpiresAtTimeSpan  *TimeSpan        `json:"expiresAtTimeSpan"`
}

func NewCreateGatheringRequestFromJson(data string) CreateGatheringRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGatheringRequestFromDict(dict)
}

func NewCreateGatheringRequestFromDict(data map[string]interface{}) CreateGatheringRequest {
	return CreateGatheringRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		AccessToken:       core.CastString(data["accessToken"]),
		Player:            NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		AttributeRanges:   CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles:   CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:      core.CastStrings(core.CastArray(data["allowUserIds"])),
		ExpiresAt:         core.CastInt64(data["expiresAt"]),
		ExpiresAtTimeSpan: NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer(),
	}
}

func (p CreateGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"player":        p.Player.ToDict(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt":         p.ExpiresAt,
		"expiresAtTimeSpan": p.ExpiresAtTimeSpan.ToDict(),
	}
}

func (p CreateGatheringRequest) Pointer() *CreateGatheringRequest {
	return &p
}

type CreateGatheringByUserIdRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	UserId             *string          `json:"userId"`
	Player             *Player          `json:"player"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	CapacityOfRoles    []CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds       []*string        `json:"allowUserIds"`
	ExpiresAt          *int64           `json:"expiresAt"`
	ExpiresAtTimeSpan  *TimeSpan        `json:"expiresAtTimeSpan"`
	TimeOffsetToken    *string          `json:"timeOffsetToken"`
}

func NewCreateGatheringByUserIdRequestFromJson(data string) CreateGatheringByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateGatheringByUserIdRequestFromDict(dict)
}

func NewCreateGatheringByUserIdRequestFromDict(data map[string]interface{}) CreateGatheringByUserIdRequest {
	return CreateGatheringByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		Player:            NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		AttributeRanges:   CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		CapacityOfRoles:   CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
		AllowUserIds:      core.CastStrings(core.CastArray(data["allowUserIds"])),
		ExpiresAt:         core.CastInt64(data["expiresAt"]),
		ExpiresAtTimeSpan: NewTimeSpanFromDict(core.CastMap(data["expiresAtTimeSpan"])).Pointer(),
		TimeOffsetToken:   core.CastString(data["timeOffsetToken"]),
	}
}

func (p CreateGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"player":        p.Player.ToDict(),
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"capacityOfRoles": CastCapacityOfRolesFromDict(
			p.CapacityOfRoles,
		),
		"allowUserIds": core.CastStringsFromDict(
			p.AllowUserIds,
		),
		"expiresAt":         p.ExpiresAt,
		"expiresAtTimeSpan": p.ExpiresAtTimeSpan.ToDict(),
		"timeOffsetToken":   p.TimeOffsetToken,
	}
}

func (p CreateGatheringByUserIdRequest) Pointer() *CreateGatheringByUserIdRequest {
	return &p
}

type UpdateGatheringRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	GatheringName      *string          `json:"gatheringName"`
	AccessToken        *string          `json:"accessToken"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
}

func NewUpdateGatheringRequestFromJson(data string) UpdateGatheringRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGatheringRequestFromDict(dict)
}

func NewUpdateGatheringRequestFromDict(data map[string]interface{}) UpdateGatheringRequest {
	return UpdateGatheringRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		AccessToken:     core.CastString(data["accessToken"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
	}
}

func (p UpdateGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
	}
}

func (p UpdateGatheringRequest) Pointer() *UpdateGatheringRequest {
	return &p
}

type UpdateGatheringByUserIdRequest struct {
	SourceRequestId    *string          `json:"sourceRequestId"`
	RequestId          *string          `json:"requestId"`
	ContextStack       *string          `json:"contextStack"`
	DuplicationAvoider *string          `json:"duplicationAvoider"`
	NamespaceName      *string          `json:"namespaceName"`
	GatheringName      *string          `json:"gatheringName"`
	UserId             *string          `json:"userId"`
	AttributeRanges    []AttributeRange `json:"attributeRanges"`
	TimeOffsetToken    *string          `json:"timeOffsetToken"`
}

func NewUpdateGatheringByUserIdRequestFromJson(data string) UpdateGatheringByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateGatheringByUserIdRequestFromDict(dict)
}

func NewUpdateGatheringByUserIdRequestFromDict(data map[string]interface{}) UpdateGatheringByUserIdRequest {
	return UpdateGatheringByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p UpdateGatheringByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"userId":        p.UserId,
		"attributeRanges": CastAttributeRangesFromDict(
			p.AttributeRanges,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p UpdateGatheringByUserIdRequest) Pointer() *UpdateGatheringByUserIdRequest {
	return &p
}

type DoMatchmakingByPlayerRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	NamespaceName           *string `json:"namespaceName"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func NewDoMatchmakingByPlayerRequestFromJson(data string) DoMatchmakingByPlayerRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingByPlayerRequestFromDict(dict)
}

func NewDoMatchmakingByPlayerRequestFromDict(data map[string]interface{}) DoMatchmakingByPlayerRequest {
	return DoMatchmakingByPlayerRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingByPlayerRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"player":                  p.Player.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingByPlayerRequest) Pointer() *DoMatchmakingByPlayerRequest {
	return &p
}

type DoMatchmakingRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	AccessToken             *string `json:"accessToken"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func NewDoMatchmakingRequestFromJson(data string) DoMatchmakingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingRequestFromDict(dict)
}

func NewDoMatchmakingRequestFromDict(data map[string]interface{}) DoMatchmakingRequest {
	return DoMatchmakingRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		AccessToken:             core.CastString(data["accessToken"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
	}
}

func (p DoMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"accessToken":             p.AccessToken,
		"player":                  p.Player.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
	}
}

func (p DoMatchmakingRequest) Pointer() *DoMatchmakingRequest {
	return &p
}

type DoMatchmakingByUserIdRequest struct {
	SourceRequestId         *string `json:"sourceRequestId"`
	RequestId               *string `json:"requestId"`
	ContextStack            *string `json:"contextStack"`
	DuplicationAvoider      *string `json:"duplicationAvoider"`
	NamespaceName           *string `json:"namespaceName"`
	UserId                  *string `json:"userId"`
	Player                  *Player `json:"player"`
	MatchmakingContextToken *string `json:"matchmakingContextToken"`
	TimeOffsetToken         *string `json:"timeOffsetToken"`
}

func NewDoMatchmakingByUserIdRequestFromJson(data string) DoMatchmakingByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDoMatchmakingByUserIdRequestFromDict(dict)
}

func NewDoMatchmakingByUserIdRequestFromDict(data map[string]interface{}) DoMatchmakingByUserIdRequest {
	return DoMatchmakingByUserIdRequest{
		NamespaceName:           core.CastString(data["namespaceName"]),
		UserId:                  core.CastString(data["userId"]),
		Player:                  NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
		MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
		TimeOffsetToken:         core.CastString(data["timeOffsetToken"]),
	}
}

func (p DoMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":           p.NamespaceName,
		"userId":                  p.UserId,
		"player":                  p.Player.ToDict(),
		"matchmakingContextToken": p.MatchmakingContextToken,
		"timeOffsetToken":         p.TimeOffsetToken,
	}
}

func (p DoMatchmakingByUserIdRequest) Pointer() *DoMatchmakingByUserIdRequest {
	return &p
}

type GetGatheringRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GatheringName   *string `json:"gatheringName"`
}

func NewGetGatheringRequestFromJson(data string) GetGatheringRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetGatheringRequestFromDict(dict)
}

func NewGetGatheringRequestFromDict(data map[string]interface{}) GetGatheringRequest {
	return GetGatheringRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
	}
}

func (p GetGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
	}
}

func (p GetGatheringRequest) Pointer() *GetGatheringRequest {
	return &p
}

type CancelMatchmakingRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	AccessToken        *string `json:"accessToken"`
}

func NewCancelMatchmakingRequestFromJson(data string) CancelMatchmakingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCancelMatchmakingRequestFromDict(dict)
}

func NewCancelMatchmakingRequestFromDict(data map[string]interface{}) CancelMatchmakingRequest {
	return CancelMatchmakingRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p CancelMatchmakingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
	}
}

func (p CancelMatchmakingRequest) Pointer() *CancelMatchmakingRequest {
	return &p
}

type CancelMatchmakingByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewCancelMatchmakingByUserIdRequestFromJson(data string) CancelMatchmakingByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCancelMatchmakingByUserIdRequestFromDict(dict)
}

func NewCancelMatchmakingByUserIdRequestFromDict(data map[string]interface{}) CancelMatchmakingByUserIdRequest {
	return CancelMatchmakingByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p CancelMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CancelMatchmakingByUserIdRequest) Pointer() *CancelMatchmakingByUserIdRequest {
	return &p
}

type EarlyCompleteRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	AccessToken        *string `json:"accessToken"`
}

func NewEarlyCompleteRequestFromJson(data string) EarlyCompleteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEarlyCompleteRequestFromDict(dict)
}

func NewEarlyCompleteRequestFromDict(data map[string]interface{}) EarlyCompleteRequest {
	return EarlyCompleteRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p EarlyCompleteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
		"accessToken":   p.AccessToken,
	}
}

func (p EarlyCompleteRequest) Pointer() *EarlyCompleteRequest {
	return &p
}

type EarlyCompleteByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	GatheringName      *string `json:"gatheringName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewEarlyCompleteByUserIdRequestFromJson(data string) EarlyCompleteByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewEarlyCompleteByUserIdRequestFromDict(dict)
}

func NewEarlyCompleteByUserIdRequestFromDict(data map[string]interface{}) EarlyCompleteByUserIdRequest {
	return EarlyCompleteByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p EarlyCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p EarlyCompleteByUserIdRequest) Pointer() *EarlyCompleteByUserIdRequest {
	return &p
}

type DeleteGatheringRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	GatheringName   *string `json:"gatheringName"`
}

func NewDeleteGatheringRequestFromJson(data string) DeleteGatheringRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteGatheringRequestFromDict(dict)
}

func NewDeleteGatheringRequestFromDict(data map[string]interface{}) DeleteGatheringRequest {
	return DeleteGatheringRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		GatheringName: core.CastString(data["gatheringName"]),
	}
}

func (p DeleteGatheringRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"gatheringName": p.GatheringName,
	}
}

func (p DeleteGatheringRequest) Pointer() *DeleteGatheringRequest {
	return &p
}

type DescribeRatingModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeRatingModelMastersRequestFromJson(data string) DescribeRatingModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingModelMastersRequestFromDict(dict)
}

func NewDescribeRatingModelMastersRequestFromDict(data map[string]interface{}) DescribeRatingModelMastersRequest {
	return DescribeRatingModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRatingModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRatingModelMastersRequest) Pointer() *DescribeRatingModelMastersRequest {
	return &p
}

type CreateRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	InitialValue    *int32  `json:"initialValue"`
	Volatility      *int32  `json:"volatility"`
}

func NewCreateRatingModelMasterRequestFromJson(data string) CreateRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateRatingModelMasterRequestFromDict(dict)
}

func NewCreateRatingModelMasterRequestFromDict(data map[string]interface{}) CreateRatingModelMasterRequest {
	return CreateRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Name:          core.CastString(data["name"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
	}
}

func (p CreateRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"initialValue":  p.InitialValue,
		"volatility":    p.Volatility,
	}
}

func (p CreateRatingModelMasterRequest) Pointer() *CreateRatingModelMasterRequest {
	return &p
}

type GetRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func NewGetRatingModelMasterRequestFromJson(data string) GetRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingModelMasterRequestFromDict(dict)
}

func NewGetRatingModelMasterRequestFromDict(data map[string]interface{}) GetRatingModelMasterRequest {
	return GetRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
	}
}

func (p GetRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingModelMasterRequest) Pointer() *GetRatingModelMasterRequest {
	return &p
}

type UpdateRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	Description     *string `json:"description"`
	Metadata        *string `json:"metadata"`
	InitialValue    *int32  `json:"initialValue"`
	Volatility      *int32  `json:"volatility"`
}

func NewUpdateRatingModelMasterRequestFromJson(data string) UpdateRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateRatingModelMasterRequestFromDict(dict)
}

func NewUpdateRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateRatingModelMasterRequest {
	return UpdateRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		Description:   core.CastString(data["description"]),
		Metadata:      core.CastString(data["metadata"]),
		InitialValue:  core.CastInt32(data["initialValue"]),
		Volatility:    core.CastInt32(data["volatility"]),
	}
}

func (p UpdateRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"description":   p.Description,
		"metadata":      p.Metadata,
		"initialValue":  p.InitialValue,
		"volatility":    p.Volatility,
	}
}

func (p UpdateRatingModelMasterRequest) Pointer() *UpdateRatingModelMasterRequest {
	return &p
}

type DeleteRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func NewDeleteRatingModelMasterRequestFromJson(data string) DeleteRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRatingModelMasterRequestFromDict(dict)
}

func NewDeleteRatingModelMasterRequestFromDict(data map[string]interface{}) DeleteRatingModelMasterRequest {
	return DeleteRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
	}
}

func (p DeleteRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p DeleteRatingModelMasterRequest) Pointer() *DeleteRatingModelMasterRequest {
	return &p
}

type DescribeRatingModelsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeRatingModelsRequestFromJson(data string) DescribeRatingModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingModelsRequestFromDict(dict)
}

func NewDescribeRatingModelsRequestFromDict(data map[string]interface{}) DescribeRatingModelsRequest {
	return DescribeRatingModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeRatingModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeRatingModelsRequest) Pointer() *DescribeRatingModelsRequest {
	return &p
}

type GetRatingModelRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
}

func NewGetRatingModelRequestFromJson(data string) GetRatingModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingModelRequestFromDict(dict)
}

func NewGetRatingModelRequestFromDict(data map[string]interface{}) GetRatingModelRequest {
	return GetRatingModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
	}
}

func (p GetRatingModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingModelRequest) Pointer() *GetRatingModelRequest {
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

type GetCurrentRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentRatingModelMasterRequestFromJson(data string) GetCurrentRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentRatingModelMasterRequestFromDict(dict)
}

func NewGetCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) GetCurrentRatingModelMasterRequest {
	return GetCurrentRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentRatingModelMasterRequest) Pointer() *GetCurrentRatingModelMasterRequest {
	return &p
}

type UpdateCurrentRatingModelMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentRatingModelMasterRequestFromJson(data string) UpdateCurrentRatingModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRatingModelMasterRequestFromDict(dict)
}

func NewUpdateCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterRequest {
	return UpdateCurrentRatingModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentRatingModelMasterRequest) Pointer() *UpdateCurrentRatingModelMasterRequest {
	return &p
}

type UpdateCurrentRatingModelMasterFromGitHubRequest struct {
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRatingModelMasterFromGitHubRequestFromJson(data string) UpdateCurrentRatingModelMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentRatingModelMasterFromGitHubRequestFromDict(dict)
}

func NewUpdateCurrentRatingModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterFromGitHubRequest {
	return UpdateCurrentRatingModelMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) Pointer() *UpdateCurrentRatingModelMasterFromGitHubRequest {
	return &p
}

type DescribeRatingsRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeRatingsRequestFromJson(data string) DescribeRatingsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingsRequestFromDict(dict)
}

func NewDescribeRatingsRequestFromDict(data map[string]interface{}) DescribeRatingsRequest {
	return DescribeRatingsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeRatingsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeRatingsRequest) Pointer() *DescribeRatingsRequest {
	return &p
}

type DescribeRatingsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeRatingsByUserIdRequestFromJson(data string) DescribeRatingsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeRatingsByUserIdRequestFromDict(dict)
}

func NewDescribeRatingsByUserIdRequestFromDict(data map[string]interface{}) DescribeRatingsByUserIdRequest {
	return DescribeRatingsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeRatingsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeRatingsByUserIdRequest) Pointer() *DescribeRatingsByUserIdRequest {
	return &p
}

type GetRatingRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	RatingName      *string `json:"ratingName"`
}

func NewGetRatingRequestFromJson(data string) GetRatingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingRequestFromDict(dict)
}

func NewGetRatingRequestFromDict(data map[string]interface{}) GetRatingRequest {
	return GetRatingRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		RatingName:    core.CastString(data["ratingName"]),
	}
}

func (p GetRatingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"ratingName":    p.RatingName,
	}
}

func (p GetRatingRequest) Pointer() *GetRatingRequest {
	return &p
}

type GetRatingByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	RatingName      *string `json:"ratingName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetRatingByUserIdRequestFromJson(data string) GetRatingByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetRatingByUserIdRequestFromDict(dict)
}

func NewGetRatingByUserIdRequestFromDict(data map[string]interface{}) GetRatingByUserIdRequest {
	return GetRatingByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		RatingName:      core.CastString(data["ratingName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetRatingByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"ratingName":      p.RatingName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetRatingByUserIdRequest) Pointer() *GetRatingByUserIdRequest {
	return &p
}

type PutResultRequest struct {
	SourceRequestId *string      `json:"sourceRequestId"`
	RequestId       *string      `json:"requestId"`
	ContextStack    *string      `json:"contextStack"`
	NamespaceName   *string      `json:"namespaceName"`
	RatingName      *string      `json:"ratingName"`
	GameResults     []GameResult `json:"gameResults"`
}

func NewPutResultRequestFromJson(data string) PutResultRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPutResultRequestFromDict(dict)
}

func NewPutResultRequestFromDict(data map[string]interface{}) PutResultRequest {
	return PutResultRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		GameResults:   CastGameResults(core.CastArray(data["gameResults"])),
	}
}

func (p PutResultRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
	}
}

func (p PutResultRequest) Pointer() *PutResultRequest {
	return &p
}

type DeleteRatingRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	RatingName         *string `json:"ratingName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteRatingRequestFromJson(data string) DeleteRatingRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteRatingRequestFromDict(dict)
}

func NewDeleteRatingRequestFromDict(data map[string]interface{}) DeleteRatingRequest {
	return DeleteRatingRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		RatingName:      core.CastString(data["ratingName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteRatingRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"ratingName":      p.RatingName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteRatingRequest) Pointer() *DeleteRatingRequest {
	return &p
}

type GetBallotRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
	AccessToken     *string `json:"accessToken"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
}

func NewGetBallotRequestFromJson(data string) GetBallotRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotRequestFromDict(dict)
}

func NewGetBallotRequestFromDict(data map[string]interface{}) GetBallotRequest {
	return GetBallotRequest{
		NamespaceName:  core.CastString(data["namespaceName"]),
		RatingName:     core.CastString(data["ratingName"]),
		GatheringName:  core.CastString(data["gatheringName"]),
		AccessToken:    core.CastString(data["accessToken"]),
		NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
		KeyId:          core.CastString(data["keyId"]),
	}
}

func (p GetBallotRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":  p.NamespaceName,
		"ratingName":     p.RatingName,
		"gatheringName":  p.GatheringName,
		"accessToken":    p.AccessToken,
		"numberOfPlayer": p.NumberOfPlayer,
		"keyId":          p.KeyId,
	}
}

func (p GetBallotRequest) Pointer() *GetBallotRequest {
	return &p
}

type GetBallotByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
	UserId          *string `json:"userId"`
	NumberOfPlayer  *int32  `json:"numberOfPlayer"`
	KeyId           *string `json:"keyId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetBallotByUserIdRequestFromJson(data string) GetBallotByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetBallotByUserIdRequestFromDict(dict)
}

func NewGetBallotByUserIdRequestFromDict(data map[string]interface{}) GetBallotByUserIdRequest {
	return GetBallotByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		RatingName:      core.CastString(data["ratingName"]),
		GatheringName:   core.CastString(data["gatheringName"]),
		UserId:          core.CastString(data["userId"]),
		NumberOfPlayer:  core.CastInt32(data["numberOfPlayer"]),
		KeyId:           core.CastString(data["keyId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetBallotByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"ratingName":      p.RatingName,
		"gatheringName":   p.GatheringName,
		"userId":          p.UserId,
		"numberOfPlayer":  p.NumberOfPlayer,
		"keyId":           p.KeyId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetBallotByUserIdRequest) Pointer() *GetBallotByUserIdRequest {
	return &p
}

type VoteRequest struct {
	SourceRequestId *string      `json:"sourceRequestId"`
	RequestId       *string      `json:"requestId"`
	ContextStack    *string      `json:"contextStack"`
	NamespaceName   *string      `json:"namespaceName"`
	BallotBody      *string      `json:"ballotBody"`
	BallotSignature *string      `json:"ballotSignature"`
	GameResults     []GameResult `json:"gameResults"`
	KeyId           *string      `json:"keyId"`
}

func NewVoteRequestFromJson(data string) VoteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteRequestFromDict(dict)
}

func NewVoteRequestFromDict(data map[string]interface{}) VoteRequest {
	return VoteRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		BallotBody:      core.CastString(data["ballotBody"]),
		BallotSignature: core.CastString(data["ballotSignature"]),
		GameResults:     CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:           core.CastString(data["keyId"]),
	}
}

func (p VoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"ballotBody":      p.BallotBody,
		"ballotSignature": p.BallotSignature,
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteRequest) Pointer() *VoteRequest {
	return &p
}

type VoteMultipleRequest struct {
	SourceRequestId *string        `json:"sourceRequestId"`
	RequestId       *string        `json:"requestId"`
	ContextStack    *string        `json:"contextStack"`
	NamespaceName   *string        `json:"namespaceName"`
	SignedBallots   []SignedBallot `json:"signedBallots"`
	GameResults     []GameResult   `json:"gameResults"`
	KeyId           *string        `json:"keyId"`
}

func NewVoteMultipleRequestFromJson(data string) VoteMultipleRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewVoteMultipleRequestFromDict(dict)
}

func NewVoteMultipleRequestFromDict(data map[string]interface{}) VoteMultipleRequest {
	return VoteMultipleRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		SignedBallots: CastSignedBallots(core.CastArray(data["signedBallots"])),
		GameResults:   CastGameResults(core.CastArray(data["gameResults"])),
		KeyId:         core.CastString(data["keyId"]),
	}
}

func (p VoteMultipleRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"signedBallots": CastSignedBallotsFromDict(
			p.SignedBallots,
		),
		"gameResults": CastGameResultsFromDict(
			p.GameResults,
		),
		"keyId": p.KeyId,
	}
}

func (p VoteMultipleRequest) Pointer() *VoteMultipleRequest {
	return &p
}

type CommitVoteRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	RatingName      *string `json:"ratingName"`
	GatheringName   *string `json:"gatheringName"`
}

func NewCommitVoteRequestFromJson(data string) CommitVoteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCommitVoteRequestFromDict(dict)
}

func NewCommitVoteRequestFromDict(data map[string]interface{}) CommitVoteRequest {
	return CommitVoteRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		RatingName:    core.CastString(data["ratingName"]),
		GatheringName: core.CastString(data["gatheringName"]),
	}
}

func (p CommitVoteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"ratingName":    p.RatingName,
		"gatheringName": p.GatheringName,
	}
}

func (p CommitVoteRequest) Pointer() *CommitVoteRequest {
	return &p
}
