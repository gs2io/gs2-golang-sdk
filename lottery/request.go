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
	QueueNamespaceId *core.String	`json:"queueNamespaceId"`
	KeyId *core.String	`json:"keyId"`
	LotteryTriggerScriptId *core.String	`json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *core.String	`json:"choicePrizeTableScriptId"`
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
	QueueNamespaceId *core.String	`json:"queueNamespaceId"`
	KeyId *core.String	`json:"keyId"`
	LotteryTriggerScriptId *core.String	`json:"lotteryTriggerScriptId"`
	ChoicePrizeTableScriptId *core.String	`json:"choicePrizeTableScriptId"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type DescribeLotteryModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateLotteryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Mode *core.String	`json:"mode"`
	Method *core.String	`json:"method"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	ChoicePrizeTableScriptId *core.String	`json:"choicePrizeTableScriptId"`
}

type GetLotteryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
}

type UpdateLotteryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Mode *core.String	`json:"mode"`
	Method *core.String	`json:"method"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	ChoicePrizeTableScriptId *core.String	`json:"choicePrizeTableScriptId"`
}

type DeleteLotteryModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
}

type DescribePrizeTableMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreatePrizeTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Prizes *[]*Prize	`json:"prizes"`
}

type GetPrizeTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
}

type UpdatePrizeTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Prizes *[]*Prize	`json:"prizes"`
}

type DeletePrizeTableMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
}

type DescribeBoxesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeBoxesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetBoxRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetBoxByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetRawBoxByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type ResetBoxRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type ResetBoxByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DescribeLotteryModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetLotteryModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
}

type DescribePrizeTablesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetPrizeTableRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PrizeTableName *core.String	`json:"prizeTableName"`
}

type DrawByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
	UserId *core.String	`json:"userId"`
	Count *int32	`json:"count"`
	Config *[]*Config	`json:"config"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DescribeProbabilitiesRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeProbabilitiesByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	LotteryName *core.String	`json:"lotteryName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DrawByStampSheetRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	StampSheet *core.String	`json:"stampSheet"`
	KeyId *core.String	`json:"keyId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetCurrentLotteryMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateCurrentLotteryMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Settings *core.String	`json:"settings"`
}

type UpdateCurrentLotteryMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}
