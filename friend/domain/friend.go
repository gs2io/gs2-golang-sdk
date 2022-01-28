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

type FriendDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    friendUserCache cache.FriendUserDomainCache
    namespaceName string
    userId string
}

func NewFriendDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) FriendDomain {
    return FriendDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        friendUserCache: cache.NewFriendUserDomainCache(),
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p FriendDomain) Load(
    request friend.GetFriendByUserIdRequest,
) (*friend.GetFriendByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetFriendByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FriendDomain) Delete(
    request friend.DeleteFriendByUserIdRequest,
) (*friend.DeleteFriendByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteFriendByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p FriendDomain) List(
    withProfile *bool,
) iterator.DescribeFriendsByUserIdIterator {
    return iterator.NewDescribeFriendsByUserIdIterator(
        p.friendUserCache,
        p.client,
        p.namespaceName,
        p.userId,
        withProfile,
    )
}
