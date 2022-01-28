package domainexchange
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

    "github.com/gs2io/gs2-golang-sdk/exchange"
    "github.com/gs2io/gs2-golang-sdk/exchange/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/exchange/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type AwaitDomain struct {
    session *core.Gs2RestSession
    client exchange.Gs2ExchangeRestClient
    awaitCache cache.AwaitDomainCache
    namespaceName string
    userId string
    awaitName string
}

func NewAwaitDomain(
    session *core.Gs2RestSession,
    awaitCache cache.AwaitDomainCache,
    namespaceName string,
    userId string,
    awaitName string,
) AwaitDomain {
    return AwaitDomain {
        session: session,
        client: exchange.Gs2ExchangeRestClient{
            Session: session,
        },
        awaitCache: awaitCache,
        namespaceName: namespaceName,
        userId: userId,
        awaitName: awaitName,
    }
}

func (p AwaitDomain) Load(
    request exchange.GetAwaitByUserIdRequest,
) (*exchange.GetAwaitByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.AwaitName = &p.awaitName
    r, err := p.client.GetAwaitByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitDomain) Acquire(
    request exchange.AcquireByUserIdRequest,
) (*exchange.AcquireByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.AwaitName = &p.awaitName
    r, err := p.client.AcquireByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitDomain) AcquireForce(
    request exchange.AcquireForceByUserIdRequest,
) (*exchange.AcquireForceByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.AwaitName = &p.awaitName
    r, err := p.client.AcquireForceByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitDomain) Skip(
    request exchange.SkipByUserIdRequest,
) (*exchange.SkipByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.AwaitName = &p.awaitName
    r, err := p.client.SkipByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitDomain) Delete(
    request exchange.DeleteAwaitByUserIdRequest,
) (*exchange.DeleteAwaitByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.AwaitName = &p.awaitName
    r, err := p.client.DeleteAwaitByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Delete(*r.Item)
    return r, nil
}
