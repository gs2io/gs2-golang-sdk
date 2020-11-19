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

package lottery

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
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	Name                     *string            `json:"name"`
	Description              *string            `json:"description"`
	QueueNamespaceId         *string            `json:"queueNamespaceId"`
	KeyId                    *string            `json:"keyId"`
	LotteryTriggerScriptId   *string            `json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *string            `json:"choicePrizeTableScriptId"`
	LogSetting               *LogSetting        `json:"logSetting"`
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
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	Description              *string            `json:"description"`
	QueueNamespaceId         *string            `json:"queueNamespaceId"`
	KeyId                    *string            `json:"keyId"`
	LotteryTriggerScriptId   *string            `json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *string            `json:"choicePrizeTableScriptId"`
	LogSetting               *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeLotteryModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateLotteryModelMasterRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	Name                     *string            `json:"name"`
	Description              *string            `json:"description"`
	Metadata                 *string            `json:"metadata"`
	Mode                     *string            `json:"mode"`
	Method                   *string            `json:"method"`
	PrizeTableName           *string            `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string            `json:"choicePrizeTableScriptId"`
}

type GetLotteryModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	LotteryName   *string            `json:"lotteryName"`
}

type UpdateLotteryModelMasterRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	LotteryName              *string            `json:"lotteryName"`
	Description              *string            `json:"description"`
	Metadata                 *string            `json:"metadata"`
	Mode                     *string            `json:"mode"`
	Method                   *string            `json:"method"`
	PrizeTableName           *string            `json:"prizeTableName"`
	ChoicePrizeTableScriptId *string            `json:"choicePrizeTableScriptId"`
}

type DeleteLotteryModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	LotteryName   *string            `json:"lotteryName"`
}

type DescribePrizeTableMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreatePrizeTableMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Name          *string            `json:"name"`
	Description   *string            `json:"description"`
	Metadata      *string            `json:"metadata"`
	Prizes        *[]*Prize          `json:"prizes"`
}

type GetPrizeTableMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	PrizeTableName *string            `json:"prizeTableName"`
}

type UpdatePrizeTableMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	PrizeTableName *string            `json:"prizeTableName"`
	Description    *string            `json:"description"`
	Metadata       *string            `json:"metadata"`
	Prizes         *[]*Prize          `json:"prizes"`
}

type DeletePrizeTableMasterRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	PrizeTableName *string            `json:"prizeTableName"`
}

type DescribeBoxesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeBoxesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetBoxRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PrizeTableName         *string            `json:"prizeTableName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetBoxByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PrizeTableName         *string            `json:"prizeTableName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetRawBoxByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PrizeTableName         *string            `json:"prizeTableName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ResetBoxRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PrizeTableName         *string            `json:"prizeTableName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type ResetBoxByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PrizeTableName         *string            `json:"prizeTableName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeLotteryModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetLotteryModelRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	LotteryName   *string            `json:"lotteryName"`
}

type DescribePrizeTablesRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetPrizeTableRequest struct {
	RequestId      *core.RequestId    `json:"requestId"`
	ContextStack   *core.ContextStack `json:"contextStack"`
	NamespaceName  *string            `json:"namespaceName"`
	PrizeTableName *string            `json:"prizeTableName"`
}

type DrawByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	LotteryName            *string            `json:"lotteryName"`
	UserId                 *string            `json:"userId"`
	Count                  *int32             `json:"count"`
	Config                 *[]*Config         `json:"config"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeProbabilitiesRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	LotteryName            *string            `json:"lotteryName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeProbabilitiesByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	LotteryName            *string            `json:"lotteryName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DrawByStampSheetRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	StampSheet             *string            `json:"stampSheet"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentLotteryMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentLotteryMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentLotteryMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}
