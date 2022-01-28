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

type SalesItemGroupMasterDomain struct {
    session *core.Gs2RestSession
    client showcase.Gs2ShowcaseRestClient
    salesItemGroupMasterCache cache.SalesItemGroupMasterDomainCache
    namespaceName string
    salesItemGroupName string
}

func NewSalesItemGroupMasterDomain(
    session *core.Gs2RestSession,
    salesItemGroupMasterCache cache.SalesItemGroupMasterDomainCache,
    namespaceName string,
    salesItemGroupName string,
) SalesItemGroupMasterDomain {
    return SalesItemGroupMasterDomain {
        session: session,
        client: showcase.Gs2ShowcaseRestClient{
            Session: session,
        },
        salesItemGroupMasterCache: salesItemGroupMasterCache,
        namespaceName: namespaceName,
        salesItemGroupName: salesItemGroupName,
    }
}

func (p SalesItemGroupMasterDomain) Load(
    request showcase.GetSalesItemGroupMasterRequest,
) (*showcase.GetSalesItemGroupMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemGroupName = &p.salesItemGroupName
    r, err := p.client.GetSalesItemGroupMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemGroupMasterCache.Update(*r.Item)
    return r, nil
}

func (p SalesItemGroupMasterDomain) Update(
    request showcase.UpdateSalesItemGroupMasterRequest,
) (*showcase.UpdateSalesItemGroupMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemGroupName = &p.salesItemGroupName
    r, err := p.client.UpdateSalesItemGroupMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemGroupMasterCache.Update(*r.Item)
    return r, nil
}

func (p SalesItemGroupMasterDomain) Delete(
    request showcase.DeleteSalesItemGroupMasterRequest,
) (*showcase.DeleteSalesItemGroupMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.SalesItemGroupName = &p.salesItemGroupName
    r, err := p.client.DeleteSalesItemGroupMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.salesItemGroupMasterCache.Delete(*r.Item)
    return r, nil
}
