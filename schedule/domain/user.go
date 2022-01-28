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

type UserDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    namespaceName string
    userId string
    triggerCache cache.TriggerDomainCache
    eventCache cache.EventDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        triggerCache: cache.NewTriggerDomainCache(),
        eventCache: cache.NewEventDomainCache(),
    }
}

func (p UserDomain) Triggers(
) iterator.DescribeTriggersByUserIdIterator {
    return iterator.NewDescribeTriggersByUserIdIterator(
        p.triggerCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Events(
) iterator.DescribeEventsByUserIdIterator {
    return iterator.NewDescribeEventsByUserIdIterator(
        p.eventCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) RawEvents(
) iterator.DescribeRawEventsIterator {
    return iterator.NewDescribeRawEventsIterator(
        p.eventCache,
        p.client,
        p.namespaceName,
    )
}

func (p UserDomain) Trigger(
    triggerName string,
) TriggerDomain {
    return NewTriggerDomain(
        p.session,
        p.triggerCache,
        p.namespaceName,
        p.userId,
        triggerName,
    )
}

func (p UserDomain) Event(
    eventName string,
) EventDomain {
    return NewEventDomain(
        p.session,
        p.eventCache,
        p.namespaceName,
        p.userId,
        eventName,
    )
}
