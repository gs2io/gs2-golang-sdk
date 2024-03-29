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

import (
	"encoding/json"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeCompletesRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCompletesRequestFromJson(data string) DescribeCompletesRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletesRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeCompletesByUserIdRequestFromJson(data string) DescribeCompletesByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCompletesByUserIdRequestFromDict(dict)
}

func NewDescribeCompletesByUserIdRequestFromDict(data map[string]interface{}) DescribeCompletesByUserIdRequest {
	return DescribeCompletesByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeCompletesByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeCompletesByUserIdRequest) Pointer() *DescribeCompletesByUserIdRequest {
	return &p
}

type CompleteRequest struct {
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	MissionGroupName   *string  `json:"missionGroupName"`
	MissionTaskName    *string  `json:"missionTaskName"`
	AccessToken        *string  `json:"accessToken"`
	Config             []Config `json:"config"`
}

func NewCompleteRequestFromJson(data string) CompleteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCompleteRequestFromDict(dict)
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
	SourceRequestId    *string  `json:"sourceRequestId"`
	RequestId          *string  `json:"requestId"`
	ContextStack       *string  `json:"contextStack"`
	DuplicationAvoider *string  `json:"duplicationAvoider"`
	NamespaceName      *string  `json:"namespaceName"`
	MissionGroupName   *string  `json:"missionGroupName"`
	MissionTaskName    *string  `json:"missionTaskName"`
	UserId             *string  `json:"userId"`
	Config             []Config `json:"config"`
	TimeOffsetToken    *string  `json:"timeOffsetToken"`
}

func NewCompleteByUserIdRequestFromJson(data string) CompleteByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCompleteByUserIdRequestFromDict(dict)
}

func NewCompleteByUserIdRequestFromDict(data map[string]interface{}) CompleteByUserIdRequest {
	return CompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		UserId:           core.CastString(data["userId"]),
		Config:           CastConfigs(core.CastArray(data["config"])),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
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
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p CompleteByUserIdRequest) Pointer() *CompleteByUserIdRequest {
	return &p
}

type ReceiveByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewReceiveByUserIdRequestFromJson(data string) ReceiveByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByUserIdRequestFromDict(dict)
}

func NewReceiveByUserIdRequestFromDict(data map[string]interface{}) ReceiveByUserIdRequest {
	return ReceiveByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		UserId:           core.CastString(data["userId"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p ReceiveByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"userId":           p.UserId,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p ReceiveByUserIdRequest) Pointer() *ReceiveByUserIdRequest {
	return &p
}

type RevertReceiveByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	MissionGroupName   *string `json:"missionGroupName"`
	MissionTaskName    *string `json:"missionTaskName"`
	UserId             *string `json:"userId"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewRevertReceiveByUserIdRequestFromJson(data string) RevertReceiveByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertReceiveByUserIdRequestFromDict(dict)
}

func NewRevertReceiveByUserIdRequestFromDict(data map[string]interface{}) RevertReceiveByUserIdRequest {
	return RevertReceiveByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		MissionTaskName:  core.CastString(data["missionTaskName"]),
		UserId:           core.CastString(data["userId"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p RevertReceiveByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"missionTaskName":  p.MissionTaskName,
		"userId":           p.UserId,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p RevertReceiveByUserIdRequest) Pointer() *RevertReceiveByUserIdRequest {
	return &p
}

type GetCompleteRequest struct {
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	AccessToken      *string `json:"accessToken"`
}

func NewGetCompleteRequestFromJson(data string) GetCompleteRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompleteRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	UserId           *string `json:"userId"`
	TimeOffsetToken  *string `json:"timeOffsetToken"`
}

func NewGetCompleteByUserIdRequestFromJson(data string) GetCompleteByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCompleteByUserIdRequestFromDict(dict)
}

func NewGetCompleteByUserIdRequestFromDict(data map[string]interface{}) GetCompleteByUserIdRequest {
	return GetCompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		UserId:           core.CastString(data["userId"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"missionGroupName": p.MissionGroupName,
		"userId":           p.UserId,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p GetCompleteByUserIdRequest) Pointer() *GetCompleteByUserIdRequest {
	return &p
}

type DeleteCompleteByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	MissionGroupName   *string `json:"missionGroupName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteCompleteByUserIdRequestFromJson(data string) DeleteCompleteByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCompleteByUserIdRequestFromDict(dict)
}

func NewDeleteCompleteByUserIdRequestFromDict(data map[string]interface{}) DeleteCompleteByUserIdRequest {
	return DeleteCompleteByUserIdRequest{
		NamespaceName:    core.CastString(data["namespaceName"]),
		UserId:           core.CastString(data["userId"]),
		MissionGroupName: core.CastString(data["missionGroupName"]),
		TimeOffsetToken:  core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteCompleteByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":    p.NamespaceName,
		"userId":           p.UserId,
		"missionGroupName": p.MissionGroupName,
		"timeOffsetToken":  p.TimeOffsetToken,
	}
}

func (p DeleteCompleteByUserIdRequest) Pointer() *DeleteCompleteByUserIdRequest {
	return &p
}

type ReceiveByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewReceiveByStampTaskRequestFromJson(data string) ReceiveByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewReceiveByStampTaskRequestFromDict(dict)
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

type RevertReceiveByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewRevertReceiveByStampSheetRequestFromJson(data string) RevertReceiveByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewRevertReceiveByStampSheetRequestFromDict(dict)
}

