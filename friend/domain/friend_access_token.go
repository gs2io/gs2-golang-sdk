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

type FriendAccessTokenDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    friendUserCache cache.FriendUserDomainCache
    namespaceName string
    accessToken auth.AccessToken
}

func NewFriendAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) FriendAccessTokenDomain {
    return FriendAccessTokenDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        friendUserCache: cache.NewFriendUserDomainCache(),
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p FriendAccessTokenDomain) Load(
    request friend.GetFriendRequest,
) (*friend.GetFriendResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.GetFriend(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FriendAccessTokenDomain) Delete(
    request friend.DeleteFriendRequest,
) (*friend.DeleteFriendResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.DeleteFriend(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FriendAccessTokenDomain) List(
    withProfile *bool,
) iterator.DescribeFriendsIterator {
    return iterator.NewDescribeFriendsIterator(
        p.friendUserCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        withProfile,
    )
}
