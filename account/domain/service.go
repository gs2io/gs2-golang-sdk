package domainaccount
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

    "github.com/gs2io/gs2-golang-sdk/account"
    "github.com/gs2io/gs2-golang-sdk/account/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/account/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = iterator.DescribeNamespacesIterator{}

type Gs2Account struct {
    session *core.Gs2RestSession
    client account.Gs2AccountRestClient
    namespaceCache cache.NamespaceDomainCache
}

func NewGs2Account(
    session *core.Gs2RestSession,
) Gs2Account {
    return Gs2Account {
        session: session,
        client: account.Gs2AccountRestClient {
            Session: session,
        },
        namespaceCache: cache.NewNamespaceDomainCache(),
    }
}

func (p Gs2Account) CreateNamespace(
    request account.CreateNamespaceRequest,
) (*account.CreateNamespaceResult, error) {
    r, err := p.client.CreateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p Gs2Account) DeleteTakeOverByUserIdentifier(
    request account.DeleteTakeOverByUserIdentifierRequest,
) (*account.DeleteTakeOverByUserIdentifierResult, error) {
    r, err := p.client.DeleteTakeOverByUserIdentifier(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Account) DoTakeOver(
    request account.DoTakeOverRequest,
) (*account.DoTakeOverResult, error) {
    r, err := p.client.DoTakeOver(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p Gs2Account) Namespaces(
) iterator.DescribeNamespacesIterator {
    return iterator.NewDescribeNamespacesIterator(
        p.namespaceCache,
        p.client,
    )
}

func (p Gs2Account) Namespace(
    namespaceName string,
) NamespaceDomain {
    return NewNamespaceDomain(
        p.session,
        p.namespaceCache,
        namespaceName,
    )
}
