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

type BlackListDomain struct {
    session *core.Gs2RestSession
    client friend.Gs2FriendRestClient
    namespaceName string
    userId string
}

func NewBlackListDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) BlackListDomain {
    return BlackListDomain {
        session: session,
        client: friend.Gs2FriendRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
    }
}

func (p BlackListDomain) Register(
    request friend.RegisterBlackListByUserIdRequest,
) (*friend.RegisterBlackListByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.RegisterBlackListByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p BlackListDomain) Unregister(
    request friend.UnregisterBlackListByUserIdRequest,
) (*friend.UnregisterBlackListByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.UnregisterBlackListByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p BlackListDomain) List(
) iterator.DescribeBlackListByUserIdIterator {
    return iterator.NewDescribeBlackListByUserIdIterator(
        p.client,
        p.namespaceName,
        p.userId,
    )
}
