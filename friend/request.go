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

package friend

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
	RequestId                  *core.RequestId      `json:"requestId"`
	ContextStack               *core.ContextStack   `json:"contextStack"`
	Name                       *string              `json:"name"`
	Description                *string              `json:"description"`
	FollowScript               *ScriptSetting       `json:"followScript"`
	UnfollowScript             *ScriptSetting       `json:"unfollowScript"`
	SendRequestScript          *ScriptSetting       `json:"sendRequestScript"`
	CancelRequestScript        *ScriptSetting       `json:"cancelRequestScript"`
	AcceptRequestScript        *ScriptSetting       `json:"acceptRequestScript"`
	RejectRequestScript        *ScriptSetting       `json:"rejectRequestScript"`
	DeleteFriendScript         *ScriptSetting       `json:"deleteFriendScript"`
	UpdateProfileScript        *ScriptSetting       `json:"updateProfileScript"`
	FollowNotification         *NotificationSetting `json:"followNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	AcceptRequestNotification  *NotificationSetting `json:"acceptRequestNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
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
	RequestId                  *core.RequestId      `json:"requestId"`
	ContextStack               *core.ContextStack   `json:"contextStack"`
	NamespaceName              *string              `json:"namespaceName"`
	Description                *string              `json:"description"`
	FollowScript               *ScriptSetting       `json:"followScript"`
	UnfollowScript             *ScriptSetting       `json:"unfollowScript"`
	SendRequestScript          *ScriptSetting       `json:"sendRequestScript"`
	CancelRequestScript        *ScriptSetting       `json:"cancelRequestScript"`
	AcceptRequestScript        *ScriptSetting       `json:"acceptRequestScript"`
	RejectRequestScript        *ScriptSetting       `json:"rejectRequestScript"`
	DeleteFriendScript         *ScriptSetting       `json:"deleteFriendScript"`
	UpdateProfileScript        *ScriptSetting       `json:"updateProfileScript"`
	FollowNotification         *NotificationSetting `json:"followNotification"`
	ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
	AcceptRequestNotification  *NotificationSetting `json:"acceptRequestNotification"`
	LogSetting                 *LogSetting          `json:"logSetting"`
}

type DeleteNamespaceRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
}

type DescribeProfilesRequest struct {
	RequestId     *core.RequestId    `json:"requestId"`
	ContextStack  *core.ContextStack `json:"contextStack"`
	NamespaceName *string            `json:"namespaceName"`
	PageToken     *string            `json:"pageToken"`
	Limit         *int64             `json:"limit"`
}

type GetProfileRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetProfileByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UpdateProfileRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	PublicProfile          *string            `json:"publicProfile"`
	FollowerProfile        *string            `json:"followerProfile"`
	FriendProfile          *string            `json:"friendProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UpdateProfileByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	PublicProfile          *string            `json:"publicProfile"`
	FollowerProfile        *string            `json:"followerProfile"`
	FriendProfile          *string            `json:"friendProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteProfileByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetPublicProfileRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeFollowsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	WithProfile            *bool              `json:"withProfile"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeFollowsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	WithProfile            *bool              `json:"withProfile"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetFollowRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	WithProfile            *bool              `json:"withProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetFollowByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	WithProfile            *bool              `json:"withProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type FollowRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type FollowByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UnfollowRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UnfollowByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeFriendsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	WithProfile            *bool              `json:"withProfile"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeFriendsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	WithProfile            *bool              `json:"withProfile"`
	PageToken              *string            `json:"pageToken"`
	Limit                  *int64             `json:"limit"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetFriendRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	WithProfile            *bool              `json:"withProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetFriendByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	WithProfile            *bool              `json:"withProfile"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteFriendRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteFriendByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeSendRequestsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeSendRequestsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetSendRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetSendRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type SendRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type SendRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DeleteRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DeleteRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeReceiveRequestsRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeReceiveRequestsByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type GetReceiveRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type GetReceiveRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type AcceptRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type AcceptRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type RejectRequestRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type RejectRequestByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	FromUserId             *string            `json:"fromUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type DescribeBlackListRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type DescribeBlackListByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type RegisterBlackListRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type RegisterBlackListByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}

type UnregisterBlackListRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
	AccessToken            *core.AccessToken  `json:"accessToken"`
}

type UnregisterBlackListByUserIdRequest struct {
	RequestId              *core.RequestId    `json:"requestId"`
	ContextStack           *core.ContextStack `json:"contextStack"`
	NamespaceName          *string            `json:"namespaceName"`
	UserId                 *string            `json:"userId"`
	TargetUserId           *string            `json:"targetUserId"`
	XGs2DuplicationAvoider *string            `json:"xGs2DuplicationAvoider"`
}
