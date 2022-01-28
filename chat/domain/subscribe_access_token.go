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

type SubscribeAccessTokenDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    subscribeCache cache.SubscribeDomainCache
    namespaceName string
    accessToken auth.AccessToken
    roomName string
}

func NewSubscribeAccessTokenDomain(
    session *core.Gs2RestSession,
    subscribeCache cache.SubscribeDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    roomName string,
) SubscribeAccessTokenDomain {
    return SubscribeAccessTokenDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        subscribeCache: subscribeCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        roomName: roomName,
    }
}

func (p SubscribeAccessTokenDomain) Subscribe(
    request chat.SubscribeRequest,
) (*chat.SubscribeResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    r, err := p.client.Subscribe(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeAccessTokenDomain) Load(
    request chat.GetSubscribeRequest,
) (*chat.GetSubscribeResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    r, err := p.client.GetSubscribe(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeAccessTokenDomain) UpdateNotificationType(
    request chat.UpdateNotificationTypeRequest,
) (*chat.UpdateNotificationTypeResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    r, err := p.client.UpdateNotificationType(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeAccessTokenDomain) Unsubscribe(
    request chat.UnsubscribeRequest,
) (*chat.UnsubscribeResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.RoomName = &p.roomName
    r, err := p.client.Unsubscribe(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}
