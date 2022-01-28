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

type UserDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    namespaceName string
    userId string
    followUserCache cache.FollowUserDomainCache
    friendUserCache cache.FriendUserDomainCache
    friendRequestCache cache.FriendRequestDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        followUserCache: cache.NewFollowUserDomainCache(),
        friendUserCache: cache.NewFriendUserDomainCache(),
        friendRequestCache: cache.NewFriendRequestDomainCache(),
    }
}

func (p UserDomain) Profile(
) ProfileDomain {
    return NewProfileDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) BlackList(
) BlackListDomain {
    return NewBlackListDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Follow(
) FollowDomain {
    return NewFollowDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Friend(
) FriendDomain {
    return NewFriendDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) SendBox(
) SendBoxDomain {
    return NewSendBoxDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Inbox(
) InboxDomain {
    return NewInboxDomain(
        p.session,
        p.namespaceName,
        p.userId,
    )
}