func NewRevertReceiveByStampSheetRequestFromDict(data map[string]interface{}) RevertReceiveByStampSheetRequest {
	return RevertReceiveByStampSheetRequest{
		StampSheet: core.CastString(data["stampSheet"]),
		KeyId:      core.CastString(data["keyId"]),
	}
}

func (p RevertReceiveByStampSheetRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampSheet": p.StampSheet,
		"keyId":      p.KeyId,
	}
}

func (p RevertReceiveByStampSheetRequest) Pointer() *RevertReceiveByStampSheetRequest {
	return &p
}

type DescribeCounterModelMastersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCounterModelMastersRequestFromJson(data string) DescribeCounterModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCounterModelMastersRequestFromDict(dict)
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
	SourceRequestId        *string             `json:"sourceRequestId"`
	RequestId              *string             `json:"requestId"`
	ContextStack           *string             `json:"contextStack"`
	NamespaceName          *string             `json:"namespaceName"`
	Name                   *string             `json:"name"`
	Metadata               *string             `json:"metadata"`
	Description            *string             `json:"description"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
}

func NewCreateCounterModelMasterRequestFromJson(data string) CreateCounterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateCounterModelMasterRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CounterName     *string `json:"counterName"`
}

func NewGetCounterModelMasterRequestFromJson(data string) GetCounterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterModelMasterRequestFromDict(dict)
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
	SourceRequestId        *string             `json:"sourceRequestId"`
	RequestId              *string             `json:"requestId"`
	ContextStack           *string             `json:"contextStack"`
	NamespaceName          *string             `json:"namespaceName"`
	CounterName            *string             `json:"counterName"`
	Metadata               *string             `json:"metadata"`
	Description            *string             `json:"description"`
	Scopes                 []CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string             `json:"challengePeriodEventId"`
}

func NewUpdateCounterModelMasterRequestFromJson(data string) UpdateCounterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCounterModelMasterRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CounterName     *string `json:"counterName"`
}

func NewDeleteCounterModelMasterRequestFromJson(data string) DeleteCounterModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCounterModelMasterRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeMissionGroupModelMastersRequestFromJson(data string) DescribeMissionGroupModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionGroupModelMastersRequestFromDict(dict)
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
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
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

func NewCreateMissionGroupModelMasterRequestFromJson(data string) CreateMissionGroupModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMissionGroupModelMasterRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
}

func NewGetMissionGroupModelMasterRequestFromJson(data string) GetMissionGroupModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionGroupModelMasterRequestFromDict(dict)
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
	SourceRequestId                 *string `json:"sourceRequestId"`
	RequestId                       *string `json:"requestId"`
	ContextStack                    *string `json:"contextStack"`
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

func NewUpdateMissionGroupModelMasterRequestFromJson(data string) UpdateMissionGroupModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMissionGroupModelMasterRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
}

func NewDeleteMissionGroupModelMasterRequestFromJson(data string) DeleteMissionGroupModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMissionGroupModelMasterRequestFromDict(dict)
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
	SourceRequestId        *string              `json:"sourceRequestId"`
	RequestId              *string              `json:"requestId"`
	ContextStack           *string              `json:"contextStack"`
	Name                   *string              `json:"name"`
	Description            *string              `json:"description"`
	TransactionSetting     *TransactionSetting  `json:"transactionSetting"`
	MissionCompleteScript  *ScriptSetting       `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting       `json:"counterIncrementScript"`
	ReceiveRewardsScript   *ScriptSetting       `json:"receiveRewardsScript"`
	CompleteNotification   *NotificationSetting `json:"completeNotification"`
	LogSetting             *LogSetting          `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
	return CreateNamespaceRequest{
		Name:                   core.CastString(data["name"]),
		Description:            core.CastString(data["description"]),
		TransactionSetting:     NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		MissionCompleteScript:  NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
		CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
		ReceiveRewardsScript:   NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
		CompleteNotification:   NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:             NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
	}
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"name":                   p.Name,
		"description":            p.Description,
		"transactionSetting":     p.TransactionSetting.ToDict(),
		"missionCompleteScript":  p.MissionCompleteScript.ToDict(),
		"counterIncrementScript": p.CounterIncrementScript.ToDict(),
		"receiveRewardsScript":   p.ReceiveRewardsScript.ToDict(),
		"completeNotification":   p.CompleteNotification.ToDict(),
		"logSetting":             p.LogSetting.ToDict(),
		"queueNamespaceId":       p.QueueNamespaceId,
		"keyId":                  p.KeyId,
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
	SourceRequestId        *string              `json:"sourceRequestId"`
	RequestId              *string              `json:"requestId"`
	ContextStack           *string              `json:"contextStack"`
	NamespaceName          *string              `json:"namespaceName"`
	Description            *string              `json:"description"`
	TransactionSetting     *TransactionSetting  `json:"transactionSetting"`
	MissionCompleteScript  *ScriptSetting       `json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting       `json:"counterIncrementScript"`
	ReceiveRewardsScript   *ScriptSetting       `json:"receiveRewardsScript"`
	CompleteNotification   *NotificationSetting `json:"completeNotification"`
	LogSetting             *LogSetting          `json:"logSetting"`
	// Deprecated: should not be used
	QueueNamespaceId *string `json:"queueNamespaceId"`
	// Deprecated: should not be used
	KeyId *string `json:"keyId"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
	return UpdateNamespaceRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		Description:            core.CastString(data["description"]),
		TransactionSetting:     NewTransactionSettingFromDict(core.CastMap(data["transactionSetting"])).Pointer(),
		MissionCompleteScript:  NewScriptSettingFromDict(core.CastMap(data["missionCompleteScript"])).Pointer(),
		CounterIncrementScript: NewScriptSettingFromDict(core.CastMap(data["counterIncrementScript"])).Pointer(),
		ReceiveRewardsScript:   NewScriptSettingFromDict(core.CastMap(data["receiveRewardsScript"])).Pointer(),
		CompleteNotification:   NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
		LogSetting:             NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
		QueueNamespaceId:       core.CastString(data["queueNamespaceId"]),
		KeyId:                  core.CastString(data["keyId"]),
	}
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":          p.NamespaceName,
		"description":            p.Description,
		"transactionSetting":     p.TransactionSetting.ToDict(),
		"missionCompleteScript":  p.MissionCompleteScript.ToDict(),
		"counterIncrementScript": p.CounterIncrementScript.ToDict(),
		"receiveRewardsScript":   p.ReceiveRewardsScript.ToDict(),
		"completeNotification":   p.CompleteNotification.ToDict(),
		"logSetting":             p.LogSetting.ToDict(),
		"queueNamespaceId":       p.QueueNamespaceId,
		"keyId":                  p.KeyId,
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

type DescribeCountersRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	AccessToken     *string `json:"accessToken"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
}

