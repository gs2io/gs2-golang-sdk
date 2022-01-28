package domainfriend
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

import (
    "github.com/gs2io/gs2-golang-sdk/auth"
    "github.com/gs2io/gs2-golang-sdk/core"

    "github.com/gs2io/gs2-golang-sdk/friend"
    "github.com/gs2io/gs2-golang-sdk/friend/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/friend/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type FollowDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    followUserCache cache.FollowUserDomainCache
    namespaceName string
    userId string
}

func NewFollowDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) FollowDomain {
    return FollowDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        followUserCache: cache.NewFollowUserDomainCache(),
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p FollowDomain) Load(
    request friend.GetFollowByUserIdRequest,
) (*friend.GetFollowByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetFollowByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowDomain) Follow(
    request friend.FollowByUserIdRequest,
) (*friend.FollowByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.FollowByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowDomain) Unfollow(
    request friend.UnfollowByUserIdRequest,
) (*friend.UnfollowByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.UnfollowByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowDomain) List(
    withProfile *bool,
) iterator.DescribeFollowsByUserIdIterator {
    return iterator.NewDescribeFollowsByUserIdIterator(
        p.followUserCache,
        p.client,
        p.namespaceName,
        p.userId,
        withProfile,
    )
}
