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

type RoomAccessTokenDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    roomCache cache.RoomDomainCache
    namespaceName string
    accessToken auth.AccessToken
    roomName string
    password *string
    messageCache cache.MessageDomainCache
}

func NewRoomAccessTokenDomain(
    session *core.Gs2RestSession,
    roomCache cache.RoomDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    roomName string,
    password *string,
) RoomAccessTokenDomain {
    return RoomAccessTokenDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        roomCache: roomCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        roomName: roomName,
        password: password,
        messageCache: cache.NewMessageDomainCache(),
    }
}

func (p RoomAccessTokenDomain) Update(
    request chat.UpdateRoomRequest,
) (*chat.UpdateRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    request.Password = p.password
    r, err := p.client.UpdateRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Update(*r.Item)
    return r, nil
}

func (p RoomAccessTokenDomain) Delete(
    request chat.DeleteRoomRequest,
) (*chat.DeleteRoomResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    r, err := p.client.DeleteRoom(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.roomCache.Delete(*r.Item)
    return r, nil
}

func (p RoomAccessTokenDomain) Post(
    request chat.PostRequest,
) (*chat.PostResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    request.Password = p.password
    r, err := p.client.Post(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p RoomAccessTokenDomain) Messages(
) iterator.DescribeMessagesIterator {
    return iterator.NewDescribeMessagesIterator(
        p.messageCache,
        p.client,
        p.namespaceName,
        p.roomName,
        p.password,
        p.accessToken,
    )
}

func (p RoomAccessTokenDomain) Message(
    messageName string,
) MessageAccessTokenDomain {
    return NewMessageAccessTokenDomain(
        p.session,
        p.messageCache,
        p.namespaceName,
        p.accessToken,
        p.roomName,
        messageName,
        p.password,
    )
}
