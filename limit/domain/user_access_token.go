package domainlimit
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

    "github.com/gs2io/gs2-golang-sdk/limit"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/limit/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    namespaceName string
    accessToken auth.AccessToken
    counterCache cache.CounterDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        counterCache: cache.NewCounterDomainCache(),
    }
}

func (p UserAccessTokenDomain) Counters(
    limitName *string,
) iterator.DescribeCountersIterator {
    return iterator.NewDescribeCountersIterator(
        p.counterCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        limitName,
    )
}

func (p UserAccessTokenDomain) Counter(
    limitName string,
    counterName string,
) CounterAccessTokenDomain {
    return NewCounterAccessTokenDomain(
        p.session,
        p.counterCache,
        p.namespaceName,
        p.accessToken,
        limitName,
        counterName,
    )
}
