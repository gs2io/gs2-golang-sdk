package domainshowcase
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

    "github.com/gs2io/gs2-golang-sdk/showcase"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/showcase/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type NamespaceDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    namespaceCache cache.NamespaceDomainCache
    namespaceName string
    salesItemMasterCache cache.SalesItemMasterDomainCache
    salesItemGroupMasterCache cache.SalesItemGroupMasterDomainCache
    showcaseMasterCache cache.ShowcaseMasterDomainCache
}

func NewNamespaceDomain(
    session *core.Gs2RestSession,
    namespaceCache cache.NamespaceDomainCache,
    namespaceName string,
) NamespaceDomain {
    return NamespaceDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        namespaceCache: namespaceCache,
        namespaceName: namespaceName,
        salesItemMasterCache: cache.NewSalesItemMasterDomainCache(),
        salesItemGroupMasterCache: cache.NewSalesItemGroupMasterDomainCache(),
        showcaseMasterCache: cache.NewShowcaseMasterDomainCache(),
    }
}

func (p NamespaceDomain) GetStatus(
    request showcase.GetNamespaceStatusRequest,
) (*showcase.GetNamespaceStatusResult, error) {
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
    request showcase.GetNamespaceRequest,
) (*showcase.GetNamespaceResult, error) {
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
    request showcase.UpdateNamespaceRequest,
) (*showcase.UpdateNamespaceResult, error) {
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
    request showcase.DeleteNamespaceRequest,
) (*showcase.DeleteNamespaceResult, error) {
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

func (p NamespaceDomain) CreateSalesItemMaster(
    request showcase.CreateSalesItemMasterRequest,
) (*showcase.CreateSalesItemMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateSalesItemMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateSalesItemGroupMaster(
    request showcase.CreateSalesItemGroupMasterRequest,
) (*showcase.CreateSalesItemGroupMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateSalesItemGroupMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) CreateShowcaseMaster(
    request showcase.CreateShowcaseMasterRequest,
) (*showcase.CreateShowcaseMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    r, err := p.client.CreateShowcaseMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p NamespaceDomain) SalesItemMasters(
) iterator.DescribeSalesItemMastersIterator {
    return iterator.NewDescribeSalesItemMastersIterator(
        p.salesItemMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) SalesItemGroupMasters(
) iterator.DescribeSalesItemGroupMastersIterator {
    return iterator.NewDescribeSalesItemGroupMastersIterator(
        p.salesItemGroupMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) ShowcaseMasters(
) iterator.DescribeShowcaseMastersIterator {
    return iterator.NewDescribeShowcaseMastersIterator(
        p.showcaseMasterCache,
        p.client,
        p.namespaceName,
    )
}

func (p NamespaceDomain) CurrentShowcaseMaster(
) CurrentShowcaseMasterDomain {
    return NewCurrentShowcaseMasterDomain(
        p.session,
        p.namespaceName,
    )
}

func (p NamespaceDomain) User(
    userId string,
) UserDomain {
    return NewUserDomain(
        p.session,
        p.namespaceName,
        userId,
    )
}

func (p NamespaceDomain) AccessToken(
    accessToken auth.AccessToken,
) UserAccessTokenDomain {
    return NewUserAccessTokenDomain(
        p.session,
        p.namespaceName,
        accessToken,
    )
}

func (p NamespaceDomain) SalesItemMaster(
    salesItemName string,
) SalesItemMasterDomain {
    return NewSalesItemMasterDomain(
        p.session,
        p.salesItemMasterCache,
        p.namespaceName,
        salesItemName,
    )
}

func (p NamespaceDomain) SalesItemGroupMaster(
    salesItemGroupName string,
) SalesItemGroupMasterDomain {
    return NewSalesItemGroupMasterDomain(
        p.session,
        p.salesItemGroupMasterCache,
        p.namespaceName,
        salesItemGroupName,
    )
}

func (p NamespaceDomain) ShowcaseMaster(
    showcaseName string,
) ShowcaseMasterDomain {
    return NewShowcaseMasterDomain(
        p.session,
        p.showcaseMasterCache,
        p.namespaceName,
        showcaseName,
    )
}
