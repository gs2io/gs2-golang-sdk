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

package mission

import "core"

type DescribeCompletesRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCompletesRequestFromDict(data map[string]interface{}) DescribeCompletesRequest {
	return DescribeCompletesRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCompletesRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCompletesRequest) Pointer() *DescribeCompletesRequest {
	return &p
}

type DescribeCompletesByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCompletesByUserIdRequestFromDict(data map[string]interface{}) DescribeCompletesByUserIdRequest {
	return DescribeCompletesByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCompletesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCompletesByUserIdRequest) Pointer() *DescribeCompletesByUserIdRequest {
	return &p
}

type CompleteRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	MissionGroupName   *string  `json:"missionGroupName"`
	MissionTaskName    *string  `json:"missionTaskName"`
	AccessToken        *string  `json:"accessToken"`
	Config             []Config `json:"config"`
}

func NewCompleteRequestFromDict(data map[string]interface{}) CompleteRequest {
	return CompleteRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		AccessToken:      core.CastString(data["accessToken"]),
		Config:           CastConfigs(core.CastArray(data["config"])),
	}
}

func (p CompleteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"accessToken":      p.AccessToken,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p CompleteRequest) Pointer() *CompleteRequest {
	return &p
}

type CompleteByUserIdRequest struct {
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	MissionGroupName   *string  `json:"missionGroupName"`
	MissionTaskName    *string  `json:"missionTaskName"`
	UserId             *string  `json:"userId"`
	Config             []Config `json:"config"`
}

func NewCompleteByUserIdRequestFromDict(data map[string]interface{}) CompleteByUserIdRequest {
	return CompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		UserId:           core.CastString(data["userId"]),
		Config:           CastConfigs(core.CastArray(data["config"])),
	}
}

func (p CompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"userId":           p.UserId,
		"config": CastConfigsFromDict(
			p.Config,
		),
	}
}

func (p CompleteByUserIdRequest) Pointer() *CompleteByUserIdRequest {
	return &p
}

type ReceiveByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
	UserId             *string `json:"userId"`
}

func NewReceiveByUserIdRequestFromDict(data map[string]interface{}) ReceiveByUserIdRequest {
	return ReceiveByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		UserId:           core.CastString(data["userId"]),
	}
}

func (p ReceiveByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"userId":           p.UserId,
	}
}

func (p ReceiveByUserIdRequest) Pointer() *ReceiveByUserIdRequest {
	return &p
}

type GetCompleteRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	AccessToken        *string `json:"accessToken"`
}

func NewGetCompleteRequestFromDict(data map[string]interface{}) GetCompleteRequest {
	return GetCompleteRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		AccessToken:      core.CastString(data["accessToken"]),
	}
}

func (p GetCompleteRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"accessToken":      p.AccessToken,
	}
}

func (p GetCompleteRequest) Pointer() *GetCompleteRequest {
	return &p
}

type GetCompleteByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	UserId             *string `json:"userId"`
}

func NewGetCompleteByUserIdRequestFromDict(data map[string]interface{}) GetCompleteByUserIdRequest {
	return GetCompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		UserId:           core.CastString(data["userId"]),
	}
}

func (p GetCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"userId":           p.UserId,
	}
}

func (p GetCompleteByUserIdRequest) Pointer() *GetCompleteByUserIdRequest {
	return &p
}

type DeleteCompleteByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MissionGroupName   *string `json:"missionGroupName"`
}

