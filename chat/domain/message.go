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

type MessageDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    messageCache cache.MessageDomainCache
    namespaceName string
    userId string
    roomName string
    messageName string
    password *string
}

func NewMessageDomain(
    session *core.Gs2RestSession,
    messageCache cache.MessageDomainCache,
    namespaceName string,
    userId string,
    roomName string,
    messageName string,
    password *string,
) MessageDomain {
    return MessageDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        messageCache: messageCache,
        namespaceName: namespaceName,
        userId: userId,
        roomName: roomName,
        messageName: messageName,
        password: password,
    }
}

func (p MessageDomain) Load(
    request chat.GetMessageByUserIdRequest,
) (*chat.GetMessageByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    request.MessageName = &p.messageName
    request.Password = p.password
    r, err := p.client.GetMessageByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Update(*r.Item)
    return r, nil
}

func (p MessageDomain) Delete(
    request chat.DeleteMessageRequest,
) (*chat.DeleteMessageResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    request.MessageName = &p.messageName
    r, err := p.client.DeleteMessage(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.messageCache.Delete(*r.Item)
    return r, nil
}
