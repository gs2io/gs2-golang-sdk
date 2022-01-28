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

type UserDomain struct {
    session *core.Gs2RestSession
    client exchange.Gs2ExchangeRestClient
    namespaceName string
    userId string
    awaitCache cache.AwaitDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: exchange.Gs2ExchangeRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        awaitCache: cache.NewAwaitDomainCache(),
    }
}

func (p UserDomain) Exchange(
    request exchange.ExchangeByUserIdRequest,
) (*exchange.ExchangeByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.ExchangeByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) CreateAwait(
    request exchange.CreateAwaitByUserIdRequest,
) (*exchange.CreateAwaitByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    r, err := p.client.CreateAwaitByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p UserDomain) Awaits(
    rateName *string,
) iterator.DescribeAwaitsByUserIdIterator {
    return iterator.NewDescribeAwaitsByUserIdIterator(
        p.awaitCache,
        p.client,
        p.namespaceName,
        p.userId,
        rateName,
    )
}

func (p UserDomain) Await(
    awaitName string,
) AwaitDomain {
    return NewAwaitDomain(
        p.session,
        p.awaitCache,
        p.namespaceName,
        p.userId,
        awaitName,
    )
}
