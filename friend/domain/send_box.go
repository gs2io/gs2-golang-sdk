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

type SendBoxDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    friendRequestCache cache.FriendRequestDomainCache
    namespaceName string
    userId string
}

func NewSendBoxDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) SendBoxDomain {
    return SendBoxDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        friendRequestCache: cache.NewFriendRequestDomainCache(),
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p SendBoxDomain) GetSendRequest(
    request friend.GetSendRequestByUserIdRequest,
) (*friend.GetSendRequestByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.GetSendRequestByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p SendBoxDomain) SendRequest(
    request friend.SendRequestByUserIdRequest,
) (*friend.SendRequestByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.SendRequestByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p SendBoxDomain) DeleteRequest(
    request friend.DeleteRequestByUserIdRequest,
) (*friend.DeleteRequestByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.DeleteRequestByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p SendBoxDomain) List(
) iterator.DescribeSendRequestsByUserIdIterator {
    return iterator.NewDescribeSendRequestsByUserIdIterator(
        p.friendRequestCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}
