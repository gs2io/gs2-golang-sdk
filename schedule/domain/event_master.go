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

type EventMasterDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    eventMasterCache cache.EventMasterDomainCache
    namespaceName string
    eventName string
}

func NewEventMasterDomain(
    session *core.Gs2RestSession,
    eventMasterCache cache.EventMasterDomainCache,
    namespaceName string,
    eventName string,
) EventMasterDomain {
    return EventMasterDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        eventMasterCache: eventMasterCache,
        namespaceName: namespaceName,
        eventName: eventName,
    }
}

func (p EventMasterDomain) Load(
    request schedule.GetEventMasterRequest,
) (*schedule.GetEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EventName = &p.eventName
    r, err := p.client.GetEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.eventMasterCache.Update(*r.Item)
    return r, nil
}

func (p EventMasterDomain) Update(
    request schedule.UpdateEventMasterRequest,
) (*schedule.UpdateEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EventName = &p.eventName
    r, err := p.client.UpdateEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.eventMasterCache.Update(*r.Item)
    return r, nil
}

func (p EventMasterDomain) Delete(
    request schedule.DeleteEventMasterRequest,
) (*schedule.DeleteEventMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.EventName = &p.eventName
    r, err := p.client.DeleteEventMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.eventMasterCache.Delete(*r.Item)
    return r, nil
}