func NewDeleteCompleteByUserIdRequestFromDict(data map[string]interface{}) DeleteCompleteByUserIdRequest {
	return DeleteCompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		UserId:           core.CastString(data["userId"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p DeleteCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"userId":           p.UserId,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p DeleteCompleteByUserIdRequest) Pointer() *DeleteCompleteByUserIdRequest {
	return &p
}

type ReceiveByStampTaskRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampTask          *string `json:"stampTask"`
	KeyId              *string `json:"keyId"`
}

func NewReceiveByStampTaskRequestFromDict(data map[string]interface{}) ReceiveByStampTaskRequest {
	return ReceiveByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p ReceiveByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p ReceiveByStampTaskRequest) Pointer() *ReceiveByStampTaskRequest {
	return &p
}

type DescribeCounterModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCounterModelMastersRequestFromDict(data map[string]interface{}) DescribeCounterModelMastersRequest {
	return DescribeCounterModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCounterModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCounterModelMastersRequest) Pointer() *DescribeCounterModelMastersRequest {
	return &p
}

type CreateCounterModelMasterRequest struct {
	RequestId              *string             `json:"requestId"`
	ContextStack           *string             `json:"contextStack"`
	DuplicationAvoider     *string             `json:"duplicationAvoider"`
	NamespaceName          *string             `json:"namespaceName"`
	Name                   *string             `json:"name"`
	Metadata               *string             `json:"metadata"`
	Description            *string             `json:"description"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
}

func NewCreateCounterModelMasterRequestFromDict(data map[string]interface{}) CreateCounterModelMasterRequest {
	return CreateCounterModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		Scopes:                 CastCounterScopeModels(core.CastArray(data["scopes"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p CreateCounterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"name":          p.Name,
		"metadata":      p.Metadata,
		"description":   p.Description,
		"scopes": CastCounterScopeModelsFromDict(
			p.Scopes,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p CreateCounterModelMasterRequest) Pointer() *CreateCounterModelMasterRequest {
	return &p
}

type GetCounterModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
}

func NewGetCounterModelMasterRequestFromDict(data map[string]interface{}) GetCounterModelMasterRequest {
	return GetCounterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p GetCounterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
	}
}

func (p GetCounterModelMasterRequest) Pointer() *GetCounterModelMasterRequest {
	return &p
}

type UpdateCounterModelMasterRequest struct {
	RequestId              *string             `json:"requestId"`
	ContextStack           *string             `json:"contextStack"`
	DuplicationAvoider     *string             `json:"duplicationAvoider"`
	NamespaceName          *string             `json:"namespaceName"`
	CounterName            *string             `json:"counterName"`
	Metadata               *string             `json:"metadata"`
	Description            *string             `json:"description"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
}

func NewUpdateCounterModelMasterRequestFromDict(data map[string]interface{}) UpdateCounterModelMasterRequest {
	return UpdateCounterModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		CounterName:            core.CastString(data["counterName"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		Scopes:                 CastCounterScopeModels(core.CastArray(data["scopes"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
	}
}

func (p UpdateCounterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
		"metadata":      p.Metadata,
		"description":   p.Description,
		"scopes": CastCounterScopeModelsFromDict(
			p.Scopes,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
	}
}

func (p UpdateCounterModelMasterRequest) Pointer() *UpdateCounterModelMasterRequest {
	return &p
}

type DeleteCounterModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
}

func NewDeleteCounterModelMasterRequestFromDict(data map[string]interface{}) DeleteCounterModelMasterRequest {
	return DeleteCounterModelMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p DeleteCounterModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
	}
}

func (p DeleteCounterModelMasterRequest) Pointer() *DeleteCounterModelMasterRequest {
	return &p
}

type DescribeMissionGroupModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeMissionGroupModelMastersRequestFromDict(data map[string]interface{}) DescribeMissionGroupModelMastersRequest {
	return DescribeMissionGroupModelMastersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeMissionGroupModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeMissionGroupModelMastersRequest) Pointer() *DescribeMissionGroupModelMastersRequest {
	return &p
}

type CreateMissionGroupModelMasterRequest struct {
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	Name                            *string `json:"name"`
	Metadata                        *string `json:"metadata"`
	Description                     *string `json:"description"`
	ResetType                       *string `json:"resetType"`
	ResetDayOfMonth                 *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string `json:"resetDayOfWeek"`
	ResetHour                       *int32  `json:"resetHour"`
	CompleteNotificationNamespaceId *string `json:"completeNotificationNamespaceId"`
}

func NewCreateMissionGroupModelMasterRequestFromDict(data map[string]interface{}) CreateMissionGroupModelMasterRequest {
	return CreateMissionGroupModelMasterRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		Name:                            core.CastString(data["name"]),
		Metadata:                        core.CastString(data["metadata"]),
		Description:                     core.CastString(data["description"]),
		ResetType:                       core.CastString(data["resetType"]),
		ResetDayOfMonth:                 core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:                  core.CastString(data["resetDayOfWeek"]),
		ResetHour:                       core.CastInt32(data["resetHour"]),
		CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
	}
}

func (p CreateMissionGroupModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"name":                            p.Name,
		"metadata":                        p.Metadata,
		"description":                     p.Description,
		"resetType":                       p.ResetType,
		"resetDayOfMonth":                 p.ResetDayOfMonth,
		"resetDayOfWeek":                  p.ResetDayOfWeek,
		"resetHour":                       p.ResetHour,
		"completeNotificationNamespaceId": p.CompleteNotificationNamespaceId,
	}
}

func (p CreateMissionGroupModelMasterRequest) Pointer() *CreateMissionGroupModelMasterRequest {
	return &p
}

type GetMissionGroupModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
}

