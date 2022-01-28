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

type InboxAccessTokenDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    friendRequestCache cache.FriendRequestDomainCache
    namespaceName string
    accessToken auth.AccessToken
}

func NewInboxAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) InboxAccessTokenDomain {
    return InboxAccessTokenDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        friendRequestCache: cache.NewFriendRequestDomainCache(),
        namespaceName: namespaceName,
        accessToken: accessToken,
    }
}

func (p InboxAccessTokenDomain) GetReceiveRequest(
    request friend.GetReceiveRequestRequest,
) (*friend.GetReceiveRequestResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.GetReceiveRequest(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p InboxAccessTokenDomain) AcceptRequest(
    request friend.AcceptRequestRequest,
) (*friend.AcceptRequestResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.AcceptRequest(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p InboxAccessTokenDomain) RejectRequest(
    request friend.RejectRequestRequest,
) (*friend.RejectRequestResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.RejectRequest(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p InboxAccessTokenDomain) List(
) iterator.DescribeReceiveRequestsIterator {
    return iterator.NewDescribeReceiveRequestsIterator(
        p.friendRequestCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}
