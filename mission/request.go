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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeMissionTaskModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateMissionTaskModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	Name *core.String	`json:"name"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	CounterName *core.String	`json:"counterName"`
	TargetValue *int64	`json:"targetValue"`
	CompleteAcquireActions *[]*AcquireAction	`json:"completeAcquireActions"`
	ChallengePeriodEventId *core.String	`json:"challengePeriodEventId"`
	PremiseMissionTaskName *core.String	`json:"premiseMissionTaskName"`
}

type GetMissionTaskModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
}

type UpdateMissionTaskModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	CounterName *core.String	`json:"counterName"`
	TargetValue *int64	`json:"targetValue"`
	CompleteAcquireActions *[]*AcquireAction	`json:"completeAcquireActions"`
	ChallengePeriodEventId *core.String	`json:"challengePeriodEventId"`
	PremiseMissionTaskName *core.String	`json:"premiseMissionTaskName"`
}

type DeleteMissionTaskModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
}

type DescribeCounterModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetCounterModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
}

type DescribeCountersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeCountersByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type IncreaseCounterByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
	UserId *core.String	`json:"userId"`
	Value *int64	`json:"value"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetCounterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetCounterByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DeleteCounterByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	CounterName *core.String	`json:"counterName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type IncreaseByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *core.String	`json:"stampSheet"`
	KeyId *core.String	`json:"keyId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DescribeCounterModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateCounterModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	Scopes *[]*CounterScopeModel	`json:"scopes"`
	ChallengePeriodEventId *core.String	`json:"challengePeriodEventId"`
}

type GetCounterModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
}

type UpdateCounterModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	Scopes *[]*CounterScopeModel	`json:"scopes"`
	ChallengePeriodEventId *core.String	`json:"challengePeriodEventId"`
}

type DeleteCounterModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CounterName *core.String	`json:"counterName"`
}

type DescribeMissionGroupModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetMissionGroupModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetCurrentMissionMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateCurrentMissionMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Settings *core.String	`json:"settings"`
}

type UpdateCurrentMissionMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeMissionGroupModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateMissionGroupModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	ResetType *core.String	`json:"resetType"`
	ResetDayOfMonth *int32	`json:"resetDayOfMonth"`
	ResetDayOfWeek *core.String	`json:"resetDayOfWeek"`
	ResetHour *int32	`json:"resetHour"`
	CompleteNotificationNamespaceId *core.String	`json:"completeNotificationNamespaceId"`
}

type GetMissionGroupModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
}

type UpdateMissionGroupModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	Metadata *core.String	`json:"metadata"`
	Description *core.String	`json:"description"`
	ResetType *core.String	`json:"resetType"`
	ResetDayOfMonth *int32	`json:"resetDayOfMonth"`
	ResetDayOfWeek *core.String	`json:"resetDayOfWeek"`
	ResetHour *int32	`json:"resetHour"`
	CompleteNotificationNamespaceId *core.String	`json:"completeNotificationNamespaceId"`
}

type DeleteMissionGroupModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
}

type DescribeCompletesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeCompletesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type CompleteRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
	Config *[]*Config	`json:"config"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CompleteByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
	UserId *core.String	`json:"userId"`
	Config *[]*Config	`json:"config"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type ReceiveByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetCompleteRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetCompleteByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DeleteCompleteByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type ReceiveByStampTaskRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampTask *core.String	`json:"stampTask"`
	KeyId *core.String	`json:"keyId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

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
	MissionCompleteScript *ScriptSetting	`json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting	`json:"counterIncrementScript"`
	ReceiveRewardsScript *ScriptSetting	`json:"receiveRewardsScript"`
	QueueNamespaceId *core.String	`json:"queueNamespaceId"`
	KeyId *core.String	`json:"keyId"`
	CompleteNotification *NotificationSetting	`json:"completeNotification"`
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
	MissionCompleteScript *ScriptSetting	`json:"missionCompleteScript"`
	CounterIncrementScript *ScriptSetting	`json:"counterIncrementScript"`
	ReceiveRewardsScript *ScriptSetting	`json:"receiveRewardsScript"`
	QueueNamespaceId *core.String	`json:"queueNamespaceId"`
	KeyId *core.String	`json:"keyId"`
	CompleteNotification *NotificationSetting	`json:"completeNotification"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type DescribeMissionTaskModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
}

type GetMissionTaskModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	MissionGroupName *core.String	`json:"missionGroupName"`
	MissionTaskName *core.String	`json:"missionTaskName"`
}
