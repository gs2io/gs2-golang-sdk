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

type CounterDomain struct {
    session *core.Gs2RestSession
    client limit.Gs2LimitRestClient
    counterCache cache.CounterDomainCache
    namespaceName string
    userId string
    limitName string
    counterName string
}

func NewCounterDomain(
    session *core.Gs2RestSession,
    counterCache cache.CounterDomainCache,
    namespaceName string,
    userId string,
    limitName string,
    counterName string,
) CounterDomain {
    return CounterDomain {
        session: session,
        client: limit.Gs2LimitRestClient{
            Session: session,
        },
        counterCache: counterCache,
        namespaceName: namespaceName,
        userId: userId,
        limitName: limitName,
        counterName: counterName,
    }
}

func (p CounterDomain) Load(
    request limit.GetCounterByUserIdRequest,
) (*limit.GetCounterByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.LimitName = &p.limitName
    request.CounterName = &p.counterName
    r, err := p.client.GetCounterByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterCache.Update(*r.Item)
    return r, nil
}

func (p CounterDomain) CountUp(
    request limit.CountUpByUserIdRequest,
) (*limit.CountUpByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.LimitName = &p.limitName
    request.CounterName = &p.counterName
    r, err := p.client.CountUpByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterCache.Update(*r.Item)
    return r, nil
}

func (p CounterDomain) Delete(
    request limit.DeleteCounterByUserIdRequest,
) (*limit.DeleteCounterByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.LimitName = &p.limitName
    request.CounterName = &p.counterName
    r, err := p.client.DeleteCounterByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.counterCache.Delete(*r.Item)
    return r, nil
}