func NewDescribeCountersRequestFromJson(data string) DescribeCountersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	UserId          *string `json:"userId"`
	PageToken       *string `json:"pageToken"`
	Limit           *int32  `json:"limit"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewDescribeCountersByUserIdRequestFromJson(data string) DescribeCountersByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCountersByUserIdRequestFromDict(dict)
}

func NewDescribeCountersByUserIdRequestFromDict(data map[string]interface{}) DescribeCountersByUserIdRequest {
	return DescribeCountersByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		PageToken:       core.CastString(data["pageToken"]),
		Limit:           core.CastInt32(data["limit"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DescribeCountersByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"pageToken":       p.PageToken,
		"limit":           p.Limit,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DescribeCountersByUserIdRequest) Pointer() *DescribeCountersByUserIdRequest {
	return &p
}

type IncreaseCounterByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
	Value              *int64  `json:"value"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewIncreaseCounterByUserIdRequestFromJson(data string) IncreaseCounterByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseCounterByUserIdRequestFromDict(dict)
}

func NewIncreaseCounterByUserIdRequestFromDict(data map[string]interface{}) IncreaseCounterByUserIdRequest {
	return IncreaseCounterByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CounterName:     core.CastString(data["counterName"]),
		UserId:          core.CastString(data["userId"]),
		Value:           core.CastInt64(data["value"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p IncreaseCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"counterName":     p.CounterName,
		"userId":          p.UserId,
		"value":           p.Value,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p IncreaseCounterByUserIdRequest) Pointer() *IncreaseCounterByUserIdRequest {
	return &p
}

type DecreaseCounterByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	CounterName        *string `json:"counterName"`
	UserId             *string `json:"userId"`
	Value              *int64  `json:"value"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDecreaseCounterByUserIdRequestFromJson(data string) DecreaseCounterByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseCounterByUserIdRequestFromDict(dict)
}

func NewDecreaseCounterByUserIdRequestFromDict(data map[string]interface{}) DecreaseCounterByUserIdRequest {
	return DecreaseCounterByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CounterName:     core.CastString(data["counterName"]),
		UserId:          core.CastString(data["userId"]),
		Value:           core.CastInt64(data["value"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DecreaseCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"counterName":     p.CounterName,
		"userId":          p.UserId,
		"value":           p.Value,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DecreaseCounterByUserIdRequest) Pointer() *DecreaseCounterByUserIdRequest {
	return &p
}

type GetCounterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CounterName     *string `json:"counterName"`
	AccessToken     *string `json:"accessToken"`
}

func NewGetCounterRequestFromJson(data string) GetCounterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CounterName     *string `json:"counterName"`
	UserId          *string `json:"userId"`
	TimeOffsetToken *string `json:"timeOffsetToken"`
}

func NewGetCounterByUserIdRequestFromJson(data string) GetCounterByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterByUserIdRequestFromDict(dict)
}

func NewGetCounterByUserIdRequestFromDict(data map[string]interface{}) GetCounterByUserIdRequest {
	return GetCounterByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		CounterName:     core.CastString(data["counterName"]),
		UserId:          core.CastString(data["userId"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p GetCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"counterName":     p.CounterName,
		"userId":          p.UserId,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p GetCounterByUserIdRequest) Pointer() *GetCounterByUserIdRequest {
	return &p
}

type DeleteCounterByUserIdRequest struct {
	SourceRequestId    *string `json:"sourceRequestId"`
	RequestId          *string `json:"requestId"`
	ContextStack       *string `json:"contextStack"`
	DuplicationAvoider *string `json:"duplicationAvoider"`
	NamespaceName      *string `json:"namespaceName"`
	UserId             *string `json:"userId"`
	CounterName        *string `json:"counterName"`
	TimeOffsetToken    *string `json:"timeOffsetToken"`
}

func NewDeleteCounterByUserIdRequestFromJson(data string) DeleteCounterByUserIdRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteCounterByUserIdRequestFromDict(dict)
}

func NewDeleteCounterByUserIdRequestFromDict(data map[string]interface{}) DeleteCounterByUserIdRequest {
	return DeleteCounterByUserIdRequest{
		NamespaceName:   core.CastString(data["namespaceName"]),
		UserId:          core.CastString(data["userId"]),
		CounterName:     core.CastString(data["counterName"]),
		TimeOffsetToken: core.CastString(data["timeOffsetToken"]),
	}
}

func (p DeleteCounterByUserIdRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"namespaceName":   p.NamespaceName,
		"userId":          p.UserId,
		"counterName":     p.CounterName,
		"timeOffsetToken": p.TimeOffsetToken,
	}
}

func (p DeleteCounterByUserIdRequest) Pointer() *DeleteCounterByUserIdRequest {
	return &p
}

type IncreaseByStampSheetRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampSheet      *string `json:"stampSheet"`
	KeyId           *string `json:"keyId"`
}

func NewIncreaseByStampSheetRequestFromJson(data string) IncreaseByStampSheetRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewIncreaseByStampSheetRequestFromDict(dict)
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

type DecreaseByStampTaskRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	StampTask       *string `json:"stampTask"`
	KeyId           *string `json:"keyId"`
}

func NewDecreaseByStampTaskRequestFromJson(data string) DecreaseByStampTaskRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDecreaseByStampTaskRequestFromDict(dict)
}

func NewDecreaseByStampTaskRequestFromDict(data map[string]interface{}) DecreaseByStampTaskRequest {
	return DecreaseByStampTaskRequest{
		StampTask: core.CastString(data["stampTask"]),
		KeyId:     core.CastString(data["keyId"]),
	}
}

func (p DecreaseByStampTaskRequest) ToDict() map[string]interface{} {
	return map[string]interface{}{
		"stampTask": p.StampTask,
		"keyId":     p.KeyId,
	}
}

func (p DecreaseByStampTaskRequest) Pointer() *DecreaseByStampTaskRequest {
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

type GetCurrentMissionMasterRequest struct {
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewGetCurrentMissionMasterRequestFromJson(data string) GetCurrentMissionMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCurrentMissionMasterRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	Settings        *string `json:"settings"`
}

func NewUpdateCurrentMissionMasterRequestFromJson(data string) UpdateCurrentMissionMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMissionMasterRequestFromDict(dict)
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
	SourceRequestId *string                `json:"sourceRequestId"`
	RequestId       *string                `json:"requestId"`
	ContextStack    *string                `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentMissionMasterFromGitHubRequestFromJson(data string) UpdateCurrentMissionMasterFromGitHubRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateCurrentMissionMasterFromGitHubRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeCounterModelsRequestFromJson(data string) DescribeCounterModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeCounterModelsRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
	CounterName     *string `json:"counterName"`
}

func NewGetCounterModelRequestFromJson(data string) GetCounterModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetCounterModelRequestFromDict(dict)
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
	SourceRequestId *string `json:"sourceRequestId"`
	RequestId       *string `json:"requestId"`
	ContextStack    *string `json:"contextStack"`
	NamespaceName   *string `json:"namespaceName"`
}

