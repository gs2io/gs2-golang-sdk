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

type InventoryAccessTokenDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    inventoryCache cache.InventoryDomainCache
    namespaceName string
    accessToken auth.AccessToken
    inventoryName string
    itemSetCache cache.ItemSetDomainCache
}

func NewInventoryAccessTokenDomain(
    session *core.Gs2RestSession,
    inventoryCache cache.InventoryDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    inventoryName string,
) InventoryAccessTokenDomain {
    return InventoryAccessTokenDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        inventoryCache: inventoryCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        inventoryName: inventoryName,
        itemSetCache: cache.NewItemSetDomainCache(),
    }
}

func (p InventoryAccessTokenDomain) Load(
    request inventory.GetInventoryRequest,
) (*inventory.GetInventoryResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    r, err := p.client.GetInventory(
        &request,
    )
    if err != nil {
        return nil, err
    }
    p.inventoryCache.Update(*r.Item)
    return r, nil
}

func (p InventoryAccessTokenDomain) ItemSets(
) iterator.DescribeItemSetsIterator {
    return iterator.NewDescribeItemSetsIterator(
        p.itemSetCache,
        p.client,
        p.namespaceName,
        p.inventoryName,
        p.accessToken,
    )
}

func (p InventoryAccessTokenDomain) ItemSet(
    itemName string,
    itemSetName string,
) ItemSetAccessTokenDomain {
    return NewItemSetAccessTokenDomain(
        p.session,
        p.itemSetCache,
        p.namespaceName,
        p.accessToken,
        p.inventoryName,
        itemName,
        itemSetName,
    )
}
