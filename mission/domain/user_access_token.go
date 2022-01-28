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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client mission.Gs2MissionRestClient
    namespaceName string
    accessToken auth.AccessToken
    completeCache cache.CompleteDomainCache
    counterCache cache.CounterDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: mission.Gs2MissionRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        completeCache: cache.NewCompleteDomainCache(),
        counterCache: cache.NewCounterDomainCache(),
    }
}

func (p UserAccessTokenDomain) Completes(
) iterator.DescribeCompletesIterator {
    return iterator.NewDescribeCompletesIterator(
        p.completeCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Counters(
) iterator.DescribeCountersIterator {
    return iterator.NewDescribeCountersIterator(
        p.counterCache,
        p.client,
        p.namespaceName,
        p.accessToken,
    )
}

func (p UserAccessTokenDomain) Counter(
    counterName string,
) CounterAccessTokenDomain {
    return NewCounterAccessTokenDomain(
        p.session,
        p.counterCache,
        p.namespaceName,
        p.accessToken,
        counterName,
    )
}

func (p UserAccessTokenDomain) Complete(
    missionGroupName string,
) CompleteAccessTokenDomain {
    return NewCompleteAccessTokenDomain(
        p.session,
        p.completeCache,
        p.namespaceName,
        p.accessToken,
        missionGroupName,
    )
}
