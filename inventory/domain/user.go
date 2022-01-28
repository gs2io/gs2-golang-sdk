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

type UserDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    namespaceName string
    userId string
    inventoryCache cache.InventoryDomainCache
}

func NewUserDomain(
    session *core.Gs2RestSession,
    namespaceName string,
    userId string,
) UserDomain {
    return UserDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        namespaceName: namespaceName,
        userId: userId,
        inventoryCache: cache.NewInventoryDomainCache(),
    }
}

func (p UserDomain) Inventories(
) iterator.DescribeInventoriesByUserIdIterator {
    return iterator.NewDescribeInventoriesByUserIdIterator(
        p.inventoryCache,
        p.client,
        p.namespaceName,
        p.userId,
    )
}

func (p UserDomain) Inventory(
    inventoryName string,
) InventoryDomain {
    return NewInventoryDomain(
        p.session,
        p.inventoryCache,
        p.namespaceName,
        p.userId,
        inventoryName,
    )
}
