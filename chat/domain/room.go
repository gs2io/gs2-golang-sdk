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

type RoomDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    roomCache cache.RoomDomainCache
    namespaceName string
    userId string
    roomName string
    password *string
    messageCache cache.MessageDomainCache
}

func NewRoomDomain(
    session *core.Gs2RestSession,
    roomCache cache.RoomDomainCache,
    namespaceName string,
    userId string,
    roomName string,
    password *string,
) RoomDomain {
    return RoomDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        roomCache: roomCache,
        namespaceName: namespaceName,
        userId: userId,
        roomName: roomName,
        password: password,
        messageCache: cache.NewMessageDomainCache(),
    }
}

func (p RoomDomain) Load(
    request chat.GetRoomRequest,
) (*chat.GetRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.RoomName = &p.roomName
    r, err := p.client.GetRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Update(*r.Item)
    return r, nil
}

func (p RoomDomain) UpdateFromBackend(
    request chat.UpdateRoomFromBackendRequest,
) (*chat.UpdateRoomFromBackendResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    request.Password = p.password
    r, err := p.client.UpdateRoomFromBackend(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Update(*r.Item)
    return r, nil
}

func (p RoomDomain) DeleteFromBackend(
    request chat.DeleteRoomFromBackendRequest,
) (*chat.DeleteRoomFromBackendResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    r, err := p.client.DeleteRoomFromBackend(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Delete(*r.Item)
    return r, nil
}

func (p RoomDomain) Post(
    request chat.PostByUserIdRequest,
) (*chat.PostByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    request.Password = p.password
    r, err := p.client.PostByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p RoomDomain) Messages(
) iterator.DescribeMessagesByUserIdIterator {
    return iterator.NewDescribeMessagesByUserIdIterator(
        p.messageCache,
        p.client,
        p.namespaceName,
        p.roomName,
        p.password,
        p.userId,
    )
}

func (p RoomDomain) Message(
    messageName string,
) MessageDomain {
    return NewMessageDomain(
        p.session,
        p.messageCache,
        p.namespaceName,
        p.userId,
        p.roomName,
        messageName,
        p.password,
    )
}
