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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client schedule.Gs2ScheduleRestClient
    namespaceName string
    accessToken auth.AccessToken
    triggerCache cache.TriggerDomainCache
    eventCache cache.EventDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: schedule.Gs2ScheduleRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        triggerCache: cache.NewTriggerDomainCache(),
        eventCache: cache.NewEventDomainCache(),
    }
}

func (p UserAccessTokenDomain) Triggers(
) iterator.DescribeTriggersIterator {
    return iterator.NewDescribeTriggersIterator(
        p.triggerCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Events(
) iterator.DescribeEventsIterator {
    return iterator.NewDescribeEventsIterator(
        p.eventCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Trigger(
    triggerName string,
) TriggerAccessTokenDomain {
    return NewTriggerAccessTokenDomain(
        p.session,
        p.triggerCache,
        p.namespaceName,
        p.accessToken,
        triggerName,
    )
}

func (p UserAccessTokenDomain) Event(
    eventName string,
) EventAccessTokenDomain {
    return NewEventAccessTokenDomain(
        p.session,
        p.eventCache,
        p.namespaceName,
        p.accessToken,
        eventName,
    )
}