func NewDescribeMissionGroupModelsRequestFromJson(data string) DescribeMissionGroupModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionGroupModelsRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
}

func NewGetMissionGroupModelRequestFromJson(data string) GetMissionGroupModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionGroupModelRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
}

func NewDescribeMissionTaskModelsRequestFromJson(data string) DescribeMissionTaskModelsRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionTaskModelsRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	MissionTaskName  *string `json:"missionTaskName"`
}

func NewGetMissionTaskModelRequestFromJson(data string) GetMissionTaskModelRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionTaskModelRequestFromDict(dict)
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	PageToken        *string `json:"pageToken"`
	Limit            *int32  `json:"limit"`
}

func NewDescribeMissionTaskModelMastersRequestFromJson(data string) DescribeMissionTaskModelMastersRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDescribeMissionTaskModelMastersRequestFromDict(dict)
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
	SourceRequestId        *string         `json:"sourceRequestId"`
	RequestId              *string         `json:"requestId"`
	ContextStack           *string         `json:"contextStack"`
	NamespaceName          *string         `json:"namespaceName"`
	MissionGroupName       *string         `json:"missionGroupName"`
	Name                   *string         `json:"name"`
	Metadata               *string         `json:"metadata"`
	Description            *string         `json:"description"`
	CounterName            *string         `json:"counterName"`
	TargetResetType        *string         `json:"targetResetType"`
	TargetValue            *int64          `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string         `json:"premiseMissionTaskName"`
}

func NewCreateMissionTaskModelMasterRequestFromJson(data string) CreateMissionTaskModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewCreateMissionTaskModelMasterRequestFromDict(dict)
}

func NewCreateMissionTaskModelMasterRequestFromDict(data map[string]interface{}) CreateMissionTaskModelMasterRequest {
	return CreateMissionTaskModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		MissionGroupName:       core.CastString(data["missionGroupName"]),
		Name:                   core.CastString(data["name"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		CounterName:            core.CastString(data["counterName"]),
		TargetResetType:        core.CastString(data["targetResetType"]),
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
		"targetResetType":  p.TargetResetType,
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	MissionTaskName  *string `json:"missionTaskName"`
}

func NewGetMissionTaskModelMasterRequestFromJson(data string) GetMissionTaskModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewGetMissionTaskModelMasterRequestFromDict(dict)
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
	SourceRequestId        *string         `json:"sourceRequestId"`
	RequestId              *string         `json:"requestId"`
	ContextStack           *string         `json:"contextStack"`
	NamespaceName          *string         `json:"namespaceName"`
	MissionGroupName       *string         `json:"missionGroupName"`
	MissionTaskName        *string         `json:"missionTaskName"`
	Metadata               *string         `json:"metadata"`
	Description            *string         `json:"description"`
	CounterName            *string         `json:"counterName"`
	TargetResetType        *string         `json:"targetResetType"`
	TargetValue            *int64          `json:"targetValue"`
	CompleteAcquireActions []AcquireAction `json:"completeAcquireActions"`
	ChallengePeriodEventId *string         `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string         `json:"premiseMissionTaskName"`
}

