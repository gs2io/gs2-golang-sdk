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
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Exchange struct {
    session *core.Gs2RestSession
    client exchange.Gs2ExchangeRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Exchange(
    session *core.Gs2RestSession,
) Gs2Exchange {
    return Gs2Exchange {
        session: session,
        client: exchange.Gs2ExchangeRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Exchange) CreateNamespace(
    request exchange.CreateNamespaceRequest,
) (*exchange.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Exchange) ExchangeByStampSheet(
    request exchange.ExchangeByStampSheetRequest,
) (*exchange.ExchangeByStampSheetResult, error) {
    r, err := p.client.ExchangeByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Exchange) CreateAwaitByStampSheet(
    request exchange.CreateAwaitByStampSheetRequest,
) (*exchange.CreateAwaitByStampSheetResult, error) {
    r, err := p.client.CreateAwaitByStampSheet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Exchange) DeleteAwaitByStampTask(
    request exchange.DeleteAwaitByStampTaskRequest,
) (*exchange.DeleteAwaitByStampTaskResult, error) {
    r, err := p.client.DeleteAwaitByStampTask(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Exchange) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Exchange) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
