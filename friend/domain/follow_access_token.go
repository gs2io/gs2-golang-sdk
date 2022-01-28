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

type FollowAccessTokenDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    followUserCache cache.FollowUserDomainCache
    namespaceName string
    accessToken auth.AccessToken
}

func NewFollowAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) FollowAccessTokenDomain {
    return FollowAccessTokenDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        followUserCache: cache.NewFollowUserDomainCache(),
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p FollowAccessTokenDomain) Load(
    request friend.GetFollowRequest,
) (*friend.GetFollowResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.GetFollow(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowAccessTokenDomain) Follow(
    request friend.FollowRequest,
) (*friend.FollowResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.Follow(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowAccessTokenDomain) Unfollow(
    request friend.UnfollowRequest,
) (*friend.UnfollowResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.Unfollow(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FollowAccessTokenDomain) List(
    withProfile *bool,
) iterator.DescribeFollowsIterator {
    return iterator.NewDescribeFollowsIterator(
        p.followUserCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        withProfile,
    )
}
