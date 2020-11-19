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

package matchmaking

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
	RequestId                                     *core.RequestId      `json:"requestId"`
	ContextStack                                  *core.ContextStack   `json:"contextStack"`
	Name                                          *string              `json:"name"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
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
	RequestId                                     *core.RequestId      `json:"requestId"`
	ContextStack                                  *core.ContextStack   `json:"contextStack"`
	NamespaceName                                 *string              `json:"namespaceName"`
	Description                                   *string              `json:"description"`
	EnableRating                                  *bool                `json:"enableRating"`
	CreateGatheringTriggerType                    *string              `json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId     *string              `json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId                *string              `json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType                *string              `json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *string              `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId            *string              `json:"completeMatchmakingTriggerScriptId"`
	JoinNotification                              *NotificationSetting `json:"joinNotification"`
	LeaveNotification                             *NotificationSetting `json:"leaveNotification"`
	CompleteNotification                          *NotificationSetting `json:"completeNotification"`
	LogSetting                                    *LogSetting          `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeGatheringsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateGatheringRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	Player                 *Player            `json:"player"`
	AttributeRanges        *[]*AttributeRange `json:"attributeRanges"`
	CapacityOfRoles        *[]*CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	ExpiresAt              *int64             `json:"expiresAt"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type CreateGatheringByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	Player                 *Player            `json:"player"`
	AttributeRanges        *[]*AttributeRange `json:"attributeRanges"`
	CapacityOfRoles        *[]*CapacityOfRole `json:"capacityOfRoles"`
	AllowUserIds           *[]string          `json:"allowUserIds"`
	ExpiresAt              *int64             `json:"expiresAt"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UpdateGatheringRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	GatheringName          *string            `json:"gatheringName"`
	AttributeRanges        *[]*AttributeRange `json:"attributeRanges"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UpdateGatheringByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	GatheringName          *string            `json:"gatheringName"`
	UserId                 *string            `json:"userId"`
	AttributeRanges        *[]*AttributeRange `json:"attributeRanges"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DoMatchmakingByPlayerRequest struct {
	RequestId               *core.RequestId    `json:"requestId"`
	ContextStack            *core.ContextStack `json:"contextStack"`
	NamespaceName           *string            `json:"namespaceName"`
	Player                  *Player            `json:"player"`
	MatchmakingContextToken *string            `json:"matchmakingContextToken"`
}

type DoMatchmakingRequest struct {
	RequestId               *core.RequestId    `json:"requestId"`
	ContextStack            *core.ContextStack `json:"contextStack"`
	NamespaceName           *string            `json:"namespaceName"`
	Player                  *Player            `json:"player"`
	MatchmakingContextToken *string            `json:"matchmakingContextToken"`
	XGs2DuplicationAvoider  *string            `json:"xGs2DuplicationAvoider"`
	AccessToken             *core.AccessToken  `json:"accessToken"`
}

type GetGatheringRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	GatheringName *string            `json:"gatheringName"`
}

type CancelMatchmakingRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	GatheringName          *string            `json:"gatheringName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type CancelMatchmakingByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	GatheringName          *string            `json:"gatheringName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteGatheringRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	GatheringName *string            `json:"gatheringName"`
}

type DescribeRatingModelMastersRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type CreateRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Name          *string            `json:"name"`
	Description   *string            `json:"description"`
	Metadata      *string            `json:"metadata"`
	Volatility    *int32             `json:"volatility"`
}

type GetRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
}

type UpdateRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
	Description   *string            `json:"description"`
	Metadata      *string            `json:"metadata"`
	Volatility    *int32             `json:"volatility"`
}

type DeleteRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
}

type DescribeRatingModelsRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetRatingModelRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
}

type ExportMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type GetCurrentRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type UpdateCurrentRatingModelMasterRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	Settings      *string            `json:"settings"`
}

type UpdateCurrentRatingModelMasterFromGitHubRequest struct {
	RequestId       *core.RequestId        `json:"requestId"`
	ContextStack    *core.ContextStack     `json:"contextStack"`
	NamespaceName   *string                `json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

type DescribeRatingsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeRatingsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetRatingRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RatingName             *string            `json:"ratingName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetRatingByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RatingName             *string            `json:"ratingName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type PutResultRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
	GameResults   *[]*GameResult     `json:"gameResults"`
}

type DeleteRatingRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	RatingName             *string            `json:"ratingName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetBallotRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RatingName             *string            `json:"ratingName"`
	GatheringName          *string            `json:"gatheringName"`
	GatheringId            *string            `json:"gatheringId"`
	NumberOfPlayer         *int32             `json:"numberOfPlayer"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetBallotByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	RatingName             *string            `json:"ratingName"`
	GatheringName          *string            `json:"gatheringName"`
	UserId                 *string            `json:"userId"`
	GatheringId            *string            `json:"gatheringId"`
	NumberOfPlayer         *int32             `json:"numberOfPlayer"`
	KeyId                  *string            `json:"keyId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type VoteRequest struct {
	RequestId       *core.RequestId    `json:"requestId"`
	ContextStack    *core.ContextStack `json:"contextStack"`
	NamespaceName   *string            `json:"namespaceName"`
	BallotBody      *string            `json:"ballotBody"`
	BallotSignature *string            `json:"ballotSignature"`
	GameResults     *[]*GameResult     `json:"gameResults"`
	KeyId           *string            `json:"keyId"`
}

type VoteMultipleRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	SignedBallots *[]*SignedBallot   `json:"signedBallots"`
	GameResults   *[]*GameResult     `json:"gameResults"`
	KeyId         *string            `json:"keyId"`
}

type CommitVoteRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	RatingName    *string            `json:"ratingName"`
	GatheringName *string            `json:"gatheringName"`
}
