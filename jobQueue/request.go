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

package jobQueue

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
	SourceRequestId  *string              `json:"sourceRequestId"`
	RequestId        *string              `json:"requestId"`
	ContextStack     *string              `json:"contextStack"`
	Name             *string              `json:"name"`
	Description      *string              `json:"description"`
	EnableAutoRun    *bool                `json:"enableAutoRun"`
	PushNotification *NotificationSetting `json:"pushNotification"`
	RunNotification  *NotificationSetting `json:"runNotification"`
	LogSetting       *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		EnableAutoRun:    core.CastBool(data["enableAutoRun"]),
		PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
		RunNotification:  NewNotificationSettingFromDict(core.CastMap(data["runNotification"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":             p.Name,
		"description":      p.Description,
		"enableAutoRun":    p.EnableAutoRun,
		"pushNotification": p.PushNotification.ToDict(),
		"runNotification":  p.RunNotification.ToDict(),
		"logSetting":       p.LogSetting.ToDict(),
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
	SourceRequestId  *string              `json:"sourceRequestId"`
	RequestId        *string              `json:"requestId"`
	ContextStack     *string              `json:"contextStack"`
	NamespaceName    *string              `json:"namespaceName"`
	Description      *string              `json:"description"`
	EnableAutoRun    *bool                `json:"enableAutoRun"`
	PushNotification *NotificationSetting `json:"pushNotification"`
	RunNotification  *NotificationSetting `json:"runNotification"`
	LogSetting       *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		Description:      core.CastString(data["description"]),
		EnableAutoRun:    core.CastBool(data["enableAutoRun"]),
		PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
		RunNotification:  NewNotificationSettingFromDict(core.CastMap(data["runNotification"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"description":      p.Description,
		"enableAutoRun":    p.EnableAutoRun,
		"pushNotification": p.PushNotification.ToDict(),
		"runNotification":  p.RunNotification.ToDict(),
		"logSetting":       p.LogSetting.ToDict(),
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

type DescribeJobsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeJobsByUserIdRequestFromJson(data string) DescribeJobsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeJobsByUserIdRequestFromDict(dict)
}

func NewDescribeJobsByUserIdRequestFromDict(data map[string]interface{}) DescribeJobsByUserIdRequest {
	return DescribeJobsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeJobsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeJobsByUserIdRequest) Pointer() *DescribeJobsByUserIdRequest {
	return &p
}

type GetJobByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	JobName         *string `json:"jobName"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetJobByUserIdRequestFromJson(data string) GetJobByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobByUserIdRequestFromDict(dict)
}

func NewGetJobByUserIdRequestFromDict(data map[string]interface{}) GetJobByUserIdRequest {
	return GetJobByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		JobName:         core.CastString(data["jobName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"jobName":         p.JobName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetJobByUserIdRequest) Pointer() *GetJobByUserIdRequest {
	return &p
}

type PushByUserIdRequest struct {
	SourceRequestId    *string    `json:"sourceRequestId"`
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	UserId             *string    `json:"userId"`
	Jobs               []JobEntry `json:"jobs"`
	TimeOffsetToken    *string    `json:"timeOffsetToken"`
}

func NewPushByUserIdRequestFromJson(data string) PushByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPushByUserIdRequestFromDict(dict)
}

func NewPushByUserIdRequestFromDict(data map[string]interface{}) PushByUserIdRequest {
	return PushByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		Jobs:            CastJobEntries(core.CastArray(data["jobs"])),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p PushByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"jobs": CastJobEntriesFromDict(
			p.Jobs,
		),
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p PushByUserIdRequest) Pointer() *PushByUserIdRequest {
	return &p
}

type RunRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
}

func NewRunRequestFromJson(data string) RunRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunRequestFromDict(dict)
}

func NewRunRequestFromDict(data map[string]interface{}) RunRequest {
	return RunRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p RunRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
	}
}

func (p RunRequest) Pointer() *RunRequest {
	return &p
}

type RunByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRunByUserIdRequestFromJson(data string) RunByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRunByUserIdRequestFromDict(dict)
}

func NewRunByUserIdRequestFromDict(data map[string]interface{}) RunByUserIdRequest {
	return RunByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p RunByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p RunByUserIdRequest) Pointer() *RunByUserIdRequest {
	return &p
}

type DeleteJobByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	JobName            *string `json:"jobName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteJobByUserIdRequestFromJson(data string) DeleteJobByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteJobByUserIdRequestFromDict(dict)
}

func NewDeleteJobByUserIdRequestFromDict(data map[string]interface{}) DeleteJobByUserIdRequest {
	return DeleteJobByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		JobName:         core.CastString(data["jobName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"jobName":         p.JobName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteJobByUserIdRequest) Pointer() *DeleteJobByUserIdRequest {
	return &p
}

type PushByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewPushByStampSheetRequestFromJson(data string) PushByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewPushByStampSheetRequestFromDict(dict)
}

func NewPushByStampSheetRequestFromDict(data map[string]interface{}) PushByStampSheetRequest {
	return PushByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p PushByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p PushByStampSheetRequest) Pointer() *PushByStampSheetRequest {
	return &p
}

type DeleteByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewDeleteByStampTaskRequestFromJson(data string) DeleteByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteByStampTaskRequestFromDict(dict)
}

func NewDeleteByStampTaskRequestFromDict(data map[string]interface{}) DeleteByStampTaskRequest {
	return DeleteByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DeleteByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DeleteByStampTaskRequest) Pointer() *DeleteByStampTaskRequest {
	return &p
}

type GetJobResultRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	JobName         *string `json:"jobName"`
	TryNumber       *int32  `json:"tryNumber"`
}

func NewGetJobResultRequestFromJson(data string) GetJobResultRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobResultRequestFromDict(dict)
}

func NewGetJobResultRequestFromDict(data map[string]interface{}) GetJobResultRequest {
	return GetJobResultRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		JobName:       core.CastString(data["jobName"]),
		TryNumber:     core.CastInt32(data["tryNumber"]),
	}
}

func (p GetJobResultRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"jobName":       p.JobName,
		"tryNumber":     p.TryNumber,
	}
}

