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
	EnableRating *bool	`json:"enableRating"`
	CreateGatheringTriggerType *core.String	`json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId *core.String	`json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId *core.String	`json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType *core.String	`json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *core.String	`json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId *core.String	`json:"completeMatchmakingTriggerScriptId"`
	JoinNotification *NotificationSetting	`json:"joinNotification"`
	LeaveNotification *NotificationSetting	`json:"leaveNotification"`
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
	EnableRating *bool	`json:"enableRating"`
	CreateGatheringTriggerType *core.String	`json:"createGatheringTriggerType"`
	CreateGatheringTriggerRealtimeNamespaceId *core.String	`json:"createGatheringTriggerRealtimeNamespaceId"`
	CreateGatheringTriggerScriptId *core.String	`json:"createGatheringTriggerScriptId"`
	CompleteMatchmakingTriggerType *core.String	`json:"completeMatchmakingTriggerType"`
	CompleteMatchmakingTriggerRealtimeNamespaceId *core.String	`json:"completeMatchmakingTriggerRealtimeNamespaceId"`
	CompleteMatchmakingTriggerScriptId *core.String	`json:"completeMatchmakingTriggerScriptId"`
	JoinNotification *NotificationSetting	`json:"joinNotification"`
	LeaveNotification *NotificationSetting	`json:"leaveNotification"`
	CompleteNotification *NotificationSetting	`json:"completeNotification"`
	LogSetting *LogSetting	`json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type DescribeGatheringsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateGatheringRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Player *Player	`json:"player"`
	AttributeRanges *[]*AttributeRange	`json:"attributeRanges"`
	CapacityOfRoles *[]*CapacityOfRole	`json:"capacityOfRoles"`
	AllowUserIds *[]core.String	`json:"allowUserIds"`
	ExpiresAt *int64	`json:"expiresAt"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CreateGatheringByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	Player *Player	`json:"player"`
	AttributeRanges *[]*AttributeRange	`json:"attributeRanges"`
	CapacityOfRoles *[]*CapacityOfRole	`json:"capacityOfRoles"`
	AllowUserIds *[]core.String	`json:"allowUserIds"`
	ExpiresAt *int64	`json:"expiresAt"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type UpdateGatheringRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
	AttributeRanges *[]*AttributeRange	`json:"attributeRanges"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type UpdateGatheringByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
	UserId *core.String	`json:"userId"`
	AttributeRanges *[]*AttributeRange	`json:"attributeRanges"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DoMatchmakingByPlayerRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Player *Player	`json:"player"`
	MatchmakingContextToken *core.String	`json:"matchmakingContextToken"`
}

type DoMatchmakingRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Player *Player	`json:"player"`
	MatchmakingContextToken *core.String	`json:"matchmakingContextToken"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetGatheringRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
}

type CancelMatchmakingRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type CancelMatchmakingByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
	UserId *core.String	`json:"userId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type DeleteGatheringRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	GatheringName *core.String	`json:"gatheringName"`
}

type DescribeRatingModelMastersRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
}

type CreateRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Name *core.String	`json:"name"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Volatility *int32	`json:"volatility"`
}

type GetRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
}

type UpdateRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	Description *core.String	`json:"description"`
	Metadata *core.String	`json:"metadata"`
	Volatility *int32	`json:"volatility"`
}

type DeleteRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
}

type DescribeRatingModelsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetRatingModelRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
}

type ExportMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type GetCurrentRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
}

type UpdateCurrentRatingModelMasterRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	Settings *core.String	`json:"settings"`
}

type UpdateCurrentRatingModelMasterFromGitHubRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	CheckoutSetting *GitHubCheckoutSetting	`json:"checkoutSetting"`
}

type DescribeRatingsRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type DescribeRatingsByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	PageToken *core.String	`json:"pageToken"`
	Limit *int64	`json:"limit"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetRatingRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetRatingByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	RatingName *core.String	`json:"ratingName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type PutResultRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	GameResults *[]*GameResult	`json:"gameResults"`
}

type DeleteRatingRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	UserId *core.String	`json:"userId"`
	RatingName *core.String	`json:"ratingName"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type GetBallotRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	GatheringName *core.String	`json:"gatheringName"`
	NumberOfPlayer *int32	`json:"numberOfPlayer"`
	KeyId *core.String	`json:"keyId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
	AccessToken *core.AccessToken	`json:"accessToken"`
}

type GetBallotByUserIdRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	GatheringName *core.String	`json:"gatheringName"`
	UserId *core.String	`json:"userId"`
	NumberOfPlayer *int32	`json:"numberOfPlayer"`
	KeyId *core.String	`json:"keyId"`
	XGs2DuplicationAvoider *core.String	`json:"xGs2DuplicationAvoider"`
}

type VoteRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	BallotBody *core.String	`json:"ballotBody"`
	BallotSignature *core.String	`json:"ballotSignature"`
	GameResults *[]*GameResult	`json:"gameResults"`
	KeyId *core.String	`json:"keyId"`
}

type VoteMultipleRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	SignedBallots *[]*SignedBallot	`json:"signedBallots"`
	GameResults *[]*GameResult	`json:"gameResults"`
	KeyId *core.String	`json:"keyId"`
}

type CommitVoteRequest struct {
	RequestId    *core.RequestId	`json:"requestId"`
	ContextStack *core.ContextStack	`json:"contextStack"`
	NamespaceName *core.String	`json:"namespaceName"`
	RatingName *core.String	`json:"ratingName"`
	GatheringName *core.String	`json:"gatheringName"`
}
