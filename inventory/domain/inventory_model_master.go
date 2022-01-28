package domaininventory
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

    "github.com/gs2io/gs2-golang-sdk/inventory"
    "github.com/gs2io/gs2-golang-sdk/inventory/domain/cache"
    "github.com/gs2io/gs2-golang-sdk/inventory/domain/iterator"
)

var _ = auth.AccessToken{}
var _ = cache.NamespaceDomainCache{}
var _ = iterator.DescribeNamespacesIterator{}

type InventoryModelMasterDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    inventoryModelMasterCache cache.InventoryModelMasterDomainCache
    namespaceName string
    inventoryName string
    itemModelMasterCache cache.ItemModelMasterDomainCache
}

func NewInventoryModelMasterDomain(
    session *core.Gs2RestSession,
    inventoryModelMasterCache cache.InventoryModelMasterDomainCache,
    namespaceName string,
    inventoryName string,
) InventoryModelMasterDomain {
    return InventoryModelMasterDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        inventoryModelMasterCache: inventoryModelMasterCache,
        namespaceName: namespaceName,
        inventoryName: inventoryName,
        itemModelMasterCache: cache.NewItemModelMasterDomainCache(),
    }
}

func (p InventoryModelMasterDomain) Load(
    request inventory.GetInventoryModelMasterRequest,
) (*inventory.GetInventoryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.InventoryName = &p.inventoryName
    r, err := p.client.GetInventoryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p InventoryModelMasterDomain) Update(
    request inventory.UpdateInventoryModelMasterRequest,
) (*inventory.UpdateInventoryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.InventoryName = &p.inventoryName
    r, err := p.client.UpdateInventoryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryModelMasterCache.Update(*r.Item)
    return r, nil
}

func (p InventoryModelMasterDomain) Delete(
    request inventory.DeleteInventoryModelMasterRequest,
) (*inventory.DeleteInventoryModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.InventoryName = &p.inventoryName
    r, err := p.client.DeleteInventoryModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryModelMasterCache.Delete(*r.Item)
    return r, nil
}

func (p InventoryModelMasterDomain) CreateItemModelMaster(
    request inventory.CreateItemModelMasterRequest,
) (*inventory.CreateItemModelMasterResult, error) {
    request.NamespaceName = &p.namespaceName
    request.InventoryName = &p.inventoryName
    r, err := p.client.CreateItemModelMaster(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p InventoryModelMasterDomain) ItemModelMasters(
) iterator.DescribeItemModelMastersIterator {
    return iterator.NewDescribeItemModelMastersIterator(
        p.itemModelMasterCache,
        p.client,
        p.namespaceName,
        p.inventoryName,
    )
}

func (p InventoryModelMasterDomain) ItemModelMaster(
    itemName string,
) ItemModelMasterDomain {
    return NewItemModelMasterDomain(
        p.session,
        p.itemModelMasterCache,
        p.namespaceName,
        p.inventoryName,
        itemName,
    )
}
