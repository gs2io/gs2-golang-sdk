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

type SalesItemMasterDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    salesItemMasterCache cache.SalesItemMasterDomainCache
    namespaceName string
    salesItemName string
}

func NewSalesItemMasterDomain(
    session *core.Gs2RestSession,
    salesItemMasterCache cache.SalesItemMasterDomainCache,
    namespaceName string,
    salesItemName string,
) SalesItemMasterDomain {
    return SalesItemMasterDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        salesItemMasterCache: salesItemMasterCache,
        namespaceName: namespaceName,
        salesItemName: salesItemName,
    }
}

func (p SalesItemMasterDomain) Load(
    request showcase.GetSalesItemMasterRequest,
) (*showcase.GetSalesItemMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemName = &p.salesItemName
    r, err := p.client.GetSalesItemMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemMasterCache.Update(*r.Item)
    return r, nil
}

func (p SalesItemMasterDomain) Update(
    request showcase.UpdateSalesItemMasterRequest,
) (*showcase.UpdateSalesItemMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemName = &p.salesItemName
    r, err := p.client.UpdateSalesItemMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemMasterCache.Update(*r.Item)
    return r, nil
}

func (p SalesItemMasterDomain) Delete(
    request showcase.DeleteSalesItemMasterRequest,
) (*showcase.DeleteSalesItemMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemName = &p.salesItemName
    r, err := p.client.DeleteSalesItemMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemMasterCache.Delete(*r.Item)
    return r, nil
}
