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

type UserDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    namespaceName string
    userId string
    counterCache cache.CounterDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        counterCache: cache.NewCounterDomainCache(),
    }
}

func (p UserDomain) Counters(
    limitName *string,
) iterator.DescribeCountersByUserIdIterator {
    return iterator.NewDescribeCountersByUserIdIterator(
        p.counterCache,
        p.client,
        p.namespaceName,
        p.userId,
        limitName,
    )
}

func (p UserDomain) Counter(
    limitName string,
    counterName string,
) CounterDomain {
    return NewCounterDomain(
        p.session,
        p.counterCache,
        p.namespaceName,
        p.userId,
        limitName,
        counterName,
    )
}
