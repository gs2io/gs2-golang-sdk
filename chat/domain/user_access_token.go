package domainchat
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

    "github.com/gs2io/gs2-golang-sdk/chat"
    "github.com/gs2io/gs2-golang-sdk/chat/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/chat/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    namespaceName string
    accessToken auth.AccessToken
    roomCache cache.RoomDomainCache
    subscribeCache cache.SubscribeDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        roomCache: cache.NewRoomDomainCache(),
        subscribeCache: cache.NewSubscribeDomainCache(),
    }
}

func (p UserAccessTokenDomain) CreateRoom(
    request chat.CreateRoomRequest,
) (*chat.CreateRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.CreateRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Subscribes(
) iterator.DescribeSubscribesIterator {
    return iterator.NewDescribeSubscribesIterator(
        p.subscribeCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Room(
    roomName string,
    password *string,
) RoomAccessTokenDomain {
    return NewRoomAccessTokenDomain(
        p.session,
        p.roomCache,
        p.namespaceName,
        p.accessToken,
        roomName,
        password,
    )
}

func (p UserAccessTokenDomain) Subscribe(
    roomName string,
) SubscribeAccessTokenDomain {
    return NewSubscribeAccessTokenDomain(
        p.session,
        p.subscribeCache,
        p.namespaceName,
        p.accessToken,
        roomName,
    )
}
