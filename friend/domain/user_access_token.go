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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    namespaceName string
    accessToken auth.AccessToken
    followUserCache cache.FollowUserDomainCache
    friendUserCache cache.FriendUserDomainCache
    friendRequestCache cache.FriendRequestDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        followUserCache: cache.NewFollowUserDomainCache(),
        friendUserCache: cache.NewFriendUserDomainCache(),
        friendRequestCache: cache.NewFriendRequestDomainCache(),
    }
}

func (p UserAccessTokenDomain) Profile(
) ProfileAccessTokenDomain {
    return NewProfileAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) BlackList(
) BlackListAccessTokenDomain {
    return NewBlackListAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Follow(
) FollowAccessTokenDomain {
    return NewFollowAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Friend(
) FriendAccessTokenDomain {
    return NewFriendAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) SendBox(
) SendBoxAccessTokenDomain {
    return NewSendBoxAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Inbox(
) InboxAccessTokenDomain {
    return NewInboxAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
    )
}
