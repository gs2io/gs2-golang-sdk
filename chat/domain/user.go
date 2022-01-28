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

type UserDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    namespaceName string
    userId string
    roomCache cache.RoomDomainCache
    subscribeCache cache.SubscribeDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        roomCache: cache.NewRoomDomainCache(),
        subscribeCache: cache.NewSubscribeDomainCache(),
    }
}

func (p UserDomain) CreateRoomFromBackend(
    request chat.CreateRoomFromBackendRequest,
) (*chat.CreateRoomFromBackendResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CreateRoomFromBackend(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Rooms(
) iterator.DescribeRoomsIterator {
    return iterator.NewDescribeRoomsIterator(
        p.roomCache,
        p.client,
        p.namespaceName,
    )
}

func (p UserDomain) Subscribes(
) iterator.DescribeSubscribesByUserIdIterator {
    return iterator.NewDescribeSubscribesByUserIdIterator(
        p.subscribeCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Room(
    roomName string,
    password *string,
) RoomDomain {
    return NewRoomDomain(
        p.session,
        p.roomCache,
        p.namespaceName,
        p.userId,
        roomName,
        password,
    )
}

func (p UserDomain) Subscribe(
    roomName string,
) SubscribeDomain {
    return NewSubscribeDomain(
        p.session,
        p.subscribeCache,
        p.namespaceName,
        p.userId,
        roomName,
    )
}
