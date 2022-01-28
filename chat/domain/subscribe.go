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

type SubscribeDomain struct {
    session *core.Gs2RestSession
    client chat.Gs2ChatRestClient
    subscribeCache cache.SubscribeDomainCache
    namespaceName string
    userId string
    roomName string
}

func NewSubscribeDomain(
    session *core.Gs2RestSession,
    subscribeCache cache.SubscribeDomainCache,
    namespaceName string,
    userId string,
    roomName string,
) SubscribeDomain {
    return SubscribeDomain {
        session: session,
        client: chat.Gs2ChatRestClient{
            Session: session,
        },
        subscribeCache: subscribeCache,
        namespaceName: namespaceName,
        userId: userId,
        roomName: roomName,
    }
}

func (p SubscribeDomain) Subscribe(
    request chat.SubscribeByUserIdRequest,
) (*chat.SubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    r, err := p.client.SubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeDomain) Load(
    request chat.GetSubscribeByUserIdRequest,
) (*chat.GetSubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    r, err := p.client.GetSubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeDomain) UpdateNotificationType(
    request chat.UpdateNotificationTypeByUserIdRequest,
) (*chat.UpdateNotificationTypeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    r, err := p.client.UpdateNotificationTypeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}

func (p SubscribeDomain) Unsubscribe(
    request chat.UnsubscribeByUserIdRequest,
) (*chat.UnsubscribeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.RoomName = &p.roomName
    r, err := p.client.UnsubscribeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.subscribeCache.Update(*r.Item)
    return r, nil
}
