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

type ItemSetDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    itemSetCache cache.ItemSetDomainCache
    namespaceName string
    userId string
    inventoryName string
    itemName string
    itemSetName string
}

func NewItemSetDomain(
    session *core.Gs2RestSession,
    itemSetCache cache.ItemSetDomainCache,
    namespaceName string,
    userId string,
    inventoryName string,
    itemName string,
    itemSetName string,
) ItemSetDomain {
    return ItemSetDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        itemSetCache: itemSetCache,
        namespaceName: namespaceName,
        userId: userId,
        inventoryName: inventoryName,
        itemName: itemName,
        itemSetName: itemSetName,
    }
}

func (p ItemSetDomain) Load(
    request inventory.GetItemSetByUserIdRequest,
) (*inventory.GetItemSetByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetItemSetByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetDomain) GetItemWithSignature(
    request inventory.GetItemWithSignatureByUserIdRequest,
) (*inventory.GetItemWithSignatureByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetItemWithSignatureByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetDomain) Acquire(
    request inventory.AcquireItemSetByUserIdRequest,
) (*inventory.AcquireItemSetByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.AcquireItemSetByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetDomain) Consume(
    request inventory.ConsumeItemSetByUserIdRequest,
) (*inventory.ConsumeItemSetByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.ConsumeItemSetByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetDomain) Delete(
    request inventory.DeleteItemSetByUserIdRequest,
) (*inventory.DeleteItemSetByUserIdResult, error) {
    request.NamespaceName = &p.namespaceName
    request.UserId = &p.userId
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.DeleteItemSetByUserId(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetDomain) ReferenceOf(
) ReferenceOfDomain {
    return NewReferenceOfDomain(
        p.session,
        p.namespaceName,
        p.userId,
        p.inventoryName,
        p.itemName,
        p.itemSetName,
    )
}
