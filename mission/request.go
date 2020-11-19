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

type DescribeCountersRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeCountersByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type IncreaseCounterByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CounterName            *string            `json:"counterName"`
	UserId                 *string            `json:"userId"`
	Value                  *int64             `json:"value"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetCounterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CounterName            *string            `json:"counterName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetCounterByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CounterName            *string            `json:"counterName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteCounterByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	CounterName            *string            `json:"counterName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type IncreaseByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId              *core.RequestId      `json:"requestId"`
	ContextStack           *core.ContextStack   `json:"contextStack"`
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
	RequestId              *core.RequestId      `json:"requestId"`
	ContextStack           *core.ContextStack   `json:"contextStack"`
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

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeCompletesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeCompletesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CompleteRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	MissionTaskName        *string            `json:"missionTaskName"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type CompleteByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	MissionTaskName        *string            `json:"missionTaskName"`
	UserId                 *string            `json:"userId"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ReceiveByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	MissionTaskName        *string            `json:"missionTaskName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetCompleteRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetCompleteByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteCompleteByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	MissionGroupName       *string            `json:"missionGroupName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ReceiveByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeMissionTaskModelsRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
}

type GetMissionTaskModelRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
	MissionTaskName  *string            `json:"missionTaskName"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentMissionMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentMissionMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentMissionMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeMissionGroupModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateMissionGroupModelMasterRequest struct {
	RequestId                       *core.RequestId    `json:"requestId"`
	ContextStack                    *core.ContextStack `json:"contextStack"`
	NamespaceName                   *string            `json:"namespaceName"`
	Name                            *string            `json:"name"`
	Metadata                        *string            `json:"metadata"`
	Description                     *string            `json:"description"`
	ResetType                       *string            `json:"resetType"`
	ResetDayOfMonth                 *int32             `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string            `json:"resetDayOfWeek"`
	ResetHour                       *int32             `json:"resetHour"`
	CompleteNotificationNamespaceId *string            `json:"completeNotificationNamespaceId"`
}

type GetMissionGroupModelMasterRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
}

type UpdateMissionGroupModelMasterRequest struct {
	RequestId                       *core.RequestId    `json:"requestId"`
	ContextStack                    *core.ContextStack `json:"contextStack"`
	NamespaceName                   *string            `json:"namespaceName"`
	MissionGroupName                *string            `json:"missionGroupName"`
	Metadata                        *string            `json:"metadata"`
	Description                     *string            `json:"description"`
	ResetType                       *string            `json:"resetType"`
	ResetDayOfMonth                 *int32             `json:"resetDayOfMonth"`
	ResetDayOfWeek                  *string            `json:"resetDayOfWeek"`
	ResetHour                       *int32             `json:"resetHour"`
	CompleteNotificationNamespaceId *string            `json:"completeNotificationNamespaceId"`
}

type DeleteMissionGroupModelMasterRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
}

type DescribeMissionTaskModelMastersRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
	PageToken        *string            `json:"pageToken"`
	Limit            *int64             `json:"limit"`
}

type CreateMissionTaskModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	Name                   *string            `json:"name"`
	Metadata               *string            `json:"metadata"`
	Description            *string            `json:"description"`
	CounterName            *string            `json:"counterName"`
	TargetValue            *int64             `json:"targetValue"`
	CompleteAcquireActions *[]*AcquireAction  `json:"completeAcquireActions"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string            `json:"premiseMissionTaskName"`
}

type GetMissionTaskModelMasterRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
	MissionTaskName  *string            `json:"missionTaskName"`
}

type UpdateMissionTaskModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	MissionGroupName       *string            `json:"missionGroupName"`
	MissionTaskName        *string            `json:"missionTaskName"`
	Metadata               *string            `json:"metadata"`
	Description            *string            `json:"description"`
	CounterName            *string            `json:"counterName"`
	TargetValue            *int64             `json:"targetValue"`
	CompleteAcquireActions *[]*AcquireAction  `json:"completeAcquireActions"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
	PremiseMissionTaskName *string            `json:"premiseMissionTaskName"`
}

type DeleteMissionTaskModelMasterRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
	MissionTaskName  *string            `json:"missionTaskName"`
}

type DescribeCounterModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateCounterModelMasterRequest struct {
	RequestId              *core.RequestId       `json:"requestId"`
	ContextStack           *core.ContextStack    `json:"contextStack"`
	NamespaceName          *string               `json:"namespaceName"`
	Name                   *string               `json:"name"`
	Metadata               *string               `json:"metadata"`
	Description            *string               `json:"description"`
	Scopes                 *[]*CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string               `json:"challengePeriodEventId"`
}

type GetCounterModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CounterName   *string            `json:"counterName"`
}

type UpdateCounterModelMasterRequest struct {
	RequestId              *core.RequestId       `json:"requestId"`
	ContextStack           *core.ContextStack    `json:"contextStack"`
	NamespaceName          *string               `json:"namespaceName"`
	CounterName            *string               `json:"counterName"`
	Metadata               *string               `json:"metadata"`
	Description            *string               `json:"description"`
	Scopes                 *[]*CounterScopeModel `json:"scopes"`
	ChallengePeriodEventId *string               `json:"challengePeriodEventId"`
}

type DeleteCounterModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CounterName   *string            `json:"counterName"`
}

type DescribeMissionGroupModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetMissionGroupModelRequest struct {
	RequestId        *core.RequestId    `json:"requestId"`
	ContextStack     *core.ContextStack `json:"contextStack"`
	NamespaceName    *string            `json:"namespaceName"`
	MissionGroupName *string            `json:"missionGroupName"`
}

type DescribeCounterModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCounterModelRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CounterName   *string            `json:"counterName"`
}
