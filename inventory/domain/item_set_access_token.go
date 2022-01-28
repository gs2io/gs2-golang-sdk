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

type ItemSetAccessTokenDomain struct {
    session *core.Gs2RestSession
    client inventory.Gs2InventoryRestClient
    itemSetCache cache.ItemSetDomainCache
    namespaceName string
    accessToken auth.AccessToken
    inventoryName string
    itemName string
    itemSetName string
}

func NewItemSetAccessTokenDomain(
    session *core.Gs2RestSession,
    itemSetCache cache.ItemSetDomainCache,
    namespaceName string,
    accessToken auth.AccessToken,
    inventoryName string,
    itemName string,
    itemSetName string,
) ItemSetAccessTokenDomain {
    return ItemSetAccessTokenDomain {
        session: session,
        client: inventory.Gs2InventoryRestClient{
            Session: session,
        },
        itemSetCache: itemSetCache,
        namespaceName: namespaceName,
        accessToken: accessToken,
        inventoryName: inventoryName,
        itemName: itemName,
        itemSetName: itemSetName,
    }
}

func (p ItemSetAccessTokenDomain) Load(
    request inventory.GetItemSetRequest,
) (*inventory.GetItemSetResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetItemSet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetAccessTokenDomain) GetItemWithSignature(
    request inventory.GetItemWithSignatureRequest,
) (*inventory.GetItemWithSignatureResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.GetItemWithSignature(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetAccessTokenDomain) Consume(
    request inventory.ConsumeItemSetRequest,
) (*inventory.ConsumeItemSetResult, error) {
    request.NamespaceName = &p.namespaceName
    request.AccessToken = p.accessToken.Token
    request.InventoryName = &p.inventoryName
    request.ItemName = &p.itemName
    request.ItemSetName = &p.itemSetName
    r, err := p.client.ConsumeItemSet(
        &request,
    )
    if err != nil {
        return nil, err
    }
    return r, nil
}

func (p ItemSetAccessTokenDomain) ReferenceOf(
) ReferenceOfAccessTokenDomain {
    return NewReferenceOfAccessTokenDomain(
        p.session,
        p.namespaceName,
        p.accessToken,
        p.inventoryName,
        p.itemName,
        p.itemSetName,
    )
}
