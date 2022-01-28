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

type UserAccessTokenDomain struct {
    session *core.Gs2RestSession
    client exchange.Gs2ExchangeRestClient
    namespaceName string
    accessToken auth.AccessToken
    awaitCache cache.AwaitDomainCache
}

func NewUserAccessTokenDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return UserAccessTokenDomain {
        session: session,
        client: exchange.Gs2ExchangeRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        accessToken: accessToken,
        awaitCache: cache.NewAwaitDomainCache(),
    }
}

func (p UserAccessTokenDomain) Exchange(
    request exchange.ExchangeRequest,
) (*exchange.ExchangeResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    r, err := p.client.Exchange(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserAccessTokenDomain) Awaits(
    rateName *string,
) iterator.DescribeAwaitsIterator {
    return iterator.NewDescribeAwaitsIterator(
        p.awaitCache,
        p.client,
        p.namespaceName,
        p.accessToken,
        rateName,
    )
}

func (p UserAccessTokenDomain) Await(
    awaitName string,
) AwaitAccessTokenDomain {
    return NewAwaitAccessTokenDomain(
        p.session,
        p.awaitCache,
        p.namespaceName,
        p.accessToken,
        awaitName,
    )
}