func NewUpdateMissionTaskModelMasterRequestFromJson(data string) UpdateMissionTaskModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewUpdateMissionTaskModelMasterRequestFromDict(dict)
}

func NewUpdateMissionTaskModelMasterRequestFromDict(data map[string]interface{}) UpdateMissionTaskModelMasterRequest {
	return UpdateMissionTaskModelMasterRequest{
		NamespaceName:          core.CastString(data["namespaceName"]),
		MissionGroupName:       core.CastString(data["missionGroupName"]),
		MissionTaskName:        core.CastString(data["missionTaskName"]),
		Metadata:               core.CastString(data["metadata"]),
		Description:            core.CastString(data["description"]),
		CounterName:            core.CastString(data["counterName"]),
		TargetResetType:        core.CastString(data["targetResetType"]),
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
		"targetResetType":  p.TargetResetType,
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
	SourceRequestId  *string `json:"sourceRequestId"`
	RequestId        *string `json:"requestId"`
	ContextStack     *string `json:"contextStack"`
	NamespaceName    *string `json:"namespaceName"`
	MissionGroupName *string `json:"missionGroupName"`
	MissionTaskName  *string `json:"missionTaskName"`
}

func NewDeleteMissionTaskModelMasterRequestFromJson(data string) DeleteMissionTaskModelMasterRequest {
	dict := map[string]interface{}{}
	_ = json.Unmarshal([]byte(data), &dict)
	return NewDeleteMissionTaskModelMasterRequestFromDict(dict)
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
