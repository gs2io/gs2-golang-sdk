package domainmission
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

    "github.com/gs2io/gs2-golang-sdk/mission"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/mission/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    namespaceName string
    userId string
    completeCache cache.CompleteDomainCache
    counterCache cache.CounterDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        completeCache: cache.NewCompleteDomainCache(),
        counterCache: cache.NewCounterDomainCache(),
    }
}

func (p UserDomain) Completes(
) iterator.DescribeCompletesByUserIdIterator {
    return iterator.NewDescribeCompletesByUserIdIterator(
        p.completeCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Counters(
) iterator.DescribeCountersByUserIdIterator {
    return iterator.NewDescribeCountersByUserIdIterator(
        p.counterCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Counter(
    counterName string,
) CounterDomain {
    return NewCounterDomain(
        p.session,
        p.counterCache,
        p.namespaceName,
        p.userId,
        counterName,
    )
}

func (p UserDomain) Complete(
    missionGroupName string,
) CompleteDomain {
    return NewCompleteDomain(
        p.session,
        p.completeCache,
        p.namespaceName,
        p.userId,
        missionGroupName,
    )
}
