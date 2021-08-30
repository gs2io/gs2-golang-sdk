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

import "core"

type DescribeNamespacesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
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
	PushNotification   *NotificationSetting `json:"pushNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:             core.CastString(data["name"]),
		Description:      core.CastString(data["description"]),
		PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":             p.Name,
		"description":      p.Description,
		"pushNotification": p.PushNotification.ToDict(),
		"logSetting":       p.LogSetting.ToDict(),
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
	PushNotification   *NotificationSetting `json:"pushNotification"`
	LogSetting         *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		Description:      core.CastString(data["description"]),
		PushNotification: NewNotificationSettingFromDict(core.CastMap(data["pushNotification"])).Pointer(),
		LogSetting:       NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"description":      p.Description,
		"pushNotification": p.PushNotification.ToDict(),
		"logSetting":       p.LogSetting.ToDict(),
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

type DescribeJobsByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeJobsByUserIdRequestFromDict(data map[string]interface{}) DescribeJobsByUserIdRequest {
	return DescribeJobsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeJobsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeJobsByUserIdRequest) Pointer() *DescribeJobsByUserIdRequest {
	return &p
}

type GetJobByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	JobName            *string `json:"jobName"`
}

func NewGetJobByUserIdRequestFromDict(data map[string]interface{}) GetJobByUserIdRequest {
	return GetJobByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		JobName:       core.CastString(data["jobName"]),
	}
}

func (p GetJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"jobName":       p.JobName,
	}
}

func (p GetJobByUserIdRequest) Pointer() *GetJobByUserIdRequest {
	return &p
}

type PushByUserIdRequest struct {
	RequestId          *string    `json:"requestId"`
	ContextStack       *string    `json:"contextStack"`
	DuplicationAvoider *string    `json:"duplicationAvoider"`
	NamespaceName      *string    `json:"namespaceName"`
	UserId             *string    `json:"userId"`
	Jobs               []JobEntry `json:"jobs"`
}

func NewPushByUserIdRequestFromDict(data map[string]interface{}) PushByUserIdRequest {
	return PushByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		Jobs:          CastJobEntries(core.CastArray(data["jobs"])),
	}
}

func (p PushByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"jobs": CastJobEntriesFromDict(
			p.Jobs,
		),
	}
}

func (p PushByUserIdRequest) Pointer() *PushByUserIdRequest {
	return &p
}

type RunRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
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
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
}

func NewRunByUserIdRequestFromDict(data map[string]interface{}) RunByUserIdRequest {
	return RunByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p RunByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
	}
}

func (p RunByUserIdRequest) Pointer() *RunByUserIdRequest {
	return &p
}

type DeleteJobByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	JobName            *string `json:"jobName"`
}

func NewDeleteJobByUserIdRequestFromDict(data map[string]interface{}) DeleteJobByUserIdRequest {
	return DeleteJobByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		JobName:       core.CastString(data["jobName"]),
	}
}

func (p DeleteJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"jobName":       p.JobName,
	}
}

func (p DeleteJobByUserIdRequest) Pointer() *DeleteJobByUserIdRequest {
	return &p
}

type PushByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
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

type DescribeDeadLetterJobsByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeDeadLetterJobsByUserIdRequestFromDict(data map[string]interface{}) DescribeDeadLetterJobsByUserIdRequest {
	return DescribeDeadLetterJobsByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeDeadLetterJobsByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeDeadLetterJobsByUserIdRequest) Pointer() *DescribeDeadLetterJobsByUserIdRequest {
	return &p
}

type GetDeadLetterJobByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	DeadLetterJobName  *string `json:"deadLetterJobName"`
}

func NewGetDeadLetterJobByUserIdRequestFromDict(data map[string]interface{}) GetDeadLetterJobByUserIdRequest {
	return GetDeadLetterJobByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		DeadLetterJobName: core.CastString(data["deadLetterJobName"]),
	}
}

func (p GetDeadLetterJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"userId":            p.UserId,
		"deadLetterJobName": p.DeadLetterJobName,
	}
}

func (p GetDeadLetterJobByUserIdRequest) Pointer() *GetDeadLetterJobByUserIdRequest {
	return &p
}

type DeleteDeadLetterJobByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	DeadLetterJobName  *string `json:"deadLetterJobName"`
}

func NewDeleteDeadLetterJobByUserIdRequestFromDict(data map[string]interface{}) DeleteDeadLetterJobByUserIdRequest {
	return DeleteDeadLetterJobByUserIdRequest{
		NamespaceName:     core.CastString(data["namespaceName"]),
		UserId:            core.CastString(data["userId"]),
		DeadLetterJobName: core.CastString(data["deadLetterJobName"]),
	}
}

func (p DeleteDeadLetterJobByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":     p.NamespaceName,
		"userId":            p.UserId,
		"deadLetterJobName": p.DeadLetterJobName,
	}
}

func (p DeleteDeadLetterJobByUserIdRequest) Pointer() *DeleteDeadLetterJobByUserIdRequest {
	return &p
}