func NewGetMissionGroupModelMasterRequestFromDict(data map[string]interface{}) GetMissionGroupModelMasterRequest {
	return GetMissionGroupModelMasterRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p GetMissionGroupModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p GetMissionGroupModelMasterRequest) Pointer() *GetMissionGroupModelMasterRequest {
	return &p
}

type UpdateMissionGroupModelMasterRequest struct {
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
	DuplicationAvoider              *string `json:"duplicationAvoider"`
	NamespaceName                   *string `json:"namespaceName"`
	MissionGroupName                *string `json:"missionGroupName"`
	Metadata                        *string `json:"metadata"`
	Description                     *string `json:"description"`
	ResetType                       *string `json:"resetType"`
	ResetDayOfMonth                 *int32  `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string `json:"resetDayOfWeek"`
	ResetHour                       *int32  `json:"resetHour"`
	CompleteNotificationNamespaceId *string `json:"completeNotificationNamespaceId"`
}

func NewUpdateMissionGroupModelMasterRequestFromDict(data map[string]interface{}) UpdateMissionGroupModelMasterRequest {
	return UpdateMissionGroupModelMasterRequest{
		NamespaceName:                   core.CastString(data["namespaceName"]),
		MissionGroupName:                core.CastString(data["missionGroupName"]),
		Metadata:                        core.CastString(data["metadata"]),
		Description:                     core.CastString(data["description"]),
		ResetType:                       core.CastString(data["resetType"]),
		ResetDayOfMonth:                 core.CastInt32(data["resetDayOfMonth"]),
		ResetDayOfWeek:                  core.CastString(data["resetDayOfWeek"]),
		ResetHour:                       core.CastInt32(data["resetHour"]),
		CompleteNotificationNamespaceId: core.CastString(data["completeNotificationNamespaceId"]),
	}
}

func (p UpdateMissionGroupModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":                   p.NamespaceName,
		"missionGroupName":                p.MissionGroupName,
		"metadata":                        p.Metadata,
		"description":                     p.Description,
		"resetType":                       p.ResetType,
		"resetDayOfMonth":                 p.ResetDayOfMonth,
		"resetDayOfWeek":                  p.ResetDayOfWeek,
		"resetHour":                       p.ResetHour,
		"completeNotificationNamespaceId": p.CompleteNotificationNamespaceId,
	}
}

func (p UpdateMissionGroupModelMasterRequest) Pointer() *UpdateMissionGroupModelMasterRequest {
	return &p
}

type DeleteMissionGroupModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
}

func NewDeleteMissionGroupModelMasterRequestFromDict(data map[string]interface{}) DeleteMissionGroupModelMasterRequest {
	return DeleteMissionGroupModelMasterRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p DeleteMissionGroupModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p DeleteMissionGroupModelMasterRequest) Pointer() *DeleteMissionGroupModelMasterRequest {
	return &p
}

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
	RequestId              *string              `json:"requestId"`
	ContextStack           *string              `json:"contextStack"`
	DuplicationAvoider     *string              `json:"duplicationAvoider"`
	Name                   *string              `json:"name"`
	Description            *string              `json:"description"`
	MissionCompleteScript  *ScriptSetting       `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting       `json:"counterIncrementScript"`
	ReceiveRewardsScript   *ScriptSetting       `json:"receiveRewardsScript"`
	QueueNamespaceId       *string              `json:"queueNamespaceId"`
	KeyId                  *string              `json:"keyId"`
	CompleteNotification   *NotificationSetting `json:"completeNotification"`
	LogSetting             *LogSetting          `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		MissionCompleteScript:  NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
		CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
		ReceiveRewardsScript:   NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		CompleteNotification:   NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:             NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                   p.Name,
		"description":            p.Description,
		"missionCompleteScript":  p.MissionCompleteScript.ToDict(),
		"counterIncrementScript": p.CounterIncrementScript.ToDict(),
		"receiveRewardsScript":   p.ReceiveRewardsScript.ToDict(),
		"queueNamespaceId":       p.QueueNamespaceId,
		"keyId":                  p.KeyId,
		"completeNotification":   p.CompleteNotification.ToDict(),
		"logSetting":             p.LogSetting.ToDict(),
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
	RequestId              *string              `json:"requestId"`
	ContextStack           *string              `json:"contextStack"`
	DuplicationAvoider     *string              `json:"duplicationAvoider"`
	NamespaceName          *string              `json:"namespaceName"`
	Description            *string              `json:"description"`
	MissionCompleteScript  *ScriptSetting       `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting       `json:"counterIncrementScript"`
	ReceiveRewardsScript   *ScriptSetting       `json:"receiveRewardsScript"`
	QueueNamespaceId       *string              `json:"queueNamespaceId"`
	KeyId                  *string              `json:"keyId"`
	CompleteNotification   *NotificationSetting `json:"completeNotification"`
	LogSetting             *LogSetting          `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		Description:            core.CastString(data["description"]),
		MissionCompleteScript:  NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
		CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
		ReceiveRewardsScript:   NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
		CompleteNotification:   NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:             NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"description":            p.Description,
		"missionCompleteScript":  p.MissionCompleteScript.ToDict(),
		"counterIncrementScript": p.CounterIncrementScript.ToDict(),
		"receiveRewardsScript":   p.ReceiveRewardsScript.ToDict(),
		"queueNamespaceId":       p.QueueNamespaceId,
		"keyId":                  p.KeyId,
		"completeNotification":   p.CompleteNotification.ToDict(),
		"logSetting":             p.LogSetting.ToDict(),
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

type DescribeCountersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	AccessToken        *string `json:"accessToken"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCountersRequestFromDict(data map[string]interface{}) DescribeCountersRequest {
	return DescribeCountersRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		AccessToken:   core.CastString(data["accessToken"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCountersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"accessToken":   p.AccessToken,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCountersRequest) Pointer() *DescribeCountersRequest {
	return &p
}

type DescribeCountersByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeCountersByUserIdRequestFromDict(data map[string]interface{}) DescribeCountersByUserIdRequest {
	return DescribeCountersByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		PageToken:     core.CastString(data["pageToken"]),
		Limit:         core.CastInt32(data["limit"]),
	}
}

func (p DescribeCountersByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"pageToken":     p.PageToken,
		"limit":         p.Limit,
	}
}

func (p DescribeCountersByUserIdRequest) Pointer() *DescribeCountersByUserIdRequest {
	return &p
}

type IncreaseCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
	Value              *int64  `json:"value"`
}

