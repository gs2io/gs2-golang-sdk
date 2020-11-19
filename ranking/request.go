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

package ranking

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
	RequestId    *core.RequestId    `json:"requestId"`
	ContextStack *core.ContextStack `json:"contextStack"`
	Name         *string            `json:"name"`
	Description  *string            `json:"description"`
	LogSetting   *LogSetting        `json:"logSetting"`
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
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Description   *string            `json:"description"`
	LogSetting    *LogSetting        `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeCategoryModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCategoryModelRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CategoryName  *string            `json:"categoryName"`
}

type DescribeCategoryModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateCategoryModelMasterRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	Name                     *string            `json:"name"`
	Description              *string            `json:"description"`
	Metadata                 *string            `json:"metadata"`
	MinimumValue             *int64             `json:"minimumValue"`
	MaximumValue             *int64             `json:"maximumValue"`
	OrderDirection           *string            `json:"orderDirection"`
	Scope                    *string            `json:"scope"`
	UniqueByUserId           *bool              `json:"uniqueByUserId"`
	CalculateIntervalMinutes *int32             `json:"calculateIntervalMinutes"`
	EntryPeriodEventId       *string            `json:"entryPeriodEventId"`
	AccessPeriodEventId      *string            `json:"accessPeriodEventId"`
}

type GetCategoryModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CategoryName  *string            `json:"categoryName"`
}

type UpdateCategoryModelMasterRequest struct {
	RequestId                *core.RequestId    `json:"requestId"`
	ContextStack             *core.ContextStack `json:"contextStack"`
	NamespaceName            *string            `json:"namespaceName"`
	CategoryName             *string            `json:"categoryName"`
	Description              *string            `json:"description"`
	Metadata                 *string            `json:"metadata"`
	MinimumValue             *int64             `json:"minimumValue"`
	MaximumValue             *int64             `json:"maximumValue"`
	OrderDirection           *string            `json:"orderDirection"`
	Scope                    *string            `json:"scope"`
	UniqueByUserId           *bool              `json:"uniqueByUserId"`
	CalculateIntervalMinutes *int32             `json:"calculateIntervalMinutes"`
	EntryPeriodEventId       *string            `json:"entryPeriodEventId"`
	AccessPeriodEventId      *string            `json:"accessPeriodEventId"`
}

type DeleteCategoryModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CategoryName  *string            `json:"categoryName"`
}

type DescribeSubscribesByCategoryNameRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeSubscribesByCategoryNameAndUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SubscribeRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type SubscribeByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetSubscribeRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetSubscribeByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UnsubscribeRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UnsubscribeByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeScoresRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	ScorerUserId           *string            `json:"scorerUserId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeScoresByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	ScorerUserId           *string            `json:"scorerUserId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetScoreRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	ScorerUserId           *string            `json:"scorerUserId"`
	UniqueId               *string            `json:"uniqueId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetScoreByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	ScorerUserId           *string            `json:"scorerUserId"`
	UniqueId               *string            `json:"uniqueId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeRankingsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	StartIndex             *int64             `json:"startIndex"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeRankingssByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	StartIndex             *int64             `json:"startIndex"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeNearRankingsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	CategoryName  *string            `json:"categoryName"`
	Score         *int64             `json:"score"`
}

type GetRankingRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	ScorerUserId           *string            `json:"scorerUserId"`
	UniqueId               *string            `json:"uniqueId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetRankingByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	ScorerUserId           *string            `json:"scorerUserId"`
	UniqueId               *string            `json:"uniqueId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PutScoreRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	Score                  *int64             `json:"score"`
	Metadata               *string            `json:"metadata"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type PutScoreByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	CategoryName           *string            `json:"categoryName"`
	UserId                 *string            `json:"userId"`
	Score                  *int64             `json:"score"`
	Metadata               *string            `json:"metadata"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentRankingMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentRankingMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentRankingMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}
