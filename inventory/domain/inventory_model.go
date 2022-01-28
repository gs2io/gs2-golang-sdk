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

type InventoryModelDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    inventoryModelCache cache.InventoryModelDomainCache
    namespaceName string
    inventoryName string
    itemModelCache cache.ItemModelDomainCache
}

func NewInventoryModelDomain(
    session *core.Gs2RestSession,
    inventoryModelCache cache.InventoryModelDomainCache,
    namespaceName string,
    inventoryName string,
) InventoryModelDomain {
    return InventoryModelDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        inventoryModelCache: inventoryModelCache,
        namespaceName: namespaceName,
        inventoryName: inventoryName,
        itemModelCache: cache.NewItemModelDomainCache(),
    }
}

func (p InventoryModelDomain) Load(
    request inventory.GetInventoryModelRequest,
) (*inventory.GetInventoryModelResult, error) {
    request.NamespaceName = &p.namespaceName
    request.InventoryName = &p.inventoryName
    r, err := p.client.GetInventoryModel(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryModelCache.Update(*r.Item)
    return r, nil
}

func (p InventoryModelDomain) ItemModels(
) iterator.DescribeItemModelsIterator {
    return iterator.NewDescribeItemModelsIterator(
        p.itemModelCache,
        p.client,
        p.namespaceName,
        p.inventoryName,
    )
}

func (p InventoryModelDomain) ItemModel(
    itemName string,
) ItemModelDomain {
    return NewItemModelDomain(
        p.session,
        p.itemModelCache,
        p.namespaceName,
        p.inventoryName,
        itemName,
    )
}