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

package quest

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	PageToken    *string            `json:"pageToken"`
	Limit        *int64             `json:"limit"`
}

type CreateNamespaceRequest struct {
	RequestId           *core.RequestId    `json:"requestId"`
	ContextStack        *core.ContextStack `json:"contextStack"`
	Name                *string            `json:"name"`
	Description         *string            `json:"description"`
	StartQuestScript    *ScriptSetting     `json:"startQuestScript"`
	CompleteQuestScript *ScriptSetting     `json:"completeQuestScript"`
	FailedQuestScript   *ScriptSetting     `json:"failedQuestScript"`
	QueueNamespaceId    *string            `json:"queueNamespaceId"`
	KeyId               *string            `json:"keyId"`
	LogSetting          *LogSetting        `json:"logSetting"`
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
	RequestId           *core.RequestId    `json:"requestId"`
	ContextStack        *core.ContextStack `json:"contextStack"`
	NamespaceName       *string            `json:"namespaceName"`
	Description         *string            `json:"description"`
	StartQuestScript    *ScriptSetting     `json:"startQuestScript"`
	CompleteQuestScript *ScriptSetting     `json:"completeQuestScript"`
	FailedQuestScript   *ScriptSetting     `json:"failedQuestScript"`
	QueueNamespaceId    *string            `json:"queueNamespaceId"`
	KeyId               *string            `json:"keyId"`
	LogSetting          *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeQuestGroupModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateQuestGroupModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Name                   *string            `json:"name"`
	Description            *string            `json:"description"`
	Metadata               *string            `json:"metadata"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
}

type GetQuestGroupModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
}

type UpdateQuestGroupModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	Description            *string            `json:"description"`
	Metadata               *string            `json:"metadata"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
}

type DeleteQuestGroupModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
}

type DescribeQuestModelMastersRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
	PageToken      *string            `json:"pageToken"`
	Limit          *int64             `json:"limit"`
}

type CreateQuestModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	Name                   *string            `json:"name"`
	Description            *string            `json:"description"`
	Metadata               *string            `json:"metadata"`
	Contents               *[]*Contents       `json:"contents"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
	ConsumeActions         *[]*ConsumeAction  `json:"consumeActions"`
	FailedAcquireActions   *[]*AcquireAction  `json:"failedAcquireActions"`
	PremiseQuestNames      *[]string          `json:"premiseQuestNames"`
}

type GetQuestModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
	QuestName      *string            `json:"questName"`
}

type UpdateQuestModelMasterRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	QuestName              *string            `json:"questName"`
	Description            *string            `json:"description"`
	Metadata               *string            `json:"metadata"`
	Contents               *[]*Contents       `json:"contents"`
	ChallengePeriodEventId *string            `json:"challengePeriodEventId"`
	ConsumeActions         *[]*ConsumeAction  `json:"consumeActions"`
	FailedAcquireActions   *[]*AcquireAction  `json:"failedAcquireActions"`
	PremiseQuestNames      *[]string          `json:"premiseQuestNames"`
}

type DeleteQuestModelMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
	QuestName      *string            `json:"questName"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentQuestMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentQuestMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentQuestMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeProgressesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	QuestModelId           *string            `json:"questModelId"`
	Force                  *bool              `json:"force"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetProgressRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type StartRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	QuestName              *string            `json:"questName"`
	Force                  *bool              `json:"force"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type StartByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	QuestName              *string            `json:"questName"`
	UserId                 *string            `json:"userId"`
	Force                  *bool              `json:"force"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type EndRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TransactionId          *string            `json:"transactionId"`
	Rewards                *[]*Reward         `json:"rewards"`
	IsComplete             *bool              `json:"isComplete"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type EndByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TransactionId          *string            `json:"transactionId"`
	Rewards                *[]*Reward         `json:"rewards"`
	IsComplete             *bool              `json:"isComplete"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteProgressRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteProgressByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type CreateProgressByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteProgressByStampTaskRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampTask              *string            `json:"stampTask"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeCompletedQuestListsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeCompletedQuestListsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetCompletedQuestListRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetCompletedQuestListByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteCompletedQuestListByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	QuestGroupName         *string            `json:"questGroupName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeQuestGroupModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetQuestGroupModelRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
}

type DescribeQuestModelsRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
}

type GetQuestModelRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	QuestGroupName *string            `json:"questGroupName"`
	QuestName      *string            `json:"questName"`
}
