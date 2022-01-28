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

type AwaitAccessTokenDomain struct {
    session *core.Gs2RestSession
    client exchange.Gs2ExchangeRestClient
    awaitCache cache.AwaitDomainCache
    namespaceName string
    accessToken auth.AccessToken
    awaitName string
}

func NewAwaitAccessTokenDomain(
    session *core.Gs2RestSession,
    awaitCache cache.AwaitDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    awaitName string,
) AwaitAccessTokenDomain {
    return AwaitAccessTokenDomain {
        session: session,
        client: exchange.Gs2ExchangeRestClient{
            Session: session,
        },
        awaitCache: awaitCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        awaitName: awaitName,
    }
}

func (p AwaitAccessTokenDomain) Load(
    request exchange.GetAwaitRequest,
) (*exchange.GetAwaitResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.AwaitName = &p.awaitName
    r, err := p.client.GetAwait(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitAccessTokenDomain) Acquire(
    request exchange.AcquireRequest,
) (*exchange.AcquireResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.AwaitName = &p.awaitName
    r, err := p.client.Acquire(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitAccessTokenDomain) Skip(
    request exchange.SkipRequest,
) (*exchange.SkipResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.AwaitName = &p.awaitName
    r, err := p.client.Skip(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Update(*r.Item)
    return r, nil
}

func (p AwaitAccessTokenDomain) Delete(
    request exchange.DeleteAwaitRequest,
) (*exchange.DeleteAwaitResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.AwaitName = &p.awaitName
    r, err := p.client.DeleteAwait(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.awaitCache.Delete(*r.Item)
    return r, nil
}
