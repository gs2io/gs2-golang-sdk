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

type InventoryDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    inventoryCache cache.InventoryDomainCache
    namespaceName string
    userId string
    inventoryName string
    itemSetCache cache.ItemSetDomainCache
}

func NewInventoryDomain(
    session *core.Gs2RestSession,
    inventoryCache cache.InventoryDomainCache,
    namespaceName string,
    userId string,
    inventoryName string,
) InventoryDomain {
    return InventoryDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        inventoryCache: inventoryCache,
        namespaceName: namespaceName,
        userId: userId,
        inventoryName: inventoryName,
        itemSetCache: cache.NewItemSetDomainCache(),
    }
}

func (p InventoryDomain) Load(
    request inventory.GetInventoryByUserIdRequest,
) (*inventory.GetInventoryByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    r, err := p.client.GetInventoryByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryCache.Update(*r.Item)
    return r, nil
}

func (p InventoryDomain) AddCapacity(
    request inventory.AddCapacityByUserIdRequest,
) (*inventory.AddCapacityByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    r, err := p.client.AddCapacityByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryCache.Update(*r.Item)
    return r, nil
}

func (p InventoryDomain) SetCapacity(
    request inventory.SetCapacityByUserIdRequest,
) (*inventory.SetCapacityByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    r, err := p.client.SetCapacityByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryCache.Update(*r.Item)
    return r, nil
}

func (p InventoryDomain) Delete(
    request inventory.DeleteInventoryByUserIdRequest,
) (*inventory.DeleteInventoryByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    r, err := p.client.DeleteInventoryByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryCache.Delete(*r.Item)
    return r, nil
}

func (p InventoryDomain) ItemSets(
) iterator.DescribeItemSetsByUserIdIterator {
    return iterator.NewDescribeItemSetsByUserIdIterator(
        p.itemSetCache,
        p.client,
        p.namespaceName,
        p.inventoryName,
        p.userId,
    )
}

func (p InventoryDomain) ItemSet(
    itemName string,
    itemSetName string,
) ItemSetDomain {
    return NewItemSetDomain(
        p.session,
        p.itemSetCache,
        p.namespaceName,
        p.userId,
        p.inventoryName,
        itemName,
        itemSetName,
    )
}
