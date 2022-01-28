package domainschedule
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

    "github.com/gs2io/gs2-golang-sdk/schedule"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/schedule/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type EventDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    eventCache cache.EventDomainCache
    namespaceName string
    userId string
    eventName string
}

func NewEventDomain(
    session *core.Gs2RestSession,
    eventCache cache.EventDomainCache,
    namespaceName string,
    userId string,
    eventName string,
) EventDomain {
    return EventDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        eventCache: eventCache,
        namespaceName: namespaceName,
        userId: userId,
        eventName: eventName,
    }
}

func (p EventDomain) Load(
    request schedule.GetEventByUserIdRequest,
) (*schedule.GetEventByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.EventName = &p.eventName
    r, err := p.client.GetEventByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.eventCache.Update(*r.Item)
    return r, nil
}

func (p EventDomain) GetRaw(
    request schedule.GetRawEventRequest,
) (*schedule.GetRawEventResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EventName = &p.eventName
    r, err := p.client.GetRawEvent(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.eventCache.Update(*r.Item)
    return r, nil
}