func (p GetJobResultRequest) Pointer() *GetJobResultRequest {
	return &p
}

type GetJobResultByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	JobName         *string `json:"jobName"`
	TryNumber       *int32  `json:"tryNumber"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetJobResultByUserIdRequestFromJson(data string) GetJobResultByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetJobResultByUserIdRequestFromDict(dict)
}

func NewGetJobResultByUserIdRequestFromDict(data map[string]interface{}) GetJobResultByUserIdRequest {
	return GetJobResultByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		JobName:         core.CastString(data["jobName"]),
		TryNumber:       core.CastInt32(data["tryNumber"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetJobResultByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"jobName":         p.JobName,
		"tryNumber":       p.TryNumber,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetJobResultByUserIdRequest) Pointer() *GetJobResultByUserIdRequest {
	return &p
}

type DescribeDeadLetterJobsByUserIdRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeDeadLetterJobsByUserIdRequestFromJson(data string) DescribeDeadLetterJobsByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeDeadLetterJobsByUserIdRequestFromDict(dict)
}

func NewDescribeDeadLetterJobsByUserIdRequestFromDict(data map[string]interface{}) DescribeDeadLetterJobsByUserIdRequest {
	return DescribeDeadLetterJobsByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeDeadLetterJobsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeDeadLetterJobsByUserIdRequest) Pointer() *DescribeDeadLetterJobsByUserIdRequest {
	return &p
}

type GetDeadLetterJobByUserIdRequest struct {
	SourceRequestId   *string `json:"sourceRequestId"`
	RequestId         *string `json:"requestId"`
	ContextStack      *string `json:"contextStack"`
	NamespaceName     *string `json:"namespaceName"`
	UserId            *string `json:"userId"`
	DeadLetterJobName *string `json:"deadLetterJobName"`
	TimeOffsetToken   *string `json:"timeOffsetToken"`
}

func NewGetDeadLetterJobByUserIdRequestFromJson(data string) GetDeadLetterJobByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetDeadLetterJobByUserIdRequestFromDict(dict)
}

func NewGetDeadLetterJobByUserIdRequestFromDict(data map[string]interface{}) GetDeadLetterJobByUserIdRequest {
	return GetDeadLetterJobByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		DeadLetterJobName: core.CastString(data["deadLetterJobName"]),
		TimeOffsetToken:   core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetDeadLetterJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"userId":            p.UserId,
		"deadLetterJobName": p.DeadLetterJobName,
		"timeOffsetToken":   p.TimeOffsetToken,
	}
}

func (p GetDeadLetterJobByUserIdRequest) Pointer() *GetDeadLetterJobByUserIdRequest {
	return &p
}

type DeleteDeadLetterJobByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	DeadLetterJobName  *string `json:"deadLetterJobName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteDeadLetterJobByUserIdRequestFromJson(data string) DeleteDeadLetterJobByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteDeadLetterJobByUserIdRequestFromDict(dict)
}

func NewDeleteDeadLetterJobByUserIdRequestFromDict(data map[string]interface{}) DeleteDeadLetterJobByUserIdRequest {
	return DeleteDeadLetterJobByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		DeadLetterJobName: core.CastString(data["deadLetterJobName"]),
		TimeOffsetToken:   core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteDeadLetterJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"userId":            p.UserId,
		"deadLetterJobName": p.DeadLetterJobName,
		"timeOffsetToken":   p.TimeOffsetToken,
	}
}

func (p DeleteDeadLetterJobByUserIdRequest) Pointer() *DeleteDeadLetterJobByUserIdRequest {
	return &p
}