func NewIncreaseCounterByUserIdRequestFromDict(data map[string]interface{}) IncreaseCounterByUserIdRequest {
	return IncreaseCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
		UserId:        core.CastString(data["userId"]),
		Value:         core.CastInt64(data["value"]),
	}
}

func (p IncreaseCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
		"userId":        p.UserId,
		"value":         p.Value,
	}
}

func (p IncreaseCounterByUserIdRequest) Pointer() *IncreaseCounterByUserIdRequest {
	return &p
}

type GetCounterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
	AccessToken        *string `json:"accessToken"`
}

func NewGetCounterRequestFromDict(data map[string]interface{}) GetCounterRequest {
	return GetCounterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
		AccessToken:   core.CastString(data["accessToken"]),
	}
}

func (p GetCounterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
		"accessToken":   p.AccessToken,
	}
}

func (p GetCounterRequest) Pointer() *GetCounterRequest {
	return &p
}

type GetCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
}

func NewGetCounterByUserIdRequestFromDict(data map[string]interface{}) GetCounterByUserIdRequest {
	return GetCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
		UserId:        core.CastString(data["userId"]),
	}
}

func (p GetCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
		"userId":        p.UserId,
	}
}

func (p GetCounterByUserIdRequest) Pointer() *GetCounterByUserIdRequest {
	return &p
}

type DeleteCounterByUserIdRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	CounterName        *string `json:"counterName"`
}

func NewDeleteCounterByUserIdRequestFromDict(data map[string]interface{}) DeleteCounterByUserIdRequest {
	return DeleteCounterByUserIdRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		UserId:        core.CastString(data["userId"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p DeleteCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"userId":        p.UserId,
		"counterName":   p.CounterName,
	}
}

func (p DeleteCounterByUserIdRequest) Pointer() *DeleteCounterByUserIdRequest {
	return &p
}

type IncreaseByStampSheetRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	StampSheet         *string `json:"stampSheet"`
	KeyId              *string `json:"keyId"`
}

func NewIncreaseByStampSheetRequestFromDict(data map[string]interface{}) IncreaseByStampSheetRequest {
	return IncreaseByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p IncreaseByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p IncreaseByStampSheetRequest) Pointer() *IncreaseByStampSheetRequest {
	return &p
}

type ExportMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
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

type GetCurrentMissionMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewGetCurrentMissionMasterRequestFromDict(data map[string]interface{}) GetCurrentMissionMasterRequest {
	return GetCurrentMissionMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p GetCurrentMissionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p GetCurrentMissionMasterRequest) Pointer() *GetCurrentMissionMasterRequest {
	return &p
}

type UpdateCurrentMissionMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	Settings           *string `json:"settings"`
}

func NewUpdateCurrentMissionMasterRequestFromDict(data map[string]interface{}) UpdateCurrentMissionMasterRequest {
	return UpdateCurrentMissionMasterRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		Settings:      core.CastString(data["settings"]),
	}
}

func (p UpdateCurrentMissionMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"settings":      p.Settings,
	}
}

func (p UpdateCurrentMissionMasterRequest) Pointer() *UpdateCurrentMissionMasterRequest {
	return &p
}

type UpdateCurrentMissionMasterFromGitHubRequest struct {
	RequestId          *string                `json:"requestId"`
	ContextStack       *string                `json:"contextStack"`
	DuplicationAvoider *string                `json:"duplicationAvoider"`
	NamespaceName      *string                `json:"namespaceName"`
	CheckoutSetting    *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentMissionMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentMissionMasterFromGitHubRequest {
	return UpdateCurrentMissionMasterFromGitHubRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
	}
}

func (p UpdateCurrentMissionMasterFromGitHubRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"checkoutSetting": p.CheckoutSetting.ToDict(),
	}
}

func (p UpdateCurrentMissionMasterFromGitHubRequest) Pointer() *UpdateCurrentMissionMasterFromGitHubRequest {
	return &p
}

type DescribeCounterModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeCounterModelsRequestFromDict(data map[string]interface{}) DescribeCounterModelsRequest {
	return DescribeCounterModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeCounterModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeCounterModelsRequest) Pointer() *DescribeCounterModelsRequest {
	return &p
}

type GetCounterModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
}

func NewGetCounterModelRequestFromDict(data map[string]interface{}) GetCounterModelRequest {
	return GetCounterModelRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
		CounterName:   core.CastString(data["counterName"]),
	}
}

func (p GetCounterModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
		"counterName":   p.CounterName,
	}
}

func (p GetCounterModelRequest) Pointer() *GetCounterModelRequest {
	return &p
}

type DescribeMissionGroupModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
}

func NewDescribeMissionGroupModelsRequestFromDict(data map[string]interface{}) DescribeMissionGroupModelsRequest {
	return DescribeMissionGroupModelsRequest{
		NamespaceName: core.CastString(data["namespaceName"]),
	}
}

func (p DescribeMissionGroupModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName": p.NamespaceName,
	}
}

func (p DescribeMissionGroupModelsRequest) Pointer() *DescribeMissionGroupModelsRequest {
	return &p
}

type GetMissionGroupModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
}

func NewGetMissionGroupModelRequestFromDict(data map[string]interface{}) GetMissionGroupModelRequest {
	return GetMissionGroupModelRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p GetMissionGroupModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p GetMissionGroupModelRequest) Pointer() *GetMissionGroupModelRequest {
	return &p
}

type DescribeMissionTaskModelsRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
}

func NewDescribeMissionTaskModelsRequestFromDict(data map[string]interface{}) DescribeMissionTaskModelsRequest {
	return DescribeMissionTaskModelsRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
	}
}

func (p DescribeMissionTaskModelsRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
	}
}

func (p DescribeMissionTaskModelsRequest) Pointer() *DescribeMissionTaskModelsRequest {
	return &p
}

type GetMissionTaskModelRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
}

func NewGetMissionTaskModelRequestFromDict(data map[string]interface{}) GetMissionTaskModelRequest {
	return GetMissionTaskModelRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
	}
}

func (p GetMissionTaskModelRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
	}
}

func (p GetMissionTaskModelRequest) Pointer() *GetMissionTaskModelRequest {
	return &p
}

type DescribeMissionTaskModelMastersRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	PageToken          *string `json:"pageToken"`
	Limit              *int32  `json:"limit"`
}

func NewDescribeMissionTaskModelMastersRequestFromDict(data map[string]interface{}) DescribeMissionTaskModelMastersRequest {
	return DescribeMissionTaskModelMastersRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		PageToken:        core.CastString(data["pageToken"]),
		Limit:            core.CastInt32(data["limit"]),
	}
}

func (p DescribeMissionTaskModelMastersRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"pageToken":        p.PageToken,
		"limit":            p.Limit,
	}
}

func (p DescribeMissionTaskModelMastersRequest) Pointer() *DescribeMissionTaskModelMastersRequest {
	return &p
}

type CreateMissionTaskModelMasterRequest struct {
	RequestId              *string         `json:"requestId"`
	ContextStack           *string         `json:"contextStack"`
	DuplicationAvoider     *string         `json:"duplicationAvoider"`
	NamespaceName          *string         `json:"namespaceName"`
	MissionGroupName       *string         `json:"missionGroupName"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	Description            *string         `json:"description"`
	CounterName            *string         `json:"counterName"`
	TargetValue            *int64          `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string         `json:"premiseMissionTaskName"`
}

func NewCreateMissionTaskModelMasterRequestFromDict(data map[string]interface{}) CreateMissionTaskModelMasterRequest {
	return CreateMissionTaskModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		MissionGroupName:       core.CastString(data["missionGroupName"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		CounterName:            core.CastString(data["counterName"]),
		TargetValue:            core.CastInt64(data["targetValue"]),
		CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		PremiseMissionTaskName: core.CastString(data["premiseMissionTaskName"]),
	}
}

func (p CreateMissionTaskModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"name":             p.Name,
		"metadata":         p.Metadata,
		"description":      p.Description,
		"counterName":      p.CounterName,
		"targetValue":      p.TargetValue,
		"completeAcquireActions": CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
		"premiseMissionTaskName": p.PremiseMissionTaskName,
	}
}

func (p CreateMissionTaskModelMasterRequest) Pointer() *CreateMissionTaskModelMasterRequest {
	return &p
}

type GetMissionTaskModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
}

func NewGetMissionTaskModelMasterRequestFromDict(data map[string]interface{}) GetMissionTaskModelMasterRequest {
	return GetMissionTaskModelMasterRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
	}
}

func (p GetMissionTaskModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
	}
}

func (p GetMissionTaskModelMasterRequest) Pointer() *GetMissionTaskModelMasterRequest {
	return &p
}

type UpdateMissionTaskModelMasterRequest struct {
	RequestId              *string         `json:"requestId"`
	ContextStack           *string         `json:"contextStack"`
	DuplicationAvoider     *string         `json:"duplicationAvoider"`
	NamespaceName          *string         `json:"namespaceName"`
	MissionGroupName       *string         `json:"missionGroupName"`
	MissionTaskName        *string         `json:"missionTaskName"`
	Metadata               *string         `json:"metadata"`
	Description            *string         `json:"description"`
	CounterName            *string         `json:"counterName"`
	TargetValue            *int64          `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string         `json:"premiseMissionTaskName"`
}

func NewUpdateMissionTaskModelMasterRequestFromDict(data map[string]interface{}) UpdateMissionTaskModelMasterRequest {
	return UpdateMissionTaskModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		MissionGroupName:       core.CastString(data["missionGroupName"]),
		MissionTaskName:        core.CastString(data["missionTaskName"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		CounterName:            core.CastString(data["counterName"]),
		TargetValue:            core.CastInt64(data["targetValue"]),
		CompleteAcquireActions: CastAcquireActions(core.CastArray(data["completeAcquireActions"])),
		ChallengePeriodEventId: core.CastString(data["challengePeriodEventId"]),
		PremiseMissionTaskName: core.CastString(data["premiseMissionTaskName"]),
	}
}

func (p UpdateMissionTaskModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"metadata":         p.Metadata,
		"description":      p.Description,
		"counterName":      p.CounterName,
		"targetValue":      p.TargetValue,
		"completeAcquireActions": CastAcquireActionsFromDict(
			p.CompleteAcquireActions,
		),
		"challengePeriodEventId": p.ChallengePeriodEventId,
		"premiseMissionTaskName": p.PremiseMissionTaskName,
	}
}

func (p UpdateMissionTaskModelMasterRequest) Pointer() *UpdateMissionTaskModelMasterRequest {
	return &p
}

type DeleteMissionTaskModelMasterRequest struct {
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
}

func NewDeleteMissionTaskModelMasterRequestFromDict(data map[string]interface{}) DeleteMissionTaskModelMasterRequest {
	return DeleteMissionTaskModelMasterRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
	}
}

func (p DeleteMissionTaskModelMasterRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
	}
}

func (p DeleteMissionTaskModelMasterRequest) Pointer() *DeleteMissionTaskModelMasterRequest {
	return &p
}
