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

import "core"

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
    return DescribeNamespacesResult {
        Items: CastNamespaces(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastNamespacesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
    return &p
}

type CreateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
    return CreateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
    return &p
}

type GetNamespaceStatusResult struct {
    Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
    return GetNamespaceStatusResult {
        Status: core.CastString(data["status"]),
    }
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "status": p.Status,
    }
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
    return &p
}

type GetNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
    return GetNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
    return &p
}

type UpdateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
    return UpdateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
    return &p
}

type DeleteNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeProfilesResult struct {
    Items []Profile `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeProfilesAsyncResult struct {
	result *DescribeProfilesResult
	err    error
}

func NewDescribeProfilesResultFromDict(data map[string]interface{}) DescribeProfilesResult {
    return DescribeProfilesResult {
        Items: CastProfiles(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeProfilesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastProfilesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeProfilesResult) Pointer() *DescribeProfilesResult {
    return &p
}

type GetProfileResult struct {
    Item *Profile `json:"item"`
}

type GetProfileAsyncResult struct {
	result *GetProfileResult
	err    error
}

func NewGetProfileResultFromDict(data map[string]interface{}) GetProfileResult {
    return GetProfileResult {
        Item: NewProfileFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetProfileResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetProfileResult) Pointer() *GetProfileResult {
    return &p
}

type GetProfileByUserIdResult struct {
    Item *Profile `json:"item"`
}

type GetProfileByUserIdAsyncResult struct {
	result *GetProfileByUserIdResult
	err    error
}

func NewGetProfileByUserIdResultFromDict(data map[string]interface{}) GetProfileByUserIdResult {
    return GetProfileByUserIdResult {
        Item: NewProfileFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetProfileByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetProfileByUserIdResult) Pointer() *GetProfileByUserIdResult {
    return &p
}

type UpdateProfileResult struct {
    Item *Profile `json:"item"`
}

type UpdateProfileAsyncResult struct {
	result *UpdateProfileResult
	err    error
}

func NewUpdateProfileResultFromDict(data map[string]interface{}) UpdateProfileResult {
    return UpdateProfileResult {
        Item: NewProfileFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateProfileResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateProfileResult) Pointer() *UpdateProfileResult {
    return &p
}

type UpdateProfileByUserIdResult struct {
    Item *Profile `json:"item"`
}

type UpdateProfileByUserIdAsyncResult struct {
	result *UpdateProfileByUserIdResult
	err    error
}

func NewUpdateProfileByUserIdResultFromDict(data map[string]interface{}) UpdateProfileByUserIdResult {
    return UpdateProfileByUserIdResult {
        Item: NewProfileFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateProfileByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateProfileByUserIdResult) Pointer() *UpdateProfileByUserIdResult {
    return &p
}

type DeleteProfileByUserIdResult struct {
}

type DeleteProfileByUserIdAsyncResult struct {
	result *DeleteProfileByUserIdResult
	err    error
}

func NewDeleteProfileByUserIdResultFromDict(data map[string]interface{}) DeleteProfileByUserIdResult {
    return DeleteProfileByUserIdResult {
    }
}

func (p DeleteProfileByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
    }
}

func (p DeleteProfileByUserIdResult) Pointer() *DeleteProfileByUserIdResult {
    return &p
}

type GetPublicProfileResult struct {
    Item *PublicProfile `json:"item"`
}

type GetPublicProfileAsyncResult struct {
	result *GetPublicProfileResult
	err    error
}

func NewGetPublicProfileResultFromDict(data map[string]interface{}) GetPublicProfileResult {
    return GetPublicProfileResult {
        Item: NewPublicProfileFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetPublicProfileResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetPublicProfileResult) Pointer() *GetPublicProfileResult {
    return &p
}

type DescribeFollowsResult struct {
    Items []FollowUser `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeFollowsAsyncResult struct {
	result *DescribeFollowsResult
	err    error
}

func NewDescribeFollowsResultFromDict(data map[string]interface{}) DescribeFollowsResult {
    return DescribeFollowsResult {
        Items: CastFollowUsers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeFollowsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFollowUsersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeFollowsResult) Pointer() *DescribeFollowsResult {
    return &p
}

type DescribeFollowsByUserIdResult struct {
    Items []FollowUser `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeFollowsByUserIdAsyncResult struct {
	result *DescribeFollowsByUserIdResult
	err    error
}

func NewDescribeFollowsByUserIdResultFromDict(data map[string]interface{}) DescribeFollowsByUserIdResult {
    return DescribeFollowsByUserIdResult {
        Items: CastFollowUsers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeFollowsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFollowUsersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeFollowsByUserIdResult) Pointer() *DescribeFollowsByUserIdResult {
    return &p
}

type GetFollowResult struct {
    Item *FollowUser `json:"item"`
}

type GetFollowAsyncResult struct {
	result *GetFollowResult
	err    error
}

func NewGetFollowResultFromDict(data map[string]interface{}) GetFollowResult {
    return GetFollowResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFollowResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFollowResult) Pointer() *GetFollowResult {
    return &p
}

type GetFollowByUserIdResult struct {
    Item *FollowUser `json:"item"`
}

type GetFollowByUserIdAsyncResult struct {
	result *GetFollowByUserIdResult
	err    error
}

func NewGetFollowByUserIdResultFromDict(data map[string]interface{}) GetFollowByUserIdResult {
    return GetFollowByUserIdResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFollowByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFollowByUserIdResult) Pointer() *GetFollowByUserIdResult {
    return &p
}

type FollowResult struct {
    Item *FollowUser `json:"item"`
}

type FollowAsyncResult struct {
	result *FollowResult
	err    error
}

func NewFollowResultFromDict(data map[string]interface{}) FollowResult {
    return FollowResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p FollowResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p FollowResult) Pointer() *FollowResult {
    return &p
}

type FollowByUserIdResult struct {
    Item *FollowUser `json:"item"`
}

type FollowByUserIdAsyncResult struct {
	result *FollowByUserIdResult
	err    error
}

func NewFollowByUserIdResultFromDict(data map[string]interface{}) FollowByUserIdResult {
    return FollowByUserIdResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p FollowByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p FollowByUserIdResult) Pointer() *FollowByUserIdResult {
    return &p
}

type UnfollowResult struct {
    Item *FollowUser `json:"item"`
}

type UnfollowAsyncResult struct {
	result *UnfollowResult
	err    error
}

func NewUnfollowResultFromDict(data map[string]interface{}) UnfollowResult {
    return UnfollowResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnfollowResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnfollowResult) Pointer() *UnfollowResult {
    return &p
}

type UnfollowByUserIdResult struct {
    Item *FollowUser `json:"item"`
}

type UnfollowByUserIdAsyncResult struct {
	result *UnfollowByUserIdResult
	err    error
}

func NewUnfollowByUserIdResultFromDict(data map[string]interface{}) UnfollowByUserIdResult {
    return UnfollowByUserIdResult {
        Item: NewFollowUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnfollowByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnfollowByUserIdResult) Pointer() *UnfollowByUserIdResult {
    return &p
}

type DescribeFriendsResult struct {
    Items []FriendUser `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeFriendsAsyncResult struct {
	result *DescribeFriendsResult
	err    error
}

func NewDescribeFriendsResultFromDict(data map[string]interface{}) DescribeFriendsResult {
    return DescribeFriendsResult {
        Items: CastFriendUsers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeFriendsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendUsersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeFriendsResult) Pointer() *DescribeFriendsResult {
    return &p
}

type DescribeFriendsByUserIdResult struct {
    Items []FriendUser `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeFriendsByUserIdAsyncResult struct {
	result *DescribeFriendsByUserIdResult
	err    error
}

func NewDescribeFriendsByUserIdResultFromDict(data map[string]interface{}) DescribeFriendsByUserIdResult {
    return DescribeFriendsByUserIdResult {
        Items: CastFriendUsers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeFriendsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendUsersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeFriendsByUserIdResult) Pointer() *DescribeFriendsByUserIdResult {
    return &p
}

type GetFriendResult struct {
    Item *FriendUser `json:"item"`
}

type GetFriendAsyncResult struct {
	result *GetFriendResult
	err    error
}

func NewGetFriendResultFromDict(data map[string]interface{}) GetFriendResult {
    return GetFriendResult {
        Item: NewFriendUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFriendResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFriendResult) Pointer() *GetFriendResult {
    return &p
}

type GetFriendByUserIdResult struct {
    Item *FriendUser `json:"item"`
}

type GetFriendByUserIdAsyncResult struct {
	result *GetFriendByUserIdResult
	err    error
}

func NewGetFriendByUserIdResultFromDict(data map[string]interface{}) GetFriendByUserIdResult {
    return GetFriendByUserIdResult {
        Item: NewFriendUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetFriendByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetFriendByUserIdResult) Pointer() *GetFriendByUserIdResult {
    return &p
}

type DeleteFriendResult struct {
    Item *FriendUser `json:"item"`
}

type DeleteFriendAsyncResult struct {
	result *DeleteFriendResult
	err    error
}

func NewDeleteFriendResultFromDict(data map[string]interface{}) DeleteFriendResult {
    return DeleteFriendResult {
        Item: NewFriendUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteFriendResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteFriendResult) Pointer() *DeleteFriendResult {
    return &p
}

type DeleteFriendByUserIdResult struct {
    Item *FriendUser `json:"item"`
}

type DeleteFriendByUserIdAsyncResult struct {
	result *DeleteFriendByUserIdResult
	err    error
}

func NewDeleteFriendByUserIdResultFromDict(data map[string]interface{}) DeleteFriendByUserIdResult {
    return DeleteFriendByUserIdResult {
        Item: NewFriendUserFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteFriendByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteFriendByUserIdResult) Pointer() *DeleteFriendByUserIdResult {
    return &p
}

type DescribeSendRequestsResult struct {
    Items []FriendRequest `json:"items"`
}

type DescribeSendRequestsAsyncResult struct {
	result *DescribeSendRequestsResult
	err    error
}

func NewDescribeSendRequestsResultFromDict(data map[string]interface{}) DescribeSendRequestsResult {
    return DescribeSendRequestsResult {
        Items: CastFriendRequests(core.CastArray(data["items"])),
    }
}

func (p DescribeSendRequestsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendRequestsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeSendRequestsResult) Pointer() *DescribeSendRequestsResult {
    return &p
}

type DescribeSendRequestsByUserIdResult struct {
    Items []FriendRequest `json:"items"`
}

type DescribeSendRequestsByUserIdAsyncResult struct {
	result *DescribeSendRequestsByUserIdResult
	err    error
}

func NewDescribeSendRequestsByUserIdResultFromDict(data map[string]interface{}) DescribeSendRequestsByUserIdResult {
    return DescribeSendRequestsByUserIdResult {
        Items: CastFriendRequests(core.CastArray(data["items"])),
    }
}

func (p DescribeSendRequestsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendRequestsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeSendRequestsByUserIdResult) Pointer() *DescribeSendRequestsByUserIdResult {
    return &p
}

type GetSendRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type GetSendRequestAsyncResult struct {
	result *GetSendRequestResult
	err    error
}

func NewGetSendRequestResultFromDict(data map[string]interface{}) GetSendRequestResult {
    return GetSendRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSendRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSendRequestResult) Pointer() *GetSendRequestResult {
    return &p
}

type GetSendRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type GetSendRequestByUserIdAsyncResult struct {
	result *GetSendRequestByUserIdResult
	err    error
}

func NewGetSendRequestByUserIdResultFromDict(data map[string]interface{}) GetSendRequestByUserIdResult {
    return GetSendRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSendRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSendRequestByUserIdResult) Pointer() *GetSendRequestByUserIdResult {
    return &p
}

type SendRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type SendRequestAsyncResult struct {
	result *SendRequestResult
	err    error
}

func NewSendRequestResultFromDict(data map[string]interface{}) SendRequestResult {
    return SendRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SendRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SendRequestResult) Pointer() *SendRequestResult {
    return &p
}

type SendRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type SendRequestByUserIdAsyncResult struct {
	result *SendRequestByUserIdResult
	err    error
}

func NewSendRequestByUserIdResultFromDict(data map[string]interface{}) SendRequestByUserIdResult {
    return SendRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p SendRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p SendRequestByUserIdResult) Pointer() *SendRequestByUserIdResult {
    return &p
}

type DeleteRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type DeleteRequestAsyncResult struct {
	result *DeleteRequestResult
	err    error
}

func NewDeleteRequestResultFromDict(data map[string]interface{}) DeleteRequestResult {
    return DeleteRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRequestResult) Pointer() *DeleteRequestResult {
    return &p
}

type DeleteRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type DeleteRequestByUserIdAsyncResult struct {
	result *DeleteRequestByUserIdResult
	err    error
}

func NewDeleteRequestByUserIdResultFromDict(data map[string]interface{}) DeleteRequestByUserIdResult {
    return DeleteRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRequestByUserIdResult) Pointer() *DeleteRequestByUserIdResult {
    return &p
}

type DescribeReceiveRequestsResult struct {
    Items []FriendRequest `json:"items"`
}

type DescribeReceiveRequestsAsyncResult struct {
	result *DescribeReceiveRequestsResult
	err    error
}

func NewDescribeReceiveRequestsResultFromDict(data map[string]interface{}) DescribeReceiveRequestsResult {
    return DescribeReceiveRequestsResult {
        Items: CastFriendRequests(core.CastArray(data["items"])),
    }
}

func (p DescribeReceiveRequestsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendRequestsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeReceiveRequestsResult) Pointer() *DescribeReceiveRequestsResult {
    return &p
}

type DescribeReceiveRequestsByUserIdResult struct {
    Items []FriendRequest `json:"items"`
}

type DescribeReceiveRequestsByUserIdAsyncResult struct {
	result *DescribeReceiveRequestsByUserIdResult
	err    error
}

func NewDescribeReceiveRequestsByUserIdResultFromDict(data map[string]interface{}) DescribeReceiveRequestsByUserIdResult {
    return DescribeReceiveRequestsByUserIdResult {
        Items: CastFriendRequests(core.CastArray(data["items"])),
    }
}

func (p DescribeReceiveRequestsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastFriendRequestsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeReceiveRequestsByUserIdResult) Pointer() *DescribeReceiveRequestsByUserIdResult {
    return &p
}

type GetReceiveRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type GetReceiveRequestAsyncResult struct {
	result *GetReceiveRequestResult
	err    error
}

func NewGetReceiveRequestResultFromDict(data map[string]interface{}) GetReceiveRequestResult {
    return GetReceiveRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetReceiveRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetReceiveRequestResult) Pointer() *GetReceiveRequestResult {
    return &p
}

type GetReceiveRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type GetReceiveRequestByUserIdAsyncResult struct {
	result *GetReceiveRequestByUserIdResult
	err    error
}

func NewGetReceiveRequestByUserIdResultFromDict(data map[string]interface{}) GetReceiveRequestByUserIdResult {
    return GetReceiveRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetReceiveRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetReceiveRequestByUserIdResult) Pointer() *GetReceiveRequestByUserIdResult {
    return &p
}

type AcceptRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type AcceptRequestAsyncResult struct {
	result *AcceptRequestResult
	err    error
}

func NewAcceptRequestResultFromDict(data map[string]interface{}) AcceptRequestResult {
    return AcceptRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcceptRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcceptRequestResult) Pointer() *AcceptRequestResult {
    return &p
}

type AcceptRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type AcceptRequestByUserIdAsyncResult struct {
	result *AcceptRequestByUserIdResult
	err    error
}

func NewAcceptRequestByUserIdResultFromDict(data map[string]interface{}) AcceptRequestByUserIdResult {
    return AcceptRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p AcceptRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p AcceptRequestByUserIdResult) Pointer() *AcceptRequestByUserIdResult {
    return &p
}

type RejectRequestResult struct {
    Item *FriendRequest `json:"item"`
}

type RejectRequestAsyncResult struct {
	result *RejectRequestResult
	err    error
}

func NewRejectRequestResultFromDict(data map[string]interface{}) RejectRequestResult {
    return RejectRequestResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p RejectRequestResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p RejectRequestResult) Pointer() *RejectRequestResult {
    return &p
}

type RejectRequestByUserIdResult struct {
    Item *FriendRequest `json:"item"`
}

type RejectRequestByUserIdAsyncResult struct {
	result *RejectRequestByUserIdResult
	err    error
}

func NewRejectRequestByUserIdResultFromDict(data map[string]interface{}) RejectRequestByUserIdResult {
    return RejectRequestByUserIdResult {
        Item: NewFriendRequestFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p RejectRequestByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p RejectRequestByUserIdResult) Pointer() *RejectRequestByUserIdResult {
    return &p
}

type DescribeBlackListResult struct {
    Items []string `json:"items"`
}

type DescribeBlackListAsyncResult struct {
	result *DescribeBlackListResult
	err    error
}

func NewDescribeBlackListResultFromDict(data map[string]interface{}) DescribeBlackListResult {
    return DescribeBlackListResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
    }
}

func (p DescribeBlackListResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeBlackListResult) Pointer() *DescribeBlackListResult {
    return &p
}

type DescribeBlackListByUserIdResult struct {
    Items []string `json:"items"`
}

type DescribeBlackListByUserIdAsyncResult struct {
	result *DescribeBlackListByUserIdResult
	err    error
}

func NewDescribeBlackListByUserIdResultFromDict(data map[string]interface{}) DescribeBlackListByUserIdResult {
    return DescribeBlackListByUserIdResult {
        Items: core.CastStrings(core.CastArray(data["items"])),
    }
}

func (p DescribeBlackListByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": core.CastStringsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeBlackListByUserIdResult) Pointer() *DescribeBlackListByUserIdResult {
    return &p
}

type RegisterBlackListResult struct {
    Item *BlackList `json:"item"`
}

type RegisterBlackListAsyncResult struct {
	result *RegisterBlackListResult
	err    error
}

func NewRegisterBlackListResultFromDict(data map[string]interface{}) RegisterBlackListResult {
    return RegisterBlackListResult {
        Item: NewBlackListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p RegisterBlackListResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p RegisterBlackListResult) Pointer() *RegisterBlackListResult {
    return &p
}

type RegisterBlackListByUserIdResult struct {
    Item *BlackList `json:"item"`
}

type RegisterBlackListByUserIdAsyncResult struct {
	result *RegisterBlackListByUserIdResult
	err    error
}

func NewRegisterBlackListByUserIdResultFromDict(data map[string]interface{}) RegisterBlackListByUserIdResult {
    return RegisterBlackListByUserIdResult {
        Item: NewBlackListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p RegisterBlackListByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p RegisterBlackListByUserIdResult) Pointer() *RegisterBlackListByUserIdResult {
    return &p
}

type UnregisterBlackListResult struct {
    Item *BlackList `json:"item"`
}

type UnregisterBlackListAsyncResult struct {
	result *UnregisterBlackListResult
	err    error
}

func NewUnregisterBlackListResultFromDict(data map[string]interface{}) UnregisterBlackListResult {
    return UnregisterBlackListResult {
        Item: NewBlackListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnregisterBlackListResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnregisterBlackListResult) Pointer() *UnregisterBlackListResult {
    return &p
}

type UnregisterBlackListByUserIdResult struct {
    Item *BlackList `json:"item"`
}

type UnregisterBlackListByUserIdAsyncResult struct {
	result *UnregisterBlackListByUserIdResult
	err    error
}

func NewUnregisterBlackListByUserIdResultFromDict(data map[string]interface{}) UnregisterBlackListByUserIdResult {
    return UnregisterBlackListByUserIdResult {
        Item: NewBlackListFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UnregisterBlackListByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UnregisterBlackListByUserIdResult) Pointer() *UnregisterBlackListByUserIdResult {
    return &p
}