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
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client account.Gs2AccountRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    accountCache cache.AccountDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: account.Gs2AccountRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        accountCache: cache.NewAccountDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request account.GetNamespaceStatusRequest,
) (*account.GetNamespaceStatusResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespaceStatus(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Load(
    request account.GetNamespaceRequest,
) (*account.GetNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.GetNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Update(
    request account.UpdateNamespaceRequest,
) (*account.UpdateNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.UpdateNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Update(*r.Item)
    return r, nil
}

func (p NamespaceDomain) Delete(
    request account.DeleteNamespaceRequest,
) (*account.DeleteNamespaceResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.DeleteNamespace(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.namespaceCache.Delete(*r.Item)
    return r, nil
}

func (p NamespaceDomain) CreateAccount(
    request account.CreateAccountRequest,
) (*account.CreateAccountResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateAccount(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) Accounts(
) iterator.DescribeAccountsIterator {
    return iterator.NewDescribeAccountsIterator(
        p.accountCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) Account(
    userId string,
) AccountDomain {
    return NewAccountDomain(
        p.session,
        p.accountCache,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) AccountAccessTokenDomain {
    return NewAccountAccessTokenDomain(
        p.session,
        p.accountCache,
        p.namespaceName,
        accessToken,
    )
}
